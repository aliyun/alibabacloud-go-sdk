// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessControlListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAccessControlListResponseBody
	GetRequestId() *string
}

type DeleteAccessControlListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 988CB45E-1643-48C0-87B4-928DDF77EA49
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccessControlListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessControlListResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessControlListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAccessControlListResponseBody) SetRequestId(v string) *DeleteAccessControlListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccessControlListResponseBody) Validate() error {
	return dara.Validate(s)
}
