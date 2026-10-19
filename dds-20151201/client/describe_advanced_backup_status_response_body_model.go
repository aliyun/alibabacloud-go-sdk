// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvancedBackupStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAdvancedBackupStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeAdvancedBackupStatusResponseBody
	GetStatus() *string
}

type DescribeAdvancedBackupStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 55A2504F-xxxx-xxxx-xxxx-C7C5E57350A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// - Enabled: Advanced backup is enabled.
	//
	// - Disabled: Advanced backup is not enabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAdvancedBackupStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvancedBackupStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdvancedBackupStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAdvancedBackupStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeAdvancedBackupStatusResponseBody) SetRequestId(v string) *DescribeAdvancedBackupStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdvancedBackupStatusResponseBody) SetStatus(v string) *DescribeAdvancedBackupStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeAdvancedBackupStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
