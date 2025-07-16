// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNewestInnerGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupInfos(v []*GetNewestInnerGroupsResponseBodyGroupInfos) *GetNewestInnerGroupsResponseBody
	GetGroupInfos() []*GetNewestInnerGroupsResponseBodyGroupInfos
	SetRequestId(v string) *GetNewestInnerGroupsResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *GetNewestInnerGroupsResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetNewestInnerGroupsResponseBody
	GetVendorType() *string
}

type GetNewestInnerGroupsResponseBody struct {
	GroupInfos []*GetNewestInnerGroupsResponseBodyGroupInfos `json:"groupInfos,omitempty" xml:"groupInfos,omitempty" type:"Repeated"`
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

func (s GetNewestInnerGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNewestInnerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *GetNewestInnerGroupsResponseBody) GetGroupInfos() []*GetNewestInnerGroupsResponseBodyGroupInfos {
	return s.GroupInfos
}

func (s *GetNewestInnerGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNewestInnerGroupsResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetNewestInnerGroupsResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetNewestInnerGroupsResponseBody) SetGroupInfos(v []*GetNewestInnerGroupsResponseBodyGroupInfos) *GetNewestInnerGroupsResponseBody {
	s.GroupInfos = v
	return s
}

func (s *GetNewestInnerGroupsResponseBody) SetRequestId(v string) *GetNewestInnerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNewestInnerGroupsResponseBody) SetVendorRequestId(v string) *GetNewestInnerGroupsResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetNewestInnerGroupsResponseBody) SetVendorType(v string) *GetNewestInnerGroupsResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetNewestInnerGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetNewestInnerGroupsResponseBodyGroupInfos struct {
	// example:
	//
	// @lADOADma*****QKA
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// 10
	MemberAmount *string `json:"MemberAmount,omitempty" xml:"MemberAmount,omitempty"`
	// example:
	//
	// cid1e******==
	OpenConversationId *string `json:"OpenConversationId,omitempty" xml:"OpenConversationId,omitempty"`
	Title              *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetNewestInnerGroupsResponseBodyGroupInfos) String() string {
	return dara.Prettify(s)
}

func (s GetNewestInnerGroupsResponseBodyGroupInfos) GoString() string {
	return s.String()
}

func (s *GetNewestInnerGroupsResponseBodyGroupInfos) GetIcon() *string {
	return s.Icon
}

func (s *GetNewestInnerGroupsResponseBodyGroupInfos) GetMemberAmount() *string {
	return s.MemberAmount
}

func (s *GetNewestInnerGroupsResponseBodyGroupInfos) GetOpenConversationId() *string {
	return s.OpenConversationId
}

func (s *GetNewestInnerGroupsResponseBodyGroupInfos) GetTitle() *string {
	return s.Title
}

func (s *GetNewestInnerGroupsResponseBodyGroupInfos) SetIcon(v string) *GetNewestInnerGroupsResponseBodyGroupInfos {
	s.Icon = &v
	return s
}

func (s *GetNewestInnerGroupsResponseBodyGroupInfos) SetMemberAmount(v string) *GetNewestInnerGroupsResponseBodyGroupInfos {
	s.MemberAmount = &v
	return s
}

func (s *GetNewestInnerGroupsResponseBodyGroupInfos) SetOpenConversationId(v string) *GetNewestInnerGroupsResponseBodyGroupInfos {
	s.OpenConversationId = &v
	return s
}

func (s *GetNewestInnerGroupsResponseBodyGroupInfos) SetTitle(v string) *GetNewestInnerGroupsResponseBodyGroupInfos {
	s.Title = &v
	return s
}

func (s *GetNewestInnerGroupsResponseBodyGroupInfos) Validate() error {
	return dara.Validate(s)
}
