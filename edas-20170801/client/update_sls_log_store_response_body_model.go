// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSlsLogStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateSlsLogStoreResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateSlsLogStoreResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSlsLogStoreResponseBody
	GetRequestId() *string
}

type UpdateSlsLogStoreResponseBody struct {
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
	// D16979DC-4D42-**************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSlsLogStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSlsLogStoreResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSlsLogStoreResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateSlsLogStoreResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSlsLogStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSlsLogStoreResponseBody) SetCode(v int32) *UpdateSlsLogStoreResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSlsLogStoreResponseBody) SetMessage(v string) *UpdateSlsLogStoreResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSlsLogStoreResponseBody) SetRequestId(v string) *UpdateSlsLogStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSlsLogStoreResponseBody) Validate() error {
	return dara.Validate(s)
}
