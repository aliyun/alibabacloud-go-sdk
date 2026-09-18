// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertRobotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAlertRobotResponseBody
	GetRequestId() *string
}

type UpdateAlertRobotResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateAlertRobotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRobotResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlertRobotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAlertRobotResponseBody) SetRequestId(v string) *UpdateAlertRobotResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAlertRobotResponseBody) Validate() error {
	return dara.Validate(s)
}
