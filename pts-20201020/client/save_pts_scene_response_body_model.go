// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSavePtsSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SavePtsSceneResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SavePtsSceneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SavePtsSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *SavePtsSceneResponseBody
	GetRequestId() *string
	SetSceneId(v string) *SavePtsSceneResponseBody
	GetSceneId() *string
	SetSuccess(v bool) *SavePtsSceneResponseBody
	GetSuccess() *bool
}

type SavePtsSceneResponseBody struct {
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
	// If the operation is successful, this parameter is not returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A8E16480-15C1-555A-922F-B736A005E52D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scenario.
	//
	// example:
	//
	// IUYAHGJ
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SavePtsSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SavePtsSceneResponseBody) GoString() string {
	return s.String()
}

func (s *SavePtsSceneResponseBody) GetCode() *string {
	return s.Code
}

func (s *SavePtsSceneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SavePtsSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SavePtsSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SavePtsSceneResponseBody) GetSceneId() *string {
	return s.SceneId
}

func (s *SavePtsSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SavePtsSceneResponseBody) SetCode(v string) *SavePtsSceneResponseBody {
	s.Code = &v
	return s
}

func (s *SavePtsSceneResponseBody) SetHttpStatusCode(v int32) *SavePtsSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SavePtsSceneResponseBody) SetMessage(v string) *SavePtsSceneResponseBody {
	s.Message = &v
	return s
}

func (s *SavePtsSceneResponseBody) SetRequestId(v string) *SavePtsSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *SavePtsSceneResponseBody) SetSceneId(v string) *SavePtsSceneResponseBody {
	s.SceneId = &v
	return s
}

func (s *SavePtsSceneResponseBody) SetSuccess(v bool) *SavePtsSceneResponseBody {
	s.Success = &v
	return s
}

func (s *SavePtsSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
