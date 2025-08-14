// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataSourceResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DeleteDataSourceResponseBody
	GetResultObject() *bool
}

type DeleteDataSourceResponseBody struct {
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

func (s DeleteDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataSourceResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DeleteDataSourceResponseBody) SetRequestId(v string) *DeleteDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataSourceResponseBody) SetResultObject(v bool) *DeleteDataSourceResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DeleteDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
