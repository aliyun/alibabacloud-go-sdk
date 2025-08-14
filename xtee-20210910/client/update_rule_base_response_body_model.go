// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRuleBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRuleBaseResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *UpdateRuleBaseResponseBody
	GetResultObject() *bool
}

type UpdateRuleBaseResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s UpdateRuleBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleBaseResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRuleBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRuleBaseResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *UpdateRuleBaseResponseBody) SetRequestId(v string) *UpdateRuleBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRuleBaseResponseBody) SetResultObject(v bool) *UpdateRuleBaseResponseBody {
	s.ResultObject = &v
	return s
}

func (s *UpdateRuleBaseResponseBody) Validate() error {
	return dara.Validate(s)
}
