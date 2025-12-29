// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRcuSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CreateRcuSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateRcuSceneResponseBody
	GetRequestId() *string
	SetResult(v bool) *CreateRcuSceneResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *CreateRcuSceneResponseBody
	GetStatusCode() *int32
}

type CreateRcuSceneResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 844BA5CE-E30A-53CB-8A11-DE1F344C846D
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

func (s CreateRcuSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRcuSceneResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRcuSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateRcuSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRcuSceneResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CreateRcuSceneResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRcuSceneResponseBody) SetMessage(v string) *CreateRcuSceneResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRcuSceneResponseBody) SetRequestId(v string) *CreateRcuSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRcuSceneResponseBody) SetResult(v bool) *CreateRcuSceneResponseBody {
	s.Result = &v
	return s
}

func (s *CreateRcuSceneResponseBody) SetStatusCode(v int32) *CreateRcuSceneResponseBody {
	s.StatusCode = &v
	return s
}

func (s *CreateRcuSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
