// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSaslUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteSaslUserResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteSaslUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSaslUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSaslUserResponseBody
	GetSuccess() *bool
}

type DeleteSaslUserResponseBody struct {
	// The HTTP status code. If the request is successful, 200 is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3CB89F5C-CD97-4C1D-BC7C-FEDEC2F4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSaslUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSaslUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSaslUserResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteSaslUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSaslUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSaslUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSaslUserResponseBody) SetCode(v int32) *DeleteSaslUserResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSaslUserResponseBody) SetMessage(v string) *DeleteSaslUserResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSaslUserResponseBody) SetRequestId(v string) *DeleteSaslUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSaslUserResponseBody) SetSuccess(v bool) *DeleteSaslUserResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSaslUserResponseBody) Validate() error {
	return dara.Validate(s)
}
