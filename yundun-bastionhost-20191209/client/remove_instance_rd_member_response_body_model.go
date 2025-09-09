// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveInstanceRdMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveInstanceRdMemberResponseBody
	GetRequestId() *string
}

type RemoveInstanceRdMemberResponseBody struct {
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveInstanceRdMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstanceRdMemberResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveInstanceRdMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveInstanceRdMemberResponseBody) SetRequestId(v string) *RemoveInstanceRdMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveInstanceRdMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
