// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListErsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnablePage(v bool) *ListErsRequest
	GetEnablePage() *bool
	SetErId(v string) *ListErsRequest
	GetErId() *string
	SetErName(v string) *ListErsRequest
	GetErName() *string
	SetInstanceId(v string) *ListErsRequest
	GetInstanceId() *string
	SetInstanceType(v string) *ListErsRequest
	GetInstanceType() *string
	SetMasterZoneId(v string) *ListErsRequest
	GetMasterZoneId() *string
	SetPageNumber(v int32) *ListErsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListErsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListErsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListErsRequest
	GetResourceGroupId() *string
	SetTag(v []*ListErsRequestTag) *ListErsRequest
	GetTag() []*ListErsRequestTag
}

type ListErsRequest struct {
	// Specifies whether to enable paged query. Valid values:
	//
	// 	- true: enables paged query.
	//
	// 	- false: Paged query is disabled.
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// The ID of the Lingjun HUB instance.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Lingjun HUB name.
	//
	// example:
	//
	// er-heyuan-main
	ErName *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	// The ID of the network instance.
	//
	// example:
	//
	// vcc-cn-209300qha01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the attached network instance. Valid values:
	//
	// 	- **VPD**
	//
	// 	- **VCC**
	//
	// example:
	//
	// VCC
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The primary zone.
	//
	// example:
	//
	// cn-wulanchabu-b
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	// The page number to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-acfmwfm33rlt6zi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// List of tags.
	Tag []*ListErsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListErsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListErsRequest) GoString() string {
	return s.String()
}

func (s *ListErsRequest) GetEnablePage() *bool {
	return s.EnablePage
}

func (s *ListErsRequest) GetErId() *string {
	return s.ErId
}

func (s *ListErsRequest) GetErName() *string {
	return s.ErName
}

func (s *ListErsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListErsRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListErsRequest) GetMasterZoneId() *string {
	return s.MasterZoneId
}

func (s *ListErsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListErsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListErsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListErsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListErsRequest) GetTag() []*ListErsRequestTag {
	return s.Tag
}

func (s *ListErsRequest) SetEnablePage(v bool) *ListErsRequest {
	s.EnablePage = &v
	return s
}

func (s *ListErsRequest) SetErId(v string) *ListErsRequest {
	s.ErId = &v
	return s
}

func (s *ListErsRequest) SetErName(v string) *ListErsRequest {
	s.ErName = &v
	return s
}

func (s *ListErsRequest) SetInstanceId(v string) *ListErsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListErsRequest) SetInstanceType(v string) *ListErsRequest {
	s.InstanceType = &v
	return s
}

func (s *ListErsRequest) SetMasterZoneId(v string) *ListErsRequest {
	s.MasterZoneId = &v
	return s
}

func (s *ListErsRequest) SetPageNumber(v int32) *ListErsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListErsRequest) SetPageSize(v int32) *ListErsRequest {
	s.PageSize = &v
	return s
}

func (s *ListErsRequest) SetRegionId(v string) *ListErsRequest {
	s.RegionId = &v
	return s
}

func (s *ListErsRequest) SetResourceGroupId(v string) *ListErsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListErsRequest) SetTag(v []*ListErsRequestTag) *ListErsRequest {
	s.Tag = v
	return s
}

func (s *ListErsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListErsRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// rg-er
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// rg-xxxxx
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListErsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListErsRequestTag) GoString() string {
	return s.String()
}

func (s *ListErsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListErsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListErsRequestTag) SetKey(v string) *ListErsRequestTag {
	s.Key = &v
	return s
}

func (s *ListErsRequestTag) SetValue(v string) *ListErsRequestTag {
	s.Value = &v
	return s
}

func (s *ListErsRequestTag) Validate() error {
	return dara.Validate(s)
}
