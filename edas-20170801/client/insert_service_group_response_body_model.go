// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertServiceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *InsertServiceGroupResponseBody
	GetCode() *int32
	SetMessage(v string) *InsertServiceGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *InsertServiceGroupResponseBody
	GetRequestId() *string
}

type InsertServiceGroupResponseBody struct {
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
	// ECD1D6FC-4307-4583-BA6F-215F3857EAF4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InsertServiceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsertServiceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *InsertServiceGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *InsertServiceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsertServiceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsertServiceGroupResponseBody) SetCode(v int32) *InsertServiceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *InsertServiceGroupResponseBody) SetMessage(v string) *InsertServiceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *InsertServiceGroupResponseBody) SetRequestId(v string) *InsertServiceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertServiceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
