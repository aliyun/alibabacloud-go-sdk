// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTraceAppByNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SearchTraceAppByNameResponseBody
	GetRequestId() *string
	SetTraceApps(v []*SearchTraceAppByNameResponseBodyTraceApps) *SearchTraceAppByNameResponseBody
	GetTraceApps() []*SearchTraceAppByNameResponseBodyTraceApps
}

type SearchTraceAppByNameResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F7781D4A-2818-41E7-B7BB-79D809E9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the monitoring task.
	TraceApps []*SearchTraceAppByNameResponseBodyTraceApps `json:"TraceApps,omitempty" xml:"TraceApps,omitempty" type:"Repeated"`
}

func (s SearchTraceAppByNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchTraceAppByNameResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchTraceAppByNameResponseBody) GetTraceApps() []*SearchTraceAppByNameResponseBodyTraceApps {
	return s.TraceApps
}

func (s *SearchTraceAppByNameResponseBody) SetRequestId(v string) *SearchTraceAppByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTraceAppByNameResponseBody) SetTraceApps(v []*SearchTraceAppByNameResponseBodyTraceApps) *SearchTraceAppByNameResponseBody {
	s.TraceApps = v
	return s
}

func (s *SearchTraceAppByNameResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchTraceAppByNameResponseBodyTraceApps struct {
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
	// 1593486786000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The aliases of the application.
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The process identifier (PID) of the application.
	//
	// example:
	//
	// a5f9bdeb-2627-4dbe-9247-****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
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
	Tags []*SearchTraceAppByNameResponseBodyTraceAppsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	// 1593486786000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 113197164949****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchTraceAppByNameResponseBodyTraceApps) String() string {
	return dara.Prettify(s)
}

func (s SearchTraceAppByNameResponseBodyTraceApps) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) GetAppId() *int64 {
	return s.AppId
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) GetAppName() *string {
	return s.AppName
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) GetLabels() []*string {
	return s.Labels
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) GetPid() *string {
	return s.Pid
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) GetShow() *bool {
	return s.Show
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) GetTags() []*SearchTraceAppByNameResponseBodyTraceAppsTags {
	return s.Tags
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) GetType() *string {
	return s.Type
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) GetUserId() *string {
	return s.UserId
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetAppId(v int64) *SearchTraceAppByNameResponseBodyTraceApps {
	s.AppId = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetAppName(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.AppName = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetCreateTime(v int64) *SearchTraceAppByNameResponseBodyTraceApps {
	s.CreateTime = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetLabels(v []*string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.Labels = v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetPid(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.Pid = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetRegionId(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.RegionId = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetResourceGroupId(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.ResourceGroupId = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetShow(v bool) *SearchTraceAppByNameResponseBodyTraceApps {
	s.Show = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetTags(v []*SearchTraceAppByNameResponseBodyTraceAppsTags) *SearchTraceAppByNameResponseBodyTraceApps {
	s.Tags = v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetType(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.Type = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetUpdateTime(v int64) *SearchTraceAppByNameResponseBodyTraceApps {
	s.UpdateTime = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetUserId(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.UserId = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) Validate() error {
	return dara.Validate(s)
}

type SearchTraceAppByNameResponseBodyTraceAppsTags struct {
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

func (s SearchTraceAppByNameResponseBodyTraceAppsTags) String() string {
	return dara.Prettify(s)
}

func (s SearchTraceAppByNameResponseBodyTraceAppsTags) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByNameResponseBodyTraceAppsTags) GetKey() *string {
	return s.Key
}

func (s *SearchTraceAppByNameResponseBodyTraceAppsTags) GetValue() *string {
	return s.Value
}

func (s *SearchTraceAppByNameResponseBodyTraceAppsTags) SetKey(v string) *SearchTraceAppByNameResponseBodyTraceAppsTags {
	s.Key = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceAppsTags) SetValue(v string) *SearchTraceAppByNameResponseBodyTraceAppsTags {
	s.Value = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceAppsTags) Validate() error {
	return dara.Validate(s)
}
