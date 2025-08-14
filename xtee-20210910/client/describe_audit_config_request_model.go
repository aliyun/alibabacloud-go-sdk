// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAuditConfigRequest
	GetLang() *string
	SetAuditRelationType(v string) *DescribeAuditConfigRequest
	GetAuditRelationType() *string
	SetRegId(v string) *DescribeAuditConfigRequest
	GetRegId() *string
}

type DescribeAuditConfigRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Audit type
	//
	// example:
	//
	// RULE
	AuditRelationType *string `json:"auditRelationType,omitempty" xml:"auditRelationType,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeAuditConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAuditConfigRequest) GetAuditRelationType() *string {
	return s.AuditRelationType
}

func (s *DescribeAuditConfigRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeAuditConfigRequest) SetLang(v string) *DescribeAuditConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAuditConfigRequest) SetAuditRelationType(v string) *DescribeAuditConfigRequest {
	s.AuditRelationType = &v
	return s
}

func (s *DescribeAuditConfigRequest) SetRegId(v string) *DescribeAuditConfigRequest {
	s.RegId = &v
	return s
}

func (s *DescribeAuditConfigRequest) Validate() error {
	return dara.Validate(s)
}
