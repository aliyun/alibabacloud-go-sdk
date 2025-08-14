// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelNameIsDuplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModelNameIsDuplicationResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *ModelNameIsDuplicationResponseBody
	GetResultObject() *bool
}

type ModelNameIsDuplicationResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	//
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
}

func (s ModelNameIsDuplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelNameIsDuplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ModelNameIsDuplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelNameIsDuplicationResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *ModelNameIsDuplicationResponseBody) SetRequestId(v string) *ModelNameIsDuplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelNameIsDuplicationResponseBody) SetResultObject(v bool) *ModelNameIsDuplicationResponseBody {
	s.ResultObject = &v
	return s
}

func (s *ModelNameIsDuplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
