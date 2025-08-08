// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRepositorySourceConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCodeVersion(v *CodeVersionReference) *RepositorySourceConfig
	GetCodeVersion() *CodeVersionReference
	SetFilter(v *EventFilterConfig) *RepositorySourceConfig
	GetFilter() *EventFilterConfig
	SetRepositoryName(v string) *RepositorySourceConfig
	GetRepositoryName() *string
}

type RepositorySourceConfig struct {
	CodeVersion *CodeVersionReference `json:"codeVersion,omitempty" xml:"codeVersion,omitempty"`
	Filter      *EventFilterConfig    `json:"filter,omitempty" xml:"filter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-repository
	RepositoryName *string `json:"repositoryName,omitempty" xml:"repositoryName,omitempty"`
}

func (s RepositorySourceConfig) String() string {
	return dara.Prettify(s)
}

func (s RepositorySourceConfig) GoString() string {
	return s.String()
}

func (s *RepositorySourceConfig) GetCodeVersion() *CodeVersionReference {
	return s.CodeVersion
}

func (s *RepositorySourceConfig) GetFilter() *EventFilterConfig {
	return s.Filter
}

func (s *RepositorySourceConfig) GetRepositoryName() *string {
	return s.RepositoryName
}

func (s *RepositorySourceConfig) SetCodeVersion(v *CodeVersionReference) *RepositorySourceConfig {
	s.CodeVersion = v
	return s
}

func (s *RepositorySourceConfig) SetFilter(v *EventFilterConfig) *RepositorySourceConfig {
	s.Filter = v
	return s
}

func (s *RepositorySourceConfig) SetRepositoryName(v string) *RepositorySourceConfig {
	s.RepositoryName = &v
	return s
}

func (s *RepositorySourceConfig) Validate() error {
	return dara.Validate(s)
}
