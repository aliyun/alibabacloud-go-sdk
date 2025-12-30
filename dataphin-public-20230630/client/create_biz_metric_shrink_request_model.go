// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBizMetricShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateBizMetricCommandShrink(v string) *CreateBizMetricShrinkRequest
	GetCreateBizMetricCommandShrink() *string
	SetOpTenantId(v int64) *CreateBizMetricShrinkRequest
	GetOpTenantId() *int64
}

type CreateBizMetricShrinkRequest struct {
	// This parameter is required.
	CreateBizMetricCommandShrink *string `json:"CreateBizMetricCommand,omitempty" xml:"CreateBizMetricCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateBizMetricShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBizMetricShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateBizMetricShrinkRequest) GetCreateBizMetricCommandShrink() *string {
	return s.CreateBizMetricCommandShrink
}

func (s *CreateBizMetricShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateBizMetricShrinkRequest) SetCreateBizMetricCommandShrink(v string) *CreateBizMetricShrinkRequest {
	s.CreateBizMetricCommandShrink = &v
	return s
}

func (s *CreateBizMetricShrinkRequest) SetOpTenantId(v int64) *CreateBizMetricShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateBizMetricShrinkRequest) Validate() error {
	return dara.Validate(s)
}
