// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFieldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFieldResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DeleteFieldResponseBody
	GetResultObject() *bool
}

type DeleteFieldResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DeleteFieldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFieldResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFieldResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFieldResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DeleteFieldResponseBody) SetRequestId(v string) *DeleteFieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFieldResponseBody) SetResultObject(v bool) *DeleteFieldResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DeleteFieldResponseBody) Validate() error {
	return dara.Validate(s)
}
