// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tezos

var (
	ProtoAlpha     = MustParseProtocolHash("ProtoALphaALphaALphaALphaALphaALphaALphaALphaDdp3zK")
	ProtoGenesis   = MustParseProtocolHash("PrihK96nBAFSxVL1GLJTVhu9YnzkMFiBeuJRPA8NwuZVZCE1L6i")
	ProtoBootstrap = MustParseProtocolHash("Ps9mPmXaRzmzk35gbAYNCAw6UXdE2qoABTHbN2oEEc1qM7CwT9P")
	ProtoV001      = MustParseProtocolHash("PtCJ7pwoxe8JasnHY8YonnLYjcVHmhiARPJvqcC6VfHT5s8k8sY")
	ProtoV002      = MustParseProtocolHash("PsYLVpVvgbLhAhoqAkMFUo6gudkJ9weNXhUYCiLDzcUpFpkk8Wt")
	ProtoV003      = MustParseProtocolHash("PsddFKi32cMJ2qPjf43Qv5GDWLDPZb3T3bF6fLKiF5HtvHNU7aP")
	ProtoV004      = MustParseProtocolHash("Pt24m4xiPbLDhVgVfABUjirbmda3yohdN82Sp9FeuAXJ4eV9otd")
	ProtoV005_2    = MustParseProtocolHash("PsBabyM1eUXZseaJdmXFApDSBqj8YBfwELoxZHHW77EMcAbbwAS")
	ProtoV006_2    = MustParseProtocolHash("PsCARTHAGazKbHtnKfLzQg3kms52kSRpgnDY982a9oYsSXRLQEb")
	ProtoV007      = MustParseProtocolHash("PsDELPH1Kxsxt8f9eWbxQeRxkjfbxoqM52jvs5Y5fBxWWh4ifpo")
	ProtoV008_2    = MustParseProtocolHash("PtEdo2ZkT9oKpimTah6x2embF25oss54njMuPzkJTEi5RqfdZFA")
	ProtoV009      = MustParseProtocolHash("PsFLorenaUUuikDWvMDr6fGBRG8kt3e3D3fHoXK1j1BFRxeSH4i")
	ProtoV010      = MustParseProtocolHash("PtGRANADsDU8R9daYKAgWnQYAJ64omN1o3KMGVCykShA97vQbvV")
	ProtoV011_2    = MustParseProtocolHash("PtHangz2aRngywmSRGGvrcTyMbbdpWdpFKuS4uMWxg2RaH9i1qx")
	ProtoV012_2    = MustParseProtocolHash("Psithaca2MLRFYargivpo7YvUr7wUDqyxrdhC5CQq78mRvimz6A")
	ProtoV013_2    = MustParseProtocolHash("PtJakart2xVj7pYXJBXrqHgd82rdkLey5ZeeGwDgPp9rhQUbSqY")
	ProtoV014      = MustParseProtocolHash("PtKathmankSpLLDALzWw7CGD2j2MtyveTwboEYokqUCP4a1LxMg")
	ProtoV015      = MustParseProtocolHash("PtLimaPtLMwfNinJi9rCfDPWea8dFgTZ1MeJ9f1m2SRic6ayiwW")
	ProtoV016      = MustParseProtocolHash("PtMumbaiiFFEGbew1rRjzSPyzRbA51Tm3RVZL5suHPxSZYDhCEc")
	ProtoV016_2    = MustParseProtocolHash("PtMumbai2TmsJHNGRkD8v8YDbtao7BLUC3wjASn1inAKLFCjaH1")

	// aliases
	PtAthens = ProtoV004
	PsBabyM1 = ProtoV005_2
	PsCARTHA = ProtoV006_2
	PsDELPH1 = ProtoV007
	PtEdo2Zk = ProtoV008_2
	PsFLoren = ProtoV009
	PtGRANAD = ProtoV010
	PtHangz2 = ProtoV011_2
	Psithaca = ProtoV012_2
	PtJakart = ProtoV013_2
	PtKathma = ProtoV014
	PtLimaPt = ProtoV015
	PtMumbai = ProtoV016_2

	Mainnet      = MustParseChainIdHash("NetXdQprcVkpaWU")
	Ghostnet     = MustParseChainIdHash("NetXnHfVqm9iesp")
	Jakartanet   = MustParseChainIdHash("NetXLH1uAxK7CCh")
	Kathmandunet = MustParseChainIdHash("NetXi2ZagzEsXbZ")
	Limanet      = MustParseChainIdHash("NetXizpkH94bocH")
	Mumbainet    = MustParseChainIdHash("NetXgbcrNtXD2yA")

	Versions = map[ProtocolHash]int{
		ProtoGenesis:   0,
		ProtoBootstrap: 0,
		ProtoV001:      1,
		ProtoV002:      2,
		ProtoV003:      3,
		ProtoV004:      4,
		ProtoV005_2:    5,
		ProtoV006_2:    6,
		ProtoV007:      7,
		ProtoV008_2:    8,
		ProtoV009:      9,
		ProtoV010:      10,
		ProtoV011_2:    11,
		ProtoV012_2:    12,
		ProtoV013_2:    13,
		ProtoV014:      14,
		ProtoV015:      15,
		ProtoV016:      16,
		ProtoV016_2:    16,
		ProtoAlpha:     17,
	}
)
