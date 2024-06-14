// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-pop/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AsymmetricDecryptRequest struct {
	// The decryption algorithm.
	//
	// This parameter is required.
	//
	// example:
	//
	// RSAES_OAEP_SHA_1
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ciphertext that you want to decrypt.
	//
	// > 	- The value is encoded in Base64.
	//
	// > 	- You can call the [AsymmetricEncrypt](https://help.aliyun.com/document_detail/148131.html) operation to generate the ciphertext.
	//
	// This parameter is required.
	//
	// example:
	//
	// BQKP+1zK6+ZEMxTP5qaVzcsgXtWplYBKm0NXdSnB5FzliFxE1bSiu4dnEIlca2JpeH7yz1/S6fed630H+hIH6DoM25fTLNcKj+mFB0Xnh9m2+HN59Mn4qyTfcUeadnfCXSWcGBouhXFwcdd2rJ3n337bzTf4jm659gZu3L0i6PLuxM9p7mqdwO0cKJPfGVfhnfMz+f4alMg79WB/NNyE2lyX7/qxvV49ObNrrJbKSFiz8Djocaf0IESNLMbfYI5bXjWkJlX92DQbKhibtQW8ZOJ//ZC6t0AWcUoKL6QDm/dg5koQalcleRinpB+QadFm894sLbVZ9+N4GVsv1W****==
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The ID of the customer master key (CMK). The ID must be globally unique.
	//
	// >  You can also set this parameter to an alias that is bound to the CMK. For more information, see [Alias overview](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The version ID of the CMK. The ID must be globally unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s AsymmetricDecryptRequest) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricDecryptRequest) GoString() string {
	return s.String()
}

func (s *AsymmetricDecryptRequest) SetAlgorithm(v string) *AsymmetricDecryptRequest {
	s.Algorithm = &v
	return s
}

func (s *AsymmetricDecryptRequest) SetCiphertextBlob(v string) *AsymmetricDecryptRequest {
	s.CiphertextBlob = &v
	return s
}

func (s *AsymmetricDecryptRequest) SetKeyId(v string) *AsymmetricDecryptRequest {
	s.KeyId = &v
	return s
}

func (s *AsymmetricDecryptRequest) SetKeyVersionId(v string) *AsymmetricDecryptRequest {
	s.KeyVersionId = &v
	return s
}

type AsymmetricDecryptResponseBody struct {
	// The ID of the CMK. The ID must be globally unique.
	//
	// >  If you set the KeyId parameter in the request to an alias, the ID of the CMK to which the alias is bound is returned.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The version ID of the CMK that is used to encrypt the plaintext.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The Base64-encoded plaintext that is generated after decryption.
	//
	// example:
	//
	// SGVsbG8gd29ybGQ=
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 475f1620-b9d3-4d35-b5c6-3fbdd941423d
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AsymmetricDecryptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricDecryptResponseBody) GoString() string {
	return s.String()
}

func (s *AsymmetricDecryptResponseBody) SetKeyId(v string) *AsymmetricDecryptResponseBody {
	s.KeyId = &v
	return s
}

func (s *AsymmetricDecryptResponseBody) SetKeyVersionId(v string) *AsymmetricDecryptResponseBody {
	s.KeyVersionId = &v
	return s
}

func (s *AsymmetricDecryptResponseBody) SetPlaintext(v string) *AsymmetricDecryptResponseBody {
	s.Plaintext = &v
	return s
}

func (s *AsymmetricDecryptResponseBody) SetRequestId(v string) *AsymmetricDecryptResponseBody {
	s.RequestId = &v
	return s
}

type AsymmetricDecryptResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsymmetricDecryptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AsymmetricDecryptResponse) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricDecryptResponse) GoString() string {
	return s.String()
}

func (s *AsymmetricDecryptResponse) SetHeaders(v map[string]*string) *AsymmetricDecryptResponse {
	s.Headers = v
	return s
}

func (s *AsymmetricDecryptResponse) SetStatusCode(v int32) *AsymmetricDecryptResponse {
	s.StatusCode = &v
	return s
}

func (s *AsymmetricDecryptResponse) SetBody(v *AsymmetricDecryptResponseBody) *AsymmetricDecryptResponse {
	s.Body = v
	return s
}

type AsymmetricEncryptRequest struct {
	// The encryption algorithm.
	//
	// This parameter is required.
	//
	// example:
	//
	// RSAES_OAEP_SHA_1
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// >  You can also set this parameter to an alias that is bound to the CMK. For more information, see [Overview of aliases](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The version ID of the CMK. The ID must be globally unique.
	//
	// >  You can call the [ListKeyVersions](https://help.aliyun.com/document_detail/133966.html) operation to query the versions of a CMK. The ID of a version is specified by the KeyVersionId parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The plaintext that you want to encrypt. The plaintext must be Base64-encoded.
	//
	// This parameter is required.
	//
	// example:
	//
	// SGVsbG8gd29ybGQ=
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
}

func (s AsymmetricEncryptRequest) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricEncryptRequest) GoString() string {
	return s.String()
}

func (s *AsymmetricEncryptRequest) SetAlgorithm(v string) *AsymmetricEncryptRequest {
	s.Algorithm = &v
	return s
}

func (s *AsymmetricEncryptRequest) SetKeyId(v string) *AsymmetricEncryptRequest {
	s.KeyId = &v
	return s
}

func (s *AsymmetricEncryptRequest) SetKeyVersionId(v string) *AsymmetricEncryptRequest {
	s.KeyVersionId = &v
	return s
}

func (s *AsymmetricEncryptRequest) SetPlaintext(v string) *AsymmetricEncryptRequest {
	s.Plaintext = &v
	return s
}

type AsymmetricEncryptResponseBody struct {
	// The Base64-encoded ciphertext that was generated after encryption.
	//
	// example:
	//
	// BQKP+1zK6+ZEMxTP5qaVzcsgXtWplYBKm0NXdSnB5FzliFxE1bSiu4dnEIlca2JpeH7yz1/S6fed630H+hIH6DoM25fTLNcKj+mFB0Xnh9m2+HN59Mn4qyTfcUeadnfCXSWcGBouhXFwcdd2rJ3n337bzTf4jm659gZu3L0i6PLuxM9p7mqdwO0cKJPfGVfhnfMz+f4alMg79WB/NNyE2lyX7/qxvV49ObNrrJbKSFiz8Djocaf0IESNLMbfYI5bXjWkJlX92DQbKhibtQW8ZOJ//ZC6t0AWcUoKL6QDm/dg5koQalcleRinpB+QadFm894sLbVZ9+N4GVsv1Wbjwg==
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// >  If you set the KeyId parameter in the request to an alias, the ID of the CMK to which the alias is bound is returned.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The version ID of the CMK that is used to encrypt the plaintext.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 475f1620-b9d3-4d35-b5c6-3fbdd941423d
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AsymmetricEncryptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricEncryptResponseBody) GoString() string {
	return s.String()
}

func (s *AsymmetricEncryptResponseBody) SetCiphertextBlob(v string) *AsymmetricEncryptResponseBody {
	s.CiphertextBlob = &v
	return s
}

func (s *AsymmetricEncryptResponseBody) SetKeyId(v string) *AsymmetricEncryptResponseBody {
	s.KeyId = &v
	return s
}

func (s *AsymmetricEncryptResponseBody) SetKeyVersionId(v string) *AsymmetricEncryptResponseBody {
	s.KeyVersionId = &v
	return s
}

func (s *AsymmetricEncryptResponseBody) SetRequestId(v string) *AsymmetricEncryptResponseBody {
	s.RequestId = &v
	return s
}

type AsymmetricEncryptResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsymmetricEncryptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AsymmetricEncryptResponse) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricEncryptResponse) GoString() string {
	return s.String()
}

func (s *AsymmetricEncryptResponse) SetHeaders(v map[string]*string) *AsymmetricEncryptResponse {
	s.Headers = v
	return s
}

func (s *AsymmetricEncryptResponse) SetStatusCode(v int32) *AsymmetricEncryptResponse {
	s.StatusCode = &v
	return s
}

func (s *AsymmetricEncryptResponse) SetBody(v *AsymmetricEncryptResponseBody) *AsymmetricEncryptResponse {
	s.Body = v
	return s
}

type AsymmetricSignRequest struct {
	// The version ID of the CMK. The ID must be globally unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// RSA_PSS_SHA_256
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The signature algorithm.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZOyIygCyaOW6GjVnihtTFtIS9PNmskdyMlNKiu****=
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The operation that you want to perform. Set the value to **AsymmetricSign**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the customer master key (CMK). The ID must be globally unique.
	//
	// >  You can also set this parameter to an alias that is bound to the CMK. For more information, see [Alias overview](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s AsymmetricSignRequest) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricSignRequest) GoString() string {
	return s.String()
}

func (s *AsymmetricSignRequest) SetAlgorithm(v string) *AsymmetricSignRequest {
	s.Algorithm = &v
	return s
}

func (s *AsymmetricSignRequest) SetDigest(v string) *AsymmetricSignRequest {
	s.Digest = &v
	return s
}

func (s *AsymmetricSignRequest) SetKeyId(v string) *AsymmetricSignRequest {
	s.KeyId = &v
	return s
}

func (s *AsymmetricSignRequest) SetKeyVersionId(v string) *AsymmetricSignRequest {
	s.KeyVersionId = &v
	return s
}

type AsymmetricSignResponseBody struct {
	// The version ID of the CMK. The ID must be globally unique.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The digest that is generated for the original message by using a hash algorithm. The hash algorithm is specified by the Algorithm parameter.
	//
	// > 	- The value is encoded in Base64.
	//
	// > 	- For more information about how to calculate message digests, see the **Preprocess signature: compute a message digest*	- section of the [Generate and verify a signature by using an asymmetric CMK](https://help.aliyun.com/document_detail/148146.html) topic.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The calculated signature.
	//
	// >  The value is encoded in Base64.
	//
	// example:
	//
	// 475f1620-b9d3-4d35-b5c6-3fbdd941423d
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// >  If you set the KeyId parameter in the request to an alias, the ID of the CMK to which the alias is bound is returned.
	//
	// example:
	//
	// M2CceNZH00ZgL9ED/ZHFp21YRAvYeZHknJUc207OCZ0N9wNn9As4z2bON3FF3je+1Nu+2+/8Zj50HpMTpzYpMp2R93cYmACCmhaYoKydxylbyGzJR8y9likZRCrkD38lRoS40aBBvv/6iRKzQuo9EGYVcel36cMNg00VmYNBy3pa1rwg3gA4l3cy6kjayZja1WGPkVhrVKsrJMdbpl0ApLjXKuD8rw1n1XLCwCUEL5eLPljTZaAveqdOFQOiZnZEGI27qIiZe7I1fN8tcz6anS/gTM7xRKE++5egEvRWlTQQTJeApnPSiUPA+8ZykNdelQsOQh5SrGoyI4A5pq****==
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AsymmetricSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricSignResponseBody) GoString() string {
	return s.String()
}

func (s *AsymmetricSignResponseBody) SetKeyId(v string) *AsymmetricSignResponseBody {
	s.KeyId = &v
	return s
}

func (s *AsymmetricSignResponseBody) SetKeyVersionId(v string) *AsymmetricSignResponseBody {
	s.KeyVersionId = &v
	return s
}

func (s *AsymmetricSignResponseBody) SetRequestId(v string) *AsymmetricSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *AsymmetricSignResponseBody) SetValue(v string) *AsymmetricSignResponseBody {
	s.Value = &v
	return s
}

type AsymmetricSignResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsymmetricSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AsymmetricSignResponse) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricSignResponse) GoString() string {
	return s.String()
}

func (s *AsymmetricSignResponse) SetHeaders(v map[string]*string) *AsymmetricSignResponse {
	s.Headers = v
	return s
}

func (s *AsymmetricSignResponse) SetStatusCode(v int32) *AsymmetricSignResponse {
	s.StatusCode = &v
	return s
}

func (s *AsymmetricSignResponse) SetBody(v *AsymmetricSignResponseBody) *AsymmetricSignResponse {
	s.Body = v
	return s
}

type AsymmetricVerifyRequest struct {
	// The signature algorithm.
	//
	// This parameter is required.
	//
	// example:
	//
	// RSA_PSS_SHA_256
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The digest that is generated for the original message by using a hash algorithm. The hash algorithm is specified by the **Algorithm*	- parameter.
	//
	// >  The value is encoded in Base64.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZOyIygCyaOW6GjVnihtTFtIS9PNmskdyMlNKiuy****=
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// >  You can also set this parameter to an alias that is bound to the CMK. For more information, see [Overview of aliases](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The version ID of the CMK. The ID must be globally unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The signature value to be verified.
	//
	// >  The value is encoded in Base64.
	//
	// This parameter is required.
	//
	// example:
	//
	// M2CceNZH00ZgL9ED/ZHFp21YRAvYeZHknJUc207OCZ0N9wNn9As4z2bON3FF3je+1Nu+2+/8Zj50HpMTpzYpMp2R93cYmACCmhaYoKydxylbyGzJR8y9likZRCrkD38lRoS40aBBvv/6iRKzQuo9EGYVcel36cMNg00VmYNBy3pa1rwg3gA4l3cy6kjayZja1WGPkVhrVKsrJMdbpl0ApLjXKuD8rw1n1XLCwCUEL5eLPljTZaAveqdOFQOiZnZEGI27qIiZe7I1fN8tcz6anS/gTM7xRKE++5egEvRWlTQQTJeApnPSiUPA+8ZykNdelQsOQh5SrGoyI4A5pq****==
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AsymmetricVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricVerifyRequest) GoString() string {
	return s.String()
}

func (s *AsymmetricVerifyRequest) SetAlgorithm(v string) *AsymmetricVerifyRequest {
	s.Algorithm = &v
	return s
}

func (s *AsymmetricVerifyRequest) SetDigest(v string) *AsymmetricVerifyRequest {
	s.Digest = &v
	return s
}

func (s *AsymmetricVerifyRequest) SetKeyId(v string) *AsymmetricVerifyRequest {
	s.KeyId = &v
	return s
}

func (s *AsymmetricVerifyRequest) SetKeyVersionId(v string) *AsymmetricVerifyRequest {
	s.KeyVersionId = &v
	return s
}

func (s *AsymmetricVerifyRequest) SetValue(v string) *AsymmetricVerifyRequest {
	s.Value = &v
	return s
}

type AsymmetricVerifyResponseBody struct {
	// The ID of the CMK. The ID must be globally unique.
	//
	// >  If you set the KeyId parameter in the request to an alias, the ID of the CMK to which the alias is bound is returned.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The version ID of the CMK that is used to encrypt the plaintext.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 475f1620-b9d3-4d35-b5c6-3fbdd941423d
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the signature passed the verification.
	//
	// example:
	//
	// true
	Value *bool `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AsymmetricVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *AsymmetricVerifyResponseBody) SetKeyId(v string) *AsymmetricVerifyResponseBody {
	s.KeyId = &v
	return s
}

func (s *AsymmetricVerifyResponseBody) SetKeyVersionId(v string) *AsymmetricVerifyResponseBody {
	s.KeyVersionId = &v
	return s
}

func (s *AsymmetricVerifyResponseBody) SetRequestId(v string) *AsymmetricVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AsymmetricVerifyResponseBody) SetValue(v bool) *AsymmetricVerifyResponseBody {
	s.Value = &v
	return s
}

type AsymmetricVerifyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsymmetricVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AsymmetricVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricVerifyResponse) GoString() string {
	return s.String()
}

func (s *AsymmetricVerifyResponse) SetHeaders(v map[string]*string) *AsymmetricVerifyResponse {
	s.Headers = v
	return s
}

func (s *AsymmetricVerifyResponse) SetStatusCode(v int32) *AsymmetricVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *AsymmetricVerifyResponse) SetBody(v *AsymmetricVerifyResponseBody) *AsymmetricVerifyResponse {
	s.Body = v
	return s
}

type CancelKeyDeletionRequest struct {
	// The ID of the CMK. The ID must be globally unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s CancelKeyDeletionRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelKeyDeletionRequest) GoString() string {
	return s.String()
}

func (s *CancelKeyDeletionRequest) SetKeyId(v string) *CancelKeyDeletionRequest {
	s.KeyId = &v
	return s
}

type CancelKeyDeletionResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3da5b8cc-8107-40ac-a170-793cd181d7b7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelKeyDeletionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelKeyDeletionResponseBody) GoString() string {
	return s.String()
}

func (s *CancelKeyDeletionResponseBody) SetRequestId(v string) *CancelKeyDeletionResponseBody {
	s.RequestId = &v
	return s
}

type CancelKeyDeletionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelKeyDeletionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelKeyDeletionResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelKeyDeletionResponse) GoString() string {
	return s.String()
}

func (s *CancelKeyDeletionResponse) SetHeaders(v map[string]*string) *CancelKeyDeletionResponse {
	s.Headers = v
	return s
}

func (s *CancelKeyDeletionResponse) SetStatusCode(v int32) *CancelKeyDeletionResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelKeyDeletionResponse) SetBody(v *CancelKeyDeletionResponseBody) *CancelKeyDeletionResponse {
	s.Body = v
	return s
}

type CertificatePrivateKeyDecryptRequest struct {
	// The encryption algorithm. Valid values:
	//
	// 	- RSAES_OAEP_SHA_1
	//
	// 	- RSAES_OAEP_SHA_256
	//
	// 	- SM2PKE
	//
	// > The SM2PKE encryption algorithm is supported only in regions in mainland China. In these regions, managed hardware security modules (HSMs) are used. For more information, see [Managed HSM overview](https://help.aliyun.com/document_detail/125803.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// RSAES_OAEP_SHA_256
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ID of the certificate. The ID must be globally unique in Certificates Manager.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The data that you want to decrypt.
	//
	// The value is encoded in Base64.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZOyIygCyaOW6Gj****MlNKiuyjfzw=
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
}

func (s CertificatePrivateKeyDecryptRequest) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrivateKeyDecryptRequest) GoString() string {
	return s.String()
}

func (s *CertificatePrivateKeyDecryptRequest) SetAlgorithm(v string) *CertificatePrivateKeyDecryptRequest {
	s.Algorithm = &v
	return s
}

func (s *CertificatePrivateKeyDecryptRequest) SetCertificateId(v string) *CertificatePrivateKeyDecryptRequest {
	s.CertificateId = &v
	return s
}

func (s *CertificatePrivateKeyDecryptRequest) SetCiphertextBlob(v string) *CertificatePrivateKeyDecryptRequest {
	s.CiphertextBlob = &v
	return s
}

type CertificatePrivateKeyDecryptResponseBody struct {
	// The ID of the certificate.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The plaintext after data is decrypted.
	//
	// The value is encoded in Base64.
	//
	// example:
	//
	// VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5979d897-d69f-4fc9-87dd-f3bb73c40b80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CertificatePrivateKeyDecryptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrivateKeyDecryptResponseBody) GoString() string {
	return s.String()
}

func (s *CertificatePrivateKeyDecryptResponseBody) SetCertificateId(v string) *CertificatePrivateKeyDecryptResponseBody {
	s.CertificateId = &v
	return s
}

func (s *CertificatePrivateKeyDecryptResponseBody) SetPlaintext(v string) *CertificatePrivateKeyDecryptResponseBody {
	s.Plaintext = &v
	return s
}

func (s *CertificatePrivateKeyDecryptResponseBody) SetRequestId(v string) *CertificatePrivateKeyDecryptResponseBody {
	s.RequestId = &v
	return s
}

type CertificatePrivateKeyDecryptResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CertificatePrivateKeyDecryptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CertificatePrivateKeyDecryptResponse) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrivateKeyDecryptResponse) GoString() string {
	return s.String()
}

func (s *CertificatePrivateKeyDecryptResponse) SetHeaders(v map[string]*string) *CertificatePrivateKeyDecryptResponse {
	s.Headers = v
	return s
}

func (s *CertificatePrivateKeyDecryptResponse) SetStatusCode(v int32) *CertificatePrivateKeyDecryptResponse {
	s.StatusCode = &v
	return s
}

func (s *CertificatePrivateKeyDecryptResponse) SetBody(v *CertificatePrivateKeyDecryptResponseBody) *CertificatePrivateKeyDecryptResponse {
	s.Body = v
	return s
}

type CertificatePrivateKeySignRequest struct {
	// The signature algorithm. Valid values:
	//
	// 	- RSA_PKCS1_SHA_256
	//
	// 	- RSA_PSS_SHA_256
	//
	// 	- ECDSA_SHA_256
	//
	// 	- SM2DSA
	//
	// >	- The SM2DSA signature algorithm is supported only in regions where managed hardware security modules (HSMs) are used in mainland China. For more information, see [Managed HSM overview](https://help.aliyun.com/document_detail/125803.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ECDSA_SHA_256
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ID of the certificate. The ID must be globally unique in Certificates Manager.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The data to be signed.
	//
	// The value is encoded in Base64. For example, if the hexadecimal data that you want to sign is `[0x31, 0x32, 0x33, 0x34]`, the Base64-encoded data is `MTIzNA==`.
	//
	// If the MessageType parameter is set to RAW, the size of the data must be less than or equal to 4 KB.
	//
	// If the size of the data is greater than 4 KB, you can set the MessageType parameter to DIGEST and set the Message parameter to the digest of the data. The digest is also called hash value. You can compute the digest of the data on an on-premises machine. Certificates Manager uses the digest that you compute in your own certificate application system. The message digest algorithm that you use must match the specified signature algorithm. Comply with the following mapping between signature algorithms and message digest algorithms:
	//
	// 	- If the signature algorithm is RSA_PKCS1_SHA_256, RSA_PSS_SHA_256, or ECDSA_SHA_256, the message digest algorithm must be SHA-256.
	//
	// 	- If the signature algorithm is SM2DSA, the message digest algorithm must be SM3.
	//
	// >  If the key type of the certificate is EC_SM2 and the MessageType parameter is set to DIGEST, the value of the Message parameter is `e` that is described in GB/T 32918.2-2016 6.1.
	//
	// This parameter is required.
	//
	// example:
	//
	// VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4=
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The type of the message. Valid values:
	//
	// 	- RAW: the raw data. This is the default value.
	//
	// 	- DIGEST: the message digest (hash value) of the raw data.
	//
	// This parameter is required.
	//
	// example:
	//
	// RAW
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
}

func (s CertificatePrivateKeySignRequest) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrivateKeySignRequest) GoString() string {
	return s.String()
}

func (s *CertificatePrivateKeySignRequest) SetAlgorithm(v string) *CertificatePrivateKeySignRequest {
	s.Algorithm = &v
	return s
}

func (s *CertificatePrivateKeySignRequest) SetCertificateId(v string) *CertificatePrivateKeySignRequest {
	s.CertificateId = &v
	return s
}

func (s *CertificatePrivateKeySignRequest) SetMessage(v string) *CertificatePrivateKeySignRequest {
	s.Message = &v
	return s
}

func (s *CertificatePrivateKeySignRequest) SetMessageType(v string) *CertificatePrivateKeySignRequest {
	s.MessageType = &v
	return s
}

type CertificatePrivateKeySignResponseBody struct {
	// The ID of the certificate.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5979d897-d69f-4fc9-87dd-f3bb73c40b80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The signature value.
	//
	// The value is encoded in Base64.
	//
	// example:
	//
	// ZOyIygCyaOW6Gj****MlNKiuyjfzw=
	SignatureValue *string `json:"SignatureValue,omitempty" xml:"SignatureValue,omitempty"`
}

func (s CertificatePrivateKeySignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrivateKeySignResponseBody) GoString() string {
	return s.String()
}

func (s *CertificatePrivateKeySignResponseBody) SetCertificateId(v string) *CertificatePrivateKeySignResponseBody {
	s.CertificateId = &v
	return s
}

func (s *CertificatePrivateKeySignResponseBody) SetRequestId(v string) *CertificatePrivateKeySignResponseBody {
	s.RequestId = &v
	return s
}

func (s *CertificatePrivateKeySignResponseBody) SetSignatureValue(v string) *CertificatePrivateKeySignResponseBody {
	s.SignatureValue = &v
	return s
}

type CertificatePrivateKeySignResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CertificatePrivateKeySignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CertificatePrivateKeySignResponse) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrivateKeySignResponse) GoString() string {
	return s.String()
}

func (s *CertificatePrivateKeySignResponse) SetHeaders(v map[string]*string) *CertificatePrivateKeySignResponse {
	s.Headers = v
	return s
}

func (s *CertificatePrivateKeySignResponse) SetStatusCode(v int32) *CertificatePrivateKeySignResponse {
	s.StatusCode = &v
	return s
}

func (s *CertificatePrivateKeySignResponse) SetBody(v *CertificatePrivateKeySignResponseBody) *CertificatePrivateKeySignResponse {
	s.Body = v
	return s
}

type CertificatePublicKeyEncryptRequest struct {
	// The encryption algorithm. Valid values:
	//
	// 	- RSAES_OAEP_SHA_1
	//
	// 	- RSAES_OAEP_SHA_256
	//
	// 	- SM2PKE
	//
	// >The SM2PKE encryption algorithm is supported only in regions in mainland China. In these regions, managed hardware security modules (HSMs) are used. For more information, see [Managed HSM overview](https://help.aliyun.com/document_detail/125803.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// RSAES_OAEP_SHA_256
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ID of the certificate. The ID must be globally unique in Certificates Manager.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The data that you want to encrypt.
	//
	// The value is encoded in Base64. For example, if the hexadecimal data that you want to encrypt is `[0x31, 0x32, 0x33, 0x34]`, the Base64-encoded data is `MTIzNA==`.
	//
	// The size of data that can be encrypted varies based on the encryption algorithm that you use:
	//
	// 	- RSAES_OAEP_SHA_1: 214 bytes
	//
	// 	- RSAES_OAEP_SHA_256: 190 bytes
	//
	// 	- SM2PKE: 6,047 bytes
	//
	// If the size of data that you want to encrypt exceeds the preceding limits, you can call the [GenerateDataKey](https://help.aliyun.com/document_detail/28948.html) operation to generate a data key to encrypt the data. Then, call the CertificatePublicKeyEncrypt operation to encrypt the data key.
	//
	// This parameter is required.
	//
	// example:
	//
	// VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4=
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
}

func (s CertificatePublicKeyEncryptRequest) String() string {
	return tea.Prettify(s)
}

func (s CertificatePublicKeyEncryptRequest) GoString() string {
	return s.String()
}

func (s *CertificatePublicKeyEncryptRequest) SetAlgorithm(v string) *CertificatePublicKeyEncryptRequest {
	s.Algorithm = &v
	return s
}

func (s *CertificatePublicKeyEncryptRequest) SetCertificateId(v string) *CertificatePublicKeyEncryptRequest {
	s.CertificateId = &v
	return s
}

func (s *CertificatePublicKeyEncryptRequest) SetPlaintext(v string) *CertificatePublicKeyEncryptRequest {
	s.Plaintext = &v
	return s
}

type CertificatePublicKeyEncryptResponseBody struct {
	// The ID of the certificate.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The ciphertext.
	//
	// The value is encoded in Base64.
	//
	// example:
	//
	// ZOyIygCyaOW6Gj****MlNKiuyjfzw=
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5979d897-d69f-4fc9-87dd-f3bb73c40b80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CertificatePublicKeyEncryptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CertificatePublicKeyEncryptResponseBody) GoString() string {
	return s.String()
}

func (s *CertificatePublicKeyEncryptResponseBody) SetCertificateId(v string) *CertificatePublicKeyEncryptResponseBody {
	s.CertificateId = &v
	return s
}

func (s *CertificatePublicKeyEncryptResponseBody) SetCiphertextBlob(v string) *CertificatePublicKeyEncryptResponseBody {
	s.CiphertextBlob = &v
	return s
}

func (s *CertificatePublicKeyEncryptResponseBody) SetRequestId(v string) *CertificatePublicKeyEncryptResponseBody {
	s.RequestId = &v
	return s
}

type CertificatePublicKeyEncryptResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CertificatePublicKeyEncryptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CertificatePublicKeyEncryptResponse) String() string {
	return tea.Prettify(s)
}

func (s CertificatePublicKeyEncryptResponse) GoString() string {
	return s.String()
}

func (s *CertificatePublicKeyEncryptResponse) SetHeaders(v map[string]*string) *CertificatePublicKeyEncryptResponse {
	s.Headers = v
	return s
}

func (s *CertificatePublicKeyEncryptResponse) SetStatusCode(v int32) *CertificatePublicKeyEncryptResponse {
	s.StatusCode = &v
	return s
}

func (s *CertificatePublicKeyEncryptResponse) SetBody(v *CertificatePublicKeyEncryptResponseBody) *CertificatePublicKeyEncryptResponse {
	s.Body = v
	return s
}

type CertificatePublicKeyVerifyRequest struct {
	// The signature algorithm. Valid values:
	//
	// 	- RSA_PKCS1_SHA_256
	//
	// 	- RSA_PSS_SHA_256
	//
	// 	- ECDSA_SHA_256
	//
	// 	- SM2DSA
	//
	// > The SM2DSA signature algorithm is supported only in regions where managed hardware security modules (HSMs) are used in the Chinese mainland. For more information, see [Managed HSM overview](https://help.aliyun.com/document_detail/125803.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ECDSA_SHA_256
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ID of the certificate. The ID must be globally unique in Certificates Manager.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The raw data that is signed.
	//
	// The value is encoded in Base64. For example, if the raw data in the hexadecimal format is `[0x31, 0x32, 0x33, 0x34]`, set this parameter to the Base64-encoded value `MTIzNA==`.
	//
	// If the MessageType parameter is set to RAW, the size of the data must be less than or equal to 4 KB.
	//
	// If the size of the data is greater than 4 KB, you can set the MessageType parameter to DIGEST and set the Message parameter to the digest of the data. The digest is also called hash value. You can compute the digest of the data on an on-premises device. Certificates Manager uses the digest that you compute in your own certificate application system. The message digest algorithm that you use must match the specified signature algorithm. Comply with the following mapping between signature algorithms and message digest algorithms:
	//
	// 	- If the signature algorithm is RSA_PKCS1_SHA_256, RSA_PSS_SHA_256, or ECDSA_SHA_256, the message digest algorithm must be SHA-256.
	//
	// 	- If the signature algorithm is SM2DSA, the message digest algorithm must be SM3.
	//
	// >  If the key type of the certificate is EC_SM2 and the MessageType parameter is set to DIGEST, the value of the Message parameter is `e` that is described in GB/T 32918.2-2016 6.1.
	//
	// This parameter is required.
	//
	// example:
	//
	// VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4=
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The type of the message. Valid values:
	//
	// 	- RAW: the raw data. This is the default value.
	//
	// 	- DIGEST: the message digest (hash value) of the raw data.
	//
	// This parameter is required.
	//
	// example:
	//
	// RAW
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The signature value.
	//
	// The value is encoded in Base64.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZOyIygCyaOW6Gj****MlNKiuyjfzw=
	SignatureValue *string `json:"SignatureValue,omitempty" xml:"SignatureValue,omitempty"`
}

func (s CertificatePublicKeyVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s CertificatePublicKeyVerifyRequest) GoString() string {
	return s.String()
}

func (s *CertificatePublicKeyVerifyRequest) SetAlgorithm(v string) *CertificatePublicKeyVerifyRequest {
	s.Algorithm = &v
	return s
}

func (s *CertificatePublicKeyVerifyRequest) SetCertificateId(v string) *CertificatePublicKeyVerifyRequest {
	s.CertificateId = &v
	return s
}

func (s *CertificatePublicKeyVerifyRequest) SetMessage(v string) *CertificatePublicKeyVerifyRequest {
	s.Message = &v
	return s
}

func (s *CertificatePublicKeyVerifyRequest) SetMessageType(v string) *CertificatePublicKeyVerifyRequest {
	s.MessageType = &v
	return s
}

func (s *CertificatePublicKeyVerifyRequest) SetSignatureValue(v string) *CertificatePublicKeyVerifyRequest {
	s.SignatureValue = &v
	return s
}

type CertificatePublicKeyVerifyResponseBody struct {
	// The ID of the certificate.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5979d897-d69f-4fc9-87dd-f3bb73c40b80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The verification result. Valid values:
	//
	// 	- true: The signature is valid.
	//
	// 	- false: The signature is invalid.
	//
	// example:
	//
	// true
	SignatureValid *bool `json:"SignatureValid,omitempty" xml:"SignatureValid,omitempty"`
}

func (s CertificatePublicKeyVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CertificatePublicKeyVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *CertificatePublicKeyVerifyResponseBody) SetCertificateId(v string) *CertificatePublicKeyVerifyResponseBody {
	s.CertificateId = &v
	return s
}

func (s *CertificatePublicKeyVerifyResponseBody) SetRequestId(v string) *CertificatePublicKeyVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CertificatePublicKeyVerifyResponseBody) SetSignatureValid(v bool) *CertificatePublicKeyVerifyResponseBody {
	s.SignatureValid = &v
	return s
}

type CertificatePublicKeyVerifyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CertificatePublicKeyVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CertificatePublicKeyVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s CertificatePublicKeyVerifyResponse) GoString() string {
	return s.String()
}

func (s *CertificatePublicKeyVerifyResponse) SetHeaders(v map[string]*string) *CertificatePublicKeyVerifyResponse {
	s.Headers = v
	return s
}

func (s *CertificatePublicKeyVerifyResponse) SetStatusCode(v int32) *CertificatePublicKeyVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *CertificatePublicKeyVerifyResponse) SetBody(v *CertificatePublicKeyVerifyResponseBody) *CertificatePublicKeyVerifyResponse {
	s.Body = v
	return s
}

type ConnectKmsInstanceRequest struct {
	// The provider of the KMS instance. Set the value to Aliyun.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	KMProvider *string `json:"KMProvider,omitempty" xml:"KMProvider,omitempty"`
	// The ID of the KMS instance that you want to enable.
	//
	// This parameter is required.
	//
	// example:
	//
	// kst-phzz64f722a1buamw0****
	KmsInstanceId *string `json:"KmsInstanceId,omitempty" xml:"KmsInstanceId,omitempty"`
	// The vSwitch in the two zones. The vSwitch must have at least one available IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1i512amda6d10a0****
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The ID of the virtual private cloud (VPC) that is associated with the KMS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp19z7cwmltad5dff****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The two zones for the KMS instance. Dual-zone deployment improves service availability and disaster recovery capabilities.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-k,cn-hangzhou-j
	ZoneIds *string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty"`
}

func (s ConnectKmsInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ConnectKmsInstanceRequest) GoString() string {
	return s.String()
}

func (s *ConnectKmsInstanceRequest) SetKMProvider(v string) *ConnectKmsInstanceRequest {
	s.KMProvider = &v
	return s
}

func (s *ConnectKmsInstanceRequest) SetKmsInstanceId(v string) *ConnectKmsInstanceRequest {
	s.KmsInstanceId = &v
	return s
}

func (s *ConnectKmsInstanceRequest) SetVSwitchIds(v string) *ConnectKmsInstanceRequest {
	s.VSwitchIds = &v
	return s
}

func (s *ConnectKmsInstanceRequest) SetVpcId(v string) *ConnectKmsInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *ConnectKmsInstanceRequest) SetZoneIds(v string) *ConnectKmsInstanceRequest {
	s.ZoneIds = &v
	return s
}

type ConnectKmsInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// d3eca5c8-a856-4347-8eb6-e1898c3fda2e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConnectKmsInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConnectKmsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ConnectKmsInstanceResponseBody) SetRequestId(v string) *ConnectKmsInstanceResponseBody {
	s.RequestId = &v
	return s
}

type ConnectKmsInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConnectKmsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConnectKmsInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ConnectKmsInstanceResponse) GoString() string {
	return s.String()
}

func (s *ConnectKmsInstanceResponse) SetHeaders(v map[string]*string) *ConnectKmsInstanceResponse {
	s.Headers = v
	return s
}

func (s *ConnectKmsInstanceResponse) SetStatusCode(v int32) *ConnectKmsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ConnectKmsInstanceResponse) SetBody(v *ConnectKmsInstanceResponseBody) *ConnectKmsInstanceResponse {
	s.Body = v
	return s
}

type CreateAliasRequest struct {
	// The alias of the CMK.
	//
	// The alias must be 1 to 255 characters in length and must contain the prefix `alias/`. The alias cannot be prefixed with the reserved word `alias/acs`.
	//
	// This parameter is required.
	//
	// example:
	//
	// alias/example
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7906979c-8e06-46a2-be2d-68e3ccbc****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s CreateAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAliasRequest) GoString() string {
	return s.String()
}

func (s *CreateAliasRequest) SetAliasName(v string) *CreateAliasRequest {
	s.AliasName = &v
	return s
}

func (s *CreateAliasRequest) SetKeyId(v string) *CreateAliasRequest {
	s.KeyId = &v
	return s
}

type CreateAliasResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 1d2baaf3-d357-46c2-832e-13560c2bd9cd
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAliasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAliasResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAliasResponseBody) SetRequestId(v string) *CreateAliasResponseBody {
	s.RequestId = &v
	return s
}

type CreateAliasResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAliasResponse) GoString() string {
	return s.String()
}

func (s *CreateAliasResponse) SetHeaders(v map[string]*string) *CreateAliasResponse {
	s.Headers = v
	return s
}

func (s *CreateAliasResponse) SetStatusCode(v int32) *CreateAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAliasResponse) SetBody(v *CreateAliasResponseBody) *CreateAliasResponse {
	s.Body = v
	return s
}

type CreateApplicationAccessPointRequest struct {
	// The authentication method. Currently, only ClientKey is supported.
	//
	// example:
	//
	// ClientKey
	AuthenticationMethod *string `json:"AuthenticationMethod,omitempty" xml:"AuthenticationMethod,omitempty"`
	// The description of the AAP.
	//
	// example:
	//
	// aap description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the AAP.
	//
	// This parameter is required.
	//
	// example:
	//
	// aap_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The permission policy.
	//
	// > You can bind up to three permission policies to each AAP.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["kst-hzz62ee817bvyyr5x****.efkd","kst-hzz62ee817bvyyr5x****.eyyp"]
	Policies *string `json:"Policies,omitempty" xml:"Policies,omitempty"`
}

func (s CreateApplicationAccessPointRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationAccessPointRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationAccessPointRequest) SetAuthenticationMethod(v string) *CreateApplicationAccessPointRequest {
	s.AuthenticationMethod = &v
	return s
}

func (s *CreateApplicationAccessPointRequest) SetDescription(v string) *CreateApplicationAccessPointRequest {
	s.Description = &v
	return s
}

func (s *CreateApplicationAccessPointRequest) SetName(v string) *CreateApplicationAccessPointRequest {
	s.Name = &v
	return s
}

func (s *CreateApplicationAccessPointRequest) SetPolicies(v string) *CreateApplicationAccessPointRequest {
	s.Policies = &v
	return s
}

type CreateApplicationAccessPointResponseBody struct {
	// The Alibaba Cloud Resource Name (ARN) of the AAP.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:119285303511****:applicationaccesspoint/aap_test
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The authentication method.
	//
	// example:
	//
	// ClientKey
	AuthenticationMethod *string `json:"AuthenticationMethod,omitempty" xml:"AuthenticationMethod,omitempty"`
	// The description of the AAP.
	//
	// example:
	//
	// aap description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the AAP.
	//
	// example:
	//
	// aap_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The permission policy.
	//
	// example:
	//
	// ["kst-hzz62ee817bvyyr5x****.efkd","kst-hzz62ee817bvyyr5x****.eyyp"]
	Policies *string `json:"Policies,omitempty" xml:"Policies,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// bcfefe15-46f0-44a3-bd96-3d422474b71a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationAccessPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationAccessPointResponseBody) SetArn(v string) *CreateApplicationAccessPointResponseBody {
	s.Arn = &v
	return s
}

func (s *CreateApplicationAccessPointResponseBody) SetAuthenticationMethod(v string) *CreateApplicationAccessPointResponseBody {
	s.AuthenticationMethod = &v
	return s
}

func (s *CreateApplicationAccessPointResponseBody) SetDescription(v string) *CreateApplicationAccessPointResponseBody {
	s.Description = &v
	return s
}

func (s *CreateApplicationAccessPointResponseBody) SetName(v string) *CreateApplicationAccessPointResponseBody {
	s.Name = &v
	return s
}

func (s *CreateApplicationAccessPointResponseBody) SetPolicies(v string) *CreateApplicationAccessPointResponseBody {
	s.Policies = &v
	return s
}

func (s *CreateApplicationAccessPointResponseBody) SetRequestId(v string) *CreateApplicationAccessPointResponseBody {
	s.RequestId = &v
	return s
}

type CreateApplicationAccessPointResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationAccessPointResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationAccessPointResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationAccessPointResponse) SetHeaders(v map[string]*string) *CreateApplicationAccessPointResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationAccessPointResponse) SetStatusCode(v int32) *CreateApplicationAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationAccessPointResponse) SetBody(v *CreateApplicationAccessPointResponseBody) *CreateApplicationAccessPointResponse {
	s.Body = v
	return s
}

type CreateCertificateRequest struct {
	// Specifies whether the private key of the certificate can be exported for use. Valid values:
	//
	// 	- true: The private key of the certificate can be exported for use. This is the default value.
	//
	// 	- false: The private key of the certificate cannot be exported for use. We recommend that you set this parameter to false to protect keys with a higher security level.
	//
	// example:
	//
	// true
	ExportablePrivateKey *bool `json:"ExportablePrivateKey,omitempty" xml:"ExportablePrivateKey,omitempty"`
	// The type of the key. Valid values:
	//
	// 	- RSA_2048
	//
	// 	- EC_P256
	//
	// 	- EC_SM2
	//
	// This parameter is required.
	//
	// example:
	//
	// RSA_2048
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The certificate subject, which is the owner of the certificate.
	//
	// Specify the value in the distinguished name (DN) format, as defined in [RFC 2253](https://tools.ietf.org/html/rfc2253?spm=a2c4g.11186623.2.13.265f1a1cGFCn3Q). A DN is a sequence of relative distinguished names (RDNs).
	//
	// RDNs are key-value pairs in the format of `attribute1=value1,attribute2=value2`. Separate multiple RDNs with commas (,).
	//
	// The Subject parameter consists of the following fields:
	//
	// 	- CN: required. The name of the certificate subject.
	//
	// 	- C: required. The two-character country or region code in the [ISO 3166-1](https://www.iso.org/obp/ui/#search/code/) standard. For example, CN indicates China.
	//
	// 	- O: required. The legal name of the enterprise, company, organization, or institution.
	//
	// 	- OU: required. The name of the department.
	//
	// 	- ST: optional. The name of the province, municipality, autonomous region, or special administrative region.
	//
	// 	- L: optional. The name of the city.
	//
	// This parameter is required.
	//
	// example:
	//
	// CN=userName,OU=kms,O=aliyun,C=CN
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// The subject alternative names.
	//
	// A domain name list is supported. A maximum of 10 domain names are supported.
	//
	// example:
	//
	// ["test1.example.com","test2.example.com"]
	SubjectAlternativeNames map[string]interface{} `json:"SubjectAlternativeNames,omitempty" xml:"SubjectAlternativeNames,omitempty"`
}

func (s CreateCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateCertificateRequest) SetExportablePrivateKey(v bool) *CreateCertificateRequest {
	s.ExportablePrivateKey = &v
	return s
}

func (s *CreateCertificateRequest) SetKeySpec(v string) *CreateCertificateRequest {
	s.KeySpec = &v
	return s
}

func (s *CreateCertificateRequest) SetSubject(v string) *CreateCertificateRequest {
	s.Subject = &v
	return s
}

func (s *CreateCertificateRequest) SetSubjectAlternativeNames(v map[string]interface{}) *CreateCertificateRequest {
	s.SubjectAlternativeNames = v
	return s
}

type CreateCertificateShrinkRequest struct {
	// Specifies whether the private key of the certificate can be exported for use. Valid values:
	//
	// 	- true: The private key of the certificate can be exported for use. This is the default value.
	//
	// 	- false: The private key of the certificate cannot be exported for use. We recommend that you set this parameter to false to protect keys with a higher security level.
	//
	// example:
	//
	// true
	ExportablePrivateKey *bool `json:"ExportablePrivateKey,omitempty" xml:"ExportablePrivateKey,omitempty"`
	// The type of the key. Valid values:
	//
	// 	- RSA_2048
	//
	// 	- EC_P256
	//
	// 	- EC_SM2
	//
	// This parameter is required.
	//
	// example:
	//
	// RSA_2048
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The certificate subject, which is the owner of the certificate.
	//
	// Specify the value in the distinguished name (DN) format, as defined in [RFC 2253](https://tools.ietf.org/html/rfc2253?spm=a2c4g.11186623.2.13.265f1a1cGFCn3Q). A DN is a sequence of relative distinguished names (RDNs).
	//
	// RDNs are key-value pairs in the format of `attribute1=value1,attribute2=value2`. Separate multiple RDNs with commas (,).
	//
	// The Subject parameter consists of the following fields:
	//
	// 	- CN: required. The name of the certificate subject.
	//
	// 	- C: required. The two-character country or region code in the [ISO 3166-1](https://www.iso.org/obp/ui/#search/code/) standard. For example, CN indicates China.
	//
	// 	- O: required. The legal name of the enterprise, company, organization, or institution.
	//
	// 	- OU: required. The name of the department.
	//
	// 	- ST: optional. The name of the province, municipality, autonomous region, or special administrative region.
	//
	// 	- L: optional. The name of the city.
	//
	// This parameter is required.
	//
	// example:
	//
	// CN=userName,OU=kms,O=aliyun,C=CN
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// The subject alternative names.
	//
	// A domain name list is supported. A maximum of 10 domain names are supported.
	//
	// example:
	//
	// ["test1.example.com","test2.example.com"]
	SubjectAlternativeNamesShrink *string `json:"SubjectAlternativeNames,omitempty" xml:"SubjectAlternativeNames,omitempty"`
}

func (s CreateCertificateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCertificateShrinkRequest) SetExportablePrivateKey(v bool) *CreateCertificateShrinkRequest {
	s.ExportablePrivateKey = &v
	return s
}

func (s *CreateCertificateShrinkRequest) SetKeySpec(v string) *CreateCertificateShrinkRequest {
	s.KeySpec = &v
	return s
}

func (s *CreateCertificateShrinkRequest) SetSubject(v string) *CreateCertificateShrinkRequest {
	s.Subject = &v
	return s
}

func (s *CreateCertificateShrinkRequest) SetSubjectAlternativeNamesShrink(v string) *CreateCertificateShrinkRequest {
	s.SubjectAlternativeNamesShrink = &v
	return s
}

type CreateCertificateResponseBody struct {
	// The Alibaba Cloud Resource Name (ARN) of the certificate.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:154035569884****:certificate/98e85c94-52d0-40c9-b3b2-afda52f4****
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the certificate. It is the globally unique identifier (GUID) of the certificate in Certificates Manager.
	//
	// example:
	//
	// 9a28de48-8d8b-484d-a766-dec4****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The CSR in the PEM format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST-----\\nMIIDADCCAegCAQAwgboxCzAJBgNVBAYTAkNOMREwDwYDVQQIEwhaaGVqaWFuZzER\\n****\\nmkj4rg==\\n-----END CERTIFICATE REQUEST-----\\n
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 15a735a1-8fe6-45cc-a64c-3c4ff839334e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCertificateResponseBody) SetArn(v string) *CreateCertificateResponseBody {
	s.Arn = &v
	return s
}

func (s *CreateCertificateResponseBody) SetCertificateId(v string) *CreateCertificateResponseBody {
	s.CertificateId = &v
	return s
}

func (s *CreateCertificateResponseBody) SetCsr(v string) *CreateCertificateResponseBody {
	s.Csr = &v
	return s
}

func (s *CreateCertificateResponseBody) SetRequestId(v string) *CreateCertificateResponseBody {
	s.RequestId = &v
	return s
}

type CreateCertificateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateResponse) GoString() string {
	return s.String()
}

func (s *CreateCertificateResponse) SetHeaders(v map[string]*string) *CreateCertificateResponse {
	s.Headers = v
	return s
}

func (s *CreateCertificateResponse) SetStatusCode(v int32) *CreateCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCertificateResponse) SetBody(v *CreateCertificateResponseBody) *CreateCertificateResponse {
	s.Body = v
	return s
}

type CreateClientKeyRequest struct {
	// The operation that you want to perform. Set the value to **CreateClientKey**.
	//
	// This parameter is required.
	//
	// example:
	//
	// aap_test
	AapName *string `json:"AapName,omitempty" xml:"AapName,omitempty"`
	// The encryption password of the client key.
	//
	// The password must be 8 to 64 characters in length and must contain at least two of the following types: digits, letters, and special characters. Special characters include `~ ! @ # $ % ^ & 	- ? _ -`.
	//
	// example:
	//
	// 2028-08-31T17:14:33Z
	NotAfter *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The end of the validity period of the client key.
	//
	// Specify the time in the ISO 8601 standard. The time must be in UTC. The time must be in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// >
	//
	// 	- If you do not configure NotAfter, the default value is the time when the client key was created plus five years.
	//
	// 	- If you configure NotAfter, you must configure NotBefore.
	//
	// example:
	//
	// 2023-08-31T17:14:33Z
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The name of the AAP.
	//
	// This parameter is required.
	//
	// example:
	//
	// bcfefe15-46f0****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s CreateClientKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClientKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateClientKeyRequest) SetAapName(v string) *CreateClientKeyRequest {
	s.AapName = &v
	return s
}

func (s *CreateClientKeyRequest) SetNotAfter(v string) *CreateClientKeyRequest {
	s.NotAfter = &v
	return s
}

func (s *CreateClientKeyRequest) SetNotBefore(v string) *CreateClientKeyRequest {
	s.NotBefore = &v
	return s
}

func (s *CreateClientKeyRequest) SetPassword(v string) *CreateClientKeyRequest {
	s.Password = &v
	return s
}

type CreateClientKeyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// KAAP.66abf237-63f6-4625-b8cf-47e1086e****
	ClientKeyId *string `json:"ClientKeyId,omitempty" xml:"ClientKeyId,omitempty"`
	// The ID of the client key.
	//
	// example:
	//
	// RSA_2048
	KeyAlgorithm *string `json:"KeyAlgorithm,omitempty" xml:"KeyAlgorithm,omitempty"`
	// The beginning of the validity period of the client key.
	//
	// example:
	//
	// 2028-08-31T17:14:33Z
	NotAfter *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The private key of the client key.
	//
	// example:
	//
	// 2023-08-31T17:14:33Z
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The algorithm that is used to encrypt the private key of the client key. Currently, only RSA_2048 is supported.
	//
	// example:
	//
	// MIIJqwIBAzCCCXcGCSqGSIb3DQEHAaCCCWgEgglkMIIJYDCCBBcGCSqGSIb3DQEHBqCCBAgwgg******
	PrivateKeyData *string `json:"PrivateKeyData,omitempty" xml:"PrivateKeyData,omitempty"`
	// The beginning of the validity period of the client key.
	//
	// Specify the time in the ISO 8601 standard. The time must be in UTC. The time must be in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// >
	//
	// 	- If you do not configure NotBefore, the default value is the time when the client key was created.
	//
	// 	- If you configure NotBefore, you must configure NotAfter.
	//
	// example:
	//
	// 2312e45f-b2fa-4c34-ad94-3eca50932916
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateClientKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClientKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClientKeyResponseBody) SetClientKeyId(v string) *CreateClientKeyResponseBody {
	s.ClientKeyId = &v
	return s
}

func (s *CreateClientKeyResponseBody) SetKeyAlgorithm(v string) *CreateClientKeyResponseBody {
	s.KeyAlgorithm = &v
	return s
}

func (s *CreateClientKeyResponseBody) SetNotAfter(v string) *CreateClientKeyResponseBody {
	s.NotAfter = &v
	return s
}

func (s *CreateClientKeyResponseBody) SetNotBefore(v string) *CreateClientKeyResponseBody {
	s.NotBefore = &v
	return s
}

func (s *CreateClientKeyResponseBody) SetPrivateKeyData(v string) *CreateClientKeyResponseBody {
	s.PrivateKeyData = &v
	return s
}

func (s *CreateClientKeyResponseBody) SetRequestId(v string) *CreateClientKeyResponseBody {
	s.RequestId = &v
	return s
}

type CreateClientKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClientKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClientKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClientKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateClientKeyResponse) SetHeaders(v map[string]*string) *CreateClientKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateClientKeyResponse) SetStatusCode(v int32) *CreateClientKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClientKeyResponse) SetBody(v *CreateClientKeyResponseBody) *CreateClientKeyResponse {
	s.Body = v
	return s
}

type CreateKeyRequest struct {
	// The ID of the KMS instance.
	//
	// > You must specify this parameter if you need to create a key for a KMS instance. If you need to create a default key of the CMK type, you do not need to specify this parameter.
	//
	// example:
	//
	// kst-bjj62d8f5e0sgtx8h****
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// The description of the key.
	//
	// The description can be 0 to 8,192 characters in length.
	//
	// example:
	//
	// key description example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable automatic key rotation. Valid values:
	//
	// - true
	//
	// - false (default)
	//
	// This parameter is valid only when the key belongs to an instance type that supports automatic rotation. For more information, see [Key rotation](https://help.aliyun.com/document_detail/2358146.html).
	//
	// example:
	//
	// true
	EnableAutomaticRotation *bool `json:"EnableAutomaticRotation,omitempty" xml:"EnableAutomaticRotation,omitempty"`
	// The key specification. The valid values vary based on the KMS instance type. For more information, see [Overview](https://help.aliyun.com/document_detail/480159.html).
	//
	// > If you do not specify a value for this parameter, the default key specification is Aliyun_AES_256.
	//
	// example:
	//
	// Aliyun_AES_256
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The usage of the key. Valid values:
	//
	// - ENCRYPT/DECRYPT
	//
	// - SIGN/VERIFY
	//
	// If the key supports signing and verification, the default value is SIGN/VERIFY. If the key does not support signing and verification, the default value is ENCRYPT/DECRYPT.
	//
	// example:
	//
	// ENCRYPT/DECRYPT
	KeyUsage *string `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty"`
	// The key material origin. Valid values:
	//
	// - Aliyun_KMS (default): KMS generates key material.
	//
	// - EXTERNAL: You import key material.
	//
	//
	// > - The value of this parameter is case-sensitive.
	//
	// > - Default keys of the customer master key (CMK) type support Aliyun_KMS and EXTERNAL. Keys in instances of the software key management type support only Aliyun_KMS. Keys in instances of the hardware key management type support Aliyun_KMS and EXTERNAL.
	//
	// > - If you set Origin to EXTERNAL, you must import key material. For more information, see [Import key material into a symmetric key](https://help.aliyun.com/document_detail/607841.html) or [Import key material into an asymmetric key](https://help.aliyun.com/document_detail/608827.html).
	//
	// example:
	//
	// Aliyun_KMS
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// You do not need to specify this parameter. KMS sets a protection level for your key.
	//
	// The protection level of the key. Valid values:
	//
	// - SOFTWARE
	//
	// - HSM
	//
	//
	// > - If DKMSInstanceId is specified, this parameter does not take effect. If your instance is an instance of the software key management type, set the value to SOFTWARE. If your instance is an instance of the hardware key management type, set the value to HSM.
	//
	// > - If you do not specify DKMSInstanceId, we recommend that you do not specify this parameter. KMS sets a protection level for your key. If managed hardware security modules (HSMs) exist in the region of your KMS instance, set the value to HSM. If managed HSMs do not exist in the region of your KMS instance, set the value to SOFTWARE. For more information, see Managed HSM overview.
	//
	// example:
	//
	// SOFTWARE
	ProtectionLevel *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	// The period of automatic key rotation. Format: integer[unit]. Unit: d (day), h (hour), m (minute), or s (second). For example, both 7d and 604800s represent a seven-day interval.
	//
	// - For a default key, set the value to 365 days.
	//
	// - For a software-protected key, set a value that ranges from 7 to 365 days.
	//
	// - A hardware-protected key does not support automatic rotation.
	//
	// > If EnableAutomaticRotation is set to true, this parameter is required.
	//
	// example:
	//
	// 365d
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	// The tag that is added to the key. A tag consists of a key-value pair.
	//
	// You can enter up to 20 tags. Enter multiple tags in the [{"TagKey":"key1","TagValue":"value1"},{"TagKey":"key2","TagValue":"value2"},..] format.
	//
	// Each tag key or tag value can be up to 128 characters in length and can contain letters, digits, forward slashes (/), backslashes (\\), underscores (_), hyphens (-), periods (.), plus signs (+), equal signs (=), colons (:), and at signs (@).
	//
	// > The tag key cannot start with aliyun or acs:.
	//
	// example:
	//
	// [{"TagKey":"disk-encryption","TagValue":"true"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateKeyRequest) SetDKMSInstanceId(v string) *CreateKeyRequest {
	s.DKMSInstanceId = &v
	return s
}

func (s *CreateKeyRequest) SetDescription(v string) *CreateKeyRequest {
	s.Description = &v
	return s
}

func (s *CreateKeyRequest) SetEnableAutomaticRotation(v bool) *CreateKeyRequest {
	s.EnableAutomaticRotation = &v
	return s
}

func (s *CreateKeyRequest) SetKeySpec(v string) *CreateKeyRequest {
	s.KeySpec = &v
	return s
}

func (s *CreateKeyRequest) SetKeyUsage(v string) *CreateKeyRequest {
	s.KeyUsage = &v
	return s
}

func (s *CreateKeyRequest) SetOrigin(v string) *CreateKeyRequest {
	s.Origin = &v
	return s
}

func (s *CreateKeyRequest) SetPolicy(v string) *CreateKeyRequest {
	s.Policy = &v
	return s
}

func (s *CreateKeyRequest) SetProtectionLevel(v string) *CreateKeyRequest {
	s.ProtectionLevel = &v
	return s
}

func (s *CreateKeyRequest) SetRotationInterval(v string) *CreateKeyRequest {
	s.RotationInterval = &v
	return s
}

func (s *CreateKeyRequest) SetTags(v string) *CreateKeyRequest {
	s.Tags = &v
	return s
}

type CreateKeyResponseBody struct {
	// The metadata of the key.
	KeyMetadata *CreateKeyResponseBodyKeyMetadata `json:"KeyMetadata,omitempty" xml:"KeyMetadata,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 381D5D33-BB8F-395F-8EE4-AE3BB4B523C4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKeyResponseBody) SetKeyMetadata(v *CreateKeyResponseBodyKeyMetadata) *CreateKeyResponseBody {
	s.KeyMetadata = v
	return s
}

func (s *CreateKeyResponseBody) SetRequestId(v string) *CreateKeyResponseBody {
	s.RequestId = &v
	return s
}

type CreateKeyResponseBodyKeyMetadata struct {
	// The Alibaba Cloud Resource Name (ARN) of the key.
	//
	// example:
	//
	// acs:kms:cn-qingdao:154035569884****:key/key-hzz62f1cb66fa42qo****
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The status of automatic key rotation. Valid values:
	//
	// - Enabled
	//
	// - Disabled
	//
	// - Suspended
	//
	// example:
	//
	// Enabled
	AutomaticRotation *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	// The date and time (UTC) when the key was created.
	//
	// example:
	//
	// 2023-03-25T10:00:00Z
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// The user who created the key.
	//
	// example:
	//
	// 154035569884****
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The ID of the KMS instance.
	//
	// example:
	//
	// kst-bjj62d8f5e0sgtx8h****
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// The time when the key is scheduled for deletion. For more information, see ScheduleKeyDeletion.
	//
	// This parameter is returned only when the value of KeyState is PendingDeletion.
	//
	// example:
	//
	// 2025-03-25T10:00:00Z
	DeleteDate *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
	// The description of the key.
	//
	// example:
	//
	// key description example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The globally unique ID of the key.
	//
	// example:
	//
	// key-hzz62f1cb66fa42qo****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The specification of the key.
	//
	// example:
	//
	// Aliyun_AES_256
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The status of the key.
	//
	// For more information, see [Impacts of key status on API operations](https://help.aliyun.com/document_detail/44211.html).
	//
	// example:
	//
	// Enabled
	KeyState *string `json:"KeyState,omitempty" xml:"KeyState,omitempty"`
	// The usage of the key.
	//
	// example:
	//
	// ENCRYPT/DECRYPT
	KeyUsage *string `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty"`
	// The time when the last rotation was performed. The time is displayed in UTC.
	//
	// For a new key, this parameter value is the time when the initial version of the key was generated.
	//
	// example:
	//
	// 2023-03-25T10:00:00Z
	LastRotationDate *string `json:"LastRotationDate,omitempty" xml:"LastRotationDate,omitempty"`
	// The time when the key material expires. The time is displayed in UTC.
	//
	// If this parameter value is empty, the key material does not expire.
	//
	// example:
	//
	// 2025-03-25T10:00:00Z
	MaterialExpireTime *string `json:"MaterialExpireTime,omitempty" xml:"MaterialExpireTime,omitempty"`
	// The time when the key is next rotated.
	//
	// This value is returned only when the value of AutomaticRotation is Enabled or Suspended.
	//
	// example:
	//
	// 2024-03-25T10:00:00Z
	NextRotationDate *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	// The key material origin.
	//
	// example:
	//
	// Aliyun_KMS
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The current primary version identifier of the key.
	//
	// example:
	//
	// 7ce1d081-06cb-42e6-aab6-5c5de030****
	PrimaryKeyVersion *string `json:"PrimaryKeyVersion,omitempty" xml:"PrimaryKeyVersion,omitempty"`
	// The protection level of the key.
	//
	// example:
	//
	// SOFTWARE
	ProtectionLevel *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	// The interval for automatic key rotation. Unit: seconds. The format is an integer value followed by the character s. For example, if the rotation period is seven days, this parameter is set to 604800s.
	//
	// This value is returned only when the value of AutomaticRotation is Enabled or Suspended.
	//
	// example:
	//
	// 31536000s
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
}

func (s CreateKeyResponseBodyKeyMetadata) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyResponseBodyKeyMetadata) GoString() string {
	return s.String()
}

func (s *CreateKeyResponseBodyKeyMetadata) SetArn(v string) *CreateKeyResponseBodyKeyMetadata {
	s.Arn = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetAutomaticRotation(v string) *CreateKeyResponseBodyKeyMetadata {
	s.AutomaticRotation = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetCreationDate(v string) *CreateKeyResponseBodyKeyMetadata {
	s.CreationDate = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetCreator(v string) *CreateKeyResponseBodyKeyMetadata {
	s.Creator = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetDKMSInstanceId(v string) *CreateKeyResponseBodyKeyMetadata {
	s.DKMSInstanceId = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetDeleteDate(v string) *CreateKeyResponseBodyKeyMetadata {
	s.DeleteDate = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetDescription(v string) *CreateKeyResponseBodyKeyMetadata {
	s.Description = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetKeyId(v string) *CreateKeyResponseBodyKeyMetadata {
	s.KeyId = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetKeySpec(v string) *CreateKeyResponseBodyKeyMetadata {
	s.KeySpec = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetKeyState(v string) *CreateKeyResponseBodyKeyMetadata {
	s.KeyState = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetKeyUsage(v string) *CreateKeyResponseBodyKeyMetadata {
	s.KeyUsage = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetLastRotationDate(v string) *CreateKeyResponseBodyKeyMetadata {
	s.LastRotationDate = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetMaterialExpireTime(v string) *CreateKeyResponseBodyKeyMetadata {
	s.MaterialExpireTime = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetNextRotationDate(v string) *CreateKeyResponseBodyKeyMetadata {
	s.NextRotationDate = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetOrigin(v string) *CreateKeyResponseBodyKeyMetadata {
	s.Origin = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetPrimaryKeyVersion(v string) *CreateKeyResponseBodyKeyMetadata {
	s.PrimaryKeyVersion = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetProtectionLevel(v string) *CreateKeyResponseBodyKeyMetadata {
	s.ProtectionLevel = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetRotationInterval(v string) *CreateKeyResponseBodyKeyMetadata {
	s.RotationInterval = &v
	return s
}

type CreateKeyResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateKeyResponse) SetHeaders(v map[string]*string) *CreateKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateKeyResponse) SetStatusCode(v int32) *CreateKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKeyResponse) SetBody(v *CreateKeyResponseBody) *CreateKeyResponse {
	s.Body = v
	return s
}

type CreateKeyVersionRequest struct {
	// The ID of the CMK. The ID must be globally unique.
	//
	// >  You can also set the value to an alias that is bound to the CMK. For more information, see [Overview of aliases](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0b30658a-ed1a-4922-b8f7-a673ca9c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s CreateKeyVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateKeyVersionRequest) SetKeyId(v string) *CreateKeyVersionRequest {
	s.KeyId = &v
	return s
}

type CreateKeyVersionResponseBody struct {
	// The metadata of the version.
	KeyVersion *CreateKeyVersionResponseBodyKeyVersion `json:"KeyVersion,omitempty" xml:"KeyVersion,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// b96f250a-4b75-498c-91be-22c6928f85be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateKeyVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKeyVersionResponseBody) SetKeyVersion(v *CreateKeyVersionResponseBodyKeyVersion) *CreateKeyVersionResponseBody {
	s.KeyVersion = v
	return s
}

func (s *CreateKeyVersionResponseBody) SetRequestId(v string) *CreateKeyVersionResponseBody {
	s.RequestId = &v
	return s
}

type CreateKeyVersionResponseBodyKeyVersion struct {
	// The date and time when the version was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-08-02T10:38:27Z
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// example:
	//
	// 0b30658a-ed1a-4922-b8f7-a673ca9c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the version.
	//
	// example:
	//
	// c0a3d5dc-0b47-4199-a050-b289349a****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s CreateKeyVersionResponseBodyKeyVersion) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyVersionResponseBodyKeyVersion) GoString() string {
	return s.String()
}

func (s *CreateKeyVersionResponseBodyKeyVersion) SetCreationDate(v string) *CreateKeyVersionResponseBodyKeyVersion {
	s.CreationDate = &v
	return s
}

func (s *CreateKeyVersionResponseBodyKeyVersion) SetKeyId(v string) *CreateKeyVersionResponseBodyKeyVersion {
	s.KeyId = &v
	return s
}

func (s *CreateKeyVersionResponseBodyKeyVersion) SetKeyVersionId(v string) *CreateKeyVersionResponseBodyKeyVersion {
	s.KeyVersionId = &v
	return s
}

type CreateKeyVersionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKeyVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKeyVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateKeyVersionResponse) SetHeaders(v map[string]*string) *CreateKeyVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateKeyVersionResponse) SetStatusCode(v int32) *CreateKeyVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKeyVersionResponse) SetBody(v *CreateKeyVersionResponseBody) *CreateKeyVersionResponse {
	s.Body = v
	return s
}

type CreateNetworkRuleRequest struct {
	// The description.
	//
	// example:
	//
	// networkrule description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the access control rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// networkrule_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The private IP address or private CIDR block. Separate multiple items with commas (,).
	//
	// example:
	//
	// ["192.10.XX.XX","192.168.XX.XX/24"]
	SourcePrivateIp *string `json:"SourcePrivateIp,omitempty" xml:"SourcePrivateIp,omitempty"`
	// The network type.
	//
	// Only private IP addresses are supported. Set the value to Private.
	//
	// This parameter is required.
	//
	// example:
	//
	// Private
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateNetworkRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkRuleRequest) SetDescription(v string) *CreateNetworkRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateNetworkRuleRequest) SetName(v string) *CreateNetworkRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateNetworkRuleRequest) SetSourcePrivateIp(v string) *CreateNetworkRuleRequest {
	s.SourcePrivateIp = &v
	return s
}

func (s *CreateNetworkRuleRequest) SetType(v string) *CreateNetworkRuleRequest {
	s.Type = &v
	return s
}

type CreateNetworkRuleResponseBody struct {
	// The ARN of the access control rule.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:119285303511****:network/networkrule_test
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The description.
	//
	// example:
	//
	// networkrule description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the access control rule.
	//
	// example:
	//
	// networkrule_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3bf02f7a-015b-4f93-be0f-cc043fda2dd3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The private IP address or private CIDR block.
	//
	// example:
	//
	// ["192.10.XX.XX","192.168.XX.XX/24"]
	SourcePrivateIp *string `json:"SourcePrivateIp,omitempty" xml:"SourcePrivateIp,omitempty"`
	// The network type.
	//
	// example:
	//
	// Private
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateNetworkRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkRuleResponseBody) SetArn(v string) *CreateNetworkRuleResponseBody {
	s.Arn = &v
	return s
}

func (s *CreateNetworkRuleResponseBody) SetDescription(v string) *CreateNetworkRuleResponseBody {
	s.Description = &v
	return s
}

func (s *CreateNetworkRuleResponseBody) SetName(v string) *CreateNetworkRuleResponseBody {
	s.Name = &v
	return s
}

func (s *CreateNetworkRuleResponseBody) SetRequestId(v string) *CreateNetworkRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkRuleResponseBody) SetSourcePrivateIp(v string) *CreateNetworkRuleResponseBody {
	s.SourcePrivateIp = &v
	return s
}

func (s *CreateNetworkRuleResponseBody) SetType(v string) *CreateNetworkRuleResponseBody {
	s.Type = &v
	return s
}

type CreateNetworkRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetworkRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetworkRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkRuleResponse) SetHeaders(v map[string]*string) *CreateNetworkRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkRuleResponse) SetStatusCode(v int32) *CreateNetworkRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetworkRuleResponse) SetBody(v *CreateNetworkRuleResponseBody) *CreateNetworkRuleResponse {
	s.Body = v
	return s
}

type CreatePolicyRequest struct {
	// The name of the access control rule.
	//
	// > For more information about how to query created access control rules, see [ListNetworkRules](https://help.aliyun.com/document_detail/2539433.html).
	//
	// example:
	//
	// {"NetworkRules":["kst-hzz62ee817bvyyr5x****.efkd","kst-hzz62ee817bvyyr5x****.eyyp"]}
	AccessControlRules *string `json:"AccessControlRules,omitempty" xml:"AccessControlRules,omitempty"`
	// The description.
	//
	// example:
	//
	// policy  description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The scope of the permission policy. You need to specify the KMS instance that you want to access.
	//
	// example:
	//
	// kst-hzz634e67d126u9p9****
	KmsInstance *string `json:"KmsInstance,omitempty" xml:"KmsInstance,omitempty"`
	// The name of the permission policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// policy_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operations that can be performed. Valid values:
	//
	// 	- RbacPermission/Template/CryptoServiceKeyUser: allows you to perform cryptographic operations.
	//
	// 	- RbacPermission/Template/CryptoServiceSecretUser: allows you to perform secret-related operations.
	//
	// You can select both.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["RbacPermission/Template/CryptoServiceKeyUser", "RbacPermission/Template/CryptoServiceSecretUser"]
	Permissions *string `json:"Permissions,omitempty" xml:"Permissions,omitempty"`
	// The key and secret that are allowed to access.
	//
	// 	- Key: Enter a key in the `key/${KeyId}` format. To allow access to all keys of a KMS instance, enter key/\\*.
	//
	// 	- Secret: Enter a secret in the `secret/${SecretName}` format. To allow access to all secrets of a KMS instance, enter secret/\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["secret/acs/ram/user/ram-secret", "secret/acs/ram/user/acr-master", "key/key-hzz63d9c8d3dfv8cv****"]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s CreatePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequest) SetAccessControlRules(v string) *CreatePolicyRequest {
	s.AccessControlRules = &v
	return s
}

func (s *CreatePolicyRequest) SetDescription(v string) *CreatePolicyRequest {
	s.Description = &v
	return s
}

func (s *CreatePolicyRequest) SetKmsInstance(v string) *CreatePolicyRequest {
	s.KmsInstance = &v
	return s
}

func (s *CreatePolicyRequest) SetName(v string) *CreatePolicyRequest {
	s.Name = &v
	return s
}

func (s *CreatePolicyRequest) SetPermissions(v string) *CreatePolicyRequest {
	s.Permissions = &v
	return s
}

func (s *CreatePolicyRequest) SetResources(v string) *CreatePolicyRequest {
	s.Resources = &v
	return s
}

type CreatePolicyResponseBody struct {
	// The name of the access control rule.
	//
	// example:
	//
	// {"NetworkRules":["kst-hzz62ee817bvyyr5x****.efkd","kst-hzz62ee817bvyyr5x****.eyyp"]}
	AccessControlRules *string `json:"AccessControlRules,omitempty" xml:"AccessControlRules,omitempty"`
	// The ARN of the permission policy.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:119285303511****:policy/policy_test
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The description.
	//
	// example:
	//
	// policy  description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The scope of the permission policy.
	//
	// example:
	//
	// kst-hzz634e67d126u9p9****
	KmsInstance *string `json:"KmsInstance,omitempty" xml:"KmsInstance,omitempty"`
	// The name of the permission policy.
	//
	// example:
	//
	// policy_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operations that can be performed.
	//
	// example:
	//
	// ["RbacPermission/Template/CryptoServiceKeyUser", "RbacPermission/Template/CryptoServiceSecretUser"]
	Permissions *string `json:"Permissions,omitempty" xml:"Permissions,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3bf02f7a-015b-4f34-be0f-c4543fda2d33
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The key and secret that are allowed to access.
	//
	// 	- `key/*` indicates that all keys of the KMS instance can be accessed.
	//
	// 	- `secret/*` indicates all secrets of the KMS instance can be accessed.
	//
	// example:
	//
	// ["secret/acs/ram/user/ram-secret", "secret/acs/ram/user/acr-master", "key/key-hzz63d9c8d3dfv8cv****"]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s CreatePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponseBody) SetAccessControlRules(v string) *CreatePolicyResponseBody {
	s.AccessControlRules = &v
	return s
}

func (s *CreatePolicyResponseBody) SetArn(v string) *CreatePolicyResponseBody {
	s.Arn = &v
	return s
}

func (s *CreatePolicyResponseBody) SetDescription(v string) *CreatePolicyResponseBody {
	s.Description = &v
	return s
}

func (s *CreatePolicyResponseBody) SetKmsInstance(v string) *CreatePolicyResponseBody {
	s.KmsInstance = &v
	return s
}

func (s *CreatePolicyResponseBody) SetName(v string) *CreatePolicyResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePolicyResponseBody) SetPermissions(v string) *CreatePolicyResponseBody {
	s.Permissions = &v
	return s
}

func (s *CreatePolicyResponseBody) SetRequestId(v string) *CreatePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolicyResponseBody) SetResources(v string) *CreatePolicyResponseBody {
	s.Resources = &v
	return s
}

type CreatePolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyResponse) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponse) SetHeaders(v map[string]*string) *CreatePolicyResponse {
	s.Headers = v
	return s
}

func (s *CreatePolicyResponse) SetStatusCode(v int32) *CreatePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePolicyResponse) SetBody(v *CreatePolicyResponseBody) *CreatePolicyResponse {
	s.Body = v
	return s
}

type CreateSecretRequest struct {
	// The version number of the secret.
	//
	// example:
	//
	// kst-bjj62d8f5e0sgtx8h****
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// Specifies whether to enable automatic rotation. Valid values:
	//
	// 	- true: specifies to enable automatic rotation.
	//
	// 	- false: specifies to disable automatic rotation. This is the default value.
	//
	// >  This parameter is valid if you set the SecretType parameter to Rds, RAMCredentials, or ECS.
	//
	// example:
	//
	// mydbinfo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether automatic rotation is enabled. Valid values:
	//
	// 	- Enabled: indicates that automatic rotation is enabled.
	//
	// 	- Disabled: indicates that automatic rotation is disabled.
	//
	// 	- Invalid: indicates that the status of automatic rotation is abnormal. In this case, Secrets Manager cannot automatically rotate the secret.
	//
	// >  This parameter is returned if you set the SecretType parameter to Rds, RAMCredentials, or ECS.
	//
	// example:
	//
	// true
	EnableAutomaticRotation *bool `json:"EnableAutomaticRotation,omitempty" xml:"EnableAutomaticRotation,omitempty"`
	// The description of the secret.
	//
	// example:
	//
	// 00aa68af-2c02-4f68-95fe-3435d330****
	EncryptionKeyId *string `json:"EncryptionKeyId,omitempty" xml:"EncryptionKeyId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// {"SecretSubType":"SingleUser", "DBInstanceId":"rm-bp1b3dd3a506e****" ,"CustomData":{}}
	ExtendedConfig map[string]interface{} `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	Policy         *string                `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The name of the secret.
	//
	// example:
	//
	// 30d
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	// The tags of the secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// The type of the secret. Valid values:
	//
	// 	- Generic: specifies a generic secret.
	//
	// 	- Rds: specifies a managed ApsaraDB RDS secret.
	//
	// 	- RAMCredentials: specifies a managed RAM secret.
	//
	// 	- ECS: specifies a managed ECS secret.
	SecretData *string `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
	// The extended configuration of the secret. This parameter specifies the properties of the secret of the specific type. The description can be up to 1,024 characters in length.
	//
	// 	- If you set the SecretType parameter to Generic, you do not need to configure this parameter.
	//
	// 	- If you set the SecretType parameter to Rds, configure the following fields for the ExtendedConfig parameter:
	//
	//     	- SecretSubType: required. The subtype of the secret. Valid values:
	//
	//         	- SingleUser: Secrets Manager manages the ApsaraDB RDS secret in single-account mode. When the secret is rotated, the password of the specified account is reset to a new random password.
	//
	//         	- DoubleUsers: Secrets Manager manages the ApsaraDB RDS secret in dual-account mode. One account is referenced by the ACSCurrent version, and the other account is referenced by the ACSPrevious version. When the secret is rotated, the password of the account referenced by the ACSPrevious version is reset to a new random password. Then, Secrets Manager switches the referenced accounts between the ACSCurrent and ACSPrevious versions.
	//
	//     	- DBInstanceId: required. The ApsaraDB RDS instance to which the ApsaraDB RDS account belongs.
	//
	//     	- CustomData: optional. The custom data. The value is a collection of key-value pairs in the JSON format. Up to 10 key-value pairs can be specified. Separate multiple key-value pairs with commas (,). Example: `{"Key1": "v1", "fds":"fdsf"}`. The default value is a pair of empty braces (`{}`).
	//
	// 	- If you set the SecretType parameter to RAMCredentials, configure the following fields for the ExtendedConfig parameter:
	//
	//     	- SecretSubType: required. The subtype of the secret. Set the value to RamUserAccessKey.
	//
	//     	- UserName: required. The name of the RAM user.
	//
	//     	- CustomData: optional. The custom data. The value is a collection of key-value pairs in the JSON format. Up to 10 key-value pairs can be specified. Separate multiple key-value pairs with commas (,). The default value is a pair of empty braces (`{}`).
	//
	// 	- If you set the SecretType parameter to ECS, configure the following fields for the ExtendedConfig parameter:
	//
	//     	- SecretSubType: required. The subtype of the secret. Valid values:
	//
	//         	- Password: the password that is used to log on to the ECS instance.
	//
	//         	- SSHKey: the SSH public key and private key that are used to log on to the ECS instance.
	//
	//     	- RegionId: required. The ID of the region in which the ECS instance resides.
	//
	//     	- InstanceId: required. The ID of the ECS instance.
	//
	//     	- CustomData: optional. The custom data. The value is a collection of key-value pairs in the JSON format. Up to 10 key-value pairs can be specified. Separate multiple key-value pairs with commas (,). The default value is a pair of empty braces (`{}`).
	//
	// >  This parameter is required if you set the SecretType parameter to Rds, RAMCredentials, or ECS.
	//
	// example:
	//
	// text
	SecretDataType *string `json:"SecretDataType,omitempty" xml:"SecretDataType,omitempty"`
	// The value of the secret that you want to create. Secrets Manager encrypts the secret value and stores the encrypted value in the initial version.
	//
	// 	- If you set the SecretType parameter to Generic that indicates a generic secret, you can customize the secret value.
	//
	// 	- If you set the SecretType parameter to Rds that indicates a managed ApsaraDB RDS secret, the secret value must be in the format of `{"Accounts":[{"AccountName":"","AccountPassword":""}]}`. In the preceding format, `AccountName` indicates the username of the account that is used to connect to your ApsaraDB RDS instance, and `AccountPassword` specifies the password of the account.
	//
	// 	- If you set the SecretType parameter to RAMCredentials that indicates a managed RAM secret, the secret value must be in the format of `{"AccessKeys":[{"AccessKeyId":"","AccessKeySecret":"",}]}`. In the preceding format, `AccessKeyId` indicates the AccessKey ID of the RAM user and `AccessKeySecret` specifies the AccessKey secret of the RAM user. You must specify all the AccessKey pairs of the RAM user.
	//
	// 	- If you set the SecretType parameter to ECS that indicates a managed ECS secret, the secret value must be in one of the following formats:
	//
	//     	- `{"UserName":"","Password": ""}`: In the format, `UserName` specifies the username that is used to log on to the ECS instance, and `Password` specifies the password that is used to log on to the ECS instance.
	//
	//     	- `{"UserName":"","PublicKey": "", "PrivateKey": ""}`: In the format, `PublicKey` indicates the SSH public key that is used to log on to the ECS instance, and `PrivateKey` specifies the SSH private key that is used to log on to the ECS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mydbconninfo
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The ID of the dedicated KMS instance.
	//
	// example:
	//
	// Rds
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// The interval for automatic rotation. Valid values: 6 hours to 8,760 hours (365 days).
	//
	// The value is in the `integer[unit]` format.
	//
	// The unit can be d (day), h (hour), m (minute), or s (second). For example, both 7d and 604800s indicate a seven-day interval.
	//
	// >  This parameter is required if you set the EnableAutomaticRotation parameter to true. This parameter is ignored if you set the EnableAutomaticRotation parameter to false or if the EnableAutomaticRotation parameter is not configured.
	//
	// example:
	//
	// [{\\"TagKey\\":\\"key1\\",\\"TagValue\\":\\"val1\\"},{\\"TagKey\\":\\"key2\\",\\"TagValue\\":\\"val2\\"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the secret value. Valid values:
	//
	// 	- text
	//
	// 	- binary
	//
	// >  If you set the SecretType parameter to Rds, RAMCredentials, or ECS, the SecretDataType parameter must be set to text.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreateSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSecretRequest) GoString() string {
	return s.String()
}

func (s *CreateSecretRequest) SetDKMSInstanceId(v string) *CreateSecretRequest {
	s.DKMSInstanceId = &v
	return s
}

func (s *CreateSecretRequest) SetDescription(v string) *CreateSecretRequest {
	s.Description = &v
	return s
}

func (s *CreateSecretRequest) SetEnableAutomaticRotation(v bool) *CreateSecretRequest {
	s.EnableAutomaticRotation = &v
	return s
}

func (s *CreateSecretRequest) SetEncryptionKeyId(v string) *CreateSecretRequest {
	s.EncryptionKeyId = &v
	return s
}

func (s *CreateSecretRequest) SetExtendedConfig(v map[string]interface{}) *CreateSecretRequest {
	s.ExtendedConfig = v
	return s
}

func (s *CreateSecretRequest) SetPolicy(v string) *CreateSecretRequest {
	s.Policy = &v
	return s
}

func (s *CreateSecretRequest) SetRotationInterval(v string) *CreateSecretRequest {
	s.RotationInterval = &v
	return s
}

func (s *CreateSecretRequest) SetSecretData(v string) *CreateSecretRequest {
	s.SecretData = &v
	return s
}

func (s *CreateSecretRequest) SetSecretDataType(v string) *CreateSecretRequest {
	s.SecretDataType = &v
	return s
}

func (s *CreateSecretRequest) SetSecretName(v string) *CreateSecretRequest {
	s.SecretName = &v
	return s
}

func (s *CreateSecretRequest) SetSecretType(v string) *CreateSecretRequest {
	s.SecretType = &v
	return s
}

func (s *CreateSecretRequest) SetTags(v string) *CreateSecretRequest {
	s.Tags = &v
	return s
}

func (s *CreateSecretRequest) SetVersionId(v string) *CreateSecretRequest {
	s.VersionId = &v
	return s
}

type CreateSecretShrinkRequest struct {
	// The version number of the secret.
	//
	// example:
	//
	// kst-bjj62d8f5e0sgtx8h****
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// Specifies whether to enable automatic rotation. Valid values:
	//
	// 	- true: specifies to enable automatic rotation.
	//
	// 	- false: specifies to disable automatic rotation. This is the default value.
	//
	// >  This parameter is valid if you set the SecretType parameter to Rds, RAMCredentials, or ECS.
	//
	// example:
	//
	// mydbinfo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether automatic rotation is enabled. Valid values:
	//
	// 	- Enabled: indicates that automatic rotation is enabled.
	//
	// 	- Disabled: indicates that automatic rotation is disabled.
	//
	// 	- Invalid: indicates that the status of automatic rotation is abnormal. In this case, Secrets Manager cannot automatically rotate the secret.
	//
	// >  This parameter is returned if you set the SecretType parameter to Rds, RAMCredentials, or ECS.
	//
	// example:
	//
	// true
	EnableAutomaticRotation *bool `json:"EnableAutomaticRotation,omitempty" xml:"EnableAutomaticRotation,omitempty"`
	// The description of the secret.
	//
	// example:
	//
	// 00aa68af-2c02-4f68-95fe-3435d330****
	EncryptionKeyId *string `json:"EncryptionKeyId,omitempty" xml:"EncryptionKeyId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// {"SecretSubType":"SingleUser", "DBInstanceId":"rm-bp1b3dd3a506e****" ,"CustomData":{}}
	ExtendedConfigShrink *string `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	Policy               *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The name of the secret.
	//
	// example:
	//
	// 30d
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	// The tags of the secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// The type of the secret. Valid values:
	//
	// 	- Generic: specifies a generic secret.
	//
	// 	- Rds: specifies a managed ApsaraDB RDS secret.
	//
	// 	- RAMCredentials: specifies a managed RAM secret.
	//
	// 	- ECS: specifies a managed ECS secret.
	SecretData *string `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
	// The extended configuration of the secret. This parameter specifies the properties of the secret of the specific type. The description can be up to 1,024 characters in length.
	//
	// 	- If you set the SecretType parameter to Generic, you do not need to configure this parameter.
	//
	// 	- If you set the SecretType parameter to Rds, configure the following fields for the ExtendedConfig parameter:
	//
	//     	- SecretSubType: required. The subtype of the secret. Valid values:
	//
	//         	- SingleUser: Secrets Manager manages the ApsaraDB RDS secret in single-account mode. When the secret is rotated, the password of the specified account is reset to a new random password.
	//
	//         	- DoubleUsers: Secrets Manager manages the ApsaraDB RDS secret in dual-account mode. One account is referenced by the ACSCurrent version, and the other account is referenced by the ACSPrevious version. When the secret is rotated, the password of the account referenced by the ACSPrevious version is reset to a new random password. Then, Secrets Manager switches the referenced accounts between the ACSCurrent and ACSPrevious versions.
	//
	//     	- DBInstanceId: required. The ApsaraDB RDS instance to which the ApsaraDB RDS account belongs.
	//
	//     	- CustomData: optional. The custom data. The value is a collection of key-value pairs in the JSON format. Up to 10 key-value pairs can be specified. Separate multiple key-value pairs with commas (,). Example: `{"Key1": "v1", "fds":"fdsf"}`. The default value is a pair of empty braces (`{}`).
	//
	// 	- If you set the SecretType parameter to RAMCredentials, configure the following fields for the ExtendedConfig parameter:
	//
	//     	- SecretSubType: required. The subtype of the secret. Set the value to RamUserAccessKey.
	//
	//     	- UserName: required. The name of the RAM user.
	//
	//     	- CustomData: optional. The custom data. The value is a collection of key-value pairs in the JSON format. Up to 10 key-value pairs can be specified. Separate multiple key-value pairs with commas (,). The default value is a pair of empty braces (`{}`).
	//
	// 	- If you set the SecretType parameter to ECS, configure the following fields for the ExtendedConfig parameter:
	//
	//     	- SecretSubType: required. The subtype of the secret. Valid values:
	//
	//         	- Password: the password that is used to log on to the ECS instance.
	//
	//         	- SSHKey: the SSH public key and private key that are used to log on to the ECS instance.
	//
	//     	- RegionId: required. The ID of the region in which the ECS instance resides.
	//
	//     	- InstanceId: required. The ID of the ECS instance.
	//
	//     	- CustomData: optional. The custom data. The value is a collection of key-value pairs in the JSON format. Up to 10 key-value pairs can be specified. Separate multiple key-value pairs with commas (,). The default value is a pair of empty braces (`{}`).
	//
	// >  This parameter is required if you set the SecretType parameter to Rds, RAMCredentials, or ECS.
	//
	// example:
	//
	// text
	SecretDataType *string `json:"SecretDataType,omitempty" xml:"SecretDataType,omitempty"`
	// The value of the secret that you want to create. Secrets Manager encrypts the secret value and stores the encrypted value in the initial version.
	//
	// 	- If you set the SecretType parameter to Generic that indicates a generic secret, you can customize the secret value.
	//
	// 	- If you set the SecretType parameter to Rds that indicates a managed ApsaraDB RDS secret, the secret value must be in the format of `{"Accounts":[{"AccountName":"","AccountPassword":""}]}`. In the preceding format, `AccountName` indicates the username of the account that is used to connect to your ApsaraDB RDS instance, and `AccountPassword` specifies the password of the account.
	//
	// 	- If you set the SecretType parameter to RAMCredentials that indicates a managed RAM secret, the secret value must be in the format of `{"AccessKeys":[{"AccessKeyId":"","AccessKeySecret":"",}]}`. In the preceding format, `AccessKeyId` indicates the AccessKey ID of the RAM user and `AccessKeySecret` specifies the AccessKey secret of the RAM user. You must specify all the AccessKey pairs of the RAM user.
	//
	// 	- If you set the SecretType parameter to ECS that indicates a managed ECS secret, the secret value must be in one of the following formats:
	//
	//     	- `{"UserName":"","Password": ""}`: In the format, `UserName` specifies the username that is used to log on to the ECS instance, and `Password` specifies the password that is used to log on to the ECS instance.
	//
	//     	- `{"UserName":"","PublicKey": "", "PrivateKey": ""}`: In the format, `PublicKey` indicates the SSH public key that is used to log on to the ECS instance, and `PrivateKey` specifies the SSH private key that is used to log on to the ECS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mydbconninfo
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The ID of the dedicated KMS instance.
	//
	// example:
	//
	// Rds
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// The interval for automatic rotation. Valid values: 6 hours to 8,760 hours (365 days).
	//
	// The value is in the `integer[unit]` format.
	//
	// The unit can be d (day), h (hour), m (minute), or s (second). For example, both 7d and 604800s indicate a seven-day interval.
	//
	// >  This parameter is required if you set the EnableAutomaticRotation parameter to true. This parameter is ignored if you set the EnableAutomaticRotation parameter to false or if the EnableAutomaticRotation parameter is not configured.
	//
	// example:
	//
	// [{\\"TagKey\\":\\"key1\\",\\"TagValue\\":\\"val1\\"},{\\"TagKey\\":\\"key2\\",\\"TagValue\\":\\"val2\\"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the secret value. Valid values:
	//
	// 	- text
	//
	// 	- binary
	//
	// >  If you set the SecretType parameter to Rds, RAMCredentials, or ECS, the SecretDataType parameter must be set to text.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreateSecretShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSecretShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSecretShrinkRequest) SetDKMSInstanceId(v string) *CreateSecretShrinkRequest {
	s.DKMSInstanceId = &v
	return s
}

func (s *CreateSecretShrinkRequest) SetDescription(v string) *CreateSecretShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateSecretShrinkRequest) SetEnableAutomaticRotation(v bool) *CreateSecretShrinkRequest {
	s.EnableAutomaticRotation = &v
	return s
}

func (s *CreateSecretShrinkRequest) SetEncryptionKeyId(v string) *CreateSecretShrinkRequest {
	s.EncryptionKeyId = &v
	return s
}

func (s *CreateSecretShrinkRequest) SetExtendedConfigShrink(v string) *CreateSecretShrinkRequest {
	s.ExtendedConfigShrink = &v
	return s
}

func (s *CreateSecretShrinkRequest) SetPolicy(v string) *CreateSecretShrinkRequest {
	s.Policy = &v
	return s
}

func (s *CreateSecretShrinkRequest) SetRotationInterval(v string) *CreateSecretShrinkRequest {
	s.RotationInterval = &v
	return s
}

func (s *CreateSecretShrinkRequest) SetSecretData(v string) *CreateSecretShrinkRequest {
	s.SecretData = &v
	return s
}

func (s *CreateSecretShrinkRequest) SetSecretDataType(v string) *CreateSecretShrinkRequest {
	s.SecretDataType = &v
	return s
}

func (s *CreateSecretShrinkRequest) SetSecretName(v string) *CreateSecretShrinkRequest {
	s.SecretName = &v
	return s
}

func (s *CreateSecretShrinkRequest) SetSecretType(v string) *CreateSecretShrinkRequest {
	s.SecretType = &v
	return s
}

func (s *CreateSecretShrinkRequest) SetTags(v string) *CreateSecretShrinkRequest {
	s.Tags = &v
	return s
}

func (s *CreateSecretShrinkRequest) SetVersionId(v string) *CreateSecretShrinkRequest {
	s.VersionId = &v
	return s
}

type CreateSecretResponseBody struct {
	// example:
	//
	// acs:kms:cn-hangzhou:154035569884****:secret/mydbconninfo
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The type of the secret. Valid values:
	//
	// 	- Generic: indicates a generic secret.
	//
	// 	- Rds: indicates a managed ApsaraDB RDS secret.
	//
	// 	- RAMCredentials: indicates a managed RAM secret.
	//
	// 	- ECS: indicates a managed ECS secret.
	//
	// example:
	//
	// Enabled
	AutomaticRotation *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	// example:
	//
	// kst-bjj62d8f5e0sgtx8h****
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// example:
	//
	// {\\"SecretSubType\\":\\"SingleUser\\", \\"DBInstanceId\\":\\"rm-uf667446pc955****\\",  \\"CustomData\\":{} }
	ExtendedConfig *string `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	// The extended configuration of the secret.
	//
	// >  This parameter is returned if you set the SecretType parameter to Rds, RAMCredentials, or ECS.
	//
	// example:
	//
	// 2022-07-06T18:22:03Z
	NextRotationDate *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	// The time when the next rotation will be performed.
	//
	// >  This parameter is returned if automatic rotation is enabled.
	//
	// example:
	//
	// 3bf02f7a-015b-4f93-be0f-cc043fda2dd3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 604800s
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	// The interval for automatic rotation.
	//
	// The value is in the `integer[unit]` format. The value of the `unit` field is fixed as s. For example, if the value is 604800s, automatic rotation is performed at a 7-day interval.
	//
	// >  This parameter is returned if automatic rotation is enabled.
	//
	// example:
	//
	// mydbconninfo
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The ID of the dedicated KMS instance.
	//
	// example:
	//
	// Rds
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the secret.
	//
	// example:
	//
	// v1
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreateSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSecretResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecretResponseBody) SetArn(v string) *CreateSecretResponseBody {
	s.Arn = &v
	return s
}

func (s *CreateSecretResponseBody) SetAutomaticRotation(v string) *CreateSecretResponseBody {
	s.AutomaticRotation = &v
	return s
}

func (s *CreateSecretResponseBody) SetDKMSInstanceId(v string) *CreateSecretResponseBody {
	s.DKMSInstanceId = &v
	return s
}

func (s *CreateSecretResponseBody) SetExtendedConfig(v string) *CreateSecretResponseBody {
	s.ExtendedConfig = &v
	return s
}

func (s *CreateSecretResponseBody) SetNextRotationDate(v string) *CreateSecretResponseBody {
	s.NextRotationDate = &v
	return s
}

func (s *CreateSecretResponseBody) SetRequestId(v string) *CreateSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecretResponseBody) SetRotationInterval(v string) *CreateSecretResponseBody {
	s.RotationInterval = &v
	return s
}

func (s *CreateSecretResponseBody) SetSecretName(v string) *CreateSecretResponseBody {
	s.SecretName = &v
	return s
}

func (s *CreateSecretResponseBody) SetSecretType(v string) *CreateSecretResponseBody {
	s.SecretType = &v
	return s
}

func (s *CreateSecretResponseBody) SetVersionId(v string) *CreateSecretResponseBody {
	s.VersionId = &v
	return s
}

type CreateSecretResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSecretResponse) GoString() string {
	return s.String()
}

func (s *CreateSecretResponse) SetHeaders(v map[string]*string) *CreateSecretResponse {
	s.Headers = v
	return s
}

func (s *CreateSecretResponse) SetStatusCode(v int32) *CreateSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecretResponse) SetBody(v *CreateSecretResponseBody) *CreateSecretResponse {
	s.Body = v
	return s
}

type DecryptRequest struct {
	// The ciphertext that you want to decrypt.
	//
	// You can generate the ciphertext by calling the following operations:
	//
	// 	- [GenerateDataKey](https://help.aliyun.com/document_detail/28948.html)
	//
	// 	- [Encrypt](https://help.aliyun.com/document_detail/28949.html)
	//
	// 	- [GenerateDataKeyWithoutPlaintext](https://help.aliyun.com/document_detail/134043.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// DZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmaaSl+TztSIMe43nbTH/Z1Wr4XfLftKhAciUmDQXuMRl4WTvKhxjMThjK****
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The JSON string that consists of key-value pairs.
	//
	// >  If you specify the EncryptionContext parameter when you call the [GenerateDataKey](https://help.aliyun.com/document_detail/28948.html), [Encrypt](https://help.aliyun.com/document_detail/28949.html), or [GenerateDataKeyWithoutPlaintext](https://help.aliyun.com/document_detail/134043.html) operation, you must specify the same context when you call the Decrypt operation. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// example:
	//
	// {"Example":"Example"}
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
}

func (s DecryptRequest) String() string {
	return tea.Prettify(s)
}

func (s DecryptRequest) GoString() string {
	return s.String()
}

func (s *DecryptRequest) SetCiphertextBlob(v string) *DecryptRequest {
	s.CiphertextBlob = &v
	return s
}

func (s *DecryptRequest) SetEncryptionContext(v map[string]interface{}) *DecryptRequest {
	s.EncryptionContext = v
	return s
}

type DecryptShrinkRequest struct {
	// The ciphertext that you want to decrypt.
	//
	// You can generate the ciphertext by calling the following operations:
	//
	// 	- [GenerateDataKey](https://help.aliyun.com/document_detail/28948.html)
	//
	// 	- [Encrypt](https://help.aliyun.com/document_detail/28949.html)
	//
	// 	- [GenerateDataKeyWithoutPlaintext](https://help.aliyun.com/document_detail/134043.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// DZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmaaSl+TztSIMe43nbTH/Z1Wr4XfLftKhAciUmDQXuMRl4WTvKhxjMThjK****
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The JSON string that consists of key-value pairs.
	//
	// >  If you specify the EncryptionContext parameter when you call the [GenerateDataKey](https://help.aliyun.com/document_detail/28948.html), [Encrypt](https://help.aliyun.com/document_detail/28949.html), or [GenerateDataKeyWithoutPlaintext](https://help.aliyun.com/document_detail/134043.html) operation, you must specify the same context when you call the Decrypt operation. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// example:
	//
	// {"Example":"Example"}
	EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
}

func (s DecryptShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DecryptShrinkRequest) GoString() string {
	return s.String()
}

func (s *DecryptShrinkRequest) SetCiphertextBlob(v string) *DecryptShrinkRequest {
	s.CiphertextBlob = &v
	return s
}

func (s *DecryptShrinkRequest) SetEncryptionContextShrink(v string) *DecryptShrinkRequest {
	s.EncryptionContextShrink = &v
	return s
}

type DecryptResponseBody struct {
	// The ID of the customer master key (CMK) that is used to decrypt the ciphertext.
	//
	// It is the GUID of the CMK.
	//
	// example:
	//
	// 202b9877-5a25-46e3-a763-e20791b5****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the CMK version that is used to decrypt the ciphertext.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The plaintext that is generated after decryption.
	//
	// example:
	//
	// tRYXuCwgja12xxO1N/gZERDDCLw9doZEQiPDk/Bv****
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 207596a2-36d3-4840-b1bd-f87044699bd7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DecryptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DecryptResponseBody) GoString() string {
	return s.String()
}

func (s *DecryptResponseBody) SetKeyId(v string) *DecryptResponseBody {
	s.KeyId = &v
	return s
}

func (s *DecryptResponseBody) SetKeyVersionId(v string) *DecryptResponseBody {
	s.KeyVersionId = &v
	return s
}

func (s *DecryptResponseBody) SetPlaintext(v string) *DecryptResponseBody {
	s.Plaintext = &v
	return s
}

func (s *DecryptResponseBody) SetRequestId(v string) *DecryptResponseBody {
	s.RequestId = &v
	return s
}

type DecryptResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DecryptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DecryptResponse) String() string {
	return tea.Prettify(s)
}

func (s DecryptResponse) GoString() string {
	return s.String()
}

func (s *DecryptResponse) SetHeaders(v map[string]*string) *DecryptResponse {
	s.Headers = v
	return s
}

func (s *DecryptResponse) SetStatusCode(v int32) *DecryptResponse {
	s.StatusCode = &v
	return s
}

func (s *DecryptResponse) SetBody(v *DecryptResponseBody) *DecryptResponse {
	s.Body = v
	return s
}

type DeleteAliasRequest struct {
	// The alias that you want to delete.
	//
	// The value must be 1 to 255 characters in length and must include the alias/ prefix.
	//
	// This parameter is required.
	//
	// example:
	//
	// alias/example
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
}

func (s DeleteAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAliasRequest) GoString() string {
	return s.String()
}

func (s *DeleteAliasRequest) SetAliasName(v string) *DeleteAliasRequest {
	s.AliasName = &v
	return s
}

type DeleteAliasResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4c8ae23f-3a42-6791-a4ba-1faa77831c28
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAliasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAliasResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAliasResponseBody) SetRequestId(v string) *DeleteAliasResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAliasResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAliasResponse) GoString() string {
	return s.String()
}

func (s *DeleteAliasResponse) SetHeaders(v map[string]*string) *DeleteAliasResponse {
	s.Headers = v
	return s
}

func (s *DeleteAliasResponse) SetStatusCode(v int32) *DeleteAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAliasResponse) SetBody(v *DeleteAliasResponseBody) *DeleteAliasResponse {
	s.Body = v
	return s
}

type DeleteApplicationAccessPointRequest struct {
	// The name of the AAP that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// aap_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteApplicationAccessPointRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationAccessPointRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationAccessPointRequest) SetName(v string) *DeleteApplicationAccessPointRequest {
	s.Name = &v
	return s
}

type DeleteApplicationAccessPointResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// bcfefe15-46f0-44a3-bd96-3d422474b71a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationAccessPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationAccessPointResponseBody) SetRequestId(v string) *DeleteApplicationAccessPointResponseBody {
	s.RequestId = &v
	return s
}

type DeleteApplicationAccessPointResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApplicationAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApplicationAccessPointResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationAccessPointResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationAccessPointResponse) SetHeaders(v map[string]*string) *DeleteApplicationAccessPointResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationAccessPointResponse) SetStatusCode(v int32) *DeleteApplicationAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationAccessPointResponse) SetBody(v *DeleteApplicationAccessPointResponseBody) *DeleteApplicationAccessPointResponse {
	s.Body = v
	return s
}

type DeleteCertificateRequest struct {
	// The ID of the certificate. It is the globally unique identifier (GUID) of the certificate in Alibaba Cloud Certificate Manager.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9a28de48-8d8b-484d-a766-dec4****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s DeleteCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteCertificateRequest) SetCertificateId(v string) *DeleteCertificateRequest {
	s.CertificateId = &v
	return s
}

type DeleteCertificateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// d97f6c33-ca26-4de2-a580-0e2fd1c5bfb0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCertificateResponseBody) SetRequestId(v string) *DeleteCertificateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCertificateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteCertificateResponse) SetHeaders(v map[string]*string) *DeleteCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteCertificateResponse) SetStatusCode(v int32) *DeleteCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCertificateResponse) SetBody(v *DeleteCertificateResponseBody) *DeleteCertificateResponse {
	s.Body = v
	return s
}

type DeleteClientKeyRequest struct {
	// The ID of the client key.
	//
	// This parameter is required.
	//
	// example:
	//
	// KAAP.66abf237-63f6-4625-b8cf-47e1086e****
	ClientKeyId *string `json:"ClientKeyId,omitempty" xml:"ClientKeyId,omitempty"`
}

func (s DeleteClientKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClientKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteClientKeyRequest) SetClientKeyId(v string) *DeleteClientKeyRequest {
	s.ClientKeyId = &v
	return s
}

type DeleteClientKeyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 2312e45f-b2fa-4c34-ad94-3eca50932916
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClientKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClientKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClientKeyResponseBody) SetRequestId(v string) *DeleteClientKeyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteClientKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClientKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClientKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClientKeyResponse) GoString() string {
	return s.String()
}

func (s *DeleteClientKeyResponse) SetHeaders(v map[string]*string) *DeleteClientKeyResponse {
	s.Headers = v
	return s
}

func (s *DeleteClientKeyResponse) SetStatusCode(v int32) *DeleteClientKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClientKeyResponse) SetBody(v *DeleteClientKeyResponseBody) *DeleteClientKeyResponse {
	s.Body = v
	return s
}

type DeleteKeyMaterialRequest struct {
	// The globally unique ID of the CMK.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s DeleteKeyMaterialRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyMaterialRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeyMaterialRequest) SetKeyId(v string) *DeleteKeyMaterialRequest {
	s.KeyId = &v
	return s
}

type DeleteKeyMaterialResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4162a6af-bc99-40b3-a552-89dcc8aaf7c8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteKeyMaterialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKeyMaterialResponseBody) SetRequestId(v string) *DeleteKeyMaterialResponseBody {
	s.RequestId = &v
	return s
}

type DeleteKeyMaterialResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKeyMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKeyMaterialResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyMaterialResponse) GoString() string {
	return s.String()
}

func (s *DeleteKeyMaterialResponse) SetHeaders(v map[string]*string) *DeleteKeyMaterialResponse {
	s.Headers = v
	return s
}

func (s *DeleteKeyMaterialResponse) SetStatusCode(v int32) *DeleteKeyMaterialResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKeyMaterialResponse) SetBody(v *DeleteKeyMaterialResponseBody) *DeleteKeyMaterialResponse {
	s.Body = v
	return s
}

type DeleteNetworkRuleRequest struct {
	// The name of the network access rule that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// networkrule_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteNetworkRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkRuleRequest) SetName(v string) *DeleteNetworkRuleRequest {
	s.Name = &v
	return s
}

type DeleteNetworkRuleResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3bf02f7a-015b-4f93-be0f-cc043fda2d4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkRuleResponseBody) SetRequestId(v string) *DeleteNetworkRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNetworkRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNetworkRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNetworkRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkRuleResponse) SetHeaders(v map[string]*string) *DeleteNetworkRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkRuleResponse) SetStatusCode(v int32) *DeleteNetworkRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNetworkRuleResponse) SetBody(v *DeleteNetworkRuleResponseBody) *DeleteNetworkRuleResponse {
	s.Body = v
	return s
}

type DeletePolicyRequest struct {
	// The name of the permission policy that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// policy_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeletePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyRequest) SetName(v string) *DeletePolicyRequest {
	s.Name = &v
	return s
}

type DeletePolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 00a26a33-d992-42f3-9012-5fd12764430f
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolicyResponseBody) SetRequestId(v string) *DeletePolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeletePolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyResponse) SetHeaders(v map[string]*string) *DeletePolicyResponse {
	s.Headers = v
	return s
}

func (s *DeletePolicyResponse) SetStatusCode(v int32) *DeletePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolicyResponse) SetBody(v *DeletePolicyResponseBody) *DeletePolicyResponse {
	s.Body = v
	return s
}

type DeleteSecretRequest struct {
	// Specifies whether to forcibly delete the secret. If this parameter is set to true, the secret cannot be recovered.
	//
	// Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default value)
	//
	// example:
	//
	// false
	ForceDeleteWithoutRecovery *string `json:"ForceDeleteWithoutRecovery,omitempty" xml:"ForceDeleteWithoutRecovery,omitempty"`
	// Specifies the recovery period of the secret if you do not forcibly delete it. Default value: 30. Unit: Days.
	//
	// example:
	//
	// 10
	RecoveryWindowInDays *string `json:"RecoveryWindowInDays,omitempty" xml:"RecoveryWindowInDays,omitempty"`
	// The name of the secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s DeleteSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecretRequest) SetForceDeleteWithoutRecovery(v string) *DeleteSecretRequest {
	s.ForceDeleteWithoutRecovery = &v
	return s
}

func (s *DeleteSecretRequest) SetRecoveryWindowInDays(v string) *DeleteSecretRequest {
	s.RecoveryWindowInDays = &v
	return s
}

func (s *DeleteSecretRequest) SetSecretName(v string) *DeleteSecretRequest {
	s.SecretName = &v
	return s
}

type DeleteSecretResponseBody struct {
	// The time when the secret is scheduled to be deleted.
	//
	// example:
	//
	// 2022-09-15T07:02:14Z
	PlannedDeleteTime *string `json:"PlannedDeleteTime,omitempty" xml:"PlannedDeleteTime,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 38bbed2a-15e0-45ad-98d4-816ad2ccf4ea
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the secret.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s DeleteSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecretResponseBody) SetPlannedDeleteTime(v string) *DeleteSecretResponseBody {
	s.PlannedDeleteTime = &v
	return s
}

func (s *DeleteSecretResponseBody) SetRequestId(v string) *DeleteSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecretResponseBody) SetSecretName(v string) *DeleteSecretResponseBody {
	s.SecretName = &v
	return s
}

type DeleteSecretResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecretResponse) SetHeaders(v map[string]*string) *DeleteSecretResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecretResponse) SetStatusCode(v int32) *DeleteSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecretResponse) SetBody(v *DeleteSecretResponseBody) *DeleteSecretResponse {
	s.Body = v
	return s
}

type DescribeAccountKmsStatusResponseBody struct {
	// The status of KMS within your Alibaba cloud account. Valid values:
	//
	// 	- Enabled: KMS is enabled.
	//
	// 	- NotEnabled: KMS is disabled.
	//
	// 	- InDebt: Your account is overdue, and KMS stops providing services.
	//
	// > If your Alibaba Cloud account is overdue, top up your account at the earliest opportunity to avoid impacts on your services.
	//
	// 	- Suspended: KMS is suspended.
	//
	// example:
	//
	// Enabled
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3ac84333-d64d-4784-a8bc-997834a7ac6c
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccountKmsStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountKmsStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountKmsStatusResponseBody) SetAccountStatus(v string) *DescribeAccountKmsStatusResponseBody {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAccountKmsStatusResponseBody) SetRequestId(v string) *DescribeAccountKmsStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAccountKmsStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountKmsStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccountKmsStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountKmsStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountKmsStatusResponse) SetHeaders(v map[string]*string) *DescribeAccountKmsStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountKmsStatusResponse) SetStatusCode(v int32) *DescribeAccountKmsStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountKmsStatusResponse) SetBody(v *DescribeAccountKmsStatusResponseBody) *DescribeAccountKmsStatusResponse {
	s.Body = v
	return s
}

type DescribeApplicationAccessPointRequest struct {
	// The name of the AAP that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// aap_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeApplicationAccessPointRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationAccessPointRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAccessPointRequest) SetName(v string) *DescribeApplicationAccessPointRequest {
	s.Name = &v
	return s
}

type DescribeApplicationAccessPointResponseBody struct {
	// The ARN of the AAP.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:119285303511****:applicationaccesspoint/aap_test
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The authentication method.
	//
	// example:
	//
	// ClientKey
	AuthenticationMethod *string `json:"AuthenticationMethod,omitempty" xml:"AuthenticationMethod,omitempty"`
	// The description.
	//
	// example:
	//
	// aap description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the AAP.
	//
	// example:
	//
	// aap_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The permission policy that is bound to the AAP.
	//
	// example:
	//
	// ["kst-hzz62ee817bvyyr5x****.efkd","kst-hzz62ee817bvyyr5x****.eyyp"]
	Policies *string `json:"Policies,omitempty" xml:"Policies,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// bcfefe15-46f0-44a3-bd96-3d422474b71a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApplicationAccessPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAccessPointResponseBody) SetArn(v string) *DescribeApplicationAccessPointResponseBody {
	s.Arn = &v
	return s
}

func (s *DescribeApplicationAccessPointResponseBody) SetAuthenticationMethod(v string) *DescribeApplicationAccessPointResponseBody {
	s.AuthenticationMethod = &v
	return s
}

func (s *DescribeApplicationAccessPointResponseBody) SetDescription(v string) *DescribeApplicationAccessPointResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeApplicationAccessPointResponseBody) SetName(v string) *DescribeApplicationAccessPointResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeApplicationAccessPointResponseBody) SetPolicies(v string) *DescribeApplicationAccessPointResponseBody {
	s.Policies = &v
	return s
}

func (s *DescribeApplicationAccessPointResponseBody) SetRequestId(v string) *DescribeApplicationAccessPointResponseBody {
	s.RequestId = &v
	return s
}

type DescribeApplicationAccessPointResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationAccessPointResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationAccessPointResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAccessPointResponse) SetHeaders(v map[string]*string) *DescribeApplicationAccessPointResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationAccessPointResponse) SetStatusCode(v int32) *DescribeApplicationAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationAccessPointResponse) SetBody(v *DescribeApplicationAccessPointResponseBody) *DescribeApplicationAccessPointResponse {
	s.Body = v
	return s
}

type DescribeCertificateRequest struct {
	// The ID of the certificate. The ID must be globally unique in Certificates Manager.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9a28de48-8d8b-484d-a766-dec4****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s DescribeCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificateRequest) GoString() string {
	return s.String()
}

func (s *DescribeCertificateRequest) SetCertificateId(v string) *DescribeCertificateRequest {
	s.CertificateId = &v
	return s
}

type DescribeCertificateResponseBody struct {
	// The Alibaba Cloud Resource Name (ARN) of the certificate.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:159498693826****:certificate/9a28de48-8d8b-484d-a766-dec4****"
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the certificate. The ID must be globally unique in Certificates Manager.
	//
	// example:
	//
	// 9a28de48-8d8b-484d-a766-dec4****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The time when the certificate was created.
	//
	// example:
	//
	// 2020-10-13T03:05:03Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// Indicates whether the private key of the certificate can be exported for use. Valid values:
	//
	// 	- true: The private key of the certificate can be exported for use. This is the default value.
	//
	// 	- false: The private key of the certificate cannot be exported for use.
	//
	// example:
	//
	// true
	ExportablePrivateKey *bool `json:"ExportablePrivateKey,omitempty" xml:"ExportablePrivateKey,omitempty"`
	// The certificate issuer in the distinguished name (DN) format.
	//
	// example:
	//
	// CN=testCA,OU=kms,O=aliyun,C=CN
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The type of the key.
	//
	// example:
	//
	// RSA_2048
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The end of the validity period of the certificate.
	//
	// example:
	//
	// 2022-10-13T03:09:00Z
	NotAfter *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The beginning of the validity period of the certificate.
	//
	// example:
	//
	// 2020-10-13T03:09:00Z
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// edb671a3-c5a1-4ebe-a1de-d748b640bdf2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the certificate.
	//
	// example:
	//
	// 12345678
	Serial *string `json:"Serial,omitempty" xml:"Serial,omitempty"`
	// The signature algorithm of the certificate. Valid values:
	//
	// 	- RSA2048-SHA256
	//
	// 	- ECDSA-SHA256
	//
	// 	- SM2-SM3
	//
	// example:
	//
	// ECDSA-SHA256
	SignatureAlgorithm *string `json:"SignatureAlgorithm,omitempty" xml:"SignatureAlgorithm,omitempty"`
	// The status of the certificate. Valid values:
	//
	// 	- PENDING: The certificate is to be imported.
	//
	// 	- ACTIVE: The certificate is enabled.
	//
	// 	- INACTIVE: The certificate is disabled.
	//
	// 	- REVOKED: The certificate is revoked.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The subject of the certificate, which is in the DN format.
	//
	// example:
	//
	// CN=userName,OU=aliyun,O=aliyun,C=CN
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// The alias of the certificate subject.
	//
	// A domain name list is supported. A maximum of 10 domain names are supported.
	SubjectAlternativeNames []*string `json:"SubjectAlternativeNames,omitempty" xml:"SubjectAlternativeNames,omitempty" type:"Repeated"`
	// The public key identifier of the certificate subject.
	//
	// example:
	//
	// 79 36 26 DE 9F F5 15 E3 56 DC ****
	SubjectKeyIdentifier *string `json:"SubjectKeyIdentifier,omitempty" xml:"SubjectKeyIdentifier,omitempty"`
	// The public key of the certificate.
	//
	// example:
	//
	// -----BEGIN PUBLIC KEY----- MIIBIjA -----END PUBLIC KEY-----
	SubjectPublicKey *string `json:"SubjectPublicKey,omitempty" xml:"SubjectPublicKey,omitempty"`
	// The tag of the certificate.
	//
	// example:
	//
	// [{\\"TagKey\\":\\"S1key1\\",\\"TagValue\\":\\"S1val1\\"},{\\"TagKey\\":\\"S1key2\\",\\"TagValue\\":\\"S2val2\\"}]
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The time when the certificate was updated.
	//
	// example:
	//
	// 2020-12-23T06:10:13Z
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s DescribeCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCertificateResponseBody) SetArn(v string) *DescribeCertificateResponseBody {
	s.Arn = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetCertificateId(v string) *DescribeCertificateResponseBody {
	s.CertificateId = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetCreatedAt(v string) *DescribeCertificateResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetExportablePrivateKey(v bool) *DescribeCertificateResponseBody {
	s.ExportablePrivateKey = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetIssuer(v string) *DescribeCertificateResponseBody {
	s.Issuer = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetKeySpec(v string) *DescribeCertificateResponseBody {
	s.KeySpec = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetNotAfter(v string) *DescribeCertificateResponseBody {
	s.NotAfter = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetNotBefore(v string) *DescribeCertificateResponseBody {
	s.NotBefore = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetRequestId(v string) *DescribeCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetSerial(v string) *DescribeCertificateResponseBody {
	s.Serial = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetSignatureAlgorithm(v string) *DescribeCertificateResponseBody {
	s.SignatureAlgorithm = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetStatus(v string) *DescribeCertificateResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetSubject(v string) *DescribeCertificateResponseBody {
	s.Subject = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetSubjectAlternativeNames(v []*string) *DescribeCertificateResponseBody {
	s.SubjectAlternativeNames = v
	return s
}

func (s *DescribeCertificateResponseBody) SetSubjectKeyIdentifier(v string) *DescribeCertificateResponseBody {
	s.SubjectKeyIdentifier = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetSubjectPublicKey(v string) *DescribeCertificateResponseBody {
	s.SubjectPublicKey = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetTags(v map[string]interface{}) *DescribeCertificateResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeCertificateResponseBody) SetUpdatedAt(v string) *DescribeCertificateResponseBody {
	s.UpdatedAt = &v
	return s
}

type DescribeCertificateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificateResponse) GoString() string {
	return s.String()
}

func (s *DescribeCertificateResponse) SetHeaders(v map[string]*string) *DescribeCertificateResponse {
	s.Headers = v
	return s
}

func (s *DescribeCertificateResponse) SetStatusCode(v int32) *DescribeCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCertificateResponse) SetBody(v *DescribeCertificateResponseBody) *DescribeCertificateResponse {
	s.Body = v
	return s
}

type DescribeKeyRequest struct {
	// The ID of the CMK. The ID must be globally unique.
	//
	// You can also set this parameter to an alias that is bound to the CMK. For more information, see [Overview of aliases](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 05754286-3ba2-4fa6-8d41-4323aca6****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s DescribeKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeKeyRequest) SetKeyId(v string) *DescribeKeyRequest {
	s.KeyId = &v
	return s
}

type DescribeKeyResponseBody struct {
	// The metadata of the CMK.
	KeyMetadata *DescribeKeyResponseBodyKeyMetadata `json:"KeyMetadata,omitempty" xml:"KeyMetadata,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// f1fdfa9d-bd49-418b-942f-8f3e3ec00a4f
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKeyResponseBody) SetKeyMetadata(v *DescribeKeyResponseBodyKeyMetadata) *DescribeKeyResponseBody {
	s.KeyMetadata = v
	return s
}

func (s *DescribeKeyResponseBody) SetRequestId(v string) *DescribeKeyResponseBody {
	s.RequestId = &v
	return s
}

type DescribeKeyResponseBodyKeyMetadata struct {
	// The Alibaba Cloud Resource Name (ARN) of the CMK.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:154035569884****:key/05754286-3ba2-4fa6-8d41-4323aca6****
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// Indicates whether automatic key rotation is enabled. Valid values:
	//
	// 	- Enabled
	//
	// 	- Disabled
	//
	// 	- Suspended
	//
	// For more information, see [Automatic key rotation](https://help.aliyun.com/document_detail/134270.html).
	//
	// >  Only symmetric CMKs support automatic key rotation.
	//
	// example:
	//
	// Disabled
	AutomaticRotation *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	// The time when the CMK was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-05-20T06:34:21Z
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// The Alibaba Cloud account that is used to create the CMK.
	//
	// example:
	//
	// 154035569884****
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The ID of the dedicated KMS instance.
	//
	// example:
	//
	// kst-bjj62d8f5e0sgtx8h****
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// The time at which the CMK is scheduled for deletion. The time is displayed in UTC.
	//
	// For more information, see [ScheduleKeyDeletion](https://help.aliyun.com/document_detail/44196.html).
	//
	// >  This parameter is returned only when the value of the KeyState parameter is PendingDeletion.
	//
	// example:
	//
	// 2021-05-26T18:22:03Z
	DeleteDate *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
	// Indicates whether deletion protection is enabled. Valid values:
	//
	// 	- Enabled
	//
	// 	- Disabled
	//
	// example:
	//
	// Enabled
	DeletionProtection *string `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The description of deletion protection.
	//
	// example:
	//
	// The CMK is being used by XXX. Deletion protection is set.
	DeletionProtectionDescription *string `json:"DeletionProtectionDescription,omitempty" xml:"DeletionProtectionDescription,omitempty"`
	// The description of the CMK.
	//
	// example:
	//
	// key description example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// example:
	//
	// 05754286-3ba2-4fa6-8d41-4323aca6****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The type of the CMK.
	//
	// example:
	//
	// Aliyun_AES_256
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The status of the CMK.
	//
	// For more information, see [Impact of CMK status on API operations](https://help.aliyun.com/document_detail/44211.html).
	//
	// example:
	//
	// Enabled
	KeyState *string `json:"KeyState,omitempty" xml:"KeyState,omitempty"`
	// The usage of the CMK.
	//
	// example:
	//
	// ENCRYPT/DECRYPT
	KeyUsage *string `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty"`
	// The time when the last rotation was performed. The time is displayed in UTC. For a new CMK, the value of this parameter is the time when the initial version of the CMK was generated.
	//
	// example:
	//
	// 2021-05-20T06:34:21Z
	LastRotationDate *string `json:"LastRotationDate,omitempty" xml:"LastRotationDate,omitempty"`
	// The time when the key material expires. The time is displayed in UTC. If this parameter value is empty, the key material does not expire.
	//
	// example:
	//
	// 2021-07-06T18:22:03Z
	MaterialExpireTime *string `json:"MaterialExpireTime,omitempty" xml:"MaterialExpireTime,omitempty"`
	// The time when the next rotation will be performed.
	//
	// >  This parameter is returned only when the value of the AutomaticRotation parameter is Enabled or Suspended.
	//
	// example:
	//
	// 2021-07-06T18:22:03Z
	NextRotationDate *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	// The source of the key material for the CMK.
	//
	// example:
	//
	// Aliyun_KMS
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The ID of the current primary key version for the symmetric CMK.
	//
	// example:
	//
	// 515e0b0a-624f-45ab-92b5-54f9b551****
	PrimaryKeyVersion *string `json:"PrimaryKeyVersion,omitempty" xml:"PrimaryKeyVersion,omitempty"`
	// The protection level of the CMK.
	//
	// example:
	//
	// HSM
	ProtectionLevel *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	// The interval for automatic key rotation.
	//
	// Unit: seconds.
	//
	// For example, if the value is 604800s, automatic key rotation is performed at a 7-day interval.
	//
	// >  This parameter is returned only when the value of the AutomaticRotation parameter is Enabled or Suspended.
	//
	// example:
	//
	// 31536000s
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
}

func (s DescribeKeyResponseBodyKeyMetadata) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyResponseBodyKeyMetadata) GoString() string {
	return s.String()
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetArn(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.Arn = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetAutomaticRotation(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.AutomaticRotation = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetCreationDate(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.CreationDate = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetCreator(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.Creator = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetDKMSInstanceId(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.DKMSInstanceId = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetDeleteDate(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.DeleteDate = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetDeletionProtection(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetDeletionProtectionDescription(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.DeletionProtectionDescription = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetDescription(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.Description = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetKeyId(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.KeyId = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetKeySpec(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.KeySpec = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetKeyState(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.KeyState = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetKeyUsage(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.KeyUsage = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetLastRotationDate(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.LastRotationDate = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetMaterialExpireTime(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.MaterialExpireTime = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetNextRotationDate(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.NextRotationDate = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetOrigin(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.Origin = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetPrimaryKeyVersion(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.PrimaryKeyVersion = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetProtectionLevel(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.ProtectionLevel = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetRotationInterval(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.RotationInterval = &v
	return s
}

type DescribeKeyResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeKeyResponse) SetHeaders(v map[string]*string) *DescribeKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeKeyResponse) SetStatusCode(v int32) *DescribeKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKeyResponse) SetBody(v *DescribeKeyResponseBody) *DescribeKeyResponse {
	s.Body = v
	return s
}

type DescribeKeyVersionRequest struct {
	// The globally unique ID of the CMK.
	//
	// You can also set this parameter to an alias that is bound to the CMK. For more information, see [Alias overview](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The globally unique ID of the CMK version.
	//
	// You can call the [ListKeyVersions](https://help.aliyun.com/document_detail/133966.html) operation to query the versions of the CMK.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s DescribeKeyVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeKeyVersionRequest) SetKeyId(v string) *DescribeKeyVersionRequest {
	s.KeyId = &v
	return s
}

func (s *DescribeKeyVersionRequest) SetKeyVersionId(v string) *DescribeKeyVersionRequest {
	s.KeyVersionId = &v
	return s
}

type DescribeKeyVersionResponseBody struct {
	// The metadata of the CMK version.
	KeyVersion *DescribeKeyVersionResponseBodyKeyVersion `json:"KeyVersion,omitempty" xml:"KeyVersion,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 7021b6ec-4be7-4d3c-8a68-1e85d4d515a0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeKeyVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKeyVersionResponseBody) SetKeyVersion(v *DescribeKeyVersionResponseBodyKeyVersion) *DescribeKeyVersionResponseBody {
	s.KeyVersion = v
	return s
}

func (s *DescribeKeyVersionResponseBody) SetRequestId(v string) *DescribeKeyVersionResponseBody {
	s.RequestId = &v
	return s
}

type DescribeKeyVersionResponseBodyKeyVersion struct {
	// The date and time when the CMK version was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2016-03-25T10:42:40Z
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// The globally unique ID of the CMK.
	//
	// >  If you set the KeyId parameter in the request to an alias of the CMK, the ID of the CMK to which the alias is bound is returned.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The globally unique ID of the CMK version.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s DescribeKeyVersionResponseBodyKeyVersion) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyVersionResponseBodyKeyVersion) GoString() string {
	return s.String()
}

func (s *DescribeKeyVersionResponseBodyKeyVersion) SetCreationDate(v string) *DescribeKeyVersionResponseBodyKeyVersion {
	s.CreationDate = &v
	return s
}

func (s *DescribeKeyVersionResponseBodyKeyVersion) SetKeyId(v string) *DescribeKeyVersionResponseBodyKeyVersion {
	s.KeyId = &v
	return s
}

func (s *DescribeKeyVersionResponseBodyKeyVersion) SetKeyVersionId(v string) *DescribeKeyVersionResponseBodyKeyVersion {
	s.KeyVersionId = &v
	return s
}

type DescribeKeyVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKeyVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKeyVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeKeyVersionResponse) SetHeaders(v map[string]*string) *DescribeKeyVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeKeyVersionResponse) SetStatusCode(v int32) *DescribeKeyVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKeyVersionResponse) SetBody(v *DescribeKeyVersionResponseBody) *DescribeKeyVersionResponse {
	s.Body = v
	return s
}

type DescribeNetworkRuleRequest struct {
	// The name of the access control rule that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// networkrule_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeNetworkRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleRequest) SetName(v string) *DescribeNetworkRuleRequest {
	s.Name = &v
	return s
}

type DescribeNetworkRuleResponseBody struct {
	// The ARN of the access control rule.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:119285303511****:network/networkrule_test
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The description.
	//
	// example:
	//
	// Creat by kst-hzz62ee817bvyyr5****
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3bf02f7a-015b-4f93-be0f-cc043fda2d33
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The private IP address or private CIDR block.
	//
	// example:
	//
	// ["192.10.XX.XX","192.168.XX.XX/24"]
	SourcePrivateIp *string `json:"SourcePrivateIp,omitempty" xml:"SourcePrivateIp,omitempty"`
	// The network type. Only private IP addresses are supported. The value is fixed as Private.
	//
	// example:
	//
	// Private
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeNetworkRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleResponseBody) SetArn(v string) *DescribeNetworkRuleResponseBody {
	s.Arn = &v
	return s
}

func (s *DescribeNetworkRuleResponseBody) SetDescription(v string) *DescribeNetworkRuleResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeNetworkRuleResponseBody) SetRequestId(v string) *DescribeNetworkRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkRuleResponseBody) SetSourcePrivateIp(v string) *DescribeNetworkRuleResponseBody {
	s.SourcePrivateIp = &v
	return s
}

func (s *DescribeNetworkRuleResponseBody) SetType(v string) *DescribeNetworkRuleResponseBody {
	s.Type = &v
	return s
}

type DescribeNetworkRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetworkRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetworkRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleResponse) SetHeaders(v map[string]*string) *DescribeNetworkRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkRuleResponse) SetStatusCode(v int32) *DescribeNetworkRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetworkRuleResponse) SetBody(v *DescribeNetworkRuleResponseBody) *DescribeNetworkRuleResponse {
	s.Body = v
	return s
}

type DescribePolicyRequest struct {
	// The name of the permission policy that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// policy_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribePolicyRequest) SetName(v string) *DescribePolicyRequest {
	s.Name = &v
	return s
}

type DescribePolicyResponseBody struct {
	// The network access rule that is associated with the permission policy.
	//
	// example:
	//
	// {"NetworkRules":["kst-hzz62ee817bvyyr5x****.efkd","kst-hzz62ee817bvyyr5x****.eyyp"]}
	AccessControlRules *string `json:"AccessControlRules,omitempty" xml:"AccessControlRules,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the permission policy.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:119285303511****:policy/policy_test
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The description.
	//
	// example:
	//
	// policy  description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The scope of the permission policy.
	//
	// example:
	//
	// kst-hzz634e67d126u9p9****
	KmsInstance *string `json:"KmsInstance,omitempty" xml:"KmsInstance,omitempty"`
	// The name of the permission policy.
	//
	// example:
	//
	// policy_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// A list of operations that can be performed.
	//
	// example:
	//
	// ["RbacPermission/Template/CryptoServiceKeyUser", "RbacPermission/Template/CryptoServiceSecretUser"]
	Permissions []*string `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// f455324b-e229-4066-9f58-9c1cf3fe83a9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of keys and secrets that are allowed to access.
	//
	// example:
	//
	// ["secret/acs/ram/user/ram-secret", "secret/acs/ram/user/acr-master", "key/key-hzz63d9c8d3dfv8cv****"]
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s DescribePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolicyResponseBody) SetAccessControlRules(v string) *DescribePolicyResponseBody {
	s.AccessControlRules = &v
	return s
}

func (s *DescribePolicyResponseBody) SetArn(v string) *DescribePolicyResponseBody {
	s.Arn = &v
	return s
}

func (s *DescribePolicyResponseBody) SetDescription(v string) *DescribePolicyResponseBody {
	s.Description = &v
	return s
}

func (s *DescribePolicyResponseBody) SetKmsInstance(v string) *DescribePolicyResponseBody {
	s.KmsInstance = &v
	return s
}

func (s *DescribePolicyResponseBody) SetName(v string) *DescribePolicyResponseBody {
	s.Name = &v
	return s
}

func (s *DescribePolicyResponseBody) SetPermissions(v []*string) *DescribePolicyResponseBody {
	s.Permissions = v
	return s
}

func (s *DescribePolicyResponseBody) SetRequestId(v string) *DescribePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolicyResponseBody) SetResources(v []*string) *DescribePolicyResponseBody {
	s.Resources = v
	return s
}

type DescribePolicyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribePolicyResponse) SetHeaders(v map[string]*string) *DescribePolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribePolicyResponse) SetStatusCode(v int32) *DescribePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolicyResponse) SetBody(v *DescribePolicyResponseBody) *DescribePolicyResponse {
	s.Body = v
	return s
}

type DescribeRegionsResponseBody struct {
	// The region.
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 815240e2-aa37-4c26-9cca-05d4df3e8fe6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeSecretRequest struct {
	// Specifies whether to return the resource tags of the secret. Valid values:
	//
	// 	- true: The resource tags are returned.
	//
	// 	- false: The resource tags are not returned. This is the default value.
	//
	// example:
	//
	// true
	FetchTags *string `json:"FetchTags,omitempty" xml:"FetchTags,omitempty"`
	// The name of the secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s DescribeSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecretRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecretRequest) SetFetchTags(v string) *DescribeSecretRequest {
	s.FetchTags = &v
	return s
}

func (s *DescribeSecretRequest) SetSecretName(v string) *DescribeSecretRequest {
	s.SecretName = &v
	return s
}

type DescribeSecretResponseBody struct {
	// The Alibaba Cloud Resource Name (ARN) of the secret.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:154035569884****:secret/secret001
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// Indicates whether automatic rotation is enabled. Valid values:
	//
	// 	- Enabled: indicates that automatic rotation is enabled.
	//
	// 	- Disabled: indicates that automatic rotation is disabled.
	//
	// 	- Invalid: indicates that the status of automatic rotation is abnormal. In this case, Secrets Manager cannot automatically rotate the secret.
	//
	// >  This parameter is returned only for a managed ApsaraDB RDS secret, a managed RAM secret, or a managed ECS secret.
	//
	// example:
	//
	// Enabled
	AutomaticRotation *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	// The time when the secret was created.
	//
	// example:
	//
	// 2022-02-21T15:39:26Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the dedicated KMS instance.
	//
	// example:
	//
	// kst-bjj62d8f5e0sgtx8h****
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// The description of the secret.
	//
	// example:
	//
	// userinfo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the customer master key (CMK) that is used to encrypt the secret value.
	//
	// example:
	//
	// 00aa68af-2c02-4f68-95fe-3435d330****
	EncryptionKeyId *string `json:"EncryptionKeyId,omitempty" xml:"EncryptionKeyId,omitempty"`
	// The extended configuration of the secret.
	//
	// >  This parameter is returned only for a managed ApsaraDB RDS secret, a managed Resource Access Management (RAM) secret, or a managed Elastic Compute Service (ECS) secret.
	//
	// example:
	//
	// {\\"SecretSubType\\":\\"SingleUser\\", \\"DBInstanceId\\":\\"rm-uf667446pc955****\\",  \\"CustomData\\":{} }
	ExtendedConfig *string `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	// The time when the last rotation was performed.
	//
	// >  This parameter is returned if the secret was rotated.
	//
	// example:
	//
	// 2022-07-05T08:22:03Z
	LastRotationDate *string `json:"LastRotationDate,omitempty" xml:"LastRotationDate,omitempty"`
	// The time when the next rotation will be performed.
	//
	// >  This parameter is returned when automatic rotation is enabled.
	//
	// example:
	//
	// 2022-07-06T18:22:03Z
	NextRotationDate *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	// The time when the secret is scheduled to be deleted.
	//
	// example:
	//
	// 2022-03-21T15:45:12Z
	PlannedDeleteTime *string `json:"PlannedDeleteTime,omitempty" xml:"PlannedDeleteTime,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 93348dfb-3627-4417-8d90-487a76a909c9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The interval for automatic rotation.
	//
	// The value is in the `integer[unit]` format. `integer` indicates the length of time. `unit`: indicates the time unit. The value of `unit` is fixed as s. For example, if the value is 604800s, automatic rotation is performed at a 7-day interval.
	//
	// >  This parameter is returned when automatic rotation is enabled.
	//
	// example:
	//
	// 3153600s
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	// The name of the secret.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The type of the secret. Valid values:
	//
	// 	- Generic: indicates a generic secret.
	//
	// 	- Rds: indicates a managed ApsaraDB RDS secret.
	//
	// 	- RAMCredentials: indicates a managed RAM secret.
	//
	// 	- ECS: indicates a managed ECS secret.
	//
	// example:
	//
	// Rds
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// The resource tags of the secret.
	//
	// This parameter is not returned if you set the FetchTags parameter to false or you do not specify the FetchTags parameter.
	Tags *DescribeSecretResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The time when the secret was updated.
	//
	// example:
	//
	// 2022-02-21T15:39:26Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecretResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecretResponseBody) SetArn(v string) *DescribeSecretResponseBody {
	s.Arn = &v
	return s
}

func (s *DescribeSecretResponseBody) SetAutomaticRotation(v string) *DescribeSecretResponseBody {
	s.AutomaticRotation = &v
	return s
}

func (s *DescribeSecretResponseBody) SetCreateTime(v string) *DescribeSecretResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeSecretResponseBody) SetDKMSInstanceId(v string) *DescribeSecretResponseBody {
	s.DKMSInstanceId = &v
	return s
}

func (s *DescribeSecretResponseBody) SetDescription(v string) *DescribeSecretResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeSecretResponseBody) SetEncryptionKeyId(v string) *DescribeSecretResponseBody {
	s.EncryptionKeyId = &v
	return s
}

func (s *DescribeSecretResponseBody) SetExtendedConfig(v string) *DescribeSecretResponseBody {
	s.ExtendedConfig = &v
	return s
}

func (s *DescribeSecretResponseBody) SetLastRotationDate(v string) *DescribeSecretResponseBody {
	s.LastRotationDate = &v
	return s
}

func (s *DescribeSecretResponseBody) SetNextRotationDate(v string) *DescribeSecretResponseBody {
	s.NextRotationDate = &v
	return s
}

func (s *DescribeSecretResponseBody) SetPlannedDeleteTime(v string) *DescribeSecretResponseBody {
	s.PlannedDeleteTime = &v
	return s
}

func (s *DescribeSecretResponseBody) SetRequestId(v string) *DescribeSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecretResponseBody) SetRotationInterval(v string) *DescribeSecretResponseBody {
	s.RotationInterval = &v
	return s
}

func (s *DescribeSecretResponseBody) SetSecretName(v string) *DescribeSecretResponseBody {
	s.SecretName = &v
	return s
}

func (s *DescribeSecretResponseBody) SetSecretType(v string) *DescribeSecretResponseBody {
	s.SecretType = &v
	return s
}

func (s *DescribeSecretResponseBody) SetTags(v *DescribeSecretResponseBodyTags) *DescribeSecretResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeSecretResponseBody) SetUpdateTime(v string) *DescribeSecretResponseBody {
	s.UpdateTime = &v
	return s
}

type DescribeSecretResponseBodyTags struct {
	Tag []*DescribeSecretResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeSecretResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecretResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeSecretResponseBodyTags) SetTag(v []*DescribeSecretResponseBodyTagsTag) *DescribeSecretResponseBodyTags {
	s.Tag = v
	return s
}

type DescribeSecretResponseBodyTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// val1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeSecretResponseBodyTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecretResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeSecretResponseBodyTagsTag) SetTagKey(v string) *DescribeSecretResponseBodyTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeSecretResponseBodyTagsTag) SetTagValue(v string) *DescribeSecretResponseBodyTagsTag {
	s.TagValue = &v
	return s
}

type DescribeSecretResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecretResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecretResponse) SetHeaders(v map[string]*string) *DescribeSecretResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecretResponse) SetStatusCode(v int32) *DescribeSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecretResponse) SetBody(v *DescribeSecretResponseBody) *DescribeSecretResponse {
	s.Body = v
	return s
}

type DisableKeyRequest struct {
	// The ID of the CMK. The ID must be globally unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s DisableKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableKeyRequest) GoString() string {
	return s.String()
}

func (s *DisableKeyRequest) SetKeyId(v string) *DisableKeyRequest {
	s.KeyId = &v
	return s
}

type DisableKeyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2fe70ce2-3303-4fd6-b3ac-472fb2705c62
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DisableKeyResponseBody) SetRequestId(v string) *DisableKeyResponseBody {
	s.RequestId = &v
	return s
}

type DisableKeyResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableKeyResponse) GoString() string {
	return s.String()
}

func (s *DisableKeyResponse) SetHeaders(v map[string]*string) *DisableKeyResponse {
	s.Headers = v
	return s
}

func (s *DisableKeyResponse) SetStatusCode(v int32) *DisableKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableKeyResponse) SetBody(v *DisableKeyResponseBody) *DisableKeyResponse {
	s.Body = v
	return s
}

type EnableKeyRequest struct {
	// The globally unique ID of the CMK.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s EnableKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableKeyRequest) GoString() string {
	return s.String()
}

func (s *EnableKeyRequest) SetKeyId(v string) *EnableKeyRequest {
	s.KeyId = &v
	return s
}

type EnableKeyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// efb1cbbd-a093-4278-bc03-639dd4fcc207
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableKeyResponseBody) GoString() string {
	return s.String()
}

func (s *EnableKeyResponseBody) SetRequestId(v string) *EnableKeyResponseBody {
	s.RequestId = &v
	return s
}

type EnableKeyResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableKeyResponse) GoString() string {
	return s.String()
}

func (s *EnableKeyResponse) SetHeaders(v map[string]*string) *EnableKeyResponse {
	s.Headers = v
	return s
}

func (s *EnableKeyResponse) SetStatusCode(v int32) *EnableKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableKeyResponse) SetBody(v *EnableKeyResponseBody) *EnableKeyResponse {
	s.Body = v
	return s
}

type EncryptRequest struct {
	// A JSON string that consists of key-value pairs. If you specify this parameter, an equivalent value is required when you call the Decrypt operation. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// example:
	//
	// {"Example":"Example"}
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// The globally unique ID of the CMK. You can also set this parameter to an alias that is bound to the CMK. For more information, see [Use aliases](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The plaintext to be encrypted. The plaintext must be Base64 encoded.
	//
	// This parameter is required.
	//
	// example:
	//
	// SGVsbG8gd29y****
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
}

func (s EncryptRequest) String() string {
	return tea.Prettify(s)
}

func (s EncryptRequest) GoString() string {
	return s.String()
}

func (s *EncryptRequest) SetEncryptionContext(v map[string]interface{}) *EncryptRequest {
	s.EncryptionContext = v
	return s
}

func (s *EncryptRequest) SetKeyId(v string) *EncryptRequest {
	s.KeyId = &v
	return s
}

func (s *EncryptRequest) SetPlaintext(v string) *EncryptRequest {
	s.Plaintext = &v
	return s
}

type EncryptShrinkRequest struct {
	// A JSON string that consists of key-value pairs. If you specify this parameter, an equivalent value is required when you call the Decrypt operation. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// example:
	//
	// {"Example":"Example"}
	EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// The globally unique ID of the CMK. You can also set this parameter to an alias that is bound to the CMK. For more information, see [Use aliases](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The plaintext to be encrypted. The plaintext must be Base64 encoded.
	//
	// This parameter is required.
	//
	// example:
	//
	// SGVsbG8gd29y****
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
}

func (s EncryptShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s EncryptShrinkRequest) GoString() string {
	return s.String()
}

func (s *EncryptShrinkRequest) SetEncryptionContextShrink(v string) *EncryptShrinkRequest {
	s.EncryptionContextShrink = &v
	return s
}

func (s *EncryptShrinkRequest) SetKeyId(v string) *EncryptShrinkRequest {
	s.KeyId = &v
	return s
}

func (s *EncryptShrinkRequest) SetPlaintext(v string) *EncryptShrinkRequest {
	s.Plaintext = &v
	return s
}

type EncryptResponseBody struct {
	// The ciphertext of the data that is encrypted by using the primary CMK version.
	//
	// example:
	//
	// DZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmaaSl+TztSIMe43nbTH/Z1Wr4XfLftKhAciUmDQXuMRl4WTvKhxjMThjK****
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The globally unique ID of the CMK. If you set the KeyId parameter to an alias, the ID of the CMK to which the alias is bound is returned.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the key version that is used to encrypt the plaintext. It is the primary version of the CMK.
	//
	// example:
	//
	// 86a9efd9-3d16-4894-bd4f-1fc43f3f****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 475f1620-b9d3-4d35-b5c6-3fbdd941423d
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EncryptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EncryptResponseBody) GoString() string {
	return s.String()
}

func (s *EncryptResponseBody) SetCiphertextBlob(v string) *EncryptResponseBody {
	s.CiphertextBlob = &v
	return s
}

func (s *EncryptResponseBody) SetKeyId(v string) *EncryptResponseBody {
	s.KeyId = &v
	return s
}

func (s *EncryptResponseBody) SetKeyVersionId(v string) *EncryptResponseBody {
	s.KeyVersionId = &v
	return s
}

func (s *EncryptResponseBody) SetRequestId(v string) *EncryptResponseBody {
	s.RequestId = &v
	return s
}

type EncryptResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EncryptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EncryptResponse) String() string {
	return tea.Prettify(s)
}

func (s EncryptResponse) GoString() string {
	return s.String()
}

func (s *EncryptResponse) SetHeaders(v map[string]*string) *EncryptResponse {
	s.Headers = v
	return s
}

func (s *EncryptResponse) SetStatusCode(v int32) *EncryptResponse {
	s.StatusCode = &v
	return s
}

func (s *EncryptResponse) SetBody(v *EncryptResponseBody) *EncryptResponse {
	s.Body = v
	return s
}

type ExportDataKeyRequest struct {
	// The ciphertext of the data key encrypted by using a CMK.
	//
	// This parameter is required.
	//
	// example:
	//
	// ODZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmS7FmDBBQ0BkKsQrtRnidtPwirmDcS0ZuJCU41xxAAWk4Z8qsADfbV0b+i6kQmlvj79dJdGOvtX69Uycs901q********
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// A JSON string that consists of key-value pairs. If you specify this parameter when you use a CMK to encrypt the data key, an equivalent value is required here. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// example:
	//
	// {"Example":"Example"}
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// A Base64-encoded public key.
	//
	// This parameter is required.
	//
	// example:
	//
	// MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAndKfC2ReLL2+y8a0+ZBBeAft/uBYo86GZiYJuflqgUzKxpyuvlo3uQkBv6b+nx+0tz8g8v7GhpPWMSW5L9mNHYsvYFsa7jTxsYdt17yj6GlUHPuMIs8hr5qbwl38IHU1iIa7nYWwE2fb3ePOvLDACRJVgGpU0yxioW80d2QD+9aU4jF5dlAahcfgsNzo2CXzCUc1+xbmNuq7Rp+H9VJB9dyYOwqnW3RhOLBo21FzpORapf0UiRlrHRpk1V6ez+aE1dofaYh/9bh0m6ioxj7j5hpZbWccuEZTMBKd+cbuBkRhJzc6Tti6qwZbDiu4fUwbZS0Tqpuo1UadiyxMW********
	PublicKeyBlob *string `json:"PublicKeyBlob,omitempty" xml:"PublicKeyBlob,omitempty"`
	// The encryption algorithm based on which you want to use the public key specified by PublicKeyBlob to encrypt the data key. For more information about encryption algorithms, see [AsymmetricDecrypt](https://help.aliyun.com/document_detail/148130.html).
	//
	// Valid values:
	//
	// 	- RSAES_OAEP_SHA_256
	//
	// 	- RSAES_OAEP_SHA_1
	//
	// 	- SM2PKE
	//
	// This parameter is required.
	//
	// example:
	//
	// RSAES_OAEP_SHA_256
	WrappingAlgorithm *string `json:"WrappingAlgorithm,omitempty" xml:"WrappingAlgorithm,omitempty"`
	// The key type of the public key specified by PublicKeyBlob. For more information about key types, see [Introduction to asymmetric keys](https://help.aliyun.com/document_detail/148147.html).
	//
	// Valid values:
	//
	// 	- RSA_2048
	//
	// 	- EC_SM2
	//
	// This parameter is required.
	//
	// example:
	//
	// RSA_2048
	WrappingKeySpec *string `json:"WrappingKeySpec,omitempty" xml:"WrappingKeySpec,omitempty"`
}

func (s ExportDataKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportDataKeyRequest) GoString() string {
	return s.String()
}

func (s *ExportDataKeyRequest) SetCiphertextBlob(v string) *ExportDataKeyRequest {
	s.CiphertextBlob = &v
	return s
}

func (s *ExportDataKeyRequest) SetEncryptionContext(v map[string]interface{}) *ExportDataKeyRequest {
	s.EncryptionContext = v
	return s
}

func (s *ExportDataKeyRequest) SetPublicKeyBlob(v string) *ExportDataKeyRequest {
	s.PublicKeyBlob = &v
	return s
}

func (s *ExportDataKeyRequest) SetWrappingAlgorithm(v string) *ExportDataKeyRequest {
	s.WrappingAlgorithm = &v
	return s
}

func (s *ExportDataKeyRequest) SetWrappingKeySpec(v string) *ExportDataKeyRequest {
	s.WrappingKeySpec = &v
	return s
}

type ExportDataKeyShrinkRequest struct {
	// The ciphertext of the data key encrypted by using a CMK.
	//
	// This parameter is required.
	//
	// example:
	//
	// ODZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmS7FmDBBQ0BkKsQrtRnidtPwirmDcS0ZuJCU41xxAAWk4Z8qsADfbV0b+i6kQmlvj79dJdGOvtX69Uycs901q********
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// A JSON string that consists of key-value pairs. If you specify this parameter when you use a CMK to encrypt the data key, an equivalent value is required here. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// example:
	//
	// {"Example":"Example"}
	EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// A Base64-encoded public key.
	//
	// This parameter is required.
	//
	// example:
	//
	// MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAndKfC2ReLL2+y8a0+ZBBeAft/uBYo86GZiYJuflqgUzKxpyuvlo3uQkBv6b+nx+0tz8g8v7GhpPWMSW5L9mNHYsvYFsa7jTxsYdt17yj6GlUHPuMIs8hr5qbwl38IHU1iIa7nYWwE2fb3ePOvLDACRJVgGpU0yxioW80d2QD+9aU4jF5dlAahcfgsNzo2CXzCUc1+xbmNuq7Rp+H9VJB9dyYOwqnW3RhOLBo21FzpORapf0UiRlrHRpk1V6ez+aE1dofaYh/9bh0m6ioxj7j5hpZbWccuEZTMBKd+cbuBkRhJzc6Tti6qwZbDiu4fUwbZS0Tqpuo1UadiyxMW********
	PublicKeyBlob *string `json:"PublicKeyBlob,omitempty" xml:"PublicKeyBlob,omitempty"`
	// The encryption algorithm based on which you want to use the public key specified by PublicKeyBlob to encrypt the data key. For more information about encryption algorithms, see [AsymmetricDecrypt](https://help.aliyun.com/document_detail/148130.html).
	//
	// Valid values:
	//
	// 	- RSAES_OAEP_SHA_256
	//
	// 	- RSAES_OAEP_SHA_1
	//
	// 	- SM2PKE
	//
	// This parameter is required.
	//
	// example:
	//
	// RSAES_OAEP_SHA_256
	WrappingAlgorithm *string `json:"WrappingAlgorithm,omitempty" xml:"WrappingAlgorithm,omitempty"`
	// The key type of the public key specified by PublicKeyBlob. For more information about key types, see [Introduction to asymmetric keys](https://help.aliyun.com/document_detail/148147.html).
	//
	// Valid values:
	//
	// 	- RSA_2048
	//
	// 	- EC_SM2
	//
	// This parameter is required.
	//
	// example:
	//
	// RSA_2048
	WrappingKeySpec *string `json:"WrappingKeySpec,omitempty" xml:"WrappingKeySpec,omitempty"`
}

func (s ExportDataKeyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportDataKeyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ExportDataKeyShrinkRequest) SetCiphertextBlob(v string) *ExportDataKeyShrinkRequest {
	s.CiphertextBlob = &v
	return s
}

func (s *ExportDataKeyShrinkRequest) SetEncryptionContextShrink(v string) *ExportDataKeyShrinkRequest {
	s.EncryptionContextShrink = &v
	return s
}

func (s *ExportDataKeyShrinkRequest) SetPublicKeyBlob(v string) *ExportDataKeyShrinkRequest {
	s.PublicKeyBlob = &v
	return s
}

func (s *ExportDataKeyShrinkRequest) SetWrappingAlgorithm(v string) *ExportDataKeyShrinkRequest {
	s.WrappingAlgorithm = &v
	return s
}

func (s *ExportDataKeyShrinkRequest) SetWrappingKeySpec(v string) *ExportDataKeyShrinkRequest {
	s.WrappingKeySpec = &v
	return s
}

type ExportDataKeyResponseBody struct {
	// The data key encrypted by using the public key and then exported.
	//
	// example:
	//
	// BQKP+1zK6+ZEMxTP5qaVzcsgXtWplYBKm0NXdSnB5FzliFxE1bSiu4dnEIlca2JpeH7yz1/S6fed630H+hIH6DoM25fTLNcKj+mFB0Xnh9m2+HN59Mn4qyTfcUeadnfCXSWcGBouhXFwcdd2rJ3n337bzTf4jm659gZu3L0i6PLuxM9p7mqdwO0cKJPfGVfhnfMz+f4alMg79WB/NNyE2lyX7/qxvV49ObNrrJbKSFiz8Djocaf0IESNLMbfYI5bXjWkJlX92DQbKhibtQW8ZOJ//ZC6t0AWcUoKL6QDm/dg5koQalcleRinpB+QadFm894sLbVZ9+N4GVs*******
	ExportedDataKey *string `json:"ExportedDataKey,omitempty" xml:"ExportedDataKey,omitempty"`
	// The ID of the CMK that is used to decrypt the specified ciphertext of the data key.
	//
	// This parameter is the globally unique ID of the CMK.
	//
	// example:
	//
	// 202b9877-5a25-46e3-a763-e20791b5****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the CMK version that is used to decrypt the specified ciphertext of the data key.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4bd560a1-729e-45f1-a3d9-b2a33d61046b
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportDataKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportDataKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ExportDataKeyResponseBody) SetExportedDataKey(v string) *ExportDataKeyResponseBody {
	s.ExportedDataKey = &v
	return s
}

func (s *ExportDataKeyResponseBody) SetKeyId(v string) *ExportDataKeyResponseBody {
	s.KeyId = &v
	return s
}

func (s *ExportDataKeyResponseBody) SetKeyVersionId(v string) *ExportDataKeyResponseBody {
	s.KeyVersionId = &v
	return s
}

func (s *ExportDataKeyResponseBody) SetRequestId(v string) *ExportDataKeyResponseBody {
	s.RequestId = &v
	return s
}

type ExportDataKeyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportDataKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportDataKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportDataKeyResponse) GoString() string {
	return s.String()
}

func (s *ExportDataKeyResponse) SetHeaders(v map[string]*string) *ExportDataKeyResponse {
	s.Headers = v
	return s
}

func (s *ExportDataKeyResponse) SetStatusCode(v int32) *ExportDataKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportDataKeyResponse) SetBody(v *ExportDataKeyResponseBody) *ExportDataKeyResponse {
	s.Body = v
	return s
}

type GenerateAndExportDataKeyRequest struct {
	// A JSON string of key-value pairs. If you specify this parameter here, an equivalent value is required when you decrypt or re-encrypt the data key. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// example:
	//
	// {"Example":"Example"}
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// The globally unique ID of the CMK. You can also set this parameter to an alias that is bound to the CMK. For more information, see [Use aliases](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The length of the data key that you want to generate. Valid values:
	//
	// 	- AES_256: a 256-bit symmetric key
	//
	// 	- AES_128: a 128-bit symmetric key
	//
	// >  We recommend that you use the KeySpec or NumberOfBytes parameter to specify the length of a data key. If both parameters are not specified, KMS generates a 256-bit data key. If both parameters are specified, KMS ignores the KeySpec parameter.
	//
	// example:
	//
	// AES_256
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The length of the data key that you want to generate.
	//
	// Valid values: 1 to 1024.
	//
	// Unit: bytes.
	//
	// example:
	//
	// 32
	NumberOfBytes *int32 `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
	// A Base64-encoded public key.
	//
	// This parameter is required.
	//
	// example:
	//
	// MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAndKfC2ReLL2+y8a0+ZBBeAft/uBYo86GZiYJuflqgUzKxpyuvlo3uQkBv6b+nx+0tz8g8v7GhpPWMSW5L9mNHYsvYFsa7jTxsYdt17yj6GlUHPuMIs8hr5qbwl38IHU1iIa7nYWwE2fb3ePOvLDACRJVgGpU0yxioW80d2QD+9aU4jF5dlAahcfgsNzo2CXzCUc1+xbmNuq7Rp+H9VJB9dyYOwqnW3RhOLBo21FzpORapf0UiRlrHRpk1V6ez+aE1dofaYh/9bh0m6ioxj7j5hpZbWccuEZTMBKd+cbuBkRhJzc6Tti6qwZbDiu4fUwbZS0Tqpuo1UadiyxMW********
	PublicKeyBlob *string `json:"PublicKeyBlob,omitempty" xml:"PublicKeyBlob,omitempty"`
	// The encryption algorithm based on which you want to use the public key specified by PublicKeyBlob to encrypt the data key. For more information about encryption algorithms, see [AsymmetricDecrypt](https://help.aliyun.com/document_detail/148130.html).
	//
	// Valid values:
	//
	// 	- RSAES_OAEP_SHA_256
	//
	// 	- RSAES_OAEP_SHA_1
	//
	// 	- SM2PKE
	//
	// This parameter is required.
	//
	// example:
	//
	// RSAES_OAEP_SHA_256
	WrappingAlgorithm *string `json:"WrappingAlgorithm,omitempty" xml:"WrappingAlgorithm,omitempty"`
	// The key type of the public key specified by PublicKeyBlob. For more information about key types, see [Introduction to asymmetric keys](https://help.aliyun.com/document_detail/148147.html).
	//
	// Valid values:
	//
	// 	- RSA_2048
	//
	// 	- EC_SM2
	//
	// This parameter is required.
	//
	// example:
	//
	// RSA_2048
	WrappingKeySpec *string `json:"WrappingKeySpec,omitempty" xml:"WrappingKeySpec,omitempty"`
}

func (s GenerateAndExportDataKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateAndExportDataKeyRequest) GoString() string {
	return s.String()
}

func (s *GenerateAndExportDataKeyRequest) SetEncryptionContext(v map[string]interface{}) *GenerateAndExportDataKeyRequest {
	s.EncryptionContext = v
	return s
}

func (s *GenerateAndExportDataKeyRequest) SetKeyId(v string) *GenerateAndExportDataKeyRequest {
	s.KeyId = &v
	return s
}

func (s *GenerateAndExportDataKeyRequest) SetKeySpec(v string) *GenerateAndExportDataKeyRequest {
	s.KeySpec = &v
	return s
}

func (s *GenerateAndExportDataKeyRequest) SetNumberOfBytes(v int32) *GenerateAndExportDataKeyRequest {
	s.NumberOfBytes = &v
	return s
}

func (s *GenerateAndExportDataKeyRequest) SetPublicKeyBlob(v string) *GenerateAndExportDataKeyRequest {
	s.PublicKeyBlob = &v
	return s
}

func (s *GenerateAndExportDataKeyRequest) SetWrappingAlgorithm(v string) *GenerateAndExportDataKeyRequest {
	s.WrappingAlgorithm = &v
	return s
}

func (s *GenerateAndExportDataKeyRequest) SetWrappingKeySpec(v string) *GenerateAndExportDataKeyRequest {
	s.WrappingKeySpec = &v
	return s
}

type GenerateAndExportDataKeyShrinkRequest struct {
	// A JSON string of key-value pairs. If you specify this parameter here, an equivalent value is required when you decrypt or re-encrypt the data key. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// example:
	//
	// {"Example":"Example"}
	EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// The globally unique ID of the CMK. You can also set this parameter to an alias that is bound to the CMK. For more information, see [Use aliases](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The length of the data key that you want to generate. Valid values:
	//
	// 	- AES_256: a 256-bit symmetric key
	//
	// 	- AES_128: a 128-bit symmetric key
	//
	// >  We recommend that you use the KeySpec or NumberOfBytes parameter to specify the length of a data key. If both parameters are not specified, KMS generates a 256-bit data key. If both parameters are specified, KMS ignores the KeySpec parameter.
	//
	// example:
	//
	// AES_256
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The length of the data key that you want to generate.
	//
	// Valid values: 1 to 1024.
	//
	// Unit: bytes.
	//
	// example:
	//
	// 32
	NumberOfBytes *int32 `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
	// A Base64-encoded public key.
	//
	// This parameter is required.
	//
	// example:
	//
	// MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAndKfC2ReLL2+y8a0+ZBBeAft/uBYo86GZiYJuflqgUzKxpyuvlo3uQkBv6b+nx+0tz8g8v7GhpPWMSW5L9mNHYsvYFsa7jTxsYdt17yj6GlUHPuMIs8hr5qbwl38IHU1iIa7nYWwE2fb3ePOvLDACRJVgGpU0yxioW80d2QD+9aU4jF5dlAahcfgsNzo2CXzCUc1+xbmNuq7Rp+H9VJB9dyYOwqnW3RhOLBo21FzpORapf0UiRlrHRpk1V6ez+aE1dofaYh/9bh0m6ioxj7j5hpZbWccuEZTMBKd+cbuBkRhJzc6Tti6qwZbDiu4fUwbZS0Tqpuo1UadiyxMW********
	PublicKeyBlob *string `json:"PublicKeyBlob,omitempty" xml:"PublicKeyBlob,omitempty"`
	// The encryption algorithm based on which you want to use the public key specified by PublicKeyBlob to encrypt the data key. For more information about encryption algorithms, see [AsymmetricDecrypt](https://help.aliyun.com/document_detail/148130.html).
	//
	// Valid values:
	//
	// 	- RSAES_OAEP_SHA_256
	//
	// 	- RSAES_OAEP_SHA_1
	//
	// 	- SM2PKE
	//
	// This parameter is required.
	//
	// example:
	//
	// RSAES_OAEP_SHA_256
	WrappingAlgorithm *string `json:"WrappingAlgorithm,omitempty" xml:"WrappingAlgorithm,omitempty"`
	// The key type of the public key specified by PublicKeyBlob. For more information about key types, see [Introduction to asymmetric keys](https://help.aliyun.com/document_detail/148147.html).
	//
	// Valid values:
	//
	// 	- RSA_2048
	//
	// 	- EC_SM2
	//
	// This parameter is required.
	//
	// example:
	//
	// RSA_2048
	WrappingKeySpec *string `json:"WrappingKeySpec,omitempty" xml:"WrappingKeySpec,omitempty"`
}

func (s GenerateAndExportDataKeyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateAndExportDataKeyShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateAndExportDataKeyShrinkRequest) SetEncryptionContextShrink(v string) *GenerateAndExportDataKeyShrinkRequest {
	s.EncryptionContextShrink = &v
	return s
}

func (s *GenerateAndExportDataKeyShrinkRequest) SetKeyId(v string) *GenerateAndExportDataKeyShrinkRequest {
	s.KeyId = &v
	return s
}

func (s *GenerateAndExportDataKeyShrinkRequest) SetKeySpec(v string) *GenerateAndExportDataKeyShrinkRequest {
	s.KeySpec = &v
	return s
}

func (s *GenerateAndExportDataKeyShrinkRequest) SetNumberOfBytes(v int32) *GenerateAndExportDataKeyShrinkRequest {
	s.NumberOfBytes = &v
	return s
}

func (s *GenerateAndExportDataKeyShrinkRequest) SetPublicKeyBlob(v string) *GenerateAndExportDataKeyShrinkRequest {
	s.PublicKeyBlob = &v
	return s
}

func (s *GenerateAndExportDataKeyShrinkRequest) SetWrappingAlgorithm(v string) *GenerateAndExportDataKeyShrinkRequest {
	s.WrappingAlgorithm = &v
	return s
}

func (s *GenerateAndExportDataKeyShrinkRequest) SetWrappingKeySpec(v string) *GenerateAndExportDataKeyShrinkRequest {
	s.WrappingKeySpec = &v
	return s
}

type GenerateAndExportDataKeyResponseBody struct {
	// The ciphertext of the data key encrypted by using the primary CMK version.
	//
	// example:
	//
	// ODZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmS7FmDBBQ0BkKsQrtRnidtPwirmDcS0ZuJCU41xxAAWk4Z8qsADfbV0b+i6kQmlvj79dJdGOvtX69Uycs901qOjop4bTS****
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The data key encrypted by using the public key and then exported.
	//
	// example:
	//
	// BQKP+1zK6+ZEMxTP5qaVzcsgXtWplYBKm0NXdSnB5FzliFxE1bSiu4dnEIlca2JpeH7yz1/S6fed630H+hIH6DoM25fTLNcKj+mFB0Xnh9m2+HN59Mn4qyTfcUeadnfCXSWcGBouhXFwcdd2rJ3n337bzTf4jm659gZu3L0i6PLuxM9p7mqdwO0cKJPfGVfhnfMz+f4alMg79WB/NNyE2lyX7/qxvV49ObNrrJbKSFiz8Djocaf0IESNLMbfYI5bXjWkJlX92DQbKhibtQW8ZOJ//ZC6t0AWcUoKL6QDm/dg5koQalcleRinpB+QadFm894sLbVZ9+N4GVs*******
	ExportedDataKey *string `json:"ExportedDataKey,omitempty" xml:"ExportedDataKey,omitempty"`
	// The globally unique ID of the CMK.
	//
	// >  If you set the KeyId parameter to an alias, the ID of the CMK to which the alias is bound is returned.
	//
	// example:
	//
	// 599fa825-17de-417e-9554-bb032cc6****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the CMK version that is used to encrypt the plaintext. It is the primary version of the CMK.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7021b6ec-4be7-4d3c-8a68-1e85d4d515a0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateAndExportDataKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateAndExportDataKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateAndExportDataKeyResponseBody) SetCiphertextBlob(v string) *GenerateAndExportDataKeyResponseBody {
	s.CiphertextBlob = &v
	return s
}

func (s *GenerateAndExportDataKeyResponseBody) SetExportedDataKey(v string) *GenerateAndExportDataKeyResponseBody {
	s.ExportedDataKey = &v
	return s
}

func (s *GenerateAndExportDataKeyResponseBody) SetKeyId(v string) *GenerateAndExportDataKeyResponseBody {
	s.KeyId = &v
	return s
}

func (s *GenerateAndExportDataKeyResponseBody) SetKeyVersionId(v string) *GenerateAndExportDataKeyResponseBody {
	s.KeyVersionId = &v
	return s
}

func (s *GenerateAndExportDataKeyResponseBody) SetRequestId(v string) *GenerateAndExportDataKeyResponseBody {
	s.RequestId = &v
	return s
}

type GenerateAndExportDataKeyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateAndExportDataKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateAndExportDataKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateAndExportDataKeyResponse) GoString() string {
	return s.String()
}

func (s *GenerateAndExportDataKeyResponse) SetHeaders(v map[string]*string) *GenerateAndExportDataKeyResponse {
	s.Headers = v
	return s
}

func (s *GenerateAndExportDataKeyResponse) SetStatusCode(v int32) *GenerateAndExportDataKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateAndExportDataKeyResponse) SetBody(v *GenerateAndExportDataKeyResponseBody) *GenerateAndExportDataKeyResponse {
	s.Body = v
	return s
}

type GenerateDataKeyRequest struct {
	// The JSON string that consists of key-value pairs.
	//
	// If you specify this parameter, an equivalent value is required when you call the [Decrypt](https://help.aliyun.com/document_detail/28950.html) operation. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// example:
	//
	// {"Example":"Example"}
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// You can also set this parameter to an alias that is bound to the CMK. For more information, see [Alias overview](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 7906979c-8e06-46a2-be2d-68e3ccbc****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The type of the data key that you want to generate. Valid values:
	//
	// 	- AES_256: a 256-bit symmetric key
	//
	// 	- AES_128: a 128-bit symmetric key
	//
	// >  We recommend that you use the KeySpec or NumberOfBytes parameter to specify the length of a data key. If none of the parameters are specified, KMS generates a 256-bit data key. If both parameters are specified, KMS ignores the KeySpec parameter.
	//
	// example:
	//
	// AES_256
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The length of the data key that you want to generate. Unit: bytes.
	//
	// Valid values: 1 to 1024.
	//
	// Default value:
	//
	// 	- If the KeySpec parameter is set to AES_256, set the value of the NumberOfBytes parameter to 32.
	//
	// 	- If the KeySpec parameter is set to AES_128, set the value of the NumberOfBytes parameter to 16.
	//
	// example:
	//
	// 256
	NumberOfBytes *int32 `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
}

func (s GenerateDataKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateDataKeyRequest) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyRequest) SetEncryptionContext(v map[string]interface{}) *GenerateDataKeyRequest {
	s.EncryptionContext = v
	return s
}

func (s *GenerateDataKeyRequest) SetKeyId(v string) *GenerateDataKeyRequest {
	s.KeyId = &v
	return s
}

func (s *GenerateDataKeyRequest) SetKeySpec(v string) *GenerateDataKeyRequest {
	s.KeySpec = &v
	return s
}

func (s *GenerateDataKeyRequest) SetNumberOfBytes(v int32) *GenerateDataKeyRequest {
	s.NumberOfBytes = &v
	return s
}

type GenerateDataKeyShrinkRequest struct {
	// The JSON string that consists of key-value pairs.
	//
	// If you specify this parameter, an equivalent value is required when you call the [Decrypt](https://help.aliyun.com/document_detail/28950.html) operation. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// example:
	//
	// {"Example":"Example"}
	EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// You can also set this parameter to an alias that is bound to the CMK. For more information, see [Alias overview](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 7906979c-8e06-46a2-be2d-68e3ccbc****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The type of the data key that you want to generate. Valid values:
	//
	// 	- AES_256: a 256-bit symmetric key
	//
	// 	- AES_128: a 128-bit symmetric key
	//
	// >  We recommend that you use the KeySpec or NumberOfBytes parameter to specify the length of a data key. If none of the parameters are specified, KMS generates a 256-bit data key. If both parameters are specified, KMS ignores the KeySpec parameter.
	//
	// example:
	//
	// AES_256
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The length of the data key that you want to generate. Unit: bytes.
	//
	// Valid values: 1 to 1024.
	//
	// Default value:
	//
	// 	- If the KeySpec parameter is set to AES_256, set the value of the NumberOfBytes parameter to 32.
	//
	// 	- If the KeySpec parameter is set to AES_128, set the value of the NumberOfBytes parameter to 16.
	//
	// example:
	//
	// 256
	NumberOfBytes *int32 `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
}

func (s GenerateDataKeyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateDataKeyShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyShrinkRequest) SetEncryptionContextShrink(v string) *GenerateDataKeyShrinkRequest {
	s.EncryptionContextShrink = &v
	return s
}

func (s *GenerateDataKeyShrinkRequest) SetKeyId(v string) *GenerateDataKeyShrinkRequest {
	s.KeyId = &v
	return s
}

func (s *GenerateDataKeyShrinkRequest) SetKeySpec(v string) *GenerateDataKeyShrinkRequest {
	s.KeySpec = &v
	return s
}

func (s *GenerateDataKeyShrinkRequest) SetNumberOfBytes(v int32) *GenerateDataKeyShrinkRequest {
	s.NumberOfBytes = &v
	return s
}

type GenerateDataKeyResponseBody struct {
	// The ciphertext of the data key that is encrypted by using the primary version of the specified CMK.
	//
	// example:
	//
	// ODZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmS7FmDBBQ0BkKsQrtRnidtPwirmDcS0ZuJCU41xxAAWk4Z8qsADfbV0b+i6kQmlvj79dJdGOvtX69Uycs901qOjop4bTS****
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// >  If you set the KeyId parameter in the request to an alias of the CMK, the ID of the CMK to which the alias is bound is returned.
	//
	// example:
	//
	// 7906979c-8e06-46a2-be2d-68e3ccbc****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the CMK version. The ID must be globally unique.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The Base64 encoded plaintext of the data key.
	//
	// example:
	//
	// QmFzZTY0IGVuY29kZWQgcGxhaW50****
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7021b6ec-4be7-4d3c-8a68-1e85d4d515a0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateDataKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateDataKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyResponseBody) SetCiphertextBlob(v string) *GenerateDataKeyResponseBody {
	s.CiphertextBlob = &v
	return s
}

func (s *GenerateDataKeyResponseBody) SetKeyId(v string) *GenerateDataKeyResponseBody {
	s.KeyId = &v
	return s
}

func (s *GenerateDataKeyResponseBody) SetKeyVersionId(v string) *GenerateDataKeyResponseBody {
	s.KeyVersionId = &v
	return s
}

func (s *GenerateDataKeyResponseBody) SetPlaintext(v string) *GenerateDataKeyResponseBody {
	s.Plaintext = &v
	return s
}

func (s *GenerateDataKeyResponseBody) SetRequestId(v string) *GenerateDataKeyResponseBody {
	s.RequestId = &v
	return s
}

type GenerateDataKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateDataKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateDataKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateDataKeyResponse) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyResponse) SetHeaders(v map[string]*string) *GenerateDataKeyResponse {
	s.Headers = v
	return s
}

func (s *GenerateDataKeyResponse) SetStatusCode(v int32) *GenerateDataKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateDataKeyResponse) SetBody(v *GenerateDataKeyResponseBody) *GenerateDataKeyResponse {
	s.Body = v
	return s
}

type GenerateDataKeyWithoutPlaintextRequest struct {
	// A JSON string that consists of key-value pairs. If you specify this parameter, an equivalent value is required when you call the Decrypt operation. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// example:
	//
	// {"Example":"Example"}
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// The globally unique ID of the CMK. You can also set this parameter to an alias that is bound to the CMK. For more information, see Use aliases.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The length of the data key that you want to generate. Valid values:
	//
	// 	- AES_256: 256-bit symmetric key
	//
	// 	- AES_128: 128-bit symmetric key
	//
	// >  We recommend that you use the KeySpec or NumberOfBytes parameter to specify the length of a data key. If both of them are not specified, KMS generates a 256-bit data key. If both of them are specified, KMS ignores the KeySpec parameter.
	//
	// example:
	//
	// AES_256
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The length of the data key that you want to generate.
	//
	// Valid values: 1 to 1024.
	//
	// Unit: bytes.
	//
	// example:
	//
	// 256
	NumberOfBytes *int32 `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
}

func (s GenerateDataKeyWithoutPlaintextRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateDataKeyWithoutPlaintextRequest) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyWithoutPlaintextRequest) SetEncryptionContext(v map[string]interface{}) *GenerateDataKeyWithoutPlaintextRequest {
	s.EncryptionContext = v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextRequest) SetKeyId(v string) *GenerateDataKeyWithoutPlaintextRequest {
	s.KeyId = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextRequest) SetKeySpec(v string) *GenerateDataKeyWithoutPlaintextRequest {
	s.KeySpec = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextRequest) SetNumberOfBytes(v int32) *GenerateDataKeyWithoutPlaintextRequest {
	s.NumberOfBytes = &v
	return s
}

type GenerateDataKeyWithoutPlaintextShrinkRequest struct {
	// A JSON string that consists of key-value pairs. If you specify this parameter, an equivalent value is required when you call the Decrypt operation. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// example:
	//
	// {"Example":"Example"}
	EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// The globally unique ID of the CMK. You can also set this parameter to an alias that is bound to the CMK. For more information, see Use aliases.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The length of the data key that you want to generate. Valid values:
	//
	// 	- AES_256: 256-bit symmetric key
	//
	// 	- AES_128: 128-bit symmetric key
	//
	// >  We recommend that you use the KeySpec or NumberOfBytes parameter to specify the length of a data key. If both of them are not specified, KMS generates a 256-bit data key. If both of them are specified, KMS ignores the KeySpec parameter.
	//
	// example:
	//
	// AES_256
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The length of the data key that you want to generate.
	//
	// Valid values: 1 to 1024.
	//
	// Unit: bytes.
	//
	// example:
	//
	// 256
	NumberOfBytes *int32 `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
}

func (s GenerateDataKeyWithoutPlaintextShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateDataKeyWithoutPlaintextShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyWithoutPlaintextShrinkRequest) SetEncryptionContextShrink(v string) *GenerateDataKeyWithoutPlaintextShrinkRequest {
	s.EncryptionContextShrink = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextShrinkRequest) SetKeyId(v string) *GenerateDataKeyWithoutPlaintextShrinkRequest {
	s.KeyId = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextShrinkRequest) SetKeySpec(v string) *GenerateDataKeyWithoutPlaintextShrinkRequest {
	s.KeySpec = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextShrinkRequest) SetNumberOfBytes(v int32) *GenerateDataKeyWithoutPlaintextShrinkRequest {
	s.NumberOfBytes = &v
	return s
}

type GenerateDataKeyWithoutPlaintextResponseBody struct {
	// The ciphertext of the data that is encrypted by using the primary CMK version.
	//
	// example:
	//
	// ODZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmS7FmDBBQ0BkKsQrtRnidtPwirmDcS0ZuJCU41xxAAWk4Z8qsADfbV0b+i6kQmlvj79dJdGOvtX69Uycs901qOjop4bTS****
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The globally unique ID of the CMK.
	//
	// >  If you set the KeyId parameter to an alias, the ID of the CMK to which the alias is bound is returned.
	//
	// example:
	//
	// 599fa825-17de-417e-9554-bb032cc6****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the key version that is used to encrypt the plaintext. It is the primary version of the CMK.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7021b6ec-4be7-4d3c-8a68-1e85d4d515a0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateDataKeyWithoutPlaintextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateDataKeyWithoutPlaintextResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyWithoutPlaintextResponseBody) SetCiphertextBlob(v string) *GenerateDataKeyWithoutPlaintextResponseBody {
	s.CiphertextBlob = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextResponseBody) SetKeyId(v string) *GenerateDataKeyWithoutPlaintextResponseBody {
	s.KeyId = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextResponseBody) SetKeyVersionId(v string) *GenerateDataKeyWithoutPlaintextResponseBody {
	s.KeyVersionId = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextResponseBody) SetRequestId(v string) *GenerateDataKeyWithoutPlaintextResponseBody {
	s.RequestId = &v
	return s
}

type GenerateDataKeyWithoutPlaintextResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateDataKeyWithoutPlaintextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateDataKeyWithoutPlaintextResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateDataKeyWithoutPlaintextResponse) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyWithoutPlaintextResponse) SetHeaders(v map[string]*string) *GenerateDataKeyWithoutPlaintextResponse {
	s.Headers = v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextResponse) SetStatusCode(v int32) *GenerateDataKeyWithoutPlaintextResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextResponse) SetBody(v *GenerateDataKeyWithoutPlaintextResponseBody) *GenerateDataKeyWithoutPlaintextResponse {
	s.Body = v
	return s
}

type GetCertificateRequest struct {
	// The ID of the certificate. It is the globally unique identifier (GUID) of the certificate in Certificates Manager.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9a28de48-8d8b-484d-a766-dec4****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s GetCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateRequest) GoString() string {
	return s.String()
}

func (s *GetCertificateRequest) SetCertificateId(v string) *GetCertificateRequest {
	s.CertificateId = &v
	return s
}

type GetCertificateResponseBody struct {
	// The certificate in the Privacy Enhanced Mail (PEM) format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----  (X.509 Certificate PEM Content)  -----END CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The certificate chain in the PEM format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----  (Sub CA Certificate PEM Content)  -----END CERTIFICATE-----  -----BEGIN CERTIFICATE-----  (Sub CA Certificate PEM Content)  -----END CERTIFICATE-----  -----BEGIN CERTIFICATE-----  (Root CA Certificate PEM Content)  -----END CERTIFICATE-----
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The ID of the certificate.
	//
	// example:
	//
	// 9a28de48-8d8b-484d-a766-dec4****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The CSR in the PEM format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST-----MIICxjCCAa4CAQAwPzELMAkGA1UEBhMCQ04xDzANBgNVBAoTBmFsaXl1bjEMMAoGA1UECxMDa21zMREwDwY****-----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// b3e104b4-0319-4a20-ab7f-9fef6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *GetCertificateResponseBody) SetCertificate(v string) *GetCertificateResponseBody {
	s.Certificate = &v
	return s
}

func (s *GetCertificateResponseBody) SetCertificateChain(v string) *GetCertificateResponseBody {
	s.CertificateChain = &v
	return s
}

func (s *GetCertificateResponseBody) SetCertificateId(v string) *GetCertificateResponseBody {
	s.CertificateId = &v
	return s
}

func (s *GetCertificateResponseBody) SetCsr(v string) *GetCertificateResponseBody {
	s.Csr = &v
	return s
}

func (s *GetCertificateResponseBody) SetRequestId(v string) *GetCertificateResponseBody {
	s.RequestId = &v
	return s
}

type GetCertificateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateResponse) GoString() string {
	return s.String()
}

func (s *GetCertificateResponse) SetHeaders(v map[string]*string) *GetCertificateResponse {
	s.Headers = v
	return s
}

func (s *GetCertificateResponse) SetStatusCode(v int32) *GetCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCertificateResponse) SetBody(v *GetCertificateResponseBody) *GetCertificateResponse {
	s.Body = v
	return s
}

type GetClientKeyRequest struct {
	// The ID of the client key.
	//
	// This parameter is required.
	//
	// example:
	//
	// KAAP.66abf237-63f6-4625-b8cf-47e1086e****
	ClientKeyId *string `json:"ClientKeyId,omitempty" xml:"ClientKeyId,omitempty"`
}

func (s GetClientKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClientKeyRequest) GoString() string {
	return s.String()
}

func (s *GetClientKeyRequest) SetClientKeyId(v string) *GetClientKeyRequest {
	s.ClientKeyId = &v
	return s
}

type GetClientKeyResponseBody struct {
	// The name of the application access point (AAP).
	//
	// example:
	//
	// aap_test
	AapName *string `json:"AapName,omitempty" xml:"AapName,omitempty"`
	// The ID of the client key.
	//
	// example:
	//
	// KAAP.66abf237-63f6-4625-b8cf-47e1086e****
	ClientKeyId *string `json:"ClientKeyId,omitempty" xml:"ClientKeyId,omitempty"`
	// The time when the client key was created.
	//
	// example:
	//
	// 2023-08-31T09:14:38Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The private key algorithm of the client key.
	//
	// example:
	//
	// RSA_2048
	KeyAlgorithm *string `json:"KeyAlgorithm,omitempty" xml:"KeyAlgorithm,omitempty"`
	// The provider of the client key.
	//
	// Currently, only Key Management Service (KMS) is supported. The value is fixed as KMS_PROVIDED.
	//
	// example:
	//
	// KMS_PROVIDED
	KeyOrigin *string `json:"KeyOrigin,omitempty" xml:"KeyOrigin,omitempty"`
	// The end of the validity period of the client key.
	//
	// example:
	//
	// 2028-08-31T17:14:33Z
	NotAfter *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The beginning of the validity period of the client key.
	//
	// example:
	//
	// 2023-08-31T17:14:33Z
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The content of the public key of the client key.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\nMIIDcjCCAlqgAwIBAgIQT/sAVRxwYp54mrw****-----END CERTIFICATE-----
	PublicKeyData *string `json:"PublicKeyData,omitempty" xml:"PublicKeyData,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 63d849a6-045b-4a57-ad9f-c5f756cea9e9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetClientKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClientKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetClientKeyResponseBody) SetAapName(v string) *GetClientKeyResponseBody {
	s.AapName = &v
	return s
}

func (s *GetClientKeyResponseBody) SetClientKeyId(v string) *GetClientKeyResponseBody {
	s.ClientKeyId = &v
	return s
}

func (s *GetClientKeyResponseBody) SetCreateTime(v string) *GetClientKeyResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetClientKeyResponseBody) SetKeyAlgorithm(v string) *GetClientKeyResponseBody {
	s.KeyAlgorithm = &v
	return s
}

func (s *GetClientKeyResponseBody) SetKeyOrigin(v string) *GetClientKeyResponseBody {
	s.KeyOrigin = &v
	return s
}

func (s *GetClientKeyResponseBody) SetNotAfter(v string) *GetClientKeyResponseBody {
	s.NotAfter = &v
	return s
}

func (s *GetClientKeyResponseBody) SetNotBefore(v string) *GetClientKeyResponseBody {
	s.NotBefore = &v
	return s
}

func (s *GetClientKeyResponseBody) SetPublicKeyData(v string) *GetClientKeyResponseBody {
	s.PublicKeyData = &v
	return s
}

func (s *GetClientKeyResponseBody) SetRequestId(v string) *GetClientKeyResponseBody {
	s.RequestId = &v
	return s
}

type GetClientKeyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClientKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClientKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClientKeyResponse) GoString() string {
	return s.String()
}

func (s *GetClientKeyResponse) SetHeaders(v map[string]*string) *GetClientKeyResponse {
	s.Headers = v
	return s
}

func (s *GetClientKeyResponse) SetStatusCode(v int32) *GetClientKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClientKeyResponse) SetBody(v *GetClientKeyResponseBody) *GetClientKeyResponse {
	s.Body = v
	return s
}

type GetKeyPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// key-hzz630494463ejqjx****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// example:
	//
	// default
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s GetKeyPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetKeyPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetKeyPolicyRequest) SetKeyId(v string) *GetKeyPolicyRequest {
	s.KeyId = &v
	return s
}

func (s *GetKeyPolicyRequest) SetPolicyName(v string) *GetKeyPolicyRequest {
	s.PolicyName = &v
	return s
}

type GetKeyPolicyResponseBody struct {
	// example:
	//
	// {"Statement": [{"Action": ["kms:*"],"Effect": "Allow","Principal": {"RAM": ["acs:ram::190325303126****:*","acs:ram::119285303511****:*"]},"Resource": ["*"],"Sid": "kms default key policy"}],"Version": "1" }
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// 381D5D33-BB8F-395F-8EE4-AE3B84B523C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetKeyPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetKeyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetKeyPolicyResponseBody) SetPolicy(v string) *GetKeyPolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *GetKeyPolicyResponseBody) SetRequestId(v string) *GetKeyPolicyResponseBody {
	s.RequestId = &v
	return s
}

type GetKeyPolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKeyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKeyPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetKeyPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetKeyPolicyResponse) SetHeaders(v map[string]*string) *GetKeyPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetKeyPolicyResponse) SetStatusCode(v int32) *GetKeyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKeyPolicyResponse) SetBody(v *GetKeyPolicyResponseBody) *GetKeyPolicyResponse {
	s.Body = v
	return s
}

type GetKmsInstanceRequest struct {
	// The ID of the KMS instance that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// kst-bjj62f5ba3dnpb6v8****
	KmsInstanceId *string `json:"KmsInstanceId,omitempty" xml:"KmsInstanceId,omitempty"`
}

func (s GetKmsInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetKmsInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetKmsInstanceRequest) SetKmsInstanceId(v string) *GetKmsInstanceRequest {
	s.KmsInstanceId = &v
	return s
}

type GetKmsInstanceResponseBody struct {
	// The details of the KMS instance.
	KmsInstance *GetKmsInstanceResponseBodyKmsInstance `json:"KmsInstance,omitempty" xml:"KmsInstance,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 46b4a94a-57d2-44b4-9810-1e87d31abb33
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetKmsInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetKmsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetKmsInstanceResponseBody) SetKmsInstance(v *GetKmsInstanceResponseBodyKmsInstance) *GetKmsInstanceResponseBody {
	s.KmsInstance = v
	return s
}

func (s *GetKmsInstanceResponseBody) SetRequestId(v string) *GetKmsInstanceResponseBody {
	s.RequestId = &v
	return s
}

type GetKmsInstanceResponseBodyKmsInstance struct {
	// A list of associated VPCs.
	//
	// >  If your self-managed applications are deployed in multiple VPCs in the same region, you can associate VPCs with the KMS instance beyond the VPC that you specify when you enable the KMS instance. The VPCs can belong to the same Alibaba Cloud account or different Alibaba Cloud accounts. After the configuration is complete, self-managed applications in the VPCs can access the specified KMS instance.
	BindVpcs *GetKmsInstanceResponseBodyKmsInstanceBindVpcs `json:"BindVpcs,omitempty" xml:"BindVpcs,omitempty" type:"Struct"`
	// The content of the certificate authority (CA) certificate of the KMS instance.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\r\\nMIIDuzCCAqOgAwIBAgIJALTKwWAjvbMiMA0GCSqGSIb3DQEBCwUAMHQxCzAJBgNV****-----END CERTIFICATE-----
	CaCertificateChainPem *string `json:"CaCertificateChainPem,omitempty" xml:"CaCertificateChainPem,omitempty"`
	// The time when the KMS instance is created.
	//
	// example:
	//
	// 2023-09-05T12:44:20Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The expiration time of the KMS instance.
	//
	// example:
	//
	// 2023-10-05T16:00:00Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The ID of the KMS instance.
	//
	// example:
	//
	// kst-bjj62f5ba3dnpb6v8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the KMS instance.
	//
	// example:
	//
	// kst-bjj62f5ba3dnpb6v8****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The number of keys that can be created for the KMS instance.
	//
	// example:
	//
	// 1000
	KeyNum *int64 `json:"KeyNum,omitempty" xml:"KeyNum,omitempty"`
	// The number of secrets that can be created for the KMS instance.
	//
	// example:
	//
	// 10
	SecretNum *string `json:"SecretNum,omitempty" xml:"SecretNum,omitempty"`
	// The computing performance of the KMS instance.
	//
	// example:
	//
	// 1000
	Spec *int64 `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The time when the KMS instance is enabled.
	//
	// example:
	//
	// 2023-09-05T12:44:19Z
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The status of the KMS instance. Valid values:
	//
	// 	- Uninitialized: The KMS instance is not enabled.
	//
	// 	- Connecting: The KMS instance is being connected.
	//
	// 	- Connected: The KMS instance is enabled.
	//
	// 	- Disconnected: The KMS instance is disconnected.
	//
	// 	- Error: The KMS instance is abnormal.
	//
	// example:
	//
	// Connected
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The virtual private cloud (VPC) with which the KMS instance is associated.
	//
	// example:
	//
	// vpc-bp19z7cwmltad5dff****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The access management quota for the KMS instance.
	//
	// example:
	//
	// 5
	VpcNum *int64 `json:"VpcNum,omitempty" xml:"VpcNum,omitempty"`
	// The vSwitch in the VPC.
	//
	// example:
	//
	// vsw-bp1i512amda6d10a0****
	VswitchIds *string `json:"VswitchIds,omitempty" xml:"VswitchIds,omitempty"`
	// The zone with which the KMS instance is associated.
	//
	// example:
	//
	// "cn-hangzhou-k",       "cn-hangzhou-j"
	ZoneIds *string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty"`
}

func (s GetKmsInstanceResponseBodyKmsInstance) String() string {
	return tea.Prettify(s)
}

func (s GetKmsInstanceResponseBodyKmsInstance) GoString() string {
	return s.String()
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetBindVpcs(v *GetKmsInstanceResponseBodyKmsInstanceBindVpcs) *GetKmsInstanceResponseBodyKmsInstance {
	s.BindVpcs = v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetCaCertificateChainPem(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.CaCertificateChainPem = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetCreateTime(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.CreateTime = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetEndDate(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.EndDate = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetInstanceId(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.InstanceId = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetInstanceName(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.InstanceName = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetKeyNum(v int64) *GetKmsInstanceResponseBodyKmsInstance {
	s.KeyNum = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetSecretNum(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.SecretNum = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetSpec(v int64) *GetKmsInstanceResponseBodyKmsInstance {
	s.Spec = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetStartDate(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.StartDate = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetStatus(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.Status = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetVpcId(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.VpcId = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetVpcNum(v int64) *GetKmsInstanceResponseBodyKmsInstance {
	s.VpcNum = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetVswitchIds(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.VswitchIds = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetZoneIds(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.ZoneIds = &v
	return s
}

type GetKmsInstanceResponseBodyKmsInstanceBindVpcs struct {
	BindVpc []*GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc `json:"BindVpc,omitempty" xml:"BindVpc,omitempty" type:"Repeated"`
}

func (s GetKmsInstanceResponseBodyKmsInstanceBindVpcs) String() string {
	return tea.Prettify(s)
}

func (s GetKmsInstanceResponseBodyKmsInstanceBindVpcs) GoString() string {
	return s.String()
}

func (s *GetKmsInstanceResponseBodyKmsInstanceBindVpcs) SetBindVpc(v []*GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc) *GetKmsInstanceResponseBodyKmsInstanceBindVpcs {
	s.BindVpc = v
	return s
}

type GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc struct {
	// The region to which the VPC belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The vSwitch in the VPC.
	//
	// example:
	//
	// vsw-bp1i512amhdje10f1****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-bp19z7djuhtad5dff****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The Alibaba Cloud account to which the VPC belongs.
	//
	// example:
	//
	// 190325303126****
	VpcOwnerId *string `json:"VpcOwnerId,omitempty" xml:"VpcOwnerId,omitempty"`
}

func (s GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc) String() string {
	return tea.Prettify(s)
}

func (s GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc) GoString() string {
	return s.String()
}

func (s *GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc) SetRegionId(v string) *GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc {
	s.RegionId = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc) SetVSwitchId(v string) *GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc {
	s.VSwitchId = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc) SetVpcId(v string) *GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc {
	s.VpcId = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc) SetVpcOwnerId(v string) *GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc {
	s.VpcOwnerId = &v
	return s
}

type GetKmsInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKmsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKmsInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetKmsInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetKmsInstanceResponse) SetHeaders(v map[string]*string) *GetKmsInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetKmsInstanceResponse) SetStatusCode(v int32) *GetKmsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKmsInstanceResponse) SetBody(v *GetKmsInstanceResponseBody) *GetKmsInstanceResponse {
	s.Body = v
	return s
}

type GetParametersForImportRequest struct {
	// The globally unique ID of the CMK.
	//
	// >  You can import key material only for CMKs whose Origin parameter is set to EXTERNAL.
	//
	// This parameter is required.
	//
	// example:
	//
	// 202b9877-5a25-46e3-a763-e20791b5****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The algorithm that is used to encrypt key material.
	//
	// This parameter is required.
	//
	// example:
	//
	// RSAES_PKCS1_V1_5
	WrappingAlgorithm *string `json:"WrappingAlgorithm,omitempty" xml:"WrappingAlgorithm,omitempty"`
	// The type of the public key that is used to encrypt key material.
	//
	// This parameter is required.
	//
	// example:
	//
	// RSA_2048
	WrappingKeySpec *string `json:"WrappingKeySpec,omitempty" xml:"WrappingKeySpec,omitempty"`
}

func (s GetParametersForImportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetParametersForImportRequest) GoString() string {
	return s.String()
}

func (s *GetParametersForImportRequest) SetKeyId(v string) *GetParametersForImportRequest {
	s.KeyId = &v
	return s
}

func (s *GetParametersForImportRequest) SetWrappingAlgorithm(v string) *GetParametersForImportRequest {
	s.WrappingAlgorithm = &v
	return s
}

func (s *GetParametersForImportRequest) SetWrappingKeySpec(v string) *GetParametersForImportRequest {
	s.WrappingKeySpec = &v
	return s
}

type GetParametersForImportResponseBody struct {
	// The token that is used to import key material.
	//
	// The token is valid for 24 hours. The value of this parameter is required when you call the [ImportKeyMaterial](https://help.aliyun.com/document_detail/68622.html) operation.
	//
	// example:
	//
	// Base64String
	ImportToken *string `json:"ImportToken,omitempty" xml:"ImportToken,omitempty"`
	// The globally unique ID of the CMK.
	//
	// The value of this parameter is required when you call the [ImportKeyMaterial](https://help.aliyun.com/document_detail/68622.html) operation.
	//
	// example:
	//
	// 202b9877-5a25-46e3-a763-e20791b5****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The public key that is used to encrypt key material.
	//
	// The public key is Base64-encoded.
	//
	// example:
	//
	// MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAlls4uIBxD0GG84C+lGBO6Dhpf1J3XimC6cPmPNaKKJMOzoX4tD+C+r7aZv8lZ3vnPfxuxvy/YwG+whUxTEEFUdqJTOIzhPfYucupqKM92crVHIuG+xtMVeHKjyTr+UrtKCsQikqHT+19yDRN/RMoo2HUx0gmEnRyXd8t3JyUXun9FdoxKA08GrsV7nodb9ZsoBLhnev7tTLcXvLyKW6XG1ZQCQm6dPnbnwLeDXR7uK0Lqn9PM28mBIdaiQUQxj2XbM1CoJA+JiyVX3Ptdb+4rqukb4Rb05B80Bs9xV/cf7FIku08l7xGhrGiQFq+DFXwQWtwihXHZxz3LhldU+4ZPwID****
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 8cdf51fd-bcd6-d79a-0ef4-e52c9b5466dc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the token expires.
	//
	// example:
	//
	// 2018-01-25T00:01:02Z
	TokenExpireTime *string `json:"TokenExpireTime,omitempty" xml:"TokenExpireTime,omitempty"`
}

func (s GetParametersForImportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetParametersForImportResponseBody) GoString() string {
	return s.String()
}

func (s *GetParametersForImportResponseBody) SetImportToken(v string) *GetParametersForImportResponseBody {
	s.ImportToken = &v
	return s
}

func (s *GetParametersForImportResponseBody) SetKeyId(v string) *GetParametersForImportResponseBody {
	s.KeyId = &v
	return s
}

func (s *GetParametersForImportResponseBody) SetPublicKey(v string) *GetParametersForImportResponseBody {
	s.PublicKey = &v
	return s
}

func (s *GetParametersForImportResponseBody) SetRequestId(v string) *GetParametersForImportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetParametersForImportResponseBody) SetTokenExpireTime(v string) *GetParametersForImportResponseBody {
	s.TokenExpireTime = &v
	return s
}

type GetParametersForImportResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetParametersForImportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetParametersForImportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetParametersForImportResponse) GoString() string {
	return s.String()
}

func (s *GetParametersForImportResponse) SetHeaders(v map[string]*string) *GetParametersForImportResponse {
	s.Headers = v
	return s
}

func (s *GetParametersForImportResponse) SetStatusCode(v int32) *GetParametersForImportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetParametersForImportResponse) SetBody(v *GetParametersForImportResponseBody) *GetParametersForImportResponse {
	s.Body = v
	return s
}

type GetPublicKeyRequest struct {
	// The globally unique ID of the CMK. You can also set this parameter to an alias that is bound to the CMK. For more information, see [Use aliases](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The globally unique ID of the CMK version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s GetPublicKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPublicKeyRequest) GoString() string {
	return s.String()
}

func (s *GetPublicKeyRequest) SetKeyId(v string) *GetPublicKeyRequest {
	s.KeyId = &v
	return s
}

func (s *GetPublicKeyRequest) SetKeyVersionId(v string) *GetPublicKeyRequest {
	s.KeyVersionId = &v
	return s
}

type GetPublicKeyResponseBody struct {
	// The globally unique ID of the CMK.
	//
	// >  If you set the KeyId parameter to the alias of the CMK, the ID of the CMK to which the alias is bound is returned.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The version of the CMK that is used to encrypt the plaintext.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The public key returned in the PEM format.
	//
	// example:
	//
	// -----BEGIN PUBLIC KEY-----\\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAs5Yu9AEgATN2/e3nUz1K\\nEy6ng8MSPutcse2/VECG/NUF9C6D4IsJ64ShzY3dcn34WYzTOe916eMJFxyrNrSw\\nHtc4UOR5AvaoRrfpgu2uq+i70/ZXrWL+pGb1hgZV8cWheIHMxwrR3IiQlM5qN7EF\\n9BdyWtyBfUGsp0Bn1VqlPc5G0x0a9xU2z9YtP994yDenNVIoIQ6Cov1lIEuwXAb2\\n7boC41ePXwD0JWt41sP+rgCmpjBx00puIG+IlnoReEgI1ZGYmK98GgA/XzmNjZiD\\nyvXJZAcM33Ue85+PkR5iHTtSEbi4QAoqpJabprUzz3Fin2j1dRrcacxGb7p31A9c\\nJQIDAQAB\\n-----END PUBLIC KEY-----\\n
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 475f1620-b9d3-4d35-b5c6-3fbdd941423d
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPublicKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPublicKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublicKeyResponseBody) SetKeyId(v string) *GetPublicKeyResponseBody {
	s.KeyId = &v
	return s
}

func (s *GetPublicKeyResponseBody) SetKeyVersionId(v string) *GetPublicKeyResponseBody {
	s.KeyVersionId = &v
	return s
}

func (s *GetPublicKeyResponseBody) SetPublicKey(v string) *GetPublicKeyResponseBody {
	s.PublicKey = &v
	return s
}

func (s *GetPublicKeyResponseBody) SetRequestId(v string) *GetPublicKeyResponseBody {
	s.RequestId = &v
	return s
}

type GetPublicKeyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPublicKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPublicKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPublicKeyResponse) GoString() string {
	return s.String()
}

func (s *GetPublicKeyResponse) SetHeaders(v map[string]*string) *GetPublicKeyResponse {
	s.Headers = v
	return s
}

func (s *GetPublicKeyResponse) SetStatusCode(v int32) *GetPublicKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPublicKeyResponse) SetBody(v *GetPublicKeyResponseBody) *GetPublicKeyResponse {
	s.Body = v
	return s
}

type GetRandomPasswordRequest struct {
	// The characters that are not included in the password to be generated.
	//
	// Valid values:
	//
	// ` Valid characters: 0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ! \\"#$%&\\"()*+,-. /:;<=>? @[\\] your_project_id} ~  `.
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// ABCabc
	ExcludeCharacters *string `json:"ExcludeCharacters,omitempty" xml:"ExcludeCharacters,omitempty"`
	// Specifies whether to exclude lowercase letters.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	ExcludeLowercase *string `json:"ExcludeLowercase,omitempty" xml:"ExcludeLowercase,omitempty"`
	// Specifies whether to exclude digits.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	ExcludeNumbers *string `json:"ExcludeNumbers,omitempty" xml:"ExcludeNumbers,omitempty"`
	// Specifies whether to exclude special characters.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	ExcludePunctuation *string `json:"ExcludePunctuation,omitempty" xml:"ExcludePunctuation,omitempty"`
	// Specifies whether to exclude uppercase letters.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	ExcludeUppercase *string `json:"ExcludeUppercase,omitempty" xml:"ExcludeUppercase,omitempty"`
	// The number of bytes that the password to be generated contains.
	//
	// Valid values: 8 to 128.
	//
	// Default value: 32
	//
	// example:
	//
	// 32
	PasswordLength *string `json:"PasswordLength,omitempty" xml:"PasswordLength,omitempty"`
	// Specifies whether to include all the preceding character types.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	RequireEachIncludedType *string `json:"RequireEachIncludedType,omitempty" xml:"RequireEachIncludedType,omitempty"`
}

func (s GetRandomPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRandomPasswordRequest) GoString() string {
	return s.String()
}

func (s *GetRandomPasswordRequest) SetExcludeCharacters(v string) *GetRandomPasswordRequest {
	s.ExcludeCharacters = &v
	return s
}

func (s *GetRandomPasswordRequest) SetExcludeLowercase(v string) *GetRandomPasswordRequest {
	s.ExcludeLowercase = &v
	return s
}

func (s *GetRandomPasswordRequest) SetExcludeNumbers(v string) *GetRandomPasswordRequest {
	s.ExcludeNumbers = &v
	return s
}

func (s *GetRandomPasswordRequest) SetExcludePunctuation(v string) *GetRandomPasswordRequest {
	s.ExcludePunctuation = &v
	return s
}

func (s *GetRandomPasswordRequest) SetExcludeUppercase(v string) *GetRandomPasswordRequest {
	s.ExcludeUppercase = &v
	return s
}

func (s *GetRandomPasswordRequest) SetPasswordLength(v string) *GetRandomPasswordRequest {
	s.PasswordLength = &v
	return s
}

func (s *GetRandomPasswordRequest) SetRequireEachIncludedType(v string) *GetRandomPasswordRequest {
	s.RequireEachIncludedType = &v
	return s
}

type GetRandomPasswordResponseBody struct {
	// The generated random password.
	//
	// example:
	//
	// IxGn>NMmNB(y?iZ<Yc,_H/{2GC\\"U****
	RandomPassword *string `json:"RandomPassword,omitempty" xml:"RandomPassword,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6b0cbe25-5e33-467e-972e-7a83c6c97604
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRandomPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRandomPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *GetRandomPasswordResponseBody) SetRandomPassword(v string) *GetRandomPasswordResponseBody {
	s.RandomPassword = &v
	return s
}

func (s *GetRandomPasswordResponseBody) SetRequestId(v string) *GetRandomPasswordResponseBody {
	s.RequestId = &v
	return s
}

type GetRandomPasswordResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRandomPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRandomPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRandomPasswordResponse) GoString() string {
	return s.String()
}

func (s *GetRandomPasswordResponse) SetHeaders(v map[string]*string) *GetRandomPasswordResponse {
	s.Headers = v
	return s
}

func (s *GetRandomPasswordResponse) SetStatusCode(v int32) *GetRandomPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRandomPasswordResponse) SetBody(v *GetRandomPasswordResponseBody) *GetRandomPasswordResponse {
	s.Body = v
	return s
}

type GetSecretPolicyRequest struct {
	// example:
	//
	// default
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// secret_test
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s GetSecretPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSecretPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetSecretPolicyRequest) SetPolicyName(v string) *GetSecretPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *GetSecretPolicyRequest) SetSecretName(v string) *GetSecretPolicyRequest {
	s.SecretName = &v
	return s
}

type GetSecretPolicyResponseBody struct {
	// example:
	//
	// {"Version":"1","Statement": [{"Sid":"kms default secret policy","Effect":"Allow","Principal":{"RAM": ["acs:ram::119285303511****:*"]},"Action":["kms:*"],"Resource": ["*"] }] }
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// 381D5D33-BB8F-395F-8EE4-AE3BB4B523C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSecretPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSecretPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecretPolicyResponseBody) SetPolicy(v string) *GetSecretPolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *GetSecretPolicyResponseBody) SetRequestId(v string) *GetSecretPolicyResponseBody {
	s.RequestId = &v
	return s
}

type GetSecretPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecretPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecretPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSecretPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetSecretPolicyResponse) SetHeaders(v map[string]*string) *GetSecretPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetSecretPolicyResponse) SetStatusCode(v int32) *GetSecretPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecretPolicyResponse) SetBody(v *GetSecretPolicyResponseBody) *GetSecretPolicyResponse {
	s.Body = v
	return s
}

type GetSecretValueRequest struct {
	// Specifies whether to obtain the extended configuration of the secret. Valid values:
	//
	// 	- true
	//
	// 	- false: This is the default value.
	//
	// >  This parameter is ignored for a generic secret.
	//
	// example:
	//
	// true
	FetchExtendedConfig *bool `json:"FetchExtendedConfig,omitempty" xml:"FetchExtendedConfig,omitempty"`
	// The name of the secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The version number of the secret value. If you specify this parameter, Secrets Manager returns the secret value of the specified version.
	//
	// >  This parameter is ignored for a managed ApsaraDB RDS secret, a managed RAM secret, or a managed ECS secret.
	//
	// example:
	//
	// 00000000000000000000000000000001
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The stage label that marks the secret version. If you specify this parameter, Secrets Manager returns the secret value of the version that is marked with the specified stage label.
	//
	// Default value: ACSCurrent.
	//
	// >  For a managed ApsaraDB RDS secret, a managed RAM secret, or a managed ECS secret, Secrets Manager can return only the secret value of the version marked with ACSPrevious or ACSCurrent.
	//
	// example:
	//
	// ACSCurrent
	VersionStage *string `json:"VersionStage,omitempty" xml:"VersionStage,omitempty"`
}

func (s GetSecretValueRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSecretValueRequest) GoString() string {
	return s.String()
}

func (s *GetSecretValueRequest) SetFetchExtendedConfig(v bool) *GetSecretValueRequest {
	s.FetchExtendedConfig = &v
	return s
}

func (s *GetSecretValueRequest) SetSecretName(v string) *GetSecretValueRequest {
	s.SecretName = &v
	return s
}

func (s *GetSecretValueRequest) SetVersionId(v string) *GetSecretValueRequest {
	s.VersionId = &v
	return s
}

func (s *GetSecretValueRequest) SetVersionStage(v string) *GetSecretValueRequest {
	s.VersionStage = &v
	return s
}

type GetSecretValueResponseBody struct {
	// Indicates whether automatic rotation is enabled. Valid values:
	//
	// 	- Enabled: indicates that automatic rotation is enabled.
	//
	// 	- Disabled: indicates that automatic rotation is disabled.
	//
	// 	- Invalid: indicates that the status of automatic rotation is abnormal. In this case, Secrets Manager cannot automatically rotate the secret.
	//
	// >  This parameter is returned only for a managed ApsaraDB RDS secret, a managed RAM secret, or a managed ECS secret.
	//
	// example:
	//
	// Enabled
	AutomaticRotation *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	// The time when the secret was created.
	//
	// example:
	//
	// 2020-02-21T15:39:26Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The extended configuration of the secret.
	//
	// >  This parameter is returned if you set the FetchExtendedConfig parameter to true. This parameter is returned only for a managed ApsaraDB RDS secret, a managed RAM secret, or a managed ECS secret.
	//
	// example:
	//
	// {\\"SecretSubType\\":\\"SingleUser\\", \\"DBInstanceId\\":\\"rm-uf667446pc955****\\",  \\"CustomData\\":{} }
	ExtendedConfig *string `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	// The time when the last rotation was performed.
	//
	// >  This parameter is returned if the secret was rotated.
	//
	// example:
	//
	// 2020-07-05T08:22:03Z
	LastRotationDate *string `json:"LastRotationDate,omitempty" xml:"LastRotationDate,omitempty"`
	// The time when the next rotation will be performed.
	//
	// >  This parameter is returned if automatic rotation is enabled.
	//
	// example:
	//
	// 2020-07-06T18:22:03Z
	NextRotationDate *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6a3e9c36-1150-4881-84d3-eb8672fcafad
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The interval for automatic rotation.
	//
	// The value is in the `integer[unit]` format. The `unit` field has a fixed value of s. For example, if the value is 604800s, automatic rotation is performed at a 7-day interval.
	//
	// >  This parameter is returned if automatic rotation is enabled.
	//
	// example:
	//
	// 604800s
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	// The secret value. Secrets Manager decrypts the ciphertext of the secret value and returns the plaintext of the secret value in this parameter.
	//
	// 	- For a generic secret, the secret value of the specified version is returned.
	//
	// 	- For a managed ApsaraDB RDS secret, the value is returned in the following format:`{"AccountName":"","AccountPassword":""}` .
	//
	// 	- For a managed RAM secret, the secret value is returned in the following format: `{"AccessKeyId":"Adfdsfd","AccessKeySecret":"fdsfdsf","GenerateTimestamp": "2016-03-25T10:42:40Z"}`.
	//
	// 	- For a managed ECS secret, the secret value is returned in one of the following formats:
	//
	//     	- `{"UserName":"root","Password":"H5asdasdsads****"}`: The secret value is returned in this format if the ECS secret is a password.
	//
	//     	- `{"UserName":"root","PublicKey":"ssh-rsa ****mKwnVix9YTFY9Rs= imported-openssh-key","PrivateKey": "d6bee1cb-2e14-4277-ba6b-73786b21****"}`: The secret value is returned in this format is the ECS secret is a pair of SSH keys. The private key is in the Privacy Enhanced Mail (PEM) format.
	//
	// example:
	//
	// testdata1
	SecretData *string `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
	// The type of the secret value. Valid values:
	//
	// 	- text
	//
	// 	- binary
	//
	// example:
	//
	// binary
	SecretDataType *string `json:"SecretDataType,omitempty" xml:"SecretDataType,omitempty"`
	// The name of the secret.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The type of the secret. Valid values:
	//
	// 	- Generic: indicates a generic secret.
	//
	// 	- Rds: indicates a managed ApsaraDB RDS secret.
	//
	// 	- RAMCredentials: indicates a managed RAM secret.
	//
	// 	- ECS: indicates a managed ECS secret.
	//
	// example:
	//
	// Generic
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// The version number of the secret value.
	//
	// example:
	//
	// 00000000000000000000000000000001
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The stage labels that mark the secret versions.
	VersionStages *GetSecretValueResponseBodyVersionStages `json:"VersionStages,omitempty" xml:"VersionStages,omitempty" type:"Struct"`
}

func (s GetSecretValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSecretValueResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecretValueResponseBody) SetAutomaticRotation(v string) *GetSecretValueResponseBody {
	s.AutomaticRotation = &v
	return s
}

func (s *GetSecretValueResponseBody) SetCreateTime(v string) *GetSecretValueResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetSecretValueResponseBody) SetExtendedConfig(v string) *GetSecretValueResponseBody {
	s.ExtendedConfig = &v
	return s
}

func (s *GetSecretValueResponseBody) SetLastRotationDate(v string) *GetSecretValueResponseBody {
	s.LastRotationDate = &v
	return s
}

func (s *GetSecretValueResponseBody) SetNextRotationDate(v string) *GetSecretValueResponseBody {
	s.NextRotationDate = &v
	return s
}

func (s *GetSecretValueResponseBody) SetRequestId(v string) *GetSecretValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecretValueResponseBody) SetRotationInterval(v string) *GetSecretValueResponseBody {
	s.RotationInterval = &v
	return s
}

func (s *GetSecretValueResponseBody) SetSecretData(v string) *GetSecretValueResponseBody {
	s.SecretData = &v
	return s
}

func (s *GetSecretValueResponseBody) SetSecretDataType(v string) *GetSecretValueResponseBody {
	s.SecretDataType = &v
	return s
}

func (s *GetSecretValueResponseBody) SetSecretName(v string) *GetSecretValueResponseBody {
	s.SecretName = &v
	return s
}

func (s *GetSecretValueResponseBody) SetSecretType(v string) *GetSecretValueResponseBody {
	s.SecretType = &v
	return s
}

func (s *GetSecretValueResponseBody) SetVersionId(v string) *GetSecretValueResponseBody {
	s.VersionId = &v
	return s
}

func (s *GetSecretValueResponseBody) SetVersionStages(v *GetSecretValueResponseBodyVersionStages) *GetSecretValueResponseBody {
	s.VersionStages = v
	return s
}

type GetSecretValueResponseBodyVersionStages struct {
	VersionStage []*string `json:"VersionStage,omitempty" xml:"VersionStage,omitempty" type:"Repeated"`
}

func (s GetSecretValueResponseBodyVersionStages) String() string {
	return tea.Prettify(s)
}

func (s GetSecretValueResponseBodyVersionStages) GoString() string {
	return s.String()
}

func (s *GetSecretValueResponseBodyVersionStages) SetVersionStage(v []*string) *GetSecretValueResponseBodyVersionStages {
	s.VersionStage = v
	return s
}

type GetSecretValueResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecretValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecretValueResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSecretValueResponse) GoString() string {
	return s.String()
}

func (s *GetSecretValueResponse) SetHeaders(v map[string]*string) *GetSecretValueResponse {
	s.Headers = v
	return s
}

func (s *GetSecretValueResponse) SetStatusCode(v int32) *GetSecretValueResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecretValueResponse) SetBody(v *GetSecretValueResponseBody) *GetSecretValueResponse {
	s.Body = v
	return s
}

type ImportKeyMaterialRequest struct {
	// Use **GetParametersForImport*	- the Returned public key and the base64-encoded key material.
	//
	// This parameter is required.
	//
	// example:
	//
	// bCPZx7I6v6KXsqEpr2OXKxuj2CCRtKdwp75Bw+BGncYqBdfjFBYRtOE6HRlT0oeiRDWzwnw9OA54OL36smDJrq4Lo9x0CyYDiuKnRkcKtMtlzW0din7Pd7IlZWWRdVueiw2qpzl7PkUWQGTdsdbzpfJJQ+qj/cRIrk/E83UGyeyytSpgnb+lu0xEYcPajRyWNsbi98N3pqqQzHXNNHO2NJqHlnQgglqTiBEjkGeKFhfKmTc3vjulIdVa3EaVIN6lwWfgx+UUYSrvbA77WDYKlDsZ4SbK2/T7za9Tp1qU7Ynqba7OKGVVj7PMbiaO80AxWZnjUMYCgEp5w7V+seOXqw==
	EncryptedKeyMaterial *string `json:"EncryptedKeyMaterial,omitempty" xml:"EncryptedKeyMaterial,omitempty"`
	// By calling **GetParametersForImport*	- the import token.
	//
	// This parameter is required.
	//
	// example:
	//
	// Base64String
	ImportToken *string `json:"ImportToken,omitempty" xml:"ImportToken,omitempty"`
	// The ID of the CMK to be imported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The time when the key material expires.
	//
	// If this parameter is not specified or set this parameter to 0, the key material does not expire.
	//
	// >  The value cannot be earlier than the time when the API is called (based on the server time).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	KeyMaterialExpireUnix *int64 `json:"KeyMaterialExpireUnix,omitempty" xml:"KeyMaterialExpireUnix,omitempty"`
}

func (s ImportKeyMaterialRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportKeyMaterialRequest) GoString() string {
	return s.String()
}

func (s *ImportKeyMaterialRequest) SetEncryptedKeyMaterial(v string) *ImportKeyMaterialRequest {
	s.EncryptedKeyMaterial = &v
	return s
}

func (s *ImportKeyMaterialRequest) SetImportToken(v string) *ImportKeyMaterialRequest {
	s.ImportToken = &v
	return s
}

func (s *ImportKeyMaterialRequest) SetKeyId(v string) *ImportKeyMaterialRequest {
	s.KeyId = &v
	return s
}

func (s *ImportKeyMaterialRequest) SetKeyMaterialExpireUnix(v int64) *ImportKeyMaterialRequest {
	s.KeyMaterialExpireUnix = &v
	return s
}

type ImportKeyMaterialResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ec1017cf-ead4-f3ca-babc-c3b34f3dbecb
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportKeyMaterialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportKeyMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *ImportKeyMaterialResponseBody) SetRequestId(v string) *ImportKeyMaterialResponseBody {
	s.RequestId = &v
	return s
}

type ImportKeyMaterialResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportKeyMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportKeyMaterialResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportKeyMaterialResponse) GoString() string {
	return s.String()
}

func (s *ImportKeyMaterialResponse) SetHeaders(v map[string]*string) *ImportKeyMaterialResponse {
	s.Headers = v
	return s
}

func (s *ImportKeyMaterialResponse) SetStatusCode(v int32) *ImportKeyMaterialResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportKeyMaterialResponse) SetBody(v *ImportKeyMaterialResponseBody) *ImportKeyMaterialResponse {
	s.Body = v
	return s
}

type ListAliasesRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 0 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAliasesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesRequest) GoString() string {
	return s.String()
}

func (s *ListAliasesRequest) SetPageNumber(v int32) *ListAliasesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAliasesRequest) SetPageSize(v int32) *ListAliasesRequest {
	s.PageSize = &v
	return s
}

type ListAliasesResponseBody struct {
	// The alias of the user.
	Aliases *ListAliasesResponseBodyAliases `json:"Aliases,omitempty" xml:"Aliases,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1b57992c-834b-4811-a889-f8bac1ba0353
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned aliases.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAliasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAliasesResponseBody) SetAliases(v *ListAliasesResponseBodyAliases) *ListAliasesResponseBody {
	s.Aliases = v
	return s
}

func (s *ListAliasesResponseBody) SetPageNumber(v int32) *ListAliasesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAliasesResponseBody) SetPageSize(v int32) *ListAliasesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAliasesResponseBody) SetRequestId(v string) *ListAliasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAliasesResponseBody) SetTotalCount(v int32) *ListAliasesResponseBody {
	s.TotalCount = &v
	return s
}

type ListAliasesResponseBodyAliases struct {
	Alias []*ListAliasesResponseBodyAliasesAlias `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Repeated"`
}

func (s ListAliasesResponseBodyAliases) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesResponseBodyAliases) GoString() string {
	return s.String()
}

func (s *ListAliasesResponseBodyAliases) SetAlias(v []*ListAliasesResponseBodyAliasesAlias) *ListAliasesResponseBodyAliases {
	s.Alias = v
	return s
}

type ListAliasesResponseBodyAliasesAlias struct {
	// The Alibaba Cloud Resource Name (ARN) of the alias.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:123456:alias/ExampleAlias1
	AliasArn *string `json:"AliasArn,omitempty" xml:"AliasArn,omitempty"`
	// The ID of the alias.
	//
	// example:
	//
	// alias/ExampleAlias1
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The CMK to which the alias belongs.
	//
	// example:
	//
	// 08c33a6f-4e0a-4a1b-a3fa-7ddfa1d****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s ListAliasesResponseBodyAliasesAlias) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesResponseBodyAliasesAlias) GoString() string {
	return s.String()
}

func (s *ListAliasesResponseBodyAliasesAlias) SetAliasArn(v string) *ListAliasesResponseBodyAliasesAlias {
	s.AliasArn = &v
	return s
}

func (s *ListAliasesResponseBodyAliasesAlias) SetAliasName(v string) *ListAliasesResponseBodyAliasesAlias {
	s.AliasName = &v
	return s
}

func (s *ListAliasesResponseBodyAliasesAlias) SetKeyId(v string) *ListAliasesResponseBodyAliasesAlias {
	s.KeyId = &v
	return s
}

type ListAliasesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAliasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAliasesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesResponse) GoString() string {
	return s.String()
}

func (s *ListAliasesResponse) SetHeaders(v map[string]*string) *ListAliasesResponse {
	s.Headers = v
	return s
}

func (s *ListAliasesResponse) SetStatusCode(v int32) *ListAliasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAliasesResponse) SetBody(v *ListAliasesResponseBody) *ListAliasesResponse {
	s.Body = v
	return s
}

type ListAliasesByKeyIdRequest struct {
	// The globally unique ID of the CMK.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The number of the page to return.
	//
	// Valid values: an integer that is greater than 0.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 0 to 101.
	//
	// Default value: 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAliasesByKeyIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesByKeyIdRequest) GoString() string {
	return s.String()
}

func (s *ListAliasesByKeyIdRequest) SetKeyId(v string) *ListAliasesByKeyIdRequest {
	s.KeyId = &v
	return s
}

func (s *ListAliasesByKeyIdRequest) SetPageNumber(v int32) *ListAliasesByKeyIdRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAliasesByKeyIdRequest) SetPageSize(v int32) *ListAliasesByKeyIdRequest {
	s.PageSize = &v
	return s
}

type ListAliasesByKeyIdResponseBody struct {
	// An array that consists of aliases.
	Aliases *ListAliasesByKeyIdResponseBodyAliases `json:"Aliases,omitempty" xml:"Aliases,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 1b57992c-834b-4811-a889-f8bac1ba0353
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned CMKs.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAliasesByKeyIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesByKeyIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListAliasesByKeyIdResponseBody) SetAliases(v *ListAliasesByKeyIdResponseBodyAliases) *ListAliasesByKeyIdResponseBody {
	s.Aliases = v
	return s
}

func (s *ListAliasesByKeyIdResponseBody) SetPageNumber(v int32) *ListAliasesByKeyIdResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAliasesByKeyIdResponseBody) SetPageSize(v int32) *ListAliasesByKeyIdResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAliasesByKeyIdResponseBody) SetRequestId(v string) *ListAliasesByKeyIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAliasesByKeyIdResponseBody) SetTotalCount(v int32) *ListAliasesByKeyIdResponseBody {
	s.TotalCount = &v
	return s
}

type ListAliasesByKeyIdResponseBodyAliases struct {
	Alias []*ListAliasesByKeyIdResponseBodyAliasesAlias `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Repeated"`
}

func (s ListAliasesByKeyIdResponseBodyAliases) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesByKeyIdResponseBodyAliases) GoString() string {
	return s.String()
}

func (s *ListAliasesByKeyIdResponseBodyAliases) SetAlias(v []*ListAliasesByKeyIdResponseBodyAliasesAlias) *ListAliasesByKeyIdResponseBodyAliases {
	s.Alias = v
	return s
}

type ListAliasesByKeyIdResponseBodyAliasesAlias struct {
	// The Alibaba Cloud Resource Name (ARN) of the alias.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:123456:alias/ExampleAlias1
	AliasArn *string `json:"AliasArn,omitempty" xml:"AliasArn,omitempty"`
	// The ID of the alias.
	//
	// example:
	//
	// alias/ExampleAlias1
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The CMK to which an alias is bound.
	//
	// example:
	//
	// 08c33a6f-4e0a-4a1b-a3fa-7ddfa1d4****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s ListAliasesByKeyIdResponseBodyAliasesAlias) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesByKeyIdResponseBodyAliasesAlias) GoString() string {
	return s.String()
}

func (s *ListAliasesByKeyIdResponseBodyAliasesAlias) SetAliasArn(v string) *ListAliasesByKeyIdResponseBodyAliasesAlias {
	s.AliasArn = &v
	return s
}

func (s *ListAliasesByKeyIdResponseBodyAliasesAlias) SetAliasName(v string) *ListAliasesByKeyIdResponseBodyAliasesAlias {
	s.AliasName = &v
	return s
}

func (s *ListAliasesByKeyIdResponseBodyAliasesAlias) SetKeyId(v string) *ListAliasesByKeyIdResponseBodyAliasesAlias {
	s.KeyId = &v
	return s
}

type ListAliasesByKeyIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAliasesByKeyIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAliasesByKeyIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesByKeyIdResponse) GoString() string {
	return s.String()
}

func (s *ListAliasesByKeyIdResponse) SetHeaders(v map[string]*string) *ListAliasesByKeyIdResponse {
	s.Headers = v
	return s
}

func (s *ListAliasesByKeyIdResponse) SetStatusCode(v int32) *ListAliasesByKeyIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAliasesByKeyIdResponse) SetBody(v *ListAliasesByKeyIdResponseBody) *ListAliasesByKeyIdResponse {
	s.Body = v
	return s
}

type ListApplicationAccessPointsRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListApplicationAccessPointsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationAccessPointsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationAccessPointsRequest) SetPageNumber(v int32) *ListApplicationAccessPointsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationAccessPointsRequest) SetPageSize(v int32) *ListApplicationAccessPointsRequest {
	s.PageSize = &v
	return s
}

type ListApplicationAccessPointsResponseBody struct {
	// A list of AAPs.
	ApplicationAccessPoints *ListApplicationAccessPointsResponseBodyApplicationAccessPoints `json:"ApplicationAccessPoints,omitempty" xml:"ApplicationAccessPoints,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// bcfefe15-46f0-44a3-bd96-3d422474b71a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationAccessPointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationAccessPointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationAccessPointsResponseBody) SetApplicationAccessPoints(v *ListApplicationAccessPointsResponseBodyApplicationAccessPoints) *ListApplicationAccessPointsResponseBody {
	s.ApplicationAccessPoints = v
	return s
}

func (s *ListApplicationAccessPointsResponseBody) SetPageNumber(v int32) *ListApplicationAccessPointsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationAccessPointsResponseBody) SetPageSize(v int32) *ListApplicationAccessPointsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApplicationAccessPointsResponseBody) SetRequestId(v string) *ListApplicationAccessPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationAccessPointsResponseBody) SetTotalCount(v int32) *ListApplicationAccessPointsResponseBody {
	s.TotalCount = &v
	return s
}

type ListApplicationAccessPointsResponseBodyApplicationAccessPoints struct {
	ApplicationAccessPoint []*ListApplicationAccessPointsResponseBodyApplicationAccessPointsApplicationAccessPoint `json:"ApplicationAccessPoint,omitempty" xml:"ApplicationAccessPoint,omitempty" type:"Repeated"`
}

func (s ListApplicationAccessPointsResponseBodyApplicationAccessPoints) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationAccessPointsResponseBodyApplicationAccessPoints) GoString() string {
	return s.String()
}

func (s *ListApplicationAccessPointsResponseBodyApplicationAccessPoints) SetApplicationAccessPoint(v []*ListApplicationAccessPointsResponseBodyApplicationAccessPointsApplicationAccessPoint) *ListApplicationAccessPointsResponseBodyApplicationAccessPoints {
	s.ApplicationAccessPoint = v
	return s
}

type ListApplicationAccessPointsResponseBodyApplicationAccessPointsApplicationAccessPoint struct {
	// The authentication method.
	//
	// example:
	//
	// ClientKey
	AuthenticationMethod *string `json:"AuthenticationMethod,omitempty" xml:"AuthenticationMethod,omitempty"`
	// The name of the AAP.
	//
	// example:
	//
	// aap_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListApplicationAccessPointsResponseBodyApplicationAccessPointsApplicationAccessPoint) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationAccessPointsResponseBodyApplicationAccessPointsApplicationAccessPoint) GoString() string {
	return s.String()
}

func (s *ListApplicationAccessPointsResponseBodyApplicationAccessPointsApplicationAccessPoint) SetAuthenticationMethod(v string) *ListApplicationAccessPointsResponseBodyApplicationAccessPointsApplicationAccessPoint {
	s.AuthenticationMethod = &v
	return s
}

func (s *ListApplicationAccessPointsResponseBodyApplicationAccessPointsApplicationAccessPoint) SetName(v string) *ListApplicationAccessPointsResponseBodyApplicationAccessPointsApplicationAccessPoint {
	s.Name = &v
	return s
}

type ListApplicationAccessPointsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationAccessPointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationAccessPointsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationAccessPointsResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationAccessPointsResponse) SetHeaders(v map[string]*string) *ListApplicationAccessPointsResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationAccessPointsResponse) SetStatusCode(v int32) *ListApplicationAccessPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationAccessPointsResponse) SetBody(v *ListApplicationAccessPointsResponseBody) *ListApplicationAccessPointsResponse {
	s.Body = v
	return s
}

type ListClientKeysRequest struct {
	// The name of the application access point (AAP).
	//
	// example:
	//
	// aap_test
	AapName *string `json:"AapName,omitempty" xml:"AapName,omitempty"`
}

func (s ListClientKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClientKeysRequest) GoString() string {
	return s.String()
}

func (s *ListClientKeysRequest) SetAapName(v string) *ListClientKeysRequest {
	s.AapName = &v
	return s
}

type ListClientKeysResponseBody struct {
	// A list of client keys.
	ClientKeys []*ListClientKeysResponseBodyClientKeys `json:"ClientKeys,omitempty" xml:"ClientKeys,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 2312e45f-b2fa-4c34-ad94-3eca50932916
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClientKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClientKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListClientKeysResponseBody) SetClientKeys(v []*ListClientKeysResponseBodyClientKeys) *ListClientKeysResponseBody {
	s.ClientKeys = v
	return s
}

func (s *ListClientKeysResponseBody) SetRequestId(v string) *ListClientKeysResponseBody {
	s.RequestId = &v
	return s
}

type ListClientKeysResponseBodyClientKeys struct {
	// The name of the AAP.
	//
	// example:
	//
	// aap_test
	AapName *string `json:"AapName,omitempty" xml:"AapName,omitempty"`
	// The ID of the client key.
	//
	// example:
	//
	// KAAP.66abf237-63f6-4625-b8cf-47e1086e****
	ClientKeyId *string `json:"ClientKeyId,omitempty" xml:"ClientKeyId,omitempty"`
	// The time when the client key was created.
	//
	// example:
	//
	// 2023-08-31T09:14:38Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The private key algorithm of the client key.
	//
	// example:
	//
	// RSA_2048
	KeyAlgorithm *string `json:"KeyAlgorithm,omitempty" xml:"KeyAlgorithm,omitempty"`
	// The provider of the client key.
	//
	// Currently, only KMS is supported. The value is fixed as KMS_PROVIDED.
	//
	// example:
	//
	// KMS_PROVIDED
	KeyOrigin *string `json:"KeyOrigin,omitempty" xml:"KeyOrigin,omitempty"`
	// The end of the validity period of the client key.
	//
	// example:
	//
	// 2028-08-31T17:14:33Z
	NotAfter *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The beginning of the validity period of the client key.
	//
	// example:
	//
	// 2023-08-31T17:14:33Z
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The public key of the client key.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\nMIIDcjCCAlqgAwIBAgIQT/sAVRxwYp54mrw****-----END CERTIFICATE-----
	PublicKeyData *string `json:"PublicKeyData,omitempty" xml:"PublicKeyData,omitempty"`
}

func (s ListClientKeysResponseBodyClientKeys) String() string {
	return tea.Prettify(s)
}

func (s ListClientKeysResponseBodyClientKeys) GoString() string {
	return s.String()
}

func (s *ListClientKeysResponseBodyClientKeys) SetAapName(v string) *ListClientKeysResponseBodyClientKeys {
	s.AapName = &v
	return s
}

func (s *ListClientKeysResponseBodyClientKeys) SetClientKeyId(v string) *ListClientKeysResponseBodyClientKeys {
	s.ClientKeyId = &v
	return s
}

func (s *ListClientKeysResponseBodyClientKeys) SetCreateTime(v string) *ListClientKeysResponseBodyClientKeys {
	s.CreateTime = &v
	return s
}

func (s *ListClientKeysResponseBodyClientKeys) SetKeyAlgorithm(v string) *ListClientKeysResponseBodyClientKeys {
	s.KeyAlgorithm = &v
	return s
}

func (s *ListClientKeysResponseBodyClientKeys) SetKeyOrigin(v string) *ListClientKeysResponseBodyClientKeys {
	s.KeyOrigin = &v
	return s
}

func (s *ListClientKeysResponseBodyClientKeys) SetNotAfter(v string) *ListClientKeysResponseBodyClientKeys {
	s.NotAfter = &v
	return s
}

func (s *ListClientKeysResponseBodyClientKeys) SetNotBefore(v string) *ListClientKeysResponseBodyClientKeys {
	s.NotBefore = &v
	return s
}

func (s *ListClientKeysResponseBodyClientKeys) SetPublicKeyData(v string) *ListClientKeysResponseBodyClientKeys {
	s.PublicKeyData = &v
	return s
}

type ListClientKeysResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClientKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClientKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClientKeysResponse) GoString() string {
	return s.String()
}

func (s *ListClientKeysResponse) SetHeaders(v map[string]*string) *ListClientKeysResponse {
	s.Headers = v
	return s
}

func (s *ListClientKeysResponse) SetStatusCode(v int32) *ListClientKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClientKeysResponse) SetBody(v *ListClientKeysResponseBody) *ListClientKeysResponse {
	s.Body = v
	return s
}

type ListKeyVersionsRequest struct {
	// The globally unique ID of the CMK. You can also set this parameter to an alias that is bound to the CMK. For more information, see [Use aliases](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0b30658a-ed1a-4922-b8f7-a673ca9c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 0 to 101.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListKeyVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListKeyVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListKeyVersionsRequest) SetKeyId(v string) *ListKeyVersionsRequest {
	s.KeyId = &v
	return s
}

func (s *ListKeyVersionsRequest) SetPageNumber(v int32) *ListKeyVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListKeyVersionsRequest) SetPageSize(v int32) *ListKeyVersionsRequest {
	s.PageSize = &v
	return s
}

type ListKeyVersionsResponseBody struct {
	// An array that consists of key versions.
	KeyVersions *ListKeyVersionsResponseBodyKeyVersions `json:"KeyVersions,omitempty" xml:"KeyVersions,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// f71204c4-53cd-4eea-b405-653ba2db7e86
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned key versions.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKeyVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListKeyVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListKeyVersionsResponseBody) SetKeyVersions(v *ListKeyVersionsResponseBodyKeyVersions) *ListKeyVersionsResponseBody {
	s.KeyVersions = v
	return s
}

func (s *ListKeyVersionsResponseBody) SetPageNumber(v int32) *ListKeyVersionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListKeyVersionsResponseBody) SetPageSize(v int32) *ListKeyVersionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListKeyVersionsResponseBody) SetRequestId(v string) *ListKeyVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKeyVersionsResponseBody) SetTotalCount(v int32) *ListKeyVersionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListKeyVersionsResponseBodyKeyVersions struct {
	KeyVersion []*ListKeyVersionsResponseBodyKeyVersionsKeyVersion `json:"KeyVersion,omitempty" xml:"KeyVersion,omitempty" type:"Repeated"`
}

func (s ListKeyVersionsResponseBodyKeyVersions) String() string {
	return tea.Prettify(s)
}

func (s ListKeyVersionsResponseBodyKeyVersions) GoString() string {
	return s.String()
}

func (s *ListKeyVersionsResponseBodyKeyVersions) SetKeyVersion(v []*ListKeyVersionsResponseBodyKeyVersionsKeyVersion) *ListKeyVersionsResponseBodyKeyVersions {
	s.KeyVersion = v
	return s
}

type ListKeyVersionsResponseBodyKeyVersionsKeyVersion struct {
	// The date and time when the CMK version was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2016-03-25T10:42:40Z
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// The globally unique ID of the CMK.
	//
	// >  If you set the KeyId parameter to the alias of the CMK, the ID of the CMK to which the alias is bound is returned.
	//
	// example:
	//
	// 0b30658a-ed1a-4922-b8f7-a673ca9c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The globally unique ID of the CMK version.
	//
	// example:
	//
	// 1e3304fd-68ac-4d5b-8886-ae5f01a1****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s ListKeyVersionsResponseBodyKeyVersionsKeyVersion) String() string {
	return tea.Prettify(s)
}

func (s ListKeyVersionsResponseBodyKeyVersionsKeyVersion) GoString() string {
	return s.String()
}

func (s *ListKeyVersionsResponseBodyKeyVersionsKeyVersion) SetCreationDate(v string) *ListKeyVersionsResponseBodyKeyVersionsKeyVersion {
	s.CreationDate = &v
	return s
}

func (s *ListKeyVersionsResponseBodyKeyVersionsKeyVersion) SetKeyId(v string) *ListKeyVersionsResponseBodyKeyVersionsKeyVersion {
	s.KeyId = &v
	return s
}

func (s *ListKeyVersionsResponseBodyKeyVersionsKeyVersion) SetKeyVersionId(v string) *ListKeyVersionsResponseBodyKeyVersionsKeyVersion {
	s.KeyVersionId = &v
	return s
}

type ListKeyVersionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKeyVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKeyVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListKeyVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListKeyVersionsResponse) SetHeaders(v map[string]*string) *ListKeyVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListKeyVersionsResponse) SetStatusCode(v int32) *ListKeyVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKeyVersionsResponse) SetBody(v *ListKeyVersionsResponseBody) *ListKeyVersionsResponse {
	s.Body = v
	return s
}

type ListKeysRequest struct {
	// The CMK filter. The filter consists of one or more key-value pairs. You can specify a maximum of 10 key-value pairs.
	//
	// 	- Key
	//
	//     	- Description: the property that you want to filter.
	//
	//     	- Type: string.
	//
	//     	- Valid values:
	//
	//         	- KeyState: the status of the CMK.
	//
	//         	- KeySpec: the type of the CMK.
	//
	//         	- KeyUsage: the usage of the CMK.
	//
	//         	- ProtectionLevel: the protection level.
	//
	//         	- CreatorType: the type of the creator.
	//
	// 	- Values
	//
	//     	- Description: the value to be included after filtering.
	//
	//     	- Format: string array.
	//
	//     	- Length: 0 to 10.
	//
	//     	- Valid values:
	//
	//         	- When Key is set to KeyState, the value can be Enabled, Disabled, PendingDeletion, or PendingImport.
	//
	//         	- When Key is set to KeySpec, the value can be Aliyun_AES_256, Aliyun_SM4, RSA_2048, EC_P256, EC_P256K, or EC_SM2.
	//
	//             Note: You can create CMKs of the EC_SM2 or Aliyun_SM4 type only in regions where State Cryptography Administration (SCA)-certified managed HSMs reside. For more information about the regions, see [Supported regions](https://help.aliyun.com/document_detail/125803.html). If your region does not support EC_SM2 or Aliyun_SM4, the two values are ignored if they are specified.
	//
	//         	- When Key is set to KeyUsage, the value can be ENCRYPT/DECRYPT or SIGN/VERIFY. ENCRYPT/DECRYPT indicates that the CMK is used to encrypt and decrypt data. SIGN/VERIFY indicates that the CMK is used to generate and verify digital signatures.
	//
	//         	- When Key is set to ProtectionLevel, the value can be SOFTWARE (software) or HSM (hardware).
	//
	//             You can set ProtectionLevel to HSM in only specific regions. For more information about the regions, see [Supported regions](https://help.aliyun.com/document_detail/125803.html). If your region does not support the value HSM, the value is ignored if the value is specified.
	//
	//         	- If Key is set to CreatorType, the value can be User or Service. User indicates that CMKs created by the current account are queried. Service indicates that CMKs automatically created by other cloud services authorized by the current account are queried.
	//
	// The logical relationship between different keys is AND, and the logical relationship between multiple items in the same key is OR. Example:
	//
	// `[ {"Key":"KeyState", "Values":["Enabled","Disabled"]}, {"Key":"KeyState", "Values":["PendingDeletion"]}, {"Key":"KeySpec", "Values":["Aliyun_AES_256"]}]`. In this example, the semantics are:`(KeyState=Enabled OR KeyState=Disabled OR KeyState=PendingDeletion) AND (KeySpec=Aliyun_AES_ 256)`.
	//
	// example:
	//
	// [{"Key":"KeyState", "Values":["Enabled","Disabled"]}]
	Filters *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListKeysRequest) GoString() string {
	return s.String()
}

func (s *ListKeysRequest) SetFilters(v string) *ListKeysRequest {
	s.Filters = &v
	return s
}

func (s *ListKeysRequest) SetPageNumber(v int32) *ListKeysRequest {
	s.PageNumber = &v
	return s
}

func (s *ListKeysRequest) SetPageSize(v int32) *ListKeysRequest {
	s.PageSize = &v
	return s
}

type ListKeysResponseBody struct {
	// An array that consists of the CMKs of the current Alibaba Cloud account in the current region.
	Keys *ListKeysResponseBodyKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 8252db58-2036-408c-a3d5-56e656dc2551
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of CMKs.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListKeysResponseBody) SetKeys(v *ListKeysResponseBodyKeys) *ListKeysResponseBody {
	s.Keys = v
	return s
}

func (s *ListKeysResponseBody) SetPageNumber(v int32) *ListKeysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListKeysResponseBody) SetPageSize(v int32) *ListKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListKeysResponseBody) SetRequestId(v string) *ListKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKeysResponseBody) SetTotalCount(v int32) *ListKeysResponseBody {
	s.TotalCount = &v
	return s
}

type ListKeysResponseBodyKeys struct {
	Key []*ListKeysResponseBodyKeysKey `json:"Key,omitempty" xml:"Key,omitempty" type:"Repeated"`
}

func (s ListKeysResponseBodyKeys) String() string {
	return tea.Prettify(s)
}

func (s ListKeysResponseBodyKeys) GoString() string {
	return s.String()
}

func (s *ListKeysResponseBodyKeys) SetKey(v []*ListKeysResponseBodyKeysKey) *ListKeysResponseBodyKeys {
	s.Key = v
	return s
}

type ListKeysResponseBodyKeysKey struct {
	// The Alibaba Cloud Resource Name (ARN) of the CMK.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:123456:key/80e9409f-78fa-42ab-84bd-83f40c81****
	KeyArn *string `json:"KeyArn,omitempty" xml:"KeyArn,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// example:
	//
	// 08c33a6f-4e0a-4a1b-a3fa-7ddfa1d4****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s ListKeysResponseBodyKeysKey) String() string {
	return tea.Prettify(s)
}

func (s ListKeysResponseBodyKeysKey) GoString() string {
	return s.String()
}

func (s *ListKeysResponseBodyKeysKey) SetKeyArn(v string) *ListKeysResponseBodyKeysKey {
	s.KeyArn = &v
	return s
}

func (s *ListKeysResponseBodyKeysKey) SetKeyId(v string) *ListKeysResponseBodyKeysKey {
	s.KeyId = &v
	return s
}

type ListKeysResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListKeysResponse) GoString() string {
	return s.String()
}

func (s *ListKeysResponse) SetHeaders(v map[string]*string) *ListKeysResponse {
	s.Headers = v
	return s
}

func (s *ListKeysResponse) SetStatusCode(v int32) *ListKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKeysResponse) SetBody(v *ListKeysResponseBody) *ListKeysResponse {
	s.Body = v
	return s
}

type ListKmsInstancesRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListKmsInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListKmsInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListKmsInstancesRequest) SetPageNumber(v int32) *ListKmsInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListKmsInstancesRequest) SetPageSize(v int32) *ListKmsInstancesRequest {
	s.PageSize = &v
	return s
}

type ListKmsInstancesResponseBody struct {
	// A list of KMS instances.
	KmsInstances *ListKmsInstancesResponseBodyKmsInstances `json:"KmsInstances,omitempty" xml:"KmsInstances,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// d3eca5c8-a856-4347-8eb6-e1898c3fda2e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of KMS instances.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKmsInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListKmsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListKmsInstancesResponseBody) SetKmsInstances(v *ListKmsInstancesResponseBodyKmsInstances) *ListKmsInstancesResponseBody {
	s.KmsInstances = v
	return s
}

func (s *ListKmsInstancesResponseBody) SetPageNumber(v int64) *ListKmsInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListKmsInstancesResponseBody) SetPageSize(v int64) *ListKmsInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListKmsInstancesResponseBody) SetRequestId(v string) *ListKmsInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKmsInstancesResponseBody) SetTotalCount(v int64) *ListKmsInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListKmsInstancesResponseBodyKmsInstances struct {
	KmsInstance []*ListKmsInstancesResponseBodyKmsInstancesKmsInstance `json:"KmsInstance,omitempty" xml:"KmsInstance,omitempty" type:"Repeated"`
}

func (s ListKmsInstancesResponseBodyKmsInstances) String() string {
	return tea.Prettify(s)
}

func (s ListKmsInstancesResponseBodyKmsInstances) GoString() string {
	return s.String()
}

func (s *ListKmsInstancesResponseBodyKmsInstances) SetKmsInstance(v []*ListKmsInstancesResponseBodyKmsInstancesKmsInstance) *ListKmsInstancesResponseBodyKmsInstances {
	s.KmsInstance = v
	return s
}

type ListKmsInstancesResponseBodyKmsInstancesKmsInstance struct {
	// The ARN of the KMS instance.
	//
	// example:
	//
	// acs:kms:pre-hangzhou:120708975881****:keystore/kst-phzz64c9f84eo32dbs****
	KmsInstanceArn *string `json:"KmsInstanceArn,omitempty" xml:"KmsInstanceArn,omitempty"`
	// The ID of the KMS instance.
	//
	// example:
	//
	// kst-phzz64c9f84eo32dbs****
	KmsInstanceId *string `json:"KmsInstanceId,omitempty" xml:"KmsInstanceId,omitempty"`
}

func (s ListKmsInstancesResponseBodyKmsInstancesKmsInstance) String() string {
	return tea.Prettify(s)
}

func (s ListKmsInstancesResponseBodyKmsInstancesKmsInstance) GoString() string {
	return s.String()
}

func (s *ListKmsInstancesResponseBodyKmsInstancesKmsInstance) SetKmsInstanceArn(v string) *ListKmsInstancesResponseBodyKmsInstancesKmsInstance {
	s.KmsInstanceArn = &v
	return s
}

func (s *ListKmsInstancesResponseBodyKmsInstancesKmsInstance) SetKmsInstanceId(v string) *ListKmsInstancesResponseBodyKmsInstancesKmsInstance {
	s.KmsInstanceId = &v
	return s
}

type ListKmsInstancesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKmsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKmsInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListKmsInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListKmsInstancesResponse) SetHeaders(v map[string]*string) *ListKmsInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListKmsInstancesResponse) SetStatusCode(v int32) *ListKmsInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKmsInstancesResponse) SetBody(v *ListKmsInstancesResponseBody) *ListKmsInstancesResponse {
	s.Body = v
	return s
}

type ListNetworkRulesRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListNetworkRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkRulesRequest) GoString() string {
	return s.String()
}

func (s *ListNetworkRulesRequest) SetPageNumber(v int32) *ListNetworkRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNetworkRulesRequest) SetPageSize(v int32) *ListNetworkRulesRequest {
	s.PageSize = &v
	return s
}

type ListNetworkRulesResponseBody struct {
	// A list of access control rules.
	NetworkRules *ListNetworkRulesResponseBodyNetworkRules `json:"NetworkRules,omitempty" xml:"NetworkRules,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3bf02f7a-015b-4f34-be0f-cc043fda2d33
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNetworkRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworkRulesResponseBody) SetNetworkRules(v *ListNetworkRulesResponseBodyNetworkRules) *ListNetworkRulesResponseBody {
	s.NetworkRules = v
	return s
}

func (s *ListNetworkRulesResponseBody) SetPageNumber(v int32) *ListNetworkRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNetworkRulesResponseBody) SetPageSize(v int32) *ListNetworkRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNetworkRulesResponseBody) SetRequestId(v string) *ListNetworkRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNetworkRulesResponseBody) SetTotalCount(v int32) *ListNetworkRulesResponseBody {
	s.TotalCount = &v
	return s
}

type ListNetworkRulesResponseBodyNetworkRules struct {
	NetworkRule []*ListNetworkRulesResponseBodyNetworkRulesNetworkRule `json:"NetworkRule,omitempty" xml:"NetworkRule,omitempty" type:"Repeated"`
}

func (s ListNetworkRulesResponseBodyNetworkRules) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkRulesResponseBodyNetworkRules) GoString() string {
	return s.String()
}

func (s *ListNetworkRulesResponseBodyNetworkRules) SetNetworkRule(v []*ListNetworkRulesResponseBodyNetworkRulesNetworkRule) *ListNetworkRulesResponseBodyNetworkRules {
	s.NetworkRule = v
	return s
}

type ListNetworkRulesResponseBodyNetworkRulesNetworkRule struct {
	// The name of the access control rule.
	//
	// example:
	//
	// networkrule_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network type. The value is fixed as Private. Self-managed applications can access KMS instances only over a private virtual private cloud (VPC).
	//
	// example:
	//
	// Private
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNetworkRulesResponseBodyNetworkRulesNetworkRule) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkRulesResponseBodyNetworkRulesNetworkRule) GoString() string {
	return s.String()
}

func (s *ListNetworkRulesResponseBodyNetworkRulesNetworkRule) SetName(v string) *ListNetworkRulesResponseBodyNetworkRulesNetworkRule {
	s.Name = &v
	return s
}

func (s *ListNetworkRulesResponseBodyNetworkRulesNetworkRule) SetType(v string) *ListNetworkRulesResponseBodyNetworkRulesNetworkRule {
	s.Type = &v
	return s
}

type ListNetworkRulesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetworkRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNetworkRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkRulesResponse) GoString() string {
	return s.String()
}

func (s *ListNetworkRulesResponse) SetHeaders(v map[string]*string) *ListNetworkRulesResponse {
	s.Headers = v
	return s
}

func (s *ListNetworkRulesResponse) SetStatusCode(v int32) *ListNetworkRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNetworkRulesResponse) SetBody(v *ListNetworkRulesResponseBody) *ListNetworkRulesResponse {
	s.Body = v
	return s
}

type ListPoliciesRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesRequest) SetPageNumber(v int32) *ListPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPoliciesRequest) SetPageSize(v int32) *ListPoliciesRequest {
	s.PageSize = &v
	return s
}

type ListPoliciesResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A list of permission policies.
	Policies *ListPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// b66ad557-9c00-4064-9c8d-b621c3263308
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBody) SetPageNumber(v int32) *ListPoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPoliciesResponseBody) SetPageSize(v int32) *ListPoliciesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPoliciesResponseBody) SetPolicies(v *ListPoliciesResponseBodyPolicies) *ListPoliciesResponseBody {
	s.Policies = v
	return s
}

func (s *ListPoliciesResponseBody) SetRequestId(v string) *ListPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPoliciesResponseBody) SetTotalCount(v int32) *ListPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

type ListPoliciesResponseBodyPolicies struct {
	Policy []*ListPoliciesResponseBodyPoliciesPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Repeated"`
}

func (s ListPoliciesResponseBodyPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBodyPolicies) SetPolicy(v []*ListPoliciesResponseBodyPoliciesPolicy) *ListPoliciesResponseBodyPolicies {
	s.Policy = v
	return s
}

type ListPoliciesResponseBodyPoliciesPolicy struct {
	// The name of the permission policy.
	//
	// example:
	//
	// policy_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListPoliciesResponseBodyPoliciesPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponseBodyPoliciesPolicy) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) SetName(v string) *ListPoliciesResponseBodyPoliciesPolicy {
	s.Name = &v
	return s
}

type ListPoliciesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponse) SetHeaders(v map[string]*string) *ListPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListPoliciesResponse) SetStatusCode(v int32) *ListPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPoliciesResponse) SetBody(v *ListPoliciesResponseBody) *ListPoliciesResponse {
	s.Body = v
	return s
}

type ListResourceTagsRequest struct {
	// The globally unique ID of the CMK.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s ListResourceTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTagsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceTagsRequest) SetKeyId(v string) *ListResourceTagsRequest {
	s.KeyId = &v
	return s
}

type ListResourceTagsResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4162a6af-bc99-40b3-a552-89dcc8aaf7c8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags of the CMK.
	Tags *ListResourceTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s ListResourceTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceTagsResponseBody) SetRequestId(v string) *ListResourceTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceTagsResponseBody) SetTags(v *ListResourceTagsResponseBodyTags) *ListResourceTagsResponseBody {
	s.Tags = v
	return s
}

type ListResourceTagsResponseBodyTags struct {
	Tag []*ListResourceTagsResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListResourceTagsResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListResourceTagsResponseBodyTags) SetTag(v []*ListResourceTagsResponseBodyTagsTag) *ListResourceTagsResponseBodyTags {
	s.Tag = v
	return s
}

type ListResourceTagsResponseBodyTagsTag struct {
	// The globally unique ID of the CMK.
	//
	// example:
	//
	// 33caea95-c3e5-4b3e-a9c6-cec76e4e****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The tag key.
	//
	// example:
	//
	// Project
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListResourceTagsResponseBodyTagsTag) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTagsResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *ListResourceTagsResponseBodyTagsTag) SetKeyId(v string) *ListResourceTagsResponseBodyTagsTag {
	s.KeyId = &v
	return s
}

func (s *ListResourceTagsResponseBodyTagsTag) SetTagKey(v string) *ListResourceTagsResponseBodyTagsTag {
	s.TagKey = &v
	return s
}

func (s *ListResourceTagsResponseBodyTagsTag) SetTagValue(v string) *ListResourceTagsResponseBodyTagsTag {
	s.TagValue = &v
	return s
}

type ListResourceTagsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTagsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceTagsResponse) SetHeaders(v map[string]*string) *ListResourceTagsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceTagsResponse) SetStatusCode(v int32) *ListResourceTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceTagsResponse) SetBody(v *ListResourceTagsResponseBody) *ListResourceTagsResponse {
	s.Body = v
	return s
}

type ListSecretVersionIdsRequest struct {
	// Specifies whether to return deprecated secret versions.
	//
	// Valid values:
	//
	// 	- false: no
	//
	// 	- true: yes
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IncludeDeprecated *string `json:"IncludeDeprecated,omitempty" xml:"IncludeDeprecated,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s ListSecretVersionIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSecretVersionIdsRequest) GoString() string {
	return s.String()
}

func (s *ListSecretVersionIdsRequest) SetIncludeDeprecated(v string) *ListSecretVersionIdsRequest {
	s.IncludeDeprecated = &v
	return s
}

func (s *ListSecretVersionIdsRequest) SetPageNumber(v int32) *ListSecretVersionIdsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSecretVersionIdsRequest) SetPageSize(v int32) *ListSecretVersionIdsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSecretVersionIdsRequest) SetSecretName(v string) *ListSecretVersionIdsRequest {
	s.SecretName = &v
	return s
}

type ListSecretVersionIdsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5b75d8b1-5b6a-4ec0-8e0c-c08befdfad47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the secret.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of secret versions.
	VersionIds *ListSecretVersionIdsResponseBodyVersionIds `json:"VersionIds,omitempty" xml:"VersionIds,omitempty" type:"Struct"`
}

func (s ListSecretVersionIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSecretVersionIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecretVersionIdsResponseBody) SetPageNumber(v int32) *ListSecretVersionIdsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSecretVersionIdsResponseBody) SetPageSize(v int32) *ListSecretVersionIdsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSecretVersionIdsResponseBody) SetRequestId(v string) *ListSecretVersionIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecretVersionIdsResponseBody) SetSecretName(v string) *ListSecretVersionIdsResponseBody {
	s.SecretName = &v
	return s
}

func (s *ListSecretVersionIdsResponseBody) SetTotalCount(v int32) *ListSecretVersionIdsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSecretVersionIdsResponseBody) SetVersionIds(v *ListSecretVersionIdsResponseBodyVersionIds) *ListSecretVersionIdsResponseBody {
	s.VersionIds = v
	return s
}

type ListSecretVersionIdsResponseBodyVersionIds struct {
	VersionId []*ListSecretVersionIdsResponseBodyVersionIdsVersionId `json:"VersionId,omitempty" xml:"VersionId,omitempty" type:"Repeated"`
}

func (s ListSecretVersionIdsResponseBodyVersionIds) String() string {
	return tea.Prettify(s)
}

func (s ListSecretVersionIdsResponseBodyVersionIds) GoString() string {
	return s.String()
}

func (s *ListSecretVersionIdsResponseBodyVersionIds) SetVersionId(v []*ListSecretVersionIdsResponseBodyVersionIdsVersionId) *ListSecretVersionIdsResponseBodyVersionIds {
	s.VersionId = v
	return s
}

type ListSecretVersionIdsResponseBodyVersionIdsVersionId struct {
	// The time when the secret version was created.
	//
	// example:
	//
	// 2020-02-21T15:39:26Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The version number.
	//
	// example:
	//
	// 00000000000000000000000000000000203
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The stage labels that mark the secret version.
	VersionStages *ListSecretVersionIdsResponseBodyVersionIdsVersionIdVersionStages `json:"VersionStages,omitempty" xml:"VersionStages,omitempty" type:"Struct"`
}

func (s ListSecretVersionIdsResponseBodyVersionIdsVersionId) String() string {
	return tea.Prettify(s)
}

func (s ListSecretVersionIdsResponseBodyVersionIdsVersionId) GoString() string {
	return s.String()
}

func (s *ListSecretVersionIdsResponseBodyVersionIdsVersionId) SetCreateTime(v string) *ListSecretVersionIdsResponseBodyVersionIdsVersionId {
	s.CreateTime = &v
	return s
}

func (s *ListSecretVersionIdsResponseBodyVersionIdsVersionId) SetVersionId(v string) *ListSecretVersionIdsResponseBodyVersionIdsVersionId {
	s.VersionId = &v
	return s
}

func (s *ListSecretVersionIdsResponseBodyVersionIdsVersionId) SetVersionStages(v *ListSecretVersionIdsResponseBodyVersionIdsVersionIdVersionStages) *ListSecretVersionIdsResponseBodyVersionIdsVersionId {
	s.VersionStages = v
	return s
}

type ListSecretVersionIdsResponseBodyVersionIdsVersionIdVersionStages struct {
	VersionStage []*string `json:"VersionStage,omitempty" xml:"VersionStage,omitempty" type:"Repeated"`
}

func (s ListSecretVersionIdsResponseBodyVersionIdsVersionIdVersionStages) String() string {
	return tea.Prettify(s)
}

func (s ListSecretVersionIdsResponseBodyVersionIdsVersionIdVersionStages) GoString() string {
	return s.String()
}

func (s *ListSecretVersionIdsResponseBodyVersionIdsVersionIdVersionStages) SetVersionStage(v []*string) *ListSecretVersionIdsResponseBodyVersionIdsVersionIdVersionStages {
	s.VersionStage = v
	return s
}

type ListSecretVersionIdsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSecretVersionIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSecretVersionIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSecretVersionIdsResponse) GoString() string {
	return s.String()
}

func (s *ListSecretVersionIdsResponse) SetHeaders(v map[string]*string) *ListSecretVersionIdsResponse {
	s.Headers = v
	return s
}

func (s *ListSecretVersionIdsResponse) SetStatusCode(v int32) *ListSecretVersionIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecretVersionIdsResponse) SetBody(v *ListSecretVersionIdsResponseBody) *ListSecretVersionIdsResponse {
	s.Body = v
	return s
}

type ListSecretsRequest struct {
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// false
	FetchTags *string `json:"FetchTags,omitempty" xml:"FetchTags,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// [{"Key":"SecretName", "Values":["Val1","Val2"]}]
	Filters *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	// The secret filter. The filter consists of one or more key-value pairs. You can specify one key-value pair or leave this parameter empty. If you use one tag key or tag value to filter resources, up to 4,000 resources can be queried. If you want to query more than 4,000 resources, call the [ListResourceTags](https://help.aliyun.com/document_detail/120090.html) operation.
	//
	// 	- Key
	//
	//     	- Description: the property that you want to filter.
	//
	//     	- Type: string.
	//
	//     	- Valid values:
	//
	//         	- SecretName: the secret name.
	//
	//         	- Description: the description of the secret.
	//
	//         	- TagKey: the tag key.
	//
	//         	- TagValue: the tag value.
	//
	// 	- Values
	//
	//     	- Description: the value to be included after filtering.
	//
	//     	- Type: string.
	//
	//     	- Length: 0 to 10.
	//
	//     	- Valid values:
	//
	//         	- If the Key field is set to SecretName, the value must be 1 to 192 characters in length and can contain letters, digits, and special characters `_ / + = . @ -`.
	//
	//         	- If the Key field is set to Description, the value must be 1 to 256 characters in length.
	//
	//         	- If the Key field is set to TagKey, the value must be 1 to 256 characters in length and can contain letters, digits, and special characters `/ _ - . + = @ :`.
	//
	//         	- If the Key field is set to TagValue, the value must be 1 to 256 characters in length and can contain letters, numbers, and special characters `/ _ - . + = @ :`.
	//
	// The logical relationship between values of the Values field in a key-value pair is OR. Example: `[ {"Key":"SecretName", "Values":["sec1","sec2"]}]`. In this example, the semantics are `SecretName=sec 1 OR SecretName=sec 2`.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSecretsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSecretsRequest) GoString() string {
	return s.String()
}

func (s *ListSecretsRequest) SetFetchTags(v string) *ListSecretsRequest {
	s.FetchTags = &v
	return s
}

func (s *ListSecretsRequest) SetFilters(v string) *ListSecretsRequest {
	s.Filters = &v
	return s
}

func (s *ListSecretsRequest) SetPageNumber(v int32) *ListSecretsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSecretsRequest) SetPageSize(v int32) *ListSecretsRequest {
	s.PageSize = &v
	return s
}

type ListSecretsResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of returned secrets.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of secrets.
	//
	// example:
	//
	// 6a6287a0-ff34-4780-a790-fdfca900557f
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the secret was updated.
	SecretList *ListSecretsResponseBodySecretList `json:"SecretList,omitempty" xml:"SecretList,omitempty" type:"Struct"`
	// The secret name.
	//
	// example:
	//
	// 55
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSecretsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSecretsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBody) SetPageNumber(v int32) *ListSecretsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSecretsResponseBody) SetPageSize(v int32) *ListSecretsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSecretsResponseBody) SetRequestId(v string) *ListSecretsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecretsResponseBody) SetSecretList(v *ListSecretsResponseBodySecretList) *ListSecretsResponseBody {
	s.SecretList = v
	return s
}

func (s *ListSecretsResponseBody) SetTotalCount(v int32) *ListSecretsResponseBody {
	s.TotalCount = &v
	return s
}

type ListSecretsResponseBodySecretList struct {
	Secret []*ListSecretsResponseBodySecretListSecret `json:"Secret,omitempty" xml:"Secret,omitempty" type:"Repeated"`
}

func (s ListSecretsResponseBodySecretList) String() string {
	return tea.Prettify(s)
}

func (s ListSecretsResponseBodySecretList) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBodySecretList) SetSecret(v []*ListSecretsResponseBodySecretListSecret) *ListSecretsResponseBodySecretList {
	s.Secret = v
	return s
}

type ListSecretsResponseBodySecretListSecret struct {
	// The tag value.
	//
	// example:
	//
	// 2022-07-17T07:59:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The resource tags of the secret.
	//
	// This parameter is not returned if you set the FetchTags parameter to false or do not specify the FetchTags parameter.
	//
	// example:
	//
	// 2022-08-17T07:59:05Z
	PlannedDeleteTime *string `json:"PlannedDeleteTime,omitempty" xml:"PlannedDeleteTime,omitempty"`
	// The type of the secret. Valid values:
	//
	// 	- Generic: indicates a generic secret.
	//
	// 	- Rds: indicates a managed ApsaraDB RDS secret.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The time when the secret was created.
	//
	// example:
	//
	// Generic
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// The tag key.
	Tags *ListSecretsResponseBodySecretListSecretTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The time when the secret is scheduled to be deleted.
	//
	// example:
	//
	// 2022-07-17T07:59:05Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListSecretsResponseBodySecretListSecret) String() string {
	return tea.Prettify(s)
}

func (s ListSecretsResponseBodySecretListSecret) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBodySecretListSecret) SetCreateTime(v string) *ListSecretsResponseBodySecretListSecret {
	s.CreateTime = &v
	return s
}

func (s *ListSecretsResponseBodySecretListSecret) SetPlannedDeleteTime(v string) *ListSecretsResponseBodySecretListSecret {
	s.PlannedDeleteTime = &v
	return s
}

func (s *ListSecretsResponseBodySecretListSecret) SetSecretName(v string) *ListSecretsResponseBodySecretListSecret {
	s.SecretName = &v
	return s
}

func (s *ListSecretsResponseBodySecretListSecret) SetSecretType(v string) *ListSecretsResponseBodySecretListSecret {
	s.SecretType = &v
	return s
}

func (s *ListSecretsResponseBodySecretListSecret) SetTags(v *ListSecretsResponseBodySecretListSecretTags) *ListSecretsResponseBodySecretListSecret {
	s.Tags = v
	return s
}

func (s *ListSecretsResponseBodySecretListSecret) SetUpdateTime(v string) *ListSecretsResponseBodySecretListSecret {
	s.UpdateTime = &v
	return s
}

type ListSecretsResponseBodySecretListSecretTags struct {
	Tag []*ListSecretsResponseBodySecretListSecretTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListSecretsResponseBodySecretListSecretTags) String() string {
	return tea.Prettify(s)
}

func (s ListSecretsResponseBodySecretListSecretTags) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBodySecretListSecretTags) SetTag(v []*ListSecretsResponseBodySecretListSecretTagsTag) *ListSecretsResponseBodySecretListSecretTags {
	s.Tag = v
	return s
}

type ListSecretsResponseBodySecretListSecretTagsTag struct {
	// example:
	//
	// key1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// val1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListSecretsResponseBodySecretListSecretTagsTag) String() string {
	return tea.Prettify(s)
}

func (s ListSecretsResponseBodySecretListSecretTagsTag) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBodySecretListSecretTagsTag) SetTagKey(v string) *ListSecretsResponseBodySecretListSecretTagsTag {
	s.TagKey = &v
	return s
}

func (s *ListSecretsResponseBodySecretListSecretTagsTag) SetTagValue(v string) *ListSecretsResponseBodySecretListSecretTagsTag {
	s.TagValue = &v
	return s
}

type ListSecretsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSecretsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSecretsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSecretsResponse) GoString() string {
	return s.String()
}

func (s *ListSecretsResponse) SetHeaders(v map[string]*string) *ListSecretsResponse {
	s.Headers = v
	return s
}

func (s *ListSecretsResponse) SetStatusCode(v int32) *ListSecretsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecretsResponse) SetBody(v *ListSecretsResponseBody) *ListSecretsResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// >  If the call does not return all result entries, the value of the NextToken parameter is returned. By default, 200 rows are returned. You can call this operation again and set the value of the parameter to the value of the parameter that is returned in the last call to implement paged query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the resource.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/601478.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// A list of resource IDs for which you want to query tags. You can enter a maximum of 50 resource IDs.
	//
	// Enter multiple resource IDs in the `["ResourceId. 1","ResourceId. 2",...]` format.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of resource whose tags you want to query. Valid value:
	//
	// 	- key
	//
	// 	- secret
	//
	// This parameter is required.
	//
	// example:
	//
	// key
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// A list of tags that you want to query. Valid values of N: 1 to 20.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The key of the tag. A tag consists of a key-value pair.
	//
	// You can enter up to 20 tags. Enter multiple tags in the `[{"Key":"key1","Value":"value1"},{"Key":"key2","Value":"value2"},..]` format.
	//
	// >  The key cannot start with aliyun or acs:.
	//
	// example:
	//
	// disk-encryption
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. A tag consists of a key-value pair.
	//
	// You can enter up to 20 tags. Enter multiple tags in the `[{"Key":"key1","Value":"value1"},{"Key":"key2","Value":"value2"},..]` format.
	//
	// example:
	//
	// true
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// 	- If NextToken is empty ("NextToken": ""), no next page exists.
	//
	// 	- If NextToken is not empty, the next query is required, and the value is the token used to start the next query.
	//
	// example:
	//
	// e71d8a535bd9cc11
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 00827261-20B7-4562-83F2-4DF39876A45A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of tags.
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	// The resource ID.
	//
	// example:
	//
	// key-hzz62f1cb66fa42qo****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// key
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag.
	//
	// example:
	//
	// disk-encryption
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// true
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type OpenKmsServiceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 3455b9b4-95c1-419d-b310-db6a53b09a39
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenKmsServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenKmsServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenKmsServiceResponseBody) SetRequestId(v string) *OpenKmsServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenKmsServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenKmsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenKmsServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenKmsServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenKmsServiceResponse) SetHeaders(v map[string]*string) *OpenKmsServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenKmsServiceResponse) SetStatusCode(v int32) *OpenKmsServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenKmsServiceResponse) SetBody(v *OpenKmsServiceResponseBody) *OpenKmsServiceResponse {
	s.Body = v
	return s
}

type PutSecretValueRequest struct {
	// The secret value. The value is encrypted and then stored in the new version.
	//
	// This parameter is required.
	//
	// example:
	//
	// importantdata
	SecretData *string `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
	// The type of the secret value. Valid values:
	//
	// 	- text: This is the default value.
	//
	// 	- binary
	//
	// example:
	//
	// text
	SecretDataType *string `json:"SecretDataType,omitempty" xml:"SecretDataType,omitempty"`
	// The name of the secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The new version of the secret value. Version numbers must be unique in each secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00000000000000000000000000000000203
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The stage labels that are used to mark the new version. If you do not specify this parameter, Secrets Manager marks the new version with ACSCurrent.
	//
	// example:
	//
	// ["ACSCurrent","ACSNext"]
	VersionStages *string `json:"VersionStages,omitempty" xml:"VersionStages,omitempty"`
}

func (s PutSecretValueRequest) String() string {
	return tea.Prettify(s)
}

func (s PutSecretValueRequest) GoString() string {
	return s.String()
}

func (s *PutSecretValueRequest) SetSecretData(v string) *PutSecretValueRequest {
	s.SecretData = &v
	return s
}

func (s *PutSecretValueRequest) SetSecretDataType(v string) *PutSecretValueRequest {
	s.SecretDataType = &v
	return s
}

func (s *PutSecretValueRequest) SetSecretName(v string) *PutSecretValueRequest {
	s.SecretName = &v
	return s
}

func (s *PutSecretValueRequest) SetVersionId(v string) *PutSecretValueRequest {
	s.VersionId = &v
	return s
}

func (s *PutSecretValueRequest) SetVersionStages(v string) *PutSecretValueRequest {
	s.VersionStages = &v
	return s
}

type PutSecretValueResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// f94ec9d3-2d10-4922-9a5c-5dcd5ebcb5e8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the secret.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The new version of the secret value.
	//
	// example:
	//
	// 00000000000000000000000000000000203
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The stage labels that are used to mark the new version.
	VersionStages *PutSecretValueResponseBodyVersionStages `json:"VersionStages,omitempty" xml:"VersionStages,omitempty" type:"Struct"`
}

func (s PutSecretValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutSecretValueResponseBody) GoString() string {
	return s.String()
}

func (s *PutSecretValueResponseBody) SetRequestId(v string) *PutSecretValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutSecretValueResponseBody) SetSecretName(v string) *PutSecretValueResponseBody {
	s.SecretName = &v
	return s
}

func (s *PutSecretValueResponseBody) SetVersionId(v string) *PutSecretValueResponseBody {
	s.VersionId = &v
	return s
}

func (s *PutSecretValueResponseBody) SetVersionStages(v *PutSecretValueResponseBodyVersionStages) *PutSecretValueResponseBody {
	s.VersionStages = v
	return s
}

type PutSecretValueResponseBodyVersionStages struct {
	VersionStage []*string `json:"VersionStage,omitempty" xml:"VersionStage,omitempty" type:"Repeated"`
}

func (s PutSecretValueResponseBodyVersionStages) String() string {
	return tea.Prettify(s)
}

func (s PutSecretValueResponseBodyVersionStages) GoString() string {
	return s.String()
}

func (s *PutSecretValueResponseBodyVersionStages) SetVersionStage(v []*string) *PutSecretValueResponseBodyVersionStages {
	s.VersionStage = v
	return s
}

type PutSecretValueResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutSecretValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutSecretValueResponse) String() string {
	return tea.Prettify(s)
}

func (s PutSecretValueResponse) GoString() string {
	return s.String()
}

func (s *PutSecretValueResponse) SetHeaders(v map[string]*string) *PutSecretValueResponse {
	s.Headers = v
	return s
}

func (s *PutSecretValueResponse) SetStatusCode(v int32) *PutSecretValueResponse {
	s.StatusCode = &v
	return s
}

func (s *PutSecretValueResponse) SetBody(v *PutSecretValueResponseBody) *PutSecretValueResponse {
	s.Body = v
	return s
}

type ReEncryptRequest struct {
	// The ciphertext that you want to re-encrypt.
	//
	// You can set this parameter to the ciphertext that is returned after a symmetric or asymmetric encryption operation.
	//
	// 	- Symmetric encryption: the ciphertext returned after you call the [Encrypt](https://help.aliyun.com/document_detail/28949.html), [GenerateDataKey](https://help.aliyun.com/document_detail/28948.html), [GenerateDataKeyWithoutPlaintext](https://help.aliyun.com/document_detail/134043.html), or [GenerateAndExportDataKey](https://help.aliyun.com/document_detail/176804.html) operation
	//
	// 	- Asymmetric encryption: the public key-encrypted ciphertext returned after you call the [GenerateAndExportDataKey](https://help.aliyun.com/document_detail/176804.html) operation, or the ciphertext encrypted by using the public key of an asymmetric key pair outside KMS
	//
	// This parameter is required.
	//
	// example:
	//
	// ODZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmS7FmDBBQ0BkKsQrtRnidtPwirmDcS0ZuJCU41xxAAWk4Z8qsADfbV0b+i6kQmlvj79dJdGOvtX69Uycs901q********
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// A JSON string that consists of key-value pairs. This parameter specifies the EncryptionContext that is used to re-encrypt the decrypted data or data key.
	//
	// example:
	//
	// {"Example":"Example"}
	DestinationEncryptionContext map[string]interface{} `json:"DestinationEncryptionContext,omitempty" xml:"DestinationEncryptionContext,omitempty"`
	// The ID of the symmetric CMK that is used to re-encrypt the ciphertext after the ciphertext is decrypted.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	DestinationKeyId *string `json:"DestinationKeyId,omitempty" xml:"DestinationKeyId,omitempty"`
	// The encryption algorithm based on which the public key is used to encrypt the ciphertext specified by CiphertextBlob. For more information about encryption algorithms, see [AsymmetricDecrypt](https://help.aliyun.com/document_detail/148130.html).
	//
	// Valid values:
	//
	// 	- RSAES_OAEP_SHA_256
	//
	// 	- RSAES_OAEP_SHA_1
	//
	// 	- SM2PKE
	//
	// >  If you set CiphertextBlob to the public key-encrypted ciphertext that is returned after an asymmetric encryption operation, specify this parameter.
	//
	// example:
	//
	// RSAES_OAEP_SHA_256
	SourceEncryptionAlgorithm *string `json:"SourceEncryptionAlgorithm,omitempty" xml:"SourceEncryptionAlgorithm,omitempty"`
	// A JSON string that consists of key-value pairs. If you specify EncryptionContext when you call the [Encrypt](https://help.aliyun.com/document_detail/28949.html), [GenerateDataKey](https://help.aliyun.com/document_detail/28948.html), [GenerateDataKeyWithoutPlaintext](https://help.aliyun.com/document_detail/134043.html), or [GenerateAndExportDataKey](https://help.aliyun.com/document_detail/176804.html) operation to encrypt the data or data key, an equivalent value is required here. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// >  If you set CiphertextBlob to the ciphertext that is returned after a symmetric encryption operation, specify this parameter.
	//
	// example:
	//
	// {"Example":"Example"}
	SourceEncryptionContext map[string]interface{} `json:"SourceEncryptionContext,omitempty" xml:"SourceEncryptionContext,omitempty"`
	// The ID of the CMK that is used to decrypt the ciphertext.
	//
	// This parameter is the globally unique ID of the CMK.
	//
	// >  If you set CiphertextBlob to the public key-encrypted ciphertext that is returned after an asymmetric encryption operation, specify this parameter.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	SourceKeyId *string `json:"SourceKeyId,omitempty" xml:"SourceKeyId,omitempty"`
	// The ID of the CMK version that is used to decrypt the ciphertext.
	//
	// >  If you set CiphertextBlob to the public key-encrypted ciphertext that is returned after an asymmetric encryption operation, specify this parameter.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	SourceKeyVersionId *string `json:"SourceKeyVersionId,omitempty" xml:"SourceKeyVersionId,omitempty"`
}

func (s ReEncryptRequest) String() string {
	return tea.Prettify(s)
}

func (s ReEncryptRequest) GoString() string {
	return s.String()
}

func (s *ReEncryptRequest) SetCiphertextBlob(v string) *ReEncryptRequest {
	s.CiphertextBlob = &v
	return s
}

func (s *ReEncryptRequest) SetDestinationEncryptionContext(v map[string]interface{}) *ReEncryptRequest {
	s.DestinationEncryptionContext = v
	return s
}

func (s *ReEncryptRequest) SetDestinationKeyId(v string) *ReEncryptRequest {
	s.DestinationKeyId = &v
	return s
}

func (s *ReEncryptRequest) SetSourceEncryptionAlgorithm(v string) *ReEncryptRequest {
	s.SourceEncryptionAlgorithm = &v
	return s
}

func (s *ReEncryptRequest) SetSourceEncryptionContext(v map[string]interface{}) *ReEncryptRequest {
	s.SourceEncryptionContext = v
	return s
}

func (s *ReEncryptRequest) SetSourceKeyId(v string) *ReEncryptRequest {
	s.SourceKeyId = &v
	return s
}

func (s *ReEncryptRequest) SetSourceKeyVersionId(v string) *ReEncryptRequest {
	s.SourceKeyVersionId = &v
	return s
}

type ReEncryptShrinkRequest struct {
	// The ciphertext that you want to re-encrypt.
	//
	// You can set this parameter to the ciphertext that is returned after a symmetric or asymmetric encryption operation.
	//
	// 	- Symmetric encryption: the ciphertext returned after you call the [Encrypt](https://help.aliyun.com/document_detail/28949.html), [GenerateDataKey](https://help.aliyun.com/document_detail/28948.html), [GenerateDataKeyWithoutPlaintext](https://help.aliyun.com/document_detail/134043.html), or [GenerateAndExportDataKey](https://help.aliyun.com/document_detail/176804.html) operation
	//
	// 	- Asymmetric encryption: the public key-encrypted ciphertext returned after you call the [GenerateAndExportDataKey](https://help.aliyun.com/document_detail/176804.html) operation, or the ciphertext encrypted by using the public key of an asymmetric key pair outside KMS
	//
	// This parameter is required.
	//
	// example:
	//
	// ODZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmS7FmDBBQ0BkKsQrtRnidtPwirmDcS0ZuJCU41xxAAWk4Z8qsADfbV0b+i6kQmlvj79dJdGOvtX69Uycs901q********
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// A JSON string that consists of key-value pairs. This parameter specifies the EncryptionContext that is used to re-encrypt the decrypted data or data key.
	//
	// example:
	//
	// {"Example":"Example"}
	DestinationEncryptionContextShrink *string `json:"DestinationEncryptionContext,omitempty" xml:"DestinationEncryptionContext,omitempty"`
	// The ID of the symmetric CMK that is used to re-encrypt the ciphertext after the ciphertext is decrypted.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	DestinationKeyId *string `json:"DestinationKeyId,omitempty" xml:"DestinationKeyId,omitempty"`
	// The encryption algorithm based on which the public key is used to encrypt the ciphertext specified by CiphertextBlob. For more information about encryption algorithms, see [AsymmetricDecrypt](https://help.aliyun.com/document_detail/148130.html).
	//
	// Valid values:
	//
	// 	- RSAES_OAEP_SHA_256
	//
	// 	- RSAES_OAEP_SHA_1
	//
	// 	- SM2PKE
	//
	// >  If you set CiphertextBlob to the public key-encrypted ciphertext that is returned after an asymmetric encryption operation, specify this parameter.
	//
	// example:
	//
	// RSAES_OAEP_SHA_256
	SourceEncryptionAlgorithm *string `json:"SourceEncryptionAlgorithm,omitempty" xml:"SourceEncryptionAlgorithm,omitempty"`
	// A JSON string that consists of key-value pairs. If you specify EncryptionContext when you call the [Encrypt](https://help.aliyun.com/document_detail/28949.html), [GenerateDataKey](https://help.aliyun.com/document_detail/28948.html), [GenerateDataKeyWithoutPlaintext](https://help.aliyun.com/document_detail/134043.html), or [GenerateAndExportDataKey](https://help.aliyun.com/document_detail/176804.html) operation to encrypt the data or data key, an equivalent value is required here. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// >  If you set CiphertextBlob to the ciphertext that is returned after a symmetric encryption operation, specify this parameter.
	//
	// example:
	//
	// {"Example":"Example"}
	SourceEncryptionContextShrink *string `json:"SourceEncryptionContext,omitempty" xml:"SourceEncryptionContext,omitempty"`
	// The ID of the CMK that is used to decrypt the ciphertext.
	//
	// This parameter is the globally unique ID of the CMK.
	//
	// >  If you set CiphertextBlob to the public key-encrypted ciphertext that is returned after an asymmetric encryption operation, specify this parameter.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	SourceKeyId *string `json:"SourceKeyId,omitempty" xml:"SourceKeyId,omitempty"`
	// The ID of the CMK version that is used to decrypt the ciphertext.
	//
	// >  If you set CiphertextBlob to the public key-encrypted ciphertext that is returned after an asymmetric encryption operation, specify this parameter.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	SourceKeyVersionId *string `json:"SourceKeyVersionId,omitempty" xml:"SourceKeyVersionId,omitempty"`
}

func (s ReEncryptShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ReEncryptShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReEncryptShrinkRequest) SetCiphertextBlob(v string) *ReEncryptShrinkRequest {
	s.CiphertextBlob = &v
	return s
}

func (s *ReEncryptShrinkRequest) SetDestinationEncryptionContextShrink(v string) *ReEncryptShrinkRequest {
	s.DestinationEncryptionContextShrink = &v
	return s
}

func (s *ReEncryptShrinkRequest) SetDestinationKeyId(v string) *ReEncryptShrinkRequest {
	s.DestinationKeyId = &v
	return s
}

func (s *ReEncryptShrinkRequest) SetSourceEncryptionAlgorithm(v string) *ReEncryptShrinkRequest {
	s.SourceEncryptionAlgorithm = &v
	return s
}

func (s *ReEncryptShrinkRequest) SetSourceEncryptionContextShrink(v string) *ReEncryptShrinkRequest {
	s.SourceEncryptionContextShrink = &v
	return s
}

func (s *ReEncryptShrinkRequest) SetSourceKeyId(v string) *ReEncryptShrinkRequest {
	s.SourceKeyId = &v
	return s
}

func (s *ReEncryptShrinkRequest) SetSourceKeyVersionId(v string) *ReEncryptShrinkRequest {
	s.SourceKeyVersionId = &v
	return s
}

type ReEncryptResponseBody struct {
	// The ciphertext re-encrypted.
	//
	// example:
	//
	// DZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmaaSl+TztSIMe43nbTH/Z1Wr4XfLftKhAciUmDQXuMRl4WTvKhxjMThjK****
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The ID of the CMK that is used to decrypt the original ciphertext.
	//
	// This parameter is the globally unique ID of the CMK.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the CMK version that is used to decrypt the original ciphertext.
	//
	// example:
	//
	// 202b9877-5a25-46e3-a763-e20791b5****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 207596a2-36d3-4840-b1bd-f87044699bd7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReEncryptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReEncryptResponseBody) GoString() string {
	return s.String()
}

func (s *ReEncryptResponseBody) SetCiphertextBlob(v string) *ReEncryptResponseBody {
	s.CiphertextBlob = &v
	return s
}

func (s *ReEncryptResponseBody) SetKeyId(v string) *ReEncryptResponseBody {
	s.KeyId = &v
	return s
}

func (s *ReEncryptResponseBody) SetKeyVersionId(v string) *ReEncryptResponseBody {
	s.KeyVersionId = &v
	return s
}

func (s *ReEncryptResponseBody) SetRequestId(v string) *ReEncryptResponseBody {
	s.RequestId = &v
	return s
}

type ReEncryptResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReEncryptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReEncryptResponse) String() string {
	return tea.Prettify(s)
}

func (s ReEncryptResponse) GoString() string {
	return s.String()
}

func (s *ReEncryptResponse) SetHeaders(v map[string]*string) *ReEncryptResponse {
	s.Headers = v
	return s
}

func (s *ReEncryptResponse) SetStatusCode(v int32) *ReEncryptResponse {
	s.StatusCode = &v
	return s
}

func (s *ReEncryptResponse) SetBody(v *ReEncryptResponseBody) *ReEncryptResponse {
	s.Body = v
	return s
}

type RestoreSecretRequest struct {
	// The name of the secret you want to restore.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s RestoreSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s RestoreSecretRequest) GoString() string {
	return s.String()
}

func (s *RestoreSecretRequest) SetSecretName(v string) *RestoreSecretRequest {
	s.SecretName = &v
	return s
}

type RestoreSecretResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// e4885adf-548f-4ca5-8075-f540bbd3a55f
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the secret.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s RestoreSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestoreSecretResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreSecretResponseBody) SetRequestId(v string) *RestoreSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestoreSecretResponseBody) SetSecretName(v string) *RestoreSecretResponseBody {
	s.SecretName = &v
	return s
}

type RestoreSecretResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s RestoreSecretResponse) GoString() string {
	return s.String()
}

func (s *RestoreSecretResponse) SetHeaders(v map[string]*string) *RestoreSecretResponse {
	s.Headers = v
	return s
}

func (s *RestoreSecretResponse) SetStatusCode(v int32) *RestoreSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreSecretResponse) SetBody(v *RestoreSecretResponseBody) *RestoreSecretResponse {
	s.Body = v
	return s
}

type RotateSecretRequest struct {
	// The name of the secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// RdsSecret/Mysql5.4/MyCred
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The version number of the secret after the secret is rotated.
	//
	// >  The version number is used to ensure the idempotence of the request. Secrets Manager uses this version number to prevent your application from creating the same version of the secret when the application retries a request. If a version number already exists, Secrets Manager ignores the request for rotation and returns a success message.
	//
	// This parameter is required.
	//
	// example:
	//
	// 000000123
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s RotateSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s RotateSecretRequest) GoString() string {
	return s.String()
}

func (s *RotateSecretRequest) SetSecretName(v string) *RotateSecretRequest {
	s.SecretName = &v
	return s
}

func (s *RotateSecretRequest) SetVersionId(v string) *RotateSecretRequest {
	s.VersionId = &v
	return s
}

type RotateSecretResponseBody struct {
	// The Alibaba Cloud Resource Name (ARN) of the secret.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:154035569884****:secret/RdsSecret/Mysql5.4/MyCred
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 10257c86-269d-43aa-aaf3-90ed4144bb7c
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the secret.
	//
	// example:
	//
	// RdsSecret/Mysql5.4/MyCred
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The version number of the secret after the secret is rotated.
	//
	// example:
	//
	// 000000123
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s RotateSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RotateSecretResponseBody) GoString() string {
	return s.String()
}

func (s *RotateSecretResponseBody) SetArn(v string) *RotateSecretResponseBody {
	s.Arn = &v
	return s
}

func (s *RotateSecretResponseBody) SetRequestId(v string) *RotateSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *RotateSecretResponseBody) SetSecretName(v string) *RotateSecretResponseBody {
	s.SecretName = &v
	return s
}

func (s *RotateSecretResponseBody) SetVersionId(v string) *RotateSecretResponseBody {
	s.VersionId = &v
	return s
}

type RotateSecretResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RotateSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RotateSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s RotateSecretResponse) GoString() string {
	return s.String()
}

func (s *RotateSecretResponse) SetHeaders(v map[string]*string) *RotateSecretResponse {
	s.Headers = v
	return s
}

func (s *RotateSecretResponse) SetStatusCode(v int32) *RotateSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *RotateSecretResponse) SetBody(v *RotateSecretResponseBody) *RotateSecretResponse {
	s.Body = v
	return s
}

type ScheduleKeyDeletionRequest struct {
	// The ID of the customer master key (CMK). The ID must be globally unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7906979c-8e06-46a2-be2d-68e3ccbc****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The scheduled period after which the CMK is deleted. During this period, the CMK is in the PendingDeletion state. After this period ends, you cannot cancel the key deletion task.
	//
	// Valid values: 7 to 366.
	//
	// Unit: days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	PendingWindowInDays *int32 `json:"PendingWindowInDays,omitempty" xml:"PendingWindowInDays,omitempty"`
}

func (s ScheduleKeyDeletionRequest) String() string {
	return tea.Prettify(s)
}

func (s ScheduleKeyDeletionRequest) GoString() string {
	return s.String()
}

func (s *ScheduleKeyDeletionRequest) SetKeyId(v string) *ScheduleKeyDeletionRequest {
	s.KeyId = &v
	return s
}

func (s *ScheduleKeyDeletionRequest) SetPendingWindowInDays(v int32) *ScheduleKeyDeletionRequest {
	s.PendingWindowInDays = &v
	return s
}

type ScheduleKeyDeletionResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3da5b8cc-8107-40ac-a170-793cd181d7b7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ScheduleKeyDeletionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScheduleKeyDeletionResponseBody) GoString() string {
	return s.String()
}

func (s *ScheduleKeyDeletionResponseBody) SetRequestId(v string) *ScheduleKeyDeletionResponseBody {
	s.RequestId = &v
	return s
}

type ScheduleKeyDeletionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScheduleKeyDeletionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScheduleKeyDeletionResponse) String() string {
	return tea.Prettify(s)
}

func (s ScheduleKeyDeletionResponse) GoString() string {
	return s.String()
}

func (s *ScheduleKeyDeletionResponse) SetHeaders(v map[string]*string) *ScheduleKeyDeletionResponse {
	s.Headers = v
	return s
}

func (s *ScheduleKeyDeletionResponse) SetStatusCode(v int32) *ScheduleKeyDeletionResponse {
	s.StatusCode = &v
	return s
}

func (s *ScheduleKeyDeletionResponse) SetBody(v *ScheduleKeyDeletionResponseBody) *ScheduleKeyDeletionResponse {
	s.Body = v
	return s
}

type SetDeletionProtectionRequest struct {
	// The description of deletion protection.
	//
	// >  This parameter takes effect only when you set the EnableDeletionProtection parameter to true.
	//
	// example:
	//
	// This key is being used by XXX service. You are protected from deletion.
	DeletionProtectionDescription *string `json:"DeletionProtectionDescription,omitempty" xml:"DeletionProtectionDescription,omitempty"`
	// Specifies whether to enable deletion protection. Valid values:
	//
	// 	- true: enables deletion protection.
	//
	// 	- false: disables deletion protection.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableDeletionProtection *bool `json:"EnableDeletionProtection,omitempty" xml:"EnableDeletionProtection,omitempty"`
	// The ARN of the CMK for which you want to set deletion protection.
	//
	// You can call the [DescribeKey](https://help.aliyun.com/document_detail/28952.html) operation to query the CMK ARN.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:123213123****:key/0225f411-b21d-46d1-be5b-93931c82****
	ProtectedResourceArn *string `json:"ProtectedResourceArn,omitempty" xml:"ProtectedResourceArn,omitempty"`
}

func (s SetDeletionProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDeletionProtectionRequest) GoString() string {
	return s.String()
}

func (s *SetDeletionProtectionRequest) SetDeletionProtectionDescription(v string) *SetDeletionProtectionRequest {
	s.DeletionProtectionDescription = &v
	return s
}

func (s *SetDeletionProtectionRequest) SetEnableDeletionProtection(v bool) *SetDeletionProtectionRequest {
	s.EnableDeletionProtection = &v
	return s
}

func (s *SetDeletionProtectionRequest) SetProtectedResourceArn(v string) *SetDeletionProtectionRequest {
	s.ProtectedResourceArn = &v
	return s
}

type SetDeletionProtectionResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3455b9b4-95c1-419d-b310-db6a53b09a39
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDeletionProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDeletionProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *SetDeletionProtectionResponseBody) SetRequestId(v string) *SetDeletionProtectionResponseBody {
	s.RequestId = &v
	return s
}

type SetDeletionProtectionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDeletionProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDeletionProtectionResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDeletionProtectionResponse) GoString() string {
	return s.String()
}

func (s *SetDeletionProtectionResponse) SetHeaders(v map[string]*string) *SetDeletionProtectionResponse {
	s.Headers = v
	return s
}

func (s *SetDeletionProtectionResponse) SetStatusCode(v int32) *SetDeletionProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDeletionProtectionResponse) SetBody(v *SetDeletionProtectionResponseBody) *SetDeletionProtectionResponse {
	s.Body = v
	return s
}

type SetKeyPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// key-hzz630494463ejqjx****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"Statement":[{"Action":["kms:*"],"Effect":"Allow","Principal":{"RAM":["acs:ram::119285303511****:*"]},"Resource":["*"],"Sid":"kms default key policy"},{"Action":["kms:List*","kms:Describe*","kms:Create*","kms:Enable*","kms:Disable*","kms:Get*","kms:Set*","kms:Update*","kms:Delete*","kms:Cancel*","kms:TagResource","kms:UntagResource","kms:ImportKeyMaterial","kms:ScheduleKeyDeletion"],"Effect":"Allow","Principal":{"RAM":["acs:ram::119285303511****:user/for_test_policy"]},"Resource":["*"]}],"Version":"1"}
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// default
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s SetKeyPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s SetKeyPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetKeyPolicyRequest) SetKeyId(v string) *SetKeyPolicyRequest {
	s.KeyId = &v
	return s
}

func (s *SetKeyPolicyRequest) SetPolicy(v string) *SetKeyPolicyRequest {
	s.Policy = &v
	return s
}

func (s *SetKeyPolicyRequest) SetPolicyName(v string) *SetKeyPolicyRequest {
	s.PolicyName = &v
	return s
}

type SetKeyPolicyResponseBody struct {
	// example:
	//
	// 381D5D33-BB8F-395F-8EE4-AE3BB4B523C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetKeyPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetKeyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SetKeyPolicyResponseBody) SetRequestId(v string) *SetKeyPolicyResponseBody {
	s.RequestId = &v
	return s
}

type SetKeyPolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetKeyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetKeyPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s SetKeyPolicyResponse) GoString() string {
	return s.String()
}

func (s *SetKeyPolicyResponse) SetHeaders(v map[string]*string) *SetKeyPolicyResponse {
	s.Headers = v
	return s
}

func (s *SetKeyPolicyResponse) SetStatusCode(v int32) *SetKeyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *SetKeyPolicyResponse) SetBody(v *SetKeyPolicyResponseBody) *SetKeyPolicyResponse {
	s.Body = v
	return s
}

type SetSecretPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {"Version":"1","Statement": [{"Sid":"kms default secret policy","Effect":"Allow","Principal":{"RAM": ["acs:ram::119285303511****:*"]},"Action":["kms:*"],"Resource": ["*"] }] }
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// default
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// secret_test
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s SetSecretPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s SetSecretPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetSecretPolicyRequest) SetPolicy(v string) *SetSecretPolicyRequest {
	s.Policy = &v
	return s
}

func (s *SetSecretPolicyRequest) SetPolicyName(v string) *SetSecretPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *SetSecretPolicyRequest) SetSecretName(v string) *SetSecretPolicyRequest {
	s.SecretName = &v
	return s
}

type SetSecretPolicyResponseBody struct {
	// example:
	//
	// 381D5D33-BB8F-395F-8EE4-AE3BB4B523C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetSecretPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetSecretPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SetSecretPolicyResponseBody) SetRequestId(v string) *SetSecretPolicyResponseBody {
	s.RequestId = &v
	return s
}

type SetSecretPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSecretPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSecretPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s SetSecretPolicyResponse) GoString() string {
	return s.String()
}

func (s *SetSecretPolicyResponse) SetHeaders(v map[string]*string) *SetSecretPolicyResponse {
	s.Headers = v
	return s
}

func (s *SetSecretPolicyResponse) SetStatusCode(v int32) *SetSecretPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSecretPolicyResponse) SetBody(v *SetSecretPolicyResponseBody) *SetSecretPolicyResponse {
	s.Body = v
	return s
}

type TagResourceRequest struct {
	// The ID of the certificate.
	//
	// >  You can configure only one of the KeyId, SecretName, and CertificateId parameters.
	//
	// example:
	//
	// 770dbe42-e146-43d1-a55a-1355db86****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The ID of the customer master key (CMK). The ID must be globally unique.
	//
	// >  You can configure only one of the KeyId, SecretName, and CertificateId parameters.
	//
	// example:
	//
	// 08c33a6f-4e0a-4a1b-a3fa-7ddf****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the secret.
	//
	// >  You can configure only one of the KeyId, SecretName, and CertificateId parameters.
	//
	// example:
	//
	// MyDbC****
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// One or more tags that you want to add. The value is in the array format.
	//
	// Tag attributes:
	//
	// 	- TagKey: the tag key.
	//
	// 	- TagValue: the tag value.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"TagKey":"S1key1","TagValue":"S1val1"},{"TagKey":"S1key2","TagValue":"S2val2"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s TagResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourceRequest) GoString() string {
	return s.String()
}

func (s *TagResourceRequest) SetCertificateId(v string) *TagResourceRequest {
	s.CertificateId = &v
	return s
}

func (s *TagResourceRequest) SetKeyId(v string) *TagResourceRequest {
	s.KeyId = &v
	return s
}

func (s *TagResourceRequest) SetSecretName(v string) *TagResourceRequest {
	s.SecretName = &v
	return s
}

func (s *TagResourceRequest) SetTags(v string) *TagResourceRequest {
	s.Tags = &v
	return s
}

type TagResourceResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4162a6af-bc99-40b3-a552-89dcc8aaf7c8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourceResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourceResponseBody) SetRequestId(v string) *TagResourceResponseBody {
	s.RequestId = &v
	return s
}

type TagResourceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourceResponse) GoString() string {
	return s.String()
}

func (s *TagResourceResponse) SetHeaders(v map[string]*string) *TagResourceResponse {
	s.Headers = v
	return s
}

func (s *TagResourceResponse) SetStatusCode(v int32) *TagResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourceResponse) SetBody(v *TagResourceResponseBody) *TagResourceResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The region ID of the resource.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/601478.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the resources to which you want to add tags. You can enter a maximum of 50 resource IDs.
	//
	// Enter multiple resource IDs in the `["ResourceId. 1","ResourceId. 2",...]` format.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource to which you want to add tags. Valid values:
	//
	// 	- key
	//
	// 	- secret
	//
	// This parameter is required.
	//
	// example:
	//
	// key
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// A list of tags. You can enter up to 20 tags.
	//
	// A tag consists of a key-value pair. Enter multiple tags in the `[{"Key":"key1","Value":"value1"},{"Key":"key2","Value":"value2"},..]` format.
	//
	// This parameter is required.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The key of the tag. A tag consists of a key-value pair.
	//
	// You can enter up to 20 tags. Enter multiple tags in the `[{"Key":"key1","Value":"value1"},{"Key":"key2","Value":"value2"},..]` format.
	//
	// Each key can be up to 128 characters in length and can contain letters, digits, forward slashes (/), backslashes (\\\\), underscores (_), hyphens (-), periods (.), plus signs (+), equal signs (=), colons (:), and at signs (@).
	//
	// >  The key cannot start with aliyun or acs:.
	//
	// example:
	//
	// disk-encryption
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. A tag consists of a key-value pair.
	//
	// You can enter up to 20 tags. Enter multiple tags in the `[{"Key":"key1","Value":"value1"},{"Key":"key2","Value":"value2"},..]` format.
	//
	// Each value can be up to 128 characters in length and can contain letters, digits, forward slashes (/), backslashes (\\\\), underscores (_), hyphens (-), periods (.), plus signs (+), equal signs (=), colons (:), and at signs (@).
	//
	// example:
	//
	// true
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 598d0219-45cd-4477-84ad-85a52d9debcf
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourceRequest struct {
	// example:
	//
	// 770dbe42-e146-43d1-a55a-1355db86****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 08c33a6f-4e0a-4a1b-a3fa-7ddf****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// example:
	//
	// MyDbC****
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["tagkey1","tagkey2"]
	TagKeys *string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
}

func (s UntagResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourceRequest) GoString() string {
	return s.String()
}

func (s *UntagResourceRequest) SetCertificateId(v string) *UntagResourceRequest {
	s.CertificateId = &v
	return s
}

func (s *UntagResourceRequest) SetKeyId(v string) *UntagResourceRequest {
	s.KeyId = &v
	return s
}

func (s *UntagResourceRequest) SetSecretName(v string) *UntagResourceRequest {
	s.SecretName = &v
	return s
}

func (s *UntagResourceRequest) SetTagKeys(v string) *UntagResourceRequest {
	s.TagKeys = &v
	return s
}

type UntagResourceResponseBody struct {
	// example:
	//
	// 4162a6af-bc99-40b3-a552-89dcc8aaf7c8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourceResponseBody) SetRequestId(v string) *UntagResourceResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourceResponse) GoString() string {
	return s.String()
}

func (s *UntagResourceResponse) SetHeaders(v map[string]*string) *UntagResourceResponse {
	s.Headers = v
	return s
}

func (s *UntagResourceResponse) SetStatusCode(v int32) *UntagResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourceResponse) SetBody(v *UntagResourceResponseBody) *UntagResourceResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from resources. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// >  This parameter takes effect only when you specify an empty tag key.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The region ID of the resource.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/601478.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the resources from which you want to remove tags. You can enter up to 50 resource IDs.
	//
	// Enter multiple resource IDs in the `["ResourceId.1","ResourceId.2",...]` format.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource from which you want to remove tags. Valid values:
	//
	// 	- key
	//
	// 	- secret
	//
	// This parameter is required.
	//
	// example:
	//
	// key
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The keys of the tags that you want to remove. You can enter up to 20 tag keys.
	//
	// Enter multiple tag keys in the `["key.1","key.2",...]` format.
	//
	// >  The tag key cannot start with aliyun or acs:.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// b1f210dc-e52c-4a86-b9dd-7492343d46c7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateAliasRequest struct {
	// The alias that you want to bind.
	//
	// The value must be 1 to 255 characters in length and must include the alias/ prefix.
	//
	// This parameter is required.
	//
	// example:
	//
	// alias/example
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s UpdateAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAliasRequest) GoString() string {
	return s.String()
}

func (s *UpdateAliasRequest) SetAliasName(v string) *UpdateAliasRequest {
	s.AliasName = &v
	return s
}

func (s *UpdateAliasRequest) SetKeyId(v string) *UpdateAliasRequest {
	s.KeyId = &v
	return s
}

type UpdateAliasResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 1d2baaf3-d357-46c2-832e-13560c2bd9cd
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAliasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAliasResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAliasResponseBody) SetRequestId(v string) *UpdateAliasResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAliasResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAliasResponse) GoString() string {
	return s.String()
}

func (s *UpdateAliasResponse) SetHeaders(v map[string]*string) *UpdateAliasResponse {
	s.Headers = v
	return s
}

func (s *UpdateAliasResponse) SetStatusCode(v int32) *UpdateAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAliasResponse) SetBody(v *UpdateAliasResponseBody) *UpdateAliasResponse {
	s.Body = v
	return s
}

type UpdateApplicationAccessPointRequest struct {
	// The description.
	//
	// example:
	//
	// aap description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the AAP that you want to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// aap_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The permission policy that you want to update.
	//
	// > You can associate up to three permission policies with each AAP.
	//
	// example:
	//
	// ["kst-hzz62ee817bvyyr5x****.efkd","kst-hzz62ee817bvyyr5x****.eyyp"]
	Policies *string `json:"Policies,omitempty" xml:"Policies,omitempty"`
}

func (s UpdateApplicationAccessPointRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationAccessPointRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationAccessPointRequest) SetDescription(v string) *UpdateApplicationAccessPointRequest {
	s.Description = &v
	return s
}

func (s *UpdateApplicationAccessPointRequest) SetName(v string) *UpdateApplicationAccessPointRequest {
	s.Name = &v
	return s
}

func (s *UpdateApplicationAccessPointRequest) SetPolicies(v string) *UpdateApplicationAccessPointRequest {
	s.Policies = &v
	return s
}

type UpdateApplicationAccessPointResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// bcfefe15-46f0-44a3-bd96-3d422474b71a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationAccessPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationAccessPointResponseBody) SetRequestId(v string) *UpdateApplicationAccessPointResponseBody {
	s.RequestId = &v
	return s
}

type UpdateApplicationAccessPointResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationAccessPointResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationAccessPointResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationAccessPointResponse) SetHeaders(v map[string]*string) *UpdateApplicationAccessPointResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationAccessPointResponse) SetStatusCode(v int32) *UpdateApplicationAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationAccessPointResponse) SetBody(v *UpdateApplicationAccessPointResponseBody) *UpdateApplicationAccessPointResponse {
	s.Body = v
	return s
}

type UpdateCertificateStatusRequest struct {
	// The ID of the certificate. The ID must be globally unique in Certificates Manager.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9a28de48-8d8b-484d-a766-dec4****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The status of the certificate. Valid values:
	//
	// 	- INACTIVE: The certificate is disabled.
	//
	// 	- ACTIVE: The certificate is enabled.
	//
	// 	- REVOKED: The certificate is revoked.
	//
	// > If the certificate is in the REVOKED state, you can use the certificate only to verify a signature, but not to generate a signature.
	//
	// This parameter is required.
	//
	// example:
	//
	// INACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateCertificateStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCertificateStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateCertificateStatusRequest) SetCertificateId(v string) *UpdateCertificateStatusRequest {
	s.CertificateId = &v
	return s
}

func (s *UpdateCertificateStatusRequest) SetStatus(v string) *UpdateCertificateStatusRequest {
	s.Status = &v
	return s
}

type UpdateCertificateStatusResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// e3f57fe0-9ded-40b0-9caf-a3815f2148c1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCertificateStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCertificateStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCertificateStatusResponseBody) SetRequestId(v string) *UpdateCertificateStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCertificateStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCertificateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCertificateStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCertificateStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateCertificateStatusResponse) SetHeaders(v map[string]*string) *UpdateCertificateStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateCertificateStatusResponse) SetStatusCode(v int32) *UpdateCertificateStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCertificateStatusResponse) SetBody(v *UpdateCertificateStatusResponseBody) *UpdateCertificateStatusResponse {
	s.Body = v
	return s
}

type UpdateKeyDescriptionRequest struct {
	// The description of the CMK. This description includes the purpose of the CMK, such as the types of data that you want to protect and applications that can use the CMK.
	//
	// This parameter is required.
	//
	// example:
	//
	// key description example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s UpdateKeyDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateKeyDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateKeyDescriptionRequest) SetDescription(v string) *UpdateKeyDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateKeyDescriptionRequest) SetKeyId(v string) *UpdateKeyDescriptionRequest {
	s.KeyId = &v
	return s
}

type UpdateKeyDescriptionResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3455b9b4-95c1-419d-b310-db6a53b09a39
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateKeyDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateKeyDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKeyDescriptionResponseBody) SetRequestId(v string) *UpdateKeyDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateKeyDescriptionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKeyDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKeyDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateKeyDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateKeyDescriptionResponse) SetHeaders(v map[string]*string) *UpdateKeyDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateKeyDescriptionResponse) SetStatusCode(v int32) *UpdateKeyDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKeyDescriptionResponse) SetBody(v *UpdateKeyDescriptionResponseBody) *UpdateKeyDescriptionResponse {
	s.Body = v
	return s
}

type UpdateKmsInstanceBindVpcRequest struct {
	// The VPC configuration. The configuration of each VPC contains the following content:
	//
	// 	- VpcId: the ID of the VPC.
	//
	// 	- VSwitchId: the vSwitch in the VPC.
	//
	// 	- RegionID: the ID of the region to which the VPC belongs.
	//
	// 	- VpcOwnerId: the Alibaba Cloud account to which the VPC belongs.
	//
	// Format: `[{"VpcId":"${VpcId}","VSwitchId":"${VSwitchId}","RegionId":"${RegionId}","VpcOwnerId":${VpcOwnerId}},..]`.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"VpcId":"vpc-bp1go9qvmj78j4f4c****","VSwitchId":"vsw-bp16c5pvvcf0fp5b9****","RegionId":"cn-hangzhou","VpcOwnerId":120708975881****},{"VpcId":"vpc-bp14c07ucxg6h1xjm****","VSwitchId":"vsw-bp1wujtnspi1l3gvu****","RegionId":"cn-hangzhou","VpcOwnerId":119285303511****}]
	BindVpcs *string `json:"BindVpcs,omitempty" xml:"BindVpcs,omitempty"`
	// The ID of the KMS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// kst-phzz64f722a1buamw0****
	KmsInstanceId *string `json:"KmsInstanceId,omitempty" xml:"KmsInstanceId,omitempty"`
}

func (s UpdateKmsInstanceBindVpcRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateKmsInstanceBindVpcRequest) GoString() string {
	return s.String()
}

func (s *UpdateKmsInstanceBindVpcRequest) SetBindVpcs(v string) *UpdateKmsInstanceBindVpcRequest {
	s.BindVpcs = &v
	return s
}

func (s *UpdateKmsInstanceBindVpcRequest) SetKmsInstanceId(v string) *UpdateKmsInstanceBindVpcRequest {
	s.KmsInstanceId = &v
	return s
}

type UpdateKmsInstanceBindVpcResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// d3eca5c8-a856-4347-8eb6-e1898c3fda2e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateKmsInstanceBindVpcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateKmsInstanceBindVpcResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKmsInstanceBindVpcResponseBody) SetRequestId(v string) *UpdateKmsInstanceBindVpcResponseBody {
	s.RequestId = &v
	return s
}

type UpdateKmsInstanceBindVpcResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKmsInstanceBindVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKmsInstanceBindVpcResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateKmsInstanceBindVpcResponse) GoString() string {
	return s.String()
}

func (s *UpdateKmsInstanceBindVpcResponse) SetHeaders(v map[string]*string) *UpdateKmsInstanceBindVpcResponse {
	s.Headers = v
	return s
}

func (s *UpdateKmsInstanceBindVpcResponse) SetStatusCode(v int32) *UpdateKmsInstanceBindVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKmsInstanceBindVpcResponse) SetBody(v *UpdateKmsInstanceBindVpcResponseBody) *UpdateKmsInstanceBindVpcResponse {
	s.Body = v
	return s
}

type UpdateNetworkRuleRequest struct {
	// The description after the update.
	//
	// example:
	//
	// Creat by kst-hzz62ee817bvyyr5****
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the access control rule that you want to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// networkrule_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The private IP address or CIDR block after the update. Separate multiple items with commas (,).
	//
	// example:
	//
	// ["192.10.XX.XX","192.168.XX.XX/24"]
	SourcePrivateIp *string `json:"SourcePrivateIp,omitempty" xml:"SourcePrivateIp,omitempty"`
}

func (s UpdateNetworkRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNetworkRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateNetworkRuleRequest) SetDescription(v string) *UpdateNetworkRuleRequest {
	s.Description = &v
	return s
}

func (s *UpdateNetworkRuleRequest) SetName(v string) *UpdateNetworkRuleRequest {
	s.Name = &v
	return s
}

func (s *UpdateNetworkRuleRequest) SetSourcePrivateIp(v string) *UpdateNetworkRuleRequest {
	s.SourcePrivateIp = &v
	return s
}

type UpdateNetworkRuleResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3bf02f7a-015b-4f34-be0f-cc043fda2d85
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNetworkRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateNetworkRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNetworkRuleResponseBody) SetRequestId(v string) *UpdateNetworkRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateNetworkRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNetworkRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNetworkRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateNetworkRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateNetworkRuleResponse) SetHeaders(v map[string]*string) *UpdateNetworkRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateNetworkRuleResponse) SetStatusCode(v int32) *UpdateNetworkRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNetworkRuleResponse) SetBody(v *UpdateNetworkRuleResponseBody) *UpdateNetworkRuleResponse {
	s.Body = v
	return s
}

type UpdatePolicyRequest struct {
	// The access control rule.
	//
	// > For more information about how to query created access control rules, see [ListNetworkRules](https://help.aliyun.com/document_detail/2539433.html).
	//
	// example:
	//
	// {"NetworkRules":["kst-hzz62ee817bvyyr5x****.efkd","kst-hzz62ee817bvyyr5x****.eyyp"]}
	AccessControlRules *string `json:"AccessControlRules,omitempty" xml:"AccessControlRules,omitempty"`
	// The description.
	//
	// example:
	//
	// policy  description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission policy that you want to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// policy_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operations that are supported by the updated policy. Valid values:
	//
	// 	- RbacPermission/Template/CryptoServiceKeyUser: allows you to perform cryptographic operations.
	//
	// 	- RbacPermission/Template/CryptoServiceSecretUser: allows you to perform secret-related operations.
	//
	// You can select both.
	//
	// example:
	//
	// ["RbacPermission/Template/CryptoServiceKeyUser", "RbacPermission/Template/CryptoServiceSecretUser"]
	Permissions *string `json:"Permissions,omitempty" xml:"Permissions,omitempty"`
	// The key and secret that are allowed to access after the update.
	//
	// 	- Key: Enter a key in the `key/${KeyId}` format. To allow access to all keys of a KMS instance, enter key/\\*.
	//
	// 	- Secret: Enter a secret in the `secret/${SecretName}` format. To allow access to all secrets of a KMS instance, enter secret/\\*.
	//
	// example:
	//
	// ["secret/acs/ram/user/ram-secret", "secret/acs/ram/user/acr-master", "key/key-hzz63d9c8d3dfv8cv****"]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s UpdatePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolicyRequest) SetAccessControlRules(v string) *UpdatePolicyRequest {
	s.AccessControlRules = &v
	return s
}

func (s *UpdatePolicyRequest) SetDescription(v string) *UpdatePolicyRequest {
	s.Description = &v
	return s
}

func (s *UpdatePolicyRequest) SetName(v string) *UpdatePolicyRequest {
	s.Name = &v
	return s
}

func (s *UpdatePolicyRequest) SetPermissions(v string) *UpdatePolicyRequest {
	s.Permissions = &v
	return s
}

func (s *UpdatePolicyRequest) SetResources(v string) *UpdatePolicyRequest {
	s.Resources = &v
	return s
}

type UpdatePolicyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// f455324b-e229-4066-9f58-9c1cf3fe83a8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePolicyResponseBody) SetRequestId(v string) *UpdatePolicyResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdatePolicyResponse) SetHeaders(v map[string]*string) *UpdatePolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdatePolicyResponse) SetStatusCode(v int32) *UpdatePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePolicyResponse) SetBody(v *UpdatePolicyResponseBody) *UpdatePolicyResponse {
	s.Body = v
	return s
}

type UpdateRotationPolicyRequest struct {
	// Specifies whether to enable automatic key rotation. Valid values:
	//
	// 	- true: enables automatic key rotation.
	//
	// 	- false: disables automatic key rotation.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableAutomaticRotation *bool `json:"EnableAutomaticRotation,omitempty" xml:"EnableAutomaticRotation,omitempty"`
	// The ID of the customer master key (CMK). The ID must be globally unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The period of automatic key rotation. Specify the value in the integer[unit] format. The following units are supported: d (day), h (hour), m (minute), and s (second). For example, you can use either 7d or 604800s to specify a seven-day period. The period can range from 7 days to 730 days.
	//
	// >  If you set the EnableAutomaticRotation parameter to true, you must also specify this parameter. If you set the EnableAutomaticRotation parameter to false, you can leave this parameter unspecified.
	//
	// example:
	//
	// 30d
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
}

func (s UpdateRotationPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRotationPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateRotationPolicyRequest) SetEnableAutomaticRotation(v bool) *UpdateRotationPolicyRequest {
	s.EnableAutomaticRotation = &v
	return s
}

func (s *UpdateRotationPolicyRequest) SetKeyId(v string) *UpdateRotationPolicyRequest {
	s.KeyId = &v
	return s
}

func (s *UpdateRotationPolicyRequest) SetRotationInterval(v string) *UpdateRotationPolicyRequest {
	s.RotationInterval = &v
	return s
}

type UpdateRotationPolicyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// efb1cbbd-a093-4278-bc03-639dd4fcc207
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRotationPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRotationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRotationPolicyResponseBody) SetRequestId(v string) *UpdateRotationPolicyResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRotationPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRotationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRotationPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRotationPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateRotationPolicyResponse) SetHeaders(v map[string]*string) *UpdateRotationPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateRotationPolicyResponse) SetStatusCode(v int32) *UpdateRotationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRotationPolicyResponse) SetBody(v *UpdateRotationPolicyResponseBody) *UpdateRotationPolicyResponse {
	s.Body = v
	return s
}

type UpdateSecretRequest struct {
	ExtendedConfig *UpdateSecretRequestExtendedConfig `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty" type:"Struct"`
	// The description of the secret.
	//
	// example:
	//
	// datainfo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s UpdateSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecretRequest) SetExtendedConfig(v *UpdateSecretRequestExtendedConfig) *UpdateSecretRequest {
	s.ExtendedConfig = v
	return s
}

func (s *UpdateSecretRequest) SetDescription(v string) *UpdateSecretRequest {
	s.Description = &v
	return s
}

func (s *UpdateSecretRequest) SetSecretName(v string) *UpdateSecretRequest {
	s.SecretName = &v
	return s
}

type UpdateSecretRequestExtendedConfig struct {
	// The custom data in the extended configuration of the secret.
	//
	// > 	- If this parameter is specified, the existing extended configuration of the secret is updated.
	//
	// > 	- This parameter is unavailable for generic secrets.
	//
	// example:
	//
	// {"DBName":"app1","Port":"3306"}
	CustomData map[string]interface{} `json:"CustomData,omitempty" xml:"CustomData,omitempty"`
}

func (s UpdateSecretRequestExtendedConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretRequestExtendedConfig) GoString() string {
	return s.String()
}

func (s *UpdateSecretRequestExtendedConfig) SetCustomData(v map[string]interface{}) *UpdateSecretRequestExtendedConfig {
	s.CustomData = v
	return s
}

type UpdateSecretShrinkRequest struct {
	ExtendedConfig *UpdateSecretShrinkRequestExtendedConfig `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty" type:"Struct"`
	// The description of the secret.
	//
	// example:
	//
	// datainfo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s UpdateSecretShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecretShrinkRequest) SetExtendedConfig(v *UpdateSecretShrinkRequestExtendedConfig) *UpdateSecretShrinkRequest {
	s.ExtendedConfig = v
	return s
}

func (s *UpdateSecretShrinkRequest) SetDescription(v string) *UpdateSecretShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateSecretShrinkRequest) SetSecretName(v string) *UpdateSecretShrinkRequest {
	s.SecretName = &v
	return s
}

type UpdateSecretShrinkRequestExtendedConfig struct {
	// The custom data in the extended configuration of the secret.
	//
	// > 	- If this parameter is specified, the existing extended configuration of the secret is updated.
	//
	// > 	- This parameter is unavailable for generic secrets.
	//
	// example:
	//
	// {"DBName":"app1","Port":"3306"}
	CustomData *string `json:"CustomData,omitempty" xml:"CustomData,omitempty"`
}

func (s UpdateSecretShrinkRequestExtendedConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretShrinkRequestExtendedConfig) GoString() string {
	return s.String()
}

func (s *UpdateSecretShrinkRequestExtendedConfig) SetCustomData(v string) *UpdateSecretShrinkRequestExtendedConfig {
	s.CustomData = &v
	return s
}

type UpdateSecretResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5b75d8b1-5b6a-4ec0-8e0c-c08befdfad47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the secret.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s UpdateSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecretResponseBody) SetRequestId(v string) *UpdateSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSecretResponseBody) SetSecretName(v string) *UpdateSecretResponseBody {
	s.SecretName = &v
	return s
}

type UpdateSecretResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretResponse) GoString() string {
	return s.String()
}

func (s *UpdateSecretResponse) SetHeaders(v map[string]*string) *UpdateSecretResponse {
	s.Headers = v
	return s
}

func (s *UpdateSecretResponse) SetStatusCode(v int32) *UpdateSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSecretResponse) SetBody(v *UpdateSecretResponseBody) *UpdateSecretResponse {
	s.Body = v
	return s
}

type UpdateSecretRotationPolicyRequest struct {
	// Specifies whether to enable automatic rotation. Valid values:
	//
	// 	- true: enables automatic rotation.
	//
	// 	- false: does not enable automatic rotation. This is the default value.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableAutomaticRotation *bool `json:"EnableAutomaticRotation,omitempty" xml:"EnableAutomaticRotation,omitempty"`
	// The interval for automatic rotation. Valid values: 6 hours to 8,760 hours (365 days).
	//
	// The value is in the `integer[unit]` format.````
	//
	// The unit can be d (day), h (hour), m (minute), or s (second). For example, both 7d and 604800s indicate a seven-day interval.
	//
	// >  This parameter is required if you set the EnableAutomaticRotation parameter to true. This parameter is ignored if you set the EnableAutomaticRotation parameter to false or does not specify the EnableAutomaticRotation parameter.
	//
	// example:
	//
	// 30d
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	// The name of the secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// RdsSecret/Mysql5.4/MyCred
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s UpdateSecretRotationPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretRotationPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecretRotationPolicyRequest) SetEnableAutomaticRotation(v bool) *UpdateSecretRotationPolicyRequest {
	s.EnableAutomaticRotation = &v
	return s
}

func (s *UpdateSecretRotationPolicyRequest) SetRotationInterval(v string) *UpdateSecretRotationPolicyRequest {
	s.RotationInterval = &v
	return s
}

func (s *UpdateSecretRotationPolicyRequest) SetSecretName(v string) *UpdateSecretRotationPolicyRequest {
	s.SecretName = &v
	return s
}

type UpdateSecretRotationPolicyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 2c124f6f-4210-499f-b88a-69f54004d2d8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the secret.
	//
	// example:
	//
	// RdsSecret/Mysql5.4/MyCred
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s UpdateSecretRotationPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretRotationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecretRotationPolicyResponseBody) SetRequestId(v string) *UpdateSecretRotationPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSecretRotationPolicyResponseBody) SetSecretName(v string) *UpdateSecretRotationPolicyResponseBody {
	s.SecretName = &v
	return s
}

type UpdateSecretRotationPolicyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSecretRotationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSecretRotationPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretRotationPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateSecretRotationPolicyResponse) SetHeaders(v map[string]*string) *UpdateSecretRotationPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateSecretRotationPolicyResponse) SetStatusCode(v int32) *UpdateSecretRotationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSecretRotationPolicyResponse) SetBody(v *UpdateSecretRotationPolicyResponseBody) *UpdateSecretRotationPolicyResponse {
	s.Body = v
	return s
}

type UpdateSecretVersionStageRequest struct {
	// The version from which you want to remove the specified stage label.
	//
	// >  You must specify at least one of the RemoveFromVersion and MoveToVersion parameters.
	//
	// example:
	//
	// 002
	MoveToVersion *string `json:"MoveToVersion,omitempty" xml:"MoveToVersion,omitempty"`
	// The specified stage label. Valid values:
	//
	// 	- ACSCurrent
	//
	// 	- ACSPrevious
	//
	// 	- Custom stage label
	//
	// example:
	//
	// 001
	RemoveFromVersion *string `json:"RemoveFromVersion,omitempty" xml:"RemoveFromVersion,omitempty"`
	// The operation that you want to perform. Set the value to **UpdateSecretVersionStage**.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The name of the secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACSCurrent
	VersionStage *string `json:"VersionStage,omitempty" xml:"VersionStage,omitempty"`
}

func (s UpdateSecretVersionStageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretVersionStageRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecretVersionStageRequest) SetMoveToVersion(v string) *UpdateSecretVersionStageRequest {
	s.MoveToVersion = &v
	return s
}

func (s *UpdateSecretVersionStageRequest) SetRemoveFromVersion(v string) *UpdateSecretVersionStageRequest {
	s.RemoveFromVersion = &v
	return s
}

func (s *UpdateSecretVersionStageRequest) SetSecretName(v string) *UpdateSecretVersionStageRequest {
	s.SecretName = &v
	return s
}

func (s *UpdateSecretVersionStageRequest) SetVersionStage(v string) *UpdateSecretVersionStageRequest {
	s.VersionStage = &v
	return s
}

type UpdateSecretVersionStageResponseBody struct {
	// The name of the secret.
	//
	// example:
	//
	// 8cad259f-4d77-40ec-bbd7-b9c47a423bb9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The version to which you want to apply the specified stage label.
	//
	// > 	- You must specify at least one of the RemoveFromVersion and MoveToVersion parameters.
	//
	// > 	- If the VersionStage parameter is set to ACSCurrent or ACSPrevious, this parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s UpdateSecretVersionStageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretVersionStageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecretVersionStageResponseBody) SetRequestId(v string) *UpdateSecretVersionStageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSecretVersionStageResponseBody) SetSecretName(v string) *UpdateSecretVersionStageResponseBody {
	s.SecretName = &v
	return s
}

type UpdateSecretVersionStageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSecretVersionStageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSecretVersionStageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretVersionStageResponse) GoString() string {
	return s.String()
}

func (s *UpdateSecretVersionStageResponse) SetHeaders(v map[string]*string) *UpdateSecretVersionStageResponse {
	s.Headers = v
	return s
}

func (s *UpdateSecretVersionStageResponse) SetStatusCode(v int32) *UpdateSecretVersionStageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSecretVersionStageResponse) SetBody(v *UpdateSecretVersionStageResponseBody) *UpdateSecretVersionStageResponse {
	s.Body = v
	return s
}

type UploadCertificateRequest struct {
	// The certificate issued by the CA, which is in the Privacy Enhanced Mail (PEM) format.
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----  (X.509 Certificate PEM Content)  -----END CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The certificate chain issued by the CA, which is in the PEM format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----  (Sub CA Certificate PEM Content)  -----END CERTIFICATE-----  -----BEGIN CERTIFICATE-----  (Sub CA Certificate PEM Content)  -----END CERTIFICATE-----  -----BEGIN CERTIFICATE-----  (Root CA Certificate PEM Content)  -----END CERTIFICATE-----
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The ID of the certificate. The ID must be globally unique in Certificates Manager.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s UploadCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadCertificateRequest) GoString() string {
	return s.String()
}

func (s *UploadCertificateRequest) SetCertificate(v string) *UploadCertificateRequest {
	s.Certificate = &v
	return s
}

func (s *UploadCertificateRequest) SetCertificateChain(v string) *UploadCertificateRequest {
	s.CertificateChain = &v
	return s
}

func (s *UploadCertificateRequest) SetCertificateId(v string) *UploadCertificateRequest {
	s.CertificateId = &v
	return s
}

type UploadCertificateResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 15a735a1-8fe6-45cc-a64c-3c4ff839334e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *UploadCertificateResponseBody) SetRequestId(v string) *UploadCertificateResponseBody {
	s.RequestId = &v
	return s
}

type UploadCertificateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadCertificateResponse) GoString() string {
	return s.String()
}

func (s *UploadCertificateResponse) SetHeaders(v map[string]*string) *UploadCertificateResponse {
	s.Headers = v
	return s
}

func (s *UploadCertificateResponse) SetStatusCode(v int32) *UploadCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadCertificateResponse) SetBody(v *UploadCertificateResponseBody) *UploadCertificateResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.ProductId = tea.String("Kms")
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("kms"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Decrypts data by using an asymmetric key.
//
// Description:
//
// This operation supports only asymmetric keys for which the **Usage*	- parameter is set to **ENCRYPT/DECRYPT**. The following table lists supported encryption algorithms.
//
// | KeySpec | Algorithm | Description | Maximum length in bytes |
//
// | ------- | --------- | ----------- | ----------------------- |
//
// | RSA_2048 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256 | 256 |
//
// | RSA_2048 | RSAES_OAEP_SHA_1 | RSAES-OAEP using SHA1 and MGF1 with SHA1 | 256 |
//
// | RSA_3072 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256 | 384 |
//
// | RSA_3072 | RSAES_OAEP_SHA_1 | RSAES-OAEP using SHA1 and MGF1 with SHA1 | 384 |
//
// | EC_SM2 | SM2PKE | SM2 public key encryption algorithm based on elliptic curves | 6144 |
//
// In this example, the asymmetric key whose ID is `5c438b18-05be-40ad-b6c2-3be6752c****` and version ID is `2ab1a983-7072-4bbc-a582-584b5bd8****` and the decryption algorithm `RSAES_OAEP_SHA_1` are used to decrypt the ciphertext `BQKP+1zK6+ZEMxTP5qaVzcsgXtWplYBKm0NXdSnB5FzliFxE1bSiu4dnEIlca2JpeH7yz1/S6fed630H+hIH6DoM25fTLNcKj+mFB0Xnh9m2+HN59Mn4qyTfcUeadnfCXSWcGBouhXFwcdd2rJ3n337bzTf4jm659gZu3L0i6PLuxM9p7mqdwO0cKJPfGVfhnfMz+f4alMg79WB/NNyE2lyX7/qxvV49ObNrrJbKSFiz8Djocaf0IESNLMbfYI5bXjWkJlX92DQbKhibtQW8ZOJ//ZC6t0AWcUoKL6QDm/dg5koQalcleRinpB+QadFm894sLbVZ9+N4GVsv1W****==`.
//
// @param request - AsymmetricDecryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AsymmetricDecryptResponse
func (client *Client) AsymmetricDecryptWithOptions(request *AsymmetricDecryptRequest, runtime *util.RuntimeOptions) (_result *AsymmetricDecryptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		query["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.CiphertextBlob)) {
		query["CiphertextBlob"] = request.CiphertextBlob
	}

	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.KeyVersionId)) {
		query["KeyVersionId"] = request.KeyVersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AsymmetricDecrypt"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AsymmetricDecryptResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AsymmetricDecryptResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Decrypts data by using an asymmetric key.
//
// Description:
//
// This operation supports only asymmetric keys for which the **Usage*	- parameter is set to **ENCRYPT/DECRYPT**. The following table lists supported encryption algorithms.
//
// | KeySpec | Algorithm | Description | Maximum length in bytes |
//
// | ------- | --------- | ----------- | ----------------------- |
//
// | RSA_2048 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256 | 256 |
//
// | RSA_2048 | RSAES_OAEP_SHA_1 | RSAES-OAEP using SHA1 and MGF1 with SHA1 | 256 |
//
// | RSA_3072 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256 | 384 |
//
// | RSA_3072 | RSAES_OAEP_SHA_1 | RSAES-OAEP using SHA1 and MGF1 with SHA1 | 384 |
//
// | EC_SM2 | SM2PKE | SM2 public key encryption algorithm based on elliptic curves | 6144 |
//
// In this example, the asymmetric key whose ID is `5c438b18-05be-40ad-b6c2-3be6752c****` and version ID is `2ab1a983-7072-4bbc-a582-584b5bd8****` and the decryption algorithm `RSAES_OAEP_SHA_1` are used to decrypt the ciphertext `BQKP+1zK6+ZEMxTP5qaVzcsgXtWplYBKm0NXdSnB5FzliFxE1bSiu4dnEIlca2JpeH7yz1/S6fed630H+hIH6DoM25fTLNcKj+mFB0Xnh9m2+HN59Mn4qyTfcUeadnfCXSWcGBouhXFwcdd2rJ3n337bzTf4jm659gZu3L0i6PLuxM9p7mqdwO0cKJPfGVfhnfMz+f4alMg79WB/NNyE2lyX7/qxvV49ObNrrJbKSFiz8Djocaf0IESNLMbfYI5bXjWkJlX92DQbKhibtQW8ZOJ//ZC6t0AWcUoKL6QDm/dg5koQalcleRinpB+QadFm894sLbVZ9+N4GVsv1W****==`.
//
// @param request - AsymmetricDecryptRequest
//
// @return AsymmetricDecryptResponse
func (client *Client) AsymmetricDecrypt(request *AsymmetricDecryptRequest) (_result *AsymmetricDecryptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AsymmetricDecryptResponse{}
	_body, _err := client.AsymmetricDecryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Encrypts data by using an asymmetric customer master key (CMK).
//
// Description:
//
// This operation is supported only for asymmetric keys for which the **Usage*	- parameter is set to **ENCRYPT/DECRYPT**. The following table lists the supported encryption algorithms:
//
// | KeySpec | Algorithm | Description | Maximum number of bytes that can be encrypted |
//
// | ------- | --------- | ----------- | --------------------------------------------- |
//
// | RSA_2048 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256 | 190 |
//
// | RSA_2048 | RSAES_OAEP_SHA_1 | RSAES-OAEP using SHA1 and MGF1 with SHA1 | 214 |
//
// | RSA_3072 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256 | 318 |
//
// | RSA_3072 | RSAES_OAEP_SHA_1 | RSAES-OAEP using SHA1 and MGF1 with SHA1 | 342 |
//
// | EC_SM2 | SM2PKE | SM2 public key encryption algorithm based on elliptic curves | 6047 |
//
// You can use the asymmetric CMK whose ID is `5c438b18-05be-40ad-b6c2-3be6752c****` and version ID is `2ab1a983-7072-4bbc-a582-584b5bd8****` and the algorithm `RSAES_OAEP_SHA_1` to encrypt the plaintext `SGVsbG8gd29ybGQ=` based on the parameter settings provided in this topic.
//
// @param request - AsymmetricEncryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AsymmetricEncryptResponse
func (client *Client) AsymmetricEncryptWithOptions(request *AsymmetricEncryptRequest, runtime *util.RuntimeOptions) (_result *AsymmetricEncryptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		query["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.KeyVersionId)) {
		query["KeyVersionId"] = request.KeyVersionId
	}

	if !tea.BoolValue(util.IsUnset(request.Plaintext)) {
		query["Plaintext"] = request.Plaintext
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AsymmetricEncrypt"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AsymmetricEncryptResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AsymmetricEncryptResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Encrypts data by using an asymmetric customer master key (CMK).
//
// Description:
//
// This operation is supported only for asymmetric keys for which the **Usage*	- parameter is set to **ENCRYPT/DECRYPT**. The following table lists the supported encryption algorithms:
//
// | KeySpec | Algorithm | Description | Maximum number of bytes that can be encrypted |
//
// | ------- | --------- | ----------- | --------------------------------------------- |
//
// | RSA_2048 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256 | 190 |
//
// | RSA_2048 | RSAES_OAEP_SHA_1 | RSAES-OAEP using SHA1 and MGF1 with SHA1 | 214 |
//
// | RSA_3072 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256 | 318 |
//
// | RSA_3072 | RSAES_OAEP_SHA_1 | RSAES-OAEP using SHA1 and MGF1 with SHA1 | 342 |
//
// | EC_SM2 | SM2PKE | SM2 public key encryption algorithm based on elliptic curves | 6047 |
//
// You can use the asymmetric CMK whose ID is `5c438b18-05be-40ad-b6c2-3be6752c****` and version ID is `2ab1a983-7072-4bbc-a582-584b5bd8****` and the algorithm `RSAES_OAEP_SHA_1` to encrypt the plaintext `SGVsbG8gd29ybGQ=` based on the parameter settings provided in this topic.
//
// @param request - AsymmetricEncryptRequest
//
// @return AsymmetricEncryptResponse
func (client *Client) AsymmetricEncrypt(request *AsymmetricEncryptRequest) (_result *AsymmetricEncryptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AsymmetricEncryptResponse{}
	_body, _err := client.AsymmetricEncryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// AsymmetricSign
//
// Description:
//
// Generates a signature by using an asymmetric key.
//
// @param request - AsymmetricSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AsymmetricSignResponse
func (client *Client) AsymmetricSignWithOptions(request *AsymmetricSignRequest, runtime *util.RuntimeOptions) (_result *AsymmetricSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		query["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.Digest)) {
		query["Digest"] = request.Digest
	}

	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.KeyVersionId)) {
		query["KeyVersionId"] = request.KeyVersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AsymmetricSign"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AsymmetricSignResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AsymmetricSignResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// AsymmetricSign
//
// Description:
//
// Generates a signature by using an asymmetric key.
//
// @param request - AsymmetricSignRequest
//
// @return AsymmetricSignResponse
func (client *Client) AsymmetricSign(request *AsymmetricSignRequest) (_result *AsymmetricSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AsymmetricSignResponse{}
	_body, _err := client.AsymmetricSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies a signature by using an asymmetric key.
//
// Description:
//
// This operation supports only asymmetric keys for which the **Usage*	- parameter is set to **SIGN/VERIFY**. The following table describes the supported signature algorithms.
//
// | KeySpec | Algorithm | Description |
//
// | ------- | --------- | ----------- |
//
// | RSA_2048 | RSA_PSS_SHA_256 | RSASSA-PSS using SHA-256 and MGF1 with SHA-256 |
//
// | RSA_2048 | RSA_PKCS1_SHA_256 | RSASSA-PKCS1-v1_5 using SHA-256 |
//
// | RSA_3072 | RSA_PSS_SHA_256 | RSASSA-PSS using SHA-256 and MGF1 with SHA-256 |
//
// | RSA_3072 | RSA_PKCS1_SHA_256 | RSASSA-PKCS1-v1_5 using SHA-256 |
//
// | EC_P256 | ECDSA_SHA_256 | ECDSA on the P-256 Curve(secp256r1) with a SHA-256 digest |
//
// | EC_P256K | ECDSA_SHA_256 | ECDSA on the P-256K Curve(secp256k1) with a SHA-256 digest |
//
// | EC_SM2 | SM2DSA | SM2 elliptic curve public key encryption algorithm |
//
// >  When you calculate the SM2 signature based on GB/T 32918, the **Digest*	- parameter is used to calculate the digest value of the combination of Z(A) and M, rather than the SM3 digest value. M indicates the original message to be signed. Z(A) indicates the hash value for User A. The hash value is defined in GB/T 32918.  In this example, the asymmetric key whose ID is `5c438b18-05be-40ad-b6c2-3be6752c****` and version ID is `2ab1a983-7072-4bbc-a582-584b5bd8****` and the signature algorithm RSA_PSS_SHA_256 are used to verify the signature `M2CceNZH00ZgL9ED/ZHFp21YRAvYeZHknJUc207OCZ0N9wNn9As4z2bON3FF3je+1Nu+2+/8Zj50HpMTpzYpMp2R93cYmACCmhaYoKydxylbyGzJR8y9likZRCrkD38lRoS40aBBvv/6iRKzQuo9EGYVcel36cMNg00VmYNBy3pa1rwg3gA4l3cy6kjayZja1WGPkVhrVKsrJMdbpl0ApLjXKuD8rw1n1XLCwCUEL5eLPljTZaAveqdOFQOiZnZEGI27qIiZe7I1fN8tcz6anS/gTM7xRKE++5egEvRWlTQQTJeApnPSiUPA+8ZykNdelQsOQh5SrGoyI4A5pq****==` of the digest `ZOyIygCyaOW6GjVnihtTFtIS9PNmskdyMlNKiuyjfzw=`.
//
// @param request - AsymmetricVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AsymmetricVerifyResponse
func (client *Client) AsymmetricVerifyWithOptions(request *AsymmetricVerifyRequest, runtime *util.RuntimeOptions) (_result *AsymmetricVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		query["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.Digest)) {
		query["Digest"] = request.Digest
	}

	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.KeyVersionId)) {
		query["KeyVersionId"] = request.KeyVersionId
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AsymmetricVerify"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AsymmetricVerifyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AsymmetricVerifyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Verifies a signature by using an asymmetric key.
//
// Description:
//
// This operation supports only asymmetric keys for which the **Usage*	- parameter is set to **SIGN/VERIFY**. The following table describes the supported signature algorithms.
//
// | KeySpec | Algorithm | Description |
//
// | ------- | --------- | ----------- |
//
// | RSA_2048 | RSA_PSS_SHA_256 | RSASSA-PSS using SHA-256 and MGF1 with SHA-256 |
//
// | RSA_2048 | RSA_PKCS1_SHA_256 | RSASSA-PKCS1-v1_5 using SHA-256 |
//
// | RSA_3072 | RSA_PSS_SHA_256 | RSASSA-PSS using SHA-256 and MGF1 with SHA-256 |
//
// | RSA_3072 | RSA_PKCS1_SHA_256 | RSASSA-PKCS1-v1_5 using SHA-256 |
//
// | EC_P256 | ECDSA_SHA_256 | ECDSA on the P-256 Curve(secp256r1) with a SHA-256 digest |
//
// | EC_P256K | ECDSA_SHA_256 | ECDSA on the P-256K Curve(secp256k1) with a SHA-256 digest |
//
// | EC_SM2 | SM2DSA | SM2 elliptic curve public key encryption algorithm |
//
// >  When you calculate the SM2 signature based on GB/T 32918, the **Digest*	- parameter is used to calculate the digest value of the combination of Z(A) and M, rather than the SM3 digest value. M indicates the original message to be signed. Z(A) indicates the hash value for User A. The hash value is defined in GB/T 32918.  In this example, the asymmetric key whose ID is `5c438b18-05be-40ad-b6c2-3be6752c****` and version ID is `2ab1a983-7072-4bbc-a582-584b5bd8****` and the signature algorithm RSA_PSS_SHA_256 are used to verify the signature `M2CceNZH00ZgL9ED/ZHFp21YRAvYeZHknJUc207OCZ0N9wNn9As4z2bON3FF3je+1Nu+2+/8Zj50HpMTpzYpMp2R93cYmACCmhaYoKydxylbyGzJR8y9likZRCrkD38lRoS40aBBvv/6iRKzQuo9EGYVcel36cMNg00VmYNBy3pa1rwg3gA4l3cy6kjayZja1WGPkVhrVKsrJMdbpl0ApLjXKuD8rw1n1XLCwCUEL5eLPljTZaAveqdOFQOiZnZEGI27qIiZe7I1fN8tcz6anS/gTM7xRKE++5egEvRWlTQQTJeApnPSiUPA+8ZykNdelQsOQh5SrGoyI4A5pq****==` of the digest `ZOyIygCyaOW6GjVnihtTFtIS9PNmskdyMlNKiuyjfzw=`.
//
// @param request - AsymmetricVerifyRequest
//
// @return AsymmetricVerifyResponse
func (client *Client) AsymmetricVerify(request *AsymmetricVerifyRequest) (_result *AsymmetricVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AsymmetricVerifyResponse{}
	_body, _err := client.AsymmetricVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// If the deletion task of a CMK is canceled, the CMK returns to the Enabled state.
//
// @param request - CancelKeyDeletionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelKeyDeletionResponse
func (client *Client) CancelKeyDeletionWithOptions(request *CancelKeyDeletionRequest, runtime *util.RuntimeOptions) (_result *CancelKeyDeletionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelKeyDeletion"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CancelKeyDeletionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CancelKeyDeletionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// If the deletion task of a CMK is canceled, the CMK returns to the Enabled state.
//
// @param request - CancelKeyDeletionRequest
//
// @return CancelKeyDeletionResponse
func (client *Client) CancelKeyDeletion(request *CancelKeyDeletionRequest) (_result *CancelKeyDeletionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelKeyDeletionResponse{}
	_body, _err := client.CancelKeyDeletionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Decrypts data by using a specific certificate.
//
// Description:
//
// Limit: The encryption algorithm in the request parameters must match the key type.
//
// The following table describes the mapping between encryption algorithms and key types.
//
// | Algorithm | Key Spec |
//
// | --------- | -------- |
//
// | RSAES_OAEP_SHA_1 | RSA_2048 |
//
// | RSAES_OAEP_SHA_256 | RSA_2048 |
//
// | SM2PKE | EC_SM2 |
//
// In this example, the certificate whose ID is `12345678-1234-1234-1234-12345678****` and the encryption algorithm `RSAES_OAEP_SHA_256` are used to decrypt the data `ZOyIygCyaOW6Gj****MlNKiuyjfzw=`.
//
// @param request - CertificatePrivateKeyDecryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CertificatePrivateKeyDecryptResponse
func (client *Client) CertificatePrivateKeyDecryptWithOptions(request *CertificatePrivateKeyDecryptRequest, runtime *util.RuntimeOptions) (_result *CertificatePrivateKeyDecryptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		query["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.CertificateId)) {
		query["CertificateId"] = request.CertificateId
	}

	if !tea.BoolValue(util.IsUnset(request.CiphertextBlob)) {
		query["CiphertextBlob"] = request.CiphertextBlob
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CertificatePrivateKeyDecrypt"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CertificatePrivateKeyDecryptResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CertificatePrivateKeyDecryptResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Decrypts data by using a specific certificate.
//
// Description:
//
// Limit: The encryption algorithm in the request parameters must match the key type.
//
// The following table describes the mapping between encryption algorithms and key types.
//
// | Algorithm | Key Spec |
//
// | --------- | -------- |
//
// | RSAES_OAEP_SHA_1 | RSA_2048 |
//
// | RSAES_OAEP_SHA_256 | RSA_2048 |
//
// | SM2PKE | EC_SM2 |
//
// In this example, the certificate whose ID is `12345678-1234-1234-1234-12345678****` and the encryption algorithm `RSAES_OAEP_SHA_256` are used to decrypt the data `ZOyIygCyaOW6Gj****MlNKiuyjfzw=`.
//
// @param request - CertificatePrivateKeyDecryptRequest
//
// @return CertificatePrivateKeyDecryptResponse
func (client *Client) CertificatePrivateKeyDecrypt(request *CertificatePrivateKeyDecryptRequest) (_result *CertificatePrivateKeyDecryptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CertificatePrivateKeyDecryptResponse{}
	_body, _err := client.CertificatePrivateKeyDecryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a signature by using a specified certificate.
//
// Description:
//
// The signature algorithm in the request parameters must match the key type. The following table describes the mapping between signature algorithms and key types.
//
// | Algorithm | Key Spec |
//
// | --------- | -------- |
//
// | RSA_PKCS1_SHA_256 | RSA_2048 |
//
// | RSA_PSS_SHA_256 | RSA_2048 |
//
// | ECDSA_SHA_256 | EC_P256 |
//
// | SM2DSA | EC_SM2 |
//
// In this example, the certificate whose ID is `12345678-1234-1234-1234-12345678****` and the signature algorithm `ECDSA_SHA_256` are used to generate a signature for the raw data `VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4=`.
//
// @param request - CertificatePrivateKeySignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CertificatePrivateKeySignResponse
func (client *Client) CertificatePrivateKeySignWithOptions(request *CertificatePrivateKeySignRequest, runtime *util.RuntimeOptions) (_result *CertificatePrivateKeySignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		query["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.CertificateId)) {
		query["CertificateId"] = request.CertificateId
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["Message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		query["MessageType"] = request.MessageType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CertificatePrivateKeySign"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CertificatePrivateKeySignResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CertificatePrivateKeySignResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Generates a signature by using a specified certificate.
//
// Description:
//
// The signature algorithm in the request parameters must match the key type. The following table describes the mapping between signature algorithms and key types.
//
// | Algorithm | Key Spec |
//
// | --------- | -------- |
//
// | RSA_PKCS1_SHA_256 | RSA_2048 |
//
// | RSA_PSS_SHA_256 | RSA_2048 |
//
// | ECDSA_SHA_256 | EC_P256 |
//
// | SM2DSA | EC_SM2 |
//
// In this example, the certificate whose ID is `12345678-1234-1234-1234-12345678****` and the signature algorithm `ECDSA_SHA_256` are used to generate a signature for the raw data `VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4=`.
//
// @param request - CertificatePrivateKeySignRequest
//
// @return CertificatePrivateKeySignResponse
func (client *Client) CertificatePrivateKeySign(request *CertificatePrivateKeySignRequest) (_result *CertificatePrivateKeySignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CertificatePrivateKeySignResponse{}
	_body, _err := client.CertificatePrivateKeySignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Encrypts data by using a specific certificate.
//
// Description:
//
// Limit: The encryption algorithm in the request parameters must match the key type.
//
// The following table describes the mapping between encryption algorithms and key types.
//
// | Algorithm | Key Spec |
//
// | --------- | -------- |
//
// | RSAES_OAEP_SHA_1 | RSA_2048 |
//
// | RSAES_OAEP_SHA_256 | RSA_2048 |
//
// | SM2PKE | EC_SM2 |
//
// In this example, the certificate whose ID is `12345678-1234-1234-1234-12345678****` and the encryption algorithm `RSAES_OAEP_SHA_256` are used to encrypt the data `VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4=`.
//
// @param request - CertificatePublicKeyEncryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CertificatePublicKeyEncryptResponse
func (client *Client) CertificatePublicKeyEncryptWithOptions(request *CertificatePublicKeyEncryptRequest, runtime *util.RuntimeOptions) (_result *CertificatePublicKeyEncryptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		query["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.CertificateId)) {
		query["CertificateId"] = request.CertificateId
	}

	if !tea.BoolValue(util.IsUnset(request.Plaintext)) {
		query["Plaintext"] = request.Plaintext
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CertificatePublicKeyEncrypt"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CertificatePublicKeyEncryptResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CertificatePublicKeyEncryptResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Encrypts data by using a specific certificate.
//
// Description:
//
// Limit: The encryption algorithm in the request parameters must match the key type.
//
// The following table describes the mapping between encryption algorithms and key types.
//
// | Algorithm | Key Spec |
//
// | --------- | -------- |
//
// | RSAES_OAEP_SHA_1 | RSA_2048 |
//
// | RSAES_OAEP_SHA_256 | RSA_2048 |
//
// | SM2PKE | EC_SM2 |
//
// In this example, the certificate whose ID is `12345678-1234-1234-1234-12345678****` and the encryption algorithm `RSAES_OAEP_SHA_256` are used to encrypt the data `VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4=`.
//
// @param request - CertificatePublicKeyEncryptRequest
//
// @return CertificatePublicKeyEncryptResponse
func (client *Client) CertificatePublicKeyEncrypt(request *CertificatePublicKeyEncryptRequest) (_result *CertificatePublicKeyEncryptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CertificatePublicKeyEncryptResponse{}
	_body, _err := client.CertificatePublicKeyEncryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies a digital signature by using a specified certificate.
//
// Description:
//
// The signature algorithm in the request parameters must match the key type. The following table describes the mapping between signature algorithms and key types.
//
// | Algorithm | Key Spec |
//
// | --------- | -------- |
//
// | RSA_PKCS1_SHA_256 | RSA_2048 |
//
// | RSA_PSS_SHA_256 | RSA_2048 |
//
// | ECDSA_SHA_256 | EC_P256 |
//
// | SM2DSA | EC_SM2 |
//
// In this example, the certificate whose ID is `12345678-1234-1234-1234-12345678****` and the signature algorithm `ECDSA_SHA_256` are used to verify the digital signature `ZOyIygCyaOW6Gj****MlNKiuyjfzw=` of the raw data `VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4=`.
//
// @param request - CertificatePublicKeyVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CertificatePublicKeyVerifyResponse
func (client *Client) CertificatePublicKeyVerifyWithOptions(request *CertificatePublicKeyVerifyRequest, runtime *util.RuntimeOptions) (_result *CertificatePublicKeyVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		query["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.CertificateId)) {
		query["CertificateId"] = request.CertificateId
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["Message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		query["MessageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.SignatureValue)) {
		query["SignatureValue"] = request.SignatureValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CertificatePublicKeyVerify"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CertificatePublicKeyVerifyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CertificatePublicKeyVerifyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Verifies a digital signature by using a specified certificate.
//
// Description:
//
// The signature algorithm in the request parameters must match the key type. The following table describes the mapping between signature algorithms and key types.
//
// | Algorithm | Key Spec |
//
// | --------- | -------- |
//
// | RSA_PKCS1_SHA_256 | RSA_2048 |
//
// | RSA_PSS_SHA_256 | RSA_2048 |
//
// | ECDSA_SHA_256 | EC_P256 |
//
// | SM2DSA | EC_SM2 |
//
// In this example, the certificate whose ID is `12345678-1234-1234-1234-12345678****` and the signature algorithm `ECDSA_SHA_256` are used to verify the digital signature `ZOyIygCyaOW6Gj****MlNKiuyjfzw=` of the raw data `VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4=`.
//
// @param request - CertificatePublicKeyVerifyRequest
//
// @return CertificatePublicKeyVerifyResponse
func (client *Client) CertificatePublicKeyVerify(request *CertificatePublicKeyVerifyRequest) (_result *CertificatePublicKeyVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CertificatePublicKeyVerifyResponse{}
	_body, _err := client.CertificatePublicKeyVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a Key Management Service (KMS) instance.
//
// Description:
//
// ### [](#)Limits
//
// You can enable only instances of the software key management type. You cannot enable instances of the hardware key management type.
//
// @param request - ConnectKmsInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConnectKmsInstanceResponse
func (client *Client) ConnectKmsInstanceWithOptions(request *ConnectKmsInstanceRequest, runtime *util.RuntimeOptions) (_result *ConnectKmsInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KMProvider)) {
		query["KMProvider"] = request.KMProvider
	}

	if !tea.BoolValue(util.IsUnset(request.KmsInstanceId)) {
		query["KmsInstanceId"] = request.KmsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchIds)) {
		query["VSwitchIds"] = request.VSwitchIds
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneIds)) {
		query["ZoneIds"] = request.ZoneIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConnectKmsInstance"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ConnectKmsInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ConnectKmsInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Enables a Key Management Service (KMS) instance.
//
// Description:
//
// ### [](#)Limits
//
// You can enable only instances of the software key management type. You cannot enable instances of the hardware key management type.
//
// @param request - ConnectKmsInstanceRequest
//
// @return ConnectKmsInstanceResponse
func (client *Client) ConnectKmsInstance(request *ConnectKmsInstanceRequest) (_result *ConnectKmsInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConnectKmsInstanceResponse{}
	_body, _err := client.ConnectKmsInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
//   Each alias can be bound to only one CMK at a time.
//
// 	- The aliases of CMKs in the same region must be unique.
//
// In this topic, an alias named `alias/example` is created for a CMK named `7906979c-8e06-46a2-be2d-68e3ccbc****`.
//
// @param request - CreateAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAliasResponse
func (client *Client) CreateAliasWithOptions(request *CreateAliasRequest, runtime *util.RuntimeOptions) (_result *CreateAliasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliasName)) {
		query["AliasName"] = request.AliasName
	}

	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAlias"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateAliasResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateAliasResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
//   Each alias can be bound to only one CMK at a time.
//
// 	- The aliases of CMKs in the same region must be unique.
//
// In this topic, an alias named `alias/example` is created for a CMK named `7906979c-8e06-46a2-be2d-68e3ccbc****`.
//
// @param request - CreateAliasRequest
//
// @return CreateAliasResponse
func (client *Client) CreateAlias(request *CreateAliasRequest) (_result *CreateAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAliasResponse{}
	_body, _err := client.CreateAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an application access point (AAP)
//
// Description:
//
// To perform cryptographic operations and retrieve secret values, self-managed applications must use a client key to access a Key Management Service (KMS) instance. The following process shows how to create a client key-based AAP:
//
// 1.Create a network access rule: You can configure the private IP addresses or private CIDR blocks that are allowed to access KMS. For more information, see [CreateNetworkRule](https://help.aliyun.com/document_detail/2539407.html).
//
// 2.Create a permission policy: You can configure the keys and secrets that are allowed to access and bind network access rules to the keys and secrets. For more information, see [CreatePolicy](https://help.aliyun.com/document_detail/2539454.html).
//
// 3.Create an AAP: You can configure an authentication method and bind a permission policy to an AAP. This topic describes how to create an AAP.
//
// 4.Create a client key: You can configure the encryption password and validity period of a client key and bind the client key to an AAP. For more information, see [CreateClientKey](https://help.aliyun.com/document_detail/2539509.html).
//
// @param request - CreateApplicationAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationAccessPointResponse
func (client *Client) CreateApplicationAccessPointWithOptions(request *CreateApplicationAccessPointRequest, runtime *util.RuntimeOptions) (_result *CreateApplicationAccessPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthenticationMethod)) {
		query["AuthenticationMethod"] = request.AuthenticationMethod
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Policies)) {
		query["Policies"] = request.Policies
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApplicationAccessPoint"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateApplicationAccessPointResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateApplicationAccessPointResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates an application access point (AAP)
//
// Description:
//
// To perform cryptographic operations and retrieve secret values, self-managed applications must use a client key to access a Key Management Service (KMS) instance. The following process shows how to create a client key-based AAP:
//
// 1.Create a network access rule: You can configure the private IP addresses or private CIDR blocks that are allowed to access KMS. For more information, see [CreateNetworkRule](https://help.aliyun.com/document_detail/2539407.html).
//
// 2.Create a permission policy: You can configure the keys and secrets that are allowed to access and bind network access rules to the keys and secrets. For more information, see [CreatePolicy](https://help.aliyun.com/document_detail/2539454.html).
//
// 3.Create an AAP: You can configure an authentication method and bind a permission policy to an AAP. This topic describes how to create an AAP.
//
// 4.Create a client key: You can configure the encryption password and validity period of a client key and bind the client key to an AAP. For more information, see [CreateClientKey](https://help.aliyun.com/document_detail/2539509.html).
//
// @param request - CreateApplicationAccessPointRequest
//
// @return CreateApplicationAccessPointResponse
func (client *Client) CreateApplicationAccessPoint(request *CreateApplicationAccessPointRequest) (_result *CreateApplicationAccessPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateApplicationAccessPointResponse{}
	_body, _err := client.CreateApplicationAccessPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// To create a certificate, you must specify the type of the asymmetric key. Certificates Manager generates a private key and returns a certificate signing request (CSR). Submit the CSR in the Privacy Enhanced Mail (PEM) format to a certificate authority (CA) to obtain the formal certificate and certificate chain. Then, call the [UploadCertificate](https://help.aliyun.com/document_detail/212136.html) operation to import the certificate into Certificates Manager.
//
// In this example, a certificate is created and the CSR is obtained.
//
// @param tmpReq - CreateCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCertificateResponse
func (client *Client) CreateCertificateWithOptions(tmpReq *CreateCertificateRequest, runtime *util.RuntimeOptions) (_result *CreateCertificateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateCertificateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SubjectAlternativeNames)) {
		request.SubjectAlternativeNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubjectAlternativeNames, tea.String("SubjectAlternativeNames"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExportablePrivateKey)) {
		query["ExportablePrivateKey"] = request.ExportablePrivateKey
	}

	if !tea.BoolValue(util.IsUnset(request.KeySpec)) {
		query["KeySpec"] = request.KeySpec
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		query["Subject"] = request.Subject
	}

	if !tea.BoolValue(util.IsUnset(request.SubjectAlternativeNamesShrink)) {
		query["SubjectAlternativeNames"] = request.SubjectAlternativeNamesShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCertificate"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateCertificateResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateCertificateResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// To create a certificate, you must specify the type of the asymmetric key. Certificates Manager generates a private key and returns a certificate signing request (CSR). Submit the CSR in the Privacy Enhanced Mail (PEM) format to a certificate authority (CA) to obtain the formal certificate and certificate chain. Then, call the [UploadCertificate](https://help.aliyun.com/document_detail/212136.html) operation to import the certificate into Certificates Manager.
//
// In this example, a certificate is created and the CSR is obtained.
//
// @param request - CreateCertificateRequest
//
// @return CreateCertificateResponse
func (client *Client) CreateCertificate(request *CreateCertificateRequest) (_result *CreateCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCertificateResponse{}
	_body, _err := client.CreateCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a client key.
//
// Description:
//
// To perform cryptographic operations and retrieve secret values, self-managed applications must use a client key to access a Key Management Service (KMS) instance. The following process shows how to create a client key-based application access point (AAP):
//
// 1.Create an access control rule: You can configure the private IP addresses or private CIDR blocks that are allowed to access a KMS instance. For more information, see [CreateNetworkRule](https://help.aliyun.com/document_detail/2539407.html).
//
// 2.Create a permission policy: You can configure the keys and secrets that are allowed to access and bind access control rules to the keys and secrets. For more information, see [CreatePolicy](https://help.aliyun.com/document_detail/2539454.html).
//
// 3.Create an AAP: You can configure an authentication method and bind a permission policy to an AAP. For more information, see [CreateApplicationAccessPoint](https://help.aliyun.com/document_detail/2539467.html).
//
// 4.Create a client key: You can configure the encryption password and validity period of a client key and bind the client key to an AAP.
//
// ### Precautions
//
// A client key has a validity period. After a client key expires, applications into which the client key is integrated cannot access the required KMS instance. You must replace the client key before the client key expires. We recommend that you delete the expired client key in KMS after the new client key is used.
//
// @param request - CreateClientKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClientKeyResponse
func (client *Client) CreateClientKeyWithOptions(request *CreateClientKeyRequest, runtime *util.RuntimeOptions) (_result *CreateClientKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AapName)) {
		query["AapName"] = request.AapName
	}

	if !tea.BoolValue(util.IsUnset(request.NotAfter)) {
		query["NotAfter"] = request.NotAfter
	}

	if !tea.BoolValue(util.IsUnset(request.NotBefore)) {
		query["NotBefore"] = request.NotBefore
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateClientKey"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateClientKeyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateClientKeyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a client key.
//
// Description:
//
// To perform cryptographic operations and retrieve secret values, self-managed applications must use a client key to access a Key Management Service (KMS) instance. The following process shows how to create a client key-based application access point (AAP):
//
// 1.Create an access control rule: You can configure the private IP addresses or private CIDR blocks that are allowed to access a KMS instance. For more information, see [CreateNetworkRule](https://help.aliyun.com/document_detail/2539407.html).
//
// 2.Create a permission policy: You can configure the keys and secrets that are allowed to access and bind access control rules to the keys and secrets. For more information, see [CreatePolicy](https://help.aliyun.com/document_detail/2539454.html).
//
// 3.Create an AAP: You can configure an authentication method and bind a permission policy to an AAP. For more information, see [CreateApplicationAccessPoint](https://help.aliyun.com/document_detail/2539467.html).
//
// 4.Create a client key: You can configure the encryption password and validity period of a client key and bind the client key to an AAP.
//
// ### Precautions
//
// A client key has a validity period. After a client key expires, applications into which the client key is integrated cannot access the required KMS instance. You must replace the client key before the client key expires. We recommend that you delete the expired client key in KMS after the new client key is used.
//
// @param request - CreateClientKeyRequest
//
// @return CreateClientKeyResponse
func (client *Client) CreateClientKey(request *CreateClientKeyRequest) (_result *CreateClientKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClientKeyResponse{}
	_body, _err := client.CreateClientKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a customer master key (CMK).
//
// Description:
//
// KMS supports common symmetric keys and asymmetric keys. For more information, see [Key types and specifications](https://help.aliyun.com/document_detail/480161.html).
//
// @param request - CreateKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKeyResponse
func (client *Client) CreateKeyWithOptions(request *CreateKeyRequest, runtime *util.RuntimeOptions) (_result *CreateKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DKMSInstanceId)) {
		query["DKMSInstanceId"] = request.DKMSInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EnableAutomaticRotation)) {
		query["EnableAutomaticRotation"] = request.EnableAutomaticRotation
	}

	if !tea.BoolValue(util.IsUnset(request.KeySpec)) {
		query["KeySpec"] = request.KeySpec
	}

	if !tea.BoolValue(util.IsUnset(request.KeyUsage)) {
		query["KeyUsage"] = request.KeyUsage
	}

	if !tea.BoolValue(util.IsUnset(request.Origin)) {
		query["Origin"] = request.Origin
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		query["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.ProtectionLevel)) {
		query["ProtectionLevel"] = request.ProtectionLevel
	}

	if !tea.BoolValue(util.IsUnset(request.RotationInterval)) {
		query["RotationInterval"] = request.RotationInterval
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateKey"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateKeyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateKeyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a customer master key (CMK).
//
// Description:
//
// KMS supports common symmetric keys and asymmetric keys. For more information, see [Key types and specifications](https://help.aliyun.com/document_detail/480161.html).
//
// @param request - CreateKeyRequest
//
// @return CreateKeyResponse
func (client *Client) CreateKey(request *CreateKeyRequest) (_result *CreateKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateKeyResponse{}
	_body, _err := client.CreateKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 为密钥创建新的密钥版本。
//
// Description:
//
//   You can create a version only for an asymmetric CMK that is in the Enabled state. You can call the [CreateKey](https://help.aliyun.com/document_detail/28947.html) operation to create an asymmetric CMK and the [DescribeKey](https://help.aliyun.com/document_detail/28952.html) operation to query the status of the CMK. The status is specified by the KeyState parameter.
//
// 	- The minimum interval for creating a version of the same CMK is seven days. You can call the [DescribeKey](https://help.aliyun.com/document_detail/28952.html) operation to query the time when the last version of a CMK was created. The time is specified by the LastRotationDate parameter.
//
// 	- If a CMK is in a private key store, you cannot create a version for the CMK.
//
// 	- You can create a maximum of 50 versions for a CMK in the same region.
//
// You can create a version for the CMK whose ID is `0b30658a-ed1a-4922-b8f7-a673ca9c****` by using the parameter settings provided in this topic.
//
// @param request - CreateKeyVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKeyVersionResponse
func (client *Client) CreateKeyVersionWithOptions(request *CreateKeyVersionRequest, runtime *util.RuntimeOptions) (_result *CreateKeyVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateKeyVersion"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateKeyVersionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateKeyVersionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 为密钥创建新的密钥版本。
//
// Description:
//
//   You can create a version only for an asymmetric CMK that is in the Enabled state. You can call the [CreateKey](https://help.aliyun.com/document_detail/28947.html) operation to create an asymmetric CMK and the [DescribeKey](https://help.aliyun.com/document_detail/28952.html) operation to query the status of the CMK. The status is specified by the KeyState parameter.
//
// 	- The minimum interval for creating a version of the same CMK is seven days. You can call the [DescribeKey](https://help.aliyun.com/document_detail/28952.html) operation to query the time when the last version of a CMK was created. The time is specified by the LastRotationDate parameter.
//
// 	- If a CMK is in a private key store, you cannot create a version for the CMK.
//
// 	- You can create a maximum of 50 versions for a CMK in the same region.
//
// You can create a version for the CMK whose ID is `0b30658a-ed1a-4922-b8f7-a673ca9c****` by using the parameter settings provided in this topic.
//
// @param request - CreateKeyVersionRequest
//
// @return CreateKeyVersionResponse
func (client *Client) CreateKeyVersion(request *CreateKeyVersionRequest) (_result *CreateKeyVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateKeyVersionResponse{}
	_body, _err := client.CreateKeyVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an access control rule to configure the private IP addresses or CIDR blocks that are allowed to access a Key Management Service (KMS) instance.
//
// Description:
//
// To perform cryptographic operations and retrieve secret values, self-managed applications must use a client key to access a KMS instance. The following process shows how to create a client key-based application access point (AAP):
//
// 1.Create an access control rule: You can configure the private IP addresses or private CIDR blocks that are allowed to access a KMS instance.
//
// 2.Create a permission policy: You can configure the keys and secrets that are allowed to access and bind access control rules to the keys and secrets. For more information, see [CreatePolicy](https://help.aliyun.com/document_detail/2539454.html).
//
// 3.Create an AAP: You can configure an authentication method and bind a permission policy to an AAP. For more information, see [CreateApplicationAccessPoint](https://help.aliyun.com/document_detail/2539467.html).
//
// 4.Create a client key: You can configure the encryption password and validity period of a client key and bind the client key to an AAP. For more information, see [CreateClientKey](https://help.aliyun.com/document_detail/2539509.html).
//
// @param request - CreateNetworkRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetworkRuleResponse
func (client *Client) CreateNetworkRuleWithOptions(request *CreateNetworkRuleRequest, runtime *util.RuntimeOptions) (_result *CreateNetworkRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SourcePrivateIp)) {
		query["SourcePrivateIp"] = request.SourcePrivateIp
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNetworkRule"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateNetworkRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateNetworkRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates an access control rule to configure the private IP addresses or CIDR blocks that are allowed to access a Key Management Service (KMS) instance.
//
// Description:
//
// To perform cryptographic operations and retrieve secret values, self-managed applications must use a client key to access a KMS instance. The following process shows how to create a client key-based application access point (AAP):
//
// 1.Create an access control rule: You can configure the private IP addresses or private CIDR blocks that are allowed to access a KMS instance.
//
// 2.Create a permission policy: You can configure the keys and secrets that are allowed to access and bind access control rules to the keys and secrets. For more information, see [CreatePolicy](https://help.aliyun.com/document_detail/2539454.html).
//
// 3.Create an AAP: You can configure an authentication method and bind a permission policy to an AAP. For more information, see [CreateApplicationAccessPoint](https://help.aliyun.com/document_detail/2539467.html).
//
// 4.Create a client key: You can configure the encryption password and validity period of a client key and bind the client key to an AAP. For more information, see [CreateClientKey](https://help.aliyun.com/document_detail/2539509.html).
//
// @param request - CreateNetworkRuleRequest
//
// @return CreateNetworkRuleResponse
func (client *Client) CreateNetworkRule(request *CreateNetworkRuleRequest) (_result *CreateNetworkRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNetworkRuleResponse{}
	_body, _err := client.CreateNetworkRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a permission policy to configure the keys and secrets that are allowed to access.
//
// Description:
//
// To perform cryptographic operations and retrieve secret values, self-managed applications must use a client key to access a Key Management Service (KMS) instance. The following process shows how to create a client key-based application access point (AAP):
//
// 1.Create an access control rule: You can configure the private IP addresses or private CIDR blocks that are allowed to access a KMS instance. For more information, see [CreateNetworkRule](https://help.aliyun.com/document_detail/2539407.html).
//
// 2.Create a permission policy: You can configure the keys and secrets that are allowed to access and bind access control rules to the keys and secrets.
//
// 3.Create an AAP: You can configure an authentication method and bind a permission policy to an AAP. For more information, see [CreateApplicationAccessPoint](https://help.aliyun.com/document_detail/2539467.html).
//
// 4.Create a client key: You can configure the encryption password and validity period of a client key and bind the client key to an AAP. For more information, see [CreateClientKey](https://help.aliyun.com/document_detail/2539509.html).
//
// @param request - CreatePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicyWithOptions(request *CreatePolicyRequest, runtime *util.RuntimeOptions) (_result *CreatePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessControlRules)) {
		query["AccessControlRules"] = request.AccessControlRules
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.KmsInstance)) {
		query["KmsInstance"] = request.KmsInstance
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Permissions)) {
		query["Permissions"] = request.Permissions
	}

	if !tea.BoolValue(util.IsUnset(request.Resources)) {
		query["Resources"] = request.Resources
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePolicy"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreatePolicyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreatePolicyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a permission policy to configure the keys and secrets that are allowed to access.
//
// Description:
//
// To perform cryptographic operations and retrieve secret values, self-managed applications must use a client key to access a Key Management Service (KMS) instance. The following process shows how to create a client key-based application access point (AAP):
//
// 1.Create an access control rule: You can configure the private IP addresses or private CIDR blocks that are allowed to access a KMS instance. For more information, see [CreateNetworkRule](https://help.aliyun.com/document_detail/2539407.html).
//
// 2.Create a permission policy: You can configure the keys and secrets that are allowed to access and bind access control rules to the keys and secrets.
//
// 3.Create an AAP: You can configure an authentication method and bind a permission policy to an AAP. For more information, see [CreateApplicationAccessPoint](https://help.aliyun.com/document_detail/2539467.html).
//
// 4.Create a client key: You can configure the encryption password and validity period of a client key and bind the client key to an AAP. For more information, see [CreateClientKey](https://help.aliyun.com/document_detail/2539509.html).
//
// @param request - CreatePolicyRequest
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicy(request *CreatePolicyRequest) (_result *CreatePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePolicyResponse{}
	_body, _err := client.CreatePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建凭据并存入凭据的初始版本。
//
// Description:
//
// The name of the secret.
//
// The value must be 1 to 64 characters in length and can contain letters, digits, underscores (_), forward slashes (/), plus signs (+), equal signs (=), periods (.), hyphens (-), and at signs (@). The following list describes the name requirements for different types of secrets:
//
// 	- If the SecretType parameter is set to Generic or Rds, the name cannot start with `acs/`.
//
// 	- If the SecretType parameter is set to RAMCredentials, set the SecretName parameter to `$Auto`. In this case, KMS automatically generates a secret name that starts with `acs/ram/user/`. The name includes the display name of RAM user.
//
// 	- If the SecretType parameter is set to ECS, the name must start with `acs/ecs/`.
//
// @param tmpReq - CreateSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecretResponse
func (client *Client) CreateSecretWithOptions(tmpReq *CreateSecretRequest, runtime *util.RuntimeOptions) (_result *CreateSecretResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateSecretShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ExtendedConfig)) {
		request.ExtendedConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtendedConfig, tea.String("ExtendedConfig"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DKMSInstanceId)) {
		query["DKMSInstanceId"] = request.DKMSInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EnableAutomaticRotation)) {
		query["EnableAutomaticRotation"] = request.EnableAutomaticRotation
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptionKeyId)) {
		query["EncryptionKeyId"] = request.EncryptionKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendedConfigShrink)) {
		query["ExtendedConfig"] = request.ExtendedConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		query["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.RotationInterval)) {
		query["RotationInterval"] = request.RotationInterval
	}

	if !tea.BoolValue(util.IsUnset(request.SecretData)) {
		query["SecretData"] = request.SecretData
	}

	if !tea.BoolValue(util.IsUnset(request.SecretDataType)) {
		query["SecretDataType"] = request.SecretDataType
	}

	if !tea.BoolValue(util.IsUnset(request.SecretName)) {
		query["SecretName"] = request.SecretName
	}

	if !tea.BoolValue(util.IsUnset(request.SecretType)) {
		query["SecretType"] = request.SecretType
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSecret"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateSecretResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateSecretResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建凭据并存入凭据的初始版本。
//
// Description:
//
// The name of the secret.
//
// The value must be 1 to 64 characters in length and can contain letters, digits, underscores (_), forward slashes (/), plus signs (+), equal signs (=), periods (.), hyphens (-), and at signs (@). The following list describes the name requirements for different types of secrets:
//
// 	- If the SecretType parameter is set to Generic or Rds, the name cannot start with `acs/`.
//
// 	- If the SecretType parameter is set to RAMCredentials, set the SecretName parameter to `$Auto`. In this case, KMS automatically generates a secret name that starts with `acs/ram/user/`. The name includes the display name of RAM user.
//
// 	- If the SecretType parameter is set to ECS, the name must start with `acs/ecs/`.
//
// @param request - CreateSecretRequest
//
// @return CreateSecretResponse
func (client *Client) CreateSecret(request *CreateSecretRequest) (_result *CreateSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSecretResponse{}
	_body, _err := client.CreateSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用Decrypt接口解密CiphertextBlob中的密文。
//
// @param tmpReq - DecryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DecryptResponse
func (client *Client) DecryptWithOptions(tmpReq *DecryptRequest, runtime *util.RuntimeOptions) (_result *DecryptResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DecryptShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EncryptionContext)) {
		request.EncryptionContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EncryptionContext, tea.String("EncryptionContext"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CiphertextBlob)) {
		query["CiphertextBlob"] = request.CiphertextBlob
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptionContextShrink)) {
		query["EncryptionContext"] = request.EncryptionContextShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Decrypt"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DecryptResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DecryptResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 调用Decrypt接口解密CiphertextBlob中的密文。
//
// @param request - DecryptRequest
//
// @return DecryptResponse
func (client *Client) Decrypt(request *DecryptRequest) (_result *DecryptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DecryptResponse{}
	_body, _err := client.DecryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAliasResponse
func (client *Client) DeleteAliasWithOptions(request *DeleteAliasRequest, runtime *util.RuntimeOptions) (_result *DeleteAliasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliasName)) {
		query["AliasName"] = request.AliasName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAlias"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteAliasResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteAliasResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DeleteAliasRequest
//
// @return DeleteAliasResponse
func (client *Client) DeleteAlias(request *DeleteAliasRequest) (_result *DeleteAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAliasResponse{}
	_body, _err := client.DeleteAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an application access point (AAP).
//
// Description:
//
// Before you delete an AAP, make sure that the AAP is no longer in use. If you delete an AAP that is in use, applications that use the AAP cannot access Key Management Service (KMS). Exercise caution when you delete an AAP.
//
// @param request - DeleteApplicationAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationAccessPointResponse
func (client *Client) DeleteApplicationAccessPointWithOptions(request *DeleteApplicationAccessPointRequest, runtime *util.RuntimeOptions) (_result *DeleteApplicationAccessPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApplicationAccessPoint"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteApplicationAccessPointResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteApplicationAccessPointResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes an application access point (AAP).
//
// Description:
//
// Before you delete an AAP, make sure that the AAP is no longer in use. If you delete an AAP that is in use, applications that use the AAP cannot access Key Management Service (KMS). Exercise caution when you delete an AAP.
//
// @param request - DeleteApplicationAccessPointRequest
//
// @return DeleteApplicationAccessPointResponse
func (client *Client) DeleteApplicationAccessPoint(request *DeleteApplicationAccessPointRequest) (_result *DeleteApplicationAccessPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteApplicationAccessPointResponse{}
	_body, _err := client.DeleteApplicationAccessPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// After the certificate and its private key and certificate chain are deleted, they cannot be restored. Proceed with caution.
//
// In this example, the certificate whose ID is `9a28de48-8d8b-484d-a766-dec4****` and its private key and certificate chain are deleted.
//
// @param request - DeleteCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCertificateResponse
func (client *Client) DeleteCertificateWithOptions(request *DeleteCertificateRequest, runtime *util.RuntimeOptions) (_result *DeleteCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertificateId)) {
		query["CertificateId"] = request.CertificateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCertificate"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteCertificateResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteCertificateResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// After the certificate and its private key and certificate chain are deleted, they cannot be restored. Proceed with caution.
//
// In this example, the certificate whose ID is `9a28de48-8d8b-484d-a766-dec4****` and its private key and certificate chain are deleted.
//
// @param request - DeleteCertificateRequest
//
// @return DeleteCertificateResponse
func (client *Client) DeleteCertificate(request *DeleteCertificateRequest) (_result *DeleteCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCertificateResponse{}
	_body, _err := client.DeleteCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// Before you delete a client key, make sure that the client key is no longer in use. If you delete a client key that is in use, applications that use the client key cannot access Key Management Service (KMS). Exercise caution when you delete a client key.
//
// @param request - DeleteClientKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClientKeyResponse
func (client *Client) DeleteClientKeyWithOptions(request *DeleteClientKeyRequest, runtime *util.RuntimeOptions) (_result *DeleteClientKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientKeyId)) {
		query["ClientKeyId"] = request.ClientKeyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteClientKey"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteClientKeyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteClientKeyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// Before you delete a client key, make sure that the client key is no longer in use. If you delete a client key that is in use, applications that use the client key cannot access Key Management Service (KMS). Exercise caution when you delete a client key.
//
// @param request - DeleteClientKeyRequest
//
// @return DeleteClientKeyResponse
func (client *Client) DeleteClientKey(request *DeleteClientKeyRequest) (_result *DeleteClientKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClientKeyResponse{}
	_body, _err := client.DeleteClientKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// This operation does not delete the CMK that is created by using the key material.
//
// If the CMK is in the PendingDeletion state, the state of the CMK and the scheduled deletion time do not change after you call this operation. If the CMK is not in the PendingDeletion state, the state of the CMK changes to PendingImport after you call this operation.
//
// After you delete the key material, you can upload only the same key material into the CMK.
//
// @param request - DeleteKeyMaterialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteKeyMaterialResponse
func (client *Client) DeleteKeyMaterialWithOptions(request *DeleteKeyMaterialRequest, runtime *util.RuntimeOptions) (_result *DeleteKeyMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteKeyMaterial"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteKeyMaterialResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteKeyMaterialResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// This operation does not delete the CMK that is created by using the key material.
//
// If the CMK is in the PendingDeletion state, the state of the CMK and the scheduled deletion time do not change after you call this operation. If the CMK is not in the PendingDeletion state, the state of the CMK changes to PendingImport after you call this operation.
//
// After you delete the key material, you can upload only the same key material into the CMK.
//
// @param request - DeleteKeyMaterialRequest
//
// @return DeleteKeyMaterialResponse
func (client *Client) DeleteKeyMaterial(request *DeleteKeyMaterialRequest) (_result *DeleteKeyMaterialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteKeyMaterialResponse{}
	_body, _err := client.DeleteKeyMaterialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a network access rule.
//
// Description:
//
// Before you delete a network access rule, make sure that the network access rule is not bound to permission policies. Otherwise, related applications cannot access Key Management Service (KMS).
//
// @param request - DeleteNetworkRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkRuleResponse
func (client *Client) DeleteNetworkRuleWithOptions(request *DeleteNetworkRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteNetworkRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNetworkRule"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteNetworkRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteNetworkRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes a network access rule.
//
// Description:
//
// Before you delete a network access rule, make sure that the network access rule is not bound to permission policies. Otherwise, related applications cannot access Key Management Service (KMS).
//
// @param request - DeleteNetworkRuleRequest
//
// @return DeleteNetworkRuleResponse
func (client *Client) DeleteNetworkRule(request *DeleteNetworkRuleRequest) (_result *DeleteNetworkRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNetworkRuleResponse{}
	_body, _err := client.DeleteNetworkRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a permission policy.
//
// Description:
//
// Before you delete a permission policy, make sure that the permission policy is not associated with application access points (AAPs). Otherwise, related applications cannot access Key Management Service (KMS).
//
// @param request - DeletePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicyWithOptions(request *DeletePolicyRequest, runtime *util.RuntimeOptions) (_result *DeletePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePolicy"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeletePolicyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeletePolicyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes a permission policy.
//
// Description:
//
// Before you delete a permission policy, make sure that the permission policy is not associated with application access points (AAPs). Otherwise, related applications cannot access Key Management Service (KMS).
//
// @param request - DeletePolicyRequest
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicy(request *DeletePolicyRequest) (_result *DeletePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePolicyResponse{}
	_body, _err := client.DeletePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// If you call this operation without specifying a recovery period, the deleted secret can be recovered within 30 days.
//
// If you specify a recovery period, the deleted secret can be recovered within the recovery period. You can also forcibly delete a secret. A forcibly deleted secret cannot be recovered.
//
// @param request - DeleteSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecretResponse
func (client *Client) DeleteSecretWithOptions(request *DeleteSecretRequest, runtime *util.RuntimeOptions) (_result *DeleteSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ForceDeleteWithoutRecovery)) {
		query["ForceDeleteWithoutRecovery"] = request.ForceDeleteWithoutRecovery
	}

	if !tea.BoolValue(util.IsUnset(request.RecoveryWindowInDays)) {
		query["RecoveryWindowInDays"] = request.RecoveryWindowInDays
	}

	if !tea.BoolValue(util.IsUnset(request.SecretName)) {
		query["SecretName"] = request.SecretName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSecret"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteSecretResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteSecretResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// If you call this operation without specifying a recovery period, the deleted secret can be recovered within 30 days.
//
// If you specify a recovery period, the deleted secret can be recovered within the recovery period. You can also forcibly delete a secret. A forcibly deleted secret cannot be recovered.
//
// @param request - DeleteSecretRequest
//
// @return DeleteSecretResponse
func (client *Client) DeleteSecret(request *DeleteSecretRequest) (_result *DeleteSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSecretResponse{}
	_body, _err := client.DeleteSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAccountKmsStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountKmsStatusResponse
func (client *Client) DescribeAccountKmsStatusWithOptions(runtime *util.RuntimeOptions) (_result *DescribeAccountKmsStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccountKmsStatus"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeAccountKmsStatusResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeAccountKmsStatusResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @return DescribeAccountKmsStatusResponse
func (client *Client) DescribeAccountKmsStatus() (_result *DescribeAccountKmsStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccountKmsStatusResponse{}
	_body, _err := client.DescribeAccountKmsStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an application access point (AAP).
//
// @param request - DescribeApplicationAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApplicationAccessPointResponse
func (client *Client) DescribeApplicationAccessPointWithOptions(request *DescribeApplicationAccessPointRequest, runtime *util.RuntimeOptions) (_result *DescribeApplicationAccessPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApplicationAccessPoint"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeApplicationAccessPointResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeApplicationAccessPointResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the details of an application access point (AAP).
//
// @param request - DescribeApplicationAccessPointRequest
//
// @return DescribeApplicationAccessPointResponse
func (client *Client) DescribeApplicationAccessPoint(request *DescribeApplicationAccessPointRequest) (_result *DescribeApplicationAccessPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApplicationAccessPointResponse{}
	_body, _err := client.DescribeApplicationAccessPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// In this example, the information about the certificate whose ID is `9a28de48-8d8b-484d-a766-dec4****` is queried. The certificate information includes the certificate ID, creation time, certificate issuer, validity period, serial number, and signature algorithm.
//
// @param request - DescribeCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCertificateResponse
func (client *Client) DescribeCertificateWithOptions(request *DescribeCertificateRequest, runtime *util.RuntimeOptions) (_result *DescribeCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertificateId)) {
		query["CertificateId"] = request.CertificateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCertificate"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeCertificateResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeCertificateResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// In this example, the information about the certificate whose ID is `9a28de48-8d8b-484d-a766-dec4****` is queried. The certificate information includes the certificate ID, creation time, certificate issuer, validity period, serial number, and signature algorithm.
//
// @param request - DescribeCertificateRequest
//
// @return DescribeCertificateResponse
func (client *Client) DescribeCertificate(request *DescribeCertificateRequest) (_result *DescribeCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCertificateResponse{}
	_body, _err := client.DescribeCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a customer master key (CMK).
//
// Description:
//
// You can query the information about the CMK `05754286-3ba2-4fa6-8d41-4323aca6****` by using parameter settings provided in this topic. The information includes the creator, creation time, status, and deletion protection status of the CMK.
//
// @param request - DescribeKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKeyResponse
func (client *Client) DescribeKeyWithOptions(request *DescribeKeyRequest, runtime *util.RuntimeOptions) (_result *DescribeKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeKey"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeKeyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeKeyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the information about a customer master key (CMK).
//
// Description:
//
// You can query the information about the CMK `05754286-3ba2-4fa6-8d41-4323aca6****` by using parameter settings provided in this topic. The information includes the creator, creation time, status, and deletion protection status of the CMK.
//
// @param request - DescribeKeyRequest
//
// @return DescribeKeyResponse
func (client *Client) DescribeKey(request *DescribeKeyRequest) (_result *DescribeKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeKeyResponse{}
	_body, _err := client.DescribeKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// This topic provides an example on how to query the information about a version of the CMK `1234abcd-12ab-34cd-56ef-12345678****`. The ID of the CMK version is `2ab1a983-7072-4bbc-a582-584b5bd8****`. The response shows that the creation time of the CMK version is `2016-03-25T10:42:40Z`.
//
// @param request - DescribeKeyVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKeyVersionResponse
func (client *Client) DescribeKeyVersionWithOptions(request *DescribeKeyVersionRequest, runtime *util.RuntimeOptions) (_result *DescribeKeyVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.KeyVersionId)) {
		query["KeyVersionId"] = request.KeyVersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeKeyVersion"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeKeyVersionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeKeyVersionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// This topic provides an example on how to query the information about a version of the CMK `1234abcd-12ab-34cd-56ef-12345678****`. The ID of the CMK version is `2ab1a983-7072-4bbc-a582-584b5bd8****`. The response shows that the creation time of the CMK version is `2016-03-25T10:42:40Z`.
//
// @param request - DescribeKeyVersionRequest
//
// @return DescribeKeyVersionResponse
func (client *Client) DescribeKeyVersion(request *DescribeKeyVersionRequest) (_result *DescribeKeyVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeKeyVersionResponse{}
	_body, _err := client.DescribeKeyVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an access control rule.
//
// @param request - DescribeNetworkRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetworkRuleResponse
func (client *Client) DescribeNetworkRuleWithOptions(request *DescribeNetworkRuleRequest, runtime *util.RuntimeOptions) (_result *DescribeNetworkRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNetworkRule"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeNetworkRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeNetworkRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the details of an access control rule.
//
// @param request - DescribeNetworkRuleRequest
//
// @return DescribeNetworkRuleResponse
func (client *Client) DescribeNetworkRule(request *DescribeNetworkRuleRequest) (_result *DescribeNetworkRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNetworkRuleResponse{}
	_body, _err := client.DescribeNetworkRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a permission policy.
//
// @param request - DescribePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolicyResponse
func (client *Client) DescribePolicyWithOptions(request *DescribePolicyRequest, runtime *util.RuntimeOptions) (_result *DescribePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePolicy"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribePolicyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribePolicyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the details of a permission policy.
//
// @param request - DescribePolicyRequest
//
// @return DescribePolicyResponse
func (client *Client) DescribePolicy(request *DescribePolicyRequest) (_result *DescribePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePolicyResponse{}
	_body, _err := client.DescribePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries available regions.
//
// Description:
//
// ## Debugging
//
// [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code of the operation for different SDKs.](https://api.aliyun.com/#product=Kms\\&api=DescribeRegions\\&type=RPC\\&version=2016-01-20)
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeRegionsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeRegionsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries available regions.
//
// Description:
//
// ## Debugging
//
// [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code of the operation for different SDKs.](https://api.aliyun.com/#product=Kms\\&api=DescribeRegions\\&type=RPC\\&version=2016-01-20)
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions() (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// This operation returns the metadata of a secret. This operation does not return the secret value.
//
// In this example, the metadata of the secret named `secret001` is queried.
//
// @param request - DescribeSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecretResponse
func (client *Client) DescribeSecretWithOptions(request *DescribeSecretRequest, runtime *util.RuntimeOptions) (_result *DescribeSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FetchTags)) {
		query["FetchTags"] = request.FetchTags
	}

	if !tea.BoolValue(util.IsUnset(request.SecretName)) {
		query["SecretName"] = request.SecretName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSecret"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeSecretResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeSecretResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// This operation returns the metadata of a secret. This operation does not return the secret value.
//
// In this example, the metadata of the secret named `secret001` is queried.
//
// @param request - DescribeSecretRequest
//
// @return DescribeSecretResponse
func (client *Client) DescribeSecret(request *DescribeSecretRequest) (_result *DescribeSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecretResponse{}
	_body, _err := client.DescribeSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// If a customer master key (CMK) is disabled, the ciphertext encrypted by using this CMK cannot be decrypted until you re-enable it. You can call the [EnableKey](https://help.aliyun.com/document_detail/35150.html) operation to enable the CMK.
//
// In this example, the CMK whose ID is `1234abcd-12ab-34cd-56ef-12345678****` is disabled.
//
// @param request - DisableKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableKeyResponse
func (client *Client) DisableKeyWithOptions(request *DisableKeyRequest, runtime *util.RuntimeOptions) (_result *DisableKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableKey"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DisableKeyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DisableKeyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// If a customer master key (CMK) is disabled, the ciphertext encrypted by using this CMK cannot be decrypted until you re-enable it. You can call the [EnableKey](https://help.aliyun.com/document_detail/35150.html) operation to enable the CMK.
//
// In this example, the CMK whose ID is `1234abcd-12ab-34cd-56ef-12345678****` is disabled.
//
// @param request - DisableKeyRequest
//
// @return DisableKeyResponse
func (client *Client) DisableKey(request *DisableKeyRequest) (_result *DisableKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableKeyResponse{}
	_body, _err := client.DisableKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - EnableKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableKeyResponse
func (client *Client) EnableKeyWithOptions(request *EnableKeyRequest, runtime *util.RuntimeOptions) (_result *EnableKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableKey"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &EnableKeyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &EnableKeyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - EnableKeyRequest
//
// @return EnableKeyResponse
func (client *Client) EnableKey(request *EnableKeyRequest) (_result *EnableKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableKeyResponse{}
	_body, _err := client.EnableKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
//   KMS uses the primary version of a specified CMK to encrypt data.
//
// 	- Only data of 6 KB or less can be encrypted. For example, you can call this operation to encrypt RSA keys, database access passwords, or other sensitive information.
//
// 	- When you migrate encrypted data across regions, you can call this operation in the destination region to encrypt the plaintext of the data key that is used to encrypt the migrated data in the source region. This way, the ciphertext of the data key is generated in the destination region. You can also call the [Decrypt](https://help.aliyun.com/document_detail/28950.html) operation to decrypt the data key.
//
// @param tmpReq - EncryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EncryptResponse
func (client *Client) EncryptWithOptions(tmpReq *EncryptRequest, runtime *util.RuntimeOptions) (_result *EncryptResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &EncryptShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EncryptionContext)) {
		request.EncryptionContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EncryptionContext, tea.String("EncryptionContext"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncryptionContextShrink)) {
		query["EncryptionContext"] = request.EncryptionContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.Plaintext)) {
		query["Plaintext"] = request.Plaintext
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Encrypt"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &EncryptResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &EncryptResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
//   KMS uses the primary version of a specified CMK to encrypt data.
//
// 	- Only data of 6 KB or less can be encrypted. For example, you can call this operation to encrypt RSA keys, database access passwords, or other sensitive information.
//
// 	- When you migrate encrypted data across regions, you can call this operation in the destination region to encrypt the plaintext of the data key that is used to encrypt the migrated data in the source region. This way, the ciphertext of the data key is generated in the destination region. You can also call the [Decrypt](https://help.aliyun.com/document_detail/28950.html) operation to decrypt the data key.
//
// @param request - EncryptRequest
//
// @return EncryptResponse
func (client *Client) Encrypt(request *EncryptRequest) (_result *EncryptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EncryptResponse{}
	_body, _err := client.EncryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// You can call the [GenerateDataKeyWithoutPlaintext](https://help.aliyun.com/document_detail/134043.html) operation to generate a data key, which is encrypted by a CMK. If you want to distribute the data key to other regions or cryptographic modules, you can call the ExportDataKey operation to use a public key to encrypt the data key.
//
// Then, you can import the ciphertext of the data key to the cryptographic module where the private key is stored. This way, the data key is securely distributed from KMS to the cryptographic module. After the data key is imported to the cryptographic module, you can use it to encrypt or decrypt data.
//
// @param tmpReq - ExportDataKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportDataKeyResponse
func (client *Client) ExportDataKeyWithOptions(tmpReq *ExportDataKeyRequest, runtime *util.RuntimeOptions) (_result *ExportDataKeyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ExportDataKeyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EncryptionContext)) {
		request.EncryptionContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EncryptionContext, tea.String("EncryptionContext"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CiphertextBlob)) {
		query["CiphertextBlob"] = request.CiphertextBlob
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptionContextShrink)) {
		query["EncryptionContext"] = request.EncryptionContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PublicKeyBlob)) {
		query["PublicKeyBlob"] = request.PublicKeyBlob
	}

	if !tea.BoolValue(util.IsUnset(request.WrappingAlgorithm)) {
		query["WrappingAlgorithm"] = request.WrappingAlgorithm
	}

	if !tea.BoolValue(util.IsUnset(request.WrappingKeySpec)) {
		query["WrappingKeySpec"] = request.WrappingKeySpec
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportDataKey"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ExportDataKeyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ExportDataKeyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// You can call the [GenerateDataKeyWithoutPlaintext](https://help.aliyun.com/document_detail/134043.html) operation to generate a data key, which is encrypted by a CMK. If you want to distribute the data key to other regions or cryptographic modules, you can call the ExportDataKey operation to use a public key to encrypt the data key.
//
// Then, you can import the ciphertext of the data key to the cryptographic module where the private key is stored. This way, the data key is securely distributed from KMS to the cryptographic module. After the data key is imported to the cryptographic module, you can use it to encrypt or decrypt data.
//
// @param request - ExportDataKeyRequest
//
// @return ExportDataKeyResponse
func (client *Client) ExportDataKey(request *ExportDataKeyRequest) (_result *ExportDataKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportDataKeyResponse{}
	_body, _err := client.ExportDataKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// We recommend that you perform the following steps to import your data key to a cryptographic module:
//
// 	- Call the GenerateAndExportDataKey operation to generate a data key and obtain both the ciphertext of the data key encrypted by using the CMK and that encrypted by using the public key.
//
// 	- Store the ciphertext of the data key encrypted by using the CMK in KMS Secrets Manager or in a storage service such as ApsaraDB. This ciphertext is used for backup and restoration.
//
// 	- Import the ciphertext of the data key encrypted by using the public key to the cryptographic module where the private key is stored. Then, you can use the data key to encrypt or decrypt data.
//
// >  The CMK that you specify in the request of this operation is only used to encrypt the data key and is not involved in the generation of the data key. KMS does not record or store the data keys randomly generated by calling this operation. You must take note of the data keys and the returned ciphertext.
//
// @param tmpReq - GenerateAndExportDataKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateAndExportDataKeyResponse
func (client *Client) GenerateAndExportDataKeyWithOptions(tmpReq *GenerateAndExportDataKeyRequest, runtime *util.RuntimeOptions) (_result *GenerateAndExportDataKeyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GenerateAndExportDataKeyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EncryptionContext)) {
		request.EncryptionContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EncryptionContext, tea.String("EncryptionContext"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncryptionContextShrink)) {
		query["EncryptionContext"] = request.EncryptionContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.KeySpec)) {
		query["KeySpec"] = request.KeySpec
	}

	if !tea.BoolValue(util.IsUnset(request.NumberOfBytes)) {
		query["NumberOfBytes"] = request.NumberOfBytes
	}

	if !tea.BoolValue(util.IsUnset(request.PublicKeyBlob)) {
		query["PublicKeyBlob"] = request.PublicKeyBlob
	}

	if !tea.BoolValue(util.IsUnset(request.WrappingAlgorithm)) {
		query["WrappingAlgorithm"] = request.WrappingAlgorithm
	}

	if !tea.BoolValue(util.IsUnset(request.WrappingKeySpec)) {
		query["WrappingKeySpec"] = request.WrappingKeySpec
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateAndExportDataKey"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GenerateAndExportDataKeyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GenerateAndExportDataKeyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// We recommend that you perform the following steps to import your data key to a cryptographic module:
//
// 	- Call the GenerateAndExportDataKey operation to generate a data key and obtain both the ciphertext of the data key encrypted by using the CMK and that encrypted by using the public key.
//
// 	- Store the ciphertext of the data key encrypted by using the CMK in KMS Secrets Manager or in a storage service such as ApsaraDB. This ciphertext is used for backup and restoration.
//
// 	- Import the ciphertext of the data key encrypted by using the public key to the cryptographic module where the private key is stored. Then, you can use the data key to encrypt or decrypt data.
//
// >  The CMK that you specify in the request of this operation is only used to encrypt the data key and is not involved in the generation of the data key. KMS does not record or store the data keys randomly generated by calling this operation. You must take note of the data keys and the returned ciphertext.
//
// @param request - GenerateAndExportDataKeyRequest
//
// @return GenerateAndExportDataKeyResponse
func (client *Client) GenerateAndExportDataKey(request *GenerateAndExportDataKeyRequest) (_result *GenerateAndExportDataKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateAndExportDataKeyResponse{}
	_body, _err := client.GenerateAndExportDataKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成一个数据密钥
//
// Description:
//
// This operation creates a random data key, encrypts the data key by using the specified customer master key (CMK), and returns the plaintext and ciphertext of the data key. You can use the plaintext of the data key to locally encrypt your data without using KMS and store the encrypted data together with the ciphertext of the data key. You can obtain the plaintext of the data key from the Plaintext parameter in the response and the ciphertext of the data key from the CiphertextBlob parameter in the response.
//
// The CMK that you specify in the request of this operation is only used to encrypt the data key and is not involved in the generation of the data key. KMS does not record or store the generated data key. Therefore, you need to store the ciphertext of the data key in persistent storage.
//
// We recommend that you locally encrypt data by performing the following steps:
//
// 1\\. Call the GenerateDataKey operation.
//
// 2\\. Use the plaintext of the data key that you obtain to locally encrypt data without using KMS. Then, delete the plaintext of the data key from the memory.
//
// 3\\. Store the encrypted data together with the ciphertext of the data key that you obtain.
//
// We recommend that you locally decrypt data by performing the following steps:
//
// 	- Call the [Decrypt](https://help.aliyun.com/document_detail/28950.html) operation to decrypt the locally stored ciphertext of the data key. The plaintext of data key is then returned.
//
// 	- Use the plaintext of the data key to locally decrypt data and then delete the plaintext of the data key from the memory.
//
// In this example, a random data key is generated for the CMK whose ID is `7906979c-8e06-46a2-be2d-68e3ccbc****`.
//
// @param tmpReq - GenerateDataKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateDataKeyResponse
func (client *Client) GenerateDataKeyWithOptions(tmpReq *GenerateDataKeyRequest, runtime *util.RuntimeOptions) (_result *GenerateDataKeyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GenerateDataKeyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EncryptionContext)) {
		request.EncryptionContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EncryptionContext, tea.String("EncryptionContext"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncryptionContextShrink)) {
		query["EncryptionContext"] = request.EncryptionContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.KeySpec)) {
		query["KeySpec"] = request.KeySpec
	}

	if !tea.BoolValue(util.IsUnset(request.NumberOfBytes)) {
		query["NumberOfBytes"] = request.NumberOfBytes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateDataKey"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GenerateDataKeyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GenerateDataKeyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 生成一个数据密钥
//
// Description:
//
// This operation creates a random data key, encrypts the data key by using the specified customer master key (CMK), and returns the plaintext and ciphertext of the data key. You can use the plaintext of the data key to locally encrypt your data without using KMS and store the encrypted data together with the ciphertext of the data key. You can obtain the plaintext of the data key from the Plaintext parameter in the response and the ciphertext of the data key from the CiphertextBlob parameter in the response.
//
// The CMK that you specify in the request of this operation is only used to encrypt the data key and is not involved in the generation of the data key. KMS does not record or store the generated data key. Therefore, you need to store the ciphertext of the data key in persistent storage.
//
// We recommend that you locally encrypt data by performing the following steps:
//
// 1\\. Call the GenerateDataKey operation.
//
// 2\\. Use the plaintext of the data key that you obtain to locally encrypt data without using KMS. Then, delete the plaintext of the data key from the memory.
//
// 3\\. Store the encrypted data together with the ciphertext of the data key that you obtain.
//
// We recommend that you locally decrypt data by performing the following steps:
//
// 	- Call the [Decrypt](https://help.aliyun.com/document_detail/28950.html) operation to decrypt the locally stored ciphertext of the data key. The plaintext of data key is then returned.
//
// 	- Use the plaintext of the data key to locally decrypt data and then delete the plaintext of the data key from the memory.
//
// In this example, a random data key is generated for the CMK whose ID is `7906979c-8e06-46a2-be2d-68e3ccbc****`.
//
// @param request - GenerateDataKeyRequest
//
// @return GenerateDataKeyResponse
func (client *Client) GenerateDataKey(request *GenerateDataKeyRequest) (_result *GenerateDataKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateDataKeyResponse{}
	_body, _err := client.GenerateDataKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a random data key, which can be used to encrypt local data.
//
// Description:
//
// This operation creates a random data key, encrypts the data key by using a specific symmetric CMK, and returns the ciphertext of the data key. This operation serves the same purpose as the [GenerateDataKey](https://help.aliyun.com/document_detail/28948.html) operation. The only difference is that this operation does not return the plaintext of the data key.
//
// The CMK that you specify in the request of this operation is only used to encrypt the data key and is not involved in the generation of the data key. KMS does not record or store the generated data key.
//
// > 	- This operation applies to the scenario when you do not need to use the data key to immediately encrypt data. Before you can use the data key to encrypt data, you must call the [Decrypt](https://help.aliyun.com/document_detail/28950.html) operation to decrypt the ciphertext of the data key.
//
// > 	- This operation is also suitable for a distributed system with different trust levels. For example, a system stores data in different partitions based on a preset trust policy. A module creates different partitions and generates different data keys for each partition in advance. This module is not involved in data production and consumption after it completes initialization of the control plane. This module is the key provider. When producing and consuming data, modules on the control plane obtain the ciphertext of the data key for a partition first. After decrypting the ciphertext of the data key, modules on the control plane use the plaintext of the data key to encrypt or decrypt data and then clear the plaintext of the data key from the memory. In such a system, the key provider does not need to obtain the plaintext of the data key. It only needs to have the permissions to call the GenerateDataKeyWithoutPlaintext operation. The data producers or consumers do not need to generate new data keys. They only need to have the permissions to call the Decrypt operation.
//
// @param tmpReq - GenerateDataKeyWithoutPlaintextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateDataKeyWithoutPlaintextResponse
func (client *Client) GenerateDataKeyWithoutPlaintextWithOptions(tmpReq *GenerateDataKeyWithoutPlaintextRequest, runtime *util.RuntimeOptions) (_result *GenerateDataKeyWithoutPlaintextResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GenerateDataKeyWithoutPlaintextShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EncryptionContext)) {
		request.EncryptionContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EncryptionContext, tea.String("EncryptionContext"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncryptionContextShrink)) {
		query["EncryptionContext"] = request.EncryptionContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.KeySpec)) {
		query["KeySpec"] = request.KeySpec
	}

	if !tea.BoolValue(util.IsUnset(request.NumberOfBytes)) {
		query["NumberOfBytes"] = request.NumberOfBytes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateDataKeyWithoutPlaintext"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GenerateDataKeyWithoutPlaintextResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GenerateDataKeyWithoutPlaintextResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Generates a random data key, which can be used to encrypt local data.
//
// Description:
//
// This operation creates a random data key, encrypts the data key by using a specific symmetric CMK, and returns the ciphertext of the data key. This operation serves the same purpose as the [GenerateDataKey](https://help.aliyun.com/document_detail/28948.html) operation. The only difference is that this operation does not return the plaintext of the data key.
//
// The CMK that you specify in the request of this operation is only used to encrypt the data key and is not involved in the generation of the data key. KMS does not record or store the generated data key.
//
// > 	- This operation applies to the scenario when you do not need to use the data key to immediately encrypt data. Before you can use the data key to encrypt data, you must call the [Decrypt](https://help.aliyun.com/document_detail/28950.html) operation to decrypt the ciphertext of the data key.
//
// > 	- This operation is also suitable for a distributed system with different trust levels. For example, a system stores data in different partitions based on a preset trust policy. A module creates different partitions and generates different data keys for each partition in advance. This module is not involved in data production and consumption after it completes initialization of the control plane. This module is the key provider. When producing and consuming data, modules on the control plane obtain the ciphertext of the data key for a partition first. After decrypting the ciphertext of the data key, modules on the control plane use the plaintext of the data key to encrypt or decrypt data and then clear the plaintext of the data key from the memory. In such a system, the key provider does not need to obtain the plaintext of the data key. It only needs to have the permissions to call the GenerateDataKeyWithoutPlaintext operation. The data producers or consumers do not need to generate new data keys. They only need to have the permissions to call the Decrypt operation.
//
// @param request - GenerateDataKeyWithoutPlaintextRequest
//
// @return GenerateDataKeyWithoutPlaintextResponse
func (client *Client) GenerateDataKeyWithoutPlaintext(request *GenerateDataKeyWithoutPlaintextRequest) (_result *GenerateDataKeyWithoutPlaintextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateDataKeyWithoutPlaintextResponse{}
	_body, _err := client.GenerateDataKeyWithoutPlaintextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// In this example, the certificate whose ID is `9a28de48-8d8b-484d-a766-dec4****` is queried. The certificate, certificate chain, certificate ID, and certificate signing request (CSR) are returned.
//
// @param request - GetCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCertificateResponse
func (client *Client) GetCertificateWithOptions(request *GetCertificateRequest, runtime *util.RuntimeOptions) (_result *GetCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertificateId)) {
		query["CertificateId"] = request.CertificateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCertificate"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetCertificateResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetCertificateResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// In this example, the certificate whose ID is `9a28de48-8d8b-484d-a766-dec4****` is queried. The certificate, certificate chain, certificate ID, and certificate signing request (CSR) are returned.
//
// @param request - GetCertificateRequest
//
// @return GetCertificateResponse
func (client *Client) GetCertificate(request *GetCertificateRequest) (_result *GetCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCertificateResponse{}
	_body, _err := client.GetCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a client key.
//
// @param request - GetClientKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClientKeyResponse
func (client *Client) GetClientKeyWithOptions(request *GetClientKeyRequest, runtime *util.RuntimeOptions) (_result *GetClientKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetClientKey"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetClientKeyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetClientKeyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the information about a client key.
//
// @param request - GetClientKeyRequest
//
// @return GetClientKeyResponse
func (client *Client) GetClientKey(request *GetClientKeyRequest) (_result *GetClientKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetClientKeyResponse{}
	_body, _err := client.GetClientKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 仅可查询名称为 default 的 Key Policy，否则提示 Not Found。
//
// @param request - GetKeyPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKeyPolicyResponse
func (client *Client) GetKeyPolicyWithOptions(request *GetKeyPolicyRequest, runtime *util.RuntimeOptions) (_result *GetKeyPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetKeyPolicy"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetKeyPolicyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetKeyPolicyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 仅可查询名称为 default 的 Key Policy，否则提示 Not Found。
//
// @param request - GetKeyPolicyRequest
//
// @return GetKeyPolicyResponse
func (client *Client) GetKeyPolicy(request *GetKeyPolicyRequest) (_result *GetKeyPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetKeyPolicyResponse{}
	_body, _err := client.GetKeyPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a Key Management Service (KMS) instance.
//
// @param request - GetKmsInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKmsInstanceResponse
func (client *Client) GetKmsInstanceWithOptions(request *GetKmsInstanceRequest, runtime *util.RuntimeOptions) (_result *GetKmsInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KmsInstanceId)) {
		query["KmsInstanceId"] = request.KmsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetKmsInstance"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetKmsInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetKmsInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the details of a Key Management Service (KMS) instance.
//
// @param request - GetKmsInstanceRequest
//
// @return GetKmsInstanceResponse
func (client *Client) GetKmsInstance(request *GetKmsInstanceRequest) (_result *GetKmsInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetKmsInstanceResponse{}
	_body, _err := client.GetKmsInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the parameters that are used to import key material for a customer master key (CMK).
//
// Description:
//
// The returned parameters can be used to call the [ImportKeyMaterial](https://www.alibabacloud.com/help/en/key-management-service/latest/importkeymaterial) operation.
//
// - You can import key material only for CMKs whose Origin parameter is set to EXTERNAL.
//
// - The public key and token that are returned by the GetParametersForImport operation must be used together. The public key and token can be used to import key material only for the CMK that is specified when you call the operation.
//
// - The public key and token that are returned vary each time you call the GetParametersForImport operation.
//
// - You must specify the type of the public key and the encryption algorithm that are used to encrypt key material. The following table lists the types of public keys and the encryption algorithms allowed for each type.
//
// | Public key type | Encryption algorithm | Description |
//
// | --------------- | -------------------- | ----------- |
//
// | RSA_2048 | RSAES_PKCS1_V1_5
//
// RSAES_OAEP_SHA_1
//
// RSAES_OAEP_SHA_256 | CMKs of all regions and all protection levels are supported.
//
// Dedicated Key Management Service (KMS) does not support RSAES_OAEP_SHA_1. |
//
// | EC_SM2 | SM2PKE | CMKs whose ProtectionLevel is set to HSM are supported. The SM2 algorithm is developed and approved by the State Cryptography Administration of China. The SM2 algorithm can be used only to import key material for a CMK whose ProtectionLevel is set to HSM. You can use the SM2 algorithm only when you enable the Managed HSM feature for KMS in the Chinese mainland. For more information, see [Overview of Managed HSM](https://www.alibabacloud.com/help/en/key-management-service/latest/managed-hsm-overview). |
//
// For more information, see [Import key material](https://www.alibabacloud.com/help/en/key-management-service/latest/import-key-material). This topic provides an example on how to query the parameters that are used to import key material for a CMK. The ID of the CMK is `1234abcd-12ab-34cd-56ef-12345678****`, the encryption algorithm is `RSAES_PKCS1_V1_5`, and the public key is of the `RSA_2048` type. The parameters that are returned include the ID of the CMK, the public key that is used to encrypt the key material, the token that is used to import the key material, and the time when the token expires.
//
// @param request - GetParametersForImportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetParametersForImportResponse
func (client *Client) GetParametersForImportWithOptions(request *GetParametersForImportRequest, runtime *util.RuntimeOptions) (_result *GetParametersForImportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.WrappingAlgorithm)) {
		query["WrappingAlgorithm"] = request.WrappingAlgorithm
	}

	if !tea.BoolValue(util.IsUnset(request.WrappingKeySpec)) {
		query["WrappingKeySpec"] = request.WrappingKeySpec
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetParametersForImport"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetParametersForImportResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetParametersForImportResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the parameters that are used to import key material for a customer master key (CMK).
//
// Description:
//
// The returned parameters can be used to call the [ImportKeyMaterial](https://www.alibabacloud.com/help/en/key-management-service/latest/importkeymaterial) operation.
//
// - You can import key material only for CMKs whose Origin parameter is set to EXTERNAL.
//
// - The public key and token that are returned by the GetParametersForImport operation must be used together. The public key and token can be used to import key material only for the CMK that is specified when you call the operation.
//
// - The public key and token that are returned vary each time you call the GetParametersForImport operation.
//
// - You must specify the type of the public key and the encryption algorithm that are used to encrypt key material. The following table lists the types of public keys and the encryption algorithms allowed for each type.
//
// | Public key type | Encryption algorithm | Description |
//
// | --------------- | -------------------- | ----------- |
//
// | RSA_2048 | RSAES_PKCS1_V1_5
//
// RSAES_OAEP_SHA_1
//
// RSAES_OAEP_SHA_256 | CMKs of all regions and all protection levels are supported.
//
// Dedicated Key Management Service (KMS) does not support RSAES_OAEP_SHA_1. |
//
// | EC_SM2 | SM2PKE | CMKs whose ProtectionLevel is set to HSM are supported. The SM2 algorithm is developed and approved by the State Cryptography Administration of China. The SM2 algorithm can be used only to import key material for a CMK whose ProtectionLevel is set to HSM. You can use the SM2 algorithm only when you enable the Managed HSM feature for KMS in the Chinese mainland. For more information, see [Overview of Managed HSM](https://www.alibabacloud.com/help/en/key-management-service/latest/managed-hsm-overview). |
//
// For more information, see [Import key material](https://www.alibabacloud.com/help/en/key-management-service/latest/import-key-material). This topic provides an example on how to query the parameters that are used to import key material for a CMK. The ID of the CMK is `1234abcd-12ab-34cd-56ef-12345678****`, the encryption algorithm is `RSAES_PKCS1_V1_5`, and the public key is of the `RSA_2048` type. The parameters that are returned include the ID of the CMK, the public key that is used to encrypt the key material, the token that is used to import the key material, and the time when the token expires.
//
// @param request - GetParametersForImportRequest
//
// @return GetParametersForImportResponse
func (client *Client) GetParametersForImport(request *GetParametersForImportRequest) (_result *GetParametersForImportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetParametersForImportResponse{}
	_body, _err := client.GetParametersForImportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPublicKeyResponse
func (client *Client) GetPublicKeyWithOptions(request *GetPublicKeyRequest, runtime *util.RuntimeOptions) (_result *GetPublicKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.KeyVersionId)) {
		query["KeyVersionId"] = request.KeyVersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPublicKey"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetPublicKeyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetPublicKeyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - GetPublicKeyRequest
//
// @return GetPublicKeyResponse
func (client *Client) GetPublicKey(request *GetPublicKeyRequest) (_result *GetPublicKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPublicKeyResponse{}
	_body, _err := client.GetPublicKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetRandomPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRandomPasswordResponse
func (client *Client) GetRandomPasswordWithOptions(request *GetRandomPasswordRequest, runtime *util.RuntimeOptions) (_result *GetRandomPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExcludeCharacters)) {
		query["ExcludeCharacters"] = request.ExcludeCharacters
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeLowercase)) {
		query["ExcludeLowercase"] = request.ExcludeLowercase
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeNumbers)) {
		query["ExcludeNumbers"] = request.ExcludeNumbers
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludePunctuation)) {
		query["ExcludePunctuation"] = request.ExcludePunctuation
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeUppercase)) {
		query["ExcludeUppercase"] = request.ExcludeUppercase
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordLength)) {
		query["PasswordLength"] = request.PasswordLength
	}

	if !tea.BoolValue(util.IsUnset(request.RequireEachIncludedType)) {
		query["RequireEachIncludedType"] = request.RequireEachIncludedType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRandomPassword"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetRandomPasswordResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetRandomPasswordResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - GetRandomPasswordRequest
//
// @return GetRandomPasswordResponse
func (client *Client) GetRandomPassword(request *GetRandomPasswordRequest) (_result *GetRandomPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRandomPasswordResponse{}
	_body, _err := client.GetRandomPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 仅可查询名称为 default 的 Secret Policy，否则提示 Not Found。
//
// @param request - GetSecretPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecretPolicyResponse
func (client *Client) GetSecretPolicyWithOptions(request *GetSecretPolicyRequest, runtime *util.RuntimeOptions) (_result *GetSecretPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.SecretName)) {
		query["SecretName"] = request.SecretName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSecretPolicy"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetSecretPolicyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetSecretPolicyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 仅可查询名称为 default 的 Secret Policy，否则提示 Not Found。
//
// @param request - GetSecretPolicyRequest
//
// @return GetSecretPolicyResponse
func (client *Client) GetSecretPolicy(request *GetSecretPolicyRequest) (_result *GetSecretPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSecretPolicyResponse{}
	_body, _err := client.GetSecretPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用GetSecretValue接口获取凭据值。
//
// Description:
//
// If you do not specify a version number or stage label, Secrets Manager returns the secret value of the version marked with ACSCurrent.
//
// If a customer master key (CMK) is specified to encrypt the secret value, you must also have the `kms:Decrypt` permission on the CMK to call the GetSecretValue operation.
//
// In this example, the value of the secret named `secret001` is obtained. The secret value is returned in the `SecretData` parameter. The secret value is `testdata1`.
//
// @param request - GetSecretValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecretValueResponse
func (client *Client) GetSecretValueWithOptions(request *GetSecretValueRequest, runtime *util.RuntimeOptions) (_result *GetSecretValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FetchExtendedConfig)) {
		query["FetchExtendedConfig"] = request.FetchExtendedConfig
	}

	if !tea.BoolValue(util.IsUnset(request.SecretName)) {
		query["SecretName"] = request.SecretName
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionStage)) {
		query["VersionStage"] = request.VersionStage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSecretValue"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetSecretValueResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetSecretValueResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 调用GetSecretValue接口获取凭据值。
//
// Description:
//
// If you do not specify a version number or stage label, Secrets Manager returns the secret value of the version marked with ACSCurrent.
//
// If a customer master key (CMK) is specified to encrypt the secret value, you must also have the `kms:Decrypt` permission on the CMK to call the GetSecretValue operation.
//
// In this example, the value of the secret named `secret001` is obtained. The secret value is returned in the `SecretData` parameter. The secret value is `testdata1`.
//
// @param request - GetSecretValueRequest
//
// @return GetSecretValueResponse
func (client *Client) GetSecretValue(request *GetSecretValueRequest) (_result *GetSecretValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSecretValueResponse{}
	_body, _err := client.GetSecretValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the ImportKeyMaterial operation to import the key material.
//
// Description:
//
// Call [CreateKey](https://help.aliyun.com/document_detail/28947.html) when creating a CMK, you can select its key material source as external. **Origin*	- set to **EXTERNAL**. This API is used to import the key material into the CMK.
//
// 	- To view the CMK **Origin**, see [DescribeKey](https://help.aliyun.com/document_detail/28952.html).
//
// 	- Before importing key material, you need to call the [GetParametersForImport](https://help.aliyun.com/document_detail/68621.html) obtain the parameters required to import the key material, including the public key and import token.
//
// > 	- The key type of the pair is **Aliyun_AES_256*	- the key material must be 256 bits. The key type must be **Aliyun_SM4*	- the CMK and key material must be 128 bits.
//
// > 	- You can set the expiration time for the key material, or you can set it to never expire.
//
// > 	- You can reimport the key material and reset the expiration time for the specified CMK at any time, but the same key material must be imported.
//
// > 	- After the imported key material expires or is deleted, the specified CMK is unavailable until the same key material are imported again.
//
// > 	- A Key material can be imported to multiple cmks, but any Data or Data Key encrypted by one CMK cannot be decrypted by another CMK.
//
// @param request - ImportKeyMaterialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportKeyMaterialResponse
func (client *Client) ImportKeyMaterialWithOptions(request *ImportKeyMaterialRequest, runtime *util.RuntimeOptions) (_result *ImportKeyMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncryptedKeyMaterial)) {
		query["EncryptedKeyMaterial"] = request.EncryptedKeyMaterial
	}

	if !tea.BoolValue(util.IsUnset(request.ImportToken)) {
		query["ImportToken"] = request.ImportToken
	}

	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.KeyMaterialExpireUnix)) {
		query["KeyMaterialExpireUnix"] = request.KeyMaterialExpireUnix
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportKeyMaterial"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ImportKeyMaterialResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ImportKeyMaterialResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Call the ImportKeyMaterial operation to import the key material.
//
// Description:
//
// Call [CreateKey](https://help.aliyun.com/document_detail/28947.html) when creating a CMK, you can select its key material source as external. **Origin*	- set to **EXTERNAL**. This API is used to import the key material into the CMK.
//
// 	- To view the CMK **Origin**, see [DescribeKey](https://help.aliyun.com/document_detail/28952.html).
//
// 	- Before importing key material, you need to call the [GetParametersForImport](https://help.aliyun.com/document_detail/68621.html) obtain the parameters required to import the key material, including the public key and import token.
//
// > 	- The key type of the pair is **Aliyun_AES_256*	- the key material must be 256 bits. The key type must be **Aliyun_SM4*	- the CMK and key material must be 128 bits.
//
// > 	- You can set the expiration time for the key material, or you can set it to never expire.
//
// > 	- You can reimport the key material and reset the expiration time for the specified CMK at any time, but the same key material must be imported.
//
// > 	- After the imported key material expires or is deleted, the specified CMK is unavailable until the same key material are imported again.
//
// > 	- A Key material can be imported to multiple cmks, but any Data or Data Key encrypted by one CMK cannot be decrypted by another CMK.
//
// @param request - ImportKeyMaterialRequest
//
// @return ImportKeyMaterialResponse
func (client *Client) ImportKeyMaterial(request *ImportKeyMaterialRequest) (_result *ImportKeyMaterialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportKeyMaterialResponse{}
	_body, _err := client.ImportKeyMaterialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all aliases in the current region for the current account.
//
// @param request - ListAliasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAliasesResponse
func (client *Client) ListAliasesWithOptions(request *ListAliasesRequest, runtime *util.RuntimeOptions) (_result *ListAliasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAliases"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListAliasesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListAliasesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries all aliases in the current region for the current account.
//
// @param request - ListAliasesRequest
//
// @return ListAliasesResponse
func (client *Client) ListAliases(request *ListAliasesRequest) (_result *ListAliasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAliasesResponse{}
	_body, _err := client.ListAliasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListAliasesByKeyIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAliasesByKeyIdResponse
func (client *Client) ListAliasesByKeyIdWithOptions(request *ListAliasesByKeyIdRequest, runtime *util.RuntimeOptions) (_result *ListAliasesByKeyIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAliasesByKeyId"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListAliasesByKeyIdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListAliasesByKeyIdResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ListAliasesByKeyIdRequest
//
// @return ListAliasesByKeyIdResponse
func (client *Client) ListAliasesByKeyId(request *ListAliasesByKeyIdRequest) (_result *ListAliasesByKeyIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAliasesByKeyIdResponse{}
	_body, _err := client.ListAliasesByKeyIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of application access points (AAPs).
//
// @param request - ListApplicationAccessPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationAccessPointsResponse
func (client *Client) ListApplicationAccessPointsWithOptions(request *ListApplicationAccessPointsRequest, runtime *util.RuntimeOptions) (_result *ListApplicationAccessPointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApplicationAccessPoints"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListApplicationAccessPointsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListApplicationAccessPointsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries a list of application access points (AAPs).
//
// @param request - ListApplicationAccessPointsRequest
//
// @return ListApplicationAccessPointsResponse
func (client *Client) ListApplicationAccessPoints(request *ListApplicationAccessPointsRequest) (_result *ListApplicationAccessPointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplicationAccessPointsResponse{}
	_body, _err := client.ListApplicationAccessPointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListClientKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClientKeysResponse
func (client *Client) ListClientKeysWithOptions(request *ListClientKeysRequest, runtime *util.RuntimeOptions) (_result *ListClientKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClientKeys"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListClientKeysResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListClientKeysResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ListClientKeysRequest
//
// @return ListClientKeysResponse
func (client *Client) ListClientKeys(request *ListClientKeysRequest) (_result *ListClientKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClientKeysResponse{}
	_body, _err := client.ListClientKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all versions of a specified CMK.
//
// @param request - ListKeyVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKeyVersionsResponse
func (client *Client) ListKeyVersionsWithOptions(request *ListKeyVersionsRequest, runtime *util.RuntimeOptions) (_result *ListKeyVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListKeyVersions"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListKeyVersionsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListKeyVersionsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries all versions of a specified CMK.
//
// @param request - ListKeyVersionsRequest
//
// @return ListKeyVersionsResponse
func (client *Client) ListKeyVersions(request *ListKeyVersionsRequest) (_result *ListKeyVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListKeyVersionsResponse{}
	_body, _err := client.ListKeyVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all customer master keys (CMKs) of the current Alibaba Cloud account in the current region.
//
// @param request - ListKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKeysResponse
func (client *Client) ListKeysWithOptions(request *ListKeysRequest, runtime *util.RuntimeOptions) (_result *ListKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListKeys"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListKeysResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListKeysResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries all customer master keys (CMKs) of the current Alibaba Cloud account in the current region.
//
// @param request - ListKeysRequest
//
// @return ListKeysResponse
func (client *Client) ListKeys(request *ListKeysRequest) (_result *ListKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListKeysResponse{}
	_body, _err := client.ListKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of Key Management Service (KMS) instances.
//
// @param request - ListKmsInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKmsInstancesResponse
func (client *Client) ListKmsInstancesWithOptions(request *ListKmsInstancesRequest, runtime *util.RuntimeOptions) (_result *ListKmsInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListKmsInstances"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListKmsInstancesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListKmsInstancesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries a list of Key Management Service (KMS) instances.
//
// @param request - ListKmsInstancesRequest
//
// @return ListKmsInstancesResponse
func (client *Client) ListKmsInstances(request *ListKmsInstancesRequest) (_result *ListKmsInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListKmsInstancesResponse{}
	_body, _err := client.ListKmsInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of access control rules.
//
// @param request - ListNetworkRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetworkRulesResponse
func (client *Client) ListNetworkRulesWithOptions(request *ListNetworkRulesRequest, runtime *util.RuntimeOptions) (_result *ListNetworkRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNetworkRules"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListNetworkRulesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListNetworkRulesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries a list of access control rules.
//
// @param request - ListNetworkRulesRequest
//
// @return ListNetworkRulesResponse
func (client *Client) ListNetworkRules(request *ListNetworkRulesRequest) (_result *ListNetworkRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNetworkRulesResponse{}
	_body, _err := client.ListNetworkRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of permission policies.
//
// @param request - ListPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPoliciesResponse
func (client *Client) ListPoliciesWithOptions(request *ListPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPolicies"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListPoliciesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListPoliciesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries a list of permission policies.
//
// @param request - ListPoliciesRequest
//
// @return ListPoliciesResponse
func (client *Client) ListPolicies(request *ListPoliciesRequest) (_result *ListPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPoliciesResponse{}
	_body, _err := client.ListPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// Request format: KeyId="string"
//
// @param request - ListResourceTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceTagsResponse
func (client *Client) ListResourceTagsWithOptions(request *ListResourceTagsRequest, runtime *util.RuntimeOptions) (_result *ListResourceTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceTags"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListResourceTagsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListResourceTagsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// Request format: KeyId="string"
//
// @param request - ListResourceTagsRequest
//
// @return ListResourceTagsResponse
func (client *Client) ListResourceTags(request *ListResourceTagsRequest) (_result *ListResourceTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceTagsResponse{}
	_body, _err := client.ListResourceTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// The secret value is not included in the returned version information. By default, deprecated secret versions are not returned.
//
// @param request - ListSecretVersionIdsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecretVersionIdsResponse
func (client *Client) ListSecretVersionIdsWithOptions(request *ListSecretVersionIdsRequest, runtime *util.RuntimeOptions) (_result *ListSecretVersionIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncludeDeprecated)) {
		query["IncludeDeprecated"] = request.IncludeDeprecated
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecretName)) {
		query["SecretName"] = request.SecretName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSecretVersionIds"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListSecretVersionIdsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListSecretVersionIdsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// The secret value is not included in the returned version information. By default, deprecated secret versions are not returned.
//
// @param request - ListSecretVersionIdsRequest
//
// @return ListSecretVersionIdsResponse
func (client *Client) ListSecretVersionIds(request *ListSecretVersionIdsRequest) (_result *ListSecretVersionIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSecretVersionIdsResponse{}
	_body, _err := client.ListSecretVersionIdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// Specifies whether to return the resource tags of the secret. Valid values:
//
// 	- true: returns the resource tags.
//
// 	- false: does not return the resource tags. This is the default value.
//
// @param request - ListSecretsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecretsResponse
func (client *Client) ListSecretsWithOptions(request *ListSecretsRequest, runtime *util.RuntimeOptions) (_result *ListSecretsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FetchTags)) {
		query["FetchTags"] = request.FetchTags
	}

	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSecrets"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListSecretsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListSecretsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// Specifies whether to return the resource tags of the secret. Valid values:
//
// 	- true: returns the resource tags.
//
// 	- false: does not return the resource tags. This is the default value.
//
// @param request - ListSecretsRequest
//
// @return ListSecretsResponse
func (client *Client) ListSecrets(request *ListSecretsRequest) (_result *ListSecretsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSecretsResponse{}
	_body, _err := client.ListSecretsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags of a key or a secret.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTagResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the tags of a key or a secret.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Activates Key Management Service (KMS) under your Alibaba cloud account.
//
// Description:
//
// When you call this operation, note that:
//
// - KMS is a paid service. For more information about the billing method, see [Billing description](https://www.alibabacloud.com/help/en/key-management-service/latest/billing-billing).
//
// - An Alibaba Cloud account can activate KMS only once.
//
// - Make sure that your Alibaba Cloud account has passed real-name authentication.
//
// @param request - OpenKmsServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenKmsServiceResponse
func (client *Client) OpenKmsServiceWithOptions(runtime *util.RuntimeOptions) (_result *OpenKmsServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("OpenKmsService"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OpenKmsServiceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OpenKmsServiceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Activates Key Management Service (KMS) under your Alibaba cloud account.
//
// Description:
//
// When you call this operation, note that:
//
// - KMS is a paid service. For more information about the billing method, see [Billing description](https://www.alibabacloud.com/help/en/key-management-service/latest/billing-billing).
//
// - An Alibaba Cloud account can activate KMS only once.
//
// - Make sure that your Alibaba Cloud account has passed real-name authentication.
//
// @return OpenKmsServiceResponse
func (client *Client) OpenKmsService() (_result *OpenKmsServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenKmsServiceResponse{}
	_body, _err := client.OpenKmsServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// This operation is used to store the secret values of new versions. It cannot be used to modify the secret value of an existing version.
//
// By default, the newly stored secret value is marked with ACSCurrent, and the mark for the previous version of the secret value is changed from ACSCurrent to ACSPrevious. If you specify the VersionStage parameter, the newly stored secret value is marked with the stage label that you specify.
//
// You must specify a version number when you call the operation. Secrets Manager performs operations based on the following rules:
//
// 	- If the specified version number does not exist in the secret, Secrets Manager creates the version and stores the secret value.
//
// 	- If the specified version number already exists in the secret and the secret value of the existing version is the same as the secret value that you specify, Secrets Manager ignores the request and returns a success message. The request is idempotent.
//
// 	- If the specified version number already exists in the secret but the secret value of the existing version is different from the secret value that you specify, Secrets Manager rejects the request and returns a failure message.
//
// Limits: This operation is available only for standard secrets.
//
// In this example, the secret value of a new version is stored into the `secret001` secret. The `VersionId` parameter is set to `00000000000000000000000000000000203` as the new version, and the `SecretData` parameter is set to `importantdata`.
//
// @param request - PutSecretValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutSecretValueResponse
func (client *Client) PutSecretValueWithOptions(request *PutSecretValueRequest, runtime *util.RuntimeOptions) (_result *PutSecretValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecretData)) {
		query["SecretData"] = request.SecretData
	}

	if !tea.BoolValue(util.IsUnset(request.SecretDataType)) {
		query["SecretDataType"] = request.SecretDataType
	}

	if !tea.BoolValue(util.IsUnset(request.SecretName)) {
		query["SecretName"] = request.SecretName
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionStages)) {
		query["VersionStages"] = request.VersionStages
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutSecretValue"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &PutSecretValueResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &PutSecretValueResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// This operation is used to store the secret values of new versions. It cannot be used to modify the secret value of an existing version.
//
// By default, the newly stored secret value is marked with ACSCurrent, and the mark for the previous version of the secret value is changed from ACSCurrent to ACSPrevious. If you specify the VersionStage parameter, the newly stored secret value is marked with the stage label that you specify.
//
// You must specify a version number when you call the operation. Secrets Manager performs operations based on the following rules:
//
// 	- If the specified version number does not exist in the secret, Secrets Manager creates the version and stores the secret value.
//
// 	- If the specified version number already exists in the secret and the secret value of the existing version is the same as the secret value that you specify, Secrets Manager ignores the request and returns a success message. The request is idempotent.
//
// 	- If the specified version number already exists in the secret but the secret value of the existing version is different from the secret value that you specify, Secrets Manager rejects the request and returns a failure message.
//
// Limits: This operation is available only for standard secrets.
//
// In this example, the secret value of a new version is stored into the `secret001` secret. The `VersionId` parameter is set to `00000000000000000000000000000000203` as the new version, and the `SecretData` parameter is set to `importantdata`.
//
// @param request - PutSecretValueRequest
//
// @return PutSecretValueResponse
func (client *Client) PutSecretValue(request *PutSecretValueRequest) (_result *PutSecretValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutSecretValueResponse{}
	_body, _err := client.PutSecretValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// You can call this operation in the following scenarios:
//
// 	- After the CMK that was used to encrypt your data is rotated, you can call this operation to use the latest CMK version to re-encrypt the data. For more information about automatic key rotation, see [Configure automatic key rotation](https://help.aliyun.com/document_detail/134270.html).
//
// 	- The CMK that was used to encrypt your data remains unchanged, but EncryptionContext is changed. In this scenario, you can call this operation to re-encrypt the data.
//
// 	- You can call this operation to use a CMK in KMS to re-encrypt data or a data key that was previously encrypted by a different CMK.
//
// To use the ReEncrypt operation, you must have two permissions:
//
// 	- kms:ReEncryptFrom on the source CMK
//
// 	- kms:ReEncryptTo on the destination CMK
//
// 	- For simplicity, you can specify kms:ReEncrypt\\	- to allow both of the preceding permissions.
//
// @param tmpReq - ReEncryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReEncryptResponse
func (client *Client) ReEncryptWithOptions(tmpReq *ReEncryptRequest, runtime *util.RuntimeOptions) (_result *ReEncryptResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ReEncryptShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DestinationEncryptionContext)) {
		request.DestinationEncryptionContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DestinationEncryptionContext, tea.String("DestinationEncryptionContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceEncryptionContext)) {
		request.SourceEncryptionContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceEncryptionContext, tea.String("SourceEncryptionContext"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CiphertextBlob)) {
		query["CiphertextBlob"] = request.CiphertextBlob
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationEncryptionContextShrink)) {
		query["DestinationEncryptionContext"] = request.DestinationEncryptionContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationKeyId)) {
		query["DestinationKeyId"] = request.DestinationKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceEncryptionAlgorithm)) {
		query["SourceEncryptionAlgorithm"] = request.SourceEncryptionAlgorithm
	}

	if !tea.BoolValue(util.IsUnset(request.SourceEncryptionContextShrink)) {
		query["SourceEncryptionContext"] = request.SourceEncryptionContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceKeyId)) {
		query["SourceKeyId"] = request.SourceKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceKeyVersionId)) {
		query["SourceKeyVersionId"] = request.SourceKeyVersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReEncrypt"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ReEncryptResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ReEncryptResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// You can call this operation in the following scenarios:
//
// 	- After the CMK that was used to encrypt your data is rotated, you can call this operation to use the latest CMK version to re-encrypt the data. For more information about automatic key rotation, see [Configure automatic key rotation](https://help.aliyun.com/document_detail/134270.html).
//
// 	- The CMK that was used to encrypt your data remains unchanged, but EncryptionContext is changed. In this scenario, you can call this operation to re-encrypt the data.
//
// 	- You can call this operation to use a CMK in KMS to re-encrypt data or a data key that was previously encrypted by a different CMK.
//
// To use the ReEncrypt operation, you must have two permissions:
//
// 	- kms:ReEncryptFrom on the source CMK
//
// 	- kms:ReEncryptTo on the destination CMK
//
// 	- For simplicity, you can specify kms:ReEncrypt\\	- to allow both of the preceding permissions.
//
// @param request - ReEncryptRequest
//
// @return ReEncryptResponse
func (client *Client) ReEncrypt(request *ReEncryptRequest) (_result *ReEncryptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReEncryptResponse{}
	_body, _err := client.ReEncryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// You can only use this operation to restore a deleted secret that is within its recovery period. If you set **ForceDeleteWithoutRecovery*	- to **true*	- when you delete the secret, you cannot restore it.
//
// @param request - RestoreSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreSecretResponse
func (client *Client) RestoreSecretWithOptions(request *RestoreSecretRequest, runtime *util.RuntimeOptions) (_result *RestoreSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecretName)) {
		query["SecretName"] = request.SecretName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestoreSecret"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RestoreSecretResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RestoreSecretResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// You can only use this operation to restore a deleted secret that is within its recovery period. If you set **ForceDeleteWithoutRecovery*	- to **true*	- when you delete the secret, you cannot restore it.
//
// @param request - RestoreSecretRequest
//
// @return RestoreSecretResponse
func (client *Client) RestoreSecret(request *RestoreSecretRequest) (_result *RestoreSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestoreSecretResponse{}
	_body, _err := client.RestoreSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// Limits:
//
// • A secret of each Alibaba Cloud account can be rotated for a maximum of 50 times per hour.
//
// • The RotateSecret operation is unavailable for standard secrets.
//
// In this example, the `RdsSecret/Mysql5.4/MyCred` secret is manually rotated, and the version number of the secret is set to `000000123` after the secret is rotated.
//
// @param request - RotateSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RotateSecretResponse
func (client *Client) RotateSecretWithOptions(request *RotateSecretRequest, runtime *util.RuntimeOptions) (_result *RotateSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecretName)) {
		query["SecretName"] = request.SecretName
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RotateSecret"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RotateSecretResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RotateSecretResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// Limits:
//
// • A secret of each Alibaba Cloud account can be rotated for a maximum of 50 times per hour.
//
// • The RotateSecret operation is unavailable for standard secrets.
//
// In this example, the `RdsSecret/Mysql5.4/MyCred` secret is manually rotated, and the version number of the secret is set to `000000123` after the secret is rotated.
//
// @param request - RotateSecretRequest
//
// @return RotateSecretResponse
func (client *Client) RotateSecret(request *RotateSecretRequest) (_result *RotateSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RotateSecretResponse{}
	_body, _err := client.RotateSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// During the scheduled period, the CMK is in the PendingDeletion state and cannot be used to encrypt data, decrypt data, or generate data keys.
//
// After a CMK is deleted, it cannot be recovered. Data that is encrypted and data keys that are generated by using the CMK cannot be decrypted. To prevent accidental deletion of CMKs, Key Management Service (KMS) allows you to only schedule key deletion tasks. You cannot directly delete CMKs. If you want to delete a CMK, call the [DisableKey](https://help.aliyun.com/document_detail/35151.html) operation to disable the CMK.
//
// When you call this operation, you must specify a scheduled period between 7 days to 366 days. The scheduled period starts from the time when you submit the request. You can call the [CancelKeyDeletion](https://help.aliyun.com/document_detail/44197.html) operation to cancel the key deletion task before the scheduled period ends.
//
// @param request - ScheduleKeyDeletionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScheduleKeyDeletionResponse
func (client *Client) ScheduleKeyDeletionWithOptions(request *ScheduleKeyDeletionRequest, runtime *util.RuntimeOptions) (_result *ScheduleKeyDeletionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.PendingWindowInDays)) {
		query["PendingWindowInDays"] = request.PendingWindowInDays
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ScheduleKeyDeletion"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ScheduleKeyDeletionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ScheduleKeyDeletionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// During the scheduled period, the CMK is in the PendingDeletion state and cannot be used to encrypt data, decrypt data, or generate data keys.
//
// After a CMK is deleted, it cannot be recovered. Data that is encrypted and data keys that are generated by using the CMK cannot be decrypted. To prevent accidental deletion of CMKs, Key Management Service (KMS) allows you to only schedule key deletion tasks. You cannot directly delete CMKs. If you want to delete a CMK, call the [DisableKey](https://help.aliyun.com/document_detail/35151.html) operation to disable the CMK.
//
// When you call this operation, you must specify a scheduled period between 7 days to 366 days. The scheduled period starts from the time when you submit the request. You can call the [CancelKeyDeletion](https://help.aliyun.com/document_detail/44197.html) operation to cancel the key deletion task before the scheduled period ends.
//
// @param request - ScheduleKeyDeletionRequest
//
// @return ScheduleKeyDeletionResponse
func (client *Client) ScheduleKeyDeletion(request *ScheduleKeyDeletionRequest) (_result *ScheduleKeyDeletionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ScheduleKeyDeletionResponse{}
	_body, _err := client.ScheduleKeyDeletionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables deletion protection for a customer master key (CMK).
//
// Description:
//
//   After you enable deletion protection for a CMK, you cannot delete the CMK. If you want to delete the CMK, you must first disable deletion protection for the CMK.
//
// 	- Before you can call the SetDeletionProtection operation, make sure that the required CMK is not in the Pending Deletion state. You can call the [DescribeKey](https://help.aliyun.com/document_detail/28952.html) operation to query the CMK status, which is specified by the KeyState parameter.
//
// You can enable deletion protection for the CMK whose Alibaba Cloud Resource Name (ARN) is `acs:kms:cn-hangzhou:123213123****:key/0225f411-b21d-46d1-be5b-93931c82****` by using parameter settings provided in this topic. The CMK ARN is specified by the ProtectedResourceArn parameter.
//
// @param request - SetDeletionProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDeletionProtectionResponse
func (client *Client) SetDeletionProtectionWithOptions(request *SetDeletionProtectionRequest, runtime *util.RuntimeOptions) (_result *SetDeletionProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeletionProtectionDescription)) {
		query["DeletionProtectionDescription"] = request.DeletionProtectionDescription
	}

	if !tea.BoolValue(util.IsUnset(request.EnableDeletionProtection)) {
		query["EnableDeletionProtection"] = request.EnableDeletionProtection
	}

	if !tea.BoolValue(util.IsUnset(request.ProtectedResourceArn)) {
		query["ProtectedResourceArn"] = request.ProtectedResourceArn
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDeletionProtection"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SetDeletionProtectionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SetDeletionProtectionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Enables or disables deletion protection for a customer master key (CMK).
//
// Description:
//
//   After you enable deletion protection for a CMK, you cannot delete the CMK. If you want to delete the CMK, you must first disable deletion protection for the CMK.
//
// 	- Before you can call the SetDeletionProtection operation, make sure that the required CMK is not in the Pending Deletion state. You can call the [DescribeKey](https://help.aliyun.com/document_detail/28952.html) operation to query the CMK status, which is specified by the KeyState parameter.
//
// You can enable deletion protection for the CMK whose Alibaba Cloud Resource Name (ARN) is `acs:kms:cn-hangzhou:123213123****:key/0225f411-b21d-46d1-be5b-93931c82****` by using parameter settings provided in this topic. The CMK ARN is specified by the ProtectedResourceArn parameter.
//
// @param request - SetDeletionProtectionRequest
//
// @return SetDeletionProtectionResponse
func (client *Client) SetDeletionProtection(request *SetDeletionProtectionRequest) (_result *SetDeletionProtectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDeletionProtectionResponse{}
	_body, _err := client.SetDeletionProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 可以设置一条 Key Policy，且名称必须为 default。
//
// @param request - SetKeyPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetKeyPolicyResponse
func (client *Client) SetKeyPolicyWithOptions(request *SetKeyPolicyRequest, runtime *util.RuntimeOptions) (_result *SetKeyPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		query["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetKeyPolicy"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SetKeyPolicyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SetKeyPolicyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 可以设置一条 Key Policy，且名称必须为 default。
//
// @param request - SetKeyPolicyRequest
//
// @return SetKeyPolicyResponse
func (client *Client) SetKeyPolicy(request *SetKeyPolicyRequest) (_result *SetKeyPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetKeyPolicyResponse{}
	_body, _err := client.SetKeyPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 可以设置一条 Secret Policy，且名称必须为 default。
//
// @param request - SetSecretPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSecretPolicyResponse
func (client *Client) SetSecretPolicyWithOptions(request *SetSecretPolicyRequest, runtime *util.RuntimeOptions) (_result *SetSecretPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		query["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.SecretName)) {
		query["SecretName"] = request.SecretName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetSecretPolicy"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SetSecretPolicyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SetSecretPolicyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 可以设置一条 Secret Policy，且名称必须为 default。
//
// @param request - SetSecretPolicyRequest
//
// @return SetSecretPolicyResponse
func (client *Client) SetSecretPolicy(request *SetSecretPolicyRequest) (_result *SetSecretPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetSecretPolicyResponse{}
	_body, _err := client.SetSecretPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// You can add up to 10 tags to a CMK, secret, or certificate.
//
// In this example, the tags `[{"TagKey":"S1key1","TagValue":"S1val1"},{"TagKey":"S1key2","TagValue":"S2val2"}]` are added to the CMK whose ID is `08c33a6f-4e0a-4a1b-a3fa-7ddf****`.
//
// @param request - TagResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourceResponse
func (client *Client) TagResourceWithOptions(request *TagResourceRequest, runtime *util.RuntimeOptions) (_result *TagResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertificateId)) {
		query["CertificateId"] = request.CertificateId
	}

	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretName)) {
		query["SecretName"] = request.SecretName
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResource"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &TagResourceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &TagResourceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// You can add up to 10 tags to a CMK, secret, or certificate.
//
// In this example, the tags `[{"TagKey":"S1key1","TagValue":"S1val1"},{"TagKey":"S1key2","TagValue":"S2val2"}]` are added to the CMK whose ID is `08c33a6f-4e0a-4a1b-a3fa-7ddf****`.
//
// @param request - TagResourceRequest
//
// @return TagResourceResponse
func (client *Client) TagResource(request *TagResourceRequest) (_result *TagResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourceResponse{}
	_body, _err := client.TagResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tags to keys or secrets.
//
// Description:
//
// You can add multiple tags to multiple keys or multiple secrets at a time.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &TagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &TagResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Adds tags to keys or secrets.
//
// Description:
//
// You can add multiple tags to multiple keys or multiple secrets at a time.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// One or more tag keys. Separate multiple tag keys with commas (,).
//
// You need to specify only the tag keys, not the tag values.
//
// Each tag key must be 1 to 128 bytes in length.
//
// @param request - UntagResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourceResponse
func (client *Client) UntagResourceWithOptions(request *UntagResourceRequest, runtime *util.RuntimeOptions) (_result *UntagResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertificateId)) {
		query["CertificateId"] = request.CertificateId
	}

	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretName)) {
		query["SecretName"] = request.SecretName
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeys)) {
		query["TagKeys"] = request.TagKeys
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResource"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UntagResourceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UntagResourceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// One or more tag keys. Separate multiple tag keys with commas (,).
//
// You need to specify only the tag keys, not the tag values.
//
// Each tag key must be 1 to 128 bytes in length.
//
// @param request - UntagResourceRequest
//
// @return UntagResourceResponse
func (client *Client) UntagResource(request *UntagResourceRequest) (_result *UntagResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourceResponse{}
	_body, _err := client.UntagResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from keys or secrets.
//
// Description:
//
// You can remove multiple tags from multiple keys or multiple secrets at a time. You cannot remove tags that start with aliyun or acs:.
//
// If you enter multiple tag keys in the request parameters and only some of the tag keys are associated with resources, the operation can be called and the tags whose keys are associated with resources are removed from the resources.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UntagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UntagResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Removes tags from keys or secrets.
//
// Description:
//
// You can remove multiple tags from multiple keys or multiple secrets at a time. You cannot remove tags that start with aliyun or acs:.
//
// If you enter multiple tag keys in the request parameters and only some of the tag keys are associated with resources, the operation can be called and the tags whose keys are associated with resources are removed from the resources.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAliasResponse
func (client *Client) UpdateAliasWithOptions(request *UpdateAliasRequest, runtime *util.RuntimeOptions) (_result *UpdateAliasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliasName)) {
		query["AliasName"] = request.AliasName
	}

	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAlias"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateAliasResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateAliasResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - UpdateAliasRequest
//
// @return UpdateAliasResponse
func (client *Client) UpdateAlias(request *UpdateAliasRequest) (_result *UpdateAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAliasResponse{}
	_body, _err := client.UpdateAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// The update takes effect immediately after an AAP information is updated. Exercise caution when you perform this operation. You can update the description of an AAP and the permission policies that are associated with the AAP. You cannot update the name of the AAP.
//
// @param request - UpdateApplicationAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationAccessPointResponse
func (client *Client) UpdateApplicationAccessPointWithOptions(request *UpdateApplicationAccessPointRequest, runtime *util.RuntimeOptions) (_result *UpdateApplicationAccessPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Policies)) {
		query["Policies"] = request.Policies
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApplicationAccessPoint"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateApplicationAccessPointResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateApplicationAccessPointResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// The update takes effect immediately after an AAP information is updated. Exercise caution when you perform this operation. You can update the description of an AAP and the permission policies that are associated with the AAP. You cannot update the name of the AAP.
//
// @param request - UpdateApplicationAccessPointRequest
//
// @return UpdateApplicationAccessPointResponse
func (client *Client) UpdateApplicationAccessPoint(request *UpdateApplicationAccessPointRequest) (_result *UpdateApplicationAccessPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateApplicationAccessPointResponse{}
	_body, _err := client.UpdateApplicationAccessPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// In this example, the status of the certificate whose ID is `9a28de48-8d8b-484d-a766-dec4****` is updated to INACTIVE.
//
// @param request - UpdateCertificateStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCertificateStatusResponse
func (client *Client) UpdateCertificateStatusWithOptions(request *UpdateCertificateStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateCertificateStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertificateId)) {
		query["CertificateId"] = request.CertificateId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCertificateStatus"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateCertificateStatusResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateCertificateStatusResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// In this example, the status of the certificate whose ID is `9a28de48-8d8b-484d-a766-dec4****` is updated to INACTIVE.
//
// @param request - UpdateCertificateStatusRequest
//
// @return UpdateCertificateStatusResponse
func (client *Client) UpdateCertificateStatus(request *UpdateCertificateStatusRequest) (_result *UpdateCertificateStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCertificateStatusResponse{}
	_body, _err := client.UpdateCertificateStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用UpdateKeyDescription接口更新主密钥的描述信息。
//
// Description:
//
// This operation replaces the description of a customer master key (CMK) with the description that you specify. The original description of the CMK is specified by the Description parameter when you call the [DescribeKey](https://help.aliyun.com/document_detail/28952.html) operation. You can call this operation to add, modify, or delete the description of a CMK.
//
// @param request - UpdateKeyDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKeyDescriptionResponse
func (client *Client) UpdateKeyDescriptionWithOptions(request *UpdateKeyDescriptionRequest, runtime *util.RuntimeOptions) (_result *UpdateKeyDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateKeyDescription"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateKeyDescriptionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateKeyDescriptionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 调用UpdateKeyDescription接口更新主密钥的描述信息。
//
// Description:
//
// This operation replaces the description of a customer master key (CMK) with the description that you specify. The original description of the CMK is specified by the Description parameter when you call the [DescribeKey](https://help.aliyun.com/document_detail/28952.html) operation. You can call this operation to add, modify, or delete the description of a CMK.
//
// @param request - UpdateKeyDescriptionRequest
//
// @return UpdateKeyDescriptionResponse
func (client *Client) UpdateKeyDescription(request *UpdateKeyDescriptionRequest) (_result *UpdateKeyDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateKeyDescriptionResponse{}
	_body, _err := client.UpdateKeyDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the virtual private cloud (VPC) that is associated with a Key Management Service (KMS) instance.
//
// Description:
//
// If your own applications are deployed in multiple VPCs in the same region, you can associate the VPCs except the VPC in which the KMS instance resides with the KMS instance. This topic describes how to configure the VPCs.
//
// The VPCs can belong to the same Alibaba Cloud account or different Alibaba Cloud accounts. After the configuration is complete, the applications in these VPCs can access the KMS instance.
//
// > If the VPCs belong to different Alibaba Cloud accounts, you must first configure resource sharing to share the vSwitches of other Alibaba Cloud accounts with the Alibaba Cloud account to which the KMS instance belongs. For more information, see [Access a KMS instance from multiple VPCs in the same region](https://help.aliyun.com/document_detail/2393236.html).
//
// @param request - UpdateKmsInstanceBindVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKmsInstanceBindVpcResponse
func (client *Client) UpdateKmsInstanceBindVpcWithOptions(request *UpdateKmsInstanceBindVpcRequest, runtime *util.RuntimeOptions) (_result *UpdateKmsInstanceBindVpcResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateKmsInstanceBindVpc"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateKmsInstanceBindVpcResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateKmsInstanceBindVpcResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Updates the virtual private cloud (VPC) that is associated with a Key Management Service (KMS) instance.
//
// Description:
//
// If your own applications are deployed in multiple VPCs in the same region, you can associate the VPCs except the VPC in which the KMS instance resides with the KMS instance. This topic describes how to configure the VPCs.
//
// The VPCs can belong to the same Alibaba Cloud account or different Alibaba Cloud accounts. After the configuration is complete, the applications in these VPCs can access the KMS instance.
//
// > If the VPCs belong to different Alibaba Cloud accounts, you must first configure resource sharing to share the vSwitches of other Alibaba Cloud accounts with the Alibaba Cloud account to which the KMS instance belongs. For more information, see [Access a KMS instance from multiple VPCs in the same region](https://help.aliyun.com/document_detail/2393236.html).
//
// @param request - UpdateKmsInstanceBindVpcRequest
//
// @return UpdateKmsInstanceBindVpcResponse
func (client *Client) UpdateKmsInstanceBindVpc(request *UpdateKmsInstanceBindVpcRequest) (_result *UpdateKmsInstanceBindVpcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateKmsInstanceBindVpcResponse{}
	_body, _err := client.UpdateKmsInstanceBindVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an access control rule.
//
// Description:
//
// - You can update only private IP addresses and description of an access control rule. You cannot update the name and network type of an access control rule.
//
// - Updating an access control rule affects all permission policies that are bound to the access control rule. Exercise caution when you perform this operation.
//
// @param request - UpdateNetworkRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNetworkRuleResponse
func (client *Client) UpdateNetworkRuleWithOptions(request *UpdateNetworkRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateNetworkRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SourcePrivateIp)) {
		query["SourcePrivateIp"] = request.SourcePrivateIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateNetworkRule"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateNetworkRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateNetworkRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Updates an access control rule.
//
// Description:
//
// - You can update only private IP addresses and description of an access control rule. You cannot update the name and network type of an access control rule.
//
// - Updating an access control rule affects all permission policies that are bound to the access control rule. Exercise caution when you perform this operation.
//
// @param request - UpdateNetworkRuleRequest
//
// @return UpdateNetworkRuleResponse
func (client *Client) UpdateNetworkRule(request *UpdateNetworkRuleRequest) (_result *UpdateNetworkRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateNetworkRuleResponse{}
	_body, _err := client.UpdateNetworkRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新一个权限策略
//
// Description:
//
// - You can update the role-based access control (RBAC) permissions, accessible resources, access control rules, and description of a permission policy. You cannot update the name or scope of a permission policy.
//
// - Updating a permission policy affects all application access points (AAPs) that are bound to the permission policy. Exercise caution when you perform this operation.
//
// @param request - UpdatePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePolicyResponse
func (client *Client) UpdatePolicyWithOptions(request *UpdatePolicyRequest, runtime *util.RuntimeOptions) (_result *UpdatePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessControlRules)) {
		query["AccessControlRules"] = request.AccessControlRules
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Permissions)) {
		query["Permissions"] = request.Permissions
	}

	if !tea.BoolValue(util.IsUnset(request.Resources)) {
		query["Resources"] = request.Resources
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePolicy"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdatePolicyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdatePolicyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新一个权限策略
//
// Description:
//
// - You can update the role-based access control (RBAC) permissions, accessible resources, access control rules, and description of a permission policy. You cannot update the name or scope of a permission policy.
//
// - Updating a permission policy affects all application access points (AAPs) that are bound to the permission policy. Exercise caution when you perform this operation.
//
// @param request - UpdatePolicyRequest
//
// @return UpdatePolicyResponse
func (client *Client) UpdatePolicy(request *UpdatePolicyRequest) (_result *UpdatePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePolicyResponse{}
	_body, _err := client.UpdatePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// When automatic key rotation is enabled, KMS automatically creates a key version after the preset rotation period arrives. In addition, KMS sets the new key version as the primary key version.
//
// An automatic key rotation policy cannot be configured for the following keys:
//
// 	- Asymmetric key
//
// 	- Service-managed key
//
// 	- Bring your own key (BYOK) that is imported into KMS
//
// 	- Key that is not in the **Enabled*	- state
//
// In this example, automatic key rotation is enabled for a CMK whose ID is `1234abcd-12ab-34cd-56ef-12345678****`. The automatic rotation period is 30 days.
//
// @param request - UpdateRotationPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRotationPolicyResponse
func (client *Client) UpdateRotationPolicyWithOptions(request *UpdateRotationPolicyRequest, runtime *util.RuntimeOptions) (_result *UpdateRotationPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnableAutomaticRotation)) {
		query["EnableAutomaticRotation"] = request.EnableAutomaticRotation
	}

	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.RotationInterval)) {
		query["RotationInterval"] = request.RotationInterval
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRotationPolicy"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateRotationPolicyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateRotationPolicyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// When automatic key rotation is enabled, KMS automatically creates a key version after the preset rotation period arrives. In addition, KMS sets the new key version as the primary key version.
//
// An automatic key rotation policy cannot be configured for the following keys:
//
// 	- Asymmetric key
//
// 	- Service-managed key
//
// 	- Bring your own key (BYOK) that is imported into KMS
//
// 	- Key that is not in the **Enabled*	- state
//
// In this example, automatic key rotation is enabled for a CMK whose ID is `1234abcd-12ab-34cd-56ef-12345678****`. The automatic rotation period is 30 days.
//
// @param request - UpdateRotationPolicyRequest
//
// @return UpdateRotationPolicyResponse
func (client *Client) UpdateRotationPolicy(request *UpdateRotationPolicyRequest) (_result *UpdateRotationPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRotationPolicyResponse{}
	_body, _err := client.UpdateRotationPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the metadata of a secret.
//
// Description:
//
// In this example, the metadata of the `secret001` secret is updated. The `Description` parameter is set to `datainfo`.
//
// @param request - UpdateSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecretResponse
func (client *Client) UpdateSecretWithOptions(request *UpdateSecretRequest, runtime *util.RuntimeOptions) (_result *UpdateSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.SecretName)) {
		query["SecretName"] = request.SecretName
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendedConfig)) {
		query["ExtendedConfig"] = request.ExtendedConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSecret"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateSecretResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateSecretResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Updates the metadata of a secret.
//
// Description:
//
// In this example, the metadata of the `secret001` secret is updated. The `Description` parameter is set to `datainfo`.
//
// @param request - UpdateSecretRequest
//
// @return UpdateSecretResponse
func (client *Client) UpdateSecret(request *UpdateSecretRequest) (_result *UpdateSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSecretResponse{}
	_body, _err := client.UpdateSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// After automatic rotation is enabled, Secrets Manager schedules the first automatic rotation by adding the preset rotation interval to the timestamp of the last rotation.
//
// Limits: The UpdateSecretRotationPolicy operation cannot be used to update the rotation policy of generic secrets.
//
// In this example, the rotation policy of the `RdsSecret/Mysql5.4/MyCred` secret is updated. The following settings are modified:
//
// 	- The `EnableAutomaticRotation` parameter is set to `true`, which indicates that automatic rotation is enabled.
//
// 	- The `RotationInterval` parameter is set to `30d`, which indicates that the interval for automatic rotation is 30 days.
//
// @param request - UpdateSecretRotationPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecretRotationPolicyResponse
func (client *Client) UpdateSecretRotationPolicyWithOptions(request *UpdateSecretRotationPolicyRequest, runtime *util.RuntimeOptions) (_result *UpdateSecretRotationPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnableAutomaticRotation)) {
		query["EnableAutomaticRotation"] = request.EnableAutomaticRotation
	}

	if !tea.BoolValue(util.IsUnset(request.RotationInterval)) {
		query["RotationInterval"] = request.RotationInterval
	}

	if !tea.BoolValue(util.IsUnset(request.SecretName)) {
		query["SecretName"] = request.SecretName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSecretRotationPolicy"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateSecretRotationPolicyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateSecretRotationPolicyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// After automatic rotation is enabled, Secrets Manager schedules the first automatic rotation by adding the preset rotation interval to the timestamp of the last rotation.
//
// Limits: The UpdateSecretRotationPolicy operation cannot be used to update the rotation policy of generic secrets.
//
// In this example, the rotation policy of the `RdsSecret/Mysql5.4/MyCred` secret is updated. The following settings are modified:
//
// 	- The `EnableAutomaticRotation` parameter is set to `true`, which indicates that automatic rotation is enabled.
//
// 	- The `RotationInterval` parameter is set to `30d`, which indicates that the interval for automatic rotation is 30 days.
//
// @param request - UpdateSecretRotationPolicyRequest
//
// @return UpdateSecretRotationPolicyResponse
func (client *Client) UpdateSecretRotationPolicy(request *UpdateSecretRotationPolicyRequest) (_result *UpdateSecretRotationPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSecretRotationPolicyResponse{}
	_body, _err := client.UpdateSecretRotationPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// UpdateSecretVersionStage
//
// Description:
//
// Updates the stage label that marks a secret version.
//
// @param request - UpdateSecretVersionStageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecretVersionStageResponse
func (client *Client) UpdateSecretVersionStageWithOptions(request *UpdateSecretVersionStageRequest, runtime *util.RuntimeOptions) (_result *UpdateSecretVersionStageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MoveToVersion)) {
		query["MoveToVersion"] = request.MoveToVersion
	}

	if !tea.BoolValue(util.IsUnset(request.RemoveFromVersion)) {
		query["RemoveFromVersion"] = request.RemoveFromVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SecretName)) {
		query["SecretName"] = request.SecretName
	}

	if !tea.BoolValue(util.IsUnset(request.VersionStage)) {
		query["VersionStage"] = request.VersionStage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSecretVersionStage"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateSecretVersionStageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateSecretVersionStageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// UpdateSecretVersionStage
//
// Description:
//
// Updates the stage label that marks a secret version.
//
// @param request - UpdateSecretVersionStageRequest
//
// @return UpdateSecretVersionStageResponse
func (client *Client) UpdateSecretVersionStage(request *UpdateSecretVersionStageRequest) (_result *UpdateSecretVersionStageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSecretVersionStageResponse{}
	_body, _err := client.UpdateSecretVersionStageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// In this example, a certificate issued by a CA is imported into Certificates Manager. The ID of the certificate in Certificates Manager is `12345678-1234-1234-1234-12345678****`.
//
// @param request - UploadCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadCertificateResponse
func (client *Client) UploadCertificateWithOptions(request *UploadCertificateRequest, runtime *util.RuntimeOptions) (_result *UploadCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Certificate)) {
		query["Certificate"] = request.Certificate
	}

	if !tea.BoolValue(util.IsUnset(request.CertificateChain)) {
		query["CertificateChain"] = request.CertificateChain
	}

	if !tea.BoolValue(util.IsUnset(request.CertificateId)) {
		query["CertificateId"] = request.CertificateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadCertificate"),
		Version:     tea.String("2016-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UploadCertificateResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UploadCertificateResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// In this example, a certificate issued by a CA is imported into Certificates Manager. The ID of the certificate in Certificates Manager is `12345678-1234-1234-1234-12345678****`.
//
// @param request - UploadCertificateRequest
//
// @return UploadCertificateResponse
func (client *Client) UploadCertificate(request *UploadCertificateRequest) (_result *UploadCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadCertificateResponse{}
	_body, _err := client.UploadCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
