// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOriginPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *CreateOriginPoolRequest
	GetEnabled() *bool
	SetName(v string) *CreateOriginPoolRequest
	GetName() *string
	SetOrigins(v []*CreateOriginPoolRequestOrigins) *CreateOriginPoolRequest
	GetOrigins() []*CreateOriginPoolRequestOrigins
	SetSiteId(v int64) *CreateOriginPoolRequest
	GetSiteId() *int64
}

type CreateOriginPoolRequest struct {
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// This parameter is required.
	Name    *string                           `json:"Name,omitempty" xml:"Name,omitempty"`
	Origins []*CreateOriginPoolRequestOrigins `json:"Origins,omitempty" xml:"Origins,omitempty" type:"Repeated"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s CreateOriginPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOriginPoolRequest) GoString() string {
	return s.String()
}

func (s *CreateOriginPoolRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateOriginPoolRequest) GetName() *string {
	return s.Name
}

func (s *CreateOriginPoolRequest) GetOrigins() []*CreateOriginPoolRequestOrigins {
	return s.Origins
}

func (s *CreateOriginPoolRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateOriginPoolRequest) SetEnabled(v bool) *CreateOriginPoolRequest {
	s.Enabled = &v
	return s
}

func (s *CreateOriginPoolRequest) SetName(v string) *CreateOriginPoolRequest {
	s.Name = &v
	return s
}

func (s *CreateOriginPoolRequest) SetOrigins(v []*CreateOriginPoolRequestOrigins) *CreateOriginPoolRequest {
	s.Origins = v
	return s
}

func (s *CreateOriginPoolRequest) SetSiteId(v int64) *CreateOriginPoolRequest {
	s.SiteId = &v
	return s
}

func (s *CreateOriginPoolRequest) Validate() error {
	if s.Origins != nil {
		for _, item := range s.Origins {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateOriginPoolRequestOrigins struct {
	Address         *string                                 `json:"Address,omitempty" xml:"Address,omitempty"`
	AuthConf        *CreateOriginPoolRequestOriginsAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	Enabled         *bool                                   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Header          interface{}                             `json:"Header,omitempty" xml:"Header,omitempty"`
	IpVersionPolicy *string                                 `json:"IpVersionPolicy,omitempty" xml:"IpVersionPolicy,omitempty"`
	Name            *string                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	Type            *string                                 `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight          *int32                                  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateOriginPoolRequestOrigins) String() string {
	return dara.Prettify(s)
}

func (s CreateOriginPoolRequestOrigins) GoString() string {
	return s.String()
}

func (s *CreateOriginPoolRequestOrigins) GetAddress() *string {
	return s.Address
}

func (s *CreateOriginPoolRequestOrigins) GetAuthConf() *CreateOriginPoolRequestOriginsAuthConf {
	return s.AuthConf
}

func (s *CreateOriginPoolRequestOrigins) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateOriginPoolRequestOrigins) GetHeader() interface{} {
	return s.Header
}

func (s *CreateOriginPoolRequestOrigins) GetIpVersionPolicy() *string {
	return s.IpVersionPolicy
}

func (s *CreateOriginPoolRequestOrigins) GetName() *string {
	return s.Name
}

func (s *CreateOriginPoolRequestOrigins) GetType() *string {
	return s.Type
}

func (s *CreateOriginPoolRequestOrigins) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateOriginPoolRequestOrigins) SetAddress(v string) *CreateOriginPoolRequestOrigins {
	s.Address = &v
	return s
}

func (s *CreateOriginPoolRequestOrigins) SetAuthConf(v *CreateOriginPoolRequestOriginsAuthConf) *CreateOriginPoolRequestOrigins {
	s.AuthConf = v
	return s
}

func (s *CreateOriginPoolRequestOrigins) SetEnabled(v bool) *CreateOriginPoolRequestOrigins {
	s.Enabled = &v
	return s
}

func (s *CreateOriginPoolRequestOrigins) SetHeader(v interface{}) *CreateOriginPoolRequestOrigins {
	s.Header = v
	return s
}

func (s *CreateOriginPoolRequestOrigins) SetIpVersionPolicy(v string) *CreateOriginPoolRequestOrigins {
	s.IpVersionPolicy = &v
	return s
}

func (s *CreateOriginPoolRequestOrigins) SetName(v string) *CreateOriginPoolRequestOrigins {
	s.Name = &v
	return s
}

func (s *CreateOriginPoolRequestOrigins) SetType(v string) *CreateOriginPoolRequestOrigins {
	s.Type = &v
	return s
}

func (s *CreateOriginPoolRequestOrigins) SetWeight(v int32) *CreateOriginPoolRequestOrigins {
	s.Weight = &v
	return s
}

func (s *CreateOriginPoolRequestOrigins) Validate() error {
	if s.AuthConf != nil {
		if err := s.AuthConf.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOriginPoolRequestOriginsAuthConf struct {
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	AuthType  *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	Region    *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateOriginPoolRequestOriginsAuthConf) String() string {
	return dara.Prettify(s)
}

func (s CreateOriginPoolRequestOriginsAuthConf) GoString() string {
	return s.String()
}

func (s *CreateOriginPoolRequestOriginsAuthConf) GetAccessKey() *string {
	return s.AccessKey
}

func (s *CreateOriginPoolRequestOriginsAuthConf) GetAuthType() *string {
	return s.AuthType
}

func (s *CreateOriginPoolRequestOriginsAuthConf) GetRegion() *string {
	return s.Region
}

func (s *CreateOriginPoolRequestOriginsAuthConf) GetSecretKey() *string {
	return s.SecretKey
}

func (s *CreateOriginPoolRequestOriginsAuthConf) GetVersion() *string {
	return s.Version
}

func (s *CreateOriginPoolRequestOriginsAuthConf) SetAccessKey(v string) *CreateOriginPoolRequestOriginsAuthConf {
	s.AccessKey = &v
	return s
}

func (s *CreateOriginPoolRequestOriginsAuthConf) SetAuthType(v string) *CreateOriginPoolRequestOriginsAuthConf {
	s.AuthType = &v
	return s
}

func (s *CreateOriginPoolRequestOriginsAuthConf) SetRegion(v string) *CreateOriginPoolRequestOriginsAuthConf {
	s.Region = &v
	return s
}

func (s *CreateOriginPoolRequestOriginsAuthConf) SetSecretKey(v string) *CreateOriginPoolRequestOriginsAuthConf {
	s.SecretKey = &v
	return s
}

func (s *CreateOriginPoolRequestOriginsAuthConf) SetVersion(v string) *CreateOriginPoolRequestOriginsAuthConf {
	s.Version = &v
	return s
}

func (s *CreateOriginPoolRequestOriginsAuthConf) Validate() error {
	return dara.Validate(s)
}
