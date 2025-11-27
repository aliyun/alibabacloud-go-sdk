// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModifyParameterLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeModifyParameterLogResponseBody
	GetDBInstanceId() *string
	SetEngine(v string) *DescribeModifyParameterLogResponseBody
	GetEngine() *string
	SetEngineVersion(v string) *DescribeModifyParameterLogResponseBody
	GetEngineVersion() *string
	SetItems(v *DescribeModifyParameterLogResponseBodyItems) *DescribeModifyParameterLogResponseBody
	GetItems() *DescribeModifyParameterLogResponseBodyItems
	SetPageNumber(v int32) *DescribeModifyParameterLogResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeModifyParameterLogResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeModifyParameterLogResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeModifyParameterLogResponseBody
	GetTotalRecordCount() *int32
}

type DescribeModifyParameterLogResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The database engine of the instance.
	//
	// example:
	//
	// mysql
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version of the instance.
	//
	// example:
	//
	// 5.6
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The log entries.
	Items *DescribeModifyParameterLogResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C8E88DED-533F-4B3C-9207-731FBF394CCA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeModifyParameterLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeModifyParameterLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeModifyParameterLogResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeModifyParameterLogResponseBody) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeModifyParameterLogResponseBody) GetItems() *DescribeModifyParameterLogResponseBodyItems {
	return s.Items
}

func (s *DescribeModifyParameterLogResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeModifyParameterLogResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeModifyParameterLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeModifyParameterLogResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeModifyParameterLogResponseBody) SetDBInstanceId(v string) *DescribeModifyParameterLogResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBody) SetEngine(v string) *DescribeModifyParameterLogResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBody) SetEngineVersion(v string) *DescribeModifyParameterLogResponseBody {
	s.EngineVersion = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBody) SetItems(v *DescribeModifyParameterLogResponseBodyItems) *DescribeModifyParameterLogResponseBody {
	s.Items = v
	return s
}

func (s *DescribeModifyParameterLogResponseBody) SetPageNumber(v int32) *DescribeModifyParameterLogResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBody) SetPageRecordCount(v int32) *DescribeModifyParameterLogResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBody) SetRequestId(v string) *DescribeModifyParameterLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBody) SetTotalRecordCount(v int32) *DescribeModifyParameterLogResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeModifyParameterLogResponseBodyItems struct {
	ParameterChangeLog []*DescribeModifyParameterLogResponseBodyItemsParameterChangeLog `json:"ParameterChangeLog,omitempty" xml:"ParameterChangeLog,omitempty" type:"Repeated"`
}

func (s DescribeModifyParameterLogResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeModifyParameterLogResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogResponseBodyItems) GetParameterChangeLog() []*DescribeModifyParameterLogResponseBodyItemsParameterChangeLog {
	return s.ParameterChangeLog
}

func (s *DescribeModifyParameterLogResponseBodyItems) SetParameterChangeLog(v []*DescribeModifyParameterLogResponseBodyItemsParameterChangeLog) *DescribeModifyParameterLogResponseBodyItems {
	s.ParameterChangeLog = v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyItems) Validate() error {
	if s.ParameterChangeLog != nil {
		for _, item := range s.ParameterChangeLog {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeModifyParameterLogResponseBodyItemsParameterChangeLog struct {
	// The time when the parameter was modified. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1584076066000
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The new value of the parameter.
	//
	// example:
	//
	// 3
	NewParameterValue *string `json:"NewParameterValue,omitempty" xml:"NewParameterValue,omitempty"`
	// The original value of the parameter.
	//
	// example:
	//
	// 8
	OldParameterValue *string `json:"OldParameterValue,omitempty" xml:"OldParameterValue,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// innodb_stats_sample_pages
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The status. Valid values:
	//
	// 	- **Applied:*	- The new value has taken effect.
	//
	// 	- **Syncing:*	- The new value is being applied and has not taken effect.
	//
	// example:
	//
	// Syncing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeModifyParameterLogResponseBodyItemsParameterChangeLog) String() string {
	return dara.Prettify(s)
}

func (s DescribeModifyParameterLogResponseBodyItemsParameterChangeLog) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogResponseBodyItemsParameterChangeLog) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeModifyParameterLogResponseBodyItemsParameterChangeLog) GetNewParameterValue() *string {
	return s.NewParameterValue
}

func (s *DescribeModifyParameterLogResponseBodyItemsParameterChangeLog) GetOldParameterValue() *string {
	return s.OldParameterValue
}

func (s *DescribeModifyParameterLogResponseBodyItemsParameterChangeLog) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeModifyParameterLogResponseBodyItemsParameterChangeLog) GetStatus() *string {
	return s.Status
}

func (s *DescribeModifyParameterLogResponseBodyItemsParameterChangeLog) SetModifyTime(v string) *DescribeModifyParameterLogResponseBodyItemsParameterChangeLog {
	s.ModifyTime = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyItemsParameterChangeLog) SetNewParameterValue(v string) *DescribeModifyParameterLogResponseBodyItemsParameterChangeLog {
	s.NewParameterValue = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyItemsParameterChangeLog) SetOldParameterValue(v string) *DescribeModifyParameterLogResponseBodyItemsParameterChangeLog {
	s.OldParameterValue = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyItemsParameterChangeLog) SetParameterName(v string) *DescribeModifyParameterLogResponseBodyItemsParameterChangeLog {
	s.ParameterName = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyItemsParameterChangeLog) SetStatus(v string) *DescribeModifyParameterLogResponseBodyItemsParameterChangeLog {
	s.Status = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyItemsParameterChangeLog) Validate() error {
	return dara.Validate(s)
}
