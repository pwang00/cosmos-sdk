package slashing

import (
	"strings"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	slashingv1beta1 "cosmossdk.io/api/cosmos/slashing/v1beta1"
)

func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: slashingv1beta1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "SigningInfo",
					Use:       "signing-info [validator-conspub]",
					Short:     "Query a validator's signing information",
					Long: strings.TrimSpace(`Use a validators' consensus public key to find the signing-info for that validator:

$ <appd> query slashing signing-info '{"@type":"/cosmos.crypto.ed25519.PubKey","key":"OauFcTKbN5Lx3fJL689cikXBqe+hcp6Y+x0rYUdR9Jk="}'
`),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "validator-conspub"},
					},
				},
				{
					RpcMethod: "SigningInfos",
					Short:     "Query signing information of all validators",
					Long: strings.TrimSpace(`signing infos of validators:

$ <appd> query slashing signing-infos
`),
				},
				{
					RpcMethod: "Params",
					Short:     "Query the current slashing parameters",
					Long: strings.TrimSpace(`Query genesis parameters for the slashing module:

$ <appd> query slashing params
`),
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: slashingv1beta1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Unjail",
					Use:       "unjail",
					Short:     "unjail validator previously jailed for downtime",
					Long: `unjail a jailed validator:

$ <appd> tx slashing unjail --from mykey
`,
				},
			},
		},
	}
}
