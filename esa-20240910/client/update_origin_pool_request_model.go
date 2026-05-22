// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOriginPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *UpdateOriginPoolRequest
	GetEnabled() *bool
	SetId(v int64) *UpdateOriginPoolRequest
	GetId() *int64
	SetOrigins(v []*UpdateOriginPoolRequestOrigins) *UpdateOriginPoolRequest
	GetOrigins() []*UpdateOriginPoolRequestOrigins
	SetSiteId(v int64) *UpdateOriginPoolRequest
	GetSiteId() *int64
}

type UpdateOriginPoolRequest struct {
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// This parameter is required.
	Id      *int64                            `json:"Id,omitempty" xml:"Id,omitempty"`
	Origins []*UpdateOriginPoolRequestOrigins `json:"Origins,omitempty" xml:"Origins,omitempty" type:"Repeated"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateOriginPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOriginPoolRequest) GoString() string {
	return s.String()
}

func (s *UpdateOriginPoolRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateOriginPoolRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateOriginPoolRequest) GetOrigins() []*UpdateOriginPoolRequestOrigins {
	return s.Origins
}

func (s *UpdateOriginPoolRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateOriginPoolRequest) SetEnabled(v bool) *UpdateOriginPoolRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateOriginPoolRequest) SetId(v int64) *UpdateOriginPoolRequest {
	s.Id = &v
	return s
}

func (s *UpdateOriginPoolRequest) SetOrigins(v []*UpdateOriginPoolRequestOrigins) *UpdateOriginPoolRequest {
	s.Origins = v
	return s
}

func (s *UpdateOriginPoolRequest) SetSiteId(v int64) *UpdateOriginPoolRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateOriginPoolRequest) Validate() error {
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

type UpdateOriginPoolRequestOrigins struct {
	Address         *string                                 `json:"Address,omitempty" xml:"Address,omitempty"`
	AuthConf        *UpdateOriginPoolRequestOriginsAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	Enabled         *bool                                   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Header          interface{}                             `json:"Header,omitempty" xml:"Header,omitempty"`
	IpVersionPolicy *string                                 `json:"IpVersionPolicy,omitempty" xml:"IpVersionPolicy,omitempty"`
	Name            *string                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	Type            *string                                 `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight          *int32                                  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateOriginPoolRequestOrigins) String() string {
	return dara.Prettify(s)
}

func (s UpdateOriginPoolRequestOrigins) GoString() string {
	return s.String()
}

func (s *UpdateOriginPoolRequestOrigins) GetAddress() *string {
	return s.Address
}

func (s *UpdateOriginPoolRequestOrigins) GetAuthConf() *UpdateOriginPoolRequestOriginsAuthConf {
	return s.AuthConf
}

func (s *UpdateOriginPoolRequestOrigins) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateOriginPoolRequestOrigins) GetHeader() interface{} {
	return s.Header
}

func (s *UpdateOriginPoolRequestOrigins) GetIpVersionPolicy() *string {
	return s.IpVersionPolicy
}

func (s *UpdateOriginPoolRequestOrigins) GetName() *string {
	return s.Name
}

func (s *UpdateOriginPoolRequestOrigins) GetType() *string {
	return s.Type
}

func (s *UpdateOriginPoolRequestOrigins) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateOriginPoolRequestOrigins) SetAddress(v string) *UpdateOriginPoolRequestOrigins {
	s.Address = &v
	return s
}

func (s *UpdateOriginPoolRequestOrigins) SetAuthConf(v *UpdateOriginPoolRequestOriginsAuthConf) *UpdateOriginPoolRequestOrigins {
	s.AuthConf = v
	return s
}

func (s *UpdateOriginPoolRequestOrigins) SetEnabled(v bool) *UpdateOriginPoolRequestOrigins {
	s.Enabled = &v
	return s
}

func (s *UpdateOriginPoolRequestOrigins) SetHeader(v interface{}) *UpdateOriginPoolRequestOrigins {
	s.Header = v
	return s
}

func (s *UpdateOriginPoolRequestOrigins) SetIpVersionPolicy(v string) *UpdateOriginPoolRequestOrigins {
	s.IpVersionPolicy = &v
	return s
}

func (s *UpdateOriginPoolRequestOrigins) SetName(v string) *UpdateOriginPoolRequestOrigins {
	s.Name = &v
	return s
}

func (s *UpdateOriginPoolRequestOrigins) SetType(v string) *UpdateOriginPoolRequestOrigins {
	s.Type = &v
	return s
}

func (s *UpdateOriginPoolRequestOrigins) SetWeight(v int32) *UpdateOriginPoolRequestOrigins {
	s.Weight = &v
	return s
}

func (s *UpdateOriginPoolRequestOrigins) Validate() error {
	if s.AuthConf != nil {
		if err := s.AuthConf.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateOriginPoolRequestOriginsAuthConf struct {
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	AuthType  *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	Region    *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateOriginPoolRequestOriginsAuthConf) String() string {
	return dara.Prettify(s)
}

func (s UpdateOriginPoolRequestOriginsAuthConf) GoString() string {
	return s.String()
}

func (s *UpdateOriginPoolRequestOriginsAuthConf) GetAccessKey() *string {
	return s.AccessKey
}

func (s *UpdateOriginPoolRequestOriginsAuthConf) GetAuthType() *string {
	return s.AuthType
}

func (s *UpdateOriginPoolRequestOriginsAuthConf) GetRegion() *string {
	return s.Region
}

func (s *UpdateOriginPoolRequestOriginsAuthConf) GetSecretKey() *string {
	return s.SecretKey
}

func (s *UpdateOriginPoolRequestOriginsAuthConf) GetVersion() *string {
	return s.Version
}

func (s *UpdateOriginPoolRequestOriginsAuthConf) SetAccessKey(v string) *UpdateOriginPoolRequestOriginsAuthConf {
	s.AccessKey = &v
	return s
}

func (s *UpdateOriginPoolRequestOriginsAuthConf) SetAuthType(v string) *UpdateOriginPoolRequestOriginsAuthConf {
	s.AuthType = &v
	return s
}

func (s *UpdateOriginPoolRequestOriginsAuthConf) SetRegion(v string) *UpdateOriginPoolRequestOriginsAuthConf {
	s.Region = &v
	return s
}

func (s *UpdateOriginPoolRequestOriginsAuthConf) SetSecretKey(v string) *UpdateOriginPoolRequestOriginsAuthConf {
	s.SecretKey = &v
	return s
}

func (s *UpdateOriginPoolRequestOriginsAuthConf) SetVersion(v string) *UpdateOriginPoolRequestOriginsAuthConf {
	s.Version = &v
	return s
}

func (s *UpdateOriginPoolRequestOriginsAuthConf) Validate() error {
	return dara.Validate(s)
}
