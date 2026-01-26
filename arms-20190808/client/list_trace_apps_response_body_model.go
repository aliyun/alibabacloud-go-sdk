// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTraceAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListTraceAppsResponseBody
	GetCode() *int32
	SetMessage(v string) *ListTraceAppsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTraceAppsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTraceAppsResponseBody
	GetSuccess() *bool
	SetTraceApps(v []*ListTraceAppsResponseBodyTraceApps) *ListTraceAppsResponseBody
	GetTraceApps() []*ListTraceAppsResponseBodyTraceApps
}

type ListTraceAppsResponseBody struct {
	// The HTTP status code returned for the request. Valid values:
	//
	// 	- `2XX`: The request is successful.
	//
	// 	- `3XX`: A redirection message is returned.
	//
	// 	- `4XX`: The request is invalid.
	//
	// 	- `5XX`: A server error occurs.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned if the request parameters are invalid.
	//
	// example:
	//
	// Internal error. Please try again. Contact the DingTalk service account if the issue persists after multiple retries.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40B10E04-81E8-4643-970D-F1B38F2E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- `true`: The call was successful.
	//
	// 	- `false`: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The list of Application Monitoring tasks.
	TraceApps []*ListTraceAppsResponseBodyTraceApps `json:"TraceApps,omitempty" xml:"TraceApps,omitempty" type:"Repeated"`
}

func (s ListTraceAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTraceAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTraceAppsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListTraceAppsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTraceAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTraceAppsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTraceAppsResponseBody) GetTraceApps() []*ListTraceAppsResponseBodyTraceApps {
	return s.TraceApps
}

func (s *ListTraceAppsResponseBody) SetCode(v int32) *ListTraceAppsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTraceAppsResponseBody) SetMessage(v string) *ListTraceAppsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTraceAppsResponseBody) SetRequestId(v string) *ListTraceAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTraceAppsResponseBody) SetSuccess(v bool) *ListTraceAppsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTraceAppsResponseBody) SetTraceApps(v []*ListTraceAppsResponseBodyTraceApps) *ListTraceAppsResponseBody {
	s.TraceApps = v
	return s
}

func (s *ListTraceAppsResponseBody) Validate() error {
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

type ListTraceAppsResponseBodyTraceApps struct {
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
	// The cluster ID.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The time when the monitoring task was created. The value is a timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1529667762000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The labels of the application.
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The language.
	//
	// example:
	//
	// PHP
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
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
	// The resource group ID.
	//
	// example:
	//
	// Resource group
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
	// The application source.
	//
	// example:
	//
	// ACK
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The tags.
	Tags []*ListTraceAppsResponseBodyTraceAppsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	// The time when the monitoring task was updated. The value is a timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1529667762000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 113197164949****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The type of the workload.
	//
	// example:
	//
	// Deployment
	WorkloadKind *string `json:"WorkloadKind,omitempty" xml:"WorkloadKind,omitempty"`
	// The name of the workload.
	//
	// example:
	//
	// nginx-deployment
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
}

func (s ListTraceAppsResponseBodyTraceApps) String() string {
	return dara.Prettify(s)
}

func (s ListTraceAppsResponseBodyTraceApps) GoString() string {
	return s.String()
}

func (s *ListTraceAppsResponseBodyTraceApps) GetAppId() *int64 {
	return s.AppId
}

func (s *ListTraceAppsResponseBodyTraceApps) GetAppName() *string {
	return s.AppName
}

func (s *ListTraceAppsResponseBodyTraceApps) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListTraceAppsResponseBodyTraceApps) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListTraceAppsResponseBodyTraceApps) GetLabels() []*string {
	return s.Labels
}

func (s *ListTraceAppsResponseBodyTraceApps) GetLanguage() *string {
	return s.Language
}

func (s *ListTraceAppsResponseBodyTraceApps) GetNamespace() *string {
	return s.Namespace
}

func (s *ListTraceAppsResponseBodyTraceApps) GetPid() *string {
	return s.Pid
}

func (s *ListTraceAppsResponseBodyTraceApps) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTraceAppsResponseBodyTraceApps) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTraceAppsResponseBodyTraceApps) GetShow() *bool {
	return s.Show
}

func (s *ListTraceAppsResponseBodyTraceApps) GetSource() *string {
	return s.Source
}

func (s *ListTraceAppsResponseBodyTraceApps) GetTags() []*ListTraceAppsResponseBodyTraceAppsTags {
	return s.Tags
}

func (s *ListTraceAppsResponseBodyTraceApps) GetType() *string {
	return s.Type
}

func (s *ListTraceAppsResponseBodyTraceApps) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListTraceAppsResponseBodyTraceApps) GetUserId() *string {
	return s.UserId
}

func (s *ListTraceAppsResponseBodyTraceApps) GetWorkloadKind() *string {
	return s.WorkloadKind
}

func (s *ListTraceAppsResponseBodyTraceApps) GetWorkloadName() *string {
	return s.WorkloadName
}

func (s *ListTraceAppsResponseBodyTraceApps) SetAppId(v int64) *ListTraceAppsResponseBodyTraceApps {
	s.AppId = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetAppName(v string) *ListTraceAppsResponseBodyTraceApps {
	s.AppName = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetClusterId(v string) *ListTraceAppsResponseBodyTraceApps {
	s.ClusterId = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetCreateTime(v int64) *ListTraceAppsResponseBodyTraceApps {
	s.CreateTime = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetLabels(v []*string) *ListTraceAppsResponseBodyTraceApps {
	s.Labels = v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetLanguage(v string) *ListTraceAppsResponseBodyTraceApps {
	s.Language = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetNamespace(v string) *ListTraceAppsResponseBodyTraceApps {
	s.Namespace = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetPid(v string) *ListTraceAppsResponseBodyTraceApps {
	s.Pid = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetRegionId(v string) *ListTraceAppsResponseBodyTraceApps {
	s.RegionId = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetResourceGroupId(v string) *ListTraceAppsResponseBodyTraceApps {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetShow(v bool) *ListTraceAppsResponseBodyTraceApps {
	s.Show = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetSource(v string) *ListTraceAppsResponseBodyTraceApps {
	s.Source = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetTags(v []*ListTraceAppsResponseBodyTraceAppsTags) *ListTraceAppsResponseBodyTraceApps {
	s.Tags = v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetType(v string) *ListTraceAppsResponseBodyTraceApps {
	s.Type = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetUpdateTime(v int64) *ListTraceAppsResponseBodyTraceApps {
	s.UpdateTime = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetUserId(v string) *ListTraceAppsResponseBodyTraceApps {
	s.UserId = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetWorkloadKind(v string) *ListTraceAppsResponseBodyTraceApps {
	s.WorkloadKind = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetWorkloadName(v string) *ListTraceAppsResponseBodyTraceApps {
	s.WorkloadName = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) Validate() error {
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

type ListTraceAppsResponseBodyTraceAppsTags struct {
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

func (s ListTraceAppsResponseBodyTraceAppsTags) String() string {
	return dara.Prettify(s)
}

func (s ListTraceAppsResponseBodyTraceAppsTags) GoString() string {
	return s.String()
}

func (s *ListTraceAppsResponseBodyTraceAppsTags) GetKey() *string {
	return s.Key
}

func (s *ListTraceAppsResponseBodyTraceAppsTags) GetValue() *string {
	return s.Value
}

func (s *ListTraceAppsResponseBodyTraceAppsTags) SetKey(v string) *ListTraceAppsResponseBodyTraceAppsTags {
	s.Key = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceAppsTags) SetValue(v string) *ListTraceAppsResponseBodyTraceAppsTags {
	s.Value = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceAppsTags) Validate() error {
	return dara.Validate(s)
}
