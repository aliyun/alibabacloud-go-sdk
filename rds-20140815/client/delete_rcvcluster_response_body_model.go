// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCVClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRCVClusterResponseBody
	GetRequestId() *string
}

type DeleteRCVClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRCVClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCVClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRCVClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRCVClusterResponseBody) SetRequestId(v string) *DeleteRCVClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRCVClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
