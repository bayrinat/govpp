// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: map.api.json

/*
 Package maps is a generated from VPP binary API module 'map'.

 It contains following objects:
	 16 services
	  1 enum
	  2 aliases
	  5 types
	  1 union
	 32 messages
*/
package maps

import api "git.fd.io/govpp.git/api"
import struc "github.com/lunixbochs/struc"
import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// Services represents VPP binary API services:
type Services interface {
	DumpMapDomain(*MapDomainDump) ([]*MapDomainDetails, error)
	DumpMapRule(*MapRuleDump) ([]*MapRuleDetails, error)
	MapAddDelRule(*MapAddDelRule) (*MapAddDelRuleReply, error)
	MapAddDomain(*MapAddDomain) (*MapAddDomainReply, error)
	MapDelDomain(*MapDelDomain) (*MapDelDomainReply, error)
	MapIfEnableDisable(*MapIfEnableDisable) (*MapIfEnableDisableReply, error)
	MapParamAddDelPreResolve(*MapParamAddDelPreResolve) (*MapParamAddDelPreResolveReply, error)
	MapParamGet(*MapParamGet) (*MapParamGetReply, error)
	MapParamSetFragmentation(*MapParamSetFragmentation) (*MapParamSetFragmentationReply, error)
	MapParamSetICMP(*MapParamSetICMP) (*MapParamSetICMPReply, error)
	MapParamSetICMP6(*MapParamSetICMP6) (*MapParamSetICMP6Reply, error)
	MapParamSetReassembly(*MapParamSetReassembly) (*MapParamSetReassemblyReply, error)
	MapParamSetSecurityCheck(*MapParamSetSecurityCheck) (*MapParamSetSecurityCheckReply, error)
	MapParamSetTCP(*MapParamSetTCP) (*MapParamSetTCPReply, error)
	MapParamSetTrafficClass(*MapParamSetTrafficClass) (*MapParamSetTrafficClassReply, error)
	MapSummaryStats(*MapSummaryStats) (*MapSummaryStatsReply, error)
}

/* Enums */

// AddressFamily represents VPP binary API enum 'address_family':
type AddressFamily uint32

const (
	ADDRESS_IP4 AddressFamily = 0
	ADDRESS_IP6 AddressFamily = 1
)

/* Aliases */

// IP4Address represents VPP binary API alias 'ip4_address':
type IP4Address [4]uint8

// IP6Address represents VPP binary API alias 'ip6_address':
type IP6Address [16]uint8

/* Types */

// Address represents VPP binary API type 'address':
type Address struct {
	Af AddressFamily
	Un AddressUnion
}

func (*Address) GetTypeName() string {
	return "address"
}
func (*Address) GetCrcString() string {
	return "09f11671"
}

// IP4Prefix represents VPP binary API type 'ip4_prefix':
type IP4Prefix struct {
	Prefix IP4Address
	Len    uint8
}

func (*IP4Prefix) GetTypeName() string {
	return "ip4_prefix"
}
func (*IP4Prefix) GetCrcString() string {
	return "ea8dc11d"
}

// IP6Prefix represents VPP binary API type 'ip6_prefix':
type IP6Prefix struct {
	Prefix IP6Address
	Len    uint8
}

func (*IP6Prefix) GetTypeName() string {
	return "ip6_prefix"
}
func (*IP6Prefix) GetCrcString() string {
	return "779fd64f"
}

// Mprefix represents VPP binary API type 'mprefix':
type Mprefix struct {
	Af               AddressFamily
	GrpAddressLength uint16
	GrpAddress       AddressUnion
	SrcAddress       AddressUnion
}

func (*Mprefix) GetTypeName() string {
	return "mprefix"
}
func (*Mprefix) GetCrcString() string {
	return "1c4cba05"
}

// Prefix represents VPP binary API type 'prefix':
type Prefix struct {
	Address       Address
	AddressLength uint8
}

func (*Prefix) GetTypeName() string {
	return "prefix"
}
func (*Prefix) GetCrcString() string {
	return "0403aebc"
}

/* Unions */

// AddressUnion represents VPP binary API union 'address_union':
type AddressUnion struct {
	Union_data [16]byte
}

func (*AddressUnion) GetTypeName() string {
	return "address_union"
}
func (*AddressUnion) GetCrcString() string {
	return "d68a2fb4"
}

func AddressUnionIP4(a IP4Address) (u AddressUnion) {
	u.SetIP4(a)
	return
}
func (u *AddressUnion) SetIP4(a IP4Address) {
	var b = new(bytes.Buffer)
	if err := struc.Pack(b, &a); err != nil {
		return
	}
	copy(u.Union_data[:], b.Bytes())
}
func (u *AddressUnion) GetIP4() (a IP4Address) {
	var b = bytes.NewReader(u.Union_data[:])
	struc.Unpack(b, &a)
	return
}

func AddressUnionIP6(a IP6Address) (u AddressUnion) {
	u.SetIP6(a)
	return
}
func (u *AddressUnion) SetIP6(a IP6Address) {
	var b = new(bytes.Buffer)
	if err := struc.Pack(b, &a); err != nil {
		return
	}
	copy(u.Union_data[:], b.Bytes())
}
func (u *AddressUnion) GetIP6() (a IP6Address) {
	var b = bytes.NewReader(u.Union_data[:])
	struc.Unpack(b, &a)
	return
}

/* Messages */

// MapAddDelRule represents VPP binary API message 'map_add_del_rule':
type MapAddDelRule struct {
	Index  uint32
	IsAdd  bool
	IP6Dst IP6Address
	Psid   uint16
}

func (*MapAddDelRule) GetMessageName() string {
	return "map_add_del_rule"
}
func (*MapAddDelRule) GetCrcString() string {
	return "e6132040"
}
func (*MapAddDelRule) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapAddDelRuleReply represents VPP binary API message 'map_add_del_rule_reply':
type MapAddDelRuleReply struct {
	Retval int32
}

func (*MapAddDelRuleReply) GetMessageName() string {
	return "map_add_del_rule_reply"
}
func (*MapAddDelRuleReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MapAddDelRuleReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapAddDomain represents VPP binary API message 'map_add_domain':
type MapAddDomain struct {
	IP6Prefix  IP6Prefix
	IP4Prefix  IP4Prefix
	IP6Src     IP6Prefix
	EaBitsLen  uint8
	PsidOffset uint8
	PsidLength uint8
	Mtu        uint16
}

func (*MapAddDomain) GetMessageName() string {
	return "map_add_domain"
}
func (*MapAddDomain) GetCrcString() string {
	return "a9358068"
}
func (*MapAddDomain) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapAddDomainReply represents VPP binary API message 'map_add_domain_reply':
type MapAddDomainReply struct {
	Index  uint32
	Retval int32
}

func (*MapAddDomainReply) GetMessageName() string {
	return "map_add_domain_reply"
}
func (*MapAddDomainReply) GetCrcString() string {
	return "3e6d4e2c"
}
func (*MapAddDomainReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapDelDomain represents VPP binary API message 'map_del_domain':
type MapDelDomain struct {
	Index uint32
}

func (*MapDelDomain) GetMessageName() string {
	return "map_del_domain"
}
func (*MapDelDomain) GetCrcString() string {
	return "8ac76db6"
}
func (*MapDelDomain) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapDelDomainReply represents VPP binary API message 'map_del_domain_reply':
type MapDelDomainReply struct {
	Retval int32
}

func (*MapDelDomainReply) GetMessageName() string {
	return "map_del_domain_reply"
}
func (*MapDelDomainReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MapDelDomainReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapDomainDetails represents VPP binary API message 'map_domain_details':
type MapDomainDetails struct {
	DomainIndex uint32
	IP6Prefix   IP6Prefix
	IP4Prefix   IP4Prefix
	IP6Src      IP6Prefix
	EaBitsLen   uint8
	PsidOffset  uint8
	PsidLength  uint8
	Flags       uint8
	Mtu         uint16
}

func (*MapDomainDetails) GetMessageName() string {
	return "map_domain_details"
}
func (*MapDomainDetails) GetCrcString() string {
	return "2a17dcb8"
}
func (*MapDomainDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapDomainDump represents VPP binary API message 'map_domain_dump':
type MapDomainDump struct{}

func (*MapDomainDump) GetMessageName() string {
	return "map_domain_dump"
}
func (*MapDomainDump) GetCrcString() string {
	return "51077d14"
}
func (*MapDomainDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapIfEnableDisable represents VPP binary API message 'map_if_enable_disable':
type MapIfEnableDisable struct {
	SwIfIndex     uint32
	IsEnable      bool
	IsTranslation bool
}

func (*MapIfEnableDisable) GetMessageName() string {
	return "map_if_enable_disable"
}
func (*MapIfEnableDisable) GetCrcString() string {
	return "61a30cd9"
}
func (*MapIfEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapIfEnableDisableReply represents VPP binary API message 'map_if_enable_disable_reply':
type MapIfEnableDisableReply struct {
	Retval int32
}

func (*MapIfEnableDisableReply) GetMessageName() string {
	return "map_if_enable_disable_reply"
}
func (*MapIfEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MapIfEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapParamAddDelPreResolve represents VPP binary API message 'map_param_add_del_pre_resolve':
type MapParamAddDelPreResolve struct {
	IsAdd        bool
	IP4NhAddress IP4Address
	IP6NhAddress IP6Address
}

func (*MapParamAddDelPreResolve) GetMessageName() string {
	return "map_param_add_del_pre_resolve"
}
func (*MapParamAddDelPreResolve) GetCrcString() string {
	return "ea9a9a4a"
}
func (*MapParamAddDelPreResolve) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapParamAddDelPreResolveReply represents VPP binary API message 'map_param_add_del_pre_resolve_reply':
type MapParamAddDelPreResolveReply struct {
	Retval int32
}

func (*MapParamAddDelPreResolveReply) GetMessageName() string {
	return "map_param_add_del_pre_resolve_reply"
}
func (*MapParamAddDelPreResolveReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MapParamAddDelPreResolveReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapParamGet represents VPP binary API message 'map_param_get':
type MapParamGet struct{}

func (*MapParamGet) GetMessageName() string {
	return "map_param_get"
}
func (*MapParamGet) GetCrcString() string {
	return "51077d14"
}
func (*MapParamGet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapParamGetReply represents VPP binary API message 'map_param_get_reply':
type MapParamGetReply struct {
	Retval                 int32
	FragInner              uint8
	FragIgnoreDf           uint8
	ICMPIP4ErrRelaySrc     IP4Address
	ICMP6EnableUnreachable bool
	IP4NhAddress           IP4Address
	IP6NhAddress           IP6Address
	IP4LifetimeMs          uint16
	IP4PoolSize            uint16
	IP4Buffers             uint32
	IP4HtRatio             float64
	IP6LifetimeMs          uint16
	IP6PoolSize            uint16
	IP6Buffers             uint32
	IP6HtRatio             float64
	SecCheckEnable         bool
	SecCheckFragments      bool
	TcCopy                 bool
	TcClass                uint8
}

func (*MapParamGetReply) GetMessageName() string {
	return "map_param_get_reply"
}
func (*MapParamGetReply) GetCrcString() string {
	return "b40e9226"
}
func (*MapParamGetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapParamSetFragmentation represents VPP binary API message 'map_param_set_fragmentation':
type MapParamSetFragmentation struct {
	Inner    bool
	IgnoreDf bool
}

func (*MapParamSetFragmentation) GetMessageName() string {
	return "map_param_set_fragmentation"
}
func (*MapParamSetFragmentation) GetCrcString() string {
	return "9ff54d90"
}
func (*MapParamSetFragmentation) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapParamSetFragmentationReply represents VPP binary API message 'map_param_set_fragmentation_reply':
type MapParamSetFragmentationReply struct {
	Retval int32
}

func (*MapParamSetFragmentationReply) GetMessageName() string {
	return "map_param_set_fragmentation_reply"
}
func (*MapParamSetFragmentationReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MapParamSetFragmentationReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapParamSetICMP represents VPP binary API message 'map_param_set_icmp':
type MapParamSetICMP struct {
	IP4ErrRelaySrc IP4Address
}

func (*MapParamSetICMP) GetMessageName() string {
	return "map_param_set_icmp"
}
func (*MapParamSetICMP) GetCrcString() string {
	return "4c0a4fd2"
}
func (*MapParamSetICMP) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapParamSetICMP6 represents VPP binary API message 'map_param_set_icmp6':
type MapParamSetICMP6 struct {
	EnableUnreachable bool
}

func (*MapParamSetICMP6) GetMessageName() string {
	return "map_param_set_icmp6"
}
func (*MapParamSetICMP6) GetCrcString() string {
	return "5d01f8c1"
}
func (*MapParamSetICMP6) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapParamSetICMP6Reply represents VPP binary API message 'map_param_set_icmp6_reply':
type MapParamSetICMP6Reply struct {
	Retval int32
}

func (*MapParamSetICMP6Reply) GetMessageName() string {
	return "map_param_set_icmp6_reply"
}
func (*MapParamSetICMP6Reply) GetCrcString() string {
	return "e8d4e804"
}
func (*MapParamSetICMP6Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapParamSetICMPReply represents VPP binary API message 'map_param_set_icmp_reply':
type MapParamSetICMPReply struct {
	Retval int32
}

func (*MapParamSetICMPReply) GetMessageName() string {
	return "map_param_set_icmp_reply"
}
func (*MapParamSetICMPReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MapParamSetICMPReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapParamSetReassembly represents VPP binary API message 'map_param_set_reassembly':
type MapParamSetReassembly struct {
	IsIP6      bool
	LifetimeMs uint16
	PoolSize   uint16
	Buffers    uint32
	HtRatio    float64
}

func (*MapParamSetReassembly) GetMessageName() string {
	return "map_param_set_reassembly"
}
func (*MapParamSetReassembly) GetCrcString() string {
	return "54172b10"
}
func (*MapParamSetReassembly) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapParamSetReassemblyReply represents VPP binary API message 'map_param_set_reassembly_reply':
type MapParamSetReassemblyReply struct {
	Retval int32
}

func (*MapParamSetReassemblyReply) GetMessageName() string {
	return "map_param_set_reassembly_reply"
}
func (*MapParamSetReassemblyReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MapParamSetReassemblyReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapParamSetSecurityCheck represents VPP binary API message 'map_param_set_security_check':
type MapParamSetSecurityCheck struct {
	Enable    bool
	Fragments bool
}

func (*MapParamSetSecurityCheck) GetMessageName() string {
	return "map_param_set_security_check"
}
func (*MapParamSetSecurityCheck) GetCrcString() string {
	return "6abe9836"
}
func (*MapParamSetSecurityCheck) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapParamSetSecurityCheckReply represents VPP binary API message 'map_param_set_security_check_reply':
type MapParamSetSecurityCheckReply struct {
	Retval int32
}

func (*MapParamSetSecurityCheckReply) GetMessageName() string {
	return "map_param_set_security_check_reply"
}
func (*MapParamSetSecurityCheckReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MapParamSetSecurityCheckReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapParamSetTCP represents VPP binary API message 'map_param_set_tcp':
type MapParamSetTCP struct {
	TCPMss uint16
}

func (*MapParamSetTCP) GetMessageName() string {
	return "map_param_set_tcp"
}
func (*MapParamSetTCP) GetCrcString() string {
	return "87a825d9"
}
func (*MapParamSetTCP) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapParamSetTCPReply represents VPP binary API message 'map_param_set_tcp_reply':
type MapParamSetTCPReply struct {
	Retval int32
}

func (*MapParamSetTCPReply) GetMessageName() string {
	return "map_param_set_tcp_reply"
}
func (*MapParamSetTCPReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MapParamSetTCPReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapParamSetTrafficClass represents VPP binary API message 'map_param_set_traffic_class':
type MapParamSetTrafficClass struct {
	Copy  bool
	Class uint8
}

func (*MapParamSetTrafficClass) GetMessageName() string {
	return "map_param_set_traffic_class"
}
func (*MapParamSetTrafficClass) GetCrcString() string {
	return "007ee563"
}
func (*MapParamSetTrafficClass) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapParamSetTrafficClassReply represents VPP binary API message 'map_param_set_traffic_class_reply':
type MapParamSetTrafficClassReply struct {
	Retval int32
}

func (*MapParamSetTrafficClassReply) GetMessageName() string {
	return "map_param_set_traffic_class_reply"
}
func (*MapParamSetTrafficClassReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MapParamSetTrafficClassReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapRuleDetails represents VPP binary API message 'map_rule_details':
type MapRuleDetails struct {
	IP6Dst IP6Address
	Psid   uint16
}

func (*MapRuleDetails) GetMessageName() string {
	return "map_rule_details"
}
func (*MapRuleDetails) GetCrcString() string {
	return "4f932665"
}
func (*MapRuleDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapRuleDump represents VPP binary API message 'map_rule_dump':
type MapRuleDump struct {
	DomainIndex uint32
}

func (*MapRuleDump) GetMessageName() string {
	return "map_rule_dump"
}
func (*MapRuleDump) GetCrcString() string {
	return "e43e6ff6"
}
func (*MapRuleDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapSummaryStats represents VPP binary API message 'map_summary_stats':
type MapSummaryStats struct{}

func (*MapSummaryStats) GetMessageName() string {
	return "map_summary_stats"
}
func (*MapSummaryStats) GetCrcString() string {
	return "51077d14"
}
func (*MapSummaryStats) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapSummaryStatsReply represents VPP binary API message 'map_summary_stats_reply':
type MapSummaryStatsReply struct {
	Retval             int32
	TotalBindings      uint64
	TotalPkts          []uint64 `struc:"[2]uint64"`
	TotalBytes         []uint64 `struc:"[2]uint64"`
	TotalIP4Fragments  uint64
	TotalSecurityCheck []uint64 `struc:"[2]uint64"`
}

func (*MapSummaryStatsReply) GetMessageName() string {
	return "map_summary_stats_reply"
}
func (*MapSummaryStatsReply) GetCrcString() string {
	return "0e4ace0e"
}
func (*MapSummaryStatsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*MapAddDelRule)(nil), "map.MapAddDelRule")
	api.RegisterMessage((*MapAddDelRuleReply)(nil), "map.MapAddDelRuleReply")
	api.RegisterMessage((*MapAddDomain)(nil), "map.MapAddDomain")
	api.RegisterMessage((*MapAddDomainReply)(nil), "map.MapAddDomainReply")
	api.RegisterMessage((*MapDelDomain)(nil), "map.MapDelDomain")
	api.RegisterMessage((*MapDelDomainReply)(nil), "map.MapDelDomainReply")
	api.RegisterMessage((*MapDomainDetails)(nil), "map.MapDomainDetails")
	api.RegisterMessage((*MapDomainDump)(nil), "map.MapDomainDump")
	api.RegisterMessage((*MapIfEnableDisable)(nil), "map.MapIfEnableDisable")
	api.RegisterMessage((*MapIfEnableDisableReply)(nil), "map.MapIfEnableDisableReply")
	api.RegisterMessage((*MapParamAddDelPreResolve)(nil), "map.MapParamAddDelPreResolve")
	api.RegisterMessage((*MapParamAddDelPreResolveReply)(nil), "map.MapParamAddDelPreResolveReply")
	api.RegisterMessage((*MapParamGet)(nil), "map.MapParamGet")
	api.RegisterMessage((*MapParamGetReply)(nil), "map.MapParamGetReply")
	api.RegisterMessage((*MapParamSetFragmentation)(nil), "map.MapParamSetFragmentation")
	api.RegisterMessage((*MapParamSetFragmentationReply)(nil), "map.MapParamSetFragmentationReply")
	api.RegisterMessage((*MapParamSetICMP)(nil), "map.MapParamSetICMP")
	api.RegisterMessage((*MapParamSetICMP6)(nil), "map.MapParamSetICMP6")
	api.RegisterMessage((*MapParamSetICMP6Reply)(nil), "map.MapParamSetICMP6Reply")
	api.RegisterMessage((*MapParamSetICMPReply)(nil), "map.MapParamSetICMPReply")
	api.RegisterMessage((*MapParamSetReassembly)(nil), "map.MapParamSetReassembly")
	api.RegisterMessage((*MapParamSetReassemblyReply)(nil), "map.MapParamSetReassemblyReply")
	api.RegisterMessage((*MapParamSetSecurityCheck)(nil), "map.MapParamSetSecurityCheck")
	api.RegisterMessage((*MapParamSetSecurityCheckReply)(nil), "map.MapParamSetSecurityCheckReply")
	api.RegisterMessage((*MapParamSetTCP)(nil), "map.MapParamSetTCP")
	api.RegisterMessage((*MapParamSetTCPReply)(nil), "map.MapParamSetTCPReply")
	api.RegisterMessage((*MapParamSetTrafficClass)(nil), "map.MapParamSetTrafficClass")
	api.RegisterMessage((*MapParamSetTrafficClassReply)(nil), "map.MapParamSetTrafficClassReply")
	api.RegisterMessage((*MapRuleDetails)(nil), "map.MapRuleDetails")
	api.RegisterMessage((*MapRuleDump)(nil), "map.MapRuleDump")
	api.RegisterMessage((*MapSummaryStats)(nil), "map.MapSummaryStats")
	api.RegisterMessage((*MapSummaryStatsReply)(nil), "map.MapSummaryStatsReply")
}

var Messages = []api.Message{
	(*MapAddDelRule)(nil),
	(*MapAddDelRuleReply)(nil),
	(*MapAddDomain)(nil),
	(*MapAddDomainReply)(nil),
	(*MapDelDomain)(nil),
	(*MapDelDomainReply)(nil),
	(*MapDomainDetails)(nil),
	(*MapDomainDump)(nil),
	(*MapIfEnableDisable)(nil),
	(*MapIfEnableDisableReply)(nil),
	(*MapParamAddDelPreResolve)(nil),
	(*MapParamAddDelPreResolveReply)(nil),
	(*MapParamGet)(nil),
	(*MapParamGetReply)(nil),
	(*MapParamSetFragmentation)(nil),
	(*MapParamSetFragmentationReply)(nil),
	(*MapParamSetICMP)(nil),
	(*MapParamSetICMP6)(nil),
	(*MapParamSetICMP6Reply)(nil),
	(*MapParamSetICMPReply)(nil),
	(*MapParamSetReassembly)(nil),
	(*MapParamSetReassemblyReply)(nil),
	(*MapParamSetSecurityCheck)(nil),
	(*MapParamSetSecurityCheckReply)(nil),
	(*MapParamSetTCP)(nil),
	(*MapParamSetTCPReply)(nil),
	(*MapParamSetTrafficClass)(nil),
	(*MapParamSetTrafficClassReply)(nil),
	(*MapRuleDetails)(nil),
	(*MapRuleDump)(nil),
	(*MapSummaryStats)(nil),
	(*MapSummaryStatsReply)(nil),
}
