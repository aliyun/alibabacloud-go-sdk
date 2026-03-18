// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosOriginInstanceBillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssetStatus(v int32) *DescribeDdosOriginInstanceBillResponseBody
	GetAssetStatus() *int32
	SetDebtStatus(v int64) *DescribeDdosOriginInstanceBillResponseBody
	GetDebtStatus() *int64
	SetFlowList(v []*DescribeDdosOriginInstanceBillResponseBodyFlowList) *DescribeDdosOriginInstanceBillResponseBody
	GetFlowList() []*DescribeDdosOriginInstanceBillResponseBodyFlowList
	SetFlowRegion(v map[string]interface{}) *DescribeDdosOriginInstanceBillResponseBody
	GetFlowRegion() map[string]interface{}
	SetInstanceId(v string) *DescribeDdosOriginInstanceBillResponseBody
	GetInstanceId() *string
	SetIpCount(v int64) *DescribeDdosOriginInstanceBillResponseBody
	GetIpCount() *int64
	SetIpCountOrFunctionList(v []*DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) *DescribeDdosOriginInstanceBillResponseBody
	GetIpCountOrFunctionList() []*DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList
	SetIpInfo(v string) *DescribeDdosOriginInstanceBillResponseBody
	GetIpInfo() *string
	SetMonthlySummaryList(v []*DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) *DescribeDdosOriginInstanceBillResponseBody
	GetMonthlySummaryList() []*DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList
	SetRequestId(v string) *DescribeDdosOriginInstanceBillResponseBody
	GetRequestId() *string
	SetStandardAssetsFlowList(v []*DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList) *DescribeDdosOriginInstanceBillResponseBody
	GetStandardAssetsFlowList() []*DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList
	SetStandardAssetsFlowRegion(v map[string]interface{}) *DescribeDdosOriginInstanceBillResponseBody
	GetStandardAssetsFlowRegion() map[string]interface{}
	SetStandardAssetsTotalFlowCn(v int64) *DescribeDdosOriginInstanceBillResponseBody
	GetStandardAssetsTotalFlowCn() *int64
	SetStandardAssetsTotalFlowOv(v int64) *DescribeDdosOriginInstanceBillResponseBody
	GetStandardAssetsTotalFlowOv() *int64
	SetStatus(v int64) *DescribeDdosOriginInstanceBillResponseBody
	GetStatus() *int64
	SetTotalFlowCn(v int64) *DescribeDdosOriginInstanceBillResponseBody
	GetTotalFlowCn() *int64
	SetTotalFlowOv(v int64) *DescribeDdosOriginInstanceBillResponseBody
	GetTotalFlowOv() *int64
}

type DescribeDdosOriginInstanceBillResponseBody struct {
	// The asset status.
	//
	// 	- **0**: No asset is added to the instance for protection.
	//
	// 	- **1**: Assets are added to the instance for protection.
	//
	// example:
	//
	// 0
	AssetStatus *int32 `json:"AssetStatus,omitempty" xml:"AssetStatus,omitempty"`
	// The payment status. Valid values:
	//
	// 	- **0**: The payment is not overdue.
	//
	// 	- **1**: The payment is overdue.
	//
	// example:
	//
	// 0
	DebtStatus *int64 `json:"DebtStatus,omitempty" xml:"DebtStatus,omitempty"`
	// The details about the traffic of EIPs with Anti-DDoS (Enhanced) enabled.
	FlowList []*DescribeDdosOriginInstanceBillResponseBodyFlowList `json:"FlowList,omitempty" xml:"FlowList,omitempty" type:"Repeated"`
	// The traffic distribution of EIPs with Anti-DDoS (Enhanced) enabled by region.
	//
	// example:
	//
	// {\\"cn-hongkong\\": 166491566}
	FlowRegion map[string]interface{} `json:"FlowRegion,omitempty" xml:"FlowRegion,omitempty"`
	// The ID of the Anti-DDoS Origin (Pay-as-you-go) instance to query.
	//
	// example:
	//
	// ddosorigin_cn-u7c3lcr9r02
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of protected IP addresses.
	//
	// example:
	//
	// 15
	IpCount *int64 `json:"IpCount,omitempty" xml:"IpCount,omitempty"`
	// The protected IP addresses and enabled features.
	IpCountOrFunctionList []*DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList `json:"IpCountOrFunctionList,omitempty" xml:"IpCountOrFunctionList,omitempty" type:"Repeated"`
	// The IP address distribution. The JSON struct contains the following fields:
	//
	// 	- **eipCnIpCount**: the number of EIPs with Anti-DDoS (Enhanced) enabled in the Chinese mainland.
	//
	// 	- **eipOvIpCount**: the number of EIPs with Anti-DDoS (Enhanced) enabled outside the Chinese mainland.
	//
	// 	- **standardAssetsCnIpCount**: the number of IP addresses of regular Alibaba Cloud services in the Chinese mainland.
	//
	// 	- **standardAssetsOvIpCount**: the number of IP addresses of regular Alibaba Cloud services outside the Chinese mainland.
	//
	// example:
	//
	// {\\"eipCnIpCount\\":6,\\"eipOvIpCount\\":17,\\"standardAssetsCnIpCount\\":2,\\"standardAssetsOvIpCount\\":0}
	IpInfo *string `json:"IpInfo,omitempty" xml:"IpInfo,omitempty"`
	// The information about the monthly summary bills.
	MonthlySummaryList []*DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList `json:"MonthlySummaryList,omitempty" xml:"MonthlySummaryList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 72155560-F343-55C8-82FE-ED4D7E4AA97E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details about the traffic of regular Alibaba Cloud services.
	StandardAssetsFlowList []*DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList `json:"StandardAssetsFlowList,omitempty" xml:"StandardAssetsFlowList,omitempty" type:"Repeated"`
	// The traffic distribution of regular Alibaba Cloud services by region.
	//
	// example:
	//
	// {\\"cn-hongkong\\": 166491566}
	StandardAssetsFlowRegion map[string]interface{} `json:"StandardAssetsFlowRegion,omitempty" xml:"StandardAssetsFlowRegion,omitempty"`
	// The total traffic of regular Alibaba Cloud services in the Chinese mainland in the current month.
	//
	// example:
	//
	// 0
	StandardAssetsTotalFlowCn *int64 `json:"StandardAssetsTotalFlowCn,omitempty" xml:"StandardAssetsTotalFlowCn,omitempty"`
	// The total traffic of regular Alibaba Cloud services outside the Chinese mainland in the current month.
	//
	// example:
	//
	// 0
	StandardAssetsTotalFlowOv *int64 `json:"StandardAssetsTotalFlowOv,omitempty" xml:"StandardAssetsTotalFlowOv,omitempty"`
	// The instance status. Valid values:
	//
	// 	- **1**: normal
	//
	// 	- **2**: expired
	//
	// 	- **3**: released
	//
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total traffic of EIPs with Anti-DDoS (Enhanced) enabled in the Chinese mainland in the current month. Unit: bytes.
	//
	// example:
	//
	// 6302081067
	TotalFlowCn *int64 `json:"TotalFlowCn,omitempty" xml:"TotalFlowCn,omitempty"`
	// The total traffic of EIPs with Anti-DDoS (Enhanced) enabled outside the Chinese mainland in the current month. Unit: bytes.
	//
	// example:
	//
	// 6204918019
	TotalFlowOv *int64 `json:"TotalFlowOv,omitempty" xml:"TotalFlowOv,omitempty"`
}

func (s DescribeDdosOriginInstanceBillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosOriginInstanceBillResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosOriginInstanceBillResponseBody) GetAssetStatus() *int32 {
	return s.AssetStatus
}

func (s *DescribeDdosOriginInstanceBillResponseBody) GetDebtStatus() *int64 {
	return s.DebtStatus
}

func (s *DescribeDdosOriginInstanceBillResponseBody) GetFlowList() []*DescribeDdosOriginInstanceBillResponseBodyFlowList {
	return s.FlowList
}

func (s *DescribeDdosOriginInstanceBillResponseBody) GetFlowRegion() map[string]interface{} {
	return s.FlowRegion
}

func (s *DescribeDdosOriginInstanceBillResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDdosOriginInstanceBillResponseBody) GetIpCount() *int64 {
	return s.IpCount
}

func (s *DescribeDdosOriginInstanceBillResponseBody) GetIpCountOrFunctionList() []*DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList {
	return s.IpCountOrFunctionList
}

func (s *DescribeDdosOriginInstanceBillResponseBody) GetIpInfo() *string {
	return s.IpInfo
}

func (s *DescribeDdosOriginInstanceBillResponseBody) GetMonthlySummaryList() []*DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList {
	return s.MonthlySummaryList
}

func (s *DescribeDdosOriginInstanceBillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDdosOriginInstanceBillResponseBody) GetStandardAssetsFlowList() []*DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList {
	return s.StandardAssetsFlowList
}

func (s *DescribeDdosOriginInstanceBillResponseBody) GetStandardAssetsFlowRegion() map[string]interface{} {
	return s.StandardAssetsFlowRegion
}

func (s *DescribeDdosOriginInstanceBillResponseBody) GetStandardAssetsTotalFlowCn() *int64 {
	return s.StandardAssetsTotalFlowCn
}

func (s *DescribeDdosOriginInstanceBillResponseBody) GetStandardAssetsTotalFlowOv() *int64 {
	return s.StandardAssetsTotalFlowOv
}

func (s *DescribeDdosOriginInstanceBillResponseBody) GetStatus() *int64 {
	return s.Status
}

func (s *DescribeDdosOriginInstanceBillResponseBody) GetTotalFlowCn() *int64 {
	return s.TotalFlowCn
}

func (s *DescribeDdosOriginInstanceBillResponseBody) GetTotalFlowOv() *int64 {
	return s.TotalFlowOv
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetAssetStatus(v int32) *DescribeDdosOriginInstanceBillResponseBody {
	s.AssetStatus = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetDebtStatus(v int64) *DescribeDdosOriginInstanceBillResponseBody {
	s.DebtStatus = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetFlowList(v []*DescribeDdosOriginInstanceBillResponseBodyFlowList) *DescribeDdosOriginInstanceBillResponseBody {
	s.FlowList = v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetFlowRegion(v map[string]interface{}) *DescribeDdosOriginInstanceBillResponseBody {
	s.FlowRegion = v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetInstanceId(v string) *DescribeDdosOriginInstanceBillResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetIpCount(v int64) *DescribeDdosOriginInstanceBillResponseBody {
	s.IpCount = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetIpCountOrFunctionList(v []*DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) *DescribeDdosOriginInstanceBillResponseBody {
	s.IpCountOrFunctionList = v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetIpInfo(v string) *DescribeDdosOriginInstanceBillResponseBody {
	s.IpInfo = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetMonthlySummaryList(v []*DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) *DescribeDdosOriginInstanceBillResponseBody {
	s.MonthlySummaryList = v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetRequestId(v string) *DescribeDdosOriginInstanceBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetStandardAssetsFlowList(v []*DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList) *DescribeDdosOriginInstanceBillResponseBody {
	s.StandardAssetsFlowList = v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetStandardAssetsFlowRegion(v map[string]interface{}) *DescribeDdosOriginInstanceBillResponseBody {
	s.StandardAssetsFlowRegion = v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetStandardAssetsTotalFlowCn(v int64) *DescribeDdosOriginInstanceBillResponseBody {
	s.StandardAssetsTotalFlowCn = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetStandardAssetsTotalFlowOv(v int64) *DescribeDdosOriginInstanceBillResponseBody {
	s.StandardAssetsTotalFlowOv = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetStatus(v int64) *DescribeDdosOriginInstanceBillResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetTotalFlowCn(v int64) *DescribeDdosOriginInstanceBillResponseBody {
	s.TotalFlowCn = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetTotalFlowOv(v int64) *DescribeDdosOriginInstanceBillResponseBody {
	s.TotalFlowOv = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) Validate() error {
	if s.FlowList != nil {
		for _, item := range s.FlowList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IpCountOrFunctionList != nil {
		for _, item := range s.IpCountOrFunctionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MonthlySummaryList != nil {
		for _, item := range s.MonthlySummaryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.StandardAssetsFlowList != nil {
		for _, item := range s.StandardAssetsFlowList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDdosOriginInstanceBillResponseBodyFlowList struct {
	// The traffic distribution by region. The JSON struct contains the following fields:
	//
	// 	- **bytes**: the traffic volume of EIPs with Anti-DDoS (Enhanced) enabled in a region. Unit: bytes.
	//
	// 	- **memberUid**: the owner account.
	//
	// 	- **instanceId**: the ID of the pay-as-you-go instance that protects the EIPs with Anti-DDoS (Enhanced) enabled.
	//
	// 	- **ip**: the EIPs with Anti-DDoS (Enhanced) enabled.
	//
	// 	- **region**: the region.
	//
	// >  If the memberUid field in the JSON struct is empty, the information about the current account is returned. The value of the bytes parameter in the outermost level of the JSON struct indicates the total traffic, and the values of the bytes parameters in inner levels indicate the traffic of the account.
	//
	// example:
	//
	// [{\\"bytes\\":79282719,\\"memberUid\\":\\"\\",\\"regionFlows\\":{\\"cn-hangzhou\\":[{\\"bytes\\":79282719,\\"instanceId\\":\\"ddosorigin_cn-u7c3lcr9r02\\",\\"ip\\":\\"47.118.168.57\\",\\"region\\":\\"cn-hangzhou\\"}]}}]
	MemberFlow *string `json:"MemberFlow,omitempty" xml:"MemberFlow,omitempty"`
	// The traffic distribution by region. The JSON struct contains the following fields:
	//
	// 	- **bytes**: the traffic volume of EIPs with Anti-DDoS (Enhanced) enabled in a region. Unit: bytes.
	//
	// 	- **instanceId**: the ID of the pay-as-you-go instance that protects the EIPs with Anti-DDoS (Enhanced) enabled.
	//
	// 	- **ip**: the EIPs with Anti-DDoS (Enhanced) enabled.
	//
	// 	- **region**: the region.
	//
	// example:
	//
	// {\\"cn-hangzhou\\":[{\\"bytes\\":0,\\"instanceId\\":\\"ddosorigin_cn-u7c3lcr9r02\\",\\"ip\\":\\"47.118.168.124\\",\\"region\\":\\"cn-hangzhou\\"}]}
	RegionFlow *string `json:"RegionFlow,omitempty" xml:"RegionFlow,omitempty"`
	// The timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1620951900
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The total IP traffic of regular Alibaba Cloud services.
	//
	// example:
	//
	// 6302081067
	TotalBillFlow *int64 `json:"TotalBillFlow,omitempty" xml:"TotalBillFlow,omitempty"`
	// The traffic of EIPs with Anti-DDoS (Enhanced) enabled. Unit: bytes.
	//
	// example:
	//
	// 6302081067
	TotalFlow *int64 `json:"TotalFlow,omitempty" xml:"TotalFlow,omitempty"`
}

func (s DescribeDdosOriginInstanceBillResponseBodyFlowList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosOriginInstanceBillResponseBodyFlowList) GoString() string {
	return s.String()
}

func (s *DescribeDdosOriginInstanceBillResponseBodyFlowList) GetMemberFlow() *string {
	return s.MemberFlow
}

func (s *DescribeDdosOriginInstanceBillResponseBodyFlowList) GetRegionFlow() *string {
	return s.RegionFlow
}

func (s *DescribeDdosOriginInstanceBillResponseBodyFlowList) GetTime() *int64 {
	return s.Time
}

func (s *DescribeDdosOriginInstanceBillResponseBodyFlowList) GetTotalBillFlow() *int64 {
	return s.TotalBillFlow
}

func (s *DescribeDdosOriginInstanceBillResponseBodyFlowList) GetTotalFlow() *int64 {
	return s.TotalFlow
}

func (s *DescribeDdosOriginInstanceBillResponseBodyFlowList) SetMemberFlow(v string) *DescribeDdosOriginInstanceBillResponseBodyFlowList {
	s.MemberFlow = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyFlowList) SetRegionFlow(v string) *DescribeDdosOriginInstanceBillResponseBodyFlowList {
	s.RegionFlow = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyFlowList) SetTime(v int64) *DescribeDdosOriginInstanceBillResponseBodyFlowList {
	s.Time = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyFlowList) SetTotalBillFlow(v int64) *DescribeDdosOriginInstanceBillResponseBodyFlowList {
	s.TotalBillFlow = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyFlowList) SetTotalFlow(v int64) *DescribeDdosOriginInstanceBillResponseBodyFlowList {
	s.TotalFlow = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyFlowList) Validate() error {
	return dara.Validate(s)
}

type DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList struct {
	// The application scope of the instance. Valid values:
	//
	// 	- **only_mainland_china**: regions in the Chinese mainland.
	//
	// 	- **global**: all regions.
	//
	// 	- **international_and_hmt**: regions outside the Chinese mainland.
	//
	// example:
	//
	// global
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// The number of IP addresses protected by the pay-as-you-go instance in the Chinese mainland.
	//
	// example:
	//
	// 5
	IpCntCn *int64 `json:"IpCntCn,omitempty" xml:"IpCntCn,omitempty"`
	// The number of IP addresses protected by the pay-as-you-go instance outside the Chinese mainland.
	//
	// example:
	//
	// 5
	IpCntOv *int64 `json:"IpCntOv,omitempty" xml:"IpCntOv,omitempty"`
	// The bill distribution by account. The JSON struct contains the following fields:
	//
	// 	- **eipCnIpCount**: the number of EIPs with Anti-DDoS (Enhanced) enabled in the Chinese mainland.
	//
	// 	- **eipOvIpCount**: the number of EIPs with Anti-DDoS (Enhanced) enabled outside the Chinese mainland.
	//
	// 	- **memberUid**: the owner account.
	//
	// 	- **standardAssetsCnIpCount**: the number of IP addresses of regular Alibaba Cloud services in the Chinese mainland.
	//
	// 	- **standardAssetsOvIpCount**: the number of IP addresses of regular Alibaba Cloud services outside the Chinese mainland.
	//
	// >  If the memberUid field in the JSON struct is empty, the information about the current account is returned.
	//
	// example:
	//
	// [{\\"eipCnIpCount\\":3,\\"eipOvIpCount\\":18,\\"memberUid\\":\\"\\",\\"standardAssetsCnIpCount\\":2,\\"standardAssetsOvIpCount\\":0},{\\"eipCnIpCount\\":3,\\"eipOvIpCount\\":0,\\"memberUid\\":\\"1776997906319249\\",\\"standardAssetsCnIpCount\\":0,\\"standardAssetsOvIpCount\\":0}]
	MemberIpCnt *string `json:"MemberIpCnt,omitempty" xml:"MemberIpCnt,omitempty"`
	// The billing time. Unit: milliseconds.
	//
	// example:
	//
	// 1680278400000
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) GoString() string {
	return s.String()
}

func (s *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) GetCoverage() *string {
	return s.Coverage
}

func (s *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) GetIpCntCn() *int64 {
	return s.IpCntCn
}

func (s *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) GetIpCntOv() *int64 {
	return s.IpCntOv
}

func (s *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) GetMemberIpCnt() *string {
	return s.MemberIpCnt
}

func (s *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) GetTime() *int64 {
	return s.Time
}

func (s *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) SetCoverage(v string) *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList {
	s.Coverage = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) SetIpCntCn(v int64) *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList {
	s.IpCntCn = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) SetIpCntOv(v int64) *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList {
	s.IpCntOv = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) SetMemberIpCnt(v string) *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList {
	s.MemberIpCnt = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) SetTime(v int64) *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList {
	s.Time = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) Validate() error {
	return dara.Validate(s)
}

type DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList struct {
	// The number of days that the instance is activated.
	//
	// example:
	//
	// 30
	EnableDays *int32 `json:"EnableDays,omitempty" xml:"EnableDays,omitempty"`
	// The total traffic of EIPs with Anti-DDoS (Enhanced) enabled in the Chinese mainland. Unit: bytes.
	//
	// example:
	//
	// 123456
	FlowCn *int64 `json:"FlowCn,omitempty" xml:"FlowCn,omitempty"`
	// The total traffic of EIPs with Anti-DDoS (Enhanced) enabled outside the Chinese mainland. Unit: bytes.
	//
	// example:
	//
	// 123456
	FlowIntl *int64 `json:"FlowIntl,omitempty" xml:"FlowIntl,omitempty"`
	// The total number of protected IP addresses in the Chinese mainland.
	//
	// >  The total number of protected IP addresses is the sum of the daily numbers of protected IP addresses in a month.
	//
	// example:
	//
	// 28
	IpCountCn *int32 `json:"IpCountCn,omitempty" xml:"IpCountCn,omitempty"`
	// The total number of protected IP addresses outside the Chinese mainland.
	//
	// >  The total number of protected IP addresses is the sum of the daily numbers of protected IP addresses in a month.
	//
	// example:
	//
	// 30
	IpCountIntl *int32 `json:"IpCountIntl,omitempty" xml:"IpCountIntl,omitempty"`
	// The ID of the member.
	//
	// example:
	//
	// 112873971277****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The total traffic of regular Alibaba Cloud services in the Chinese mainland. Unit: bytes.
	//
	// example:
	//
	// 123456
	StandardAssetsFlowCn *int64 `json:"StandardAssetsFlowCn,omitempty" xml:"StandardAssetsFlowCn,omitempty"`
	// The total traffic of regular Alibaba Cloud services outside the Chinese mainland. Unit: bytes.
	//
	// example:
	//
	// 123456
	StandardAssetsFlowIntl *int64 `json:"StandardAssetsFlowIntl,omitempty" xml:"StandardAssetsFlowIntl,omitempty"`
	// The ID of the administrator account.
	//
	// example:
	//
	// 102518028277****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) GoString() string {
	return s.String()
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) GetEnableDays() *int32 {
	return s.EnableDays
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) GetFlowCn() *int64 {
	return s.FlowCn
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) GetFlowIntl() *int64 {
	return s.FlowIntl
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) GetIpCountCn() *int32 {
	return s.IpCountCn
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) GetIpCountIntl() *int32 {
	return s.IpCountIntl
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) GetStandardAssetsFlowCn() *int64 {
	return s.StandardAssetsFlowCn
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) GetStandardAssetsFlowIntl() *int64 {
	return s.StandardAssetsFlowIntl
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) GetUid() *string {
	return s.Uid
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) SetEnableDays(v int32) *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList {
	s.EnableDays = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) SetFlowCn(v int64) *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList {
	s.FlowCn = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) SetFlowIntl(v int64) *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList {
	s.FlowIntl = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) SetIpCountCn(v int32) *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList {
	s.IpCountCn = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) SetIpCountIntl(v int32) *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList {
	s.IpCountIntl = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) SetMemberUid(v string) *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList {
	s.MemberUid = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) SetStandardAssetsFlowCn(v int64) *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList {
	s.StandardAssetsFlowCn = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) SetStandardAssetsFlowIntl(v int64) *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList {
	s.StandardAssetsFlowIntl = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) SetUid(v string) *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList {
	s.Uid = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) Validate() error {
	return dara.Validate(s)
}

type DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList struct {
	// The traffic distribution by region. The JSON struct contains the following fields:
	//
	// 	- **bytes**: the traffic volume of regular Alibaba Cloud services in a region. Unit: bytes.
	//
	// 	- **memberUid**: the owner account.
	//
	// 	- **instanceId**: the ID of the pay-as-you-go instance that protects the regular Alibaba Cloud services.
	//
	// 	- **ip**: the IP address of the regular Alibaba Cloud service protected by the Anti-DDoS Origin instance.
	//
	// 	- **region**: the region.
	//
	// >  If the memberUid field in the JSON struct is empty, the information about the current account is returned. The value of the bytes parameter in the outermost level of the JSON struct indicates the total traffic, and the values of the bytes parameters in inner levels indicate the traffic of the account.
	//
	// example:
	//
	// [{\\"bytes\\":79282719,\\"memberUid\\":\\"\\",\\"regionFlows\\":{\\"cn-hangzhou\\":[{\\"bytes\\":79282719,\\"instanceId\\":\\"ddosorigin_cn-u7c3lcr9r02\\",\\"ip\\":\\"47.118.168.57\\",\\"region\\":\\"cn-hangzhou\\"}]}}]
	MemberFlow *string `json:"MemberFlow,omitempty" xml:"MemberFlow,omitempty"`
	// The traffic distribution by region. The JSON struct contains the following fields:
	//
	// 	- **bytes**: the traffic volume of regular Alibaba Cloud services in a region. Unit: bytes.
	//
	// 	- **instanceId**: the ID of the pay-as-you-go instance that protects the regular Alibaba Cloud services.
	//
	// 	- **ip**: the IP address protected by the Anti-DDoS Origin instance.
	//
	// 	- **region**: the region.
	//
	// example:
	//
	// {\\"cn-hangzhou\\":[{\\"bytes\\":0,\\"instanceId\\":\\"ddosorigin_cn-u7c3lcr9r02\\",\\"ip\\":\\"47.118.168.124\\",\\"region\\":\\"cn-hangzhou\\"}]}
	RegionFlow *string `json:"RegionFlow,omitempty" xml:"RegionFlow,omitempty"`
	// The timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1679846400000
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The traffic of regular Alibaba Cloud services. Unit: bytes.
	//
	// example:
	//
	// 6302081067
	TotalFlow *int64 `json:"TotalFlow,omitempty" xml:"TotalFlow,omitempty"`
}

func (s DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList) GoString() string {
	return s.String()
}

func (s *DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList) GetMemberFlow() *string {
	return s.MemberFlow
}

func (s *DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList) GetRegionFlow() *string {
	return s.RegionFlow
}

func (s *DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList) GetTime() *int64 {
	return s.Time
}

func (s *DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList) GetTotalFlow() *int64 {
	return s.TotalFlow
}

func (s *DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList) SetMemberFlow(v string) *DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList {
	s.MemberFlow = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList) SetRegionFlow(v string) *DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList {
	s.RegionFlow = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList) SetTime(v int64) *DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList {
	s.Time = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList) SetTotalFlow(v int64) *DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList {
	s.TotalFlow = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList) Validate() error {
	return dara.Validate(s)
}
