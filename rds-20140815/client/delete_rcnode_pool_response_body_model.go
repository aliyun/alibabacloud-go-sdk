// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCNodePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRCNodePoolResponseBody
	GetRequestId() *string
}

type DeleteRCNodePoolResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRCNodePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCNodePoolResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRCNodePoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRCNodePoolResponseBody) SetRequestId(v string) *DeleteRCNodePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRCNodePoolResponseBody) Validate() error {
	return dara.Validate(s)
}
