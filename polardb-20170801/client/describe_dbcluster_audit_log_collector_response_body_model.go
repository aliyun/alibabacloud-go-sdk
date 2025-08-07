// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterAuditLogCollectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCollectorStatus(v string) *DescribeDBClusterAuditLogCollectorResponseBody
	GetCollectorStatus() *string
	SetRequestId(v string) *DescribeDBClusterAuditLogCollectorResponseBody
	GetRequestId() *string
}

type DescribeDBClusterAuditLogCollectorResponseBody struct {
	// The status of SQL collector. Valid values:
	//
	// 	- Enable
	//
	// 	- Disabled
	//
	// example:
	//
	// Disabled
	CollectorStatus *string `json:"CollectorStatus,omitempty" xml:"CollectorStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 59011D2B-2A38-4207-A86C-72BC1F882D19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterAuditLogCollectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAuditLogCollectorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAuditLogCollectorResponseBody) GetCollectorStatus() *string {
	return s.CollectorStatus
}

func (s *DescribeDBClusterAuditLogCollectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterAuditLogCollectorResponseBody) SetCollectorStatus(v string) *DescribeDBClusterAuditLogCollectorResponseBody {
	s.CollectorStatus = &v
	return s
}

func (s *DescribeDBClusterAuditLogCollectorResponseBody) SetRequestId(v string) *DescribeDBClusterAuditLogCollectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterAuditLogCollectorResponseBody) Validate() error {
	return dara.Validate(s)
}
