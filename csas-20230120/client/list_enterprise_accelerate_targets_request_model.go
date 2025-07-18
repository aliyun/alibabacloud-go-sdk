// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterpriseAccelerateTargetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListEnterpriseAccelerateTargetsRequest
	GetCurrentPage() *int64
	SetEapId(v string) *ListEnterpriseAccelerateTargetsRequest
	GetEapId() *string
	SetPageSize(v int64) *ListEnterpriseAccelerateTargetsRequest
	GetPageSize() *int64
	SetTarget(v string) *ListEnterpriseAccelerateTargetsRequest
	GetTarget() *string
}

type ListEnterpriseAccelerateTargetsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eap-424ba3f47660425c
	EapId *string `json:"EapId,omitempty" xml:"EapId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// googleapis.com
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ListEnterpriseAccelerateTargetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseAccelerateTargetsRequest) GoString() string {
	return s.String()
}

func (s *ListEnterpriseAccelerateTargetsRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListEnterpriseAccelerateTargetsRequest) GetEapId() *string {
	return s.EapId
}

func (s *ListEnterpriseAccelerateTargetsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListEnterpriseAccelerateTargetsRequest) GetTarget() *string {
	return s.Target
}

func (s *ListEnterpriseAccelerateTargetsRequest) SetCurrentPage(v int64) *ListEnterpriseAccelerateTargetsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListEnterpriseAccelerateTargetsRequest) SetEapId(v string) *ListEnterpriseAccelerateTargetsRequest {
	s.EapId = &v
	return s
}

func (s *ListEnterpriseAccelerateTargetsRequest) SetPageSize(v int64) *ListEnterpriseAccelerateTargetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEnterpriseAccelerateTargetsRequest) SetTarget(v string) *ListEnterpriseAccelerateTargetsRequest {
	s.Target = &v
	return s
}

func (s *ListEnterpriseAccelerateTargetsRequest) Validate() error {
	return dara.Validate(s)
}
