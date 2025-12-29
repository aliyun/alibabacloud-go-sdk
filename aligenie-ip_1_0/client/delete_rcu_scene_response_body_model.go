// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRcuSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteRcuSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteRcuSceneResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteRcuSceneResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *DeleteRcuSceneResponseBody
	GetStatusCode() *int32
}

type DeleteRcuSceneResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4F61A7B7-409C-525D-AFDB-238A4E88925A
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

func (s DeleteRcuSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRcuSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRcuSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteRcuSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRcuSceneResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteRcuSceneResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRcuSceneResponseBody) SetMessage(v string) *DeleteRcuSceneResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRcuSceneResponseBody) SetRequestId(v string) *DeleteRcuSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRcuSceneResponseBody) SetResult(v bool) *DeleteRcuSceneResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteRcuSceneResponseBody) SetStatusCode(v int32) *DeleteRcuSceneResponseBody {
	s.StatusCode = &v
	return s
}

func (s *DeleteRcuSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
