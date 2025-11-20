// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchInnerGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupInfos(v []*SearchInnerGroupsResponseBodyGroupInfos) *SearchInnerGroupsResponseBody
	GetGroupInfos() []*SearchInnerGroupsResponseBodyGroupInfos
	SetRequestId(v string) *SearchInnerGroupsResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *SearchInnerGroupsResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *SearchInnerGroupsResponseBody
	GetVendorType() *string
}

type SearchInnerGroupsResponseBody struct {
	GroupInfos []*SearchInnerGroupsResponseBodyGroupInfos `json:"groupInfos,omitempty" xml:"groupInfos,omitempty" type:"Repeated"`
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

func (s SearchInnerGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchInnerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchInnerGroupsResponseBody) GetGroupInfos() []*SearchInnerGroupsResponseBodyGroupInfos {
	return s.GroupInfos
}

func (s *SearchInnerGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchInnerGroupsResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *SearchInnerGroupsResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *SearchInnerGroupsResponseBody) SetGroupInfos(v []*SearchInnerGroupsResponseBodyGroupInfos) *SearchInnerGroupsResponseBody {
	s.GroupInfos = v
	return s
}

func (s *SearchInnerGroupsResponseBody) SetRequestId(v string) *SearchInnerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchInnerGroupsResponseBody) SetVendorRequestId(v string) *SearchInnerGroupsResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *SearchInnerGroupsResponseBody) SetVendorType(v string) *SearchInnerGroupsResponseBody {
	s.VendorType = &v
	return s
}

func (s *SearchInnerGroupsResponseBody) Validate() error {
	if s.GroupInfos != nil {
		for _, item := range s.GroupInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchInnerGroupsResponseBodyGroupInfos struct {
	// example:
	//
	// @lAD*****
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// 10
	MemberAmount *string `json:"MemberAmount,omitempty" xml:"MemberAmount,omitempty"`
	// example:
	//
	// cid13*****==
	OpenConversationId *string `json:"OpenConversationId,omitempty" xml:"OpenConversationId,omitempty"`
	Title              *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SearchInnerGroupsResponseBodyGroupInfos) String() string {
	return dara.Prettify(s)
}

func (s SearchInnerGroupsResponseBodyGroupInfos) GoString() string {
	return s.String()
}

func (s *SearchInnerGroupsResponseBodyGroupInfos) GetIcon() *string {
	return s.Icon
}

func (s *SearchInnerGroupsResponseBodyGroupInfos) GetMemberAmount() *string {
	return s.MemberAmount
}

func (s *SearchInnerGroupsResponseBodyGroupInfos) GetOpenConversationId() *string {
	return s.OpenConversationId
}

func (s *SearchInnerGroupsResponseBodyGroupInfos) GetTitle() *string {
	return s.Title
}

func (s *SearchInnerGroupsResponseBodyGroupInfos) SetIcon(v string) *SearchInnerGroupsResponseBodyGroupInfos {
	s.Icon = &v
	return s
}

func (s *SearchInnerGroupsResponseBodyGroupInfos) SetMemberAmount(v string) *SearchInnerGroupsResponseBodyGroupInfos {
	s.MemberAmount = &v
	return s
}

func (s *SearchInnerGroupsResponseBodyGroupInfos) SetOpenConversationId(v string) *SearchInnerGroupsResponseBodyGroupInfos {
	s.OpenConversationId = &v
	return s
}

func (s *SearchInnerGroupsResponseBodyGroupInfos) SetTitle(v string) *SearchInnerGroupsResponseBodyGroupInfos {
	s.Title = &v
	return s
}

func (s *SearchInnerGroupsResponseBodyGroupInfos) Validate() error {
	return dara.Validate(s)
}
