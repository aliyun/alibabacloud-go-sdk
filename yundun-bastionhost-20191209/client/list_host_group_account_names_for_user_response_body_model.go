// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostGroupAccountNamesForUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostAccountNames(v []*string) *ListHostGroupAccountNamesForUserResponseBody
	GetHostAccountNames() []*string
	SetRequestId(v string) *ListHostGroupAccountNamesForUserResponseBody
	GetRequestId() *string
}

type ListHostGroupAccountNamesForUserResponseBody struct {
	// An array that consists of the names of host accounts.
	HostAccountNames []*string `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHostGroupAccountNamesForUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHostGroupAccountNamesForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostGroupAccountNamesForUserResponseBody) GetHostAccountNames() []*string {
	return s.HostAccountNames
}

func (s *ListHostGroupAccountNamesForUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHostGroupAccountNamesForUserResponseBody) SetHostAccountNames(v []*string) *ListHostGroupAccountNamesForUserResponseBody {
	s.HostAccountNames = v
	return s
}

func (s *ListHostGroupAccountNamesForUserResponseBody) SetRequestId(v string) *ListHostGroupAccountNamesForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserResponseBody) Validate() error {
	return dara.Validate(s)
}
