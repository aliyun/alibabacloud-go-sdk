// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberId(v string) *GetMemberRequest
	GetMemberId() *string
	SetUserId(v string) *GetMemberRequest
	GetUserId() *string
}

type GetMemberRequest struct {
	// The member UID. You must specify either UserId or MemberId. You cannot specify both.
	//
	// example:
	//
	// 145883-21513926******88039
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The user UID. For more information about how to view the user UID, see [ListWorkspaceUsers](https://help.aliyun.com/document_detail/449133.html). You must specify either UserId or MemberId. You cannot specify both.
	//
	// example:
	//
	// 21513926******88039
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMemberRequest) GoString() string {
	return s.String()
}

func (s *GetMemberRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *GetMemberRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetMemberRequest) SetMemberId(v string) *GetMemberRequest {
	s.MemberId = &v
	return s
}

func (s *GetMemberRequest) SetUserId(v string) *GetMemberRequest {
	s.UserId = &v
	return s
}

func (s *GetMemberRequest) Validate() error {
	return dara.Validate(s)
}
