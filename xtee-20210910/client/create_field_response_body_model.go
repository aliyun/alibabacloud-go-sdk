// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFieldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateFieldResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CreateFieldResponseBody
	GetResultObject() *bool
}

type CreateFieldResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object.
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s CreateFieldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFieldResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFieldResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFieldResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CreateFieldResponseBody) SetRequestId(v string) *CreateFieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFieldResponseBody) SetResultObject(v bool) *CreateFieldResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CreateFieldResponseBody) Validate() error {
	return dara.Validate(s)
}
