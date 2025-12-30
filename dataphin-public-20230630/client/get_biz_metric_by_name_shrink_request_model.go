// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBizMetricByNameShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizMetricByNameQueryShrink(v string) *GetBizMetricByNameShrinkRequest
	GetBizMetricByNameQueryShrink() *string
	SetOpTenantId(v int64) *GetBizMetricByNameShrinkRequest
	GetOpTenantId() *int64
}

type GetBizMetricByNameShrinkRequest struct {
	// This parameter is required.
	BizMetricByNameQueryShrink *string `json:"BizMetricByNameQuery,omitempty" xml:"BizMetricByNameQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetBizMetricByNameShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBizMetricByNameShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetBizMetricByNameShrinkRequest) GetBizMetricByNameQueryShrink() *string {
	return s.BizMetricByNameQueryShrink
}

func (s *GetBizMetricByNameShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetBizMetricByNameShrinkRequest) SetBizMetricByNameQueryShrink(v string) *GetBizMetricByNameShrinkRequest {
	s.BizMetricByNameQueryShrink = &v
	return s
}

func (s *GetBizMetricByNameShrinkRequest) SetOpTenantId(v int64) *GetBizMetricByNameShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetBizMetricByNameShrinkRequest) Validate() error {
	return dara.Validate(s)
}
