// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserConfigResponseBody
	GetRequestId() *string
}

type DeleteUserConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// dsjk****dfjksdf
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserConfigResponseBody) SetRequestId(v string) *DeleteUserConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
