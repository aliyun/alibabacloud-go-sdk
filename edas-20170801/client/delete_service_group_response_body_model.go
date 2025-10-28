// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteServiceGroupResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteServiceGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteServiceGroupResponseBody
	GetRequestId() *string
}

type DeleteServiceGroupResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ECD1D6FC-4307-4583-BA6F-215F38****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteServiceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteServiceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceGroupResponseBody) SetCode(v int32) *DeleteServiceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteServiceGroupResponseBody) SetMessage(v string) *DeleteServiceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceGroupResponseBody) SetRequestId(v string) *DeleteServiceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
