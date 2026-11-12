// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v string) *ListInstanceStatusRequest
	GetInstanceIds() *string
	SetPageNumber(v int32) *ListInstanceStatusRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstanceStatusRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListInstanceStatusRequest
	GetRegionId() *string
}

type ListInstanceStatusRequest struct {
	// The IDs of the simple application servers. The value can be a JSON array that consists of up to 100 simple application server IDs. Separate multiple server IDs with commas (,).
	//
	// example:
	//
	// ["2ad1ae67295445f598017499dc****", "2ad1ae67295445f598017123dc****"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application servers. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstanceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceStatusRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *ListInstanceStatusRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstanceStatusRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstanceStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstanceStatusRequest) SetInstanceIds(v string) *ListInstanceStatusRequest {
	s.InstanceIds = &v
	return s
}

func (s *ListInstanceStatusRequest) SetPageNumber(v int32) *ListInstanceStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceStatusRequest) SetPageSize(v int32) *ListInstanceStatusRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceStatusRequest) SetRegionId(v string) *ListInstanceStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstanceStatusRequest) Validate() error {
	return dara.Validate(s)
}
