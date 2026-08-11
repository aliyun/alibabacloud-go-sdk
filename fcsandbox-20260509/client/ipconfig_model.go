// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIPConfig interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *IPConfig
	GetDescription() *string
	SetIpAddress(v string) *IPConfig
	GetIpAddress() *string
}

type IPConfig struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	IpAddress   *string `json:"ipAddress,omitempty" xml:"ipAddress,omitempty"`
}

func (s IPConfig) String() string {
	return dara.Prettify(s)
}

func (s IPConfig) GoString() string {
	return s.String()
}

func (s *IPConfig) GetDescription() *string {
	return s.Description
}

func (s *IPConfig) GetIpAddress() *string {
	return s.IpAddress
}

func (s *IPConfig) SetDescription(v string) *IPConfig {
	s.Description = &v
	return s
}

func (s *IPConfig) SetIpAddress(v string) *IPConfig {
	s.IpAddress = &v
	return s
}

func (s *IPConfig) Validate() error {
	return dara.Validate(s)
}
