// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeploySDGRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentType(v string) *DeploySDGRequest
	GetDeploymentType() *string
	SetInstanceIds(v []*string) *DeploySDGRequest
	GetInstanceIds() []*string
	SetSDGId(v string) *DeploySDGRequest
	GetSDGId() *string
}

type DeploySDGRequest struct {
	// The SDG deployment type. Valid values:
	//
	// - common: (default) read/write deployment. Data updates are written to the cloud disk.
	//
	// - overlay: read/write splitting deployment. The SDG content is read-only, and data updates are written to the local storage of the instance.
	//
	// example:
	//
	// common
	DeploymentType *string `json:"DeploymentType,omitempty" xml:"DeploymentType,omitempty"`
	// The list of device resource IDs to which the SDG is deployed. Batch deployment is supported, with a maximum of 100 instances per request.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The SDG ID. A cloud disk is created based on this SDG ID and mounted to the specified instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sdg-xxxxx
	SDGId *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
}

func (s DeploySDGRequest) String() string {
	return dara.Prettify(s)
}

func (s DeploySDGRequest) GoString() string {
	return s.String()
}

func (s *DeploySDGRequest) GetDeploymentType() *string {
	return s.DeploymentType
}

func (s *DeploySDGRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DeploySDGRequest) GetSDGId() *string {
	return s.SDGId
}

func (s *DeploySDGRequest) SetDeploymentType(v string) *DeploySDGRequest {
	s.DeploymentType = &v
	return s
}

func (s *DeploySDGRequest) SetInstanceIds(v []*string) *DeploySDGRequest {
	s.InstanceIds = v
	return s
}

func (s *DeploySDGRequest) SetSDGId(v string) *DeploySDGRequest {
	s.SDGId = &v
	return s
}

func (s *DeploySDGRequest) Validate() error {
	return dara.Validate(s)
}
