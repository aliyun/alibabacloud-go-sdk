// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryNetListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *QueryNetListResponseBody
	GetCode() *int64
	SetData(v *QueryNetListResponseBodyData) *QueryNetListResponseBody
	GetData() *QueryNetListResponseBodyData
	SetMessage(v string) *QueryNetListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryNetListResponseBody
	GetRequestId() *string
}

type QueryNetListResponseBody struct {
	// The status code.
	//
	// - 200: The request was successful.
	//
	// - Other codes: The request failed.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The CIDR blocks.
	Data *QueryNetListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response parameters.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 24B652B5-AEFF-3F03-9114-00D053C42277
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryNetListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryNetListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryNetListResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *QueryNetListResponseBody) GetData() *QueryNetListResponseBodyData {
	return s.Data
}

func (s *QueryNetListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryNetListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryNetListResponseBody) SetCode(v int64) *QueryNetListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryNetListResponseBody) SetData(v *QueryNetListResponseBodyData) *QueryNetListResponseBody {
	s.Data = v
	return s
}

func (s *QueryNetListResponseBody) SetMessage(v string) *QueryNetListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryNetListResponseBody) SetRequestId(v string) *QueryNetListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryNetListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryNetListResponseBodyData struct {
	// The CIDR blocks.
	Nets []*QueryNetListResponseBodyDataNets `json:"Nets,omitempty" xml:"Nets,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	Num *int64 `json:"Num,omitempty" xml:"Num,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryNetListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryNetListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryNetListResponseBodyData) GetNets() []*QueryNetListResponseBodyDataNets {
	return s.Nets
}

func (s *QueryNetListResponseBodyData) GetNum() *int64 {
	return s.Num
}

func (s *QueryNetListResponseBodyData) GetPage() *int64 {
	return s.Page
}

func (s *QueryNetListResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *QueryNetListResponseBodyData) SetNets(v []*QueryNetListResponseBodyDataNets) *QueryNetListResponseBodyData {
	s.Nets = v
	return s
}

func (s *QueryNetListResponseBodyData) SetNum(v int64) *QueryNetListResponseBodyData {
	s.Num = &v
	return s
}

func (s *QueryNetListResponseBodyData) SetPage(v int64) *QueryNetListResponseBodyData {
	s.Page = &v
	return s
}

func (s *QueryNetListResponseBodyData) SetTotal(v int64) *QueryNetListResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryNetListResponseBodyData) Validate() error {
	if s.Nets != nil {
		for _, item := range s.Nets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryNetListResponseBodyDataNets struct {
	// The DDoS mitigation configuration of the CIDR block.
	DDoSDefense *QueryNetListResponseBodyDataNetsDDoSDefense `json:"DDoSDefense,omitempty" xml:"DDoSDefense,omitempty" type:"Struct"`
	// The advertising details.
	Declared []*QueryNetListResponseBodyDataNetsDeclared `json:"Declared,omitempty" xml:"Declared,omitempty" type:"Repeated"`
	// The advertising status of the CIDR block. Valid values:
	//
	// - 0: The CIDR block is not advertised.
	//
	// - 1: The CIDR block is advertised.
	//
	// example:
	//
	// 0
	DeclaredState *int32 `json:"DeclaredState,omitempty" xml:"DeclaredState,omitempty"`
	// Indicates whether the forwarding configuration takes effect. Valid values:
	//
	// - 0: The forwarding configuration does not take effect.
	//
	// - 1: The forwarding configuration takes effect.
	//
	// example:
	//
	// 1
	FwdEffect *int64 `json:"FwdEffect,omitempty" xml:"FwdEffect,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-09-05 13:52:50
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2024-09-05 13:59:55
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// The scheduling mode. Valid values:
	//
	// - netflow-auto
	//
	// - manual
	//
	// example:
	//
	// manual
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The CIDR block of the anti-DDoS diversion instance.
	//
	// example:
	//
	// 192.168.XX.XX/24
	Net *string `json:"Net,omitempty" xml:"Net,omitempty"`
	// Indicates whether the CIDR block needs to be extended. Valid values:
	//
	// - 0: The CIDR block does not need to be extended.
	//
	// - 1: The CIDR block needs to be extended.
	//
	// example:
	//
	// 0
	NetExtend *string `json:"NetExtend,omitempty" xml:"NetExtend,omitempty"`
	// The primary CIDR block of the anti-DDoS diversion instance for which an extended CIDR block is configured. If no extended CIDR blocks are configured for the anti-DDoS diversion instance, this parameter is left empty.
	//
	// example:
	//
	// 192.168.XX.XX/22
	NetMain *string `json:"NetMain,omitempty" xml:"NetMain,omitempty"`
	// The type of the CIDR block.
	//
	// example:
	//
	// ipv4
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The review status. Valid values:
	//
	// - 0: not reviewed.
	//
	// - 1: reviewed.
	//
	// - 2: under review.
	//
	// example:
	//
	// 2
	Permit *int32 `json:"Permit,omitempty" xml:"Permit,omitempty"`
	// The ID of the anti-DDoS diversion instance.
	//
	// example:
	//
	// ddos_diversion_public_cn-xxxxxxxxxxxxx
	SaleId *string `json:"SaleId,omitempty" xml:"SaleId,omitempty"`
	// The reinjection type. Valid values:
	//
	// - aliyun_line
	//
	// - pccw_line
	//
	// - gre_line
	//
	// - private_line
	//
	// example:
	//
	// aliyun_line
	UpstreamType *string `json:"UpstreamType,omitempty" xml:"UpstreamType,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 57690941xxxxxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryNetListResponseBodyDataNets) String() string {
	return dara.Prettify(s)
}

func (s QueryNetListResponseBodyDataNets) GoString() string {
	return s.String()
}

func (s *QueryNetListResponseBodyDataNets) GetDDoSDefense() *QueryNetListResponseBodyDataNetsDDoSDefense {
	return s.DDoSDefense
}

func (s *QueryNetListResponseBodyDataNets) GetDeclared() []*QueryNetListResponseBodyDataNetsDeclared {
	return s.Declared
}

func (s *QueryNetListResponseBodyDataNets) GetDeclaredState() *int32 {
	return s.DeclaredState
}

func (s *QueryNetListResponseBodyDataNets) GetFwdEffect() *int64 {
	return s.FwdEffect
}

func (s *QueryNetListResponseBodyDataNets) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryNetListResponseBodyDataNets) GetGmtModify() *string {
	return s.GmtModify
}

func (s *QueryNetListResponseBodyDataNets) GetMode() *string {
	return s.Mode
}

func (s *QueryNetListResponseBodyDataNets) GetNet() *string {
	return s.Net
}

func (s *QueryNetListResponseBodyDataNets) GetNetExtend() *string {
	return s.NetExtend
}

func (s *QueryNetListResponseBodyDataNets) GetNetMain() *string {
	return s.NetMain
}

func (s *QueryNetListResponseBodyDataNets) GetNetType() *string {
	return s.NetType
}

func (s *QueryNetListResponseBodyDataNets) GetPermit() *int32 {
	return s.Permit
}

func (s *QueryNetListResponseBodyDataNets) GetSaleId() *string {
	return s.SaleId
}

func (s *QueryNetListResponseBodyDataNets) GetUpstreamType() *string {
	return s.UpstreamType
}

func (s *QueryNetListResponseBodyDataNets) GetUserId() *string {
	return s.UserId
}

func (s *QueryNetListResponseBodyDataNets) SetDDoSDefense(v *QueryNetListResponseBodyDataNetsDDoSDefense) *QueryNetListResponseBodyDataNets {
	s.DDoSDefense = v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetDeclared(v []*QueryNetListResponseBodyDataNetsDeclared) *QueryNetListResponseBodyDataNets {
	s.Declared = v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetDeclaredState(v int32) *QueryNetListResponseBodyDataNets {
	s.DeclaredState = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetFwdEffect(v int64) *QueryNetListResponseBodyDataNets {
	s.FwdEffect = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetGmtCreate(v string) *QueryNetListResponseBodyDataNets {
	s.GmtCreate = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetGmtModify(v string) *QueryNetListResponseBodyDataNets {
	s.GmtModify = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetMode(v string) *QueryNetListResponseBodyDataNets {
	s.Mode = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetNet(v string) *QueryNetListResponseBodyDataNets {
	s.Net = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetNetExtend(v string) *QueryNetListResponseBodyDataNets {
	s.NetExtend = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetNetMain(v string) *QueryNetListResponseBodyDataNets {
	s.NetMain = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetNetType(v string) *QueryNetListResponseBodyDataNets {
	s.NetType = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetPermit(v int32) *QueryNetListResponseBodyDataNets {
	s.Permit = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetSaleId(v string) *QueryNetListResponseBodyDataNets {
	s.SaleId = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetUpstreamType(v string) *QueryNetListResponseBodyDataNets {
	s.UpstreamType = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetUserId(v string) *QueryNetListResponseBodyDataNets {
	s.UserId = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) Validate() error {
	if s.DDoSDefense != nil {
		if err := s.DDoSDefense.Validate(); err != nil {
			return err
		}
	}
	if s.Declared != nil {
		for _, item := range s.Declared {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryNetListResponseBodyDataNetsDDoSDefense struct {
	// The configuration of traffic scrubbing.
	CleanTh         *QueryNetListResponseBodyDataNetsDDoSDefenseCleanTh         `json:"CleanTh,omitempty" xml:"CleanTh,omitempty" type:"Struct"`
	CustomNetPolicy *QueryNetListResponseBodyDataNetsDDoSDefenseCustomNetPolicy `json:"CustomNetPolicy,omitempty" xml:"CustomNetPolicy,omitempty" type:"Struct"`
	// The configuration of the mitigation policy.
	DjPolicy *QueryNetListResponseBodyDataNetsDDoSDefenseDjPolicy `json:"DjPolicy,omitempty" xml:"DjPolicy,omitempty" type:"Struct"`
	// The configuration of blackhole filtering.
	HoleTh *QueryNetListResponseBodyDataNetsDDoSDefenseHoleTh `json:"HoleTh,omitempty" xml:"HoleTh,omitempty" type:"Struct"`
}

func (s QueryNetListResponseBodyDataNetsDDoSDefense) String() string {
	return dara.Prettify(s)
}

func (s QueryNetListResponseBodyDataNetsDDoSDefense) GoString() string {
	return s.String()
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefense) GetCleanTh() *QueryNetListResponseBodyDataNetsDDoSDefenseCleanTh {
	return s.CleanTh
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefense) GetCustomNetPolicy() *QueryNetListResponseBodyDataNetsDDoSDefenseCustomNetPolicy {
	return s.CustomNetPolicy
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefense) GetDjPolicy() *QueryNetListResponseBodyDataNetsDDoSDefenseDjPolicy {
	return s.DjPolicy
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefense) GetHoleTh() *QueryNetListResponseBodyDataNetsDDoSDefenseHoleTh {
	return s.HoleTh
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefense) SetCleanTh(v *QueryNetListResponseBodyDataNetsDDoSDefenseCleanTh) *QueryNetListResponseBodyDataNetsDDoSDefense {
	s.CleanTh = v
	return s
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefense) SetCustomNetPolicy(v *QueryNetListResponseBodyDataNetsDDoSDefenseCustomNetPolicy) *QueryNetListResponseBodyDataNetsDDoSDefense {
	s.CustomNetPolicy = v
	return s
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefense) SetDjPolicy(v *QueryNetListResponseBodyDataNetsDDoSDefenseDjPolicy) *QueryNetListResponseBodyDataNetsDDoSDefense {
	s.DjPolicy = v
	return s
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefense) SetHoleTh(v *QueryNetListResponseBodyDataNetsDDoSDefenseHoleTh) *QueryNetListResponseBodyDataNetsDDoSDefense {
	s.HoleTh = v
	return s
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefense) Validate() error {
	if s.CleanTh != nil {
		if err := s.CleanTh.Validate(); err != nil {
			return err
		}
	}
	if s.CustomNetPolicy != nil {
		if err := s.CustomNetPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.DjPolicy != nil {
		if err := s.DjPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.HoleTh != nil {
		if err := s.HoleTh.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryNetListResponseBodyDataNetsDDoSDefenseCleanTh struct {
	// The traffic scrubbing threshold in Mbit/s.
	//
	// example:
	//
	// 0
	Mbps *int32 `json:"Mbps,omitempty" xml:"Mbps,omitempty"`
	// The traffic scrubbing threshold in packets per second (pps).
	//
	// example:
	//
	// 0
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
}

func (s QueryNetListResponseBodyDataNetsDDoSDefenseCleanTh) String() string {
	return dara.Prettify(s)
}

func (s QueryNetListResponseBodyDataNetsDDoSDefenseCleanTh) GoString() string {
	return s.String()
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefenseCleanTh) GetMbps() *int32 {
	return s.Mbps
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefenseCleanTh) GetPps() *int32 {
	return s.Pps
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefenseCleanTh) SetMbps(v int32) *QueryNetListResponseBodyDataNetsDDoSDefenseCleanTh {
	s.Mbps = &v
	return s
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefenseCleanTh) SetPps(v int32) *QueryNetListResponseBodyDataNetsDDoSDefenseCleanTh {
	s.Pps = &v
	return s
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefenseCleanTh) Validate() error {
	return dara.Validate(s)
}

type QueryNetListResponseBodyDataNetsDDoSDefenseCustomNetPolicy struct {
	// example:
	//
	// test
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// example:
	//
	// BROAD_TMPL_1710000000_0001
	TmplId *string `json:"TmplId,omitempty" xml:"TmplId,omitempty"`
}

func (s QueryNetListResponseBodyDataNetsDDoSDefenseCustomNetPolicy) String() string {
	return dara.Prettify(s)
}

func (s QueryNetListResponseBodyDataNetsDDoSDefenseCustomNetPolicy) GoString() string {
	return s.String()
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefenseCustomNetPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefenseCustomNetPolicy) GetTmplId() *string {
	return s.TmplId
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefenseCustomNetPolicy) SetPolicyName(v string) *QueryNetListResponseBodyDataNetsDDoSDefenseCustomNetPolicy {
	s.PolicyName = &v
	return s
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefenseCustomNetPolicy) SetTmplId(v string) *QueryNetListResponseBodyDataNetsDDoSDefenseCustomNetPolicy {
	s.TmplId = &v
	return s
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefenseCustomNetPolicy) Validate() error {
	return dara.Validate(s)
}

type QueryNetListResponseBodyDataNetsDDoSDefenseDjPolicy struct {
	// The name of the mitigation policy.
	//
	// example:
	//
	// gf_origin_xxxxx
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s QueryNetListResponseBodyDataNetsDDoSDefenseDjPolicy) String() string {
	return dara.Prettify(s)
}

func (s QueryNetListResponseBodyDataNetsDDoSDefenseDjPolicy) GoString() string {
	return s.String()
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefenseDjPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefenseDjPolicy) SetPolicyName(v string) *QueryNetListResponseBodyDataNetsDDoSDefenseDjPolicy {
	s.PolicyName = &v
	return s
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefenseDjPolicy) Validate() error {
	return dara.Validate(s)
}

type QueryNetListResponseBodyDataNetsDDoSDefenseHoleTh struct {
	// The blackhole filtering threshold in Mbit/s. The value 0 indicates best-effort protection.
	//
	// example:
	//
	// 0
	ThreshMbps *int32 `json:"ThreshMbps,omitempty" xml:"ThreshMbps,omitempty"`
}

func (s QueryNetListResponseBodyDataNetsDDoSDefenseHoleTh) String() string {
	return dara.Prettify(s)
}

func (s QueryNetListResponseBodyDataNetsDDoSDefenseHoleTh) GoString() string {
	return s.String()
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefenseHoleTh) GetThreshMbps() *int32 {
	return s.ThreshMbps
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefenseHoleTh) SetThreshMbps(v int32) *QueryNetListResponseBodyDataNetsDDoSDefenseHoleTh {
	s.ThreshMbps = &v
	return s
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefenseHoleTh) Validate() error {
	return dara.Validate(s)
}

type QueryNetListResponseBodyDataNetsDeclared struct {
	// Indicates whether the CIDR block is advertised. Valid values:
	//
	// - 0: The CIDR block is not advertised.
	//
	// - 1: The CIDR block is advertised.
	//
	// example:
	//
	// 0
	Declared *string `json:"Declared,omitempty" xml:"Declared,omitempty"`
	// The region in which the CIDR block is advertised. Valid values:
	//
	// - oe26: China (Hong Kong).
	//
	// - us50: US (Virginia).
	//
	// - oe24: US (Silicon Valley).
	//
	// - oe28: Singapore.
	//
	// - jp141: Japan (Tokyo).
	//
	// - Gb145: UK (London).
	//
	// - oi39: Germany (Frankfurt).
	//
	// example:
	//
	// oe26
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s QueryNetListResponseBodyDataNetsDeclared) String() string {
	return dara.Prettify(s)
}

func (s QueryNetListResponseBodyDataNetsDeclared) GoString() string {
	return s.String()
}

func (s *QueryNetListResponseBodyDataNetsDeclared) GetDeclared() *string {
	return s.Declared
}

func (s *QueryNetListResponseBodyDataNetsDeclared) GetRegion() *string {
	return s.Region
}

func (s *QueryNetListResponseBodyDataNetsDeclared) SetDeclared(v string) *QueryNetListResponseBodyDataNetsDeclared {
	s.Declared = &v
	return s
}

func (s *QueryNetListResponseBodyDataNetsDeclared) SetRegion(v string) *QueryNetListResponseBodyDataNetsDeclared {
	s.Region = &v
	return s
}

func (s *QueryNetListResponseBodyDataNetsDeclared) Validate() error {
	return dara.Validate(s)
}
