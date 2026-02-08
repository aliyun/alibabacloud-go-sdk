// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *PublishScriptRequest
	GetDescription() *string
	SetInstanceId(v string) *PublishScriptRequest
	GetInstanceId() *string
	SetScriptId(v string) *PublishScriptRequest
	GetScriptId() *string
}

type PublishScriptRequest struct {
	// Description of this scenario publishing operation
	//
	// This parameter is required.
	//
	// example:
	//
	// 本次发布了xxxx内容
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// ID of the scenario to publish
	//
	// This parameter is required.
	//
	// example:
	//
	// 3677fe8b-276f-4541-babf-b9d3082a31ba
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s PublishScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishScriptRequest) GoString() string {
	return s.String()
}

func (s *PublishScriptRequest) GetDescription() *string {
	return s.Description
}

func (s *PublishScriptRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PublishScriptRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *PublishScriptRequest) SetDescription(v string) *PublishScriptRequest {
	s.Description = &v
	return s
}

func (s *PublishScriptRequest) SetInstanceId(v string) *PublishScriptRequest {
	s.InstanceId = &v
	return s
}

func (s *PublishScriptRequest) SetScriptId(v string) *PublishScriptRequest {
	s.ScriptId = &v
	return s
}

func (s *PublishScriptRequest) Validate() error {
	return dara.Validate(s)
}
