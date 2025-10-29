// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpApiInfoByName interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *HttpApiInfoByName
	GetGatewayId() *string
	SetName(v string) *HttpApiInfoByName
	GetName() *string
	SetType(v string) *HttpApiInfoByName
	GetType() *string
	SetVersionEnabled(v bool) *HttpApiInfoByName
	GetVersionEnabled() *bool
	SetVersionedHttpApis(v []*HttpApiApiInfo) *HttpApiInfoByName
	GetVersionedHttpApis() []*HttpApiApiInfo
}

type HttpApiInfoByName struct {
	// example:
	//
	// gw-xx
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// example:
	//
	// test-api
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Http
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// true
	VersionEnabled    *bool             `json:"versionEnabled,omitempty" xml:"versionEnabled,omitempty"`
	VersionedHttpApis []*HttpApiApiInfo `json:"versionedHttpApis,omitempty" xml:"versionedHttpApis,omitempty" type:"Repeated"`
}

func (s HttpApiInfoByName) String() string {
	return dara.Prettify(s)
}

func (s HttpApiInfoByName) GoString() string {
	return s.String()
}

func (s *HttpApiInfoByName) GetGatewayId() *string {
	return s.GatewayId
}

func (s *HttpApiInfoByName) GetName() *string {
	return s.Name
}

func (s *HttpApiInfoByName) GetType() *string {
	return s.Type
}

func (s *HttpApiInfoByName) GetVersionEnabled() *bool {
	return s.VersionEnabled
}

func (s *HttpApiInfoByName) GetVersionedHttpApis() []*HttpApiApiInfo {
	return s.VersionedHttpApis
}

func (s *HttpApiInfoByName) SetGatewayId(v string) *HttpApiInfoByName {
	s.GatewayId = &v
	return s
}

func (s *HttpApiInfoByName) SetName(v string) *HttpApiInfoByName {
	s.Name = &v
	return s
}

func (s *HttpApiInfoByName) SetType(v string) *HttpApiInfoByName {
	s.Type = &v
	return s
}

func (s *HttpApiInfoByName) SetVersionEnabled(v bool) *HttpApiInfoByName {
	s.VersionEnabled = &v
	return s
}

func (s *HttpApiInfoByName) SetVersionedHttpApis(v []*HttpApiApiInfo) *HttpApiInfoByName {
	s.VersionedHttpApis = v
	return s
}

func (s *HttpApiInfoByName) Validate() error {
	if s.VersionedHttpApis != nil {
		for _, item := range s.VersionedHttpApis {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
