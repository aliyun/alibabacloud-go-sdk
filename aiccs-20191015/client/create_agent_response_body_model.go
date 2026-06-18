// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAgentResponseBody
	GetCode() *string
	SetData(v int64) *CreateAgentResponseBody
	GetData() *int64
	SetHttpStatusCode(v int64) *CreateAgentResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *CreateAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAgentResponseBody
	GetSuccess() *bool
}

type CreateAgentResponseBody struct {
	// Status code. A value of "Success" indicates that the request succeeded.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Agent ID.
	//
	// example:
	//
	// 2578****
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
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded. Valid values:
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAgentResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateAgentResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *CreateAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAgentResponseBody) SetCode(v string) *CreateAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAgentResponseBody) SetData(v int64) *CreateAgentResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAgentResponseBody) SetHttpStatusCode(v int64) *CreateAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAgentResponseBody) SetMessage(v string) *CreateAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAgentResponseBody) SetRequestId(v string) *CreateAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentResponseBody) SetSuccess(v bool) *CreateAgentResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
