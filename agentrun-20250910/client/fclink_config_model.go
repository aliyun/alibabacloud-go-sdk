// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFCLinkConfig interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionName(v string) *FCLinkConfig
	GetFunctionName() *string
	SetVersion(v string) *FCLinkConfig
	GetVersion() *string
}

type FCLinkConfig struct {
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	Version      *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s FCLinkConfig) String() string {
	return dara.Prettify(s)
}

func (s FCLinkConfig) GoString() string {
	return s.String()
}

func (s *FCLinkConfig) GetFunctionName() *string {
	return s.FunctionName
}

func (s *FCLinkConfig) GetVersion() *string {
	return s.Version
}

func (s *FCLinkConfig) SetFunctionName(v string) *FCLinkConfig {
	s.FunctionName = &v
	return s
}

func (s *FCLinkConfig) SetVersion(v string) *FCLinkConfig {
	s.Version = &v
	return s
}

func (s *FCLinkConfig) Validate() error {
	return dara.Validate(s)
}
