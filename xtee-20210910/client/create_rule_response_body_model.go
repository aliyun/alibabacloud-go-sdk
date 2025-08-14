// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRuleResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CreateRuleResponseBody
	GetResultObject() *bool
}

type CreateRuleResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return result.
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s CreateRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRuleResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CreateRuleResponseBody) SetRequestId(v string) *CreateRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRuleResponseBody) SetResultObject(v bool) *CreateRuleResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CreateRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
