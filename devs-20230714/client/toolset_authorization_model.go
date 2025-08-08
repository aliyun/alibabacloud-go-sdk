// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iToolsetAuthorization interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfig(v map[string]*string) *ToolsetAuthorization
	GetAuthConfig() map[string]*string
	SetType(v string) *ToolsetAuthorization
	GetType() *string
}

type ToolsetAuthorization struct {
	AuthConfig map[string]*string `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// example:
	//
	// apiKey
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ToolsetAuthorization) String() string {
	return dara.Prettify(s)
}

func (s ToolsetAuthorization) GoString() string {
	return s.String()
}

func (s *ToolsetAuthorization) GetAuthConfig() map[string]*string {
	return s.AuthConfig
}

func (s *ToolsetAuthorization) GetType() *string {
	return s.Type
}

func (s *ToolsetAuthorization) SetAuthConfig(v map[string]*string) *ToolsetAuthorization {
	s.AuthConfig = v
	return s
}

func (s *ToolsetAuthorization) SetType(v string) *ToolsetAuthorization {
	s.Type = &v
	return s
}

func (s *ToolsetAuthorization) Validate() error {
	return dara.Validate(s)
}
