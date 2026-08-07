// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScriptProfileTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetScriptProfileTemplateRequest
	GetInstanceId() *string
	SetTemplateId(v string) *GetScriptProfileTemplateRequest
	GetTemplateId() *string
}

type GetScriptProfileTemplateRequest struct {
	// 实例ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 模板ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b59
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetScriptProfileTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScriptProfileTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetScriptProfileTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetScriptProfileTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetScriptProfileTemplateRequest) SetInstanceId(v string) *GetScriptProfileTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *GetScriptProfileTemplateRequest) SetTemplateId(v string) *GetScriptProfileTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetScriptProfileTemplateRequest) Validate() error {
	return dara.Validate(s)
}
