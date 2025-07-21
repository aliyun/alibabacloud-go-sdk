// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskByParamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *QueryTaskByParamResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryTaskByParamResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryTaskByParamResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryTaskByParamResponseBody
	GetTotalCount() *int32
	SetData(v *QueryTaskByParamResponseBodyData) *QueryTaskByParamResponseBody
	GetData() *QueryTaskByParamResponseBodyData
}

type QueryTaskByParamResponseBody struct {
	// Current page number
	//
	// example:
	//
	// 3
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total count
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Returned results
	Data *QueryTaskByParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s QueryTaskByParamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskByParamResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskByParamResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryTaskByParamResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryTaskByParamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTaskByParamResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryTaskByParamResponseBody) GetData() *QueryTaskByParamResponseBodyData {
	return s.Data
}

func (s *QueryTaskByParamResponseBody) SetPageNumber(v int32) *QueryTaskByParamResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryTaskByParamResponseBody) SetPageSize(v int32) *QueryTaskByParamResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTaskByParamResponseBody) SetRequestId(v string) *QueryTaskByParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTaskByParamResponseBody) SetTotalCount(v int32) *QueryTaskByParamResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryTaskByParamResponseBody) SetData(v *QueryTaskByParamResponseBodyData) *QueryTaskByParamResponseBody {
	s.Data = v
	return s
}

func (s *QueryTaskByParamResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryTaskByParamResponseBodyData struct {
	Task []*QueryTaskByParamResponseBodyDataTask `json:"task,omitempty" xml:"task,omitempty" type:"Repeated"`
}

func (s QueryTaskByParamResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskByParamResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTaskByParamResponseBodyData) GetTask() []*QueryTaskByParamResponseBodyDataTask {
	return s.Task
}

func (s *QueryTaskByParamResponseBodyData) SetTask(v []*QueryTaskByParamResponseBodyDataTask) *QueryTaskByParamResponseBodyData {
	s.Task = v
	return s
}

func (s *QueryTaskByParamResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryTaskByParamResponseBodyDataTask struct {
	// Address type, sending address: 1; random address: 0;
	//
	// example:
	//
	// 0
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// Creation time
	//
	// example:
	//
	// 2022-04-18T10:36Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// dedicated IP pool ID.
	//
	// example:
	//
	// xxx
	IpPoolId *string `json:"IpPoolId,omitempty" xml:"IpPoolId,omitempty"`
	// dedicated IP pool name.
	//
	// example:
	//
	// test
	IpPoolName *string `json:"IpPoolName,omitempty" xml:"IpPoolName,omitempty"`
	// Receiver\\"s name
	//
	// example:
	//
	// TKP000442-333
	ReceiversName *string `json:"ReceiversName,omitempty" xml:"ReceiversName,omitempty"`
	// Request count
	//
	// example:
	//
	// 1
	RequestCount *string `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	// Tag
	//
	// example:
	//
	// 202201
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// Task ID
	//
	// example:
	//
	// 1054296
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Task status, sent successfully: 1
	//
	// example:
	//
	// 1
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// Template name
	//
	// example:
	//
	// Short Simple
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Creation time in UTC format
	//
	// example:
	//
	// 1569734892
	UtcCreateTime *int64 `json:"UtcCreateTime,omitempty" xml:"UtcCreateTime,omitempty"`
}

func (s QueryTaskByParamResponseBodyDataTask) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskByParamResponseBodyDataTask) GoString() string {
	return s.String()
}

func (s *QueryTaskByParamResponseBodyDataTask) GetAddressType() *string {
	return s.AddressType
}

func (s *QueryTaskByParamResponseBodyDataTask) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryTaskByParamResponseBodyDataTask) GetIpPoolId() *string {
	return s.IpPoolId
}

func (s *QueryTaskByParamResponseBodyDataTask) GetIpPoolName() *string {
	return s.IpPoolName
}

func (s *QueryTaskByParamResponseBodyDataTask) GetReceiversName() *string {
	return s.ReceiversName
}

func (s *QueryTaskByParamResponseBodyDataTask) GetRequestCount() *string {
	return s.RequestCount
}

func (s *QueryTaskByParamResponseBodyDataTask) GetTagName() *string {
	return s.TagName
}

func (s *QueryTaskByParamResponseBodyDataTask) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryTaskByParamResponseBodyDataTask) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *QueryTaskByParamResponseBodyDataTask) GetTemplateName() *string {
	return s.TemplateName
}

func (s *QueryTaskByParamResponseBodyDataTask) GetUtcCreateTime() *int64 {
	return s.UtcCreateTime
}

func (s *QueryTaskByParamResponseBodyDataTask) SetAddressType(v string) *QueryTaskByParamResponseBodyDataTask {
	s.AddressType = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetCreateTime(v string) *QueryTaskByParamResponseBodyDataTask {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetIpPoolId(v string) *QueryTaskByParamResponseBodyDataTask {
	s.IpPoolId = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetIpPoolName(v string) *QueryTaskByParamResponseBodyDataTask {
	s.IpPoolName = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetReceiversName(v string) *QueryTaskByParamResponseBodyDataTask {
	s.ReceiversName = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetRequestCount(v string) *QueryTaskByParamResponseBodyDataTask {
	s.RequestCount = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetTagName(v string) *QueryTaskByParamResponseBodyDataTask {
	s.TagName = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetTaskId(v string) *QueryTaskByParamResponseBodyDataTask {
	s.TaskId = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetTaskStatus(v string) *QueryTaskByParamResponseBodyDataTask {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetTemplateName(v string) *QueryTaskByParamResponseBodyDataTask {
	s.TemplateName = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetUtcCreateTime(v int64) *QueryTaskByParamResponseBodyDataTask {
	s.UtcCreateTime = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) Validate() error {
	return dara.Validate(s)
}
