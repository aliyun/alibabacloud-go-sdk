// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserConnectTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeUserConnectTimeResponseBody
	GetCount() *int32
	SetData(v []*DescribeUserConnectTimeResponseBodyData) *DescribeUserConnectTimeResponseBody
	GetData() []*DescribeUserConnectTimeResponseBodyData
	SetNextToken(v string) *DescribeUserConnectTimeResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeUserConnectTimeResponseBody
	GetRequestId() *string
}

type DescribeUserConnectTimeResponseBody struct {
	Count     *int32                                     `json:"Count,omitempty" xml:"Count,omitempty"`
	Data      []*DescribeUserConnectTimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	NextToken *string                                    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserConnectTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserConnectTimeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserConnectTimeResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeUserConnectTimeResponseBody) GetData() []*DescribeUserConnectTimeResponseBodyData {
	return s.Data
}

func (s *DescribeUserConnectTimeResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeUserConnectTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserConnectTimeResponseBody) SetCount(v int32) *DescribeUserConnectTimeResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeUserConnectTimeResponseBody) SetData(v []*DescribeUserConnectTimeResponseBodyData) *DescribeUserConnectTimeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeUserConnectTimeResponseBody) SetNextToken(v string) *DescribeUserConnectTimeResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeUserConnectTimeResponseBody) SetRequestId(v string) *DescribeUserConnectTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserConnectTimeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUserConnectTimeResponseBodyData struct {
	EndConnectTime   *string `json:"EndConnectTime,omitempty" xml:"EndConnectTime,omitempty"`
	EndUserId        *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	OversoldGroupId  *string `json:"OversoldGroupId,omitempty" xml:"OversoldGroupId,omitempty"`
	StartConnectTime *string `json:"StartConnectTime,omitempty" xml:"StartConnectTime,omitempty"`
	UserDesktopId    *string `json:"UserDesktopId,omitempty" xml:"UserDesktopId,omitempty"`
	UserGroupId      *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DescribeUserConnectTimeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserConnectTimeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeUserConnectTimeResponseBodyData) GetEndConnectTime() *string {
	return s.EndConnectTime
}

func (s *DescribeUserConnectTimeResponseBodyData) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeUserConnectTimeResponseBodyData) GetOversoldGroupId() *string {
	return s.OversoldGroupId
}

func (s *DescribeUserConnectTimeResponseBodyData) GetStartConnectTime() *string {
	return s.StartConnectTime
}

func (s *DescribeUserConnectTimeResponseBodyData) GetUserDesktopId() *string {
	return s.UserDesktopId
}

func (s *DescribeUserConnectTimeResponseBodyData) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *DescribeUserConnectTimeResponseBodyData) SetEndConnectTime(v string) *DescribeUserConnectTimeResponseBodyData {
	s.EndConnectTime = &v
	return s
}

func (s *DescribeUserConnectTimeResponseBodyData) SetEndUserId(v string) *DescribeUserConnectTimeResponseBodyData {
	s.EndUserId = &v
	return s
}

func (s *DescribeUserConnectTimeResponseBodyData) SetOversoldGroupId(v string) *DescribeUserConnectTimeResponseBodyData {
	s.OversoldGroupId = &v
	return s
}

func (s *DescribeUserConnectTimeResponseBodyData) SetStartConnectTime(v string) *DescribeUserConnectTimeResponseBodyData {
	s.StartConnectTime = &v
	return s
}

func (s *DescribeUserConnectTimeResponseBodyData) SetUserDesktopId(v string) *DescribeUserConnectTimeResponseBodyData {
	s.UserDesktopId = &v
	return s
}

func (s *DescribeUserConnectTimeResponseBodyData) SetUserGroupId(v string) *DescribeUserConnectTimeResponseBodyData {
	s.UserGroupId = &v
	return s
}

func (s *DescribeUserConnectTimeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
