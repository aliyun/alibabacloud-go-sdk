// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRunningAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListRunningAppsResponseBody
	GetRequestId() *string
	SetRunningCloudApps(v []*ListRunningAppsResponseBodyRunningCloudApps) *ListRunningAppsResponseBody
	GetRunningCloudApps() []*ListRunningAppsResponseBodyRunningCloudApps
}

type ListRunningAppsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 2DC3521C-3820-5EA5-9A9A-00BB7AF4E8E5
	RequestId        *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RunningCloudApps []*ListRunningAppsResponseBodyRunningCloudApps `json:"RunningCloudApps,omitempty" xml:"RunningCloudApps,omitempty" type:"Repeated"`
}

func (s ListRunningAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRunningAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRunningAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRunningAppsResponseBody) GetRunningCloudApps() []*ListRunningAppsResponseBodyRunningCloudApps {
	return s.RunningCloudApps
}

func (s *ListRunningAppsResponseBody) SetRequestId(v string) *ListRunningAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRunningAppsResponseBody) SetRunningCloudApps(v []*ListRunningAppsResponseBodyRunningCloudApps) *ListRunningAppsResponseBody {
	s.RunningCloudApps = v
	return s
}

func (s *ListRunningAppsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRunningAppsResponseBodyRunningCloudApps struct {
	// example:
	//
	// ca-dln05y44ze6esfl8x
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// aig-dk8p95irk9xs5xi6a
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// example:
	//
	// ai-gc1gemx6vpa6vlync
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// example:
	//
	// alihealth-keeper
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 11.1
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// test1.0
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	// example:
	//
	// 87
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// https://app-icon-shanghai.oss-cn-shanghai.aliyuncs.com/tenant/187465/18_bf1.jpg
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1642748400
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListRunningAppsResponseBodyRunningCloudApps) String() string {
	return dara.Prettify(s)
}

func (s ListRunningAppsResponseBodyRunningCloudApps) GoString() string {
	return s.String()
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) GetAppId() *string {
	return s.AppId
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) GetAppName() *string {
	return s.AppName
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) GetAppVersionName() *string {
	return s.AppVersionName
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) GetDuration() *int64 {
	return s.Duration
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) GetIconUrl() *string {
	return s.IconUrl
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) GetOsType() *string {
	return s.OsType
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) GetRegionId() *string {
	return s.RegionId
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) GetStartTime() *string {
	return s.StartTime
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetAppId(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.AppId = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetAppInstanceGroupId(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetAppInstanceId(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.AppInstanceId = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetAppName(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.AppName = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetAppVersion(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.AppVersion = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetAppVersionName(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.AppVersionName = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetDuration(v int64) *ListRunningAppsResponseBodyRunningCloudApps {
	s.Duration = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetIconUrl(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.IconUrl = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetOsType(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.OsType = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetRegionId(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.RegionId = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetStartTime(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.StartTime = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) Validate() error {
	return dara.Validate(s)
}
