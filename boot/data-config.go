package boot

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("1f8b08000000000002ff7490c14ac340108627a9adcd5e542878748857d9e0552c78cf2507bd34b62069b289243b659382955ec4f7f0243e8e771fc0f79089d82c1a0712989df97e862f0a077b131803c0f1e56304561d800709e9ac50c1da94b2a1aabcb91e82f332db2e8410a7a832cca94ab12e9a1493a596224e967a2e1011d7a6c429fa79d3acea8b205019cfef29d74a26da67ba43395dc4dcf7b39499bb2a9564540b2ac295a1870d6664f806953652c4edd38eaf718ab1f03c2ba45d900505fe59df20d1812fbc7914ee8fdf369f871f00c05fa7e7fc979ed14e4f6be57db65d306caf44a1e34e069d5d3b98edfed4eb13ffff736da7f41df15d470057270efc3d6938e2b10b2e3c03c0adc3dd57000000ffff66305bbdf5010000"); err != nil {
		panic(err)
	}
}
