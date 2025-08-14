// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecallRuleAuditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *RecallRuleAuditRequest
	GetId() *int64
	SetRegId(v string) *RecallRuleAuditRequest
	GetRegId() *string
}

type RecallRuleAuditRequest struct {
	// Primary key ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s RecallRuleAuditRequest) String() string {
	return dara.Prettify(s)
}

func (s RecallRuleAuditRequest) GoString() string {
	return s.String()
}

func (s *RecallRuleAuditRequest) GetId() *int64 {
	return s.Id
}

func (s *RecallRuleAuditRequest) GetRegId() *string {
	return s.RegId
}

func (s *RecallRuleAuditRequest) SetId(v int64) *RecallRuleAuditRequest {
	s.Id = &v
	return s
}

func (s *RecallRuleAuditRequest) SetRegId(v string) *RecallRuleAuditRequest {
	s.RegId = &v
	return s
}

func (s *RecallRuleAuditRequest) Validate() error {
	return dara.Validate(s)
}
