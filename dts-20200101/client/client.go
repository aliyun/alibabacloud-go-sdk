// This file is auto-generated, don't edit it. Thanks.
package client

import (
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
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-qingdao":                  dara.String("dts.aliyuncs.com"),
		"cn-beijing":                  dara.String("dts.aliyuncs.com"),
		"cn-zhangjiakou":              dara.String("dts.aliyuncs.com"),
		"cn-huhehaote":                dara.String("dts.aliyuncs.com"),
		"cn-hangzhou":                 dara.String("dts.aliyuncs.com"),
		"cn-shanghai":                 dara.String("dts.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("dts.aliyuncs.com"),
		"cn-hongkong":                 dara.String("dts.aliyuncs.com"),
		"ap-southeast-1":              dara.String("dts.aliyuncs.com"),
		"ap-southeast-2":              dara.String("dts.aliyuncs.com"),
		"ap-southeast-3":              dara.String("dts.aliyuncs.com"),
		"ap-southeast-5":              dara.String("dts.aliyuncs.com"),
		"eu-west-1":                   dara.String("dts.aliyuncs.com"),
		"us-west-1":                   dara.String("dts.aliyuncs.com"),
		"us-east-1":                   dara.String("dts.aliyuncs.com"),
		"eu-central-1":                dara.String("dts.aliyuncs.com"),
		"me-east-1":                   dara.String("dts.aliyuncs.com"),
		"ap-south-1":                  dara.String("dts.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("dts.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("dts.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("dts.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("dts.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("dts.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("dts.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("dts.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("dts.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("dts.aliyuncs.com"),
		"cn-chengdu":                  dara.String("dts.aliyuncs.com"),
		"cn-edge-1":                   dara.String("dts.aliyuncs.com"),
		"cn-fujian":                   dara.String("dts.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("dts.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("dts.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("dts.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("dts.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("dts.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("dts.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("dts.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("dts.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("dts.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("dts.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("dts.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("dts.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("dts.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("dts.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("dts.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("dts.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("dts.aliyuncs.com"),
		"cn-wuhan":                    dara.String("dts.aliyuncs.com"),
		"cn-wulanchabu":               dara.String("dts.aliyuncs.com"),
		"cn-yushanfang":               dara.String("dts.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("dts.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("dts.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("dts.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("dts.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("dts.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("dts"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) _postOSSObject(bucketName *string, form map[string]interface{}) (_result map[string]interface{}, _err error) {
	request_ := dara.NewRequest()
	boundary := dara.GetBoundary()
	request_.Protocol = dara.String("HTTPS")
	request_.Method = dara.String("POST")
	request_.Pathname = dara.String("/")
	request_.Headers = map[string]*string{
		"host":       dara.String(dara.ToString(form["host"])),
		"date":       openapiutil.GetDateUTCString(),
		"user-agent": openapiutil.GetUserAgent(dara.String("")),
	}
	request_.Headers["content-type"] = dara.String("multipart/form-data; boundary=" + boundary)
	request_.Body = dara.ToFileForm(form, boundary)
	response_, _err := dara.DoRequest(request_, nil)
	if _err != nil {
		return nil, _err
	}

	_result, _err = _postOSSObject_opResponse(response_)
	if _err != nil {
		return nil, _err
	}

	return _result, nil
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
// Configures a data migration or synchronization task.
//
// @param request - ConfigureDtsJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureDtsJobResponse
func (client *Client) ConfigureDtsJobWithOptions(request *ConfigureDtsJobRequest, runtime *dara.RuntimeOptions) (_result *ConfigureDtsJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Checkpoint) {
		query["Checkpoint"] = request.Checkpoint
	}

	if !dara.IsNil(request.DataCheckConfigure) {
		query["DataCheckConfigure"] = request.DataCheckConfigure
	}

	if !dara.IsNil(request.DataInitialization) {
		query["DataInitialization"] = request.DataInitialization
	}

	if !dara.IsNil(request.DataSynchronization) {
		query["DataSynchronization"] = request.DataSynchronization
	}

	if !dara.IsNil(request.DedicatedClusterId) {
		query["DedicatedClusterId"] = request.DedicatedClusterId
	}

	if !dara.IsNil(request.DelayNotice) {
		query["DelayNotice"] = request.DelayNotice
	}

	if !dara.IsNil(request.DelayPhone) {
		query["DelayPhone"] = request.DelayPhone
	}

	if !dara.IsNil(request.DelayRuleTime) {
		query["DelayRuleTime"] = request.DelayRuleTime
	}

	if !dara.IsNil(request.DestCaCertificateOssUrl) {
		query["DestCaCertificateOssUrl"] = request.DestCaCertificateOssUrl
	}

	if !dara.IsNil(request.DestCaCertificatePassword) {
		query["DestCaCertificatePassword"] = request.DestCaCertificatePassword
	}

	if !dara.IsNil(request.DestClientCertOssUrl) {
		query["DestClientCertOssUrl"] = request.DestClientCertOssUrl
	}

	if !dara.IsNil(request.DestClientKeyOssUrl) {
		query["DestClientKeyOssUrl"] = request.DestClientKeyOssUrl
	}

	if !dara.IsNil(request.DestClientPassword) {
		query["DestClientPassword"] = request.DestClientPassword
	}

	if !dara.IsNil(request.DestPrimaryVswId) {
		query["DestPrimaryVswId"] = request.DestPrimaryVswId
	}

	if !dara.IsNil(request.DestSecondaryVswId) {
		query["DestSecondaryVswId"] = request.DestSecondaryVswId
	}

	if !dara.IsNil(request.DestinationEndpointDataBaseName) {
		query["DestinationEndpointDataBaseName"] = request.DestinationEndpointDataBaseName
	}

	if !dara.IsNil(request.DestinationEndpointEngineName) {
		query["DestinationEndpointEngineName"] = request.DestinationEndpointEngineName
	}

	if !dara.IsNil(request.DestinationEndpointIP) {
		query["DestinationEndpointIP"] = request.DestinationEndpointIP
	}

	if !dara.IsNil(request.DestinationEndpointInstanceID) {
		query["DestinationEndpointInstanceID"] = request.DestinationEndpointInstanceID
	}

	if !dara.IsNil(request.DestinationEndpointInstanceType) {
		query["DestinationEndpointInstanceType"] = request.DestinationEndpointInstanceType
	}

	if !dara.IsNil(request.DestinationEndpointOracleSID) {
		query["DestinationEndpointOracleSID"] = request.DestinationEndpointOracleSID
	}

	if !dara.IsNil(request.DestinationEndpointOwnerID) {
		query["DestinationEndpointOwnerID"] = request.DestinationEndpointOwnerID
	}

	if !dara.IsNil(request.DestinationEndpointPassword) {
		query["DestinationEndpointPassword"] = request.DestinationEndpointPassword
	}

	if !dara.IsNil(request.DestinationEndpointPort) {
		query["DestinationEndpointPort"] = request.DestinationEndpointPort
	}

	if !dara.IsNil(request.DestinationEndpointRegion) {
		query["DestinationEndpointRegion"] = request.DestinationEndpointRegion
	}

	if !dara.IsNil(request.DestinationEndpointRole) {
		query["DestinationEndpointRole"] = request.DestinationEndpointRole
	}

	if !dara.IsNil(request.DestinationEndpointUserName) {
		query["DestinationEndpointUserName"] = request.DestinationEndpointUserName
	}

	if !dara.IsNil(request.DisasterRecoveryJob) {
		query["DisasterRecoveryJob"] = request.DisasterRecoveryJob
	}

	if !dara.IsNil(request.DtsBisLabel) {
		query["DtsBisLabel"] = request.DtsBisLabel
	}

	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.DtsJobName) {
		query["DtsJobName"] = request.DtsJobName
	}

	if !dara.IsNil(request.ErrorNotice) {
		query["ErrorNotice"] = request.ErrorNotice
	}

	if !dara.IsNil(request.ErrorPhone) {
		query["ErrorPhone"] = request.ErrorPhone
	}

	if !dara.IsNil(request.FileOssUrl) {
		query["FileOssUrl"] = request.FileOssUrl
	}

	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.MaxDu) {
		query["MaxDu"] = request.MaxDu
	}

	if !dara.IsNil(request.MinDu) {
		query["MinDu"] = request.MinDu
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SourceEndpointDatabaseName) {
		query["SourceEndpointDatabaseName"] = request.SourceEndpointDatabaseName
	}

	if !dara.IsNil(request.SourceEndpointEngineName) {
		query["SourceEndpointEngineName"] = request.SourceEndpointEngineName
	}

	if !dara.IsNil(request.SourceEndpointIP) {
		query["SourceEndpointIP"] = request.SourceEndpointIP
	}

	if !dara.IsNil(request.SourceEndpointInstanceID) {
		query["SourceEndpointInstanceID"] = request.SourceEndpointInstanceID
	}

	if !dara.IsNil(request.SourceEndpointInstanceType) {
		query["SourceEndpointInstanceType"] = request.SourceEndpointInstanceType
	}

	if !dara.IsNil(request.SourceEndpointOracleSID) {
		query["SourceEndpointOracleSID"] = request.SourceEndpointOracleSID
	}

	if !dara.IsNil(request.SourceEndpointOwnerID) {
		query["SourceEndpointOwnerID"] = request.SourceEndpointOwnerID
	}

	if !dara.IsNil(request.SourceEndpointPassword) {
		query["SourceEndpointPassword"] = request.SourceEndpointPassword
	}

	if !dara.IsNil(request.SourceEndpointPort) {
		query["SourceEndpointPort"] = request.SourceEndpointPort
	}

	if !dara.IsNil(request.SourceEndpointRegion) {
		query["SourceEndpointRegion"] = request.SourceEndpointRegion
	}

	if !dara.IsNil(request.SourceEndpointRole) {
		query["SourceEndpointRole"] = request.SourceEndpointRole
	}

	if !dara.IsNil(request.SourceEndpointUserName) {
		query["SourceEndpointUserName"] = request.SourceEndpointUserName
	}

	if !dara.IsNil(request.SourceEndpointVSwitchID) {
		query["SourceEndpointVSwitchID"] = request.SourceEndpointVSwitchID
	}

	if !dara.IsNil(request.SrcCaCertificateOssUrl) {
		query["SrcCaCertificateOssUrl"] = request.SrcCaCertificateOssUrl
	}

	if !dara.IsNil(request.SrcCaCertificatePassword) {
		query["SrcCaCertificatePassword"] = request.SrcCaCertificatePassword
	}

	if !dara.IsNil(request.SrcClientCertOssUrl) {
		query["SrcClientCertOssUrl"] = request.SrcClientCertOssUrl
	}

	if !dara.IsNil(request.SrcClientKeyOssUrl) {
		query["SrcClientKeyOssUrl"] = request.SrcClientKeyOssUrl
	}

	if !dara.IsNil(request.SrcClientPassword) {
		query["SrcClientPassword"] = request.SrcClientPassword
	}

	if !dara.IsNil(request.SrcPrimaryVswId) {
		query["SrcPrimaryVswId"] = request.SrcPrimaryVswId
	}

	if !dara.IsNil(request.SrcSecondaryVswId) {
		query["SrcSecondaryVswId"] = request.SrcSecondaryVswId
	}

	if !dara.IsNil(request.StructureInitialization) {
		query["StructureInitialization"] = request.StructureInitialization
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DbList) {
		body["DbList"] = request.DbList
	}

	if !dara.IsNil(request.Reserve) {
		body["Reserve"] = request.Reserve
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigureDtsJob"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigureDtsJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a data migration or synchronization task.
//
// @param request - ConfigureDtsJobRequest
//
// @return ConfigureDtsJobResponse
func (client *Client) ConfigureDtsJob(request *ConfigureDtsJobRequest) (_result *ConfigureDtsJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigureDtsJobResponse{}
	_body, _err := client.ConfigureDtsJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigureDtsJobAdvance(request *ConfigureDtsJobAdvanceRequest, runtime *dara.RuntimeOptions) (_result *ConfigureDtsJobResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("Dts"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	configureDtsJobReq := &ConfigureDtsJobRequest{}
	openapiutil.Convert(request, configureDtsJobReq)
	if !dara.IsNil(request.FileOssUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileOssUrlObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		configureDtsJobReq.FileOssUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	configureDtsJobResp, _err := client.ConfigureDtsJobWithOptions(configureDtsJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = configureDtsJobResp
	return _result, _err
}

// Summary:
//
// Configures a data migration task.
//
// @param request - ConfigureMigrationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureMigrationJobResponse
func (client *Client) ConfigureMigrationJobWithOptions(request *ConfigureMigrationJobRequest, runtime *dara.RuntimeOptions) (_result *ConfigureMigrationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.Checkpoint) {
		query["Checkpoint"] = request.Checkpoint
	}

	if !dara.IsNil(request.MigrationJobId) {
		query["MigrationJobId"] = request.MigrationJobId
	}

	if !dara.IsNil(request.MigrationJobName) {
		query["MigrationJobName"] = request.MigrationJobName
	}

	if !dara.IsNil(request.MigrationReserved) {
		query["MigrationReserved"] = request.MigrationReserved
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.DestinationEndpoint) {
		query["DestinationEndpoint"] = request.DestinationEndpoint
	}

	if !dara.IsNil(request.MigrationMode) {
		query["MigrationMode"] = request.MigrationMode
	}

	if !dara.IsNil(request.SourceEndpoint) {
		query["SourceEndpoint"] = request.SourceEndpoint
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MigrationObject) {
		body["MigrationObject"] = request.MigrationObject
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigureMigrationJob"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigureMigrationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a data migration task.
//
// @param request - ConfigureMigrationJobRequest
//
// @return ConfigureMigrationJobResponse
func (client *Client) ConfigureMigrationJob(request *ConfigureMigrationJobRequest) (_result *ConfigureMigrationJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigureMigrationJobResponse{}
	_body, _err := client.ConfigureMigrationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures alert settings to monitor a data migration instance.
//
// @param request - ConfigureMigrationJobAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureMigrationJobAlertResponse
func (client *Client) ConfigureMigrationJobAlertWithOptions(request *ConfigureMigrationJobAlertRequest, runtime *dara.RuntimeOptions) (_result *ConfigureMigrationJobAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.DelayAlertPhone) {
		query["DelayAlertPhone"] = request.DelayAlertPhone
	}

	if !dara.IsNil(request.DelayAlertStatus) {
		query["DelayAlertStatus"] = request.DelayAlertStatus
	}

	if !dara.IsNil(request.DelayOverSeconds) {
		query["DelayOverSeconds"] = request.DelayOverSeconds
	}

	if !dara.IsNil(request.ErrorAlertPhone) {
		query["ErrorAlertPhone"] = request.ErrorAlertPhone
	}

	if !dara.IsNil(request.ErrorAlertStatus) {
		query["ErrorAlertStatus"] = request.ErrorAlertStatus
	}

	if !dara.IsNil(request.MigrationJobId) {
		query["MigrationJobId"] = request.MigrationJobId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigureMigrationJobAlert"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigureMigrationJobAlertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures alert settings to monitor a data migration instance.
//
// @param request - ConfigureMigrationJobAlertRequest
//
// @return ConfigureMigrationJobAlertResponse
func (client *Client) ConfigureMigrationJobAlert(request *ConfigureMigrationJobAlertRequest) (_result *ConfigureMigrationJobAlertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigureMigrationJobAlertResponse{}
	_body, _err := client.ConfigureMigrationJobAlertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a change tracking task.
//
// Description:
//
// >  You can preview related API operation parameters when you configure a change tracking task in the Data Transmission Service (DTS) console. This helps you configure the request parameters of this API operation. For more information, see [Preview the request parameters of API operations](https://help.aliyun.com/document_detail/2851612.html).
//
// @param request - ConfigureSubscriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureSubscriptionResponse
func (client *Client) ConfigureSubscriptionWithOptions(request *ConfigureSubscriptionRequest, runtime *dara.RuntimeOptions) (_result *ConfigureSubscriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Checkpoint) {
		query["Checkpoint"] = request.Checkpoint
	}

	if !dara.IsNil(request.DbList) {
		query["DbList"] = request.DbList
	}

	if !dara.IsNil(request.DedicatedClusterId) {
		query["DedicatedClusterId"] = request.DedicatedClusterId
	}

	if !dara.IsNil(request.DelayNotice) {
		query["DelayNotice"] = request.DelayNotice
	}

	if !dara.IsNil(request.DelayPhone) {
		query["DelayPhone"] = request.DelayPhone
	}

	if !dara.IsNil(request.DelayRuleTime) {
		query["DelayRuleTime"] = request.DelayRuleTime
	}

	if !dara.IsNil(request.DtsBisLabel) {
		query["DtsBisLabel"] = request.DtsBisLabel
	}

	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.DtsJobName) {
		query["DtsJobName"] = request.DtsJobName
	}

	if !dara.IsNil(request.ErrorNotice) {
		query["ErrorNotice"] = request.ErrorNotice
	}

	if !dara.IsNil(request.ErrorPhone) {
		query["ErrorPhone"] = request.ErrorPhone
	}

	if !dara.IsNil(request.MaxDu) {
		query["MaxDu"] = request.MaxDu
	}

	if !dara.IsNil(request.MinDu) {
		query["MinDu"] = request.MinDu
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Reserve) {
		query["Reserve"] = request.Reserve
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SourceEndpointDatabaseName) {
		query["SourceEndpointDatabaseName"] = request.SourceEndpointDatabaseName
	}

	if !dara.IsNil(request.SourceEndpointEngineName) {
		query["SourceEndpointEngineName"] = request.SourceEndpointEngineName
	}

	if !dara.IsNil(request.SourceEndpointIP) {
		query["SourceEndpointIP"] = request.SourceEndpointIP
	}

	if !dara.IsNil(request.SourceEndpointInstanceID) {
		query["SourceEndpointInstanceID"] = request.SourceEndpointInstanceID
	}

	if !dara.IsNil(request.SourceEndpointInstanceType) {
		query["SourceEndpointInstanceType"] = request.SourceEndpointInstanceType
	}

	if !dara.IsNil(request.SourceEndpointOracleSID) {
		query["SourceEndpointOracleSID"] = request.SourceEndpointOracleSID
	}

	if !dara.IsNil(request.SourceEndpointOwnerID) {
		query["SourceEndpointOwnerID"] = request.SourceEndpointOwnerID
	}

	if !dara.IsNil(request.SourceEndpointPassword) {
		query["SourceEndpointPassword"] = request.SourceEndpointPassword
	}

	if !dara.IsNil(request.SourceEndpointPort) {
		query["SourceEndpointPort"] = request.SourceEndpointPort
	}

	if !dara.IsNil(request.SourceEndpointRegion) {
		query["SourceEndpointRegion"] = request.SourceEndpointRegion
	}

	if !dara.IsNil(request.SourceEndpointRole) {
		query["SourceEndpointRole"] = request.SourceEndpointRole
	}

	if !dara.IsNil(request.SourceEndpointUserName) {
		query["SourceEndpointUserName"] = request.SourceEndpointUserName
	}

	if !dara.IsNil(request.SrcCaCertificateOssUrl) {
		query["SrcCaCertificateOssUrl"] = request.SrcCaCertificateOssUrl
	}

	if !dara.IsNil(request.SrcCaCertificatePassword) {
		query["SrcCaCertificatePassword"] = request.SrcCaCertificatePassword
	}

	if !dara.IsNil(request.SrcClientCertOssUrl) {
		query["SrcClientCertOssUrl"] = request.SrcClientCertOssUrl
	}

	if !dara.IsNil(request.SrcClientKeyOssUrl) {
		query["SrcClientKeyOssUrl"] = request.SrcClientKeyOssUrl
	}

	if !dara.IsNil(request.SrcClientPassword) {
		query["SrcClientPassword"] = request.SrcClientPassword
	}

	if !dara.IsNil(request.SubscriptionDataTypeDDL) {
		query["SubscriptionDataTypeDDL"] = request.SubscriptionDataTypeDDL
	}

	if !dara.IsNil(request.SubscriptionDataTypeDML) {
		query["SubscriptionDataTypeDML"] = request.SubscriptionDataTypeDML
	}

	if !dara.IsNil(request.SubscriptionInstanceNetworkType) {
		query["SubscriptionInstanceNetworkType"] = request.SubscriptionInstanceNetworkType
	}

	if !dara.IsNil(request.SubscriptionInstanceVPCId) {
		query["SubscriptionInstanceVPCId"] = request.SubscriptionInstanceVPCId
	}

	if !dara.IsNil(request.SubscriptionInstanceVSwitchId) {
		query["SubscriptionInstanceVSwitchId"] = request.SubscriptionInstanceVSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigureSubscription"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigureSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a change tracking task.
//
// Description:
//
// >  You can preview related API operation parameters when you configure a change tracking task in the Data Transmission Service (DTS) console. This helps you configure the request parameters of this API operation. For more information, see [Preview the request parameters of API operations](https://help.aliyun.com/document_detail/2851612.html).
//
// @param request - ConfigureSubscriptionRequest
//
// @return ConfigureSubscriptionResponse
func (client *Client) ConfigureSubscription(request *ConfigureSubscriptionRequest) (_result *ConfigureSubscriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigureSubscriptionResponse{}
	_body, _err := client.ConfigureSubscriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a change tracking instance of the previous version.
//
// Description:
//
// Before you call this operation, you must call the [CreateSubscriptionInstance](https://help.aliyun.com/document_detail/49436.html) operation to create a change tracking instance.
//
// @param request - ConfigureSubscriptionInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureSubscriptionInstanceResponse
func (client *Client) ConfigureSubscriptionInstanceWithOptions(request *ConfigureSubscriptionInstanceRequest, runtime *dara.RuntimeOptions) (_result *ConfigureSubscriptionInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SubscriptionInstanceId) {
		query["SubscriptionInstanceId"] = request.SubscriptionInstanceId
	}

	if !dara.IsNil(request.SubscriptionInstanceName) {
		query["SubscriptionInstanceName"] = request.SubscriptionInstanceName
	}

	if !dara.IsNil(request.SubscriptionInstanceNetworkType) {
		query["SubscriptionInstanceNetworkType"] = request.SubscriptionInstanceNetworkType
	}

	if !dara.IsNil(request.SourceEndpoint) {
		query["SourceEndpoint"] = request.SourceEndpoint
	}

	if !dara.IsNil(request.SubscriptionDataType) {
		query["SubscriptionDataType"] = request.SubscriptionDataType
	}

	if !dara.IsNil(request.SubscriptionInstance) {
		query["SubscriptionInstance"] = request.SubscriptionInstance
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SubscriptionObject) {
		body["SubscriptionObject"] = request.SubscriptionObject
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigureSubscriptionInstance"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigureSubscriptionInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a change tracking instance of the previous version.
//
// Description:
//
// Before you call this operation, you must call the [CreateSubscriptionInstance](https://help.aliyun.com/document_detail/49436.html) operation to create a change tracking instance.
//
// @param request - ConfigureSubscriptionInstanceRequest
//
// @return ConfigureSubscriptionInstanceResponse
func (client *Client) ConfigureSubscriptionInstance(request *ConfigureSubscriptionInstanceRequest) (_result *ConfigureSubscriptionInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigureSubscriptionInstanceResponse{}
	_body, _err := client.ConfigureSubscriptionInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures alert settings to monitor a change tracking instance.
//
// @param request - ConfigureSubscriptionInstanceAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureSubscriptionInstanceAlertResponse
func (client *Client) ConfigureSubscriptionInstanceAlertWithOptions(request *ConfigureSubscriptionInstanceAlertRequest, runtime *dara.RuntimeOptions) (_result *ConfigureSubscriptionInstanceAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.DelayAlertPhone) {
		query["DelayAlertPhone"] = request.DelayAlertPhone
	}

	if !dara.IsNil(request.DelayAlertStatus) {
		query["DelayAlertStatus"] = request.DelayAlertStatus
	}

	if !dara.IsNil(request.DelayOverSeconds) {
		query["DelayOverSeconds"] = request.DelayOverSeconds
	}

	if !dara.IsNil(request.ErrorAlertPhone) {
		query["ErrorAlertPhone"] = request.ErrorAlertPhone
	}

	if !dara.IsNil(request.ErrorAlertStatus) {
		query["ErrorAlertStatus"] = request.ErrorAlertStatus
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SubscriptionInstanceId) {
		query["SubscriptionInstanceId"] = request.SubscriptionInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigureSubscriptionInstanceAlert"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigureSubscriptionInstanceAlertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures alert settings to monitor a change tracking instance.
//
// @param request - ConfigureSubscriptionInstanceAlertRequest
//
// @return ConfigureSubscriptionInstanceAlertResponse
func (client *Client) ConfigureSubscriptionInstanceAlert(request *ConfigureSubscriptionInstanceAlertRequest) (_result *ConfigureSubscriptionInstanceAlertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigureSubscriptionInstanceAlertResponse{}
	_body, _err := client.ConfigureSubscriptionInstanceAlertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a data synchronization task.
//
// Description:
//
// Before you call this operation, you must call the [CreateSynchronizationJob](https://help.aliyun.com/document_detail/49446.html) operation to create a data synchronization instance.
//
// >
//
//   - After you call this operation to configure a data synchronization task, the task will be automatically started and prechecked. You do not need to call the [StartSynchronizationJob](https://help.aliyun.com/document_detail/49448.html) operation to start the task.
//
//   - A data synchronization task may fail to be started due to precheck failures. You can call the [DescribeSynchronizationJobStatus](https://help.aliyun.com/document_detail/49453.html) operation to query the status of the task. Then, you can change parameter settings based on the error messages about the precheck failures. After you fix the issue, you must call the [StartSynchronizationJob](https://help.aliyun.com/document_detail/49448.html) operation to restart the data synchronization task.
//
// @param request - ConfigureSynchronizationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureSynchronizationJobResponse
func (client *Client) ConfigureSynchronizationJobWithOptions(request *ConfigureSynchronizationJobRequest, runtime *dara.RuntimeOptions) (_result *ConfigureSynchronizationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.Checkpoint) {
		query["Checkpoint"] = request.Checkpoint
	}

	if !dara.IsNil(request.DataInitialization) {
		query["DataInitialization"] = request.DataInitialization
	}

	if !dara.IsNil(request.MigrationReserved) {
		query["MigrationReserved"] = request.MigrationReserved
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StructureInitialization) {
		query["StructureInitialization"] = request.StructureInitialization
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !dara.IsNil(request.SynchronizationJobId) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	if !dara.IsNil(request.SynchronizationJobName) {
		query["SynchronizationJobName"] = request.SynchronizationJobName
	}

	if !dara.IsNil(request.DestinationEndpoint) {
		query["DestinationEndpoint"] = request.DestinationEndpoint
	}

	if !dara.IsNil(request.PartitionKey) {
		query["PartitionKey"] = request.PartitionKey
	}

	if !dara.IsNil(request.SourceEndpoint) {
		query["SourceEndpoint"] = request.SourceEndpoint
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SynchronizationObjects) {
		body["SynchronizationObjects"] = request.SynchronizationObjects
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigureSynchronizationJob"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigureSynchronizationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a data synchronization task.
//
// Description:
//
// Before you call this operation, you must call the [CreateSynchronizationJob](https://help.aliyun.com/document_detail/49446.html) operation to create a data synchronization instance.
//
// >
//
//   - After you call this operation to configure a data synchronization task, the task will be automatically started and prechecked. You do not need to call the [StartSynchronizationJob](https://help.aliyun.com/document_detail/49448.html) operation to start the task.
//
//   - A data synchronization task may fail to be started due to precheck failures. You can call the [DescribeSynchronizationJobStatus](https://help.aliyun.com/document_detail/49453.html) operation to query the status of the task. Then, you can change parameter settings based on the error messages about the precheck failures. After you fix the issue, you must call the [StartSynchronizationJob](https://help.aliyun.com/document_detail/49448.html) operation to restart the data synchronization task.
//
// @param request - ConfigureSynchronizationJobRequest
//
// @return ConfigureSynchronizationJobResponse
func (client *Client) ConfigureSynchronizationJob(request *ConfigureSynchronizationJobRequest) (_result *ConfigureSynchronizationJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigureSynchronizationJobResponse{}
	_body, _err := client.ConfigureSynchronizationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures alert settings to monitor a data synchronization instance.
//
// @param request - ConfigureSynchronizationJobAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureSynchronizationJobAlertResponse
func (client *Client) ConfigureSynchronizationJobAlertWithOptions(request *ConfigureSynchronizationJobAlertRequest, runtime *dara.RuntimeOptions) (_result *ConfigureSynchronizationJobAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.DelayAlertPhone) {
		query["DelayAlertPhone"] = request.DelayAlertPhone
	}

	if !dara.IsNil(request.DelayAlertStatus) {
		query["DelayAlertStatus"] = request.DelayAlertStatus
	}

	if !dara.IsNil(request.DelayOverSeconds) {
		query["DelayOverSeconds"] = request.DelayOverSeconds
	}

	if !dara.IsNil(request.ErrorAlertPhone) {
		query["ErrorAlertPhone"] = request.ErrorAlertPhone
	}

	if !dara.IsNil(request.ErrorAlertStatus) {
		query["ErrorAlertStatus"] = request.ErrorAlertStatus
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !dara.IsNil(request.SynchronizationJobId) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigureSynchronizationJobAlert"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigureSynchronizationJobAlertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures alert settings to monitor a data synchronization instance.
//
// @param request - ConfigureSynchronizationJobAlertRequest
//
// @return ConfigureSynchronizationJobAlertResponse
func (client *Client) ConfigureSynchronizationJobAlert(request *ConfigureSynchronizationJobAlertRequest) (_result *ConfigureSynchronizationJobAlertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigureSynchronizationJobAlertResponse{}
	_body, _err := client.ConfigureSynchronizationJobAlertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables image matching for a data synchronization instance.
//
// @param request - ConfigureSynchronizationJobReplicatorCompareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureSynchronizationJobReplicatorCompareResponse
func (client *Client) ConfigureSynchronizationJobReplicatorCompareWithOptions(request *ConfigureSynchronizationJobReplicatorCompareRequest, runtime *dara.RuntimeOptions) (_result *ConfigureSynchronizationJobReplicatorCompareResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !dara.IsNil(request.SynchronizationJobId) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	if !dara.IsNil(request.SynchronizationReplicatorCompareEnable) {
		query["SynchronizationReplicatorCompareEnable"] = request.SynchronizationReplicatorCompareEnable
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigureSynchronizationJobReplicatorCompare"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigureSynchronizationJobReplicatorCompareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables image matching for a data synchronization instance.
//
// @param request - ConfigureSynchronizationJobReplicatorCompareRequest
//
// @return ConfigureSynchronizationJobReplicatorCompareResponse
func (client *Client) ConfigureSynchronizationJobReplicatorCompare(request *ConfigureSynchronizationJobReplicatorCompareRequest) (_result *ConfigureSynchronizationJobReplicatorCompareResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigureSynchronizationJobReplicatorCompareResponse{}
	_body, _err := client.ConfigureSynchronizationJobReplicatorCompareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Transfers resource groups of instance resources.
//
// @param request - ConvertInstanceResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConvertInstanceResourceGroupResponse
func (client *Client) ConvertInstanceResourceGroupWithOptions(request *ConvertInstanceResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ConvertInstanceResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ZeroEtlJob) {
		query["ZeroEtlJob"] = request.ZeroEtlJob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConvertInstanceResourceGroup"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConvertInstanceResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Transfers resource groups of instance resources.
//
// @param request - ConvertInstanceResourceGroupRequest
//
// @return ConvertInstanceResourceGroupResponse
func (client *Client) ConvertInstanceResourceGroup(request *ConvertInstanceResourceGroupRequest) (_result *ConvertInstanceResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConvertInstanceResourceGroupResponse{}
	_body, _err := client.ConvertInstanceResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Counts tasks by condition.
//
// @param request - CountJobByConditionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CountJobByConditionResponse
func (client *Client) CountJobByConditionWithOptions(request *CountJobByConditionRequest, runtime *dara.RuntimeOptions) (_result *CountJobByConditionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestDbType) {
		query["DestDbType"] = request.DestDbType
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SrcDbType) {
		query["SrcDbType"] = request.SrcDbType
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CountJobByCondition"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CountJobByConditionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Counts tasks by condition.
//
// @param request - CountJobByConditionRequest
//
// @return CountJobByConditionResponse
func (client *Client) CountJobByCondition(request *CountJobByConditionRequest) (_result *CountJobByConditionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CountJobByConditionResponse{}
	_body, _err := client.CountJobByConditionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a consumer group for a change tracking task. Downstream clients can use the consumer group to consume tracked data.
//
// @param request - CreateConsumerChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConsumerChannelResponse
func (client *Client) CreateConsumerChannelWithOptions(request *CreateConsumerChannelRequest, runtime *dara.RuntimeOptions) (_result *CreateConsumerChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerGroupName) {
		query["ConsumerGroupName"] = request.ConsumerGroupName
	}

	if !dara.IsNil(request.ConsumerGroupPassword) {
		query["ConsumerGroupPassword"] = request.ConsumerGroupPassword
	}

	if !dara.IsNil(request.ConsumerGroupUserName) {
		query["ConsumerGroupUserName"] = request.ConsumerGroupUserName
	}

	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConsumerChannel"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConsumerChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a consumer group for a change tracking task. Downstream clients can use the consumer group to consume tracked data.
//
// @param request - CreateConsumerChannelRequest
//
// @return CreateConsumerChannelResponse
func (client *Client) CreateConsumerChannel(request *CreateConsumerChannelRequest) (_result *CreateConsumerChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateConsumerChannelResponse{}
	_body, _err := client.CreateConsumerChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a consumer group for a change tracking instance.
//
// @param request - CreateConsumerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConsumerGroupResponse
func (client *Client) CreateConsumerGroupWithOptions(request *CreateConsumerGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateConsumerGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ConsumerGroupName) {
		query["ConsumerGroupName"] = request.ConsumerGroupName
	}

	if !dara.IsNil(request.ConsumerGroupPassword) {
		query["ConsumerGroupPassword"] = request.ConsumerGroupPassword
	}

	if !dara.IsNil(request.ConsumerGroupUserName) {
		query["ConsumerGroupUserName"] = request.ConsumerGroupUserName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SubscriptionInstanceId) {
		query["SubscriptionInstanceId"] = request.SubscriptionInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConsumerGroup"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConsumerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a consumer group for a change tracking instance.
//
// @param request - CreateConsumerGroupRequest
//
// @return CreateConsumerGroupResponse
func (client *Client) CreateConsumerGroup(request *CreateConsumerGroupRequest) (_result *CreateConsumerGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateConsumerGroupResponse{}
	_body, _err := client.CreateConsumerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an alert rule.
//
// @param request - CreateDedicatedClusterMonitorRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDedicatedClusterMonitorRuleResponse
func (client *Client) CreateDedicatedClusterMonitorRuleWithOptions(request *CreateDedicatedClusterMonitorRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateDedicatedClusterMonitorRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CpuAlarmThreshold) {
		query["CpuAlarmThreshold"] = request.CpuAlarmThreshold
	}

	if !dara.IsNil(request.DedicatedClusterId) {
		query["DedicatedClusterId"] = request.DedicatedClusterId
	}

	if !dara.IsNil(request.DiskAlarmThreshold) {
		query["DiskAlarmThreshold"] = request.DiskAlarmThreshold
	}

	if !dara.IsNil(request.DuAlarmThreshold) {
		query["DuAlarmThreshold"] = request.DuAlarmThreshold
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MemAlarmThreshold) {
		query["MemAlarmThreshold"] = request.MemAlarmThreshold
	}

	if !dara.IsNil(request.NoticeSwitch) {
		query["NoticeSwitch"] = request.NoticeSwitch
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Phones) {
		query["Phones"] = request.Phones
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDedicatedClusterMonitorRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDedicatedClusterMonitorRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an alert rule.
//
// @param request - CreateDedicatedClusterMonitorRuleRequest
//
// @return CreateDedicatedClusterMonitorRuleResponse
func (client *Client) CreateDedicatedClusterMonitorRule(request *CreateDedicatedClusterMonitorRuleRequest) (_result *CreateDedicatedClusterMonitorRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDedicatedClusterMonitorRuleResponse{}
	_body, _err := client.CreateDedicatedClusterMonitorRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看工作流任务结果
//
// @param request - CreateDocParserJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDocParserJobResponse
func (client *Client) CreateDocParserJobWithOptions(request *CreateDocParserJobRequest, runtime *dara.RuntimeOptions) (_result *CreateDocParserJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResultType) {
		query["ResultType"] = request.ResultType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDocParserJob"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDocParserJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看工作流任务结果
//
// @param request - CreateDocParserJobRequest
//
// @return CreateDocParserJobResponse
func (client *Client) CreateDocParserJob(request *CreateDocParserJobRequest) (_result *CreateDocParserJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDocParserJobResponse{}
	_body, _err := client.CreateDocParserJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDocParserJobAdvance(request *CreateDocParserJobAdvanceRequest, runtime *dara.RuntimeOptions) (_result *CreateDocParserJobResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("Dts"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	createDocParserJobReq := &CreateDocParserJobRequest{}
	openapiutil.Convert(request, createDocParserJobReq)
	if !dara.IsNil(request.FileUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		createDocParserJobReq.FileUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	createDocParserJobResp, _err := client.CreateDocParserJobWithOptions(createDocParserJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = createDocParserJobResp
	return _result, _err
}

// Summary:
//
// Purchases a Data Transmission Service (DTS) instance.
//
// Description:
//
//	  Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/zh/product/apsaradb-for-mongodb/pricing) of DTS.
//
//		- If you want to run a DTS task on a DTS dedicated cluster, you must configure the task before you purchase a DTS instance. You can call the [ConfigureDtsJob](https://help.aliyun.com/document_detail/208399.html) operation to configure a DTS task.
//
// @param request - CreateDtsInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDtsInstanceResponse
func (client *Client) CreateDtsInstanceWithOptions(request *CreateDtsInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateDtsInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoStart) {
		query["AutoStart"] = request.AutoStart
	}

	if !dara.IsNil(request.ComputeUnit) {
		query["ComputeUnit"] = request.ComputeUnit
	}

	if !dara.IsNil(request.DatabaseCount) {
		query["DatabaseCount"] = request.DatabaseCount
	}

	if !dara.IsNil(request.DestinationEndpointEngineName) {
		query["DestinationEndpointEngineName"] = request.DestinationEndpointEngineName
	}

	if !dara.IsNil(request.DestinationRegion) {
		query["DestinationRegion"] = request.DestinationRegion
	}

	if !dara.IsNil(request.DtsRegion) {
		query["DtsRegion"] = request.DtsRegion
	}

	if !dara.IsNil(request.Du) {
		query["Du"] = request.Du
	}

	if !dara.IsNil(request.FeeType) {
		query["FeeType"] = request.FeeType
	}

	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.MaxDu) {
		query["MaxDu"] = request.MaxDu
	}

	if !dara.IsNil(request.MinDu) {
		query["MinDu"] = request.MinDu
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.Quantity) {
		query["Quantity"] = request.Quantity
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SourceEndpointEngineName) {
		query["SourceEndpointEngineName"] = request.SourceEndpointEngineName
	}

	if !dara.IsNil(request.SourceRegion) {
		query["SourceRegion"] = request.SourceRegion
	}

	if !dara.IsNil(request.SyncArchitecture) {
		query["SyncArchitecture"] = request.SyncArchitecture
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDtsInstance"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDtsInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Purchases a Data Transmission Service (DTS) instance.
//
// Description:
//
//	  Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/zh/product/apsaradb-for-mongodb/pricing) of DTS.
//
//		- If you want to run a DTS task on a DTS dedicated cluster, you must configure the task before you purchase a DTS instance. You can call the [ConfigureDtsJob](https://help.aliyun.com/document_detail/208399.html) operation to configure a DTS task.
//
// @param request - CreateDtsInstanceRequest
//
// @return CreateDtsInstanceResponse
func (client *Client) CreateDtsInstance(request *CreateDtsInstanceRequest) (_result *CreateDtsInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDtsInstanceResponse{}
	_body, _err := client.CreateDtsInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or modifies an alert rule for a Data Transmission Service (DTS) task.
//
// Description:
//
// DTS provides the following metrics for DTS tasks:************
//
//   - **Latency**: DTS monitors the latency of a DTS task. If the latency of the task exceeds the specified threshold, an alert is triggered. The threshold is specified in units of seconds.
//
//   - **Status**: DTS monitors the status of a DTS task. If the state of the task changes to **Error*	- or **Restore**, an alert is triggered.
//
//   - **Full Timeout**: DTS monitors the duration of a DTS task. If the duration of the task exceeds the specified threshold, an alert is triggered. The threshold is specified in units of hours.
//
// @param request - CreateJobMonitorRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJobMonitorRuleResponse
func (client *Client) CreateJobMonitorRuleWithOptions(request *CreateJobMonitorRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateJobMonitorRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DelayRuleTime) {
		query["DelayRuleTime"] = request.DelayRuleTime
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.NoticeValue) {
		query["NoticeValue"] = request.NoticeValue
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.Phone) {
		query["Phone"] = request.Phone
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.Times) {
		query["Times"] = request.Times
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateJobMonitorRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateJobMonitorRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies an alert rule for a Data Transmission Service (DTS) task.
//
// Description:
//
// DTS provides the following metrics for DTS tasks:************
//
//   - **Latency**: DTS monitors the latency of a DTS task. If the latency of the task exceeds the specified threshold, an alert is triggered. The threshold is specified in units of seconds.
//
//   - **Status**: DTS monitors the status of a DTS task. If the state of the task changes to **Error*	- or **Restore**, an alert is triggered.
//
//   - **Full Timeout**: DTS monitors the duration of a DTS task. If the duration of the task exceeds the specified threshold, an alert is triggered. The threshold is specified in units of hours.
//
// @param request - CreateJobMonitorRuleRequest
//
// @return CreateJobMonitorRuleResponse
func (client *Client) CreateJobMonitorRule(request *CreateJobMonitorRuleRequest) (_result *CreateJobMonitorRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateJobMonitorRuleResponse{}
	_body, _err := client.CreateJobMonitorRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Purchases a data migration instance.
//
// @param request - CreateMigrationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMigrationJobResponse
func (client *Client) CreateMigrationJobWithOptions(request *CreateMigrationJobRequest, runtime *dara.RuntimeOptions) (_result *CreateMigrationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.MigrationJobClass) {
		query["MigrationJobClass"] = request.MigrationJobClass
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMigrationJob"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMigrationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Purchases a data migration instance.
//
// @param request - CreateMigrationJobRequest
//
// @return CreateMigrationJobResponse
func (client *Client) CreateMigrationJob(request *CreateMigrationJobRequest) (_result *CreateMigrationJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMigrationJobResponse{}
	_body, _err := client.CreateMigrationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a reverse task for a data synchronization or migration task.
//
// Description:
//
// 调用接口创建的反向任务会立即进行预检查，预检查通过后会进行增量数据采集，增量数据写入模块不会运行（需要调用**StartReverseWriter**接口运行）。
//
// > 创建的反向任务固定为同步任务，且只有增量写入模块。
//
// @param request - CreateReverseDtsJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateReverseDtsJobResponse
func (client *Client) CreateReverseDtsJobWithOptions(request *CreateReverseDtsJobRequest, runtime *dara.RuntimeOptions) (_result *CreateReverseDtsJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ShardPassword) {
		query["ShardPassword"] = request.ShardPassword
	}

	if !dara.IsNil(request.ShardUsername) {
		query["ShardUsername"] = request.ShardUsername
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateReverseDtsJob"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateReverseDtsJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a reverse task for a data synchronization or migration task.
//
// Description:
//
// 调用接口创建的反向任务会立即进行预检查，预检查通过后会进行增量数据采集，增量数据写入模块不会运行（需要调用**StartReverseWriter**接口运行）。
//
// > 创建的反向任务固定为同步任务，且只有增量写入模块。
//
// @param request - CreateReverseDtsJobRequest
//
// @return CreateReverseDtsJobResponse
func (client *Client) CreateReverseDtsJob(request *CreateReverseDtsJobRequest) (_result *CreateReverseDtsJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateReverseDtsJobResponse{}
	_body, _err := client.CreateReverseDtsJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Purchases a change tracking instance.
//
// @param request - CreateSubscriptionInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSubscriptionInstanceResponse
func (client *Client) CreateSubscriptionInstanceWithOptions(request *CreateSubscriptionInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateSubscriptionInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	if !dara.IsNil(request.SourceEndpoint) {
		query["SourceEndpoint"] = request.SourceEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSubscriptionInstance"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSubscriptionInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Purchases a change tracking instance.
//
// @param request - CreateSubscriptionInstanceRequest
//
// @return CreateSubscriptionInstanceResponse
func (client *Client) CreateSubscriptionInstance(request *CreateSubscriptionInstanceRequest) (_result *CreateSubscriptionInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSubscriptionInstanceResponse{}
	_body, _err := client.CreateSubscriptionInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data synchronization instance.
//
// @param request - CreateSynchronizationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSynchronizationJobResponse
func (client *Client) CreateSynchronizationJobWithOptions(request *CreateSynchronizationJobRequest, runtime *dara.RuntimeOptions) (_result *CreateSynchronizationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceCount) {
		query["DBInstanceCount"] = request.DBInstanceCount
	}

	if !dara.IsNil(request.DestRegion) {
		query["DestRegion"] = request.DestRegion
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SourceRegion) {
		query["SourceRegion"] = request.SourceRegion
	}

	if !dara.IsNil(request.SynchronizationJobClass) {
		query["SynchronizationJobClass"] = request.SynchronizationJobClass
	}

	if !dara.IsNil(request.Topology) {
		query["Topology"] = request.Topology
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	if !dara.IsNil(request.NetworkType) {
		query["networkType"] = request.NetworkType
	}

	if !dara.IsNil(request.DestinationEndpoint) {
		query["DestinationEndpoint"] = request.DestinationEndpoint
	}

	if !dara.IsNil(request.SourceEndpoint) {
		query["SourceEndpoint"] = request.SourceEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSynchronizationJob"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSynchronizationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data synchronization instance.
//
// @param request - CreateSynchronizationJobRequest
//
// @return CreateSynchronizationJobResponse
func (client *Client) CreateSynchronizationJob(request *CreateSynchronizationJobRequest) (_result *CreateSynchronizationJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSynchronizationJobResponse{}
	_body, _err := client.CreateSynchronizationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a consumer group.
//
// @param request - DeleteConsumerChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConsumerChannelResponse
func (client *Client) DeleteConsumerChannelWithOptions(request *DeleteConsumerChannelRequest, runtime *dara.RuntimeOptions) (_result *DeleteConsumerChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerGroupId) {
		query["ConsumerGroupId"] = request.ConsumerGroupId
	}

	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConsumerChannel"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConsumerChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a consumer group.
//
// @param request - DeleteConsumerChannelRequest
//
// @return DeleteConsumerChannelResponse
func (client *Client) DeleteConsumerChannel(request *DeleteConsumerChannelRequest) (_result *DeleteConsumerChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteConsumerChannelResponse{}
	_body, _err := client.DeleteConsumerChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a consumer group.
//
// @param request - DeleteConsumerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConsumerGroupResponse
func (client *Client) DeleteConsumerGroupWithOptions(request *DeleteConsumerGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteConsumerGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ConsumerGroupID) {
		query["ConsumerGroupID"] = request.ConsumerGroupID
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SubscriptionInstanceId) {
		query["SubscriptionInstanceId"] = request.SubscriptionInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConsumerGroup"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConsumerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a consumer group.
//
// @param request - DeleteConsumerGroupRequest
//
// @return DeleteConsumerGroupResponse
func (client *Client) DeleteConsumerGroup(request *DeleteConsumerGroupRequest) (_result *DeleteConsumerGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteConsumerGroupResponse{}
	_body, _err := client.DeleteConsumerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data migration, data synchronization, or change tracking task.
//
// @param request - DeleteDtsJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDtsJobResponse
func (client *Client) DeleteDtsJobWithOptions(request *DeleteDtsJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteDtsJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !dara.IsNil(request.ZeroEtlJob) {
		query["ZeroEtlJob"] = request.ZeroEtlJob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDtsJob"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDtsJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data migration, data synchronization, or change tracking task.
//
// @param request - DeleteDtsJobRequest
//
// @return DeleteDtsJobResponse
func (client *Client) DeleteDtsJob(request *DeleteDtsJobRequest) (_result *DeleteDtsJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDtsJobResponse{}
	_body, _err := client.DeleteDtsJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes multiple data migration, data synchronization, or change tracking tasks.
//
// @param request - DeleteDtsJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDtsJobsResponse
func (client *Client) DeleteDtsJobsWithOptions(request *DeleteDtsJobsRequest, runtime *dara.RuntimeOptions) (_result *DeleteDtsJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobIds) {
		query["DtsJobIds"] = request.DtsJobIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ZeroEtlJob) {
		query["ZeroEtlJob"] = request.ZeroEtlJob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDtsJobs"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDtsJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple data migration, data synchronization, or change tracking tasks.
//
// @param request - DeleteDtsJobsRequest
//
// @return DeleteDtsJobsResponse
func (client *Client) DeleteDtsJobs(request *DeleteDtsJobsRequest) (_result *DeleteDtsJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDtsJobsResponse{}
	_body, _err := client.DeleteDtsJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a data migration instance.
//
// @param request - DeleteMigrationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMigrationJobResponse
func (client *Client) DeleteMigrationJobWithOptions(request *DeleteMigrationJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteMigrationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.MigrationJobId) {
		query["MigrationJobId"] = request.MigrationJobId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMigrationJob"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMigrationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a data migration instance.
//
// @param request - DeleteMigrationJobRequest
//
// @return DeleteMigrationJobResponse
func (client *Client) DeleteMigrationJob(request *DeleteMigrationJobRequest) (_result *DeleteMigrationJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMigrationJobResponse{}
	_body, _err := client.DeleteMigrationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a change tracking instance.
//
// @param request - DeleteSubscriptionInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSubscriptionInstanceResponse
func (client *Client) DeleteSubscriptionInstanceWithOptions(request *DeleteSubscriptionInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteSubscriptionInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SubscriptionInstanceId) {
		query["SubscriptionInstanceId"] = request.SubscriptionInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSubscriptionInstance"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSubscriptionInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a change tracking instance.
//
// @param request - DeleteSubscriptionInstanceRequest
//
// @return DeleteSubscriptionInstanceResponse
func (client *Client) DeleteSubscriptionInstance(request *DeleteSubscriptionInstanceRequest) (_result *DeleteSubscriptionInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSubscriptionInstanceResponse{}
	_body, _err := client.DeleteSubscriptionInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a data synchronization instance.
//
// @param request - DeleteSynchronizationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSynchronizationJobResponse
func (client *Client) DeleteSynchronizationJobWithOptions(request *DeleteSynchronizationJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteSynchronizationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SynchronizationJobId) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSynchronizationJob"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSynchronizationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a data synchronization instance.
//
// @param request - DeleteSynchronizationJobRequest
//
// @return DeleteSynchronizationJobResponse
func (client *Client) DeleteSynchronizationJob(request *DeleteSynchronizationJobRequest) (_result *DeleteSynchronizationJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSynchronizationJobResponse{}
	_body, _err := client.DeleteSynchronizationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询数据投递链路store账号
//
// @param request - DescribeChannelAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelAccountResponse
func (client *Client) DescribeChannelAccountWithOptions(request *DescribeChannelAccountRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannelAccount"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据投递链路store账号
//
// @param request - DescribeChannelAccountRequest
//
// @return DescribeChannelAccountResponse
func (client *Client) DescribeChannelAccount(request *DescribeChannelAccountRequest) (_result *DescribeChannelAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeChannelAccountResponse{}
	_body, _err := client.DescribeChannelAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 请求所有数据校验任务数据
//
// @param request - DescribeCheckJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCheckJobsResponse
func (client *Client) DescribeCheckJobsWithOptions(request *DescribeCheckJobsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCheckJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckJobId) {
		query["CheckJobId"] = request.CheckJobId
	}

	if !dara.IsNil(request.CheckType) {
		query["CheckType"] = request.CheckType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCheckJobs"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCheckJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 请求所有数据校验任务数据
//
// @param request - DescribeCheckJobsRequest
//
// @return DescribeCheckJobsResponse
func (client *Client) DescribeCheckJobs(request *DescribeCheckJobsRequest) (_result *DescribeCheckJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCheckJobsResponse{}
	_body, _err := client.DescribeCheckJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries operation logs of a Data Transmission Service (DTS) dedicated cluster.
//
// @param request - DescribeClusterOperateLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterOperateLogsResponse
func (client *Client) DescribeClusterOperateLogsWithOptions(request *DescribeClusterOperateLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterOperateLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		body["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DedicatedClusterId) {
		body["DedicatedClusterId"] = request.DedicatedClusterId
	}

	if !dara.IsNil(request.DtsJobId) {
		body["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerID) {
		body["OwnerID"] = request.OwnerID
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterOperateLogs"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterOperateLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries operation logs of a Data Transmission Service (DTS) dedicated cluster.
//
// @param request - DescribeClusterOperateLogsRequest
//
// @return DescribeClusterOperateLogsResponse
func (client *Client) DescribeClusterOperateLogs(request *DescribeClusterOperateLogsRequest) (_result *DescribeClusterOperateLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeClusterOperateLogsResponse{}
	_body, _err := client.DescribeClusterOperateLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resource usage of a cluster.
//
// @param request - DescribeClusterUsedUtilizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterUsedUtilizationResponse
func (client *Client) DescribeClusterUsedUtilizationWithOptions(request *DescribeClusterUsedUtilizationRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterUsedUtilizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		body["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DedicatedClusterId) {
		body["DedicatedClusterId"] = request.DedicatedClusterId
	}

	if !dara.IsNil(request.DtsJobId) {
		body["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.Env) {
		body["Env"] = request.Env
	}

	if !dara.IsNil(request.MetricType) {
		body["MetricType"] = request.MetricType
	}

	if !dara.IsNil(request.OwnerID) {
		body["OwnerID"] = request.OwnerID
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityToken) {
		body["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterUsedUtilization"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterUsedUtilizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource usage of a cluster.
//
// @param request - DescribeClusterUsedUtilizationRequest
//
// @return DescribeClusterUsedUtilizationResponse
func (client *Client) DescribeClusterUsedUtilization(request *DescribeClusterUsedUtilizationRequest) (_result *DescribeClusterUsedUtilizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeClusterUsedUtilizationResponse{}
	_body, _err := client.DescribeClusterUsedUtilizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the connectivity of Data Transmission Service (DTS) servers to the source and destination databases.
//
// @param request - DescribeConnectionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConnectionStatusResponse
func (client *Client) DescribeConnectionStatusWithOptions(request *DescribeConnectionStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeConnectionStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestinationEndpointArchitecture) {
		query["DestinationEndpointArchitecture"] = request.DestinationEndpointArchitecture
	}

	if !dara.IsNil(request.DestinationEndpointDatabaseName) {
		query["DestinationEndpointDatabaseName"] = request.DestinationEndpointDatabaseName
	}

	if !dara.IsNil(request.DestinationEndpointEngineName) {
		query["DestinationEndpointEngineName"] = request.DestinationEndpointEngineName
	}

	if !dara.IsNil(request.DestinationEndpointIP) {
		query["DestinationEndpointIP"] = request.DestinationEndpointIP
	}

	if !dara.IsNil(request.DestinationEndpointInstanceID) {
		query["DestinationEndpointInstanceID"] = request.DestinationEndpointInstanceID
	}

	if !dara.IsNil(request.DestinationEndpointInstanceType) {
		query["DestinationEndpointInstanceType"] = request.DestinationEndpointInstanceType
	}

	if !dara.IsNil(request.DestinationEndpointOracleSID) {
		query["DestinationEndpointOracleSID"] = request.DestinationEndpointOracleSID
	}

	if !dara.IsNil(request.DestinationEndpointPassword) {
		query["DestinationEndpointPassword"] = request.DestinationEndpointPassword
	}

	if !dara.IsNil(request.DestinationEndpointPort) {
		query["DestinationEndpointPort"] = request.DestinationEndpointPort
	}

	if !dara.IsNil(request.DestinationEndpointRegion) {
		query["DestinationEndpointRegion"] = request.DestinationEndpointRegion
	}

	if !dara.IsNil(request.DestinationEndpointUserName) {
		query["DestinationEndpointUserName"] = request.DestinationEndpointUserName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SourceEndpointArchitecture) {
		query["SourceEndpointArchitecture"] = request.SourceEndpointArchitecture
	}

	if !dara.IsNil(request.SourceEndpointDatabaseName) {
		query["SourceEndpointDatabaseName"] = request.SourceEndpointDatabaseName
	}

	if !dara.IsNil(request.SourceEndpointEngineName) {
		query["SourceEndpointEngineName"] = request.SourceEndpointEngineName
	}

	if !dara.IsNil(request.SourceEndpointIP) {
		query["SourceEndpointIP"] = request.SourceEndpointIP
	}

	if !dara.IsNil(request.SourceEndpointInstanceID) {
		query["SourceEndpointInstanceID"] = request.SourceEndpointInstanceID
	}

	if !dara.IsNil(request.SourceEndpointInstanceType) {
		query["SourceEndpointInstanceType"] = request.SourceEndpointInstanceType
	}

	if !dara.IsNil(request.SourceEndpointOracleSID) {
		query["SourceEndpointOracleSID"] = request.SourceEndpointOracleSID
	}

	if !dara.IsNil(request.SourceEndpointPassword) {
		query["SourceEndpointPassword"] = request.SourceEndpointPassword
	}

	if !dara.IsNil(request.SourceEndpointPort) {
		query["SourceEndpointPort"] = request.SourceEndpointPort
	}

	if !dara.IsNil(request.SourceEndpointRegion) {
		query["SourceEndpointRegion"] = request.SourceEndpointRegion
	}

	if !dara.IsNil(request.SourceEndpointUserName) {
		query["SourceEndpointUserName"] = request.SourceEndpointUserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeConnectionStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeConnectionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the connectivity of Data Transmission Service (DTS) servers to the source and destination databases.
//
// @param request - DescribeConnectionStatusRequest
//
// @return DescribeConnectionStatusResponse
func (client *Client) DescribeConnectionStatus(request *DescribeConnectionStatusRequest) (_result *DescribeConnectionStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeConnectionStatusResponse{}
	_body, _err := client.DescribeConnectionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of a consumer group, including the consumer group ID, consumer group name, username, and message latency.
//
// @param request - DescribeConsumerChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConsumerChannelResponse
func (client *Client) DescribeConsumerChannelWithOptions(request *DescribeConsumerChannelRequest, runtime *dara.RuntimeOptions) (_result *DescribeConsumerChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentChannelId) {
		query["ParentChannelId"] = request.ParentChannelId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeConsumerChannel"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeConsumerChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of a consumer group, including the consumer group ID, consumer group name, username, and message latency.
//
// @param request - DescribeConsumerChannelRequest
//
// @return DescribeConsumerChannelResponse
func (client *Client) DescribeConsumerChannel(request *DescribeConsumerChannelRequest) (_result *DescribeConsumerChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeConsumerChannelResponse{}
	_body, _err := client.DescribeConsumerChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of consumer groups in a change tracking instance.
//
// @param request - DescribeConsumerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConsumerGroupResponse
func (client *Client) DescribeConsumerGroupWithOptions(request *DescribeConsumerGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeConsumerGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SubscriptionInstanceId) {
		query["SubscriptionInstanceId"] = request.SubscriptionInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeConsumerGroup"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeConsumerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of consumer groups in a change tracking instance.
//
// @param request - DescribeConsumerGroupRequest
//
// @return DescribeConsumerGroupResponse
func (client *Client) DescribeConsumerGroup(request *DescribeConsumerGroupRequest) (_result *DescribeConsumerGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeConsumerGroupResponse{}
	_body, _err := client.DescribeConsumerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the CIDR blocks of DTS servers.
//
// @param request - DescribeDTSIPRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDTSIPResponse
func (client *Client) DescribeDTSIPWithOptions(request *DescribeDTSIPRequest, runtime *dara.RuntimeOptions) (_result *DescribeDTSIPResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestinationEndpointRegion) {
		query["DestinationEndpointRegion"] = request.DestinationEndpointRegion
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SourceEndpointRegion) {
		query["SourceEndpointRegion"] = request.SourceEndpointRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDTSIP"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDTSIPResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the CIDR blocks of DTS servers.
//
// @param request - DescribeDTSIPRequest
//
// @return DescribeDTSIPResponse
func (client *Client) DescribeDTSIP(request *DescribeDTSIPRequest) (_result *DescribeDTSIPResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDTSIPResponse{}
	_body, _err := client.DescribeDTSIPWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the download URL of the data consistency verification report.
//
// @param request - DescribeDataCheckReportUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataCheckReportUrlResponse
func (client *Client) DescribeDataCheckReportUrlWithOptions(request *DescribeDataCheckReportUrlRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataCheckReportUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckType) {
		query["CheckType"] = request.CheckType
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TbName) {
		query["TbName"] = request.TbName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataCheckReportUrl"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataCheckReportUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the download URL of the data consistency verification report.
//
// @param request - DescribeDataCheckReportUrlRequest
//
// @return DescribeDataCheckReportUrlResponse
func (client *Client) DescribeDataCheckReportUrl(request *DescribeDataCheckReportUrlRequest) (_result *DescribeDataCheckReportUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataCheckReportUrlResponse{}
	_body, _err := client.DescribeDataCheckReportUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a data verification task.
//
// @param request - DescribeDataCheckTableDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataCheckTableDetailsResponse
func (client *Client) DescribeDataCheckTableDetailsWithOptions(request *DescribeDataCheckTableDetailsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataCheckTableDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckType) {
		query["CheckType"] = request.CheckType
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataCheckTableDetails"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataCheckTableDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a data verification task.
//
// @param request - DescribeDataCheckTableDetailsRequest
//
// @return DescribeDataCheckTableDetailsResponse
func (client *Client) DescribeDataCheckTableDetails(request *DescribeDataCheckTableDetailsRequest) (_result *DescribeDataCheckTableDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataCheckTableDetailsResponse{}
	_body, _err := client.DescribeDataCheckTableDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about inconsistent data in the data verification task.
//
// @param request - DescribeDataCheckTableDiffDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataCheckTableDiffDetailsResponse
func (client *Client) DescribeDataCheckTableDiffDetailsWithOptions(request *DescribeDataCheckTableDiffDetailsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataCheckTableDiffDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckType) {
		query["CheckType"] = request.CheckType
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TbName) {
		query["TbName"] = request.TbName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataCheckTableDiffDetails"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataCheckTableDiffDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about inconsistent data in the data verification task.
//
// @param request - DescribeDataCheckTableDiffDetailsRequest
//
// @return DescribeDataCheckTableDiffDetailsResponse
func (client *Client) DescribeDataCheckTableDiffDetails(request *DescribeDataCheckTableDiffDetailsRequest) (_result *DescribeDataCheckTableDiffDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataCheckTableDiffDetailsResponse{}
	_body, _err := client.DescribeDataCheckTableDiffDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a dedicated cluster.
//
// @param request - DescribeDedicatedClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDedicatedClusterResponse
func (client *Client) DescribeDedicatedClusterWithOptions(request *DescribeDedicatedClusterRequest, runtime *dara.RuntimeOptions) (_result *DescribeDedicatedClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DedicatedClusterId) {
		query["DedicatedClusterId"] = request.DedicatedClusterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDedicatedCluster"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDedicatedClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a dedicated cluster.
//
// @param request - DescribeDedicatedClusterRequest
//
// @return DescribeDedicatedClusterResponse
func (client *Client) DescribeDedicatedCluster(request *DescribeDedicatedClusterRequest) (_result *DescribeDedicatedClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDedicatedClusterResponse{}
	_body, _err := client.DescribeDedicatedClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an alert rule.
//
// @param request - DescribeDedicatedClusterMonitorRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDedicatedClusterMonitorRuleResponse
func (client *Client) DescribeDedicatedClusterMonitorRuleWithOptions(request *DescribeDedicatedClusterMonitorRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribeDedicatedClusterMonitorRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DedicatedClusterId) {
		query["DedicatedClusterId"] = request.DedicatedClusterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDedicatedClusterMonitorRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDedicatedClusterMonitorRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an alert rule.
//
// @param request - DescribeDedicatedClusterMonitorRuleRequest
//
// @return DescribeDedicatedClusterMonitorRuleResponse
func (client *Client) DescribeDedicatedClusterMonitorRule(request *DescribeDedicatedClusterMonitorRuleRequest) (_result *DescribeDedicatedClusterMonitorRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDedicatedClusterMonitorRuleResponse{}
	_body, _err := client.DescribeDedicatedClusterMonitorRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看工作流任务结果
//
// @param request - DescribeDocParserJobResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDocParserJobResultResponse
func (client *Client) DescribeDocParserJobResultWithOptions(request *DescribeDocParserJobResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeDocParserJobResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDocParserJobResult"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDocParserJobResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看工作流任务结果
//
// @param request - DescribeDocParserJobResultRequest
//
// @return DescribeDocParserJobResultResponse
func (client *Client) DescribeDocParserJobResult(request *DescribeDocParserJobResultRequest) (_result *DescribeDocParserJobResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDocParserJobResultResponse{}
	_body, _err := client.DescribeDocParserJobResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看工作流任务结果
//
// @param request - DescribeDocParserJobStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDocParserJobStatusResponse
func (client *Client) DescribeDocParserJobStatusWithOptions(request *DescribeDocParserJobStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeDocParserJobStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDocParserJobStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDocParserJobStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看工作流任务结果
//
// @param request - DescribeDocParserJobStatusRequest
//
// @return DescribeDocParserJobStatusResponse
func (client *Client) DescribeDocParserJobStatus(request *DescribeDocParserJobStatusRequest) (_result *DescribeDocParserJobStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDocParserJobStatusResponse{}
	_body, _err := client.DescribeDocParserJobStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of extract, transform, and load (ETL) tasks.
//
// @param request - DescribeDtsEtlJobVersionInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDtsEtlJobVersionInfoResponse
func (client *Client) DescribeDtsEtlJobVersionInfoWithOptions(request *DescribeDtsEtlJobVersionInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDtsEtlJobVersionInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDtsEtlJobVersionInfo"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDtsEtlJobVersionInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of extract, transform, and load (ETL) tasks.
//
// @param request - DescribeDtsEtlJobVersionInfoRequest
//
// @return DescribeDtsEtlJobVersionInfoResponse
func (client *Client) DescribeDtsEtlJobVersionInfo(request *DescribeDtsEtlJobVersionInfoRequest) (_result *DescribeDtsEtlJobVersionInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDtsEtlJobVersionInfoResponse{}
	_body, _err := client.DescribeDtsEtlJobVersionInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询DTS任务配置
//
// @param request - DescribeDtsJobConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDtsJobConfigResponse
func (client *Client) DescribeDtsJobConfigWithOptions(request *DescribeDtsJobConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDtsJobConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.ForAcceleration) {
		query["ForAcceleration"] = request.ForAcceleration
	}

	if !dara.IsNil(request.Module) {
		query["Module"] = request.Module
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDtsJobConfig"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDtsJobConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询DTS任务配置
//
// @param request - DescribeDtsJobConfigRequest
//
// @return DescribeDtsJobConfigResponse
func (client *Client) DescribeDtsJobConfig(request *DescribeDtsJobConfigRequest) (_result *DescribeDtsJobConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDtsJobConfigResponse{}
	_body, _err := client.DescribeDtsJobConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The latency of incremental data migration or synchronization.
//
// \\\\\\\\\\>  If you query data migration tasks, the unit of this parameter is milliseconds. If you query data synchronization tasks, the unit of this parameter is seconds.
//
// @param request - DescribeDtsJobDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDtsJobDetailResponse
func (client *Client) DescribeDtsJobDetailWithOptions(request *DescribeDtsJobDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeDtsJobDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsInstanceID) {
		query["DtsInstanceID"] = request.DtsInstanceID
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SyncSubJobHistory) {
		query["SyncSubJobHistory"] = request.SyncSubJobHistory
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !dara.IsNil(request.ZeroEtlJob) {
		query["ZeroEtlJob"] = request.ZeroEtlJob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDtsJobDetail"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDtsJobDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The latency of incremental data migration or synchronization.
//
// \\\\\\\\\\>  If you query data migration tasks, the unit of this parameter is milliseconds. If you query data synchronization tasks, the unit of this parameter is seconds.
//
// @param request - DescribeDtsJobDetailRequest
//
// @return DescribeDtsJobDetailResponse
func (client *Client) DescribeDtsJobDetail(request *DescribeDtsJobDetailRequest) (_result *DescribeDtsJobDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDtsJobDetailResponse{}
	_body, _err := client.DescribeDtsJobDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of Data Transmission Service (DTS) tasks and the details of each task.
//
// Description:
//
// ## Debugging
//
// [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code of the operation for different SDKs.](https://api.aliyun.com/#product=Dts\\&api=DescribeDtsJobs\\&type=RPC\\&version=2020-01-01)
//
// @param request - DescribeDtsJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDtsJobsResponse
func (client *Client) DescribeDtsJobsWithOptions(request *DescribeDtsJobsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDtsJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DedicatedClusterId) {
		query["DedicatedClusterId"] = request.DedicatedClusterId
	}

	if !dara.IsNil(request.DestProductType) {
		query["DestProductType"] = request.DestProductType
	}

	if !dara.IsNil(request.DtsBisLabel) {
		query["DtsBisLabel"] = request.DtsBisLabel
	}

	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.OrderColumn) {
		query["OrderColumn"] = request.OrderColumn
	}

	if !dara.IsNil(request.OrderDirection) {
		query["OrderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SrcProductType) {
		query["SrcProductType"] = request.SrcProductType
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.WithoutDbList) {
		query["WithoutDbList"] = request.WithoutDbList
	}

	if !dara.IsNil(request.ZeroEtlJob) {
		query["ZeroEtlJob"] = request.ZeroEtlJob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDtsJobs"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDtsJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of Data Transmission Service (DTS) tasks and the details of each task.
//
// Description:
//
// ## Debugging
//
// [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code of the operation for different SDKs.](https://api.aliyun.com/#product=Dts\\&api=DescribeDtsJobs\\&type=RPC\\&version=2020-01-01)
//
// @param request - DescribeDtsJobsRequest
//
// @return DescribeDtsJobsResponse
func (client *Client) DescribeDtsJobs(request *DescribeDtsJobsRequest) (_result *DescribeDtsJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDtsJobsResponse{}
	_body, _err := client.DescribeDtsJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the logs of a data migration or synchronization task.
//
// @param request - DescribeDtsServiceLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDtsServiceLogResponse
func (client *Client) DescribeDtsServiceLogWithOptions(request *DescribeDtsServiceLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeDtsServiceLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.SubJobType) {
		query["SubJobType"] = request.SubJobType
	}

	if !dara.IsNil(request.ZeroEtlJob) {
		query["ZeroEtlJob"] = request.ZeroEtlJob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDtsServiceLog"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDtsServiceLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logs of a data migration or synchronization task.
//
// @param request - DescribeDtsServiceLogRequest
//
// @return DescribeDtsServiceLogResponse
func (client *Client) DescribeDtsServiceLog(request *DescribeDtsServiceLogRequest) (_result *DescribeDtsServiceLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDtsServiceLogResponse{}
	_body, _err := client.DescribeDtsServiceLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the task that changes the database connection settings.
//
// @param request - DescribeEndpointSwitchStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEndpointSwitchStatusResponse
func (client *Client) DescribeEndpointSwitchStatusWithOptions(request *DescribeEndpointSwitchStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeEndpointSwitchStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEndpointSwitchStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEndpointSwitchStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the task that changes the database connection settings.
//
// @param request - DescribeEndpointSwitchStatusRequest
//
// @return DescribeEndpointSwitchStatusResponse
func (client *Client) DescribeEndpointSwitchStatus(request *DescribeEndpointSwitchStatusRequest) (_result *DescribeEndpointSwitchStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEndpointSwitchStatusResponse{}
	_body, _err := client.DescribeEndpointSwitchStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the logs of extract, transform, and load (ETL) tasks.
//
// @param request - DescribeEtlJobLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEtlJobLogsResponse
func (client *Client) DescribeEtlJobLogsWithOptions(request *DescribeEtlJobLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeEtlJobLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEtlJobLogs"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEtlJobLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logs of extract, transform, and load (ETL) tasks.
//
// @param request - DescribeEtlJobLogsRequest
//
// @return DescribeEtlJobLogsResponse
func (client *Client) DescribeEtlJobLogs(request *DescribeEtlJobLogsRequest) (_result *DescribeEtlJobLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEtlJobLogsResponse{}
	_body, _err := client.DescribeEtlJobLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries full data migration tasks.
//
// @param request - DescribeFullProcessListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFullProcessListResponse
func (client *Client) DescribeFullProcessListWithOptions(request *DescribeFullProcessListRequest, runtime *dara.RuntimeOptions) (_result *DescribeFullProcessListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ZeroEtlJob) {
		query["ZeroEtlJob"] = request.ZeroEtlJob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFullProcessList"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFullProcessListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries full data migration tasks.
//
// @param request - DescribeFullProcessListRequest
//
// @return DescribeFullProcessListResponse
func (client *Client) DescribeFullProcessList(request *DescribeFullProcessListRequest) (_result *DescribeFullProcessListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFullProcessListResponse{}
	_body, _err := client.DescribeFullProcessListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询GAD实例列表
//
// @param request - DescribeGadInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGadInstancesResponse
func (client *Client) DescribeGadInstancesWithOptions(request *DescribeGadInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeGadInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbEngineTypes) {
		query["DbEngineTypes"] = request.DbEngineTypes
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.MasterDbInstanceId) {
		query["MasterDbInstanceId"] = request.MasterDbInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SlaveDbInstanceId) {
		query["SlaveDbInstanceId"] = request.SlaveDbInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGadInstances"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGadInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询GAD实例列表
//
// @param request - DescribeGadInstancesRequest
//
// @return DescribeGadInstancesResponse
func (client *Client) DescribeGadInstances(request *DescribeGadInstancesRequest) (_result *DescribeGadInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGadInstancesResponse{}
	_body, _err := client.DescribeGadInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of initial data synchronization, including the information about the schemas and historical data of the object to be synchronized.
//
// @param request - DescribeInitializationStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInitializationStatusResponse
func (client *Client) DescribeInitializationStatusWithOptions(request *DescribeInitializationStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeInitializationStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SynchronizationJobId) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInitializationStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInitializationStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of initial data synchronization, including the information about the schemas and historical data of the object to be synchronized.
//
// @param request - DescribeInitializationStatusRequest
//
// @return DescribeInitializationStatusResponse
func (client *Client) DescribeInitializationStatus(request *DescribeInitializationStatusRequest) (_result *DescribeInitializationStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInitializationStatusResponse{}
	_body, _err := client.DescribeInitializationStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the monitoring rules of a Data Transmission Service (DTS) task.
//
// @param request - DescribeJobMonitorRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeJobMonitorRuleResponse
func (client *Client) DescribeJobMonitorRuleWithOptions(request *DescribeJobMonitorRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribeJobMonitorRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeJobMonitorRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeJobMonitorRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the monitoring rules of a Data Transmission Service (DTS) task.
//
// @param request - DescribeJobMonitorRuleRequest
//
// @return DescribeJobMonitorRuleResponse
func (client *Client) DescribeJobMonitorRule(request *DescribeJobMonitorRuleRequest) (_result *DescribeJobMonitorRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeJobMonitorRuleResponse{}
	_body, _err := client.DescribeJobMonitorRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the metrics of a cluster.
//
// @param request - DescribeMetricListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricListResponse
func (client *Client) DescribeMetricListWithOptions(request *DescribeMetricListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		body["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DtsJobId) {
		body["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Env) {
		body["Env"] = request.Env
	}

	if !dara.IsNil(request.MetricName) {
		body["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.MetricType) {
		body["MetricType"] = request.MetricType
	}

	if !dara.IsNil(request.OwnerID) {
		body["OwnerID"] = request.OwnerID
	}

	if !dara.IsNil(request.Param) {
		body["Param"] = request.Param
	}

	if !dara.IsNil(request.Period) {
		body["Period"] = request.Period
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricList"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metrics of a cluster.
//
// @param request - DescribeMetricListRequest
//
// @return DescribeMetricListResponse
func (client *Client) DescribeMetricList(request *DescribeMetricListRequest) (_result *DescribeMetricListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMetricListResponse{}
	_body, _err := client.DescribeMetricListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the alert settings of a data migration instance.
//
// @param request - DescribeMigrationJobAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMigrationJobAlertResponse
func (client *Client) DescribeMigrationJobAlertWithOptions(request *DescribeMigrationJobAlertRequest, runtime *dara.RuntimeOptions) (_result *DescribeMigrationJobAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.MigrationJobId) {
		query["MigrationJobId"] = request.MigrationJobId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMigrationJobAlert"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMigrationJobAlertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the alert settings of a data migration instance.
//
// @param request - DescribeMigrationJobAlertRequest
//
// @return DescribeMigrationJobAlertResponse
func (client *Client) DescribeMigrationJobAlert(request *DescribeMigrationJobAlertRequest) (_result *DescribeMigrationJobAlertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMigrationJobAlertResponse{}
	_body, _err := client.DescribeMigrationJobAlertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a data migration task.
//
// @param request - DescribeMigrationJobDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMigrationJobDetailResponse
func (client *Client) DescribeMigrationJobDetailWithOptions(request *DescribeMigrationJobDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeMigrationJobDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.MigrationJobId) {
		query["MigrationJobId"] = request.MigrationJobId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.MigrationMode) {
		query["MigrationMode"] = request.MigrationMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMigrationJobDetail"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMigrationJobDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a data migration task.
//
// @param request - DescribeMigrationJobDetailRequest
//
// @return DescribeMigrationJobDetailResponse
func (client *Client) DescribeMigrationJobDetail(request *DescribeMigrationJobDetailRequest) (_result *DescribeMigrationJobDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMigrationJobDetailResponse{}
	_body, _err := client.DescribeMigrationJobDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a data migration task.
//
// @param request - DescribeMigrationJobStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMigrationJobStatusResponse
func (client *Client) DescribeMigrationJobStatusWithOptions(request *DescribeMigrationJobStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeMigrationJobStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.MigrationJobId) {
		query["MigrationJobId"] = request.MigrationJobId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMigrationJobStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMigrationJobStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a data migration task.
//
// @param request - DescribeMigrationJobStatusRequest
//
// @return DescribeMigrationJobStatusResponse
func (client *Client) DescribeMigrationJobStatus(request *DescribeMigrationJobStatusRequest) (_result *DescribeMigrationJobStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMigrationJobStatusResponse{}
	_body, _err := client.DescribeMigrationJobStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of data migration instances and the details of each instance.
//
// @param request - DescribeMigrationJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMigrationJobsResponse
func (client *Client) DescribeMigrationJobsWithOptions(request *DescribeMigrationJobsRequest, runtime *dara.RuntimeOptions) (_result *DescribeMigrationJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.MigrationJobName) {
		query["MigrationJobName"] = request.MigrationJobName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMigrationJobs"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMigrationJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of data migration instances and the details of each instance.
//
// @param request - DescribeMigrationJobsRequest
//
// @return DescribeMigrationJobsResponse
func (client *Client) DescribeMigrationJobs(request *DescribeMigrationJobsRequest) (_result *DescribeMigrationJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMigrationJobsResponse{}
	_body, _err := client.DescribeMigrationJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询预检查创建GAD订单任务结果
//
// @param request - DescribePreCheckCreateGadOrderResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePreCheckCreateGadOrderResultResponse
func (client *Client) DescribePreCheckCreateGadOrderResultWithOptions(request *DescribePreCheckCreateGadOrderResultRequest, runtime *dara.RuntimeOptions) (_result *DescribePreCheckCreateGadOrderResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePreCheckCreateGadOrderResult"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePreCheckCreateGadOrderResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询预检查创建GAD订单任务结果
//
// @param request - DescribePreCheckCreateGadOrderResultRequest
//
// @return DescribePreCheckCreateGadOrderResultResponse
func (client *Client) DescribePreCheckCreateGadOrderResult(request *DescribePreCheckCreateGadOrderResultRequest) (_result *DescribePreCheckCreateGadOrderResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePreCheckCreateGadOrderResultResponse{}
	_body, _err := client.DescribePreCheckCreateGadOrderResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a Data Transmission Service (DTS) subtask that performs precheck, schema migration, initial schema synchronization, full data migration, initial full data synchronization, incremental data migration, or incremental data synchronization.
//
// @param request - DescribePreCheckStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePreCheckStatusResponse
func (client *Client) DescribePreCheckStatusWithOptions(request *DescribePreCheckStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribePreCheckStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.JobCode) {
		query["JobCode"] = request.JobCode
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StructPhase) {
		query["StructPhase"] = request.StructPhase
	}

	if !dara.IsNil(request.StructType) {
		query["StructType"] = request.StructType
	}

	if !dara.IsNil(request.ZeroEtlJob) {
		query["ZeroEtlJob"] = request.ZeroEtlJob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePreCheckStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePreCheckStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a Data Transmission Service (DTS) subtask that performs precheck, schema migration, initial schema synchronization, full data migration, initial full data synchronization, incremental data migration, or incremental data synchronization.
//
// @param request - DescribePreCheckStatusRequest
//
// @return DescribePreCheckStatusResponse
func (client *Client) DescribePreCheckStatus(request *DescribePreCheckStatusRequest) (_result *DescribePreCheckStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePreCheckStatusResponse{}
	_body, _err := client.DescribePreCheckStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the alert settings of a change tracking instance.
//
// @param request - DescribeSubscriptionInstanceAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSubscriptionInstanceAlertResponse
func (client *Client) DescribeSubscriptionInstanceAlertWithOptions(request *DescribeSubscriptionInstanceAlertRequest, runtime *dara.RuntimeOptions) (_result *DescribeSubscriptionInstanceAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SubscriptionInstanceId) {
		query["SubscriptionInstanceId"] = request.SubscriptionInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSubscriptionInstanceAlert"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSubscriptionInstanceAlertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the alert settings of a change tracking instance.
//
// @param request - DescribeSubscriptionInstanceAlertRequest
//
// @return DescribeSubscriptionInstanceAlertResponse
func (client *Client) DescribeSubscriptionInstanceAlert(request *DescribeSubscriptionInstanceAlertRequest) (_result *DescribeSubscriptionInstanceAlertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSubscriptionInstanceAlertResponse{}
	_body, _err := client.DescribeSubscriptionInstanceAlertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a change tracking instance.
//
// @param request - DescribeSubscriptionInstanceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSubscriptionInstanceStatusResponse
func (client *Client) DescribeSubscriptionInstanceStatusWithOptions(request *DescribeSubscriptionInstanceStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeSubscriptionInstanceStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SubscriptionInstanceId) {
		query["SubscriptionInstanceId"] = request.SubscriptionInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSubscriptionInstanceStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSubscriptionInstanceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a change tracking instance.
//
// @param request - DescribeSubscriptionInstanceStatusRequest
//
// @return DescribeSubscriptionInstanceStatusResponse
func (client *Client) DescribeSubscriptionInstanceStatus(request *DescribeSubscriptionInstanceStatusRequest) (_result *DescribeSubscriptionInstanceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSubscriptionInstanceStatusResponse{}
	_body, _err := client.DescribeSubscriptionInstanceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of change tracking instances and the details of each instance.
//
// @param request - DescribeSubscriptionInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSubscriptionInstancesResponse
func (client *Client) DescribeSubscriptionInstancesWithOptions(request *DescribeSubscriptionInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSubscriptionInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SubscriptionInstanceName) {
		query["SubscriptionInstanceName"] = request.SubscriptionInstanceName
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSubscriptionInstances"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSubscriptionInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of change tracking instances and the details of each instance.
//
// @param request - DescribeSubscriptionInstancesRequest
//
// @return DescribeSubscriptionInstancesResponse
func (client *Client) DescribeSubscriptionInstances(request *DescribeSubscriptionInstancesRequest) (_result *DescribeSubscriptionInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSubscriptionInstancesResponse{}
	_body, _err := client.DescribeSubscriptionInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the subtasks in a distributed change tracking task for a PolarDB-X 1.0 instance.
//
// Description:
//
//	  When Data Transmission Service (DTS) tracks data changes from a PolarDB-X 1.0 instance, data is distributed across the attached ApsaraDB RDS for MySQL instances. DTS runs a subtask for each ApsaraDB RDS for MySQL instance. You can call this operation to query the details of the subtasks in a distributed change tracking task.
//
//		- You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the ID of the change tracking instance and the ID of the consumer group.
//
// @param tmpReq - DescribeSubscriptionMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSubscriptionMetaResponse
func (client *Client) DescribeSubscriptionMetaWithOptions(tmpReq *DescribeSubscriptionMetaRequest, runtime *dara.RuntimeOptions) (_result *DescribeSubscriptionMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeSubscriptionMetaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SubMigrationJobIds) {
		request.SubMigrationJobIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubMigrationJobIds, dara.String("SubMigrationJobIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Topics) {
		request.TopicsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Topics, dara.String("Topics"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Sid) {
		query["Sid"] = request.Sid
	}

	if !dara.IsNil(request.SubMigrationJobIdsShrink) {
		query["SubMigrationJobIds"] = request.SubMigrationJobIdsShrink
	}

	if !dara.IsNil(request.TopicsShrink) {
		query["Topics"] = request.TopicsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSubscriptionMeta"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSubscriptionMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the subtasks in a distributed change tracking task for a PolarDB-X 1.0 instance.
//
// Description:
//
//	  When Data Transmission Service (DTS) tracks data changes from a PolarDB-X 1.0 instance, data is distributed across the attached ApsaraDB RDS for MySQL instances. DTS runs a subtask for each ApsaraDB RDS for MySQL instance. You can call this operation to query the details of the subtasks in a distributed change tracking task.
//
//		- You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the ID of the change tracking instance and the ID of the consumer group.
//
// @param request - DescribeSubscriptionMetaRequest
//
// @return DescribeSubscriptionMetaResponse
func (client *Client) DescribeSubscriptionMeta(request *DescribeSubscriptionMetaRequest) (_result *DescribeSubscriptionMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSubscriptionMetaResponse{}
	_body, _err := client.DescribeSubscriptionMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看同步和迁移任务的增量写入延迟信息
//
// @param request - DescribeSyncStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSyncStatusResponse
func (client *Client) DescribeSyncStatusWithOptions(request *DescribeSyncStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeSyncStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSyncStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSyncStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看同步和迁移任务的增量写入延迟信息
//
// @param request - DescribeSyncStatusRequest
//
// @return DescribeSyncStatusResponse
func (client *Client) DescribeSyncStatus(request *DescribeSyncStatusRequest) (_result *DescribeSyncStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSyncStatusResponse{}
	_body, _err := client.DescribeSyncStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the alert settings of a data synchronization instance.
//
// @param request - DescribeSynchronizationJobAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSynchronizationJobAlertResponse
func (client *Client) DescribeSynchronizationJobAlertWithOptions(request *DescribeSynchronizationJobAlertRequest, runtime *dara.RuntimeOptions) (_result *DescribeSynchronizationJobAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !dara.IsNil(request.SynchronizationJobId) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSynchronizationJobAlert"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSynchronizationJobAlertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the alert settings of a data synchronization instance.
//
// @param request - DescribeSynchronizationJobAlertRequest
//
// @return DescribeSynchronizationJobAlertResponse
func (client *Client) DescribeSynchronizationJobAlert(request *DescribeSynchronizationJobAlertRequest) (_result *DescribeSynchronizationJobAlertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSynchronizationJobAlertResponse{}
	_body, _err := client.DescribeSynchronizationJobAlertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether image matching is enabled for a data synchronization instance.
//
// @param request - DescribeSynchronizationJobReplicatorCompareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSynchronizationJobReplicatorCompareResponse
func (client *Client) DescribeSynchronizationJobReplicatorCompareWithOptions(request *DescribeSynchronizationJobReplicatorCompareRequest, runtime *dara.RuntimeOptions) (_result *DescribeSynchronizationJobReplicatorCompareResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !dara.IsNil(request.SynchronizationJobId) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSynchronizationJobReplicatorCompare"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSynchronizationJobReplicatorCompareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether image matching is enabled for a data synchronization instance.
//
// @param request - DescribeSynchronizationJobReplicatorCompareRequest
//
// @return DescribeSynchronizationJobReplicatorCompareResponse
func (client *Client) DescribeSynchronizationJobReplicatorCompare(request *DescribeSynchronizationJobReplicatorCompareRequest) (_result *DescribeSynchronizationJobReplicatorCompareResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSynchronizationJobReplicatorCompareResponse{}
	_body, _err := client.DescribeSynchronizationJobReplicatorCompareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a data synchronization instance.
//
// @param request - DescribeSynchronizationJobStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSynchronizationJobStatusResponse
func (client *Client) DescribeSynchronizationJobStatusWithOptions(request *DescribeSynchronizationJobStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeSynchronizationJobStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !dara.IsNil(request.SynchronizationJobId) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSynchronizationJobStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSynchronizationJobStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a data synchronization instance.
//
// @param request - DescribeSynchronizationJobStatusRequest
//
// @return DescribeSynchronizationJobStatusResponse
func (client *Client) DescribeSynchronizationJobStatus(request *DescribeSynchronizationJobStatusRequest) (_result *DescribeSynchronizationJobStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSynchronizationJobStatusResponse{}
	_body, _err := client.DescribeSynchronizationJobStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of one or more data synchronization instances.
//
// @param request - DescribeSynchronizationJobStatusListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSynchronizationJobStatusListResponse
func (client *Client) DescribeSynchronizationJobStatusListWithOptions(request *DescribeSynchronizationJobStatusListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSynchronizationJobStatusListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SynchronizationJobIdListJsonStr) {
		query["SynchronizationJobIdListJsonStr"] = request.SynchronizationJobIdListJsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSynchronizationJobStatusList"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSynchronizationJobStatusListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of one or more data synchronization instances.
//
// @param request - DescribeSynchronizationJobStatusListRequest
//
// @return DescribeSynchronizationJobStatusListResponse
func (client *Client) DescribeSynchronizationJobStatusList(request *DescribeSynchronizationJobStatusListRequest) (_result *DescribeSynchronizationJobStatusListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSynchronizationJobStatusListResponse{}
	_body, _err := client.DescribeSynchronizationJobStatusListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of data synchronization instances and the details of each instance.
//
// @param request - DescribeSynchronizationJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSynchronizationJobsResponse
func (client *Client) DescribeSynchronizationJobsWithOptions(request *DescribeSynchronizationJobsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSynchronizationJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SynchronizationJobName) {
		query["SynchronizationJobName"] = request.SynchronizationJobName
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSynchronizationJobs"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSynchronizationJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of data synchronization instances and the details of each instance.
//
// @param request - DescribeSynchronizationJobsRequest
//
// @return DescribeSynchronizationJobsResponse
func (client *Client) DescribeSynchronizationJobs(request *DescribeSynchronizationJobsRequest) (_result *DescribeSynchronizationJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSynchronizationJobsResponse{}
	_body, _err := client.DescribeSynchronizationJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the task that changes the objects to be synchronized.
//
// @param request - DescribeSynchronizationObjectModifyStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSynchronizationObjectModifyStatusResponse
func (client *Client) DescribeSynchronizationObjectModifyStatusWithOptions(request *DescribeSynchronizationObjectModifyStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeSynchronizationObjectModifyStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSynchronizationObjectModifyStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSynchronizationObjectModifyStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the task that changes the objects to be synchronized.
//
// @param request - DescribeSynchronizationObjectModifyStatusRequest
//
// @return DescribeSynchronizationObjectModifyStatusResponse
func (client *Client) DescribeSynchronizationObjectModifyStatus(request *DescribeSynchronizationObjectModifyStatusRequest) (_result *DescribeSynchronizationObjectModifyStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSynchronizationObjectModifyStatusResponse{}
	_body, _err := client.DescribeSynchronizationObjectModifyStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all the tags added to a data migration, data synchronization, or change tracking instance.
//
// @param request - DescribeTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagKeysResponse
func (client *Client) DescribeTagKeysWithOptions(request *DescribeTagKeysRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTagKeys"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all the tags added to a data migration, data synchronization, or change tracking instance.
//
// @param request - DescribeTagKeysRequest
//
// @return DescribeTagKeysResponse
func (client *Client) DescribeTagKeys(request *DescribeTagKeysRequest) (_result *DescribeTagKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTagKeysResponse{}
	_body, _err := client.DescribeTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all the tag values of a tag bound to a data migration, data synchronization, or change tracking instance.
//
// @param request - DescribeTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagValuesResponse
func (client *Client) DescribeTagValuesWithOptions(request *DescribeTagValuesRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagValuesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTagValues"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all the tag values of a tag bound to a data migration, data synchronization, or change tracking instance.
//
// @param request - DescribeTagValuesRequest
//
// @return DescribeTagValuesResponse
func (client *Client) DescribeTagValues(request *DescribeTagValuesRequest) (_result *DescribeTagValuesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTagValuesResponse{}
	_body, _err := client.DescribeTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移除从角色
//
// @param request - DetachGadInstanceDbMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachGadInstanceDbMemberResponse
func (client *Client) DetachGadInstanceDbMemberWithOptions(request *DetachGadInstanceDbMemberRequest, runtime *dara.RuntimeOptions) (_result *DetachGadInstanceDbMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SlaveDbInstanceId) {
		query["SlaveDbInstanceId"] = request.SlaveDbInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachGadInstanceDbMember"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachGadInstanceDbMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移除从角色
//
// @param request - DetachGadInstanceDbMemberRequest
//
// @return DetachGadInstanceDbMemberResponse
func (client *Client) DetachGadInstanceDbMember(request *DetachGadInstanceDbMemberRequest) (_result *DetachGadInstanceDbMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachGadInstanceDbMemberResponse{}
	_body, _err := client.DetachGadInstanceDbMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initializes a built-in account on a node of an active geo-redundancy database cluster. Data Transmission Service (DTS) uses the built-in account to connect to the node and perform data synchronization tasks.
//
// Description:
//
//	  The node must be an ApsaraDB RDS for MySQL instance or a self-managed MySQL database that is connected over Cloud Enterprise Network (CEN).
//
//		- This operation is used to initialize the built-in account named rdsdt_dtsacct on a node of an active geo-redundancy database cluster. DTS uses this account to connect to the node and perform data synchronization tasks.
//
// @param request - InitDtsRdsInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitDtsRdsInstanceResponse
func (client *Client) InitDtsRdsInstanceWithOptions(request *InitDtsRdsInstanceRequest, runtime *dara.RuntimeOptions) (_result *InitDtsRdsInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.EndpointCenId) {
		query["EndpointCenId"] = request.EndpointCenId
	}

	if !dara.IsNil(request.EndpointInstanceId) {
		query["EndpointInstanceId"] = request.EndpointInstanceId
	}

	if !dara.IsNil(request.EndpointInstanceType) {
		query["EndpointInstanceType"] = request.EndpointInstanceType
	}

	if !dara.IsNil(request.EndpointRegion) {
		query["EndpointRegion"] = request.EndpointRegion
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitDtsRdsInstance"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitDtsRdsInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initializes a built-in account on a node of an active geo-redundancy database cluster. Data Transmission Service (DTS) uses the built-in account to connect to the node and perform data synchronization tasks.
//
// Description:
//
//	  The node must be an ApsaraDB RDS for MySQL instance or a self-managed MySQL database that is connected over Cloud Enterprise Network (CEN).
//
//		- This operation is used to initialize the built-in account named rdsdt_dtsacct on a node of an active geo-redundancy database cluster. DTS uses this account to connect to the node and perform data synchronization tasks.
//
// @param request - InitDtsRdsInstanceRequest
//
// @return InitDtsRdsInstanceResponse
func (client *Client) InitDtsRdsInstance(request *InitDtsRdsInstanceRequest) (_result *InitDtsRdsInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InitDtsRdsInstanceResponse{}
	_body, _err := client.InitDtsRdsInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all clusters that are created within an Alibaba Cloud account. You can also query clusters based on the specified conditions.
//
// @param request - ListDedicatedClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDedicatedClusterResponse
func (client *Client) ListDedicatedClusterWithOptions(request *ListDedicatedClusterRequest, runtime *dara.RuntimeOptions) (_result *ListDedicatedClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderColumn) {
		query["OrderColumn"] = request.OrderColumn
	}

	if !dara.IsNil(request.OrderDirection) {
		query["OrderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDedicatedCluster"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDedicatedClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all clusters that are created within an Alibaba Cloud account. You can also query clusters based on the specified conditions.
//
// @param request - ListDedicatedClusterRequest
//
// @return ListDedicatedClusterResponse
func (client *Client) ListDedicatedCluster(request *ListDedicatedClusterRequest) (_result *ListDedicatedClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDedicatedClusterResponse{}
	_body, _err := client.ListDedicatedClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询JobStep列表。
//
// @param request - ListJobStepRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobStepResponse
func (client *Client) ListJobStepWithOptions(request *ListJobStepRequest, runtime *dara.RuntimeOptions) (_result *ListJobStepResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ZeroEtlJob) {
		query["ZeroEtlJob"] = request.ZeroEtlJob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobStep"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobStepResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询JobStep列表。
//
// @param request - ListJobStepRequest
//
// @return ListJobStepResponse
func (client *Client) ListJobStep(request *ListJobStepRequest) (_result *ListJobStepResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListJobStepResponse{}
	_body, _err := client.ListJobStepWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are bound to specific data migration, data synchronization, or change tracking instances, or queries the instances to which specific tags are bound.
//
// Description:
//
// ***
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
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
		Version:     dara.String("2020-01-01"),
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
// Queries the tags that are bound to specific data migration, data synchronization, or change tracking instances, or queries the instances to which specific tags are bound.
//
// Description:
//
// ***
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
// Modifies the information of a consumer group, including the consumer group name, username, and password.
//
// @param request - ModifyConsumerChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyConsumerChannelResponse
func (client *Client) ModifyConsumerChannelWithOptions(request *ModifyConsumerChannelRequest, runtime *dara.RuntimeOptions) (_result *ModifyConsumerChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerGroupId) {
		query["ConsumerGroupId"] = request.ConsumerGroupId
	}

	if !dara.IsNil(request.ConsumerGroupName) {
		query["ConsumerGroupName"] = request.ConsumerGroupName
	}

	if !dara.IsNil(request.ConsumerGroupPassword) {
		query["ConsumerGroupPassword"] = request.ConsumerGroupPassword
	}

	if !dara.IsNil(request.ConsumerGroupUserName) {
		query["ConsumerGroupUserName"] = request.ConsumerGroupUserName
	}

	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyConsumerChannel"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyConsumerChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information of a consumer group, including the consumer group name, username, and password.
//
// @param request - ModifyConsumerChannelRequest
//
// @return ModifyConsumerChannelResponse
func (client *Client) ModifyConsumerChannel(request *ModifyConsumerChannelRequest) (_result *ModifyConsumerChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyConsumerChannelResponse{}
	_body, _err := client.ModifyConsumerChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modifies the password of a consumer group
//
// @param request - ModifyConsumerGroupPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyConsumerGroupPasswordResponse
func (client *Client) ModifyConsumerGroupPasswordWithOptions(request *ModifyConsumerGroupPasswordRequest, runtime *dara.RuntimeOptions) (_result *ModifyConsumerGroupPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ConsumerGroupID) {
		query["ConsumerGroupID"] = request.ConsumerGroupID
	}

	if !dara.IsNil(request.ConsumerGroupName) {
		query["ConsumerGroupName"] = request.ConsumerGroupName
	}

	if !dara.IsNil(request.ConsumerGroupPassword) {
		query["ConsumerGroupPassword"] = request.ConsumerGroupPassword
	}

	if !dara.IsNil(request.ConsumerGroupUserName) {
		query["ConsumerGroupUserName"] = request.ConsumerGroupUserName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SubscriptionInstanceId) {
		query["SubscriptionInstanceId"] = request.SubscriptionInstanceId
	}

	if !dara.IsNil(request.ConsumerGroupNewPassword) {
		query["consumerGroupNewPassword"] = request.ConsumerGroupNewPassword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyConsumerGroupPassword"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyConsumerGroupPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modifies the password of a consumer group
//
// @param request - ModifyConsumerGroupPasswordRequest
//
// @return ModifyConsumerGroupPasswordResponse
func (client *Client) ModifyConsumerGroupPassword(request *ModifyConsumerGroupPasswordRequest) (_result *ModifyConsumerGroupPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyConsumerGroupPasswordResponse{}
	_body, _err := client.ModifyConsumerGroupPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the consumption checkpoint of a change tracking instance.
//
// @param request - ModifyConsumptionTimestampRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyConsumptionTimestampResponse
func (client *Client) ModifyConsumptionTimestampWithOptions(request *ModifyConsumptionTimestampRequest, runtime *dara.RuntimeOptions) (_result *ModifyConsumptionTimestampResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ConsumptionTimestamp) {
		query["ConsumptionTimestamp"] = request.ConsumptionTimestamp
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SubscriptionInstanceId) {
		query["SubscriptionInstanceId"] = request.SubscriptionInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyConsumptionTimestamp"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyConsumptionTimestampResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the consumption checkpoint of a change tracking instance.
//
// @param request - ModifyConsumptionTimestampRequest
//
// @return ModifyConsumptionTimestampResponse
func (client *Client) ModifyConsumptionTimestamp(request *ModifyConsumptionTimestampRequest) (_result *ModifyConsumptionTimestampResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyConsumptionTimestampResponse{}
	_body, _err := client.ModifyConsumptionTimestampWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a cluster.
//
// Description:
//
// You can modify only the overcommit ratio.
//
// @param request - ModifyDedicatedClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDedicatedClusterResponse
func (client *Client) ModifyDedicatedClusterWithOptions(request *ModifyDedicatedClusterRequest, runtime *dara.RuntimeOptions) (_result *ModifyDedicatedClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DedicatedClusterId) {
		query["DedicatedClusterId"] = request.DedicatedClusterId
	}

	if !dara.IsNil(request.DedicatedClusterName) {
		query["DedicatedClusterName"] = request.DedicatedClusterName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OversoldRatio) {
		query["OversoldRatio"] = request.OversoldRatio
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDedicatedCluster"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDedicatedClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a cluster.
//
// Description:
//
// You can modify only the overcommit ratio.
//
// @param request - ModifyDedicatedClusterRequest
//
// @return ModifyDedicatedClusterResponse
func (client *Client) ModifyDedicatedCluster(request *ModifyDedicatedClusterRequest) (_result *ModifyDedicatedClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDedicatedClusterResponse{}
	_body, _err := client.ModifyDedicatedClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a data synchronization task.
//
// Description:
//
// When you configure a data synchronization task in the Data Transmission Service (DTS) console, you can move the pointer over **Next: Save Task Settings and Precheck*	- in the **Advanced Settings*	- step and click **Preview OpenAPI parameters*	- to view the parameters that are used to configure the task by calling an API operation.
//
// @param tmpReq - ModifyDtsJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDtsJobResponse
func (client *Client) ModifyDtsJobWithOptions(tmpReq *ModifyDtsJobRequest, runtime *dara.RuntimeOptions) (_result *ModifyDtsJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyDtsJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DbList) {
		request.DbListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DbList, dara.String("DbList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DataInitialization) {
		query["DataInitialization"] = request.DataInitialization
	}

	if !dara.IsNil(request.DataSynchronization) {
		query["DataSynchronization"] = request.DataSynchronization
	}

	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.FileOssUrl) {
		query["FileOssUrl"] = request.FileOssUrl
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StructureInitialization) {
		query["StructureInitialization"] = request.StructureInitialization
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !dara.IsNil(request.ZeroEtlJob) {
		query["ZeroEtlJob"] = request.ZeroEtlJob
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DbListShrink) {
		body["DbList"] = request.DbListShrink
	}

	if !dara.IsNil(request.EtlOperatorColumnReference) {
		body["EtlOperatorColumnReference"] = request.EtlOperatorColumnReference
	}

	if !dara.IsNil(request.FilterTableName) {
		body["FilterTableName"] = request.FilterTableName
	}

	if !dara.IsNil(request.ModifyTypeEnum) {
		body["ModifyTypeEnum"] = request.ModifyTypeEnum
	}

	if !dara.IsNil(request.Reserved) {
		body["Reserved"] = request.Reserved
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDtsJob"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDtsJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a data synchronization task.
//
// Description:
//
// When you configure a data synchronization task in the Data Transmission Service (DTS) console, you can move the pointer over **Next: Save Task Settings and Precheck*	- in the **Advanced Settings*	- step and click **Preview OpenAPI parameters*	- to view the parameters that are used to configure the task by calling an API operation.
//
// @param request - ModifyDtsJobRequest
//
// @return ModifyDtsJobResponse
func (client *Client) ModifyDtsJob(request *ModifyDtsJobRequest) (_result *ModifyDtsJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDtsJobResponse{}
	_body, _err := client.ModifyDtsJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDtsJobAdvance(request *ModifyDtsJobAdvanceRequest, runtime *dara.RuntimeOptions) (_result *ModifyDtsJobResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("Dts"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	modifyDtsJobReq := &ModifyDtsJobRequest{}
	openapiutil.Convert(request, modifyDtsJobReq)
	if !dara.IsNil(request.FileOssUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileOssUrlObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		modifyDtsJobReq.FileOssUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	modifyDtsJobResp, _err := client.ModifyDtsJobWithOptions(modifyDtsJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = modifyDtsJobResp
	return _result, _err
}

// Summary:
//
// Modifies the parameters of a Data Transmission Service (DTS) task.
//
// @param request - ModifyDtsJobConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDtsJobConfigResponse
func (client *Client) ModifyDtsJobConfigWithOptions(request *ModifyDtsJobConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyDtsJobConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDtsJobConfig"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDtsJobConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the parameters of a Data Transmission Service (DTS) task.
//
// @param request - ModifyDtsJobConfigRequest
//
// @return ModifyDtsJobConfigResponse
func (client *Client) ModifyDtsJobConfig(request *ModifyDtsJobConfigRequest) (_result *ModifyDtsJobConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDtsJobConfigResponse{}
	_body, _err := client.ModifyDtsJobConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the dedicated cluster on which a Data Transmission Service (DTS) task runs.
//
// Description:
//
// > After a DTS task is migrated from a dedicated cluster to a shared cluster, the task is billed on a pay-as-you-go basis.
//
// @param request - ModifyDtsJobDedicatedClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDtsJobDedicatedClusterResponse
func (client *Client) ModifyDtsJobDedicatedClusterWithOptions(request *ModifyDtsJobDedicatedClusterRequest, runtime *dara.RuntimeOptions) (_result *ModifyDtsJobDedicatedClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DedicatedClusterId) {
		query["DedicatedClusterId"] = request.DedicatedClusterId
	}

	if !dara.IsNil(request.DtsJobIds) {
		query["DtsJobIds"] = request.DtsJobIds
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDtsJobDedicatedCluster"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDtsJobDedicatedClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the dedicated cluster on which a Data Transmission Service (DTS) task runs.
//
// Description:
//
// > After a DTS task is migrated from a dedicated cluster to a shared cluster, the task is billed on a pay-as-you-go basis.
//
// @param request - ModifyDtsJobDedicatedClusterRequest
//
// @return ModifyDtsJobDedicatedClusterResponse
func (client *Client) ModifyDtsJobDedicatedCluster(request *ModifyDtsJobDedicatedClusterRequest) (_result *ModifyDtsJobDedicatedClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDtsJobDedicatedClusterResponse{}
	_body, _err := client.ModifyDtsJobDedicatedClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the upper limit of DTS units (DUs) for a Data Transmission Service (DTS) task.
//
// Description:
//
//	  DTS allows you to upgrade or downgrade the configurations of DTS instances in a dedicated cluster. You can adjust the resources that are occupied for task execution to dynamically adjust the number of tasks that can be scheduled in the cluster. This way, you can reduce the total number of DUs required for the cluster or release DUs.
//
//		- Before you modify the upper limit of DUs for a DTS task, make sure that sufficient DUs are available.
//
// @param request - ModifyDtsJobDuLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDtsJobDuLimitResponse
func (client *Client) ModifyDtsJobDuLimitWithOptions(request *ModifyDtsJobDuLimitRequest, runtime *dara.RuntimeOptions) (_result *ModifyDtsJobDuLimitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.DuLimit) {
		query["DuLimit"] = request.DuLimit
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDtsJobDuLimit"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDtsJobDuLimitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the upper limit of DTS units (DUs) for a Data Transmission Service (DTS) task.
//
// Description:
//
//	  DTS allows you to upgrade or downgrade the configurations of DTS instances in a dedicated cluster. You can adjust the resources that are occupied for task execution to dynamically adjust the number of tasks that can be scheduled in the cluster. This way, you can reduce the total number of DUs required for the cluster or release DUs.
//
//		- Before you modify the upper limit of DUs for a DTS task, make sure that sufficient DUs are available.
//
// @param request - ModifyDtsJobDuLimitRequest
//
// @return ModifyDtsJobDuLimitResponse
func (client *Client) ModifyDtsJobDuLimit(request *ModifyDtsJobDuLimitRequest) (_result *ModifyDtsJobDuLimitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDtsJobDuLimitResponse{}
	_body, _err := client.ModifyDtsJobDuLimitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the source or destination database instance of a data synchronization or migration task in Data Transmission Service (DTS).
//
// Description:
//
// >  After the database is changed, Data Transmission Service (DTS) rolls back the incremental write offset for 10 seconds. If the synchronized or migrated data does not have a primary key, make sure that no data is written to the source database while the source or destination database is being replaced. Otherwise, duplicate data may exist.
//
// @param request - ModifyDtsJobEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDtsJobEndpointResponse
func (client *Client) ModifyDtsJobEndpointWithOptions(request *ModifyDtsJobEndpointRequest, runtime *dara.RuntimeOptions) (_result *ModifyDtsJobEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunUid) {
		query["AliyunUid"] = request.AliyunUid
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.Endpoint) {
		query["Endpoint"] = request.Endpoint
	}

	if !dara.IsNil(request.EndpointInstanceId) {
		query["EndpointInstanceId"] = request.EndpointInstanceId
	}

	if !dara.IsNil(request.EndpointInstanceType) {
		query["EndpointInstanceType"] = request.EndpointInstanceType
	}

	if !dara.IsNil(request.EndpointIp) {
		query["EndpointIp"] = request.EndpointIp
	}

	if !dara.IsNil(request.EndpointPort) {
		query["EndpointPort"] = request.EndpointPort
	}

	if !dara.IsNil(request.EndpointRegionId) {
		query["EndpointRegionId"] = request.EndpointRegionId
	}

	if !dara.IsNil(request.ModifyAccount) {
		query["ModifyAccount"] = request.ModifyAccount
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RoleName) {
		query["RoleName"] = request.RoleName
	}

	if !dara.IsNil(request.ShardPassword) {
		query["ShardPassword"] = request.ShardPassword
	}

	if !dara.IsNil(request.ShardUsername) {
		query["ShardUsername"] = request.ShardUsername
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	if !dara.IsNil(request.ZeroEtlJob) {
		query["ZeroEtlJob"] = request.ZeroEtlJob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDtsJobEndpoint"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDtsJobEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the source or destination database instance of a data synchronization or migration task in Data Transmission Service (DTS).
//
// Description:
//
// >  After the database is changed, Data Transmission Service (DTS) rolls back the incremental write offset for 10 seconds. If the synchronized or migrated data does not have a primary key, make sure that no data is written to the source database while the source or destination database is being replaced. Otherwise, duplicate data may exist.
//
// @param request - ModifyDtsJobEndpointRequest
//
// @return ModifyDtsJobEndpointResponse
func (client *Client) ModifyDtsJobEndpoint(request *ModifyDtsJobEndpointRequest) (_result *ModifyDtsJobEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDtsJobEndpointResponse{}
	_body, _err := client.ModifyDtsJobEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the name of a Data Transmission Service (DTS) task.
//
// @param request - ModifyDtsJobNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDtsJobNameResponse
func (client *Client) ModifyDtsJobNameWithOptions(request *ModifyDtsJobNameRequest, runtime *dara.RuntimeOptions) (_result *ModifyDtsJobNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.DtsJobName) {
		query["DtsJobName"] = request.DtsJobName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ZeroEtlJob) {
		query["ZeroEtlJob"] = request.ZeroEtlJob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDtsJobName"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDtsJobNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the name of a Data Transmission Service (DTS) task.
//
// @param request - ModifyDtsJobNameRequest
//
// @return ModifyDtsJobNameResponse
func (client *Client) ModifyDtsJobName(request *ModifyDtsJobNameRequest) (_result *ModifyDtsJobNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDtsJobNameResponse{}
	_body, _err := client.ModifyDtsJobNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the password of the account used to log on to the source or destination database in a Data Transmission Service (DTS) task.
//
// @param request - ModifyDtsJobPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDtsJobPasswordResponse
func (client *Client) ModifyDtsJobPasswordWithOptions(request *ModifyDtsJobPasswordRequest, runtime *dara.RuntimeOptions) (_result *ModifyDtsJobPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.Endpoint) {
		query["Endpoint"] = request.Endpoint
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.ZeroEtlJob) {
		query["ZeroEtlJob"] = request.ZeroEtlJob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDtsJobPassword"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDtsJobPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the password of the account used to log on to the source or destination database in a Data Transmission Service (DTS) task.
//
// @param request - ModifyDtsJobPasswordRequest
//
// @return ModifyDtsJobPasswordResponse
func (client *Client) ModifyDtsJobPassword(request *ModifyDtsJobPasswordRequest) (_result *ModifyDtsJobPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDtsJobPasswordResponse{}
	_body, _err := client.ModifyDtsJobPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables throttling for data synchronization and data migration.
//
// @param request - ModifyDynamicConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDynamicConfigResponse
func (client *Client) ModifyDynamicConfigWithOptions(request *ModifyDynamicConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyDynamicConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigList) {
		query["ConfigList"] = request.ConfigList
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.EnableLimit) {
		query["EnableLimit"] = request.EnableLimit
	}

	if !dara.IsNil(request.JobCode) {
		query["JobCode"] = request.JobCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDynamicConfig"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDynamicConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables throttling for data synchronization and data migration.
//
// @param request - ModifyDynamicConfigRequest
//
// @return ModifyDynamicConfigResponse
func (client *Client) ModifyDynamicConfig(request *ModifyDynamicConfigRequest) (_result *ModifyDynamicConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDynamicConfigResponse{}
	_body, _err := client.ModifyDynamicConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改GAD实例名称
//
// @param request - ModifyGadInstanceNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGadInstanceNameResponse
func (client *Client) ModifyGadInstanceNameWithOptions(request *ModifyGadInstanceNameRequest, runtime *dara.RuntimeOptions) (_result *ModifyGadInstanceNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyGadInstanceName"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyGadInstanceNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改GAD实例名称
//
// @param request - ModifyGadInstanceNameRequest
//
// @return ModifyGadInstanceNameResponse
func (client *Client) ModifyGadInstanceName(request *ModifyGadInstanceNameRequest) (_result *ModifyGadInstanceNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyGadInstanceNameResponse{}
	_body, _err := client.ModifyGadInstanceNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改增量写入的位点。
//
// @param request - ModifyJobStepCheckpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyJobStepCheckpointResponse
func (client *Client) ModifyJobStepCheckpointWithOptions(request *ModifyJobStepCheckpointRequest, runtime *dara.RuntimeOptions) (_result *ModifyJobStepCheckpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.JobStepId) {
		query["JobStepId"] = request.JobStepId
	}

	if !dara.IsNil(request.NewCheckPoint) {
		query["NewCheckPoint"] = request.NewCheckPoint
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyJobStepCheckpoint"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyJobStepCheckpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改增量写入的位点。
//
// @param request - ModifyJobStepCheckpointRequest
//
// @return ModifyJobStepCheckpointResponse
func (client *Client) ModifyJobStepCheckpoint(request *ModifyJobStepCheckpointRequest) (_result *ModifyJobStepCheckpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyJobStepCheckpointResponse{}
	_body, _err := client.ModifyJobStepCheckpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about a change tracking task.
//
// Description:
//
// >  You can preview related API operation parameters when you modify the information about a change tracking task in the Data Transmission Service (DTS) console. This helps you configure the request parameters of this API operation. For more information, see [Preview the request parameters of API operations](https://help.aliyun.com/document_detail/2851612.html).
//
// @param request - ModifySubscriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySubscriptionResponse
func (client *Client) ModifySubscriptionWithOptions(request *ModifySubscriptionRequest, runtime *dara.RuntimeOptions) (_result *ModifySubscriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbList) {
		query["DbList"] = request.DbList
	}

	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.ModifyType) {
		query["ModifyType"] = request.ModifyType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Reserved) {
		query["Reserved"] = request.Reserved
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SubscriptionDataTypeDDL) {
		query["SubscriptionDataTypeDDL"] = request.SubscriptionDataTypeDDL
	}

	if !dara.IsNil(request.SubscriptionDataTypeDML) {
		query["SubscriptionDataTypeDML"] = request.SubscriptionDataTypeDML
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySubscription"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a change tracking task.
//
// Description:
//
// >  You can preview related API operation parameters when you modify the information about a change tracking task in the Data Transmission Service (DTS) console. This helps you configure the request parameters of this API operation. For more information, see [Preview the request parameters of API operations](https://help.aliyun.com/document_detail/2851612.html).
//
// @param request - ModifySubscriptionRequest
//
// @return ModifySubscriptionResponse
func (client *Client) ModifySubscription(request *ModifySubscriptionRequest) (_result *ModifySubscriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySubscriptionResponse{}
	_body, _err := client.ModifySubscriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the objects for change tracking.
//
// @param request - ModifySubscriptionObjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySubscriptionObjectResponse
func (client *Client) ModifySubscriptionObjectWithOptions(request *ModifySubscriptionObjectRequest, runtime *dara.RuntimeOptions) (_result *ModifySubscriptionObjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SubscriptionInstanceId) {
		query["SubscriptionInstanceId"] = request.SubscriptionInstanceId
	}

	if !dara.IsNil(request.SubscriptionObject) {
		query["SubscriptionObject"] = request.SubscriptionObject
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySubscriptionObject"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySubscriptionObjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the objects for change tracking.
//
// @param request - ModifySubscriptionObjectRequest
//
// @return ModifySubscriptionObjectResponse
func (client *Client) ModifySubscriptionObject(request *ModifySubscriptionObjectRequest) (_result *ModifySubscriptionObjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySubscriptionObjectResponse{}
	_body, _err := client.ModifySubscriptionObjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the objects to be synchronized.
//
// @param request - ModifySynchronizationObjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySynchronizationObjectResponse
func (client *Client) ModifySynchronizationObjectWithOptions(request *ModifySynchronizationObjectRequest, runtime *dara.RuntimeOptions) (_result *ModifySynchronizationObjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !dara.IsNil(request.SynchronizationJobId) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SynchronizationObjects) {
		body["SynchronizationObjects"] = request.SynchronizationObjects
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySynchronizationObject"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySynchronizationObjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the objects to be synchronized.
//
// @param request - ModifySynchronizationObjectRequest
//
// @return ModifySynchronizationObjectResponse
func (client *Client) ModifySynchronizationObject(request *ModifySynchronizationObjectRequest) (_result *ModifySynchronizationObjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySynchronizationObjectResponse{}
	_body, _err := client.ModifySynchronizationObjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 预检查创建GAD订单
//
// @param request - PreCheckCreateGadOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreCheckCreateGadOrderResponse
func (client *Client) PreCheckCreateGadOrderWithOptions(request *PreCheckCreateGadOrderRequest, runtime *dara.RuntimeOptions) (_result *PreCheckCreateGadOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MasterDatabaseName) {
		query["MasterDatabaseName"] = request.MasterDatabaseName
	}

	if !dara.IsNil(request.MasterEngineArchType) {
		query["MasterEngineArchType"] = request.MasterEngineArchType
	}

	if !dara.IsNil(request.MasterShardAccountName) {
		query["MasterShardAccountName"] = request.MasterShardAccountName
	}

	if !dara.IsNil(request.MasterShardAccountPassword) {
		query["MasterShardAccountPassword"] = request.MasterShardAccountPassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SlaveDatabaseName) {
		query["SlaveDatabaseName"] = request.SlaveDatabaseName
	}

	if !dara.IsNil(request.SlaveDbInstanceId) {
		query["SlaveDbInstanceId"] = request.SlaveDbInstanceId
	}

	if !dara.IsNil(request.SlaveDbInstanceRegion) {
		query["SlaveDbInstanceRegion"] = request.SlaveDbInstanceRegion
	}

	if !dara.IsNil(request.SlaveEngineArchType) {
		query["SlaveEngineArchType"] = request.SlaveEngineArchType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreCheckCreateGadOrder"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PreCheckCreateGadOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预检查创建GAD订单
//
// @param request - PreCheckCreateGadOrderRequest
//
// @return PreCheckCreateGadOrderResponse
func (client *Client) PreCheckCreateGadOrder(request *PreCheckCreateGadOrderRequest) (_result *PreCheckCreateGadOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PreCheckCreateGadOrderResponse{}
	_body, _err := client.PreCheckCreateGadOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renews a Data Transmission Service (DTS) instance. This API operation is available only for subscription instances.
//
// @param request - RenewInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewInstanceResponse
func (client *Client) RenewInstanceWithOptions(request *RenewInstanceRequest, runtime *dara.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BuyCount) {
		query["BuyCount"] = request.BuyCount
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewInstance"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews a Data Transmission Service (DTS) instance. This API operation is available only for subscription instances.
//
// @param request - RenewInstanceRequest
//
// @return RenewInstanceResponse
func (client *Client) RenewInstance(request *RenewInstanceRequest) (_result *RenewInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewInstanceResponse{}
	_body, _err := client.RenewInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets a data synchronization or change tracking task.
//
// Description:
//
// >  If you clear the configurations of a data synchronization or change tracking task, DTS deletes the task. Then, DTS creates another task. The task is in the Not Configured state. You must call the [ConfigureDtsJob](https://help.aliyun.com/document_detail/208399.html) operation reconfigure the task.
//
// @param request - ResetDtsJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetDtsJobResponse
func (client *Client) ResetDtsJobWithOptions(request *ResetDtsJobRequest, runtime *dara.RuntimeOptions) (_result *ResetDtsJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetDtsJob"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetDtsJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets a data synchronization or change tracking task.
//
// Description:
//
// >  If you clear the configurations of a data synchronization or change tracking task, DTS deletes the task. Then, DTS creates another task. The task is in the Not Configured state. You must call the [ConfigureDtsJob](https://help.aliyun.com/document_detail/208399.html) operation reconfigure the task.
//
// @param request - ResetDtsJobRequest
//
// @return ResetDtsJobResponse
func (client *Client) ResetDtsJob(request *ResetDtsJobRequest) (_result *ResetDtsJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetDtsJobResponse{}
	_body, _err := client.ResetDtsJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Clears the configurations of a data synchronization task.
//
// Description:
//
// >  If you clear the configurations of a data synchronization task, the task will be released. To start the task again, you must call the **ConfigureSynchronizationJob*	- operation to reconfigure the task.
//
// @param request - ResetSynchronizationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetSynchronizationJobResponse
func (client *Client) ResetSynchronizationJobWithOptions(request *ResetSynchronizationJobRequest, runtime *dara.RuntimeOptions) (_result *ResetSynchronizationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !dara.IsNil(request.SynchronizationJobId) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetSynchronizationJob"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetSynchronizationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clears the configurations of a data synchronization task.
//
// Description:
//
// >  If you clear the configurations of a data synchronization task, the task will be released. To start the task again, you must call the **ConfigureSynchronizationJob*	- operation to reconfigure the task.
//
// @param request - ResetSynchronizationJobRequest
//
// @return ResetSynchronizationJobResponse
func (client *Client) ResetSynchronizationJob(request *ResetSynchronizationJobRequest) (_result *ResetSynchronizationJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetSynchronizationJobResponse{}
	_body, _err := client.ResetSynchronizationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调转双向任务的方向
//
// @param request - ReverseTwoWayDirectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReverseTwoWayDirectionResponse
func (client *Client) ReverseTwoWayDirectionWithOptions(request *ReverseTwoWayDirectionRequest, runtime *dara.RuntimeOptions) (_result *ReverseTwoWayDirectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.IgnoreErrorSubJob) {
		query["IgnoreErrorSubJob"] = request.IgnoreErrorSubJob
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReverseTwoWayDirection"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReverseTwoWayDirectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调转双向任务的方向
//
// @param request - ReverseTwoWayDirectionRequest
//
// @return ReverseTwoWayDirectionResponse
func (client *Client) ReverseTwoWayDirection(request *ReverseTwoWayDirectionRequest) (_result *ReverseTwoWayDirectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReverseTwoWayDirectionResponse{}
	_body, _err := client.ReverseTwoWayDirectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Ignores the precheck items that a data migration or synchronization task may fail to pass.
//
// @param request - ShieldPrecheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ShieldPrecheckResponse
func (client *Client) ShieldPrecheckWithOptions(request *ShieldPrecheckRequest, runtime *dara.RuntimeOptions) (_result *ShieldPrecheckResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.PrecheckItems) {
		query["PrecheckItems"] = request.PrecheckItems
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ShieldPrecheck"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ShieldPrecheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Ignores the precheck items that a data migration or synchronization task may fail to pass.
//
// @param request - ShieldPrecheckRequest
//
// @return ShieldPrecheckResponse
func (client *Client) ShieldPrecheck(request *ShieldPrecheckRequest) (_result *ShieldPrecheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ShieldPrecheckResponse{}
	_body, _err := client.ShieldPrecheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The tables that do not need to be synchronized in a full data synchronization are skipped.
//
// @param request - SkipFullJobTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SkipFullJobTableResponse
func (client *Client) SkipFullJobTableWithOptions(request *SkipFullJobTableRequest, runtime *dara.RuntimeOptions) (_result *SkipFullJobTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.JobProgressId) {
		query["JobProgressId"] = request.JobProgressId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ZeroEtlJob) {
		query["ZeroEtlJob"] = request.ZeroEtlJob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SkipFullJobTable"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SkipFullJobTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The tables that do not need to be synchronized in a full data synchronization are skipped.
//
// @param request - SkipFullJobTableRequest
//
// @return SkipFullJobTableResponse
func (client *Client) SkipFullJobTable(request *SkipFullJobTableRequest) (_result *SkipFullJobTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SkipFullJobTableResponse{}
	_body, _err := client.SkipFullJobTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Skips one or more precheck items.
//
// @param request - SkipPreCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SkipPreCheckResponse
func (client *Client) SkipPreCheckWithOptions(request *SkipPreCheckRequest, runtime *dara.RuntimeOptions) (_result *SkipPreCheckResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Skip) {
		query["Skip"] = request.Skip
	}

	if !dara.IsNil(request.SkipPreCheckItems) {
		query["SkipPreCheckItems"] = request.SkipPreCheckItems
	}

	if !dara.IsNil(request.SkipPreCheckNames) {
		query["SkipPreCheckNames"] = request.SkipPreCheckNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SkipPreCheck"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SkipPreCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Skips one or more precheck items.
//
// @param request - SkipPreCheckRequest
//
// @return SkipPreCheckResponse
func (client *Client) SkipPreCheck(request *SkipPreCheckRequest) (_result *SkipPreCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SkipPreCheckResponse{}
	_body, _err := client.SkipPreCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a data migration, data synchronization, or change tracking task.
//
// @param request - StartDtsJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDtsJobResponse
func (client *Client) StartDtsJobWithOptions(request *StartDtsJobRequest, runtime *dara.RuntimeOptions) (_result *StartDtsJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !dara.IsNil(request.ZeroEtlJob) {
		query["ZeroEtlJob"] = request.ZeroEtlJob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartDtsJob"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDtsJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a data migration, data synchronization, or change tracking task.
//
// @param request - StartDtsJobRequest
//
// @return StartDtsJobResponse
func (client *Client) StartDtsJob(request *StartDtsJobRequest) (_result *StartDtsJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartDtsJobResponse{}
	_body, _err := client.StartDtsJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts multiple data migration or data synchronization tasks.
//
// @param request - StartDtsJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDtsJobsResponse
func (client *Client) StartDtsJobsWithOptions(request *StartDtsJobsRequest, runtime *dara.RuntimeOptions) (_result *StartDtsJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobIds) {
		query["DtsJobIds"] = request.DtsJobIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ZeroEtlJob) {
		query["ZeroEtlJob"] = request.ZeroEtlJob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartDtsJobs"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDtsJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts multiple data migration or data synchronization tasks.
//
// @param request - StartDtsJobsRequest
//
// @return StartDtsJobsResponse
func (client *Client) StartDtsJobs(request *StartDtsJobsRequest) (_result *StartDtsJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartDtsJobsResponse{}
	_body, _err := client.StartDtsJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a data migration task.
//
// @param request - StartMigrationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartMigrationJobResponse
func (client *Client) StartMigrationJobWithOptions(request *StartMigrationJobRequest, runtime *dara.RuntimeOptions) (_result *StartMigrationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.MigrationJobId) {
		query["MigrationJobId"] = request.MigrationJobId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartMigrationJob"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartMigrationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a data migration task.
//
// @param request - StartMigrationJobRequest
//
// @return StartMigrationJobResponse
func (client *Client) StartMigrationJob(request *StartMigrationJobRequest) (_result *StartMigrationJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartMigrationJobResponse{}
	_body, _err := client.StartMigrationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts the reverse task that is created by calling the CreateReverseDtsJob operation.
//
// Description:
//
// Before you call this operation, make sure that your instance is not released and is paused. You can check the status of the instance in the Data Transmission Service (DTS) console or by calling the [DescribeDtsJobDetail](https://help.aliyun.com/document_detail/208925.html) operation.
//
// @param request - StartReverseWriterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartReverseWriterResponse
func (client *Client) StartReverseWriterWithOptions(request *StartReverseWriterRequest, runtime *dara.RuntimeOptions) (_result *StartReverseWriterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckPoint) {
		query["CheckPoint"] = request.CheckPoint
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartReverseWriter"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartReverseWriterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts the reverse task that is created by calling the CreateReverseDtsJob operation.
//
// Description:
//
// Before you call this operation, make sure that your instance is not released and is paused. You can check the status of the instance in the Data Transmission Service (DTS) console or by calling the [DescribeDtsJobDetail](https://help.aliyun.com/document_detail/208925.html) operation.
//
// @param request - StartReverseWriterRequest
//
// @return StartReverseWriterResponse
func (client *Client) StartReverseWriter(request *StartReverseWriterRequest) (_result *StartReverseWriterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartReverseWriterResponse{}
	_body, _err := client.StartReverseWriterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a change tracking task.
//
// @param request - StartSubscriptionInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartSubscriptionInstanceResponse
func (client *Client) StartSubscriptionInstanceWithOptions(request *StartSubscriptionInstanceRequest, runtime *dara.RuntimeOptions) (_result *StartSubscriptionInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SubscriptionInstanceId) {
		query["SubscriptionInstanceId"] = request.SubscriptionInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartSubscriptionInstance"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartSubscriptionInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a change tracking task.
//
// @param request - StartSubscriptionInstanceRequest
//
// @return StartSubscriptionInstanceResponse
func (client *Client) StartSubscriptionInstance(request *StartSubscriptionInstanceRequest) (_result *StartSubscriptionInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartSubscriptionInstanceResponse{}
	_body, _err := client.StartSubscriptionInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a data synchronization task.
//
// @param request - StartSynchronizationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartSynchronizationJobResponse
func (client *Client) StartSynchronizationJobWithOptions(request *StartSynchronizationJobRequest, runtime *dara.RuntimeOptions) (_result *StartSynchronizationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !dara.IsNil(request.SynchronizationJobId) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartSynchronizationJob"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartSynchronizationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a data synchronization task.
//
// @param request - StartSynchronizationJobRequest
//
// @return StartSynchronizationJobResponse
func (client *Client) StartSynchronizationJob(request *StartSynchronizationJobRequest) (_result *StartSynchronizationJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartSynchronizationJobResponse{}
	_body, _err := client.StartSynchronizationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a cluster.
//
// @param request - StopDedicatedClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDedicatedClusterResponse
func (client *Client) StopDedicatedClusterWithOptions(request *StopDedicatedClusterRequest, runtime *dara.RuntimeOptions) (_result *StopDedicatedClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DedicatedClusterId) {
		query["DedicatedClusterId"] = request.DedicatedClusterId
	}

	if !dara.IsNil(request.DedicatedClusterName) {
		query["DedicatedClusterName"] = request.DedicatedClusterName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopDedicatedCluster"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDedicatedClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a cluster.
//
// @param request - StopDedicatedClusterRequest
//
// @return StopDedicatedClusterResponse
func (client *Client) StopDedicatedCluster(request *StopDedicatedClusterRequest) (_result *StopDedicatedClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopDedicatedClusterResponse{}
	_body, _err := client.StopDedicatedClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a data migration, data synchronization, or change tracking task.
//
// @param request - StopDtsJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDtsJobResponse
func (client *Client) StopDtsJobWithOptions(request *StopDtsJobRequest, runtime *dara.RuntimeOptions) (_result *StopDtsJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !dara.IsNil(request.ZeroEtlJob) {
		query["ZeroEtlJob"] = request.ZeroEtlJob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopDtsJob"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDtsJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a data migration, data synchronization, or change tracking task.
//
// @param request - StopDtsJobRequest
//
// @return StopDtsJobResponse
func (client *Client) StopDtsJob(request *StopDtsJobRequest) (_result *StopDtsJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopDtsJobResponse{}
	_body, _err := client.StopDtsJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops multiple data migration or data synchronization tasks.
//
// @param request - StopDtsJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDtsJobsResponse
func (client *Client) StopDtsJobsWithOptions(request *StopDtsJobsRequest, runtime *dara.RuntimeOptions) (_result *StopDtsJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobIds) {
		query["DtsJobIds"] = request.DtsJobIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ZeroEtlJob) {
		query["ZeroEtlJob"] = request.ZeroEtlJob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopDtsJobs"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDtsJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops multiple data migration or data synchronization tasks.
//
// @param request - StopDtsJobsRequest
//
// @return StopDtsJobsResponse
func (client *Client) StopDtsJobs(request *StopDtsJobsRequest) (_result *StopDtsJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopDtsJobsResponse{}
	_body, _err := client.StopDtsJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a data migration task that is in the Migrating state.
//
// @param request - StopMigrationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopMigrationJobResponse
func (client *Client) StopMigrationJobWithOptions(request *StopMigrationJobRequest, runtime *dara.RuntimeOptions) (_result *StopMigrationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.MigrationJobId) {
		query["MigrationJobId"] = request.MigrationJobId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopMigrationJob"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopMigrationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a data migration task that is in the Migrating state.
//
// @param request - StopMigrationJobRequest
//
// @return StopMigrationJobResponse
func (client *Client) StopMigrationJob(request *StopMigrationJobRequest) (_result *StopMigrationJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopMigrationJobResponse{}
	_body, _err := client.StopMigrationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of migrated or synchronized objects in a Data Transmission Service (DTS) task.
//
// @param request - SummaryJobDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SummaryJobDetailResponse
func (client *Client) SummaryJobDetailWithOptions(request *SummaryJobDetailRequest, runtime *dara.RuntimeOptions) (_result *SummaryJobDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.JobCode) {
		query["JobCode"] = request.JobCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StructType) {
		query["StructType"] = request.StructType
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !dara.IsNil(request.ZeroEtlJob) {
		query["ZeroEtlJob"] = request.ZeroEtlJob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SummaryJobDetail"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SummaryJobDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of migrated or synchronized objects in a Data Transmission Service (DTS) task.
//
// @param request - SummaryJobDetailRequest
//
// @return SummaryJobDetailResponse
func (client *Client) SummaryJobDetail(request *SummaryJobDetailRequest) (_result *SummaryJobDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SummaryJobDetailResponse{}
	_body, _err := client.SummaryJobDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Pauses a data migration, data synchronization, or change tracking task.
//
// Description:
//
// ***
//
// @param request - SuspendDtsJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendDtsJobResponse
func (client *Client) SuspendDtsJobWithOptions(request *SuspendDtsJobRequest, runtime *dara.RuntimeOptions) (_result *SuspendDtsJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !dara.IsNil(request.ZeroEtlJob) {
		query["ZeroEtlJob"] = request.ZeroEtlJob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendDtsJob"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendDtsJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pauses a data migration, data synchronization, or change tracking task.
//
// Description:
//
// ***
//
// @param request - SuspendDtsJobRequest
//
// @return SuspendDtsJobResponse
func (client *Client) SuspendDtsJob(request *SuspendDtsJobRequest) (_result *SuspendDtsJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SuspendDtsJobResponse{}
	_body, _err := client.SuspendDtsJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Suspends multiple Data Transmission Service (DTS) tasks.
//
// @param request - SuspendDtsJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendDtsJobsResponse
func (client *Client) SuspendDtsJobsWithOptions(request *SuspendDtsJobsRequest, runtime *dara.RuntimeOptions) (_result *SuspendDtsJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobIds) {
		query["DtsJobIds"] = request.DtsJobIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ZeroEtlJob) {
		query["ZeroEtlJob"] = request.ZeroEtlJob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendDtsJobs"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendDtsJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Suspends multiple Data Transmission Service (DTS) tasks.
//
// @param request - SuspendDtsJobsRequest
//
// @return SuspendDtsJobsResponse
func (client *Client) SuspendDtsJobs(request *SuspendDtsJobsRequest) (_result *SuspendDtsJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SuspendDtsJobsResponse{}
	_body, _err := client.SuspendDtsJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Pauses a data migration task.
//
// @param request - SuspendMigrationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendMigrationJobResponse
func (client *Client) SuspendMigrationJobWithOptions(request *SuspendMigrationJobRequest, runtime *dara.RuntimeOptions) (_result *SuspendMigrationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.MigrationJobId) {
		query["MigrationJobId"] = request.MigrationJobId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendMigrationJob"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendMigrationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pauses a data migration task.
//
// @param request - SuspendMigrationJobRequest
//
// @return SuspendMigrationJobResponse
func (client *Client) SuspendMigrationJob(request *SuspendMigrationJobRequest) (_result *SuspendMigrationJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SuspendMigrationJobResponse{}
	_body, _err := client.SuspendMigrationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Pauses a data synchronization task.
//
// Description:
//
// >
//
//   - When you call this operation, the data synchronization task must be in the Synchronizing state.
//
//   - We recommend that you do not pause a data synchronization task for more than 6 hours. Otherwise, the task cannot be started again.
//
//   - If the billing method is pay-as-you-go, DTS charges a fee even when the task is paused. This is because DTS only stops writing data to the destination database. DTS continues to pull the logs of the source database so that the task can resume quickly after it is restarted. Therefore, data synchronization consumes resources such as the bandwidth of the source database.
//
// @param request - SuspendSynchronizationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendSynchronizationJobResponse
func (client *Client) SuspendSynchronizationJobWithOptions(request *SuspendSynchronizationJobRequest, runtime *dara.RuntimeOptions) (_result *SuspendSynchronizationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !dara.IsNil(request.SynchronizationJobId) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendSynchronizationJob"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendSynchronizationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pauses a data synchronization task.
//
// Description:
//
// >
//
//   - When you call this operation, the data synchronization task must be in the Synchronizing state.
//
//   - We recommend that you do not pause a data synchronization task for more than 6 hours. Otherwise, the task cannot be started again.
//
//   - If the billing method is pay-as-you-go, DTS charges a fee even when the task is paused. This is because DTS only stops writing data to the destination database. DTS continues to pull the logs of the source database so that the task can resume quickly after it is restarted. Therefore, data synchronization consumes resources such as the bandwidth of the source database.
//
// @param request - SuspendSynchronizationJobRequest
//
// @return SuspendSynchronizationJobResponse
func (client *Client) SuspendSynchronizationJob(request *SuspendSynchronizationJobRequest) (_result *SuspendSynchronizationJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SuspendSynchronizationJobResponse{}
	_body, _err := client.SuspendSynchronizationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 物理迁移任务切换上云
//
// @param request - SwitchPhysicalDtsJobToCloudRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchPhysicalDtsJobToCloudResponse
func (client *Client) SwitchPhysicalDtsJobToCloudWithOptions(request *SwitchPhysicalDtsJobToCloudRequest, runtime *dara.RuntimeOptions) (_result *SwitchPhysicalDtsJobToCloudResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsInstanceId) {
		query["DtsInstanceId"] = request.DtsInstanceId
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchPhysicalDtsJobToCloud"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchPhysicalDtsJobToCloudResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 物理迁移任务切换上云
//
// @param request - SwitchPhysicalDtsJobToCloudRequest
//
// @return SwitchPhysicalDtsJobToCloudResponse
func (client *Client) SwitchPhysicalDtsJobToCloud(request *SwitchPhysicalDtsJobToCloudRequest) (_result *SwitchPhysicalDtsJobToCloudResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchPhysicalDtsJobToCloudResponse{}
	_body, _err := client.SwitchPhysicalDtsJobToCloudWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// After you perform a primary/secondary switchover on the source or destination database, you can call the SwitchSynchronizationEndpoint operation to transfer the connection settings to Data Transmission Service (DTS). DTS will restart the data synchronization task from the breakpoint.
//
// @param request - SwitchSynchronizationEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchSynchronizationEndpointResponse
func (client *Client) SwitchSynchronizationEndpointWithOptions(request *SwitchSynchronizationEndpointRequest, runtime *dara.RuntimeOptions) (_result *SwitchSynchronizationEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SynchronizationDirection) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !dara.IsNil(request.SynchronizationJobId) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	if !dara.IsNil(request.Endpoint) {
		query["Endpoint"] = request.Endpoint
	}

	if !dara.IsNil(request.SourceEndpoint) {
		query["SourceEndpoint"] = request.SourceEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchSynchronizationEndpoint"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchSynchronizationEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// After you perform a primary/secondary switchover on the source or destination database, you can call the SwitchSynchronizationEndpoint operation to transfer the connection settings to Data Transmission Service (DTS). DTS will restart the data synchronization task from the breakpoint.
//
// @param request - SwitchSynchronizationEndpointRequest
//
// @return SwitchSynchronizationEndpointResponse
func (client *Client) SwitchSynchronizationEndpoint(request *SwitchSynchronizationEndpointRequest) (_result *SwitchSynchronizationEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchSynchronizationEndpointResponse{}
	_body, _err := client.SwitchSynchronizationEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tags to data migration, data synchronization, or change tracking instances.
//
// Description:
//
// If you have a large number of instances, you can create multiple tags and add these tags to the instances. Then, you can query instances by tag.
//
//   - A tag consists of a key and a value. Each key must be unique in a region within an Alibaba Cloud account. Different keys can be mapped to the same value.
//
//   - If the tag that you specify does not exist, this tag is automatically created and added to the specified instance.
//
//   - If the key of the specified tag is the same as that of an existing tag, the specified tag overwrites the existing tag.
//
//   - You can add up to 20 tags to an instance.
//
//   - You can add tags to up to 50 instances in each request.
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
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
		Version:     dara.String("2020-01-01"),
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
// Adds tags to data migration, data synchronization, or change tracking instances.
//
// Description:
//
// If you have a large number of instances, you can create multiple tags and add these tags to the instances. Then, you can query instances by tag.
//
//   - A tag consists of a key and a value. Each key must be unique in a region within an Alibaba Cloud account. Different keys can be mapped to the same value.
//
//   - If the tag that you specify does not exist, this tag is automatically created and added to the specified instance.
//
//   - If the key of the specified tag is the same as that of an existing tag, the specified tag overwrites the existing tag.
//
//   - You can add up to 20 tags to an instance.
//
//   - You can add tags to up to 50 instances in each request.
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
// Upgrades or downgrades a Data Transmission Service (DTS) instance.
//
// @param request - TransferInstanceClassRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferInstanceClassResponse
func (client *Client) TransferInstanceClassWithOptions(request *TransferInstanceClassRequest, runtime *dara.RuntimeOptions) (_result *TransferInstanceClassResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferInstanceClass"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferInstanceClassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades or downgrades a Data Transmission Service (DTS) instance.
//
// @param request - TransferInstanceClassRequest
//
// @return TransferInstanceClassResponse
func (client *Client) TransferInstanceClass(request *TransferInstanceClassRequest) (_result *TransferInstanceClassResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TransferInstanceClassResponse{}
	_body, _err := client.TransferInstanceClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the billing method of a Data Transmission Service (DTS) instance.
//
// Description:
//
// Before you call this operation, make sure that you fully understand the [billing](https://www.alibabacloud.com/zh/product/data-transmission-service/pricing) of DTS.
//
//   - To prevent resource waste, make sure that the billing method of your DTS instances has to be changed.
//
//   - Data migration instances only support the pay-as-you-go billing method.
//
// @param request - TransferPayTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferPayTypeResponse
func (client *Client) TransferPayTypeWithOptions(request *TransferPayTypeRequest, runtime *dara.RuntimeOptions) (_result *TransferPayTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.BuyCount) {
		query["BuyCount"] = request.BuyCount
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.MaxDu) {
		query["MaxDu"] = request.MaxDu
	}

	if !dara.IsNil(request.MinDu) {
		query["MinDu"] = request.MinDu
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferPayType"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferPayTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the billing method of a Data Transmission Service (DTS) instance.
//
// Description:
//
// Before you call this operation, make sure that you fully understand the [billing](https://www.alibabacloud.com/zh/product/data-transmission-service/pricing) of DTS.
//
//   - To prevent resource waste, make sure that the billing method of your DTS instances has to be changed.
//
//   - Data migration instances only support the pay-as-you-go billing method.
//
// @param request - TransferPayTypeRequest
//
// @return TransferPayTypeResponse
func (client *Client) TransferPayType(request *TransferPayTypeRequest) (_result *TransferPayTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TransferPayTypeResponse{}
	_body, _err := client.TransferPayTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds tags from one or more data migration, data synchronization, or change tracking instances.
//
// Description:
//
// >  If a tag is unbound from an instance and is not bound to other instances, the tag is deleted.
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
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
		Version:     dara.String("2020-01-01"),
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
// Unbinds tags from one or more data migration, data synchronization, or change tracking instances.
//
// Description:
//
// >  If a tag is unbound from an instance and is not bound to other instances, the tag is deleted.
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
// Upgrades the synchronization topology of a data synchronization instance from one-way synchronization to two-way synchronization. This operation is supported only for pay-as-you-go synchronization instances.
//
// Description:
//
// Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/zh/product/data-transmission-service/pricing) of Data Transmission Service (DTS)
//
// When you call this operation, take note of the following information:
//
//   - The source and destination databases of the data synchronization task are both **MySQL*	- databases.
//
//   - The synchronization topology of the data synchronization task is **one-way synchronization**.
//
//   - The data synchronization task is in the **Synchronizing*	- state.
//
//   - The upgrade operation causes data synchronization latency of about 5 seconds. We recommend that you perform this operation during off-peak hours.
//
// @param request - UpgradeTwoWayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeTwoWayResponse
func (client *Client) UpgradeTwoWayWithOptions(request *UpgradeTwoWayRequest, runtime *dara.RuntimeOptions) (_result *UpgradeTwoWayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeTwoWay"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeTwoWayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades the synchronization topology of a data synchronization instance from one-way synchronization to two-way synchronization. This operation is supported only for pay-as-you-go synchronization instances.
//
// Description:
//
// Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/zh/product/data-transmission-service/pricing) of Data Transmission Service (DTS)
//
// When you call this operation, take note of the following information:
//
//   - The source and destination databases of the data synchronization task are both **MySQL*	- databases.
//
//   - The synchronization topology of the data synchronization task is **one-way synchronization**.
//
//   - The data synchronization task is in the **Synchronizing*	- state.
//
//   - The upgrade operation causes data synchronization latency of about 5 seconds. We recommend that you perform this operation during off-peak hours.
//
// @param request - UpgradeTwoWayRequest
//
// @return UpgradeTwoWayResponse
func (client *Client) UpgradeTwoWay(request *UpgradeTwoWayRequest) (_result *UpgradeTwoWayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeTwoWayResponse{}
	_body, _err := client.UpgradeTwoWayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// If the \\\\*\\\\*source or destination instance\\\\*\\\\	- is a \\\\*\\\\*self-managed database\\\\*\\\\	- or a \\\\*\\\\*third-party cloud database\\\\*\\\\*, you need to call this operation to query the CIDR blocks of DTS servers. Then, you need to add the CIDR blocks of DTS servers to the security settings of the source or destination instance, for example, the firewall of your database. For more information, see \\[Add the CIDR blocks of DTS servers to the security settings of on-premises databases]\\\\(~~176627~~).
//
// \\\\>  If the \\\\*\\\\*source or destination database\\\\*\\\\	- is an \\\\*\\\\*ApsaraDB database instance\\\\*\\\\	- (such as RDS instance and ApsaraDB for MongoDB instance) or a \\\\*\\\\*self-managed database hosted on Elastic Compute Service (ECS)\\\\*\\\\*, you do not need to add the CIDR blocks. When you click \\\\*\\\\*Set Whitelist and Next\\\\*\\\\	- in the DTS console, DTS automatically adds the CIDR blocks of DTS servers to the security settings of the source or destination instance.
//
// Description:
//
// The operation that you want to perform. Set the value to **WhiteIpList**.
//
// @param request - WhiteIpListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WhiteIpListResponse
func (client *Client) WhiteIpListWithOptions(request *WhiteIpListRequest, runtime *dara.RuntimeOptions) (_result *WhiteIpListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestAliyunUid) {
		query["DestAliyunUid"] = request.DestAliyunUid
	}

	if !dara.IsNil(request.DestPrimaryVswId) {
		query["DestPrimaryVswId"] = request.DestPrimaryVswId
	}

	if !dara.IsNil(request.DestRoleName) {
		query["DestRoleName"] = request.DestRoleName
	}

	if !dara.IsNil(request.DestSecondaryVswId) {
		query["DestSecondaryVswId"] = request.DestSecondaryVswId
	}

	if !dara.IsNil(request.DestVpcId) {
		query["DestVpcId"] = request.DestVpcId
	}

	if !dara.IsNil(request.DestinationRegion) {
		query["DestinationRegion"] = request.DestinationRegion
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SrcAliyunUid) {
		query["SrcAliyunUid"] = request.SrcAliyunUid
	}

	if !dara.IsNil(request.SrcPrimaryVswId) {
		query["SrcPrimaryVswId"] = request.SrcPrimaryVswId
	}

	if !dara.IsNil(request.SrcRoleName) {
		query["SrcRoleName"] = request.SrcRoleName
	}

	if !dara.IsNil(request.SrcSecondaryVswId) {
		query["SrcSecondaryVswId"] = request.SrcSecondaryVswId
	}

	if !dara.IsNil(request.SrcVpcId) {
		query["SrcVpcId"] = request.SrcVpcId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.ZeroEtlJob) {
		query["ZeroEtlJob"] = request.ZeroEtlJob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("WhiteIpList"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &WhiteIpListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If the \\\\*\\\\*source or destination instance\\\\*\\\\	- is a \\\\*\\\\*self-managed database\\\\*\\\\	- or a \\\\*\\\\*third-party cloud database\\\\*\\\\*, you need to call this operation to query the CIDR blocks of DTS servers. Then, you need to add the CIDR blocks of DTS servers to the security settings of the source or destination instance, for example, the firewall of your database. For more information, see \\[Add the CIDR blocks of DTS servers to the security settings of on-premises databases]\\\\(~~176627~~).
//
// \\\\>  If the \\\\*\\\\*source or destination database\\\\*\\\\	- is an \\\\*\\\\*ApsaraDB database instance\\\\*\\\\	- (such as RDS instance and ApsaraDB for MongoDB instance) or a \\\\*\\\\*self-managed database hosted on Elastic Compute Service (ECS)\\\\*\\\\*, you do not need to add the CIDR blocks. When you click \\\\*\\\\*Set Whitelist and Next\\\\*\\\\	- in the DTS console, DTS automatically adds the CIDR blocks of DTS servers to the security settings of the source or destination instance.
//
// Description:
//
// The operation that you want to perform. Set the value to **WhiteIpList**.
//
// @param request - WhiteIpListRequest
//
// @return WhiteIpListResponse
func (client *Client) WhiteIpList(request *WhiteIpListRequest) (_result *WhiteIpListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &WhiteIpListResponse{}
	_body, _err := client.WhiteIpListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func _postOSSObject_opResponse(response_ *dara.Response) (_result map[string]interface{}, _err error) {
	var respMap map[string]interface{}
	bodyStr, _err := dara.ReadAsString(response_.Body)
	if _err != nil {
		return _result, _err
	}

	if (dara.IntValue(response_.StatusCode) >= 400) && (dara.IntValue(response_.StatusCode) < 600) {
		respMap = dara.ParseXml(bodyStr, nil)
		err := dara.ToMap(respMap["Error"])
		_err = &openapi.ClientError{
			Code:    dara.String(dara.ToString(err["Code"])),
			Message: dara.String(dara.ToString(err["Message"])),
			Data: map[string]interface{}{
				"httpCode":  dara.IntValue(response_.StatusCode),
				"requestId": dara.ToString(err["RequestId"]),
				"hostId":    dara.ToString(err["HostId"]),
			},
		}
		return _result, _err
	}

	respMap = dara.ParseXml(bodyStr, nil)
	_result = make(map[string]interface{})
	_err = dara.Convert(dara.ToMap(respMap), &_result)

	return _result, _err
}
