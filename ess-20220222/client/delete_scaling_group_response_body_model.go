// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScalingGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteScalingGroupResponseBody
	GetRequestId() *string
}

type DeleteScalingGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScalingGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteScalingGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScalingGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteScalingGroupResponseBody) SetRequestId(v string) *DeleteScalingGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScalingGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
