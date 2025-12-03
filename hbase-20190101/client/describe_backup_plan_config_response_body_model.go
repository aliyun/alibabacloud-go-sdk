// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPlanConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFullBackupCycle(v int32) *DescribeBackupPlanConfigResponseBody
	GetFullBackupCycle() *int32
	SetMinHFileBackupCount(v int32) *DescribeBackupPlanConfigResponseBody
	GetMinHFileBackupCount() *int32
	SetNextFullBackupDate(v string) *DescribeBackupPlanConfigResponseBody
	GetNextFullBackupDate() *string
	SetRequestId(v string) *DescribeBackupPlanConfigResponseBody
	GetRequestId() *string
	SetTables(v *DescribeBackupPlanConfigResponseBodyTables) *DescribeBackupPlanConfigResponseBody
	GetTables() *DescribeBackupPlanConfigResponseBodyTables
}

type DescribeBackupPlanConfigResponseBody struct {
	// example:
	//
	// 7
	FullBackupCycle *int32 `json:"FullBackupCycle,omitempty" xml:"FullBackupCycle,omitempty"`
	// example:
	//
	// 3
	MinHFileBackupCount *int32 `json:"MinHFileBackupCount,omitempty" xml:"MinHFileBackupCount,omitempty"`
	// example:
	//
	// 2020-11-09T18:00:00Z
	NextFullBackupDate *string `json:"NextFullBackupDate,omitempty" xml:"NextFullBackupDate,omitempty"`
	// example:
	//
	// 33A23201-6038-4A6A-B76A-61047EA04E6A
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tables    *DescribeBackupPlanConfigResponseBodyTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Struct"`
}

func (s DescribeBackupPlanConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlanConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanConfigResponseBody) GetFullBackupCycle() *int32 {
	return s.FullBackupCycle
}

func (s *DescribeBackupPlanConfigResponseBody) GetMinHFileBackupCount() *int32 {
	return s.MinHFileBackupCount
}

func (s *DescribeBackupPlanConfigResponseBody) GetNextFullBackupDate() *string {
	return s.NextFullBackupDate
}

func (s *DescribeBackupPlanConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupPlanConfigResponseBody) GetTables() *DescribeBackupPlanConfigResponseBodyTables {
	return s.Tables
}

func (s *DescribeBackupPlanConfigResponseBody) SetFullBackupCycle(v int32) *DescribeBackupPlanConfigResponseBody {
	s.FullBackupCycle = &v
	return s
}

func (s *DescribeBackupPlanConfigResponseBody) SetMinHFileBackupCount(v int32) *DescribeBackupPlanConfigResponseBody {
	s.MinHFileBackupCount = &v
	return s
}

func (s *DescribeBackupPlanConfigResponseBody) SetNextFullBackupDate(v string) *DescribeBackupPlanConfigResponseBody {
	s.NextFullBackupDate = &v
	return s
}

func (s *DescribeBackupPlanConfigResponseBody) SetRequestId(v string) *DescribeBackupPlanConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPlanConfigResponseBody) SetTables(v *DescribeBackupPlanConfigResponseBodyTables) *DescribeBackupPlanConfigResponseBody {
	s.Tables = v
	return s
}

func (s *DescribeBackupPlanConfigResponseBody) Validate() error {
	if s.Tables != nil {
		if err := s.Tables.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupPlanConfigResponseBodyTables struct {
	Table []*string `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeBackupPlanConfigResponseBodyTables) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlanConfigResponseBodyTables) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanConfigResponseBodyTables) GetTable() []*string {
	return s.Table
}

func (s *DescribeBackupPlanConfigResponseBodyTables) SetTable(v []*string) *DescribeBackupPlanConfigResponseBodyTables {
	s.Table = v
	return s
}

func (s *DescribeBackupPlanConfigResponseBodyTables) Validate() error {
	return dara.Validate(s)
}
