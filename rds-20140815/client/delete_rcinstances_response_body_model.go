// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRCInstancesResponseBody
	GetRequestId() *string
}

type DeleteRCInstancesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRCInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRCInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRCInstancesResponseBody) SetRequestId(v string) *DeleteRCInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRCInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
