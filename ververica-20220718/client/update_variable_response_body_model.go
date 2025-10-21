// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateVariableResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *Variable) *UpdateVariableResponseBody
	GetData() *Variable
	SetErrorCode(v string) *UpdateVariableResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateVariableResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *UpdateVariableResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *UpdateVariableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateVariableResponseBody
	GetSuccess() *bool
}

type UpdateVariableResponseBody struct {
	// example:
	//
	// “”
	AccessDeniedDetail *string   `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	Data               *Variable `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 1EF03B0C-F44F-47AD-BB48-D002D0F7B8C9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVariableResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVariableResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateVariableResponseBody) GetData() *Variable {
	return s.Data
}

func (s *UpdateVariableResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateVariableResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateVariableResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVariableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateVariableResponseBody) SetAccessDeniedDetail(v string) *UpdateVariableResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateVariableResponseBody) SetData(v *Variable) *UpdateVariableResponseBody {
	s.Data = v
	return s
}

func (s *UpdateVariableResponseBody) SetErrorCode(v string) *UpdateVariableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateVariableResponseBody) SetErrorMessage(v string) *UpdateVariableResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateVariableResponseBody) SetHttpCode(v int32) *UpdateVariableResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateVariableResponseBody) SetRequestId(v string) *UpdateVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVariableResponseBody) SetSuccess(v bool) *UpdateVariableResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateVariableResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
