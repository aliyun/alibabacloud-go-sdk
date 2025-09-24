// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberIds(v string) *DeleteMembersRequest
	GetMemberIds() *string
}

type DeleteMembersRequest struct {
	// The list of member IDs. Separate multiple member IDs with commas (,). You can call [ListMembers](https://help.aliyun.com/document_detail/449135.html) to obtain the member ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 145883-21513926******88039,145883-2769726******87513
	MemberIds *string `json:"MemberIds,omitempty" xml:"MemberIds,omitempty"`
}

func (s DeleteMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMembersRequest) GoString() string {
	return s.String()
}

func (s *DeleteMembersRequest) GetMemberIds() *string {
	return s.MemberIds
}

func (s *DeleteMembersRequest) SetMemberIds(v string) *DeleteMembersRequest {
	s.MemberIds = &v
	return s
}

func (s *DeleteMembersRequest) Validate() error {
	return dara.Validate(s)
}
