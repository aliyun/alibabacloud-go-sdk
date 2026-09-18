// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertRobotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAlertRobotsResponseBody
	GetRequestId() *string
}

type DeleteAlertRobotsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteAlertRobotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertRobotsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertRobotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAlertRobotsResponseBody) SetRequestId(v string) *DeleteAlertRobotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlertRobotsResponseBody) Validate() error {
	return dara.Validate(s)
}
