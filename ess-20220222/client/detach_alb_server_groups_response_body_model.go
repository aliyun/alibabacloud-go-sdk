// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachAlbServerGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachAlbServerGroupsResponseBody
	GetRequestId() *string
	SetScalingActivityId(v string) *DetachAlbServerGroupsResponseBody
	GetScalingActivityId() *string
}

type DetachAlbServerGroupsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scaling activity during which the ALB server group is detached from the scaling group and the existing ECS instances are removed from the ALB server group. This parameter has a return value only if you set `ForceDetach` to `true`.
	//
	// example:
	//
	// asa-2ze6wxj8vsohn6j9****
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s DetachAlbServerGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachAlbServerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DetachAlbServerGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachAlbServerGroupsResponseBody) GetScalingActivityId() *string {
	return s.ScalingActivityId
}

func (s *DetachAlbServerGroupsResponseBody) SetRequestId(v string) *DetachAlbServerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachAlbServerGroupsResponseBody) SetScalingActivityId(v string) *DetachAlbServerGroupsResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *DetachAlbServerGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
