// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderServiceProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetIdentityProviderServiceProviderResponseBody
	GetCode() *string
	SetData(v *GetIdentityProviderServiceProviderResponseBodyData) *GetIdentityProviderServiceProviderResponseBody
	GetData() *GetIdentityProviderServiceProviderResponseBodyData
	SetHttpStatusCode(v int32) *GetIdentityProviderServiceProviderResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetIdentityProviderServiceProviderResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetIdentityProviderServiceProviderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetIdentityProviderServiceProviderResponseBody
	GetSuccess() *bool
}

type GetIdentityProviderServiceProviderResponseBody struct {
	// 失败时结合 HTTP 状态及 message 处理。
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 结果结构见本接口返回字段；异步写接口的返回表示受理，不保证后台任务已完成。
	Data *GetIdentityProviderServiceProviderResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 本次请求的处理状态，不等同后台任务最终结果。
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// 请求受理或错误消息；异步任务结果需按对应查询接口核对。
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 用于关联本次调用与排障，失败时请保留。
	//
	// example:
	//
	// request-example
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 表示本次请求处理成功；异步写入仍需继续查询任务结果。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetIdentityProviderServiceProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderServiceProviderResponseBody) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderServiceProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetIdentityProviderServiceProviderResponseBody) GetData() *GetIdentityProviderServiceProviderResponseBodyData {
	return s.Data
}

func (s *GetIdentityProviderServiceProviderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetIdentityProviderServiceProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetIdentityProviderServiceProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIdentityProviderServiceProviderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetIdentityProviderServiceProviderResponseBody) SetCode(v string) *GetIdentityProviderServiceProviderResponseBody {
	s.Code = &v
	return s
}

func (s *GetIdentityProviderServiceProviderResponseBody) SetData(v *GetIdentityProviderServiceProviderResponseBodyData) *GetIdentityProviderServiceProviderResponseBody {
	s.Data = v
	return s
}

func (s *GetIdentityProviderServiceProviderResponseBody) SetHttpStatusCode(v int32) *GetIdentityProviderServiceProviderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetIdentityProviderServiceProviderResponseBody) SetMessage(v string) *GetIdentityProviderServiceProviderResponseBody {
	s.Message = &v
	return s
}

func (s *GetIdentityProviderServiceProviderResponseBody) SetRequestId(v string) *GetIdentityProviderServiceProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIdentityProviderServiceProviderResponseBody) SetSuccess(v bool) *GetIdentityProviderServiceProviderResponseBody {
	s.Success = &v
	return s
}

func (s *GetIdentityProviderServiceProviderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentityProviderServiceProviderResponseBodyData struct {
	// 我方 SAML 断言消费地址，配置到外部 IdP。
	//
	// example:
	//
	// https://sp.example.com/saml/acs
	AcsUrl *string `json:"acsUrl,omitempty" xml:"acsUrl,omitempty"`
	// 我方 SP 的实体标识，配置到外部 IdP；不是外部 IdP Entity ID。
	//
	// example:
	//
	// https://sp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// 我方 SP Metadata XML，供外部 IdP 导入；不能作为 IdP XML 提交 SAML 配置。
	//
	// example:
	//
	// <EntityDescriptor xmlns="urn:oasis:names:tc:SAML:2.0:metadata" entityID="https://sp.example.com/entity"><SPSSODescriptor protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol"><AssertionConsumerService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://sp.example.com/saml/acs" index="0"/></SPSSODescriptor></EntityDescriptor>
	MetadataXml *string `json:"metadataXml,omitempty" xml:"metadataXml,omitempty"`
	// 查询参数必填，仅支持精确值 SAML；响应返回 SAML。
	//
	// example:
	//
	// SAML
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s GetIdentityProviderServiceProviderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderServiceProviderResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderServiceProviderResponseBodyData) GetAcsUrl() *string {
	return s.AcsUrl
}

func (s *GetIdentityProviderServiceProviderResponseBodyData) GetEntityId() *string {
	return s.EntityId
}

func (s *GetIdentityProviderServiceProviderResponseBodyData) GetMetadataXml() *string {
	return s.MetadataXml
}

func (s *GetIdentityProviderServiceProviderResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *GetIdentityProviderServiceProviderResponseBodyData) SetAcsUrl(v string) *GetIdentityProviderServiceProviderResponseBodyData {
	s.AcsUrl = &v
	return s
}

func (s *GetIdentityProviderServiceProviderResponseBodyData) SetEntityId(v string) *GetIdentityProviderServiceProviderResponseBodyData {
	s.EntityId = &v
	return s
}

func (s *GetIdentityProviderServiceProviderResponseBodyData) SetMetadataXml(v string) *GetIdentityProviderServiceProviderResponseBodyData {
	s.MetadataXml = &v
	return s
}

func (s *GetIdentityProviderServiceProviderResponseBodyData) SetProtocol(v string) *GetIdentityProviderServiceProviderResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *GetIdentityProviderServiceProviderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
