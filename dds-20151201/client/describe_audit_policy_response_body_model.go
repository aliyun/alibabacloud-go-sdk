// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogAuditStatus(v string) *DescribeAuditPolicyResponseBody
	GetLogAuditStatus() *string
	SetRequestId(v string) *DescribeAuditPolicyResponseBody
	GetRequestId() *string
}

type DescribeAuditPolicyResponseBody struct {
	// The status of the audit log.
	//
	// - Enable: The audit log is enabled.
	//
	// - Disabled: The audit log is disabled.
	//
	// The default value is Disabled.
	//
	// example:
	//
	// Enable
	LogAuditStatus *string `json:"LogAuditStatus,omitempty" xml:"LogAuditStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 111E7B16-0A87-4CBA-B271-F34AD61E099F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAuditPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditPolicyResponseBody) GetLogAuditStatus() *string {
	return s.LogAuditStatus
}

func (s *DescribeAuditPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAuditPolicyResponseBody) SetLogAuditStatus(v string) *DescribeAuditPolicyResponseBody {
	s.LogAuditStatus = &v
	return s
}

func (s *DescribeAuditPolicyResponseBody) SetRequestId(v string) *DescribeAuditPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
