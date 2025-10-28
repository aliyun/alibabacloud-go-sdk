// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecureTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetSecureTokenResponseBody
	GetCode() *int32
	SetMessage(v string) *GetSecureTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSecureTokenResponseBody
	GetRequestId() *string
	SetSecureToken(v *GetSecureTokenResponseBodySecureToken) *GetSecureTokenResponseBody
	GetSecureToken() *GetSecureTokenResponseBodySecureToken
}

type GetSecureTokenResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned for the request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D16979DC-4D42-*************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned security token.
	SecureToken *GetSecureTokenResponseBodySecureToken `json:"SecureToken,omitempty" xml:"SecureToken,omitempty" type:"Struct"`
}

func (s GetSecureTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecureTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecureTokenResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetSecureTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSecureTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecureTokenResponseBody) GetSecureToken() *GetSecureTokenResponseBodySecureToken {
	return s.SecureToken
}

func (s *GetSecureTokenResponseBody) SetCode(v int32) *GetSecureTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetSecureTokenResponseBody) SetMessage(v string) *GetSecureTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetSecureTokenResponseBody) SetRequestId(v string) *GetSecureTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecureTokenResponseBody) SetSecureToken(v *GetSecureTokenResponseBodySecureToken) *GetSecureTokenResponseBody {
	s.SecureToken = v
	return s
}

func (s *GetSecureTokenResponseBody) Validate() error {
	if s.SecureToken != nil {
		if err := s.SecureToken.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSecureTokenResponseBodySecureToken struct {
	// The AccessKey ID used in the namespace.
	//
	// example:
	//
	// f676f1**************
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The address of Address Server associated with the namespace.
	//
	// example:
	//
	// addr-****-****.edas.aliyun.com
	AddressServerHost *string `json:"AddressServerHost,omitempty" xml:"AddressServerHost,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-shenzhen
	BelongRegion *string `json:"BelongRegion,omitempty" xml:"BelongRegion,omitempty"`
	// The description of the namespace.
	//
	// example:
	//
	// ”“
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the Alibaba Cloud account that activated Enterprise Distributed Application Service (EDAS).
	//
	// example:
	//
	// 11727458********
	EdasId *string `json:"EdasId,omitempty" xml:"EdasId,omitempty"`
	// The ID of the security token.
	//
	// example:
	//
	// 7279
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the MSE instance.
	//
	// example:
	//
	// mse_prepaid_public_cn-tl32p******
	MseInstanceId *string `json:"MseInstanceId,omitempty" xml:"MseInstanceId,omitempty"`
	// The public endpoint of the MSE registry.
	//
	// example:
	//
	// mse-aa2******-p.nacos-ans.mse.aliyuncs.com
	MseInternetAddress *string `json:"MseInternetAddress,omitempty" xml:"MseInternetAddress,omitempty"`
	// The private endpoint of the MSE registry.
	//
	// example:
	//
	// mse-72******-nacos-ans.mse.aliyuncs.com
	MseIntranetAddress *string `json:"MseIntranetAddress,omitempty" xml:"MseIntranetAddress,omitempty"`
	// The type of the Microservices Engine (MSE) registry.
	//
	// 	- default: the shared registry of EDAS
	//
	// 	- exclusive_mse: MSE Nacos registry
	//
	// example:
	//
	// exclusive_mse
	MseRegistryType *string `json:"MseRegistryType,omitempty" xml:"MseRegistryType,omitempty"`
	// The ID of the region where the namespace resides.
	//
	// example:
	//
	// cn-shenzhen:x*******
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region where the namespace resides.
	//
	// example:
	//
	// x******
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The AccessKey secret used in the namespace.
	//
	// example:
	//
	// gOSgbgR2R*************
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The tenant ID of the namespace.
	//
	// example:
	//
	// 401b7bc8-9441-4693-****-************
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// edas_com***_****@******-*****.***
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetSecureTokenResponseBodySecureToken) String() string {
	return dara.Prettify(s)
}

func (s GetSecureTokenResponseBodySecureToken) GoString() string {
	return s.String()
}

func (s *GetSecureTokenResponseBodySecureToken) GetAccessKey() *string {
	return s.AccessKey
}

func (s *GetSecureTokenResponseBodySecureToken) GetAddressServerHost() *string {
	return s.AddressServerHost
}

func (s *GetSecureTokenResponseBodySecureToken) GetBelongRegion() *string {
	return s.BelongRegion
}

func (s *GetSecureTokenResponseBodySecureToken) GetDescription() *string {
	return s.Description
}

func (s *GetSecureTokenResponseBodySecureToken) GetEdasId() *string {
	return s.EdasId
}

func (s *GetSecureTokenResponseBodySecureToken) GetId() *int64 {
	return s.Id
}

func (s *GetSecureTokenResponseBodySecureToken) GetMseInstanceId() *string {
	return s.MseInstanceId
}

func (s *GetSecureTokenResponseBodySecureToken) GetMseInternetAddress() *string {
	return s.MseInternetAddress
}

func (s *GetSecureTokenResponseBodySecureToken) GetMseIntranetAddress() *string {
	return s.MseIntranetAddress
}

func (s *GetSecureTokenResponseBodySecureToken) GetMseRegistryType() *string {
	return s.MseRegistryType
}

func (s *GetSecureTokenResponseBodySecureToken) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSecureTokenResponseBodySecureToken) GetRegionName() *string {
	return s.RegionName
}

func (s *GetSecureTokenResponseBodySecureToken) GetSecretKey() *string {
	return s.SecretKey
}

func (s *GetSecureTokenResponseBodySecureToken) GetTenantId() *string {
	return s.TenantId
}

func (s *GetSecureTokenResponseBodySecureToken) GetUserId() *string {
	return s.UserId
}

func (s *GetSecureTokenResponseBodySecureToken) SetAccessKey(v string) *GetSecureTokenResponseBodySecureToken {
	s.AccessKey = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetAddressServerHost(v string) *GetSecureTokenResponseBodySecureToken {
	s.AddressServerHost = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetBelongRegion(v string) *GetSecureTokenResponseBodySecureToken {
	s.BelongRegion = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetDescription(v string) *GetSecureTokenResponseBodySecureToken {
	s.Description = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetEdasId(v string) *GetSecureTokenResponseBodySecureToken {
	s.EdasId = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetId(v int64) *GetSecureTokenResponseBodySecureToken {
	s.Id = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetMseInstanceId(v string) *GetSecureTokenResponseBodySecureToken {
	s.MseInstanceId = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetMseInternetAddress(v string) *GetSecureTokenResponseBodySecureToken {
	s.MseInternetAddress = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetMseIntranetAddress(v string) *GetSecureTokenResponseBodySecureToken {
	s.MseIntranetAddress = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetMseRegistryType(v string) *GetSecureTokenResponseBodySecureToken {
	s.MseRegistryType = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetRegionId(v string) *GetSecureTokenResponseBodySecureToken {
	s.RegionId = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetRegionName(v string) *GetSecureTokenResponseBodySecureToken {
	s.RegionName = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetSecretKey(v string) *GetSecureTokenResponseBodySecureToken {
	s.SecretKey = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetTenantId(v string) *GetSecureTokenResponseBodySecureToken {
	s.TenantId = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetUserId(v string) *GetSecureTokenResponseBodySecureToken {
	s.UserId = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) Validate() error {
	return dara.Validate(s)
}
