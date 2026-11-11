// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishedAppInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppModels(v []*ListPublishedAppInfoResponseBodyAppModels) *ListPublishedAppInfoResponseBody
	GetAppModels() []*ListPublishedAppInfoResponseBodyAppModels
	SetNextToken(v string) *ListPublishedAppInfoResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListPublishedAppInfoResponseBody
	GetRequestId() *string
}

type ListPublishedAppInfoResponseBody struct {
	// The list of published applications.
	AppModels []*ListPublishedAppInfoResponseBodyAppModels `json:"AppModels,omitempty" xml:"AppModels,omitempty" type:"Repeated"`
	// The starting token for the next query.
	//
	// example:
	//
	// 2NVfhLfgy5b3J5iJyoLQ6x4EULMg1hbhgB9NfnvdK9oj5zwxd17j4TuQkZze3RvhEvBinZYjknujF3Q1M
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DB70F8FE-63A3-587B-8560-CEC258E8B944
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPublishedAppInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublishedAppInfoResponseBody) GetAppModels() []*ListPublishedAppInfoResponseBodyAppModels {
	return s.AppModels
}

func (s *ListPublishedAppInfoResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPublishedAppInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPublishedAppInfoResponseBody) SetAppModels(v []*ListPublishedAppInfoResponseBodyAppModels) *ListPublishedAppInfoResponseBody {
	s.AppModels = v
	return s
}

func (s *ListPublishedAppInfoResponseBody) SetNextToken(v string) *ListPublishedAppInfoResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPublishedAppInfoResponseBody) SetRequestId(v string) *ListPublishedAppInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublishedAppInfoResponseBody) Validate() error {
	if s.AppModels != nil {
		for _, item := range s.AppModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPublishedAppInfoResponseBodyAppModels struct {
	// The application image ID. You can obtain the ID from the Image Management page in the WUYING Cloud App console.
	//
	// example:
	//
	// img-f37nddbjc1lje14st
	AppCenterImageId *string `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	// The application ID.
	//
	// example:
	//
	// ca-fxwp4koyr5y2sp4mz
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// Microsoft Word
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The theme color of the application.
	//
	// example:
	//
	// #FFFFFF
	AppThemeColor *string `json:"AppThemeColor,omitempty" xml:"AppThemeColor,omitempty"`
	// The application version number.
	//
	// example:
	//
	// R2021a
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The application version name.
	//
	// example:
	//
	// v1.0
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	// The authorization time of the application.
	//
	// The return value is a serialized datetime string from the server in the format of `EEE MMM dd HH:mm:ss zzz yyyy`, which represents the day of the week, month, day, hour:minute:second, time zone abbreviation, and year in sequence. For example, `Wed Jan 14 20:47:02 CST 2026`, where `CST` is the time zone abbreviation of the server\\"s time zone (China Standard Time, UTC+08:00).
	//
	// example:
	//
	// Wed Jan 14 20:47:02 CST 2026
	AuthTime *string `json:"AuthTime,omitempty" xml:"AuthTime,omitempty"`
	// The specific category information of the application, which is the secondary category.
	//
	// Valid values:
	//
	// - 1: All applications
	//
	// - 2: Design software
	//
	// - 3: Programming and development
	//
	// - 4: Office essentials
	//
	// - 5: Input methods
	//
	// - 6: Image and video
	//
	// - 7: System tools
	//
	// - 8: Browsers
	//
	// example:
	//
	// 2
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The primary category of the application. Valid values:
	//
	// - 0: Cloud application
	//
	// - 1: Cloud game
	//
	// example:
	//
	// 1
	CategoryType *int64 `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// The URL of the application icon.
	//
	// example:
	//
	// https://app-streaming-icon-prod-shanghai.oss-cn-shanghai.aliyuncs.com/tenant/1973619010349344/1634523814270_Matlab.png
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	// Indicates whether the user has permissions.
	//
	// The value is True.
	//
	// example:
	//
	// True
	IsAuth *bool `json:"IsAuth,omitempty" xml:"IsAuth,omitempty"`
	// Indicates whether the application is in use.
	//
	// example:
	//
	// True
	UsedInSession *bool `json:"UsedInSession,omitempty" xml:"UsedInSession,omitempty"`
}

func (s ListPublishedAppInfoResponseBodyAppModels) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAppInfoResponseBodyAppModels) GoString() string {
	return s.String()
}

func (s *ListPublishedAppInfoResponseBodyAppModels) GetAppCenterImageId() *string {
	return s.AppCenterImageId
}

func (s *ListPublishedAppInfoResponseBodyAppModels) GetAppId() *string {
	return s.AppId
}

func (s *ListPublishedAppInfoResponseBodyAppModels) GetAppName() *string {
	return s.AppName
}

func (s *ListPublishedAppInfoResponseBodyAppModels) GetAppThemeColor() *string {
	return s.AppThemeColor
}

func (s *ListPublishedAppInfoResponseBodyAppModels) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ListPublishedAppInfoResponseBodyAppModels) GetAppVersionName() *string {
	return s.AppVersionName
}

func (s *ListPublishedAppInfoResponseBodyAppModels) GetAuthTime() *string {
	return s.AuthTime
}

func (s *ListPublishedAppInfoResponseBodyAppModels) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *ListPublishedAppInfoResponseBodyAppModels) GetCategoryType() *int64 {
	return s.CategoryType
}

func (s *ListPublishedAppInfoResponseBodyAppModels) GetIconUrl() *string {
	return s.IconUrl
}

func (s *ListPublishedAppInfoResponseBodyAppModels) GetIsAuth() *bool {
	return s.IsAuth
}

func (s *ListPublishedAppInfoResponseBodyAppModels) GetUsedInSession() *bool {
	return s.UsedInSession
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetAppCenterImageId(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.AppCenterImageId = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetAppId(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.AppId = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetAppName(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.AppName = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetAppThemeColor(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.AppThemeColor = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetAppVersion(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.AppVersion = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetAppVersionName(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.AppVersionName = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetAuthTime(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.AuthTime = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetCategoryId(v int64) *ListPublishedAppInfoResponseBodyAppModels {
	s.CategoryId = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetCategoryType(v int64) *ListPublishedAppInfoResponseBodyAppModels {
	s.CategoryType = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetIconUrl(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.IconUrl = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetIsAuth(v bool) *ListPublishedAppInfoResponseBodyAppModels {
	s.IsAuth = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetUsedInSession(v bool) *ListPublishedAppInfoResponseBodyAppModels {
	s.UsedInSession = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) Validate() error {
	return dara.Validate(s)
}
