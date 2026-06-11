// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDevicePageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceList(v []*DevicePageResponseBodyDeviceList) *DevicePageResponseBody
	GetDeviceList() []*DevicePageResponseBodyDeviceList
	SetPageNumber(v int32) *DevicePageResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DevicePageResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DevicePageResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DevicePageResponseBody
	GetTotalCount() *int32
}

type DevicePageResponseBody struct {
	DeviceList []*DevicePageResponseBodyDeviceList `json:"DeviceList,omitempty" xml:"DeviceList,omitempty" type:"Repeated"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DevicePageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DevicePageResponseBody) GoString() string {
	return s.String()
}

func (s *DevicePageResponseBody) GetDeviceList() []*DevicePageResponseBodyDeviceList {
	return s.DeviceList
}

func (s *DevicePageResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DevicePageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DevicePageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DevicePageResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DevicePageResponseBody) SetDeviceList(v []*DevicePageResponseBodyDeviceList) *DevicePageResponseBody {
	s.DeviceList = v
	return s
}

func (s *DevicePageResponseBody) SetPageNumber(v int32) *DevicePageResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DevicePageResponseBody) SetPageSize(v int32) *DevicePageResponseBody {
	s.PageSize = &v
	return s
}

func (s *DevicePageResponseBody) SetRequestId(v string) *DevicePageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DevicePageResponseBody) SetTotalCount(v int32) *DevicePageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DevicePageResponseBody) Validate() error {
	if s.DeviceList != nil {
		for _, item := range s.DeviceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DevicePageResponseBodyDeviceList struct {
	ActiveTime     *string `json:"ActiveTime,omitempty" xml:"ActiveTime,omitempty"`
	AliyunUid      *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	DailyCount     *int32  `json:"DailyCount,omitempty" xml:"DailyCount,omitempty"`
	DeviceName     *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	OrderType      *int32  `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	SubAliyunUid   *string `json:"SubAliyunUid,omitempty" xml:"SubAliyunUid,omitempty"`
	TotalCount     *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UserActiveTime *string `json:"UserActiveTime,omitempty" xml:"UserActiveTime,omitempty"`
	WorkspaceId    *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DevicePageResponseBodyDeviceList) String() string {
	return dara.Prettify(s)
}

func (s DevicePageResponseBodyDeviceList) GoString() string {
	return s.String()
}

func (s *DevicePageResponseBodyDeviceList) GetActiveTime() *string {
	return s.ActiveTime
}

func (s *DevicePageResponseBodyDeviceList) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *DevicePageResponseBodyDeviceList) GetAppId() *string {
	return s.AppId
}

func (s *DevicePageResponseBodyDeviceList) GetDailyCount() *int32 {
	return s.DailyCount
}

func (s *DevicePageResponseBodyDeviceList) GetDeviceName() *string {
	return s.DeviceName
}

func (s *DevicePageResponseBodyDeviceList) GetOrderType() *int32 {
	return s.OrderType
}

func (s *DevicePageResponseBodyDeviceList) GetStatus() *int32 {
	return s.Status
}

func (s *DevicePageResponseBodyDeviceList) GetSubAliyunUid() *string {
	return s.SubAliyunUid
}

func (s *DevicePageResponseBodyDeviceList) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DevicePageResponseBodyDeviceList) GetUserActiveTime() *string {
	return s.UserActiveTime
}

func (s *DevicePageResponseBodyDeviceList) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DevicePageResponseBodyDeviceList) SetActiveTime(v string) *DevicePageResponseBodyDeviceList {
	s.ActiveTime = &v
	return s
}

func (s *DevicePageResponseBodyDeviceList) SetAliyunUid(v string) *DevicePageResponseBodyDeviceList {
	s.AliyunUid = &v
	return s
}

func (s *DevicePageResponseBodyDeviceList) SetAppId(v string) *DevicePageResponseBodyDeviceList {
	s.AppId = &v
	return s
}

func (s *DevicePageResponseBodyDeviceList) SetDailyCount(v int32) *DevicePageResponseBodyDeviceList {
	s.DailyCount = &v
	return s
}

func (s *DevicePageResponseBodyDeviceList) SetDeviceName(v string) *DevicePageResponseBodyDeviceList {
	s.DeviceName = &v
	return s
}

func (s *DevicePageResponseBodyDeviceList) SetOrderType(v int32) *DevicePageResponseBodyDeviceList {
	s.OrderType = &v
	return s
}

func (s *DevicePageResponseBodyDeviceList) SetStatus(v int32) *DevicePageResponseBodyDeviceList {
	s.Status = &v
	return s
}

func (s *DevicePageResponseBodyDeviceList) SetSubAliyunUid(v string) *DevicePageResponseBodyDeviceList {
	s.SubAliyunUid = &v
	return s
}

func (s *DevicePageResponseBodyDeviceList) SetTotalCount(v int32) *DevicePageResponseBodyDeviceList {
	s.TotalCount = &v
	return s
}

func (s *DevicePageResponseBodyDeviceList) SetUserActiveTime(v string) *DevicePageResponseBodyDeviceList {
	s.UserActiveTime = &v
	return s
}

func (s *DevicePageResponseBodyDeviceList) SetWorkspaceId(v string) *DevicePageResponseBodyDeviceList {
	s.WorkspaceId = &v
	return s
}

func (s *DevicePageResponseBodyDeviceList) Validate() error {
	return dara.Validate(s)
}
