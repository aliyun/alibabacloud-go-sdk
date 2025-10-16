// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopOversoldUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeDesktopOversoldUserResponseBody
	GetCount() *int32
	SetData(v []*DescribeDesktopOversoldUserResponseBodyData) *DescribeDesktopOversoldUserResponseBody
	GetData() []*DescribeDesktopOversoldUserResponseBodyData
	SetNextToken(v string) *DescribeDesktopOversoldUserResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeDesktopOversoldUserResponseBody
	GetRequestId() *string
}

type DescribeDesktopOversoldUserResponseBody struct {
	Count     *int32                                         `json:"Count,omitempty" xml:"Count,omitempty"`
	Data      []*DescribeDesktopOversoldUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	NextToken *string                                        `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDesktopOversoldUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopOversoldUserResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopOversoldUserResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeDesktopOversoldUserResponseBody) GetData() []*DescribeDesktopOversoldUserResponseBodyData {
	return s.Data
}

func (s *DescribeDesktopOversoldUserResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDesktopOversoldUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDesktopOversoldUserResponseBody) SetCount(v int32) *DescribeDesktopOversoldUserResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeDesktopOversoldUserResponseBody) SetData(v []*DescribeDesktopOversoldUserResponseBodyData) *DescribeDesktopOversoldUserResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDesktopOversoldUserResponseBody) SetNextToken(v string) *DescribeDesktopOversoldUserResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopOversoldUserResponseBody) SetRequestId(v string) *DescribeDesktopOversoldUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopOversoldUserResponseBody) Validate() error {
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

type DescribeDesktopOversoldUserResponseBodyData struct {
	EndUserId       *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	OversoldGroupId *string `json:"OversoldGroupId,omitempty" xml:"OversoldGroupId,omitempty"`
	UserDesktopId   *string `json:"UserDesktopId,omitempty" xml:"UserDesktopId,omitempty"`
	UserGroupId     *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DescribeDesktopOversoldUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopOversoldUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDesktopOversoldUserResponseBodyData) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeDesktopOversoldUserResponseBodyData) GetOversoldGroupId() *string {
	return s.OversoldGroupId
}

func (s *DescribeDesktopOversoldUserResponseBodyData) GetUserDesktopId() *string {
	return s.UserDesktopId
}

func (s *DescribeDesktopOversoldUserResponseBodyData) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *DescribeDesktopOversoldUserResponseBodyData) SetEndUserId(v string) *DescribeDesktopOversoldUserResponseBodyData {
	s.EndUserId = &v
	return s
}

func (s *DescribeDesktopOversoldUserResponseBodyData) SetOversoldGroupId(v string) *DescribeDesktopOversoldUserResponseBodyData {
	s.OversoldGroupId = &v
	return s
}

func (s *DescribeDesktopOversoldUserResponseBodyData) SetUserDesktopId(v string) *DescribeDesktopOversoldUserResponseBodyData {
	s.UserDesktopId = &v
	return s
}

func (s *DescribeDesktopOversoldUserResponseBodyData) SetUserGroupId(v string) *DescribeDesktopOversoldUserResponseBodyData {
	s.UserGroupId = &v
	return s
}

func (s *DescribeDesktopOversoldUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}
