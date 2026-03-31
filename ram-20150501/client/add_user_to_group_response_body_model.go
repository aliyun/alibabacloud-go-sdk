// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddUserToGroupResponseBody
	GetRequestId() *string
}

type AddUserToGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1B968853-B423-63A6-FE1F-45E81BC2AD61
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserToGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUserToGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserToGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUserToGroupResponseBody) SetRequestId(v string) *AddUserToGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserToGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
