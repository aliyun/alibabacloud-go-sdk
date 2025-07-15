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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
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
	// The ID of the region where the IPsec server is created.
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
