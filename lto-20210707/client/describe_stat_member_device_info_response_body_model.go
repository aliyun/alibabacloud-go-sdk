// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStatMemberDeviceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeStatMemberDeviceInfoResponseBody
	GetCode() *string
	SetData(v []*DescribeStatMemberDeviceInfoResponseBodyData) *DescribeStatMemberDeviceInfoResponseBody
	GetData() []*DescribeStatMemberDeviceInfoResponseBodyData
	SetHttpStatusCode(v int32) *DescribeStatMemberDeviceInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeStatMemberDeviceInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeStatMemberDeviceInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeStatMemberDeviceInfoResponseBody
	GetSuccess() *bool
}

type DescribeStatMemberDeviceInfoResponseBody struct {
	Code           *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*DescribeStatMemberDeviceInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeStatMemberDeviceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStatMemberDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStatMemberDeviceInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeStatMemberDeviceInfoResponseBody) GetData() []*DescribeStatMemberDeviceInfoResponseBodyData {
	return s.Data
}

func (s *DescribeStatMemberDeviceInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeStatMemberDeviceInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeStatMemberDeviceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStatMemberDeviceInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeStatMemberDeviceInfoResponseBody) SetCode(v string) *DescribeStatMemberDeviceInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeStatMemberDeviceInfoResponseBody) SetData(v []*DescribeStatMemberDeviceInfoResponseBodyData) *DescribeStatMemberDeviceInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeStatMemberDeviceInfoResponseBody) SetHttpStatusCode(v int32) *DescribeStatMemberDeviceInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeStatMemberDeviceInfoResponseBody) SetMessage(v string) *DescribeStatMemberDeviceInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeStatMemberDeviceInfoResponseBody) SetRequestId(v string) *DescribeStatMemberDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStatMemberDeviceInfoResponseBody) SetSuccess(v bool) *DescribeStatMemberDeviceInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeStatMemberDeviceInfoResponseBody) Validate() error {
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

type DescribeStatMemberDeviceInfoResponseBodyData struct {
	AuthorizedCount *string `json:"AuthorizedCount,omitempty" xml:"AuthorizedCount,omitempty"`
	BizChainCount   *string `json:"BizChainCount,omitempty" xml:"BizChainCount,omitempty"`
	MemberId        *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName      *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	UsedCount       *string `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
}

func (s DescribeStatMemberDeviceInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeStatMemberDeviceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeStatMemberDeviceInfoResponseBodyData) GetAuthorizedCount() *string {
	return s.AuthorizedCount
}

func (s *DescribeStatMemberDeviceInfoResponseBodyData) GetBizChainCount() *string {
	return s.BizChainCount
}

func (s *DescribeStatMemberDeviceInfoResponseBodyData) GetMemberId() *string {
	return s.MemberId
}

func (s *DescribeStatMemberDeviceInfoResponseBodyData) GetMemberName() *string {
	return s.MemberName
}

func (s *DescribeStatMemberDeviceInfoResponseBodyData) GetUsedCount() *string {
	return s.UsedCount
}

func (s *DescribeStatMemberDeviceInfoResponseBodyData) SetAuthorizedCount(v string) *DescribeStatMemberDeviceInfoResponseBodyData {
	s.AuthorizedCount = &v
	return s
}

func (s *DescribeStatMemberDeviceInfoResponseBodyData) SetBizChainCount(v string) *DescribeStatMemberDeviceInfoResponseBodyData {
	s.BizChainCount = &v
	return s
}

func (s *DescribeStatMemberDeviceInfoResponseBodyData) SetMemberId(v string) *DescribeStatMemberDeviceInfoResponseBodyData {
	s.MemberId = &v
	return s
}

func (s *DescribeStatMemberDeviceInfoResponseBodyData) SetMemberName(v string) *DescribeStatMemberDeviceInfoResponseBodyData {
	s.MemberName = &v
	return s
}

func (s *DescribeStatMemberDeviceInfoResponseBodyData) SetUsedCount(v string) *DescribeStatMemberDeviceInfoResponseBodyData {
	s.UsedCount = &v
	return s
}

func (s *DescribeStatMemberDeviceInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
