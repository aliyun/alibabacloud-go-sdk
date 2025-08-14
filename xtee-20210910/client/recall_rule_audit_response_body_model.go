// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecallRuleAuditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RecallRuleAuditResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *RecallRuleAuditResponseBody
	GetResultObject() *bool
}

type RecallRuleAuditResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 4C7DC1C8-557F-5483-9E96-74D77A15EEB5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s RecallRuleAuditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecallRuleAuditResponseBody) GoString() string {
	return s.String()
}

func (s *RecallRuleAuditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecallRuleAuditResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *RecallRuleAuditResponseBody) SetRequestId(v string) *RecallRuleAuditResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecallRuleAuditResponseBody) SetResultObject(v bool) *RecallRuleAuditResponseBody {
	s.ResultObject = &v
	return s
}

func (s *RecallRuleAuditResponseBody) Validate() error {
	return dara.Validate(s)
}
