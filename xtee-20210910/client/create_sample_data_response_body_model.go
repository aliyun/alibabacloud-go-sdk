// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSampleDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSampleDataResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CreateSampleDataResponseBody
	GetResultObject() *bool
}

type CreateSampleDataResponseBody struct {
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

func (s CreateSampleDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSampleDataResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSampleDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSampleDataResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CreateSampleDataResponseBody) SetRequestId(v string) *CreateSampleDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSampleDataResponseBody) SetResultObject(v bool) *CreateSampleDataResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CreateSampleDataResponseBody) Validate() error {
	return dara.Validate(s)
}
