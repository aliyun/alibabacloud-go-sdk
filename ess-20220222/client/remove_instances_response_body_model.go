// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveInstancesResponseBody
	GetRequestId() *string
	SetScalingActivityId(v string) *RemoveInstancesResponseBody
	GetScalingActivityId() *string
}

type RemoveInstancesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scaling activity.
	//
	// example:
	//
	// asa-bp175o6f6ego3r2j****
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s RemoveInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveInstancesResponseBody) GetScalingActivityId() *string {
	return s.ScalingActivityId
}

func (s *RemoveInstancesResponseBody) SetRequestId(v string) *RemoveInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveInstancesResponseBody) SetScalingActivityId(v string) *RemoveInstancesResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *RemoveInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
