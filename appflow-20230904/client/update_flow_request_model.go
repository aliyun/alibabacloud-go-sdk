// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *UpdateFlowRequest
	GetEnabled() *bool
	SetFlowDesc(v string) *UpdateFlowRequest
	GetFlowDesc() *string
	SetFlowId(v string) *UpdateFlowRequest
	GetFlowId() *string
	SetFlowName(v string) *UpdateFlowRequest
	GetFlowName() *string
	SetFlowTemplate(v string) *UpdateFlowRequest
	GetFlowTemplate() *string
	SetFlowVersion(v string) *UpdateFlowRequest
	GetFlowVersion() *string
}

type UpdateFlowRequest struct {
	// example:
	//
	// true
	Enabled  *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	FlowDesc *string `json:"FlowDesc,omitempty" xml:"FlowDesc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// flow-15550241a1e942c29835
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
	// 9
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
}

func (s UpdateFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlowRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateFlowRequest) GetFlowDesc() *string {
	return s.FlowDesc
}

func (s *UpdateFlowRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *UpdateFlowRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *UpdateFlowRequest) GetFlowTemplate() *string {
	return s.FlowTemplate
}

func (s *UpdateFlowRequest) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *UpdateFlowRequest) SetEnabled(v bool) *UpdateFlowRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateFlowRequest) SetFlowDesc(v string) *UpdateFlowRequest {
	s.FlowDesc = &v
	return s
}

func (s *UpdateFlowRequest) SetFlowId(v string) *UpdateFlowRequest {
	s.FlowId = &v
	return s
}

func (s *UpdateFlowRequest) SetFlowName(v string) *UpdateFlowRequest {
	s.FlowName = &v
	return s
}

func (s *UpdateFlowRequest) SetFlowTemplate(v string) *UpdateFlowRequest {
	s.FlowTemplate = &v
	return s
}

func (s *UpdateFlowRequest) SetFlowVersion(v string) *UpdateFlowRequest {
	s.FlowVersion = &v
	return s
}

func (s *UpdateFlowRequest) Validate() error {
	return dara.Validate(s)
}
