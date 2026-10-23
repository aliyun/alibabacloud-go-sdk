// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootCloudPhoneNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *RebootCloudPhoneNodesRequest
	GetForce() *bool
	SetNodeIds(v []*string) *RebootCloudPhoneNodesRequest
	GetNodeIds() []*string
}

type RebootCloudPhoneNodesRequest struct {
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// example:
	//
	// cpn-05e4d1cec2urdh1j****
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
}

func (s RebootCloudPhoneNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootCloudPhoneNodesRequest) GoString() string {
	return s.String()
}

func (s *RebootCloudPhoneNodesRequest) GetForce() *bool {
	return s.Force
}

func (s *RebootCloudPhoneNodesRequest) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *RebootCloudPhoneNodesRequest) SetForce(v bool) *RebootCloudPhoneNodesRequest {
	s.Force = &v
	return s
}

func (s *RebootCloudPhoneNodesRequest) SetNodeIds(v []*string) *RebootCloudPhoneNodesRequest {
	s.NodeIds = v
	return s
}

func (s *RebootCloudPhoneNodesRequest) Validate() error {
	return dara.Validate(s)
}
