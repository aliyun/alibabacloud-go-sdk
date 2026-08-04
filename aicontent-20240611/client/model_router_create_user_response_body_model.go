// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterCreateUserResponseBodyData) *ModelRouterCreateUserResponseBody
	GetData() *ModelRouterCreateUserResponseBodyData
	SetErrCode(v string) *ModelRouterCreateUserResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterCreateUserResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterCreateUserResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterCreateUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterCreateUserResponseBody
	GetSuccess() *bool
}

type ModelRouterCreateUserResponseBody struct {
	// The data object.
	//
	// example:
	//
	// { "userId": 30001 }
	Data *ModelRouterCreateUserResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The fault code.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unknown error
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterCreateUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateUserResponseBody) GetData() *ModelRouterCreateUserResponseBodyData {
	return s.Data
}

func (s *ModelRouterCreateUserResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterCreateUserResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterCreateUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterCreateUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterCreateUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterCreateUserResponseBody) SetData(v *ModelRouterCreateUserResponseBodyData) *ModelRouterCreateUserResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterCreateUserResponseBody) SetErrCode(v string) *ModelRouterCreateUserResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterCreateUserResponseBody) SetErrMessage(v string) *ModelRouterCreateUserResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterCreateUserResponseBody) SetHttpStatusCode(v int32) *ModelRouterCreateUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterCreateUserResponseBody) SetRequestId(v string) *ModelRouterCreateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterCreateUserResponseBody) SetSuccess(v bool) *ModelRouterCreateUserResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterCreateUserResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterCreateUserResponseBodyData struct {
	// The user ID. This ID is used as the userId addressing key for subsequent member API operations.
	//
	// example:
	//
	// 30001
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ModelRouterCreateUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateUserResponseBodyData) GetUserId() *int64 {
	return s.UserId
}

func (s *ModelRouterCreateUserResponseBodyData) SetUserId(v int64) *ModelRouterCreateUserResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ModelRouterCreateUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}
