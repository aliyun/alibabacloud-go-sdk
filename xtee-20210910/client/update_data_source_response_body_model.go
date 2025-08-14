// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataSourceResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *UpdateDataSourceResponseBody
	GetResultObject() *bool
}

type UpdateDataSourceResponseBody struct {
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

func (s UpdateDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataSourceResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *UpdateDataSourceResponseBody) SetRequestId(v string) *UpdateDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataSourceResponseBody) SetResultObject(v bool) *UpdateDataSourceResponseBody {
	s.ResultObject = &v
	return s
}

func (s *UpdateDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
