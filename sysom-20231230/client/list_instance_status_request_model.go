// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int64) *ListInstanceStatusRequest
	GetCurrent() *int64
	SetInstance(v string) *ListInstanceStatusRequest
	GetInstance() *string
	SetPageSize(v int64) *ListInstanceStatusRequest
	GetPageSize() *int64
	SetRegion(v string) *ListInstanceStatusRequest
	GetRegion() *string
	SetStatus(v string) *ListInstanceStatusRequest
	GetStatus() *string
}

type ListInstanceStatusRequest struct {
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// example:
	//
	// i-wz9b9vucz1iubsz355rh
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// cn-shenzhen
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListInstanceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceStatusRequest) GetCurrent() *int64 {
	return s.Current
}

func (s *ListInstanceStatusRequest) GetInstance() *string {
	return s.Instance
}

func (s *ListInstanceStatusRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListInstanceStatusRequest) GetRegion() *string {
	return s.Region
}

func (s *ListInstanceStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *ListInstanceStatusRequest) SetCurrent(v int64) *ListInstanceStatusRequest {
	s.Current = &v
	return s
}

func (s *ListInstanceStatusRequest) SetInstance(v string) *ListInstanceStatusRequest {
	s.Instance = &v
	return s
}

func (s *ListInstanceStatusRequest) SetPageSize(v int64) *ListInstanceStatusRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceStatusRequest) SetRegion(v string) *ListInstanceStatusRequest {
	s.Region = &v
	return s
}

func (s *ListInstanceStatusRequest) SetStatus(v string) *ListInstanceStatusRequest {
	s.Status = &v
	return s
}

func (s *ListInstanceStatusRequest) Validate() error {
	return dara.Validate(s)
}
