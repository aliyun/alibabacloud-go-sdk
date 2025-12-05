// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPtsSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyPtsSceneResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyPtsSceneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyPtsSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyPtsSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyPtsSceneResponseBody
	GetSuccess() *bool
}

type ModifyPtsSceneResponseBody struct {
	// The system status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code that is returned.
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
	// 449ADAFB-8DA4-4317-A284-4922D04DE828
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

func (s ModifyPtsSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPtsSceneResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPtsSceneResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyPtsSceneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyPtsSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyPtsSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPtsSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyPtsSceneResponseBody) SetCode(v string) *ModifyPtsSceneResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyPtsSceneResponseBody) SetHttpStatusCode(v int32) *ModifyPtsSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyPtsSceneResponseBody) SetMessage(v string) *ModifyPtsSceneResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyPtsSceneResponseBody) SetRequestId(v string) *ModifyPtsSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPtsSceneResponseBody) SetSuccess(v bool) *ModifyPtsSceneResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyPtsSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
