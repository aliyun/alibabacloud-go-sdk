// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTargetResponseBody
	GetRequestId() *string
	SetTargets(v []*DescribeTargetResponseBodyTargets) *DescribeTargetResponseBody
	GetTargets() []*DescribeTargetResponseBodyTargets
	SetTotalCount(v int32) *DescribeTargetResponseBody
	GetTotalCount() *int32
}

type DescribeTargetResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 79CFF74D-E967-5407-8A78-EE03B925FDAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the server.
	Targets []*DescribeTargetResponseBodyTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTargetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTargetResponseBody) GetTargets() []*DescribeTargetResponseBodyTargets {
	return s.Targets
}

func (s *DescribeTargetResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTargetResponseBody) SetRequestId(v string) *DescribeTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTargetResponseBody) SetTargets(v []*DescribeTargetResponseBodyTargets) *DescribeTargetResponseBody {
	s.Targets = v
	return s
}

func (s *DescribeTargetResponseBody) SetTotalCount(v int32) *DescribeTargetResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTargetResponseBody) Validate() error {
	if s.Targets != nil {
		for _, item := range s.Targets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTargetResponseBodyTargets struct {
	// The flag that is added to the server. This parameter can be empty.
	//
	// example:
	//
	// del
	Flag *string `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The UUID of the server or the ID of the server group.
	//
	// example:
	//
	// 5c5f0169-3527-40a2-b5ff-0bc1db8f****
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The type of the object. Valid values:
	//
	// 	- **uuid**: a server
	//
	// 	- **groupId**: a server group
	//
	// example:
	//
	// uuid
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DescribeTargetResponseBodyTargets) String() string {
	return dara.Prettify(s)
}

func (s DescribeTargetResponseBodyTargets) GoString() string {
	return s.String()
}

func (s *DescribeTargetResponseBodyTargets) GetFlag() *string {
	return s.Flag
}

func (s *DescribeTargetResponseBodyTargets) GetTarget() *string {
	return s.Target
}

func (s *DescribeTargetResponseBodyTargets) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeTargetResponseBodyTargets) SetFlag(v string) *DescribeTargetResponseBodyTargets {
	s.Flag = &v
	return s
}

func (s *DescribeTargetResponseBodyTargets) SetTarget(v string) *DescribeTargetResponseBodyTargets {
	s.Target = &v
	return s
}

func (s *DescribeTargetResponseBodyTargets) SetTargetType(v string) *DescribeTargetResponseBodyTargets {
	s.TargetType = &v
	return s
}

func (s *DescribeTargetResponseBodyTargets) Validate() error {
	return dara.Validate(s)
}
