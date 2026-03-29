// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataDiagnosesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListDataDiagnosesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListDataDiagnosesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataDiagnosesRequest
	GetPageSize() *int32
	SetTypes(v []*string) *ListDataDiagnosesRequest
	GetTypes() []*string
}

type ListDataDiagnosesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Types    []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
}

func (s ListDataDiagnosesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataDiagnosesRequest) GoString() string {
	return s.String()
}

func (s *ListDataDiagnosesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDataDiagnosesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataDiagnosesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataDiagnosesRequest) GetTypes() []*string {
	return s.Types
}

func (s *ListDataDiagnosesRequest) SetInstanceId(v string) *ListDataDiagnosesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDataDiagnosesRequest) SetPageNumber(v int32) *ListDataDiagnosesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataDiagnosesRequest) SetPageSize(v int32) *ListDataDiagnosesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataDiagnosesRequest) SetTypes(v []*string) *ListDataDiagnosesRequest {
	s.Types = v
	return s
}

func (s *ListDataDiagnosesRequest) Validate() error {
	return dara.Validate(s)
}
