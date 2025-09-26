// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAuthType(v string) *TriggerConfig
	GetAuthType() *string
	SetMethods(v []*string) *TriggerConfig
	GetMethods() []*string
}

type TriggerConfig struct {
	AuthType *string   `json:"authType,omitempty" xml:"authType,omitempty"`
	Methods  []*string `json:"methods,omitempty" xml:"methods,omitempty" type:"Repeated"`
}

func (s TriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s TriggerConfig) GoString() string {
	return s.String()
}

func (s *TriggerConfig) GetAuthType() *string {
	return s.AuthType
}

func (s *TriggerConfig) GetMethods() []*string {
	return s.Methods
}

func (s *TriggerConfig) SetAuthType(v string) *TriggerConfig {
	s.AuthType = &v
	return s
}

func (s *TriggerConfig) SetMethods(v []*string) *TriggerConfig {
	s.Methods = v
	return s
}

func (s *TriggerConfig) Validate() error {
	return dara.Validate(s)
}
