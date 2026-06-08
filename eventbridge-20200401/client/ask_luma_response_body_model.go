// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAskLumaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AskLumaResponseBody
	GetCode() *string
	SetData(v *AskLumaResult) *AskLumaResponseBody
	GetData() *AskLumaResult
	SetMessage(v string) *AskLumaResponseBody
	GetMessage() *string
	SetRequestId(v string) *AskLumaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AskLumaResponseBody
	GetSuccess() *bool
}

type AskLumaResponseBody struct {
	// example:
	//
	// AgentNotFound
	Code *string        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AskLumaResult `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Agent with name \\"xxx\\" not found for account 1186xxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 97FB3BAE-XXXXXX-36435495B7EB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AskLumaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AskLumaResponseBody) GoString() string {
	return s.String()
}

func (s *AskLumaResponseBody) GetCode() *string {
	return s.Code
}

func (s *AskLumaResponseBody) GetData() *AskLumaResult {
	return s.Data
}

func (s *AskLumaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AskLumaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AskLumaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AskLumaResponseBody) SetCode(v string) *AskLumaResponseBody {
	s.Code = &v
	return s
}

func (s *AskLumaResponseBody) SetData(v *AskLumaResult) *AskLumaResponseBody {
	s.Data = v
	return s
}

func (s *AskLumaResponseBody) SetMessage(v string) *AskLumaResponseBody {
	s.Message = &v
	return s
}

func (s *AskLumaResponseBody) SetRequestId(v string) *AskLumaResponseBody {
	s.RequestId = &v
	return s
}

func (s *AskLumaResponseBody) SetSuccess(v bool) *AskLumaResponseBody {
	s.Success = &v
	return s
}

func (s *AskLumaResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
