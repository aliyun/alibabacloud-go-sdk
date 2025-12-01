// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstalledAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApps(v []*ListInstalledAppsResponseBodyApps) *ListInstalledAppsResponseBody
	GetApps() []*ListInstalledAppsResponseBodyApps
	SetRequestId(v string) *ListInstalledAppsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListInstalledAppsResponseBody
	GetTotalCount() *int64
}

type ListInstalledAppsResponseBody struct {
	// The information about the application.
	Apps []*ListInstalledAppsResponseBodyApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	// The unique ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 94
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstalledAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstalledAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstalledAppsResponseBody) GetApps() []*ListInstalledAppsResponseBodyApps {
	return s.Apps
}

func (s *ListInstalledAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstalledAppsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListInstalledAppsResponseBody) SetApps(v []*ListInstalledAppsResponseBodyApps) *ListInstalledAppsResponseBody {
	s.Apps = v
	return s
}

func (s *ListInstalledAppsResponseBody) SetRequestId(v string) *ListInstalledAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstalledAppsResponseBody) SetTotalCount(v int64) *ListInstalledAppsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstalledAppsResponseBody) Validate() error {
	if s.Apps != nil {
		for _, item := range s.Apps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstalledAppsResponseBodyApps struct {
	// The name of the application.
	//
	// example:
	//
	// test_app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application version.
	//
	// example:
	//
	// 2.0.1
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
}

func (s ListInstalledAppsResponseBodyApps) String() string {
	return dara.Prettify(s)
}

func (s ListInstalledAppsResponseBodyApps) GoString() string {
	return s.String()
}

func (s *ListInstalledAppsResponseBodyApps) GetAppName() *string {
	return s.AppName
}

func (s *ListInstalledAppsResponseBodyApps) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ListInstalledAppsResponseBodyApps) SetAppName(v string) *ListInstalledAppsResponseBodyApps {
	s.AppName = &v
	return s
}

func (s *ListInstalledAppsResponseBodyApps) SetAppVersion(v string) *ListInstalledAppsResponseBodyApps {
	s.AppVersion = &v
	return s
}

func (s *ListInstalledAppsResponseBodyApps) Validate() error {
	return dara.Validate(s)
}
