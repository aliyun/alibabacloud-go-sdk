// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppList(v *DescribeAppsResponseBodyAppList) *DescribeAppsResponseBody
	GetAppList() *DescribeAppsResponseBodyAppList
	SetRequestId(v string) *DescribeAppsResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *DescribeAppsResponseBody
	GetTotalNum() *int32
	SetTotalPage(v int32) *DescribeAppsResponseBody
	GetTotalPage() *int32
}

type DescribeAppsResponseBody struct {
	AppList *DescribeAppsResponseBodyAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Struct"`
	// example:
	//
	// 6159ba01-6687-4fb2-a831-f0cd8d188648
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBody) GetAppList() *DescribeAppsResponseBodyAppList {
	return s.AppList
}

func (s *DescribeAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppsResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *DescribeAppsResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeAppsResponseBody) SetAppList(v *DescribeAppsResponseBodyAppList) *DescribeAppsResponseBody {
	s.AppList = v
	return s
}

func (s *DescribeAppsResponseBody) SetRequestId(v string) *DescribeAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppsResponseBody) SetTotalNum(v int32) *DescribeAppsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeAppsResponseBody) SetTotalPage(v int32) *DescribeAppsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeAppsResponseBody) Validate() error {
	if s.AppList != nil {
		if err := s.AppList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAppsResponseBodyAppList struct {
	App []*DescribeAppsResponseBodyAppListApp `json:"App,omitempty" xml:"App,omitempty" type:"Repeated"`
}

func (s DescribeAppsResponseBodyAppList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyAppList) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyAppList) GetApp() []*DescribeAppsResponseBodyAppListApp {
	return s.App
}

func (s *DescribeAppsResponseBodyAppList) SetApp(v []*DescribeAppsResponseBodyAppListApp) *DescribeAppsResponseBodyAppList {
	s.App = v
	return s
}

func (s *DescribeAppsResponseBodyAppList) Validate() error {
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

type DescribeAppsResponseBodyAppListApp struct {
	// example:
	//
	// rgf1****"
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// Default AppName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// universal
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// paybyduration
	BillType *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	// example:
	//
	// 2020-01-09T02:02:29Z
	CreateTime   *string                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ServiceAreas *DescribeAppsResponseBodyAppListAppServiceAreas `json:"ServiceAreas,omitempty" xml:"ServiceAreas,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Status  *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAppsResponseBodyAppListApp) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyAppListApp) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyAppListApp) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppsResponseBodyAppListApp) GetAppName() *string {
	return s.AppName
}

func (s *DescribeAppsResponseBodyAppListApp) GetAppType() *string {
	return s.AppType
}

func (s *DescribeAppsResponseBodyAppListApp) GetBillType() *string {
	return s.BillType
}

func (s *DescribeAppsResponseBodyAppListApp) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAppsResponseBodyAppListApp) GetServiceAreas() *DescribeAppsResponseBodyAppListAppServiceAreas {
	return s.ServiceAreas
}

func (s *DescribeAppsResponseBodyAppListApp) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeAppsResponseBodyAppListApp) GetVersion() *string {
	return s.Version
}

func (s *DescribeAppsResponseBodyAppListApp) SetAppId(v string) *DescribeAppsResponseBodyAppListApp {
	s.AppId = &v
	return s
}

func (s *DescribeAppsResponseBodyAppListApp) SetAppName(v string) *DescribeAppsResponseBodyAppListApp {
	s.AppName = &v
	return s
}

func (s *DescribeAppsResponseBodyAppListApp) SetAppType(v string) *DescribeAppsResponseBodyAppListApp {
	s.AppType = &v
	return s
}

func (s *DescribeAppsResponseBodyAppListApp) SetBillType(v string) *DescribeAppsResponseBodyAppListApp {
	s.BillType = &v
	return s
}

func (s *DescribeAppsResponseBodyAppListApp) SetCreateTime(v string) *DescribeAppsResponseBodyAppListApp {
	s.CreateTime = &v
	return s
}

func (s *DescribeAppsResponseBodyAppListApp) SetServiceAreas(v *DescribeAppsResponseBodyAppListAppServiceAreas) *DescribeAppsResponseBodyAppListApp {
	s.ServiceAreas = v
	return s
}

func (s *DescribeAppsResponseBodyAppListApp) SetStatus(v int32) *DescribeAppsResponseBodyAppListApp {
	s.Status = &v
	return s
}

func (s *DescribeAppsResponseBodyAppListApp) SetVersion(v string) *DescribeAppsResponseBodyAppListApp {
	s.Version = &v
	return s
}

func (s *DescribeAppsResponseBodyAppListApp) Validate() error {
	if s.ServiceAreas != nil {
		if err := s.ServiceAreas.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAppsResponseBodyAppListAppServiceAreas struct {
	ServiceArea []*string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty" type:"Repeated"`
}

func (s DescribeAppsResponseBodyAppListAppServiceAreas) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyAppListAppServiceAreas) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyAppListAppServiceAreas) GetServiceArea() []*string {
	return s.ServiceArea
}

func (s *DescribeAppsResponseBodyAppListAppServiceAreas) SetServiceArea(v []*string) *DescribeAppsResponseBodyAppListAppServiceAreas {
	s.ServiceArea = v
	return s
}

func (s *DescribeAppsResponseBodyAppListAppServiceAreas) Validate() error {
	return dara.Validate(s)
}
