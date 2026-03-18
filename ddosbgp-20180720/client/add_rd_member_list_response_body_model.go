// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRdMemberListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddRdMemberListResponseBody
	GetRequestId() *string
}

type AddRdMemberListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddRdMemberListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddRdMemberListResponseBody) GoString() string {
	return s.String()
}

func (s *AddRdMemberListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddRdMemberListResponseBody) SetRequestId(v string) *AddRdMemberListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRdMemberListResponseBody) Validate() error {
	return dara.Validate(s)
}
