// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSampleDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSampleDataResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DeleteSampleDataResponseBody
	GetResultObject() *bool
}

type DeleteSampleDataResponseBody struct {
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

func (s DeleteSampleDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSampleDataResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSampleDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSampleDataResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DeleteSampleDataResponseBody) SetRequestId(v string) *DeleteSampleDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSampleDataResponseBody) SetResultObject(v bool) *DeleteSampleDataResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DeleteSampleDataResponseBody) Validate() error {
	return dara.Validate(s)
}
