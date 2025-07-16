// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNotInOrgList(v []*string) *AddWorkspaceMembersResponseBody
	GetNotInOrgList() []*string
	SetRequestId(v string) *AddWorkspaceMembersResponseBody
	GetRequestId() *string
}

type AddWorkspaceMembersResponseBody struct {
	NotInOrgList []*string `json:"NotInOrgList,omitempty" xml:"NotInOrgList,omitempty" type:"Repeated"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AddWorkspaceMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersResponseBody) GetNotInOrgList() []*string {
	return s.NotInOrgList
}

func (s *AddWorkspaceMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddWorkspaceMembersResponseBody) SetNotInOrgList(v []*string) *AddWorkspaceMembersResponseBody {
	s.NotInOrgList = v
	return s
}

func (s *AddWorkspaceMembersResponseBody) SetRequestId(v string) *AddWorkspaceMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddWorkspaceMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
