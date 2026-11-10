// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevealServiceCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RevealServiceCredentialResponseBody
	GetCode() *string
	SetData(v *RevealServiceCredentialResponseBodyData) *RevealServiceCredentialResponseBody
	GetData() *RevealServiceCredentialResponseBodyData
	SetMessage(v string) *RevealServiceCredentialResponseBody
	GetMessage() *string
	SetRequestId(v string) *RevealServiceCredentialResponseBody
	GetRequestId() *string
}

type RevealServiceCredentialResponseBody struct {
	// The response code. The value is Ok if the request is successful.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The currently active credential data.
	//
	// example:
	//
	// {"serviceId":"svc-xxxxxxxx","provider":"qwen","credentialType":"apiKey","configRevision":"1","credential":{"apiKeyGenerateMode":"Custom","apiKeys":["sk-xxxxxxxx"]}}
	Data *RevealServiceCredentialResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message. The value is success if the request is successful.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RevealServiceCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevealServiceCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *RevealServiceCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *RevealServiceCredentialResponseBody) GetData() *RevealServiceCredentialResponseBodyData {
	return s.Data
}

func (s *RevealServiceCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RevealServiceCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevealServiceCredentialResponseBody) SetCode(v string) *RevealServiceCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *RevealServiceCredentialResponseBody) SetData(v *RevealServiceCredentialResponseBodyData) *RevealServiceCredentialResponseBody {
	s.Data = v
	return s
}

func (s *RevealServiceCredentialResponseBody) SetMessage(v string) *RevealServiceCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *RevealServiceCredentialResponseBody) SetRequestId(v string) *RevealServiceCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevealServiceCredentialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RevealServiceCredentialResponseBodyData struct {
	// The current configuration revision. The value is legacy for historical objects.
	//
	// example:
	//
	// 1
	ConfigRevision *string `json:"configRevision,omitempty" xml:"configRevision,omitempty"`
	// The credential object of the currently active type. Only the fields corresponding to the credentialType are populated.
	//
	// example:
	//
	// {"apiKeyGenerateMode":"Custom","apiKeys":["sk-xxxxxxxx"]}
	Credential *RevealServiceCredentialResponseBodyDataCredential `json:"credential,omitempty" xml:"credential,omitempty" type:"Struct"`
	// The currently active credential type. Valid values: apiKey, bedrock, vertex, and kling.
	//
	// example:
	//
	// apiKey
	CredentialType *string `json:"credentialType,omitempty" xml:"credentialType,omitempty"`
	// The name of the provider, such as qwen, openai, bedrock, vertex, kling, or pai-eas.
	//
	// example:
	//
	// qwen
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// The unique ID of the service.
	//
	// example:
	//
	// svc-xxxxxxxx
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

func (s RevealServiceCredentialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RevealServiceCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *RevealServiceCredentialResponseBodyData) GetConfigRevision() *string {
	return s.ConfigRevision
}

func (s *RevealServiceCredentialResponseBodyData) GetCredential() *RevealServiceCredentialResponseBodyDataCredential {
	return s.Credential
}

func (s *RevealServiceCredentialResponseBodyData) GetCredentialType() *string {
	return s.CredentialType
}

func (s *RevealServiceCredentialResponseBodyData) GetProvider() *string {
	return s.Provider
}

func (s *RevealServiceCredentialResponseBodyData) GetServiceId() *string {
	return s.ServiceId
}

func (s *RevealServiceCredentialResponseBodyData) SetConfigRevision(v string) *RevealServiceCredentialResponseBodyData {
	s.ConfigRevision = &v
	return s
}

func (s *RevealServiceCredentialResponseBodyData) SetCredential(v *RevealServiceCredentialResponseBodyDataCredential) *RevealServiceCredentialResponseBodyData {
	s.Credential = v
	return s
}

func (s *RevealServiceCredentialResponseBodyData) SetCredentialType(v string) *RevealServiceCredentialResponseBodyData {
	s.CredentialType = &v
	return s
}

func (s *RevealServiceCredentialResponseBodyData) SetProvider(v string) *RevealServiceCredentialResponseBodyData {
	s.Provider = &v
	return s
}

func (s *RevealServiceCredentialResponseBodyData) SetServiceId(v string) *RevealServiceCredentialResponseBodyData {
	s.ServiceId = &v
	return s
}

func (s *RevealServiceCredentialResponseBodyData) Validate() error {
	if s.Credential != nil {
		if err := s.Credential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RevealServiceCredentialResponseBodyDataCredential struct {
	// The API key generation mode. Valid values: Custom and Reference. Historical null values are treated as Custom.
	//
	// example:
	//
	// Custom
	ApiKeyGenerateMode *string `json:"apiKeyGenerateMode,omitempty" xml:"apiKeyGenerateMode,omitempty"`
	// The list of Key Management Service (KMS) reference names. This parameter is returned only in Reference mode. The reference values are not returned.
	//
	// example:
	//
	// ["my-kms-ref"]
	ApiKeyReferences []*string `json:"apiKeyReferences,omitempty" xml:"apiKeyReferences,omitempty" type:"Repeated"`
	// The list of custom API keys. This parameter is returned only when apiKeyGenerateMode is set to Custom. The keys are returned in the original order without deduplication.
	//
	// example:
	//
	// ["sk-xxxxxxxx","sk-yyyyyyyy"]
	ApiKeys []*string `json:"apiKeys,omitempty" xml:"apiKeys,omitempty" type:"Repeated"`
	// The AWS access key ID for Bedrock.
	//
	// example:
	//
	// AKIAxxxxxxxx
	AwsAccessKey *string `json:"awsAccessKey,omitempty" xml:"awsAccessKey,omitempty"`
	// The AWS region for Bedrock.
	//
	// example:
	//
	// us-east-1
	AwsRegion *string `json:"awsRegion,omitempty" xml:"awsRegion,omitempty"`
	// The AWS secret access key for Bedrock.
	//
	// example:
	//
	// wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
	AwsSecretKey *string `json:"awsSecretKey,omitempty" xml:"awsSecretKey,omitempty"`
	// The source of the credential. The value is pai-eas when the credential is backfilled by the PAI-EAS system.
	//
	// example:
	//
	// pai-eas
	CredentialSource *string `json:"credentialSource,omitempty" xml:"credentialSource,omitempty"`
	// The Bedrock endpoint mode. The value is runtime for historical null values.
	//
	// example:
	//
	// runtime
	EndpointMode *string `json:"endpointMode,omitempty" xml:"endpointMode,omitempty"`
	// The access key ID for Kling.
	//
	// example:
	//
	// AKxxxxxxxx
	KlingAccessKey *string `json:"klingAccessKey,omitempty" xml:"klingAccessKey,omitempty"`
	// The secret access key for Kling.
	//
	// example:
	//
	// sk-kling-9f2c1d7e
	KlingSecretKey *string `json:"klingSecretKey,omitempty" xml:"klingSecretKey,omitempty"`
	// The original JSON text of the Vertex service account. The text is neither parsed nor reordered.
	//
	// example:
	//
	// {"type":"service_account","private_key":"********"}
	VertexAuthKey *string `json:"vertexAuthKey,omitempty" xml:"vertexAuthKey,omitempty"`
	// The Vertex project ID.
	//
	// example:
	//
	// my-gcp-project
	VertexProjectId *string `json:"vertexProjectId,omitempty" xml:"vertexProjectId,omitempty"`
	// The region for Vertex.
	//
	// example:
	//
	// us-central1
	VertexRegion *string `json:"vertexRegion,omitempty" xml:"vertexRegion,omitempty"`
}

func (s RevealServiceCredentialResponseBodyDataCredential) String() string {
	return dara.Prettify(s)
}

func (s RevealServiceCredentialResponseBodyDataCredential) GoString() string {
	return s.String()
}

func (s *RevealServiceCredentialResponseBodyDataCredential) GetApiKeyGenerateMode() *string {
	return s.ApiKeyGenerateMode
}

func (s *RevealServiceCredentialResponseBodyDataCredential) GetApiKeyReferences() []*string {
	return s.ApiKeyReferences
}

func (s *RevealServiceCredentialResponseBodyDataCredential) GetApiKeys() []*string {
	return s.ApiKeys
}

func (s *RevealServiceCredentialResponseBodyDataCredential) GetAwsAccessKey() *string {
	return s.AwsAccessKey
}

func (s *RevealServiceCredentialResponseBodyDataCredential) GetAwsRegion() *string {
	return s.AwsRegion
}

func (s *RevealServiceCredentialResponseBodyDataCredential) GetAwsSecretKey() *string {
	return s.AwsSecretKey
}

func (s *RevealServiceCredentialResponseBodyDataCredential) GetCredentialSource() *string {
	return s.CredentialSource
}

func (s *RevealServiceCredentialResponseBodyDataCredential) GetEndpointMode() *string {
	return s.EndpointMode
}

func (s *RevealServiceCredentialResponseBodyDataCredential) GetKlingAccessKey() *string {
	return s.KlingAccessKey
}

func (s *RevealServiceCredentialResponseBodyDataCredential) GetKlingSecretKey() *string {
	return s.KlingSecretKey
}

func (s *RevealServiceCredentialResponseBodyDataCredential) GetVertexAuthKey() *string {
	return s.VertexAuthKey
}

func (s *RevealServiceCredentialResponseBodyDataCredential) GetVertexProjectId() *string {
	return s.VertexProjectId
}

func (s *RevealServiceCredentialResponseBodyDataCredential) GetVertexRegion() *string {
	return s.VertexRegion
}

func (s *RevealServiceCredentialResponseBodyDataCredential) SetApiKeyGenerateMode(v string) *RevealServiceCredentialResponseBodyDataCredential {
	s.ApiKeyGenerateMode = &v
	return s
}

func (s *RevealServiceCredentialResponseBodyDataCredential) SetApiKeyReferences(v []*string) *RevealServiceCredentialResponseBodyDataCredential {
	s.ApiKeyReferences = v
	return s
}

func (s *RevealServiceCredentialResponseBodyDataCredential) SetApiKeys(v []*string) *RevealServiceCredentialResponseBodyDataCredential {
	s.ApiKeys = v
	return s
}

func (s *RevealServiceCredentialResponseBodyDataCredential) SetAwsAccessKey(v string) *RevealServiceCredentialResponseBodyDataCredential {
	s.AwsAccessKey = &v
	return s
}

func (s *RevealServiceCredentialResponseBodyDataCredential) SetAwsRegion(v string) *RevealServiceCredentialResponseBodyDataCredential {
	s.AwsRegion = &v
	return s
}

func (s *RevealServiceCredentialResponseBodyDataCredential) SetAwsSecretKey(v string) *RevealServiceCredentialResponseBodyDataCredential {
	s.AwsSecretKey = &v
	return s
}

func (s *RevealServiceCredentialResponseBodyDataCredential) SetCredentialSource(v string) *RevealServiceCredentialResponseBodyDataCredential {
	s.CredentialSource = &v
	return s
}

func (s *RevealServiceCredentialResponseBodyDataCredential) SetEndpointMode(v string) *RevealServiceCredentialResponseBodyDataCredential {
	s.EndpointMode = &v
	return s
}

func (s *RevealServiceCredentialResponseBodyDataCredential) SetKlingAccessKey(v string) *RevealServiceCredentialResponseBodyDataCredential {
	s.KlingAccessKey = &v
	return s
}

func (s *RevealServiceCredentialResponseBodyDataCredential) SetKlingSecretKey(v string) *RevealServiceCredentialResponseBodyDataCredential {
	s.KlingSecretKey = &v
	return s
}

func (s *RevealServiceCredentialResponseBodyDataCredential) SetVertexAuthKey(v string) *RevealServiceCredentialResponseBodyDataCredential {
	s.VertexAuthKey = &v
	return s
}

func (s *RevealServiceCredentialResponseBodyDataCredential) SetVertexProjectId(v string) *RevealServiceCredentialResponseBodyDataCredential {
	s.VertexProjectId = &v
	return s
}

func (s *RevealServiceCredentialResponseBodyDataCredential) SetVertexRegion(v string) *RevealServiceCredentialResponseBodyDataCredential {
	s.VertexRegion = &v
	return s
}

func (s *RevealServiceCredentialResponseBodyDataCredential) Validate() error {
	return dara.Validate(s)
}
