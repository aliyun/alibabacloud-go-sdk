// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountFactoryBaselineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineItems(v []*CreateAccountFactoryBaselineRequestBaselineItems) *CreateAccountFactoryBaselineRequest
	GetBaselineItems() []*CreateAccountFactoryBaselineRequestBaselineItems
	SetBaselineName(v string) *CreateAccountFactoryBaselineRequest
	GetBaselineName() *string
	SetDescription(v string) *CreateAccountFactoryBaselineRequest
	GetDescription() *string
	SetRegionId(v string) *CreateAccountFactoryBaselineRequest
	GetRegionId() *string
}

type CreateAccountFactoryBaselineRequest struct {
	// An array that contains the baseline items.
	//
	// You can call the [ListAccountFactoryBaselineItems](~~ListAccountFactoryBaselineItems~~) operation to query a list of baseline items supported by the account factory in Cloud Governance Center.
	BaselineItems []*CreateAccountFactoryBaselineRequestBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
	// The name of the baseline.
	//
	// example:
	//
	// Default
	BaselineName *string `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	// The description of the baseline.
	//
	// example:
	//
	// Default Baseline.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateAccountFactoryBaselineRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountFactoryBaselineRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountFactoryBaselineRequest) GetBaselineItems() []*CreateAccountFactoryBaselineRequestBaselineItems {
	return s.BaselineItems
}

func (s *CreateAccountFactoryBaselineRequest) GetBaselineName() *string {
	return s.BaselineName
}

func (s *CreateAccountFactoryBaselineRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAccountFactoryBaselineRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAccountFactoryBaselineRequest) SetBaselineItems(v []*CreateAccountFactoryBaselineRequestBaselineItems) *CreateAccountFactoryBaselineRequest {
	s.BaselineItems = v
	return s
}

func (s *CreateAccountFactoryBaselineRequest) SetBaselineName(v string) *CreateAccountFactoryBaselineRequest {
	s.BaselineName = &v
	return s
}

func (s *CreateAccountFactoryBaselineRequest) SetDescription(v string) *CreateAccountFactoryBaselineRequest {
	s.Description = &v
	return s
}

func (s *CreateAccountFactoryBaselineRequest) SetRegionId(v string) *CreateAccountFactoryBaselineRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAccountFactoryBaselineRequest) Validate() error {
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

type CreateAccountFactoryBaselineRequestBaselineItems struct {
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

func (s CreateAccountFactoryBaselineRequestBaselineItems) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountFactoryBaselineRequestBaselineItems) GoString() string {
	return s.String()
}

func (s *CreateAccountFactoryBaselineRequestBaselineItems) GetConfig() *string {
	return s.Config
}

func (s *CreateAccountFactoryBaselineRequestBaselineItems) GetName() *string {
	return s.Name
}

func (s *CreateAccountFactoryBaselineRequestBaselineItems) GetVersion() *string {
	return s.Version
}

func (s *CreateAccountFactoryBaselineRequestBaselineItems) SetConfig(v string) *CreateAccountFactoryBaselineRequestBaselineItems {
	s.Config = &v
	return s
}

func (s *CreateAccountFactoryBaselineRequestBaselineItems) SetName(v string) *CreateAccountFactoryBaselineRequestBaselineItems {
	s.Name = &v
	return s
}

func (s *CreateAccountFactoryBaselineRequestBaselineItems) SetVersion(v string) *CreateAccountFactoryBaselineRequestBaselineItems {
	s.Version = &v
	return s
}

func (s *CreateAccountFactoryBaselineRequestBaselineItems) Validate() error {
	return dara.Validate(s)
}
