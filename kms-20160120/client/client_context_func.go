// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

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
func (client *Client) AsymmetricDecryptWithContext(ctx context.Context, request *AsymmetricDecryptRequest, runtime *dara.RuntimeOptions) (_result *AsymmetricDecryptResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.CiphertextBlob) {
		query["CiphertextBlob"] = request.CiphertextBlob
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.KeyVersionId) {
		query["KeyVersionId"] = request.KeyVersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AsymmetricDecrypt"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AsymmetricDecryptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) AsymmetricEncryptWithContext(ctx context.Context, request *AsymmetricEncryptRequest, runtime *dara.RuntimeOptions) (_result *AsymmetricEncryptResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.KeyVersionId) {
		query["KeyVersionId"] = request.KeyVersionId
	}

	if !dara.IsNil(request.Plaintext) {
		query["Plaintext"] = request.Plaintext
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AsymmetricEncrypt"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AsymmetricEncryptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AsymmetricSign
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
func (client *Client) AsymmetricSignWithContext(ctx context.Context, request *AsymmetricSignRequest, runtime *dara.RuntimeOptions) (_result *AsymmetricSignResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.Digest) {
		query["Digest"] = request.Digest
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.KeyVersionId) {
		query["KeyVersionId"] = request.KeyVersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AsymmetricSign"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AsymmetricSignResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) AsymmetricVerifyWithContext(ctx context.Context, request *AsymmetricVerifyRequest, runtime *dara.RuntimeOptions) (_result *AsymmetricVerifyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.Digest) {
		query["Digest"] = request.Digest
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.KeyVersionId) {
		query["KeyVersionId"] = request.KeyVersionId
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AsymmetricVerify"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AsymmetricVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) CancelKeyDeletionWithContext(ctx context.Context, request *CancelKeyDeletionRequest, runtime *dara.RuntimeOptions) (_result *CancelKeyDeletionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelKeyDeletion"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelKeyDeletionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) CertificatePrivateKeyDecryptWithContext(ctx context.Context, request *CertificatePrivateKeyDecryptRequest, runtime *dara.RuntimeOptions) (_result *CertificatePrivateKeyDecryptResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.CertificateId) {
		query["CertificateId"] = request.CertificateId
	}

	if !dara.IsNil(request.CiphertextBlob) {
		query["CiphertextBlob"] = request.CiphertextBlob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CertificatePrivateKeyDecrypt"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CertificatePrivateKeyDecryptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) CertificatePrivateKeySignWithContext(ctx context.Context, request *CertificatePrivateKeySignRequest, runtime *dara.RuntimeOptions) (_result *CertificatePrivateKeySignResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.CertificateId) {
		query["CertificateId"] = request.CertificateId
	}

	if !dara.IsNil(request.Message) {
		query["Message"] = request.Message
	}

	if !dara.IsNil(request.MessageType) {
		query["MessageType"] = request.MessageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CertificatePrivateKeySign"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CertificatePrivateKeySignResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) CertificatePublicKeyEncryptWithContext(ctx context.Context, request *CertificatePublicKeyEncryptRequest, runtime *dara.RuntimeOptions) (_result *CertificatePublicKeyEncryptResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.CertificateId) {
		query["CertificateId"] = request.CertificateId
	}

	if !dara.IsNil(request.Plaintext) {
		query["Plaintext"] = request.Plaintext
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CertificatePublicKeyEncrypt"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CertificatePublicKeyEncryptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) CertificatePublicKeyVerifyWithContext(ctx context.Context, request *CertificatePublicKeyVerifyRequest, runtime *dara.RuntimeOptions) (_result *CertificatePublicKeyVerifyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.CertificateId) {
		query["CertificateId"] = request.CertificateId
	}

	if !dara.IsNil(request.Message) {
		query["Message"] = request.Message
	}

	if !dara.IsNil(request.MessageType) {
		query["MessageType"] = request.MessageType
	}

	if !dara.IsNil(request.SignatureValue) {
		query["SignatureValue"] = request.SignatureValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CertificatePublicKeyVerify"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CertificatePublicKeyVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) ConnectKmsInstanceWithContext(ctx context.Context, request *ConnectKmsInstanceRequest, runtime *dara.RuntimeOptions) (_result *ConnectKmsInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KMProvider) {
		query["KMProvider"] = request.KMProvider
	}

	if !dara.IsNil(request.KmsInstanceId) {
		query["KmsInstanceId"] = request.KmsInstanceId
	}

	if !dara.IsNil(request.VSwitchIds) {
		query["VSwitchIds"] = request.VSwitchIds
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneIds) {
		query["ZoneIds"] = request.ZoneIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConnectKmsInstance"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConnectKmsInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
//	  Each alias can be bound to only one CMK at a time.
//
//		- The aliases of CMKs in the same region must be unique.
//
// In this topic, an alias named `alias/example` is created for a CMK named `7906979c-8e06-46a2-be2d-68e3ccbc****`.
//
// @param request - CreateAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAliasResponse
func (client *Client) CreateAliasWithContext(ctx context.Context, request *CreateAliasRequest, runtime *dara.RuntimeOptions) (_result *CreateAliasResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliasName) {
		query["AliasName"] = request.AliasName
	}

	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAlias"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAliasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) CreateApplicationAccessPointWithContext(ctx context.Context, request *CreateApplicationAccessPointRequest, runtime *dara.RuntimeOptions) (_result *CreateApplicationAccessPointResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthenticationMethod) {
		query["AuthenticationMethod"] = request.AuthenticationMethod
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Policies) {
		query["Policies"] = request.Policies
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApplicationAccessPoint"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApplicationAccessPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) CreateCertificateWithContext(ctx context.Context, tmpReq *CreateCertificateRequest, runtime *dara.RuntimeOptions) (_result *CreateCertificateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateCertificateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SubjectAlternativeNames) {
		request.SubjectAlternativeNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubjectAlternativeNames, dara.String("SubjectAlternativeNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ExportablePrivateKey) {
		query["ExportablePrivateKey"] = request.ExportablePrivateKey
	}

	if !dara.IsNil(request.KeySpec) {
		query["KeySpec"] = request.KeySpec
	}

	if !dara.IsNil(request.Subject) {
		query["Subject"] = request.Subject
	}

	if !dara.IsNil(request.SubjectAlternativeNamesShrink) {
		query["SubjectAlternativeNames"] = request.SubjectAlternativeNamesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCertificate"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) CreateClientKeyWithContext(ctx context.Context, request *CreateClientKeyRequest, runtime *dara.RuntimeOptions) (_result *CreateClientKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AapName) {
		query["AapName"] = request.AapName
	}

	if !dara.IsNil(request.NotAfter) {
		query["NotAfter"] = request.NotAfter
	}

	if !dara.IsNil(request.NotBefore) {
		query["NotBefore"] = request.NotBefore
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateClientKey"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClientKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) CreateKeyWithContext(ctx context.Context, request *CreateKeyRequest, runtime *dara.RuntimeOptions) (_result *CreateKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DKMSInstanceId) {
		query["DKMSInstanceId"] = request.DKMSInstanceId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnableAutomaticRotation) {
		query["EnableAutomaticRotation"] = request.EnableAutomaticRotation
	}

	if !dara.IsNil(request.KeySpec) {
		query["KeySpec"] = request.KeySpec
	}

	if !dara.IsNil(request.KeyStorageMechanism) {
		query["KeyStorageMechanism"] = request.KeyStorageMechanism
	}

	if !dara.IsNil(request.KeyUsage) {
		query["KeyUsage"] = request.KeyUsage
	}

	if !dara.IsNil(request.Origin) {
		query["Origin"] = request.Origin
	}

	if !dara.IsNil(request.Policy) {
		query["Policy"] = request.Policy
	}

	if !dara.IsNil(request.ProtectionLevel) {
		query["ProtectionLevel"] = request.ProtectionLevel
	}

	if !dara.IsNil(request.RotationInterval) {
		query["RotationInterval"] = request.RotationInterval
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateKey"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 为密钥创建新的密钥版本。
//
// Description:
//
//	  You can create a version only for an asymmetric CMK that is in the Enabled state. You can call the [CreateKey](https://help.aliyun.com/document_detail/28947.html) operation to create an asymmetric CMK and the [DescribeKey](https://help.aliyun.com/document_detail/28952.html) operation to query the status of the CMK. The status is specified by the KeyState parameter.
//
//		- The minimum interval for creating a version of the same CMK is seven days. You can call the [DescribeKey](https://help.aliyun.com/document_detail/28952.html) operation to query the time when the last version of a CMK was created. The time is specified by the LastRotationDate parameter.
//
//		- If a CMK is in a private key store, you cannot create a version for the CMK.
//
//		- You can create a maximum of 50 versions for a CMK in the same region.
//
// You can create a version for the CMK whose ID is `0b30658a-ed1a-4922-b8f7-a673ca9c****` by using the parameter settings provided in this topic.
//
// @param request - CreateKeyVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKeyVersionResponse
func (client *Client) CreateKeyVersionWithContext(ctx context.Context, request *CreateKeyVersionRequest, runtime *dara.RuntimeOptions) (_result *CreateKeyVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateKeyVersion"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateKeyVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) CreateNetworkRuleWithContext(ctx context.Context, request *CreateNetworkRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateNetworkRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.SourcePrivateIp) {
		query["SourcePrivateIp"] = request.SourcePrivateIp
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNetworkRule"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNetworkRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) CreatePolicyWithContext(ctx context.Context, request *CreatePolicyRequest, runtime *dara.RuntimeOptions) (_result *CreatePolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessControlRules) {
		query["AccessControlRules"] = request.AccessControlRules
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.KmsInstance) {
		query["KmsInstance"] = request.KmsInstance
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Permissions) {
		query["Permissions"] = request.Permissions
	}

	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePolicy"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
//   - If the SecretType parameter is set to Generic or Rds, the name cannot start with `acs/`.
//
//   - If the SecretType parameter is set to RAMCredentials, set the SecretName parameter to `$Auto`. In this case, KMS automatically generates a secret name that starts with `acs/ram/user/`. The name includes the display name of RAM user.
//
//   - If the SecretType parameter is set to ECS, the name must start with `acs/ecs/`.
//
// @param tmpReq - CreateSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecretResponse
func (client *Client) CreateSecretWithContext(ctx context.Context, tmpReq *CreateSecretRequest, runtime *dara.RuntimeOptions) (_result *CreateSecretResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateSecretShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExtendedConfig) {
		request.ExtendedConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtendedConfig, dara.String("ExtendedConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DKMSInstanceId) {
		query["DKMSInstanceId"] = request.DKMSInstanceId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnableAutomaticRotation) {
		query["EnableAutomaticRotation"] = request.EnableAutomaticRotation
	}

	if !dara.IsNil(request.EncryptionKeyId) {
		query["EncryptionKeyId"] = request.EncryptionKeyId
	}

	if !dara.IsNil(request.ExtendedConfigShrink) {
		query["ExtendedConfig"] = request.ExtendedConfigShrink
	}

	if !dara.IsNil(request.Policy) {
		query["Policy"] = request.Policy
	}

	if !dara.IsNil(request.RotationInterval) {
		query["RotationInterval"] = request.RotationInterval
	}

	if !dara.IsNil(request.SecretData) {
		query["SecretData"] = request.SecretData
	}

	if !dara.IsNil(request.SecretDataType) {
		query["SecretDataType"] = request.SecretDataType
	}

	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	if !dara.IsNil(request.SecretType) {
		query["SecretType"] = request.SecretType
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSecret"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSecretResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) DecryptWithContext(ctx context.Context, tmpReq *DecryptRequest, runtime *dara.RuntimeOptions) (_result *DecryptResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DecryptShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EncryptionContext) {
		request.EncryptionContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EncryptionContext, dara.String("EncryptionContext"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CiphertextBlob) {
		query["CiphertextBlob"] = request.CiphertextBlob
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EncryptionContextShrink) {
		query["EncryptionContext"] = request.EncryptionContextShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Decrypt"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DecryptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAliasResponse
func (client *Client) DeleteAliasWithContext(ctx context.Context, request *DeleteAliasRequest, runtime *dara.RuntimeOptions) (_result *DeleteAliasResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliasName) {
		query["AliasName"] = request.AliasName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAlias"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAliasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) DeleteApplicationAccessPointWithContext(ctx context.Context, request *DeleteApplicationAccessPointRequest, runtime *dara.RuntimeOptions) (_result *DeleteApplicationAccessPointResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApplicationAccessPoint"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApplicationAccessPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) DeleteCertificateWithContext(ctx context.Context, request *DeleteCertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertificateId) {
		query["CertificateId"] = request.CertificateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCertificate"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) DeleteClientKeyWithContext(ctx context.Context, request *DeleteClientKeyRequest, runtime *dara.RuntimeOptions) (_result *DeleteClientKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientKeyId) {
		query["ClientKeyId"] = request.ClientKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteClientKey"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteClientKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) DeleteKeyMaterialWithContext(ctx context.Context, request *DeleteKeyMaterialRequest, runtime *dara.RuntimeOptions) (_result *DeleteKeyMaterialResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteKeyMaterial"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteKeyMaterialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) DeleteNetworkRuleWithContext(ctx context.Context, request *DeleteNetworkRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteNetworkRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNetworkRule"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNetworkRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) DeletePolicyWithContext(ctx context.Context, request *DeletePolicyRequest, runtime *dara.RuntimeOptions) (_result *DeletePolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolicy"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) DeleteSecretWithContext(ctx context.Context, request *DeleteSecretRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecretResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ForceDeleteWithoutRecovery) {
		query["ForceDeleteWithoutRecovery"] = request.ForceDeleteWithoutRecovery
	}

	if !dara.IsNil(request.RecoveryWindowInDays) {
		query["RecoveryWindowInDays"] = request.RecoveryWindowInDays
	}

	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSecret"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSecretResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) DescribeApplicationAccessPointWithContext(ctx context.Context, request *DescribeApplicationAccessPointRequest, runtime *dara.RuntimeOptions) (_result *DescribeApplicationAccessPointResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApplicationAccessPoint"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApplicationAccessPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) DescribeCertificateWithContext(ctx context.Context, request *DescribeCertificateRequest, runtime *dara.RuntimeOptions) (_result *DescribeCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertificateId) {
		query["CertificateId"] = request.CertificateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCertificate"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) DescribeKeyWithContext(ctx context.Context, request *DescribeKeyRequest, runtime *dara.RuntimeOptions) (_result *DescribeKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeKey"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) DescribeKeyVersionWithContext(ctx context.Context, request *DescribeKeyVersionRequest, runtime *dara.RuntimeOptions) (_result *DescribeKeyVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.KeyVersionId) {
		query["KeyVersionId"] = request.KeyVersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeKeyVersion"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeKeyVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) DescribeNetworkRuleWithContext(ctx context.Context, request *DescribeNetworkRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetworkRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNetworkRule"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNetworkRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) DescribePolicyWithContext(ctx context.Context, request *DescribePolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribePolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePolicy"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) DescribeSecretWithContext(ctx context.Context, request *DescribeSecretRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecretResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FetchTags) {
		query["FetchTags"] = request.FetchTags
	}

	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecret"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecretResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) DisableKeyWithContext(ctx context.Context, request *DisableKeyRequest, runtime *dara.RuntimeOptions) (_result *DisableKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableKey"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EnableKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableKeyResponse
func (client *Client) EnableKeyWithContext(ctx context.Context, request *EnableKeyRequest, runtime *dara.RuntimeOptions) (_result *EnableKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableKey"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
//	  KMS uses the primary version of a specified CMK to encrypt data.
//
//		- Only data of 6 KB or less can be encrypted. For example, you can call this operation to encrypt RSA keys, database access passwords, or other sensitive information.
//
//		- When you migrate encrypted data across regions, you can call this operation in the destination region to encrypt the plaintext of the data key that is used to encrypt the migrated data in the source region. This way, the ciphertext of the data key is generated in the destination region. You can also call the [Decrypt](https://help.aliyun.com/document_detail/28950.html) operation to decrypt the data key.
//
// @param tmpReq - EncryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EncryptResponse
func (client *Client) EncryptWithContext(ctx context.Context, tmpReq *EncryptRequest, runtime *dara.RuntimeOptions) (_result *EncryptResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &EncryptShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EncryptionContext) {
		request.EncryptionContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EncryptionContext, dara.String("EncryptionContext"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EncryptionContextShrink) {
		query["EncryptionContext"] = request.EncryptionContextShrink
	}

	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.Plaintext) {
		query["Plaintext"] = request.Plaintext
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Encrypt"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EncryptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) ExportDataKeyWithContext(ctx context.Context, tmpReq *ExportDataKeyRequest, runtime *dara.RuntimeOptions) (_result *ExportDataKeyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ExportDataKeyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EncryptionContext) {
		request.EncryptionContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EncryptionContext, dara.String("EncryptionContext"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CiphertextBlob) {
		query["CiphertextBlob"] = request.CiphertextBlob
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EncryptionContextShrink) {
		query["EncryptionContext"] = request.EncryptionContextShrink
	}

	if !dara.IsNil(request.PublicKeyBlob) {
		query["PublicKeyBlob"] = request.PublicKeyBlob
	}

	if !dara.IsNil(request.WrappingAlgorithm) {
		query["WrappingAlgorithm"] = request.WrappingAlgorithm
	}

	if !dara.IsNil(request.WrappingKeySpec) {
		query["WrappingKeySpec"] = request.WrappingKeySpec
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportDataKey"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportDataKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// We recommend that you perform the following steps to import your data key to a cryptographic module:
//
//   - Call the GenerateAndExportDataKey operation to generate a data key and obtain both the ciphertext of the data key encrypted by using the CMK and that encrypted by using the public key.
//
//   - Store the ciphertext of the data key encrypted by using the CMK in KMS Secrets Manager or in a storage service such as ApsaraDB. This ciphertext is used for backup and restoration.
//
//   - Import the ciphertext of the data key encrypted by using the public key to the cryptographic module where the private key is stored. Then, you can use the data key to encrypt or decrypt data.
//
// >  The CMK that you specify in the request of this operation is only used to encrypt the data key and is not involved in the generation of the data key. KMS does not record or store the data keys randomly generated by calling this operation. You must take note of the data keys and the returned ciphertext.
//
// @param tmpReq - GenerateAndExportDataKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateAndExportDataKeyResponse
func (client *Client) GenerateAndExportDataKeyWithContext(ctx context.Context, tmpReq *GenerateAndExportDataKeyRequest, runtime *dara.RuntimeOptions) (_result *GenerateAndExportDataKeyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GenerateAndExportDataKeyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EncryptionContext) {
		request.EncryptionContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EncryptionContext, dara.String("EncryptionContext"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EncryptionContextShrink) {
		query["EncryptionContext"] = request.EncryptionContextShrink
	}

	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.KeySpec) {
		query["KeySpec"] = request.KeySpec
	}

	if !dara.IsNil(request.NumberOfBytes) {
		query["NumberOfBytes"] = request.NumberOfBytes
	}

	if !dara.IsNil(request.PublicKeyBlob) {
		query["PublicKeyBlob"] = request.PublicKeyBlob
	}

	if !dara.IsNil(request.WrappingAlgorithm) {
		query["WrappingAlgorithm"] = request.WrappingAlgorithm
	}

	if !dara.IsNil(request.WrappingKeySpec) {
		query["WrappingKeySpec"] = request.WrappingKeySpec
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateAndExportDataKey"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateAndExportDataKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
//   - Call the [Decrypt](https://help.aliyun.com/document_detail/28950.html) operation to decrypt the locally stored ciphertext of the data key. The plaintext of data key is then returned.
//
//   - Use the plaintext of the data key to locally decrypt data and then delete the plaintext of the data key from the memory.
//
// In this example, a random data key is generated for the CMK whose ID is `7906979c-8e06-46a2-be2d-68e3ccbc****`.
//
// @param tmpReq - GenerateDataKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateDataKeyResponse
func (client *Client) GenerateDataKeyWithContext(ctx context.Context, tmpReq *GenerateDataKeyRequest, runtime *dara.RuntimeOptions) (_result *GenerateDataKeyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GenerateDataKeyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EncryptionContext) {
		request.EncryptionContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EncryptionContext, dara.String("EncryptionContext"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EncryptionContextShrink) {
		query["EncryptionContext"] = request.EncryptionContextShrink
	}

	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.KeySpec) {
		query["KeySpec"] = request.KeySpec
	}

	if !dara.IsNil(request.NumberOfBytes) {
		query["NumberOfBytes"] = request.NumberOfBytes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateDataKey"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateDataKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) GenerateDataKeyWithoutPlaintextWithContext(ctx context.Context, tmpReq *GenerateDataKeyWithoutPlaintextRequest, runtime *dara.RuntimeOptions) (_result *GenerateDataKeyWithoutPlaintextResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GenerateDataKeyWithoutPlaintextShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EncryptionContext) {
		request.EncryptionContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EncryptionContext, dara.String("EncryptionContext"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EncryptionContextShrink) {
		query["EncryptionContext"] = request.EncryptionContextShrink
	}

	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.KeySpec) {
		query["KeySpec"] = request.KeySpec
	}

	if !dara.IsNil(request.NumberOfBytes) {
		query["NumberOfBytes"] = request.NumberOfBytes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateDataKeyWithoutPlaintext"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateDataKeyWithoutPlaintextResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) GetCertificateWithContext(ctx context.Context, request *GetCertificateRequest, runtime *dara.RuntimeOptions) (_result *GetCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertificateId) {
		query["CertificateId"] = request.CertificateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCertificate"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) GetClientKeyWithContext(ctx context.Context, request *GetClientKeyRequest, runtime *dara.RuntimeOptions) (_result *GetClientKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetClientKey"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClientKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) GetKeyPolicyWithContext(ctx context.Context, request *GetKeyPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetKeyPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetKeyPolicy"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetKeyPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) GetKmsInstanceWithContext(ctx context.Context, request *GetKmsInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetKmsInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KmsInstanceId) {
		query["KmsInstanceId"] = request.KmsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetKmsInstance"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetKmsInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) GetParametersForImportWithContext(ctx context.Context, request *GetParametersForImportRequest, runtime *dara.RuntimeOptions) (_result *GetParametersForImportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.WrappingAlgorithm) {
		query["WrappingAlgorithm"] = request.WrappingAlgorithm
	}

	if !dara.IsNil(request.WrappingKeySpec) {
		query["WrappingKeySpec"] = request.WrappingKeySpec
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetParametersForImport"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetParametersForImportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPublicKeyResponse
func (client *Client) GetPublicKeyWithContext(ctx context.Context, request *GetPublicKeyRequest, runtime *dara.RuntimeOptions) (_result *GetPublicKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.KeyVersionId) {
		query["KeyVersionId"] = request.KeyVersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPublicKey"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPublicKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetRandomPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRandomPasswordResponse
func (client *Client) GetRandomPasswordWithContext(ctx context.Context, request *GetRandomPasswordRequest, runtime *dara.RuntimeOptions) (_result *GetRandomPasswordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExcludeCharacters) {
		query["ExcludeCharacters"] = request.ExcludeCharacters
	}

	if !dara.IsNil(request.ExcludeLowercase) {
		query["ExcludeLowercase"] = request.ExcludeLowercase
	}

	if !dara.IsNil(request.ExcludeNumbers) {
		query["ExcludeNumbers"] = request.ExcludeNumbers
	}

	if !dara.IsNil(request.ExcludePunctuation) {
		query["ExcludePunctuation"] = request.ExcludePunctuation
	}

	if !dara.IsNil(request.ExcludeUppercase) {
		query["ExcludeUppercase"] = request.ExcludeUppercase
	}

	if !dara.IsNil(request.PasswordLength) {
		query["PasswordLength"] = request.PasswordLength
	}

	if !dara.IsNil(request.RequireEachIncludedType) {
		query["RequireEachIncludedType"] = request.RequireEachIncludedType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRandomPassword"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRandomPasswordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) GetSecretPolicyWithContext(ctx context.Context, request *GetSecretPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetSecretPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecretPolicy"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecretPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) GetSecretValueWithContext(ctx context.Context, request *GetSecretValueRequest, runtime *dara.RuntimeOptions) (_result *GetSecretValueResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FetchExtendedConfig) {
		query["FetchExtendedConfig"] = request.FetchExtendedConfig
	}

	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	if !dara.IsNil(request.VersionStage) {
		query["VersionStage"] = request.VersionStage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecretValue"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecretValueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
//   - To view the CMK **Origin**, see [DescribeKey](https://help.aliyun.com/document_detail/28952.html).
//
//   - Before importing key material, you need to call the [GetParametersForImport](https://help.aliyun.com/document_detail/68621.html) obtain the parameters required to import the key material, including the public key and import token.
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
func (client *Client) ImportKeyMaterialWithContext(ctx context.Context, request *ImportKeyMaterialRequest, runtime *dara.RuntimeOptions) (_result *ImportKeyMaterialResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EncryptedKeyMaterial) {
		query["EncryptedKeyMaterial"] = request.EncryptedKeyMaterial
	}

	if !dara.IsNil(request.ImportToken) {
		query["ImportToken"] = request.ImportToken
	}

	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.KeyMaterialExpireUnix) {
		query["KeyMaterialExpireUnix"] = request.KeyMaterialExpireUnix
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportKeyMaterial"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportKeyMaterialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) ListAliasesWithContext(ctx context.Context, request *ListAliasesRequest, runtime *dara.RuntimeOptions) (_result *ListAliasesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAliases"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAliasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListAliasesByKeyIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAliasesByKeyIdResponse
func (client *Client) ListAliasesByKeyIdWithContext(ctx context.Context, request *ListAliasesByKeyIdRequest, runtime *dara.RuntimeOptions) (_result *ListAliasesByKeyIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAliasesByKeyId"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAliasesByKeyIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) ListApplicationAccessPointsWithContext(ctx context.Context, request *ListApplicationAccessPointsRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationAccessPointsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationAccessPoints"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationAccessPointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListClientKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClientKeysResponse
func (client *Client) ListClientKeysWithContext(ctx context.Context, request *ListClientKeysRequest, runtime *dara.RuntimeOptions) (_result *ListClientKeysResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClientKeys"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClientKeysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) ListKeyVersionsWithContext(ctx context.Context, request *ListKeyVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListKeyVersionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListKeyVersions"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListKeyVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) ListKeysWithContext(ctx context.Context, request *ListKeysRequest, runtime *dara.RuntimeOptions) (_result *ListKeysResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListKeys"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListKeysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) ListKmsInstancesWithContext(ctx context.Context, request *ListKmsInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListKmsInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListKmsInstances"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListKmsInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) ListNetworkRulesWithContext(ctx context.Context, request *ListNetworkRulesRequest, runtime *dara.RuntimeOptions) (_result *ListNetworkRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNetworkRules"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNetworkRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) ListPoliciesWithContext(ctx context.Context, request *ListPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListPoliciesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicies"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) ListResourceTagsWithContext(ctx context.Context, request *ListResourceTagsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceTagsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceTags"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceTagsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) ListSecretVersionIdsWithContext(ctx context.Context, request *ListSecretVersionIdsRequest, runtime *dara.RuntimeOptions) (_result *ListSecretVersionIdsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IncludeDeprecated) {
		query["IncludeDeprecated"] = request.IncludeDeprecated
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSecretVersionIds"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSecretVersionIdsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// Specifies whether to return the resource tags of the secret. Valid values:
//
//   - true: returns the resource tags.
//
//   - false: does not return the resource tags. This is the default value.
//
// @param request - ListSecretsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecretsResponse
func (client *Client) ListSecretsWithContext(ctx context.Context, request *ListSecretsRequest, runtime *dara.RuntimeOptions) (_result *ListSecretsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FetchTags) {
		query["FetchTags"] = request.FetchTags
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSecrets"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSecretsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
//   - If the specified version number does not exist in the secret, Secrets Manager creates the version and stores the secret value.
//
//   - If the specified version number already exists in the secret and the secret value of the existing version is the same as the secret value that you specify, Secrets Manager ignores the request and returns a success message. The request is idempotent.
//
//   - If the specified version number already exists in the secret but the secret value of the existing version is different from the secret value that you specify, Secrets Manager rejects the request and returns a failure message.
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
func (client *Client) PutSecretValueWithContext(ctx context.Context, request *PutSecretValueRequest, runtime *dara.RuntimeOptions) (_result *PutSecretValueResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecretData) {
		query["SecretData"] = request.SecretData
	}

	if !dara.IsNil(request.SecretDataType) {
		query["SecretDataType"] = request.SecretDataType
	}

	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	if !dara.IsNil(request.VersionStages) {
		query["VersionStages"] = request.VersionStages
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutSecretValue"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutSecretValueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// You can call this operation in the following scenarios:
//
//   - After the CMK that was used to encrypt your data is rotated, you can call this operation to use the latest CMK version to re-encrypt the data. For more information about automatic key rotation, see [Configure automatic key rotation](https://help.aliyun.com/document_detail/134270.html).
//
//   - The CMK that was used to encrypt your data remains unchanged, but EncryptionContext is changed. In this scenario, you can call this operation to re-encrypt the data.
//
//   - You can call this operation to use a CMK in KMS to re-encrypt data or a data key that was previously encrypted by a different CMK.
//
// To use the ReEncrypt operation, you must have two permissions:
//
//   - kms:ReEncryptFrom on the source CMK
//
//   - kms:ReEncryptTo on the destination CMK
//
//   - For simplicity, you can specify kms:ReEncrypt\\	- to allow both of the preceding permissions.
//
// @param tmpReq - ReEncryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReEncryptResponse
func (client *Client) ReEncryptWithContext(ctx context.Context, tmpReq *ReEncryptRequest, runtime *dara.RuntimeOptions) (_result *ReEncryptResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ReEncryptShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DestinationEncryptionContext) {
		request.DestinationEncryptionContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DestinationEncryptionContext, dara.String("DestinationEncryptionContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceEncryptionContext) {
		request.SourceEncryptionContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceEncryptionContext, dara.String("SourceEncryptionContext"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CiphertextBlob) {
		query["CiphertextBlob"] = request.CiphertextBlob
	}

	if !dara.IsNil(request.DestinationEncryptionContextShrink) {
		query["DestinationEncryptionContext"] = request.DestinationEncryptionContextShrink
	}

	if !dara.IsNil(request.DestinationKeyId) {
		query["DestinationKeyId"] = request.DestinationKeyId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.SourceEncryptionAlgorithm) {
		query["SourceEncryptionAlgorithm"] = request.SourceEncryptionAlgorithm
	}

	if !dara.IsNil(request.SourceEncryptionContextShrink) {
		query["SourceEncryptionContext"] = request.SourceEncryptionContextShrink
	}

	if !dara.IsNil(request.SourceKeyId) {
		query["SourceKeyId"] = request.SourceKeyId
	}

	if !dara.IsNil(request.SourceKeyVersionId) {
		query["SourceKeyVersionId"] = request.SourceKeyVersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReEncrypt"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReEncryptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 仅后付费实例支持释放，预付费实例需要从用户中心-退订管理释放。
//
// @param request - ReleaseKmsInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseKmsInstanceResponse
func (client *Client) ReleaseKmsInstanceWithContext(ctx context.Context, request *ReleaseKmsInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReleaseKmsInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ForceDeleteWithoutBackup) {
		query["ForceDeleteWithoutBackup"] = request.ForceDeleteWithoutBackup
	}

	if !dara.IsNil(request.KmsInstanceId) {
		query["KmsInstanceId"] = request.KmsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseKmsInstance"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseKmsInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) RestoreSecretWithContext(ctx context.Context, request *RestoreSecretRequest, runtime *dara.RuntimeOptions) (_result *RestoreSecretResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestoreSecret"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestoreSecretResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) RotateSecretWithContext(ctx context.Context, request *RotateSecretRequest, runtime *dara.RuntimeOptions) (_result *RotateSecretResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RotateSecret"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RotateSecretResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) ScheduleKeyDeletionWithContext(ctx context.Context, request *ScheduleKeyDeletionRequest, runtime *dara.RuntimeOptions) (_result *ScheduleKeyDeletionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.PendingWindowInDays) {
		query["PendingWindowInDays"] = request.PendingWindowInDays
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ScheduleKeyDeletion"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ScheduleKeyDeletionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables deletion protection for a customer master key (CMK).
//
// Description:
//
//	  After you enable deletion protection for a CMK, you cannot delete the CMK. If you want to delete the CMK, you must first disable deletion protection for the CMK.
//
//		- Before you can call the SetDeletionProtection operation, make sure that the required CMK is not in the Pending Deletion state. You can call the [DescribeKey](https://help.aliyun.com/document_detail/28952.html) operation to query the CMK status, which is specified by the KeyState parameter.
//
// You can enable deletion protection for the CMK whose Alibaba Cloud Resource Name (ARN) is `acs:kms:cn-hangzhou:123213123****:key/0225f411-b21d-46d1-be5b-93931c82****` by using parameter settings provided in this topic. The CMK ARN is specified by the ProtectedResourceArn parameter.
//
// @param request - SetDeletionProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDeletionProtectionResponse
func (client *Client) SetDeletionProtectionWithContext(ctx context.Context, request *SetDeletionProtectionRequest, runtime *dara.RuntimeOptions) (_result *SetDeletionProtectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeletionProtectionDescription) {
		query["DeletionProtectionDescription"] = request.DeletionProtectionDescription
	}

	if !dara.IsNil(request.EnableDeletionProtection) {
		query["EnableDeletionProtection"] = request.EnableDeletionProtection
	}

	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.ProtectedResourceArn) {
		query["ProtectedResourceArn"] = request.ProtectedResourceArn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDeletionProtection"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDeletionProtectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) SetKeyPolicyWithContext(ctx context.Context, request *SetKeyPolicyRequest, runtime *dara.RuntimeOptions) (_result *SetKeyPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.Policy) {
		query["Policy"] = request.Policy
	}

	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetKeyPolicy"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetKeyPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) SetSecretPolicyWithContext(ctx context.Context, request *SetSecretPolicyRequest, runtime *dara.RuntimeOptions) (_result *SetSecretPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Policy) {
		query["Policy"] = request.Policy
	}

	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetSecretPolicy"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetSecretPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) TagResourceWithContext(ctx context.Context, request *TagResourceRequest, runtime *dara.RuntimeOptions) (_result *TagResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertificateId) {
		query["CertificateId"] = request.CertificateId
	}

	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResource"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) UntagResourceWithContext(ctx context.Context, request *UntagResourceRequest, runtime *dara.RuntimeOptions) (_result *UntagResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertificateId) {
		query["CertificateId"] = request.CertificateId
	}

	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	if !dara.IsNil(request.TagKeys) {
		query["TagKeys"] = request.TagKeys
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResource"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAliasResponse
func (client *Client) UpdateAliasWithContext(ctx context.Context, request *UpdateAliasRequest, runtime *dara.RuntimeOptions) (_result *UpdateAliasResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliasName) {
		query["AliasName"] = request.AliasName
	}

	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAlias"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAliasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) UpdateApplicationAccessPointWithContext(ctx context.Context, request *UpdateApplicationAccessPointRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationAccessPointResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Policies) {
		query["Policies"] = request.Policies
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApplicationAccessPoint"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationAccessPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) UpdateCertificateStatusWithContext(ctx context.Context, request *UpdateCertificateStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateCertificateStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertificateId) {
		query["CertificateId"] = request.CertificateId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCertificateStatus"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCertificateStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) UpdateKeyDescriptionWithContext(ctx context.Context, request *UpdateKeyDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateKeyDescriptionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateKeyDescription"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateKeyDescriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) UpdateKmsInstanceBindVpcWithContext(ctx context.Context, request *UpdateKmsInstanceBindVpcRequest, runtime *dara.RuntimeOptions) (_result *UpdateKmsInstanceBindVpcResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateKmsInstanceBindVpc"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateKmsInstanceBindVpcResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) UpdateNetworkRuleWithContext(ctx context.Context, request *UpdateNetworkRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateNetworkRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.SourcePrivateIp) {
		query["SourcePrivateIp"] = request.SourcePrivateIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNetworkRule"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNetworkRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) UpdatePolicyWithContext(ctx context.Context, request *UpdatePolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdatePolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessControlRules) {
		query["AccessControlRules"] = request.AccessControlRules
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Permissions) {
		query["Permissions"] = request.Permissions
	}

	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePolicy"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// When automatic key rotation is enabled, KMS automatically creates a key version after the preset rotation period arrives. In addition, KMS sets the new key version as the primary key version.
//
// An automatic key rotation policy cannot be configured for the following keys:
//
//   - Asymmetric key
//
//   - Service-managed key
//
//   - Bring your own key (BYOK) that is imported into KMS
//
//   - Key that is not in the **Enabled*	- state
//
// In this example, automatic key rotation is enabled for a CMK whose ID is `1234abcd-12ab-34cd-56ef-12345678****`. The automatic rotation period is 30 days.
//
// @param request - UpdateRotationPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRotationPolicyResponse
func (client *Client) UpdateRotationPolicyWithContext(ctx context.Context, request *UpdateRotationPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateRotationPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableAutomaticRotation) {
		query["EnableAutomaticRotation"] = request.EnableAutomaticRotation
	}

	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.RotationInterval) {
		query["RotationInterval"] = request.RotationInterval
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRotationPolicy"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRotationPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) UpdateSecretWithContext(ctx context.Context, request *UpdateSecretRequest, runtime *dara.RuntimeOptions) (_result *UpdateSecretResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	if !dara.IsNil(request.ExtendedConfig) {
		query["ExtendedConfig"] = request.ExtendedConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSecret"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSecretResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
//   - The `EnableAutomaticRotation` parameter is set to `true`, which indicates that automatic rotation is enabled.
//
//   - The `RotationInterval` parameter is set to `30d`, which indicates that the interval for automatic rotation is 30 days.
//
// @param request - UpdateSecretRotationPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecretRotationPolicyResponse
func (client *Client) UpdateSecretRotationPolicyWithContext(ctx context.Context, request *UpdateSecretRotationPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateSecretRotationPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableAutomaticRotation) {
		query["EnableAutomaticRotation"] = request.EnableAutomaticRotation
	}

	if !dara.IsNil(request.RotationInterval) {
		query["RotationInterval"] = request.RotationInterval
	}

	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSecretRotationPolicy"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSecretRotationPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # UpdateSecretVersionStage
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
func (client *Client) UpdateSecretVersionStageWithContext(ctx context.Context, request *UpdateSecretVersionStageRequest, runtime *dara.RuntimeOptions) (_result *UpdateSecretVersionStageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MoveToVersion) {
		query["MoveToVersion"] = request.MoveToVersion
	}

	if !dara.IsNil(request.RemoveFromVersion) {
		query["RemoveFromVersion"] = request.RemoveFromVersion
	}

	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	if !dara.IsNil(request.VersionStage) {
		query["VersionStage"] = request.VersionStage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSecretVersionStage"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSecretVersionStageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) UploadCertificateWithContext(ctx context.Context, request *UploadCertificateRequest, runtime *dara.RuntimeOptions) (_result *UploadCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Certificate) {
		query["Certificate"] = request.Certificate
	}

	if !dara.IsNil(request.CertificateChain) {
		query["CertificateChain"] = request.CertificateChain
	}

	if !dara.IsNil(request.CertificateId) {
		query["CertificateId"] = request.CertificateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadCertificate"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
