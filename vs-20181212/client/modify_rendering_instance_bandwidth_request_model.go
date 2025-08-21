// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRenderingInstanceBandwidthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxEgressBandwidth(v int32) *ModifyRenderingInstanceBandwidthRequest
	GetMaxEgressBandwidth() *int32
	SetMaxIngressBandwidth(v int32) *ModifyRenderingInstanceBandwidthRequest
	GetMaxIngressBandwidth() *int32
	SetRenderingInstanceId(v string) *ModifyRenderingInstanceBandwidthRequest
	GetRenderingInstanceId() *string
}

type ModifyRenderingInstanceBandwidthRequest struct {
	// example:
	//
	// 100
	MaxEgressBandwidth *int32 `json:"MaxEgressBandwidth,omitempty" xml:"MaxEgressBandwidth,omitempty"`
	// example:
	//
	// 100
	MaxIngressBandwidth *int32 `json:"MaxIngressBandwidth,omitempty" xml:"MaxIngressBandwidth,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s ModifyRenderingInstanceBandwidthRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRenderingInstanceBandwidthRequest) GoString() string {
	return s.String()
}

func (s *ModifyRenderingInstanceBandwidthRequest) GetMaxEgressBandwidth() *int32 {
	return s.MaxEgressBandwidth
}

func (s *ModifyRenderingInstanceBandwidthRequest) GetMaxIngressBandwidth() *int32 {
	return s.MaxIngressBandwidth
}

func (s *ModifyRenderingInstanceBandwidthRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ModifyRenderingInstanceBandwidthRequest) SetMaxEgressBandwidth(v int32) *ModifyRenderingInstanceBandwidthRequest {
	s.MaxEgressBandwidth = &v
	return s
}

func (s *ModifyRenderingInstanceBandwidthRequest) SetMaxIngressBandwidth(v int32) *ModifyRenderingInstanceBandwidthRequest {
	s.MaxIngressBandwidth = &v
	return s
}

func (s *ModifyRenderingInstanceBandwidthRequest) SetRenderingInstanceId(v string) *ModifyRenderingInstanceBandwidthRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *ModifyRenderingInstanceBandwidthRequest) Validate() error {
	return dara.Validate(s)
}
