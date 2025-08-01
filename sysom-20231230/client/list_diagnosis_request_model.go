// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnosisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int64) *ListDiagnosisRequest
	GetCurrent() *int64
	SetPageSize(v int64) *ListDiagnosisRequest
	GetPageSize() *int64
	SetParams(v string) *ListDiagnosisRequest
	GetParams() *string
	SetServiceName(v string) *ListDiagnosisRequest
	GetServiceName() *string
	SetStatus(v string) *ListDiagnosisRequest
	GetStatus() *string
}

type ListDiagnosisRequest struct {
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// example:
	//
	// 10
	PageSize    *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Params      *string `json:"params,omitempty" xml:"params,omitempty"`
	ServiceName *string `json:"service_name,omitempty" xml:"service_name,omitempty"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListDiagnosisRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *ListDiagnosisRequest) GetCurrent() *int64 {
	return s.Current
}

func (s *ListDiagnosisRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDiagnosisRequest) GetParams() *string {
	return s.Params
}

func (s *ListDiagnosisRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListDiagnosisRequest) GetStatus() *string {
	return s.Status
}

func (s *ListDiagnosisRequest) SetCurrent(v int64) *ListDiagnosisRequest {
	s.Current = &v
	return s
}

func (s *ListDiagnosisRequest) SetPageSize(v int64) *ListDiagnosisRequest {
	s.PageSize = &v
	return s
}

func (s *ListDiagnosisRequest) SetParams(v string) *ListDiagnosisRequest {
	s.Params = &v
	return s
}

func (s *ListDiagnosisRequest) SetServiceName(v string) *ListDiagnosisRequest {
	s.ServiceName = &v
	return s
}

func (s *ListDiagnosisRequest) SetStatus(v string) *ListDiagnosisRequest {
	s.Status = &v
	return s
}

func (s *ListDiagnosisRequest) Validate() error {
	return dara.Validate(s)
}
