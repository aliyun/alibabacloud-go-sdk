// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePtsSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePtsSceneResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreatePtsSceneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreatePtsSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePtsSceneResponseBody
	GetRequestId() *string
	SetSceneId(v string) *CreatePtsSceneResponseBody
	GetSceneId() *string
	SetSuccess(v bool) *CreatePtsSceneResponseBody
	GetSuccess() *bool
}

type CreatePtsSceneResponseBody struct {
	// The system status code. If the request was successful, no data is returned.
	//
	// example:
	//
	// 4001
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. If the request was successful, no data is returned.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. If the request was successful, no data is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F7D2CE0-AE4C-4143-955A-8E4595AF86A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the created scenario.
	//
	// example:
	//
	// SDR3CX
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePtsSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePtsSceneResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePtsSceneResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePtsSceneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreatePtsSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePtsSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePtsSceneResponseBody) GetSceneId() *string {
	return s.SceneId
}

func (s *CreatePtsSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePtsSceneResponseBody) SetCode(v string) *CreatePtsSceneResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePtsSceneResponseBody) SetHttpStatusCode(v int32) *CreatePtsSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreatePtsSceneResponseBody) SetMessage(v string) *CreatePtsSceneResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePtsSceneResponseBody) SetRequestId(v string) *CreatePtsSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePtsSceneResponseBody) SetSceneId(v string) *CreatePtsSceneResponseBody {
	s.SceneId = &v
	return s
}

func (s *CreatePtsSceneResponseBody) SetSuccess(v bool) *CreatePtsSceneResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePtsSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
