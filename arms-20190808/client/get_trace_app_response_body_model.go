// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTraceAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTraceAppResponseBody
	GetRequestId() *string
	SetTraceApp(v *GetTraceAppResponseBodyTraceApp) *GetTraceAppResponseBody
	GetTraceApp() *GetTraceAppResponseBodyTraceApp
}

type GetTraceAppResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D80ADAAC-8C32-5479-BD14-C28CF832****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the array object.
	TraceApp *GetTraceAppResponseBodyTraceApp `json:"TraceApp,omitempty" xml:"TraceApp,omitempty" type:"Struct"`
}

func (s GetTraceAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTraceAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetTraceAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTraceAppResponseBody) GetTraceApp() *GetTraceAppResponseBodyTraceApp {
	return s.TraceApp
}

func (s *GetTraceAppResponseBody) SetRequestId(v string) *GetTraceAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTraceAppResponseBody) SetTraceApp(v *GetTraceAppResponseBodyTraceApp) *GetTraceAppResponseBody {
	s.TraceApp = v
	return s
}

func (s *GetTraceAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTraceAppResponseBodyTraceApp struct {
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
	// arms-k8s-demo
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// c905d1364c2dd4b6284a3f41790c4****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The timestamp generated when the task was created.
	//
	// example:
	//
	// 1576599253000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The aliases of the application.
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The programming language of the application.
	//
	// example:
	//
	// java
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The process identifier (PID) of the application.
	//
	// example:
	//
	// b590lhguqs@d8deedfa9bf****
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
	// The source of the application.
	//
	// example:
	//
	// ACSK8S
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// A list of key-value pairs.
	Tags []*GetTraceAppResponseBodyTraceAppTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	// 1635700348000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 113197164949****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetTraceAppResponseBodyTraceApp) String() string {
	return dara.Prettify(s)
}

func (s GetTraceAppResponseBodyTraceApp) GoString() string {
	return s.String()
}

func (s *GetTraceAppResponseBodyTraceApp) GetAppId() *int64 {
	return s.AppId
}

func (s *GetTraceAppResponseBodyTraceApp) GetAppName() *string {
	return s.AppName
}

func (s *GetTraceAppResponseBodyTraceApp) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetTraceAppResponseBodyTraceApp) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetTraceAppResponseBodyTraceApp) GetLabels() []*string {
	return s.Labels
}

func (s *GetTraceAppResponseBodyTraceApp) GetLanguage() *string {
	return s.Language
}

func (s *GetTraceAppResponseBodyTraceApp) GetPid() *string {
	return s.Pid
}

func (s *GetTraceAppResponseBodyTraceApp) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTraceAppResponseBodyTraceApp) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetTraceAppResponseBodyTraceApp) GetShow() *bool {
	return s.Show
}

func (s *GetTraceAppResponseBodyTraceApp) GetSource() *string {
	return s.Source
}

func (s *GetTraceAppResponseBodyTraceApp) GetTags() []*GetTraceAppResponseBodyTraceAppTags {
	return s.Tags
}

func (s *GetTraceAppResponseBodyTraceApp) GetType() *string {
	return s.Type
}

func (s *GetTraceAppResponseBodyTraceApp) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetTraceAppResponseBodyTraceApp) GetUserId() *string {
	return s.UserId
}

func (s *GetTraceAppResponseBodyTraceApp) SetAppId(v int64) *GetTraceAppResponseBodyTraceApp {
	s.AppId = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetAppName(v string) *GetTraceAppResponseBodyTraceApp {
	s.AppName = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetClusterId(v string) *GetTraceAppResponseBodyTraceApp {
	s.ClusterId = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetCreateTime(v int64) *GetTraceAppResponseBodyTraceApp {
	s.CreateTime = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetLabels(v []*string) *GetTraceAppResponseBodyTraceApp {
	s.Labels = v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetLanguage(v string) *GetTraceAppResponseBodyTraceApp {
	s.Language = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetPid(v string) *GetTraceAppResponseBodyTraceApp {
	s.Pid = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetRegionId(v string) *GetTraceAppResponseBodyTraceApp {
	s.RegionId = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetResourceGroupId(v string) *GetTraceAppResponseBodyTraceApp {
	s.ResourceGroupId = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetShow(v bool) *GetTraceAppResponseBodyTraceApp {
	s.Show = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetSource(v string) *GetTraceAppResponseBodyTraceApp {
	s.Source = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetTags(v []*GetTraceAppResponseBodyTraceAppTags) *GetTraceAppResponseBodyTraceApp {
	s.Tags = v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetType(v string) *GetTraceAppResponseBodyTraceApp {
	s.Type = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetUpdateTime(v int64) *GetTraceAppResponseBodyTraceApp {
	s.UpdateTime = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetUserId(v string) *GetTraceAppResponseBodyTraceApp {
	s.UserId = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) Validate() error {
	return dara.Validate(s)
}

type GetTraceAppResponseBodyTraceAppTags struct {
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

func (s GetTraceAppResponseBodyTraceAppTags) String() string {
	return dara.Prettify(s)
}

func (s GetTraceAppResponseBodyTraceAppTags) GoString() string {
	return s.String()
}

func (s *GetTraceAppResponseBodyTraceAppTags) GetKey() *string {
	return s.Key
}

func (s *GetTraceAppResponseBodyTraceAppTags) GetValue() *string {
	return s.Value
}

func (s *GetTraceAppResponseBodyTraceAppTags) SetKey(v string) *GetTraceAppResponseBodyTraceAppTags {
	s.Key = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceAppTags) SetValue(v string) *GetTraceAppResponseBodyTraceAppTags {
	s.Value = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceAppTags) Validate() error {
	return dara.Validate(s)
}
