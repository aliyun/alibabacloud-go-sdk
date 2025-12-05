// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveOpenJMeterSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveOpenJMeterSceneResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SaveOpenJMeterSceneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SaveOpenJMeterSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveOpenJMeterSceneResponseBody
	GetRequestId() *string
	SetSceneId(v string) *SaveOpenJMeterSceneResponseBody
	GetSceneId() *string
	SetSuccess(v bool) *SaveOpenJMeterSceneResponseBody
	GetSuccess() *bool
}

type SaveOpenJMeterSceneResponseBody struct {
	// The system status code. If the request was successful, this parameter is not returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. If the request was successful, this parameter is not returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. If the request was successful, this parameter is not returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A8E16480-15C1-555A-922F-B736A005E52D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the created or updated scenario.
	//
	// example:
	//
	// DYYPZIH
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// Indicates whether the request was successful. Valid values:
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

func (s SaveOpenJMeterSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveOpenJMeterSceneResponseBody) GoString() string {
	return s.String()
}

func (s *SaveOpenJMeterSceneResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveOpenJMeterSceneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SaveOpenJMeterSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveOpenJMeterSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveOpenJMeterSceneResponseBody) GetSceneId() *string {
	return s.SceneId
}

func (s *SaveOpenJMeterSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveOpenJMeterSceneResponseBody) SetCode(v string) *SaveOpenJMeterSceneResponseBody {
	s.Code = &v
	return s
}

func (s *SaveOpenJMeterSceneResponseBody) SetHttpStatusCode(v int32) *SaveOpenJMeterSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveOpenJMeterSceneResponseBody) SetMessage(v string) *SaveOpenJMeterSceneResponseBody {
	s.Message = &v
	return s
}

func (s *SaveOpenJMeterSceneResponseBody) SetRequestId(v string) *SaveOpenJMeterSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveOpenJMeterSceneResponseBody) SetSceneId(v string) *SaveOpenJMeterSceneResponseBody {
	s.SceneId = &v
	return s
}

func (s *SaveOpenJMeterSceneResponseBody) SetSuccess(v bool) *SaveOpenJMeterSceneResponseBody {
	s.Success = &v
	return s
}

func (s *SaveOpenJMeterSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
