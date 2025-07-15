// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopOversoldUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeDesktopOversoldUserGroupResponseBody
	GetCount() *int32
	SetData(v []*DescribeDesktopOversoldUserGroupResponseBodyData) *DescribeDesktopOversoldUserGroupResponseBody
	GetData() []*DescribeDesktopOversoldUserGroupResponseBodyData
	SetNextToken(v string) *DescribeDesktopOversoldUserGroupResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeDesktopOversoldUserGroupResponseBody
	GetRequestId() *string
}

type DescribeDesktopOversoldUserGroupResponseBody struct {
	Count     *int32                                              `json:"Count,omitempty" xml:"Count,omitempty"`
	Data      []*DescribeDesktopOversoldUserGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	NextToken *string                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDesktopOversoldUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopOversoldUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopOversoldUserGroupResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeDesktopOversoldUserGroupResponseBody) GetData() []*DescribeDesktopOversoldUserGroupResponseBodyData {
	return s.Data
}

func (s *DescribeDesktopOversoldUserGroupResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDesktopOversoldUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDesktopOversoldUserGroupResponseBody) SetCount(v int32) *DescribeDesktopOversoldUserGroupResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeDesktopOversoldUserGroupResponseBody) SetData(v []*DescribeDesktopOversoldUserGroupResponseBodyData) *DescribeDesktopOversoldUserGroupResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDesktopOversoldUserGroupResponseBody) SetNextToken(v string) *DescribeDesktopOversoldUserGroupResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopOversoldUserGroupResponseBody) SetRequestId(v string) *DescribeDesktopOversoldUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopOversoldUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDesktopOversoldUserGroupResponseBodyData struct {
	ImageId         *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OversoldGroupId *string `json:"OversoldGroupId,omitempty" xml:"OversoldGroupId,omitempty"`
	PolicyGroupId   *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	UserGroupId     *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DescribeDesktopOversoldUserGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopOversoldUserGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDesktopOversoldUserGroupResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeDesktopOversoldUserGroupResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DescribeDesktopOversoldUserGroupResponseBodyData) GetOversoldGroupId() *string {
	return s.OversoldGroupId
}

func (s *DescribeDesktopOversoldUserGroupResponseBodyData) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *DescribeDesktopOversoldUserGroupResponseBodyData) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *DescribeDesktopOversoldUserGroupResponseBodyData) SetImageId(v string) *DescribeDesktopOversoldUserGroupResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *DescribeDesktopOversoldUserGroupResponseBodyData) SetName(v string) *DescribeDesktopOversoldUserGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeDesktopOversoldUserGroupResponseBodyData) SetOversoldGroupId(v string) *DescribeDesktopOversoldUserGroupResponseBodyData {
	s.OversoldGroupId = &v
	return s
}

func (s *DescribeDesktopOversoldUserGroupResponseBodyData) SetPolicyGroupId(v string) *DescribeDesktopOversoldUserGroupResponseBodyData {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeDesktopOversoldUserGroupResponseBodyData) SetUserGroupId(v string) *DescribeDesktopOversoldUserGroupResponseBodyData {
	s.UserGroupId = &v
	return s
}

func (s *DescribeDesktopOversoldUserGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
