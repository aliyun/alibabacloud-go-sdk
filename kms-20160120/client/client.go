// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AsymmetricDecryptRequest struct {
	// The decryption algorithm.
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ciphertext that you want to decrypt.
	//
	// >
	// *   The value is encoded in Base64.
	// *   You can call the [AsymmetricEncrypt](~~148131~~) operation to generate the ciphertext.
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The ID of the customer master key (CMK). The ID must be globally unique.
	//
	// >  You can also set this parameter to an alias that is bound to the CMK. For more information, see [Alias overview](~~68522~~).
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The version ID of the CMK. The ID must be globally unique.
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
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The version ID of the CMK that is used to encrypt the plaintext.
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The Base64-encoded plaintext that is generated after decryption.
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AsymmetricDecryptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// >  You can also set this parameter to an alias that is bound to the CMK. For more information, see [Overview of aliases](~~68522~~).
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The version ID of the CMK. The ID must be globally unique.
	//
	// >  You can call the [ListKeyVersions](~~133966~~) operation to query the versions of a CMK. The ID of a version is specified by the KeyVersionId parameter.
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The plaintext that you want to encrypt. The plaintext must be Base64-encoded.
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
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// >  If you set the KeyId parameter in the request to an alias, the ID of the CMK to which the alias is bound is returned.
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The version ID of the CMK that is used to encrypt the plaintext.
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AsymmetricEncryptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The signature algorithm.
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The digest that is generated for the original message by using a hash algorithm. The hash algorithm is specified by the Algorithm parameter.
	//
	// >
	// *   The value is encoded in Base64.
	// *   For more information about how to calculate message digests, see the **Preprocess signature: compute a message digest** section of the [Generate and verify a signature by using an asymmetric CMK](~~148146~~) topic.
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The ID of the customer master key (CMK). The ID must be globally unique.
	//
	// >  You can also set this parameter to an alias that is bound to the CMK. For more information, see [Alias overview](~~68522~~).
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The version ID of the CMK. The ID must be globally unique.
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
	// The ID of the CMK. The ID must be globally unique.
	//
	// >  If you set the KeyId parameter in the request to an alias, the ID of the CMK to which the alias is bound is returned.
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The version ID of the CMK. The ID must be globally unique.
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The calculated signature.
	//
	// >  The value is encoded in Base64.
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AsymmetricSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The digest that is generated for the original message by using a hash algorithm. The hash algorithm is specified by the **Algorithm** parameter.
	//
	// >  The value is encoded in Base64.
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// >  You can also set this parameter to an alias that is bound to the CMK. For more information, see [Overview of aliases](~~68522~~).
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The version ID of the CMK. The ID must be globally unique.
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The signature value to be verified.
	//
	// >  The value is encoded in Base64.
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
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The version ID of the CMK that is used to encrypt the plaintext.
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the signature passed the verification.
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AsymmetricVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelKeyDeletionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// *   RSAES_OAEP_SHA\_1
	//
	// *   RSAES_OAEP_SHA\_256
	//
	// *   SM2PKE
	//
	// > The SM2PKE encryption algorithm is supported only in regions in mainland China. In these regions, managed hardware security modules (HSMs) are used. For more information, see [Managed HSM overview](~~125803~~).
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ID of the certificate. The ID must be globally unique in Certificates Manager.
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The data that you want to decrypt.
	//
	// The value is encoded in Base64.
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
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The plaintext after data is decrypted.
	//
	// The value is encoded in Base64.
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CertificatePrivateKeyDecryptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// *   RSA_PKCS1\_SHA\_256
	//
	// *   RSA_PSS_SHA\_256
	//
	// *   ECDSA_SHA\_256
	//
	// *   SM2DSA
	//
	// >* The SM2DSA signature algorithm is supported only in regions where managed hardware security modules (HSMs) are used in mainland China. For more information, see [Managed HSM overview](~~125803~~).
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ID of the certificate. The ID must be globally unique in Certificates Manager.
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The data to be signed.
	//
	// The value is encoded in Base64. For example, if the hexadecimal data that you want to sign is `[0x31, 0x32, 0x33, 0x34]`, the Base64-encoded data is `MTIzNA==`.
	//
	// If the MessageType parameter is set to RAW, the size of the data must be less than or equal to 4 KB.
	//
	// If the size of the data is greater than 4 KB, you can set the MessageType parameter to DIGEST and set the Message parameter to the digest of the data. The digest is also called hash value. You can compute the digest of the data on an on-premises machine. Certificates Manager uses the digest that you compute in your own certificate application system. The message digest algorithm that you use must match the specified signature algorithm. Comply with the following mapping between signature algorithms and message digest algorithms:
	//
	// *   If the signature algorithm is RSA_PKCS1\_SHA\_256, RSA_PSS_SHA\_256, or ECDSA_SHA\_256, the message digest algorithm must be SHA-256.
	// *   If the signature algorithm is SM2DSA, the message digest algorithm must be SM3.
	//
	// >  If the key type of the certificate is EC_SM2 and the MessageType parameter is set to DIGEST, the value of the Message parameter is `e` that is described in GB/T 32918.2-2016 6.1.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The type of the message. Valid values:
	//
	// *   RAW: the raw data. This is the default value.
	// *   DIGEST: the message digest (hash value) of the raw data.
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
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The signature value.
	//
	// The value is encoded in Base64.
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CertificatePrivateKeySignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// *   RSAES_OAEP_SHA\_1
	//
	// *   RSAES_OAEP_SHA\_256
	//
	// *   SM2PKE
	//
	// >The SM2PKE encryption algorithm is supported only in regions in mainland China. In these regions, managed hardware security modules (HSMs) are used. For more information, see [Managed HSM overview](~~125803~~).
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ID of the certificate. The ID must be globally unique in Certificates Manager.
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The data that you want to encrypt.
	//
	// The value is encoded in Base64. For example, if the hexadecimal data that you want to encrypt is `[0x31, 0x32, 0x33, 0x34]`, the Base64-encoded data is `MTIzNA==`.
	//
	// The size of data that can be encrypted varies based on the encryption algorithm that you use:
	//
	// *   RSAES_OAEP_SHA\_1: 214 bytes
	// *   RSAES_OAEP_SHA\_256: 190 bytes
	// *   SM2PKE: 6,047 bytes
	//
	// If the size of data that you want to encrypt exceeds the preceding limits, you can call the [GenerateDataKey](~~28948~~) operation to generate a data key to encrypt the data. Then, call the CertificatePublicKeyEncrypt operation to encrypt the data key.
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
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The ciphertext.
	//
	// The value is encoded in Base64.
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CertificatePublicKeyEncryptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// *   RSA_PKCS1\_SHA\_256
	//
	// *   RSA_PSS_SHA\_256
	//
	// *   ECDSA_SHA\_256
	//
	// *   SM2DSA
	//
	// > The SM2DSA signature algorithm is supported only in regions where managed hardware security modules (HSMs) are used in the Chinese mainland. For more information, see [Managed HSM overview](~~125803~~).
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ID of the certificate. The ID must be globally unique in Certificates Manager.
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The raw data that is signed.
	//
	// The value is encoded in Base64. For example, if the raw data in the hexadecimal format is `[0x31, 0x32, 0x33, 0x34]`, set this parameter to the Base64-encoded value `MTIzNA==`.
	//
	// If the MessageType parameter is set to RAW, the size of the data must be less than or equal to 4 KB.
	//
	// If the size of the data is greater than 4 KB, you can set the MessageType parameter to DIGEST and set the Message parameter to the digest of the data. The digest is also called hash value. You can compute the digest of the data on an on-premises device. Certificates Manager uses the digest that you compute in your own certificate application system. The message digest algorithm that you use must match the specified signature algorithm. Comply with the following mapping between signature algorithms and message digest algorithms:
	//
	// *   If the signature algorithm is RSA_PKCS1\_SHA\_256, RSA_PSS_SHA\_256, or ECDSA_SHA\_256, the message digest algorithm must be SHA-256.
	// *   If the signature algorithm is SM2DSA, the message digest algorithm must be SM3.
	//
	// >  If the key type of the certificate is EC_SM2 and the MessageType parameter is set to DIGEST, the value of the Message parameter is `e` that is described in GB/T 32918.2-2016 6.1.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The type of the message. Valid values:
	//
	// *   RAW: the raw data. This is the default value.
	// *   DIGEST: the message digest (hash value) of the raw data.
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The signature value.
	//
	// The value is encoded in Base64.
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
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The verification result. Valid values:
	//
	// *   true: The signature is valid.
	// *   false: The signature is invalid.
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CertificatePublicKeyVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateAliasRequest struct {
	// The alias of the CMK.
	//
	// The alias must be 1 to 255 characters in length and must contain the prefix `alias/`. The alias cannot be prefixed with the reserved word `alias/acs`.
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAliasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateCertificateRequest struct {
	// Specifies whether the private key of the certificate can be exported for use. Valid values:
	//
	// *   true: The private key of the certificate can be exported for use. This is the default value.
	// *   false: The private key of the certificate cannot be exported for use. We recommend that you set this parameter to false to protect keys with a higher security level.
	ExportablePrivateKey *bool `json:"ExportablePrivateKey,omitempty" xml:"ExportablePrivateKey,omitempty"`
	// The type of the key. Valid values:
	//
	// *   RSA\_2048
	// *   EC_P256
	// *   EC_SM2
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The certificate subject, which is the owner of the certificate.
	//
	// Specify the value in the distinguished name (DN) format, as defined in [RFC 2253](https://tools.ietf.org/html/rfc2253?spm=a2c4g.11186623.2.13.265f1a1cGFCn3Q). A DN is a sequence of relative distinguished names (RDNs).
	//
	// RDNs are key-value pairs in the format of `attribute1=value1,attribute2=value2`. Separate multiple RDNs with commas (,).
	//
	// The Subject parameter consists of the following fields:
	//
	// *   CN: required. The name of the certificate subject.
	// *   C: required. The two-character country or region code in the [ISO 3166-1](https://www.iso.org/obp/ui/#search/code/) standard. For example, CN indicates China.
	// *   O: required. The legal name of the enterprise, company, organization, or institution.
	// *   OU: required. The name of the department.
	// *   ST: optional. The name of the province, municipality, autonomous region, or special administrative region.
	// *   L: optional. The name of the city.
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// The subject alternative names.
	//
	// A domain name list is supported. A maximum of 10 domain names are supported.
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
	// *   true: The private key of the certificate can be exported for use. This is the default value.
	// *   false: The private key of the certificate cannot be exported for use. We recommend that you set this parameter to false to protect keys with a higher security level.
	ExportablePrivateKey *bool `json:"ExportablePrivateKey,omitempty" xml:"ExportablePrivateKey,omitempty"`
	// The type of the key. Valid values:
	//
	// *   RSA\_2048
	// *   EC_P256
	// *   EC_SM2
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The certificate subject, which is the owner of the certificate.
	//
	// Specify the value in the distinguished name (DN) format, as defined in [RFC 2253](https://tools.ietf.org/html/rfc2253?spm=a2c4g.11186623.2.13.265f1a1cGFCn3Q). A DN is a sequence of relative distinguished names (RDNs).
	//
	// RDNs are key-value pairs in the format of `attribute1=value1,attribute2=value2`. Separate multiple RDNs with commas (,).
	//
	// The Subject parameter consists of the following fields:
	//
	// *   CN: required. The name of the certificate subject.
	// *   C: required. The two-character country or region code in the [ISO 3166-1](https://www.iso.org/obp/ui/#search/code/) standard. For example, CN indicates China.
	// *   O: required. The legal name of the enterprise, company, organization, or institution.
	// *   OU: required. The name of the department.
	// *   ST: optional. The name of the province, municipality, autonomous region, or special administrative region.
	// *   L: optional. The name of the city.
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// The subject alternative names.
	//
	// A domain name list is supported. A maximum of 10 domain names are supported.
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
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the certificate. It is the globally unique identifier (GUID) of the certificate in Certificates Manager.
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The CSR in the PEM format.
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The ID of the request.
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateKeyRequest struct {
	// The ID of the dedicated KMS instance.
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// The description of the CMK.
	//
	// The description can be 0 to 8,192 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable automatic key rotation. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: false.
	//
	// >  If the Origin parameter is set to EXTERNAL or the KeySpec parameter is set to an asymmetric CMK type, automatic key rotation is not supported.
	EnableAutomaticRotation *bool `json:"EnableAutomaticRotation,omitempty" xml:"EnableAutomaticRotation,omitempty"`
	// The type of the CMK. Valid values:
	//
	// *   Aliyun_AES\_256
	// *   Aliyun_AES\_128
	// *   Aliyun_AES\_192
	// *   Aliyun_SM4
	// *   RSA\_2048
	// *   RSA\_3072
	// *   EC_P256
	// *   EC_P256K
	// *   EC_SM2
	//
	// >
	// *   The default type of the CMK is Aliyun_AES\_256.
	// *   Only Dedicated KMS supports Aliyun_AES\_128 and Aliyun_AES\_192.
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The usage of the CMK. Valid values:
	//
	// *   ENCRYPT/DECRYPT: encrypts or decrypts data.
	// *   SIGN/VERIFY: generates or verifies a digital signature.
	//
	// If the CMK supports signature verification, the default value is SIGN/VERIFY. If the CMK does not support signature verification, the default value is ENCRYPT/DECRYPT.
	KeyUsage *string `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty"`
	// The source of key material. Valid values:
	//
	// *   Aliyun_KMS (default value)
	// *   EXTERNAL
	//
	// >
	// *   The value of this parameter is case-sensitive.
	// *   If you set the KeySpec parameter to an asymmetric CMK type, you are not allowed to set the Origin parameter to EXTERNAL.
	// *   If you set the Origin parameter to EXTERNAL, you must import key material. For more information, see [Import key material](~~68523~~).
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The protection level of the CMK. Valid values:
	//
	// *   SOFTWARE
	// *   HSM
	//
	// Default value: SOFTWARE.
	//
	// >
	// *   The value of this parameter is case-sensitive.
	// *   Assume that you set this parameter to HSM. If you set the Origin parameter to Aliyun_KMS, the CMK is created in a managed HSM. If you set the Origin parameter to EXTERNAL, you can import an external key into the managed HSM.
	ProtectionLevel *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	// The period of automatic key rotation. Specify the value in the integer\[unit] format. Unit: d (day), h (hour), m (minute), or s (second). For example, you can use either 7d or 604800s to specify a seven-day period. The period can range from 7 days to 730 days.
	//
	// >  If you set the EnableAutomaticRotation parameter to true, you must also specify this parameter. If you set the EnableAutomaticRotation parameter to false, you can leave this parameter unspecified.
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
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

func (s *CreateKeyRequest) SetProtectionLevel(v string) *CreateKeyRequest {
	s.ProtectionLevel = &v
	return s
}

func (s *CreateKeyRequest) SetRotationInterval(v string) *CreateKeyRequest {
	s.RotationInterval = &v
	return s
}

type CreateKeyResponseBody struct {
	// The metadata of the CMK.
	KeyMetadata *CreateKeyResponseBodyKeyMetadata `json:"KeyMetadata,omitempty" xml:"KeyMetadata,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	// The Alibaba Cloud Resource Name (ARN) of the CMK.
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// Indicates whether automatic key rotation is enabled. Valid values:
	//
	// *   Enabled: Automatic key rotation is enabled.
	// *   Disabled: Automatic key rotation is disabled.
	// *   Suspended: Automatic key rotation is suspended. For more information, see [Automatic key rotation](~~134270~~).
	//
	// >  Automatic key rotation is available only for symmetric CMKs.
	AutomaticRotation *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	// The date and time when the CMK was created. The time is displayed in UTC.
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// The creator of the CMK.
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The ID of the dedicated KMS instance.
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// The time when the CMK is scheduled for deletion.
	//
	// For more information, see [ScheduleKeyDeletion](~~44196~~).
	//
	// >  This value is returned only when the value of the KeyState parameter is PendingDeletion.
	DeleteDate *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
	// The description of the CMK.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The type of the CMK.
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The status of the CMK.
	//
	// For more information, see [Impact of CMK status on API operations](~~44211~~).
	KeyState *string `json:"KeyState,omitempty" xml:"KeyState,omitempty"`
	// The usage of the CMK.
	KeyUsage *string `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty"`
	// The time when the last rotation was performed. The time is displayed in UTC.
	//
	// For a new CMK, this parameter value is the time when the initial version of the CMK was generated.
	LastRotationDate *string `json:"LastRotationDate,omitempty" xml:"LastRotationDate,omitempty"`
	// The time when the key material expires. The time is displayed in UTC.
	//
	// If this parameter value is empty, the key material does not expire.
	MaterialExpireTime *string `json:"MaterialExpireTime,omitempty" xml:"MaterialExpireTime,omitempty"`
	// The time when the next rotation will be performed.
	//
	// >  This value is returned only when the value of the AutomaticRotation parameter is Enabled or Suspended.
	NextRotationDate *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	// The source of the key material for the CMK.
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The ID of the current primary key version of the symmetric CMK.
	//
	// >
	// *   The primary key version of a symmetric CMK is an active encryption key. KMS uses the primary key version of a specified CMK to encrypt data.
	// *   This parameter is unavailable for asymmetric CMKs.
	PrimaryKeyVersion *string `json:"PrimaryKeyVersion,omitempty" xml:"PrimaryKeyVersion,omitempty"`
	// The protection level of the CMK.
	ProtectionLevel *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	// The period of automatic key rotation. Unit: seconds. The value is in the format of an integer followed by the letter s. For example, if the rotation period is seven days, this parameter is set to 604800s. This value is returned only when the value of the AutomaticRotation parameter is Enabled or Suspended.
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// >  You can also set the value to an alias that is bound to the CMK. For more information, see [Overview of aliases](~~68522~~).
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
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the version.
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateKeyVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateSecretRequest struct {
	// The ID of the dedicated KMS instance.
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// The description of the secret.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable automatic rotation. Valid values:
	//
	// *   true: specifies to enable automatic rotation.
	// *   false: specifies to disable automatic rotation. This is the default value.
	//
	// >  This parameter is valid if you set the SecretType parameter to Rds, RAMCredentials, or ECS.
	EnableAutomaticRotation *bool `json:"EnableAutomaticRotation,omitempty" xml:"EnableAutomaticRotation,omitempty"`
	// The ID of the CMK that is used to encrypt the secret value.
	//
	// If the DKMSInstanceId parameter is empty, Secrets Manager uses a CMK that is created by Dedicated KMS to encrypt and protect secrets. If the DKMSInstanceId parameter is not empty, specify the CMK of the dedicated KMS instance to encrypt and protect secrets.
	//
	// >  The CMK must be a symmetric CMK.
	EncryptionKeyId *string `json:"EncryptionKeyId,omitempty" xml:"EncryptionKeyId,omitempty"`
	// The extended configuration of the secret. This parameter specifies the properties of the secret of the specific type. The description can be up to 1,024 characters in length.
	//
	// *   If you set the SecretType parameter to Generic, you do not need to configure this parameter.
	//
	// *   If you set the SecretType parameter to Rds, configure the following fields for the ExtendedConfig parameter:
	//
	//     *   SecretSubType: required. The subtype of the secret. Valid values:
	//
	//         *   SingleUser: Secrets Manager manages the ApsaraDB RDS secret in single-account mode. When the secret is rotated, the password of the specified account is reset to a new random password.
	//         *   DoubleUsers: Secrets Manager manages the ApsaraDB RDS secret in dual-account mode. One account is referenced by the ACSCurrent version, and the other account is referenced by the ACSPrevious version. When the secret is rotated, the password of the account referenced by the ACSPrevious version is reset to a new random password. Then, Secrets Manager switches the referenced accounts between the ACSCurrent and ACSPrevious versions.
	//
	//     *   DBInstanceId: required. The ApsaraDB RDS instance to which the ApsaraDB RDS account belongs.
	//
	//     *   CustomData: optional. The custom data. The value is a collection of key-value pairs in the JSON format. Up to 10 key-value pairs can be specified. Separate multiple key-value pairs with commas (,). Example: `{"Key1": "v1", "fds":"fdsf"}`. The default value is a pair of empty braces (`{}`).
	//
	// *   If you set the SecretType parameter to RAMCredentials, configure the following fields for the ExtendedConfig parameter:
	//
	//     *   SecretSubType: required. The subtype of the secret. Set the value to RamUserAccessKey.
	//     *   UserName: required. The name of the RAM user.
	//     *   CustomData: optional. The custom data. The value is a collection of key-value pairs in the JSON format. Up to 10 key-value pairs can be specified. Separate multiple key-value pairs with commas (,). The default value is a pair of empty braces (`{}`).
	//
	// *   If you set the SecretType parameter to ECS, configure the following fields for the ExtendedConfig parameter:
	//
	//     *   SecretSubType: required. The subtype of the secret. Valid values:
	//
	//         *   Password: the password that is used to log on to the ECS instance.
	//         *   SSHKey: the SSH public key and private key that are used to log on to the ECS instance.
	//
	//     *   RegionId: required. The ID of the region in which the ECS instance resides.
	//
	//     *   InstanceId: required. The ID of the ECS instance.
	//
	//     *   CustomData: optional. The custom data. The value is a collection of key-value pairs in the JSON format. Up to 10 key-value pairs can be specified. Separate multiple key-value pairs with commas (,). The default value is a pair of empty braces (`{}`).
	//
	// >  This parameter is required if you set the SecretType parameter to Rds, RAMCredentials, or ECS.
	ExtendedConfig map[string]interface{} `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	// The interval for automatic rotation. Valid values: 6 hours to 8,760 hours (365 days).
	//
	// The value is in the `integer[unit]` format.
	//
	// The unit can be d (day), h (hour), m (minute), or s (second). For example, both 7d and 604800s indicate a seven-day interval.
	//
	// >  This parameter is required if you set the EnableAutomaticRotation parameter to true. This parameter is ignored if you set the EnableAutomaticRotation parameter to false or if the EnableAutomaticRotation parameter is not configured.
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	// The value of the secret that you want to create. Secrets Manager encrypts the secret value and stores the encrypted value in the initial version.
	//
	// *   If you set the SecretType parameter to Generic that indicates a generic secret, you can customize the secret value.
	//
	// *   If you set the SecretType parameter to Rds that indicates a managed ApsaraDB RDS secret, the secret value must be in the format of `{"Accounts":[{"AccountName":"","AccountPassword":""}]}`. In the preceding format, `AccountName` indicates the username of the account that is used to connect to your ApsaraDB RDS instance, and `AccountPassword` specifies the password of the account.
	//
	// *   If you set the SecretType parameter to RAMCredentials that indicates a managed RAM secret, the secret value must be in the format of `{"AccessKeys":[{"AccessKeyId":"","AccessKeySecret":"",}]}`. In the preceding format, `AccessKeyId` indicates the AccessKey ID of the RAM user and `AccessKeySecret` specifies the AccessKey secret of the RAM user. You must specify all the AccessKey pairs of the RAM user.
	//
	// *   If you set the SecretType parameter to ECS that indicates a managed ECS secret, the secret value must be in one of the following formats:
	//
	//     *   `{"UserName":"","Password": ""}`: In the format, `UserName` specifies the username that is used to log on to the ECS instance, and `Password` specifies the password that is used to log on to the ECS instance.
	//     *   `{"UserName":"","PublicKey": "", "PrivateKey": ""}`: In the format, `PublicKey` indicates the SSH public key that is used to log on to the ECS instance, and `PrivateKey` specifies the SSH private key that is used to log on to the ECS instance.
	SecretData *string `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
	// The type of the secret value. Valid values:
	//
	// *   text
	// *   binary
	//
	// >  If you set the SecretType parameter to Rds, RAMCredentials, or ECS, the SecretDataType parameter must be set to text.
	SecretDataType *string `json:"SecretDataType,omitempty" xml:"SecretDataType,omitempty"`
	// The name of the secret.
	//
	// The value must be 1 to 64 characters in length and can contain letters, digits, underscores (\_), forward slashes (/), plus signs (+), equal signs (=), periods (.), hyphens (-), and at signs (@). The following list describes the name requirements for different types of secrets:
	//
	// *   If the SecretType parameter is set to Generic or Rds, the name cannot start with `acs/`.
	// *   If the SecretType parameter is set to RAMCredentials, set the SecretName parameter to `$Auto`. In this case, KMS automatically generates a secret name that starts with `acs/ram/user/`. The name includes the display name of RAM user.
	// *   If the SecretType parameter is set to ECS, the name must start with `acs/ecs/`.
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The type of the secret. Valid values:
	//
	// *   Generic: specifies a generic secret.
	// *   Rds: specifies a managed ApsaraDB RDS secret.
	// *   RAMCredentials: specifies a managed RAM secret.
	// *   ECS: specifies a managed ECS secret.
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// The tags of the secret.
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The initial version number. Version numbers are unique in each secret.
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
	// The ID of the dedicated KMS instance.
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// The description of the secret.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable automatic rotation. Valid values:
	//
	// *   true: specifies to enable automatic rotation.
	// *   false: specifies to disable automatic rotation. This is the default value.
	//
	// >  This parameter is valid if you set the SecretType parameter to Rds, RAMCredentials, or ECS.
	EnableAutomaticRotation *bool `json:"EnableAutomaticRotation,omitempty" xml:"EnableAutomaticRotation,omitempty"`
	// The ID of the CMK that is used to encrypt the secret value.
	//
	// If the DKMSInstanceId parameter is empty, Secrets Manager uses a CMK that is created by Dedicated KMS to encrypt and protect secrets. If the DKMSInstanceId parameter is not empty, specify the CMK of the dedicated KMS instance to encrypt and protect secrets.
	//
	// >  The CMK must be a symmetric CMK.
	EncryptionKeyId *string `json:"EncryptionKeyId,omitempty" xml:"EncryptionKeyId,omitempty"`
	// The extended configuration of the secret. This parameter specifies the properties of the secret of the specific type. The description can be up to 1,024 characters in length.
	//
	// *   If you set the SecretType parameter to Generic, you do not need to configure this parameter.
	//
	// *   If you set the SecretType parameter to Rds, configure the following fields for the ExtendedConfig parameter:
	//
	//     *   SecretSubType: required. The subtype of the secret. Valid values:
	//
	//         *   SingleUser: Secrets Manager manages the ApsaraDB RDS secret in single-account mode. When the secret is rotated, the password of the specified account is reset to a new random password.
	//         *   DoubleUsers: Secrets Manager manages the ApsaraDB RDS secret in dual-account mode. One account is referenced by the ACSCurrent version, and the other account is referenced by the ACSPrevious version. When the secret is rotated, the password of the account referenced by the ACSPrevious version is reset to a new random password. Then, Secrets Manager switches the referenced accounts between the ACSCurrent and ACSPrevious versions.
	//
	//     *   DBInstanceId: required. The ApsaraDB RDS instance to which the ApsaraDB RDS account belongs.
	//
	//     *   CustomData: optional. The custom data. The value is a collection of key-value pairs in the JSON format. Up to 10 key-value pairs can be specified. Separate multiple key-value pairs with commas (,). Example: `{"Key1": "v1", "fds":"fdsf"}`. The default value is a pair of empty braces (`{}`).
	//
	// *   If you set the SecretType parameter to RAMCredentials, configure the following fields for the ExtendedConfig parameter:
	//
	//     *   SecretSubType: required. The subtype of the secret. Set the value to RamUserAccessKey.
	//     *   UserName: required. The name of the RAM user.
	//     *   CustomData: optional. The custom data. The value is a collection of key-value pairs in the JSON format. Up to 10 key-value pairs can be specified. Separate multiple key-value pairs with commas (,). The default value is a pair of empty braces (`{}`).
	//
	// *   If you set the SecretType parameter to ECS, configure the following fields for the ExtendedConfig parameter:
	//
	//     *   SecretSubType: required. The subtype of the secret. Valid values:
	//
	//         *   Password: the password that is used to log on to the ECS instance.
	//         *   SSHKey: the SSH public key and private key that are used to log on to the ECS instance.
	//
	//     *   RegionId: required. The ID of the region in which the ECS instance resides.
	//
	//     *   InstanceId: required. The ID of the ECS instance.
	//
	//     *   CustomData: optional. The custom data. The value is a collection of key-value pairs in the JSON format. Up to 10 key-value pairs can be specified. Separate multiple key-value pairs with commas (,). The default value is a pair of empty braces (`{}`).
	//
	// >  This parameter is required if you set the SecretType parameter to Rds, RAMCredentials, or ECS.
	ExtendedConfigShrink *string `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	// The interval for automatic rotation. Valid values: 6 hours to 8,760 hours (365 days).
	//
	// The value is in the `integer[unit]` format.
	//
	// The unit can be d (day), h (hour), m (minute), or s (second). For example, both 7d and 604800s indicate a seven-day interval.
	//
	// >  This parameter is required if you set the EnableAutomaticRotation parameter to true. This parameter is ignored if you set the EnableAutomaticRotation parameter to false or if the EnableAutomaticRotation parameter is not configured.
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	// The value of the secret that you want to create. Secrets Manager encrypts the secret value and stores the encrypted value in the initial version.
	//
	// *   If you set the SecretType parameter to Generic that indicates a generic secret, you can customize the secret value.
	//
	// *   If you set the SecretType parameter to Rds that indicates a managed ApsaraDB RDS secret, the secret value must be in the format of `{"Accounts":[{"AccountName":"","AccountPassword":""}]}`. In the preceding format, `AccountName` indicates the username of the account that is used to connect to your ApsaraDB RDS instance, and `AccountPassword` specifies the password of the account.
	//
	// *   If you set the SecretType parameter to RAMCredentials that indicates a managed RAM secret, the secret value must be in the format of `{"AccessKeys":[{"AccessKeyId":"","AccessKeySecret":"",}]}`. In the preceding format, `AccessKeyId` indicates the AccessKey ID of the RAM user and `AccessKeySecret` specifies the AccessKey secret of the RAM user. You must specify all the AccessKey pairs of the RAM user.
	//
	// *   If you set the SecretType parameter to ECS that indicates a managed ECS secret, the secret value must be in one of the following formats:
	//
	//     *   `{"UserName":"","Password": ""}`: In the format, `UserName` specifies the username that is used to log on to the ECS instance, and `Password` specifies the password that is used to log on to the ECS instance.
	//     *   `{"UserName":"","PublicKey": "", "PrivateKey": ""}`: In the format, `PublicKey` indicates the SSH public key that is used to log on to the ECS instance, and `PrivateKey` specifies the SSH private key that is used to log on to the ECS instance.
	SecretData *string `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
	// The type of the secret value. Valid values:
	//
	// *   text
	// *   binary
	//
	// >  If you set the SecretType parameter to Rds, RAMCredentials, or ECS, the SecretDataType parameter must be set to text.
	SecretDataType *string `json:"SecretDataType,omitempty" xml:"SecretDataType,omitempty"`
	// The name of the secret.
	//
	// The value must be 1 to 64 characters in length and can contain letters, digits, underscores (\_), forward slashes (/), plus signs (+), equal signs (=), periods (.), hyphens (-), and at signs (@). The following list describes the name requirements for different types of secrets:
	//
	// *   If the SecretType parameter is set to Generic or Rds, the name cannot start with `acs/`.
	// *   If the SecretType parameter is set to RAMCredentials, set the SecretName parameter to `$Auto`. In this case, KMS automatically generates a secret name that starts with `acs/ram/user/`. The name includes the display name of RAM user.
	// *   If the SecretType parameter is set to ECS, the name must start with `acs/ecs/`.
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The type of the secret. Valid values:
	//
	// *   Generic: specifies a generic secret.
	// *   Rds: specifies a managed ApsaraDB RDS secret.
	// *   RAMCredentials: specifies a managed RAM secret.
	// *   ECS: specifies a managed ECS secret.
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// The tags of the secret.
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The initial version number. Version numbers are unique in each secret.
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
	// The Alibaba Cloud Resource Name (ARN) of the secret.
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// Indicates whether automatic rotation is enabled. Valid values:
	//
	// *   Enabled: indicates that automatic rotation is enabled.
	// *   Disabled: indicates that automatic rotation is disabled.
	// *   Invalid: indicates that the status of automatic rotation is abnormal. In this case, Secrets Manager cannot automatically rotate the secret.
	//
	// >  This parameter is returned if you set the SecretType parameter to Rds, RAMCredentials, or ECS.
	AutomaticRotation *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	// The ID of the dedicated KMS instance.
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// The extended configuration of the secret.
	//
	// >  This parameter is returned if you set the SecretType parameter to Rds, RAMCredentials, or ECS.
	ExtendedConfig *string `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	// The time when the next rotation will be performed.
	//
	// >  This parameter is returned if automatic rotation is enabled.
	NextRotationDate *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The interval for automatic rotation.
	//
	// The value is in the `integer[unit]` format. The value of the `unit` field is fixed as s. For example, if the value is 604800s, automatic rotation is performed at a 7-day interval.
	//
	// >  This parameter is returned if automatic rotation is enabled.
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	// The name of the secret.
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The type of the secret. Valid values:
	//
	// *   Generic: indicates a generic secret.
	// *   Rds: indicates a managed ApsaraDB RDS secret.
	// *   RAMCredentials: indicates a managed RAM secret.
	// *   ECS: indicates a managed ECS secret.
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// The version number of the secret.
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSecretResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// *   [GenerateDataKey](~~28948~~)
	// *   [Encrypt](~~28949~~)
	// *   [GenerateDataKeyWithoutPlaintext](~~134043~~)
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The JSON string that consists of key-value pairs.
	//
	// >  If you specify the EncryptionContext parameter when you call the [GenerateDataKey](~~28948~~), [Encrypt](~~28949~~), or [GenerateDataKeyWithoutPlaintext](~~134043~~) operation, you must specify the same context when you call the Decrypt operation. For more information, see [EncryptionContext](~~42975~~).
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
	// *   [GenerateDataKey](~~28948~~)
	// *   [Encrypt](~~28949~~)
	// *   [GenerateDataKeyWithoutPlaintext](~~134043~~)
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The JSON string that consists of key-value pairs.
	//
	// >  If you specify the EncryptionContext parameter when you call the [GenerateDataKey](~~28948~~), [Encrypt](~~28949~~), or [GenerateDataKeyWithoutPlaintext](~~134043~~) operation, you must specify the same context when you call the Decrypt operation. For more information, see [EncryptionContext](~~42975~~).
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
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the CMK version that is used to decrypt the ciphertext.
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The plaintext that is generated after decryption.
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	// The ID of the request.
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DecryptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAliasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeleteCertificateRequest struct {
	// The ID of the certificate. It is the globally unique identifier (GUID) of the certificate in Alibaba Cloud Certificate Manager.
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeleteKeyMaterialRequest struct {
	// The globally unique ID of the CMK.
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteKeyMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeleteSecretRequest struct {
	// Specifies whether to forcibly delete the secret. If this parameter is set to true, the secret cannot be recovered.
	//
	// Valid values:
	//
	// *   **true**
	// *   **false** (default value)
	ForceDeleteWithoutRecovery *string `json:"ForceDeleteWithoutRecovery,omitempty" xml:"ForceDeleteWithoutRecovery,omitempty"`
	// Specifies the recovery period of the secret if you do not forcibly delete it. Default value: 30. Unit: Days.
	RecoveryWindowInDays *string `json:"RecoveryWindowInDays,omitempty" xml:"RecoveryWindowInDays,omitempty"`
	// The name of the secret.
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
	PlannedDeleteTime *string `json:"PlannedDeleteTime,omitempty" xml:"PlannedDeleteTime,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the secret.
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSecretResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// *   Enabled: KMS is enabled.
	//
	// *   NotEnabled: KMS is disabled.
	//
	// *   InDebt: Your account is overdue, and KMS stops providing services.
	//
	// > If your Alibaba Cloud account is overdue, top up your account at the earliest opportunity to avoid impacts on your services.
	//
	// *   Suspended: KMS is suspended.
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAccountKmsStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeCertificateRequest struct {
	// The ID of the certificate. The ID must be globally unique in Certificates Manager.
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
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the certificate. The ID must be globally unique in Certificates Manager.
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The time when the certificate was created.
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// Indicates whether the private key of the certificate can be exported for use. Valid values:
	//
	// *   true: The private key of the certificate can be exported for use. This is the default value.
	// *   false: The private key of the certificate cannot be exported for use.
	ExportablePrivateKey *bool `json:"ExportablePrivateKey,omitempty" xml:"ExportablePrivateKey,omitempty"`
	// The certificate issuer in the distinguished name (DN) format.
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The type of the key.
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The end of the validity period of the certificate.
	NotAfter *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The beginning of the validity period of the certificate.
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the certificate.
	Serial *string `json:"Serial,omitempty" xml:"Serial,omitempty"`
	// The signature algorithm of the certificate. Valid values:
	//
	// *   RSA2048-SHA256
	// *   ECDSA-SHA256
	// *   SM2-SM3
	SignatureAlgorithm *string `json:"SignatureAlgorithm,omitempty" xml:"SignatureAlgorithm,omitempty"`
	// The status of the certificate. Valid values:
	//
	// *   PENDING: The certificate is to be imported.
	// *   ACTIVE: The certificate is enabled.
	// *   INACTIVE: The certificate is disabled.
	// *   REVOKED: The certificate is revoked.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The subject of the certificate, which is in the DN format.
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// The alias of the certificate subject.
	//
	// A domain name list is supported. A maximum of 10 domain names are supported.
	SubjectAlternativeNames []*string `json:"SubjectAlternativeNames,omitempty" xml:"SubjectAlternativeNames,omitempty" type:"Repeated"`
	// The public key identifier of the certificate subject.
	SubjectKeyIdentifier *string `json:"SubjectKeyIdentifier,omitempty" xml:"SubjectKeyIdentifier,omitempty"`
	// The public key of the certificate.
	SubjectPublicKey *string `json:"SubjectPublicKey,omitempty" xml:"SubjectPublicKey,omitempty"`
	// The tag of the certificate.
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The time when the certificate was updated.
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// You can also set this parameter to an alias that is bound to the CMK. For more information, see [Overview of aliases](~~68522~~).
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
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// Indicates whether automatic key rotation is enabled. Valid values:
	//
	// *   Enabled
	// *   Disabled
	// *   Suspended
	//
	// For more information, see [Automatic key rotation](~~134270~~).
	//
	// >  Only symmetric CMKs support automatic key rotation.
	AutomaticRotation *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	// The time when the CMK was created. The time is displayed in UTC.
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// The Alibaba Cloud account that is used to create the CMK.
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The ID of the dedicated KMS instance.
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// The time at which the CMK is scheduled for deletion. The time is displayed in UTC.
	//
	// For more information, see [ScheduleKeyDeletion](~~44196~~).
	//
	// >  This parameter is returned only when the value of the KeyState parameter is PendingDeletion.
	DeleteDate *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
	// Indicates whether deletion protection is enabled. Valid values:
	//
	// *   Enabled
	// *   Disabled
	DeletionProtection *string `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The description of deletion protection.
	DeletionProtectionDescription *string `json:"DeletionProtectionDescription,omitempty" xml:"DeletionProtectionDescription,omitempty"`
	// The description of the CMK.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The type of the CMK.
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The status of the CMK.
	//
	// For more information, see [Impact of CMK status on API operations](~~44211~~).
	KeyState *string `json:"KeyState,omitempty" xml:"KeyState,omitempty"`
	// The usage of the CMK.
	KeyUsage *string `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty"`
	// The time when the last rotation was performed. The time is displayed in UTC. For a new CMK, the value of this parameter is the time when the initial version of the CMK was generated.
	LastRotationDate *string `json:"LastRotationDate,omitempty" xml:"LastRotationDate,omitempty"`
	// The time when the key material expires. The time is displayed in UTC. If this parameter value is empty, the key material does not expire.
	MaterialExpireTime *string `json:"MaterialExpireTime,omitempty" xml:"MaterialExpireTime,omitempty"`
	// The time when the next rotation will be performed.
	//
	// >  This parameter is returned only when the value of the AutomaticRotation parameter is Enabled or Suspended.
	NextRotationDate *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	// The source of the key material for the CMK.
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The ID of the current primary key version for the symmetric CMK.
	PrimaryKeyVersion *string `json:"PrimaryKeyVersion,omitempty" xml:"PrimaryKeyVersion,omitempty"`
	// The protection level of the CMK.
	ProtectionLevel *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	// The interval for automatic key rotation.
	//
	// Unit: seconds.
	//
	// For example, if the value is 604800s, automatic key rotation is performed at a 7-day interval.
	//
	// >  This parameter is returned only when the value of the AutomaticRotation parameter is Enabled or Suspended.
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// You can also set this parameter to an alias that is bound to the CMK. For more information, see [Alias overview](~~68522~~).
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The globally unique ID of the CMK version.
	//
	// You can call the [ListKeyVersions](~~133966~~) operation to query the versions of the CMK.
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
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// The globally unique ID of the CMK.
	//
	// >  If you set the KeyId parameter in the request to an alias of the CMK, the ID of the CMK to which the alias is bound is returned.
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The globally unique ID of the CMK version.
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeKeyVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeRegionsResponseBody struct {
	// The region.
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The ID of the request.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// *   true: The resource tags are returned.
	// *   false: The resource tags are not returned. This is the default value.
	FetchTags *string `json:"FetchTags,omitempty" xml:"FetchTags,omitempty"`
	// The name of the secret.
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
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// Indicates whether automatic rotation is enabled. Valid values:
	//
	// *   Enabled: indicates that automatic rotation is enabled.
	// *   Disabled: indicates that automatic rotation is disabled.
	// *   Invalid: indicates that the status of automatic rotation is abnormal. In this case, Secrets Manager cannot automatically rotate the secret.
	//
	// >  This parameter is returned only for a managed ApsaraDB RDS secret, a managed RAM secret, or a managed ECS secret.
	AutomaticRotation *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	// The time when the secret was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the dedicated KMS instance.
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// The description of the secret.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the customer master key (CMK) that is used to encrypt the secret value.
	EncryptionKeyId *string `json:"EncryptionKeyId,omitempty" xml:"EncryptionKeyId,omitempty"`
	// The extended configuration of the secret.
	//
	// >  This parameter is returned only for a managed ApsaraDB RDS secret, a managed Resource Access Management (RAM) secret, or a managed Elastic Compute Service (ECS) secret.
	ExtendedConfig *string `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	// The time when the last rotation was performed.
	//
	// >  This parameter is returned if the secret was rotated.
	LastRotationDate *string `json:"LastRotationDate,omitempty" xml:"LastRotationDate,omitempty"`
	// The time when the next rotation will be performed.
	//
	// >  This parameter is returned when automatic rotation is enabled.
	NextRotationDate *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	// The time when the secret is scheduled to be deleted.
	PlannedDeleteTime *string `json:"PlannedDeleteTime,omitempty" xml:"PlannedDeleteTime,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The interval for automatic rotation.
	//
	// The value is in the `integer[unit]` format. `integer` indicates the length of time. `unit`: indicates the time unit. The value of `unit` is fixed as s. For example, if the value is 604800s, automatic rotation is performed at a 7-day interval.
	//
	// >  This parameter is returned when automatic rotation is enabled.
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	// The name of the secret.
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The type of the secret. Valid values:
	//
	// *   Generic: indicates a generic secret.
	// *   Rds: indicates a managed ApsaraDB RDS secret.
	// *   RAMCredentials: indicates a managed RAM secret.
	// *   ECS: indicates a managed ECS secret.
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// The resource tags of the secret.
	//
	// This parameter is not returned if you set the FetchTags parameter to false or you do not specify the FetchTags parameter.
	Tags *DescribeSecretResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The time when the secret was updated.
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
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSecretResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// A JSON string that consists of key-value pairs. If you specify this parameter, an equivalent value is required when you call the Decrypt operation. For more information, see [EncryptionContext](~~42975~~).
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// The globally unique ID of the CMK. You can also set this parameter to an alias that is bound to the CMK. For more information, see [Use aliases](~~68522~~).
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The plaintext to be encrypted. The plaintext must be Base64 encoded.
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
	// A JSON string that consists of key-value pairs. If you specify this parameter, an equivalent value is required when you call the Decrypt operation. For more information, see [EncryptionContext](~~42975~~).
	EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// The globally unique ID of the CMK. You can also set this parameter to an alias that is bound to the CMK. For more information, see [Use aliases](~~68522~~).
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The plaintext to be encrypted. The plaintext must be Base64 encoded.
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
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The globally unique ID of the CMK. If you set the KeyId parameter to an alias, the ID of the CMK to which the alias is bound is returned.
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the key version that is used to encrypt the plaintext. It is the primary version of the CMK.
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The ID of the request.
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EncryptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// A JSON string that consists of key-value pairs. If you specify this parameter when you use a CMK to encrypt the data key, an equivalent value is required here. For more information, see [EncryptionContext](~~42975~~).
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// A Base64-encoded public key.
	PublicKeyBlob *string `json:"PublicKeyBlob,omitempty" xml:"PublicKeyBlob,omitempty"`
	// The encryption algorithm based on which you want to use the public key specified by PublicKeyBlob to encrypt the data key. For more information about encryption algorithms, see [AsymmetricDecrypt](~~148130~~).
	//
	// Valid values:
	//
	// *   RSAES_OAEP_SHA\_256
	// *   RSAES_OAEP_SHA\_1
	// *   SM2PKE
	WrappingAlgorithm *string `json:"WrappingAlgorithm,omitempty" xml:"WrappingAlgorithm,omitempty"`
	// The key type of the public key specified by PublicKeyBlob. For more information about key types, see [Introduction to asymmetric keys](~~148147~~).
	//
	// Valid values:
	//
	// *   RSA\_2048
	// *   EC_SM2
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
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// A JSON string that consists of key-value pairs. If you specify this parameter when you use a CMK to encrypt the data key, an equivalent value is required here. For more information, see [EncryptionContext](~~42975~~).
	EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// A Base64-encoded public key.
	PublicKeyBlob *string `json:"PublicKeyBlob,omitempty" xml:"PublicKeyBlob,omitempty"`
	// The encryption algorithm based on which you want to use the public key specified by PublicKeyBlob to encrypt the data key. For more information about encryption algorithms, see [AsymmetricDecrypt](~~148130~~).
	//
	// Valid values:
	//
	// *   RSAES_OAEP_SHA\_256
	// *   RSAES_OAEP_SHA\_1
	// *   SM2PKE
	WrappingAlgorithm *string `json:"WrappingAlgorithm,omitempty" xml:"WrappingAlgorithm,omitempty"`
	// The key type of the public key specified by PublicKeyBlob. For more information about key types, see [Introduction to asymmetric keys](~~148147~~).
	//
	// Valid values:
	//
	// *   RSA\_2048
	// *   EC_SM2
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
	ExportedDataKey *string `json:"ExportedDataKey,omitempty" xml:"ExportedDataKey,omitempty"`
	// The ID of the CMK that is used to decrypt the specified ciphertext of the data key.
	//
	// This parameter is the globally unique ID of the CMK.
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the CMK version that is used to decrypt the specified ciphertext of the data key.
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The ID of the request.
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExportDataKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// A JSON string of key-value pairs. If you specify this parameter here, an equivalent value is required when you decrypt or re-encrypt the data key. For more information, see [EncryptionContext](~~42975~~).
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// The globally unique ID of the CMK. You can also set this parameter to an alias that is bound to the CMK. For more information, see [Use aliases](~~68522~~).
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The length of the data key that you want to generate. Valid values:
	//
	// *   AES\_256: a 256-bit symmetric key
	// *   AES\_128: a 128-bit symmetric key
	//
	// >  We recommend that you use the KeySpec or NumberOfBytes parameter to specify the length of a data key. If both parameters are not specified, KMS generates a 256-bit data key. If both parameters are specified, KMS ignores the KeySpec parameter.
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The length of the data key that you want to generate.
	//
	// Valid values: 1 to 1024.
	//
	// Unit: bytes.
	NumberOfBytes *int32 `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
	// A Base64-encoded public key.
	PublicKeyBlob *string `json:"PublicKeyBlob,omitempty" xml:"PublicKeyBlob,omitempty"`
	// The encryption algorithm based on which you want to use the public key specified by PublicKeyBlob to encrypt the data key. For more information about encryption algorithms, see [AsymmetricDecrypt](~~148130~~).
	//
	// Valid values:
	//
	// *   RSAES_OAEP_SHA\_256
	// *   RSAES_OAEP_SHA\_1
	// *   SM2PKE
	WrappingAlgorithm *string `json:"WrappingAlgorithm,omitempty" xml:"WrappingAlgorithm,omitempty"`
	// The key type of the public key specified by PublicKeyBlob. For more information about key types, see [Introduction to asymmetric keys](~~148147~~).
	//
	// Valid values:
	//
	// *   RSA\_2048
	// *   EC_SM2
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
	// A JSON string of key-value pairs. If you specify this parameter here, an equivalent value is required when you decrypt or re-encrypt the data key. For more information, see [EncryptionContext](~~42975~~).
	EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// The globally unique ID of the CMK. You can also set this parameter to an alias that is bound to the CMK. For more information, see [Use aliases](~~68522~~).
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The length of the data key that you want to generate. Valid values:
	//
	// *   AES\_256: a 256-bit symmetric key
	// *   AES\_128: a 128-bit symmetric key
	//
	// >  We recommend that you use the KeySpec or NumberOfBytes parameter to specify the length of a data key. If both parameters are not specified, KMS generates a 256-bit data key. If both parameters are specified, KMS ignores the KeySpec parameter.
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The length of the data key that you want to generate.
	//
	// Valid values: 1 to 1024.
	//
	// Unit: bytes.
	NumberOfBytes *int32 `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
	// A Base64-encoded public key.
	PublicKeyBlob *string `json:"PublicKeyBlob,omitempty" xml:"PublicKeyBlob,omitempty"`
	// The encryption algorithm based on which you want to use the public key specified by PublicKeyBlob to encrypt the data key. For more information about encryption algorithms, see [AsymmetricDecrypt](~~148130~~).
	//
	// Valid values:
	//
	// *   RSAES_OAEP_SHA\_256
	// *   RSAES_OAEP_SHA\_1
	// *   SM2PKE
	WrappingAlgorithm *string `json:"WrappingAlgorithm,omitempty" xml:"WrappingAlgorithm,omitempty"`
	// The key type of the public key specified by PublicKeyBlob. For more information about key types, see [Introduction to asymmetric keys](~~148147~~).
	//
	// Valid values:
	//
	// *   RSA\_2048
	// *   EC_SM2
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
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The data key encrypted by using the public key and then exported.
	ExportedDataKey *string `json:"ExportedDataKey,omitempty" xml:"ExportedDataKey,omitempty"`
	// The globally unique ID of the CMK.
	//
	// >  If you set the KeyId parameter to an alias, the ID of the CMK to which the alias is bound is returned.
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the CMK version that is used to encrypt the plaintext. It is the primary version of the CMK.
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The ID of the request.
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateAndExportDataKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// If you specify this parameter, an equivalent value is required when you call the [Decrypt](~~28950~~) operation. For more information, see [EncryptionContext](~~42975~~).
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// You can also set this parameter to an alias that is bound to the CMK. For more information, see [Alias overview](~~68522~~).
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The type of the data key that you want to generate. Valid values:
	//
	// *   AES\_256: a 256-bit symmetric key
	// *   AES\_128: a 128-bit symmetric key
	//
	// >  We recommend that you use the KeySpec or NumberOfBytes parameter to specify the length of a data key. If none of the parameters are specified, KMS generates a 256-bit data key. If both parameters are specified, KMS ignores the KeySpec parameter.
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The length of the data key that you want to generate. Unit: bytes.
	//
	// Valid values: 1 to 1024.
	//
	// Default value:
	//
	// *   If the KeySpec parameter is set to AES\_256, set the value of the NumberOfBytes parameter to 32.
	// *   If the KeySpec parameter is set to AES\_128, set the value of the NumberOfBytes parameter to 16.
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
	// If you specify this parameter, an equivalent value is required when you call the [Decrypt](~~28950~~) operation. For more information, see [EncryptionContext](~~42975~~).
	EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// You can also set this parameter to an alias that is bound to the CMK. For more information, see [Alias overview](~~68522~~).
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The type of the data key that you want to generate. Valid values:
	//
	// *   AES\_256: a 256-bit symmetric key
	// *   AES\_128: a 128-bit symmetric key
	//
	// >  We recommend that you use the KeySpec or NumberOfBytes parameter to specify the length of a data key. If none of the parameters are specified, KMS generates a 256-bit data key. If both parameters are specified, KMS ignores the KeySpec parameter.
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The length of the data key that you want to generate. Unit: bytes.
	//
	// Valid values: 1 to 1024.
	//
	// Default value:
	//
	// *   If the KeySpec parameter is set to AES\_256, set the value of the NumberOfBytes parameter to 32.
	// *   If the KeySpec parameter is set to AES\_128, set the value of the NumberOfBytes parameter to 16.
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
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// >  If you set the KeyId parameter in the request to an alias of the CMK, the ID of the CMK to which the alias is bound is returned.
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the CMK version. The ID must be globally unique.
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The Base64 encoded plaintext of the data key.
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateDataKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// A JSON string that consists of key-value pairs. If you specify this parameter, an equivalent value is required when you call the Decrypt operation. For more information, see [EncryptionContext](~~42975~~).
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// The globally unique ID of the CMK. You can also set this parameter to an alias that is bound to the CMK. For more information, see Use aliases.
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The length of the data key that you want to generate. Valid values:
	//
	// *   AES\_256: 256-bit symmetric key
	// *   AES\_128: 128-bit symmetric key
	//
	// >  We recommend that you use the KeySpec or NumberOfBytes parameter to specify the length of a data key. If both of them are not specified, KMS generates a 256-bit data key. If both of them are specified, KMS ignores the KeySpec parameter.
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The length of the data key that you want to generate.
	//
	// Valid values: 1 to 1024.
	//
	// Unit: bytes.
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
	// A JSON string that consists of key-value pairs. If you specify this parameter, an equivalent value is required when you call the Decrypt operation. For more information, see [EncryptionContext](~~42975~~).
	EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// The globally unique ID of the CMK. You can also set this parameter to an alias that is bound to the CMK. For more information, see Use aliases.
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The length of the data key that you want to generate. Valid values:
	//
	// *   AES\_256: 256-bit symmetric key
	// *   AES\_128: 128-bit symmetric key
	//
	// >  We recommend that you use the KeySpec or NumberOfBytes parameter to specify the length of a data key. If both of them are not specified, KMS generates a 256-bit data key. If both of them are specified, KMS ignores the KeySpec parameter.
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The length of the data key that you want to generate.
	//
	// Valid values: 1 to 1024.
	//
	// Unit: bytes.
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
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The globally unique ID of the CMK.
	//
	// >  If you set the KeyId parameter to an alias, the ID of the CMK to which the alias is bound is returned.
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the key version that is used to encrypt the plaintext. It is the primary version of the CMK.
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The ID of the request.
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateDataKeyWithoutPlaintextResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The certificate chain in the PEM format.
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The ID of the certificate.
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The CSR in the PEM format.
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The ID of the request.
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetParametersForImportRequest struct {
	// The globally unique ID of the CMK.
	//
	// >  You can import key material only for CMKs whose Origin parameter is set to EXTERNAL.
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The algorithm that is used to encrypt key material.
	WrappingAlgorithm *string `json:"WrappingAlgorithm,omitempty" xml:"WrappingAlgorithm,omitempty"`
	// The type of the public key that is used to encrypt key material.
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
	// The token is valid for 24 hours. The value of this parameter is required when you call the [ImportKeyMaterial](~~68622~~) operation.
	ImportToken *string `json:"ImportToken,omitempty" xml:"ImportToken,omitempty"`
	// The globally unique ID of the CMK.
	//
	// The value of this parameter is required when you call the [ImportKeyMaterial](~~68622~~) operation.
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The public key that is used to encrypt key material.
	//
	// The public key is Base64-encoded.
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the token expires.
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetParametersForImportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The globally unique ID of the CMK. You can also set this parameter to an alias that is bound to the CMK. For more information, see [Use aliases](~~68522~~).
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The globally unique ID of the CMK version.
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
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The version of the CMK that is used to encrypt the plaintext.
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The public key returned in the PEM format.
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	// The ID of the request.
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPublicKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// ` Valid characters: 0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ! \"#$%&\"()*+,-. /:;<=>? @[\] your_project_id} ~  `.
	//
	// This parameter is empty by default.
	ExcludeCharacters *string `json:"ExcludeCharacters,omitempty" xml:"ExcludeCharacters,omitempty"`
	// Specifies whether to exclude lowercase letters.
	//
	// Valid values:
	//
	// *   true
	// *   false
	ExcludeLowercase *string `json:"ExcludeLowercase,omitempty" xml:"ExcludeLowercase,omitempty"`
	// Specifies whether to exclude digits.
	//
	// Valid values:
	//
	// *   true
	// *   false
	ExcludeNumbers *string `json:"ExcludeNumbers,omitempty" xml:"ExcludeNumbers,omitempty"`
	// Specifies whether to exclude special characters.
	//
	// Valid values:
	//
	// *   true
	// *   false
	ExcludePunctuation *string `json:"ExcludePunctuation,omitempty" xml:"ExcludePunctuation,omitempty"`
	// Specifies whether to exclude uppercase letters.
	//
	// Valid values:
	//
	// *   true
	// *   false
	ExcludeUppercase *string `json:"ExcludeUppercase,omitempty" xml:"ExcludeUppercase,omitempty"`
	// The number of bytes that the password to be generated contains.
	//
	// Valid values: 8 to 128.
	//
	// Default value: 32
	PasswordLength *string `json:"PasswordLength,omitempty" xml:"PasswordLength,omitempty"`
	// Specifies whether to include all the preceding character types.
	//
	// Valid values:
	//
	// *   true
	// *   false
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
	RandomPassword *string `json:"RandomPassword,omitempty" xml:"RandomPassword,omitempty"`
	// The ID of the request.
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRandomPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetSecretValueRequest struct {
	// Specifies whether to obtain the extended configuration of the secret. Valid values:
	//
	// *   true
	// *   false: This is the default value.
	//
	// >  This parameter is ignored for a generic secret.
	FetchExtendedConfig *bool `json:"FetchExtendedConfig,omitempty" xml:"FetchExtendedConfig,omitempty"`
	// The name of the secret.
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The version number of the secret value. If you specify this parameter, Secrets Manager returns the secret value of the specified version.
	//
	// >  This parameter is ignored for a managed ApsaraDB RDS secret, a managed RAM secret, or a managed ECS secret.
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The stage label that marks the secret version. If you specify this parameter, Secrets Manager returns the secret value of the version that is marked with the specified stage label.
	//
	// Default value: ACSCurrent.
	//
	// >  For a managed ApsaraDB RDS secret, a managed RAM secret, or a managed ECS secret, Secrets Manager can return only the secret value of the version marked with ACSPrevious or ACSCurrent.
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
	// *   Enabled: indicates that automatic rotation is enabled.
	// *   Disabled: indicates that automatic rotation is disabled.
	// *   Invalid: indicates that the status of automatic rotation is abnormal. In this case, Secrets Manager cannot automatically rotate the secret.
	//
	// >  This parameter is returned only for a managed ApsaraDB RDS secret, a managed RAM secret, or a managed ECS secret.
	AutomaticRotation *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	// The time when the secret was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The extended configuration of the secret.
	//
	// >  This parameter is returned if you set the FetchExtendedConfig parameter to true. This parameter is returned only for a managed ApsaraDB RDS secret, a managed RAM secret, or a managed ECS secret.
	ExtendedConfig *string `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	// The time when the last rotation was performed.
	//
	// >  This parameter is returned if the secret was rotated.
	LastRotationDate *string `json:"LastRotationDate,omitempty" xml:"LastRotationDate,omitempty"`
	// The time when the next rotation will be performed.
	//
	// >  This parameter is returned if automatic rotation is enabled.
	NextRotationDate *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The interval for automatic rotation.
	//
	// The value is in the `integer[unit]` format. The `unit` field has a fixed value of s. For example, if the value is 604800s, automatic rotation is performed at a 7-day interval.
	//
	// >  This parameter is returned if automatic rotation is enabled.
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	// The secret value. Secrets Manager decrypts the ciphertext of the secret value and returns the plaintext of the secret value in this parameter.
	//
	// *   For a generic secret, the secret value of the specified version is returned.
	//
	// *   For a managed ApsaraDB RDS secret, the value is returned in the following format:`{"AccountName":"","AccountPassword":""}` .
	//
	// *   For a managed RAM secret, the secret value is returned in the following format: `{"AccessKeyId":"Adfdsfd","AccessKeySecret":"fdsfdsf","GenerateTimestamp": "2016-03-25T10:42:40Z"}`.
	//
	// *   For a managed ECS secret, the secret value is returned in one of the following formats:
	//
	//     *   `{"UserName":"root","Password":"H5asdasdsads****"}`: The secret value is returned in this format if the ECS secret is a password.
	//     *   `{"UserName":"root","PublicKey":"ssh-rsa ****mKwnVix9YTFY9Rs= imported-openssh-key","PrivateKey": "d6bee1cb-2e14-4277-ba6b-73786b21****"}`: The secret value is returned in this format is the ECS secret is a pair of SSH keys. The private key is in the Privacy Enhanced Mail (PEM) format.
	SecretData *string `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
	// The type of the secret value. Valid values:
	//
	// *   text
	// *   binary
	SecretDataType *string `json:"SecretDataType,omitempty" xml:"SecretDataType,omitempty"`
	// The name of the secret.
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The type of the secret. Valid values:
	//
	// *   Generic: indicates a generic secret.
	// *   Rds: indicates a managed ApsaraDB RDS secret.
	// *   RAMCredentials: indicates a managed RAM secret.
	// *   ECS: indicates a managed ECS secret.
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// The version number of the secret value.
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSecretValueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// Use** GetParametersForImport** the Returned public key and the base64-encoded key material.
	EncryptedKeyMaterial *string `json:"EncryptedKeyMaterial,omitempty" xml:"EncryptedKeyMaterial,omitempty"`
	// By calling** GetParametersForImport** the import token.
	ImportToken *string `json:"ImportToken,omitempty" xml:"ImportToken,omitempty"`
	// The ID of the CMK to be imported.
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The time when the key material expires.
	//
	// If this parameter is not specified or set this parameter to 0, the key material does not expire.
	//
	// >  The value cannot be earlier than the time when the API is called (based on the server time).
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ImportKeyMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 0 to 100.
	//
	// Default value: 10.
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
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned aliases.
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
	AliasArn *string `json:"AliasArn,omitempty" xml:"AliasArn,omitempty"`
	// The ID of the alias.
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The CMK to which the alias belongs.
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAliasesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The number of the page to return.
	//
	// Valid values: an integer that is greater than 0.
	//
	// Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 0 to 101.
	//
	// Default value: 10
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
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned CMKs.
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
	AliasArn *string `json:"AliasArn,omitempty" xml:"AliasArn,omitempty"`
	// The ID of the alias.
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The CMK to which an alias is bound.
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAliasesByKeyIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListKeyVersionsRequest struct {
	// The globally unique ID of the CMK. You can also set this parameter to an alias that is bound to the CMK. For more information, see [Use aliases](~~68522~~).
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 0 to 101.
	//
	// Default value: 10.
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
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned key versions.
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
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// The globally unique ID of the CMK.
	//
	// >  If you set the KeyId parameter to the alias of the CMK, the ID of the CMK to which the alias is bound is returned.
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The globally unique ID of the CMK version.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListKeyVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// *   Key
	//
	//     *   Description: the property that you want to filter.
	//
	//     *   Type: string.
	//
	//     *   Valid values:
	//
	//         *   KeyState: the status of the CMK.
	//         *   KeySpec: the type of the CMK.
	//         *   KeyUsage: the usage of the CMK.
	//         *   ProtectionLevel: the protection level.
	//         *   CreatorType: the type of the creator.
	//
	// *   Values
	//
	//     *   Description: the value to be included after filtering.
	//
	//     *   Format: string array.
	//
	//     *   Length: 0 to 10.
	//
	//     *   Valid values:
	//
	//         *   When Key is set to KeyState, the value can be Enabled, Disabled, PendingDeletion, or PendingImport.
	//
	//         *   When Key is set to KeySpec, the value can be Aliyun_AES\_256, Aliyun_SM4, RSA\_2048, EC_P256, EC_P256K, or EC_SM2.
	//
	//             Note: You can create CMKs of the EC_SM2 or Aliyun_SM4 type only in regions where State Cryptography Administration (SCA)-certified managed HSMs reside. For more information about the regions, see [Supported regions](~~125803~~). If your region does not support EC_SM2 or Aliyun_SM4, the two values are ignored if they are specified.
	//
	//         *   When Key is set to KeyUsage, the value can be ENCRYPT/DECRYPT or SIGN/VERIFY. ENCRYPT/DECRYPT indicates that the CMK is used to encrypt and decrypt data. SIGN/VERIFY indicates that the CMK is used to generate and verify digital signatures.
	//
	//         *   When Key is set to ProtectionLevel, the value can be SOFTWARE (software) or HSM (hardware).
	//
	//             You can set ProtectionLevel to HSM in only specific regions. For more information about the regions, see [Supported regions](~~125803~~). If your region does not support the value HSM, the value is ignored if the value is specified.
	//
	//         *   If Key is set to CreatorType, the value can be User or Service. User indicates that CMKs created by the current account are queried. Service indicates that CMKs automatically created by other cloud services authorized by the current account are queried.
	//
	// The logical relationship between different keys is AND, and the logical relationship between multiple items in the same key is OR. Example:
	//
	// `[ {"Key":"KeyState", "Values":["Enabled","Disabled"]}, {"Key":"KeyState", "Values":["PendingDeletion"]}, {"Key":"KeySpec", "Values":["Aliyun_AES_256"]}]`. In this example, the semantics are:`(KeyState=Enabled OR KeyState=Disabled OR KeyState=PendingDeletion) AND (KeySpec=Aliyun_AES_ 256)`.
	Filters *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10
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
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of CMKs.
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
	KeyArn *string `json:"KeyArn,omitempty" xml:"KeyArn,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListResourceTagsRequest struct {
	// The globally unique ID of the CMK.
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
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The tag key.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResourceTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// *   false: no
	// *   true: yes
	//
	// Default value: false.
	IncludeDeprecated *string `json:"IncludeDeprecated,omitempty" xml:"IncludeDeprecated,omitempty"`
	// The number of the page to return. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the secret.
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
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the secret.
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The number of entries returned on the current page.
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
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The version number.
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSecretVersionIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// Specifies whether to return the resource tags of the secret. Valid values:
	//
	// *   true: returns the resource tags.
	// *   false: does not return the resource tags. This is the default value.
	FetchTags *string `json:"FetchTags,omitempty" xml:"FetchTags,omitempty"`
	// The secret filter. The filter consists of one or more key-value pairs. You can specify one key-value pair or leave this parameter empty. If you use one tag key or tag value to filter resources, up to 4,000 resources can be queried. If you want to query more than 4,000 resources, call the [ListResourceTags](~~120090~~) operation.
	//
	// *   Key
	//
	//     *   Description: the property that you want to filter.
	//
	//     *   Type: string.
	//
	//     *   Valid values:
	//
	//         *   SecretName: the secret name.
	//         *   Description: the description of the secret.
	//         *   TagKey: the tag key.
	//         *   TagValue: the tag value.
	//
	// *   Values
	//
	//     *   Description: the value to be included after filtering.
	//
	//     *   Type: string.
	//
	//     *   Length: 0 to 10.
	//
	//     *   Valid values:
	//
	//         *   If the Key field is set to SecretName, the value must be 1 to 192 characters in length and can contain letters, digits, and special characters `_ / + = . @ -`.
	//         *   If the Key field is set to Description, the value must be 1 to 256 characters in length.
	//         *   If the Key field is set to TagKey, the value must be 1 to 256 characters in length and can contain letters, digits, and special characters `/ _ - . + = @ :`.
	//         *   If the Key field is set to TagValue, the value must be 1 to 256 characters in length and can contain letters, numbers, and special characters `/ _ - . + = @ :`.
	//
	// The logical relationship between values of the Values field in a key-value pair is OR. Example: `[ {"Key":"SecretName", "Values":["sec1","sec2"]}]`. In this example, the semantics are `SecretName=sec 1 OR SecretName=sec 2`.
	Filters *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
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
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of secrets.
	SecretList *ListSecretsResponseBodySecretList `json:"SecretList,omitempty" xml:"SecretList,omitempty" type:"Struct"`
	// The number of returned secrets.
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
	// The time when the secret was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the secret is scheduled to be deleted.
	PlannedDeleteTime *string `json:"PlannedDeleteTime,omitempty" xml:"PlannedDeleteTime,omitempty"`
	// The secret name.
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The type of the secret. Valid values:
	//
	// *   Generic: indicates a generic secret.
	// *   Rds: indicates a managed ApsaraDB RDS secret.
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// The resource tags of the secret.
	//
	// This parameter is not returned if you set the FetchTags parameter to false or do not specify the FetchTags parameter.
	Tags *ListSecretsResponseBodySecretListSecretTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The time when the secret was updated.
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
	// The tag key.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSecretsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type OpenKmsServiceResponseBody struct {
	// The ID of the request.
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenKmsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	SecretData *string `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
	// The type of the secret value. Valid values:
	//
	// *   text: This is the default value.
	// *   binary
	SecretDataType *string `json:"SecretDataType,omitempty" xml:"SecretDataType,omitempty"`
	// The name of the secret.
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The new version of the secret value. Version numbers must be unique in each secret.
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The stage labels that are used to mark the new version. If you do not specify this parameter, Secrets Manager marks the new version with ACSCurrent.
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the secret.
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The new version of the secret value.
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutSecretValueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// *   Symmetric encryption: the ciphertext returned after you call the [Encrypt](~~28949~~), [GenerateDataKey](~~28948~~), [GenerateDataKeyWithoutPlaintext](~~134043~~), or [GenerateAndExportDataKey](~~176804~~) operation
	// *   Asymmetric encryption: the public key-encrypted ciphertext returned after you call the [GenerateAndExportDataKey](~~176804~~) operation, or the ciphertext encrypted by using the public key of an asymmetric key pair outside KMS
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// A JSON string that consists of key-value pairs. This parameter specifies the EncryptionContext that is used to re-encrypt the decrypted data or data key.
	DestinationEncryptionContext map[string]interface{} `json:"DestinationEncryptionContext,omitempty" xml:"DestinationEncryptionContext,omitempty"`
	// The ID of the symmetric CMK that is used to re-encrypt the ciphertext after the ciphertext is decrypted.
	DestinationKeyId *string `json:"DestinationKeyId,omitempty" xml:"DestinationKeyId,omitempty"`
	// The encryption algorithm based on which the public key is used to encrypt the ciphertext specified by CiphertextBlob. For more information about encryption algorithms, see [AsymmetricDecrypt](~~148130~~).
	//
	// Valid values:
	//
	// *   RSAES_OAEP_SHA\_256
	// *   RSAES_OAEP_SHA\_1
	// *   SM2PKE
	//
	// >  If you set CiphertextBlob to the public key-encrypted ciphertext that is returned after an asymmetric encryption operation, specify this parameter.
	SourceEncryptionAlgorithm *string `json:"SourceEncryptionAlgorithm,omitempty" xml:"SourceEncryptionAlgorithm,omitempty"`
	// A JSON string that consists of key-value pairs. If you specify EncryptionContext when you call the [Encrypt](~~28949~~), [GenerateDataKey](~~28948~~), [GenerateDataKeyWithoutPlaintext](~~134043~~), or [GenerateAndExportDataKey](~~176804~~) operation to encrypt the data or data key, an equivalent value is required here. For more information, see [EncryptionContext](~~42975~~).
	//
	// >  If you set CiphertextBlob to the ciphertext that is returned after a symmetric encryption operation, specify this parameter.
	SourceEncryptionContext map[string]interface{} `json:"SourceEncryptionContext,omitempty" xml:"SourceEncryptionContext,omitempty"`
	// The ID of the CMK that is used to decrypt the ciphertext.
	//
	// This parameter is the globally unique ID of the CMK.
	//
	// >  If you set CiphertextBlob to the public key-encrypted ciphertext that is returned after an asymmetric encryption operation, specify this parameter.
	SourceKeyId *string `json:"SourceKeyId,omitempty" xml:"SourceKeyId,omitempty"`
	// The ID of the CMK version that is used to decrypt the ciphertext.
	//
	// >  If you set CiphertextBlob to the public key-encrypted ciphertext that is returned after an asymmetric encryption operation, specify this parameter.
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
	// *   Symmetric encryption: the ciphertext returned after you call the [Encrypt](~~28949~~), [GenerateDataKey](~~28948~~), [GenerateDataKeyWithoutPlaintext](~~134043~~), or [GenerateAndExportDataKey](~~176804~~) operation
	// *   Asymmetric encryption: the public key-encrypted ciphertext returned after you call the [GenerateAndExportDataKey](~~176804~~) operation, or the ciphertext encrypted by using the public key of an asymmetric key pair outside KMS
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// A JSON string that consists of key-value pairs. This parameter specifies the EncryptionContext that is used to re-encrypt the decrypted data or data key.
	DestinationEncryptionContextShrink *string `json:"DestinationEncryptionContext,omitempty" xml:"DestinationEncryptionContext,omitempty"`
	// The ID of the symmetric CMK that is used to re-encrypt the ciphertext after the ciphertext is decrypted.
	DestinationKeyId *string `json:"DestinationKeyId,omitempty" xml:"DestinationKeyId,omitempty"`
	// The encryption algorithm based on which the public key is used to encrypt the ciphertext specified by CiphertextBlob. For more information about encryption algorithms, see [AsymmetricDecrypt](~~148130~~).
	//
	// Valid values:
	//
	// *   RSAES_OAEP_SHA\_256
	// *   RSAES_OAEP_SHA\_1
	// *   SM2PKE
	//
	// >  If you set CiphertextBlob to the public key-encrypted ciphertext that is returned after an asymmetric encryption operation, specify this parameter.
	SourceEncryptionAlgorithm *string `json:"SourceEncryptionAlgorithm,omitempty" xml:"SourceEncryptionAlgorithm,omitempty"`
	// A JSON string that consists of key-value pairs. If you specify EncryptionContext when you call the [Encrypt](~~28949~~), [GenerateDataKey](~~28948~~), [GenerateDataKeyWithoutPlaintext](~~134043~~), or [GenerateAndExportDataKey](~~176804~~) operation to encrypt the data or data key, an equivalent value is required here. For more information, see [EncryptionContext](~~42975~~).
	//
	// >  If you set CiphertextBlob to the ciphertext that is returned after a symmetric encryption operation, specify this parameter.
	SourceEncryptionContextShrink *string `json:"SourceEncryptionContext,omitempty" xml:"SourceEncryptionContext,omitempty"`
	// The ID of the CMK that is used to decrypt the ciphertext.
	//
	// This parameter is the globally unique ID of the CMK.
	//
	// >  If you set CiphertextBlob to the public key-encrypted ciphertext that is returned after an asymmetric encryption operation, specify this parameter.
	SourceKeyId *string `json:"SourceKeyId,omitempty" xml:"SourceKeyId,omitempty"`
	// The ID of the CMK version that is used to decrypt the ciphertext.
	//
	// >  If you set CiphertextBlob to the public key-encrypted ciphertext that is returned after an asymmetric encryption operation, specify this parameter.
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
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The ID of the CMK that is used to decrypt the original ciphertext.
	//
	// This parameter is the globally unique ID of the CMK.
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the CMK version that is used to decrypt the original ciphertext.
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The ID of the request.
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReEncryptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the secret.
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RestoreSecretResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The version number of the secret after the secret is rotated.
	//
	// >  The version number is used to ensure the idempotence of the request. Secrets Manager uses this version number to prevent your application from creating the same version of the secret when the application retries a request. If a version number already exists, Secrets Manager ignores the request for rotation and returns a success message.
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
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the secret.
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The version number of the secret after the secret is rotated.
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RotateSecretResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The scheduled period after which the CMK is deleted. During this period, the CMK is in the PendingDeletion state. After this period ends, you cannot cancel the key deletion task.
	//
	// Valid values: 7 to 366.
	//
	// Unit: days.
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ScheduleKeyDeletionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DeletionProtectionDescription *string `json:"DeletionProtectionDescription,omitempty" xml:"DeletionProtectionDescription,omitempty"`
	// Specifies whether to enable deletion protection. Valid values:
	//
	// *   true: enables deletion protection.
	// *   false: disables deletion protection.
	EnableDeletionProtection *bool `json:"EnableDeletionProtection,omitempty" xml:"EnableDeletionProtection,omitempty"`
	// The ARN of the CMK for which you want to set deletion protection.
	//
	// You can call the [DescribeKey](~~28952~~) operation to query the CMK ARN.
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDeletionProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type TagResourceRequest struct {
	// The ID of the certificate.
	//
	// >  You can configure only one of the KeyId, SecretName, and CertificateId parameters.
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The ID of the customer master key (CMK). The ID must be globally unique.
	//
	// >  You can configure only one of the KeyId, SecretName, and CertificateId parameters.
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the secret.
	//
	// >  You can configure only one of the KeyId, SecretName, and CertificateId parameters.
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// One or more tags that you want to add. The value is in the array format.
	//
	// Tag attributes:
	//
	// *   TagKey: the tag key.
	// *   TagValue: the tag value.
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UntagResourceRequest struct {
	// The ID of the certificate.
	//
	// >  You can configure only one of the KeyId, SecretName, and CertificateId parameters.
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The ID of the customer master key (CMK). The ID must be globally unique.
	//
	// >  You can configure only one of the KeyId, SecretName, and CertificateId parameters.
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the secret.
	//
	// >  You can configure only one of the KeyId, SecretName, and CertificateId parameters.
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// One or more tag keys. Separate multiple tag keys with commas (,).
	//
	// You need to specify only the tag keys, not the tag values.
	//
	// Each tag key must be 1 to 128 bytes in length.
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
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UpdateAliasRequest struct {
	// The alias that you want to bind.
	//
	// The value must be 1 to 255 characters in length and must include the alias/ prefix.
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAliasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UpdateCertificateStatusRequest struct {
	// The ID of the certificate. The ID must be globally unique in Certificates Manager.
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The status of the certificate. Valid values:
	//
	// *   INACTIVE: The certificate is disabled.
	//
	// *   ACTIVE: The certificate is enabled.
	//
	// *   REVOKED: The certificate is revoked.
	//
	// > If the certificate is in the REVOKED state, you can use the certificate only to verify a signature, but not to generate a signature.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateCertificateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateKeyDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UpdateRotationPolicyRequest struct {
	// Specifies whether to enable automatic key rotation. Valid values:
	//
	// *   true: enables automatic key rotation.
	// *   false: disables automatic key rotation.
	EnableAutomaticRotation *bool `json:"EnableAutomaticRotation,omitempty" xml:"EnableAutomaticRotation,omitempty"`
	// The ID of the customer master key (CMK). The ID must be globally unique.
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The period of automatic key rotation. Specify the value in the integer\[unit] format. The following units are supported: d (day), h (hour), m (minute), and s (second). For example, you can use either 7d or 604800s to specify a seven-day period. The period can range from 7 days to 730 days.
	//
	// >  If you set the EnableAutomaticRotation parameter to true, you must also specify this parameter. If you set the EnableAutomaticRotation parameter to false, you can leave this parameter unspecified.
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRotationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the secret.
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
	// >
	// *   If this parameter is specified, the existing extended configuration of the secret is updated.
	// *   This parameter is unavailable for generic secrets.
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
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the secret.
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
	// >
	// *   If this parameter is specified, the existing extended configuration of the secret is updated.
	// *   This parameter is unavailable for generic secrets.
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the secret.
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSecretResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// *   true: enables automatic rotation.
	// *   false: does not enable automatic rotation. This is the default value.
	EnableAutomaticRotation *bool `json:"EnableAutomaticRotation,omitempty" xml:"EnableAutomaticRotation,omitempty"`
	// The interval for automatic rotation. Valid values: 6 hours to 8,760 hours (365 days).
	//
	// The value is in the `integer[unit]` format.````
	//
	// The unit can be d (day), h (hour), m (minute), or s (second). For example, both 7d and 604800s indicate a seven-day interval.
	//
	// >  This parameter is required if you set the EnableAutomaticRotation parameter to true. This parameter is ignored if you set the EnableAutomaticRotation parameter to false or does not specify the EnableAutomaticRotation parameter.
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	// The name of the secret.
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the secret.
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSecretRotationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The version to which you want to apply the specified stage label.
	//
	// >
	// *   You must specify at least one of the RemoveFromVersion and MoveToVersion parameters.
	// *   If the VersionStage parameter is set to ACSCurrent or ACSPrevious, this parameter is required.
	MoveToVersion *string `json:"MoveToVersion,omitempty" xml:"MoveToVersion,omitempty"`
	// The version from which you want to remove the specified stage label.
	//
	// >  You must specify at least one of the RemoveFromVersion and MoveToVersion parameters.
	RemoveFromVersion *string `json:"RemoveFromVersion,omitempty" xml:"RemoveFromVersion,omitempty"`
	// The name of the secret.
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The specified stage label. Valid values:
	//
	// *   ACSCurrent
	// *   ACSPrevious
	// *   Custom stage label
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
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the secret.
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSecretVersionStageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The certificate chain issued by the CA, which is in the PEM format.
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The ID of the certificate. The ID must be globally unique in Certificates Manager.
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

/**
 * This operation supports only asymmetric keys for which the **Usage** parameter is set to **ENCRYPT/DECRYPT**. The following table lists supported encryption algorithms.
 * | KeySpec | Algorithm | Description | Maximum length in bytes |
 * | ------- | --------- | ----------- | ----------------------- |
 * | RSA_2048 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256 | 256 |
 * | RSA_2048 | RSAES_OAEP_SHA_1 | RSAES-OAEP using SHA1 and MGF1 with SHA1 | 256 |
 * | RSA_3072 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256 | 384 |
 * | RSA_3072 | RSAES_OAEP_SHA_1 | RSAES-OAEP using SHA1 and MGF1 with SHA1 | 384 |
 * | EC_SM2 | SM2PKE | SM2 public key encryption algorithm based on elliptic curves | 6144 |
 * In this example, the asymmetric key whose ID is `5c438b18-05be-40ad-b6c2-3be6752c****` and version ID is `2ab1a983-7072-4bbc-a582-584b5bd8****` and the decryption algorithm `RSAES_OAEP_SHA_1` are used to decrypt the ciphertext `BQKP+1zK6+ZEMxTP5qaVzcsgXtWplYBKm0NXdSnB5FzliFxE1bSiu4dnEIlca2JpeH7yz1/S6fed630H+hIH6DoM25fTLNcKj+mFB0Xnh9m2+HN59Mn4qyTfcUeadnfCXSWcGBouhXFwcdd2rJ3n337bzTf4jm659gZu3L0i6PLuxM9p7mqdwO0cKJPfGVfhnfMz+f4alMg79WB/NNyE2lyX7/qxvV49ObNrrJbKSFiz8Djocaf0IESNLMbfYI5bXjWkJlX92DQbKhibtQW8ZOJ//ZC6t0AWcUoKL6QDm/dg5koQalcleRinpB+QadFm894sLbVZ9+N4GVsv1W****==`.
 *
 * @param request AsymmetricDecryptRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AsymmetricDecryptResponse
 */
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
	_result = &AsymmetricDecryptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation supports only asymmetric keys for which the **Usage** parameter is set to **ENCRYPT/DECRYPT**. The following table lists supported encryption algorithms.
 * | KeySpec | Algorithm | Description | Maximum length in bytes |
 * | ------- | --------- | ----------- | ----------------------- |
 * | RSA_2048 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256 | 256 |
 * | RSA_2048 | RSAES_OAEP_SHA_1 | RSAES-OAEP using SHA1 and MGF1 with SHA1 | 256 |
 * | RSA_3072 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256 | 384 |
 * | RSA_3072 | RSAES_OAEP_SHA_1 | RSAES-OAEP using SHA1 and MGF1 with SHA1 | 384 |
 * | EC_SM2 | SM2PKE | SM2 public key encryption algorithm based on elliptic curves | 6144 |
 * In this example, the asymmetric key whose ID is `5c438b18-05be-40ad-b6c2-3be6752c****` and version ID is `2ab1a983-7072-4bbc-a582-584b5bd8****` and the decryption algorithm `RSAES_OAEP_SHA_1` are used to decrypt the ciphertext `BQKP+1zK6+ZEMxTP5qaVzcsgXtWplYBKm0NXdSnB5FzliFxE1bSiu4dnEIlca2JpeH7yz1/S6fed630H+hIH6DoM25fTLNcKj+mFB0Xnh9m2+HN59Mn4qyTfcUeadnfCXSWcGBouhXFwcdd2rJ3n337bzTf4jm659gZu3L0i6PLuxM9p7mqdwO0cKJPfGVfhnfMz+f4alMg79WB/NNyE2lyX7/qxvV49ObNrrJbKSFiz8Djocaf0IESNLMbfYI5bXjWkJlX92DQbKhibtQW8ZOJ//ZC6t0AWcUoKL6QDm/dg5koQalcleRinpB+QadFm894sLbVZ9+N4GVsv1W****==`.
 *
 * @param request AsymmetricDecryptRequest
 * @return AsymmetricDecryptResponse
 */
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

/**
 * This operation is supported only for asymmetric keys for which the **Usage** parameter is set to **ENCRYPT/DECRYPT**. The following table lists the supported encryption algorithms:
 * | KeySpec | Algorithm | Description | Maximum number of bytes that can be encrypted |
 * | ------- | --------- | ----------- | --------------------------------------------- |
 * | RSA_2048 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256 | 190 |
 * | RSA_2048 | RSAES_OAEP_SHA_1 | RSAES-OAEP using SHA1 and MGF1 with SHA1 | 214 |
 * | RSA_3072 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256 | 318 |
 * | RSA_3072 | RSAES_OAEP_SHA_1 | RSAES-OAEP using SHA1 and MGF1 with SHA1 | 342 |
 * | EC_SM2 | SM2PKE | SM2 public key encryption algorithm based on elliptic curves | 6047 |
 * You can use the asymmetric CMK whose ID is `5c438b18-05be-40ad-b6c2-3be6752c****` and version ID is `2ab1a983-7072-4bbc-a582-584b5bd8****` and the algorithm `RSAES_OAEP_SHA_1` to encrypt the plaintext `SGVsbG8gd29ybGQ=` based on the parameter settings provided in this topic.
 *
 * @param request AsymmetricEncryptRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AsymmetricEncryptResponse
 */
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
	_result = &AsymmetricEncryptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation is supported only for asymmetric keys for which the **Usage** parameter is set to **ENCRYPT/DECRYPT**. The following table lists the supported encryption algorithms:
 * | KeySpec | Algorithm | Description | Maximum number of bytes that can be encrypted |
 * | ------- | --------- | ----------- | --------------------------------------------- |
 * | RSA_2048 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256 | 190 |
 * | RSA_2048 | RSAES_OAEP_SHA_1 | RSAES-OAEP using SHA1 and MGF1 with SHA1 | 214 |
 * | RSA_3072 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256 | 318 |
 * | RSA_3072 | RSAES_OAEP_SHA_1 | RSAES-OAEP using SHA1 and MGF1 with SHA1 | 342 |
 * | EC_SM2 | SM2PKE | SM2 public key encryption algorithm based on elliptic curves | 6047 |
 * You can use the asymmetric CMK whose ID is `5c438b18-05be-40ad-b6c2-3be6752c****` and version ID is `2ab1a983-7072-4bbc-a582-584b5bd8****` and the algorithm `RSAES_OAEP_SHA_1` to encrypt the plaintext `SGVsbG8gd29ybGQ=` based on the parameter settings provided in this topic.
 *
 * @param request AsymmetricEncryptRequest
 * @return AsymmetricEncryptResponse
 */
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

/**
 * This operation supports only asymmetric keys for which the **Usage** parameter is set to **SIGN/VERIFY**. The following table describes the supported signature algorithms.
 * | KeySpec | Algorithm | Description |
 * | ------- | --------- | ----------- |
 * | RSA_2048 | RSA_PSS_SHA_256 | RSASSA-PSS using SHA-256 and MGF1 with SHA-256 |
 * | RSA_2048 | RSA_PKCS1_SHA_256 | RSASSA-PKCS1-v1_5 using SHA-256 |
 * | RSA_3072 | RSA_PSS_SHA_256 | RSASSA-PSS using SHA-256 and MGF1 with SHA-256 |
 * | RSA_3072 | RSA_PKCS1_SHA_256 | RSASSA-PKCS1-v1_5 using SHA-256 |
 * | EC_P256 | ECDSA_SHA_256 | ECDSA on the P-256 Curve(secp256r1) with a SHA-256 digest |
 * | EC_P256K | ECDSA_SHA_256 | ECDSA on the P-256K Curve(secp256k1) with a SHA-256 digest |
 * | EC_SM2 | SM2DSA | SM2 public key encryption algorithm based on elliptic curves cryptography (ECC) |
 * >  According to GB/T 32918.2 "Information security technology-Public key cryptographic algorithm SM2 based on elliptic curves-Part 2: Digital signature algorithm", when you calculate the SM2 signature, the **Digest** parameter is used to calculate the digest value of the combination of Z(A) and M, rather than the SM3 digest value. M indicates the original message to be signed. Z(A) indicates the hash value for User A. The hash value is defined in GB/T GB/T 32918.2. In this example, the asymmetric key whose ID is `5c438b18-05be-40ad-b6c2-3be6752c****` and version ID is `2ab1a983-7072-4bbc-a582-584b5bd8****` and the signature algorithm `RSA_PSS_SHA_256` are used to generate a signature for the digest `ZOyIygCyaOW6GjVnihtTFtIS9PNmskdyMlNKiuy****=`.
 *
 * @param request AsymmetricSignRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AsymmetricSignResponse
 */
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
	_result = &AsymmetricSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation supports only asymmetric keys for which the **Usage** parameter is set to **SIGN/VERIFY**. The following table describes the supported signature algorithms.
 * | KeySpec | Algorithm | Description |
 * | ------- | --------- | ----------- |
 * | RSA_2048 | RSA_PSS_SHA_256 | RSASSA-PSS using SHA-256 and MGF1 with SHA-256 |
 * | RSA_2048 | RSA_PKCS1_SHA_256 | RSASSA-PKCS1-v1_5 using SHA-256 |
 * | RSA_3072 | RSA_PSS_SHA_256 | RSASSA-PSS using SHA-256 and MGF1 with SHA-256 |
 * | RSA_3072 | RSA_PKCS1_SHA_256 | RSASSA-PKCS1-v1_5 using SHA-256 |
 * | EC_P256 | ECDSA_SHA_256 | ECDSA on the P-256 Curve(secp256r1) with a SHA-256 digest |
 * | EC_P256K | ECDSA_SHA_256 | ECDSA on the P-256K Curve(secp256k1) with a SHA-256 digest |
 * | EC_SM2 | SM2DSA | SM2 public key encryption algorithm based on elliptic curves cryptography (ECC) |
 * >  According to GB/T 32918.2 "Information security technology-Public key cryptographic algorithm SM2 based on elliptic curves-Part 2: Digital signature algorithm", when you calculate the SM2 signature, the **Digest** parameter is used to calculate the digest value of the combination of Z(A) and M, rather than the SM3 digest value. M indicates the original message to be signed. Z(A) indicates the hash value for User A. The hash value is defined in GB/T GB/T 32918.2. In this example, the asymmetric key whose ID is `5c438b18-05be-40ad-b6c2-3be6752c****` and version ID is `2ab1a983-7072-4bbc-a582-584b5bd8****` and the signature algorithm `RSA_PSS_SHA_256` are used to generate a signature for the digest `ZOyIygCyaOW6GjVnihtTFtIS9PNmskdyMlNKiuy****=`.
 *
 * @param request AsymmetricSignRequest
 * @return AsymmetricSignResponse
 */
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

/**
 * This operation supports only asymmetric keys for which the **Usage** parameter is set to **SIGN/VERIFY**. The following table describes the supported signature algorithms.
 * | KeySpec | Algorithm | Description |
 * | ------- | --------- | ----------- |
 * | RSA_2048 | RSA_PSS_SHA_256 | RSASSA-PSS using SHA-256 and MGF1 with SHA-256 |
 * | RSA_2048 | RSA_PKCS1_SHA_256 | RSASSA-PKCS1-v1_5 using SHA-256 |
 * | RSA_3072 | RSA_PSS_SHA_256 | RSASSA-PSS using SHA-256 and MGF1 with SHA-256 |
 * | RSA_3072 | RSA_PKCS1_SHA_256 | RSASSA-PKCS1-v1_5 using SHA-256 |
 * | EC_P256 | ECDSA_SHA_256 | ECDSA on the P-256 Curve(secp256r1) with a SHA-256 digest |
 * | EC_P256K | ECDSA_SHA_256 | ECDSA on the P-256K Curve(secp256k1) with a SHA-256 digest |
 * | EC_SM2 | SM2DSA | SM2 elliptic curve public key encryption algorithm |
 * >  When you calculate the SM2 signature based on GB/T 32918, the **Digest** parameter is used to calculate the digest value of the combination of Z(A) and M, rather than the SM3 digest value. M indicates the original message to be signed. Z(A) indicates the hash value for User A. The hash value is defined in GB/T 32918.  In this example, the asymmetric key whose ID is `5c438b18-05be-40ad-b6c2-3be6752c****` and version ID is `2ab1a983-7072-4bbc-a582-584b5bd8****` and the signature algorithm RSA_PSS_SHA_256 are used to verify the signature `M2CceNZH00ZgL9ED/ZHFp21YRAvYeZHknJUc207OCZ0N9wNn9As4z2bON3FF3je+1Nu+2+/8Zj50HpMTpzYpMp2R93cYmACCmhaYoKydxylbyGzJR8y9likZRCrkD38lRoS40aBBvv/6iRKzQuo9EGYVcel36cMNg00VmYNBy3pa1rwg3gA4l3cy6kjayZja1WGPkVhrVKsrJMdbpl0ApLjXKuD8rw1n1XLCwCUEL5eLPljTZaAveqdOFQOiZnZEGI27qIiZe7I1fN8tcz6anS/gTM7xRKE++5egEvRWlTQQTJeApnPSiUPA+8ZykNdelQsOQh5SrGoyI4A5pq****==` of the digest `ZOyIygCyaOW6GjVnihtTFtIS9PNmskdyMlNKiuyjfzw=`.
 *
 * @param request AsymmetricVerifyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AsymmetricVerifyResponse
 */
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
	_result = &AsymmetricVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation supports only asymmetric keys for which the **Usage** parameter is set to **SIGN/VERIFY**. The following table describes the supported signature algorithms.
 * | KeySpec | Algorithm | Description |
 * | ------- | --------- | ----------- |
 * | RSA_2048 | RSA_PSS_SHA_256 | RSASSA-PSS using SHA-256 and MGF1 with SHA-256 |
 * | RSA_2048 | RSA_PKCS1_SHA_256 | RSASSA-PKCS1-v1_5 using SHA-256 |
 * | RSA_3072 | RSA_PSS_SHA_256 | RSASSA-PSS using SHA-256 and MGF1 with SHA-256 |
 * | RSA_3072 | RSA_PKCS1_SHA_256 | RSASSA-PKCS1-v1_5 using SHA-256 |
 * | EC_P256 | ECDSA_SHA_256 | ECDSA on the P-256 Curve(secp256r1) with a SHA-256 digest |
 * | EC_P256K | ECDSA_SHA_256 | ECDSA on the P-256K Curve(secp256k1) with a SHA-256 digest |
 * | EC_SM2 | SM2DSA | SM2 elliptic curve public key encryption algorithm |
 * >  When you calculate the SM2 signature based on GB/T 32918, the **Digest** parameter is used to calculate the digest value of the combination of Z(A) and M, rather than the SM3 digest value. M indicates the original message to be signed. Z(A) indicates the hash value for User A. The hash value is defined in GB/T 32918.  In this example, the asymmetric key whose ID is `5c438b18-05be-40ad-b6c2-3be6752c****` and version ID is `2ab1a983-7072-4bbc-a582-584b5bd8****` and the signature algorithm RSA_PSS_SHA_256 are used to verify the signature `M2CceNZH00ZgL9ED/ZHFp21YRAvYeZHknJUc207OCZ0N9wNn9As4z2bON3FF3je+1Nu+2+/8Zj50HpMTpzYpMp2R93cYmACCmhaYoKydxylbyGzJR8y9likZRCrkD38lRoS40aBBvv/6iRKzQuo9EGYVcel36cMNg00VmYNBy3pa1rwg3gA4l3cy6kjayZja1WGPkVhrVKsrJMdbpl0ApLjXKuD8rw1n1XLCwCUEL5eLPljTZaAveqdOFQOiZnZEGI27qIiZe7I1fN8tcz6anS/gTM7xRKE++5egEvRWlTQQTJeApnPSiUPA+8ZykNdelQsOQh5SrGoyI4A5pq****==` of the digest `ZOyIygCyaOW6GjVnihtTFtIS9PNmskdyMlNKiuyjfzw=`.
 *
 * @param request AsymmetricVerifyRequest
 * @return AsymmetricVerifyResponse
 */
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

/**
 * If the deletion task of a CMK is canceled, the CMK returns to the Enabled state.
 *
 * @param request CancelKeyDeletionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CancelKeyDeletionResponse
 */
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
	_result = &CancelKeyDeletionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If the deletion task of a CMK is canceled, the CMK returns to the Enabled state.
 *
 * @param request CancelKeyDeletionRequest
 * @return CancelKeyDeletionResponse
 */
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

/**
 * Limit: The encryption algorithm in the request parameters must match the key type.
 * The following table describes the mapping between encryption algorithms and key types.
 * | Algorithm | Key Spec |
 * | --------- | -------- |
 * | RSAES_OAEP_SHA_1 | RSA_2048 |
 * | RSAES_OAEP_SHA_256 | RSA_2048 |
 * | SM2PKE | EC_SM2 |
 * In this example, the certificate whose ID is `12345678-1234-1234-1234-12345678****` and the encryption algorithm `RSAES_OAEP_SHA_256` are used to decrypt the data `ZOyIygCyaOW6Gj****MlNKiuyjfzw=`.
 *
 * @param request CertificatePrivateKeyDecryptRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CertificatePrivateKeyDecryptResponse
 */
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
	_result = &CertificatePrivateKeyDecryptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Limit: The encryption algorithm in the request parameters must match the key type.
 * The following table describes the mapping between encryption algorithms and key types.
 * | Algorithm | Key Spec |
 * | --------- | -------- |
 * | RSAES_OAEP_SHA_1 | RSA_2048 |
 * | RSAES_OAEP_SHA_256 | RSA_2048 |
 * | SM2PKE | EC_SM2 |
 * In this example, the certificate whose ID is `12345678-1234-1234-1234-12345678****` and the encryption algorithm `RSAES_OAEP_SHA_256` are used to decrypt the data `ZOyIygCyaOW6Gj****MlNKiuyjfzw=`.
 *
 * @param request CertificatePrivateKeyDecryptRequest
 * @return CertificatePrivateKeyDecryptResponse
 */
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

/**
 * The signature algorithm in the request parameters must match the key type. The following table describes the mapping between signature algorithms and key types.
 * | Algorithm | Key Spec |
 * | --------- | -------- |
 * | RSA_PKCS1_SHA_256 | RSA_2048 |
 * | RSA_PSS_SHA_256 | RSA_2048 |
 * | ECDSA_SHA_256 | EC_P256 |
 * | SM2DSA | EC_SM2 |
 * In this example, the certificate whose ID is `12345678-1234-1234-1234-12345678****` and the signature algorithm `ECDSA_SHA_256` are used to generate a signature for the raw data `VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4=`.
 *
 * @param request CertificatePrivateKeySignRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CertificatePrivateKeySignResponse
 */
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
	_result = &CertificatePrivateKeySignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The signature algorithm in the request parameters must match the key type. The following table describes the mapping between signature algorithms and key types.
 * | Algorithm | Key Spec |
 * | --------- | -------- |
 * | RSA_PKCS1_SHA_256 | RSA_2048 |
 * | RSA_PSS_SHA_256 | RSA_2048 |
 * | ECDSA_SHA_256 | EC_P256 |
 * | SM2DSA | EC_SM2 |
 * In this example, the certificate whose ID is `12345678-1234-1234-1234-12345678****` and the signature algorithm `ECDSA_SHA_256` are used to generate a signature for the raw data `VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4=`.
 *
 * @param request CertificatePrivateKeySignRequest
 * @return CertificatePrivateKeySignResponse
 */
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

/**
 * Limit: The encryption algorithm in the request parameters must match the key type.
 * The following table describes the mapping between encryption algorithms and key types.
 * | Algorithm | Key Spec |
 * | --------- | -------- |
 * | RSAES_OAEP_SHA_1 | RSA_2048 |
 * | RSAES_OAEP_SHA_256 | RSA_2048 |
 * | SM2PKE | EC_SM2 |
 * In this example, the certificate whose ID is `12345678-1234-1234-1234-12345678****` and the encryption algorithm `RSAES_OAEP_SHA_256` are used to encrypt the data `VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4=`.
 *
 * @param request CertificatePublicKeyEncryptRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CertificatePublicKeyEncryptResponse
 */
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
	_result = &CertificatePublicKeyEncryptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Limit: The encryption algorithm in the request parameters must match the key type.
 * The following table describes the mapping between encryption algorithms and key types.
 * | Algorithm | Key Spec |
 * | --------- | -------- |
 * | RSAES_OAEP_SHA_1 | RSA_2048 |
 * | RSAES_OAEP_SHA_256 | RSA_2048 |
 * | SM2PKE | EC_SM2 |
 * In this example, the certificate whose ID is `12345678-1234-1234-1234-12345678****` and the encryption algorithm `RSAES_OAEP_SHA_256` are used to encrypt the data `VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4=`.
 *
 * @param request CertificatePublicKeyEncryptRequest
 * @return CertificatePublicKeyEncryptResponse
 */
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

/**
 * The signature algorithm in the request parameters must match the key type. The following table describes the mapping between signature algorithms and key types.
 * | Algorithm | Key Spec |
 * | --------- | -------- |
 * | RSA_PKCS1_SHA_256 | RSA_2048 |
 * | RSA_PSS_SHA_256 | RSA_2048 |
 * | ECDSA_SHA_256 | EC_P256 |
 * | SM2DSA | EC_SM2 |
 * In this example, the certificate whose ID is `12345678-1234-1234-1234-12345678****` and the signature algorithm `ECDSA_SHA_256` are used to verify the digital signature `ZOyIygCyaOW6Gj****MlNKiuyjfzw=` of the raw data `VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4=`.
 *
 * @param request CertificatePublicKeyVerifyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CertificatePublicKeyVerifyResponse
 */
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
	_result = &CertificatePublicKeyVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The signature algorithm in the request parameters must match the key type. The following table describes the mapping between signature algorithms and key types.
 * | Algorithm | Key Spec |
 * | --------- | -------- |
 * | RSA_PKCS1_SHA_256 | RSA_2048 |
 * | RSA_PSS_SHA_256 | RSA_2048 |
 * | ECDSA_SHA_256 | EC_P256 |
 * | SM2DSA | EC_SM2 |
 * In this example, the certificate whose ID is `12345678-1234-1234-1234-12345678****` and the signature algorithm `ECDSA_SHA_256` are used to verify the digital signature `ZOyIygCyaOW6Gj****MlNKiuyjfzw=` of the raw data `VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4=`.
 *
 * @param request CertificatePublicKeyVerifyRequest
 * @return CertificatePublicKeyVerifyResponse
 */
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

/**
 * *   Each alias can be bound to only one CMK at a time.
 * *   The aliases of CMKs in the same region must be unique.
 * In this topic, an alias named `alias/example` is created for a CMK named `7906979c-8e06-46a2-be2d-68e3ccbc****`.
 *
 * @param request CreateAliasRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateAliasResponse
 */
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
	_result = &CreateAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   Each alias can be bound to only one CMK at a time.
 * *   The aliases of CMKs in the same region must be unique.
 * In this topic, an alias named `alias/example` is created for a CMK named `7906979c-8e06-46a2-be2d-68e3ccbc****`.
 *
 * @param request CreateAliasRequest
 * @return CreateAliasResponse
 */
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

/**
 * To create a certificate, you must specify the type of the asymmetric key. Certificates Manager generates a private key and returns a certificate signing request (CSR). Submit the CSR in the Privacy Enhanced Mail (PEM) format to a certificate authority (CA) to obtain the formal certificate and certificate chain. Then, call the [UploadCertificate](~~212136~~) operation to import the certificate into Certificates Manager.
 * In this example, a certificate is created and the CSR is obtained.
 *
 * @param tmpReq CreateCertificateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateCertificateResponse
 */
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
	_result = &CreateCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * To create a certificate, you must specify the type of the asymmetric key. Certificates Manager generates a private key and returns a certificate signing request (CSR). Submit the CSR in the Privacy Enhanced Mail (PEM) format to a certificate authority (CA) to obtain the formal certificate and certificate chain. Then, call the [UploadCertificate](~~212136~~) operation to import the certificate into Certificates Manager.
 * In this example, a certificate is created and the CSR is obtained.
 *
 * @param request CreateCertificateRequest
 * @return CreateCertificateResponse
 */
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

/**
 * A CMK can be symmetric or asymmetric. Symmetric CMKs are used to generate data keys that can be used to encrypt large amounts of data. You can also use symmetric CMKs to encrypt small volume of data that is less than 6 KB. For more information, see [GenerateDataKey](https://www.alibabacloud.com/help/en/key-management-service/latest/generatedatakey). Asymmetric CMKs are used to encrypt data, decrypt data, generate digital signatures, and verify digital signatures. However, you cannot use asymmetric CMKs to generate data keys.
 * The following table describes different types of CMKs and the operations that are supported by the CMKs.
 * | CMK category | CMK type | Description | Encryption and decryption | Signature generation and verification |
 * | ------------ | -------- | ----------- | ------------------------- | ------------------------------------- |
 * | Symmetric CMK | Aliyun_AES_256 | An advanced Encryption Standard (AES) CMK with a length of 256 bits. | Yes | No |
 * | Symmetric CMK | Aliyun_AES_128 | An AES CMK with a length of 128 bits. Only Dedicated KMS supports this CMK type. | Yes | No |
 * | Symmetric CMK | Aliyun_AES_192 | An AES CMK with a length of 192 bits. Only Dedicated KMS supports this CMK type. | Yes | No |
 * | Symmetric CMK | Aliyun_SM4 | SM4 CMK. | Yes | No |
 * | Asymmetric CMK | RSA_2048 | Rivest-Shamir-Adleman (RSA) CMK with a length of 2,048 bits. | Supported | Supported |
 * | Asymmetric CMK | RSA_3072 | RSA CMK with a length of 3,072 bits. | Supported | Supported |
 * | Asymmetric CMK | EC_P256 | National Institute of Standards and Technology (NIST)-recommended elliptic curve P-256 (secp256r1). | Not supported | Supported |
 * | Asymmetric CMK | EC_P256K | Standards for Efficient Cryptography Group (SECG) elliptic curve secp256k1 | Not supported | Supported |
 * | Asymmetric CMK | EC_SM2 | 256-bit elliptic curves over the prime field that is defined in GB/T 32918. | Supported | Supported |
 * > - If the value of the KeySpec parameter that is used to create a symmetric CMK is prefixed with `Aliyun_`, a standard cryptographic algorithm is used, but non-standard ciphertext is generated. An asymmetric CMK can be used to generate standard ciphertext or signatures.
 * - You can use an RSA CMK to perform one of the two types of operations: encrypt and decrypt data, and generate and verify signatures. You cannot use the RSA CMK to perform both two types of operations.
 *
 * @param request CreateKeyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateKeyResponse
 */
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

	if !tea.BoolValue(util.IsUnset(request.ProtectionLevel)) {
		query["ProtectionLevel"] = request.ProtectionLevel
	}

	if !tea.BoolValue(util.IsUnset(request.RotationInterval)) {
		query["RotationInterval"] = request.RotationInterval
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
	_result = &CreateKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * A CMK can be symmetric or asymmetric. Symmetric CMKs are used to generate data keys that can be used to encrypt large amounts of data. You can also use symmetric CMKs to encrypt small volume of data that is less than 6 KB. For more information, see [GenerateDataKey](https://www.alibabacloud.com/help/en/key-management-service/latest/generatedatakey). Asymmetric CMKs are used to encrypt data, decrypt data, generate digital signatures, and verify digital signatures. However, you cannot use asymmetric CMKs to generate data keys.
 * The following table describes different types of CMKs and the operations that are supported by the CMKs.
 * | CMK category | CMK type | Description | Encryption and decryption | Signature generation and verification |
 * | ------------ | -------- | ----------- | ------------------------- | ------------------------------------- |
 * | Symmetric CMK | Aliyun_AES_256 | An advanced Encryption Standard (AES) CMK with a length of 256 bits. | Yes | No |
 * | Symmetric CMK | Aliyun_AES_128 | An AES CMK with a length of 128 bits. Only Dedicated KMS supports this CMK type. | Yes | No |
 * | Symmetric CMK | Aliyun_AES_192 | An AES CMK with a length of 192 bits. Only Dedicated KMS supports this CMK type. | Yes | No |
 * | Symmetric CMK | Aliyun_SM4 | SM4 CMK. | Yes | No |
 * | Asymmetric CMK | RSA_2048 | Rivest-Shamir-Adleman (RSA) CMK with a length of 2,048 bits. | Supported | Supported |
 * | Asymmetric CMK | RSA_3072 | RSA CMK with a length of 3,072 bits. | Supported | Supported |
 * | Asymmetric CMK | EC_P256 | National Institute of Standards and Technology (NIST)-recommended elliptic curve P-256 (secp256r1). | Not supported | Supported |
 * | Asymmetric CMK | EC_P256K | Standards for Efficient Cryptography Group (SECG) elliptic curve secp256k1 | Not supported | Supported |
 * | Asymmetric CMK | EC_SM2 | 256-bit elliptic curves over the prime field that is defined in GB/T 32918. | Supported | Supported |
 * > - If the value of the KeySpec parameter that is used to create a symmetric CMK is prefixed with `Aliyun_`, a standard cryptographic algorithm is used, but non-standard ciphertext is generated. An asymmetric CMK can be used to generate standard ciphertext or signatures.
 * - You can use an RSA CMK to perform one of the two types of operations: encrypt and decrypt data, and generate and verify signatures. You cannot use the RSA CMK to perform both two types of operations.
 *
 * @param request CreateKeyRequest
 * @return CreateKeyResponse
 */
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

/**
 * *   You can create a version only for an asymmetric CMK that is in the Enabled state. You can call the [CreateKey](~~28947~~) operation to create an asymmetric CMK and the [DescribeKey](~~28952~~) operation to query the status of the CMK. The status is specified by the KeyState parameter.
 * *   The minimum interval for creating a version of the same CMK is seven days. You can call the [DescribeKey](~~28952~~) operation to query the time when the last version of a CMK was created. The time is specified by the LastRotationDate parameter.
 * *   If a CMK is in a private key store, you cannot create a version for the CMK.
 * *   You can create a maximum of 50 versions for a CMK in the same region.
 * You can create a version for the CMK whose ID is `0b30658a-ed1a-4922-b8f7-a673ca9c****` by using the parameter settings provided in this topic.
 *
 * @param request CreateKeyVersionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateKeyVersionResponse
 */
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
	_result = &CreateKeyVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   You can create a version only for an asymmetric CMK that is in the Enabled state. You can call the [CreateKey](~~28947~~) operation to create an asymmetric CMK and the [DescribeKey](~~28952~~) operation to query the status of the CMK. The status is specified by the KeyState parameter.
 * *   The minimum interval for creating a version of the same CMK is seven days. You can call the [DescribeKey](~~28952~~) operation to query the time when the last version of a CMK was created. The time is specified by the LastRotationDate parameter.
 * *   If a CMK is in a private key store, you cannot create a version for the CMK.
 * *   You can create a maximum of 50 versions for a CMK in the same region.
 * You can create a version for the CMK whose ID is `0b30658a-ed1a-4922-b8f7-a673ca9c****` by using the parameter settings provided in this topic.
 *
 * @param request CreateKeyVersionRequest
 * @return CreateKeyVersionResponse
 */
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

/**
 * You must specify the secret name, the secret value that is stored in the initial version, and the version number. The initial version is labeled as ACSCurrent.
 * You can specify a symmetric customer master key (CMK) as the encryption key to encrypt the secret value. If you do not specify an encryption key, Secrets Manager creates a CMK to encrypt the secret value. This CMK is used as the default encryption key for all secrets that are created by your Alibaba Cloud account in the current region. Secrets Manager encrypts only the secret value of each version. Secrets Manager does not encrypt the metadata such as the secret name, version number, or state label.
 * To use a specified CMK to encrypt the secret value, you must have the `kms:GenerateDataKey` permission on the CMK.
 * In this example, a generic secret named `mydbconninfo` is created. The initial version number of the secret is specified in the `VersionId` parameter, and the value is `v1`. The secret value is specified in the `SecretData` parameter, and the value is `{"user":"root","passwd":"****"}`.
 *
 * @param tmpReq CreateSecretRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateSecretResponse
 */
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
	_result = &CreateSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You must specify the secret name, the secret value that is stored in the initial version, and the version number. The initial version is labeled as ACSCurrent.
 * You can specify a symmetric customer master key (CMK) as the encryption key to encrypt the secret value. If you do not specify an encryption key, Secrets Manager creates a CMK to encrypt the secret value. This CMK is used as the default encryption key for all secrets that are created by your Alibaba Cloud account in the current region. Secrets Manager encrypts only the secret value of each version. Secrets Manager does not encrypt the metadata such as the secret name, version number, or state label.
 * To use a specified CMK to encrypt the secret value, you must have the `kms:GenerateDataKey` permission on the CMK.
 * In this example, a generic secret named `mydbconninfo` is created. The initial version number of the secret is specified in the `VersionId` parameter, and the value is `v1`. The secret value is specified in the `SecretData` parameter, and the value is `{"user":"root","passwd":"****"}`.
 *
 * @param request CreateSecretRequest
 * @return CreateSecretResponse
 */
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
	_result = &DecryptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &DeleteAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

/**
 * After the certificate and its private key and certificate chain are deleted, they cannot be restored. Proceed with caution.
 * In this example, the certificate whose ID is `9a28de48-8d8b-484d-a766-dec4****` and its private key and certificate chain are deleted.
 *
 * @param request DeleteCertificateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteCertificateResponse
 */
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
	_result = &DeleteCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * After the certificate and its private key and certificate chain are deleted, they cannot be restored. Proceed with caution.
 * In this example, the certificate whose ID is `9a28de48-8d8b-484d-a766-dec4****` and its private key and certificate chain are deleted.
 *
 * @param request DeleteCertificateRequest
 * @return DeleteCertificateResponse
 */
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

/**
 * This operation does not delete the CMK that is created by using the key material.
 * If the CMK is in the PendingDeletion state, the state of the CMK and the scheduled deletion time do not change after you call this operation. If the CMK is not in the PendingDeletion state, the state of the CMK changes to PendingImport after you call this operation.
 * After you delete the key material, you can upload only the same key material into the CMK.
 *
 * @param request DeleteKeyMaterialRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteKeyMaterialResponse
 */
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
	_result = &DeleteKeyMaterialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation does not delete the CMK that is created by using the key material.
 * If the CMK is in the PendingDeletion state, the state of the CMK and the scheduled deletion time do not change after you call this operation. If the CMK is not in the PendingDeletion state, the state of the CMK changes to PendingImport after you call this operation.
 * After you delete the key material, you can upload only the same key material into the CMK.
 *
 * @param request DeleteKeyMaterialRequest
 * @return DeleteKeyMaterialResponse
 */
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

/**
 * If you call this operation without specifying a recovery period, the deleted secret can be recovered within 30 days.
 * If you specify a recovery period, the deleted secret can be recovered within the recovery period. You can also forcibly delete a secret. A forcibly deleted secret cannot be recovered.
 *
 * @param request DeleteSecretRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteSecretResponse
 */
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
	_result = &DeleteSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you call this operation without specifying a recovery period, the deleted secret can be recovered within 30 days.
 * If you specify a recovery period, the deleted secret can be recovered within the recovery period. You can also forcibly delete a secret. A forcibly deleted secret cannot be recovered.
 *
 * @param request DeleteSecretRequest
 * @return DeleteSecretResponse
 */
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
	_result = &DescribeAccountKmsStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

/**
 * In this example, the information about the certificate whose ID is `9a28de48-8d8b-484d-a766-dec4****` is queried. The certificate information includes the certificate ID, creation time, certificate issuer, validity period, serial number, and signature algorithm.
 *
 * @param request DescribeCertificateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeCertificateResponse
 */
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
	_result = &DescribeCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the information about the certificate whose ID is `9a28de48-8d8b-484d-a766-dec4****` is queried. The certificate information includes the certificate ID, creation time, certificate issuer, validity period, serial number, and signature algorithm.
 *
 * @param request DescribeCertificateRequest
 * @return DescribeCertificateResponse
 */
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

/**
 * You can query the information about the CMK `05754286-3ba2-4fa6-8d41-4323aca6****` by using parameter settings provided in this topic. The information includes the creator, creation time, status, and deletion protection status of the CMK.
 *
 * @param request DescribeKeyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeKeyResponse
 */
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
	_result = &DescribeKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can query the information about the CMK `05754286-3ba2-4fa6-8d41-4323aca6****` by using parameter settings provided in this topic. The information includes the creator, creation time, status, and deletion protection status of the CMK.
 *
 * @param request DescribeKeyRequest
 * @return DescribeKeyResponse
 */
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

/**
 * This topic provides an example on how to query the information about a version of the CMK `1234abcd-12ab-34cd-56ef-12345678****`. The ID of the CMK version is `2ab1a983-7072-4bbc-a582-584b5bd8****`. The response shows that the creation time of the CMK version is `2016-03-25T10:42:40Z`.
 *
 * @param request DescribeKeyVersionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeKeyVersionResponse
 */
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
	_result = &DescribeKeyVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to query the information about a version of the CMK `1234abcd-12ab-34cd-56ef-12345678****`. The ID of the CMK version is `2ab1a983-7072-4bbc-a582-584b5bd8****`. The response shows that the creation time of the CMK version is `2016-03-25T10:42:40Z`.
 *
 * @param request DescribeKeyVersionRequest
 * @return DescribeKeyVersionResponse
 */
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
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

/**
 * This operation returns the metadata of a secret. This operation does not return the secret value.
 * In this example, the metadata of the secret named `secret001` is queried.
 *
 * @param request DescribeSecretRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeSecretResponse
 */
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
	_result = &DescribeSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation returns the metadata of a secret. This operation does not return the secret value.
 * In this example, the metadata of the secret named `secret001` is queried.
 *
 * @param request DescribeSecretRequest
 * @return DescribeSecretResponse
 */
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

/**
 * If a customer master key (CMK) is disabled, the ciphertext encrypted by using this CMK cannot be decrypted until you re-enable it. You can call the [EnableKey](~~35150~~) operation to enable the CMK.
 * In this example, the CMK whose ID is `1234abcd-12ab-34cd-56ef-12345678****` is disabled.
 *
 * @param request DisableKeyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DisableKeyResponse
 */
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
	_result = &DisableKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If a customer master key (CMK) is disabled, the ciphertext encrypted by using this CMK cannot be decrypted until you re-enable it. You can call the [EnableKey](~~35150~~) operation to enable the CMK.
 * In this example, the CMK whose ID is `1234abcd-12ab-34cd-56ef-12345678****` is disabled.
 *
 * @param request DisableKeyRequest
 * @return DisableKeyResponse
 */
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
	_result = &EnableKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

/**
 * *   KMS uses the primary version of a specified CMK to encrypt data.
 * *   Only data of 6 KB or less can be encrypted. For example, you can call this operation to encrypt RSA keys, database access passwords, or other sensitive information.
 * *   When you migrate encrypted data across regions, you can call this operation in the destination region to encrypt the plaintext of the data key that is used to encrypt the migrated data in the source region. This way, the ciphertext of the data key is generated in the destination region. You can also call the [Decrypt](~~28950~~) operation to decrypt the data key.
 *
 * @param tmpReq EncryptRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return EncryptResponse
 */
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
	_result = &EncryptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   KMS uses the primary version of a specified CMK to encrypt data.
 * *   Only data of 6 KB or less can be encrypted. For example, you can call this operation to encrypt RSA keys, database access passwords, or other sensitive information.
 * *   When you migrate encrypted data across regions, you can call this operation in the destination region to encrypt the plaintext of the data key that is used to encrypt the migrated data in the source region. This way, the ciphertext of the data key is generated in the destination region. You can also call the [Decrypt](~~28950~~) operation to decrypt the data key.
 *
 * @param request EncryptRequest
 * @return EncryptResponse
 */
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

/**
 * You can call the [GenerateDataKeyWithoutPlaintext](~~134043~~) operation to generate a data key, which is encrypted by a CMK. If you want to distribute the data key to other regions or cryptographic modules, you can call the ExportDataKey operation to use a public key to encrypt the data key.
 * Then, you can import the ciphertext of the data key to the cryptographic module where the private key is stored. This way, the data key is securely distributed from KMS to the cryptographic module. After the data key is imported to the cryptographic module, you can use it to encrypt or decrypt data.
 *
 * @param tmpReq ExportDataKeyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ExportDataKeyResponse
 */
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
	_result = &ExportDataKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the [GenerateDataKeyWithoutPlaintext](~~134043~~) operation to generate a data key, which is encrypted by a CMK. If you want to distribute the data key to other regions or cryptographic modules, you can call the ExportDataKey operation to use a public key to encrypt the data key.
 * Then, you can import the ciphertext of the data key to the cryptographic module where the private key is stored. This way, the data key is securely distributed from KMS to the cryptographic module. After the data key is imported to the cryptographic module, you can use it to encrypt or decrypt data.
 *
 * @param request ExportDataKeyRequest
 * @return ExportDataKeyResponse
 */
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

/**
 * We recommend that you perform the following steps to import your data key to a cryptographic module:
 * *   Call the GenerateAndExportDataKey operation to generate a data key and obtain both the ciphertext of the data key encrypted by using the CMK and that encrypted by using the public key.
 * *   Store the ciphertext of the data key encrypted by using the CMK in KMS Secrets Manager or in a storage service such as ApsaraDB. This ciphertext is used for backup and restoration.
 * *   Import the ciphertext of the data key encrypted by using the public key to the cryptographic module where the private key is stored. Then, you can use the data key to encrypt or decrypt data.
 * >  The CMK that you specify in the request of this operation is only used to encrypt the data key and is not involved in the generation of the data key. KMS does not record or store the data keys randomly generated by calling this operation. You must take note of the data keys and the returned ciphertext.
 *
 * @param tmpReq GenerateAndExportDataKeyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GenerateAndExportDataKeyResponse
 */
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
	_result = &GenerateAndExportDataKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * We recommend that you perform the following steps to import your data key to a cryptographic module:
 * *   Call the GenerateAndExportDataKey operation to generate a data key and obtain both the ciphertext of the data key encrypted by using the CMK and that encrypted by using the public key.
 * *   Store the ciphertext of the data key encrypted by using the CMK in KMS Secrets Manager or in a storage service such as ApsaraDB. This ciphertext is used for backup and restoration.
 * *   Import the ciphertext of the data key encrypted by using the public key to the cryptographic module where the private key is stored. Then, you can use the data key to encrypt or decrypt data.
 * >  The CMK that you specify in the request of this operation is only used to encrypt the data key and is not involved in the generation of the data key. KMS does not record or store the data keys randomly generated by calling this operation. You must take note of the data keys and the returned ciphertext.
 *
 * @param request GenerateAndExportDataKeyRequest
 * @return GenerateAndExportDataKeyResponse
 */
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

/**
 * This operation creates a random data key, encrypts the data key by using the specified customer master key (CMK), and returns the plaintext and ciphertext of the data key. You can use the plaintext of the data key to locally encrypt your data without using KMS and store the encrypted data together with the ciphertext of the data key. You can obtain the plaintext of the data key from the Plaintext parameter in the response and the ciphertext of the data key from the CiphertextBlob parameter in the response.
 * The CMK that you specify in the request of this operation is only used to encrypt the data key and is not involved in the generation of the data key. KMS does not record or store the generated data key. Therefore, you need to store the ciphertext of the data key in persistent storage.
 * We recommend that you locally encrypt data by performing the following steps:
 * 1\\. Call the GenerateDataKey operation.
 * 2\\. Use the plaintext of the data key that you obtain to locally encrypt data without using KMS. Then, delete the plaintext of the data key from the memory.
 * 3\\. Store the encrypted data together with the ciphertext of the data key that you obtain.
 * We recommend that you locally decrypt data by performing the following steps:
 * *   Call the [Decrypt](~~28950~~) operation to decrypt the locally stored ciphertext of the data key. The plaintext of data key is then returned.
 * *   Use the plaintext of the data key to locally decrypt data and then delete the plaintext of the data key from the memory.
 * In this example, a random data key is generated for the CMK whose ID is `7906979c-8e06-46a2-be2d-68e3ccbc****`.
 *
 * @param tmpReq GenerateDataKeyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GenerateDataKeyResponse
 */
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
	_result = &GenerateDataKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation creates a random data key, encrypts the data key by using the specified customer master key (CMK), and returns the plaintext and ciphertext of the data key. You can use the plaintext of the data key to locally encrypt your data without using KMS and store the encrypted data together with the ciphertext of the data key. You can obtain the plaintext of the data key from the Plaintext parameter in the response and the ciphertext of the data key from the CiphertextBlob parameter in the response.
 * The CMK that you specify in the request of this operation is only used to encrypt the data key and is not involved in the generation of the data key. KMS does not record or store the generated data key. Therefore, you need to store the ciphertext of the data key in persistent storage.
 * We recommend that you locally encrypt data by performing the following steps:
 * 1\\. Call the GenerateDataKey operation.
 * 2\\. Use the plaintext of the data key that you obtain to locally encrypt data without using KMS. Then, delete the plaintext of the data key from the memory.
 * 3\\. Store the encrypted data together with the ciphertext of the data key that you obtain.
 * We recommend that you locally decrypt data by performing the following steps:
 * *   Call the [Decrypt](~~28950~~) operation to decrypt the locally stored ciphertext of the data key. The plaintext of data key is then returned.
 * *   Use the plaintext of the data key to locally decrypt data and then delete the plaintext of the data key from the memory.
 * In this example, a random data key is generated for the CMK whose ID is `7906979c-8e06-46a2-be2d-68e3ccbc****`.
 *
 * @param request GenerateDataKeyRequest
 * @return GenerateDataKeyResponse
 */
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

/**
 * This operation creates a random data key, encrypts the data key by using a specific symmetric CMK, and returns the ciphertext of the data key. This operation serves the same purpose as the [GenerateDataKey](~~28948~~) operation. The only difference is that this operation does not return the plaintext of the data key.
 * The CMK that you specify in the request of this operation is only used to encrypt the data key and is not involved in the generation of the data key. KMS does not record or store the generated data key.
 * >
 * *   This operation applies to the scenario when you do not need to use the data key to immediately encrypt data. Before you can use the data key to encrypt data, you must call the [Decrypt](~~28950~~) operation to decrypt the ciphertext of the data key.
 * *   This operation is also suitable for a distributed system with different trust levels. For example, a system stores data in different partitions based on a preset trust policy. A module creates different partitions and generates different data keys for each partition in advance. This module is not involved in data production and consumption after it completes initialization of the control plane. This module is the key provider. When producing and consuming data, modules on the control plane obtain the ciphertext of the data key for a partition first. After decrypting the ciphertext of the data key, modules on the control plane use the plaintext of the data key to encrypt or decrypt data and then clear the plaintext of the data key from the memory. In such a system, the key provider does not need to obtain the plaintext of the data key. It only needs to have the permissions to call the GenerateDataKeyWithoutPlaintext operation. The data producers or consumers do not need to generate new data keys. They only need to have the permissions to call the Decrypt operation.
 *
 * @param tmpReq GenerateDataKeyWithoutPlaintextRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GenerateDataKeyWithoutPlaintextResponse
 */
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
	_result = &GenerateDataKeyWithoutPlaintextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation creates a random data key, encrypts the data key by using a specific symmetric CMK, and returns the ciphertext of the data key. This operation serves the same purpose as the [GenerateDataKey](~~28948~~) operation. The only difference is that this operation does not return the plaintext of the data key.
 * The CMK that you specify in the request of this operation is only used to encrypt the data key and is not involved in the generation of the data key. KMS does not record or store the generated data key.
 * >
 * *   This operation applies to the scenario when you do not need to use the data key to immediately encrypt data. Before you can use the data key to encrypt data, you must call the [Decrypt](~~28950~~) operation to decrypt the ciphertext of the data key.
 * *   This operation is also suitable for a distributed system with different trust levels. For example, a system stores data in different partitions based on a preset trust policy. A module creates different partitions and generates different data keys for each partition in advance. This module is not involved in data production and consumption after it completes initialization of the control plane. This module is the key provider. When producing and consuming data, modules on the control plane obtain the ciphertext of the data key for a partition first. After decrypting the ciphertext of the data key, modules on the control plane use the plaintext of the data key to encrypt or decrypt data and then clear the plaintext of the data key from the memory. In such a system, the key provider does not need to obtain the plaintext of the data key. It only needs to have the permissions to call the GenerateDataKeyWithoutPlaintext operation. The data producers or consumers do not need to generate new data keys. They only need to have the permissions to call the Decrypt operation.
 *
 * @param request GenerateDataKeyWithoutPlaintextRequest
 * @return GenerateDataKeyWithoutPlaintextResponse
 */
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

/**
 * In this example, the certificate whose ID is `9a28de48-8d8b-484d-a766-dec4****` is queried. The certificate, certificate chain, certificate ID, and certificate signing request (CSR) are returned.
 *
 * @param request GetCertificateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetCertificateResponse
 */
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
	_result = &GetCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the certificate whose ID is `9a28de48-8d8b-484d-a766-dec4****` is queried. The certificate, certificate chain, certificate ID, and certificate signing request (CSR) are returned.
 *
 * @param request GetCertificateRequest
 * @return GetCertificateResponse
 */
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

/**
 * The returned parameters can be used to call the [ImportKeyMaterial](https://www.alibabacloud.com/help/en/key-management-service/latest/importkeymaterial) operation.
 * - You can import key material only for CMKs whose Origin parameter is set to EXTERNAL.
 * - The public key and token that are returned by the GetParametersForImport operation must be used together. The public key and token can be used to import key material only for the CMK that is specified when you call the operation.
 * - The public key and token that are returned vary each time you call the GetParametersForImport operation.
 * - You must specify the type of the public key and the encryption algorithm that are used to encrypt key material. The following table lists the types of public keys and the encryption algorithms allowed for each type.
 * | Public key type | Encryption algorithm | Description |
 * | --------------- | -------------------- | ----------- |
 * | RSA_2048 | RSAES_PKCS1_V1_5
 * RSAES_OAEP_SHA_1
 * RSAES_OAEP_SHA_256 | CMKs of all regions and all protection levels are supported.
 * Dedicated Key Management Service (KMS) does not support RSAES_OAEP_SHA_1. |
 * | EC_SM2 | SM2PKE | CMKs whose ProtectionLevel is set to HSM are supported. The SM2 algorithm is developed and approved by the State Cryptography Administration of China. The SM2 algorithm can be used only to import key material for a CMK whose ProtectionLevel is set to HSM. You can use the SM2 algorithm only when you enable the Managed HSM feature for KMS in the Chinese mainland. For more information, see [Overview of Managed HSM](https://www.alibabacloud.com/help/en/key-management-service/latest/managed-hsm-overview). |
 * For more information, see [Import key material](https://www.alibabacloud.com/help/en/key-management-service/latest/import-key-material). This topic provides an example on how to query the parameters that are used to import key material for a CMK. The ID of the CMK is `1234abcd-12ab-34cd-56ef-12345678****`, the encryption algorithm is `RSAES_PKCS1_V1_5`, and the public key is of the `RSA_2048` type. The parameters that are returned include the ID of the CMK, the public key that is used to encrypt the key material, the token that is used to import the key material, and the time when the token expires.
 *
 * @param request GetParametersForImportRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetParametersForImportResponse
 */
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
	_result = &GetParametersForImportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The returned parameters can be used to call the [ImportKeyMaterial](https://www.alibabacloud.com/help/en/key-management-service/latest/importkeymaterial) operation.
 * - You can import key material only for CMKs whose Origin parameter is set to EXTERNAL.
 * - The public key and token that are returned by the GetParametersForImport operation must be used together. The public key and token can be used to import key material only for the CMK that is specified when you call the operation.
 * - The public key and token that are returned vary each time you call the GetParametersForImport operation.
 * - You must specify the type of the public key and the encryption algorithm that are used to encrypt key material. The following table lists the types of public keys and the encryption algorithms allowed for each type.
 * | Public key type | Encryption algorithm | Description |
 * | --------------- | -------------------- | ----------- |
 * | RSA_2048 | RSAES_PKCS1_V1_5
 * RSAES_OAEP_SHA_1
 * RSAES_OAEP_SHA_256 | CMKs of all regions and all protection levels are supported.
 * Dedicated Key Management Service (KMS) does not support RSAES_OAEP_SHA_1. |
 * | EC_SM2 | SM2PKE | CMKs whose ProtectionLevel is set to HSM are supported. The SM2 algorithm is developed and approved by the State Cryptography Administration of China. The SM2 algorithm can be used only to import key material for a CMK whose ProtectionLevel is set to HSM. You can use the SM2 algorithm only when you enable the Managed HSM feature for KMS in the Chinese mainland. For more information, see [Overview of Managed HSM](https://www.alibabacloud.com/help/en/key-management-service/latest/managed-hsm-overview). |
 * For more information, see [Import key material](https://www.alibabacloud.com/help/en/key-management-service/latest/import-key-material). This topic provides an example on how to query the parameters that are used to import key material for a CMK. The ID of the CMK is `1234abcd-12ab-34cd-56ef-12345678****`, the encryption algorithm is `RSAES_PKCS1_V1_5`, and the public key is of the `RSA_2048` type. The parameters that are returned include the ID of the CMK, the public key that is used to encrypt the key material, the token that is used to import the key material, and the time when the token expires.
 *
 * @param request GetParametersForImportRequest
 * @return GetParametersForImportResponse
 */
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
	_result = &GetPublicKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetRandomPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

/**
 * If you do not specify a version number or stage label, Secrets Manager returns the secret value of the version marked with ACSCurrent.
 * If a customer master key (CMK) is specified to encrypt the secret value, you must also have the `kms:Decrypt` permission on the CMK to call the GetSecretValue operation.
 * In this example, the value of the secret named `secret001` is obtained. The secret value is returned in the `SecretData` parameter. The secret value is `testdata1`.
 *
 * @param request GetSecretValueRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetSecretValueResponse
 */
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
	_result = &GetSecretValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you do not specify a version number or stage label, Secrets Manager returns the secret value of the version marked with ACSCurrent.
 * If a customer master key (CMK) is specified to encrypt the secret value, you must also have the `kms:Decrypt` permission on the CMK to call the GetSecretValue operation.
 * In this example, the value of the secret named `secret001` is obtained. The secret value is returned in the `SecretData` parameter. The secret value is `testdata1`.
 *
 * @param request GetSecretValueRequest
 * @return GetSecretValueResponse
 */
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

/**
 * Call [CreateKey](~~28947~~) when creating a CMK, you can select its key material source as external.** Origin** set to** EXTERNAL**. This API is used to import the key material into the CMK.
 * *   To view the CMK **Origin**, see [DescribeKey](~~28952~~).
 * *   Before importing key material, you need to call the [GetParametersForImport](~~68621~~) obtain the parameters required to import the key material, including the public key and import token.
 * >
 * *   The key type of the pair is** Aliyun_AES\\_256** the key material must be 256 bits. The key type must be** Aliyun_SM4** the CMK and key material must be 128 bits.
 * *   You can set the expiration time for the key material, or you can set it to never expire.
 * *   You can reimport the key material and reset the expiration time for the specified CMK at any time, but the same key material must be imported.
 * *   After the imported key material expires or is deleted, the specified CMK is unavailable until the same key material are imported again.
 * *   A Key material can be imported to multiple cmks, but any Data or Data Key encrypted by one CMK cannot be decrypted by another CMK.
 *
 * @param request ImportKeyMaterialRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ImportKeyMaterialResponse
 */
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
	_result = &ImportKeyMaterialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Call [CreateKey](~~28947~~) when creating a CMK, you can select its key material source as external.** Origin** set to** EXTERNAL**. This API is used to import the key material into the CMK.
 * *   To view the CMK **Origin**, see [DescribeKey](~~28952~~).
 * *   Before importing key material, you need to call the [GetParametersForImport](~~68621~~) obtain the parameters required to import the key material, including the public key and import token.
 * >
 * *   The key type of the pair is** Aliyun_AES\\_256** the key material must be 256 bits. The key type must be** Aliyun_SM4** the CMK and key material must be 128 bits.
 * *   You can set the expiration time for the key material, or you can set it to never expire.
 * *   You can reimport the key material and reset the expiration time for the specified CMK at any time, but the same key material must be imported.
 * *   After the imported key material expires or is deleted, the specified CMK is unavailable until the same key material are imported again.
 * *   A Key material can be imported to multiple cmks, but any Data or Data Key encrypted by one CMK cannot be decrypted by another CMK.
 *
 * @param request ImportKeyMaterialRequest
 * @return ImportKeyMaterialResponse
 */
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
	_result = &ListAliasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListAliasesByKeyIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListKeyVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

/**
 * Request format: KeyId="string"
 *
 * @param request ListResourceTagsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListResourceTagsResponse
 */
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
	_result = &ListResourceTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Request format: KeyId="string"
 *
 * @param request ListResourceTagsRequest
 * @return ListResourceTagsResponse
 */
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

/**
 * The secret value is not included in the returned version information. By default, deprecated secret versions are not returned.
 *
 * @param request ListSecretVersionIdsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListSecretVersionIdsResponse
 */
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
	_result = &ListSecretVersionIdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The secret value is not included in the returned version information. By default, deprecated secret versions are not returned.
 *
 * @param request ListSecretVersionIdsRequest
 * @return ListSecretVersionIdsResponse
 */
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

/**
 * This operation returns the metadata information about the secrets and does not return encrypted secret values.
 * In this example, the secrets created by the current account in the current region are returned. The `PageNumber` parameter is set to `1`, and the `PageSize` parameter is set to `2`, which indicates that two secrets are to be returned on Page 1.
 *
 * @param request ListSecretsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListSecretsResponse
 */
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
	_result = &ListSecretsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation returns the metadata information about the secrets and does not return encrypted secret values.
 * In this example, the secrets created by the current account in the current region are returned. The `PageNumber` parameter is set to `1`, and the `PageSize` parameter is set to `2`, which indicates that two secrets are to be returned on Page 1.
 *
 * @param request ListSecretsRequest
 * @return ListSecretsResponse
 */
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

/**
 * When you call this operation, note that:
 * - KMS is a paid service. For more information about the billing method, see [Billing description](https://www.alibabacloud.com/help/en/key-management-service/latest/billing-billing).
 * - An Alibaba Cloud account can activate KMS only once.
 * - Make sure that your Alibaba Cloud account has passed real-name authentication.
 *
 * @param request OpenKmsServiceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return OpenKmsServiceResponse
 */
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
	_result = &OpenKmsServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * When you call this operation, note that:
 * - KMS is a paid service. For more information about the billing method, see [Billing description](https://www.alibabacloud.com/help/en/key-management-service/latest/billing-billing).
 * - An Alibaba Cloud account can activate KMS only once.
 * - Make sure that your Alibaba Cloud account has passed real-name authentication.
 *
 * @return OpenKmsServiceResponse
 */
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

/**
 * This operation is used to store the secret values of new versions. It cannot be used to modify the secret value of an existing version.
 * By default, the newly stored secret value is marked with ACSCurrent, and the mark for the previous version of the secret value is changed from ACSCurrent to ACSPrevious. If you specify the VersionStage parameter, the newly stored secret value is marked with the stage label that you specify.
 * You must specify a version number when you call the operation. Secrets Manager performs operations based on the following rules:
 * *   If the specified version number does not exist in the secret, Secrets Manager creates the version and stores the secret value.
 * *   If the specified version number already exists in the secret and the secret value of the existing version is the same as the secret value that you specify, Secrets Manager ignores the request and returns a success message. The request is idempotent.
 * *   If the specified version number already exists in the secret but the secret value of the existing version is different from the secret value that you specify, Secrets Manager rejects the request and returns a failure message.
 * Limits: This operation is available only for standard secrets.
 * In this example, the secret value of a new version is stored into the `secret001` secret. The `VersionId` parameter is set to `00000000000000000000000000000000203` as the new version, and the `SecretData` parameter is set to `importantdata`.
 *
 * @param request PutSecretValueRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return PutSecretValueResponse
 */
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
	_result = &PutSecretValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation is used to store the secret values of new versions. It cannot be used to modify the secret value of an existing version.
 * By default, the newly stored secret value is marked with ACSCurrent, and the mark for the previous version of the secret value is changed from ACSCurrent to ACSPrevious. If you specify the VersionStage parameter, the newly stored secret value is marked with the stage label that you specify.
 * You must specify a version number when you call the operation. Secrets Manager performs operations based on the following rules:
 * *   If the specified version number does not exist in the secret, Secrets Manager creates the version and stores the secret value.
 * *   If the specified version number already exists in the secret and the secret value of the existing version is the same as the secret value that you specify, Secrets Manager ignores the request and returns a success message. The request is idempotent.
 * *   If the specified version number already exists in the secret but the secret value of the existing version is different from the secret value that you specify, Secrets Manager rejects the request and returns a failure message.
 * Limits: This operation is available only for standard secrets.
 * In this example, the secret value of a new version is stored into the `secret001` secret. The `VersionId` parameter is set to `00000000000000000000000000000000203` as the new version, and the `SecretData` parameter is set to `importantdata`.
 *
 * @param request PutSecretValueRequest
 * @return PutSecretValueResponse
 */
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

/**
 * You can call this operation in the following scenarios:
 * *   After the CMK that was used to encrypt your data is rotated, you can call this operation to use the latest CMK version to re-encrypt the data. For more information about automatic key rotation, see [Configure automatic key rotation](~~134270~~).
 * *   The CMK that was used to encrypt your data remains unchanged, but EncryptionContext is changed. In this scenario, you can call this operation to re-encrypt the data.
 * *   You can call this operation to use a CMK in KMS to re-encrypt data or a data key that was previously encrypted by a different CMK.
 * To use the ReEncrypt operation, you must have two permissions:
 * *   kms:ReEncryptFrom on the source CMK
 * *   kms:ReEncryptTo on the destination CMK
 * *   For simplicity, you can specify kms:ReEncrypt\\* to allow both of the preceding permissions.
 *
 * @param tmpReq ReEncryptRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ReEncryptResponse
 */
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
	_result = &ReEncryptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation in the following scenarios:
 * *   After the CMK that was used to encrypt your data is rotated, you can call this operation to use the latest CMK version to re-encrypt the data. For more information about automatic key rotation, see [Configure automatic key rotation](~~134270~~).
 * *   The CMK that was used to encrypt your data remains unchanged, but EncryptionContext is changed. In this scenario, you can call this operation to re-encrypt the data.
 * *   You can call this operation to use a CMK in KMS to re-encrypt data or a data key that was previously encrypted by a different CMK.
 * To use the ReEncrypt operation, you must have two permissions:
 * *   kms:ReEncryptFrom on the source CMK
 * *   kms:ReEncryptTo on the destination CMK
 * *   For simplicity, you can specify kms:ReEncrypt\\* to allow both of the preceding permissions.
 *
 * @param request ReEncryptRequest
 * @return ReEncryptResponse
 */
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

/**
 * You can only use this operation to restore a deleted secret that is within its recovery period. If you set **ForceDeleteWithoutRecovery** to **true** when you delete the secret, you cannot restore it.
 *
 * @param request RestoreSecretRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RestoreSecretResponse
 */
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
	_result = &RestoreSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can only use this operation to restore a deleted secret that is within its recovery period. If you set **ForceDeleteWithoutRecovery** to **true** when you delete the secret, you cannot restore it.
 *
 * @param request RestoreSecretRequest
 * @return RestoreSecretResponse
 */
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

/**
 * Limits:
 * • A secret of each Alibaba Cloud account can be rotated for a maximum of 50 times per hour.
 * • The RotateSecret operation is unavailable for standard secrets.
 * In this example, the `RdsSecret/Mysql5.4/MyCred` secret is manually rotated, and the version number of the secret is set to `000000123` after the secret is rotated.
 *
 * @param request RotateSecretRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RotateSecretResponse
 */
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
	_result = &RotateSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Limits:
 * • A secret of each Alibaba Cloud account can be rotated for a maximum of 50 times per hour.
 * • The RotateSecret operation is unavailable for standard secrets.
 * In this example, the `RdsSecret/Mysql5.4/MyCred` secret is manually rotated, and the version number of the secret is set to `000000123` after the secret is rotated.
 *
 * @param request RotateSecretRequest
 * @return RotateSecretResponse
 */
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

/**
 * During the scheduled period, the CMK is in the PendingDeletion state and cannot be used to encrypt data, decrypt data, or generate data keys.
 * After a CMK is deleted, it cannot be recovered. Data that is encrypted and data keys that are generated by using the CMK cannot be decrypted. To prevent accidental deletion of CMKs, Key Management Service (KMS) allows you to only schedule key deletion tasks. You cannot directly delete CMKs. If you want to delete a CMK, call the [DisableKey](~~35151~~) operation to disable the CMK.
 * When you call this operation, you must specify a scheduled period between 7 days to 366 days. The scheduled period starts from the time when you submit the request. You can call the [CancelKeyDeletion](~~44197~~) operation to cancel the key deletion task before the scheduled period ends.
 *
 * @param request ScheduleKeyDeletionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ScheduleKeyDeletionResponse
 */
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
	_result = &ScheduleKeyDeletionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * During the scheduled period, the CMK is in the PendingDeletion state and cannot be used to encrypt data, decrypt data, or generate data keys.
 * After a CMK is deleted, it cannot be recovered. Data that is encrypted and data keys that are generated by using the CMK cannot be decrypted. To prevent accidental deletion of CMKs, Key Management Service (KMS) allows you to only schedule key deletion tasks. You cannot directly delete CMKs. If you want to delete a CMK, call the [DisableKey](~~35151~~) operation to disable the CMK.
 * When you call this operation, you must specify a scheduled period between 7 days to 366 days. The scheduled period starts from the time when you submit the request. You can call the [CancelKeyDeletion](~~44197~~) operation to cancel the key deletion task before the scheduled period ends.
 *
 * @param request ScheduleKeyDeletionRequest
 * @return ScheduleKeyDeletionResponse
 */
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

/**
 * *   After you enable deletion protection for a CMK, you cannot delete the CMK. If you want to delete the CMK, you must first disable deletion protection for the CMK.
 * *   Before you can call the SetDeletionProtection operation, make sure that the required CMK is not in the Pending Deletion state. You can call the [DescribeKey](~~28952~~) operation to query the CMK status, which is specified by the KeyState parameter.
 * You can enable deletion protection for the CMK whose Alibaba Cloud Resource Name (ARN) is `acs:kms:cn-hangzhou:123213123****:key/0225f411-b21d-46d1-be5b-93931c82****` by using parameter settings provided in this topic. The CMK ARN is specified by the ProtectedResourceArn parameter.
 *
 * @param request SetDeletionProtectionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SetDeletionProtectionResponse
 */
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
	_result = &SetDeletionProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   After you enable deletion protection for a CMK, you cannot delete the CMK. If you want to delete the CMK, you must first disable deletion protection for the CMK.
 * *   Before you can call the SetDeletionProtection operation, make sure that the required CMK is not in the Pending Deletion state. You can call the [DescribeKey](~~28952~~) operation to query the CMK status, which is specified by the KeyState parameter.
 * You can enable deletion protection for the CMK whose Alibaba Cloud Resource Name (ARN) is `acs:kms:cn-hangzhou:123213123****:key/0225f411-b21d-46d1-be5b-93931c82****` by using parameter settings provided in this topic. The CMK ARN is specified by the ProtectedResourceArn parameter.
 *
 * @param request SetDeletionProtectionRequest
 * @return SetDeletionProtectionResponse
 */
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

/**
 * You can add up to 10 tags to a CMK, secret, or certificate.
 * In this example, the tags `[{"TagKey":"S1key1","TagValue":"S1val1"},{"TagKey":"S1key2","TagValue":"S2val2"}]` are added to the CMK whose ID is `08c33a6f-4e0a-4a1b-a3fa-7ddf****`.
 *
 * @param request TagResourceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return TagResourceResponse
 */
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
	_result = &TagResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can add up to 10 tags to a CMK, secret, or certificate.
 * In this example, the tags `[{"TagKey":"S1key1","TagValue":"S1val1"},{"TagKey":"S1key2","TagValue":"S2val2"}]` are added to the CMK whose ID is `08c33a6f-4e0a-4a1b-a3fa-7ddf****`.
 *
 * @param request TagResourceRequest
 * @return TagResourceResponse
 */
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

/**
 * In this example, the tags whose tag keys are tagkey1 and tagkey2 are removed from the CMK whose ID is `08c33a6f-4e0a-4a1b-a3fa-7ddf****`.
 *
 * @param request UntagResourceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UntagResourceResponse
 */
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
	_result = &UntagResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the tags whose tag keys are tagkey1 and tagkey2 are removed from the CMK whose ID is `08c33a6f-4e0a-4a1b-a3fa-7ddf****`.
 *
 * @param request UntagResourceRequest
 * @return UntagResourceResponse
 */
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
	_result = &UpdateAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

/**
 * In this example, the status of the certificate whose ID is `9a28de48-8d8b-484d-a766-dec4****` is updated to INACTIVE.
 *
 * @param request UpdateCertificateStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateCertificateStatusResponse
 */
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
	_result = &UpdateCertificateStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the status of the certificate whose ID is `9a28de48-8d8b-484d-a766-dec4****` is updated to INACTIVE.
 *
 * @param request UpdateCertificateStatusRequest
 * @return UpdateCertificateStatusResponse
 */
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

/**
 * This operation replaces the description of a customer master key (CMK) with the description that you specify. The original description of the CMK is specified by the Description parameter when you call the [DescribeKey](~~28952~~) operation. You can call this operation to add, modify, or delete the description of a CMK.
 *
 * @param request UpdateKeyDescriptionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateKeyDescriptionResponse
 */
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
	_result = &UpdateKeyDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation replaces the description of a customer master key (CMK) with the description that you specify. The original description of the CMK is specified by the Description parameter when you call the [DescribeKey](~~28952~~) operation. You can call this operation to add, modify, or delete the description of a CMK.
 *
 * @param request UpdateKeyDescriptionRequest
 * @return UpdateKeyDescriptionResponse
 */
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

/**
 * When automatic key rotation is enabled, KMS automatically creates a key version after the preset rotation period arrives. In addition, KMS sets the new key version as the primary key version.
 * An automatic key rotation policy cannot be configured for the following keys:
 * *   Asymmetric key
 * *   Service-managed key
 * *   Bring your own key (BYOK) that is imported into KMS
 * *   Key that is not in the **Enabled** state
 * In this example, automatic key rotation is enabled for a CMK whose ID is `1234abcd-12ab-34cd-56ef-12345678****`. The automatic rotation period is 30 days.
 *
 * @param request UpdateRotationPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateRotationPolicyResponse
 */
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
	_result = &UpdateRotationPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * When automatic key rotation is enabled, KMS automatically creates a key version after the preset rotation period arrives. In addition, KMS sets the new key version as the primary key version.
 * An automatic key rotation policy cannot be configured for the following keys:
 * *   Asymmetric key
 * *   Service-managed key
 * *   Bring your own key (BYOK) that is imported into KMS
 * *   Key that is not in the **Enabled** state
 * In this example, automatic key rotation is enabled for a CMK whose ID is `1234abcd-12ab-34cd-56ef-12345678****`. The automatic rotation period is 30 days.
 *
 * @param request UpdateRotationPolicyRequest
 * @return UpdateRotationPolicyResponse
 */
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

/**
 * In this example, the metadata of the `secret001` secret is updated. The `Description` parameter is set to `datainfo`.
 *
 * @param request UpdateSecretRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateSecretResponse
 */
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
	_result = &UpdateSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the metadata of the `secret001` secret is updated. The `Description` parameter is set to `datainfo`.
 *
 * @param request UpdateSecretRequest
 * @return UpdateSecretResponse
 */
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

/**
 * After automatic rotation is enabled, Secrets Manager schedules the first automatic rotation by adding the preset rotation interval to the timestamp of the last rotation.
 * Limits: The UpdateSecretRotationPolicy operation cannot be used to update the rotation policy of generic secrets.
 * In this example, the rotation policy of the `RdsSecret/Mysql5.4/MyCred` secret is updated. The following settings are modified:
 * *   The `EnableAutomaticRotation` parameter is set to `true`, which indicates that automatic rotation is enabled.
 * *   The `RotationInterval` parameter is set to `30d`, which indicates that the interval for automatic rotation is 30 days.
 *
 * @param request UpdateSecretRotationPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateSecretRotationPolicyResponse
 */
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
	_result = &UpdateSecretRotationPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * After automatic rotation is enabled, Secrets Manager schedules the first automatic rotation by adding the preset rotation interval to the timestamp of the last rotation.
 * Limits: The UpdateSecretRotationPolicy operation cannot be used to update the rotation policy of generic secrets.
 * In this example, the rotation policy of the `RdsSecret/Mysql5.4/MyCred` secret is updated. The following settings are modified:
 * *   The `EnableAutomaticRotation` parameter is set to `true`, which indicates that automatic rotation is enabled.
 * *   The `RotationInterval` parameter is set to `30d`, which indicates that the interval for automatic rotation is 30 days.
 *
 * @param request UpdateSecretRotationPolicyRequest
 * @return UpdateSecretRotationPolicyResponse
 */
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

/**
 * You can use this operation to achieve the following purposes:
 * *   Use a specified stage label to mark a new secret version.
 * *   Remove a specific stage label from an existing secret version.
 * Limits: This operation is available only for generic secrets.
 * In this example, the stage label that marks the version of the `secret001` secret is updated. The stage label `ACSCurrent` is used to mark the `002` version.
 *
 * @param request UpdateSecretVersionStageRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateSecretVersionStageResponse
 */
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
	_result = &UpdateSecretVersionStageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can use this operation to achieve the following purposes:
 * *   Use a specified stage label to mark a new secret version.
 * *   Remove a specific stage label from an existing secret version.
 * Limits: This operation is available only for generic secrets.
 * In this example, the stage label that marks the version of the `secret001` secret is updated. The stage label `ACSCurrent` is used to mark the `002` version.
 *
 * @param request UpdateSecretVersionStageRequest
 * @return UpdateSecretVersionStageResponse
 */
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

/**
 * In this example, a certificate issued by a CA is imported into Certificates Manager. The ID of the certificate in Certificates Manager is `12345678-1234-1234-1234-12345678****`.
 *
 * @param request UploadCertificateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UploadCertificateResponse
 */
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
	_result = &UploadCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, a certificate issued by a CA is imported into Certificates Manager. The ID of the certificate in Certificates Manager is `12345678-1234-1234-1234-12345678****`.
 *
 * @param request UploadCertificateRequest
 * @return UploadCertificateResponse
 */
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
