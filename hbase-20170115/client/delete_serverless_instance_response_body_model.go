// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServerlessInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteServerlessInstanceResponseBody
	GetRequestId() *string
}

type DeleteServerlessInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServerlessInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServerlessInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServerlessInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServerlessInstanceResponseBody) SetRequestId(v string) *DeleteServerlessInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServerlessInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
