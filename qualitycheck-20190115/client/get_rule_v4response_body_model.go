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
	// Result code. A value of **200*	- indicates success. Other values indicate failure. Use this field to identify the cause of failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Rule information.
	Data *RulesInfo `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP status code returned by the request. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error details if the request failed. Returns **successful*	- if the request succeeded.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Error details if the request failed. Use this field when returning multiple messages.
	Messages []*string `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded. Set to true for success. Set to false or null for failure.
	//
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
