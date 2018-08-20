package bitcoind

type BlockchainInfo struct {
	Blocks               float64 `json:"blocks"`
	Initialblockdownload bool    `json:"initialblockdownload"`
	Mediantime           int64   `json:"mediantime"`
}
