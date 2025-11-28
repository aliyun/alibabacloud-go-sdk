// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTotalStatInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeTotalStatInfoResponseBody
	GetCode() *string
	SetData(v *DescribeTotalStatInfoResponseBodyData) *DescribeTotalStatInfoResponseBody
	GetData() *DescribeTotalStatInfoResponseBodyData
	SetHttpStatusCode(v int32) *DescribeTotalStatInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeTotalStatInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeTotalStatInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeTotalStatInfoResponseBody
	GetSuccess() *bool
}

type DescribeTotalStatInfoResponseBody struct {
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeTotalStatInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeTotalStatInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTotalStatInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTotalStatInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeTotalStatInfoResponseBody) GetData() *DescribeTotalStatInfoResponseBodyData {
	return s.Data
}

func (s *DescribeTotalStatInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeTotalStatInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeTotalStatInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTotalStatInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeTotalStatInfoResponseBody) SetCode(v string) *DescribeTotalStatInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTotalStatInfoResponseBody) SetData(v *DescribeTotalStatInfoResponseBodyData) *DescribeTotalStatInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTotalStatInfoResponseBody) SetHttpStatusCode(v int32) *DescribeTotalStatInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeTotalStatInfoResponseBody) SetMessage(v string) *DescribeTotalStatInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTotalStatInfoResponseBody) SetRequestId(v string) *DescribeTotalStatInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTotalStatInfoResponseBody) SetSuccess(v bool) *DescribeTotalStatInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTotalStatInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTotalStatInfoResponseBodyData struct {
	AuthorizedCount *int64 `json:"AuthorizedCount,omitempty" xml:"AuthorizedCount,omitempty"`
	TotalCount      *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UsedCount       *int64 `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
}

func (s DescribeTotalStatInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeTotalStatInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTotalStatInfoResponseBodyData) GetAuthorizedCount() *int64 {
	return s.AuthorizedCount
}

func (s *DescribeTotalStatInfoResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeTotalStatInfoResponseBodyData) GetUsedCount() *int64 {
	return s.UsedCount
}

func (s *DescribeTotalStatInfoResponseBodyData) SetAuthorizedCount(v int64) *DescribeTotalStatInfoResponseBodyData {
	s.AuthorizedCount = &v
	return s
}

func (s *DescribeTotalStatInfoResponseBodyData) SetTotalCount(v int64) *DescribeTotalStatInfoResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeTotalStatInfoResponseBodyData) SetUsedCount(v int64) *DescribeTotalStatInfoResponseBodyData {
	s.UsedCount = &v
	return s
}

func (s *DescribeTotalStatInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
