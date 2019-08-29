package traverser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	t.Run("without any select properties", func(t *testing.T) {
		sp := SelectProperties{}
		assert.Equal(t, false, sp.HasRefs(), "indicates no refs are present")
	})

	t.Run("with only primitive select properties", func(t *testing.T) {
		sp := SelectProperties{
			SelectProperty{
				IsPrimitive: true,
				Name:        "Foo",
			},
			SelectProperty{
				IsPrimitive: true,
				Name:        "Bar",
			},
		}

		assert.Equal(t, false, sp.HasRefs(), "indicates no refs are present")

		resolve, err := sp.ShouldResolve([]string{"inCountry", "Country"})
		require.Nil(t, err)
		assert.Equal(t, false, resolve)

	})

	t.Run("with a ref prop", func(t *testing.T) {
		sp := SelectProperties{
			SelectProperty{
				IsPrimitive: true,
				Name:        "name",
			},
			SelectProperty{
				IsPrimitive: false,
				Name:        "inCity",
				Refs: []SelectClass{
					SelectClass{
						ClassName: "City",
						RefProperties: SelectProperties{
							SelectProperty{
								Name:        "name",
								IsPrimitive: true,
							},
							SelectProperty{
								Name:        "inCountry",
								IsPrimitive: false,
								Refs: []SelectClass{
									SelectClass{
										ClassName: "Country",
										RefProperties: SelectProperties{
											SelectProperty{
												Name:        "name",
												IsPrimitive: true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}

		t.Run("checking for refs", func(t *testing.T) {
			assert.Equal(t, true, sp.HasRefs(), "indicates refs are present")
		})

		t.Run("checking valid single level ref", func(t *testing.T) {
			resolve, err := sp.ShouldResolve([]string{"inCity", "City"})
			require.Nil(t, err)
			assert.Equal(t, true, resolve)
		})

		t.Run("checking invalid single level ref", func(t *testing.T) {
			resolve, err := sp.ShouldResolve([]string{"inCity", "Town"})
			require.Nil(t, err)
			assert.Equal(t, false, resolve)
		})

		t.Run("checking valid nested ref", func(t *testing.T) {
			resolve, err := sp.ShouldResolve([]string{"inCity", "City", "inCountry", "Country"})
			require.Nil(t, err)
			assert.Equal(t, true, resolve)
		})

		t.Run("checking invalid nested level refs", func(t *testing.T) {
			resolve, err := sp.ShouldResolve([]string{"inCity", "Town", "inCountry", "Country"})
			require.Nil(t, err)
			assert.Equal(t, false, resolve)

			resolve, err = sp.ShouldResolve([]string{"inCity", "City", "inCountry", "Land"})
			require.Nil(t, err)
			assert.Equal(t, false, resolve)
		})

		t.Run("selecting a specific prop", func(t *testing.T) {
			prop := sp.FindProperty("inCity")
			assert.Equal(t, prop, &sp[1])
		})
	})

}