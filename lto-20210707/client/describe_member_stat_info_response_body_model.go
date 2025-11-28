// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMemberStatInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeMemberStatInfoResponseBody
	GetCode() *string
	SetData(v []*DescribeMemberStatInfoResponseBodyData) *DescribeMemberStatInfoResponseBody
	GetData() []*DescribeMemberStatInfoResponseBodyData
	SetHttpStatusCode(v int32) *DescribeMemberStatInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeMemberStatInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeMemberStatInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMemberStatInfoResponseBody
	GetSuccess() *bool
}

type DescribeMemberStatInfoResponseBody struct {
	Code           *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*DescribeMemberStatInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMemberStatInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMemberStatInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMemberStatInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMemberStatInfoResponseBody) GetData() []*DescribeMemberStatInfoResponseBodyData {
	return s.Data
}

func (s *DescribeMemberStatInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeMemberStatInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMemberStatInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMemberStatInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMemberStatInfoResponseBody) SetCode(v string) *DescribeMemberStatInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMemberStatInfoResponseBody) SetData(v []*DescribeMemberStatInfoResponseBodyData) *DescribeMemberStatInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMemberStatInfoResponseBody) SetHttpStatusCode(v int32) *DescribeMemberStatInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeMemberStatInfoResponseBody) SetMessage(v string) *DescribeMemberStatInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMemberStatInfoResponseBody) SetRequestId(v string) *DescribeMemberStatInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMemberStatInfoResponseBody) SetSuccess(v bool) *DescribeMemberStatInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMemberStatInfoResponseBody) Validate() error {
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

type DescribeMemberStatInfoResponseBodyData struct {
	AuthorizedCount *int64  `json:"AuthorizedCount,omitempty" xml:"AuthorizedCount,omitempty"`
	BizChainCount   *int32  `json:"BizChainCount,omitempty" xml:"BizChainCount,omitempty"`
	MemberId        *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName      *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	UsedCount       *int64  `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
}

func (s DescribeMemberStatInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMemberStatInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMemberStatInfoResponseBodyData) GetAuthorizedCount() *int64 {
	return s.AuthorizedCount
}

func (s *DescribeMemberStatInfoResponseBodyData) GetBizChainCount() *int32 {
	return s.BizChainCount
}

func (s *DescribeMemberStatInfoResponseBodyData) GetMemberId() *string {
	return s.MemberId
}

func (s *DescribeMemberStatInfoResponseBodyData) GetMemberName() *string {
	return s.MemberName
}

func (s *DescribeMemberStatInfoResponseBodyData) GetUsedCount() *int64 {
	return s.UsedCount
}

func (s *DescribeMemberStatInfoResponseBodyData) SetAuthorizedCount(v int64) *DescribeMemberStatInfoResponseBodyData {
	s.AuthorizedCount = &v
	return s
}

func (s *DescribeMemberStatInfoResponseBodyData) SetBizChainCount(v int32) *DescribeMemberStatInfoResponseBodyData {
	s.BizChainCount = &v
	return s
}

func (s *DescribeMemberStatInfoResponseBodyData) SetMemberId(v string) *DescribeMemberStatInfoResponseBodyData {
	s.MemberId = &v
	return s
}

func (s *DescribeMemberStatInfoResponseBodyData) SetMemberName(v string) *DescribeMemberStatInfoResponseBodyData {
	s.MemberName = &v
	return s
}

func (s *DescribeMemberStatInfoResponseBodyData) SetUsedCount(v int64) *DescribeMemberStatInfoResponseBodyData {
	s.UsedCount = &v
	return s
}

func (s *DescribeMemberStatInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
