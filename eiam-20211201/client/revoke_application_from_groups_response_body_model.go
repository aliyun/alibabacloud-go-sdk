// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeApplicationFromGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeApplicationFromGroupsResponseBody
	GetRequestId() *string
}

type RevokeApplicationFromGroupsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeApplicationFromGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeApplicationFromGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeApplicationFromGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeApplicationFromGroupsResponseBody) SetRequestId(v string) *RevokeApplicationFromGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeApplicationFromGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
