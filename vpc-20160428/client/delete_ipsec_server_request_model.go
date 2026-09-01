// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpsecServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteIpsecServerRequest
	GetClientToken() *string
	SetDryRun(v string) *DeleteIpsecServerRequest
	GetDryRun() *string
	SetIpsecServerId(v string) *DeleteIpsecServerRequest
	GetIpsecServerId() *string
	SetRegionId(v string) *DeleteIpsecServerRequest
	GetRegionId() *string
}

type DeleteIpsecServerRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without deleting the IPsec server. The system checks the required parameters, request syntax, and business restrictions. If the check fails, the corresponding error message is returned. If the check succeeds, `DryRunOperation` is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. After the check succeeds, the IPsec server is directly deleted.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the IPsec server.
	//
	// This parameter is required.
	//
	// example:
	//
	// iss-bp1jougp8cfsbo8y9****
	IpsecServerId *string `json:"IpsecServerId,omitempty" xml:"IpsecServerId,omitempty"`
	// The region ID of the IPsec server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteIpsecServerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpsecServerRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpsecServerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteIpsecServerRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *DeleteIpsecServerRequest) GetIpsecServerId() *string {
	return s.IpsecServerId
}

func (s *DeleteIpsecServerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteIpsecServerRequest) SetClientToken(v string) *DeleteIpsecServerRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIpsecServerRequest) SetDryRun(v string) *DeleteIpsecServerRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteIpsecServerRequest) SetIpsecServerId(v string) *DeleteIpsecServerRequest {
	s.IpsecServerId = &v
	return s
}

func (s *DeleteIpsecServerRequest) SetRegionId(v string) *DeleteIpsecServerRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpsecServerRequest) Validate() error {
	return dara.Validate(s)
}
