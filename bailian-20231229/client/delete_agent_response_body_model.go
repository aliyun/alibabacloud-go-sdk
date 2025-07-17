// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAgentResponseBody
	GetCode() *string
	SetData(v string) *DeleteAgentResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteAgentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAgentResponseBody
	GetSuccess() *bool
}

type DeleteAgentResponseBody struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	Data           *string `json:"data,omitempty" xml:"data,omitempty"`
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAgentResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteAgentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAgentResponseBody) SetCode(v string) *DeleteAgentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAgentResponseBody) SetData(v string) *DeleteAgentResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAgentResponseBody) SetHttpStatusCode(v int32) *DeleteAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteAgentResponseBody) SetMessage(v string) *DeleteAgentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAgentResponseBody) SetRequestId(v string) *DeleteAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgentResponseBody) SetSuccess(v bool) *DeleteAgentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
