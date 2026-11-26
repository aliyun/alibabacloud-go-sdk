// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScriptHotWordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateScriptHotWordResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateScriptHotWordResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateScriptHotWordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateScriptHotWordResponseBody
	GetSuccess() *bool
}

type UpdateScriptHotWordResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateScriptHotWordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptHotWordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScriptHotWordResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateScriptHotWordResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateScriptHotWordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateScriptHotWordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateScriptHotWordResponseBody) SetErrorCode(v string) *UpdateScriptHotWordResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateScriptHotWordResponseBody) SetErrorMessage(v string) *UpdateScriptHotWordResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateScriptHotWordResponseBody) SetRequestId(v string) *UpdateScriptHotWordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateScriptHotWordResponseBody) SetSuccess(v bool) *UpdateScriptHotWordResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateScriptHotWordResponseBody) Validate() error {
	return dara.Validate(s)
}
