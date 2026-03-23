// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRCDeploymentSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateRCDeploymentSetRequest
	GetClientToken() *string
	SetDeploymentSetName(v string) *CreateRCDeploymentSetRequest
	GetDeploymentSetName() *string
	SetDescription(v string) *CreateRCDeploymentSetRequest
	GetDescription() *string
	SetGroupCount(v int64) *CreateRCDeploymentSetRequest
	GetGroupCount() *int64
	SetOnUnableToRedeployFailedInstance(v string) *CreateRCDeploymentSetRequest
	GetOnUnableToRedeployFailedInstance() *string
	SetRegionId(v string) *CreateRCDeploymentSetRequest
	GetRegionId() *string
	SetStrategy(v string) *CreateRCDeploymentSetRequest
	GetStrategy() *string
	SetTag(v []*CreateRCDeploymentSetRequestTag) *CreateRCDeploymentSetRequest
	GetTag() []*CreateRCDeploymentSetRequestTag
}

type CreateRCDeploymentSetRequest struct {
	ClientToken                      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DeploymentSetName                *string `json:"DeploymentSetName,omitempty" xml:"DeploymentSetName,omitempty"`
	Description                      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupCount                       *int64  `json:"GroupCount,omitempty" xml:"GroupCount,omitempty"`
	OnUnableToRedeployFailedInstance *string `json:"OnUnableToRedeployFailedInstance,omitempty" xml:"OnUnableToRedeployFailedInstance,omitempty"`
	// This parameter is required.
	RegionId *string                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Strategy *string                            `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	Tag      []*CreateRCDeploymentSetRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateRCDeploymentSetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRCDeploymentSetRequest) GoString() string {
	return s.String()
}

func (s *CreateRCDeploymentSetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRCDeploymentSetRequest) GetDeploymentSetName() *string {
	return s.DeploymentSetName
}

func (s *CreateRCDeploymentSetRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRCDeploymentSetRequest) GetGroupCount() *int64 {
	return s.GroupCount
}

func (s *CreateRCDeploymentSetRequest) GetOnUnableToRedeployFailedInstance() *string {
	return s.OnUnableToRedeployFailedInstance
}

func (s *CreateRCDeploymentSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRCDeploymentSetRequest) GetStrategy() *string {
	return s.Strategy
}

func (s *CreateRCDeploymentSetRequest) GetTag() []*CreateRCDeploymentSetRequestTag {
	return s.Tag
}

func (s *CreateRCDeploymentSetRequest) SetClientToken(v string) *CreateRCDeploymentSetRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRCDeploymentSetRequest) SetDeploymentSetName(v string) *CreateRCDeploymentSetRequest {
	s.DeploymentSetName = &v
	return s
}

func (s *CreateRCDeploymentSetRequest) SetDescription(v string) *CreateRCDeploymentSetRequest {
	s.Description = &v
	return s
}

func (s *CreateRCDeploymentSetRequest) SetGroupCount(v int64) *CreateRCDeploymentSetRequest {
	s.GroupCount = &v
	return s
}

func (s *CreateRCDeploymentSetRequest) SetOnUnableToRedeployFailedInstance(v string) *CreateRCDeploymentSetRequest {
	s.OnUnableToRedeployFailedInstance = &v
	return s
}

func (s *CreateRCDeploymentSetRequest) SetRegionId(v string) *CreateRCDeploymentSetRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRCDeploymentSetRequest) SetStrategy(v string) *CreateRCDeploymentSetRequest {
	s.Strategy = &v
	return s
}

func (s *CreateRCDeploymentSetRequest) SetTag(v []*CreateRCDeploymentSetRequestTag) *CreateRCDeploymentSetRequest {
	s.Tag = v
	return s
}

func (s *CreateRCDeploymentSetRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateRCDeploymentSetRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRCDeploymentSetRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateRCDeploymentSetRequestTag) GoString() string {
	return s.String()
}

func (s *CreateRCDeploymentSetRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateRCDeploymentSetRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateRCDeploymentSetRequestTag) SetKey(v string) *CreateRCDeploymentSetRequestTag {
	s.Key = &v
	return s
}

func (s *CreateRCDeploymentSetRequestTag) SetValue(v string) *CreateRCDeploymentSetRequestTag {
	s.Value = &v
	return s
}

func (s *CreateRCDeploymentSetRequestTag) Validate() error {
	return dara.Validate(s)
}
