// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAppointmentElectZookeeperLeaderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelAppointmentElectZookeeperLeaderResponseBody
	GetRequestId() *string
}

type CancelAppointmentElectZookeeperLeaderResponseBody struct {
	// example:
	//
	// 1F488A93-83FD-540F-9B67-0333AF64E6A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelAppointmentElectZookeeperLeaderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelAppointmentElectZookeeperLeaderResponseBody) GoString() string {
	return s.String()
}

func (s *CancelAppointmentElectZookeeperLeaderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelAppointmentElectZookeeperLeaderResponseBody) SetRequestId(v string) *CancelAppointmentElectZookeeperLeaderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelAppointmentElectZookeeperLeaderResponseBody) Validate() error {
	return dara.Validate(s)
}
