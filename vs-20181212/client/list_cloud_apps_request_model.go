// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListCloudAppsRequest
	GetAppId() *string
	SetAppName(v string) *ListCloudAppsRequest
	GetAppName() *string
	SetAppVersion(v string) *ListCloudAppsRequest
	GetAppVersion() *string
	SetEndTime(v string) *ListCloudAppsRequest
	GetEndTime() *string
	SetLatestVersionOnly(v bool) *ListCloudAppsRequest
	GetLatestVersionOnly() *bool
	SetPageNumber(v int64) *ListCloudAppsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListCloudAppsRequest
	GetPageSize() *int64
	SetPkgLabel(v string) *ListCloudAppsRequest
	GetPkgLabel() *string
	SetPkgType(v string) *ListCloudAppsRequest
	GetPkgType() *string
	SetStartTime(v string) *ListCloudAppsRequest
	GetStartTime() *string
	SetStatus(v string) *ListCloudAppsRequest
	GetStatus() *string
}

type ListCloudAppsRequest struct {
	// The cloud application ID, which corresponds to a unique application package.
	//
	// example:
	//
	// cap-b06b26edfhytbn b94a75ae1a79efc90eb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// com.aaa.bbb
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application version.
	//
	// example:
	//
	// 1.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The start time for time range filtering. Specify the time in the ISO 8601 standard in UTC. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2021-09-30T02:23:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether to return only the latest submitted version of each application. Default value: false.
	//
	// example:
	//
	// false
	LatestVersionOnly *bool `json:"LatestVersionOnly,omitempty" xml:"LatestVersionOnly,omitempty"`
	// The page number of the query list. Minimum value: 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query. Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The cloud application label.
	//
	// 1. Valid values:
	//
	//    a. hot
	//
	//    b. game
	//
	//    c. app
	//
	// 2. Special case:
	//
	//    a. To query applications without a label, enter "NULL".
	//
	// example:
	//
	// app
	PkgLabel *string `json:"PkgLabel,omitempty" xml:"PkgLabel,omitempty"`
	// The installation package type. Valid values:
	//
	// 1. android
	//
	// 2. win
	//
	// 3. android_appmarket
	//
	// example:
	//
	// android
	PkgType *string `json:"PkgType,omitempty" xml:"PkgType,omitempty"`
	// The start time for time range filtering. Specify the time in the ISO 8601 standard in UTC. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2021-09-29T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The application upload status. Valid values:
	//
	// 1. Success: desired state, succeeded.
	//
	// 2. Failed: desired state, failed.
	//
	// 3. Created
	//
	// 4. Doing
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCloudAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAppsRequest) GoString() string {
	return s.String()
}

func (s *ListCloudAppsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListCloudAppsRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListCloudAppsRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ListCloudAppsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListCloudAppsRequest) GetLatestVersionOnly() *bool {
	return s.LatestVersionOnly
}

func (s *ListCloudAppsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListCloudAppsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCloudAppsRequest) GetPkgLabel() *string {
	return s.PkgLabel
}

func (s *ListCloudAppsRequest) GetPkgType() *string {
	return s.PkgType
}

func (s *ListCloudAppsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListCloudAppsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListCloudAppsRequest) SetAppId(v string) *ListCloudAppsRequest {
	s.AppId = &v
	return s
}

func (s *ListCloudAppsRequest) SetAppName(v string) *ListCloudAppsRequest {
	s.AppName = &v
	return s
}

func (s *ListCloudAppsRequest) SetAppVersion(v string) *ListCloudAppsRequest {
	s.AppVersion = &v
	return s
}

func (s *ListCloudAppsRequest) SetEndTime(v string) *ListCloudAppsRequest {
	s.EndTime = &v
	return s
}

func (s *ListCloudAppsRequest) SetLatestVersionOnly(v bool) *ListCloudAppsRequest {
	s.LatestVersionOnly = &v
	return s
}

func (s *ListCloudAppsRequest) SetPageNumber(v int64) *ListCloudAppsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCloudAppsRequest) SetPageSize(v int64) *ListCloudAppsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloudAppsRequest) SetPkgLabel(v string) *ListCloudAppsRequest {
	s.PkgLabel = &v
	return s
}

func (s *ListCloudAppsRequest) SetPkgType(v string) *ListCloudAppsRequest {
	s.PkgType = &v
	return s
}

func (s *ListCloudAppsRequest) SetStartTime(v string) *ListCloudAppsRequest {
	s.StartTime = &v
	return s
}

func (s *ListCloudAppsRequest) SetStatus(v string) *ListCloudAppsRequest {
	s.Status = &v
	return s
}

func (s *ListCloudAppsRequest) Validate() error {
	return dara.Validate(s)
}
