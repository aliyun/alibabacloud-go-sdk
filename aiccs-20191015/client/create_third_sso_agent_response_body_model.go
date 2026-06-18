// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateThirdSsoAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateThirdSsoAgentResponseBody
	GetCode() *string
	SetData(v int64) *CreateThirdSsoAgentResponseBody
	GetData() *int64
	SetHttpStatusCode(v int64) *CreateThirdSsoAgentResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *CreateThirdSsoAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateThirdSsoAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateThirdSsoAgentResponseBody
	GetSuccess() *bool
}

type CreateThirdSsoAgentResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the created agent.
	//
	// example:
	//
	// 123456
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateThirdSsoAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateThirdSsoAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateThirdSsoAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateThirdSsoAgentResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateThirdSsoAgentResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *CreateThirdSsoAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateThirdSsoAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateThirdSsoAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateThirdSsoAgentResponseBody) SetCode(v string) *CreateThirdSsoAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateThirdSsoAgentResponseBody) SetData(v int64) *CreateThirdSsoAgentResponseBody {
	s.Data = &v
	return s
}

func (s *CreateThirdSsoAgentResponseBody) SetHttpStatusCode(v int64) *CreateThirdSsoAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateThirdSsoAgentResponseBody) SetMessage(v string) *CreateThirdSsoAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateThirdSsoAgentResponseBody) SetRequestId(v string) *CreateThirdSsoAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateThirdSsoAgentResponseBody) SetSuccess(v bool) *CreateThirdSsoAgentResponseBody {
	s.Success = &v
	return s
}

func (s *CreateThirdSsoAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
