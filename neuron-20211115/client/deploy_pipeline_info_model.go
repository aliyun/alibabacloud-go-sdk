// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployPipelineInfo interface {
	dara.Model
	String() string
	GoString() string
	SetDeployInstanceInfos(v []*DeployInstanceInfo) *DeployPipelineInfo
	GetDeployInstanceInfos() []*DeployInstanceInfo
	SetId(v string) *DeployPipelineInfo
	GetId() *string
	SetName(v string) *DeployPipelineInfo
	GetName() *string
}

type DeployPipelineInfo struct {
	DeployInstanceInfos []*DeployInstanceInfo `json:"deployInstanceInfos,omitempty" xml:"deployInstanceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 528f9a66-cbe3-4cd8-91c0-bc4260a13d87
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// Batch 1 Change
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DeployPipelineInfo) String() string {
	return dara.Prettify(s)
}

func (s DeployPipelineInfo) GoString() string {
	return s.String()
}

func (s *DeployPipelineInfo) GetDeployInstanceInfos() []*DeployInstanceInfo {
	return s.DeployInstanceInfos
}

func (s *DeployPipelineInfo) GetId() *string {
	return s.Id
}

func (s *DeployPipelineInfo) GetName() *string {
	return s.Name
}

func (s *DeployPipelineInfo) SetDeployInstanceInfos(v []*DeployInstanceInfo) *DeployPipelineInfo {
	s.DeployInstanceInfos = v
	return s
}

func (s *DeployPipelineInfo) SetId(v string) *DeployPipelineInfo {
	s.Id = &v
	return s
}

func (s *DeployPipelineInfo) SetName(v string) *DeployPipelineInfo {
	s.Name = &v
	return s
}

func (s *DeployPipelineInfo) Validate() error {
	if s.DeployInstanceInfos != nil {
		for _, item := range s.DeployInstanceInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
