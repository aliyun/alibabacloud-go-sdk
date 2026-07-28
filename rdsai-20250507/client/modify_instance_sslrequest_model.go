// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceSSLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchName(v string) *ModifyInstanceSSLRequest
	GetBranchName() *string
	SetCAType(v string) *ModifyInstanceSSLRequest
	GetCAType() *string
	SetInstanceName(v string) *ModifyInstanceSSLRequest
	GetInstanceName() *string
	SetRegionId(v string) *ModifyInstanceSSLRequest
	GetRegionId() *string
	SetSSLEnabled(v int32) *ModifyInstanceSSLRequest
	GetSSLEnabled() *int32
	SetServerCert(v string) *ModifyInstanceSSLRequest
	GetServerCert() *string
	SetServerKey(v string) *ModifyInstanceSSLRequest
	GetServerKey() *string
}

type ModifyInstanceSSLRequest struct {
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// The certificate type. Currently, only **custom*	- is supported, which indicates that a custom certificate is used.
	//
	// > This parameter is required when **SSLEnabled*	- is set to **1**.
	//
	// example:
	//
	// custom
	CAType *string `json:"CAType,omitempty" xml:"CAType,omitempty"`
	// The instance ID of the AI application.
	//
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to enable or disable SSL. Valid values:
	//
	// 	- **1**: Enable.
	//
	// 	- **0**: Disable.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SSLEnabled *int32 `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	// The custom certificate content.
	//
	// > This parameter is required when **CAType*	- is set to **custom**.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----MIID*****QqEP-----END CERTIFICATE-----
	ServerCert *string `json:"ServerCert,omitempty" xml:"ServerCert,omitempty"`
	// The private key of the certificate.
	//
	// > This parameter is required when **CAType*	- is set to **custom**.
	//
	// example:
	//
	// -----BEGIN PRIVATE KEY-----MIIE****ihfg==-----END PRIVATE KEY-----
	ServerKey *string `json:"ServerKey,omitempty" xml:"ServerKey,omitempty"`
}

func (s ModifyInstanceSSLRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSSLRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *ModifyInstanceSSLRequest) GetCAType() *string {
	return s.CAType
}

func (s *ModifyInstanceSSLRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceSSLRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceSSLRequest) GetSSLEnabled() *int32 {
	return s.SSLEnabled
}

func (s *ModifyInstanceSSLRequest) GetServerCert() *string {
	return s.ServerCert
}

func (s *ModifyInstanceSSLRequest) GetServerKey() *string {
	return s.ServerKey
}

func (s *ModifyInstanceSSLRequest) SetBranchName(v string) *ModifyInstanceSSLRequest {
	s.BranchName = &v
	return s
}

func (s *ModifyInstanceSSLRequest) SetCAType(v string) *ModifyInstanceSSLRequest {
	s.CAType = &v
	return s
}

func (s *ModifyInstanceSSLRequest) SetInstanceName(v string) *ModifyInstanceSSLRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceSSLRequest) SetRegionId(v string) *ModifyInstanceSSLRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceSSLRequest) SetSSLEnabled(v int32) *ModifyInstanceSSLRequest {
	s.SSLEnabled = &v
	return s
}

func (s *ModifyInstanceSSLRequest) SetServerCert(v string) *ModifyInstanceSSLRequest {
	s.ServerCert = &v
	return s
}

func (s *ModifyInstanceSSLRequest) SetServerKey(v string) *ModifyInstanceSSLRequest {
	s.ServerKey = &v
	return s
}

func (s *ModifyInstanceSSLRequest) Validate() error {
	return dara.Validate(s)
}
