// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAccessControlListEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddAccessControlListEntryResponseBody
	GetRequestId() *string
}

type AddAccessControlListEntryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 988CB45E-1643-48C0-87B4-928DDF77EA4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAccessControlListEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAccessControlListEntryResponseBody) GoString() string {
	return s.String()
}

func (s *AddAccessControlListEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAccessControlListEntryResponseBody) SetRequestId(v string) *AddAccessControlListEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAccessControlListEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
