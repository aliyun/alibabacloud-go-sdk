// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDashboardMemberApiInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDashboardMemberApiInfoResponseBody
	GetCode() *string
	SetData(v []*DescribeDashboardMemberApiInfoResponseBodyData) *DescribeDashboardMemberApiInfoResponseBody
	GetData() []*DescribeDashboardMemberApiInfoResponseBodyData
	SetHttpStatusCode(v int32) *DescribeDashboardMemberApiInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeDashboardMemberApiInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDashboardMemberApiInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDashboardMemberApiInfoResponseBody
	GetSuccess() *bool
}

type DescribeDashboardMemberApiInfoResponseBody struct {
	Code           *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*DescribeDashboardMemberApiInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDashboardMemberApiInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardMemberApiInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDashboardMemberApiInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDashboardMemberApiInfoResponseBody) GetData() []*DescribeDashboardMemberApiInfoResponseBodyData {
	return s.Data
}

func (s *DescribeDashboardMemberApiInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDashboardMemberApiInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDashboardMemberApiInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDashboardMemberApiInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDashboardMemberApiInfoResponseBody) SetCode(v string) *DescribeDashboardMemberApiInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDashboardMemberApiInfoResponseBody) SetData(v []*DescribeDashboardMemberApiInfoResponseBodyData) *DescribeDashboardMemberApiInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDashboardMemberApiInfoResponseBody) SetHttpStatusCode(v int32) *DescribeDashboardMemberApiInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDashboardMemberApiInfoResponseBody) SetMessage(v string) *DescribeDashboardMemberApiInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDashboardMemberApiInfoResponseBody) SetRequestId(v string) *DescribeDashboardMemberApiInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDashboardMemberApiInfoResponseBody) SetSuccess(v bool) *DescribeDashboardMemberApiInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDashboardMemberApiInfoResponseBody) Validate() error {
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

type DescribeDashboardMemberApiInfoResponseBodyData struct {
	MemberInfoList []*DescribeDashboardMemberApiInfoResponseBodyDataMemberInfoList `json:"MemberInfoList,omitempty" xml:"MemberInfoList,omitempty" type:"Repeated"`
	MemberName     *string                                                         `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
}

func (s DescribeDashboardMemberApiInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardMemberApiInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDashboardMemberApiInfoResponseBodyData) GetMemberInfoList() []*DescribeDashboardMemberApiInfoResponseBodyDataMemberInfoList {
	return s.MemberInfoList
}

func (s *DescribeDashboardMemberApiInfoResponseBodyData) GetMemberName() *string {
	return s.MemberName
}

func (s *DescribeDashboardMemberApiInfoResponseBodyData) SetMemberInfoList(v []*DescribeDashboardMemberApiInfoResponseBodyDataMemberInfoList) *DescribeDashboardMemberApiInfoResponseBodyData {
	s.MemberInfoList = v
	return s
}

func (s *DescribeDashboardMemberApiInfoResponseBodyData) SetMemberName(v string) *DescribeDashboardMemberApiInfoResponseBodyData {
	s.MemberName = &v
	return s
}

func (s *DescribeDashboardMemberApiInfoResponseBodyData) Validate() error {
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

type DescribeDashboardMemberApiInfoResponseBodyDataMemberInfoList struct {
	ApiInvokeCount *int64 `json:"ApiInvokeCount,omitempty" xml:"ApiInvokeCount,omitempty"`
	Time           *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeDashboardMemberApiInfoResponseBodyDataMemberInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardMemberApiInfoResponseBodyDataMemberInfoList) GoString() string {
	return s.String()
}

func (s *DescribeDashboardMemberApiInfoResponseBodyDataMemberInfoList) GetApiInvokeCount() *int64 {
	return s.ApiInvokeCount
}

func (s *DescribeDashboardMemberApiInfoResponseBodyDataMemberInfoList) GetTime() *int64 {
	return s.Time
}

func (s *DescribeDashboardMemberApiInfoResponseBodyDataMemberInfoList) SetApiInvokeCount(v int64) *DescribeDashboardMemberApiInfoResponseBodyDataMemberInfoList {
	s.ApiInvokeCount = &v
	return s
}

func (s *DescribeDashboardMemberApiInfoResponseBodyDataMemberInfoList) SetTime(v int64) *DescribeDashboardMemberApiInfoResponseBodyDataMemberInfoList {
	s.Time = &v
	return s
}

func (s *DescribeDashboardMemberApiInfoResponseBodyDataMemberInfoList) Validate() error {
	return dara.Validate(s)
}
