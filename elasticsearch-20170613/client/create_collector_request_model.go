// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCollectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollectorPaths(v []*string) *CreateCollectorRequest
	GetCollectorPaths() []*string
	SetConfigs(v []*CreateCollectorRequestConfigs) *CreateCollectorRequest
	GetConfigs() []*CreateCollectorRequestConfigs
	SetDryRun(v bool) *CreateCollectorRequest
	GetDryRun() *bool
	SetExtendConfigs(v []map[string]interface{}) *CreateCollectorRequest
	GetExtendConfigs() []map[string]interface{}
	SetName(v string) *CreateCollectorRequest
	GetName() *string
	SetResType(v string) *CreateCollectorRequest
	GetResType() *string
	SetResVersion(v string) *CreateCollectorRequest
	GetResVersion() *string
	SetVpcId(v string) *CreateCollectorRequest
	GetVpcId() *string
	SetClientToken(v string) *CreateCollectorRequest
	GetClientToken() *string
}

type CreateCollectorRequest struct {
	CollectorPaths []*string `json:"collectorPaths,omitempty" xml:"collectorPaths,omitempty" type:"Repeated"`
	// This parameter is required.
	Configs []*CreateCollectorRequestConfigs `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// This parameter is required.
	ExtendConfigs []map[string]interface{} `json:"extendConfigs,omitempty" xml:"extendConfigs,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// ct-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fileBeat
	ResType *string `json:"resType,omitempty" xml:"resType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6.8.5_with_community
	ResVersion *string `json:"resVersion,omitempty" xml:"resVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp12nu14urf0upaf*****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// The ID of the created crawer.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateCollectorRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCollectorRequest) GoString() string {
	return s.String()
}

func (s *CreateCollectorRequest) GetCollectorPaths() []*string {
	return s.CollectorPaths
}

func (s *CreateCollectorRequest) GetConfigs() []*CreateCollectorRequestConfigs {
	return s.Configs
}

func (s *CreateCollectorRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateCollectorRequest) GetExtendConfigs() []map[string]interface{} {
	return s.ExtendConfigs
}

func (s *CreateCollectorRequest) GetName() *string {
	return s.Name
}

func (s *CreateCollectorRequest) GetResType() *string {
	return s.ResType
}

func (s *CreateCollectorRequest) GetResVersion() *string {
	return s.ResVersion
}

func (s *CreateCollectorRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateCollectorRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCollectorRequest) SetCollectorPaths(v []*string) *CreateCollectorRequest {
	s.CollectorPaths = v
	return s
}

func (s *CreateCollectorRequest) SetConfigs(v []*CreateCollectorRequestConfigs) *CreateCollectorRequest {
	s.Configs = v
	return s
}

func (s *CreateCollectorRequest) SetDryRun(v bool) *CreateCollectorRequest {
	s.DryRun = &v
	return s
}

func (s *CreateCollectorRequest) SetExtendConfigs(v []map[string]interface{}) *CreateCollectorRequest {
	s.ExtendConfigs = v
	return s
}

func (s *CreateCollectorRequest) SetName(v string) *CreateCollectorRequest {
	s.Name = &v
	return s
}

func (s *CreateCollectorRequest) SetResType(v string) *CreateCollectorRequest {
	s.ResType = &v
	return s
}

func (s *CreateCollectorRequest) SetResVersion(v string) *CreateCollectorRequest {
	s.ResVersion = &v
	return s
}

func (s *CreateCollectorRequest) SetVpcId(v string) *CreateCollectorRequest {
	s.VpcId = &v
	return s
}

func (s *CreateCollectorRequest) SetClientToken(v string) *CreateCollectorRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCollectorRequest) Validate() error {
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCollectorRequestConfigs struct {
	// This parameter is required.
	//
	// example:
	//
	// "filebeat.inputs:xxx"
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// filebeat.yml
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s CreateCollectorRequestConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateCollectorRequestConfigs) GoString() string {
	return s.String()
}

func (s *CreateCollectorRequestConfigs) GetContent() *string {
	return s.Content
}

func (s *CreateCollectorRequestConfigs) GetFileName() *string {
	return s.FileName
}

func (s *CreateCollectorRequestConfigs) SetContent(v string) *CreateCollectorRequestConfigs {
	s.Content = &v
	return s
}

func (s *CreateCollectorRequestConfigs) SetFileName(v string) *CreateCollectorRequestConfigs {
	s.FileName = &v
	return s
}

func (s *CreateCollectorRequestConfigs) Validate() error {
	return dara.Validate(s)
}
