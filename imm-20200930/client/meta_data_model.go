// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaData interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifier(v string) *MetaData
	GetIdentifier() *string
	SetProvider(v string) *MetaData
	GetProvider() *string
	SetVersion(v string) *MetaData
	GetVersion() *string
}

type MetaData struct {
	// example:
	//
	// detection
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// example:
	//
	// imm
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// example:
	//
	// v1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s MetaData) String() string {
	return dara.Prettify(s)
}

func (s MetaData) GoString() string {
	return s.String()
}

func (s *MetaData) GetIdentifier() *string {
	return s.Identifier
}

func (s *MetaData) GetProvider() *string {
	return s.Provider
}

func (s *MetaData) GetVersion() *string {
	return s.Version
}

func (s *MetaData) SetIdentifier(v string) *MetaData {
	s.Identifier = &v
	return s
}

func (s *MetaData) SetProvider(v string) *MetaData {
	s.Provider = &v
	return s
}

func (s *MetaData) SetVersion(v string) *MetaData {
	s.Version = &v
	return s
}

func (s *MetaData) Validate() error {
	return dara.Validate(s)
}
