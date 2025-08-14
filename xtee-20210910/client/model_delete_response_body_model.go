// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModelDeleteResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *ModelDeleteResponseBody
	GetResultObject() *bool
}

type ModelDeleteResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Deletion result.
	//
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
}

func (s ModelDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *ModelDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelDeleteResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *ModelDeleteResponseBody) SetRequestId(v string) *ModelDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelDeleteResponseBody) SetResultObject(v bool) *ModelDeleteResponseBody {
	s.ResultObject = &v
	return s
}

func (s *ModelDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
