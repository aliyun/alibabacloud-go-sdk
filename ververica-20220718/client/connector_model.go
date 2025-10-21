// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConnector interface {
	dara.Model
	String() string
	GoString() string
	SetCreator(v string) *Connector
	GetCreator() *string
	SetCreatorName(v string) *Connector
	GetCreatorName() *string
	SetDependencies(v []*string) *Connector
	GetDependencies() []*string
	SetLookup(v bool) *Connector
	GetLookup() *bool
	SetModifier(v string) *Connector
	GetModifier() *string
	SetModifierName(v string) *Connector
	GetModifierName() *string
	SetName(v string) *Connector
	GetName() *string
	SetProperties(v []*Property) *Connector
	GetProperties() []*Property
	SetSink(v bool) *Connector
	GetSink() *bool
	SetSource(v bool) *Connector
	GetSource() *bool
	SetSupportedFormats(v []*string) *Connector
	GetSupportedFormats() []*string
	SetType(v string) *Connector
	GetType() *string
}

type Connector struct {
	Creator          *string     `json:"creator,omitempty" xml:"creator,omitempty"`
	CreatorName      *string     `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	Dependencies     []*string   `json:"dependencies,omitempty" xml:"dependencies,omitempty" type:"Repeated"`
	Lookup           *bool       `json:"lookup,omitempty" xml:"lookup,omitempty"`
	Modifier         *string     `json:"modifier,omitempty" xml:"modifier,omitempty"`
	ModifierName     *string     `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
	Name             *string     `json:"name,omitempty" xml:"name,omitempty"`
	Properties       []*Property `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
	Sink             *bool       `json:"sink,omitempty" xml:"sink,omitempty"`
	Source           *bool       `json:"source,omitempty" xml:"source,omitempty"`
	SupportedFormats []*string   `json:"supportedFormats,omitempty" xml:"supportedFormats,omitempty" type:"Repeated"`
	Type             *string     `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Connector) String() string {
	return dara.Prettify(s)
}

func (s Connector) GoString() string {
	return s.String()
}

func (s *Connector) GetCreator() *string {
	return s.Creator
}

func (s *Connector) GetCreatorName() *string {
	return s.CreatorName
}

func (s *Connector) GetDependencies() []*string {
	return s.Dependencies
}

func (s *Connector) GetLookup() *bool {
	return s.Lookup
}

func (s *Connector) GetModifier() *string {
	return s.Modifier
}

func (s *Connector) GetModifierName() *string {
	return s.ModifierName
}

func (s *Connector) GetName() *string {
	return s.Name
}

func (s *Connector) GetProperties() []*Property {
	return s.Properties
}

func (s *Connector) GetSink() *bool {
	return s.Sink
}

func (s *Connector) GetSource() *bool {
	return s.Source
}

func (s *Connector) GetSupportedFormats() []*string {
	return s.SupportedFormats
}

func (s *Connector) GetType() *string {
	return s.Type
}

func (s *Connector) SetCreator(v string) *Connector {
	s.Creator = &v
	return s
}

func (s *Connector) SetCreatorName(v string) *Connector {
	s.CreatorName = &v
	return s
}

func (s *Connector) SetDependencies(v []*string) *Connector {
	s.Dependencies = v
	return s
}

func (s *Connector) SetLookup(v bool) *Connector {
	s.Lookup = &v
	return s
}

func (s *Connector) SetModifier(v string) *Connector {
	s.Modifier = &v
	return s
}

func (s *Connector) SetModifierName(v string) *Connector {
	s.ModifierName = &v
	return s
}

func (s *Connector) SetName(v string) *Connector {
	s.Name = &v
	return s
}

func (s *Connector) SetProperties(v []*Property) *Connector {
	s.Properties = v
	return s
}

func (s *Connector) SetSink(v bool) *Connector {
	s.Sink = &v
	return s
}

func (s *Connector) SetSource(v bool) *Connector {
	s.Source = &v
	return s
}

func (s *Connector) SetSupportedFormats(v []*string) *Connector {
	s.SupportedFormats = v
	return s
}

func (s *Connector) SetType(v string) *Connector {
	s.Type = &v
	return s
}

func (s *Connector) Validate() error {
	if s.Properties != nil {
		for _, item := range s.Properties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
