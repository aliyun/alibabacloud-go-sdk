// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpnGatewayDiagnoseResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetVpnGatewayDiagnoseResultRequest
	GetClientToken() *string
	SetDiagnoseId(v string) *GetVpnGatewayDiagnoseResultRequest
	GetDiagnoseId() *string
	SetRegionId(v string) *GetVpnGatewayDiagnoseResultRequest
	GetRegionId() *string
	SetVpnGatewayId(v string) *GetVpnGatewayDiagnoseResultRequest
	GetVpnGatewayId() *string
}

type GetVpnGatewayDiagnoseResultRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the client token. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The diagnostic ID.
	//
	// The diagnostic ID is returned when you call the [DiagnoseVpnGateway](https://help.aliyun.com/document_detail/469751.html) operation.
	//
	// example:
	//
	// vpndgn-uf6kuxbe3iv028k3s****
	DiagnoseId *string `json:"DiagnoseId,omitempty" xml:"DiagnoseId,omitempty"`
	// The region ID of the VPN gateway instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The VPN gateway instance ID.
	//
	// example:
	//
	// vpn-uf6fzwp0ck3frwtbk****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s GetVpnGatewayDiagnoseResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVpnGatewayDiagnoseResultRequest) GoString() string {
	return s.String()
}

func (s *GetVpnGatewayDiagnoseResultRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetVpnGatewayDiagnoseResultRequest) GetDiagnoseId() *string {
	return s.DiagnoseId
}

func (s *GetVpnGatewayDiagnoseResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVpnGatewayDiagnoseResultRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *GetVpnGatewayDiagnoseResultRequest) SetClientToken(v string) *GetVpnGatewayDiagnoseResultRequest {
	s.ClientToken = &v
	return s
}

func (s *GetVpnGatewayDiagnoseResultRequest) SetDiagnoseId(v string) *GetVpnGatewayDiagnoseResultRequest {
	s.DiagnoseId = &v
	return s
}

func (s *GetVpnGatewayDiagnoseResultRequest) SetRegionId(v string) *GetVpnGatewayDiagnoseResultRequest {
	s.RegionId = &v
	return s
}

func (s *GetVpnGatewayDiagnoseResultRequest) SetVpnGatewayId(v string) *GetVpnGatewayDiagnoseResultRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *GetVpnGatewayDiagnoseResultRequest) Validate() error {
	return dara.Validate(s)
}
