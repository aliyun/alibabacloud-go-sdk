// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachServerGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachServerGroupsResponseBody
	GetRequestId() *string
	SetScalingActivityId(v string) *AttachServerGroupsResponseBody
	GetScalingActivityId() *string
}

type AttachServerGroupsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scaling activity. During the scaling activity, the server group is attached to the scaling group and the existing ECS instances or elastic container instances in the scaling group are added to the server group.
	//
	// >  This parameter is returned only if you set ForceAttach to true.
	//
	// example:
	//
	// asa-bp1c9djwrgxjyk31****
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s AttachServerGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachServerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *AttachServerGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachServerGroupsResponseBody) GetScalingActivityId() *string {
	return s.ScalingActivityId
}

func (s *AttachServerGroupsResponseBody) SetRequestId(v string) *AttachServerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachServerGroupsResponseBody) SetScalingActivityId(v string) *AttachServerGroupsResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *AttachServerGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
