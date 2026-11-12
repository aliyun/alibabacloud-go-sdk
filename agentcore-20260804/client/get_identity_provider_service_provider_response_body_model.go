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
	// The business result code. If the request fails, handle the error based on the HTTP status code and the message.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The SP information. For the result structure, see the response fields of this operation. The response of an asynchronous write operation indicates that the request is accepted, which does not guarantee that the background task is complete.
	Data *GetIdentityProviderServiceProviderResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The processing status of the current request. This status is not equivalent to the final result of the background task.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The message about request acceptance or errors. Check the asynchronous task result by calling the corresponding query operation.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. You can use the ID to associate the call with troubleshooting. Keep the ID if the request fails.
	//
	// example:
	//
	// request-example
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the current request is successful. A value of true indicates that the request is processed successfully. For asynchronous write operations, you must continue to query the task results.
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
	// The SAML assertion consumer URL of the local SP. Configure it in the external IdP.
	//
	// example:
	//
	// https://sp.example.com/saml/acs
	AcsUrl *string `json:"acsUrl,omitempty" xml:"acsUrl,omitempty"`
	// The entity identifier of the local SP. Configure it in the external IdP. This is not the Entity ID of the external IdP.
	//
	// example:
	//
	// https://sp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// The SP metadata XML of the local SP. Use it for import in the external IdP. Do not submit it as the IdP XML for SAML configuration.
	//
	// example:
	//
	// <EntityDescriptor xmlns="urn:oasis:names:tc:SAML:2.0:metadata" entityID="https://sp.example.com/entity"><SPSSODescriptor protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol"><AssertionConsumerService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://sp.example.com/saml/acs" index="0"/></SPSSODescriptor></EntityDescriptor>
	MetadataXml *string `json:"metadataXml,omitempty" xml:"metadataXml,omitempty"`
	// Required. The query parameter. Only the exact value SAML is supported. The response returns SAML.
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
