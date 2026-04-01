// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v *ListApplicationsResponseBodyApplications) *ListApplicationsResponseBody
	GetApplications() *ListApplicationsResponseBodyApplications
	SetPageNumber(v int32) *ListApplicationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListApplicationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListApplicationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListApplicationsResponseBody
	GetTotalCount() *int32
}

type ListApplicationsResponseBody struct {
	Applications *ListApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 50373E71-7710-4620-8AAB-133CCE49451C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 49
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBody) GetApplications() *ListApplicationsResponseBodyApplications {
	return s.Applications
}

func (s *ListApplicationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApplicationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListApplicationsResponseBody) SetApplications(v *ListApplicationsResponseBodyApplications) *ListApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *ListApplicationsResponseBody) SetPageNumber(v int32) *ListApplicationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationsResponseBody) SetPageSize(v int32) *ListApplicationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsResponseBody) SetRequestId(v string) *ListApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsResponseBody) SetTotalCount(v int32) *ListApplicationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationsResponseBody) Validate() error {
	if s.Applications != nil {
		if err := s.Applications.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApplicationsResponseBodyApplications struct {
	Application []*ListApplicationsResponseBodyApplicationsApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Repeated"`
}

func (s ListApplicationsResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplications) GetApplication() []*ListApplicationsResponseBodyApplicationsApplication {
	return s.Application
}

func (s *ListApplicationsResponseBodyApplications) SetApplication(v []*ListApplicationsResponseBodyApplicationsApplication) *ListApplicationsResponseBodyApplications {
	s.Application = v
	return s
}

func (s *ListApplicationsResponseBodyApplications) Validate() error {
	if s.Application != nil {
		for _, item := range s.Application {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationsResponseBodyApplicationsApplication struct {
	AppList     *ListApplicationsResponseBodyApplicationsApplicationAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Struct"`
	ClusterName *string                                                     `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
}

func (s ListApplicationsResponseBodyApplicationsApplication) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyApplicationsApplication) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplicationsApplication) GetAppList() *ListApplicationsResponseBodyApplicationsApplicationAppList {
	return s.AppList
}

func (s *ListApplicationsResponseBodyApplicationsApplication) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetAppList(v *ListApplicationsResponseBodyApplicationsApplicationAppList) *ListApplicationsResponseBodyApplicationsApplication {
	s.AppList = v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetClusterName(v string) *ListApplicationsResponseBodyApplicationsApplication {
	s.ClusterName = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) Validate() error {
	if s.AppList != nil {
		if err := s.AppList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApplicationsResponseBodyApplicationsApplicationAppList struct {
	App []*ListApplicationsResponseBodyApplicationsApplicationAppListApp `json:"App,omitempty" xml:"App,omitempty" type:"Repeated"`
}

func (s ListApplicationsResponseBodyApplicationsApplicationAppList) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyApplicationsApplicationAppList) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplicationsApplicationAppList) GetApp() []*ListApplicationsResponseBodyApplicationsApplicationAppListApp {
	return s.App
}

func (s *ListApplicationsResponseBodyApplicationsApplicationAppList) SetApp(v []*ListApplicationsResponseBodyApplicationsApplicationAppListApp) *ListApplicationsResponseBodyApplicationsApplicationAppList {
	s.App = v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplicationAppList) Validate() error {
	if s.App != nil {
		for _, item := range s.App {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationsResponseBodyApplicationsApplicationAppListApp struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppInfo *string `json:"AppInfo,omitempty" xml:"AppInfo,omitempty"`
}

func (s ListApplicationsResponseBodyApplicationsApplicationAppListApp) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyApplicationsApplicationAppListApp) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplicationsApplicationAppListApp) GetAppId() *string {
	return s.AppId
}

func (s *ListApplicationsResponseBodyApplicationsApplicationAppListApp) GetAppInfo() *string {
	return s.AppInfo
}

func (s *ListApplicationsResponseBodyApplicationsApplicationAppListApp) SetAppId(v string) *ListApplicationsResponseBodyApplicationsApplicationAppListApp {
	s.AppId = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplicationAppListApp) SetAppInfo(v string) *ListApplicationsResponseBodyApplicationsApplicationAppListApp {
	s.AppInfo = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplicationAppListApp) Validate() error {
	return dara.Validate(s)
}
