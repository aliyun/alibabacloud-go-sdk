// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewComponentCrdSchemaCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *PreviewComponentCrdSchemaCmd
	GetAddress() *string
	SetConfiguration(v []*ConfigModel) *PreviewComponentCrdSchemaCmd
	GetConfiguration() []*ConfigModel
	SetCredentials(v string) *PreviewComponentCrdSchemaCmd
	GetCredentials() *string
	SetId(v int64) *PreviewComponentCrdSchemaCmd
	GetId() *int64
	SetIsCustom(v bool) *PreviewComponentCrdSchemaCmd
	GetIsCustom() *bool
	SetName(v string) *PreviewComponentCrdSchemaCmd
	GetName() *string
	SetReesourceVersion(v string) *PreviewComponentCrdSchemaCmd
	GetReesourceVersion() *string
	SetResourceId(v int64) *PreviewComponentCrdSchemaCmd
	GetResourceId() *int64
	SetScope(v string) *PreviewComponentCrdSchemaCmd
	GetScope() *string
	SetTemplateId(v int64) *PreviewComponentCrdSchemaCmd
	GetTemplateId() *int64
	SetType(v string) *PreviewComponentCrdSchemaCmd
	GetType() *string
}

type PreviewComponentCrdSchemaCmd struct {
	Address          *string        `json:"address,omitempty" xml:"address,omitempty"`
	Configuration    []*ConfigModel `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Repeated"`
	Credentials      *string        `json:"credentials,omitempty" xml:"credentials,omitempty"`
	Id               *int64         `json:"id,omitempty" xml:"id,omitempty"`
	IsCustom         *bool          `json:"isCustom,omitempty" xml:"isCustom,omitempty"`
	Name             *string        `json:"name,omitempty" xml:"name,omitempty"`
	ReesourceVersion *string        `json:"reesourceVersion,omitempty" xml:"reesourceVersion,omitempty"`
	ResourceId       *int64         `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	Scope            *string        `json:"scope,omitempty" xml:"scope,omitempty"`
	TemplateId       *int64         `json:"templateId,omitempty" xml:"templateId,omitempty"`
	Type             *string        `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PreviewComponentCrdSchemaCmd) String() string {
	return dara.Prettify(s)
}

func (s PreviewComponentCrdSchemaCmd) GoString() string {
	return s.String()
}

func (s *PreviewComponentCrdSchemaCmd) GetAddress() *string {
	return s.Address
}

func (s *PreviewComponentCrdSchemaCmd) GetConfiguration() []*ConfigModel {
	return s.Configuration
}

func (s *PreviewComponentCrdSchemaCmd) GetCredentials() *string {
	return s.Credentials
}

func (s *PreviewComponentCrdSchemaCmd) GetId() *int64 {
	return s.Id
}

func (s *PreviewComponentCrdSchemaCmd) GetIsCustom() *bool {
	return s.IsCustom
}

func (s *PreviewComponentCrdSchemaCmd) GetName() *string {
	return s.Name
}

func (s *PreviewComponentCrdSchemaCmd) GetReesourceVersion() *string {
	return s.ReesourceVersion
}

func (s *PreviewComponentCrdSchemaCmd) GetResourceId() *int64 {
	return s.ResourceId
}

func (s *PreviewComponentCrdSchemaCmd) GetScope() *string {
	return s.Scope
}

func (s *PreviewComponentCrdSchemaCmd) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *PreviewComponentCrdSchemaCmd) GetType() *string {
	return s.Type
}

func (s *PreviewComponentCrdSchemaCmd) SetAddress(v string) *PreviewComponentCrdSchemaCmd {
	s.Address = &v
	return s
}

func (s *PreviewComponentCrdSchemaCmd) SetConfiguration(v []*ConfigModel) *PreviewComponentCrdSchemaCmd {
	s.Configuration = v
	return s
}

func (s *PreviewComponentCrdSchemaCmd) SetCredentials(v string) *PreviewComponentCrdSchemaCmd {
	s.Credentials = &v
	return s
}

func (s *PreviewComponentCrdSchemaCmd) SetId(v int64) *PreviewComponentCrdSchemaCmd {
	s.Id = &v
	return s
}

func (s *PreviewComponentCrdSchemaCmd) SetIsCustom(v bool) *PreviewComponentCrdSchemaCmd {
	s.IsCustom = &v
	return s
}

func (s *PreviewComponentCrdSchemaCmd) SetName(v string) *PreviewComponentCrdSchemaCmd {
	s.Name = &v
	return s
}

func (s *PreviewComponentCrdSchemaCmd) SetReesourceVersion(v string) *PreviewComponentCrdSchemaCmd {
	s.ReesourceVersion = &v
	return s
}

func (s *PreviewComponentCrdSchemaCmd) SetResourceId(v int64) *PreviewComponentCrdSchemaCmd {
	s.ResourceId = &v
	return s
}

func (s *PreviewComponentCrdSchemaCmd) SetScope(v string) *PreviewComponentCrdSchemaCmd {
	s.Scope = &v
	return s
}

func (s *PreviewComponentCrdSchemaCmd) SetTemplateId(v int64) *PreviewComponentCrdSchemaCmd {
	s.TemplateId = &v
	return s
}

func (s *PreviewComponentCrdSchemaCmd) SetType(v string) *PreviewComponentCrdSchemaCmd {
	s.Type = &v
	return s
}

func (s *PreviewComponentCrdSchemaCmd) Validate() error {
	if s.Configuration != nil {
		for _, item := range s.Configuration {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
