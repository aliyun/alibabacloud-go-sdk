// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteInstanceMembersResponseBody
	GetRequestId() *string
}

type DeleteInstanceMembersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 03E8AA70-0CC9-42EA-97AA-EA68377930B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstanceMembersResponseBody) SetRequestId(v string) *DeleteInstanceMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
