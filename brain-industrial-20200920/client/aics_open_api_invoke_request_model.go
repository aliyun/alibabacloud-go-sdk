// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAicsOpenApiInvokeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *AicsOpenApiInvokeRequest
	GetJobId() *string
	SetNodeId(v string) *AicsOpenApiInvokeRequest
	GetNodeId() *string
	SetParam(v map[string]interface{}) *AicsOpenApiInvokeRequest
	GetParam() map[string]interface{}
	SetServiceId(v string) *AicsOpenApiInvokeRequest
	GetServiceId() *string
	SetType(v string) *AicsOpenApiInvokeRequest
	GetType() *string
}

type AicsOpenApiInvokeRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 119397
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// {"a":1}
	Param map[string]interface{} `json:"Param,omitempty" xml:"Param,omitempty"`
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

func (s AicsOpenApiInvokeRequest) String() string {
	return dara.Prettify(s)
}

func (s AicsOpenApiInvokeRequest) GoString() string {
	return s.String()
}

func (s *AicsOpenApiInvokeRequest) GetJobId() *string {
	return s.JobId
}

func (s *AicsOpenApiInvokeRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *AicsOpenApiInvokeRequest) GetParam() map[string]interface{} {
	return s.Param
}

func (s *AicsOpenApiInvokeRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *AicsOpenApiInvokeRequest) GetType() *string {
	return s.Type
}

func (s *AicsOpenApiInvokeRequest) SetJobId(v string) *AicsOpenApiInvokeRequest {
	s.JobId = &v
	return s
}

func (s *AicsOpenApiInvokeRequest) SetNodeId(v string) *AicsOpenApiInvokeRequest {
	s.NodeId = &v
	return s
}

func (s *AicsOpenApiInvokeRequest) SetParam(v map[string]interface{}) *AicsOpenApiInvokeRequest {
	s.Param = v
	return s
}

func (s *AicsOpenApiInvokeRequest) SetServiceId(v string) *AicsOpenApiInvokeRequest {
	s.ServiceId = &v
	return s
}

func (s *AicsOpenApiInvokeRequest) SetType(v string) *AicsOpenApiInvokeRequest {
	s.Type = &v
	return s
}

func (s *AicsOpenApiInvokeRequest) Validate() error {
	return dara.Validate(s)
}
