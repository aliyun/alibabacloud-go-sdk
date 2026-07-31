// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReservedInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyReservedInstanceAttributeResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyReservedInstanceAttributeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyReservedInstanceAttributeResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyReservedInstanceAttributeResponseBody
	GetRequestId() *string
}

type ModifyReservedInstanceAttributeResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned for the instance operation. A return value of Success indicates success. For more information, see the error codes section.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyReservedInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyReservedInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyReservedInstanceAttributeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyReservedInstanceAttributeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyReservedInstanceAttributeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyReservedInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyReservedInstanceAttributeResponseBody) SetCode(v string) *ModifyReservedInstanceAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyReservedInstanceAttributeResponseBody) SetHttpStatusCode(v int32) *ModifyReservedInstanceAttributeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyReservedInstanceAttributeResponseBody) SetMessage(v string) *ModifyReservedInstanceAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyReservedInstanceAttributeResponseBody) SetRequestId(v string) *ModifyReservedInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyReservedInstanceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
