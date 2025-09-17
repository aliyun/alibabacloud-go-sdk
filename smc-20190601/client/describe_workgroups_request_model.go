// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWorkgroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DescribeWorkgroupsRequest
	GetName() *string
	SetOwnerId(v int64) *DescribeWorkgroupsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeWorkgroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeWorkgroupsRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeWorkgroupsRequest
	GetResourceOwnerAccount() *string
	SetStatus(v string) *DescribeWorkgroupsRequest
	GetStatus() *string
	SetTag(v []*DescribeWorkgroupsRequestTag) *DescribeWorkgroupsRequest
	GetTag() []*DescribeWorkgroupsRequestTag
	SetWorkgroupId(v []*string) *DescribeWorkgroupsRequest
	GetWorkgroupId() []*string
}

type DescribeWorkgroupsRequest struct {
	// The name of the workgroup.
	//
	// example:
	//
	// test
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The state of the workgroup. Valid values:
	//
	// 	- NotStarted
	//
	// 	- InProgress
	//
	// 	- Cutover
	//
	// 	- Completed
	//
	// example:
	//
	// InProgress
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the reserved instance. You can specify up to 20 tags. If you specify multiple tags, the tag keys cannot be duplicated.
	Tag []*DescribeWorkgroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The workgroup IDs. You can specify up to 50 workgroup IDs.
	WorkgroupId []*string `json:"WorkgroupId,omitempty" xml:"WorkgroupId,omitempty" type:"Repeated"`
}

func (s DescribeWorkgroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkgroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWorkgroupsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeWorkgroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeWorkgroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeWorkgroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWorkgroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeWorkgroupsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeWorkgroupsRequest) GetTag() []*DescribeWorkgroupsRequestTag {
	return s.Tag
}

func (s *DescribeWorkgroupsRequest) GetWorkgroupId() []*string {
	return s.WorkgroupId
}

func (s *DescribeWorkgroupsRequest) SetName(v string) *DescribeWorkgroupsRequest {
	s.Name = &v
	return s
}

func (s *DescribeWorkgroupsRequest) SetOwnerId(v int64) *DescribeWorkgroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeWorkgroupsRequest) SetPageNumber(v int32) *DescribeWorkgroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeWorkgroupsRequest) SetPageSize(v int32) *DescribeWorkgroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWorkgroupsRequest) SetResourceOwnerAccount(v string) *DescribeWorkgroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeWorkgroupsRequest) SetStatus(v string) *DescribeWorkgroupsRequest {
	s.Status = &v
	return s
}

func (s *DescribeWorkgroupsRequest) SetTag(v []*DescribeWorkgroupsRequestTag) *DescribeWorkgroupsRequest {
	s.Tag = v
	return s
}

func (s *DescribeWorkgroupsRequest) SetWorkgroupId(v []*string) *DescribeWorkgroupsRequest {
	s.WorkgroupId = v
	return s
}

func (s *DescribeWorkgroupsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeWorkgroupsRequestTag struct {
	// The tag key of the workgroup. You cannot specify an empty string as a tag key. The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the workgroup. The tag value can be up to 128 characters in length, cannot be an empty string, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeWorkgroupsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkgroupsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeWorkgroupsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeWorkgroupsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeWorkgroupsRequestTag) SetKey(v string) *DescribeWorkgroupsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeWorkgroupsRequestTag) SetValue(v string) *DescribeWorkgroupsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeWorkgroupsRequestTag) Validate() error {
	return dara.Validate(s)
}
