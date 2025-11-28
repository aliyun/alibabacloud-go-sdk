// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMemberTotalStatInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeMemberTotalStatInfoResponseBody
	GetCode() *string
	SetData(v *DescribeMemberTotalStatInfoResponseBodyData) *DescribeMemberTotalStatInfoResponseBody
	GetData() *DescribeMemberTotalStatInfoResponseBodyData
	SetHttpStatusCode(v int32) *DescribeMemberTotalStatInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeMemberTotalStatInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeMemberTotalStatInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMemberTotalStatInfoResponseBody
	GetSuccess() *bool
}

type DescribeMemberTotalStatInfoResponseBody struct {
	Code           *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeMemberTotalStatInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMemberTotalStatInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMemberTotalStatInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMemberTotalStatInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMemberTotalStatInfoResponseBody) GetData() *DescribeMemberTotalStatInfoResponseBodyData {
	return s.Data
}

func (s *DescribeMemberTotalStatInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeMemberTotalStatInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMemberTotalStatInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMemberTotalStatInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMemberTotalStatInfoResponseBody) SetCode(v string) *DescribeMemberTotalStatInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMemberTotalStatInfoResponseBody) SetData(v *DescribeMemberTotalStatInfoResponseBodyData) *DescribeMemberTotalStatInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMemberTotalStatInfoResponseBody) SetHttpStatusCode(v int32) *DescribeMemberTotalStatInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeMemberTotalStatInfoResponseBody) SetMessage(v string) *DescribeMemberTotalStatInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMemberTotalStatInfoResponseBody) SetRequestId(v string) *DescribeMemberTotalStatInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMemberTotalStatInfoResponseBody) SetSuccess(v bool) *DescribeMemberTotalStatInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMemberTotalStatInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMemberTotalStatInfoResponseBodyData struct {
	AuthorizedCount *int64 `json:"AuthorizedCount,omitempty" xml:"AuthorizedCount,omitempty"`
	UsedCount       *int64 `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
}

func (s DescribeMemberTotalStatInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMemberTotalStatInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMemberTotalStatInfoResponseBodyData) GetAuthorizedCount() *int64 {
	return s.AuthorizedCount
}

func (s *DescribeMemberTotalStatInfoResponseBodyData) GetUsedCount() *int64 {
	return s.UsedCount
}

func (s *DescribeMemberTotalStatInfoResponseBodyData) SetAuthorizedCount(v int64) *DescribeMemberTotalStatInfoResponseBodyData {
	s.AuthorizedCount = &v
	return s
}

func (s *DescribeMemberTotalStatInfoResponseBodyData) SetUsedCount(v int64) *DescribeMemberTotalStatInfoResponseBodyData {
	s.UsedCount = &v
	return s
}

func (s *DescribeMemberTotalStatInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
