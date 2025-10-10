// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAclAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAclAttributeResponseBody
	GetRequestId() *string
}

type UpdateAclAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAclAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAclAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAclAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAclAttributeResponseBody) SetRequestId(v string) *UpdateAclAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAclAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
