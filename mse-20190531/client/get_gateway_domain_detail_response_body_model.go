// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayDomainDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetGatewayDomainDetailResponseBody
	GetCode() *int32
	SetData(v *GetGatewayDomainDetailResponseBodyData) *GetGatewayDomainDetailResponseBody
	GetData() *GetGatewayDomainDetailResponseBodyData
	SetHttpStatusCode(v int32) *GetGatewayDomainDetailResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetGatewayDomainDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGatewayDomainDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetGatewayDomainDetailResponseBody
	GetSuccess() *bool
}

type GetGatewayDomainDetailResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 403
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetGatewayDomainDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 88B83302-CD88-54D3-8DF2-208BFDC73F0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGatewayDomainDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayDomainDetailResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetGatewayDomainDetailResponseBody) GetData() *GetGatewayDomainDetailResponseBodyData {
	return s.Data
}

func (s *GetGatewayDomainDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetGatewayDomainDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGatewayDomainDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGatewayDomainDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetGatewayDomainDetailResponseBody) SetCode(v int32) *GetGatewayDomainDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBody) SetData(v *GetGatewayDomainDetailResponseBodyData) *GetGatewayDomainDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayDomainDetailResponseBody) SetHttpStatusCode(v int32) *GetGatewayDomainDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBody) SetMessage(v string) *GetGatewayDomainDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBody) SetRequestId(v string) *GetGatewayDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBody) SetSuccess(v bool) *GetGatewayDomainDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayDomainDetailResponseBodyData struct {
	// The start time.
	//
	// example:
	//
	// 2021-04-01 02:35:12
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The algorithm.
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2021-04-01 02:35:12
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The ID of the certificate.
	//
	// example:
	//
	// 234-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// test
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The public domain name.
	//
	// example:
	//
	// name
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The ID of the gateway.
	//
	// example:
	//
	// 1
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-7ea3da97b96543e19f6c597cd4a9****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2021-04-01 02:35:12
	GmtAfter *string `json:"GmtAfter,omitempty" xml:"GmtAfter,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2021-04-01 02:35:12
	GmtBefore *string `json:"GmtBefore,omitempty" xml:"GmtBefore,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-01-07T10:07:57.000+0000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2022-01-07T10:07:57.000+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether `HTTP/2` is enabled.
	//
	// 	- `open`: `HTTP/2` is enabled.
	//
	// 	- `close`: `HTTP/2` is disabled.
	//
	// 	- `globalConfig`: Global configurations are used.
	//
	// example:
	//
	// close
	Http2 *string `json:"Http2,omitempty" xml:"Http2,omitempty"`
	// The ID.
	//
	// example:
	//
	// 12
	Id        *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	IsManaged *bool  `json:"IsManaged,omitempty" xml:"IsManaged,omitempty"`
	// The issuer.
	//
	// example:
	//
	// Istio
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// Indicates whether HTTPS is forcibly used.
	//
	// example:
	//
	// true
	MustHttps *bool `json:"MustHttps,omitempty" xml:"MustHttps,omitempty"`
	// The domain name.
	//
	// example:
	//
	// test.com
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The protocol of the gateway.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The name of the extended field.
	//
	// example:
	//
	// test.com
	Sans                  *string                                                      `json:"Sans,omitempty" xml:"Sans,omitempty"`
	TlsCipherSuitesConfig *GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig `json:"TlsCipherSuitesConfig,omitempty" xml:"TlsCipherSuitesConfig,omitempty" type:"Struct"`
	// The maximum version of Transport Layer Security (TLS).
	//
	// example:
	//
	// TLS 1.3
	TlsMax *string `json:"TlsMax,omitempty" xml:"TlsMax,omitempty"`
	// The minimum version of TLS.
	//
	// example:
	//
	// TLS 1.0
	TlsMin *string `json:"TlsMin,omitempty" xml:"TlsMin,omitempty"`
}

func (s GetGatewayDomainDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayDomainDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayDomainDetailResponseBodyData) GetAfterDate() *int64 {
	return s.AfterDate
}

func (s *GetGatewayDomainDetailResponseBodyData) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *GetGatewayDomainDetailResponseBodyData) GetBeforeDate() *int64 {
	return s.BeforeDate
}

func (s *GetGatewayDomainDetailResponseBodyData) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *GetGatewayDomainDetailResponseBodyData) GetCertName() *string {
	return s.CertName
}

func (s *GetGatewayDomainDetailResponseBodyData) GetCommonName() *string {
	return s.CommonName
}

func (s *GetGatewayDomainDetailResponseBodyData) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *GetGatewayDomainDetailResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetGatewayDomainDetailResponseBodyData) GetGmtAfter() *string {
	return s.GmtAfter
}

func (s *GetGatewayDomainDetailResponseBodyData) GetGmtBefore() *string {
	return s.GmtBefore
}

func (s *GetGatewayDomainDetailResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetGatewayDomainDetailResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetGatewayDomainDetailResponseBodyData) GetHttp2() *string {
	return s.Http2
}

func (s *GetGatewayDomainDetailResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetGatewayDomainDetailResponseBodyData) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *GetGatewayDomainDetailResponseBodyData) GetIssuer() *string {
	return s.Issuer
}

func (s *GetGatewayDomainDetailResponseBodyData) GetMustHttps() *bool {
	return s.MustHttps
}

func (s *GetGatewayDomainDetailResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetGatewayDomainDetailResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *GetGatewayDomainDetailResponseBodyData) GetSans() *string {
	return s.Sans
}

func (s *GetGatewayDomainDetailResponseBodyData) GetTlsCipherSuitesConfig() *GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig {
	return s.TlsCipherSuitesConfig
}

func (s *GetGatewayDomainDetailResponseBodyData) GetTlsMax() *string {
	return s.TlsMax
}

func (s *GetGatewayDomainDetailResponseBodyData) GetTlsMin() *string {
	return s.TlsMin
}

func (s *GetGatewayDomainDetailResponseBodyData) SetAfterDate(v int64) *GetGatewayDomainDetailResponseBodyData {
	s.AfterDate = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetAlgorithm(v string) *GetGatewayDomainDetailResponseBodyData {
	s.Algorithm = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetBeforeDate(v int64) *GetGatewayDomainDetailResponseBodyData {
	s.BeforeDate = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetCertIdentifier(v string) *GetGatewayDomainDetailResponseBodyData {
	s.CertIdentifier = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetCertName(v string) *GetGatewayDomainDetailResponseBodyData {
	s.CertName = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetCommonName(v string) *GetGatewayDomainDetailResponseBodyData {
	s.CommonName = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetGatewayId(v int64) *GetGatewayDomainDetailResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetGatewayUniqueId(v string) *GetGatewayDomainDetailResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetGmtAfter(v string) *GetGatewayDomainDetailResponseBodyData {
	s.GmtAfter = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetGmtBefore(v string) *GetGatewayDomainDetailResponseBodyData {
	s.GmtBefore = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetGmtCreate(v string) *GetGatewayDomainDetailResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetGmtModified(v string) *GetGatewayDomainDetailResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetHttp2(v string) *GetGatewayDomainDetailResponseBodyData {
	s.Http2 = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetId(v int64) *GetGatewayDomainDetailResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetIsManaged(v bool) *GetGatewayDomainDetailResponseBodyData {
	s.IsManaged = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetIssuer(v string) *GetGatewayDomainDetailResponseBodyData {
	s.Issuer = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetMustHttps(v bool) *GetGatewayDomainDetailResponseBodyData {
	s.MustHttps = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetName(v string) *GetGatewayDomainDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetProtocol(v string) *GetGatewayDomainDetailResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetSans(v string) *GetGatewayDomainDetailResponseBodyData {
	s.Sans = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetTlsCipherSuitesConfig(v *GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig) *GetGatewayDomainDetailResponseBodyData {
	s.TlsCipherSuitesConfig = v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetTlsMax(v string) *GetGatewayDomainDetailResponseBodyData {
	s.TlsMax = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetTlsMin(v string) *GetGatewayDomainDetailResponseBodyData {
	s.TlsMin = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) Validate() error {
	if s.TlsCipherSuitesConfig != nil {
		if err := s.TlsCipherSuitesConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig struct {
	ConfigType      *string   `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	TlsCipherSuites []*string `json:"TlsCipherSuites,omitempty" xml:"TlsCipherSuites,omitempty" type:"Repeated"`
}

func (s GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig) GoString() string {
	return s.String()
}

func (s *GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig) GetTlsCipherSuites() []*string {
	return s.TlsCipherSuites
}

func (s *GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig) SetConfigType(v string) *GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig {
	s.ConfigType = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig) SetTlsCipherSuites(v []*string) *GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig {
	s.TlsCipherSuites = v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig) Validate() error {
	return dara.Validate(s)
}
