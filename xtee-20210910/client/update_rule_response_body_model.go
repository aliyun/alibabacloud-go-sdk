// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRuleResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *UpdateRuleResponseBody
	GetResultObject() *bool
}

type UpdateRuleResponseBody struct {
	// Request ID
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

func (s UpdateRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRuleResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *UpdateRuleResponseBody) SetRequestId(v string) *UpdateRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRuleResponseBody) SetResultObject(v bool) *UpdateRuleResponseBody {
	s.ResultObject = &v
	return s
}

func (s *UpdateRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
