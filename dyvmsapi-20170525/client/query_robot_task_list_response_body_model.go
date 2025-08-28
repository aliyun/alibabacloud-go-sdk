// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRobotTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryRobotTaskListResponseBody
	GetCode() *string
	SetData(v string) *QueryRobotTaskListResponseBody
	GetData() *string
	SetMessage(v string) *QueryRobotTaskListResponseBody
	GetMessage() *string
	SetPageNo(v string) *QueryRobotTaskListResponseBody
	GetPageNo() *string
	SetPageSize(v string) *QueryRobotTaskListResponseBody
	GetPageSize() *string
	SetRequestId(v string) *QueryRobotTaskListResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *QueryRobotTaskListResponseBody
	GetTotalCount() *string
}

type QueryRobotTaskListResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The robocall tasks, in the JSON format.
	//
	// 	- **id**: the unique ID of the robocall task.
	//
	// 	- **taskName**: the task name.
	//
	// 	- **robotId**: the robot ID.
	//
	// 	- **robotName**: the robot name.
	//
	// 	- **status**: the task state.
	//
	// 	- **scheduleType**: the scheduling type. Valid values: **SINGLE*	- and **ORDER**. The value SINGLE indicates that the task was started immediately after it was created. The value ORDER indicates that the task was started at a scheduled time.
	//
	// 	- **createTime**: the time when the task was created, in the yyyy.MM.dd HH:mm:ss format.
	//
	// 	- **completeTime**: the time when the task was completed, in the yyyy.MM.dd HH:mm:ss format.
	//
	// 	- **fireTime**: the time when the task was started, in the yyyy.MM.dd HH:mm:ss format.
	//
	// 	- **totalCount**: the total number of calls processed.
	//
	// 	- **finishCount**: the number of calls completed.
	//
	// example:
	//
	// [ {"id": 1045001, "taskName": "Test Template", "robotId": 1000000075003, "robotName": "robot", "status": "INIT","scheduleType": "SINGLE", "createTime": "2019.06.14 14:55:23", "completeTime": "2019.06.14 14:55:23", "fireTime": "2019.06.14 14:55:23", "totalCount": 1000, "finishCount": 998} ]
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of tasks.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryRobotTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRobotTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryRobotTaskListResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryRobotTaskListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryRobotTaskListResponseBody) GetPageNo() *string {
	return s.PageNo
}

func (s *QueryRobotTaskListResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *QueryRobotTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRobotTaskListResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *QueryRobotTaskListResponseBody) SetCode(v string) *QueryRobotTaskListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRobotTaskListResponseBody) SetData(v string) *QueryRobotTaskListResponseBody {
	s.Data = &v
	return s
}

func (s *QueryRobotTaskListResponseBody) SetMessage(v string) *QueryRobotTaskListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRobotTaskListResponseBody) SetPageNo(v string) *QueryRobotTaskListResponseBody {
	s.PageNo = &v
	return s
}

func (s *QueryRobotTaskListResponseBody) SetPageSize(v string) *QueryRobotTaskListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryRobotTaskListResponseBody) SetRequestId(v string) *QueryRobotTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRobotTaskListResponseBody) SetTotalCount(v string) *QueryRobotTaskListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryRobotTaskListResponseBody) Validate() error {
	return dara.Validate(s)
}
