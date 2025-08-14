// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveCostCenterShareRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateShareRuleList(v []*SaveCostCenterShareRuleRequestCreateShareRuleList) *SaveCostCenterShareRuleRequest
	GetCreateShareRuleList() []*SaveCostCenterShareRuleRequestCreateShareRuleList
	SetModifyShareRuleList(v []*SaveCostCenterShareRuleRequestModifyShareRuleList) *SaveCostCenterShareRuleRequest
	GetModifyShareRuleList() []*SaveCostCenterShareRuleRequestModifyShareRuleList
	SetNbid(v string) *SaveCostCenterShareRuleRequest
	GetNbid() *string
	SetOwnerAccountId(v int64) *SaveCostCenterShareRuleRequest
	GetOwnerAccountId() *int64
	SetRemoveShareRuleList(v []*int64) *SaveCostCenterShareRuleRequest
	GetRemoveShareRuleList() []*int64
}

type SaveCostCenterShareRuleRequest struct {
	CreateShareRuleList []*SaveCostCenterShareRuleRequestCreateShareRuleList `json:"CreateShareRuleList,omitempty" xml:"CreateShareRuleList,omitempty" type:"Repeated"`
	ModifyShareRuleList []*SaveCostCenterShareRuleRequestModifyShareRuleList `json:"ModifyShareRuleList,omitempty" xml:"ModifyShareRuleList,omitempty" type:"Repeated"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// 1977800748053695
	OwnerAccountId      *int64   `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
	RemoveShareRuleList []*int64 `json:"RemoveShareRuleList,omitempty" xml:"RemoveShareRuleList,omitempty" type:"Repeated"`
}

func (s SaveCostCenterShareRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveCostCenterShareRuleRequest) GoString() string {
	return s.String()
}

func (s *SaveCostCenterShareRuleRequest) GetCreateShareRuleList() []*SaveCostCenterShareRuleRequestCreateShareRuleList {
	return s.CreateShareRuleList
}

func (s *SaveCostCenterShareRuleRequest) GetModifyShareRuleList() []*SaveCostCenterShareRuleRequestModifyShareRuleList {
	return s.ModifyShareRuleList
}

func (s *SaveCostCenterShareRuleRequest) GetNbid() *string {
	return s.Nbid
}

func (s *SaveCostCenterShareRuleRequest) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *SaveCostCenterShareRuleRequest) GetRemoveShareRuleList() []*int64 {
	return s.RemoveShareRuleList
}

func (s *SaveCostCenterShareRuleRequest) SetCreateShareRuleList(v []*SaveCostCenterShareRuleRequestCreateShareRuleList) *SaveCostCenterShareRuleRequest {
	s.CreateShareRuleList = v
	return s
}

func (s *SaveCostCenterShareRuleRequest) SetModifyShareRuleList(v []*SaveCostCenterShareRuleRequestModifyShareRuleList) *SaveCostCenterShareRuleRequest {
	s.ModifyShareRuleList = v
	return s
}

func (s *SaveCostCenterShareRuleRequest) SetNbid(v string) *SaveCostCenterShareRuleRequest {
	s.Nbid = &v
	return s
}

func (s *SaveCostCenterShareRuleRequest) SetOwnerAccountId(v int64) *SaveCostCenterShareRuleRequest {
	s.OwnerAccountId = &v
	return s
}

func (s *SaveCostCenterShareRuleRequest) SetRemoveShareRuleList(v []*int64) *SaveCostCenterShareRuleRequest {
	s.RemoveShareRuleList = v
	return s
}

func (s *SaveCostCenterShareRuleRequest) Validate() error {
	return dara.Validate(s)
}

type SaveCostCenterShareRuleRequestCreateShareRuleList struct {
	FromCostCenterList []*int64   `json:"FromCostCenterList,omitempty" xml:"FromCostCenterList,omitempty" type:"Repeated"`
	ShareRatioList     []*float64 `json:"ShareRatioList,omitempty" xml:"ShareRatioList,omitempty" type:"Repeated"`
	// example:
	//
	// test
	ShareRuleName *string `json:"ShareRuleName,omitempty" xml:"ShareRuleName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RATIO
	ShareType        *string  `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	ToCostCenterList []*int64 `json:"ToCostCenterList,omitempty" xml:"ToCostCenterList,omitempty" type:"Repeated"`
}

func (s SaveCostCenterShareRuleRequestCreateShareRuleList) String() string {
	return dara.Prettify(s)
}

func (s SaveCostCenterShareRuleRequestCreateShareRuleList) GoString() string {
	return s.String()
}

func (s *SaveCostCenterShareRuleRequestCreateShareRuleList) GetFromCostCenterList() []*int64 {
	return s.FromCostCenterList
}

func (s *SaveCostCenterShareRuleRequestCreateShareRuleList) GetShareRatioList() []*float64 {
	return s.ShareRatioList
}

func (s *SaveCostCenterShareRuleRequestCreateShareRuleList) GetShareRuleName() *string {
	return s.ShareRuleName
}

func (s *SaveCostCenterShareRuleRequestCreateShareRuleList) GetShareType() *string {
	return s.ShareType
}

func (s *SaveCostCenterShareRuleRequestCreateShareRuleList) GetToCostCenterList() []*int64 {
	return s.ToCostCenterList
}

func (s *SaveCostCenterShareRuleRequestCreateShareRuleList) SetFromCostCenterList(v []*int64) *SaveCostCenterShareRuleRequestCreateShareRuleList {
	s.FromCostCenterList = v
	return s
}

func (s *SaveCostCenterShareRuleRequestCreateShareRuleList) SetShareRatioList(v []*float64) *SaveCostCenterShareRuleRequestCreateShareRuleList {
	s.ShareRatioList = v
	return s
}

func (s *SaveCostCenterShareRuleRequestCreateShareRuleList) SetShareRuleName(v string) *SaveCostCenterShareRuleRequestCreateShareRuleList {
	s.ShareRuleName = &v
	return s
}

func (s *SaveCostCenterShareRuleRequestCreateShareRuleList) SetShareType(v string) *SaveCostCenterShareRuleRequestCreateShareRuleList {
	s.ShareType = &v
	return s
}

func (s *SaveCostCenterShareRuleRequestCreateShareRuleList) SetToCostCenterList(v []*int64) *SaveCostCenterShareRuleRequestCreateShareRuleList {
	s.ToCostCenterList = v
	return s
}

func (s *SaveCostCenterShareRuleRequestCreateShareRuleList) Validate() error {
	return dara.Validate(s)
}

type SaveCostCenterShareRuleRequestModifyShareRuleList struct {
	FromCostCenterList []*int64   `json:"FromCostCenterList,omitempty" xml:"FromCostCenterList,omitempty" type:"Repeated"`
	ShareRatioList     []*float64 `json:"ShareRatioList,omitempty" xml:"ShareRatioList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1828
	ShareRuleId *int64 `json:"ShareRuleId,omitempty" xml:"ShareRuleId,omitempty"`
	// example:
	//
	// test
	ShareRuleName *string `json:"ShareRuleName,omitempty" xml:"ShareRuleName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CUSTOM
	ShareType        *string  `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	ToCostCenterList []*int64 `json:"ToCostCenterList,omitempty" xml:"ToCostCenterList,omitempty" type:"Repeated"`
}

func (s SaveCostCenterShareRuleRequestModifyShareRuleList) String() string {
	return dara.Prettify(s)
}

func (s SaveCostCenterShareRuleRequestModifyShareRuleList) GoString() string {
	return s.String()
}

func (s *SaveCostCenterShareRuleRequestModifyShareRuleList) GetFromCostCenterList() []*int64 {
	return s.FromCostCenterList
}

func (s *SaveCostCenterShareRuleRequestModifyShareRuleList) GetShareRatioList() []*float64 {
	return s.ShareRatioList
}

func (s *SaveCostCenterShareRuleRequestModifyShareRuleList) GetShareRuleId() *int64 {
	return s.ShareRuleId
}

func (s *SaveCostCenterShareRuleRequestModifyShareRuleList) GetShareRuleName() *string {
	return s.ShareRuleName
}

func (s *SaveCostCenterShareRuleRequestModifyShareRuleList) GetShareType() *string {
	return s.ShareType
}

func (s *SaveCostCenterShareRuleRequestModifyShareRuleList) GetToCostCenterList() []*int64 {
	return s.ToCostCenterList
}

func (s *SaveCostCenterShareRuleRequestModifyShareRuleList) SetFromCostCenterList(v []*int64) *SaveCostCenterShareRuleRequestModifyShareRuleList {
	s.FromCostCenterList = v
	return s
}

func (s *SaveCostCenterShareRuleRequestModifyShareRuleList) SetShareRatioList(v []*float64) *SaveCostCenterShareRuleRequestModifyShareRuleList {
	s.ShareRatioList = v
	return s
}

func (s *SaveCostCenterShareRuleRequestModifyShareRuleList) SetShareRuleId(v int64) *SaveCostCenterShareRuleRequestModifyShareRuleList {
	s.ShareRuleId = &v
	return s
}

func (s *SaveCostCenterShareRuleRequestModifyShareRuleList) SetShareRuleName(v string) *SaveCostCenterShareRuleRequestModifyShareRuleList {
	s.ShareRuleName = &v
	return s
}

func (s *SaveCostCenterShareRuleRequestModifyShareRuleList) SetShareType(v string) *SaveCostCenterShareRuleRequestModifyShareRuleList {
	s.ShareType = &v
	return s
}

func (s *SaveCostCenterShareRuleRequestModifyShareRuleList) SetToCostCenterList(v []*int64) *SaveCostCenterShareRuleRequestModifyShareRuleList {
	s.ToCostCenterList = v
	return s
}

func (s *SaveCostCenterShareRuleRequestModifyShareRuleList) Validate() error {
	return dara.Validate(s)
}
