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
	// The ID of the request.
	//
	// example:
	//
	// EF924FE4-2EDD-4CD3-89EC-34E4708574E7
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
