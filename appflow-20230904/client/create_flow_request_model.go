// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowDesc(v string) *CreateFlowRequest
	GetFlowDesc() *string
	SetFlowId(v string) *CreateFlowRequest
	GetFlowId() *string
	SetFlowName(v string) *CreateFlowRequest
	GetFlowName() *string
	SetFlowTemplate(v string) *CreateFlowRequest
	GetFlowTemplate() *string
	SetLaunchStatus(v bool) *CreateFlowRequest
	GetLaunchStatus() *bool
	SetParameters(v []*CreateFlowRequestParameters) *CreateFlowRequest
	GetParameters() []*CreateFlowRequestParameters
	SetTag(v []*CreateFlowRequestTag) *CreateFlowRequest
	GetTag() []*CreateFlowRequestTag
	SetTemplateId(v string) *CreateFlowRequest
	GetTemplateId() *string
}

type CreateFlowRequest struct {
	FlowDesc *string `json:"FlowDesc,omitempty" xml:"FlowDesc,omitempty"`
	// example:
	//
	// flow-647da8e366a74d1cab6e
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// This parameter is required.
	FlowName     *string                        `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	FlowTemplate *string                        `json:"FlowTemplate,omitempty" xml:"FlowTemplate,omitempty"`
	LaunchStatus *bool                          `json:"LaunchStatus,omitempty" xml:"LaunchStatus,omitempty"`
	Parameters   []*CreateFlowRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	Tag          []*CreateFlowRequestTag        `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// example:
	//
	// tl-715d93e708b546b7b464
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowRequest) GetFlowDesc() *string {
	return s.FlowDesc
}

func (s *CreateFlowRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *CreateFlowRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *CreateFlowRequest) GetFlowTemplate() *string {
	return s.FlowTemplate
}

func (s *CreateFlowRequest) GetLaunchStatus() *bool {
	return s.LaunchStatus
}

func (s *CreateFlowRequest) GetParameters() []*CreateFlowRequestParameters {
	return s.Parameters
}

func (s *CreateFlowRequest) GetTag() []*CreateFlowRequestTag {
	return s.Tag
}

func (s *CreateFlowRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateFlowRequest) SetFlowDesc(v string) *CreateFlowRequest {
	s.FlowDesc = &v
	return s
}

func (s *CreateFlowRequest) SetFlowId(v string) *CreateFlowRequest {
	s.FlowId = &v
	return s
}

func (s *CreateFlowRequest) SetFlowName(v string) *CreateFlowRequest {
	s.FlowName = &v
	return s
}

func (s *CreateFlowRequest) SetFlowTemplate(v string) *CreateFlowRequest {
	s.FlowTemplate = &v
	return s
}

func (s *CreateFlowRequest) SetLaunchStatus(v bool) *CreateFlowRequest {
	s.LaunchStatus = &v
	return s
}

func (s *CreateFlowRequest) SetParameters(v []*CreateFlowRequestParameters) *CreateFlowRequest {
	s.Parameters = v
	return s
}

func (s *CreateFlowRequest) SetTag(v []*CreateFlowRequestTag) *CreateFlowRequest {
	s.Tag = v
	return s
}

func (s *CreateFlowRequest) SetTemplateId(v string) *CreateFlowRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateFlowRequest) Validate() error {
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateFlowRequestParameters struct {
	// example:
	//
	// dingdingAuthId
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// example:
	//
	// uac-cdd8e1cfde534b4db482
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateFlowRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowRequestParameters) GoString() string {
	return s.String()
}

func (s *CreateFlowRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *CreateFlowRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *CreateFlowRequestParameters) SetParameterKey(v string) *CreateFlowRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *CreateFlowRequestParameters) SetParameterValue(v string) *CreateFlowRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *CreateFlowRequestParameters) Validate() error {
	return dara.Validate(s)
}

type CreateFlowRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateFlowRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowRequestTag) GoString() string {
	return s.String()
}

func (s *CreateFlowRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateFlowRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateFlowRequestTag) SetKey(v string) *CreateFlowRequestTag {
	s.Key = &v
	return s
}

func (s *CreateFlowRequestTag) SetValue(v string) *CreateFlowRequestTag {
	s.Value = &v
	return s
}

func (s *CreateFlowRequestTag) Validate() error {
	return dara.Validate(s)
}
