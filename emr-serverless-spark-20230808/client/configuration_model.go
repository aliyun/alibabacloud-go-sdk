// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetConfigFileName(v string) *Configuration
	GetConfigFileName() *string
	SetConfigItemKey(v string) *Configuration
	GetConfigItemKey() *string
	SetConfigItemValue(v string) *Configuration
	GetConfigItemValue() *string
}

type Configuration struct {
	ConfigFileName  *string `json:"configFileName,omitempty" xml:"configFileName,omitempty"`
	ConfigItemKey   *string `json:"configItemKey,omitempty" xml:"configItemKey,omitempty"`
	ConfigItemValue *string `json:"configItemValue,omitempty" xml:"configItemValue,omitempty"`
}

func (s Configuration) String() string {
	return dara.Prettify(s)
}

func (s Configuration) GoString() string {
	return s.String()
}

func (s *Configuration) GetConfigFileName() *string {
	return s.ConfigFileName
}

func (s *Configuration) GetConfigItemKey() *string {
	return s.ConfigItemKey
}

func (s *Configuration) GetConfigItemValue() *string {
	return s.ConfigItemValue
}

func (s *Configuration) SetConfigFileName(v string) *Configuration {
	s.ConfigFileName = &v
	return s
}

func (s *Configuration) SetConfigItemKey(v string) *Configuration {
	s.ConfigItemKey = &v
	return s
}

func (s *Configuration) SetConfigItemValue(v string) *Configuration {
	s.ConfigItemValue = &v
	return s
}

func (s *Configuration) Validate() error {
	return dara.Validate(s)
}
