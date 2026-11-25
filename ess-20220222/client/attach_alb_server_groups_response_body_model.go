// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachAlbServerGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachAlbServerGroupsResponseBody
	GetRequestId() *string
	SetScalingActivityId(v string) *AttachAlbServerGroupsResponseBody
	GetScalingActivityId() *string
}

type AttachAlbServerGroupsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scaling activity that is triggered when the ALB server group is added and the ECS instances or elastic container instances in the scaling group are added to the ALB server group.
	//
	// > This parameter is returned only when `ForceAttach` is set to `true`.
	//
	// example:
	//
	// asa-2ze6wxj8vsohn6j9****
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s AttachAlbServerGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachAlbServerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *AttachAlbServerGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachAlbServerGroupsResponseBody) GetScalingActivityId() *string {
	return s.ScalingActivityId
}

func (s *AttachAlbServerGroupsResponseBody) SetRequestId(v string) *AttachAlbServerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachAlbServerGroupsResponseBody) SetScalingActivityId(v string) *AttachAlbServerGroupsResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *AttachAlbServerGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
