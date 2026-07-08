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
	// The description of the flow.
	//
	// example:
	//
	// 在钉钉中使用OpenClaw(MoltBot、MoltBot)
	FlowDesc *string `json:"FlowDesc,omitempty" xml:"FlowDesc,omitempty"`
	// The ID of the flow. This parameter is required when you update a flow or create a new flow version.
	//
	// example:
	//
	// flow-647da8e366a74d1cab6e
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The name of the flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 企业微信自建应用大模型自动回复
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The content of the template.
	//
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
	// The publication status of the flow: True for published, False for unpublished.
	//
	// example:
	//
	// true
	LaunchStatus *bool `json:"LaunchStatus,omitempty" xml:"LaunchStatus,omitempty"`
	// The parameters for the template.
	//
	// You can specify up to 200 parameters.
	//
	// > This parameter is optional. If you use this parameter, you must specify both ParameterKey and ParameterValue for each entry.
	Parameters []*CreateFlowRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The object tags to which the rule applies. You can specify multiple tags.
	Tag []*CreateFlowRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the template. Specify this parameter when you create a flow from a template in the Template Center.
	//
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
	// The name of a parameter defined in the template. If no parameter name or value is specified, ROS uses the default value defined in the template.
	//
	// The maximum value of N is 200.<br>
	//
	// The name must be 1 to 128 characters in length, cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// > Parameters is optional. If you specify Parameters, you must specify both Parameters.N.ParameterKey and Parameters.N.ParameterValue.
	//
	// example:
	//
	// dingdingAuthId
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value for the parameter.
	//
	// example:
	//
	// uac-xxxxxxx
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
	// The tag key. You can filter the cluster list by tag. You can specify up to 20 tag pairs. The number N in each tag pair must be unique and a consecutive integer starting from 1. The value corresponding to `Tag.N.Key` is `Tag.N.Value`.
	//
	// > The tag key can be up to 64 characters long and cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	//
	// example:
	//
	// CreateBy
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The authentication content.
	//
	// example:
	//
	// zhangsan
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
