// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListTaskFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudListTaskFileResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudListTaskFileResponseBody
	GetCode() *string
	SetData(v *CloudListTaskFileResponseBodyData) *CloudListTaskFileResponseBody
	GetData() *CloudListTaskFileResponseBodyData
	SetMessage(v string) *CloudListTaskFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudListTaskFileResponseBody
	GetRequestId() *string
}

type CloudListTaskFileResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudListTaskFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6086693B-2250-17CE-A41F-06259AB6DB1B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudListTaskFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudListTaskFileResponseBody) GoString() string {
	return s.String()
}

func (s *CloudListTaskFileResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudListTaskFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudListTaskFileResponseBody) GetData() *CloudListTaskFileResponseBodyData {
	return s.Data
}

func (s *CloudListTaskFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudListTaskFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudListTaskFileResponseBody) SetAccessDeniedDetail(v string) *CloudListTaskFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudListTaskFileResponseBody) SetCode(v string) *CloudListTaskFileResponseBody {
	s.Code = &v
	return s
}

func (s *CloudListTaskFileResponseBody) SetData(v *CloudListTaskFileResponseBodyData) *CloudListTaskFileResponseBody {
	s.Data = v
	return s
}

func (s *CloudListTaskFileResponseBody) SetMessage(v string) *CloudListTaskFileResponseBody {
	s.Message = &v
	return s
}

func (s *CloudListTaskFileResponseBody) SetRequestId(v string) *CloudListTaskFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudListTaskFileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudListTaskFileResponseBodyData struct {
	// 任务批次信息
	TaskFileList []*CloudListTaskFileResponseBodyDataTaskFileList `json:"TaskFileList,omitempty" xml:"TaskFileList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CloudListTaskFileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudListTaskFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudListTaskFileResponseBodyData) GetTaskFileList() []*CloudListTaskFileResponseBodyDataTaskFileList {
	return s.TaskFileList
}

func (s *CloudListTaskFileResponseBodyData) GetTotalCount() *string {
	return s.TotalCount
}

func (s *CloudListTaskFileResponseBodyData) SetTaskFileList(v []*CloudListTaskFileResponseBodyDataTaskFileList) *CloudListTaskFileResponseBodyData {
	s.TaskFileList = v
	return s
}

func (s *CloudListTaskFileResponseBodyData) SetTotalCount(v string) *CloudListTaskFileResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *CloudListTaskFileResponseBodyData) Validate() error {
	if s.TaskFileList != nil {
		for _, item := range s.TaskFileList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CloudListTaskFileResponseBodyDataTaskFileList struct {
	// 创建批次时间，格式为： yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-04-20 10:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 企业编号
	//
	// example:
	//
	// 7000002
	EnterpriseId *string `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 批次Id
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 导入条数
	//
	// example:
	//
	// 3
	ImportCount *string `json:"ImportCount,omitempty" xml:"ImportCount,omitempty"`
	// 批次名称
	//
	// example:
	//
	// name2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 优先级，默认为0，值越大优先级越高
	//
	// example:
	//
	// 0
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// 重复条数
	//
	// example:
	//
	// 0
	RepeatCount *string `json:"RepeatCount,omitempty" xml:"RepeatCount,omitempty"`
	// 黑名单条数
	//
	// example:
	//
	// 0
	RestrictCount *string `json:"RestrictCount,omitempty" xml:"RestrictCount,omitempty"`
	// 批次状态，0：未导入，1：导入完成 ，2：加载到缓存，3：呼叫完，4：已冻结
	//
	// example:
	//
	// 0
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 任务Id
	//
	// example:
	//
	// 27
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CloudListTaskFileResponseBodyDataTaskFileList) String() string {
	return dara.Prettify(s)
}

func (s CloudListTaskFileResponseBodyDataTaskFileList) GoString() string {
	return s.String()
}

func (s *CloudListTaskFileResponseBodyDataTaskFileList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudListTaskFileResponseBodyDataTaskFileList) GetEnterpriseId() *string {
	return s.EnterpriseId
}

func (s *CloudListTaskFileResponseBodyDataTaskFileList) GetId() *string {
	return s.Id
}

func (s *CloudListTaskFileResponseBodyDataTaskFileList) GetImportCount() *string {
	return s.ImportCount
}

func (s *CloudListTaskFileResponseBodyDataTaskFileList) GetName() *string {
	return s.Name
}

func (s *CloudListTaskFileResponseBodyDataTaskFileList) GetPriority() *string {
	return s.Priority
}

func (s *CloudListTaskFileResponseBodyDataTaskFileList) GetRepeatCount() *string {
	return s.RepeatCount
}

func (s *CloudListTaskFileResponseBodyDataTaskFileList) GetRestrictCount() *string {
	return s.RestrictCount
}

func (s *CloudListTaskFileResponseBodyDataTaskFileList) GetStatus() *string {
	return s.Status
}

func (s *CloudListTaskFileResponseBodyDataTaskFileList) GetTaskId() *string {
	return s.TaskId
}

func (s *CloudListTaskFileResponseBodyDataTaskFileList) SetCreateTime(v string) *CloudListTaskFileResponseBodyDataTaskFileList {
	s.CreateTime = &v
	return s
}

func (s *CloudListTaskFileResponseBodyDataTaskFileList) SetEnterpriseId(v string) *CloudListTaskFileResponseBodyDataTaskFileList {
	s.EnterpriseId = &v
	return s
}

func (s *CloudListTaskFileResponseBodyDataTaskFileList) SetId(v string) *CloudListTaskFileResponseBodyDataTaskFileList {
	s.Id = &v
	return s
}

func (s *CloudListTaskFileResponseBodyDataTaskFileList) SetImportCount(v string) *CloudListTaskFileResponseBodyDataTaskFileList {
	s.ImportCount = &v
	return s
}

func (s *CloudListTaskFileResponseBodyDataTaskFileList) SetName(v string) *CloudListTaskFileResponseBodyDataTaskFileList {
	s.Name = &v
	return s
}

func (s *CloudListTaskFileResponseBodyDataTaskFileList) SetPriority(v string) *CloudListTaskFileResponseBodyDataTaskFileList {
	s.Priority = &v
	return s
}

func (s *CloudListTaskFileResponseBodyDataTaskFileList) SetRepeatCount(v string) *CloudListTaskFileResponseBodyDataTaskFileList {
	s.RepeatCount = &v
	return s
}

func (s *CloudListTaskFileResponseBodyDataTaskFileList) SetRestrictCount(v string) *CloudListTaskFileResponseBodyDataTaskFileList {
	s.RestrictCount = &v
	return s
}

func (s *CloudListTaskFileResponseBodyDataTaskFileList) SetStatus(v string) *CloudListTaskFileResponseBodyDataTaskFileList {
	s.Status = &v
	return s
}

func (s *CloudListTaskFileResponseBodyDataTaskFileList) SetTaskId(v string) *CloudListTaskFileResponseBodyDataTaskFileList {
	s.TaskId = &v
	return s
}

func (s *CloudListTaskFileResponseBodyDataTaskFileList) Validate() error {
	return dara.Validate(s)
}
