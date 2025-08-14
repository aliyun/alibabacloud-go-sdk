// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRuleStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRuleStatusResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *ModifyRuleStatusResponseBody
	GetResultObject() *bool
}

type ModifyRuleStatusResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return object.
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s ModifyRuleStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRuleStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRuleStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRuleStatusResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *ModifyRuleStatusResponseBody) SetRequestId(v string) *ModifyRuleStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRuleStatusResponseBody) SetResultObject(v bool) *ModifyRuleStatusResponseBody {
	s.ResultObject = &v
	return s
}

func (s *ModifyRuleStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
