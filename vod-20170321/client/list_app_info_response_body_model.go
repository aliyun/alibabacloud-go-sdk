// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppInfoList(v []*ListAppInfoResponseBodyAppInfoList) *ListAppInfoResponseBody
	GetAppInfoList() []*ListAppInfoResponseBodyAppInfoList
	SetRequestId(v string) *ListAppInfoResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListAppInfoResponseBody
	GetTotal() *int32
}

type ListAppInfoResponseBody struct {
	// The list of application information.
	AppInfoList []*ListAppInfoResponseBodyAppInfoList `json:"AppInfoList,omitempty" xml:"AppInfoList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4A13-4D5C-D7393642****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListAppInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppInfoResponseBody) GetAppInfoList() []*ListAppInfoResponseBodyAppInfoList {
	return s.AppInfoList
}

func (s *ListAppInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppInfoResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListAppInfoResponseBody) SetAppInfoList(v []*ListAppInfoResponseBodyAppInfoList) *ListAppInfoResponseBody {
	s.AppInfoList = v
	return s
}

func (s *ListAppInfoResponseBody) SetRequestId(v string) *ListAppInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppInfoResponseBody) SetTotal(v int32) *ListAppInfoResponseBody {
	s.Total = &v
	return s
}

func (s *ListAppInfoResponseBody) Validate() error {
	if s.AppInfoList != nil {
		for _, item := range s.AppInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppInfoResponseBodyAppInfoList struct {
	// The application ID.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The time when the application was created. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2019-03-01T08:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The application description.
	//
	// example:
	//
	// my first app.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the application was last modified. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format (UTC).
	//
	// example:
	//
	// 2019-03-01T09:00:00Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The region.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzko7fsuj****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The application status. Valid values:
	//
	// - **Normal**
	//
	// - **Disable**
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The application type. Valid values:
	//
	// - **System**: system default.
	//
	// - **Custom**: user-created.
	//
	// example:
	//
	// System
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAppInfoResponseBodyAppInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListAppInfoResponseBodyAppInfoList) GoString() string {
	return s.String()
}

func (s *ListAppInfoResponseBodyAppInfoList) GetAppId() *string {
	return s.AppId
}

func (s *ListAppInfoResponseBodyAppInfoList) GetAppName() *string {
	return s.AppName
}

func (s *ListAppInfoResponseBodyAppInfoList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListAppInfoResponseBodyAppInfoList) GetDescription() *string {
	return s.Description
}

func (s *ListAppInfoResponseBodyAppInfoList) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *ListAppInfoResponseBodyAppInfoList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAppInfoResponseBodyAppInfoList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListAppInfoResponseBodyAppInfoList) GetStatus() *string {
	return s.Status
}

func (s *ListAppInfoResponseBodyAppInfoList) GetType() *string {
	return s.Type
}

func (s *ListAppInfoResponseBodyAppInfoList) SetAppId(v string) *ListAppInfoResponseBodyAppInfoList {
	s.AppId = &v
	return s
}

func (s *ListAppInfoResponseBodyAppInfoList) SetAppName(v string) *ListAppInfoResponseBodyAppInfoList {
	s.AppName = &v
	return s
}

func (s *ListAppInfoResponseBodyAppInfoList) SetCreationTime(v string) *ListAppInfoResponseBodyAppInfoList {
	s.CreationTime = &v
	return s
}

func (s *ListAppInfoResponseBodyAppInfoList) SetDescription(v string) *ListAppInfoResponseBodyAppInfoList {
	s.Description = &v
	return s
}

func (s *ListAppInfoResponseBodyAppInfoList) SetModificationTime(v string) *ListAppInfoResponseBodyAppInfoList {
	s.ModificationTime = &v
	return s
}

func (s *ListAppInfoResponseBodyAppInfoList) SetRegionId(v string) *ListAppInfoResponseBodyAppInfoList {
	s.RegionId = &v
	return s
}

func (s *ListAppInfoResponseBodyAppInfoList) SetResourceGroupId(v string) *ListAppInfoResponseBodyAppInfoList {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAppInfoResponseBodyAppInfoList) SetStatus(v string) *ListAppInfoResponseBodyAppInfoList {
	s.Status = &v
	return s
}

func (s *ListAppInfoResponseBodyAppInfoList) SetType(v string) *ListAppInfoResponseBodyAppInfoList {
	s.Type = &v
	return s
}

func (s *ListAppInfoResponseBodyAppInfoList) Validate() error {
	return dara.Validate(s)
}
