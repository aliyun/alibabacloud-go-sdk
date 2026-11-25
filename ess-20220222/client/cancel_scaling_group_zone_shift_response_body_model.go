// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelScalingGroupZoneShiftResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelScalingGroupZoneShiftResponseBody
	GetRequestId() *string
}

type CancelScalingGroupZoneShiftResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelScalingGroupZoneShiftResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelScalingGroupZoneShiftResponseBody) GoString() string {
	return s.String()
}

func (s *CancelScalingGroupZoneShiftResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelScalingGroupZoneShiftResponseBody) SetRequestId(v string) *CancelScalingGroupZoneShiftResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelScalingGroupZoneShiftResponseBody) Validate() error {
	return dara.Validate(s)
}
