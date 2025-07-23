// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAicsOpenApiInvokeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *AicsOpenApiInvokeShrinkRequest
	GetNodeId() *string
	SetParamShrink(v string) *AicsOpenApiInvokeShrinkRequest
	GetParamShrink() *string
	SetServiceId(v string) *AicsOpenApiInvokeShrinkRequest
	GetServiceId() *string
	SetType(v string) *AicsOpenApiInvokeShrinkRequest
	GetType() *string
}

type AicsOpenApiInvokeShrinkRequest struct {
	// example:
	//
	// 119397
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// {"a":1}
	ParamShrink *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ae5f9884c9914ed7af72b07e6c1616f9
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// EXPERIMENT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AicsOpenApiInvokeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AicsOpenApiInvokeShrinkRequest) GoString() string {
	return s.String()
}

func (s *AicsOpenApiInvokeShrinkRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *AicsOpenApiInvokeShrinkRequest) GetParamShrink() *string {
	return s.ParamShrink
}

func (s *AicsOpenApiInvokeShrinkRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *AicsOpenApiInvokeShrinkRequest) GetType() *string {
	return s.Type
}

func (s *AicsOpenApiInvokeShrinkRequest) SetNodeId(v string) *AicsOpenApiInvokeShrinkRequest {
	s.NodeId = &v
	return s
}

func (s *AicsOpenApiInvokeShrinkRequest) SetParamShrink(v string) *AicsOpenApiInvokeShrinkRequest {
	s.ParamShrink = &v
	return s
}

func (s *AicsOpenApiInvokeShrinkRequest) SetServiceId(v string) *AicsOpenApiInvokeShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *AicsOpenApiInvokeShrinkRequest) SetType(v string) *AicsOpenApiInvokeShrinkRequest {
	s.Type = &v
	return s
}

func (s *AicsOpenApiInvokeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
