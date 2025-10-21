// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWordGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteWordGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteWordGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteWordGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteWordGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteWordGroupResponseBody
	GetSuccess() *bool
}

type DeleteWordGroupResponseBody struct {
	// Result code, 200 indicates success; others indicate failure.
	//
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// If an error occurs, returns the error message.
	//
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Possible values:
	//
	// - True indicates success.
	//
	// - False indicates failure.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWordGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWordGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWordGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteWordGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteWordGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteWordGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWordGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteWordGroupResponseBody) SetCode(v string) *DeleteWordGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWordGroupResponseBody) SetHttpStatusCode(v int32) *DeleteWordGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteWordGroupResponseBody) SetMessage(v string) *DeleteWordGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWordGroupResponseBody) SetRequestId(v string) *DeleteWordGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWordGroupResponseBody) SetSuccess(v bool) *DeleteWordGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteWordGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
