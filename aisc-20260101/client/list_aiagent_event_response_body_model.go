// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIAgentEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListAIAgentEventResponseBodyData) *ListAIAgentEventResponseBody
	GetData() []*ListAIAgentEventResponseBodyData
	SetPageInfo(v *ListAIAgentEventResponseBodyPageInfo) *ListAIAgentEventResponseBody
	GetPageInfo() *ListAIAgentEventResponseBodyPageInfo
	SetRequestId(v string) *ListAIAgentEventResponseBody
	GetRequestId() *string
}

type ListAIAgentEventResponseBody struct {
	// The list of event information returned.
	Data []*ListAIAgentEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListAIAgentEventResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 9FDE3D6F-26BD-5937-B0E5-8F47962B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAIAgentEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAIAgentEventResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIAgentEventResponseBody) GetData() []*ListAIAgentEventResponseBodyData {
	return s.Data
}

func (s *ListAIAgentEventResponseBody) GetPageInfo() *ListAIAgentEventResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListAIAgentEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAIAgentEventResponseBody) SetData(v []*ListAIAgentEventResponseBodyData) *ListAIAgentEventResponseBody {
	s.Data = v
	return s
}

func (s *ListAIAgentEventResponseBody) SetPageInfo(v *ListAIAgentEventResponseBodyPageInfo) *ListAIAgentEventResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListAIAgentEventResponseBody) SetRequestId(v string) *ListAIAgentEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAIAgentEventResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAIAgentEventResponseBodyData struct {
	// The application ID.
	//
	// example:
	//
	// 5zuzvcfe
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// prod-chagee-bc-activity-elespin
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The asset name.
	//
	// example:
	//
	// 25.2.2.83
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// The type of the risky asset. Valid values:
	//
	// 1. rag
	//
	// 2. internet
	//
	// 3. datasets
	//
	// 4. tool
	//
	// 5. model
	//
	// 6. skill
	//
	// 7. app
	//
	// 8. identity
	//
	// example:
	//
	// tool
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The check time.
	//
	// example:
	//
	// 1763949968
	CheckTime *string `json:"CheckTime,omitempty" xml:"CheckTime,omitempty"`
	// The event handling time.
	//
	// example:
	//
	// 1763949968
	HandleTime *string `json:"HandleTime,omitempty" xml:"HandleTime,omitempty"`
	// The primary key ID.
	//
	// example:
	//
	// 17616
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The infrastructure instance ID.
	//
	// example:
	//
	// i-test
	InfraInstanceId *string `json:"InfraInstanceId,omitempty" xml:"InfraInstanceId,omitempty"`
	// The public IP address of the infrastructure.
	//
	// example:
	//
	// 1.2.3.4
	InfraInternetIp *string `json:"InfraInternetIp,omitempty" xml:"InfraInternetIp,omitempty"`
	// The private IP address of the infrastructure.
	//
	// example:
	//
	// 10.0.0.3
	InfraIntranetIp *string `json:"InfraIntranetIp,omitempty" xml:"InfraIntranetIp,omitempty"`
	// The infrastructure name.
	//
	// example:
	//
	// test
	InfraName *string `json:"InfraName,omitempty" xml:"InfraName,omitempty"`
	// The infrastructure region.
	//
	// example:
	//
	// cn-shanghai
	InfraRegionId *string `json:"InfraRegionId,omitempty" xml:"InfraRegionId,omitempty"`
	// The infrastructure type.
	//
	// example:
	//
	// ECS
	InfraType *string `json:"InfraType,omitempty" xml:"InfraType,omitempty"`
	// The risk description.
	//
	// example:
	//
	// The workflow does not have AI security guardrails enabled, which may lead to compliance violations, prompt injection and bypass, sensitive data leaks, and other risks
	RiskDesc *string `json:"RiskDesc,omitempty" xml:"RiskDesc,omitempty"`
	// The risk level of the detected alert. Valid values:
	//
	// - **high**: High.
	//
	// - **medium**: Medium.
	//
	// - **low**: Low.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The risk name.
	//
	// example:
	//
	// Weak password
	RiskName *string `json:"RiskName,omitempty" xml:"RiskName,omitempty"`
	SkillId  *int64  `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The event source. Valid values:
	//
	// 1. cspm
	//
	// 2. aiguard
	//
	// 3. SASE
	//
	// 4. SAS
	//
	// 5. Agent-Runtime-Guard
	//
	// example:
	//
	// SASE
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status. Valid values:
	//
	// 1. unhandled: Pending.
	//
	// 2. handling: Being processed.
	//
	// 3. fixed: Fixed.
	//
	// 4. ignored: Ignored.
	//
	// 5. rescanned: Rescanned.
	//
	// example:
	//
	// unhandled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The cloud asset vendor. Valid values:
	//
	// - **DIFY**: DIFY.
	//
	// - **BAILIAN**: BAILIAN.
	//
	// - **VOLCAI**: VOLCAI.
	//
	// - **AGENTRUN**: AGENTRUN.
	//
	// - **PAI**: PAI.
	//
	// - **OpenClaw**: OpenClaw.
	//
	// example:
	//
	// DIFY
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ListAIAgentEventResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAIAgentEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAIAgentEventResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *ListAIAgentEventResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *ListAIAgentEventResponseBodyData) GetAssetName() *string {
	return s.AssetName
}

func (s *ListAIAgentEventResponseBodyData) GetAssetType() *string {
	return s.AssetType
}

func (s *ListAIAgentEventResponseBodyData) GetCheckTime() *string {
	return s.CheckTime
}

func (s *ListAIAgentEventResponseBodyData) GetHandleTime() *string {
	return s.HandleTime
}

func (s *ListAIAgentEventResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListAIAgentEventResponseBodyData) GetInfraInstanceId() *string {
	return s.InfraInstanceId
}

func (s *ListAIAgentEventResponseBodyData) GetInfraInternetIp() *string {
	return s.InfraInternetIp
}

func (s *ListAIAgentEventResponseBodyData) GetInfraIntranetIp() *string {
	return s.InfraIntranetIp
}

func (s *ListAIAgentEventResponseBodyData) GetInfraName() *string {
	return s.InfraName
}

func (s *ListAIAgentEventResponseBodyData) GetInfraRegionId() *string {
	return s.InfraRegionId
}

func (s *ListAIAgentEventResponseBodyData) GetInfraType() *string {
	return s.InfraType
}

func (s *ListAIAgentEventResponseBodyData) GetRiskDesc() *string {
	return s.RiskDesc
}

func (s *ListAIAgentEventResponseBodyData) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListAIAgentEventResponseBodyData) GetRiskName() *string {
	return s.RiskName
}

func (s *ListAIAgentEventResponseBodyData) GetSkillId() *int64 {
	return s.SkillId
}

func (s *ListAIAgentEventResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *ListAIAgentEventResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListAIAgentEventResponseBodyData) GetVendor() *string {
	return s.Vendor
}

func (s *ListAIAgentEventResponseBodyData) SetAppId(v string) *ListAIAgentEventResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListAIAgentEventResponseBodyData) SetAppName(v string) *ListAIAgentEventResponseBodyData {
	s.AppName = &v
	return s
}

func (s *ListAIAgentEventResponseBodyData) SetAssetName(v string) *ListAIAgentEventResponseBodyData {
	s.AssetName = &v
	return s
}

func (s *ListAIAgentEventResponseBodyData) SetAssetType(v string) *ListAIAgentEventResponseBodyData {
	s.AssetType = &v
	return s
}

func (s *ListAIAgentEventResponseBodyData) SetCheckTime(v string) *ListAIAgentEventResponseBodyData {
	s.CheckTime = &v
	return s
}

func (s *ListAIAgentEventResponseBodyData) SetHandleTime(v string) *ListAIAgentEventResponseBodyData {
	s.HandleTime = &v
	return s
}

func (s *ListAIAgentEventResponseBodyData) SetId(v int64) *ListAIAgentEventResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAIAgentEventResponseBodyData) SetInfraInstanceId(v string) *ListAIAgentEventResponseBodyData {
	s.InfraInstanceId = &v
	return s
}

func (s *ListAIAgentEventResponseBodyData) SetInfraInternetIp(v string) *ListAIAgentEventResponseBodyData {
	s.InfraInternetIp = &v
	return s
}

func (s *ListAIAgentEventResponseBodyData) SetInfraIntranetIp(v string) *ListAIAgentEventResponseBodyData {
	s.InfraIntranetIp = &v
	return s
}

func (s *ListAIAgentEventResponseBodyData) SetInfraName(v string) *ListAIAgentEventResponseBodyData {
	s.InfraName = &v
	return s
}

func (s *ListAIAgentEventResponseBodyData) SetInfraRegionId(v string) *ListAIAgentEventResponseBodyData {
	s.InfraRegionId = &v
	return s
}

func (s *ListAIAgentEventResponseBodyData) SetInfraType(v string) *ListAIAgentEventResponseBodyData {
	s.InfraType = &v
	return s
}

func (s *ListAIAgentEventResponseBodyData) SetRiskDesc(v string) *ListAIAgentEventResponseBodyData {
	s.RiskDesc = &v
	return s
}

func (s *ListAIAgentEventResponseBodyData) SetRiskLevel(v string) *ListAIAgentEventResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *ListAIAgentEventResponseBodyData) SetRiskName(v string) *ListAIAgentEventResponseBodyData {
	s.RiskName = &v
	return s
}

func (s *ListAIAgentEventResponseBodyData) SetSkillId(v int64) *ListAIAgentEventResponseBodyData {
	s.SkillId = &v
	return s
}

func (s *ListAIAgentEventResponseBodyData) SetSource(v string) *ListAIAgentEventResponseBodyData {
	s.Source = &v
	return s
}

func (s *ListAIAgentEventResponseBodyData) SetStatus(v string) *ListAIAgentEventResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListAIAgentEventResponseBodyData) SetVendor(v string) *ListAIAgentEventResponseBodyData {
	s.Vendor = &v
	return s
}

func (s *ListAIAgentEventResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListAIAgentEventResponseBodyPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The NextToken value returned when the NextToken-based pagination method is used.
	//
	// example:
	//
	// d6yVpGGP9cH8f9AWtqEXqOawJdolFvFeqJJSIPnYLoGc7/XPd5nbDfZcn1mJCj66Ep3Gbr55tl4NuBtNwsc0A0qvqC2Onfm9h2QmtG8HhaulnPkGmBnhntKqJmpRptTU
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records in the query result.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAIAgentEventResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAIAgentEventResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListAIAgentEventResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAIAgentEventResponseBodyPageInfo) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAIAgentEventResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAIAgentEventResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAIAgentEventResponseBodyPageInfo) SetCurrentPage(v int32) *ListAIAgentEventResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListAIAgentEventResponseBodyPageInfo) SetNextToken(v string) *ListAIAgentEventResponseBodyPageInfo {
	s.NextToken = &v
	return s
}

func (s *ListAIAgentEventResponseBodyPageInfo) SetPageSize(v int32) *ListAIAgentEventResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListAIAgentEventResponseBodyPageInfo) SetTotalCount(v int32) *ListAIAgentEventResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListAIAgentEventResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
