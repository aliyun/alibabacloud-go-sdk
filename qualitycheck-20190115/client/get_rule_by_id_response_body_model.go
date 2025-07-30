// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuleByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRuleByIdResponseBody
	GetCode() *string
	SetData(v *RulesInfo) *GetRuleByIdResponseBody
	GetData() *RulesInfo
	SetHttpStatusCode(v int32) *GetRuleByIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetRuleByIdResponseBody
	GetMessage() *string
	SetMessages(v []*string) *GetRuleByIdResponseBody
	GetMessages() []*string
	SetRequestId(v string) *GetRuleByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRuleByIdResponseBody
	GetSuccess() *bool
}

type GetRuleByIdResponseBody struct {
	// example:
	//
	// 200
	Code *string    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *RulesInfo `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages []*string `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// example:
	//
	// 3CEA0495-341B-4482-9AD9-8191EF4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRuleByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRuleByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuleByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRuleByIdResponseBody) GetData() *RulesInfo {
	return s.Data
}

func (s *GetRuleByIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetRuleByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRuleByIdResponseBody) GetMessages() []*string {
	return s.Messages
}

func (s *GetRuleByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRuleByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRuleByIdResponseBody) SetCode(v string) *GetRuleByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetRuleByIdResponseBody) SetData(v *RulesInfo) *GetRuleByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetRuleByIdResponseBody) SetHttpStatusCode(v int32) *GetRuleByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRuleByIdResponseBody) SetMessage(v string) *GetRuleByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetRuleByIdResponseBody) SetMessages(v []*string) *GetRuleByIdResponseBody {
	s.Messages = v
	return s
}

func (s *GetRuleByIdResponseBody) SetRequestId(v string) *GetRuleByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRuleByIdResponseBody) SetSuccess(v bool) *GetRuleByIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetRuleByIdResponseBody) Validate() error {
	return dara.Validate(s)
}
