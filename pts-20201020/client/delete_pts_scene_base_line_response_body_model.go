// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePtsSceneBaseLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeletePtsSceneBaseLineResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeletePtsSceneBaseLineResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeletePtsSceneBaseLineResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeletePtsSceneBaseLineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeletePtsSceneBaseLineResponseBody
	GetSuccess() *bool
}

type DeletePtsSceneBaseLineResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4F7D2XE0-AE4C-4143-955A-8E4595AF86A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeletePtsSceneBaseLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePtsSceneBaseLineResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePtsSceneBaseLineResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeletePtsSceneBaseLineResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeletePtsSceneBaseLineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePtsSceneBaseLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePtsSceneBaseLineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeletePtsSceneBaseLineResponseBody) SetCode(v string) *DeletePtsSceneBaseLineResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePtsSceneBaseLineResponseBody) SetHttpStatusCode(v int32) *DeletePtsSceneBaseLineResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeletePtsSceneBaseLineResponseBody) SetMessage(v string) *DeletePtsSceneBaseLineResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePtsSceneBaseLineResponseBody) SetRequestId(v string) *DeletePtsSceneBaseLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePtsSceneBaseLineResponseBody) SetSuccess(v bool) *DeletePtsSceneBaseLineResponseBody {
	s.Success = &v
	return s
}

func (s *DeletePtsSceneBaseLineResponseBody) Validate() error {
	return dara.Validate(s)
}
