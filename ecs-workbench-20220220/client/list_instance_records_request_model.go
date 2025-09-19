// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListInstanceRecordsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListInstanceRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstanceRecordsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListInstanceRecordsRequest
	GetRegionId() *string
}

type ListInstanceRecordsRequest struct {
	// example:
	//
	// i-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstanceRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceRecordsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstanceRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstanceRecordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstanceRecordsRequest) SetInstanceId(v string) *ListInstanceRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceRecordsRequest) SetPageNumber(v int32) *ListInstanceRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceRecordsRequest) SetPageSize(v int32) *ListInstanceRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceRecordsRequest) SetRegionId(v string) *ListInstanceRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstanceRecordsRequest) Validate() error {
	return dara.Validate(s)
}
