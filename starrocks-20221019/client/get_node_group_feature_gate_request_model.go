// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeGroupFeatureGateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetNodeGroupFeatureGateRequest
	GetInstanceId() *string
	SetNodeGroupId(v string) *GetNodeGroupFeatureGateRequest
	GetNodeGroupId() *string
}

type GetNodeGroupFeatureGateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ng-d332aa8bca48****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s GetNodeGroupFeatureGateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeGroupFeatureGateRequest) GoString() string {
	return s.String()
}

func (s *GetNodeGroupFeatureGateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetNodeGroupFeatureGateRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *GetNodeGroupFeatureGateRequest) SetInstanceId(v string) *GetNodeGroupFeatureGateRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNodeGroupFeatureGateRequest) SetNodeGroupId(v string) *GetNodeGroupFeatureGateRequest {
	s.NodeGroupId = &v
	return s
}

func (s *GetNodeGroupFeatureGateRequest) Validate() error {
	return dara.Validate(s)
}
