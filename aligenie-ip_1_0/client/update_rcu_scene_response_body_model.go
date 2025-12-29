// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRcuSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateRcuSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateRcuSceneResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateRcuSceneResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *UpdateRcuSceneResponseBody
	GetStatusCode() *int32
}

type UpdateRcuSceneResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3A680F3A-6672-5A47-AB28-12BBCD80C679
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s UpdateRcuSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRcuSceneResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRcuSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateRcuSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRcuSceneResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateRcuSceneResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRcuSceneResponseBody) SetMessage(v string) *UpdateRcuSceneResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRcuSceneResponseBody) SetRequestId(v string) *UpdateRcuSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRcuSceneResponseBody) SetResult(v bool) *UpdateRcuSceneResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateRcuSceneResponseBody) SetStatusCode(v int32) *UpdateRcuSceneResponseBody {
	s.StatusCode = &v
	return s
}

func (s *UpdateRcuSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
