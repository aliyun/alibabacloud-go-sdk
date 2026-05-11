// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdListJsonString(v string) *ListInstancesRequest
	GetInstanceIdListJsonString() *string
	SetName(v string) *ListInstancesRequest
	GetName() *string
	SetNluServiceTypeListJsonString(v string) *ListInstancesRequest
	GetNluServiceTypeListJsonString() *string
	SetNumber(v string) *ListInstancesRequest
	GetNumber() *string
	SetPageNumber(v int32) *ListInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstancesRequest
	GetPageSize() *int32
	SetStatus(v string) *ListInstancesRequest
	GetStatus() *string
	SetUnionInstanceId(v string) *ListInstancesRequest
	GetUnionInstanceId() *string
	SetUnionSource(v string) *ListInstancesRequest
	GetUnionSource() *string
}

type ListInstancesRequest struct {
	InstanceIdListJsonString *string `json:"InstanceIdListJsonString,omitempty" xml:"InstanceIdListJsonString,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// [Managed]
	NluServiceTypeListJsonString *string `json:"NluServiceTypeListJsonString,omitempty" xml:"NluServiceTypeListJsonString,omitempty"`
	Number                       *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UnionInstanceId *string `json:"UnionInstanceId,omitempty" xml:"UnionInstanceId,omitempty"`
	UnionSource     *string `json:"UnionSource,omitempty" xml:"UnionSource,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) GetInstanceIdListJsonString() *string {
	return s.InstanceIdListJsonString
}

func (s *ListInstancesRequest) GetName() *string {
	return s.Name
}

func (s *ListInstancesRequest) GetNluServiceTypeListJsonString() *string {
	return s.NluServiceTypeListJsonString
}

func (s *ListInstancesRequest) GetNumber() *string {
	return s.Number
}

func (s *ListInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesRequest) GetUnionInstanceId() *string {
	return s.UnionInstanceId
}

func (s *ListInstancesRequest) GetUnionSource() *string {
	return s.UnionSource
}

func (s *ListInstancesRequest) SetInstanceIdListJsonString(v string) *ListInstancesRequest {
	s.InstanceIdListJsonString = &v
	return s
}

func (s *ListInstancesRequest) SetName(v string) *ListInstancesRequest {
	s.Name = &v
	return s
}

func (s *ListInstancesRequest) SetNluServiceTypeListJsonString(v string) *ListInstancesRequest {
	s.NluServiceTypeListJsonString = &v
	return s
}

func (s *ListInstancesRequest) SetNumber(v string) *ListInstancesRequest {
	s.Number = &v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int32) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int32) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetStatus(v string) *ListInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListInstancesRequest) SetUnionInstanceId(v string) *ListInstancesRequest {
	s.UnionInstanceId = &v
	return s
}

func (s *ListInstancesRequest) SetUnionSource(v string) *ListInstancesRequest {
	s.UnionSource = &v
	return s
}

func (s *ListInstancesRequest) Validate() error {
	return dara.Validate(s)
}
