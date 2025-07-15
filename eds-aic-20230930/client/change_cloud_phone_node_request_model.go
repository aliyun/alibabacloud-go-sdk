// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeCloudPhoneNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceType(v string) *ChangeCloudPhoneNodeRequest
	GetInstanceType() *string
	SetNodeId(v string) *ChangeCloudPhoneNodeRequest
	GetNodeId() *string
	SetPhoneCount(v int32) *ChangeCloudPhoneNodeRequest
	GetPhoneCount() *int32
}

type ChangeCloudPhoneNodeRequest struct {
	// example:
	//
	// ac.max
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// cpn-0ugbptfu473fy****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// 20
	PhoneCount *int32 `json:"PhoneCount,omitempty" xml:"PhoneCount,omitempty"`
}

func (s ChangeCloudPhoneNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeCloudPhoneNodeRequest) GoString() string {
	return s.String()
}

func (s *ChangeCloudPhoneNodeRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ChangeCloudPhoneNodeRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ChangeCloudPhoneNodeRequest) GetPhoneCount() *int32 {
	return s.PhoneCount
}

func (s *ChangeCloudPhoneNodeRequest) SetInstanceType(v string) *ChangeCloudPhoneNodeRequest {
	s.InstanceType = &v
	return s
}

func (s *ChangeCloudPhoneNodeRequest) SetNodeId(v string) *ChangeCloudPhoneNodeRequest {
	s.NodeId = &v
	return s
}

func (s *ChangeCloudPhoneNodeRequest) SetPhoneCount(v int32) *ChangeCloudPhoneNodeRequest {
	s.PhoneCount = &v
	return s
}

func (s *ChangeCloudPhoneNodeRequest) Validate() error {
	return dara.Validate(s)
}
