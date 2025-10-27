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
	// List of cloud product assets in the attack path.
	AttackPathAssetList []*ListAttackPathEventRequestAttackPathAssetList `json:"AttackPathAssetList,omitempty" xml:"AttackPathAssetList,omitempty" type:"Repeated"`
	// Specifies from which page of the returned results the query results should be displayed. The default value is 1, indicating that the display starts from the first page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Timestamp of the end time. Unit: milliseconds.
	//
	// example:
	//
	// 1668064495000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Sets the language type for requests and received messages, with the default being **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of data entries displayed per page in a paginated query. The default value is **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Description of the path name.
	//
	// > You can call [ListAvailableAttackPath](~~ListAvailableAttackPath~~) to query the path name description.
	//
	// example:
	//
	// ECS Instance Can Obtain Long-term Access Credential by Enabling Console Logon for RAM User
	PathNameDesc *string `json:"PathNameDesc,omitempty" xml:"PathNameDesc,omitempty"`
	// Path type.
	//
	// > You can call [ListAvailableAttackPath](~~ListAvailableAttackPath~~) to query the path type.
	//
	// example:
	//
	// role_escalation
	PathType *string `json:"PathType,omitempty" xml:"PathType,omitempty"`
	// List of risk level information.
	RiskLevelList []*string `json:"RiskLevelList,omitempty" xml:"RiskLevelList,omitempty" type:"Repeated"`
	// Timestamp of the start time. Unit: milliseconds.
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
	// Subtype of the cloud product asset.
	//
	// > You can call [ListSupportAttackPathAsset](~~ListSupportAttackPathAsset~~) to query the subtype of the cloud product asset.
	//
	// example:
	//
	// 2
	AssetSubType *int32 `json:"AssetSubType,omitempty" xml:"AssetSubType,omitempty"`
	// Type of the cloud product asset.
	//
	// > You can call [ListSupportAttackPathAsset](~~ListSupportAttackPathAsset~~) to query the type of the cloud product asset.
	//
	// example:
	//
	// 17
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// Node type, values:
	//
	// - **start**: start point.
	//
	// - **end**: end point.
	//
	// example:
	//
	// start
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// Vendor of the cloud product asset.
	//
	// > You can call [ListSupportAttackPathAsset](~~ListSupportAttackPathAsset~~) to query the vendor of the cloud product asset.
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
