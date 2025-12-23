// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIpamMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddIpamMembersResponseBody
	GetRequestId() *string
}

type AddIpamMembersResponseBody struct {
	// example:
	//
	// BB2C39DE-CEB8-595A-981A-F2EFCBE7324E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddIpamMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddIpamMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddIpamMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddIpamMembersResponseBody) SetRequestId(v string) *AddIpamMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddIpamMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
