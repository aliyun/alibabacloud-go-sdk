// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTraceAppByPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageBean(v *SearchTraceAppByPageResponseBodyPageBean) *SearchTraceAppByPageResponseBody
	GetPageBean() *SearchTraceAppByPageResponseBodyPageBean
	SetRequestId(v string) *SearchTraceAppByPageResponseBody
	GetRequestId() *string
}

type SearchTraceAppByPageResponseBody struct {
	// The information about the array object.
	PageBean *SearchTraceAppByPageResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4B446DF2-3DDD-4B5B-8E3F-D5225120****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchTraceAppByPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchTraceAppByPageResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageResponseBody) GetPageBean() *SearchTraceAppByPageResponseBodyPageBean {
	return s.PageBean
}

func (s *SearchTraceAppByPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchTraceAppByPageResponseBody) SetPageBean(v *SearchTraceAppByPageResponseBodyPageBean) *SearchTraceAppByPageResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchTraceAppByPageResponseBody) SetRequestId(v string) *SearchTraceAppByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTraceAppByPageResponseBody) Validate() error {
	if s.PageBean != nil {
		if err := s.PageBean.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchTraceAppByPageResponseBodyPageBean struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about the monitoring task.
	TraceApps []*SearchTraceAppByPageResponseBodyPageBeanTraceApps `json:"TraceApps,omitempty" xml:"TraceApps,omitempty" type:"Repeated"`
}

func (s SearchTraceAppByPageResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s SearchTraceAppByPageResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageResponseBodyPageBean) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchTraceAppByPageResponseBodyPageBean) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchTraceAppByPageResponseBodyPageBean) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SearchTraceAppByPageResponseBodyPageBean) GetTraceApps() []*SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	return s.TraceApps
}

func (s *SearchTraceAppByPageResponseBodyPageBean) SetPageNumber(v int32) *SearchTraceAppByPageResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBean) SetPageSize(v int32) *SearchTraceAppByPageResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBean) SetTotalCount(v int32) *SearchTraceAppByPageResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBean) SetTraceApps(v []*SearchTraceAppByPageResponseBodyPageBeanTraceApps) *SearchTraceAppByPageResponseBodyPageBean {
	s.TraceApps = v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBean) Validate() error {
	if s.TraceApps != nil {
		for _, item := range s.TraceApps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchTraceAppByPageResponseBodyPageBeanTraceApps struct {
	// The application ID.
	//
	// example:
	//
	// 123
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The timestamp generated when the task was created.
	//
	// example:
	//
	// 1531291867000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The aliases of the application.
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The process identifier (PID) of the application.
	//
	// example:
	//
	// atc889zkcf@d8deedfa9bf****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the application is displayed in the Application Real-Time Monitoring Service (ARMS) console. Valid values:
	//
	// 	- `true`: yes
	//
	// 	- `false`: no
	//
	// example:
	//
	// true
	Show *bool `json:"Show,omitempty" xml:"Show,omitempty"`
	// A list of tags.
	Tags []*SearchTraceAppByPageResponseBodyPageBeanTraceAppsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the monitoring task. Valid values:
	//
	// 	- `TRACE`: Application Monitoring
	//
	// 	- `RETCODE`: Browser Monitoring
	//
	// example:
	//
	// TRACE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The timestamp generated when the task information was updated.
	//
	// example:
	//
	// 1531291867000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 113197164949****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchTraceAppByPageResponseBodyPageBeanTraceApps) String() string {
	return dara.Prettify(s)
}

func (s SearchTraceAppByPageResponseBodyPageBeanTraceApps) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) GetAppId() *int64 {
	return s.AppId
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) GetAppName() *string {
	return s.AppName
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) GetLabels() []*string {
	return s.Labels
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) GetPid() *string {
	return s.Pid
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) GetShow() *bool {
	return s.Show
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) GetTags() []*SearchTraceAppByPageResponseBodyPageBeanTraceAppsTags {
	return s.Tags
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) GetType() *string {
	return s.Type
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) GetUserId() *string {
	return s.UserId
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetAppId(v int64) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.AppId = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetAppName(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.AppName = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetCreateTime(v int64) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.CreateTime = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetLabels(v []*string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.Labels = v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetPid(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.Pid = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetRegionId(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.RegionId = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetResourceGroupId(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.ResourceGroupId = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetShow(v bool) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.Show = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetTags(v []*SearchTraceAppByPageResponseBodyPageBeanTraceAppsTags) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.Tags = v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetType(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.Type = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetUpdateTime(v int64) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.UpdateTime = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetUserId(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.UserId = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchTraceAppByPageResponseBodyPageBeanTraceAppsTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchTraceAppByPageResponseBodyPageBeanTraceAppsTags) String() string {
	return dara.Prettify(s)
}

func (s SearchTraceAppByPageResponseBodyPageBeanTraceAppsTags) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceAppsTags) GetKey() *string {
	return s.Key
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceAppsTags) GetValue() *string {
	return s.Value
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceAppsTags) SetKey(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceAppsTags {
	s.Key = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceAppsTags) SetValue(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceAppsTags {
	s.Value = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceAppsTags) Validate() error {
	return dara.Validate(s)
}
