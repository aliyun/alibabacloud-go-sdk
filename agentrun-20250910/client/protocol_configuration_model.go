// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProtocolConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetProtocolSettings(v []*ProtocolSettings) *ProtocolConfiguration
	GetProtocolSettings() []*ProtocolSettings
	SetType(v string) *ProtocolConfiguration
	GetType() *string
}

type ProtocolConfiguration struct {
	// 详细的协议配置信息
	ProtocolSettings []*ProtocolSettings `json:"protocolSettings,omitempty" xml:"protocolSettings,omitempty" type:"Repeated"`
	// Deprecated
	//
	// example:
	//
	// HTTP
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ProtocolConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ProtocolConfiguration) GoString() string {
	return s.String()
}

func (s *ProtocolConfiguration) GetProtocolSettings() []*ProtocolSettings {
	return s.ProtocolSettings
}

func (s *ProtocolConfiguration) GetType() *string {
	return s.Type
}

func (s *ProtocolConfiguration) SetProtocolSettings(v []*ProtocolSettings) *ProtocolConfiguration {
	s.ProtocolSettings = v
	return s
}

func (s *ProtocolConfiguration) SetType(v string) *ProtocolConfiguration {
	s.Type = &v
	return s
}

func (s *ProtocolConfiguration) Validate() error {
	if s.ProtocolSettings != nil {
		for _, item := range s.ProtocolSettings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
