// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelIsUsingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModelIsUsingResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *ModelIsUsingResponseBody
	GetResultObject() *bool
}

type ModelIsUsingResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Result object.
	//
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
}

func (s ModelIsUsingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelIsUsingResponseBody) GoString() string {
	return s.String()
}

func (s *ModelIsUsingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelIsUsingResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *ModelIsUsingResponseBody) SetRequestId(v string) *ModelIsUsingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelIsUsingResponseBody) SetResultObject(v bool) *ModelIsUsingResponseBody {
	s.ResultObject = &v
	return s
}

func (s *ModelIsUsingResponseBody) Validate() error {
	return dara.Validate(s)
}
