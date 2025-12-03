// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApps(v *DescribeAppsResponseBodyApps) *DescribeAppsResponseBody
	GetApps() *DescribeAppsResponseBodyApps
	SetPageNumber(v int32) *DescribeAppsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAppsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAppsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAppsResponseBody
	GetTotalCount() *int32
}

type DescribeAppsResponseBody struct {
	// The returned app information. It is an array consisting of AppItem data.
	Apps *DescribeAppsResponseBodyApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ015
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBody) GetApps() *DescribeAppsResponseBodyApps {
	return s.Apps
}

func (s *DescribeAppsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAppsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAppsResponseBody) SetApps(v *DescribeAppsResponseBodyApps) *DescribeAppsResponseBody {
	s.Apps = v
	return s
}

func (s *DescribeAppsResponseBody) SetPageNumber(v int32) *DescribeAppsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppsResponseBody) SetPageSize(v int32) *DescribeAppsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAppsResponseBody) SetRequestId(v string) *DescribeAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppsResponseBody) SetTotalCount(v int32) *DescribeAppsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAppsResponseBody) Validate() error {
	if s.Apps != nil {
		if err := s.Apps.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAppsResponseBodyApps struct {
	AppItem []*DescribeAppsResponseBodyAppsAppItem `json:"AppItem,omitempty" xml:"AppItem,omitempty" type:"Repeated"`
}

func (s DescribeAppsResponseBodyApps) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyApps) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyApps) GetAppItem() []*DescribeAppsResponseBodyAppsAppItem {
	return s.AppItem
}

func (s *DescribeAppsResponseBodyApps) SetAppItem(v []*DescribeAppsResponseBodyAppsAppItem) *DescribeAppsResponseBodyApps {
	s.AppItem = v
	return s
}

func (s *DescribeAppsResponseBodyApps) Validate() error {
	if s.AppItem != nil {
		for _, item := range s.AppItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAppsResponseBodyAppsAppItem struct {
	// The ID of the app.
	//
	// example:
	//
	// 20112314518278
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the app.
	//
	// example:
	//
	// CreateApptest
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The description of the app.
	//
	// example:
	//
	// App test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s DescribeAppsResponseBodyAppsAppItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyAppsAppItem) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyAppsAppItem) GetAppId() *int64 {
	return s.AppId
}

func (s *DescribeAppsResponseBodyAppsAppItem) GetAppName() *string {
	return s.AppName
}

func (s *DescribeAppsResponseBodyAppsAppItem) GetDescription() *string {
	return s.Description
}

func (s *DescribeAppsResponseBodyAppsAppItem) SetAppId(v int64) *DescribeAppsResponseBodyAppsAppItem {
	s.AppId = &v
	return s
}

func (s *DescribeAppsResponseBodyAppsAppItem) SetAppName(v string) *DescribeAppsResponseBodyAppsAppItem {
	s.AppName = &v
	return s
}

func (s *DescribeAppsResponseBodyAppsAppItem) SetDescription(v string) *DescribeAppsResponseBodyAppsAppItem {
	s.Description = &v
	return s
}

func (s *DescribeAppsResponseBodyAppsAppItem) Validate() error {
	return dara.Validate(s)
}
