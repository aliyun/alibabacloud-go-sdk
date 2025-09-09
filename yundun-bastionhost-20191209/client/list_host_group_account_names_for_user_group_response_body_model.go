// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostGroupAccountNamesForUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostAccountNames(v []*string) *ListHostGroupAccountNamesForUserGroupResponseBody
	GetHostAccountNames() []*string
	SetRequestId(v string) *ListHostGroupAccountNamesForUserGroupResponseBody
	GetRequestId() *string
}

type ListHostGroupAccountNamesForUserGroupResponseBody struct {
	// The names of host accounts returned.
	HostAccountNames []*string `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHostGroupAccountNamesForUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHostGroupAccountNamesForUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostGroupAccountNamesForUserGroupResponseBody) GetHostAccountNames() []*string {
	return s.HostAccountNames
}

func (s *ListHostGroupAccountNamesForUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHostGroupAccountNamesForUserGroupResponseBody) SetHostAccountNames(v []*string) *ListHostGroupAccountNamesForUserGroupResponseBody {
	s.HostAccountNames = v
	return s
}

func (s *ListHostGroupAccountNamesForUserGroupResponseBody) SetRequestId(v string) *ListHostGroupAccountNamesForUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
