// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListGatewayDomainResponseBody
	GetCode() *int32
	SetData(v []*ListGatewayDomainResponseBodyData) *ListGatewayDomainResponseBody
	GetData() []*ListGatewayDomainResponseBodyData
	SetHttpStatusCode(v int32) *ListGatewayDomainResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListGatewayDomainResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGatewayDomainResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListGatewayDomainResponseBody
	GetSuccess() *bool
}

type ListGatewayDomainResponseBody struct {
	// The return code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The node data.
	Data []*ListGatewayDomainResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3F5D8E93-CA66-57F1-8BCF-A223E11B6B91
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - `true`: The request is successful.
	//
	// - `false`: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGatewayDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayDomainResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayDomainResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListGatewayDomainResponseBody) GetData() []*ListGatewayDomainResponseBodyData {
	return s.Data
}

func (s *ListGatewayDomainResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListGatewayDomainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGatewayDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayDomainResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGatewayDomainResponseBody) SetCode(v int32) *ListGatewayDomainResponseBody {
	s.Code = &v
	return s
}

func (s *ListGatewayDomainResponseBody) SetData(v []*ListGatewayDomainResponseBodyData) *ListGatewayDomainResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayDomainResponseBody) SetHttpStatusCode(v int32) *ListGatewayDomainResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGatewayDomainResponseBody) SetMessage(v string) *ListGatewayDomainResponseBody {
	s.Message = &v
	return s
}

func (s *ListGatewayDomainResponseBody) SetRequestId(v string) *ListGatewayDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayDomainResponseBody) SetSuccess(v bool) *ListGatewayDomainResponseBody {
	s.Success = &v
	return s
}

func (s *ListGatewayDomainResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGatewayDomainResponseBodyData struct {
	// **[Deprecated]*	- The expiration time of the certificate. Go to the domain name details page to view the information.
	//
	// example:
	//
	// 2031-03-30 02:35:12
	CertBeforeDate *string `json:"CertBeforeDate,omitempty" xml:"CertBeforeDate,omitempty"`
	// The certificate ID.
	//
	// example:
	//
	// 3452-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The route remark (Ingress).
	Comment *ListGatewayDomainResponseBodyDataComment `json:"Comment,omitempty" xml:"Comment,omitempty" type:"Struct"`
	// Indicates whether WAF protection is enabled for the domain name.
	EnableWaf *bool `json:"EnableWaf,omitempty" xml:"EnableWaf,omitempty"`
	// The gateway ID.
	//
	// example:
	//
	// 12
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The time when the domain name record was created.
	//
	// example:
	//
	// 2031-03-30 02:35:12
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the domain name record was updated.
	//
	// example:
	//
	// 2031-03-30 02:35:12
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether HTTP/2 is enabled. Valid values:
	//
	// - `open`: HTTP/2 is enabled.
	//
	// - `close`: HTTP/2 is disabled.
	//
	// - `globalConfig`: The global configuration is used.
	//
	// example:
	//
	// close
	Http2 *string `json:"Http2,omitempty" xml:"Http2,omitempty"`
	// The domain name ID.
	//
	// example:
	//
	// 243
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the domain name is managed.
	IsManaged *bool `json:"IsManaged,omitempty" xml:"IsManaged,omitempty"`
	// Indicates whether HTTPS is forced.
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
	// The protocol.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The domain name status. Valid values:
	//
	// - 0: unpublished
	//
	// - 2: publishing
	//
	// - 3: published
	//
	// - 4: editing
	//
	// - 5: taking offline
	//
	// - 6: unavailable
	//
	// example:
	//
	// 3
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the gateway supports WAF protection.
	SupportWaf *bool `json:"SupportWaf,omitempty" xml:"SupportWaf,omitempty"`
	// The TLS cipher suite configuration, including the configuration type and the list of cipher suite names.
	TlsCipherSuitesConfig *ListGatewayDomainResponseBodyDataTlsCipherSuitesConfig `json:"TlsCipherSuitesConfig,omitempty" xml:"TlsCipherSuitesConfig,omitempty" type:"Struct"`
	// The maximum TLS version.
	//
	// example:
	//
	// TLS 1.3
	TlsMax *string `json:"TlsMax,omitempty" xml:"TlsMax,omitempty"`
	// The minimum TLS version.
	//
	// example:
	//
	// TLS 1.0
	TlsMin *string `json:"TlsMin,omitempty" xml:"TlsMin,omitempty"`
	// The source type of the domain name. Valid values:
	//
	// - Op: created in the console
	//
	// - Ingress: created by MSE Ingress
	//
	// example:
	//
	// Op
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListGatewayDomainResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayDomainResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayDomainResponseBodyData) GetCertBeforeDate() *string {
	return s.CertBeforeDate
}

func (s *ListGatewayDomainResponseBodyData) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *ListGatewayDomainResponseBodyData) GetComment() *ListGatewayDomainResponseBodyDataComment {
	return s.Comment
}

func (s *ListGatewayDomainResponseBodyData) GetEnableWaf() *bool {
	return s.EnableWaf
}

func (s *ListGatewayDomainResponseBodyData) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *ListGatewayDomainResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListGatewayDomainResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListGatewayDomainResponseBodyData) GetHttp2() *string {
	return s.Http2
}

func (s *ListGatewayDomainResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListGatewayDomainResponseBodyData) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *ListGatewayDomainResponseBodyData) GetMustHttps() *bool {
	return s.MustHttps
}

func (s *ListGatewayDomainResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListGatewayDomainResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *ListGatewayDomainResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ListGatewayDomainResponseBodyData) GetSupportWaf() *bool {
	return s.SupportWaf
}

func (s *ListGatewayDomainResponseBodyData) GetTlsCipherSuitesConfig() *ListGatewayDomainResponseBodyDataTlsCipherSuitesConfig {
	return s.TlsCipherSuitesConfig
}

func (s *ListGatewayDomainResponseBodyData) GetTlsMax() *string {
	return s.TlsMax
}

func (s *ListGatewayDomainResponseBodyData) GetTlsMin() *string {
	return s.TlsMin
}

func (s *ListGatewayDomainResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListGatewayDomainResponseBodyData) SetCertBeforeDate(v string) *ListGatewayDomainResponseBodyData {
	s.CertBeforeDate = &v
	return s
}

func (s *ListGatewayDomainResponseBodyData) SetCertIdentifier(v string) *ListGatewayDomainResponseBodyData {
	s.CertIdentifier = &v
	return s
}

func (s *ListGatewayDomainResponseBodyData) SetComment(v *ListGatewayDomainResponseBodyDataComment) *ListGatewayDomainResponseBodyData {
	s.Comment = v
	return s
}

func (s *ListGatewayDomainResponseBodyData) SetEnableWaf(v bool) *ListGatewayDomainResponseBodyData {
	s.EnableWaf = &v
	return s
}

func (s *ListGatewayDomainResponseBodyData) SetGatewayId(v int64) *ListGatewayDomainResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayDomainResponseBodyData) SetGmtCreate(v string) *ListGatewayDomainResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListGatewayDomainResponseBodyData) SetGmtModified(v string) *ListGatewayDomainResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListGatewayDomainResponseBodyData) SetHttp2(v string) *ListGatewayDomainResponseBodyData {
	s.Http2 = &v
	return s
}

func (s *ListGatewayDomainResponseBodyData) SetId(v int64) *ListGatewayDomainResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListGatewayDomainResponseBodyData) SetIsManaged(v bool) *ListGatewayDomainResponseBodyData {
	s.IsManaged = &v
	return s
}

func (s *ListGatewayDomainResponseBodyData) SetMustHttps(v bool) *ListGatewayDomainResponseBodyData {
	s.MustHttps = &v
	return s
}

func (s *ListGatewayDomainResponseBodyData) SetName(v string) *ListGatewayDomainResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListGatewayDomainResponseBodyData) SetProtocol(v string) *ListGatewayDomainResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *ListGatewayDomainResponseBodyData) SetStatus(v int32) *ListGatewayDomainResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListGatewayDomainResponseBodyData) SetSupportWaf(v bool) *ListGatewayDomainResponseBodyData {
	s.SupportWaf = &v
	return s
}

func (s *ListGatewayDomainResponseBodyData) SetTlsCipherSuitesConfig(v *ListGatewayDomainResponseBodyDataTlsCipherSuitesConfig) *ListGatewayDomainResponseBodyData {
	s.TlsCipherSuitesConfig = v
	return s
}

func (s *ListGatewayDomainResponseBodyData) SetTlsMax(v string) *ListGatewayDomainResponseBodyData {
	s.TlsMax = &v
	return s
}

func (s *ListGatewayDomainResponseBodyData) SetTlsMin(v string) *ListGatewayDomainResponseBodyData {
	s.TlsMin = &v
	return s
}

func (s *ListGatewayDomainResponseBodyData) SetType(v string) *ListGatewayDomainResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListGatewayDomainResponseBodyData) Validate() error {
	if s.Comment != nil {
		if err := s.Comment.Validate(); err != nil {
			return err
		}
	}
	if s.TlsCipherSuitesConfig != nil {
		if err := s.TlsCipherSuitesConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayDomainResponseBodyDataComment struct {
	// The status.
	//
	// example:
	//
	// Error
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListGatewayDomainResponseBodyDataComment) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayDomainResponseBodyDataComment) GoString() string {
	return s.String()
}

func (s *ListGatewayDomainResponseBodyDataComment) GetStatus() *string {
	return s.Status
}

func (s *ListGatewayDomainResponseBodyDataComment) SetStatus(v string) *ListGatewayDomainResponseBodyDataComment {
	s.Status = &v
	return s
}

func (s *ListGatewayDomainResponseBodyDataComment) Validate() error {
	return dara.Validate(s)
}

type ListGatewayDomainResponseBodyDataTlsCipherSuitesConfig struct {
	// The TLS encrypt cipher suite configuration type. DEFAULT indicates the default configurations, and CUSTOM indicates a custom configuration.
	//
	// example:
	//
	// DEFAULT
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The list of TLS cipher suite names.
	TlsCipherSuites []*string `json:"TlsCipherSuites,omitempty" xml:"TlsCipherSuites,omitempty" type:"Repeated"`
}

func (s ListGatewayDomainResponseBodyDataTlsCipherSuitesConfig) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayDomainResponseBodyDataTlsCipherSuitesConfig) GoString() string {
	return s.String()
}

func (s *ListGatewayDomainResponseBodyDataTlsCipherSuitesConfig) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListGatewayDomainResponseBodyDataTlsCipherSuitesConfig) GetTlsCipherSuites() []*string {
	return s.TlsCipherSuites
}

func (s *ListGatewayDomainResponseBodyDataTlsCipherSuitesConfig) SetConfigType(v string) *ListGatewayDomainResponseBodyDataTlsCipherSuitesConfig {
	s.ConfigType = &v
	return s
}

func (s *ListGatewayDomainResponseBodyDataTlsCipherSuitesConfig) SetTlsCipherSuites(v []*string) *ListGatewayDomainResponseBodyDataTlsCipherSuitesConfig {
	s.TlsCipherSuites = v
	return s
}

func (s *ListGatewayDomainResponseBodyDataTlsCipherSuitesConfig) Validate() error {
	return dara.Validate(s)
}
