// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateAccessApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePrivateAccessApplicationResponseBody
	GetRequestId() *string
}

type UpdatePrivateAccessApplicationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// FD724DBC-CD76-5235-BF76-59C51B73296D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePrivateAccessApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateAccessApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePrivateAccessApplicationResponseBody) SetRequestId(v string) *UpdatePrivateAccessApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePrivateAccessApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
