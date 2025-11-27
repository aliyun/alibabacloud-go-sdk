// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRCInstanceResponseBody
	GetRequestId() *string
}

type DeleteRCInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRCInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRCInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRCInstanceResponseBody) SetRequestId(v string) *DeleteRCInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRCInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
