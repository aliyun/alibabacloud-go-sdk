// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBizMetricShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteBizMetricCommandShrink(v string) *DeleteBizMetricShrinkRequest
	GetDeleteBizMetricCommandShrink() *string
	SetOpTenantId(v int64) *DeleteBizMetricShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *DeleteBizMetricShrinkRequest
	GetOpUserId() *string
}

type DeleteBizMetricShrinkRequest struct {
	// The delete request.
	//
	// This parameter is required.
	DeleteBizMetricCommandShrink *string `json:"DeleteBizMetricCommand,omitempty" xml:"DeleteBizMetricCommand,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s DeleteBizMetricShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBizMetricShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteBizMetricShrinkRequest) GetDeleteBizMetricCommandShrink() *string {
	return s.DeleteBizMetricCommandShrink
}

func (s *DeleteBizMetricShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteBizMetricShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *DeleteBizMetricShrinkRequest) SetDeleteBizMetricCommandShrink(v string) *DeleteBizMetricShrinkRequest {
	s.DeleteBizMetricCommandShrink = &v
	return s
}

func (s *DeleteBizMetricShrinkRequest) SetOpTenantId(v int64) *DeleteBizMetricShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteBizMetricShrinkRequest) SetOpUserId(v string) *DeleteBizMetricShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *DeleteBizMetricShrinkRequest) Validate() error {
	return dara.Validate(s)
}
