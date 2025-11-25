// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclBackupListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackups(v []*DescribeAclBackupListResponseBodyBackups) *DescribeAclBackupListResponseBody
	GetBackups() []*DescribeAclBackupListResponseBodyBackups
	SetRequestId(v string) *DescribeAclBackupListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAclBackupListResponseBody
	GetTotalCount() *int32
}

type DescribeAclBackupListResponseBody struct {
	Backups []*DescribeAclBackupListResponseBodyBackups `json:"Backups,omitempty" xml:"Backups,omitempty" type:"Repeated"`
	// example:
	//
	// 6C9105F2-9F31-5A62-8D52-FA65A3E5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 32
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAclBackupListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclBackupListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAclBackupListResponseBody) GetBackups() []*DescribeAclBackupListResponseBodyBackups {
	return s.Backups
}

func (s *DescribeAclBackupListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAclBackupListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAclBackupListResponseBody) SetBackups(v []*DescribeAclBackupListResponseBodyBackups) *DescribeAclBackupListResponseBody {
	s.Backups = v
	return s
}

func (s *DescribeAclBackupListResponseBody) SetRequestId(v string) *DescribeAclBackupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAclBackupListResponseBody) SetTotalCount(v int32) *DescribeAclBackupListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAclBackupListResponseBody) Validate() error {
	if s.Backups != nil {
		for _, item := range s.Backups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAclBackupListResponseBodyBackups struct {
	// example:
	//
	// 10
	AclCount *int32 `json:"AclCount,omitempty" xml:"AclCount,omitempty"`
	// example:
	//
	// 1743683400
	BackUpTime *int64 `json:"BackUpTime,omitempty" xml:"BackUpTime,omitempty"`
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s DescribeAclBackupListResponseBodyBackups) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclBackupListResponseBodyBackups) GoString() string {
	return s.String()
}

func (s *DescribeAclBackupListResponseBodyBackups) GetAclCount() *int32 {
	return s.AclCount
}

func (s *DescribeAclBackupListResponseBodyBackups) GetBackUpTime() *int64 {
	return s.BackUpTime
}

func (s *DescribeAclBackupListResponseBodyBackups) GetDescription() *string {
	return s.Description
}

func (s *DescribeAclBackupListResponseBodyBackups) SetAclCount(v int32) *DescribeAclBackupListResponseBodyBackups {
	s.AclCount = &v
	return s
}

func (s *DescribeAclBackupListResponseBodyBackups) SetBackUpTime(v int64) *DescribeAclBackupListResponseBodyBackups {
	s.BackUpTime = &v
	return s
}

func (s *DescribeAclBackupListResponseBodyBackups) SetDescription(v string) *DescribeAclBackupListResponseBodyBackups {
	s.Description = &v
	return s
}

func (s *DescribeAclBackupListResponseBodyBackups) Validate() error {
	return dara.Validate(s)
}
