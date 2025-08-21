// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ModifyInstanceAttributeResponseBody
	GetCode() *int32
	SetRequestId(v string) *ModifyInstanceAttributeResponseBody
	GetRequestId() *string
}

type ModifyInstanceAttributeResponseBody struct {
	// The returned service code. 0 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ModifyInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceAttributeResponseBody) SetCode(v int32) *ModifyInstanceAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyInstanceAttributeResponseBody) SetRequestId(v string) *ModifyInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
