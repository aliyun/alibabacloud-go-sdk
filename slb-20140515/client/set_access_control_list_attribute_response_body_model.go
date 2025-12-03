// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAccessControlListAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetAccessControlListAttributeResponseBody
	GetRequestId() *string
}

type SetAccessControlListAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 988CB45E-1643-48C0-87B4-928DDF77EA49
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAccessControlListAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAccessControlListAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *SetAccessControlListAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAccessControlListAttributeResponseBody) SetRequestId(v string) *SetAccessControlListAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAccessControlListAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
