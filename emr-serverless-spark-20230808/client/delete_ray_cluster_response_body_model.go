// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRayClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRayClusterResponseBody
	GetRequestId() *string
}

type DeleteRayClusterResponseBody struct {
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteRayClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRayClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRayClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRayClusterResponseBody) SetRequestId(v string) *DeleteRayClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRayClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
