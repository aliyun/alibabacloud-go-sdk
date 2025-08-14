// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEventStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyEventStatusResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *ModifyEventStatusResponseBody
	GetResultObject() *bool
}

type ModifyEventStatusResponseBody struct {
	// Request ID
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s ModifyEventStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyEventStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEventStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyEventStatusResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *ModifyEventStatusResponseBody) SetRequestId(v string) *ModifyEventStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyEventStatusResponseBody) SetResultObject(v bool) *ModifyEventStatusResponseBody {
	s.ResultObject = &v
	return s
}

func (s *ModifyEventStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
