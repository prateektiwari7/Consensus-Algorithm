import "github.com/ethereum/go-ethereum/crypto"
import "github.com/ethereum/go-ethereum/p2p"

nodekey, _ := crypto.GenerateKey()
srv := p2p.Server{
	MaxPeers:   10,
	PrivateKey: nodekey,
	Name:       "my node name",
	ListenAddr: ":30300",
	Protocols:  []p2p.Protocol{},
}
srv.Start()