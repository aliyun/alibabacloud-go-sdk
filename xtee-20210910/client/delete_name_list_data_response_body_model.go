// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNameListDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNameListDataResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DeleteNameListDataResponseBody
	GetResultObject() *bool
}

type DeleteNameListDataResponseBody struct {
	// Request ID
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

func (s DeleteNameListDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNameListDataResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNameListDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNameListDataResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DeleteNameListDataResponseBody) SetRequestId(v string) *DeleteNameListDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNameListDataResponseBody) SetResultObject(v bool) *DeleteNameListDataResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DeleteNameListDataResponseBody) Validate() error {
	return dara.Validate(s)
}
