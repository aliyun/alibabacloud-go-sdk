// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNASMountConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnableTLS(v bool) *NASMountConfig
	GetEnableTLS() *bool
	SetMountDir(v string) *NASMountConfig
	GetMountDir() *string
	SetServerAddr(v string) *NASMountConfig
	GetServerAddr() *string
}

type NASMountConfig struct {
	// Specifies whether to enable encryption in transit. This option is supported only for general-purpose NAS.
	EnableTLS *bool `json:"enableTLS,omitempty" xml:"enableTLS,omitempty"`
	// Specifies the local mount directory.
	//
	// example:
	//
	// /home/test
	MountDir *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	// Specifies the NAS server address.
	//
	// example:
	//
	// ***-uni85.cn-hangzhou.nas.com:/
	ServerAddr *string `json:"serverAddr,omitempty" xml:"serverAddr,omitempty"`
}

func (s NASMountConfig) String() string {
	return dara.Prettify(s)
}

func (s NASMountConfig) GoString() string {
	return s.String()
}

func (s *NASMountConfig) GetEnableTLS() *bool {
	return s.EnableTLS
}

func (s *NASMountConfig) GetMountDir() *string {
	return s.MountDir
}

func (s *NASMountConfig) GetServerAddr() *string {
	return s.ServerAddr
}

func (s *NASMountConfig) SetEnableTLS(v bool) *NASMountConfig {
	s.EnableTLS = &v
	return s
}

func (s *NASMountConfig) SetMountDir(v string) *NASMountConfig {
	s.MountDir = &v
	return s
}

func (s *NASMountConfig) SetServerAddr(v string) *NASMountConfig {
	s.ServerAddr = &v
	return s
}

func (s *NASMountConfig) Validate() error {
	return dara.Validate(s)
}
