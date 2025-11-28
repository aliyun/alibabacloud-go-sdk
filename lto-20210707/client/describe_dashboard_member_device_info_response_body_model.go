// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDashboardMemberDeviceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDashboardMemberDeviceInfoResponseBody
	GetCode() *string
	SetData(v []*DescribeDashboardMemberDeviceInfoResponseBodyData) *DescribeDashboardMemberDeviceInfoResponseBody
	GetData() []*DescribeDashboardMemberDeviceInfoResponseBodyData
	SetHttpStatusCode(v int32) *DescribeDashboardMemberDeviceInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeDashboardMemberDeviceInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDashboardMemberDeviceInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDashboardMemberDeviceInfoResponseBody
	GetSuccess() *bool
}

type DescribeDashboardMemberDeviceInfoResponseBody struct {
	Code           *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*DescribeDashboardMemberDeviceInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDashboardMemberDeviceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardMemberDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDashboardMemberDeviceInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDashboardMemberDeviceInfoResponseBody) GetData() []*DescribeDashboardMemberDeviceInfoResponseBodyData {
	return s.Data
}

func (s *DescribeDashboardMemberDeviceInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDashboardMemberDeviceInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDashboardMemberDeviceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDashboardMemberDeviceInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDashboardMemberDeviceInfoResponseBody) SetCode(v string) *DescribeDashboardMemberDeviceInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDashboardMemberDeviceInfoResponseBody) SetData(v []*DescribeDashboardMemberDeviceInfoResponseBodyData) *DescribeDashboardMemberDeviceInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDashboardMemberDeviceInfoResponseBody) SetHttpStatusCode(v int32) *DescribeDashboardMemberDeviceInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDashboardMemberDeviceInfoResponseBody) SetMessage(v string) *DescribeDashboardMemberDeviceInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDashboardMemberDeviceInfoResponseBody) SetRequestId(v string) *DescribeDashboardMemberDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDashboardMemberDeviceInfoResponseBody) SetSuccess(v bool) *DescribeDashboardMemberDeviceInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDashboardMemberDeviceInfoResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDashboardMemberDeviceInfoResponseBodyData struct {
	MemberInfoList []*DescribeDashboardMemberDeviceInfoResponseBodyDataMemberInfoList `json:"MemberInfoList,omitempty" xml:"MemberInfoList,omitempty" type:"Repeated"`
	MemberName     *string                                                            `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
}

func (s DescribeDashboardMemberDeviceInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardMemberDeviceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDashboardMemberDeviceInfoResponseBodyData) GetMemberInfoList() []*DescribeDashboardMemberDeviceInfoResponseBodyDataMemberInfoList {
	return s.MemberInfoList
}

func (s *DescribeDashboardMemberDeviceInfoResponseBodyData) GetMemberName() *string {
	return s.MemberName
}

func (s *DescribeDashboardMemberDeviceInfoResponseBodyData) SetMemberInfoList(v []*DescribeDashboardMemberDeviceInfoResponseBodyDataMemberInfoList) *DescribeDashboardMemberDeviceInfoResponseBodyData {
	s.MemberInfoList = v
	return s
}

func (s *DescribeDashboardMemberDeviceInfoResponseBodyData) SetMemberName(v string) *DescribeDashboardMemberDeviceInfoResponseBodyData {
	s.MemberName = &v
	return s
}

func (s *DescribeDashboardMemberDeviceInfoResponseBodyData) Validate() error {
	if s.MemberInfoList != nil {
		for _, item := range s.MemberInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDashboardMemberDeviceInfoResponseBodyDataMemberInfoList struct {
	DeviceCount *int32 `json:"DeviceCount,omitempty" xml:"DeviceCount,omitempty"`
	Time        *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeDashboardMemberDeviceInfoResponseBodyDataMemberInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardMemberDeviceInfoResponseBodyDataMemberInfoList) GoString() string {
	return s.String()
}

func (s *DescribeDashboardMemberDeviceInfoResponseBodyDataMemberInfoList) GetDeviceCount() *int32 {
	return s.DeviceCount
}

func (s *DescribeDashboardMemberDeviceInfoResponseBodyDataMemberInfoList) GetTime() *int64 {
	return s.Time
}

func (s *DescribeDashboardMemberDeviceInfoResponseBodyDataMemberInfoList) SetDeviceCount(v int32) *DescribeDashboardMemberDeviceInfoResponseBodyDataMemberInfoList {
	s.DeviceCount = &v
	return s
}

func (s *DescribeDashboardMemberDeviceInfoResponseBodyDataMemberInfoList) SetTime(v int64) *DescribeDashboardMemberDeviceInfoResponseBodyDataMemberInfoList {
	s.Time = &v
	return s
}

func (s *DescribeDashboardMemberDeviceInfoResponseBodyDataMemberInfoList) Validate() error {
	return dara.Validate(s)
}
