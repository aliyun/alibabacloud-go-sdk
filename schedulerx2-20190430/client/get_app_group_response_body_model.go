// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetAppGroupResponseBody
	GetCode() *int32
	SetData(v *GetAppGroupResponseBodyData) *GetAppGroupResponseBody
	GetData() *GetAppGroupResponseBodyData
	SetMessage(v string) *GetAppGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAppGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAppGroupResponseBody
	GetSuccess() *bool
}

type GetAppGroupResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the application group.
	Data *GetAppGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The additional information that is returned.
	//
	// example:
	//
	// app is not existed, groupId=xxxx, namesapce=xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAppGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetAppGroupResponseBody) GetData() *GetAppGroupResponseBodyData {
	return s.Data
}

func (s *GetAppGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAppGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAppGroupResponseBody) SetCode(v int32) *GetAppGroupResponseBody {
	s.Code = &v
	return s
}

func (s *GetAppGroupResponseBody) SetData(v *GetAppGroupResponseBodyData) *GetAppGroupResponseBody {
	s.Data = v
	return s
}

func (s *GetAppGroupResponseBody) SetMessage(v string) *GetAppGroupResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppGroupResponseBody) SetRequestId(v string) *GetAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppGroupResponseBody) SetSuccess(v bool) *GetAppGroupResponseBody {
	s.Success = &v
	return s
}

func (s *GetAppGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppGroupResponseBodyData struct {
	// The AppKey of the application.
	//
	// example:
	//
	// QI4lWMZ+xk1rNB67jFUhaw==
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// DocTest
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application version. 1: Basic version, 2: Professional version.
	//
	// example:
	//
	// 2
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The number of jobs that are configured for the application group.
	//
	// example:
	//
	// 1
	CurJobs *int32 `json:"CurJobs,omitempty" xml:"CurJobs,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The maximum number of jobs that can be configured for the application group.
	//
	// example:
	//
	// 1000
	MaxJobs *int32 `json:"MaxJobs,omitempty" xml:"MaxJobs,omitempty"`
	// The alert notification configurations.
	//
	// >  For more information about this parameter, see the following **additional information about request parameters**.
	//
	// example:
	//
	// {"sendChannel":"sms,mail,ding"}
	MonitorConfigJson *string `json:"MonitorConfigJson,omitempty" xml:"MonitorConfigJson,omitempty"`
	// The alert contact configurations.
	//
	// >  For more information about this parameter, see the following **additional information about request parameters**.
	//
	// example:
	//
	// [ {"name": "Peter"}, {"name": "Paul"} ]
	MonitorContactsJson *string `json:"MonitorContactsJson,omitempty" xml:"MonitorContactsJson,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// test-workday-notification
	NotificationPolicyName *string `json:"NotificationPolicyName,omitempty" xml:"NotificationPolicyName,omitempty"`
}

func (s GetAppGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAppGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAppGroupResponseBodyData) GetAppKey() *string {
	return s.AppKey
}

func (s *GetAppGroupResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *GetAppGroupResponseBodyData) GetAppVersion() *string {
	return s.AppVersion
}

func (s *GetAppGroupResponseBodyData) GetCurJobs() *int32 {
	return s.CurJobs
}

func (s *GetAppGroupResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetAppGroupResponseBodyData) GetGroupId() *string {
	return s.GroupId
}

func (s *GetAppGroupResponseBodyData) GetMaxJobs() *int32 {
	return s.MaxJobs
}

func (s *GetAppGroupResponseBodyData) GetMonitorConfigJson() *string {
	return s.MonitorConfigJson
}

func (s *GetAppGroupResponseBodyData) GetMonitorContactsJson() *string {
	return s.MonitorContactsJson
}

func (s *GetAppGroupResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *GetAppGroupResponseBodyData) GetNotificationPolicyName() *string {
	return s.NotificationPolicyName
}

func (s *GetAppGroupResponseBodyData) SetAppKey(v string) *GetAppGroupResponseBodyData {
	s.AppKey = &v
	return s
}

func (s *GetAppGroupResponseBodyData) SetAppName(v string) *GetAppGroupResponseBodyData {
	s.AppName = &v
	return s
}

func (s *GetAppGroupResponseBodyData) SetAppVersion(v string) *GetAppGroupResponseBodyData {
	s.AppVersion = &v
	return s
}

func (s *GetAppGroupResponseBodyData) SetCurJobs(v int32) *GetAppGroupResponseBodyData {
	s.CurJobs = &v
	return s
}

func (s *GetAppGroupResponseBodyData) SetDescription(v string) *GetAppGroupResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetAppGroupResponseBodyData) SetGroupId(v string) *GetAppGroupResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *GetAppGroupResponseBodyData) SetMaxJobs(v int32) *GetAppGroupResponseBodyData {
	s.MaxJobs = &v
	return s
}

func (s *GetAppGroupResponseBodyData) SetMonitorConfigJson(v string) *GetAppGroupResponseBodyData {
	s.MonitorConfigJson = &v
	return s
}

func (s *GetAppGroupResponseBodyData) SetMonitorContactsJson(v string) *GetAppGroupResponseBodyData {
	s.MonitorContactsJson = &v
	return s
}

func (s *GetAppGroupResponseBodyData) SetNamespace(v string) *GetAppGroupResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *GetAppGroupResponseBodyData) SetNotificationPolicyName(v string) *GetAppGroupResponseBodyData {
	s.NotificationPolicyName = &v
	return s
}

func (s *GetAppGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
