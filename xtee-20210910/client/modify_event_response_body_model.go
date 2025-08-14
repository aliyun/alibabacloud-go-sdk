// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyEventResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *ModifyEventResponseBody
	GetResultObject() *bool
}

type ModifyEventResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s ModifyEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyEventResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyEventResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *ModifyEventResponseBody) SetRequestId(v string) *ModifyEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyEventResponseBody) SetResultObject(v bool) *ModifyEventResponseBody {
	s.ResultObject = &v
	return s
}

func (s *ModifyEventResponseBody) Validate() error {
	return dara.Validate(s)
}
