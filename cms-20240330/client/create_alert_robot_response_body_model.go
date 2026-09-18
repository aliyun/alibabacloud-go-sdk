// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertRobotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertRobotId(v string) *CreateAlertRobotResponseBody
	GetAlertRobotId() *string
	SetRequestId(v string) *CreateAlertRobotResponseBody
	GetRequestId() *string
}

type CreateAlertRobotResponseBody struct {
	// The robot ID.
	//
	// example:
	//
	// testId
	AlertRobotId *string `json:"alertRobotId,omitempty" xml:"alertRobotId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateAlertRobotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertRobotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlertRobotResponseBody) GetAlertRobotId() *string {
	return s.AlertRobotId
}

func (s *CreateAlertRobotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAlertRobotResponseBody) SetAlertRobotId(v string) *CreateAlertRobotResponseBody {
	s.AlertRobotId = &v
	return s
}

func (s *CreateAlertRobotResponseBody) SetRequestId(v string) *CreateAlertRobotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAlertRobotResponseBody) Validate() error {
	return dara.Validate(s)
}
