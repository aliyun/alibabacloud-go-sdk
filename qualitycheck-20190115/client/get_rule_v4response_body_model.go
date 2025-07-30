// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuleV4ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRuleV4ResponseBody
	GetCode() *string
	SetData(v *RulesInfo) *GetRuleV4ResponseBody
	GetData() *RulesInfo
	SetHttpStatusCode(v int32) *GetRuleV4ResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetRuleV4ResponseBody
	GetMessage() *string
	SetMessages(v []*string) *GetRuleV4ResponseBody
	GetMessages() []*string
	SetRequestId(v string) *GetRuleV4ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRuleV4ResponseBody
	GetSuccess() *bool
}

type GetRuleV4ResponseBody struct {
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
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRuleV4ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRuleV4ResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuleV4ResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRuleV4ResponseBody) GetData() *RulesInfo {
	return s.Data
}

func (s *GetRuleV4ResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetRuleV4ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRuleV4ResponseBody) GetMessages() []*string {
	return s.Messages
}

func (s *GetRuleV4ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRuleV4ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRuleV4ResponseBody) SetCode(v string) *GetRuleV4ResponseBody {
	s.Code = &v
	return s
}

func (s *GetRuleV4ResponseBody) SetData(v *RulesInfo) *GetRuleV4ResponseBody {
	s.Data = v
	return s
}

func (s *GetRuleV4ResponseBody) SetHttpStatusCode(v int32) *GetRuleV4ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRuleV4ResponseBody) SetMessage(v string) *GetRuleV4ResponseBody {
	s.Message = &v
	return s
}

func (s *GetRuleV4ResponseBody) SetMessages(v []*string) *GetRuleV4ResponseBody {
	s.Messages = v
	return s
}

func (s *GetRuleV4ResponseBody) SetRequestId(v string) *GetRuleV4ResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRuleV4ResponseBody) SetSuccess(v bool) *GetRuleV4ResponseBody {
	s.Success = &v
	return s
}

func (s *GetRuleV4ResponseBody) Validate() error {
	return dara.Validate(s)
}
