// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportUnReadCountShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequestShrink(v string) *GetReportUnReadCountShrinkRequest
	GetRequestShrink() *string
	SetTenantContextShrink(v string) *GetReportUnReadCountShrinkRequest
	GetTenantContextShrink() *string
}

type GetReportUnReadCountShrinkRequest struct {
	// example:
	//
	// null
	RequestShrink       *string `json:"Request,omitempty" xml:"Request,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetReportUnReadCountShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetReportUnReadCountShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetReportUnReadCountShrinkRequest) GetRequestShrink() *string {
	return s.RequestShrink
}

func (s *GetReportUnReadCountShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetReportUnReadCountShrinkRequest) SetRequestShrink(v string) *GetReportUnReadCountShrinkRequest {
	s.RequestShrink = &v
	return s
}

func (s *GetReportUnReadCountShrinkRequest) SetTenantContextShrink(v string) *GetReportUnReadCountShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetReportUnReadCountShrinkRequest) Validate() error {
	return dara.Validate(s)
}
