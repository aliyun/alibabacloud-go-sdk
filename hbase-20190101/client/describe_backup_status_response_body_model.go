// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupStatus(v string) *DescribeBackupStatusResponseBody
	GetBackupStatus() *string
	SetBdsClusterId(v string) *DescribeBackupStatusResponseBody
	GetBdsClusterId() *string
	SetClusterId(v string) *DescribeBackupStatusResponseBody
	GetClusterId() *string
	SetRequestId(v string) *DescribeBackupStatusResponseBody
	GetRequestId() *string
}

type DescribeBackupStatusResponseBody struct {
	// example:
	//
	// opened
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// example:
	//
	// bds-m5e54q06ceyhxxxx
	BdsClusterId *string `json:"BdsClusterId,omitempty" xml:"BdsClusterId,omitempty"`
	// example:
	//
	// ld-m5eznlga4k5bcxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// F7E71430-A825-470A-B40B-DF3F3AAC9BEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupStatusResponseBody) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeBackupStatusResponseBody) GetBdsClusterId() *string {
	return s.BdsClusterId
}

func (s *DescribeBackupStatusResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeBackupStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupStatusResponseBody) SetBackupStatus(v string) *DescribeBackupStatusResponseBody {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupStatusResponseBody) SetBdsClusterId(v string) *DescribeBackupStatusResponseBody {
	s.BdsClusterId = &v
	return s
}

func (s *DescribeBackupStatusResponseBody) SetClusterId(v string) *DescribeBackupStatusResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupStatusResponseBody) SetRequestId(v string) *DescribeBackupStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
