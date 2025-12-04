// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpdsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnablePage(v bool) *ListVpdsRequest
	GetEnablePage() *bool
	SetFilterErId(v string) *ListVpdsRequest
	GetFilterErId() *string
	SetForSelect(v bool) *ListVpdsRequest
	GetForSelect() *bool
	SetPageNumber(v int32) *ListVpdsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVpdsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListVpdsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListVpdsRequest
	GetResourceGroupId() *string
	SetStatus(v string) *ListVpdsRequest
	GetStatus() *string
	SetTag(v []*ListVpdsRequestTag) *ListVpdsRequest
	GetTag() []*ListVpdsRequestTag
	SetVpdId(v string) *ListVpdsRequest
	GetVpdId() *string
	SetVpdName(v string) *ListVpdsRequest
	GetVpdName() *string
	SetWithDependence(v bool) *ListVpdsRequest
	GetWithDependence() *bool
	SetWithoutVcc(v bool) *ListVpdsRequest
	GetWithoutVcc() *bool
}

type ListVpdsRequest struct {
	// Specifies whether to enable paged query.
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// Queries the network segments of Lingjun that are not bound to a specified Lingjun HUB.
	//
	// example:
	//
	// er-kkopgtne
	FilterErId *string `json:"FilterErId,omitempty" xml:"FilterErId,omitempty"`
	// If you select a drop-down list, only the basic information (including the instance ID and instance name) is returned. Possible values:
	//
	// 	- **true**: Select Query Use from the drop-down list.
	//
	// 	- **false**: Normal queries are used.
	//
	// example:
	//
	// true
	ForSelect *bool `json:"ForSelect,omitempty" xml:"ForSelect,omitempty"`
	// The page number of the page to return. Start value: 1 Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/94475.htm?spm=a2c4g.11186623.0.0.29e15d7akXhpuu).
	//
	// example:
	//
	// rg-aeky5f3qx6ceapq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the CLB instance. Valid values:
	//
	// 	- **Available**: Normal.
	//
	// 	- **Not Available**: Not available.
	//
	// 	- **Executing**: The task is being executed.
	//
	// 	- **Deleting**: The account is being deleted
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags.
	Tag []*ListVpdsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPD instance.
	//
	// example:
	//
	// vpd-fuliephf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The name of the VPD instance.
	//
	// example:
	//
	// vpd-1
	VpdName *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
	// Specifies whether to include the dependent resource information. We recommend that you do not query the dependent resource information when you query by page. You can query the dependent resource information separately when you delete it. Possible values:
	//
	// 	- **true**: with dependency information.
	//
	// 	- **false**: does not include dependency information.
	//
	// example:
	//
	// false
	WithDependence *bool `json:"WithDependence,omitempty" xml:"WithDependence,omitempty"`
	// Queries the information about a Lingjun CIDR block that is not bound to a Lingjun connection. Possible values:
	//
	// 	- **true**: filters out VPDs that have been bound to VCC
	//
	// 	- **false**: does not filter VPD that has been bound to VCC
	//
	// example:
	//
	// true
	WithoutVcc *bool `json:"WithoutVcc,omitempty" xml:"WithoutVcc,omitempty"`
}

func (s ListVpdsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpdsRequest) GoString() string {
	return s.String()
}

func (s *ListVpdsRequest) GetEnablePage() *bool {
	return s.EnablePage
}

func (s *ListVpdsRequest) GetFilterErId() *string {
	return s.FilterErId
}

func (s *ListVpdsRequest) GetForSelect() *bool {
	return s.ForSelect
}

func (s *ListVpdsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVpdsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVpdsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpdsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVpdsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListVpdsRequest) GetTag() []*ListVpdsRequestTag {
	return s.Tag
}

func (s *ListVpdsRequest) GetVpdId() *string {
	return s.VpdId
}

func (s *ListVpdsRequest) GetVpdName() *string {
	return s.VpdName
}

func (s *ListVpdsRequest) GetWithDependence() *bool {
	return s.WithDependence
}

func (s *ListVpdsRequest) GetWithoutVcc() *bool {
	return s.WithoutVcc
}

func (s *ListVpdsRequest) SetEnablePage(v bool) *ListVpdsRequest {
	s.EnablePage = &v
	return s
}

func (s *ListVpdsRequest) SetFilterErId(v string) *ListVpdsRequest {
	s.FilterErId = &v
	return s
}

func (s *ListVpdsRequest) SetForSelect(v bool) *ListVpdsRequest {
	s.ForSelect = &v
	return s
}

func (s *ListVpdsRequest) SetPageNumber(v int32) *ListVpdsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVpdsRequest) SetPageSize(v int32) *ListVpdsRequest {
	s.PageSize = &v
	return s
}

func (s *ListVpdsRequest) SetRegionId(v string) *ListVpdsRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpdsRequest) SetResourceGroupId(v string) *ListVpdsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpdsRequest) SetStatus(v string) *ListVpdsRequest {
	s.Status = &v
	return s
}

func (s *ListVpdsRequest) SetTag(v []*ListVpdsRequestTag) *ListVpdsRequest {
	s.Tag = v
	return s
}

func (s *ListVpdsRequest) SetVpdId(v string) *ListVpdsRequest {
	s.VpdId = &v
	return s
}

func (s *ListVpdsRequest) SetVpdName(v string) *ListVpdsRequest {
	s.VpdName = &v
	return s
}

func (s *ListVpdsRequest) SetWithDependence(v bool) *ListVpdsRequest {
	s.WithDependence = &v
	return s
}

func (s *ListVpdsRequest) SetWithoutVcc(v bool) *ListVpdsRequest {
	s.WithoutVcc = &v
	return s
}

func (s *ListVpdsRequest) Validate() error {
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

type ListVpdsRequestTag struct {
	// The tag key of the VPN attachment.
	//
	// You cannot specify an empty string as a tag key. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// tag-vpd-region
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the VPN connection.
	//
	// The tag value can be empty or a string of up to 128 characters. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each key-value pair must be unique. You can specify values for at most 20 tag keys in each call.
	//
	// example:
	//
	// wulanchabu
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVpdsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListVpdsRequestTag) GoString() string {
	return s.String()
}

func (s *ListVpdsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListVpdsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListVpdsRequestTag) SetKey(v string) *ListVpdsRequestTag {
	s.Key = &v
	return s
}

func (s *ListVpdsRequestTag) SetValue(v string) *ListVpdsRequestTag {
	s.Value = &v
	return s
}

func (s *ListVpdsRequestTag) Validate() error {
	return dara.Validate(s)
}
