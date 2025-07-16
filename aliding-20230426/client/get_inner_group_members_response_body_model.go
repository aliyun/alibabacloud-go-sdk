// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInnerGroupMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHasMore(v bool) *GetInnerGroupMembersResponseBody
	GetHasMore() *bool
	SetNextToken(v string) *GetInnerGroupMembersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetInnerGroupMembersResponseBody
	GetRequestId() *string
	SetUserIds(v []*string) *GetInnerGroupMembersResponseBody
	GetUserIds() []*string
	SetVendorRequestId(v string) *GetInnerGroupMembersResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetInnerGroupMembersResponseBody
	GetVendorType() *string
}

type GetInnerGroupMembersResponseBody struct {
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// cdf***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// ["012345"]
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetInnerGroupMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInnerGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *GetInnerGroupMembersResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *GetInnerGroupMembersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetInnerGroupMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInnerGroupMembersResponseBody) GetUserIds() []*string {
	return s.UserIds
}

func (s *GetInnerGroupMembersResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetInnerGroupMembersResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetInnerGroupMembersResponseBody) SetHasMore(v bool) *GetInnerGroupMembersResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetInnerGroupMembersResponseBody) SetNextToken(v string) *GetInnerGroupMembersResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetInnerGroupMembersResponseBody) SetRequestId(v string) *GetInnerGroupMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInnerGroupMembersResponseBody) SetUserIds(v []*string) *GetInnerGroupMembersResponseBody {
	s.UserIds = v
	return s
}

func (s *GetInnerGroupMembersResponseBody) SetVendorRequestId(v string) *GetInnerGroupMembersResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetInnerGroupMembersResponseBody) SetVendorType(v string) *GetInnerGroupMembersResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetInnerGroupMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
