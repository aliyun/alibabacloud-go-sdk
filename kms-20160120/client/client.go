// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-pop/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

type Client struct {
	openapi.Client
	DisableSDKError *bool
	EnableValidate  *bool
}

func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.ProductId = dara.String("Kms")
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"us-west-1":             dara.String("kms.us-west-1.aliyuncs.com"),
		"us-east-1":             dara.String("kms.us-east-1.aliyuncs.com"),
		"me-east-1":             dara.String("kms.me-east-1.aliyuncs.com"),
		"me-central-1":          dara.String("kms.me-central-1.aliyuncs.com"),
		"eu-west-1":             dara.String("kms.eu-west-1.aliyuncs.com"),
		"eu-central-1":          dara.String("kms.eu-central-1.aliyuncs.com"),
		"cn-zhengzhou-jva":      dara.String("kms.cn-zhengzhou-jva.aliyuncs.com"),
		"cn-zhangjiakou":        dara.String("kms.cn-zhangjiakou.aliyuncs.com"),
		"cn-wulanchabu":         dara.String("kms.cn-wulanchabu.aliyuncs.com"),
		"cn-wuhan-lr":           dara.String("kms.cn-wuhan-lr.aliyuncs.com"),
		"cn-shenzhen-finance-1": dara.String("kms.cn-shenzhen-finance-1.aliyuncs.com"),
		"cn-shenzhen":           dara.String("kms.cn-shenzhen.aliyuncs.com"),
		"cn-shanghai-finance-1": dara.String("kms.cn-shanghai-finance-1.aliyuncs.com"),
		"cn-shanghai":           dara.String("kms.cn-shanghai.aliyuncs.com"),
		"cn-qingdao":            dara.String("kms.cn-qingdao.aliyuncs.com"),
		"cn-huhehaote":          dara.String("kms.cn-huhehaote.aliyuncs.com"),
		"cn-hongkong":           dara.String("kms.cn-hongkong.aliyuncs.com"),
		"cn-heyuan":             dara.String("kms.cn-heyuan.aliyuncs.com"),
		"cn-hangzhou-finance":   dara.String("kms.cn-hangzhou-finance.aliyuncs.com"),
		"cn-hangzhou":           dara.String("kms.cn-hangzhou.aliyuncs.com"),
		"cn-guangzhou":          dara.String("kms.cn-guangzhou.aliyuncs.com"),
		"cn-fuzhou":             dara.String("kms.cn-fuzhou.aliyuncs.com"),
		"cn-chengdu":            dara.String("kms.cn-chengdu.aliyuncs.com"),
		"cn-beijing-finance-1":  dara.String("kms.cn-beijing-finance-1.aliyuncs.com"),
		"cn-beijing":            dara.String("kms.cn-beijing.aliyuncs.com"),
		"ap-southeast-7":        dara.String("kms.ap-southeast-7.aliyuncs.com"),
		"ap-southeast-6":        dara.String("kms.ap-southeast-6.aliyuncs.com"),
		"ap-southeast-5":        dara.String("kms.ap-southeast-5.aliyuncs.com"),
		"ap-southeast-3":        dara.String("kms.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-1":        dara.String("kms.ap-southeast-1.aliyuncs.com"),
		"ap-northeast-2":        dara.String("kms.ap-northeast-2.aliyuncs.com"),
		"ap-northeast-1":        dara.String("kms.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("kms"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !dara.IsNil(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		_result = endpointMap[dara.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Decrypts data by using the private key of an asymmetric CMK.
//
// Description:
//
// ### Usage notes
//
// - For information about the access policy required for a RAM user or RAM role to call this API operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation can be called through a shared gateway or a dedicated gateway. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//   - Shared gateway: You can access KMS over the Internet or a VPC. To access KMS over the Internet, you must enable the public endpoint. For more information, see [Access keys in a KMS instance over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
//   - Dedicated gateway: You can access KMS using the private endpoint of KMS (`<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com`).
//
// ### QPS limits
//
// - If you use a shared gateway, the queries per second (QPS) limit for each Alibaba Cloud account is 200. If the QPS exceeds the limit, the API call is throttled. This can affect your business. We recommend that you plan your calls to avoid exceeding this limit.
//
// - If you use a dedicated gateway, the QPS limit for each Alibaba Cloud account is subject to the performance specifications of your KMS instance. For more information, see [Performance metrics](https://help.aliyun.com/document_detail/480120.html).
//
// ### Description
//
// This operation supports only asymmetric keys for which the **Usage*	- parameter is set to **ENCRYPT/DECRYPT**. The following table describes the supported encryption algorithms.
//
// | KeySpec   | Algorithm             | Description                                        | Ciphertext length (bytes) |
//
// | --------- | --------------------- | -------------------------------------------------- | ------------------------- |
//
// | RSA_2048 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256     | 256                       |
//
// | RSA_2048 | RSAES_OAEP_SHA_1   | RSAES-OAEP using SHA1 and MGF1 with SHA1           | 256                       |
//
// | RSA_3072 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256     | 384                       |
//
// | RSA_3072 | RSAES_OAEP_SHA_1   | RSAES-OAEP using SHA1 and MGF1 with SHA1           | 384                       |
//
// | EC_SM2   | SM2PKE                | SM2 elliptic curve public key encryption algorithm | Maximum 6,144             |
//
// This topic provides an example of how to use the asymmetric key whose ID is `key-hzz630494463ejqjx****` and version ID is `2ab1a983-7072-4bbc-a582-584b5bd8****` to decrypt the ciphertext `BQKP+1zK6+ZEMxTP5qaVzcsgXtWplYBKm0NXdSnB5FzliFxE1bSiu4dnEIlca2JpeH7yz1/S6fed630H+hIH6DoM25fTLNcKj+mFB0Xnh9m2+HN59Mn4qyTfcUeadnfCXSWcGBouhXFwcdd2rJ3n337bzTf4jm659gZu3L0i6PLuxM9p7mqdwO0cKJPfGVfhnfMz+f4alMg79WB/NNyE2lyX7/qxvV49ObNrrJbKSFiz8Djocaf0IESNLMbfYI5bXjWkJlX92DQbKhibtQW8ZOJ//ZC6t0AWcUoKL6QDm/dg5koQalcleRinpB+QadFm894sLbVZ9+N4GVsv1W****==` using the `RSAES_OAEP_SHA_1` decryption algorithm.
//
// @param request - AsymmetricDecryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AsymmetricDecryptResponse
func (client *Client) AsymmetricDecryptWithOptions(request *AsymmetricDecryptRequest, runtime *dara.RuntimeOptions) (_result *AsymmetricDecryptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Decrypts data by using the private key of an asymmetric CMK.
//
// Description:
//
// ### Usage notes
//
// - For information about the access policy required for a RAM user or RAM role to call this API operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation can be called through a shared gateway or a dedicated gateway. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//   - Shared gateway: You can access KMS over the Internet or a VPC. To access KMS over the Internet, you must enable the public endpoint. For more information, see [Access keys in a KMS instance over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
//   - Dedicated gateway: You can access KMS using the private endpoint of KMS (`<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com`).
//
// ### QPS limits
//
// - If you use a shared gateway, the queries per second (QPS) limit for each Alibaba Cloud account is 200. If the QPS exceeds the limit, the API call is throttled. This can affect your business. We recommend that you plan your calls to avoid exceeding this limit.
//
// - If you use a dedicated gateway, the QPS limit for each Alibaba Cloud account is subject to the performance specifications of your KMS instance. For more information, see [Performance metrics](https://help.aliyun.com/document_detail/480120.html).
//
// ### Description
//
// This operation supports only asymmetric keys for which the **Usage*	- parameter is set to **ENCRYPT/DECRYPT**. The following table describes the supported encryption algorithms.
//
// | KeySpec   | Algorithm             | Description                                        | Ciphertext length (bytes) |
//
// | --------- | --------------------- | -------------------------------------------------- | ------------------------- |
//
// | RSA_2048 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256     | 256                       |
//
// | RSA_2048 | RSAES_OAEP_SHA_1   | RSAES-OAEP using SHA1 and MGF1 with SHA1           | 256                       |
//
// | RSA_3072 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256     | 384                       |
//
// | RSA_3072 | RSAES_OAEP_SHA_1   | RSAES-OAEP using SHA1 and MGF1 with SHA1           | 384                       |
//
// | EC_SM2   | SM2PKE                | SM2 elliptic curve public key encryption algorithm | Maximum 6,144             |
//
// This topic provides an example of how to use the asymmetric key whose ID is `key-hzz630494463ejqjx****` and version ID is `2ab1a983-7072-4bbc-a582-584b5bd8****` to decrypt the ciphertext `BQKP+1zK6+ZEMxTP5qaVzcsgXtWplYBKm0NXdSnB5FzliFxE1bSiu4dnEIlca2JpeH7yz1/S6fed630H+hIH6DoM25fTLNcKj+mFB0Xnh9m2+HN59Mn4qyTfcUeadnfCXSWcGBouhXFwcdd2rJ3n337bzTf4jm659gZu3L0i6PLuxM9p7mqdwO0cKJPfGVfhnfMz+f4alMg79WB/NNyE2lyX7/qxvV49ObNrrJbKSFiz8Djocaf0IESNLMbfYI5bXjWkJlX92DQbKhibtQW8ZOJ//ZC6t0AWcUoKL6QDm/dg5koQalcleRinpB+QadFm894sLbVZ9+N4GVsv1W****==` using the `RSAES_OAEP_SHA_1` decryption algorithm.
//
// @param request - AsymmetricDecryptRequest
//
// @return AsymmetricDecryptResponse
func (client *Client) AsymmetricDecrypt(request *AsymmetricDecryptRequest) (_result *AsymmetricDecryptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Encrypts data by using the public key of an asymmetric CMK.
//
// Description:
//
// ### Precautions
//
// - For information about the permissions that are required to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation can be called through a shared gateway or a dedicated gateway. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//   - Shared gateway: You can access KMS over the Internet or using a VPC domain name. To access KMS over the Internet, you must enable Internet access. For more information, see [Access a key in a KMS instance over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
//   - Dedicated gateway: You can access KMS using the private endpoint of KMS (`<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com`).
//
// ### QPS limits
//
// - If you use a shared gateway: The number of queries per second (QPS) for a single user is limited to 200. If the limit is exceeded, API calls are throttled. This may affect your business. We recommend that you plan your API calls to avoid exceeding this limit.
//
// - If you use a dedicated gateway: The QPS limit for a single user depends on the computing performance specifications of your KMS instance. For more information, see [Performance metrics](https://help.aliyun.com/document_detail/480120.html).
//
// ### Description
//
// This operation supports only asymmetric keys that have the **Usage*	- parameter set to **ENCRYPT/DECRYPT**. The following table describes the supported encryption algorithms.
//
// | KeySpec   | Algorithm             | Description                                        | Maximum number of bytes that can be encrypted |
//
// | --------- | --------------------- | -------------------------------------------------- | --------------------------------------------- |
//
// | RSA_2048 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256     | 190                                           |
//
// | RSA_2048 | RSAES_OAEP_SHA_1   | RSAES-OAEP using SHA1 and MGF1 with SHA1           | 214                                           |
//
// | RSA_3072 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256     | 318                                           |
//
// | RSA_3072 | RSAES_OAEP_SHA_1   | RSAES-OAEP using SHA1 and MGF1 with SHA1           | 342                                           |
//
// | EC_SM2   | SM2PKE                | SM2 elliptic curve public key encryption algorithm | 6047                                          |
//
// In this example, the plaintext `SGVsbG8gd29ybGQ=` is encrypted using an asymmetric key with the key ID `key-hzz630494463ejqjx****`, the key version ID `2ab1a983-7072-4bbc-a582-584b5bd8****`, and the `RSAES_OAEP_SHA_1` encryption algorithm.
//
// @param request - AsymmetricEncryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AsymmetricEncryptResponse
func (client *Client) AsymmetricEncryptWithOptions(request *AsymmetricEncryptRequest, runtime *dara.RuntimeOptions) (_result *AsymmetricEncryptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Encrypts data by using the public key of an asymmetric CMK.
//
// Description:
//
// ### Precautions
//
// - For information about the permissions that are required to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation can be called through a shared gateway or a dedicated gateway. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//   - Shared gateway: You can access KMS over the Internet or using a VPC domain name. To access KMS over the Internet, you must enable Internet access. For more information, see [Access a key in a KMS instance over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
//   - Dedicated gateway: You can access KMS using the private endpoint of KMS (`<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com`).
//
// ### QPS limits
//
// - If you use a shared gateway: The number of queries per second (QPS) for a single user is limited to 200. If the limit is exceeded, API calls are throttled. This may affect your business. We recommend that you plan your API calls to avoid exceeding this limit.
//
// - If you use a dedicated gateway: The QPS limit for a single user depends on the computing performance specifications of your KMS instance. For more information, see [Performance metrics](https://help.aliyun.com/document_detail/480120.html).
//
// ### Description
//
// This operation supports only asymmetric keys that have the **Usage*	- parameter set to **ENCRYPT/DECRYPT**. The following table describes the supported encryption algorithms.
//
// | KeySpec   | Algorithm             | Description                                        | Maximum number of bytes that can be encrypted |
//
// | --------- | --------------------- | -------------------------------------------------- | --------------------------------------------- |
//
// | RSA_2048 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256     | 190                                           |
//
// | RSA_2048 | RSAES_OAEP_SHA_1   | RSAES-OAEP using SHA1 and MGF1 with SHA1           | 214                                           |
//
// | RSA_3072 | RSAES_OAEP_SHA_256 | RSAES-OAEP using SHA-256 and MGF1 with SHA-256     | 318                                           |
//
// | RSA_3072 | RSAES_OAEP_SHA_1   | RSAES-OAEP using SHA1 and MGF1 with SHA1           | 342                                           |
//
// | EC_SM2   | SM2PKE                | SM2 elliptic curve public key encryption algorithm | 6047                                          |
//
// In this example, the plaintext `SGVsbG8gd29ybGQ=` is encrypted using an asymmetric key with the key ID `key-hzz630494463ejqjx****`, the key version ID `2ab1a983-7072-4bbc-a582-584b5bd8****`, and the `RSAES_OAEP_SHA_1` encryption algorithm.
//
// @param request - AsymmetricEncryptRequest
//
// @return AsymmetricEncryptResponse
func (client *Client) AsymmetricEncrypt(request *AsymmetricEncryptRequest) (_result *AsymmetricEncryptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Generates a digital signature by using an asymmetric CMK.
//
// Description:
//
// ### Precautions
//
// - For information about the access policies that are required for a RAM user or RAM role to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation can be called through a shared gateway or a dedicated gateway. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//   - Shared gateway: You can access KMS over the Internet or a VPC. This method requires you to enable Internet access. For more information, see [Access keys in a KMS instance over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
//   - Dedicated gateway: You can access KMS using the private endpoint of KMS (`<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com`).
//
// ### QPS limits
//
// - Shared gateway: This operation is limited to 200 queries per second (QPS) for each user. If the limit is exceeded, API calls are throttled, which may affect your business. We recommend that you call this operation at a reasonable rate.
//
// - Dedicated gateway: The QPS for each user is limited by the performance specifications of your KMS instance. For more information, see [Performance metrics](https://help.aliyun.com/document_detail/480120.html).
//
// ### Description
//
// This operation supports only asymmetric keys for which the **Usage*	- parameter is set to **SIGN/VERIFY**. The following table describes the supported signature algorithms.
//
// | KeySpec   | Algorithm            | Description                                                |
//
// | --------- | -------------------- | ---------------------------------------------------------- |
//
// | RSA_2048 | RSA_PSS_SHA_256   | RSASSA-PSS using SHA-256 and MGF1 with SHA-256             |
//
// | RSA_2048 | RSA_PKCS1_SHA_256 | RSASSA-PKCS1-v1_5 using SHA-256                           |
//
// | RSA_3072 | RSA_PSS_SHA_256   | RSASSA-PSS using SHA-256 and MGF1 with SHA-256             |
//
// | RSA_3072 | RSA_PKCS1_SHA_256 | RSASSA-PKCS1-v1_5 using SHA-256                           |
//
// | EC_P256  | ECDSA_SHA_256      | ECDSA on the P-256 Curve(secp256r1) with a SHA-256 digest  |
//
// | EC_P256K | ECDSA_SHA_256      | ECDSA on the P-256K Curve(secp256k1) with a SHA-256 digest |
//
// | EC_SM2   | SM2DSA               | SM2 elliptic curve digital signature algorithm             |
//
// > According to the GB/T 32918.2 standard "Information security technology - SM2 elliptic curve public key cryptography - Part 2: Digital signature algorithm", when you calculate an SM2 signature, the value of the **Digest*	- parameter is not the SM3 hash value of the original message. Instead, the value is the SM3 hash value of the result of concatenating Z(A) and M. M is the original message to be signed. Z(A) is the hash value of user A, as defined in GB/T 32918.2.
//
// This topic provides an example of how to use an asymmetric key with the key ID `5c438b18-05be-40ad-b6c2-3be6752c****` and the key version ID `2ab1a983-7072-4bbc-a582-584b5bd8****` to sign the digest `ZOyIygCyaOW6GjVnihtTFtIS9PNmskdyMlNKiuy****=` using the `RSA_PSS_SHA_256` signature algorithm.
//
// @param request - AsymmetricSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AsymmetricSignResponse
func (client *Client) AsymmetricSignWithOptions(request *AsymmetricSignRequest, runtime *dara.RuntimeOptions) (_result *AsymmetricSignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a digital signature by using an asymmetric CMK.
//
// Description:
//
// ### Precautions
//
// - For information about the access policies that are required for a RAM user or RAM role to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation can be called through a shared gateway or a dedicated gateway. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//   - Shared gateway: You can access KMS over the Internet or a VPC. This method requires you to enable Internet access. For more information, see [Access keys in a KMS instance over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
//   - Dedicated gateway: You can access KMS using the private endpoint of KMS (`<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com`).
//
// ### QPS limits
//
// - Shared gateway: This operation is limited to 200 queries per second (QPS) for each user. If the limit is exceeded, API calls are throttled, which may affect your business. We recommend that you call this operation at a reasonable rate.
//
// - Dedicated gateway: The QPS for each user is limited by the performance specifications of your KMS instance. For more information, see [Performance metrics](https://help.aliyun.com/document_detail/480120.html).
//
// ### Description
//
// This operation supports only asymmetric keys for which the **Usage*	- parameter is set to **SIGN/VERIFY**. The following table describes the supported signature algorithms.
//
// | KeySpec   | Algorithm            | Description                                                |
//
// | --------- | -------------------- | ---------------------------------------------------------- |
//
// | RSA_2048 | RSA_PSS_SHA_256   | RSASSA-PSS using SHA-256 and MGF1 with SHA-256             |
//
// | RSA_2048 | RSA_PKCS1_SHA_256 | RSASSA-PKCS1-v1_5 using SHA-256                           |
//
// | RSA_3072 | RSA_PSS_SHA_256   | RSASSA-PSS using SHA-256 and MGF1 with SHA-256             |
//
// | RSA_3072 | RSA_PKCS1_SHA_256 | RSASSA-PKCS1-v1_5 using SHA-256                           |
//
// | EC_P256  | ECDSA_SHA_256      | ECDSA on the P-256 Curve(secp256r1) with a SHA-256 digest  |
//
// | EC_P256K | ECDSA_SHA_256      | ECDSA on the P-256K Curve(secp256k1) with a SHA-256 digest |
//
// | EC_SM2   | SM2DSA               | SM2 elliptic curve digital signature algorithm             |
//
// > According to the GB/T 32918.2 standard "Information security technology - SM2 elliptic curve public key cryptography - Part 2: Digital signature algorithm", when you calculate an SM2 signature, the value of the **Digest*	- parameter is not the SM3 hash value of the original message. Instead, the value is the SM3 hash value of the result of concatenating Z(A) and M. M is the original message to be signed. Z(A) is the hash value of user A, as defined in GB/T 32918.2.
//
// This topic provides an example of how to use an asymmetric key with the key ID `5c438b18-05be-40ad-b6c2-3be6752c****` and the key version ID `2ab1a983-7072-4bbc-a582-584b5bd8****` to sign the digest `ZOyIygCyaOW6GjVnihtTFtIS9PNmskdyMlNKiuy****=` using the `RSA_PSS_SHA_256` signature algorithm.
//
// @param request - AsymmetricSignRequest
//
// @return AsymmetricSignResponse
func (client *Client) AsymmetricSign(request *AsymmetricSignRequest) (_result *AsymmetricSignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Verifies a digital signature by using the public key of an asymmetric CMK.
//
// Description:
//
// ### Precautions
//
// - For information about the access policy required for a RAM user or RAM role to call this API operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation can be called through a shared gateway or a dedicated gateway. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//   - Shared gateway: You can access KMS over the Internet or through a VPC. To access KMS over the Internet, you must enable the public endpoint. For more information, see [Access KMS instances over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
//   - Dedicated gateway: You can access KMS using the private endpoint of KMS (`<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com`).
//
// ### QPS limits
//
// - If you use a shared gateway, the queries per second (QPS) limit for this operation is 200 for a single user. If you exceed this limit, API calls are throttled, which may impact your business. We recommend that you manage your call frequency to stay within the QPS limit.
//
// - If you use a dedicated gateway, the QPS limit for this operation for a single user is determined by the computing performance specifications of your KMS instance. For more information, see [Performance metrics](https://help.aliyun.com/document_detail/480120.html).
//
// ### Description
//
// This operation supports only asymmetric keys for which the **Usage*	- parameter is set to **SIGN/VERIFY**. The following table lists the supported signature algorithms.
//
// | KeySpec   | Algorithm            | Description                                                |
//
// | --------- | -------------------- | ---------------------------------------------------------- |
//
// | RSA_2048 | RSA_PSS_SHA_256   | RSASSA-PSS using SHA-256 and MGF1 with SHA-256             |
//
// | RSA_2048 | RSA_PKCS1_SHA_256 | RSASSA-PKCS1-v1_5 using SHA-256                           |
//
// | RSA_3072 | RSA_PSS_SHA_256   | RSASSA-PSS using SHA-256 and MGF1 with SHA-256             |
//
// | RSA_3072 | RSA_PKCS1_SHA_256 | RSASSA-PKCS1-v1_5 using SHA-256                           |
//
// | EC_P256  | ECDSA_SHA_256      | ECDSA on the P-256 Curve(secp256r1) with a SHA-256 digest  |
//
// | EC_P256K | ECDSA_SHA_256      | ECDSA on the P-256K Curve(secp256k1) with a SHA-256 digest |
//
// | EC_SM2   | SM2DSA               | SM2 elliptic curve digital signature algorithm             |
//
// > In accordance with the GBT32918 standard, when an SM2 signature is calculated, the value of the **Digest*	- parameter is not the SM3 hash value of the original message. Instead, the value is the SM3 hash value of the result generated by concatenating Z(A) and M. In this formula, M is the original message to be signed, and Z(A) is the hash value of user A as defined in GBT32918.
//
// This topic provides an example of how to use an asymmetric key with the key ID \\`5c438b18-05be-40ad-b6c2-3be6752c\\*\\*\\*\\*\\` and the key version ID \\`2ab1a983-7072-4bbc-a582-584b5bd8\\*\\*\\*\\*\\` to verify the signature \\`M2CceNZH00ZgL9ED/ZHFp21YRAvYeZHknJUc207OCZ0N9wNn9As4z2bON3FF3je+1Nu+2+/8Zj50HpMTpzYpMp2R93cYmACCmhaYoKydxylbyGzJR8y9likZRCrkD38lRoS40aBBvv/6iRKzQuo9EGYVcel36cMNg00VmYNBy3pa1rwg3gA4l3cy6kjayZja1WGPkVhrVKsrJMdbpl0ApLjXKuD8rw1n1XLCwCUEL5eLPljTZaAveqdOFQOiZnZEGI27qIiZe7I1fN8tcz6anS/gTM7xRKE++5egEvRWlTQQTJeApnPSiUPA+8ZykNdelQsOQh5SrGoyI4A5pq\\*\\*\\*\\*==\\` for the digest \\`ZOyIygCyaOW6GjVnihtTFtIS9PNmskdyMlNKiuyjfzw=\\` using the RSA_PSS_SHA_256 signature algorithm.
//
// @param request - AsymmetricVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AsymmetricVerifyResponse
func (client *Client) AsymmetricVerifyWithOptions(request *AsymmetricVerifyRequest, runtime *dara.RuntimeOptions) (_result *AsymmetricVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies a digital signature by using the public key of an asymmetric CMK.
//
// Description:
//
// ### Precautions
//
// - For information about the access policy required for a RAM user or RAM role to call this API operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation can be called through a shared gateway or a dedicated gateway. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//   - Shared gateway: You can access KMS over the Internet or through a VPC. To access KMS over the Internet, you must enable the public endpoint. For more information, see [Access KMS instances over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
//   - Dedicated gateway: You can access KMS using the private endpoint of KMS (`<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com`).
//
// ### QPS limits
//
// - If you use a shared gateway, the queries per second (QPS) limit for this operation is 200 for a single user. If you exceed this limit, API calls are throttled, which may impact your business. We recommend that you manage your call frequency to stay within the QPS limit.
//
// - If you use a dedicated gateway, the QPS limit for this operation for a single user is determined by the computing performance specifications of your KMS instance. For more information, see [Performance metrics](https://help.aliyun.com/document_detail/480120.html).
//
// ### Description
//
// This operation supports only asymmetric keys for which the **Usage*	- parameter is set to **SIGN/VERIFY**. The following table lists the supported signature algorithms.
//
// | KeySpec   | Algorithm            | Description                                                |
//
// | --------- | -------------------- | ---------------------------------------------------------- |
//
// | RSA_2048 | RSA_PSS_SHA_256   | RSASSA-PSS using SHA-256 and MGF1 with SHA-256             |
//
// | RSA_2048 | RSA_PKCS1_SHA_256 | RSASSA-PKCS1-v1_5 using SHA-256                           |
//
// | RSA_3072 | RSA_PSS_SHA_256   | RSASSA-PSS using SHA-256 and MGF1 with SHA-256             |
//
// | RSA_3072 | RSA_PKCS1_SHA_256 | RSASSA-PKCS1-v1_5 using SHA-256                           |
//
// | EC_P256  | ECDSA_SHA_256      | ECDSA on the P-256 Curve(secp256r1) with a SHA-256 digest  |
//
// | EC_P256K | ECDSA_SHA_256      | ECDSA on the P-256K Curve(secp256k1) with a SHA-256 digest |
//
// | EC_SM2   | SM2DSA               | SM2 elliptic curve digital signature algorithm             |
//
// > In accordance with the GBT32918 standard, when an SM2 signature is calculated, the value of the **Digest*	- parameter is not the SM3 hash value of the original message. Instead, the value is the SM3 hash value of the result generated by concatenating Z(A) and M. In this formula, M is the original message to be signed, and Z(A) is the hash value of user A as defined in GBT32918.
//
// This topic provides an example of how to use an asymmetric key with the key ID \\`5c438b18-05be-40ad-b6c2-3be6752c\\*\\*\\*\\*\\` and the key version ID \\`2ab1a983-7072-4bbc-a582-584b5bd8\\*\\*\\*\\*\\` to verify the signature \\`M2CceNZH00ZgL9ED/ZHFp21YRAvYeZHknJUc207OCZ0N9wNn9As4z2bON3FF3je+1Nu+2+/8Zj50HpMTpzYpMp2R93cYmACCmhaYoKydxylbyGzJR8y9likZRCrkD38lRoS40aBBvv/6iRKzQuo9EGYVcel36cMNg00VmYNBy3pa1rwg3gA4l3cy6kjayZja1WGPkVhrVKsrJMdbpl0ApLjXKuD8rw1n1XLCwCUEL5eLPljTZaAveqdOFQOiZnZEGI27qIiZe7I1fN8tcz6anS/gTM7xRKE++5egEvRWlTQQTJeApnPSiUPA+8ZykNdelQsOQh5SrGoyI4A5pq\\*\\*\\*\\*==\\` for the digest \\`ZOyIygCyaOW6GjVnihtTFtIS9PNmskdyMlNKiuyjfzw=\\` using the RSA_PSS_SHA_256 signature algorithm.
//
// @param request - AsymmetricVerifyRequest
//
// @return AsymmetricVerifyResponse
func (client *Client) AsymmetricVerify(request *AsymmetricVerifyRequest) (_result *AsymmetricVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AsymmetricVerifyResponse{}
	_body, _err := client.AsymmetricVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels the deletion task of a CMK.
//
// Description:
//
// If the deletion task of a CMK is canceled, the CMK returns to the Enabled state.
//
// @param request - CancelKeyDeletionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelKeyDeletionResponse
func (client *Client) CancelKeyDeletionWithOptions(request *CancelKeyDeletionRequest, runtime *dara.RuntimeOptions) (_result *CancelKeyDeletionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the deletion task of a CMK.
//
// Description:
//
// If the deletion task of a CMK is canceled, the CMK returns to the Enabled state.
//
// @param request - CancelKeyDeletionRequest
//
// @return CancelKeyDeletionResponse
func (client *Client) CancelKeyDeletion(request *CancelKeyDeletionRequest) (_result *CancelKeyDeletionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) ConnectKmsInstanceWithOptions(request *ConnectKmsInstanceRequest, runtime *dara.RuntimeOptions) (_result *ConnectKmsInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ConnectKmsInstanceResponse
func (client *Client) ConnectKmsInstance(request *ConnectKmsInstanceRequest) (_result *ConnectKmsInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConnectKmsInstanceResponse{}
	_body, _err := client.ConnectKmsInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an alias for a key.
//
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
func (client *Client) CreateAliasWithOptions(request *CreateAliasRequest, runtime *dara.RuntimeOptions) (_result *CreateAliasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an alias for a key.
//
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
// @return CreateAliasResponse
func (client *Client) CreateAlias(request *CreateAliasRequest) (_result *CreateAliasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) CreateApplicationAccessPointWithOptions(request *CreateApplicationAccessPointRequest, runtime *dara.RuntimeOptions) (_result *CreateApplicationAccessPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateApplicationAccessPointResponse
func (client *Client) CreateApplicationAccessPoint(request *CreateApplicationAccessPointRequest) (_result *CreateApplicationAccessPointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApplicationAccessPointResponse{}
	_body, _err := client.CreateApplicationAccessPointWithOptions(request, runtime)
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
func (client *Client) CreateClientKeyWithOptions(request *CreateClientKeyRequest, runtime *dara.RuntimeOptions) (_result *CreateClientKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateClientKeyResponse
func (client *Client) CreateClientKey(request *CreateClientKeyRequest) (_result *CreateClientKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Creates a customer master key (CMK) for envelope encryption, digital signatures, or other cryptographic operations.
//
// Description:
//
// - For information about the access policies required for a RAM user or RAM role to call this OpenAPI operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - Alibaba Cloud Key Management Service (KMS) supports common specifications for symmetric and asymmetric keys. For more information, see [Key management types and key specifications](https://help.aliyun.com/document_detail/480161.html).
//
// @param request - CreateKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKeyResponse
func (client *Client) CreateKeyWithOptions(request *CreateKeyRequest, runtime *dara.RuntimeOptions) (_result *CreateKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a customer master key (CMK) for envelope encryption, digital signatures, or other cryptographic operations.
//
// Description:
//
// - For information about the access policies required for a RAM user or RAM role to call this OpenAPI operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - Alibaba Cloud Key Management Service (KMS) supports common specifications for symmetric and asymmetric keys. For more information, see [Key management types and key specifications](https://help.aliyun.com/document_detail/480161.html).
//
// @param request - CreateKeyRequest
//
// @return CreateKeyResponse
func (client *Client) CreateKey(request *CreateKeyRequest) (_result *CreateKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Creates a version for a customer master key (CMK).
//
// Description:
//
// - You can create a version only for an asymmetric CMK that is in the Enabled state. You can call the [CreateKey](https://help.aliyun.com/document_detail/28947.html) operation to create an asymmetric CMK and the [DescribeKey](https://help.aliyun.com/document_detail/28952.html) operation to query the status of the CMK. The status is specified by the KeyState parameter.
//
// - The minimum interval for creating a version of the same CMK is seven days. You can call the [DescribeKey](https://help.aliyun.com/document_detail/28952.html) operation to query the time when the last version of a CMK was created. The time is specified by the LastRotationDate parameter.
//
// - If a CMK is in a private key store, you cannot create a version for the CMK.
//
// - You can create a maximum of 50 versions for a CMK in the same region.
//
// You can create a version for the CMK whose ID is `0b30658a-ed1a-4922-b8f7-a673ca9c****` by using the parameter settings provided in this topic.
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - CreateKeyVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKeyVersionResponse
func (client *Client) CreateKeyVersionWithOptions(request *CreateKeyVersionRequest, runtime *dara.RuntimeOptions) (_result *CreateKeyVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a version for a customer master key (CMK).
//
// Description:
//
// - You can create a version only for an asymmetric CMK that is in the Enabled state. You can call the [CreateKey](https://help.aliyun.com/document_detail/28947.html) operation to create an asymmetric CMK and the [DescribeKey](https://help.aliyun.com/document_detail/28952.html) operation to query the status of the CMK. The status is specified by the KeyState parameter.
//
// - The minimum interval for creating a version of the same CMK is seven days. You can call the [DescribeKey](https://help.aliyun.com/document_detail/28952.html) operation to query the time when the last version of a CMK was created. The time is specified by the LastRotationDate parameter.
//
// - If a CMK is in a private key store, you cannot create a version for the CMK.
//
// - You can create a maximum of 50 versions for a CMK in the same region.
//
// You can create a version for the CMK whose ID is `0b30658a-ed1a-4922-b8f7-a673ca9c****` by using the parameter settings provided in this topic.
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - CreateKeyVersionRequest
//
// @return CreateKeyVersionResponse
func (client *Client) CreateKeyVersion(request *CreateKeyVersionRequest) (_result *CreateKeyVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Creates a network access rule to configure the private IP addresses or private CIDR blocks that are allowed to access a Key Management Service (KMS) instance.
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
func (client *Client) CreateNetworkRuleWithOptions(request *CreateNetworkRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateNetworkRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a network access rule to configure the private IP addresses or private CIDR blocks that are allowed to access a Key Management Service (KMS) instance.
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
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) CreatePolicyWithOptions(request *CreatePolicyRequest, runtime *dara.RuntimeOptions) (_result *CreatePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreatePolicyResponse
func (client *Client) CreatePolicy(request *CreatePolicyRequest) (_result *CreatePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Creates a secret and stores its initial version.
//
// Description:
//
// - For information about the access policy required for a Resource Access Management (RAM) user or RAM role to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - Specify the secret name, the secret value for the initial version, and the version number. The initial version is marked with the ACSCurrent stage label.
//
// - Key Management Service (KMS) uses the key that you specify to encrypt the secret value. The key and the secret must be in the same KMS instance. The key must be a symmetric key.
//
//	> KMS encrypts the secret value of each version. Metadata such as the secret name, version number, and version stage labels are not encrypted.
//
// - Before you encrypt the secret value, you must have the `kms:GenerateDataKey` permission on the key.
//
// This topic provides an example of how to create an RDS secret. The secret is named `mydbconninfo`. The `VersionId` of the initial version is `v1`. The `SecretData` is `{"Accounts":[{"AccountName":"user1","AccountPassword":"****"}]}`.
//
// @param tmpReq - CreateSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecretResponse
func (client *Client) CreateSecretWithOptions(tmpReq *CreateSecretRequest, runtime *dara.RuntimeOptions) (_result *CreateSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a secret and stores its initial version.
//
// Description:
//
// - For information about the access policy required for a Resource Access Management (RAM) user or RAM role to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - Specify the secret name, the secret value for the initial version, and the version number. The initial version is marked with the ACSCurrent stage label.
//
// - Key Management Service (KMS) uses the key that you specify to encrypt the secret value. The key and the secret must be in the same KMS instance. The key must be a symmetric key.
//
//	> KMS encrypts the secret value of each version. Metadata such as the secret name, version number, and version stage labels are not encrypted.
//
// - Before you encrypt the secret value, you must have the `kms:GenerateDataKey` permission on the key.
//
// This topic provides an example of how to create an RDS secret. The secret is named `mydbconninfo`. The `VersionId` of the initial version is `v1`. The `SecretData` is `{"Accounts":[{"AccountName":"user1","AccountPassword":"****"}]}`.
//
// @param request - CreateSecretRequest
//
// @return CreateSecretResponse
func (client *Client) CreateSecret(request *CreateSecretRequest) (_result *CreateSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Decrypts ciphertext that was encrypted by using a CMK.
//
// Description:
//
// ### Precautions
//
// - For information about the access policy required for a RAM user or RAM role to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation can be called through a shared gateway or a dedicated gateway. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//   - Shared gateway: You can access KMS over the Internet or using a VPC domain name. To use a shared gateway, you must enable Internet access. For more information, see [Access keys in a KMS instance over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
//   - Dedicated gateway: You can access KMS using the private endpoint of KMS (`<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com`).
//
// ### QPS limits
//
// - Shared gateway: The queries per second (QPS) limit for a single user for this operation is 1,000. If this limit is exceeded, API calls are throttled, which may affect your business. We recommend that you plan your calls accordingly.
//
// - Dedicated gateway: The QPS limit for a single user for this operation is subject to the performance specifications of your KMS instance. For more information, see [Performance metrics](https://help.aliyun.com/document_detail/480120.html).
//
// @param tmpReq - DecryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DecryptResponse
func (client *Client) DecryptWithOptions(tmpReq *DecryptRequest, runtime *dara.RuntimeOptions) (_result *DecryptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.Recipient) {
		query["Recipient"] = request.Recipient
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Decrypts ciphertext that was encrypted by using a CMK.
//
// Description:
//
// ### Precautions
//
// - For information about the access policy required for a RAM user or RAM role to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation can be called through a shared gateway or a dedicated gateway. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//   - Shared gateway: You can access KMS over the Internet or using a VPC domain name. To use a shared gateway, you must enable Internet access. For more information, see [Access keys in a KMS instance over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
//   - Dedicated gateway: You can access KMS using the private endpoint of KMS (`<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com`).
//
// ### QPS limits
//
// - Shared gateway: The queries per second (QPS) limit for a single user for this operation is 1,000. If this limit is exceeded, API calls are throttled, which may affect your business. We recommend that you plan your calls accordingly.
//
// - Dedicated gateway: The QPS limit for a single user for this operation is subject to the performance specifications of your KMS instance. For more information, see [Performance metrics](https://help.aliyun.com/document_detail/480120.html).
//
// @param request - DecryptRequest
//
// @return DecryptResponse
func (client *Client) Decrypt(request *DecryptRequest) (_result *DecryptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DecryptResponse{}
	_body, _err := client.DecryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an alias.
//
// @param request - DeleteAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAliasResponse
func (client *Client) DeleteAliasWithOptions(request *DeleteAliasRequest, runtime *dara.RuntimeOptions) (_result *DeleteAliasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an alias.
//
// @param request - DeleteAliasRequest
//
// @return DeleteAliasResponse
func (client *Client) DeleteAlias(request *DeleteAliasRequest) (_result *DeleteAliasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// - For information about the access policy that a RAM user or RAM role requires to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - Before you delete an AAP, make sure that it is no longer in use. If you delete an AAP that is in use, related applications cannot access KMS. Proceed with caution.
//
// @param request - DeleteApplicationAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationAccessPointResponse
func (client *Client) DeleteApplicationAccessPointWithOptions(request *DeleteApplicationAccessPointRequest, runtime *dara.RuntimeOptions) (_result *DeleteApplicationAccessPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - For information about the access policy that a RAM user or RAM role requires to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - Before you delete an AAP, make sure that it is no longer in use. If you delete an AAP that is in use, related applications cannot access KMS. Proceed with caution.
//
// @param request - DeleteApplicationAccessPointRequest
//
// @return DeleteApplicationAccessPointResponse
func (client *Client) DeleteApplicationAccessPoint(request *DeleteApplicationAccessPointRequest) (_result *DeleteApplicationAccessPointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApplicationAccessPointResponse{}
	_body, _err := client.DeleteApplicationAccessPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a client key.
//
// Description:
//
// - For information about the access policy required for a RAM user or RAM role to call this API operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - Before you delete a ClientKey, make sure that it is no longer in use. Deleting a ClientKey that is in use prevents related applications from accessing KMS. Proceed with caution.
//
// @param request - DeleteClientKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClientKeyResponse
func (client *Client) DeleteClientKeyWithOptions(request *DeleteClientKeyRequest, runtime *dara.RuntimeOptions) (_result *DeleteClientKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a client key.
//
// Description:
//
// - For information about the access policy required for a RAM user or RAM role to call this API operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - Before you delete a ClientKey, make sure that it is no longer in use. Deleting a ClientKey that is in use prevents related applications from accessing KMS. Proceed with caution.
//
// @param request - DeleteClientKeyRequest
//
// @return DeleteClientKeyResponse
func (client *Client) DeleteClientKey(request *DeleteClientKeyRequest) (_result *DeleteClientKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteClientKeyResponse{}
	_body, _err := client.DeleteClientKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the imported key material from a CMK. After deletion, the CMK enters the PendingImport state until you re-import key material.
//
// Description:
//
// This operation does not delete the CMK that is created by using the key material.
//
// If the CMK is in the PendingDeletion state, the state of the CMK and the scheduled deletion time do not change after you call this operation. If the CMK is not in the PendingDeletion state, the state of the CMK changes to PendingImport after you call this operation.
//
// After you delete the key material, you can upload only the same key material into the CMK.
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - DeleteKeyMaterialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteKeyMaterialResponse
func (client *Client) DeleteKeyMaterialWithOptions(request *DeleteKeyMaterialRequest, runtime *dara.RuntimeOptions) (_result *DeleteKeyMaterialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the imported key material from a CMK. After deletion, the CMK enters the PendingImport state until you re-import key material.
//
// Description:
//
// This operation does not delete the CMK that is created by using the key material.
//
// If the CMK is in the PendingDeletion state, the state of the CMK and the scheduled deletion time do not change after you call this operation. If the CMK is not in the PendingDeletion state, the state of the CMK changes to PendingImport after you call this operation.
//
// After you delete the key material, you can upload only the same key material into the CMK.
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - DeleteKeyMaterialRequest
//
// @return DeleteKeyMaterialResponse
func (client *Client) DeleteKeyMaterial(request *DeleteKeyMaterialRequest) (_result *DeleteKeyMaterialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// - For information about the access policy that is required to call this OpenAPI as a Resource Access Management (RAM) user or RAM role, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - Before deleting the network control rule, ensure that it is not attached to any access policies. Otherwise, related applications cannot access KMS as expected.
//
// @param request - DeleteNetworkRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkRuleResponse
func (client *Client) DeleteNetworkRuleWithOptions(request *DeleteNetworkRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteNetworkRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - For information about the access policy that is required to call this OpenAPI as a Resource Access Management (RAM) user or RAM role, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - Before deleting the network control rule, ensure that it is not attached to any access policies. Otherwise, related applications cannot access KMS as expected.
//
// @param request - DeleteNetworkRuleRequest
//
// @return DeleteNetworkRuleResponse
func (client *Client) DeleteNetworkRule(request *DeleteNetworkRuleRequest) (_result *DeleteNetworkRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) DeletePolicyWithOptions(request *DeletePolicyRequest, runtime *dara.RuntimeOptions) (_result *DeletePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeletePolicyResponse
func (client *Client) DeletePolicy(request *DeletePolicyRequest) (_result *DeletePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePolicyResponse{}
	_body, _err := client.DeletePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a secret.
//
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
func (client *Client) DeleteSecretWithOptions(request *DeleteSecretRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a secret.
//
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
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSecretResponse{}
	_body, _err := client.DeleteSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of Key Management Service (KMS) within your Alibaba Cloud account.
//
// Description:
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountKmsStatusResponse
func (client *Client) DescribeAccountKmsStatusWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeAccountKmsStatusResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccountKmsStatus"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccountKmsStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of Key Management Service (KMS) within your Alibaba Cloud account.
//
// Description:
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @return DescribeAccountKmsStatusResponse
func (client *Client) DescribeAccountKmsStatus() (_result *DescribeAccountKmsStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Retrieves the details of an application access point (AAP).
//
// Description:
//
// For information about the access policy that a Resource Access Management (RAM) user or RAM role must have to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - DescribeApplicationAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApplicationAccessPointResponse
func (client *Client) DescribeApplicationAccessPointWithOptions(request *DescribeApplicationAccessPointRequest, runtime *dara.RuntimeOptions) (_result *DescribeApplicationAccessPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an application access point (AAP).
//
// Description:
//
// For information about the access policy that a Resource Access Management (RAM) user or RAM role must have to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - DescribeApplicationAccessPointRequest
//
// @return DescribeApplicationAccessPointResponse
func (client *Client) DescribeApplicationAccessPoint(request *DescribeApplicationAccessPointRequest) (_result *DescribeApplicationAccessPointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApplicationAccessPointResponse{}
	_body, _err := client.DescribeApplicationAccessPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the metadata of a CMK, such as the key state, usage, and rotation configuration.
//
// Description:
//
// You can query the information about the CMK `05754286-3ba2-4fa6-8d41-4323aca6****` by using parameter settings provided in this topic. The information includes the creator, creation time, status, and deletion protection status of the CMK.
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - DescribeKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKeyResponse
func (client *Client) DescribeKeyWithOptions(request *DescribeKeyRequest, runtime *dara.RuntimeOptions) (_result *DescribeKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metadata of a CMK, such as the key state, usage, and rotation configuration.
//
// Description:
//
// You can query the information about the CMK `05754286-3ba2-4fa6-8d41-4323aca6****` by using parameter settings provided in this topic. The information includes the creator, creation time, status, and deletion protection status of the CMK.
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - DescribeKeyRequest
//
// @return DescribeKeyResponse
func (client *Client) DescribeKey(request *DescribeKeyRequest) (_result *DescribeKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeKeyResponse{}
	_body, _err := client.DescribeKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the metadata of a specific CMK version.
//
// Description:
//
// This topic provides an example on how to query the information about a version of the CMK `1234abcd-12ab-34cd-56ef-12345678****`. The ID of the CMK version is `2ab1a983-7072-4bbc-a582-584b5bd8****`. The response shows that the creation time of the CMK version is `2016-03-25T10:42:40Z`.
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - DescribeKeyVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKeyVersionResponse
func (client *Client) DescribeKeyVersionWithOptions(request *DescribeKeyVersionRequest, runtime *dara.RuntimeOptions) (_result *DescribeKeyVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metadata of a specific CMK version.
//
// Description:
//
// This topic provides an example on how to query the information about a version of the CMK `1234abcd-12ab-34cd-56ef-12345678****`. The ID of the CMK version is `2ab1a983-7072-4bbc-a582-584b5bd8****`. The response shows that the creation time of the CMK version is `2016-03-25T10:42:40Z`.
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - DescribeKeyVersionRequest
//
// @return DescribeKeyVersionResponse
func (client *Client) DescribeKeyVersion(request *DescribeKeyVersionRequest) (_result *DescribeKeyVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Retrieves the details of a network access rule.
//
// Description:
//
// For information about the required access policy for a Resource Access Management (RAM) user or RAM role to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - DescribeNetworkRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetworkRuleResponse
func (client *Client) DescribeNetworkRuleWithOptions(request *DescribeNetworkRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetworkRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a network access rule.
//
// Description:
//
// For information about the required access policy for a Resource Access Management (RAM) user or RAM role to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - DescribeNetworkRuleRequest
//
// @return DescribeNetworkRuleResponse
func (client *Client) DescribeNetworkRule(request *DescribeNetworkRuleRequest) (_result *DescribeNetworkRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Retrieves the details of a permission policy.
//
// Description:
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - DescribePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolicyResponse
func (client *Client) DescribePolicyWithOptions(request *DescribePolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a permission policy.
//
// Description:
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - DescribePolicyRequest
//
// @return DescribePolicyResponse
func (client *Client) DescribePolicy(request *DescribePolicyRequest) (_result *DescribePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Queries the regions where KMS is available.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the regions where KMS is available.
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions() (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the metadata of a secret.
//
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
func (client *Client) DescribeSecretWithOptions(request *DescribeSecretRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metadata of a secret.
//
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
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSecretResponse{}
	_body, _err := client.DescribeSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables a key.
//
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
func (client *Client) DisableKeyWithOptions(request *DisableKeyRequest, runtime *dara.RuntimeOptions) (_result *DisableKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a key.
//
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
	runtime := &dara.RuntimeOptions{}
	_result = &DisableKeyResponse{}
	_body, _err := client.DisableKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a key to encrypt and decrypt data.
//
// @param request - EnableKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableKeyResponse
func (client *Client) EnableKeyWithOptions(request *EnableKeyRequest, runtime *dara.RuntimeOptions) (_result *EnableKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a key to encrypt and decrypt data.
//
// @param request - EnableKeyRequest
//
// @return EnableKeyResponse
func (client *Client) EnableKey(request *EnableKeyRequest) (_result *EnableKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableKeyResponse{}
	_body, _err := client.EnableKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Encrypts plaintext by using a symmetric CMK.
//
// Description:
//
// ### Precautions
//
// - For information about the access policy required to allow a RAM user or RAM role to use this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation is accessible through a shared gateway or a dedicated gateway. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//   - Shared gateway: You can access KMS over the Internet or a VPC. To use this method, you must enable Internet access. For more information, see [Access keys in a KMS instance over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
//   - Dedicated gateway: You can access KMS using the private endpoint of KMS (`<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com`).
//
// ### QPS limits
//
// - When accessed through a shared gateway, the queries per second (QPS) limit for a single user is 1,000. If the limit is exceeded, requests are throttled, which can affect your business. We recommend that you stay within this limit to avoid throttling.
//
// - When accessed through a dedicated gateway, the QPS limit for a single user is subject to the computing performance specifications of your KMS instance. For more information, see [Performance metrics](https://help.aliyun.com/document_detail/480120.html).
//
// ### Description
//
// - KMS encrypts the specified data using the primary version of a specified key.
//
// - You can encrypt a maximum of 6 KB of data, such as an RSA key, a database password, or other sensitive information.
//
// - If you migrate encrypted data from one region to another, you can call the Encrypt operation in the destination region to re-encrypt the plaintext data key from the source region. This generates a new encrypted data key. You can then call the [Decrypt](https://help.aliyun.com/document_detail/28950.html) operation to decrypt this new key in the destination region.
//
// @param tmpReq - EncryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EncryptResponse
func (client *Client) EncryptWithOptions(tmpReq *EncryptRequest, runtime *dara.RuntimeOptions) (_result *EncryptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Encrypts plaintext by using a symmetric CMK.
//
// Description:
//
// ### Precautions
//
// - For information about the access policy required to allow a RAM user or RAM role to use this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation is accessible through a shared gateway or a dedicated gateway. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//   - Shared gateway: You can access KMS over the Internet or a VPC. To use this method, you must enable Internet access. For more information, see [Access keys in a KMS instance over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
//   - Dedicated gateway: You can access KMS using the private endpoint of KMS (`<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com`).
//
// ### QPS limits
//
// - When accessed through a shared gateway, the queries per second (QPS) limit for a single user is 1,000. If the limit is exceeded, requests are throttled, which can affect your business. We recommend that you stay within this limit to avoid throttling.
//
// - When accessed through a dedicated gateway, the QPS limit for a single user is subject to the computing performance specifications of your KMS instance. For more information, see [Performance metrics](https://help.aliyun.com/document_detail/480120.html).
//
// ### Description
//
// - KMS encrypts the specified data using the primary version of a specified key.
//
// - You can encrypt a maximum of 6 KB of data, such as an RSA key, a database password, or other sensitive information.
//
// - If you migrate encrypted data from one region to another, you can call the Encrypt operation in the destination region to re-encrypt the plaintext data key from the source region. This generates a new encrypted data key. You can then call the [Decrypt](https://help.aliyun.com/document_detail/28950.html) operation to decrypt this new key in the destination region.
//
// @param request - EncryptRequest
//
// @return EncryptResponse
func (client *Client) Encrypt(request *EncryptRequest) (_result *EncryptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EncryptResponse{}
	_body, _err := client.EncryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Exports a data key encrypted by a CMK. The data key is re-encrypted by a public key that you specify for secure transmission.
//
// Description:
//
// ### Precautions
//
// - For information about the access policy required for a RAM user or RAM role to use this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation is accessible through a shared gateway or a dedicated gateway. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//   - Shared gateway: You can access KMS using the public endpoint or a VPC endpoint. To use the public endpoint, you must first enable it. For more information, see [Access the key in a KMS instance over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
//   - Dedicated gateway: You can access KMS using the private endpoint of the KMS instance: `<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com`.
//
// ### Description
//
// After you call the [GenerateDataKeyWithoutPlaintext](https://help.aliyun.com/document_detail/134043.html) operation to obtain a data key encrypted by a master key (CMK), you can use the ExportDataKey operation to distribute the data key to other regions or cryptographic modules. The ExportDataKey operation returns the ciphertext of the data key, re-encrypted with the specified public key.
//
// You can import the exported ciphertext into the cryptographic module that holds the corresponding private key. This process lets you securely distribute the data key from KMS to a cryptographic module. After the data key is imported into the cryptographic module, you can use it to encrypt or decrypt data.
//
// @param tmpReq - ExportDataKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportDataKeyResponse
func (client *Client) ExportDataKeyWithOptions(tmpReq *ExportDataKeyRequest, runtime *dara.RuntimeOptions) (_result *ExportDataKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Exports a data key encrypted by a CMK. The data key is re-encrypted by a public key that you specify for secure transmission.
//
// Description:
//
// ### Precautions
//
// - For information about the access policy required for a RAM user or RAM role to use this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation is accessible through a shared gateway or a dedicated gateway. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//   - Shared gateway: You can access KMS using the public endpoint or a VPC endpoint. To use the public endpoint, you must first enable it. For more information, see [Access the key in a KMS instance over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
//   - Dedicated gateway: You can access KMS using the private endpoint of the KMS instance: `<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com`.
//
// ### Description
//
// After you call the [GenerateDataKeyWithoutPlaintext](https://help.aliyun.com/document_detail/134043.html) operation to obtain a data key encrypted by a master key (CMK), you can use the ExportDataKey operation to distribute the data key to other regions or cryptographic modules. The ExportDataKey operation returns the ciphertext of the data key, re-encrypted with the specified public key.
//
// You can import the exported ciphertext into the cryptographic module that holds the corresponding private key. This process lets you securely distribute the data key from KMS to a cryptographic module. After the data key is imported into the cryptographic module, you can use it to encrypt or decrypt data.
//
// @param request - ExportDataKeyRequest
//
// @return ExportDataKeyResponse
func (client *Client) ExportDataKey(request *ExportDataKeyRequest) (_result *ExportDataKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportDataKeyResponse{}
	_body, _err := client.ExportDataKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a random data key, encrypts it by using a CMK and a public key that you specify, and returns both ciphertexts.
//
// Description:
//
// ### Notes
//
// - For more information about the access policy required for a RAM user or RAM role to use this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation is accessible through a shared gateway or a dedicated gateway. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//   - Shared gateway: You can access KMS over the Internet or a VPC. To access KMS over the Internet, you must enable Internet access. For more information, see [Access KMS instances over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
//   - Dedicated gateway: You can access KMS using the private endpoint of KMS (`<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com`).
//
// ### Description
//
// We recommend that you import the data key to a cryptographic module for data encryption and data decryption as follows:
//
// 1\\. Call the GenerateAndExportDataKey operation to obtain the data key encrypted by a KMS key and a specified public key.
//
// 2\\. Save the ciphertext of the data key that is encrypted by the KMS key to KMS or a storage service, such as ApsaraDB, for key backup and recovery.
//
// 3\\. Import the ciphertext of the data key that is encrypted by the public key to the cryptographic module that contains the corresponding private key. This process distributes the key from KMS to the cryptographic module. You can then use the data key to encrypt and decrypt data.
//
// > The KMS key that you specify in the request is used only to encrypt the data key and is not used to generate the data key. KMS does not record or store the randomly generated data key. You are responsible for recording the data key or its ciphertext.
//
// @param tmpReq - GenerateAndExportDataKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateAndExportDataKeyResponse
func (client *Client) GenerateAndExportDataKeyWithOptions(tmpReq *GenerateAndExportDataKeyRequest, runtime *dara.RuntimeOptions) (_result *GenerateAndExportDataKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a random data key, encrypts it by using a CMK and a public key that you specify, and returns both ciphertexts.
//
// Description:
//
// ### Notes
//
// - For more information about the access policy required for a RAM user or RAM role to use this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation is accessible through a shared gateway or a dedicated gateway. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//   - Shared gateway: You can access KMS over the Internet or a VPC. To access KMS over the Internet, you must enable Internet access. For more information, see [Access KMS instances over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
//   - Dedicated gateway: You can access KMS using the private endpoint of KMS (`<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com`).
//
// ### Description
//
// We recommend that you import the data key to a cryptographic module for data encryption and data decryption as follows:
//
// 1\\. Call the GenerateAndExportDataKey operation to obtain the data key encrypted by a KMS key and a specified public key.
//
// 2\\. Save the ciphertext of the data key that is encrypted by the KMS key to KMS or a storage service, such as ApsaraDB, for key backup and recovery.
//
// 3\\. Import the ciphertext of the data key that is encrypted by the public key to the cryptographic module that contains the corresponding private key. This process distributes the key from KMS to the cryptographic module. You can then use the data key to encrypt and decrypt data.
//
// > The KMS key that you specify in the request is used only to encrypt the data key and is not used to generate the data key. KMS does not record or store the randomly generated data key. You are responsible for recording the data key or its ciphertext.
//
// @param request - GenerateAndExportDataKeyRequest
//
// @return GenerateAndExportDataKeyResponse
func (client *Client) GenerateAndExportDataKey(request *GenerateAndExportDataKeyRequest) (_result *GenerateAndExportDataKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Generates a random data key for envelope encryption. The data key is returned in both plaintext and ciphertext forms.
//
// Description:
//
// - For information about the permissions that are required to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation can be called using a shared gateway or a dedicated gateway. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//   - Shared gateway: You can access KMS over the Internet or a VPC. To access KMS over the Internet, you must enable the public endpoint. For more information, see [Access keys in a KMS instance over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
//   - Dedicated gateway: You can access KMS using the private endpoint of KMS (`<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com`).
//
// ### QPS limits
//
// - If you use a shared gateway to call this operation, the queries per second (QPS) limit for a single user is 1,000. If the limit is exceeded, API calls are throttled. This may affect your business. We recommend that you call this operation at a reasonable rate.
//
// - If you use a dedicated gateway to call this operation, the QPS limit for a single user is based on the computing performance of your KMS instance. For more information, see [Performance metrics](https://help.aliyun.com/document_detail/480120.html).
//
// ### Description
//
// This operation generates a random data key, encrypts the data key using the specified customer master key (CMK), and returns the plaintext and ciphertext of the data key. You can use the plaintext of the data key to encrypt data locally and outside of KMS. When you store the encrypted data, you must also store the ciphertext of the data key. You can obtain the plaintext of the data key from the Plaintext field and the ciphertext of the data key from the CiphertextBlob field in the response.
//
// The CMK that you specify in the request is used only to encrypt the data key. It is not involved in the generation of the data key. KMS does not record or store the randomly generated data key. You are responsible for the persistence of the ciphertext of the data key.
//
// We recommend that you perform the following steps to encrypt data locally:
//
// 1\\. Call the GenerateDataKey operation to obtain a data key for data encryption.
//
// 2\\. Use the plaintext of the data key returned in the Plaintext field of the response to encrypt data locally. Then, clear the plaintext of the data key from memory.
//
// 3\\. Store the ciphertext of the data key returned in the CiphertextBlob field of the response together with the encrypted data.
//
// To decrypt data locally:
//
// - Call the [Decrypt](https://help.aliyun.com/document_detail/28950.html) operation to decrypt the stored ciphertext of the data key. This operation returns the plaintext of the data key.
//
// - Use the plaintext of the data key to decrypt data locally. Then, clear the plaintext of the data key from memory.
//
// This topic provides an example of how to generate a random data key for a key with the ID `key-hzz630494463ejqjx****`.
//
// @param tmpReq - GenerateDataKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateDataKeyResponse
func (client *Client) GenerateDataKeyWithOptions(tmpReq *GenerateDataKeyRequest, runtime *dara.RuntimeOptions) (_result *GenerateDataKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.Recipient) {
		query["Recipient"] = request.Recipient
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a random data key for envelope encryption. The data key is returned in both plaintext and ciphertext forms.
//
// Description:
//
// - For information about the permissions that are required to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation can be called using a shared gateway or a dedicated gateway. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//   - Shared gateway: You can access KMS over the Internet or a VPC. To access KMS over the Internet, you must enable the public endpoint. For more information, see [Access keys in a KMS instance over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
//   - Dedicated gateway: You can access KMS using the private endpoint of KMS (`<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com`).
//
// ### QPS limits
//
// - If you use a shared gateway to call this operation, the queries per second (QPS) limit for a single user is 1,000. If the limit is exceeded, API calls are throttled. This may affect your business. We recommend that you call this operation at a reasonable rate.
//
// - If you use a dedicated gateway to call this operation, the QPS limit for a single user is based on the computing performance of your KMS instance. For more information, see [Performance metrics](https://help.aliyun.com/document_detail/480120.html).
//
// ### Description
//
// This operation generates a random data key, encrypts the data key using the specified customer master key (CMK), and returns the plaintext and ciphertext of the data key. You can use the plaintext of the data key to encrypt data locally and outside of KMS. When you store the encrypted data, you must also store the ciphertext of the data key. You can obtain the plaintext of the data key from the Plaintext field and the ciphertext of the data key from the CiphertextBlob field in the response.
//
// The CMK that you specify in the request is used only to encrypt the data key. It is not involved in the generation of the data key. KMS does not record or store the randomly generated data key. You are responsible for the persistence of the ciphertext of the data key.
//
// We recommend that you perform the following steps to encrypt data locally:
//
// 1\\. Call the GenerateDataKey operation to obtain a data key for data encryption.
//
// 2\\. Use the plaintext of the data key returned in the Plaintext field of the response to encrypt data locally. Then, clear the plaintext of the data key from memory.
//
// 3\\. Store the ciphertext of the data key returned in the CiphertextBlob field of the response together with the encrypted data.
//
// To decrypt data locally:
//
// - Call the [Decrypt](https://help.aliyun.com/document_detail/28950.html) operation to decrypt the stored ciphertext of the data key. This operation returns the plaintext of the data key.
//
// - Use the plaintext of the data key to decrypt data locally. Then, clear the plaintext of the data key from memory.
//
// This topic provides an example of how to generate a random data key for a key with the ID `key-hzz630494463ejqjx****`.
//
// @param request - GenerateDataKeyRequest
//
// @return GenerateDataKeyResponse
func (client *Client) GenerateDataKey(request *GenerateDataKeyRequest) (_result *GenerateDataKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Generates a random data key in only ciphertext form, without the plaintext copy.
//
// Description:
//
// ### Precautions
//
// - For information about the access policy that a RAM user or RAM role needs to use this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation is accessible through a shared gateway or a dedicated gateway. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//   - Shared gateway: Access KMS over the Internet or through a VPC domain name. This method requires Internet access to be enabled. For more information, see [Access keys in a KMS instance over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
//   - Dedicated gateway: Access KMS through a KMS private endpoint (`<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com`).
//
// ### QPS limits
//
// - Calls through a shared gateway: The queries per second (QPS) limit for a single user is 1,000. If you exceed this limit, requests are throttled, which may affect your business. We recommend that you stay within this limit.
//
// - Calls through a dedicated gateway: The QPS limit for a single user depends on the computing performance of your KMS instance. For more information, see [Performance metrics](https://help.aliyun.com/document_detail/480120.html).
//
// ### Details
//
// This operation generates a random data key, encrypts it with a specified symmetric customer master key (CMK), and returns the ciphertext of the data key. This operation provides the same features as [GenerateDataKey](https://help.aliyun.com/document_detail/28948.html). The only difference is that this operation does not return the plaintext of the data key.
//
// The CMK that you specify in the request is used only to encrypt the data key. It is not used to generate the data key. KMS does not record or store the randomly generated data key.
//
// > - This operation is suitable for systems that do not need to immediately use the data key for data encryption. When encryption is required, the system calls the [Decrypt](https://help.aliyun.com/document_detail/28950.html) API to decrypt the ciphertext of the data key.
//
// >
//
// > - This operation is also suitable for distributed systems with different trust levels. For example, your system stores data in different partitions based on a defined policy. A module pre-creates these data partitions and generates a unique data key for each one. After this module initializes the control plane, it acts as a key distributor and does not produce or consume data. When data plane modules produce and consume data, they first retrieve the ciphertext of the data key for a partition. They then decrypt the ciphertext and use the plaintext data key to encrypt or decrypt data. Finally, they purge the plaintext data key from memory. In such a system, the key distributor does not need to access the plaintext of the data key. It only requires the \\`GenerateDataKeyWithoutPlaintext\\` permission for the relevant CMK. Data producers and consumers do not need to generate new data keys. They only require the \\`Decrypt\\` permission for the relevant CMK.
//
// @param tmpReq - GenerateDataKeyWithoutPlaintextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateDataKeyWithoutPlaintextResponse
func (client *Client) GenerateDataKeyWithoutPlaintextWithOptions(tmpReq *GenerateDataKeyWithoutPlaintextRequest, runtime *dara.RuntimeOptions) (_result *GenerateDataKeyWithoutPlaintextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a random data key in only ciphertext form, without the plaintext copy.
//
// Description:
//
// ### Precautions
//
// - For information about the access policy that a RAM user or RAM role needs to use this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation is accessible through a shared gateway or a dedicated gateway. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//   - Shared gateway: Access KMS over the Internet or through a VPC domain name. This method requires Internet access to be enabled. For more information, see [Access keys in a KMS instance over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
//   - Dedicated gateway: Access KMS through a KMS private endpoint (`<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com`).
//
// ### QPS limits
//
// - Calls through a shared gateway: The queries per second (QPS) limit for a single user is 1,000. If you exceed this limit, requests are throttled, which may affect your business. We recommend that you stay within this limit.
//
// - Calls through a dedicated gateway: The QPS limit for a single user depends on the computing performance of your KMS instance. For more information, see [Performance metrics](https://help.aliyun.com/document_detail/480120.html).
//
// ### Details
//
// This operation generates a random data key, encrypts it with a specified symmetric customer master key (CMK), and returns the ciphertext of the data key. This operation provides the same features as [GenerateDataKey](https://help.aliyun.com/document_detail/28948.html). The only difference is that this operation does not return the plaintext of the data key.
//
// The CMK that you specify in the request is used only to encrypt the data key. It is not used to generate the data key. KMS does not record or store the randomly generated data key.
//
// > - This operation is suitable for systems that do not need to immediately use the data key for data encryption. When encryption is required, the system calls the [Decrypt](https://help.aliyun.com/document_detail/28950.html) API to decrypt the ciphertext of the data key.
//
// >
//
// > - This operation is also suitable for distributed systems with different trust levels. For example, your system stores data in different partitions based on a defined policy. A module pre-creates these data partitions and generates a unique data key for each one. After this module initializes the control plane, it acts as a key distributor and does not produce or consume data. When data plane modules produce and consume data, they first retrieve the ciphertext of the data key for a partition. They then decrypt the ciphertext and use the plaintext data key to encrypt or decrypt data. Finally, they purge the plaintext data key from memory. In such a system, the key distributor does not need to access the plaintext of the data key. It only requires the \\`GenerateDataKeyWithoutPlaintext\\` permission for the relevant CMK. Data producers and consumers do not need to generate new data keys. They only require the \\`Decrypt\\` permission for the relevant CMK.
//
// @param request - GenerateDataKeyWithoutPlaintextRequest
//
// @return GenerateDataKeyWithoutPlaintextResponse
func (client *Client) GenerateDataKeyWithoutPlaintext(request *GenerateDataKeyWithoutPlaintextRequest) (_result *GenerateDataKeyWithoutPlaintextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateDataKeyWithoutPlaintextResponse{}
	_body, _err := client.GenerateDataKeyWithoutPlaintextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates an HMAC message authentication code for a message by using a specified key.
//
// Description:
//
// For details about the access policy required when a RAM user or RAM role invokes this operation, refer to access control.
//
// This operation can be invoked through a shared gateway or a dedicated gateway. For more information, refer to Alibaba Cloud SDK.
//
// - Shared gateway: Access KMS through a public or VPC endpoint. This method requires you to enable the public network access switch. For more information, refer to accessing keys in a KMS instance over the Internet.
//
// - Dedicated gateway: Access KMS through a KMS private endpoint (<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com).
//
// @param request - GenerateMacRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateMacResponse
func (client *Client) GenerateMacWithOptions(request *GenerateMacRequest, runtime *dara.RuntimeOptions) (_result *GenerateMacResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.Message) {
		query["Message"] = request.Message
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateMac"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateMacResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates an HMAC message authentication code for a message by using a specified key.
//
// Description:
//
// For details about the access policy required when a RAM user or RAM role invokes this operation, refer to access control.
//
// This operation can be invoked through a shared gateway or a dedicated gateway. For more information, refer to Alibaba Cloud SDK.
//
// - Shared gateway: Access KMS through a public or VPC endpoint. This method requires you to enable the public network access switch. For more information, refer to accessing keys in a KMS instance over the Internet.
//
// - Dedicated gateway: Access KMS through a KMS private endpoint (<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com).
//
// @param request - GenerateMacRequest
//
// @return GenerateMacResponse
func (client *Client) GenerateMac(request *GenerateMacRequest) (_result *GenerateMacResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateMacResponse{}
	_body, _err := client.GenerateMacWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves information about a client key.
//
// Description:
//
// For information about the access policy required for a RAM user or RAM role to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - GetClientKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClientKeyResponse
func (client *Client) GetClientKeyWithOptions(request *GetClientKeyRequest, runtime *dara.RuntimeOptions) (_result *GetClientKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about a client key.
//
// Description:
//
// For information about the access policy required for a RAM user or RAM role to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - GetClientKeyRequest
//
// @return GetClientKeyResponse
func (client *Client) GetClientKey(request *GetClientKeyRequest) (_result *GetClientKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Queries the default KMS instance in a specified region.
//
// Description:
//
// - For information about the access policy that is required to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This API is for users who migrate from KMS 1.0 to KMS 3.0. After the migration is complete, if you create an Asset without specifying a KMS instance, the Asset is created in the default KMS instance.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDefaultKmsInstanceResponse
func (client *Client) GetDefaultKmsInstanceWithOptions(runtime *dara.RuntimeOptions) (_result *GetDefaultKmsInstanceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetDefaultKmsInstance"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDefaultKmsInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the default KMS instance in a specified region.
//
// Description:
//
// - For information about the access policy that is required to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This API is for users who migrate from KMS 1.0 to KMS 3.0. After the migration is complete, if you create an Asset without specifying a KMS instance, the Asset is created in the default KMS instance.
//
// @return GetDefaultKmsInstanceResponse
func (client *Client) GetDefaultKmsInstance() (_result *GetDefaultKmsInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDefaultKmsInstanceResponse{}
	_body, _err := client.GetDefaultKmsInstanceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the key policy of a CMK in a KMS instance.
//
// Description:
//
// - For more information about the access policy required for a Resource Access Management (RAM) user or RAM role to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - Because the key policy name can only be set to default, you must set the PolicyName parameter to default when you query the key policy. Otherwise, a `Not Found` error is returned.
//
// @param request - GetKeyPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKeyPolicyResponse
func (client *Client) GetKeyPolicyWithOptions(request *GetKeyPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetKeyPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the key policy of a CMK in a KMS instance.
//
// Description:
//
// - For more information about the access policy required for a Resource Access Management (RAM) user or RAM role to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - Because the key policy name can only be set to default, you must set the PolicyName parameter to default when you query the key policy. Otherwise, a `Not Found` error is returned.
//
// @param request - GetKeyPolicyRequest
//
// @return GetKeyPolicyResponse
func (client *Client) GetKeyPolicy(request *GetKeyPolicyRequest) (_result *GetKeyPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Retrieves the details of a KMS instance.
//
// Description:
//
// Refer to [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html) for the access policy required to call this OpenAPI as a RAM user or RAM role.
//
// @param request - GetKmsInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKmsInstanceResponse
func (client *Client) GetKmsInstanceWithOptions(request *GetKmsInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetKmsInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a KMS instance.
//
// Description:
//
// Refer to [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html) for the access policy required to call this OpenAPI as a RAM user or RAM role.
//
// @param request - GetKmsInstanceRequest
//
// @return GetKmsInstanceResponse
func (client *Client) GetKmsInstance(request *GetKmsInstanceRequest) (_result *GetKmsInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Queries the quota usage and limits for a KMS instance.
//
// @param request - GetKmsInstanceQuotaInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKmsInstanceQuotaInfosResponse
func (client *Client) GetKmsInstanceQuotaInfosWithOptions(request *GetKmsInstanceQuotaInfosRequest, runtime *dara.RuntimeOptions) (_result *GetKmsInstanceQuotaInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KmsInstanceId) {
		query["KmsInstanceId"] = request.KmsInstanceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetKmsInstanceQuotaInfos"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetKmsInstanceQuotaInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quota usage and limits for a KMS instance.
//
// @param request - GetKmsInstanceQuotaInfosRequest
//
// @return GetKmsInstanceQuotaInfosResponse
func (client *Client) GetKmsInstanceQuotaInfos(request *GetKmsInstanceQuotaInfosRequest) (_result *GetKmsInstanceQuotaInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetKmsInstanceQuotaInfosResponse{}
	_body, _err := client.GetKmsInstanceQuotaInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the plaintext of a managed data key (DK).
//
// Description:
//
// For information about the access policy that must be granted to a RAM user or RAM role to call this operation, refer to access control.
//
// This operation can be called through a shared gateway. For more information, refer to Alibaba Cloud SDK.
//
// - Shared gateway: Access KMS through public or VPC endpoints.
//
// @param request - GetManagedDataKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetManagedDataKeyResponse
func (client *Client) GetManagedDataKeyWithOptions(request *GetManagedDataKeyRequest, runtime *dara.RuntimeOptions) (_result *GetManagedDataKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataKeyName) {
		query["DataKeyName"] = request.DataKeyName
	}

	if !dara.IsNil(request.DataKeyVersionId) {
		query["DataKeyVersionId"] = request.DataKeyVersionId
	}

	if !dara.IsNil(request.UseLatest) {
		query["UseLatest"] = request.UseLatest
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetManagedDataKey"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetManagedDataKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the plaintext of a managed data key (DK).
//
// Description:
//
// For information about the access policy that must be granted to a RAM user or RAM role to call this operation, refer to access control.
//
// This operation can be called through a shared gateway. For more information, refer to Alibaba Cloud SDK.
//
// - Shared gateway: Access KMS through public or VPC endpoints.
//
// @param request - GetManagedDataKeyRequest
//
// @return GetManagedDataKeyResponse
func (client *Client) GetManagedDataKey(request *GetManagedDataKeyRequest) (_result *GetManagedDataKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetManagedDataKeyResponse{}
	_body, _err := client.GetManagedDataKeyWithOptions(request, runtime)
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
func (client *Client) GetParametersForImportWithOptions(request *GetParametersForImportRequest, runtime *dara.RuntimeOptions) (_result *GetParametersForImportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetParametersForImportResponse
func (client *Client) GetParametersForImport(request *GetParametersForImportRequest) (_result *GetParametersForImportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetParametersForImportResponse{}
	_body, _err := client.GetParametersForImportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the public key of an asymmetric key. You can use the public key to encrypt data or verify a signature on your device.
//
// Description:
//
// - For more information about the access policy required for a RAM user or RAM role to call this OpenAPI operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation supports both shared gateways and dedicated gateways. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//   - Shared gateway: You can access KMS over the Internet or using a VPC domain name. If you access KMS over the Internet, you must enable Internet access. For more information, see [Access keys in a KMS instance over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
//   - Dedicated gateway: You can access KMS using the private endpoint of KMS (`<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com`).
//
// @param request - GetPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPublicKeyResponse
func (client *Client) GetPublicKeyWithOptions(request *GetPublicKeyRequest, runtime *dara.RuntimeOptions) (_result *GetPublicKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the public key of an asymmetric key. You can use the public key to encrypt data or verify a signature on your device.
//
// Description:
//
// - For more information about the access policy required for a RAM user or RAM role to call this OpenAPI operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation supports both shared gateways and dedicated gateways. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//   - Shared gateway: You can access KMS over the Internet or using a VPC domain name. If you access KMS over the Internet, you must enable Internet access. For more information, see [Access keys in a KMS instance over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
//   - Dedicated gateway: You can access KMS using the private endpoint of KMS (`<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com`).
//
// @param request - GetPublicKeyRequest
//
// @return GetPublicKeyResponse
func (client *Client) GetPublicKey(request *GetPublicKeyRequest) (_result *GetPublicKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPublicKeyResponse{}
	_body, _err := client.GetPublicKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a random password string.
//
// Description:
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - GetRandomPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRandomPasswordResponse
func (client *Client) GetRandomPasswordWithOptions(request *GetRandomPasswordRequest, runtime *dara.RuntimeOptions) (_result *GetRandomPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a random password string.
//
// Description:
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - GetRandomPasswordRequest
//
// @return GetRandomPasswordResponse
func (client *Client) GetRandomPassword(request *GetRandomPasswordRequest) (_result *GetRandomPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Queries the access policy of a specified secret in a KMS instance.
//
// Description:
//
// - For information about the access policy required for a RAM user or RAM role to call this OpenAPI, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - A credential policy name can be set only to default. Therefore, you must set the PolicyName parameter to default when you query the credential policy. Otherwise, a `Not Found` error is returned.
//
// @param request - GetSecretPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecretPolicyResponse
func (client *Client) GetSecretPolicyWithOptions(request *GetSecretPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetSecretPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access policy of a specified secret in a KMS instance.
//
// Description:
//
// - For information about the access policy required for a RAM user or RAM role to call this OpenAPI, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - A credential policy name can be set only to default. Therefore, you must set the PolicyName parameter to default when you query the credential policy. Otherwise, a `Not Found` error is returned.
//
// @param request - GetSecretPolicyRequest
//
// @return GetSecretPolicyResponse
func (client *Client) GetSecretPolicy(request *GetSecretPolicyRequest) (_result *GetSecretPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Retrieve the 
//
// Description:
//
// - For details about the access policy that must be granted to a Resource Access Management (RAM) user or RAM role to invoke this OpenAPI operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - If you do not specify a version number or version status, Key Management Service (KMS) returns the credential value of the version marked as ACSCurrent by default.
//
// - If a customer-managed key is used to protect the credential value, the caller must also have the `kms:Decrypt` permission on the corresponding master key.
//
// This topic provides a sample request to retrieve the credential value of a credential named `secret001`. The returned result shows that the credential value `SecretData` is `testdata1`.
//
// @param request - GetSecretValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecretValueResponse
func (client *Client) GetSecretValueWithOptions(request *GetSecretValueRequest, runtime *dara.RuntimeOptions) (_result *GetSecretValueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the credential value.
//
// Description:
//
// - For details about the access policy that must be granted to a Resource Access Management (RAM) user or RAM role to invoke this OpenAPI operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - If you do not specify a version number or version status, Key Management Service (KMS) returns the credential value of the version marked as ACSCurrent by default.
//
// - If a customer-managed key is used to protect the credential value, the caller must also have the `kms:Decrypt` permission on the corresponding master key.
//
// This topic provides a sample request to retrieve the credential value of a credential named `secret001`. The returned result shows that the credential value `SecretData` is `testdata1`.
//
// @param request - GetSecretValueRequest
//
// @return GetSecretValueResponse
func (client *Client) GetSecretValue(request *GetSecretValueRequest) (_result *GetSecretValueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Imports externally generated key material into a CMK whose origin is EXTERNAL.
//
// Description:
//
// When you call [CreateKey](https://help.aliyun.com/document_detail/28947.html) to create a CMK, you can set **Origin*	- to **EXTERNAL*	- to specify that the key material comes from an external source. Use this operation to import the key material into such a CMK.
//
// - To view the CMK **Origin**, see [DescribeKey](https://help.aliyun.com/document_detail/28952.html).
//
// - Before importing key material, call [GetParametersForImport](https://help.aliyun.com/document_detail/68621.html) to obtain the parameters required for the import, including the public key and import token.
//
// > 	- For a CMK of type **Aliyun_AES_256**, the key material must be 256 bits. For a CMK of type **Aliyun_SM4**, the key material must be 128 bits.
//
// >
//
// > 	- You can set the expiration time for the key material, or you can set it to never expire.
//
// >
//
// > 	- You can reimport the key material and reset the expiration time for the specified CMK at any time, but the same key material must be imported.
//
// >
//
// > 	- After the imported key material expires or is deleted, the specified CMK becomes unavailable until the same key material is imported again.
//
// >
//
// > 	- The same key material can be imported into multiple CMKs, but data or data keys encrypted by one CMK cannot be decrypted by another CMK.
//
// For more information about the access policy required for a RAM user or RAM role to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - ImportKeyMaterialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportKeyMaterialResponse
func (client *Client) ImportKeyMaterialWithOptions(request *ImportKeyMaterialRequest, runtime *dara.RuntimeOptions) (_result *ImportKeyMaterialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports externally generated key material into a CMK whose origin is EXTERNAL.
//
// Description:
//
// When you call [CreateKey](https://help.aliyun.com/document_detail/28947.html) to create a CMK, you can set **Origin*	- to **EXTERNAL*	- to specify that the key material comes from an external source. Use this operation to import the key material into such a CMK.
//
// - To view the CMK **Origin**, see [DescribeKey](https://help.aliyun.com/document_detail/28952.html).
//
// - Before importing key material, call [GetParametersForImport](https://help.aliyun.com/document_detail/68621.html) to obtain the parameters required for the import, including the public key and import token.
//
// > 	- For a CMK of type **Aliyun_AES_256**, the key material must be 256 bits. For a CMK of type **Aliyun_SM4**, the key material must be 128 bits.
//
// >
//
// > 	- You can set the expiration time for the key material, or you can set it to never expire.
//
// >
//
// > 	- You can reimport the key material and reset the expiration time for the specified CMK at any time, but the same key material must be imported.
//
// >
//
// > 	- After the imported key material expires or is deleted, the specified CMK becomes unavailable until the same key material is imported again.
//
// >
//
// > 	- The same key material can be imported into multiple CMKs, but data or data keys encrypted by one CMK cannot be decrypted by another CMK.
//
// For more information about the access policy required for a RAM user or RAM role to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - ImportKeyMaterialRequest
//
// @return ImportKeyMaterialResponse
func (client *Client) ImportKeyMaterial(request *ImportKeyMaterialRequest) (_result *ImportKeyMaterialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Description:
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - ListAliasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAliasesResponse
func (client *Client) ListAliasesWithOptions(request *ListAliasesRequest, runtime *dara.RuntimeOptions) (_result *ListAliasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
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
// Description:
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - ListAliasesRequest
//
// @return ListAliasesResponse
func (client *Client) ListAliases(request *ListAliasesRequest) (_result *ListAliasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAliasesResponse{}
	_body, _err := client.ListAliasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all aliases that are bound to a key.
//
// @param request - ListAliasesByKeyIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAliasesByKeyIdResponse
func (client *Client) ListAliasesByKeyIdWithOptions(request *ListAliasesByKeyIdRequest, runtime *dara.RuntimeOptions) (_result *ListAliasesByKeyIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all aliases that are bound to a key.
//
// @param request - ListAliasesByKeyIdRequest
//
// @return ListAliasesByKeyIdResponse
func (client *Client) ListAliasesByKeyId(request *ListAliasesByKeyIdRequest) (_result *ListAliasesByKeyIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Queries all application access points (AAPs) in the current region.
//
// Description:
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - ListApplicationAccessPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationAccessPointsResponse
func (client *Client) ListApplicationAccessPointsWithOptions(request *ListApplicationAccessPointsRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationAccessPointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all application access points (AAPs) in the current region.
//
// Description:
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - ListApplicationAccessPointsRequest
//
// @return ListApplicationAccessPointsResponse
func (client *Client) ListApplicationAccessPoints(request *ListApplicationAccessPointsRequest) (_result *ListApplicationAccessPointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApplicationAccessPointsResponse{}
	_body, _err := client.ListApplicationAccessPointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all client keys within an AAP.
//
// Description:
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - ListClientKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClientKeysResponse
func (client *Client) ListClientKeysWithOptions(request *ListClientKeysRequest, runtime *dara.RuntimeOptions) (_result *ListClientKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all client keys within an AAP.
//
// Description:
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - ListClientKeysRequest
//
// @return ListClientKeysResponse
func (client *Client) ListClientKeys(request *ListClientKeysRequest) (_result *ListClientKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Description:
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - ListKeyVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKeyVersionsResponse
func (client *Client) ListKeyVersionsWithOptions(request *ListKeyVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListKeyVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
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
// Description:
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - ListKeyVersionsRequest
//
// @return ListKeyVersionsResponse
func (client *Client) ListKeyVersions(request *ListKeyVersionsRequest) (_result *ListKeyVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Queries the IDs and ARNs of all CMKs in the current region.
//
// Description:
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - ListKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKeysResponse
func (client *Client) ListKeysWithOptions(request *ListKeysRequest, runtime *dara.RuntimeOptions) (_result *ListKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IDs and ARNs of all CMKs in the current region.
//
// Description:
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - ListKeysRequest
//
// @return ListKeysResponse
func (client *Client) ListKeys(request *ListKeysRequest) (_result *ListKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Queries all KMS instances in the current region.
//
// Description:
//
// For more information about the access policy required for a RAM user or RAM role to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - ListKmsInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKmsInstancesResponse
func (client *Client) ListKmsInstancesWithOptions(request *ListKmsInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListKmsInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all KMS instances in the current region.
//
// Description:
//
// For more information about the access policy required for a RAM user or RAM role to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - ListKmsInstancesRequest
//
// @return ListKmsInstancesResponse
func (client *Client) ListKmsInstances(request *ListKmsInstancesRequest) (_result *ListKmsInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Lists the key versions of a specified managed data key (DK).
//
// Description:
//
// For details about the access policy required when a RAM user or RAM role invokes this operation, refer to access control.
//
// This operation can be invoked through the shared gateway. For more information, refer to Alibaba Cloud SDK.
//
// - Shared gateway: Access KMS through public or VPC endpoints.
//
// @param request - ListManagedDataKeyVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListManagedDataKeyVersionsResponse
func (client *Client) ListManagedDataKeyVersionsWithOptions(request *ListManagedDataKeyVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListManagedDataKeyVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataKeyName) {
		query["DataKeyName"] = request.DataKeyName
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
		Action:      dara.String("ListManagedDataKeyVersions"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListManagedDataKeyVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the key versions of a specified managed data key (DK).
//
// Description:
//
// For details about the access policy required when a RAM user or RAM role invokes this operation, refer to access control.
//
// This operation can be invoked through the shared gateway. For more information, refer to Alibaba Cloud SDK.
//
// - Shared gateway: Access KMS through public or VPC endpoints.
//
// @param request - ListManagedDataKeyVersionsRequest
//
// @return ListManagedDataKeyVersionsResponse
func (client *Client) ListManagedDataKeyVersions(request *ListManagedDataKeyVersionsRequest) (_result *ListManagedDataKeyVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListManagedDataKeyVersionsResponse{}
	_body, _err := client.ListManagedDataKeyVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all network access rules in the current region.
//
// Description:
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - ListNetworkRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetworkRulesResponse
func (client *Client) ListNetworkRulesWithOptions(request *ListNetworkRulesRequest, runtime *dara.RuntimeOptions) (_result *ListNetworkRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all network access rules in the current region.
//
// Description:
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - ListNetworkRulesRequest
//
// @return ListNetworkRulesResponse
func (client *Client) ListNetworkRules(request *ListNetworkRulesRequest) (_result *ListNetworkRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Queries all permission policies in the current region.
//
// Description:
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - ListPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPoliciesResponse
func (client *Client) ListPoliciesWithOptions(request *ListPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListPoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all permission policies in the current region.
//
// Description:
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - ListPoliciesRequest
//
// @return ListPoliciesResponse
func (client *Client) ListPolicies(request *ListPoliciesRequest) (_result *ListPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPoliciesResponse{}
	_body, _err := client.ListPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags of a customer master key (CMK).
//
// Description:
//
// Request format: KeyId="string"
//
// @param request - ListResourceTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceTagsResponse
func (client *Client) ListResourceTagsWithOptions(request *ListResourceTagsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags of a customer master key (CMK).
//
// Description:
//
// Request format: KeyId="string"
//
// @param request - ListResourceTagsRequest
//
// @return ListResourceTagsResponse
func (client *Client) ListResourceTags(request *ListResourceTagsRequest) (_result *ListResourceTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourceTagsResponse{}
	_body, _err := client.ListResourceTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all version IDs and stage labels of a specified secret.
//
// Description:
//
// - For more information about the access policy required for a RAM user or RAM role to call this OpenAPI, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - The version information does not include the secret value. By default, this operation returns only the secret versions that are marked with a version stage.
//
// @param request - ListSecretVersionIdsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecretVersionIdsResponse
func (client *Client) ListSecretVersionIdsWithOptions(request *ListSecretVersionIdsRequest, runtime *dara.RuntimeOptions) (_result *ListSecretVersionIdsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all version IDs and stage labels of a specified secret.
//
// Description:
//
// - For more information about the access policy required for a RAM user or RAM role to call this OpenAPI, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - The version information does not include the secret value. By default, this operation returns only the secret versions that are marked with a version stage.
//
// @param request - ListSecretVersionIdsRequest
//
// @return ListSecretVersionIdsResponse
func (client *Client) ListSecretVersionIds(request *ListSecretVersionIdsRequest) (_result *ListSecretVersionIdsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSecretVersionIdsResponse{}
	_body, _err := client.ListSecretVersionIdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all secrets in the current region.
//
// Description:
//
// - To call this operation, the RAM user or RAM role must be granted the required policy. For more information, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation returns only secret metadata, not the secret values.
//
// This example shows how to query secrets created by the current user in the current region. `PageNumber` is set to `1` and `PageSize` is set to `2`, returning metadata for two secrets.
//
// @param request - ListSecretsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecretsResponse
func (client *Client) ListSecretsWithOptions(request *ListSecretsRequest, runtime *dara.RuntimeOptions) (_result *ListSecretsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all secrets in the current region.
//
// Description:
//
// - To call this operation, the RAM user or RAM role must be granted the required policy. For more information, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation returns only secret metadata, not the secret values.
//
// This example shows how to query secrets created by the current user in the current region. `PageNumber` is set to `1` and `PageSize` is set to `2`, returning metadata for two secrets.
//
// @param request - ListSecretsRequest
//
// @return ListSecretsResponse
func (client *Client) ListSecrets(request *ListSecretsRequest) (_result *ListSecretsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Lists the tags that are bound to a key or a secret.
//
// Description:
//
// For more information about the access policy required for a RAM user or RAM role to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the tags that are bound to a key or a secret.
//
// Description:
//
// For more information about the access policy required for a RAM user or RAM role to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Activates Key Management Service (KMS) for your Alibaba Cloud account.
//
// Description:
//
// - For more information about the access policies that a RAM user or RAM role needs to call this OpenAPI, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - KMS is a paid service. For more information about billing, see [Billing](https://help.aliyun.com/document_detail/52608.html).
//
// - You can activate the service for an Alibaba Cloud account only once.
//
// - Make sure that your Alibaba Cloud account has completed real-name verification.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenKmsServiceResponse
func (client *Client) OpenKmsServiceWithOptions(runtime *dara.RuntimeOptions) (_result *OpenKmsServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("OpenKmsService"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenKmsServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates Key Management Service (KMS) for your Alibaba Cloud account.
//
// Description:
//
// - For more information about the access policies that a RAM user or RAM role needs to call this OpenAPI, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - KMS is a paid service. For more information about billing, see [Billing](https://help.aliyun.com/document_detail/52608.html).
//
// - You can activate the service for an Alibaba Cloud account only once.
//
// - Make sure that your Alibaba Cloud account has completed real-name verification.
//
// @return OpenKmsServiceResponse
func (client *Client) OpenKmsService() (_result *OpenKmsServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenKmsServiceResponse{}
	_body, _err := client.OpenKmsServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stores a new version of a secret value for a generic secret.
//
// Description:
//
// - For information about the access policy required for a RAM user or RAM role to call this OpenAPI operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation supports only generic secrets. Each generic secret can have a maximum of 10 versions. If the number of versions exceeds the limit, KMS deletes the earliest version.
//
// - By default, the new secret value is marked with ACSCurrent, and the previous version that was marked with ACSCurrent is marked with ACSPrevious. You can specify the VersionStage parameter to overwrite this default behavior.
//
// - This operation stores a new version of a secret value. You cannot use it to modify an existing version of a secret value. You must specify a version number when you store a new version. KMS processes requests based on the following rules:
//
//   - If the version number does not exist in the secret, KMS creates a new version and stores the secret value.
//
//   - If the version number already exists in the secret, KMS compares the secret value in the request with the stored value. If the values are the same, the request is ignored and a success message is returned. This makes the operation idempotent. If the values are different, the request is rejected.
//
// This topic provides an example of how to store a new version of a secret value for the secret named `secret001`. The new version number (`VersionId`) is `v3` and the secret value (`SecretData`) is `importantdata`.
//
// @param request - PutSecretValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutSecretValueResponse
func (client *Client) PutSecretValueWithOptions(request *PutSecretValueRequest, runtime *dara.RuntimeOptions) (_result *PutSecretValueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stores a new version of a secret value for a generic secret.
//
// Description:
//
// - For information about the access policy required for a RAM user or RAM role to call this OpenAPI operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation supports only generic secrets. Each generic secret can have a maximum of 10 versions. If the number of versions exceeds the limit, KMS deletes the earliest version.
//
// - By default, the new secret value is marked with ACSCurrent, and the previous version that was marked with ACSCurrent is marked with ACSPrevious. You can specify the VersionStage parameter to overwrite this default behavior.
//
// - This operation stores a new version of a secret value. You cannot use it to modify an existing version of a secret value. You must specify a version number when you store a new version. KMS processes requests based on the following rules:
//
//   - If the version number does not exist in the secret, KMS creates a new version and stores the secret value.
//
//   - If the version number already exists in the secret, KMS compares the secret value in the request with the stored value. If the values are the same, the request is ignored and a success message is returned. This makes the operation idempotent. If the values are different, the request is rejected.
//
// This topic provides an example of how to store a new version of a secret value for the secret named `secret001`. The new version number (`VersionId`) is `v3` and the secret value (`SecretData`) is `importantdata`.
//
// @param request - PutSecretValueRequest
//
// @return PutSecretValueResponse
func (client *Client) PutSecretValue(request *PutSecretValueRequest) (_result *PutSecretValueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutSecretValueResponse{}
	_body, _err := client.PutSecretValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Re-encrypts ciphertext under a different CMK without exposing the plaintext.
//
// Description:
//
// ### Notes
//
// - For more information about the access policy required to grant a RAM user or RAM role the permission to use this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation is accessible only through a shared gateway, not a dedicated gateway. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//	When using a shared gateway, you access KMS through an Internet or a VPC domain name. This method requires Internet access to be enabled. For more information, see [Access keys in a KMS instance over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
// ### QPS limits
//
// This operation is accessible only through a shared gateway. The single-user queries per second (QPS) limit is 750. If this limit is exceeded, requests are throttled, which may affect your business. We recommend that you stay within the specified limit.
//
// ### Details
//
// You can use the ReEncrypt operation in the following scenarios:
//
// - After a customer master key (CMK) is rotated, you can use the latest key version to re-encrypt data. For more information about automatic key rotation, see [Automatic key rotation](https://help.aliyun.com/document_detail/134270.html).
//
// - You can re-encrypt data by changing the encryption context without changing the master key.
//
// - You can re-encrypt data or a data key that is encrypted by one master key with another master key in KMS.
//
// The ReEncrypt operation requires the following permissions:
//
// - The kms:ReEncryptFrom permission for the source master key.
//
// - The kms:ReEncryptTo permission for the destination master key.
//
// - You can use kms:ReEncrypt\\	- to grant both permissions.
//
// @param tmpReq - ReEncryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReEncryptResponse
func (client *Client) ReEncryptWithOptions(tmpReq *ReEncryptRequest, runtime *dara.RuntimeOptions) (_result *ReEncryptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Re-encrypts ciphertext under a different CMK without exposing the plaintext.
//
// Description:
//
// ### Notes
//
// - For more information about the access policy required to grant a RAM user or RAM role the permission to use this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation is accessible only through a shared gateway, not a dedicated gateway. For more information, see [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/601559.html).
//
//	When using a shared gateway, you access KMS through an Internet or a VPC domain name. This method requires Internet access to be enabled. For more information, see [Access keys in a KMS instance over the Internet](https://help.aliyun.com/document_detail/2856718.html).
//
// ### QPS limits
//
// This operation is accessible only through a shared gateway. The single-user queries per second (QPS) limit is 750. If this limit is exceeded, requests are throttled, which may affect your business. We recommend that you stay within the specified limit.
//
// ### Details
//
// You can use the ReEncrypt operation in the following scenarios:
//
// - After a customer master key (CMK) is rotated, you can use the latest key version to re-encrypt data. For more information about automatic key rotation, see [Automatic key rotation](https://help.aliyun.com/document_detail/134270.html).
//
// - You can re-encrypt data by changing the encryption context without changing the master key.
//
// - You can re-encrypt data or a data key that is encrypted by one master key with another master key in KMS.
//
// The ReEncrypt operation requires the following permissions:
//
// - The kms:ReEncryptFrom permission for the source master key.
//
// - The kms:ReEncryptTo permission for the destination master key.
//
// - You can use kms:ReEncrypt\\	- to grant both permissions.
//
// @param request - ReEncryptRequest
//
// @return ReEncryptResponse
func (client *Client) ReEncrypt(request *ReEncryptRequest) (_result *ReEncryptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReEncryptResponse{}
	_body, _err := client.ReEncryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a pay-as-you-go KMS instance.
//
// Description:
//
// - For information about the access policy that is required to allow a RAM user or RAM role to call this OpenAPI operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - Subscription instances cannot be manually released. You can only unsubscribe from them under specific conditions. To unsubscribe, go to the Expenses and Costs page in the console and choose Unsubscribe. For more information, see [Unsubscription policy](https://help.aliyun.com/document_detail/600418.html).
//
// - After you release an instance, all resources in the instance are also released. Resources that are encrypted using keys in the instance cannot be decrypted, and credentials cannot be retrieved. Before you release an instance, make sure that no data is encrypted by the keys in the instance and no services call the credentials. This prevents service interruptions.
//
// - If your instance is a software key management instance, we recommend that you back up the resources of the instance before you release it. The backed-up resources can be recovered. For more information, see [Backup management](https://help.aliyun.com/document_detail/2357488.html).
//
// - The billing epoch is daily. Therefore, after you release a pay-as-you-go instance, the bill for the previous day is pushed before 12:00 on the next day.
//
// - Before you release a KMS instance, we recommend that you check whether deletion protection is enabled for the instance in the console. If deletion protection is enabled, disable it in the console before you release the instance. For more information, see [Manage a KMS instance](https://help.aliyun.com/document_detail/604735.html).
//
// @param request - ReleaseKmsInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseKmsInstanceResponse
func (client *Client) ReleaseKmsInstanceWithOptions(request *ReleaseKmsInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReleaseKmsInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a pay-as-you-go KMS instance.
//
// Description:
//
// - For information about the access policy that is required to allow a RAM user or RAM role to call this OpenAPI operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - Subscription instances cannot be manually released. You can only unsubscribe from them under specific conditions. To unsubscribe, go to the Expenses and Costs page in the console and choose Unsubscribe. For more information, see [Unsubscription policy](https://help.aliyun.com/document_detail/600418.html).
//
// - After you release an instance, all resources in the instance are also released. Resources that are encrypted using keys in the instance cannot be decrypted, and credentials cannot be retrieved. Before you release an instance, make sure that no data is encrypted by the keys in the instance and no services call the credentials. This prevents service interruptions.
//
// - If your instance is a software key management instance, we recommend that you back up the resources of the instance before you release it. The backed-up resources can be recovered. For more information, see [Backup management](https://help.aliyun.com/document_detail/2357488.html).
//
// - The billing epoch is daily. Therefore, after you release a pay-as-you-go instance, the bill for the previous day is pushed before 12:00 on the next day.
//
// - Before you release a KMS instance, we recommend that you check whether deletion protection is enabled for the instance in the console. If deletion protection is enabled, disable it in the console before you release the instance. For more information, see [Manage a KMS instance](https://help.aliyun.com/document_detail/604735.html).
//
// @param request - ReleaseKmsInstanceRequest
//
// @return ReleaseKmsInstanceResponse
func (client *Client) ReleaseKmsInstance(request *ReleaseKmsInstanceRequest) (_result *ReleaseKmsInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseKmsInstanceResponse{}
	_body, _err := client.ReleaseKmsInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restores a deleted secret.
//
// Description:
//
// You can only use this operation to restore a deleted secret that is within its recovery period. If you set **ForceDeleteWithoutRecovery*	- to **true*	- when you delete the secret, you cannot restore it.
//
// @param request - RestoreSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreSecretResponse
func (client *Client) RestoreSecretWithOptions(request *RestoreSecretRequest, runtime *dara.RuntimeOptions) (_result *RestoreSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores a deleted secret.
//
// Description:
//
// You can only use this operation to restore a deleted secret that is within its recovery period. If you set **ForceDeleteWithoutRecovery*	- to **true*	- when you delete the secret, you cannot restore it.
//
// @param request - RestoreSecretRequest
//
// @return RestoreSecretResponse
func (client *Client) RestoreSecret(request *RestoreSecretRequest) (_result *RestoreSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestoreSecretResponse{}
	_body, _err := client.RestoreSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Immediately rotates a secret.
//
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
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - RotateSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RotateSecretResponse
func (client *Client) RotateSecretWithOptions(request *RotateSecretRequest, runtime *dara.RuntimeOptions) (_result *RotateSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Immediately rotates a secret.
//
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
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - RotateSecretRequest
//
// @return RotateSecretResponse
func (client *Client) RotateSecret(request *RotateSecretRequest) (_result *RotateSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RotateSecretResponse{}
	_body, _err := client.RotateSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified customer master key (CMK).
//
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
func (client *Client) ScheduleKeyDeletionWithOptions(request *ScheduleKeyDeletionRequest, runtime *dara.RuntimeOptions) (_result *ScheduleKeyDeletionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified customer master key (CMK).
//
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
	runtime := &dara.RuntimeOptions{}
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
// - For details about the access policy that must be granted to a Resource Access Management (RAM) user or RAM role to invoke this operation, see [Access control](https://help.aliyun.com/document_detail/2767210.html).
//
// - After you enable deletion protection for a CMK, you cannot delete the CMK. To delete the CMK, disable deletion protection first.
//
// - Before you invoke the SetDeletionProtection operation, make sure that the CMK is not in the PendingDeletion state. You can invoke the [DescribeKey](https://help.aliyun.com/document_detail/28952.html) operation to query the status (KeyState) of the CMK.
//
// @param request - SetDeletionProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDeletionProtectionResponse
func (client *Client) SetDeletionProtectionWithOptions(request *SetDeletionProtectionRequest, runtime *dara.RuntimeOptions) (_result *SetDeletionProtectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.KmsInstanceId) {
		query["KmsInstanceId"] = request.KmsInstanceId
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - For details about the access policy that must be granted to a Resource Access Management (RAM) user or RAM role to invoke this operation, see [Access control](https://help.aliyun.com/document_detail/2767210.html).
//
// - After you enable deletion protection for a CMK, you cannot delete the CMK. To delete the CMK, disable deletion protection first.
//
// - Before you invoke the SetDeletionProtection operation, make sure that the CMK is not in the PendingDeletion state. You can invoke the [DescribeKey](https://help.aliyun.com/document_detail/28952.html) operation to query the status (KeyState) of the CMK.
//
// @param request - SetDeletionProtectionRequest
//
// @return SetDeletionProtectionResponse
func (client *Client) SetDeletionProtection(request *SetDeletionProtectionRequest) (_result *SetDeletionProtectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Sets the key policy for a CMK in a KMS instance.
//
// Description:
//
// - For information about the access policy required for a RAM user or RAM role to call this API operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - For more information about key policies, see [Key policy overview](https://help.aliyun.com/document_detail/2716468.html).
//
// @param request - SetKeyPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetKeyPolicyResponse
func (client *Client) SetKeyPolicyWithOptions(request *SetKeyPolicyRequest, runtime *dara.RuntimeOptions) (_result *SetKeyPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the key policy for a CMK in a KMS instance.
//
// Description:
//
// - For information about the access policy required for a RAM user or RAM role to call this API operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - For more information about key policies, see [Key policy overview](https://help.aliyun.com/document_detail/2716468.html).
//
// @param request - SetKeyPolicyRequest
//
// @return SetKeyPolicyResponse
func (client *Client) SetKeyPolicy(request *SetKeyPolicyRequest) (_result *SetKeyPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Sets the access policy for a secret in a KMS instance.
//
// Description:
//
// - For information about the access policy that a RAM user or RAM role requires to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - For more information about secret policies, see [Secret policy overview](https://help.aliyun.com/document_detail/2716465.html).
//
// @param request - SetSecretPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSecretPolicyResponse
func (client *Client) SetSecretPolicyWithOptions(request *SetSecretPolicyRequest, runtime *dara.RuntimeOptions) (_result *SetSecretPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the access policy for a secret in a KMS instance.
//
// Description:
//
// - For information about the access policy that a RAM user or RAM role requires to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - For more information about secret policies, see [Secret policy overview](https://help.aliyun.com/document_detail/2716465.html).
//
// @param request - SetSecretPolicyRequest
//
// @return SetSecretPolicyResponse
func (client *Client) SetSecretPolicy(request *SetSecretPolicyRequest) (_result *SetSecretPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetSecretPolicyResponse{}
	_body, _err := client.SetSecretPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tags to a CMK, secret, or certificate.
//
// Description:
//
// You can add up to 10 tags to a CMK, secret, or certificate.
//
// In this example, the tags `[{"TagKey":"S1key1","TagValue":"S1val1"},{"TagKey":"S1key2","TagValue":"S2val2"}]` are added to the CMK whose ID is `08c33a6f-4e0a-4a1b-a3fa-7ddf****`.
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - TagResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourceResponse
func (client *Client) TagResourceWithOptions(request *TagResourceRequest, runtime *dara.RuntimeOptions) (_result *TagResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to a CMK, secret, or certificate.
//
// Description:
//
// You can add up to 10 tags to a CMK, secret, or certificate.
//
// In this example, the tags `[{"TagKey":"S1key1","TagValue":"S1val1"},{"TagKey":"S1key2","TagValue":"S2val2"}]` are added to the CMK whose ID is `08c33a6f-4e0a-4a1b-a3fa-7ddf****`.
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - TagResourceRequest
//
// @return TagResourceResponse
func (client *Client) TagResource(request *TagResourceRequest) (_result *TagResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Adds tags to one or more keys or secrets.
//
// Description:
//
// You can add multiple tags to multiple keys or multiple secrets at a time.
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to one or more keys or secrets.
//
// Description:
//
// You can add multiple tags to multiple keys or multiple secrets at a time.
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from a CMK, secret, or certificate.
//
// Description:
//
// For information about the access policy that is required for a RAM user or RAM role to call this OpenAPI operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// This topic provides an example of how to detach tags with the tag keys tagkey1 and tagkey2 from the key with the ID `08c33a6f-4e0a-4a1b-a3fa-7ddf****`.
//
// @param request - UntagResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourceResponse
func (client *Client) UntagResourceWithOptions(request *UntagResourceRequest, runtime *dara.RuntimeOptions) (_result *UntagResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from a CMK, secret, or certificate.
//
// Description:
//
// For information about the access policy that is required for a RAM user or RAM role to call this OpenAPI operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// This topic provides an example of how to detach tags with the tag keys tagkey1 and tagkey2 from the key with the ID `08c33a6f-4e0a-4a1b-a3fa-7ddf****`.
//
// @param request - UntagResourceRequest
//
// @return UntagResourceResponse
func (client *Client) UntagResource(request *UntagResourceRequest) (_result *UntagResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds an existing alias to a different customer master key (CMK) ID.
//
// @param request - UpdateAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAliasResponse
func (client *Client) UpdateAliasWithOptions(request *UpdateAliasRequest, runtime *dara.RuntimeOptions) (_result *UpdateAliasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds an existing alias to a different customer master key (CMK) ID.
//
// @param request - UpdateAliasRequest
//
// @return UpdateAliasResponse
func (client *Client) UpdateAlias(request *UpdateAliasRequest) (_result *UpdateAliasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAliasResponse{}
	_body, _err := client.UpdateAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about an application access point (AAP).
//
// Description:
//
// The update takes effect immediately after an AAP information is updated. Exercise caution when you perform this operation. You can update the description of an AAP and the permission policies that are associated with the AAP. You cannot update the name of the AAP.
//
// @param request - UpdateApplicationAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationAccessPointResponse
func (client *Client) UpdateApplicationAccessPointWithOptions(request *UpdateApplicationAccessPointRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationAccessPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about an application access point (AAP).
//
// Description:
//
// The update takes effect immediately after an AAP information is updated. Exercise caution when you perform this operation. You can update the description of an AAP and the permission policies that are associated with the AAP. You cannot update the name of the AAP.
//
// @param request - UpdateApplicationAccessPointRequest
//
// @return UpdateApplicationAccessPointResponse
func (client *Client) UpdateApplicationAccessPoint(request *UpdateApplicationAccessPointRequest) (_result *UpdateApplicationAccessPointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateApplicationAccessPointResponse{}
	_body, _err := client.UpdateApplicationAccessPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the description of a CMK.
//
// Description:
//
// This operation replaces the description of a customer master key (CMK) with the description that you specify. The original description of the CMK is specified by the Description parameter when you call the [DescribeKey](https://help.aliyun.com/document_detail/28952.html) operation. Use this operation to add, modify, or delete the description of a CMK.
//
// For more information about the access policy required for a RAM user or RAM role to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - UpdateKeyDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKeyDescriptionResponse
func (client *Client) UpdateKeyDescriptionWithOptions(request *UpdateKeyDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateKeyDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the description of a CMK.
//
// Description:
//
// This operation replaces the description of a customer master key (CMK) with the description that you specify. The original description of the CMK is specified by the Description parameter when you call the [DescribeKey](https://help.aliyun.com/document_detail/28952.html) operation. Use this operation to add, modify, or delete the description of a CMK.
//
// For more information about the access policy required for a RAM user or RAM role to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - UpdateKeyDescriptionRequest
//
// @return UpdateKeyDescriptionResponse
func (client *Client) UpdateKeyDescription(request *UpdateKeyDescriptionRequest) (_result *UpdateKeyDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Updates the VPC bindings of a KMS instance.
//
// Description:
//
// - For information about the access policy required for a Resource Access Management (RAM) user or RAM role to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - If your self-managed application is deployed in multiple VPCs in the same region, you can bind the KMS instance to additional VPCs. These VPCs are in addition to the one that you specified when you enabled the instance.
//
//	These VPCs can belong to the same Alibaba Cloud account or different Alibaba Cloud accounts. After the configuration is complete, self-managed applications in these VPCs can access the specified KMS instance.
//
//	> If the VPCs belong to different Alibaba Cloud accounts, you must first configure resource sharing to share the vSwitch resources from those accounts with the Alibaba Cloud account that owns the KMS instance. For more information, see [Access a KMS instance from multiple VPCs in the same region](https://help.aliyun.com/document_detail/2393236.html).
//
// @param request - UpdateKmsInstanceBindVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKmsInstanceBindVpcResponse
func (client *Client) UpdateKmsInstanceBindVpcWithOptions(request *UpdateKmsInstanceBindVpcRequest, runtime *dara.RuntimeOptions) (_result *UpdateKmsInstanceBindVpcResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BindVpcs) {
		query["BindVpcs"] = request.BindVpcs
	}

	if !dara.IsNil(request.KmsInstanceId) {
		query["KmsInstanceId"] = request.KmsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateKmsInstanceBindVpc"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateKmsInstanceBindVpcResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the VPC bindings of a KMS instance.
//
// Description:
//
// - For information about the access policy required for a Resource Access Management (RAM) user or RAM role to call this operation, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - If your self-managed application is deployed in multiple VPCs in the same region, you can bind the KMS instance to additional VPCs. These VPCs are in addition to the one that you specified when you enabled the instance.
//
//	These VPCs can belong to the same Alibaba Cloud account or different Alibaba Cloud accounts. After the configuration is complete, self-managed applications in these VPCs can access the specified KMS instance.
//
//	> If the VPCs belong to different Alibaba Cloud accounts, you must first configure resource sharing to share the vSwitch resources from those accounts with the Alibaba Cloud account that owns the KMS instance. For more information, see [Access a KMS instance from multiple VPCs in the same region](https://help.aliyun.com/document_detail/2393236.html).
//
// @param request - UpdateKmsInstanceBindVpcRequest
//
// @return UpdateKmsInstanceBindVpcResponse
func (client *Client) UpdateKmsInstanceBindVpc(request *UpdateKmsInstanceBindVpcRequest) (_result *UpdateKmsInstanceBindVpcResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Updates a network access rule.
//
// Description:
//
// - For more information about the access policy required for a Resource Access Management (RAM) user or RAM role to call this operation, see [Access control](https://help.aliyun.com/document_detail/2767210.html).
//
// - You can only modify the private IP addresses and description of a network control rule. The name and network type cannot be modified.
//
// - When you update a network control rule, the access policies attached to it are also updated. Proceed with caution.
//
// @param request - UpdateNetworkRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNetworkRuleResponse
func (client *Client) UpdateNetworkRuleWithOptions(request *UpdateNetworkRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateNetworkRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a network access rule.
//
// Description:
//
// - For more information about the access policy required for a Resource Access Management (RAM) user or RAM role to call this operation, see [Access control](https://help.aliyun.com/document_detail/2767210.html).
//
// - You can only modify the private IP addresses and description of a network control rule. The name and network type cannot be modified.
//
// - When you update a network control rule, the access policies attached to it are also updated. Proceed with caution.
//
// @param request - UpdateNetworkRuleRequest
//
// @return UpdateNetworkRuleResponse
func (client *Client) UpdateNetworkRule(request *UpdateNetworkRuleRequest) (_result *UpdateNetworkRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Updates a permission policy.
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
func (client *Client) UpdatePolicyWithOptions(request *UpdatePolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdatePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a permission policy.
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
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePolicyResponse{}
	_body, _err := client.UpdatePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the automatic rotation policy of a CMK.
//
// Description:
//
// When automatic key rotation is enabled, KMS automatically creates a key version after the preset rotation period arrives. In addition, KMS sets the new key version as the primary key version.
//
// An automatic key rotation policy cannot be configured for the following keys:
//
// - Asymmetric key
//
// - Service-managed key
//
// - Bring your own key (BYOK) that is imported into KMS
//
// - Key that is not in the **Enabled*	- state
//
// In this example, automatic key rotation is enabled for a CMK whose ID is `1234abcd-12ab-34cd-56ef-12345678****`. The automatic rotation period is 30 days.
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - UpdateRotationPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRotationPolicyResponse
func (client *Client) UpdateRotationPolicyWithOptions(request *UpdateRotationPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateRotationPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the automatic rotation policy of a CMK.
//
// Description:
//
// When automatic key rotation is enabled, KMS automatically creates a key version after the preset rotation period arrives. In addition, KMS sets the new key version as the primary key version.
//
// An automatic key rotation policy cannot be configured for the following keys:
//
// - Asymmetric key
//
// - Service-managed key
//
// - Bring your own key (BYOK) that is imported into KMS
//
// - Key that is not in the **Enabled*	- state
//
// In this example, automatic key rotation is enabled for a CMK whose ID is `1234abcd-12ab-34cd-56ef-12345678****`. The automatic rotation period is 30 days.
//
// For more information about the access policy required by a RAM user or RAM role to call this API, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// @param request - UpdateRotationPolicyRequest
//
// @return UpdateRotationPolicyResponse
func (client *Client) UpdateRotationPolicy(request *UpdateRotationPolicyRequest) (_result *UpdateRotationPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) UpdateSecretWithOptions(request *UpdateSecretRequest, runtime *dara.RuntimeOptions) (_result *UpdateSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateSecretResponse
func (client *Client) UpdateSecret(request *UpdateSecretRequest) (_result *UpdateSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSecretResponse{}
	_body, _err := client.UpdateSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the rotation policy of a secret.
//
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
func (client *Client) UpdateSecretRotationPolicyWithOptions(request *UpdateSecretRotationPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateSecretRotationPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the rotation policy of a secret.
//
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
// @return UpdateSecretRotationPolicyResponse
func (client *Client) UpdateSecretRotationPolicy(request *UpdateSecretRotationPolicyRequest) (_result *UpdateSecretRotationPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// Moves a version stage label to a different version of a secret.
//
// Description:
//
// - For more information about the access policy that is required to call this operation as a Resource Access Management (RAM) user or RAM role, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation supports only generic secrets. You can perform the following operations:
//
//   - Add a version stage to a specified secret version.
//
//   - Remove a version stage from a specified secret version.
//
//   - Remove a version stage from a specified secret version and attach it to another secret version.
//
// - The total number of version stages for each generic secret cannot exceed 8.
//
// This topic provides an example of how to update the version stage of the secret named `secret001`. In this example, the `ACSCurrent` stage is used to mark version `002`.
//
// @param request - UpdateSecretVersionStageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecretVersionStageResponse
func (client *Client) UpdateSecretVersionStageWithOptions(request *UpdateSecretVersionStageRequest, runtime *dara.RuntimeOptions) (_result *UpdateSecretVersionStageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a version stage label to a different version of a secret.
//
// Description:
//
// - For more information about the access policy that is required to call this operation as a Resource Access Management (RAM) user or RAM role, see [Resource Access Management](https://help.aliyun.com/document_detail/2767210.html).
//
// - This operation supports only generic secrets. You can perform the following operations:
//
//   - Add a version stage to a specified secret version.
//
//   - Remove a version stage from a specified secret version.
//
//   - Remove a version stage from a specified secret version and attach it to another secret version.
//
// - The total number of version stages for each generic secret cannot exceed 8.
//
// This topic provides an example of how to update the version stage of the secret named `secret001`. In this example, the `ACSCurrent` stage is used to mark version `002`.
//
// @param request - UpdateSecretVersionStageRequest
//
// @return UpdateSecretVersionStageResponse
func (client *Client) UpdateSecretVersionStage(request *UpdateSecretVersionStageRequest) (_result *UpdateSecretVersionStageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSecretVersionStageResponse{}
	_body, _err := client.UpdateSecretVersionStageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies the HMAC message authentication code of a specific message by using a specified key.
//
// Description:
//
// For details about the access policy required when a RAM user or RAM role invokes this operation, refer to access control.
//
// This operation can be invoked through a shared gateway or a dedicated gateway. For more information, refer to Alibaba Cloud SDK.
//
// - Shared gateway: Access KMS through a public or VPC endpoint. This method requires you to enable the public network access switch. For more information, refer to accessing keys in a KMS instance over the Internet.
//
// - Dedicated gateway: Access KMS through a KMS private endpoint (<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com).
//
// @param request - VerifyMacRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyMacResponse
func (client *Client) VerifyMacWithOptions(request *VerifyMacRequest, runtime *dara.RuntimeOptions) (_result *VerifyMacResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.Mac) {
		query["Mac"] = request.Mac
	}

	if !dara.IsNil(request.Message) {
		query["Message"] = request.Message
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyMac"),
		Version:     dara.String("2016-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyMacResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the HMAC message authentication code of a specific message by using a specified key.
//
// Description:
//
// For details about the access policy required when a RAM user or RAM role invokes this operation, refer to access control.
//
// This operation can be invoked through a shared gateway or a dedicated gateway. For more information, refer to Alibaba Cloud SDK.
//
// - Shared gateway: Access KMS through a public or VPC endpoint. This method requires you to enable the public network access switch. For more information, refer to accessing keys in a KMS instance over the Internet.
//
// - Dedicated gateway: Access KMS through a KMS private endpoint (<YOUR_KMS_INSTANCE_ID>.cryptoservice.kms.aliyuncs.com).
//
// @param request - VerifyMacRequest
//
// @return VerifyMacResponse
func (client *Client) VerifyMac(request *VerifyMacRequest) (_result *VerifyMacResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifyMacResponse{}
	_body, _err := client.VerifyMacWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
