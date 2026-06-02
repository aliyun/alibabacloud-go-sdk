// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLaunchFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowDesc(v string) *LaunchFlowRequest
	GetFlowDesc() *string
	SetFlowId(v string) *LaunchFlowRequest
	GetFlowId() *string
	SetFlowName(v string) *LaunchFlowRequest
	GetFlowName() *string
	SetFlowTemplate(v string) *LaunchFlowRequest
	GetFlowTemplate() *string
	SetFlowVersion(v int32) *LaunchFlowRequest
	GetFlowVersion() *int32
}

type LaunchFlowRequest struct {
	FlowDesc *string `json:"FlowDesc,omitempty" xml:"FlowDesc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// flow-6a3acc07d51541b0b836
	FlowId   *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// example:
	//
	// {
	//
	//   "FormatVersion": "appflow-2025-07-01",
	//
	//   "Nodes": [
	//
	//      ]
	//
	// }
	FlowTemplate *string `json:"FlowTemplate,omitempty" xml:"FlowTemplate,omitempty"`
	// example:
	//
	// 1
	FlowVersion *int32 `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
}

func (s LaunchFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s LaunchFlowRequest) GoString() string {
	return s.String()
}

func (s *LaunchFlowRequest) GetFlowDesc() *string {
	return s.FlowDesc
}

func (s *LaunchFlowRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *LaunchFlowRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *LaunchFlowRequest) GetFlowTemplate() *string {
	return s.FlowTemplate
}

func (s *LaunchFlowRequest) GetFlowVersion() *int32 {
	return s.FlowVersion
}

func (s *LaunchFlowRequest) SetFlowDesc(v string) *LaunchFlowRequest {
	s.FlowDesc = &v
	return s
}

func (s *LaunchFlowRequest) SetFlowId(v string) *LaunchFlowRequest {
	s.FlowId = &v
	return s
}

func (s *LaunchFlowRequest) SetFlowName(v string) *LaunchFlowRequest {
	s.FlowName = &v
	return s
}

func (s *LaunchFlowRequest) SetFlowTemplate(v string) *LaunchFlowRequest {
	s.FlowTemplate = &v
	return s
}

func (s *LaunchFlowRequest) SetFlowVersion(v int32) *LaunchFlowRequest {
	s.FlowVersion = &v
	return s
}

func (s *LaunchFlowRequest) Validate() error {
	return dara.Validate(s)
}
