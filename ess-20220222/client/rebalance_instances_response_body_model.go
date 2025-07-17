// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebalanceInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RebalanceInstancesResponseBody
	GetRequestId() *string
	SetScalingActivityId(v string) *RebalanceInstancesResponseBody
	GetScalingActivityId() *string
}

type RebalanceInstancesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scaling activity.
	//
	// example:
	//
	// asa-kjgffgdfadah****
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s RebalanceInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebalanceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RebalanceInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebalanceInstancesResponseBody) GetScalingActivityId() *string {
	return s.ScalingActivityId
}

func (s *RebalanceInstancesResponseBody) SetRequestId(v string) *RebalanceInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebalanceInstancesResponseBody) SetScalingActivityId(v string) *RebalanceInstancesResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *RebalanceInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
