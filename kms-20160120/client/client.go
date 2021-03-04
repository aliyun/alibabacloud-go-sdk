// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AsymmetricDecryptRequest struct {
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	KeyId          *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeyVersionId   *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	Algorithm      *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
}

func (s AsymmetricDecryptRequest) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricDecryptRequest) GoString() string {
	return s.String()
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

func (s *AsymmetricDecryptRequest) SetAlgorithm(v string) *AsymmetricDecryptRequest {
	s.Algorithm = &v
	return s
}

type AsymmetricDecryptResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Plaintext    *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s AsymmetricDecryptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricDecryptResponseBody) GoString() string {
	return s.String()
}

func (s *AsymmetricDecryptResponseBody) SetRequestId(v string) *AsymmetricDecryptResponseBody {
	s.RequestId = &v
	return s
}

func (s *AsymmetricDecryptResponseBody) SetPlaintext(v string) *AsymmetricDecryptResponseBody {
	s.Plaintext = &v
	return s
}

func (s *AsymmetricDecryptResponseBody) SetKeyId(v string) *AsymmetricDecryptResponseBody {
	s.KeyId = &v
	return s
}

func (s *AsymmetricDecryptResponseBody) SetKeyVersionId(v string) *AsymmetricDecryptResponseBody {
	s.KeyVersionId = &v
	return s
}

type AsymmetricDecryptResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AsymmetricDecryptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AsymmetricDecryptResponse) SetBody(v *AsymmetricDecryptResponseBody) *AsymmetricDecryptResponse {
	s.Body = v
	return s
}

type AsymmetricEncryptRequest struct {
	Plaintext    *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	Algorithm    *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
}

func (s AsymmetricEncryptRequest) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricEncryptRequest) GoString() string {
	return s.String()
}

func (s *AsymmetricEncryptRequest) SetPlaintext(v string) *AsymmetricEncryptRequest {
	s.Plaintext = &v
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

func (s *AsymmetricEncryptRequest) SetAlgorithm(v string) *AsymmetricEncryptRequest {
	s.Algorithm = &v
	return s
}

type AsymmetricEncryptResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	KeyId          *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeyVersionId   *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s AsymmetricEncryptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricEncryptResponseBody) GoString() string {
	return s.String()
}

func (s *AsymmetricEncryptResponseBody) SetRequestId(v string) *AsymmetricEncryptResponseBody {
	s.RequestId = &v
	return s
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

type AsymmetricEncryptResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AsymmetricEncryptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AsymmetricEncryptResponse) SetBody(v *AsymmetricEncryptResponseBody) *AsymmetricEncryptResponse {
	s.Body = v
	return s
}

type AsymmetricSignRequest struct {
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	Algorithm    *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Digest       *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
}

func (s AsymmetricSignRequest) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricSignRequest) GoString() string {
	return s.String()
}

func (s *AsymmetricSignRequest) SetKeyId(v string) *AsymmetricSignRequest {
	s.KeyId = &v
	return s
}

func (s *AsymmetricSignRequest) SetKeyVersionId(v string) *AsymmetricSignRequest {
	s.KeyVersionId = &v
	return s
}

func (s *AsymmetricSignRequest) SetAlgorithm(v string) *AsymmetricSignRequest {
	s.Algorithm = &v
	return s
}

func (s *AsymmetricSignRequest) SetDigest(v string) *AsymmetricSignRequest {
	s.Digest = &v
	return s
}

type AsymmetricSignResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s AsymmetricSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricSignResponseBody) GoString() string {
	return s.String()
}

func (s *AsymmetricSignResponseBody) SetRequestId(v string) *AsymmetricSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *AsymmetricSignResponseBody) SetValue(v string) *AsymmetricSignResponseBody {
	s.Value = &v
	return s
}

func (s *AsymmetricSignResponseBody) SetKeyId(v string) *AsymmetricSignResponseBody {
	s.KeyId = &v
	return s
}

func (s *AsymmetricSignResponseBody) SetKeyVersionId(v string) *AsymmetricSignResponseBody {
	s.KeyVersionId = &v
	return s
}

type AsymmetricSignResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AsymmetricSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AsymmetricSignResponse) SetBody(v *AsymmetricSignResponseBody) *AsymmetricSignResponse {
	s.Body = v
	return s
}

type AsymmetricVerifyRequest struct {
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	Algorithm    *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Digest       *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AsymmetricVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricVerifyRequest) GoString() string {
	return s.String()
}

func (s *AsymmetricVerifyRequest) SetKeyId(v string) *AsymmetricVerifyRequest {
	s.KeyId = &v
	return s
}

func (s *AsymmetricVerifyRequest) SetKeyVersionId(v string) *AsymmetricVerifyRequest {
	s.KeyVersionId = &v
	return s
}

func (s *AsymmetricVerifyRequest) SetAlgorithm(v string) *AsymmetricVerifyRequest {
	s.Algorithm = &v
	return s
}

func (s *AsymmetricVerifyRequest) SetDigest(v string) *AsymmetricVerifyRequest {
	s.Digest = &v
	return s
}

func (s *AsymmetricVerifyRequest) SetValue(v string) *AsymmetricVerifyRequest {
	s.Value = &v
	return s
}

type AsymmetricVerifyResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Value        *bool   `json:"Value,omitempty" xml:"Value,omitempty"`
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s AsymmetricVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *AsymmetricVerifyResponseBody) SetRequestId(v string) *AsymmetricVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AsymmetricVerifyResponseBody) SetValue(v bool) *AsymmetricVerifyResponseBody {
	s.Value = &v
	return s
}

func (s *AsymmetricVerifyResponseBody) SetKeyId(v string) *AsymmetricVerifyResponseBody {
	s.KeyId = &v
	return s
}

func (s *AsymmetricVerifyResponseBody) SetKeyVersionId(v string) *AsymmetricVerifyResponseBody {
	s.KeyVersionId = &v
	return s
}

type AsymmetricVerifyResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AsymmetricVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AsymmetricVerifyResponse) SetBody(v *AsymmetricVerifyResponseBody) *AsymmetricVerifyResponse {
	s.Body = v
	return s
}

type CancelKeyDeletionRequest struct {
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelKeyDeletionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CancelKeyDeletionResponse) SetBody(v *CancelKeyDeletionResponseBody) *CancelKeyDeletionResponse {
	s.Body = v
	return s
}

type CertificatePrivateKeyDecryptRequest struct {
	CertificateId  *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	Algorithm      *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
}

func (s CertificatePrivateKeyDecryptRequest) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrivateKeyDecryptRequest) GoString() string {
	return s.String()
}

func (s *CertificatePrivateKeyDecryptRequest) SetCertificateId(v string) *CertificatePrivateKeyDecryptRequest {
	s.CertificateId = &v
	return s
}

func (s *CertificatePrivateKeyDecryptRequest) SetAlgorithm(v string) *CertificatePrivateKeyDecryptRequest {
	s.Algorithm = &v
	return s
}

func (s *CertificatePrivateKeyDecryptRequest) SetCiphertextBlob(v string) *CertificatePrivateKeyDecryptRequest {
	s.CiphertextBlob = &v
	return s
}

type CertificatePrivateKeyDecryptResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	Plaintext     *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
}

func (s CertificatePrivateKeyDecryptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrivateKeyDecryptResponseBody) GoString() string {
	return s.String()
}

func (s *CertificatePrivateKeyDecryptResponseBody) SetRequestId(v string) *CertificatePrivateKeyDecryptResponseBody {
	s.RequestId = &v
	return s
}

func (s *CertificatePrivateKeyDecryptResponseBody) SetCertificateId(v string) *CertificatePrivateKeyDecryptResponseBody {
	s.CertificateId = &v
	return s
}

func (s *CertificatePrivateKeyDecryptResponseBody) SetPlaintext(v string) *CertificatePrivateKeyDecryptResponseBody {
	s.Plaintext = &v
	return s
}

type CertificatePrivateKeyDecryptResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CertificatePrivateKeyDecryptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CertificatePrivateKeyDecryptResponse) SetBody(v *CertificatePrivateKeyDecryptResponseBody) *CertificatePrivateKeyDecryptResponse {
	s.Body = v
	return s
}

type CertificatePrivateKeySignRequest struct {
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	Algorithm     *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	MessageType   *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s CertificatePrivateKeySignRequest) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrivateKeySignRequest) GoString() string {
	return s.String()
}

func (s *CertificatePrivateKeySignRequest) SetCertificateId(v string) *CertificatePrivateKeySignRequest {
	s.CertificateId = &v
	return s
}

func (s *CertificatePrivateKeySignRequest) SetAlgorithm(v string) *CertificatePrivateKeySignRequest {
	s.Algorithm = &v
	return s
}

func (s *CertificatePrivateKeySignRequest) SetMessageType(v string) *CertificatePrivateKeySignRequest {
	s.MessageType = &v
	return s
}

func (s *CertificatePrivateKeySignRequest) SetMessage(v string) *CertificatePrivateKeySignRequest {
	s.Message = &v
	return s
}

type CertificatePrivateKeySignResponseBody struct {
	SignatureValue *string `json:"SignatureValue,omitempty" xml:"SignatureValue,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CertificateId  *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s CertificatePrivateKeySignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrivateKeySignResponseBody) GoString() string {
	return s.String()
}

func (s *CertificatePrivateKeySignResponseBody) SetSignatureValue(v string) *CertificatePrivateKeySignResponseBody {
	s.SignatureValue = &v
	return s
}

func (s *CertificatePrivateKeySignResponseBody) SetRequestId(v string) *CertificatePrivateKeySignResponseBody {
	s.RequestId = &v
	return s
}

func (s *CertificatePrivateKeySignResponseBody) SetCertificateId(v string) *CertificatePrivateKeySignResponseBody {
	s.CertificateId = &v
	return s
}

type CertificatePrivateKeySignResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CertificatePrivateKeySignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CertificatePrivateKeySignResponse) SetBody(v *CertificatePrivateKeySignResponseBody) *CertificatePrivateKeySignResponse {
	s.Body = v
	return s
}

type CertificatePublicKeyEncryptRequest struct {
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	Algorithm     *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Plaintext     *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
}

func (s CertificatePublicKeyEncryptRequest) String() string {
	return tea.Prettify(s)
}

func (s CertificatePublicKeyEncryptRequest) GoString() string {
	return s.String()
}

func (s *CertificatePublicKeyEncryptRequest) SetCertificateId(v string) *CertificatePublicKeyEncryptRequest {
	s.CertificateId = &v
	return s
}

func (s *CertificatePublicKeyEncryptRequest) SetAlgorithm(v string) *CertificatePublicKeyEncryptRequest {
	s.Algorithm = &v
	return s
}

func (s *CertificatePublicKeyEncryptRequest) SetPlaintext(v string) *CertificatePublicKeyEncryptRequest {
	s.Plaintext = &v
	return s
}

type CertificatePublicKeyEncryptResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	CertificateId  *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s CertificatePublicKeyEncryptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CertificatePublicKeyEncryptResponseBody) GoString() string {
	return s.String()
}

func (s *CertificatePublicKeyEncryptResponseBody) SetRequestId(v string) *CertificatePublicKeyEncryptResponseBody {
	s.RequestId = &v
	return s
}

func (s *CertificatePublicKeyEncryptResponseBody) SetCiphertextBlob(v string) *CertificatePublicKeyEncryptResponseBody {
	s.CiphertextBlob = &v
	return s
}

func (s *CertificatePublicKeyEncryptResponseBody) SetCertificateId(v string) *CertificatePublicKeyEncryptResponseBody {
	s.CertificateId = &v
	return s
}

type CertificatePublicKeyEncryptResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CertificatePublicKeyEncryptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CertificatePublicKeyEncryptResponse) SetBody(v *CertificatePublicKeyEncryptResponseBody) *CertificatePublicKeyEncryptResponse {
	s.Body = v
	return s
}

type CertificatePublicKeyVerifyRequest struct {
	CertificateId  *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	Algorithm      *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	MessageType    *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	SignatureValue *string `json:"SignatureValue,omitempty" xml:"SignatureValue,omitempty"`
}

func (s CertificatePublicKeyVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s CertificatePublicKeyVerifyRequest) GoString() string {
	return s.String()
}

func (s *CertificatePublicKeyVerifyRequest) SetCertificateId(v string) *CertificatePublicKeyVerifyRequest {
	s.CertificateId = &v
	return s
}

func (s *CertificatePublicKeyVerifyRequest) SetAlgorithm(v string) *CertificatePublicKeyVerifyRequest {
	s.Algorithm = &v
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
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CertificateId  *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	SignatureValid *bool   `json:"SignatureValid,omitempty" xml:"SignatureValid,omitempty"`
}

func (s CertificatePublicKeyVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CertificatePublicKeyVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *CertificatePublicKeyVerifyResponseBody) SetRequestId(v string) *CertificatePublicKeyVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CertificatePublicKeyVerifyResponseBody) SetCertificateId(v string) *CertificatePublicKeyVerifyResponseBody {
	s.CertificateId = &v
	return s
}

func (s *CertificatePublicKeyVerifyResponseBody) SetSignatureValid(v bool) *CertificatePublicKeyVerifyResponseBody {
	s.SignatureValid = &v
	return s
}

type CertificatePublicKeyVerifyResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CertificatePublicKeyVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CertificatePublicKeyVerifyResponse) SetBody(v *CertificatePublicKeyVerifyResponseBody) *CertificatePublicKeyVerifyResponse {
	s.Body = v
	return s
}

type CreateAliasRequest struct {
	KeyId     *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
}

func (s CreateAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAliasRequest) GoString() string {
	return s.String()
}

func (s *CreateAliasRequest) SetKeyId(v string) *CreateAliasRequest {
	s.KeyId = &v
	return s
}

func (s *CreateAliasRequest) SetAliasName(v string) *CreateAliasRequest {
	s.AliasName = &v
	return s
}

type CreateAliasResponseBody struct {
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAliasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateAliasResponse) SetBody(v *CreateAliasResponseBody) *CreateAliasResponse {
	s.Body = v
	return s
}

type CreateCertificateRequest struct {
	Subject                 *string                `json:"Subject,omitempty" xml:"Subject,omitempty"`
	SubjectAlternativeNames map[string]interface{} `json:"SubjectAlternativeNames,omitempty" xml:"SubjectAlternativeNames,omitempty"`
	KeySpec                 *string                `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	ProtectionLevel         *string                `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
}

func (s CreateCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateCertificateRequest) SetSubject(v string) *CreateCertificateRequest {
	s.Subject = &v
	return s
}

func (s *CreateCertificateRequest) SetSubjectAlternativeNames(v map[string]interface{}) *CreateCertificateRequest {
	s.SubjectAlternativeNames = v
	return s
}

func (s *CreateCertificateRequest) SetKeySpec(v string) *CreateCertificateRequest {
	s.KeySpec = &v
	return s
}

func (s *CreateCertificateRequest) SetProtectionLevel(v string) *CreateCertificateRequest {
	s.ProtectionLevel = &v
	return s
}

type CreateCertificateShrinkRequest struct {
	Subject                       *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	SubjectAlternativeNamesShrink *string `json:"SubjectAlternativeNames,omitempty" xml:"SubjectAlternativeNames,omitempty"`
	KeySpec                       *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	ProtectionLevel               *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
}

func (s CreateCertificateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCertificateShrinkRequest) SetSubject(v string) *CreateCertificateShrinkRequest {
	s.Subject = &v
	return s
}

func (s *CreateCertificateShrinkRequest) SetSubjectAlternativeNamesShrink(v string) *CreateCertificateShrinkRequest {
	s.SubjectAlternativeNamesShrink = &v
	return s
}

func (s *CreateCertificateShrinkRequest) SetKeySpec(v string) *CreateCertificateShrinkRequest {
	s.KeySpec = &v
	return s
}

func (s *CreateCertificateShrinkRequest) SetProtectionLevel(v string) *CreateCertificateShrinkRequest {
	s.ProtectionLevel = &v
	return s
}

type CreateCertificateResponseBody struct {
	Csr           *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	Arn           *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
}

func (s CreateCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCertificateResponseBody) SetCsr(v string) *CreateCertificateResponseBody {
	s.Csr = &v
	return s
}

func (s *CreateCertificateResponseBody) SetRequestId(v string) *CreateCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCertificateResponseBody) SetCertificateId(v string) *CreateCertificateResponseBody {
	s.CertificateId = &v
	return s
}

func (s *CreateCertificateResponseBody) SetArn(v string) *CreateCertificateResponseBody {
	s.Arn = &v
	return s
}

type CreateCertificateResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateCertificateResponse) SetBody(v *CreateCertificateResponseBody) *CreateCertificateResponse {
	s.Body = v
	return s
}

type CreateKeyRequest struct {
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	KeyUsage                *string `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty"`
	Origin                  *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	ProtectionLevel         *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	EnableAutomaticRotation *bool   `json:"EnableAutomaticRotation,omitempty" xml:"EnableAutomaticRotation,omitempty"`
	RotationInterval        *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	KeySpec                 *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
}

func (s CreateKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateKeyRequest) SetDescription(v string) *CreateKeyRequest {
	s.Description = &v
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

func (s *CreateKeyRequest) SetEnableAutomaticRotation(v bool) *CreateKeyRequest {
	s.EnableAutomaticRotation = &v
	return s
}

func (s *CreateKeyRequest) SetRotationInterval(v string) *CreateKeyRequest {
	s.RotationInterval = &v
	return s
}

func (s *CreateKeyRequest) SetKeySpec(v string) *CreateKeyRequest {
	s.KeySpec = &v
	return s
}

type CreateKeyResponseBody struct {
	RequestId   *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	KeyMetadata *CreateKeyResponseBodyKeyMetadata `json:"KeyMetadata,omitempty" xml:"KeyMetadata,omitempty" type:"Struct"`
}

func (s CreateKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKeyResponseBody) SetRequestId(v string) *CreateKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKeyResponseBody) SetKeyMetadata(v *CreateKeyResponseBodyKeyMetadata) *CreateKeyResponseBody {
	s.KeyMetadata = v
	return s
}

type CreateKeyResponseBodyKeyMetadata struct {
	KeyId              *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	NextRotationDate   *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	KeyState           *string `json:"KeyState,omitempty" xml:"KeyState,omitempty"`
	RotationInterval   *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	Arn                *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	Creator            *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DeleteDate         *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
	LastRotationDate   *string `json:"LastRotationDate,omitempty" xml:"LastRotationDate,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PrimaryKeyVersion  *string `json:"PrimaryKeyVersion,omitempty" xml:"PrimaryKeyVersion,omitempty"`
	Origin             *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	KeySpec            *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	MaterialExpireTime *string `json:"MaterialExpireTime,omitempty" xml:"MaterialExpireTime,omitempty"`
	AutomaticRotation  *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	ProtectionLevel    *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	KeyUsage           *string `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty"`
	CreationDate       *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
}

func (s CreateKeyResponseBodyKeyMetadata) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyResponseBodyKeyMetadata) GoString() string {
	return s.String()
}

func (s *CreateKeyResponseBodyKeyMetadata) SetKeyId(v string) *CreateKeyResponseBodyKeyMetadata {
	s.KeyId = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetNextRotationDate(v string) *CreateKeyResponseBodyKeyMetadata {
	s.NextRotationDate = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetKeyState(v string) *CreateKeyResponseBodyKeyMetadata {
	s.KeyState = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetRotationInterval(v string) *CreateKeyResponseBodyKeyMetadata {
	s.RotationInterval = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetArn(v string) *CreateKeyResponseBodyKeyMetadata {
	s.Arn = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetCreator(v string) *CreateKeyResponseBodyKeyMetadata {
	s.Creator = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetDeleteDate(v string) *CreateKeyResponseBodyKeyMetadata {
	s.DeleteDate = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetLastRotationDate(v string) *CreateKeyResponseBodyKeyMetadata {
	s.LastRotationDate = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetDescription(v string) *CreateKeyResponseBodyKeyMetadata {
	s.Description = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetPrimaryKeyVersion(v string) *CreateKeyResponseBodyKeyMetadata {
	s.PrimaryKeyVersion = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetOrigin(v string) *CreateKeyResponseBodyKeyMetadata {
	s.Origin = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetKeySpec(v string) *CreateKeyResponseBodyKeyMetadata {
	s.KeySpec = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetMaterialExpireTime(v string) *CreateKeyResponseBodyKeyMetadata {
	s.MaterialExpireTime = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetAutomaticRotation(v string) *CreateKeyResponseBodyKeyMetadata {
	s.AutomaticRotation = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetProtectionLevel(v string) *CreateKeyResponseBodyKeyMetadata {
	s.ProtectionLevel = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetKeyUsage(v string) *CreateKeyResponseBodyKeyMetadata {
	s.KeyUsage = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetCreationDate(v string) *CreateKeyResponseBodyKeyMetadata {
	s.CreationDate = &v
	return s
}

type CreateKeyResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateKeyResponse) SetBody(v *CreateKeyResponseBody) *CreateKeyResponse {
	s.Body = v
	return s
}

type CreateKeyVersionRequest struct {
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
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	KeyVersion *CreateKeyVersionResponseBodyKeyVersion `json:"KeyVersion,omitempty" xml:"KeyVersion,omitempty" type:"Struct"`
}

func (s CreateKeyVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKeyVersionResponseBody) SetRequestId(v string) *CreateKeyVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKeyVersionResponseBody) SetKeyVersion(v *CreateKeyVersionResponseBodyKeyVersion) *CreateKeyVersionResponseBody {
	s.KeyVersion = v
	return s
}

type CreateKeyVersionResponseBodyKeyVersion struct {
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
}

func (s CreateKeyVersionResponseBodyKeyVersion) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyVersionResponseBodyKeyVersion) GoString() string {
	return s.String()
}

func (s *CreateKeyVersionResponseBodyKeyVersion) SetKeyVersionId(v string) *CreateKeyVersionResponseBodyKeyVersion {
	s.KeyVersionId = &v
	return s
}

func (s *CreateKeyVersionResponseBodyKeyVersion) SetKeyId(v string) *CreateKeyVersionResponseBodyKeyVersion {
	s.KeyId = &v
	return s
}

func (s *CreateKeyVersionResponseBodyKeyVersion) SetCreationDate(v string) *CreateKeyVersionResponseBodyKeyVersion {
	s.CreationDate = &v
	return s
}

type CreateKeyVersionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateKeyVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateKeyVersionResponse) SetBody(v *CreateKeyVersionResponseBody) *CreateKeyVersionResponse {
	s.Body = v
	return s
}

type CreateSecretRequest struct {
	SecretName              *string                `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	VersionId               *string                `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	EncryptionKeyId         *string                `json:"EncryptionKeyId,omitempty" xml:"EncryptionKeyId,omitempty"`
	SecretData              *string                `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
	SecretDataType          *string                `json:"SecretDataType,omitempty" xml:"SecretDataType,omitempty"`
	Description             *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	Tags                    *string                `json:"Tags,omitempty" xml:"Tags,omitempty"`
	SecretType              *string                `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	ExtendedConfig          map[string]interface{} `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	EnableAutomaticRotation *bool                  `json:"EnableAutomaticRotation,omitempty" xml:"EnableAutomaticRotation,omitempty"`
	RotationInterval        *string                `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
}

func (s CreateSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSecretRequest) GoString() string {
	return s.String()
}

func (s *CreateSecretRequest) SetSecretName(v string) *CreateSecretRequest {
	s.SecretName = &v
	return s
}

func (s *CreateSecretRequest) SetVersionId(v string) *CreateSecretRequest {
	s.VersionId = &v
	return s
}

func (s *CreateSecretRequest) SetEncryptionKeyId(v string) *CreateSecretRequest {
	s.EncryptionKeyId = &v
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

func (s *CreateSecretRequest) SetDescription(v string) *CreateSecretRequest {
	s.Description = &v
	return s
}

func (s *CreateSecretRequest) SetTags(v string) *CreateSecretRequest {
	s.Tags = &v
	return s
}

func (s *CreateSecretRequest) SetSecretType(v string) *CreateSecretRequest {
	s.SecretType = &v
	return s
}

func (s *CreateSecretRequest) SetExtendedConfig(v map[string]interface{}) *CreateSecretRequest {
	s.ExtendedConfig = v
	return s
}

func (s *CreateSecretRequest) SetEnableAutomaticRotation(v bool) *CreateSecretRequest {
	s.EnableAutomaticRotation = &v
	return s
}

func (s *CreateSecretRequest) SetRotationInterval(v string) *CreateSecretRequest {
	s.RotationInterval = &v
	return s
}

type CreateSecretShrinkRequest struct {
	SecretName              *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	VersionId               *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	EncryptionKeyId         *string `json:"EncryptionKeyId,omitempty" xml:"EncryptionKeyId,omitempty"`
	SecretData              *string `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
	SecretDataType          *string `json:"SecretDataType,omitempty" xml:"SecretDataType,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Tags                    *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	SecretType              *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	ExtendedConfigShrink    *string `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	EnableAutomaticRotation *bool   `json:"EnableAutomaticRotation,omitempty" xml:"EnableAutomaticRotation,omitempty"`
	RotationInterval        *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
}

func (s CreateSecretShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSecretShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSecretShrinkRequest) SetSecretName(v string) *CreateSecretShrinkRequest {
	s.SecretName = &v
	return s
}

func (s *CreateSecretShrinkRequest) SetVersionId(v string) *CreateSecretShrinkRequest {
	s.VersionId = &v
	return s
}

func (s *CreateSecretShrinkRequest) SetEncryptionKeyId(v string) *CreateSecretShrinkRequest {
	s.EncryptionKeyId = &v
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

func (s *CreateSecretShrinkRequest) SetDescription(v string) *CreateSecretShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateSecretShrinkRequest) SetTags(v string) *CreateSecretShrinkRequest {
	s.Tags = &v
	return s
}

func (s *CreateSecretShrinkRequest) SetSecretType(v string) *CreateSecretShrinkRequest {
	s.SecretType = &v
	return s
}

func (s *CreateSecretShrinkRequest) SetExtendedConfigShrink(v string) *CreateSecretShrinkRequest {
	s.ExtendedConfigShrink = &v
	return s
}

func (s *CreateSecretShrinkRequest) SetEnableAutomaticRotation(v bool) *CreateSecretShrinkRequest {
	s.EnableAutomaticRotation = &v
	return s
}

func (s *CreateSecretShrinkRequest) SetRotationInterval(v string) *CreateSecretShrinkRequest {
	s.RotationInterval = &v
	return s
}

type CreateSecretResponseBody struct {
	SecretName        *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	VersionId         *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	NextRotationDate  *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecretType        *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	RotationInterval  *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	ExtendedConfig    *string `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	Arn               *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	AutomaticRotation *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
}

func (s CreateSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSecretResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecretResponseBody) SetSecretName(v string) *CreateSecretResponseBody {
	s.SecretName = &v
	return s
}

func (s *CreateSecretResponseBody) SetVersionId(v string) *CreateSecretResponseBody {
	s.VersionId = &v
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

func (s *CreateSecretResponseBody) SetSecretType(v string) *CreateSecretResponseBody {
	s.SecretType = &v
	return s
}

func (s *CreateSecretResponseBody) SetRotationInterval(v string) *CreateSecretResponseBody {
	s.RotationInterval = &v
	return s
}

func (s *CreateSecretResponseBody) SetExtendedConfig(v string) *CreateSecretResponseBody {
	s.ExtendedConfig = &v
	return s
}

func (s *CreateSecretResponseBody) SetArn(v string) *CreateSecretResponseBody {
	s.Arn = &v
	return s
}

func (s *CreateSecretResponseBody) SetAutomaticRotation(v string) *CreateSecretResponseBody {
	s.AutomaticRotation = &v
	return s
}

type CreateSecretResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSecretResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateSecretResponse) SetBody(v *CreateSecretResponseBody) *CreateSecretResponse {
	s.Body = v
	return s
}

type DecryptRequest struct {
	CiphertextBlob    *string                `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
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
	CiphertextBlob          *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
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
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Plaintext    *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s DecryptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DecryptResponseBody) GoString() string {
	return s.String()
}

func (s *DecryptResponseBody) SetRequestId(v string) *DecryptResponseBody {
	s.RequestId = &v
	return s
}

func (s *DecryptResponseBody) SetPlaintext(v string) *DecryptResponseBody {
	s.Plaintext = &v
	return s
}

func (s *DecryptResponseBody) SetKeyId(v string) *DecryptResponseBody {
	s.KeyId = &v
	return s
}

func (s *DecryptResponseBody) SetKeyVersionId(v string) *DecryptResponseBody {
	s.KeyVersionId = &v
	return s
}

type DecryptResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DecryptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DecryptResponse) SetBody(v *DecryptResponseBody) *DecryptResponse {
	s.Body = v
	return s
}

type DeleteAliasRequest struct {
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAliasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteAliasResponse) SetBody(v *DeleteAliasResponseBody) *DeleteAliasResponse {
	s.Body = v
	return s
}

type DeleteCertificateRequest struct {
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteCertificateResponse) SetBody(v *DeleteCertificateResponseBody) *DeleteCertificateResponse {
	s.Body = v
	return s
}

type DeleteKeyMaterialRequest struct {
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteKeyMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteKeyMaterialResponse) SetBody(v *DeleteKeyMaterialResponseBody) *DeleteKeyMaterialResponse {
	s.Body = v
	return s
}

type DeleteSecretRequest struct {
	SecretName                 *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	ForceDeleteWithoutRecovery *string `json:"ForceDeleteWithoutRecovery,omitempty" xml:"ForceDeleteWithoutRecovery,omitempty"`
	RecoveryWindowInDays       *string `json:"RecoveryWindowInDays,omitempty" xml:"RecoveryWindowInDays,omitempty"`
}

func (s DeleteSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecretRequest) SetSecretName(v string) *DeleteSecretRequest {
	s.SecretName = &v
	return s
}

func (s *DeleteSecretRequest) SetForceDeleteWithoutRecovery(v string) *DeleteSecretRequest {
	s.ForceDeleteWithoutRecovery = &v
	return s
}

func (s *DeleteSecretRequest) SetRecoveryWindowInDays(v string) *DeleteSecretRequest {
	s.RecoveryWindowInDays = &v
	return s
}

type DeleteSecretResponseBody struct {
	SecretName        *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PlannedDeleteTime *string `json:"PlannedDeleteTime,omitempty" xml:"PlannedDeleteTime,omitempty"`
}

func (s DeleteSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecretResponseBody) SetSecretName(v string) *DeleteSecretResponseBody {
	s.SecretName = &v
	return s
}

func (s *DeleteSecretResponseBody) SetRequestId(v string) *DeleteSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecretResponseBody) SetPlannedDeleteTime(v string) *DeleteSecretResponseBody {
	s.PlannedDeleteTime = &v
	return s
}

type DeleteSecretResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSecretResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteSecretResponse) SetBody(v *DeleteSecretResponseBody) *DeleteSecretResponse {
	s.Body = v
	return s
}

type DescribeAccountKmsStatusResponseBody struct {
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAccountKmsStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeAccountKmsStatusResponse) SetBody(v *DescribeAccountKmsStatusResponseBody) *DescribeAccountKmsStatusResponse {
	s.Body = v
	return s
}

type DescribeCertificateRequest struct {
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
	Status                  *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	ProtectionLevel         *string                `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	RequestId               *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Issuer                  *string                `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	CertificateId           *string                `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	CreatedAt               *string                `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	KeySpec                 *string                `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	SubjectAlternativeNames []*string              `json:"SubjectAlternativeNames,omitempty" xml:"SubjectAlternativeNames,omitempty" type:"Repeated"`
	SignatureAlgorithm      *string                `json:"SignatureAlgorithm,omitempty" xml:"SignatureAlgorithm,omitempty"`
	SubjectKeyIdentifier    *string                `json:"SubjectKeyIdentifier,omitempty" xml:"SubjectKeyIdentifier,omitempty"`
	NotAfter                *string                `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	UpdatedAt               *string                `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Subject                 *string                `json:"Subject,omitempty" xml:"Subject,omitempty"`
	Serial                  *string                `json:"Serial,omitempty" xml:"Serial,omitempty"`
	SubjectPublicKey        *string                `json:"SubjectPublicKey,omitempty" xml:"SubjectPublicKey,omitempty"`
	Arn                     *string                `json:"Arn,omitempty" xml:"Arn,omitempty"`
	NotBefore               *string                `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	Tags                    map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s DescribeCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCertificateResponseBody) SetStatus(v string) *DescribeCertificateResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetProtectionLevel(v string) *DescribeCertificateResponseBody {
	s.ProtectionLevel = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetRequestId(v string) *DescribeCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetIssuer(v string) *DescribeCertificateResponseBody {
	s.Issuer = &v
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

func (s *DescribeCertificateResponseBody) SetKeySpec(v string) *DescribeCertificateResponseBody {
	s.KeySpec = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetSubjectAlternativeNames(v []*string) *DescribeCertificateResponseBody {
	s.SubjectAlternativeNames = v
	return s
}

func (s *DescribeCertificateResponseBody) SetSignatureAlgorithm(v string) *DescribeCertificateResponseBody {
	s.SignatureAlgorithm = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetSubjectKeyIdentifier(v string) *DescribeCertificateResponseBody {
	s.SubjectKeyIdentifier = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetNotAfter(v string) *DescribeCertificateResponseBody {
	s.NotAfter = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetUpdatedAt(v string) *DescribeCertificateResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetSubject(v string) *DescribeCertificateResponseBody {
	s.Subject = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetSerial(v string) *DescribeCertificateResponseBody {
	s.Serial = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetSubjectPublicKey(v string) *DescribeCertificateResponseBody {
	s.SubjectPublicKey = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetArn(v string) *DescribeCertificateResponseBody {
	s.Arn = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetNotBefore(v string) *DescribeCertificateResponseBody {
	s.NotBefore = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetTags(v map[string]interface{}) *DescribeCertificateResponseBody {
	s.Tags = v
	return s
}

type DescribeCertificateResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeCertificateResponse) SetBody(v *DescribeCertificateResponseBody) *DescribeCertificateResponse {
	s.Body = v
	return s
}

type DescribeKeyRequest struct {
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
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	KeyMetadata *DescribeKeyResponseBodyKeyMetadata `json:"KeyMetadata,omitempty" xml:"KeyMetadata,omitempty" type:"Struct"`
}

func (s DescribeKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKeyResponseBody) SetRequestId(v string) *DescribeKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKeyResponseBody) SetKeyMetadata(v *DescribeKeyResponseBodyKeyMetadata) *DescribeKeyResponseBody {
	s.KeyMetadata = v
	return s
}

type DescribeKeyResponseBodyKeyMetadata struct {
	KeyId              *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	NextRotationDate   *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	KeyState           *string `json:"KeyState,omitempty" xml:"KeyState,omitempty"`
	RotationInterval   *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	Arn                *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	Creator            *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DeleteDate         *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
	LastRotationDate   *string `json:"LastRotationDate,omitempty" xml:"LastRotationDate,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PrimaryKeyVersion  *string `json:"PrimaryKeyVersion,omitempty" xml:"PrimaryKeyVersion,omitempty"`
	Origin             *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	KeySpec            *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	MaterialExpireTime *string `json:"MaterialExpireTime,omitempty" xml:"MaterialExpireTime,omitempty"`
	AutomaticRotation  *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	ProtectionLevel    *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	KeyUsage           *string `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty"`
	CreationDate       *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
}

func (s DescribeKeyResponseBodyKeyMetadata) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyResponseBodyKeyMetadata) GoString() string {
	return s.String()
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetKeyId(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.KeyId = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetNextRotationDate(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.NextRotationDate = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetKeyState(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.KeyState = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetRotationInterval(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.RotationInterval = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetArn(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.Arn = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetCreator(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.Creator = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetDeleteDate(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.DeleteDate = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetLastRotationDate(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.LastRotationDate = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetDescription(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.Description = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetPrimaryKeyVersion(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.PrimaryKeyVersion = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetOrigin(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.Origin = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetKeySpec(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.KeySpec = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetMaterialExpireTime(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.MaterialExpireTime = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetAutomaticRotation(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.AutomaticRotation = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetProtectionLevel(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.ProtectionLevel = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetKeyUsage(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.KeyUsage = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetCreationDate(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.CreationDate = &v
	return s
}

type DescribeKeyResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeKeyResponse) SetBody(v *DescribeKeyResponseBody) *DescribeKeyResponse {
	s.Body = v
	return s
}

type DescribeKeyVersionRequest struct {
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
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
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	KeyVersion *DescribeKeyVersionResponseBodyKeyVersion `json:"KeyVersion,omitempty" xml:"KeyVersion,omitempty" type:"Struct"`
}

func (s DescribeKeyVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKeyVersionResponseBody) SetRequestId(v string) *DescribeKeyVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKeyVersionResponseBody) SetKeyVersion(v *DescribeKeyVersionResponseBodyKeyVersion) *DescribeKeyVersionResponseBody {
	s.KeyVersion = v
	return s
}

type DescribeKeyVersionResponseBodyKeyVersion struct {
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
}

func (s DescribeKeyVersionResponseBodyKeyVersion) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyVersionResponseBodyKeyVersion) GoString() string {
	return s.String()
}

func (s *DescribeKeyVersionResponseBodyKeyVersion) SetKeyVersionId(v string) *DescribeKeyVersionResponseBodyKeyVersion {
	s.KeyVersionId = &v
	return s
}

func (s *DescribeKeyVersionResponseBodyKeyVersion) SetKeyId(v string) *DescribeKeyVersionResponseBodyKeyVersion {
	s.KeyId = &v
	return s
}

func (s *DescribeKeyVersionResponseBodyKeyVersion) SetCreationDate(v string) *DescribeKeyVersionResponseBodyKeyVersion {
	s.CreationDate = &v
	return s
}

type DescribeKeyVersionResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeKeyVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeKeyVersionResponse) SetBody(v *DescribeKeyVersionResponseBody) *DescribeKeyVersionResponse {
	s.Body = v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeSecretRequest struct {
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	FetchTags  *string `json:"FetchTags,omitempty" xml:"FetchTags,omitempty"`
}

func (s DescribeSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecretRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecretRequest) SetSecretName(v string) *DescribeSecretRequest {
	s.SecretName = &v
	return s
}

func (s *DescribeSecretRequest) SetFetchTags(v string) *DescribeSecretRequest {
	s.FetchTags = &v
	return s
}

type DescribeSecretResponseBody struct {
	Description       *string                         `json:"Description,omitempty" xml:"Description,omitempty"`
	RotationInterval  *string                         `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	LastRotationDate  *string                         `json:"LastRotationDate,omitempty" xml:"LastRotationDate,omitempty"`
	RequestId         *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecretType        *string                         `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	CreateTime        *string                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	AutomaticRotation *string                         `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	SecretName        *string                         `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	NextRotationDate  *string                         `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	UpdateTime        *string                         `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ExtendedConfig    *string                         `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	PlannedDeleteTime *string                         `json:"PlannedDeleteTime,omitempty" xml:"PlannedDeleteTime,omitempty"`
	Arn               *string                         `json:"Arn,omitempty" xml:"Arn,omitempty"`
	Tags              *DescribeSecretResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	EncryptionKeyId   *string                         `json:"EncryptionKeyId,omitempty" xml:"EncryptionKeyId,omitempty"`
}

func (s DescribeSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecretResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecretResponseBody) SetDescription(v string) *DescribeSecretResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeSecretResponseBody) SetRotationInterval(v string) *DescribeSecretResponseBody {
	s.RotationInterval = &v
	return s
}

func (s *DescribeSecretResponseBody) SetLastRotationDate(v string) *DescribeSecretResponseBody {
	s.LastRotationDate = &v
	return s
}

func (s *DescribeSecretResponseBody) SetRequestId(v string) *DescribeSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecretResponseBody) SetSecretType(v string) *DescribeSecretResponseBody {
	s.SecretType = &v
	return s
}

func (s *DescribeSecretResponseBody) SetCreateTime(v string) *DescribeSecretResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeSecretResponseBody) SetAutomaticRotation(v string) *DescribeSecretResponseBody {
	s.AutomaticRotation = &v
	return s
}

func (s *DescribeSecretResponseBody) SetSecretName(v string) *DescribeSecretResponseBody {
	s.SecretName = &v
	return s
}

func (s *DescribeSecretResponseBody) SetNextRotationDate(v string) *DescribeSecretResponseBody {
	s.NextRotationDate = &v
	return s
}

func (s *DescribeSecretResponseBody) SetUpdateTime(v string) *DescribeSecretResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeSecretResponseBody) SetExtendedConfig(v string) *DescribeSecretResponseBody {
	s.ExtendedConfig = &v
	return s
}

func (s *DescribeSecretResponseBody) SetPlannedDeleteTime(v string) *DescribeSecretResponseBody {
	s.PlannedDeleteTime = &v
	return s
}

func (s *DescribeSecretResponseBody) SetArn(v string) *DescribeSecretResponseBody {
	s.Arn = &v
	return s
}

func (s *DescribeSecretResponseBody) SetTags(v *DescribeSecretResponseBodyTags) *DescribeSecretResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeSecretResponseBody) SetEncryptionKeyId(v string) *DescribeSecretResponseBody {
	s.EncryptionKeyId = &v
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
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s DescribeSecretResponseBodyTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecretResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeSecretResponseBodyTagsTag) SetTagValue(v string) *DescribeSecretResponseBodyTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeSecretResponseBodyTagsTag) SetTagKey(v string) *DescribeSecretResponseBodyTagsTag {
	s.TagKey = &v
	return s
}

type DescribeSecretResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSecretResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeSecretResponse) SetBody(v *DescribeSecretResponseBody) *DescribeSecretResponse {
	s.Body = v
	return s
}

type DescribeServiceResponseBody struct {
	ProtectionLevels *DescribeServiceResponseBodyProtectionLevels `json:"ProtectionLevels,omitempty" xml:"ProtectionLevels,omitempty" type:"Struct"`
	RequestId        *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	KeySpecs         *DescribeServiceResponseBodyKeySpecs         `json:"KeySpecs,omitempty" xml:"KeySpecs,omitempty" type:"Struct"`
}

func (s DescribeServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceResponseBody) SetProtectionLevels(v *DescribeServiceResponseBodyProtectionLevels) *DescribeServiceResponseBody {
	s.ProtectionLevels = v
	return s
}

func (s *DescribeServiceResponseBody) SetRequestId(v string) *DescribeServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceResponseBody) SetKeySpecs(v *DescribeServiceResponseBodyKeySpecs) *DescribeServiceResponseBody {
	s.KeySpecs = v
	return s
}

type DescribeServiceResponseBodyProtectionLevels struct {
	ProtectionLevel []*DescribeServiceResponseBodyProtectionLevelsProtectionLevel `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty" type:"Repeated"`
}

func (s DescribeServiceResponseBodyProtectionLevels) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceResponseBodyProtectionLevels) GoString() string {
	return s.String()
}

func (s *DescribeServiceResponseBodyProtectionLevels) SetProtectionLevel(v []*DescribeServiceResponseBodyProtectionLevelsProtectionLevel) *DescribeServiceResponseBodyProtectionLevels {
	s.ProtectionLevel = v
	return s
}

type DescribeServiceResponseBodyProtectionLevelsProtectionLevel struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeServiceResponseBodyProtectionLevelsProtectionLevel) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceResponseBodyProtectionLevelsProtectionLevel) GoString() string {
	return s.String()
}

func (s *DescribeServiceResponseBodyProtectionLevelsProtectionLevel) SetType(v string) *DescribeServiceResponseBodyProtectionLevelsProtectionLevel {
	s.Type = &v
	return s
}

type DescribeServiceResponseBodyKeySpecs struct {
	KeySpec []*DescribeServiceResponseBodyKeySpecsKeySpec `json:"KeySpec,omitempty" xml:"KeySpec,omitempty" type:"Repeated"`
}

func (s DescribeServiceResponseBodyKeySpecs) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceResponseBodyKeySpecs) GoString() string {
	return s.String()
}

func (s *DescribeServiceResponseBodyKeySpecs) SetKeySpec(v []*DescribeServiceResponseBodyKeySpecsKeySpec) *DescribeServiceResponseBodyKeySpecs {
	s.KeySpec = v
	return s
}

type DescribeServiceResponseBodyKeySpecsKeySpec struct {
	SupportedProtectionLevels *DescribeServiceResponseBodyKeySpecsKeySpecSupportedProtectionLevels `json:"SupportedProtectionLevels,omitempty" xml:"SupportedProtectionLevels,omitempty" type:"Struct"`
	Name                      *string                                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	Usages                    *DescribeServiceResponseBodyKeySpecsKeySpecUsages                    `json:"Usages,omitempty" xml:"Usages,omitempty" type:"Struct"`
}

func (s DescribeServiceResponseBodyKeySpecsKeySpec) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceResponseBodyKeySpecsKeySpec) GoString() string {
	return s.String()
}

func (s *DescribeServiceResponseBodyKeySpecsKeySpec) SetSupportedProtectionLevels(v *DescribeServiceResponseBodyKeySpecsKeySpecSupportedProtectionLevels) *DescribeServiceResponseBodyKeySpecsKeySpec {
	s.SupportedProtectionLevels = v
	return s
}

func (s *DescribeServiceResponseBodyKeySpecsKeySpec) SetName(v string) *DescribeServiceResponseBodyKeySpecsKeySpec {
	s.Name = &v
	return s
}

func (s *DescribeServiceResponseBodyKeySpecsKeySpec) SetUsages(v *DescribeServiceResponseBodyKeySpecsKeySpecUsages) *DescribeServiceResponseBodyKeySpecsKeySpec {
	s.Usages = v
	return s
}

type DescribeServiceResponseBodyKeySpecsKeySpecSupportedProtectionLevels struct {
	SupportedProtectionLevel []*string `json:"SupportedProtectionLevel,omitempty" xml:"SupportedProtectionLevel,omitempty" type:"Repeated"`
}

func (s DescribeServiceResponseBodyKeySpecsKeySpecSupportedProtectionLevels) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceResponseBodyKeySpecsKeySpecSupportedProtectionLevels) GoString() string {
	return s.String()
}

func (s *DescribeServiceResponseBodyKeySpecsKeySpecSupportedProtectionLevels) SetSupportedProtectionLevel(v []*string) *DescribeServiceResponseBodyKeySpecsKeySpecSupportedProtectionLevels {
	s.SupportedProtectionLevel = v
	return s
}

type DescribeServiceResponseBodyKeySpecsKeySpecUsages struct {
	Usage []*string `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Repeated"`
}

func (s DescribeServiceResponseBodyKeySpecsKeySpecUsages) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceResponseBodyKeySpecsKeySpecUsages) GoString() string {
	return s.String()
}

func (s *DescribeServiceResponseBodyKeySpecsKeySpecUsages) SetUsage(v []*string) *DescribeServiceResponseBodyKeySpecsKeySpecUsages {
	s.Usage = v
	return s
}

type DescribeServiceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceResponse) SetHeaders(v map[string]*string) *DescribeServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceResponse) SetBody(v *DescribeServiceResponseBody) *DescribeServiceResponse {
	s.Body = v
	return s
}

type DisableKeyRequest struct {
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
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DisableKeyResponse) SetBody(v *DisableKeyResponseBody) *DisableKeyResponse {
	s.Body = v
	return s
}

type EnableKeyRequest struct {
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *EnableKeyResponse) SetBody(v *EnableKeyResponseBody) *EnableKeyResponse {
	s.Body = v
	return s
}

type EncryptRequest struct {
	KeyId             *string                `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Plaintext         *string                `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
}

func (s EncryptRequest) String() string {
	return tea.Prettify(s)
}

func (s EncryptRequest) GoString() string {
	return s.String()
}

func (s *EncryptRequest) SetKeyId(v string) *EncryptRequest {
	s.KeyId = &v
	return s
}

func (s *EncryptRequest) SetPlaintext(v string) *EncryptRequest {
	s.Plaintext = &v
	return s
}

func (s *EncryptRequest) SetEncryptionContext(v map[string]interface{}) *EncryptRequest {
	s.EncryptionContext = v
	return s
}

type EncryptShrinkRequest struct {
	KeyId                   *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Plaintext               *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
}

func (s EncryptShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s EncryptShrinkRequest) GoString() string {
	return s.String()
}

func (s *EncryptShrinkRequest) SetKeyId(v string) *EncryptShrinkRequest {
	s.KeyId = &v
	return s
}

func (s *EncryptShrinkRequest) SetPlaintext(v string) *EncryptShrinkRequest {
	s.Plaintext = &v
	return s
}

func (s *EncryptShrinkRequest) SetEncryptionContextShrink(v string) *EncryptShrinkRequest {
	s.EncryptionContextShrink = &v
	return s
}

type EncryptResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	KeyId          *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeyVersionId   *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s EncryptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EncryptResponseBody) GoString() string {
	return s.String()
}

func (s *EncryptResponseBody) SetRequestId(v string) *EncryptResponseBody {
	s.RequestId = &v
	return s
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

type EncryptResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EncryptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *EncryptResponse) SetBody(v *EncryptResponseBody) *EncryptResponse {
	s.Body = v
	return s
}

type ExportCertificateRequest struct {
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	Passphrase    *string `json:"Passphrase,omitempty" xml:"Passphrase,omitempty"`
	ExportFormat  *string `json:"ExportFormat,omitempty" xml:"ExportFormat,omitempty"`
}

func (s ExportCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportCertificateRequest) GoString() string {
	return s.String()
}

func (s *ExportCertificateRequest) SetCertificateId(v string) *ExportCertificateRequest {
	s.CertificateId = &v
	return s
}

func (s *ExportCertificateRequest) SetPassphrase(v string) *ExportCertificateRequest {
	s.Passphrase = &v
	return s
}

func (s *ExportCertificateRequest) SetExportFormat(v string) *ExportCertificateRequest {
	s.ExportFormat = &v
	return s
}

type ExportCertificateResponseBody struct {
	PrivateKey       *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CertificateId    *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	PKCS12Blob       *string `json:"PKCS12Blob,omitempty" xml:"PKCS12Blob,omitempty"`
	Certificate      *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
}

func (s ExportCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ExportCertificateResponseBody) SetPrivateKey(v string) *ExportCertificateResponseBody {
	s.PrivateKey = &v
	return s
}

func (s *ExportCertificateResponseBody) SetRequestId(v string) *ExportCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportCertificateResponseBody) SetCertificateId(v string) *ExportCertificateResponseBody {
	s.CertificateId = &v
	return s
}

func (s *ExportCertificateResponseBody) SetCertificateChain(v string) *ExportCertificateResponseBody {
	s.CertificateChain = &v
	return s
}

func (s *ExportCertificateResponseBody) SetPKCS12Blob(v string) *ExportCertificateResponseBody {
	s.PKCS12Blob = &v
	return s
}

func (s *ExportCertificateResponseBody) SetCertificate(v string) *ExportCertificateResponseBody {
	s.Certificate = &v
	return s
}

type ExportCertificateResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExportCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExportCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportCertificateResponse) GoString() string {
	return s.String()
}

func (s *ExportCertificateResponse) SetHeaders(v map[string]*string) *ExportCertificateResponse {
	s.Headers = v
	return s
}

func (s *ExportCertificateResponse) SetBody(v *ExportCertificateResponseBody) *ExportCertificateResponse {
	s.Body = v
	return s
}

type ExportDataKeyRequest struct {
	CiphertextBlob    *string                `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	PublicKeyBlob     *string                `json:"PublicKeyBlob,omitempty" xml:"PublicKeyBlob,omitempty"`
	WrappingKeySpec   *string                `json:"WrappingKeySpec,omitempty" xml:"WrappingKeySpec,omitempty"`
	WrappingAlgorithm *string                `json:"WrappingAlgorithm,omitempty" xml:"WrappingAlgorithm,omitempty"`
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

func (s *ExportDataKeyRequest) SetWrappingKeySpec(v string) *ExportDataKeyRequest {
	s.WrappingKeySpec = &v
	return s
}

func (s *ExportDataKeyRequest) SetWrappingAlgorithm(v string) *ExportDataKeyRequest {
	s.WrappingAlgorithm = &v
	return s
}

type ExportDataKeyShrinkRequest struct {
	CiphertextBlob          *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	PublicKeyBlob           *string `json:"PublicKeyBlob,omitempty" xml:"PublicKeyBlob,omitempty"`
	WrappingKeySpec         *string `json:"WrappingKeySpec,omitempty" xml:"WrappingKeySpec,omitempty"`
	WrappingAlgorithm       *string `json:"WrappingAlgorithm,omitempty" xml:"WrappingAlgorithm,omitempty"`
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

func (s *ExportDataKeyShrinkRequest) SetWrappingKeySpec(v string) *ExportDataKeyShrinkRequest {
	s.WrappingKeySpec = &v
	return s
}

func (s *ExportDataKeyShrinkRequest) SetWrappingAlgorithm(v string) *ExportDataKeyShrinkRequest {
	s.WrappingAlgorithm = &v
	return s
}

type ExportDataKeyResponseBody struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ExportedDataKey *string `json:"ExportedDataKey,omitempty" xml:"ExportedDataKey,omitempty"`
	KeyId           *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeyVersionId    *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s ExportDataKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportDataKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ExportDataKeyResponseBody) SetRequestId(v string) *ExportDataKeyResponseBody {
	s.RequestId = &v
	return s
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

type ExportDataKeyResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExportDataKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ExportDataKeyResponse) SetBody(v *ExportDataKeyResponseBody) *ExportDataKeyResponse {
	s.Body = v
	return s
}

type GenerateAndExportDataKeyRequest struct {
	KeyId             *string                `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeySpec           *string                `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	NumberOfBytes     *int32                 `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	PublicKeyBlob     *string                `json:"PublicKeyBlob,omitempty" xml:"PublicKeyBlob,omitempty"`
	WrappingKeySpec   *string                `json:"WrappingKeySpec,omitempty" xml:"WrappingKeySpec,omitempty"`
	WrappingAlgorithm *string                `json:"WrappingAlgorithm,omitempty" xml:"WrappingAlgorithm,omitempty"`
}

func (s GenerateAndExportDataKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateAndExportDataKeyRequest) GoString() string {
	return s.String()
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

func (s *GenerateAndExportDataKeyRequest) SetEncryptionContext(v map[string]interface{}) *GenerateAndExportDataKeyRequest {
	s.EncryptionContext = v
	return s
}

func (s *GenerateAndExportDataKeyRequest) SetPublicKeyBlob(v string) *GenerateAndExportDataKeyRequest {
	s.PublicKeyBlob = &v
	return s
}

func (s *GenerateAndExportDataKeyRequest) SetWrappingKeySpec(v string) *GenerateAndExportDataKeyRequest {
	s.WrappingKeySpec = &v
	return s
}

func (s *GenerateAndExportDataKeyRequest) SetWrappingAlgorithm(v string) *GenerateAndExportDataKeyRequest {
	s.WrappingAlgorithm = &v
	return s
}

type GenerateAndExportDataKeyShrinkRequest struct {
	KeyId                   *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeySpec                 *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	NumberOfBytes           *int32  `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
	EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	PublicKeyBlob           *string `json:"PublicKeyBlob,omitempty" xml:"PublicKeyBlob,omitempty"`
	WrappingKeySpec         *string `json:"WrappingKeySpec,omitempty" xml:"WrappingKeySpec,omitempty"`
	WrappingAlgorithm       *string `json:"WrappingAlgorithm,omitempty" xml:"WrappingAlgorithm,omitempty"`
}

func (s GenerateAndExportDataKeyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateAndExportDataKeyShrinkRequest) GoString() string {
	return s.String()
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

func (s *GenerateAndExportDataKeyShrinkRequest) SetEncryptionContextShrink(v string) *GenerateAndExportDataKeyShrinkRequest {
	s.EncryptionContextShrink = &v
	return s
}

func (s *GenerateAndExportDataKeyShrinkRequest) SetPublicKeyBlob(v string) *GenerateAndExportDataKeyShrinkRequest {
	s.PublicKeyBlob = &v
	return s
}

func (s *GenerateAndExportDataKeyShrinkRequest) SetWrappingKeySpec(v string) *GenerateAndExportDataKeyShrinkRequest {
	s.WrappingKeySpec = &v
	return s
}

func (s *GenerateAndExportDataKeyShrinkRequest) SetWrappingAlgorithm(v string) *GenerateAndExportDataKeyShrinkRequest {
	s.WrappingAlgorithm = &v
	return s
}

type GenerateAndExportDataKeyResponseBody struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CiphertextBlob  *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	ExportedDataKey *string `json:"ExportedDataKey,omitempty" xml:"ExportedDataKey,omitempty"`
	KeyId           *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeyVersionId    *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s GenerateAndExportDataKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateAndExportDataKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateAndExportDataKeyResponseBody) SetRequestId(v string) *GenerateAndExportDataKeyResponseBody {
	s.RequestId = &v
	return s
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

type GenerateAndExportDataKeyResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateAndExportDataKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GenerateAndExportDataKeyResponse) SetBody(v *GenerateAndExportDataKeyResponseBody) *GenerateAndExportDataKeyResponse {
	s.Body = v
	return s
}

type GenerateDataKeyRequest struct {
	KeyId             *string                `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeySpec           *string                `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	NumberOfBytes     *int32                 `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
}

func (s GenerateDataKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateDataKeyRequest) GoString() string {
	return s.String()
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

func (s *GenerateDataKeyRequest) SetEncryptionContext(v map[string]interface{}) *GenerateDataKeyRequest {
	s.EncryptionContext = v
	return s
}

type GenerateDataKeyShrinkRequest struct {
	KeyId                   *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeySpec                 *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	NumberOfBytes           *int32  `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
	EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
}

func (s GenerateDataKeyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateDataKeyShrinkRequest) GoString() string {
	return s.String()
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

func (s *GenerateDataKeyShrinkRequest) SetEncryptionContextShrink(v string) *GenerateDataKeyShrinkRequest {
	s.EncryptionContextShrink = &v
	return s
}

type GenerateDataKeyResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	Plaintext      *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	KeyId          *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeyVersionId   *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s GenerateDataKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateDataKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyResponseBody) SetRequestId(v string) *GenerateDataKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateDataKeyResponseBody) SetCiphertextBlob(v string) *GenerateDataKeyResponseBody {
	s.CiphertextBlob = &v
	return s
}

func (s *GenerateDataKeyResponseBody) SetPlaintext(v string) *GenerateDataKeyResponseBody {
	s.Plaintext = &v
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

type GenerateDataKeyResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateDataKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GenerateDataKeyResponse) SetBody(v *GenerateDataKeyResponseBody) *GenerateDataKeyResponse {
	s.Body = v
	return s
}

type GenerateDataKeyWithoutPlaintextRequest struct {
	KeyId             *string                `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeySpec           *string                `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	NumberOfBytes     *int32                 `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
}

func (s GenerateDataKeyWithoutPlaintextRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateDataKeyWithoutPlaintextRequest) GoString() string {
	return s.String()
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

func (s *GenerateDataKeyWithoutPlaintextRequest) SetEncryptionContext(v map[string]interface{}) *GenerateDataKeyWithoutPlaintextRequest {
	s.EncryptionContext = v
	return s
}

type GenerateDataKeyWithoutPlaintextShrinkRequest struct {
	KeyId                   *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeySpec                 *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	NumberOfBytes           *int32  `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
	EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
}

func (s GenerateDataKeyWithoutPlaintextShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateDataKeyWithoutPlaintextShrinkRequest) GoString() string {
	return s.String()
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

func (s *GenerateDataKeyWithoutPlaintextShrinkRequest) SetEncryptionContextShrink(v string) *GenerateDataKeyWithoutPlaintextShrinkRequest {
	s.EncryptionContextShrink = &v
	return s
}

type GenerateDataKeyWithoutPlaintextResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	KeyId          *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeyVersionId   *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s GenerateDataKeyWithoutPlaintextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateDataKeyWithoutPlaintextResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyWithoutPlaintextResponseBody) SetRequestId(v string) *GenerateDataKeyWithoutPlaintextResponseBody {
	s.RequestId = &v
	return s
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

type GenerateDataKeyWithoutPlaintextResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateDataKeyWithoutPlaintextResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GenerateDataKeyWithoutPlaintextResponse) SetBody(v *GenerateDataKeyWithoutPlaintextResponseBody) *GenerateDataKeyWithoutPlaintextResponse {
	s.Body = v
	return s
}

type GetCertificateRequest struct {
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
	Csr              *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CertificateId    *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	Certificate      *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
}

func (s GetCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *GetCertificateResponseBody) SetCsr(v string) *GetCertificateResponseBody {
	s.Csr = &v
	return s
}

func (s *GetCertificateResponseBody) SetRequestId(v string) *GetCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCertificateResponseBody) SetCertificateId(v string) *GetCertificateResponseBody {
	s.CertificateId = &v
	return s
}

func (s *GetCertificateResponseBody) SetCertificateChain(v string) *GetCertificateResponseBody {
	s.CertificateChain = &v
	return s
}

func (s *GetCertificateResponseBody) SetCertificate(v string) *GetCertificateResponseBody {
	s.Certificate = &v
	return s
}

type GetCertificateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetCertificateResponse) SetBody(v *GetCertificateResponseBody) *GetCertificateResponse {
	s.Body = v
	return s
}

type GetParametersForImportRequest struct {
	KeyId             *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	WrappingAlgorithm *string `json:"WrappingAlgorithm,omitempty" xml:"WrappingAlgorithm,omitempty"`
	WrappingKeySpec   *string `json:"WrappingKeySpec,omitempty" xml:"WrappingKeySpec,omitempty"`
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
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PublicKey       *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	KeyId           *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	TokenExpireTime *string `json:"TokenExpireTime,omitempty" xml:"TokenExpireTime,omitempty"`
	ImportToken     *string `json:"ImportToken,omitempty" xml:"ImportToken,omitempty"`
}

func (s GetParametersForImportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetParametersForImportResponseBody) GoString() string {
	return s.String()
}

func (s *GetParametersForImportResponseBody) SetRequestId(v string) *GetParametersForImportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetParametersForImportResponseBody) SetPublicKey(v string) *GetParametersForImportResponseBody {
	s.PublicKey = &v
	return s
}

func (s *GetParametersForImportResponseBody) SetKeyId(v string) *GetParametersForImportResponseBody {
	s.KeyId = &v
	return s
}

func (s *GetParametersForImportResponseBody) SetTokenExpireTime(v string) *GetParametersForImportResponseBody {
	s.TokenExpireTime = &v
	return s
}

func (s *GetParametersForImportResponseBody) SetImportToken(v string) *GetParametersForImportResponseBody {
	s.ImportToken = &v
	return s
}

type GetParametersForImportResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetParametersForImportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetParametersForImportResponse) SetBody(v *GetParametersForImportResponseBody) *GetParametersForImportResponse {
	s.Body = v
	return s
}

type GetPublicKeyRequest struct {
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
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
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PublicKey    *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s GetPublicKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPublicKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublicKeyResponseBody) SetRequestId(v string) *GetPublicKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPublicKeyResponseBody) SetPublicKey(v string) *GetPublicKeyResponseBody {
	s.PublicKey = &v
	return s
}

func (s *GetPublicKeyResponseBody) SetKeyId(v string) *GetPublicKeyResponseBody {
	s.KeyId = &v
	return s
}

func (s *GetPublicKeyResponseBody) SetKeyVersionId(v string) *GetPublicKeyResponseBody {
	s.KeyVersionId = &v
	return s
}

type GetPublicKeyResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPublicKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetPublicKeyResponse) SetBody(v *GetPublicKeyResponseBody) *GetPublicKeyResponse {
	s.Body = v
	return s
}

type GetRandomPasswordRequest struct {
	PasswordLength          *string `json:"PasswordLength,omitempty" xml:"PasswordLength,omitempty"`
	ExcludeCharacters       *string `json:"ExcludeCharacters,omitempty" xml:"ExcludeCharacters,omitempty"`
	ExcludeLowercase        *string `json:"ExcludeLowercase,omitempty" xml:"ExcludeLowercase,omitempty"`
	ExcludeUppercase        *string `json:"ExcludeUppercase,omitempty" xml:"ExcludeUppercase,omitempty"`
	ExcludeNumbers          *string `json:"ExcludeNumbers,omitempty" xml:"ExcludeNumbers,omitempty"`
	ExcludePunctuation      *string `json:"ExcludePunctuation,omitempty" xml:"ExcludePunctuation,omitempty"`
	RequireEachIncludedType *string `json:"RequireEachIncludedType,omitempty" xml:"RequireEachIncludedType,omitempty"`
}

func (s GetRandomPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRandomPasswordRequest) GoString() string {
	return s.String()
}

func (s *GetRandomPasswordRequest) SetPasswordLength(v string) *GetRandomPasswordRequest {
	s.PasswordLength = &v
	return s
}

func (s *GetRandomPasswordRequest) SetExcludeCharacters(v string) *GetRandomPasswordRequest {
	s.ExcludeCharacters = &v
	return s
}

func (s *GetRandomPasswordRequest) SetExcludeLowercase(v string) *GetRandomPasswordRequest {
	s.ExcludeLowercase = &v
	return s
}

func (s *GetRandomPasswordRequest) SetExcludeUppercase(v string) *GetRandomPasswordRequest {
	s.ExcludeUppercase = &v
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

func (s *GetRandomPasswordRequest) SetRequireEachIncludedType(v string) *GetRandomPasswordRequest {
	s.RequireEachIncludedType = &v
	return s
}

type GetRandomPasswordResponseBody struct {
	RandomPassword *string `json:"RandomPassword,omitempty" xml:"RandomPassword,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRandomPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetRandomPasswordResponse) SetBody(v *GetRandomPasswordResponseBody) *GetRandomPasswordResponse {
	s.Body = v
	return s
}

type GetSecretValueRequest struct {
	SecretName          *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	VersionStage        *string `json:"VersionStage,omitempty" xml:"VersionStage,omitempty"`
	VersionId           *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	FetchExtendedConfig *bool   `json:"FetchExtendedConfig,omitempty" xml:"FetchExtendedConfig,omitempty"`
}

func (s GetSecretValueRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSecretValueRequest) GoString() string {
	return s.String()
}

func (s *GetSecretValueRequest) SetSecretName(v string) *GetSecretValueRequest {
	s.SecretName = &v
	return s
}

func (s *GetSecretValueRequest) SetVersionStage(v string) *GetSecretValueRequest {
	s.VersionStage = &v
	return s
}

func (s *GetSecretValueRequest) SetVersionId(v string) *GetSecretValueRequest {
	s.VersionId = &v
	return s
}

func (s *GetSecretValueRequest) SetFetchExtendedConfig(v bool) *GetSecretValueRequest {
	s.FetchExtendedConfig = &v
	return s
}

type GetSecretValueResponseBody struct {
	VersionId         *string                                  `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	RotationInterval  *string                                  `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	LastRotationDate  *string                                  `json:"LastRotationDate,omitempty" xml:"LastRotationDate,omitempty"`
	RequestId         *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecretType        *string                                  `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	CreateTime        *string                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SecretDataType    *string                                  `json:"SecretDataType,omitempty" xml:"SecretDataType,omitempty"`
	VersionStages     *GetSecretValueResponseBodyVersionStages `json:"VersionStages,omitempty" xml:"VersionStages,omitempty" type:"Struct"`
	AutomaticRotation *string                                  `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	SecretName        *string                                  `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	NextRotationDate  *string                                  `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	ExtendedConfig    *string                                  `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	SecretData        *string                                  `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
}

func (s GetSecretValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSecretValueResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecretValueResponseBody) SetVersionId(v string) *GetSecretValueResponseBody {
	s.VersionId = &v
	return s
}

func (s *GetSecretValueResponseBody) SetRotationInterval(v string) *GetSecretValueResponseBody {
	s.RotationInterval = &v
	return s
}

func (s *GetSecretValueResponseBody) SetLastRotationDate(v string) *GetSecretValueResponseBody {
	s.LastRotationDate = &v
	return s
}

func (s *GetSecretValueResponseBody) SetRequestId(v string) *GetSecretValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecretValueResponseBody) SetSecretType(v string) *GetSecretValueResponseBody {
	s.SecretType = &v
	return s
}

func (s *GetSecretValueResponseBody) SetCreateTime(v string) *GetSecretValueResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetSecretValueResponseBody) SetSecretDataType(v string) *GetSecretValueResponseBody {
	s.SecretDataType = &v
	return s
}

func (s *GetSecretValueResponseBody) SetVersionStages(v *GetSecretValueResponseBodyVersionStages) *GetSecretValueResponseBody {
	s.VersionStages = v
	return s
}

func (s *GetSecretValueResponseBody) SetAutomaticRotation(v string) *GetSecretValueResponseBody {
	s.AutomaticRotation = &v
	return s
}

func (s *GetSecretValueResponseBody) SetSecretName(v string) *GetSecretValueResponseBody {
	s.SecretName = &v
	return s
}

func (s *GetSecretValueResponseBody) SetNextRotationDate(v string) *GetSecretValueResponseBody {
	s.NextRotationDate = &v
	return s
}

func (s *GetSecretValueResponseBody) SetExtendedConfig(v string) *GetSecretValueResponseBody {
	s.ExtendedConfig = &v
	return s
}

func (s *GetSecretValueResponseBody) SetSecretData(v string) *GetSecretValueResponseBody {
	s.SecretData = &v
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSecretValueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetSecretValueResponse) SetBody(v *GetSecretValueResponseBody) *GetSecretValueResponse {
	s.Body = v
	return s
}

type ImportCertificateRequest struct {
	PKCS12Blob *string `json:"PKCS12Blob,omitempty" xml:"PKCS12Blob,omitempty"`
	Passphrase *string `json:"Passphrase,omitempty" xml:"Passphrase,omitempty"`
}

func (s ImportCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportCertificateRequest) GoString() string {
	return s.String()
}

func (s *ImportCertificateRequest) SetPKCS12Blob(v string) *ImportCertificateRequest {
	s.PKCS12Blob = &v
	return s
}

func (s *ImportCertificateRequest) SetPassphrase(v string) *ImportCertificateRequest {
	s.Passphrase = &v
	return s
}

type ImportCertificateResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	Arn           *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
}

func (s ImportCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ImportCertificateResponseBody) SetRequestId(v string) *ImportCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportCertificateResponseBody) SetCertificateId(v string) *ImportCertificateResponseBody {
	s.CertificateId = &v
	return s
}

func (s *ImportCertificateResponseBody) SetArn(v string) *ImportCertificateResponseBody {
	s.Arn = &v
	return s
}

type ImportCertificateResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImportCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportCertificateResponse) GoString() string {
	return s.String()
}

func (s *ImportCertificateResponse) SetHeaders(v map[string]*string) *ImportCertificateResponse {
	s.Headers = v
	return s
}

func (s *ImportCertificateResponse) SetBody(v *ImportCertificateResponseBody) *ImportCertificateResponse {
	s.Body = v
	return s
}

type ImportEncryptionCertificateRequest struct {
	CertificateId         *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	EncryptedSymmetricKey *string `json:"EncryptedSymmetricKey,omitempty" xml:"EncryptedSymmetricKey,omitempty"`
	EncryptedPrivateKey   *string `json:"EncryptedPrivateKey,omitempty" xml:"EncryptedPrivateKey,omitempty"`
	SymmetricAlgorithm    *string `json:"SymmetricAlgorithm,omitempty" xml:"SymmetricAlgorithm,omitempty"`
	AsymmetricAlgorithm   *string `json:"AsymmetricAlgorithm,omitempty" xml:"AsymmetricAlgorithm,omitempty"`
	Certificate           *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	CertificateChain      *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
}

func (s ImportEncryptionCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportEncryptionCertificateRequest) GoString() string {
	return s.String()
}

func (s *ImportEncryptionCertificateRequest) SetCertificateId(v string) *ImportEncryptionCertificateRequest {
	s.CertificateId = &v
	return s
}

func (s *ImportEncryptionCertificateRequest) SetEncryptedSymmetricKey(v string) *ImportEncryptionCertificateRequest {
	s.EncryptedSymmetricKey = &v
	return s
}

func (s *ImportEncryptionCertificateRequest) SetEncryptedPrivateKey(v string) *ImportEncryptionCertificateRequest {
	s.EncryptedPrivateKey = &v
	return s
}

func (s *ImportEncryptionCertificateRequest) SetSymmetricAlgorithm(v string) *ImportEncryptionCertificateRequest {
	s.SymmetricAlgorithm = &v
	return s
}

func (s *ImportEncryptionCertificateRequest) SetAsymmetricAlgorithm(v string) *ImportEncryptionCertificateRequest {
	s.AsymmetricAlgorithm = &v
	return s
}

func (s *ImportEncryptionCertificateRequest) SetCertificate(v string) *ImportEncryptionCertificateRequest {
	s.Certificate = &v
	return s
}

func (s *ImportEncryptionCertificateRequest) SetCertificateChain(v string) *ImportEncryptionCertificateRequest {
	s.CertificateChain = &v
	return s
}

type ImportEncryptionCertificateResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	Arn           *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
}

func (s ImportEncryptionCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportEncryptionCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ImportEncryptionCertificateResponseBody) SetRequestId(v string) *ImportEncryptionCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportEncryptionCertificateResponseBody) SetCertificateId(v string) *ImportEncryptionCertificateResponseBody {
	s.CertificateId = &v
	return s
}

func (s *ImportEncryptionCertificateResponseBody) SetArn(v string) *ImportEncryptionCertificateResponseBody {
	s.Arn = &v
	return s
}

type ImportEncryptionCertificateResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImportEncryptionCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportEncryptionCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportEncryptionCertificateResponse) GoString() string {
	return s.String()
}

func (s *ImportEncryptionCertificateResponse) SetHeaders(v map[string]*string) *ImportEncryptionCertificateResponse {
	s.Headers = v
	return s
}

func (s *ImportEncryptionCertificateResponse) SetBody(v *ImportEncryptionCertificateResponseBody) *ImportEncryptionCertificateResponse {
	s.Body = v
	return s
}

type ImportKeyMaterialRequest struct {
	KeyId                 *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	EncryptedKeyMaterial  *string `json:"EncryptedKeyMaterial,omitempty" xml:"EncryptedKeyMaterial,omitempty"`
	ImportToken           *string `json:"ImportToken,omitempty" xml:"ImportToken,omitempty"`
	KeyMaterialExpireUnix *int64  `json:"KeyMaterialExpireUnix,omitempty" xml:"KeyMaterialExpireUnix,omitempty"`
}

func (s ImportKeyMaterialRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportKeyMaterialRequest) GoString() string {
	return s.String()
}

func (s *ImportKeyMaterialRequest) SetKeyId(v string) *ImportKeyMaterialRequest {
	s.KeyId = &v
	return s
}

func (s *ImportKeyMaterialRequest) SetEncryptedKeyMaterial(v string) *ImportKeyMaterialRequest {
	s.EncryptedKeyMaterial = &v
	return s
}

func (s *ImportKeyMaterialRequest) SetImportToken(v string) *ImportKeyMaterialRequest {
	s.ImportToken = &v
	return s
}

func (s *ImportKeyMaterialRequest) SetKeyMaterialExpireUnix(v int64) *ImportKeyMaterialRequest {
	s.KeyMaterialExpireUnix = &v
	return s
}

type ImportKeyMaterialResponseBody struct {
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImportKeyMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ImportKeyMaterialResponse) SetBody(v *ImportKeyMaterialResponseBody) *ImportKeyMaterialResponse {
	s.Body = v
	return s
}

type ListAliasesRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	TotalCount *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Aliases    *ListAliasesResponseBodyAliases `json:"Aliases,omitempty" xml:"Aliases,omitempty" type:"Struct"`
}

func (s ListAliasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAliasesResponseBody) SetTotalCount(v int32) *ListAliasesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAliasesResponseBody) SetRequestId(v string) *ListAliasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAliasesResponseBody) SetPageSize(v int32) *ListAliasesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAliasesResponseBody) SetPageNumber(v int32) *ListAliasesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAliasesResponseBody) SetAliases(v *ListAliasesResponseBodyAliases) *ListAliasesResponseBody {
	s.Aliases = v
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
	AliasArn  *string `json:"AliasArn,omitempty" xml:"AliasArn,omitempty"`
	KeyId     *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
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

func (s *ListAliasesResponseBodyAliasesAlias) SetKeyId(v string) *ListAliasesResponseBodyAliasesAlias {
	s.KeyId = &v
	return s
}

func (s *ListAliasesResponseBodyAliasesAlias) SetAliasName(v string) *ListAliasesResponseBodyAliasesAlias {
	s.AliasName = &v
	return s
}

type ListAliasesResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAliasesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListAliasesResponse) SetBody(v *ListAliasesResponseBody) *ListAliasesResponse {
	s.Body = v
	return s
}

type ListAliasesByKeyIdRequest struct {
	KeyId      *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Aliases    *ListAliasesByKeyIdResponseBodyAliases `json:"Aliases,omitempty" xml:"Aliases,omitempty" type:"Struct"`
}

func (s ListAliasesByKeyIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesByKeyIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListAliasesByKeyIdResponseBody) SetTotalCount(v int32) *ListAliasesByKeyIdResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAliasesByKeyIdResponseBody) SetRequestId(v string) *ListAliasesByKeyIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAliasesByKeyIdResponseBody) SetPageSize(v int32) *ListAliasesByKeyIdResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAliasesByKeyIdResponseBody) SetPageNumber(v int32) *ListAliasesByKeyIdResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAliasesByKeyIdResponseBody) SetAliases(v *ListAliasesByKeyIdResponseBodyAliases) *ListAliasesByKeyIdResponseBody {
	s.Aliases = v
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
	AliasArn  *string `json:"AliasArn,omitempty" xml:"AliasArn,omitempty"`
	KeyId     *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
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

func (s *ListAliasesByKeyIdResponseBodyAliasesAlias) SetKeyId(v string) *ListAliasesByKeyIdResponseBodyAliasesAlias {
	s.KeyId = &v
	return s
}

func (s *ListAliasesByKeyIdResponseBodyAliasesAlias) SetAliasName(v string) *ListAliasesByKeyIdResponseBodyAliasesAlias {
	s.AliasName = &v
	return s
}

type ListAliasesByKeyIdResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAliasesByKeyIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListAliasesByKeyIdResponse) SetBody(v *ListAliasesByKeyIdResponseBody) *ListAliasesByKeyIdResponse {
	s.Body = v
	return s
}

type ListCertificatesRequest struct {
	Subject    *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	Issuer     *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCertificatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListCertificatesRequest) SetSubject(v string) *ListCertificatesRequest {
	s.Subject = &v
	return s
}

func (s *ListCertificatesRequest) SetIssuer(v string) *ListCertificatesRequest {
	s.Issuer = &v
	return s
}

func (s *ListCertificatesRequest) SetStatus(v string) *ListCertificatesRequest {
	s.Status = &v
	return s
}

func (s *ListCertificatesRequest) SetPageNumber(v int32) *ListCertificatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCertificatesRequest) SetPageSize(v int32) *ListCertificatesRequest {
	s.PageSize = &v
	return s
}

type ListCertificatesResponseBody struct {
	PageSize               *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId              *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber             *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	CertificateSummaryList *ListCertificatesResponseBodyCertificateSummaryList `json:"CertificateSummaryList,omitempty" xml:"CertificateSummaryList,omitempty" type:"Struct"`
	TotalSize              *int32                                              `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListCertificatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCertificatesResponseBody) SetPageSize(v int32) *ListCertificatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCertificatesResponseBody) SetRequestId(v string) *ListCertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCertificatesResponseBody) SetPageNumber(v int32) *ListCertificatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCertificatesResponseBody) SetCertificateSummaryList(v *ListCertificatesResponseBodyCertificateSummaryList) *ListCertificatesResponseBody {
	s.CertificateSummaryList = v
	return s
}

func (s *ListCertificatesResponseBody) SetTotalSize(v int32) *ListCertificatesResponseBody {
	s.TotalSize = &v
	return s
}

type ListCertificatesResponseBodyCertificateSummaryList struct {
	CertificateSummary []*ListCertificatesResponseBodyCertificateSummaryListCertificateSummary `json:"CertificateSummary,omitempty" xml:"CertificateSummary,omitempty" type:"Repeated"`
}

func (s ListCertificatesResponseBodyCertificateSummaryList) String() string {
	return tea.Prettify(s)
}

func (s ListCertificatesResponseBodyCertificateSummaryList) GoString() string {
	return s.String()
}

func (s *ListCertificatesResponseBodyCertificateSummaryList) SetCertificateSummary(v []*ListCertificatesResponseBodyCertificateSummaryListCertificateSummary) *ListCertificatesResponseBodyCertificateSummaryList {
	s.CertificateSummary = v
	return s
}

type ListCertificatesResponseBodyCertificateSummaryListCertificateSummary struct {
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	KeySpec         *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	Subject         *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	ProtectionLevel *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	Issuer          *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	NotBefore       *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	NotAfter        *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	CertificateId   *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s ListCertificatesResponseBodyCertificateSummaryListCertificateSummary) String() string {
	return tea.Prettify(s)
}

func (s ListCertificatesResponseBodyCertificateSummaryListCertificateSummary) GoString() string {
	return s.String()
}

func (s *ListCertificatesResponseBodyCertificateSummaryListCertificateSummary) SetStatus(v string) *ListCertificatesResponseBodyCertificateSummaryListCertificateSummary {
	s.Status = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateSummaryListCertificateSummary) SetKeySpec(v string) *ListCertificatesResponseBodyCertificateSummaryListCertificateSummary {
	s.KeySpec = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateSummaryListCertificateSummary) SetSubject(v string) *ListCertificatesResponseBodyCertificateSummaryListCertificateSummary {
	s.Subject = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateSummaryListCertificateSummary) SetProtectionLevel(v string) *ListCertificatesResponseBodyCertificateSummaryListCertificateSummary {
	s.ProtectionLevel = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateSummaryListCertificateSummary) SetIssuer(v string) *ListCertificatesResponseBodyCertificateSummaryListCertificateSummary {
	s.Issuer = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateSummaryListCertificateSummary) SetNotBefore(v string) *ListCertificatesResponseBodyCertificateSummaryListCertificateSummary {
	s.NotBefore = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateSummaryListCertificateSummary) SetNotAfter(v string) *ListCertificatesResponseBodyCertificateSummaryListCertificateSummary {
	s.NotAfter = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateSummaryListCertificateSummary) SetCertificateId(v string) *ListCertificatesResponseBodyCertificateSummaryListCertificateSummary {
	s.CertificateId = &v
	return s
}

type ListCertificatesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCertificatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCertificatesResponse) GoString() string {
	return s.String()
}

func (s *ListCertificatesResponse) SetHeaders(v map[string]*string) *ListCertificatesResponse {
	s.Headers = v
	return s
}

func (s *ListCertificatesResponse) SetBody(v *ListCertificatesResponseBody) *ListCertificatesResponse {
	s.Body = v
	return s
}

type ListKeysRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Filters    *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
}

func (s ListKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListKeysRequest) GoString() string {
	return s.String()
}

func (s *ListKeysRequest) SetPageNumber(v int32) *ListKeysRequest {
	s.PageNumber = &v
	return s
}

func (s *ListKeysRequest) SetPageSize(v int32) *ListKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListKeysRequest) SetFilters(v string) *ListKeysRequest {
	s.Filters = &v
	return s
}

type ListKeysResponseBody struct {
	TotalCount *int32                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Keys       *ListKeysResponseBodyKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Struct"`
}

func (s ListKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListKeysResponseBody) SetTotalCount(v int32) *ListKeysResponseBody {
	s.TotalCount = &v
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

func (s *ListKeysResponseBody) SetPageNumber(v int32) *ListKeysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListKeysResponseBody) SetKeys(v *ListKeysResponseBodyKeys) *ListKeysResponseBody {
	s.Keys = v
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
	KeyArn *string `json:"KeyArn,omitempty" xml:"KeyArn,omitempty"`
	KeyId  *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
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
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListKeysResponse) SetBody(v *ListKeysResponseBody) *ListKeysResponse {
	s.Body = v
	return s
}

type ListKeyVersionsRequest struct {
	KeyId      *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	KeyVersions *ListKeyVersionsResponseBodyKeyVersions `json:"KeyVersions,omitempty" xml:"KeyVersions,omitempty" type:"Struct"`
	TotalCount  *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber  *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
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

func (s *ListKeyVersionsResponseBody) SetTotalCount(v int32) *ListKeyVersionsResponseBody {
	s.TotalCount = &v
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

func (s *ListKeyVersionsResponseBody) SetPageNumber(v int32) *ListKeyVersionsResponseBody {
	s.PageNumber = &v
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
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
}

func (s ListKeyVersionsResponseBodyKeyVersionsKeyVersion) String() string {
	return tea.Prettify(s)
}

func (s ListKeyVersionsResponseBodyKeyVersionsKeyVersion) GoString() string {
	return s.String()
}

func (s *ListKeyVersionsResponseBodyKeyVersionsKeyVersion) SetKeyVersionId(v string) *ListKeyVersionsResponseBodyKeyVersionsKeyVersion {
	s.KeyVersionId = &v
	return s
}

func (s *ListKeyVersionsResponseBodyKeyVersionsKeyVersion) SetKeyId(v string) *ListKeyVersionsResponseBodyKeyVersionsKeyVersion {
	s.KeyId = &v
	return s
}

func (s *ListKeyVersionsResponseBodyKeyVersionsKeyVersion) SetCreationDate(v string) *ListKeyVersionsResponseBodyKeyVersionsKeyVersion {
	s.CreationDate = &v
	return s
}

type ListKeyVersionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListKeyVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListKeyVersionsResponse) SetBody(v *ListKeyVersionsResponseBody) *ListKeyVersionsResponse {
	s.Body = v
	return s
}

type ListResourceTagsRequest struct {
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
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      *ListResourceTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
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
	KeyId    *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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

func (s *ListResourceTagsResponseBodyTagsTag) SetTagValue(v string) *ListResourceTagsResponseBodyTagsTag {
	s.TagValue = &v
	return s
}

func (s *ListResourceTagsResponseBodyTagsTag) SetTagKey(v string) *ListResourceTagsResponseBodyTagsTag {
	s.TagKey = &v
	return s
}

type ListResourceTagsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListResourceTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListResourceTagsResponse) SetBody(v *ListResourceTagsResponseBody) *ListResourceTagsResponse {
	s.Body = v
	return s
}

type ListSecretsRequest struct {
	FetchTags  *string `json:"FetchTags,omitempty" xml:"FetchTags,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Filters    *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
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

func (s *ListSecretsRequest) SetPageNumber(v int32) *ListSecretsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSecretsRequest) SetPageSize(v int32) *ListSecretsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSecretsRequest) SetFilters(v string) *ListSecretsRequest {
	s.Filters = &v
	return s
}

type ListSecretsResponseBody struct {
	TotalCount *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecretList *ListSecretsResponseBodySecretList `json:"SecretList,omitempty" xml:"SecretList,omitempty" type:"Struct"`
	PageNumber *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s ListSecretsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSecretsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBody) SetTotalCount(v int32) *ListSecretsResponseBody {
	s.TotalCount = &v
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

func (s *ListSecretsResponseBody) SetPageNumber(v int32) *ListSecretsResponseBody {
	s.PageNumber = &v
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
	UpdateTime        *string                                      `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	CreateTime        *string                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SecretName        *string                                      `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	Tags              *ListSecretsResponseBodySecretListSecretTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	SecretType        *string                                      `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	PlannedDeleteTime *string                                      `json:"PlannedDeleteTime,omitempty" xml:"PlannedDeleteTime,omitempty"`
}

func (s ListSecretsResponseBodySecretListSecret) String() string {
	return tea.Prettify(s)
}

func (s ListSecretsResponseBodySecretListSecret) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBodySecretListSecret) SetUpdateTime(v string) *ListSecretsResponseBodySecretListSecret {
	s.UpdateTime = &v
	return s
}

func (s *ListSecretsResponseBodySecretListSecret) SetCreateTime(v string) *ListSecretsResponseBodySecretListSecret {
	s.CreateTime = &v
	return s
}

func (s *ListSecretsResponseBodySecretListSecret) SetSecretName(v string) *ListSecretsResponseBodySecretListSecret {
	s.SecretName = &v
	return s
}

func (s *ListSecretsResponseBodySecretListSecret) SetTags(v *ListSecretsResponseBodySecretListSecretTags) *ListSecretsResponseBodySecretListSecret {
	s.Tags = v
	return s
}

func (s *ListSecretsResponseBodySecretListSecret) SetSecretType(v string) *ListSecretsResponseBodySecretListSecret {
	s.SecretType = &v
	return s
}

func (s *ListSecretsResponseBodySecretListSecret) SetPlannedDeleteTime(v string) *ListSecretsResponseBodySecretListSecret {
	s.PlannedDeleteTime = &v
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
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListSecretsResponseBodySecretListSecretTagsTag) String() string {
	return tea.Prettify(s)
}

func (s ListSecretsResponseBodySecretListSecretTagsTag) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBodySecretListSecretTagsTag) SetTagValue(v string) *ListSecretsResponseBodySecretListSecretTagsTag {
	s.TagValue = &v
	return s
}

func (s *ListSecretsResponseBodySecretListSecretTagsTag) SetTagKey(v string) *ListSecretsResponseBodySecretListSecretTagsTag {
	s.TagKey = &v
	return s
}

type ListSecretsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSecretsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListSecretsResponse) SetBody(v *ListSecretsResponseBody) *ListSecretsResponse {
	s.Body = v
	return s
}

type ListSecretVersionIdsRequest struct {
	SecretName        *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	IncludeDeprecated *string `json:"IncludeDeprecated,omitempty" xml:"IncludeDeprecated,omitempty"`
	PageNumber        *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSecretVersionIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSecretVersionIdsRequest) GoString() string {
	return s.String()
}

func (s *ListSecretVersionIdsRequest) SetSecretName(v string) *ListSecretVersionIdsRequest {
	s.SecretName = &v
	return s
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

type ListSecretVersionIdsResponseBody struct {
	SecretName *string                                     `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	TotalCount *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	VersionIds *ListSecretVersionIdsResponseBodyVersionIds `json:"VersionIds,omitempty" xml:"VersionIds,omitempty" type:"Struct"`
}

func (s ListSecretVersionIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSecretVersionIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecretVersionIdsResponseBody) SetSecretName(v string) *ListSecretVersionIdsResponseBody {
	s.SecretName = &v
	return s
}

func (s *ListSecretVersionIdsResponseBody) SetTotalCount(v int32) *ListSecretVersionIdsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSecretVersionIdsResponseBody) SetRequestId(v string) *ListSecretVersionIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecretVersionIdsResponseBody) SetPageSize(v int32) *ListSecretVersionIdsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSecretVersionIdsResponseBody) SetPageNumber(v int32) *ListSecretVersionIdsResponseBody {
	s.PageNumber = &v
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
	CreateTime    *string                                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	VersionId     *string                                                           `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSecretVersionIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListSecretVersionIdsResponse) SetBody(v *ListSecretVersionIdsResponseBody) *ListSecretVersionIdsResponse {
	s.Body = v
	return s
}

type OpenKmsServiceResponseBody struct {
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenKmsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OpenKmsServiceResponse) SetBody(v *OpenKmsServiceResponseBody) *OpenKmsServiceResponse {
	s.Body = v
	return s
}

type PutSecretValueRequest struct {
	VersionId      *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	SecretName     *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	SecretData     *string `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
	SecretDataType *string `json:"SecretDataType,omitempty" xml:"SecretDataType,omitempty"`
	VersionStages  *string `json:"VersionStages,omitempty" xml:"VersionStages,omitempty"`
}

func (s PutSecretValueRequest) String() string {
	return tea.Prettify(s)
}

func (s PutSecretValueRequest) GoString() string {
	return s.String()
}

func (s *PutSecretValueRequest) SetVersionId(v string) *PutSecretValueRequest {
	s.VersionId = &v
	return s
}

func (s *PutSecretValueRequest) SetSecretName(v string) *PutSecretValueRequest {
	s.SecretName = &v
	return s
}

func (s *PutSecretValueRequest) SetSecretData(v string) *PutSecretValueRequest {
	s.SecretData = &v
	return s
}

func (s *PutSecretValueRequest) SetSecretDataType(v string) *PutSecretValueRequest {
	s.SecretDataType = &v
	return s
}

func (s *PutSecretValueRequest) SetVersionStages(v string) *PutSecretValueRequest {
	s.VersionStages = &v
	return s
}

type PutSecretValueResponseBody struct {
	SecretName    *string                                  `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	VersionId     *string                                  `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	RequestId     *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VersionStages *PutSecretValueResponseBodyVersionStages `json:"VersionStages,omitempty" xml:"VersionStages,omitempty" type:"Struct"`
}

func (s PutSecretValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutSecretValueResponseBody) GoString() string {
	return s.String()
}

func (s *PutSecretValueResponseBody) SetSecretName(v string) *PutSecretValueResponseBody {
	s.SecretName = &v
	return s
}

func (s *PutSecretValueResponseBody) SetVersionId(v string) *PutSecretValueResponseBody {
	s.VersionId = &v
	return s
}

func (s *PutSecretValueResponseBody) SetRequestId(v string) *PutSecretValueResponseBody {
	s.RequestId = &v
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutSecretValueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PutSecretValueResponse) SetBody(v *PutSecretValueResponseBody) *PutSecretValueResponse {
	s.Body = v
	return s
}

type ReEncryptRequest struct {
	CiphertextBlob               *string                `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	SourceKeyId                  *string                `json:"SourceKeyId,omitempty" xml:"SourceKeyId,omitempty"`
	SourceKeyVersionId           *string                `json:"SourceKeyVersionId,omitempty" xml:"SourceKeyVersionId,omitempty"`
	SourceEncryptionAlgorithm    *string                `json:"SourceEncryptionAlgorithm,omitempty" xml:"SourceEncryptionAlgorithm,omitempty"`
	SourceEncryptionContext      map[string]interface{} `json:"SourceEncryptionContext,omitempty" xml:"SourceEncryptionContext,omitempty"`
	DestinationKeyId             *string                `json:"DestinationKeyId,omitempty" xml:"DestinationKeyId,omitempty"`
	DestinationEncryptionContext map[string]interface{} `json:"DestinationEncryptionContext,omitempty" xml:"DestinationEncryptionContext,omitempty"`
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

func (s *ReEncryptRequest) SetSourceKeyId(v string) *ReEncryptRequest {
	s.SourceKeyId = &v
	return s
}

func (s *ReEncryptRequest) SetSourceKeyVersionId(v string) *ReEncryptRequest {
	s.SourceKeyVersionId = &v
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

func (s *ReEncryptRequest) SetDestinationKeyId(v string) *ReEncryptRequest {
	s.DestinationKeyId = &v
	return s
}

func (s *ReEncryptRequest) SetDestinationEncryptionContext(v map[string]interface{}) *ReEncryptRequest {
	s.DestinationEncryptionContext = v
	return s
}

type ReEncryptShrinkRequest struct {
	CiphertextBlob                     *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	SourceKeyId                        *string `json:"SourceKeyId,omitempty" xml:"SourceKeyId,omitempty"`
	SourceKeyVersionId                 *string `json:"SourceKeyVersionId,omitempty" xml:"SourceKeyVersionId,omitempty"`
	SourceEncryptionAlgorithm          *string `json:"SourceEncryptionAlgorithm,omitempty" xml:"SourceEncryptionAlgorithm,omitempty"`
	SourceEncryptionContextShrink      *string `json:"SourceEncryptionContext,omitempty" xml:"SourceEncryptionContext,omitempty"`
	DestinationKeyId                   *string `json:"DestinationKeyId,omitempty" xml:"DestinationKeyId,omitempty"`
	DestinationEncryptionContextShrink *string `json:"DestinationEncryptionContext,omitempty" xml:"DestinationEncryptionContext,omitempty"`
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

func (s *ReEncryptShrinkRequest) SetSourceKeyId(v string) *ReEncryptShrinkRequest {
	s.SourceKeyId = &v
	return s
}

func (s *ReEncryptShrinkRequest) SetSourceKeyVersionId(v string) *ReEncryptShrinkRequest {
	s.SourceKeyVersionId = &v
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

func (s *ReEncryptShrinkRequest) SetDestinationKeyId(v string) *ReEncryptShrinkRequest {
	s.DestinationKeyId = &v
	return s
}

func (s *ReEncryptShrinkRequest) SetDestinationEncryptionContextShrink(v string) *ReEncryptShrinkRequest {
	s.DestinationEncryptionContextShrink = &v
	return s
}

type ReEncryptResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	KeyId          *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeyVersionId   *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s ReEncryptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReEncryptResponseBody) GoString() string {
	return s.String()
}

func (s *ReEncryptResponseBody) SetRequestId(v string) *ReEncryptResponseBody {
	s.RequestId = &v
	return s
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

type ReEncryptResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReEncryptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ReEncryptResponse) SetBody(v *ReEncryptResponseBody) *ReEncryptResponse {
	s.Body = v
	return s
}

type RestoreSecretRequest struct {
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
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestoreSecretResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreSecretResponseBody) SetSecretName(v string) *RestoreSecretResponseBody {
	s.SecretName = &v
	return s
}

func (s *RestoreSecretResponseBody) SetRequestId(v string) *RestoreSecretResponseBody {
	s.RequestId = &v
	return s
}

type RestoreSecretResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestoreSecretResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RestoreSecretResponse) SetBody(v *RestoreSecretResponseBody) *RestoreSecretResponse {
	s.Body = v
	return s
}

type RotateSecretRequest struct {
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	VersionId  *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
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
	VersionId  *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Arn        *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
}

func (s RotateSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RotateSecretResponseBody) GoString() string {
	return s.String()
}

func (s *RotateSecretResponseBody) SetVersionId(v string) *RotateSecretResponseBody {
	s.VersionId = &v
	return s
}

func (s *RotateSecretResponseBody) SetSecretName(v string) *RotateSecretResponseBody {
	s.SecretName = &v
	return s
}

func (s *RotateSecretResponseBody) SetRequestId(v string) *RotateSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *RotateSecretResponseBody) SetArn(v string) *RotateSecretResponseBody {
	s.Arn = &v
	return s
}

type RotateSecretResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RotateSecretResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RotateSecretResponse) SetBody(v *RotateSecretResponseBody) *RotateSecretResponse {
	s.Body = v
	return s
}

type ScheduleKeyDeletionRequest struct {
	KeyId               *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	PendingWindowInDays *int32  `json:"PendingWindowInDays,omitempty" xml:"PendingWindowInDays,omitempty"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ScheduleKeyDeletionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ScheduleKeyDeletionResponse) SetBody(v *ScheduleKeyDeletionResponseBody) *ScheduleKeyDeletionResponse {
	s.Body = v
	return s
}

type TagResourceRequest struct {
	KeyId         *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Tags          *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	SecretName    *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s TagResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourceRequest) GoString() string {
	return s.String()
}

func (s *TagResourceRequest) SetKeyId(v string) *TagResourceRequest {
	s.KeyId = &v
	return s
}

func (s *TagResourceRequest) SetTags(v string) *TagResourceRequest {
	s.Tags = &v
	return s
}

func (s *TagResourceRequest) SetSecretName(v string) *TagResourceRequest {
	s.SecretName = &v
	return s
}

func (s *TagResourceRequest) SetCertificateId(v string) *TagResourceRequest {
	s.CertificateId = &v
	return s
}

type TagResourceResponseBody struct {
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TagResourceResponse) SetBody(v *TagResourceResponseBody) *TagResourceResponse {
	s.Body = v
	return s
}

type UntagResourceRequest struct {
	KeyId         *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	TagKeys       *string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
	SecretName    *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s UntagResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourceRequest) GoString() string {
	return s.String()
}

func (s *UntagResourceRequest) SetKeyId(v string) *UntagResourceRequest {
	s.KeyId = &v
	return s
}

func (s *UntagResourceRequest) SetTagKeys(v string) *UntagResourceRequest {
	s.TagKeys = &v
	return s
}

func (s *UntagResourceRequest) SetSecretName(v string) *UntagResourceRequest {
	s.SecretName = &v
	return s
}

func (s *UntagResourceRequest) SetCertificateId(v string) *UntagResourceRequest {
	s.CertificateId = &v
	return s
}

type UntagResourceResponseBody struct {
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UntagResourceResponse) SetBody(v *UntagResourceResponseBody) *UntagResourceResponse {
	s.Body = v
	return s
}

type UpdateAliasRequest struct {
	KeyId     *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
}

func (s UpdateAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAliasRequest) GoString() string {
	return s.String()
}

func (s *UpdateAliasRequest) SetKeyId(v string) *UpdateAliasRequest {
	s.KeyId = &v
	return s
}

func (s *UpdateAliasRequest) SetAliasName(v string) *UpdateAliasRequest {
	s.AliasName = &v
	return s
}

type UpdateAliasResponseBody struct {
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAliasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateAliasResponse) SetBody(v *UpdateAliasResponseBody) *UpdateAliasResponse {
	s.Body = v
	return s
}

type UpdateCertificateStatusRequest struct {
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCertificateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateCertificateStatusResponse) SetBody(v *UpdateCertificateStatusResponseBody) *UpdateCertificateStatusResponse {
	s.Body = v
	return s
}

type UpdateKeyDescriptionRequest struct {
	KeyId       *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s UpdateKeyDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateKeyDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateKeyDescriptionRequest) SetKeyId(v string) *UpdateKeyDescriptionRequest {
	s.KeyId = &v
	return s
}

func (s *UpdateKeyDescriptionRequest) SetDescription(v string) *UpdateKeyDescriptionRequest {
	s.Description = &v
	return s
}

type UpdateKeyDescriptionResponseBody struct {
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateKeyDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateKeyDescriptionResponse) SetBody(v *UpdateKeyDescriptionResponseBody) *UpdateKeyDescriptionResponse {
	s.Body = v
	return s
}

type UpdateRotationPolicyRequest struct {
	KeyId                   *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	EnableAutomaticRotation *bool   `json:"EnableAutomaticRotation,omitempty" xml:"EnableAutomaticRotation,omitempty"`
	RotationInterval        *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
}

func (s UpdateRotationPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRotationPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateRotationPolicyRequest) SetKeyId(v string) *UpdateRotationPolicyRequest {
	s.KeyId = &v
	return s
}

func (s *UpdateRotationPolicyRequest) SetEnableAutomaticRotation(v bool) *UpdateRotationPolicyRequest {
	s.EnableAutomaticRotation = &v
	return s
}

func (s *UpdateRotationPolicyRequest) SetRotationInterval(v string) *UpdateRotationPolicyRequest {
	s.RotationInterval = &v
	return s
}

type UpdateRotationPolicyResponseBody struct {
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateRotationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateRotationPolicyResponse) SetBody(v *UpdateRotationPolicyResponseBody) *UpdateRotationPolicyResponse {
	s.Body = v
	return s
}

type UpdateSecretRotationPolicyRequest struct {
	SecretName              *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	EnableAutomaticRotation *bool   `json:"EnableAutomaticRotation,omitempty" xml:"EnableAutomaticRotation,omitempty"`
	RotationInterval        *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
}

func (s UpdateSecretRotationPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretRotationPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecretRotationPolicyRequest) SetSecretName(v string) *UpdateSecretRotationPolicyRequest {
	s.SecretName = &v
	return s
}

func (s *UpdateSecretRotationPolicyRequest) SetEnableAutomaticRotation(v bool) *UpdateSecretRotationPolicyRequest {
	s.EnableAutomaticRotation = &v
	return s
}

func (s *UpdateSecretRotationPolicyRequest) SetRotationInterval(v string) *UpdateSecretRotationPolicyRequest {
	s.RotationInterval = &v
	return s
}

type UpdateSecretRotationPolicyResponseBody struct {
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSecretRotationPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretRotationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecretRotationPolicyResponseBody) SetSecretName(v string) *UpdateSecretRotationPolicyResponseBody {
	s.SecretName = &v
	return s
}

func (s *UpdateSecretRotationPolicyResponseBody) SetRequestId(v string) *UpdateSecretRotationPolicyResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSecretRotationPolicyResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSecretRotationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateSecretRotationPolicyResponse) SetBody(v *UpdateSecretRotationPolicyResponseBody) *UpdateSecretRotationPolicyResponse {
	s.Body = v
	return s
}

type UpdateSecretVersionStageRequest struct {
	SecretName        *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	VersionStage      *string `json:"VersionStage,omitempty" xml:"VersionStage,omitempty"`
	RemoveFromVersion *string `json:"RemoveFromVersion,omitempty" xml:"RemoveFromVersion,omitempty"`
	MoveToVersion     *string `json:"MoveToVersion,omitempty" xml:"MoveToVersion,omitempty"`
}

func (s UpdateSecretVersionStageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretVersionStageRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecretVersionStageRequest) SetSecretName(v string) *UpdateSecretVersionStageRequest {
	s.SecretName = &v
	return s
}

func (s *UpdateSecretVersionStageRequest) SetVersionStage(v string) *UpdateSecretVersionStageRequest {
	s.VersionStage = &v
	return s
}

func (s *UpdateSecretVersionStageRequest) SetRemoveFromVersion(v string) *UpdateSecretVersionStageRequest {
	s.RemoveFromVersion = &v
	return s
}

func (s *UpdateSecretVersionStageRequest) SetMoveToVersion(v string) *UpdateSecretVersionStageRequest {
	s.MoveToVersion = &v
	return s
}

type UpdateSecretVersionStageResponseBody struct {
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSecretVersionStageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretVersionStageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecretVersionStageResponseBody) SetSecretName(v string) *UpdateSecretVersionStageResponseBody {
	s.SecretName = &v
	return s
}

func (s *UpdateSecretVersionStageResponseBody) SetRequestId(v string) *UpdateSecretVersionStageResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSecretVersionStageResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSecretVersionStageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateSecretVersionStageResponse) SetBody(v *UpdateSecretVersionStageResponseBody) *UpdateSecretVersionStageResponse {
	s.Body = v
	return s
}

type UploadCertificateRequest struct {
	CertificateId    *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	Certificate      *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
}

func (s UploadCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadCertificateRequest) GoString() string {
	return s.String()
}

func (s *UploadCertificateRequest) SetCertificateId(v string) *UploadCertificateRequest {
	s.CertificateId = &v
	return s
}

func (s *UploadCertificateRequest) SetCertificate(v string) *UploadCertificateRequest {
	s.Certificate = &v
	return s
}

func (s *UploadCertificateRequest) SetCertificateChain(v string) *UploadCertificateRequest {
	s.CertificateChain = &v
	return s
}

type UploadCertificateResponseBody struct {
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (client *Client) AsymmetricDecryptWithOptions(request *AsymmetricDecryptRequest, runtime *util.RuntimeOptions) (_result *AsymmetricDecryptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AsymmetricDecryptResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AsymmetricDecrypt"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) AsymmetricEncryptWithOptions(request *AsymmetricEncryptRequest, runtime *util.RuntimeOptions) (_result *AsymmetricEncryptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AsymmetricEncryptResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AsymmetricEncrypt"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) AsymmetricSignWithOptions(request *AsymmetricSignRequest, runtime *util.RuntimeOptions) (_result *AsymmetricSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AsymmetricSignResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AsymmetricSign"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) AsymmetricVerifyWithOptions(request *AsymmetricVerifyRequest, runtime *util.RuntimeOptions) (_result *AsymmetricVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AsymmetricVerifyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AsymmetricVerify"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CancelKeyDeletionWithOptions(request *CancelKeyDeletionRequest, runtime *util.RuntimeOptions) (_result *CancelKeyDeletionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelKeyDeletionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelKeyDeletion"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CertificatePrivateKeyDecryptWithOptions(request *CertificatePrivateKeyDecryptRequest, runtime *util.RuntimeOptions) (_result *CertificatePrivateKeyDecryptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CertificatePrivateKeyDecryptResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CertificatePrivateKeyDecrypt"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CertificatePrivateKeySignWithOptions(request *CertificatePrivateKeySignRequest, runtime *util.RuntimeOptions) (_result *CertificatePrivateKeySignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CertificatePrivateKeySignResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CertificatePrivateKeySign"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CertificatePublicKeyEncryptWithOptions(request *CertificatePublicKeyEncryptRequest, runtime *util.RuntimeOptions) (_result *CertificatePublicKeyEncryptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CertificatePublicKeyEncryptResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CertificatePublicKeyEncrypt"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CertificatePublicKeyVerifyWithOptions(request *CertificatePublicKeyVerifyRequest, runtime *util.RuntimeOptions) (_result *CertificatePublicKeyVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CertificatePublicKeyVerifyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CertificatePublicKeyVerify"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CreateAliasWithOptions(request *CreateAliasRequest, runtime *util.RuntimeOptions) (_result *CreateAliasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAliasResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAlias"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCertificateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCertificate"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CreateKeyWithOptions(request *CreateKeyRequest, runtime *util.RuntimeOptions) (_result *CreateKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateKeyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateKey"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CreateKeyVersionWithOptions(request *CreateKeyVersionRequest, runtime *util.RuntimeOptions) (_result *CreateKeyVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateKeyVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateKeyVersion"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSecretResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSecret"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DecryptResponse{}
	_body, _err := client.DoRPCRequest(tea.String("Decrypt"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAliasResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAlias"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteCertificateWithOptions(request *DeleteCertificateRequest, runtime *util.RuntimeOptions) (_result *DeleteCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCertificateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCertificate"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DeleteKeyMaterialWithOptions(request *DeleteKeyMaterialRequest, runtime *util.RuntimeOptions) (_result *DeleteKeyMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteKeyMaterialResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteKeyMaterial"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DeleteSecretWithOptions(request *DeleteSecretRequest, runtime *util.RuntimeOptions) (_result *DeleteSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSecretResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSecret"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &DescribeAccountKmsStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAccountKmsStatus"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeCertificateWithOptions(request *DescribeCertificateRequest, runtime *util.RuntimeOptions) (_result *DescribeCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCertificateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCertificate"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeKeyWithOptions(request *DescribeKeyRequest, runtime *util.RuntimeOptions) (_result *DescribeKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeKeyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeKey"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeKeyVersionWithOptions(request *DescribeKeyVersionRequest, runtime *util.RuntimeOptions) (_result *DescribeKeyVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeKeyVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeKeyVersion"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeSecretWithOptions(request *DescribeSecretRequest, runtime *util.RuntimeOptions) (_result *DescribeSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSecretResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSecret"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeServiceWithOptions(runtime *util.RuntimeOptions) (_result *DescribeServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeService"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeService() (_result *DescribeServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServiceResponse{}
	_body, _err := client.DescribeServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableKeyWithOptions(request *DisableKeyRequest, runtime *util.RuntimeOptions) (_result *DisableKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableKeyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableKey"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableKeyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableKey"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EncryptResponse{}
	_body, _err := client.DoRPCRequest(tea.String("Encrypt"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ExportCertificateWithOptions(request *ExportCertificateRequest, runtime *util.RuntimeOptions) (_result *ExportCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExportCertificateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExportCertificate"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExportCertificate(request *ExportCertificateRequest) (_result *ExportCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportCertificateResponse{}
	_body, _err := client.ExportCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExportDataKeyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExportDataKey"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateAndExportDataKeyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateAndExportDataKey"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateDataKeyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateDataKey"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateDataKeyWithoutPlaintextResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateDataKeyWithoutPlaintext"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) GetCertificateWithOptions(request *GetCertificateRequest, runtime *util.RuntimeOptions) (_result *GetCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetCertificateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCertificate"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) GetParametersForImportWithOptions(request *GetParametersForImportRequest, runtime *util.RuntimeOptions) (_result *GetParametersForImportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetParametersForImportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetParametersForImport"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPublicKeyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPublicKey"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetRandomPasswordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetRandomPassword"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetSecretValueWithOptions(request *GetSecretValueRequest, runtime *util.RuntimeOptions) (_result *GetSecretValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSecretValueResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSecretValue"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ImportCertificateWithOptions(request *ImportCertificateRequest, runtime *util.RuntimeOptions) (_result *ImportCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ImportCertificateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ImportCertificate"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportCertificate(request *ImportCertificateRequest) (_result *ImportCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportCertificateResponse{}
	_body, _err := client.ImportCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportEncryptionCertificateWithOptions(request *ImportEncryptionCertificateRequest, runtime *util.RuntimeOptions) (_result *ImportEncryptionCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ImportEncryptionCertificateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ImportEncryptionCertificate"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportEncryptionCertificate(request *ImportEncryptionCertificateRequest) (_result *ImportEncryptionCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportEncryptionCertificateResponse{}
	_body, _err := client.ImportEncryptionCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportKeyMaterialWithOptions(request *ImportKeyMaterialRequest, runtime *util.RuntimeOptions) (_result *ImportKeyMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ImportKeyMaterialResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ImportKeyMaterial"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAliasesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAliases"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAliasesByKeyIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAliasesByKeyId"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListCertificatesWithOptions(request *ListCertificatesRequest, runtime *util.RuntimeOptions) (_result *ListCertificatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListCertificatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCertificates"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCertificates(request *ListCertificatesRequest) (_result *ListCertificatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCertificatesResponse{}
	_body, _err := client.ListCertificatesWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListKeys"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListKeyVersionsWithOptions(request *ListKeyVersionsRequest, runtime *util.RuntimeOptions) (_result *ListKeyVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListKeyVersionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListKeyVersions"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListResourceTagsWithOptions(request *ListResourceTagsRequest, runtime *util.RuntimeOptions) (_result *ListResourceTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListResourceTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListResourceTags"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ListSecretsWithOptions(request *ListSecretsRequest, runtime *util.RuntimeOptions) (_result *ListSecretsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSecretsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSecrets"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ListSecretVersionIdsWithOptions(request *ListSecretVersionIdsRequest, runtime *util.RuntimeOptions) (_result *ListSecretVersionIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSecretVersionIdsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSecretVersionIds"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) OpenKmsServiceWithOptions(runtime *util.RuntimeOptions) (_result *OpenKmsServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &OpenKmsServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OpenKmsService"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) PutSecretValueWithOptions(request *PutSecretValueRequest, runtime *util.RuntimeOptions) (_result *PutSecretValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutSecretValueResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutSecretValue"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ReEncryptWithOptions(tmpReq *ReEncryptRequest, runtime *util.RuntimeOptions) (_result *ReEncryptResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ReEncryptShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SourceEncryptionContext)) {
		request.SourceEncryptionContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceEncryptionContext, tea.String("SourceEncryptionContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DestinationEncryptionContext)) {
		request.DestinationEncryptionContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DestinationEncryptionContext, tea.String("DestinationEncryptionContext"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReEncryptResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReEncrypt"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) RestoreSecretWithOptions(request *RestoreSecretRequest, runtime *util.RuntimeOptions) (_result *RestoreSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RestoreSecretResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RestoreSecret"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) RotateSecretWithOptions(request *RotateSecretRequest, runtime *util.RuntimeOptions) (_result *RotateSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RotateSecretResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RotateSecret"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ScheduleKeyDeletionWithOptions(request *ScheduleKeyDeletionRequest, runtime *util.RuntimeOptions) (_result *ScheduleKeyDeletionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ScheduleKeyDeletionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ScheduleKeyDeletion"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) TagResourceWithOptions(request *TagResourceRequest, runtime *util.RuntimeOptions) (_result *TagResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResource"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) UntagResourceWithOptions(request *UntagResourceRequest, runtime *util.RuntimeOptions) (_result *UntagResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResource"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAliasResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAlias"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpdateCertificateStatusWithOptions(request *UpdateCertificateStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateCertificateStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCertificateStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCertificateStatus"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) UpdateKeyDescriptionWithOptions(request *UpdateKeyDescriptionRequest, runtime *util.RuntimeOptions) (_result *UpdateKeyDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateKeyDescriptionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateKeyDescription"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) UpdateRotationPolicyWithOptions(request *UpdateRotationPolicyRequest, runtime *util.RuntimeOptions) (_result *UpdateRotationPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateRotationPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateRotationPolicy"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) UpdateSecretRotationPolicyWithOptions(request *UpdateSecretRotationPolicyRequest, runtime *util.RuntimeOptions) (_result *UpdateSecretRotationPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateSecretRotationPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateSecretRotationPolicy"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) UpdateSecretVersionStageWithOptions(request *UpdateSecretVersionStageRequest, runtime *util.RuntimeOptions) (_result *UpdateSecretVersionStageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateSecretVersionStageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateSecretVersionStage"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) UploadCertificateWithOptions(request *UploadCertificateRequest, runtime *util.RuntimeOptions) (_result *UploadCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UploadCertificateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UploadCertificate"), tea.String("2016-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
