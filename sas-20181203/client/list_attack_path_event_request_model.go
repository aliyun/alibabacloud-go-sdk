// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAttackPathEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttackPathAssetList(v []*ListAttackPathEventRequestAttackPathAssetList) *ListAttackPathEventRequest
	GetAttackPathAssetList() []*ListAttackPathEventRequestAttackPathAssetList
	SetCurrentPage(v int32) *ListAttackPathEventRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *ListAttackPathEventRequest
	GetEndTime() *int64
	SetLang(v string) *ListAttackPathEventRequest
	GetLang() *string
	SetPageSize(v int32) *ListAttackPathEventRequest
	GetPageSize() *int32
	SetPathNameDesc(v string) *ListAttackPathEventRequest
	GetPathNameDesc() *string
	SetPathType(v string) *ListAttackPathEventRequest
	GetPathType() *string
	SetRiskLevelList(v []*string) *ListAttackPathEventRequest
	GetRiskLevelList() []*string
	SetStartTime(v int64) *ListAttackPathEventRequest
	GetStartTime() *int64
}

type ListAttackPathEventRequest struct {
	// The list of cloud service assets in the attack path.
	AttackPathAssetList []*ListAttackPathEventRequestAttackPathAssetList `json:"AttackPathAssetList,omitempty" xml:"AttackPathAssetList,omitempty" type:"Repeated"`
	// The page number of the results to return. Default value: 1, which indicates the first page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end time as a timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1668064495000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries per page in a paged query. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The path name description.
	//
	// > Call [ListAvailableAttackPath](~~ListAvailableAttackPath~~) to query path name descriptions.
	//
	// example:
	//
	// ECS Instance Can Obtain Long-term Access Credential by Enabling Console Logon for RAM User
	PathNameDesc *string `json:"PathNameDesc,omitempty" xml:"PathNameDesc,omitempty"`
	// The path type.
	//
	// > Call [ListAvailableAttackPath](~~ListAvailableAttackPath~~) to query path types.
	//
	// example:
	//
	// role_escalation
	PathType *string `json:"PathType,omitempty" xml:"PathType,omitempty"`
	// The list of risk levels.
	RiskLevelList []*string `json:"RiskLevelList,omitempty" xml:"RiskLevelList,omitempty" type:"Repeated"`
	// The start time as a timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1666886400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAttackPathEventRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAttackPathEventRequest) GoString() string {
	return s.String()
}

func (s *ListAttackPathEventRequest) GetAttackPathAssetList() []*ListAttackPathEventRequestAttackPathAssetList {
	return s.AttackPathAssetList
}

func (s *ListAttackPathEventRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAttackPathEventRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListAttackPathEventRequest) GetLang() *string {
	return s.Lang
}

func (s *ListAttackPathEventRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAttackPathEventRequest) GetPathNameDesc() *string {
	return s.PathNameDesc
}

func (s *ListAttackPathEventRequest) GetPathType() *string {
	return s.PathType
}

func (s *ListAttackPathEventRequest) GetRiskLevelList() []*string {
	return s.RiskLevelList
}

func (s *ListAttackPathEventRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListAttackPathEventRequest) SetAttackPathAssetList(v []*ListAttackPathEventRequestAttackPathAssetList) *ListAttackPathEventRequest {
	s.AttackPathAssetList = v
	return s
}

func (s *ListAttackPathEventRequest) SetCurrentPage(v int32) *ListAttackPathEventRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAttackPathEventRequest) SetEndTime(v int64) *ListAttackPathEventRequest {
	s.EndTime = &v
	return s
}

func (s *ListAttackPathEventRequest) SetLang(v string) *ListAttackPathEventRequest {
	s.Lang = &v
	return s
}

func (s *ListAttackPathEventRequest) SetPageSize(v int32) *ListAttackPathEventRequest {
	s.PageSize = &v
	return s
}

func (s *ListAttackPathEventRequest) SetPathNameDesc(v string) *ListAttackPathEventRequest {
	s.PathNameDesc = &v
	return s
}

func (s *ListAttackPathEventRequest) SetPathType(v string) *ListAttackPathEventRequest {
	s.PathType = &v
	return s
}

func (s *ListAttackPathEventRequest) SetRiskLevelList(v []*string) *ListAttackPathEventRequest {
	s.RiskLevelList = v
	return s
}

func (s *ListAttackPathEventRequest) SetStartTime(v int64) *ListAttackPathEventRequest {
	s.StartTime = &v
	return s
}

func (s *ListAttackPathEventRequest) Validate() error {
	if s.AttackPathAssetList != nil {
		for _, item := range s.AttackPathAssetList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAttackPathEventRequestAttackPathAssetList struct {
	// The subtype of the cloud service asset.
	//
	// > Call [ListSupportAttackPathAsset](~~ListSupportAttackPathAsset~~) to query the subtypes of cloud service assets.
	//
	// example:
	//
	// 2
	AssetSubType *int32 `json:"AssetSubType,omitempty" xml:"AssetSubType,omitempty"`
	// The type of the cloud service asset.
	//
	// > Call [ListSupportAttackPathAsset](~~ListSupportAttackPathAsset~~) to query the types of cloud service assets.
	//
	// example:
	//
	// 17
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The node type. Valid values:
	//
	// - **start**: start node.
	//
	// - **end**: end node.
	//
	// example:
	//
	// start
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The vendor of the cloud service asset.
	//
	// > Call [ListSupportAttackPathAsset](~~ListSupportAttackPathAsset~~) to query the vendors of cloud service assets.
	//
	// example:
	//
	// 0
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ListAttackPathEventRequestAttackPathAssetList) String() string {
	return dara.Prettify(s)
}

func (s ListAttackPathEventRequestAttackPathAssetList) GoString() string {
	return s.String()
}

func (s *ListAttackPathEventRequestAttackPathAssetList) GetAssetSubType() *int32 {
	return s.AssetSubType
}

func (s *ListAttackPathEventRequestAttackPathAssetList) GetAssetType() *int32 {
	return s.AssetType
}

func (s *ListAttackPathEventRequestAttackPathAssetList) GetNodeType() *string {
	return s.NodeType
}

func (s *ListAttackPathEventRequestAttackPathAssetList) GetVendor() *int32 {
	return s.Vendor
}

func (s *ListAttackPathEventRequestAttackPathAssetList) SetAssetSubType(v int32) *ListAttackPathEventRequestAttackPathAssetList {
	s.AssetSubType = &v
	return s
}

func (s *ListAttackPathEventRequestAttackPathAssetList) SetAssetType(v int32) *ListAttackPathEventRequestAttackPathAssetList {
	s.AssetType = &v
	return s
}

func (s *ListAttackPathEventRequestAttackPathAssetList) SetNodeType(v string) *ListAttackPathEventRequestAttackPathAssetList {
	s.NodeType = &v
	return s
}

func (s *ListAttackPathEventRequestAttackPathAssetList) SetVendor(v int32) *ListAttackPathEventRequestAttackPathAssetList {
	s.Vendor = &v
	return s
}

func (s *ListAttackPathEventRequestAttackPathAssetList) Validate() error {
	return dara.Validate(s)
}
