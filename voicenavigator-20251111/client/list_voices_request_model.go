// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListVoicesRequest
	GetInstanceId() *string
	SetNlsAccessType(v string) *ListVoicesRequest
	GetNlsAccessType() *string
	SetNlsEngine(v string) *ListVoicesRequest
	GetNlsEngine() *string
	SetPageNumber(v int32) *ListVoicesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVoicesRequest
	GetPageSize() *int32
}

type ListVoicesRequest struct {
	// example:
	//
	// e48e45dd-e47a-4744-a063-f08cbebb1c5a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// MANAGED
	NlsAccessType *string `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	// example:
	//
	// BAILIAN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListVoicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVoicesRequest) GoString() string {
	return s.String()
}

func (s *ListVoicesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVoicesRequest) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *ListVoicesRequest) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *ListVoicesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVoicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVoicesRequest) SetInstanceId(v string) *ListVoicesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListVoicesRequest) SetNlsAccessType(v string) *ListVoicesRequest {
	s.NlsAccessType = &v
	return s
}

func (s *ListVoicesRequest) SetNlsEngine(v string) *ListVoicesRequest {
	s.NlsEngine = &v
	return s
}

func (s *ListVoicesRequest) SetPageNumber(v int32) *ListVoicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVoicesRequest) SetPageSize(v int32) *ListVoicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListVoicesRequest) Validate() error {
	return dara.Validate(s)
}
