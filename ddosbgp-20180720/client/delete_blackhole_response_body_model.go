// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBlackholeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBlackholeResponseBody
	GetRequestId() *string
}

type DeleteBlackholeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBlackholeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBlackholeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBlackholeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBlackholeResponseBody) SetRequestId(v string) *DeleteBlackholeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBlackholeResponseBody) Validate() error {
	return dara.Validate(s)
}
