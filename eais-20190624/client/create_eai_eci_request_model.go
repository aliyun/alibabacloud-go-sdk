// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEaiEciRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateEaiEciRequest
	GetClientToken() *string
	SetEaisName(v string) *CreateEaiEciRequest
	GetEaisName() *string
	SetEaisType(v string) *CreateEaiEciRequest
	GetEaisType() *string
	SetEci(v *CreateEaiEciRequestEci) *CreateEaiEciRequest
	GetEci() *CreateEaiEciRequestEci
	SetRegionId(v string) *CreateEaiEciRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateEaiEciRequest
	GetResourceGroupId() *string
	SetSecurityGroupId(v string) *CreateEaiEciRequest
	GetSecurityGroupId() *string
	SetTag(v []*CreateEaiEciRequestTag) *CreateEaiEciRequest
	GetTag() []*CreateEaiEciRequestTag
	SetVSwitchId(v string) *CreateEaiEciRequest
	GetVSwitchId() *string
}

type CreateEaiEciRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// eais-test01
	EaisName *string `json:"EaisName,omitempty" xml:"EaisName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eais.ei-a6.2xlarge
	EaisType *string                 `json:"EaisType,omitempty" xml:"EaisType,omitempty"`
	Eci      *CreateEaiEciRequestEci `json:"Eci,omitempty" xml:"Eci,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmvpuy4a5****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sg-uf66jeqopgqa9hdn****
	SecurityGroupId *string                   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Tag             []*CreateEaiEciRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-uf6h3rbwbm90urjwa****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateEaiEciRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiEciRequest) GoString() string {
	return s.String()
}

func (s *CreateEaiEciRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEaiEciRequest) GetEaisName() *string {
	return s.EaisName
}

func (s *CreateEaiEciRequest) GetEaisType() *string {
	return s.EaisType
}

func (s *CreateEaiEciRequest) GetEci() *CreateEaiEciRequestEci {
	return s.Eci
}

func (s *CreateEaiEciRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEaiEciRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateEaiEciRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateEaiEciRequest) GetTag() []*CreateEaiEciRequestTag {
	return s.Tag
}

func (s *CreateEaiEciRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateEaiEciRequest) SetClientToken(v string) *CreateEaiEciRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEaiEciRequest) SetEaisName(v string) *CreateEaiEciRequest {
	s.EaisName = &v
	return s
}

func (s *CreateEaiEciRequest) SetEaisType(v string) *CreateEaiEciRequest {
	s.EaisType = &v
	return s
}

func (s *CreateEaiEciRequest) SetEci(v *CreateEaiEciRequestEci) *CreateEaiEciRequest {
	s.Eci = v
	return s
}

func (s *CreateEaiEciRequest) SetRegionId(v string) *CreateEaiEciRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEaiEciRequest) SetResourceGroupId(v string) *CreateEaiEciRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEaiEciRequest) SetSecurityGroupId(v string) *CreateEaiEciRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEaiEciRequest) SetTag(v []*CreateEaiEciRequestTag) *CreateEaiEciRequest {
	s.Tag = v
	return s
}

func (s *CreateEaiEciRequest) SetVSwitchId(v string) *CreateEaiEciRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateEaiEciRequest) Validate() error {
	if s.Eci != nil {
		if err := s.Eci.Validate(); err != nil {
			return err
		}
	}
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

type CreateEaiEciRequestEci struct {
	Container *CreateEaiEciRequestEciContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	// example:
	//
	// eip-uf66jeqopgqa9hdn****
	EipId *string `json:"EipId,omitempty" xml:"EipId,omitempty"`
	// example:
	//
	// test-nginx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ecs.c5.xlarge
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 00c7****-rivj.cn-hangzhou.extreme.nas.aliyuncs.com:/share
	Volume *string `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s CreateEaiEciRequestEci) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiEciRequestEci) GoString() string {
	return s.String()
}

func (s *CreateEaiEciRequestEci) GetContainer() *CreateEaiEciRequestEciContainer {
	return s.Container
}

func (s *CreateEaiEciRequestEci) GetEipId() *string {
	return s.EipId
}

func (s *CreateEaiEciRequestEci) GetName() *string {
	return s.Name
}

func (s *CreateEaiEciRequestEci) GetType() *string {
	return s.Type
}

func (s *CreateEaiEciRequestEci) GetVolume() *string {
	return s.Volume
}

func (s *CreateEaiEciRequestEci) SetContainer(v *CreateEaiEciRequestEciContainer) *CreateEaiEciRequestEci {
	s.Container = v
	return s
}

func (s *CreateEaiEciRequestEci) SetEipId(v string) *CreateEaiEciRequestEci {
	s.EipId = &v
	return s
}

func (s *CreateEaiEciRequestEci) SetName(v string) *CreateEaiEciRequestEci {
	s.Name = &v
	return s
}

func (s *CreateEaiEciRequestEci) SetType(v string) *CreateEaiEciRequestEci {
	s.Type = &v
	return s
}

func (s *CreateEaiEciRequestEci) SetVolume(v string) *CreateEaiEciRequestEci {
	s.Volume = &v
	return s
}

func (s *CreateEaiEciRequestEci) Validate() error {
	if s.Container != nil {
		if err := s.Container.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEaiEciRequestEciContainer struct {
	// example:
	//
	// 100
	Arg *string `json:"Arg,omitempty" xml:"Arg,omitempty"`
	// example:
	//
	// sleep
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// example:
	//
	// nginx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// example:
	//
	// test1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// /mnt/eais=eais,/models=eais/models
	Volumes *string `json:"Volumes,omitempty" xml:"Volumes,omitempty"`
}

func (s CreateEaiEciRequestEciContainer) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiEciRequestEciContainer) GoString() string {
	return s.String()
}

func (s *CreateEaiEciRequestEciContainer) GetArg() *string {
	return s.Arg
}

func (s *CreateEaiEciRequestEciContainer) GetCommand() *string {
	return s.Command
}

func (s *CreateEaiEciRequestEciContainer) GetImage() *string {
	return s.Image
}

func (s *CreateEaiEciRequestEciContainer) GetName() *string {
	return s.Name
}

func (s *CreateEaiEciRequestEciContainer) GetVolumes() *string {
	return s.Volumes
}

func (s *CreateEaiEciRequestEciContainer) SetArg(v string) *CreateEaiEciRequestEciContainer {
	s.Arg = &v
	return s
}

func (s *CreateEaiEciRequestEciContainer) SetCommand(v string) *CreateEaiEciRequestEciContainer {
	s.Command = &v
	return s
}

func (s *CreateEaiEciRequestEciContainer) SetImage(v string) *CreateEaiEciRequestEciContainer {
	s.Image = &v
	return s
}

func (s *CreateEaiEciRequestEciContainer) SetName(v string) *CreateEaiEciRequestEciContainer {
	s.Name = &v
	return s
}

func (s *CreateEaiEciRequestEciContainer) SetVolumes(v string) *CreateEaiEciRequestEciContainer {
	s.Volumes = &v
	return s
}

func (s *CreateEaiEciRequestEciContainer) Validate() error {
	return dara.Validate(s)
}

type CreateEaiEciRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEaiEciRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiEciRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEaiEciRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateEaiEciRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateEaiEciRequestTag) SetKey(v string) *CreateEaiEciRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEaiEciRequestTag) SetValue(v string) *CreateEaiEciRequestTag {
	s.Value = &v
	return s
}

func (s *CreateEaiEciRequestTag) Validate() error {
	return dara.Validate(s)
}
