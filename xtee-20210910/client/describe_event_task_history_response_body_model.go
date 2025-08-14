// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventTaskHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeEventTaskHistoryResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeEventTaskHistoryResponseBodyResultObject) *DescribeEventTaskHistoryResponseBody
	GetResultObject() []*DescribeEventTaskHistoryResponseBodyResultObject
}

type DescribeEventTaskHistoryResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject []*DescribeEventTaskHistoryResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeEventTaskHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventTaskHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventTaskHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventTaskHistoryResponseBody) GetResultObject() []*DescribeEventTaskHistoryResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeEventTaskHistoryResponseBody) SetRequestId(v string) *DescribeEventTaskHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventTaskHistoryResponseBody) SetResultObject(v []*DescribeEventTaskHistoryResponseBodyResultObject) *DescribeEventTaskHistoryResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeEventTaskHistoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEventTaskHistoryResponseBodyResultObject struct {
	// Task code.
	//
	// example:
	//
	// de_aoxcdy9473
	TaskCode *string `json:"taskCode,omitempty" xml:"taskCode,omitempty"`
	// Task name
	//
	// example:
	//
	// 仿真任务
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// Task status.
	//
	// example:
	//
	// SUCCESSFUL
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
	// OSS download URL
	//
	// example:
	//
	// https://xxxxx-oss-xxxxx.xxxxxx.aliyuncs.com/xx/xx/xxx/xxxxxx.csv?Expires=1753433384&OSSAccessKeyId=xxxxxxxxx&Signature=%2F%xxxxxxxxxxxx%3D
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s DescribeEventTaskHistoryResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventTaskHistoryResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeEventTaskHistoryResponseBodyResultObject) GetTaskCode() *string {
	return s.TaskCode
}

func (s *DescribeEventTaskHistoryResponseBodyResultObject) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeEventTaskHistoryResponseBodyResultObject) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeEventTaskHistoryResponseBodyResultObject) GetUrl() *string {
	return s.Url
}

func (s *DescribeEventTaskHistoryResponseBodyResultObject) SetTaskCode(v string) *DescribeEventTaskHistoryResponseBodyResultObject {
	s.TaskCode = &v
	return s
}

func (s *DescribeEventTaskHistoryResponseBodyResultObject) SetTaskName(v string) *DescribeEventTaskHistoryResponseBodyResultObject {
	s.TaskName = &v
	return s
}

func (s *DescribeEventTaskHistoryResponseBodyResultObject) SetTaskStatus(v string) *DescribeEventTaskHistoryResponseBodyResultObject {
	s.TaskStatus = &v
	return s
}

func (s *DescribeEventTaskHistoryResponseBodyResultObject) SetUrl(v string) *DescribeEventTaskHistoryResponseBodyResultObject {
	s.Url = &v
	return s
}

func (s *DescribeEventTaskHistoryResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
