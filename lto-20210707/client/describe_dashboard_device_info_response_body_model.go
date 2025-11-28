// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDashboardDeviceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDashboardDeviceInfoResponseBody
	GetCode() *string
	SetData(v *DescribeDashboardDeviceInfoResponseBodyData) *DescribeDashboardDeviceInfoResponseBody
	GetData() *DescribeDashboardDeviceInfoResponseBodyData
	SetHttpStatusCode(v int32) *DescribeDashboardDeviceInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeDashboardDeviceInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDashboardDeviceInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDashboardDeviceInfoResponseBody
	GetSuccess() *bool
}

type DescribeDashboardDeviceInfoResponseBody struct {
	Code           *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeDashboardDeviceInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDashboardDeviceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDashboardDeviceInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDashboardDeviceInfoResponseBody) GetData() *DescribeDashboardDeviceInfoResponseBodyData {
	return s.Data
}

func (s *DescribeDashboardDeviceInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDashboardDeviceInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDashboardDeviceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDashboardDeviceInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDashboardDeviceInfoResponseBody) SetCode(v string) *DescribeDashboardDeviceInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDashboardDeviceInfoResponseBody) SetData(v *DescribeDashboardDeviceInfoResponseBodyData) *DescribeDashboardDeviceInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDashboardDeviceInfoResponseBody) SetHttpStatusCode(v int32) *DescribeDashboardDeviceInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDashboardDeviceInfoResponseBody) SetMessage(v string) *DescribeDashboardDeviceInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDashboardDeviceInfoResponseBody) SetRequestId(v string) *DescribeDashboardDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDashboardDeviceInfoResponseBody) SetSuccess(v bool) *DescribeDashboardDeviceInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDashboardDeviceInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDashboardDeviceInfoResponseBodyData struct {
	AuthorizedCount *int64 `json:"AuthorizedCount,omitempty" xml:"AuthorizedCount,omitempty"`
	UsedCount       *int64 `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
}

func (s DescribeDashboardDeviceInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardDeviceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDashboardDeviceInfoResponseBodyData) GetAuthorizedCount() *int64 {
	return s.AuthorizedCount
}

func (s *DescribeDashboardDeviceInfoResponseBodyData) GetUsedCount() *int64 {
	return s.UsedCount
}

func (s *DescribeDashboardDeviceInfoResponseBodyData) SetAuthorizedCount(v int64) *DescribeDashboardDeviceInfoResponseBodyData {
	s.AuthorizedCount = &v
	return s
}

func (s *DescribeDashboardDeviceInfoResponseBodyData) SetUsedCount(v int64) *DescribeDashboardDeviceInfoResponseBodyData {
	s.UsedCount = &v
	return s
}

func (s *DescribeDashboardDeviceInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
