// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentProfileTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAgentProfileTemplateResponseBody
	GetCode() *string
	SetData(v *GetAgentProfileTemplateResponseBodyData) *GetAgentProfileTemplateResponseBody
	GetData() *GetAgentProfileTemplateResponseBodyData
	SetHttpStatusCode(v int32) *GetAgentProfileTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAgentProfileTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgentProfileTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgentProfileTemplateResponseBody
	GetSuccess() *bool
}

type GetAgentProfileTemplateResponseBody struct {
	// example:
	//
	// OK
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAgentProfileTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgentProfileTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentProfileTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentProfileTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentProfileTemplateResponseBody) GetData() *GetAgentProfileTemplateResponseBodyData {
	return s.Data
}

func (s *GetAgentProfileTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAgentProfileTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentProfileTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentProfileTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgentProfileTemplateResponseBody) SetCode(v string) *GetAgentProfileTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentProfileTemplateResponseBody) SetData(v *GetAgentProfileTemplateResponseBodyData) *GetAgentProfileTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentProfileTemplateResponseBody) SetHttpStatusCode(v int32) *GetAgentProfileTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAgentProfileTemplateResponseBody) SetMessage(v string) *GetAgentProfileTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentProfileTemplateResponseBody) SetRequestId(v string) *GetAgentProfileTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentProfileTemplateResponseBody) SetSuccess(v bool) *GetAgentProfileTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentProfileTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentProfileTemplateResponseBodyData struct {
	// example:
	//
	// default-survey
	AgentProfileTemplateId *string `json:"AgentProfileTemplateId,omitempty" xml:"AgentProfileTemplateId,omitempty"`
	// example:
	//
	// 1720766491000
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PromptSchema *string `json:"PromptSchema,omitempty" xml:"PromptSchema,omitempty"`
	// example:
	//
	// 1720766491000
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetAgentProfileTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentProfileTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentProfileTemplateResponseBodyData) GetAgentProfileTemplateId() *string {
	return s.AgentProfileTemplateId
}

func (s *GetAgentProfileTemplateResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAgentProfileTemplateResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetAgentProfileTemplateResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetAgentProfileTemplateResponseBodyData) GetPromptSchema() *string {
	return s.PromptSchema
}

func (s *GetAgentProfileTemplateResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetAgentProfileTemplateResponseBodyData) SetAgentProfileTemplateId(v string) *GetAgentProfileTemplateResponseBodyData {
	s.AgentProfileTemplateId = &v
	return s
}

func (s *GetAgentProfileTemplateResponseBodyData) SetCreateTime(v string) *GetAgentProfileTemplateResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetAgentProfileTemplateResponseBodyData) SetDescription(v string) *GetAgentProfileTemplateResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetAgentProfileTemplateResponseBodyData) SetName(v string) *GetAgentProfileTemplateResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetAgentProfileTemplateResponseBodyData) SetPromptSchema(v string) *GetAgentProfileTemplateResponseBodyData {
	s.PromptSchema = &v
	return s
}

func (s *GetAgentProfileTemplateResponseBodyData) SetUpdateTime(v string) *GetAgentProfileTemplateResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetAgentProfileTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
