// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDataSourceResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CreateDataSourceResponseBody
	GetResultObject() *bool
}

type CreateDataSourceResponseBody struct {
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

func (s CreateDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataSourceResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CreateDataSourceResponseBody) SetRequestId(v string) *CreateDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataSourceResponseBody) SetResultObject(v bool) *CreateDataSourceResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CreateDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
