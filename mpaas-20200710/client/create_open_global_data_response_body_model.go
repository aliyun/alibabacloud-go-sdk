// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpenGlobalDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateOpenGlobalDataResponseBody
	GetRequestId() *string
	SetResult(v bool) *CreateOpenGlobalDataResponseBody
	GetResult() *bool
	SetResultCode(v string) *CreateOpenGlobalDataResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *CreateOpenGlobalDataResponseBody
	GetResultMessage() *string
}

type CreateOpenGlobalDataResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result        *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateOpenGlobalDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenGlobalDataResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOpenGlobalDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOpenGlobalDataResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CreateOpenGlobalDataResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateOpenGlobalDataResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateOpenGlobalDataResponseBody) SetRequestId(v string) *CreateOpenGlobalDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOpenGlobalDataResponseBody) SetResult(v bool) *CreateOpenGlobalDataResponseBody {
	s.Result = &v
	return s
}

func (s *CreateOpenGlobalDataResponseBody) SetResultCode(v string) *CreateOpenGlobalDataResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateOpenGlobalDataResponseBody) SetResultMessage(v string) *CreateOpenGlobalDataResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateOpenGlobalDataResponseBody) Validate() error {
	return dara.Validate(s)
}
