// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPathConfig interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *PathConfig
	GetApplicationName() *string
	SetPath(v string) *PathConfig
	GetPath() *string
}

type PathConfig struct {
	ApplicationName *string `json:"applicationName,omitempty" xml:"applicationName,omitempty"`
	Path            *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s PathConfig) String() string {
	return dara.Prettify(s)
}

func (s PathConfig) GoString() string {
	return s.String()
}

func (s *PathConfig) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *PathConfig) GetPath() *string {
	return s.Path
}

func (s *PathConfig) SetApplicationName(v string) *PathConfig {
	s.ApplicationName = &v
	return s
}

func (s *PathConfig) SetPath(v string) *PathConfig {
	s.Path = &v
	return s
}

func (s *PathConfig) Validate() error {
	return dara.Validate(s)
}
