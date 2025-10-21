// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteModelInstanceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteModelInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteModelInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteModelInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteModelInstanceResponseBody
	GetSuccess() *bool
}

type DeleteModelInstanceResponseBody struct {
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteModelInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteModelInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteModelInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteModelInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteModelInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteModelInstanceResponseBody) SetCode(v string) *DeleteModelInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteModelInstanceResponseBody) SetHttpStatusCode(v int32) *DeleteModelInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteModelInstanceResponseBody) SetMessage(v string) *DeleteModelInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteModelInstanceResponseBody) SetRequestId(v string) *DeleteModelInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteModelInstanceResponseBody) SetSuccess(v bool) *DeleteModelInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteModelInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
