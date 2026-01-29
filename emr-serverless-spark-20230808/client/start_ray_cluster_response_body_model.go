// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRayClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartRayClusterResponseBody
	GetRequestId() *string
}

type StartRayClusterResponseBody struct {
	// example:
	//
	// 8CE06D75-E6A2-505D-9B4B-31DEE3D98A04
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StartRayClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartRayClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StartRayClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartRayClusterResponseBody) SetRequestId(v string) *StartRayClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartRayClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
