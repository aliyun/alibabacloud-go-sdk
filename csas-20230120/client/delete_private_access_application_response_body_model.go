// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateAccessApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePrivateAccessApplicationResponseBody
	GetRequestId() *string
}

type DeletePrivateAccessApplicationResponseBody struct {
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePrivateAccessApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateAccessApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrivateAccessApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePrivateAccessApplicationResponseBody) SetRequestId(v string) *DeletePrivateAccessApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrivateAccessApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
