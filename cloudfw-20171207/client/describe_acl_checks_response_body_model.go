// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclChecksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckRecords(v *DescribeAclChecksResponseBodyCheckRecords) *DescribeAclChecksResponseBody
	GetCheckRecords() *DescribeAclChecksResponseBodyCheckRecords
	SetRequestId(v string) *DescribeAclChecksResponseBody
	GetRequestId() *string
}

type DescribeAclChecksResponseBody struct {
	CheckRecords *DescribeAclChecksResponseBodyCheckRecords `json:"CheckRecords,omitempty" xml:"CheckRecords,omitempty" type:"Struct"`
	// example:
	//
	// 9AABB1B7-C81F-5158-9EF9-B2DD5D3D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAclChecksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclChecksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAclChecksResponseBody) GetCheckRecords() *DescribeAclChecksResponseBodyCheckRecords {
	return s.CheckRecords
}

func (s *DescribeAclChecksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAclChecksResponseBody) SetCheckRecords(v *DescribeAclChecksResponseBodyCheckRecords) *DescribeAclChecksResponseBody {
	s.CheckRecords = v
	return s
}

func (s *DescribeAclChecksResponseBody) SetRequestId(v string) *DescribeAclChecksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAclChecksResponseBody) Validate() error {
	if s.CheckRecords != nil {
		if err := s.CheckRecords.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAclChecksResponseBodyCheckRecords struct {
	// example:
	//
	// Internet
	AclType *string                                             `json:"AclType,omitempty" xml:"AclType,omitempty"`
	Records []*DescribeAclChecksResponseBodyCheckRecordsRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s DescribeAclChecksResponseBodyCheckRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclChecksResponseBodyCheckRecords) GoString() string {
	return s.String()
}

func (s *DescribeAclChecksResponseBodyCheckRecords) GetAclType() *string {
	return s.AclType
}

func (s *DescribeAclChecksResponseBodyCheckRecords) GetRecords() []*DescribeAclChecksResponseBodyCheckRecordsRecords {
	return s.Records
}

func (s *DescribeAclChecksResponseBodyCheckRecords) SetAclType(v string) *DescribeAclChecksResponseBodyCheckRecords {
	s.AclType = &v
	return s
}

func (s *DescribeAclChecksResponseBodyCheckRecords) SetRecords(v []*DescribeAclChecksResponseBodyCheckRecordsRecords) *DescribeAclChecksResponseBodyCheckRecords {
	s.Records = v
	return s
}

func (s *DescribeAclChecksResponseBodyCheckRecords) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAclChecksResponseBodyCheckRecordsRecords struct {
	// example:
	//
	// 1
	AclPendingCount *int64 `json:"AclPendingCount,omitempty" xml:"AclPendingCount,omitempty"`
	// example:
	//
	// 10
	AclTotalCount *int64 `json:"AclTotalCount,omitempty" xml:"AclTotalCount,omitempty"`
	// example:
	//
	// AddressBookDomainValid
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// example:
	//
	// Checked
	CheckStatus *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// example:
	//
	// AddressBookGather
	CheckType *string `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// example:
	//
	// 1724982259
	LastCheckTime *string `json:"LastCheckTime,omitempty" xml:"LastCheckTime,omitempty"`
	// example:
	//
	// Medium
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// task-c92d4544ef7b6a42
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeAclChecksResponseBodyCheckRecordsRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclChecksResponseBodyCheckRecordsRecords) GoString() string {
	return s.String()
}

func (s *DescribeAclChecksResponseBodyCheckRecordsRecords) GetAclPendingCount() *int64 {
	return s.AclPendingCount
}

func (s *DescribeAclChecksResponseBodyCheckRecordsRecords) GetAclTotalCount() *int64 {
	return s.AclTotalCount
}

func (s *DescribeAclChecksResponseBodyCheckRecordsRecords) GetCheckName() *string {
	return s.CheckName
}

func (s *DescribeAclChecksResponseBodyCheckRecordsRecords) GetCheckStatus() *string {
	return s.CheckStatus
}

func (s *DescribeAclChecksResponseBodyCheckRecordsRecords) GetCheckType() *string {
	return s.CheckType
}

func (s *DescribeAclChecksResponseBodyCheckRecordsRecords) GetLastCheckTime() *string {
	return s.LastCheckTime
}

func (s *DescribeAclChecksResponseBodyCheckRecordsRecords) GetLevel() *string {
	return s.Level
}

func (s *DescribeAclChecksResponseBodyCheckRecordsRecords) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeAclChecksResponseBodyCheckRecordsRecords) SetAclPendingCount(v int64) *DescribeAclChecksResponseBodyCheckRecordsRecords {
	s.AclPendingCount = &v
	return s
}

func (s *DescribeAclChecksResponseBodyCheckRecordsRecords) SetAclTotalCount(v int64) *DescribeAclChecksResponseBodyCheckRecordsRecords {
	s.AclTotalCount = &v
	return s
}

func (s *DescribeAclChecksResponseBodyCheckRecordsRecords) SetCheckName(v string) *DescribeAclChecksResponseBodyCheckRecordsRecords {
	s.CheckName = &v
	return s
}

func (s *DescribeAclChecksResponseBodyCheckRecordsRecords) SetCheckStatus(v string) *DescribeAclChecksResponseBodyCheckRecordsRecords {
	s.CheckStatus = &v
	return s
}

func (s *DescribeAclChecksResponseBodyCheckRecordsRecords) SetCheckType(v string) *DescribeAclChecksResponseBodyCheckRecordsRecords {
	s.CheckType = &v
	return s
}

func (s *DescribeAclChecksResponseBodyCheckRecordsRecords) SetLastCheckTime(v string) *DescribeAclChecksResponseBodyCheckRecordsRecords {
	s.LastCheckTime = &v
	return s
}

func (s *DescribeAclChecksResponseBodyCheckRecordsRecords) SetLevel(v string) *DescribeAclChecksResponseBodyCheckRecordsRecords {
	s.Level = &v
	return s
}

func (s *DescribeAclChecksResponseBodyCheckRecordsRecords) SetTaskId(v string) *DescribeAclChecksResponseBodyCheckRecordsRecords {
	s.TaskId = &v
	return s
}

func (s *DescribeAclChecksResponseBodyCheckRecordsRecords) Validate() error {
	return dara.Validate(s)
}
