// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSelectItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSelectItemResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeSelectItemResponseBodyResultObject) *DescribeSelectItemResponseBody
	GetResultObject() *DescribeSelectItemResponseBodyResultObject
}

type DescribeSelectItemResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject *DescribeSelectItemResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeSelectItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSelectItemResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSelectItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSelectItemResponseBody) GetResultObject() *DescribeSelectItemResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeSelectItemResponseBody) SetRequestId(v string) *DescribeSelectItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSelectItemResponseBody) SetResultObject(v *DescribeSelectItemResponseBodyResultObject) *DescribeSelectItemResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeSelectItemResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSelectItemResponseBodyResultObject struct {
	// Monitoring status list.
	MonitorStatusList []*string `json:"monitorStatusList,omitempty" xml:"monitorStatusList,omitempty" type:"Repeated"`
	// Task ID list.
	TaskIdList []*string `json:"taskIdList,omitempty" xml:"taskIdList,omitempty" type:"Repeated"`
}

func (s DescribeSelectItemResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeSelectItemResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeSelectItemResponseBodyResultObject) GetMonitorStatusList() []*string {
	return s.MonitorStatusList
}

func (s *DescribeSelectItemResponseBodyResultObject) GetTaskIdList() []*string {
	return s.TaskIdList
}

func (s *DescribeSelectItemResponseBodyResultObject) SetMonitorStatusList(v []*string) *DescribeSelectItemResponseBodyResultObject {
	s.MonitorStatusList = v
	return s
}

func (s *DescribeSelectItemResponseBodyResultObject) SetTaskIdList(v []*string) *DescribeSelectItemResponseBodyResultObject {
	s.TaskIdList = v
	return s
}

func (s *DescribeSelectItemResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
