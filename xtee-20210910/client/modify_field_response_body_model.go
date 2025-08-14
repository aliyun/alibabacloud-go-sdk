// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFieldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyFieldResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *ModifyFieldResponseBody
	GetResultObject() *bool
}

type ModifyFieldResponseBody struct {
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

func (s ModifyFieldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyFieldResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFieldResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyFieldResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *ModifyFieldResponseBody) SetRequestId(v string) *ModifyFieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFieldResponseBody) SetResultObject(v bool) *ModifyFieldResponseBody {
	s.ResultObject = &v
	return s
}

func (s *ModifyFieldResponseBody) Validate() error {
	return dara.Validate(s)
}
