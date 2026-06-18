// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTaskResponseBody
	GetCode() *string
	SetData(v *ListTaskResponseBodyData) *ListTaskResponseBody
	GetData() *ListTaskResponseBodyData
	SetMessage(v string) *ListTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTaskResponseBody
	GetSuccess() *bool
}

type ListTaskResponseBody struct {
	// Request status code. A return value of OK indicates that the request succeeded.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// List of job data.
	Data *ListTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API was invoked successfully. Valid values:
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTaskResponseBody) GetData() *ListTaskResponseBodyData {
	return s.Data
}

func (s *ListTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTaskResponseBody) SetCode(v string) *ListTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ListTaskResponseBody) SetData(v *ListTaskResponseBodyData) *ListTaskResponseBody {
	s.Data = v
	return s
}

func (s *ListTaskResponseBody) SetMessage(v string) *ListTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ListTaskResponseBody) SetRequestId(v string) *ListTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskResponseBody) SetSuccess(v bool) *ListTaskResponseBody {
	s.Success = &v
	return s
}

func (s *ListTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTaskResponseBodyData struct {
	// Current page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// List of job information.
	Record []*ListTaskResponseBodyDataRecord `json:"Record,omitempty" xml:"Record,omitempty" type:"Repeated"`
	// Total number of jobs.
	//
	// example:
	//
	// 50
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTaskResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListTaskResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTaskResponseBodyData) GetRecord() []*ListTaskResponseBodyDataRecord {
	return s.Record
}

func (s *ListTaskResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListTaskResponseBodyData) SetPageNo(v int64) *ListTaskResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListTaskResponseBodyData) SetPageSize(v int64) *ListTaskResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTaskResponseBodyData) SetRecord(v []*ListTaskResponseBodyDataRecord) *ListTaskResponseBodyData {
	s.Record = v
	return s
}

func (s *ListTaskResponseBodyData) SetTotal(v int64) *ListTaskResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListTaskResponseBodyData) Validate() error {
	if s.Record != nil {
		for _, item := range s.Record {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTaskResponseBodyDataRecord struct {
	// Number of completed calls.
	//
	// example:
	//
	// 1
	CompleteCount *int32 `json:"CompleteCount,omitempty" xml:"CompleteCount,omitempty"`
	// Job start time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1618477232000
	FireTime *string `json:"FireTime,omitempty" xml:"FireTime,omitempty"`
	// Creation Time of the job. Format: UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1618477232000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The unique job ID for the robot calling task.
	//
	// example:
	//
	// 12****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the specified robot, which is the script ID.
	//
	// example:
	//
	// 12****
	RobotId *int64 `json:"RobotId,omitempty" xml:"RobotId,omitempty"`
	// Robot Name.
	//
	// example:
	//
	// 机器人
	RobotName *string `json:"RobotName,omitempty" xml:"RobotName,omitempty"`
	// Task Status.
	//
	// example:
	//
	// RELEASE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Task Name.
	//
	// example:
	//
	// 测试任务
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// Total number of processed calls.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTaskResponseBodyDataRecord) String() string {
	return dara.Prettify(s)
}

func (s ListTaskResponseBodyDataRecord) GoString() string {
	return s.String()
}

func (s *ListTaskResponseBodyDataRecord) GetCompleteCount() *int32 {
	return s.CompleteCount
}

func (s *ListTaskResponseBodyDataRecord) GetFireTime() *string {
	return s.FireTime
}

func (s *ListTaskResponseBodyDataRecord) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListTaskResponseBodyDataRecord) GetId() *int64 {
	return s.Id
}

func (s *ListTaskResponseBodyDataRecord) GetRobotId() *int64 {
	return s.RobotId
}

func (s *ListTaskResponseBodyDataRecord) GetRobotName() *string {
	return s.RobotName
}

func (s *ListTaskResponseBodyDataRecord) GetStatus() *string {
	return s.Status
}

func (s *ListTaskResponseBodyDataRecord) GetTaskName() *string {
	return s.TaskName
}

func (s *ListTaskResponseBodyDataRecord) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTaskResponseBodyDataRecord) SetCompleteCount(v int32) *ListTaskResponseBodyDataRecord {
	s.CompleteCount = &v
	return s
}

func (s *ListTaskResponseBodyDataRecord) SetFireTime(v string) *ListTaskResponseBodyDataRecord {
	s.FireTime = &v
	return s
}

func (s *ListTaskResponseBodyDataRecord) SetGmtCreate(v string) *ListTaskResponseBodyDataRecord {
	s.GmtCreate = &v
	return s
}

func (s *ListTaskResponseBodyDataRecord) SetId(v int64) *ListTaskResponseBodyDataRecord {
	s.Id = &v
	return s
}

func (s *ListTaskResponseBodyDataRecord) SetRobotId(v int64) *ListTaskResponseBodyDataRecord {
	s.RobotId = &v
	return s
}

func (s *ListTaskResponseBodyDataRecord) SetRobotName(v string) *ListTaskResponseBodyDataRecord {
	s.RobotName = &v
	return s
}

func (s *ListTaskResponseBodyDataRecord) SetStatus(v string) *ListTaskResponseBodyDataRecord {
	s.Status = &v
	return s
}

func (s *ListTaskResponseBodyDataRecord) SetTaskName(v string) *ListTaskResponseBodyDataRecord {
	s.TaskName = &v
	return s
}

func (s *ListTaskResponseBodyDataRecord) SetTotalCount(v int32) *ListTaskResponseBodyDataRecord {
	s.TotalCount = &v
	return s
}

func (s *ListTaskResponseBodyDataRecord) Validate() error {
	return dara.Validate(s)
}
