// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAuthRuleResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *UpdateAuthRuleResponseBody
	GetResultObject() *bool
}

type UpdateAuthRuleResponseBody struct {
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

func (s UpdateAuthRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAuthRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAuthRuleResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *UpdateAuthRuleResponseBody) SetRequestId(v string) *UpdateAuthRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAuthRuleResponseBody) SetResultObject(v bool) *UpdateAuthRuleResponseBody {
	s.ResultObject = &v
	return s
}

func (s *UpdateAuthRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
