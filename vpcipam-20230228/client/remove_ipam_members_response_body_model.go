// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveIpamMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveIpamMembersResponseBody
	GetRequestId() *string
}

type RemoveIpamMembersResponseBody struct {
	// example:
	//
	// 3F814C37-B032-5477-AF5A-2925D0593CD4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveIpamMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveIpamMembersResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveIpamMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveIpamMembersResponseBody) SetRequestId(v string) *RemoveIpamMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveIpamMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
