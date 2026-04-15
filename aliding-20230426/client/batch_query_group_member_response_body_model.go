// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryGroupMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHasMore(v bool) *BatchQueryGroupMemberResponseBody
	GetHasMore() *bool
	SetMemberUserIds(v []*string) *BatchQueryGroupMemberResponseBody
	GetMemberUserIds() []*string
	SetNextToken(v string) *BatchQueryGroupMemberResponseBody
	GetNextToken() *string
	SetRequestId(v string) *BatchQueryGroupMemberResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *BatchQueryGroupMemberResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *BatchQueryGroupMemberResponseBody
	GetVendorType() *string
}

type BatchQueryGroupMemberResponseBody struct {
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// ["012345"]
	MemberUserIds []*string `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// weqrwereqsadqaadfafa
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s BatchQueryGroupMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryGroupMemberResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *BatchQueryGroupMemberResponseBody) GetMemberUserIds() []*string {
	return s.MemberUserIds
}

func (s *BatchQueryGroupMemberResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *BatchQueryGroupMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchQueryGroupMemberResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *BatchQueryGroupMemberResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *BatchQueryGroupMemberResponseBody) SetHasMore(v bool) *BatchQueryGroupMemberResponseBody {
	s.HasMore = &v
	return s
}

func (s *BatchQueryGroupMemberResponseBody) SetMemberUserIds(v []*string) *BatchQueryGroupMemberResponseBody {
	s.MemberUserIds = v
	return s
}

func (s *BatchQueryGroupMemberResponseBody) SetNextToken(v string) *BatchQueryGroupMemberResponseBody {
	s.NextToken = &v
	return s
}

func (s *BatchQueryGroupMemberResponseBody) SetRequestId(v string) *BatchQueryGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchQueryGroupMemberResponseBody) SetVendorRequestId(v string) *BatchQueryGroupMemberResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *BatchQueryGroupMemberResponseBody) SetVendorType(v string) *BatchQueryGroupMemberResponseBody {
	s.VendorType = &v
	return s
}

func (s *BatchQueryGroupMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
