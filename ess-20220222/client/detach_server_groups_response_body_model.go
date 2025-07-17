// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachServerGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachServerGroupsResponseBody
	GetRequestId() *string
	SetScalingActivityId(v string) *DetachServerGroupsResponseBody
	GetScalingActivityId() *string
}

type DetachServerGroupsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6EF9BFEE-FE07-4627-B8FB-14326FB9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scaling activity. During the scaling activity, the server group is detached from the scaling group and the existing servers, which are the ECS instances or elastic container instances in the scaling group, are removed from the server group.
	//
	// >  This parameter is returned only if you set `ForceDetach` to `true`.
	//
	// example:
	//
	// asa-bp1gbswjhjrw8tko****
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s DetachServerGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachServerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DetachServerGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachServerGroupsResponseBody) GetScalingActivityId() *string {
	return s.ScalingActivityId
}

func (s *DetachServerGroupsResponseBody) SetRequestId(v string) *DetachServerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachServerGroupsResponseBody) SetScalingActivityId(v string) *DetachServerGroupsResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *DetachServerGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
