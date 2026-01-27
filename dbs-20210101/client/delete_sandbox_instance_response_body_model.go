// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSandboxInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSandboxInstanceResponseBody
	GetCode() *string
	SetData(v string) *DeleteSandboxInstanceResponseBody
	GetData() *string
	SetErrCode(v string) *DeleteSandboxInstanceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteSandboxInstanceResponseBody
	GetErrMessage() *string
	SetMessage(v string) *DeleteSandboxInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSandboxInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteSandboxInstanceResponseBody
	GetSuccess() *string
}

type DeleteSandboxInstanceResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// Param.NotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// operation forbidden due to sandbox is creating.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4F1888AC-1138-4995-B9FE-D2734F61C058
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSandboxInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSandboxInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSandboxInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSandboxInstanceResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteSandboxInstanceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteSandboxInstanceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteSandboxInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSandboxInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSandboxInstanceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteSandboxInstanceResponseBody) SetCode(v string) *DeleteSandboxInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSandboxInstanceResponseBody) SetData(v string) *DeleteSandboxInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteSandboxInstanceResponseBody) SetErrCode(v string) *DeleteSandboxInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteSandboxInstanceResponseBody) SetErrMessage(v string) *DeleteSandboxInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteSandboxInstanceResponseBody) SetMessage(v string) *DeleteSandboxInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSandboxInstanceResponseBody) SetRequestId(v string) *DeleteSandboxInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSandboxInstanceResponseBody) SetSuccess(v string) *DeleteSandboxInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSandboxInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
