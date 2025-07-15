// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCoordinationForMonitoringResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelCoordinationForMonitoringResponseBody
	GetRequestId() *string
}

type CancelCoordinationForMonitoringResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelCoordinationForMonitoringResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelCoordinationForMonitoringResponseBody) GoString() string {
	return s.String()
}

func (s *CancelCoordinationForMonitoringResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelCoordinationForMonitoringResponseBody) SetRequestId(v string) *CancelCoordinationForMonitoringResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelCoordinationForMonitoringResponseBody) Validate() error {
	return dara.Validate(s)
}
