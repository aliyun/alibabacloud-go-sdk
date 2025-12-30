// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBizMetricShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateBizMetricShrinkRequest
	GetOpTenantId() *int64
	SetUpdateBizMetricCommandShrink(v string) *UpdateBizMetricShrinkRequest
	GetUpdateBizMetricCommandShrink() *string
}

type UpdateBizMetricShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateBizMetricCommandShrink *string `json:"UpdateBizMetricCommand,omitempty" xml:"UpdateBizMetricCommand,omitempty"`
}

func (s UpdateBizMetricShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizMetricShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateBizMetricShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateBizMetricShrinkRequest) GetUpdateBizMetricCommandShrink() *string {
	return s.UpdateBizMetricCommandShrink
}

func (s *UpdateBizMetricShrinkRequest) SetOpTenantId(v int64) *UpdateBizMetricShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateBizMetricShrinkRequest) SetUpdateBizMetricCommandShrink(v string) *UpdateBizMetricShrinkRequest {
	s.UpdateBizMetricCommandShrink = &v
	return s
}

func (s *UpdateBizMetricShrinkRequest) Validate() error {
	return dara.Validate(s)
}
