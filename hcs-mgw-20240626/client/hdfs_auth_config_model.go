// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHdfsAuthConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAuthType(v string) *HdfsAuthConfig
	GetAuthType() *string
	SetUserName(v string) *HdfsAuthConfig
	GetUserName() *string
}

type HdfsAuthConfig struct {
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s HdfsAuthConfig) String() string {
	return dara.Prettify(s)
}

func (s HdfsAuthConfig) GoString() string {
	return s.String()
}

func (s *HdfsAuthConfig) GetAuthType() *string {
	return s.AuthType
}

func (s *HdfsAuthConfig) GetUserName() *string {
	return s.UserName
}

func (s *HdfsAuthConfig) SetAuthType(v string) *HdfsAuthConfig {
	s.AuthType = &v
	return s
}

func (s *HdfsAuthConfig) SetUserName(v string) *HdfsAuthConfig {
	s.UserName = &v
	return s
}

func (s *HdfsAuthConfig) Validate() error {
	return dara.Validate(s)
}
