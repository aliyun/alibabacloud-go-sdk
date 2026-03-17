// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiagnoseSmartAccessGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DiagnoseSmartAccessGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DiagnoseSmartAccessGatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DiagnoseSmartAccessGatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DiagnoseSmartAccessGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DiagnoseSmartAccessGatewayRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DiagnoseSmartAccessGatewayRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DiagnoseSmartAccessGatewayRequest
	GetSmartAGSn() *string
}

type DiagnoseSmartAccessGatewayRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-1um5x5nwhilymw****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The serial number (SN) of the SAG device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sage62x022502****
	SmartAGSn *string `json:"SmartAGSn,omitempty" xml:"SmartAGSn,omitempty"`
}

func (s DiagnoseSmartAccessGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DiagnoseSmartAccessGatewayRequest) GoString() string {
	return s.String()
}

func (s *DiagnoseSmartAccessGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DiagnoseSmartAccessGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DiagnoseSmartAccessGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DiagnoseSmartAccessGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DiagnoseSmartAccessGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DiagnoseSmartAccessGatewayRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DiagnoseSmartAccessGatewayRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DiagnoseSmartAccessGatewayRequest) SetOwnerAccount(v string) *DiagnoseSmartAccessGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DiagnoseSmartAccessGatewayRequest) SetOwnerId(v int64) *DiagnoseSmartAccessGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *DiagnoseSmartAccessGatewayRequest) SetRegionId(v string) *DiagnoseSmartAccessGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *DiagnoseSmartAccessGatewayRequest) SetResourceOwnerAccount(v string) *DiagnoseSmartAccessGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DiagnoseSmartAccessGatewayRequest) SetResourceOwnerId(v int64) *DiagnoseSmartAccessGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DiagnoseSmartAccessGatewayRequest) SetSmartAGId(v string) *DiagnoseSmartAccessGatewayRequest {
	s.SmartAGId = &v
	return s
}

func (s *DiagnoseSmartAccessGatewayRequest) SetSmartAGSn(v string) *DiagnoseSmartAccessGatewayRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DiagnoseSmartAccessGatewayRequest) Validate() error {
	return dara.Validate(s)
}
