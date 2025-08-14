// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRulePriorityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRulePriorityResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *ModifyRulePriorityResponseBody
	GetResultObject() *bool
}

type ModifyRulePriorityResponseBody struct {
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

func (s ModifyRulePriorityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRulePriorityResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRulePriorityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRulePriorityResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *ModifyRulePriorityResponseBody) SetRequestId(v string) *ModifyRulePriorityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRulePriorityResponseBody) SetResultObject(v bool) *ModifyRulePriorityResponseBody {
	s.ResultObject = &v
	return s
}

func (s *ModifyRulePriorityResponseBody) Validate() error {
	return dara.Validate(s)
}
