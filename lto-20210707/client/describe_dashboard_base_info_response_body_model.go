// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDashboardBaseInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDashboardBaseInfoResponseBody
	GetCode() *string
	SetData(v *DescribeDashboardBaseInfoResponseBodyData) *DescribeDashboardBaseInfoResponseBody
	GetData() *DescribeDashboardBaseInfoResponseBodyData
	SetHttpStatusCode(v int32) *DescribeDashboardBaseInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeDashboardBaseInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDashboardBaseInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDashboardBaseInfoResponseBody
	GetSuccess() *bool
}

type DescribeDashboardBaseInfoResponseBody struct {
	Code           *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeDashboardBaseInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDashboardBaseInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardBaseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDashboardBaseInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDashboardBaseInfoResponseBody) GetData() *DescribeDashboardBaseInfoResponseBodyData {
	return s.Data
}

func (s *DescribeDashboardBaseInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDashboardBaseInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDashboardBaseInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDashboardBaseInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDashboardBaseInfoResponseBody) SetCode(v string) *DescribeDashboardBaseInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDashboardBaseInfoResponseBody) SetData(v *DescribeDashboardBaseInfoResponseBodyData) *DescribeDashboardBaseInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDashboardBaseInfoResponseBody) SetHttpStatusCode(v int32) *DescribeDashboardBaseInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDashboardBaseInfoResponseBody) SetMessage(v string) *DescribeDashboardBaseInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDashboardBaseInfoResponseBody) SetRequestId(v string) *DescribeDashboardBaseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDashboardBaseInfoResponseBody) SetSuccess(v bool) *DescribeDashboardBaseInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDashboardBaseInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDashboardBaseInfoResponseBodyData struct {
	ApiInvokeCount *int64 `json:"ApiInvokeCount,omitempty" xml:"ApiInvokeCount,omitempty"`
	BizChainCount  *int32 `json:"BizChainCount,omitempty" xml:"BizChainCount,omitempty"`
	DeviceCount    *int32 `json:"DeviceCount,omitempty" xml:"DeviceCount,omitempty"`
	MemberCount    *int32 `json:"MemberCount,omitempty" xml:"MemberCount,omitempty"`
}

func (s DescribeDashboardBaseInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardBaseInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDashboardBaseInfoResponseBodyData) GetApiInvokeCount() *int64 {
	return s.ApiInvokeCount
}

func (s *DescribeDashboardBaseInfoResponseBodyData) GetBizChainCount() *int32 {
	return s.BizChainCount
}

func (s *DescribeDashboardBaseInfoResponseBodyData) GetDeviceCount() *int32 {
	return s.DeviceCount
}

func (s *DescribeDashboardBaseInfoResponseBodyData) GetMemberCount() *int32 {
	return s.MemberCount
}

func (s *DescribeDashboardBaseInfoResponseBodyData) SetApiInvokeCount(v int64) *DescribeDashboardBaseInfoResponseBodyData {
	s.ApiInvokeCount = &v
	return s
}

func (s *DescribeDashboardBaseInfoResponseBodyData) SetBizChainCount(v int32) *DescribeDashboardBaseInfoResponseBodyData {
	s.BizChainCount = &v
	return s
}

func (s *DescribeDashboardBaseInfoResponseBodyData) SetDeviceCount(v int32) *DescribeDashboardBaseInfoResponseBodyData {
	s.DeviceCount = &v
	return s
}

func (s *DescribeDashboardBaseInfoResponseBodyData) SetMemberCount(v int32) *DescribeDashboardBaseInfoResponseBodyData {
	s.MemberCount = &v
	return s
}

func (s *DescribeDashboardBaseInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
