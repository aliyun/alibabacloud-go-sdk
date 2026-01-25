// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstancesSSLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCAType(v string) *ModifyInstancesSSLRequest
	GetCAType() *string
	SetInstanceNames(v []*string) *ModifyInstancesSSLRequest
	GetInstanceNames() []*string
	SetRegionId(v string) *ModifyInstancesSSLRequest
	GetRegionId() *string
	SetSSLEnabled(v int32) *ModifyInstancesSSLRequest
	GetSSLEnabled() *int32
	SetServerCert(v string) *ModifyInstancesSSLRequest
	GetServerCert() *string
	SetServerKey(v string) *ModifyInstancesSSLRequest
	GetServerKey() *string
}

type ModifyInstancesSSLRequest struct {
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
	InstanceNames []*string `json:"InstanceNames,omitempty" xml:"InstanceNames,omitempty" type:"Repeated"`
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

func (s ModifyInstancesSSLRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstancesSSLRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstancesSSLRequest) GetCAType() *string {
	return s.CAType
}

func (s *ModifyInstancesSSLRequest) GetInstanceNames() []*string {
	return s.InstanceNames
}

func (s *ModifyInstancesSSLRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstancesSSLRequest) GetSSLEnabled() *int32 {
	return s.SSLEnabled
}

func (s *ModifyInstancesSSLRequest) GetServerCert() *string {
	return s.ServerCert
}

func (s *ModifyInstancesSSLRequest) GetServerKey() *string {
	return s.ServerKey
}

func (s *ModifyInstancesSSLRequest) SetCAType(v string) *ModifyInstancesSSLRequest {
	s.CAType = &v
	return s
}

func (s *ModifyInstancesSSLRequest) SetInstanceNames(v []*string) *ModifyInstancesSSLRequest {
	s.InstanceNames = v
	return s
}

func (s *ModifyInstancesSSLRequest) SetRegionId(v string) *ModifyInstancesSSLRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstancesSSLRequest) SetSSLEnabled(v int32) *ModifyInstancesSSLRequest {
	s.SSLEnabled = &v
	return s
}

func (s *ModifyInstancesSSLRequest) SetServerCert(v string) *ModifyInstancesSSLRequest {
	s.ServerCert = &v
	return s
}

func (s *ModifyInstancesSSLRequest) SetServerKey(v string) *ModifyInstancesSSLRequest {
	s.ServerKey = &v
	return s
}

func (s *ModifyInstancesSSLRequest) Validate() error {
	return dara.Validate(s)
}
