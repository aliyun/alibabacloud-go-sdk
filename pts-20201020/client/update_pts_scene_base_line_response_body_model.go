// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePtsSceneBaseLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdatePtsSceneBaseLineResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdatePtsSceneBaseLineResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdatePtsSceneBaseLineResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdatePtsSceneBaseLineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdatePtsSceneBaseLineResponseBody
	GetSuccess() *bool
}

type UpdatePtsSceneBaseLineResponseBody struct {
	// The system status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F7D2CE0-AE4C-4143-955A-8E4595AF86A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// 	- true
	//
	// 	- false:
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdatePtsSceneBaseLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePtsSceneBaseLineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePtsSceneBaseLineResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdatePtsSceneBaseLineResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdatePtsSceneBaseLineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePtsSceneBaseLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePtsSceneBaseLineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePtsSceneBaseLineResponseBody) SetCode(v string) *UpdatePtsSceneBaseLineResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePtsSceneBaseLineResponseBody) SetHttpStatusCode(v int32) *UpdatePtsSceneBaseLineResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdatePtsSceneBaseLineResponseBody) SetMessage(v string) *UpdatePtsSceneBaseLineResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePtsSceneBaseLineResponseBody) SetRequestId(v string) *UpdatePtsSceneBaseLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePtsSceneBaseLineResponseBody) SetSuccess(v bool) *UpdatePtsSceneBaseLineResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePtsSceneBaseLineResponseBody) Validate() error {
	return dara.Validate(s)
}
