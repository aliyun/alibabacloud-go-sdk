// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRdMemberListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRdMemberListResponseBody
	GetRequestId() *string
}

type DeleteRdMemberListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A2D6D5FB-FA07-41A8-B093-A2B7B26E72F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRdMemberListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRdMemberListResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRdMemberListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRdMemberListResponseBody) SetRequestId(v string) *DeleteRdMemberListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRdMemberListResponseBody) Validate() error {
	return dara.Validate(s)
}
