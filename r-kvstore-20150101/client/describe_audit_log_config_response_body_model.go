// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditLogConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDbAudit(v string) *DescribeAuditLogConfigResponseBody
	GetDbAudit() *string
	SetRequestId(v string) *DescribeAuditLogConfigResponseBody
	GetRequestId() *string
	SetRetention(v string) *DescribeAuditLogConfigResponseBody
	GetRetention() *string
}

type DescribeAuditLogConfigResponseBody struct {
	// Indicates whether audit logging is enabled. Valid values:
	//
	// 	- **true**: Enabled.
	//
	// 	- **false**: Disabled.
	//
	// > You can call the [ModifyAuditLogConfig](https://help.aliyun.com/document_detail/473829.html) operation to enable or disable audit logging for the instance.
	//
	// example:
	//
	// true
	DbAudit *string `json:"DbAudit,omitempty" xml:"DbAudit,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2BE6E619-A657-42E3-AD2D-18F8428A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The retention period of audit logs. Unit: days.
	//
	// example:
	//
	// 5
	Retention *string `json:"Retention,omitempty" xml:"Retention,omitempty"`
}

func (s DescribeAuditLogConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogConfigResponseBody) GetDbAudit() *string {
	return s.DbAudit
}

func (s *DescribeAuditLogConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAuditLogConfigResponseBody) GetRetention() *string {
	return s.Retention
}

func (s *DescribeAuditLogConfigResponseBody) SetDbAudit(v string) *DescribeAuditLogConfigResponseBody {
	s.DbAudit = &v
	return s
}

func (s *DescribeAuditLogConfigResponseBody) SetRequestId(v string) *DescribeAuditLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditLogConfigResponseBody) SetRetention(v string) *DescribeAuditLogConfigResponseBody {
	s.Retention = &v
	return s
}

func (s *DescribeAuditLogConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
