// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInstanceRdMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddInstanceRdMemberResponseBody
	GetRequestId() *string
}

type AddInstanceRdMemberResponseBody struct {
	// example:
	//
	// 4F6C075F-FC86-476E-943B-097BD4E12948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddInstanceRdMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddInstanceRdMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddInstanceRdMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddInstanceRdMemberResponseBody) SetRequestId(v string) *AddInstanceRdMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddInstanceRdMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
