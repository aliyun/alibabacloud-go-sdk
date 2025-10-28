// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoIncrementUsageStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetAutoIncrementUsageStatisticResponseBody
	GetCode() *int64
	SetData(v *GetAutoIncrementUsageStatisticResponseBodyData) *GetAutoIncrementUsageStatisticResponseBody
	GetData() *GetAutoIncrementUsageStatisticResponseBodyData
	SetMessage(v string) *GetAutoIncrementUsageStatisticResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAutoIncrementUsageStatisticResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAutoIncrementUsageStatisticResponseBody
	GetSuccess() *bool
}

type GetAutoIncrementUsageStatisticResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetAutoIncrementUsageStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request is successful, **Successful*	- is returned. Otherwise, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0A74B755-98B7-59DB-8724-1321B394****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAutoIncrementUsageStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAutoIncrementUsageStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutoIncrementUsageStatisticResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetAutoIncrementUsageStatisticResponseBody) GetData() *GetAutoIncrementUsageStatisticResponseBodyData {
	return s.Data
}

func (s *GetAutoIncrementUsageStatisticResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAutoIncrementUsageStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAutoIncrementUsageStatisticResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAutoIncrementUsageStatisticResponseBody) SetCode(v int64) *GetAutoIncrementUsageStatisticResponseBody {
	s.Code = &v
	return s
}

func (s *GetAutoIncrementUsageStatisticResponseBody) SetData(v *GetAutoIncrementUsageStatisticResponseBodyData) *GetAutoIncrementUsageStatisticResponseBody {
	s.Data = v
	return s
}

func (s *GetAutoIncrementUsageStatisticResponseBody) SetMessage(v string) *GetAutoIncrementUsageStatisticResponseBody {
	s.Message = &v
	return s
}

func (s *GetAutoIncrementUsageStatisticResponseBody) SetRequestId(v string) *GetAutoIncrementUsageStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutoIncrementUsageStatisticResponseBody) SetSuccess(v bool) *GetAutoIncrementUsageStatisticResponseBody {
	s.Success = &v
	return s
}

func (s *GetAutoIncrementUsageStatisticResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAutoIncrementUsageStatisticResponseBodyData struct {
	// The usage details of auto-increment IDs.
	AutoIncrementUsageList []*GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList `json:"AutoIncrementUsageList,omitempty" xml:"AutoIncrementUsageList,omitempty" type:"Repeated"`
	// The error message returned if the task fails.
	//
	// example:
	//
	// the given database name list invalid, none of the database names in the list exist on the instance
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// Indicates whether the task is complete. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Finish *bool `json:"Finish,omitempty" xml:"Finish,omitempty"`
	// The task status. Valid values:
	//
	// 	- **INIT**: The task is being initialized.
	//
	// 	- **RUNNING**: The task is being executed.
	//
	// 	- **SUCCESS**: The task succeeds.
	//
	// 	- **FAIL**: The task fails.
	//
	// example:
	//
	// INIT
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The time when the request was made. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1697183353000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetAutoIncrementUsageStatisticResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAutoIncrementUsageStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAutoIncrementUsageStatisticResponseBodyData) GetAutoIncrementUsageList() []*GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList {
	return s.AutoIncrementUsageList
}

func (s *GetAutoIncrementUsageStatisticResponseBodyData) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *GetAutoIncrementUsageStatisticResponseBodyData) GetFinish() *bool {
	return s.Finish
}

func (s *GetAutoIncrementUsageStatisticResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetAutoIncrementUsageStatisticResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetAutoIncrementUsageStatisticResponseBodyData) SetAutoIncrementUsageList(v []*GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList) *GetAutoIncrementUsageStatisticResponseBodyData {
	s.AutoIncrementUsageList = v
	return s
}

func (s *GetAutoIncrementUsageStatisticResponseBodyData) SetErrorInfo(v string) *GetAutoIncrementUsageStatisticResponseBodyData {
	s.ErrorInfo = &v
	return s
}

func (s *GetAutoIncrementUsageStatisticResponseBodyData) SetFinish(v bool) *GetAutoIncrementUsageStatisticResponseBodyData {
	s.Finish = &v
	return s
}

func (s *GetAutoIncrementUsageStatisticResponseBodyData) SetTaskStatus(v string) *GetAutoIncrementUsageStatisticResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *GetAutoIncrementUsageStatisticResponseBodyData) SetTimestamp(v int64) *GetAutoIncrementUsageStatisticResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *GetAutoIncrementUsageStatisticResponseBodyData) Validate() error {
	if s.AutoIncrementUsageList != nil {
		for _, item := range s.AutoIncrementUsageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList struct {
	// The latest auto-increment ID.
	//
	// example:
	//
	// 2147483647
	AutoIncrementCurrentValue *int64 `json:"AutoIncrementCurrentValue,omitempty" xml:"AutoIncrementCurrentValue,omitempty"`
	// The usage ratio of auto-increment IDs.
	//
	// example:
	//
	// 1
	AutoIncrementRatio *float64 `json:"AutoIncrementRatio,omitempty" xml:"AutoIncrementRatio,omitempty"`
	// The column name.
	//
	// example:
	//
	// id
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The database name.
	//
	// example:
	//
	// db01
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The maximum auto-increment ID that is supported by the current data type.
	//
	// example:
	//
	// 2147483647
	MaximumValue *int64 `json:"MaximumValue,omitempty" xml:"MaximumValue,omitempty"`
	// The table name.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList) String() string {
	return dara.Prettify(s)
}

func (s GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList) GoString() string {
	return s.String()
}

func (s *GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList) GetAutoIncrementCurrentValue() *int64 {
	return s.AutoIncrementCurrentValue
}

func (s *GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList) GetAutoIncrementRatio() *float64 {
	return s.AutoIncrementRatio
}

func (s *GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList) GetColumnName() *string {
	return s.ColumnName
}

func (s *GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList) GetDbName() *string {
	return s.DbName
}

func (s *GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList) GetMaximumValue() *int64 {
	return s.MaximumValue
}

func (s *GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList) GetTableName() *string {
	return s.TableName
}

func (s *GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList) SetAutoIncrementCurrentValue(v int64) *GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList {
	s.AutoIncrementCurrentValue = &v
	return s
}

func (s *GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList) SetAutoIncrementRatio(v float64) *GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList {
	s.AutoIncrementRatio = &v
	return s
}

func (s *GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList) SetColumnName(v string) *GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList {
	s.ColumnName = &v
	return s
}

func (s *GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList) SetDbName(v string) *GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList {
	s.DbName = &v
	return s
}

func (s *GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList) SetMaximumValue(v int64) *GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList {
	s.MaximumValue = &v
	return s
}

func (s *GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList) SetTableName(v string) *GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList {
	s.TableName = &v
	return s
}

func (s *GetAutoIncrementUsageStatisticResponseBodyDataAutoIncrementUsageList) Validate() error {
	return dara.Validate(s)
}
