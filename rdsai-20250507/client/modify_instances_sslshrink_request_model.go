// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstancesSSLShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCAType(v string) *ModifyInstancesSSLShrinkRequest
	GetCAType() *string
	SetInstanceNamesShrink(v string) *ModifyInstancesSSLShrinkRequest
	GetInstanceNamesShrink() *string
	SetRegionId(v string) *ModifyInstancesSSLShrinkRequest
	GetRegionId() *string
	SetSSLEnabled(v int32) *ModifyInstancesSSLShrinkRequest
	GetSSLEnabled() *int32
	SetServerCert(v string) *ModifyInstancesSSLShrinkRequest
	GetServerCert() *string
	SetServerKey(v string) *ModifyInstancesSSLShrinkRequest
	GetServerKey() *string
}

type ModifyInstancesSSLShrinkRequest struct {
	// example:
	//
	// custom
	CAType *string `json:"CAType,omitempty" xml:"CAType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [
	//
	//     "ra-supabase-xxx",
	//
	//     "ra-supabase-xxx"
	//
	//   ]
	InstanceNamesShrink *string `json:"InstanceNames,omitempty" xml:"InstanceNames,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SSLEnabled *int32 `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	// example:
	//
	// -----BEGIN CERTIFICATE-----MIID*****QqEP-----END CERTIFICATE-----
	ServerCert *string `json:"ServerCert,omitempty" xml:"ServerCert,omitempty"`
	// example:
	//
	// -----BEGIN PRIVATE KEY-----MIIE****ihfg==-----END PRIVATE KEY-----
	ServerKey *string `json:"ServerKey,omitempty" xml:"ServerKey,omitempty"`
}

func (s ModifyInstancesSSLShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstancesSSLShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstancesSSLShrinkRequest) GetCAType() *string {
	return s.CAType
}

func (s *ModifyInstancesSSLShrinkRequest) GetInstanceNamesShrink() *string {
	return s.InstanceNamesShrink
}

func (s *ModifyInstancesSSLShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstancesSSLShrinkRequest) GetSSLEnabled() *int32 {
	return s.SSLEnabled
}

func (s *ModifyInstancesSSLShrinkRequest) GetServerCert() *string {
	return s.ServerCert
}

func (s *ModifyInstancesSSLShrinkRequest) GetServerKey() *string {
	return s.ServerKey
}

func (s *ModifyInstancesSSLShrinkRequest) SetCAType(v string) *ModifyInstancesSSLShrinkRequest {
	s.CAType = &v
	return s
}

func (s *ModifyInstancesSSLShrinkRequest) SetInstanceNamesShrink(v string) *ModifyInstancesSSLShrinkRequest {
	s.InstanceNamesShrink = &v
	return s
}

func (s *ModifyInstancesSSLShrinkRequest) SetRegionId(v string) *ModifyInstancesSSLShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstancesSSLShrinkRequest) SetSSLEnabled(v int32) *ModifyInstancesSSLShrinkRequest {
	s.SSLEnabled = &v
	return s
}

func (s *ModifyInstancesSSLShrinkRequest) SetServerCert(v string) *ModifyInstancesSSLShrinkRequest {
	s.ServerCert = &v
	return s
}

func (s *ModifyInstancesSSLShrinkRequest) SetServerKey(v string) *ModifyInstancesSSLShrinkRequest {
	s.ServerKey = &v
	return s
}

func (s *ModifyInstancesSSLShrinkRequest) Validate() error {
	return dara.Validate(s)
}
