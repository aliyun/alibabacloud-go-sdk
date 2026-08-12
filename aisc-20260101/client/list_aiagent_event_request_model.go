// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIAgentEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListAIAgentEventRequest
	GetAppId() *string
	SetAppName(v string) *ListAIAgentEventRequest
	GetAppName() *string
	SetAssetName(v string) *ListAIAgentEventRequest
	GetAssetName() *string
	SetAssetType(v string) *ListAIAgentEventRequest
	GetAssetType() *string
	SetCurrentPage(v int32) *ListAIAgentEventRequest
	GetCurrentPage() *int32
	SetInfraInstanceId(v string) *ListAIAgentEventRequest
	GetInfraInstanceId() *string
	SetInfraName(v string) *ListAIAgentEventRequest
	GetInfraName() *string
	SetInfraRegionId(v string) *ListAIAgentEventRequest
	GetInfraRegionId() *string
	SetLang(v string) *ListAIAgentEventRequest
	GetLang() *string
	SetPageSize(v int32) *ListAIAgentEventRequest
	GetPageSize() *int32
	SetRiskLevel(v string) *ListAIAgentEventRequest
	GetRiskLevel() *string
	SetRiskName(v string) *ListAIAgentEventRequest
	GetRiskName() *string
	SetSource(v string) *ListAIAgentEventRequest
	GetSource() *string
	SetStatus(v string) *ListAIAgentEventRequest
	GetStatus() *string
	SetStatusList(v []*string) *ListAIAgentEventRequest
	GetStatusList() []*string
	SetVendor(v string) *ListAIAgentEventRequest
	GetVendor() *string
}

type ListAIAgentEventRequest struct {
	// The ID of the agent application.
	//
	// example:
	//
	// 99f30e6b-8374-4a45-8830-439f178c5463
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Filters the agent list by application name.
	//
	// example:
	//
	// erH
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The asset name.
	//
	// example:
	//
	// 13.115.192.70
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// The type of the agent asset. Valid values:
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
	// identity
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The infrastructure instance ID.
	//
	// example:
	//
	// i-test
	InfraInstanceId *string `json:"InfraInstanceId,omitempty" xml:"InfraInstanceId,omitempty"`
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
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The risk level of the check item to query. Valid values:
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
	// The risk name. Fuzzy match is supported.
	//
	// example:
	//
	// defense
	RiskName *string `json:"RiskName,omitempty" xml:"RiskName,omitempty"`
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
	// CSPM
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The event status. Valid values:
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
	// fixed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of statuses.
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
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

func (s ListAIAgentEventRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAIAgentEventRequest) GoString() string {
	return s.String()
}

func (s *ListAIAgentEventRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListAIAgentEventRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListAIAgentEventRequest) GetAssetName() *string {
	return s.AssetName
}

func (s *ListAIAgentEventRequest) GetAssetType() *string {
	return s.AssetType
}

func (s *ListAIAgentEventRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAIAgentEventRequest) GetInfraInstanceId() *string {
	return s.InfraInstanceId
}

func (s *ListAIAgentEventRequest) GetInfraName() *string {
	return s.InfraName
}

func (s *ListAIAgentEventRequest) GetInfraRegionId() *string {
	return s.InfraRegionId
}

func (s *ListAIAgentEventRequest) GetLang() *string {
	return s.Lang
}

func (s *ListAIAgentEventRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAIAgentEventRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListAIAgentEventRequest) GetRiskName() *string {
	return s.RiskName
}

func (s *ListAIAgentEventRequest) GetSource() *string {
	return s.Source
}

func (s *ListAIAgentEventRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAIAgentEventRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *ListAIAgentEventRequest) GetVendor() *string {
	return s.Vendor
}

func (s *ListAIAgentEventRequest) SetAppId(v string) *ListAIAgentEventRequest {
	s.AppId = &v
	return s
}

func (s *ListAIAgentEventRequest) SetAppName(v string) *ListAIAgentEventRequest {
	s.AppName = &v
	return s
}

func (s *ListAIAgentEventRequest) SetAssetName(v string) *ListAIAgentEventRequest {
	s.AssetName = &v
	return s
}

func (s *ListAIAgentEventRequest) SetAssetType(v string) *ListAIAgentEventRequest {
	s.AssetType = &v
	return s
}

func (s *ListAIAgentEventRequest) SetCurrentPage(v int32) *ListAIAgentEventRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAIAgentEventRequest) SetInfraInstanceId(v string) *ListAIAgentEventRequest {
	s.InfraInstanceId = &v
	return s
}

func (s *ListAIAgentEventRequest) SetInfraName(v string) *ListAIAgentEventRequest {
	s.InfraName = &v
	return s
}

func (s *ListAIAgentEventRequest) SetInfraRegionId(v string) *ListAIAgentEventRequest {
	s.InfraRegionId = &v
	return s
}

func (s *ListAIAgentEventRequest) SetLang(v string) *ListAIAgentEventRequest {
	s.Lang = &v
	return s
}

func (s *ListAIAgentEventRequest) SetPageSize(v int32) *ListAIAgentEventRequest {
	s.PageSize = &v
	return s
}

func (s *ListAIAgentEventRequest) SetRiskLevel(v string) *ListAIAgentEventRequest {
	s.RiskLevel = &v
	return s
}

func (s *ListAIAgentEventRequest) SetRiskName(v string) *ListAIAgentEventRequest {
	s.RiskName = &v
	return s
}

func (s *ListAIAgentEventRequest) SetSource(v string) *ListAIAgentEventRequest {
	s.Source = &v
	return s
}

func (s *ListAIAgentEventRequest) SetStatus(v string) *ListAIAgentEventRequest {
	s.Status = &v
	return s
}

func (s *ListAIAgentEventRequest) SetStatusList(v []*string) *ListAIAgentEventRequest {
	s.StatusList = v
	return s
}

func (s *ListAIAgentEventRequest) SetVendor(v string) *ListAIAgentEventRequest {
	s.Vendor = &v
	return s
}

func (s *ListAIAgentEventRequest) Validate() error {
	return dara.Validate(s)
}
