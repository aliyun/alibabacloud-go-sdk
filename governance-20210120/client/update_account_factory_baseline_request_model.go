// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccountFactoryBaselineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineId(v string) *UpdateAccountFactoryBaselineRequest
	GetBaselineId() *string
	SetBaselineItems(v []*UpdateAccountFactoryBaselineRequestBaselineItems) *UpdateAccountFactoryBaselineRequest
	GetBaselineItems() []*UpdateAccountFactoryBaselineRequestBaselineItems
	SetBaselineName(v string) *UpdateAccountFactoryBaselineRequest
	GetBaselineName() *string
	SetDescription(v string) *UpdateAccountFactoryBaselineRequest
	GetDescription() *string
	SetRegionId(v string) *UpdateAccountFactoryBaselineRequest
	GetRegionId() *string
}

type UpdateAccountFactoryBaselineRequest struct {
	// The baseline ID.
	//
	// example:
	//
	// afb-bp1pq3emlkt27vsj****
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The baseline items.
	//
	// You can call the [ListAccountFactoryBaselineItems](~~ListAccountFactoryBaselineItems~~) operation to query a list of baseline items supported by the account factory in Cloud Governance Center.
	BaselineItems []*UpdateAccountFactoryBaselineRequestBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
	// The name of the baseline.
	BaselineName *string `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	// The description of the baseline.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateAccountFactoryBaselineRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccountFactoryBaselineRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccountFactoryBaselineRequest) GetBaselineId() *string {
	return s.BaselineId
}

func (s *UpdateAccountFactoryBaselineRequest) GetBaselineItems() []*UpdateAccountFactoryBaselineRequestBaselineItems {
	return s.BaselineItems
}

func (s *UpdateAccountFactoryBaselineRequest) GetBaselineName() *string {
	return s.BaselineName
}

func (s *UpdateAccountFactoryBaselineRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAccountFactoryBaselineRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAccountFactoryBaselineRequest) SetBaselineId(v string) *UpdateAccountFactoryBaselineRequest {
	s.BaselineId = &v
	return s
}

func (s *UpdateAccountFactoryBaselineRequest) SetBaselineItems(v []*UpdateAccountFactoryBaselineRequestBaselineItems) *UpdateAccountFactoryBaselineRequest {
	s.BaselineItems = v
	return s
}

func (s *UpdateAccountFactoryBaselineRequest) SetBaselineName(v string) *UpdateAccountFactoryBaselineRequest {
	s.BaselineName = &v
	return s
}

func (s *UpdateAccountFactoryBaselineRequest) SetDescription(v string) *UpdateAccountFactoryBaselineRequest {
	s.Description = &v
	return s
}

func (s *UpdateAccountFactoryBaselineRequest) SetRegionId(v string) *UpdateAccountFactoryBaselineRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAccountFactoryBaselineRequest) Validate() error {
	if s.BaselineItems != nil {
		for _, item := range s.BaselineItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateAccountFactoryBaselineRequestBaselineItems struct {
	// The configurations of the baseline item. The value of this parameter is a JSON string.
	//
	// example:
	//
	// {\\"EnabledServices\\":[\\"CEN_TR\\",\\"CDT\\",\\"CMS\\",\\"KMS\\"]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The name of the baseline item.
	//
	// example:
	//
	// ACS-BP_ACCOUNT_FACTORY_VPC
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the baseline item.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateAccountFactoryBaselineRequestBaselineItems) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccountFactoryBaselineRequestBaselineItems) GoString() string {
	return s.String()
}

func (s *UpdateAccountFactoryBaselineRequestBaselineItems) GetConfig() *string {
	return s.Config
}

func (s *UpdateAccountFactoryBaselineRequestBaselineItems) GetName() *string {
	return s.Name
}

func (s *UpdateAccountFactoryBaselineRequestBaselineItems) GetVersion() *string {
	return s.Version
}

func (s *UpdateAccountFactoryBaselineRequestBaselineItems) SetConfig(v string) *UpdateAccountFactoryBaselineRequestBaselineItems {
	s.Config = &v
	return s
}

func (s *UpdateAccountFactoryBaselineRequestBaselineItems) SetName(v string) *UpdateAccountFactoryBaselineRequestBaselineItems {
	s.Name = &v
	return s
}

func (s *UpdateAccountFactoryBaselineRequestBaselineItems) SetVersion(v string) *UpdateAccountFactoryBaselineRequestBaselineItems {
	s.Version = &v
	return s
}

func (s *UpdateAccountFactoryBaselineRequestBaselineItems) Validate() error {
	return dara.Validate(s)
}
