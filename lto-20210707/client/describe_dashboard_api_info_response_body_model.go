// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDashboardApiInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDashboardApiInfoResponseBody
	GetCode() *string
	SetData(v *DescribeDashboardApiInfoResponseBodyData) *DescribeDashboardApiInfoResponseBody
	GetData() *DescribeDashboardApiInfoResponseBodyData
	SetHttpStatusCode(v int32) *DescribeDashboardApiInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeDashboardApiInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDashboardApiInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDashboardApiInfoResponseBody
	GetSuccess() *bool
}

type DescribeDashboardApiInfoResponseBody struct {
	Code           *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeDashboardApiInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDashboardApiInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardApiInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDashboardApiInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDashboardApiInfoResponseBody) GetData() *DescribeDashboardApiInfoResponseBodyData {
	return s.Data
}

func (s *DescribeDashboardApiInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDashboardApiInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDashboardApiInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDashboardApiInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDashboardApiInfoResponseBody) SetCode(v string) *DescribeDashboardApiInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDashboardApiInfoResponseBody) SetData(v *DescribeDashboardApiInfoResponseBodyData) *DescribeDashboardApiInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDashboardApiInfoResponseBody) SetHttpStatusCode(v int32) *DescribeDashboardApiInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDashboardApiInfoResponseBody) SetMessage(v string) *DescribeDashboardApiInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDashboardApiInfoResponseBody) SetRequestId(v string) *DescribeDashboardApiInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDashboardApiInfoResponseBody) SetSuccess(v bool) *DescribeDashboardApiInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDashboardApiInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDashboardApiInfoResponseBodyData struct {
	AuthorizedCount *int64 `json:"AuthorizedCount,omitempty" xml:"AuthorizedCount,omitempty"`
	UsedCount       *int64 `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
}

func (s DescribeDashboardApiInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardApiInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDashboardApiInfoResponseBodyData) GetAuthorizedCount() *int64 {
	return s.AuthorizedCount
}

func (s *DescribeDashboardApiInfoResponseBodyData) GetUsedCount() *int64 {
	return s.UsedCount
}

func (s *DescribeDashboardApiInfoResponseBodyData) SetAuthorizedCount(v int64) *DescribeDashboardApiInfoResponseBodyData {
	s.AuthorizedCount = &v
	return s
}

func (s *DescribeDashboardApiInfoResponseBodyData) SetUsedCount(v int64) *DescribeDashboardApiInfoResponseBodyData {
	s.UsedCount = &v
	return s
}

func (s *DescribeDashboardApiInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
