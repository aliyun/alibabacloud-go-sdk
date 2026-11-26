// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScriptHotWordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteScriptHotWordResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteScriptHotWordResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteScriptHotWordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteScriptHotWordResponseBody
	GetSuccess() *bool
}

type DeleteScriptHotWordResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteScriptHotWordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteScriptHotWordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScriptHotWordResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteScriptHotWordResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteScriptHotWordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteScriptHotWordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteScriptHotWordResponseBody) SetErrorCode(v string) *DeleteScriptHotWordResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteScriptHotWordResponseBody) SetErrorMessage(v string) *DeleteScriptHotWordResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteScriptHotWordResponseBody) SetRequestId(v string) *DeleteScriptHotWordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScriptHotWordResponseBody) SetSuccess(v bool) *DeleteScriptHotWordResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteScriptHotWordResponseBody) Validate() error {
	return dara.Validate(s)
}
