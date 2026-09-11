// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Configures a data migration or synchronization task.
//
// Description:
//
// - You can perform the required pre-configurations in the console and then preview the corresponding OpenAPI parameter information to help you specify request parameters. For more information, see [Preview OpenAPI request parameters](https://help.aliyun.com/document_detail/2851612.html).
//
// - Tasks on dedicated clusters support only the configure-before-purchase mode and do not support cross-region tasks.
//
// @param request - ConfigureDtsJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureDtsJobResponse
func (client *Client) ConfigureDtsJobWithContext(ctx context.Context, request *ConfigureDtsJobRequest, runtime *dara.RuntimeOptions) (_result *ConfigureDtsJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a legacy data migration task.
//
// @param request - ConfigureMigrationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureMigrationJobResponse
func (client *Client) ConfigureMigrationJobWithContext(ctx context.Context, request *ConfigureMigrationJobRequest, runtime *dara.RuntimeOptions) (_result *ConfigureMigrationJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures monitoring alerts to monitor the latency and exception status of a data migration task.
//
// @param request - ConfigureMigrationJobAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureMigrationJobAlertResponse
func (client *Client) ConfigureMigrationJobAlertWithContext(ctx context.Context, request *ConfigureMigrationJobAlertRequest, runtime *dara.RuntimeOptions) (_result *ConfigureMigrationJobAlertResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a DTS change tracking task.
//
// Description:
//
// > You can perform the required pre-configurations in the console and then preview the corresponding OpenAPI parameter information to help you specify request parameters. For more information, see [Preview OpenAPI request parameters](https://help.aliyun.com/document_detail/2851612.html).
//
// @param request - ConfigureSubscriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureSubscriptionResponse
func (client *Client) ConfigureSubscriptionWithContext(ctx context.Context, request *ConfigureSubscriptionRequest, runtime *dara.RuntimeOptions) (_result *ConfigureSubscriptionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a change tracking channel. This is a legacy operation.
//
// Description:
//
// Before you call this operation, you must call the [CreateSubscriptionInstance](https://help.aliyun.com/document_detail/49436.html) operation to create a change tracking instance.
//
// > In the **Advanced Settings*	- step of the console, move the pointer over the **Next: Save the task and perform a precheck*	- button, and then click **Preview OpenAPI parameters*	- in the tooltip to view the parameter information for configuring this instance by using API operations.
//
// @param request - ConfigureSubscriptionInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureSubscriptionInstanceResponse
func (client *Client) ConfigureSubscriptionInstanceWithContext(ctx context.Context, request *ConfigureSubscriptionInstanceRequest, runtime *dara.RuntimeOptions) (_result *ConfigureSubscriptionInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures monitoring and alerting to monitor the latency and exception status of a change tracking channel.
//
// @param request - ConfigureSubscriptionInstanceAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureSubscriptionInstanceAlertResponse
func (client *Client) ConfigureSubscriptionInstanceAlertWithContext(ctx context.Context, request *ConfigureSubscriptionInstanceAlertRequest, runtime *dara.RuntimeOptions) (_result *ConfigureSubscriptionInstanceAlertResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a data synchronization task by using the previous version.
//
// Description:
//
// Before you call this operation, you must call the [CreateSynchronizationJob](https://help.aliyun.com/document_detail/49446.html) operation to create a data synchronization instance.
//
// > - After this operation is called, the data synchronization instance automatically starts and performs a precheck. You do not need to call the [StartSynchronizationJob](https://help.aliyun.com/document_detail/49448.html) operation to start the instance.
//
// - If the data synchronization instance fails to start, the precheck may have failed. You can call the [DescribeSynchronizationJobStatus](https://help.aliyun.com/document_detail/49453.html) operation to query the status of the data synchronization instance, obtain the error message of the precheck failure, and adjust the parameters. After the adjustment, you can call the [StartSynchronizationJob](https://help.aliyun.com/document_detail/49448.html) operation to restart the data synchronization instance.
//
// @param request - ConfigureSynchronizationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureSynchronizationJobResponse
func (client *Client) ConfigureSynchronizationJobWithContext(ctx context.Context, request *ConfigureSynchronizationJobRequest, runtime *dara.RuntimeOptions) (_result *ConfigureSynchronizationJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures monitoring and alerting to monitor the latency and exception status of a synchronization task.
//
// @param request - ConfigureSynchronizationJobAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureSynchronizationJobAlertResponse
func (client *Client) ConfigureSynchronizationJobAlertWithContext(ctx context.Context, request *ConfigureSynchronizationJobAlertRequest, runtime *dara.RuntimeOptions) (_result *ConfigureSynchronizationJobAlertResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the full image matching switch for a data synchronization instance.
//
// @param request - ConfigureSynchronizationJobReplicatorCompareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureSynchronizationJobReplicatorCompareResponse
func (client *Client) ConfigureSynchronizationJobReplicatorCompareWithContext(ctx context.Context, request *ConfigureSynchronizationJobReplicatorCompareRequest, runtime *dara.RuntimeOptions) (_result *ConfigureSynchronizationJobReplicatorCompareResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Transfers a DTS instance to a different resource group.
//
// @param request - ConvertInstanceResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConvertInstanceResourceGroupResponse
func (client *Client) ConvertInstanceResourceGroupWithContext(ctx context.Context, request *ConvertInstanceResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ConvertInstanceResourceGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the count of tasks by conditions.
//
// @param request - CountJobByConditionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CountJobByConditionResponse
func (client *Client) CountJobByConditionWithContext(ctx context.Context, request *CountJobByConditionRequest, runtime *dara.RuntimeOptions) (_result *CountJobByConditionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a consumer group for a change tracking task (new version).
//
// @param request - CreateConsumerChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConsumerChannelResponse
func (client *Client) CreateConsumerChannelWithContext(ctx context.Context, request *CreateConsumerChannelRequest, runtime *dara.RuntimeOptions) (_result *CreateConsumerChannelResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConsumerGroupResponse
func (client *Client) CreateConsumerGroupWithContext(ctx context.Context, request *CreateConsumerGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateConsumerGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an alert rule by calling the CreateDedicatedClusterMonitorRule operation.
//
// @param request - CreateDedicatedClusterMonitorRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDedicatedClusterMonitorRuleResponse
func (client *Client) CreateDedicatedClusterMonitorRuleWithContext(ctx context.Context, request *CreateDedicatedClusterMonitorRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateDedicatedClusterMonitorRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a document parsing task.
//
// Description:
//
// Calling this operation creates a document parsing task and returns a task ID (DtsJobId).
//
// > - This operation relies on Object Storage Service (OSS) for file transfer. We recommend that you call this operation by using an SDK. The CreateDocParserJobAdvance operation automatically encapsulates the file transfer process.
//
// > - After you obtain the DtsJobId response parameter, you can call the DescribeDocParserJobStatus operation to query the execution status of the document parsing task, and call the DescribeDocParserJobResult operation to obtain the output of the document parsing task.
//
// @param request - CreateDocParserJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDocParserJobResponse
func (client *Client) CreateDocParserJobWithContext(ctx context.Context, request *CreateDocParserJobRequest, runtime *dara.RuntimeOptions) (_result *CreateDocParserJobResponse, _err error) {
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

	if !dara.IsNil(request.RagInstanceId) {
		query["RagInstanceId"] = request.RagInstanceId
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Purchases a DTS instance by calling the CreateDtsInstance operation.
//
// Description:
//
// <props="china">
//
// - Before invoking this operation, make sure that you fully understand the billing methods and [pricing](https://www.aliyun.com/price/product#/dts/detail) of Data Transmission Service (DTS).
//
// <props="intl">
//
// - Before invoking this operation, make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/product/data-transmission-service/pricing) of Data Transmission Service (DTS).
//
// - Nodes on a dedicated cluster support only the workflow of configuring a node before purchasing an instance. You can invoke the [ConfigureDtsJob](https://help.aliyun.com/document_detail/208399.html) operation to configure a node.
//
// @param request - CreateDtsInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDtsInstanceResponse
func (client *Client) CreateDtsInstanceWithContext(ctx context.Context, request *CreateDtsInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateDtsInstanceResponse, _err error) {
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

	if !dara.IsNil(request.InsightModule) {
		query["InsightModule"] = request.InsightModule
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies an alert rule for a DTS task.
//
// Description:
//
// DTS currently supports the following alert metrics: **Latency**, **Migration Status**, and **Full Migration Duration**:
//
// - **Latency**: Monitors incremental data migration latency. An alert is triggered when the migration latency, synchronization latency, or change tracking latency exceeds the specified threshold (in seconds).
//
// - **Migration Status**: Monitors the task status. An alert is triggered when the task status is **Error*	- or **Recovered**.
//
// - **Full Migration Duration**: Monitors the duration of full data migration. An alert is triggered when the duration exceeds the specified threshold (in hours).
//
// @param request - CreateJobMonitorRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJobMonitorRuleResponse
func (client *Client) CreateJobMonitorRuleWithContext(ctx context.Context, request *CreateJobMonitorRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateJobMonitorRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMigrationJobResponse
func (client *Client) CreateMigrationJobWithContext(ctx context.Context, request *CreateMigrationJobRequest, runtime *dara.RuntimeOptions) (_result *CreateMigrationJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a reverse task for a specified synchronization or migration task by calling the CreateReverseDtsJob operation.
//
// Description:
//
// The reverse task created by calling this operation immediately starts a precheck. After the precheck is passed, incremental data collection begins, but the incremental data write module does not run. You must call the **StartReverseWriter*	- operation to start it.
//
// > The created reverse task is a synchronization task that contains only the incremental write module.
//
// @param request - CreateReverseDtsJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateReverseDtsJobResponse
func (client *Client) CreateReverseDtsJobWithContext(ctx context.Context, request *CreateReverseDtsJobRequest, runtime *dara.RuntimeOptions) (_result *CreateReverseDtsJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a change tracking channel. This is a legacy operation.
//
// @param request - CreateSubscriptionInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSubscriptionInstanceResponse
func (client *Client) CreateSubscriptionInstanceWithContext(ctx context.Context, request *CreateSubscriptionInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateSubscriptionInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data synchronization job instance. This is a legacy API operation.
//
// @param request - CreateSynchronizationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSynchronizationJobResponse
func (client *Client) CreateSynchronizationJobWithContext(ctx context.Context, request *CreateSynchronizationJobRequest, runtime *dara.RuntimeOptions) (_result *CreateSynchronizationJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a consumer group of a change tracking task (new version).
//
// @param request - DeleteConsumerChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConsumerChannelResponse
func (client *Client) DeleteConsumerChannelWithContext(ctx context.Context, request *DeleteConsumerChannelRequest, runtime *dara.RuntimeOptions) (_result *DeleteConsumerChannelResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a consumer group from a change tracking channel.
//
// @param request - DeleteConsumerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConsumerGroupResponse
func (client *Client) DeleteConsumerGroupWithContext(ctx context.Context, request *DeleteConsumerGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteConsumerGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a data migration, synchronization, or change tracking instance.
//
// Description:
//
// > <props="china"><ph>Subscription DTS instances cannot be released by calling this API operation. You can release them by unsubscribing. For more information, see [Release a DTS instance](https://help.aliyun.com/document_detail/289054.html).</ph><props="intl"><ph>Subscription DTS instances cannot be released.</ph>.
//
// @param request - DeleteDtsJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDtsJobResponse
func (client *Client) DeleteDtsJobWithContext(ctx context.Context, request *DeleteDtsJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteDtsJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases data migration, data synchronization, or change tracking tasks in batches by calling the DeleteDtsJobs operation.
//
// Description:
//
// > <props="china"><ph>Subscription DTS instances cannot be released by calling API operations. You can release them by unsubscribing. For more information, see [Release a DTS instance](https://help.aliyun.com/document_detail/289054.html).</ph><props="intl"><ph>Subscription DTS instances cannot be released.</ph>.
//
// @param request - DeleteDtsJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDtsJobsResponse
func (client *Client) DeleteDtsJobsWithContext(ctx context.Context, request *DeleteDtsJobsRequest, runtime *dara.RuntimeOptions) (_result *DeleteDtsJobsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMigrationJobResponse
func (client *Client) DeleteMigrationJobWithContext(ctx context.Context, request *DeleteMigrationJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteMigrationJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases the channel of a change tracking instance.
//
// @param request - DeleteSubscriptionInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSubscriptionInstanceResponse
func (client *Client) DeleteSubscriptionInstanceWithContext(ctx context.Context, request *DeleteSubscriptionInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteSubscriptionInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSynchronizationJobResponse
func (client *Client) DeleteSynchronizationJobWithContext(ctx context.Context, request *DeleteSynchronizationJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteSynchronizationJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelAccountResponse
func (client *Client) DescribeChannelAccountWithContext(ctx context.Context, request *DescribeChannelAccountRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelAccountResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Contains data validation tasks associated with data migration tasks and data synchronization tasks.
//
// @param request - DescribeCheckJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCheckJobsResponse
func (client *Client) DescribeCheckJobsWithContext(ctx context.Context, request *DescribeCheckJobsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCheckJobsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the log information of a DTS cluster by calling the DescribeClusterOperateLogs operation.
//
// @param request - DescribeClusterOperateLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterOperateLogsResponse
func (client *Client) DescribeClusterOperateLogsWithContext(ctx context.Context, request *DescribeClusterOperateLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterOperateLogsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the current usage of a cluster by calling the DescribeClusterUsedUtilization operation.
//
// @param request - DescribeClusterUsedUtilizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterUsedUtilizationResponse
func (client *Client) DescribeClusterUsedUtilizationWithContext(ctx context.Context, request *DescribeClusterUsedUtilizationRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterUsedUtilizationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Tests the connectivity between the execution node of a data migration task and the source and destination databases.
//
// @param request - DescribeConnectionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConnectionStatusResponse
func (client *Client) DescribeConnectionStatusWithContext(ctx context.Context, request *DescribeConnectionStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeConnectionStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the consumer group information of a DTS change tracking task, such as the consumer group ID, name, account, and consumption latency.
//
// @param request - DescribeConsumerChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConsumerChannelResponse
func (client *Client) DescribeConsumerChannelWithContext(ctx context.Context, request *DescribeConsumerChannelRequest, runtime *dara.RuntimeOptions) (_result *DescribeConsumerChannelResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConsumerGroupResponse
func (client *Client) DescribeConsumerGroupWithContext(ctx context.Context, request *DescribeConsumerGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeConsumerGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the DTS IP addresses that must be added to the whitelists of both the source and destination databases.
//
// @param request - DescribeDTSIPRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDTSIPResponse
func (client *Client) DescribeDTSIPWithContext(ctx context.Context, request *DescribeDTSIPRequest, runtime *dara.RuntimeOptions) (_result *DescribeDTSIPResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the download URL for the list data of inconsistent data.
//
// @param request - DescribeDataCheckReportUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataCheckReportUrlResponse
func (client *Client) DescribeDataCheckReportUrlWithContext(ctx context.Context, request *DescribeDataCheckReportUrlRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataCheckReportUrlResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries data consistency verification results at the table level.
//
// @param request - DescribeDataCheckTableDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataCheckTableDetailsResponse
func (client *Client) DescribeDataCheckTableDetailsWithContext(ctx context.Context, request *DescribeDataCheckTableDetailsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataCheckTableDetailsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists data inconsistency results grouped by inconsistent data.
//
// @param request - DescribeDataCheckTableDiffDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataCheckTableDiffDetailsResponse
func (client *Client) DescribeDataCheckTableDiffDetailsWithContext(ctx context.Context, request *DescribeDataCheckTableDiffDetailsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataCheckTableDiffDetailsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified cluster by calling the DescribeDedicatedCluster operation.
//
// @param request - DescribeDedicatedClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDedicatedClusterResponse
func (client *Client) DescribeDedicatedClusterWithContext(ctx context.Context, request *DescribeDedicatedClusterRequest, runtime *dara.RuntimeOptions) (_result *DescribeDedicatedClusterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries alert rules by calling the DescribeDedicatedClusterMonitorRule operation.
//
// @param request - DescribeDedicatedClusterMonitorRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDedicatedClusterMonitorRuleResponse
func (client *Client) DescribeDedicatedClusterMonitorRuleWithContext(ctx context.Context, request *DescribeDedicatedClusterMonitorRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribeDedicatedClusterMonitorRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the result of a document parsing task.
//
// Description:
//
// This operation has rate limits. Calls that exceed the limits are rejected.
//
// - The cumulative call threshold per region is 100 calls per second.
//
// - The call threshold per account per region is 5 calls per second.
//
// @param request - DescribeDocParserJobResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDocParserJobResultResponse
func (client *Client) DescribeDocParserJobResultWithContext(ctx context.Context, request *DescribeDocParserJobResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeDocParserJobResultResponse, _err error) {
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

	if !dara.IsNil(request.RagInstanceId) {
		query["RagInstanceId"] = request.RagInstanceId
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution status of a document parsing task.
//
// Description:
//
// This operation has call frequency limits. Calls that exceed the limits are rejected.
//
// - The cumulative call threshold for a single region is 200 calls per second.
//
// - The call threshold for a single account in a single region is 20 calls per second.
//
// @param request - DescribeDocParserJobStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDocParserJobStatusResponse
func (client *Client) DescribeDocParserJobStatusWithContext(ctx context.Context, request *DescribeDocParserJobStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeDocParserJobStatusResponse, _err error) {
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

	if !dara.IsNil(request.RagInstanceId) {
		query["RagInstanceId"] = request.RagInstanceId
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an ETL task.
//
// @param request - DescribeDtsEtlJobVersionInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDtsEtlJobVersionInfoResponse
func (client *Client) DescribeDtsEtlJobVersionInfoWithContext(ctx context.Context, request *DescribeDtsEtlJobVersionInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDtsEtlJobVersionInfoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDtsJobConfigResponse
func (client *Client) DescribeDtsJobConfigWithContext(ctx context.Context, request *DescribeDtsJobConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDtsJobConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a DTS task by calling DescribeDtsJobDetail.
//
// Description:
//
// This operation has rate limits. Calls that exceed the limits are rejected.
//
// - The cumulative threshold for calls in a single region is 160 calls per second.
//
// - The threshold for calls by a single account in a single region is 40 calls per second.
//
// @param request - DescribeDtsJobDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDtsJobDetailResponse
func (client *Client) DescribeDtsJobDetailWithContext(ctx context.Context, request *DescribeDtsJobDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeDtsJobDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbObjectOutputType) {
		query["DbObjectOutputType"] = request.DbObjectOutputType
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of DTS tasks and the execution details of each task.
//
// Description:
//
// This operation has rate limits. Calls that exceed the limits are rejected.
//
// - The cumulative threshold for calls in a single region is 200 calls per second.
//
// - The threshold for calls by a single account in a single region is 20 calls per second.
//
// @param request - DescribeDtsJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDtsJobsResponse
func (client *Client) DescribeDtsJobsWithContext(ctx context.Context, request *DescribeDtsJobsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDtsJobsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the log information of a data migration or synchronization task.
//
// @param request - DescribeDtsServiceLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDtsServiceLogResponse
func (client *Client) DescribeDtsServiceLogWithContext(ctx context.Context, request *DescribeDtsServiceLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeDtsServiceLogResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution status of a task.
//
// @param request - DescribeEndpointSwitchStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEndpointSwitchStatusResponse
func (client *Client) DescribeEndpointSwitchStatusWithContext(ctx context.Context, request *DescribeEndpointSwitchStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeEndpointSwitchStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the running logs of an ETL task.
//
// @param request - DescribeEtlJobLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEtlJobLogsResponse
func (client *Client) DescribeEtlJobLogsWithContext(ctx context.Context, request *DescribeEtlJobLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeEtlJobLogsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the running details of a full data migration task.
//
// @param request - DescribeFullProcessListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFullProcessListResponse
func (client *Client) DescribeFullProcessListWithContext(ctx context.Context, request *DescribeFullProcessListRequest, runtime *dara.RuntimeOptions) (_result *DescribeFullProcessListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Global Active Database (GAD) instances.
//
// @param request - DescribeGadInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGadInstancesResponse
func (client *Client) DescribeGadInstancesWithContext(ctx context.Context, request *DescribeGadInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeGadInstancesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the initialization status. This is an earlier version of the operation.
//
// @param request - DescribeInitializationStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInitializationStatusResponse
func (client *Client) DescribeInitializationStatusWithContext(ctx context.Context, request *DescribeInitializationStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeInitializationStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the alert rules of a DTS task by calling DescribeJobMonitorRule.
//
// @param request - DescribeJobMonitorRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeJobMonitorRuleResponse
func (client *Client) DescribeJobMonitorRuleWithContext(ctx context.Context, request *DescribeJobMonitorRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribeJobMonitorRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries cluster monitoring information by calling the DescribeMetricList operation.
//
// @param request - DescribeMetricListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricListResponse
func (client *Client) DescribeMetricListWithContext(ctx context.Context, request *DescribeMetricListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the monitoring and alert settings of a data migration task.
//
// @param request - DescribeMigrationJobAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMigrationJobAlertResponse
func (client *Client) DescribeMigrationJobAlertWithContext(ctx context.Context, request *DescribeMigrationJobAlertRequest, runtime *dara.RuntimeOptions) (_result *DescribeMigrationJobAlertResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution details of a data migration task. This is a legacy operation.
//
// @param request - DescribeMigrationJobDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMigrationJobDetailResponse
func (client *Client) DescribeMigrationJobDetailWithContext(ctx context.Context, request *DescribeMigrationJobDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeMigrationJobDetailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a data migration task. This is a legacy operation.
//
// @param request - DescribeMigrationJobStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMigrationJobStatusResponse
func (client *Client) DescribeMigrationJobStatusWithContext(ctx context.Context, request *DescribeMigrationJobStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeMigrationJobStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of data migration instances and details of each migration instance.
//
// @param request - DescribeMigrationJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMigrationJobsResponse
func (client *Client) DescribeMigrationJobsWithContext(ctx context.Context, request *DescribeMigrationJobsRequest, runtime *dara.RuntimeOptions) (_result *DescribeMigrationJobsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the task result of a precheck for creating a Global Active Database (GAD) order node.
//
// @param request - DescribePreCheckCreateGadOrderResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePreCheckCreateGadOrderResultResponse
func (client *Client) DescribePreCheckCreateGadOrderResultWithContext(ctx context.Context, request *DescribePreCheckCreateGadOrderResultRequest, runtime *dara.RuntimeOptions) (_result *DescribePreCheckCreateGadOrderResultResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution details of subtasks of a DTS task, including precheck, schema migration or synchronization, full data migration or synchronization, and incremental data migration or synchronization.
//
// @param request - DescribePreCheckStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePreCheckStatusResponse
func (client *Client) DescribePreCheckStatusWithContext(ctx context.Context, request *DescribePreCheckStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribePreCheckStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the monitoring and alerting settings of a change tracking instance.
//
// @param request - DescribeSubscriptionInstanceAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSubscriptionInstanceAlertResponse
func (client *Client) DescribeSubscriptionInstanceAlertWithContext(ctx context.Context, request *DescribeSubscriptionInstanceAlertRequest, runtime *dara.RuntimeOptions) (_result *DescribeSubscriptionInstanceAlertResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the instance status details of a change tracking channel. This is a legacy operation.
//
// @param request - DescribeSubscriptionInstanceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSubscriptionInstanceStatusResponse
func (client *Client) DescribeSubscriptionInstanceStatusWithContext(ctx context.Context, request *DescribeSubscriptionInstanceStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeSubscriptionInstanceStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSubscriptionInstancesResponse
func (client *Client) DescribeSubscriptionInstancesWithContext(ctx context.Context, request *DescribeSubscriptionInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSubscriptionInstancesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about subtasks of a PolarDB-X 1.0 distributed change tracking task.
//
// Description:
//
// <props="china">
//
// - Because a PolarDB-X 1.0 change tracking task is a distributed change tracking task, each ApsaraDB RDS for MySQL instance associated with the task corresponds to a change tracking subtask. You can call this operation to query the information about change tracking subtasks.
//
// - You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the instance ID, consumer group ID, and other information about a PolarDB-X 1.0 change tracking task.
//
// <props="intl">
//
// - Because a DRDS change tracking task is a distributed change tracking task, each ApsaraDB RDS for MySQL instance associated with the task corresponds to a change tracking subtask. You can call this operation to query the information about change tracking subtasks.
//
// - You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the instance ID, consumer group ID, and other information about a DRDS change tracking task.
//
// .
//
// @param tmpReq - DescribeSubscriptionMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSubscriptionMetaResponse
func (client *Client) DescribeSubscriptionMetaWithContext(ctx context.Context, tmpReq *DescribeSubscriptionMetaRequest, runtime *dara.RuntimeOptions) (_result *DescribeSubscriptionMetaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSyncStatusResponse
func (client *Client) DescribeSyncStatusWithContext(ctx context.Context, request *DescribeSyncStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeSyncStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the monitoring and alerting settings of a synchronization task.
//
// @param request - DescribeSynchronizationJobAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSynchronizationJobAlertResponse
func (client *Client) DescribeSynchronizationJobAlertWithContext(ctx context.Context, request *DescribeSynchronizationJobAlertRequest, runtime *dara.RuntimeOptions) (_result *DescribeSynchronizationJobAlertResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the current image matching switch configuration. This is a legacy operation.
//
// @param request - DescribeSynchronizationJobReplicatorCompareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSynchronizationJobReplicatorCompareResponse
func (client *Client) DescribeSynchronizationJobReplicatorCompareWithContext(ctx context.Context, request *DescribeSynchronizationJobReplicatorCompareRequest, runtime *dara.RuntimeOptions) (_result *DescribeSynchronizationJobReplicatorCompareResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the running status of a data synchronization task. This is a legacy API operation.
//
// @param request - DescribeSynchronizationJobStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSynchronizationJobStatusResponse
func (client *Client) DescribeSynchronizationJobStatusWithContext(ctx context.Context, request *DescribeSynchronizationJobStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeSynchronizationJobStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status list of synchronization jobs. This is a legacy operation.
//
// @param request - DescribeSynchronizationJobStatusListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSynchronizationJobStatusListResponse
func (client *Client) DescribeSynchronizationJobStatusListWithContext(ctx context.Context, request *DescribeSynchronizationJobStatusListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSynchronizationJobStatusListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of data synchronization instances and the details of each instance by calling DescribeSynchronizationJobs.
//
// @param request - DescribeSynchronizationJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSynchronizationJobsResponse
func (client *Client) DescribeSynchronizationJobsWithContext(ctx context.Context, request *DescribeSynchronizationJobsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSynchronizationJobsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution status of a task that modifies synchronization objects. This is a legacy operation.
//
// @param request - DescribeSynchronizationObjectModifyStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSynchronizationObjectModifyStatusResponse
func (client *Client) DescribeSynchronizationObjectModifyStatusWithContext(ctx context.Context, request *DescribeSynchronizationObjectModifyStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeSynchronizationObjectModifyStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all tags that are bound to a data migration, data synchronization, or change tracking instance.
//
// @param request - DescribeTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagKeysResponse
func (client *Client) DescribeTagKeysWithContext(ctx context.Context, request *DescribeTagKeysRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagKeysResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all values of a tag key that is attached to a data migration, data synchronization, or change tracking instance.
//
// @param request - DescribeTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagValuesResponse
func (client *Client) DescribeTagValuesWithContext(ctx context.Context, request *DescribeTagValuesRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagValuesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a secondary role.
//
// @param request - DetachGadInstanceDbMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachGadInstanceDbMemberResponse
func (client *Client) DetachGadInstanceDbMemberWithContext(ctx context.Context, request *DetachGadInstanceDbMemberRequest, runtime *dara.RuntimeOptions) (_result *DetachGadInstanceDbMemberResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initializes a built-in account in a node of an active geo-redundancy database cluster. Data Transmission Service (DTS) uses this account to connect to the node and perform synchronization tasks.
//
// Description:
//
// - The unit node must be an ApsaraDB RDS for MySQL instance or a self-managed MySQL database connected through Cloud Enterprise Network (CEN).
//
// - This operation initializes a built-in account named rdsdt_dtsacct in a unit node of an active geo-redundancy database cluster. DTS uses this account to connect to the node and perform synchronization tasks.
//
// @param request - InitDtsRdsInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitDtsRdsInstanceResponse
func (client *Client) InitDtsRdsInstanceWithContext(ctx context.Context, request *InitDtsRdsInstanceRequest, runtime *dara.RuntimeOptions) (_result *InitDtsRdsInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all clusters created by the current user. You can also filter specific clusters based on specified conditions.
//
// @param request - ListDedicatedClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDedicatedClusterResponse
func (client *Client) ListDedicatedClusterWithContext(ctx context.Context, request *ListDedicatedClusterRequest, runtime *dara.RuntimeOptions) (_result *ListDedicatedClusterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the JobStep list
//
// @param request - ListJobStepRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobStepResponse
func (client *Client) ListJobStepWithContext(ctx context.Context, request *ListJobStepRequest, runtime *dara.RuntimeOptions) (_result *ListJobStepResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags bound to data migration, data synchronization, and change tracking instances. You can also query the instances bound to specific tags.
//
// Description:
//
// ***.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information of a consumer group in a change tracking channel (new version).
//
// @param request - ModifyConsumerChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyConsumerChannelResponse
func (client *Client) ModifyConsumerChannelWithContext(ctx context.Context, request *ModifyConsumerChannelRequest, runtime *dara.RuntimeOptions) (_result *ModifyConsumerChannelResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the password of a consumer group. This is a legacy operation.
//
// @param request - ModifyConsumerGroupPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyConsumerGroupPasswordResponse
func (client *Client) ModifyConsumerGroupPasswordWithContext(ctx context.Context, request *ModifyConsumerGroupPasswordRequest, runtime *dara.RuntimeOptions) (_result *ModifyConsumerGroupPasswordResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the consumption checkpoint of a change tracking instance channel.
//
// @param request - ModifyConsumptionTimestampRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyConsumptionTimestampResponse
func (client *Client) ModifyConsumptionTimestampWithContext(ctx context.Context, request *ModifyConsumptionTimestampRequest, runtime *dara.RuntimeOptions) (_result *ModifyConsumptionTimestampResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a dedicated cluster by calling the ModifyDedicatedCluster operation.
//
// Description:
//
// Currently, only the overcommit ratio can be modified.
//
// @param request - ModifyDedicatedClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDedicatedClusterResponse
func (client *Client) ModifyDedicatedClusterWithContext(ctx context.Context, request *ModifyDedicatedClusterRequest, runtime *dara.RuntimeOptions) (_result *ModifyDedicatedClusterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a data synchronization task by calling the ModifyDtsJob operation.
//
// Description:
//
// > You can preconfigure settings in the console as needed, and then preview the corresponding OpenAPI parameter information to help you specify request parameters. For more information, see [Preview OpenAPI request parameters](https://help.aliyun.com/document_detail/2851612.html).
//
// @param tmpReq - ModifyDtsJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDtsJobResponse
func (client *Client) ModifyDtsJobWithContext(ctx context.Context, tmpReq *ModifyDtsJobRequest, runtime *dara.RuntimeOptions) (_result *ModifyDtsJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the parameters of a DTS task by calling the ModifyDtsJobConfig operation.
//
// @param request - ModifyDtsJobConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDtsJobConfigResponse
func (client *Client) ModifyDtsJobConfigWithContext(ctx context.Context, request *ModifyDtsJobConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyDtsJobConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the dedicated cluster on which a task runs.
//
// Description:
//
// > After a migration task is changed from a dedicated cluster to a public cluster, the billing method of the task changes to pay-as-you-go, and billing starts.
//
// @param request - ModifyDtsJobDedicatedClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDtsJobDedicatedClusterResponse
func (client *Client) ModifyDtsJobDedicatedClusterWithContext(ctx context.Context, request *ModifyDtsJobDedicatedClusterRequest, runtime *dara.RuntimeOptions) (_result *ModifyDtsJobDedicatedClusterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the DU upper limit of a DTS task by calling the ModifyDtsJobDuLimit operation.
//
// Description:
//
// - DTS instances in a dedicated cluster must support specification changes. By changing the resources consumed by a task at runtime, you can dynamically adjust the number of schedulable tasks in the current cluster, thereby deducting or releasing the total number of DUs in the cluster.
//
// - Before modifying the DU upper limit of a task, ensure that sufficient resources are available.
//
// @param request - ModifyDtsJobDuLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDtsJobDuLimitResponse
func (client *Client) ModifyDtsJobDuLimitWithContext(ctx context.Context, request *ModifyDtsJobDuLimitRequest, runtime *dara.RuntimeOptions) (_result *ModifyDtsJobDuLimitResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the source or destination instance of a DTS synchronization or migration task.
//
// Description:
//
// > After the database instance is modified, the DTS incremental write module rolls back writes by 10 seconds. If the data being synchronized or migrated does not have a primary key, stop writing data to the business associated with the source instance during the database instance replacement. Otherwise, duplicate data may occur.
//
// @param request - ModifyDtsJobEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDtsJobEndpointResponse
func (client *Client) ModifyDtsJobEndpointWithContext(ctx context.Context, request *ModifyDtsJobEndpointRequest, runtime *dara.RuntimeOptions) (_result *ModifyDtsJobEndpointResponse, _err error) {
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

	if !dara.IsNil(request.EndpointPrimaryVswId) {
		query["EndpointPrimaryVswId"] = request.EndpointPrimaryVswId
	}

	if !dara.IsNil(request.EndpointRegionId) {
		query["EndpointRegionId"] = request.EndpointRegionId
	}

	if !dara.IsNil(request.EndpointSecondaryVswId) {
		query["EndpointSecondaryVswId"] = request.EndpointSecondaryVswId
	}

	if !dara.IsNil(request.EndpointVpcId) {
		query["EndpointVpcId"] = request.EndpointVpcId
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name of a DTS task by calling ModifyDtsJobName.
//
// @param request - ModifyDtsJobNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDtsJobNameResponse
func (client *Client) ModifyDtsJobNameWithContext(ctx context.Context, request *ModifyDtsJobNameRequest, runtime *dara.RuntimeOptions) (_result *ModifyDtsJobNameResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the password of a DTS task (new version).
//
// @param request - ModifyDtsJobPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDtsJobPasswordResponse
func (client *Client) ModifyDtsJobPasswordWithContext(ctx context.Context, request *ModifyDtsJobPasswordRequest, runtime *dara.RuntimeOptions) (_result *ModifyDtsJobPasswordResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adjusts the migration rate of a data synchronization or migration instance.
//
// @param request - ModifyDynamicConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDynamicConfigResponse
func (client *Client) ModifyDynamicConfigWithContext(ctx context.Context, request *ModifyDynamicConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyDynamicConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name of a Global Active Database (GAD) instance.
//
// @param request - ModifyGadInstanceNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGadInstanceNameResponse
func (client *Client) ModifyGadInstanceNameWithContext(ctx context.Context, request *ModifyGadInstanceNameRequest, runtime *dara.RuntimeOptions) (_result *ModifyGadInstanceNameResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the offset for incremental data writing.
//
// @param request - ModifyJobStepCheckpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyJobStepCheckpointResponse
func (client *Client) ModifyJobStepCheckpointWithContext(ctx context.Context, request *ModifyJobStepCheckpointRequest, runtime *dara.RuntimeOptions) (_result *ModifyJobStepCheckpointResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a change tracking task (new version).
//
// Description:
//
// > You can perform the required preconfigurations in the console and then preview the corresponding OpenAPI parameter information to help you fill in the request parameters. For more information, see [Preview OpenAPI request parameters](https://help.aliyun.com/document_detail/2851612.html).
//
// @param request - ModifySubscriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySubscriptionResponse
func (client *Client) ModifySubscriptionWithContext(ctx context.Context, request *ModifySubscriptionRequest, runtime *dara.RuntimeOptions) (_result *ModifySubscriptionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the subscription objects of a change tracking task. This is a legacy operation.
//
// @param request - ModifySubscriptionObjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySubscriptionObjectResponse
func (client *Client) ModifySubscriptionObjectWithContext(ctx context.Context, request *ModifySubscriptionObjectRequest, runtime *dara.RuntimeOptions) (_result *ModifySubscriptionObjectResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the synchronization objects in a data synchronization job instance. This is a legacy operation.
//
// @param request - ModifySynchronizationObjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySynchronizationObjectResponse
func (client *Client) ModifySynchronizationObjectWithContext(ctx context.Context, request *ModifySynchronizationObjectRequest, runtime *dara.RuntimeOptions) (_result *ModifySynchronizationObjectResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Prechecks an order for creating a Global Active Database (GAD) instance group.
//
// @param request - PreCheckCreateGadOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreCheckCreateGadOrderResponse
func (client *Client) PreCheckCreateGadOrderWithContext(ctx context.Context, request *PreCheckCreateGadOrderRequest, runtime *dara.RuntimeOptions) (_result *PreCheckCreateGadOrderResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Promote a geo-disaster recovery instance from the secondary role to the primary role
//
// @param request - PromoteToMasterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PromoteToMasterResponse
func (client *Client) PromoteToMasterWithContext(ctx context.Context, request *PromoteToMasterRequest, runtime *dara.RuntimeOptions) (_result *PromoteToMasterResponse, _err error) {
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

	if !dara.IsNil(request.MasterDbInstanceId) {
		query["MasterDbInstanceId"] = request.MasterDbInstanceId
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
		Action:      dara.String("PromoteToMaster"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PromoteToMasterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews a DTS instance. This operation is applicable only to subscription DTS instances.
//
// @param request - RenewInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewInstanceResponse
func (client *Client) RenewInstanceWithContext(ctx context.Context, request *RenewInstanceRequest, runtime *dara.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// > After the configuration of a data synchronization or change tracking task is cleared, the original task is deleted. DTS creates a new unconfigured task. You must call the [ConfigureDtsJob](https://help.aliyun.com/document_detail/208399.html) operation to reconfigure the task.
//
// @param request - ResetDtsJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetDtsJobResponse
func (client *Client) ResetDtsJobWithContext(ctx context.Context, request *ResetDtsJobRequest, runtime *dara.RuntimeOptions) (_result *ResetDtsJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the configuration of a data synchronization task.
//
// Description:
//
// > After you reset the configuration of a data synchronization task, the original synchronization task is released. You must call the **ConfigureSynchronizationJob*	- operation to reconfigure the synchronization task before you can start the task.
//
// @param request - ResetSynchronizationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetSynchronizationJobResponse
func (client *Client) ResetSynchronizationJobWithContext(ctx context.Context, request *ResetSynchronizationJobRequest, runtime *dara.RuntimeOptions) (_result *ResetSynchronizationJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReverseTwoWayDirectionResponse
func (client *Client) ReverseTwoWayDirectionWithContext(ctx context.Context, request *ReverseTwoWayDirectionRequest, runtime *dara.RuntimeOptions) (_result *ReverseTwoWayDirectionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Skips the precheck for a legacy data migration or synchronization task.
//
// @param request - ShieldPrecheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ShieldPrecheckResponse
func (client *Client) ShieldPrecheckWithContext(ctx context.Context, request *ShieldPrecheckRequest, runtime *dara.RuntimeOptions) (_result *ShieldPrecheckResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Skips tables that do not need to be synchronized during the full data synchronization phase.
//
// @param request - SkipFullJobTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SkipFullJobTableResponse
func (client *Client) SkipFullJobTableWithContext(ctx context.Context, request *SkipFullJobTableRequest, runtime *dara.RuntimeOptions) (_result *SkipFullJobTableResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Suppresses or unsuppresses precheck alert items.
//
// @param request - SkipPreCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SkipPreCheckResponse
func (client *Client) SkipPreCheckWithContext(ctx context.Context, request *SkipPreCheckRequest, runtime *dara.RuntimeOptions) (_result *SkipPreCheckResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a data migration, data synchronization, or change tracking task by calling the StartDtsJob operation.
//
// @param request - StartDtsJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDtsJobResponse
func (client *Client) StartDtsJobWithContext(ctx context.Context, request *StartDtsJobRequest, runtime *dara.RuntimeOptions) (_result *StartDtsJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts multiple data migration or synchronization tasks in a batch by calling the StartDtsJobs operation.
//
// @param request - StartDtsJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDtsJobsResponse
func (client *Client) StartDtsJobsWithContext(ctx context.Context, request *StartDtsJobsRequest, runtime *dara.RuntimeOptions) (_result *StartDtsJobsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a data migration task of Data Transmission Service (DTS).
//
// @param request - StartMigrationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartMigrationJobResponse
func (client *Client) StartMigrationJobWithContext(ctx context.Context, request *StartMigrationJobRequest, runtime *dara.RuntimeOptions) (_result *StartMigrationJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a reverse task that is created by calling the CreateReverseDtsJob operation.
//
// Description:
//
// Before you call this operation, check the status of the reverse task in the console or by calling [DescribeDtsJobDetail](https://help.aliyun.com/document_detail/208925.html). Make sure that the task has not been released and is in the paused state.
//
// @param request - StartReverseWriterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartReverseWriterResponse
func (client *Client) StartReverseWriterWithContext(ctx context.Context, request *StartReverseWriterRequest, runtime *dara.RuntimeOptions) (_result *StartReverseWriterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts the channel of a change tracking instance. This is a legacy operation.
//
// @param request - StartSubscriptionInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartSubscriptionInstanceResponse
func (client *Client) StartSubscriptionInstanceWithContext(ctx context.Context, request *StartSubscriptionInstanceRequest, runtime *dara.RuntimeOptions) (_result *StartSubscriptionInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartSynchronizationJobResponse
func (client *Client) StartSynchronizationJobWithContext(ctx context.Context, request *StartSynchronizationJobRequest, runtime *dara.RuntimeOptions) (_result *StartSynchronizationJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a cluster by calling the StopDedicatedCluster operation.
//
// @param request - StopDedicatedClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDedicatedClusterResponse
func (client *Client) StopDedicatedClusterWithContext(ctx context.Context, request *StopDedicatedClusterRequest, runtime *dara.RuntimeOptions) (_result *StopDedicatedClusterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a data migration, data synchronization, or change tracking task by calling StopDtsJob.
//
// @param request - StopDtsJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDtsJobResponse
func (client *Client) StopDtsJobWithContext(ctx context.Context, request *StopDtsJobRequest, runtime *dara.RuntimeOptions) (_result *StopDtsJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops multiple DTS tasks at a time.
//
// @param request - StopDtsJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDtsJobsResponse
func (client *Client) StopDtsJobsWithContext(ctx context.Context, request *StopDtsJobsRequest, runtime *dara.RuntimeOptions) (_result *StopDtsJobsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Ends a data migration task that is in a migration state.
//
// @param request - StopMigrationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopMigrationJobResponse
func (client *Client) StopMigrationJobWithContext(ctx context.Context, request *StopMigrationJobRequest, runtime *dara.RuntimeOptions) (_result *StopMigrationJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of tables migrated in a Data Transmission Service (DTS) data migration or synchronization task.
//
// @param request - SummaryJobDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SummaryJobDetailResponse
func (client *Client) SummaryJobDetailWithContext(ctx context.Context, request *SummaryJobDetailRequest, runtime *dara.RuntimeOptions) (_result *SummaryJobDetailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Suspends a data migration or synchronization task. Change tracking tasks are not supported. Change tracking instances do not support the suspend capability. Do not call this operation on change tracking instances.
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
func (client *Client) SuspendDtsJobWithContext(ctx context.Context, request *SuspendDtsJobRequest, runtime *dara.RuntimeOptions) (_result *SuspendDtsJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Suspends multiple DTS tasks at a time.
//
// @param request - SuspendDtsJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendDtsJobsResponse
func (client *Client) SuspendDtsJobsWithContext(ctx context.Context, request *SuspendDtsJobsRequest, runtime *dara.RuntimeOptions) (_result *SuspendDtsJobsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Suspends a data migration task that is in progress.
//
// @param request - SuspendMigrationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendMigrationJobResponse
func (client *Client) SuspendMigrationJobWithContext(ctx context.Context, request *SuspendMigrationJobRequest, runtime *dara.RuntimeOptions) (_result *SuspendMigrationJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pauses a data synchronization task that is in the Synchronizing state.
//
// Description:
//
// > - When you call this operation, the synchronization task must be in the Synchronizing state.
//
// - A synchronization task cannot be paused for more than 6 hours. Otherwise, the task cannot be restarted.
//
// - DTS continues to charge fees for a pay-as-you-go synchronization task even if the task is paused. This is because DTS only pauses writing data to the destination instance but continues to pull logs from the source instance to ensure quick resumption when the task is restarted. Therefore, the task still consumes resources such as bandwidth of the source database.
//
// @param request - SuspendSynchronizationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendSynchronizationJobResponse
func (client *Client) SuspendSynchronizationJobWithContext(ctx context.Context, request *SuspendSynchronizationJobRequest, runtime *dara.RuntimeOptions) (_result *SuspendSynchronizationJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs physical migration of an MSSQL database to Alibaba Cloud.
//
// @param request - SwitchPhysicalDtsJobToCloudRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchPhysicalDtsJobToCloudResponse
func (client *Client) SwitchPhysicalDtsJobToCloudWithContext(ctx context.Context, request *SwitchPhysicalDtsJobToCloudRequest, runtime *dara.RuntimeOptions) (_result *SwitchPhysicalDtsJobToCloudResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Passes the connection information of the new database to DTS after a primary/secondary switchover. DTS restarts data synchronization from the checkpoint.
//
// @param request - SwitchSynchronizationEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchSynchronizationEndpointResponse
func (client *Client) SwitchSynchronizationEndpointWithContext(ctx context.Context, request *SwitchSynchronizationEndpointRequest, runtime *dara.RuntimeOptions) (_result *SwitchSynchronizationEndpointResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds tags to one or more data migration, synchronization, and change tracking instances by calling the TagResources operation.
//
// Description:
//
// If you have a large number of instances, you can create multiple tags and attach different tags to instances for categorization. Then, you can filter instances by tag.
//
// - A tag consists of a key-value pair. Tag keys must be unique within the same Alibaba Cloud account and region. Tag values do not have this restriction.
//
// - If the specified tag does not exist, the tag is automatically created and attached to the destination instance.
//
// - If the instance already has a tag with the same key, the existing tag is overwritten.
//
// - You can attach up to 20 tags to each instance.
//
// - You can invoke the operation to attach tags to up to 50 instances at a time.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades or downgrades the specifications of a DTS instance.
//
// Description:
//
// > - Downgrading DTS instance specifications is no longer supported.
//
// - If the source of a DTS instance is Redis 6.0 and incremental data updates exist, do not perform an upgrade. Otherwise, the DTS instance may fail and cannot be recovered. You must reconfigure the instance after a failure.
//
// @param request - TransferInstanceClassRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferInstanceClassResponse
func (client *Client) TransferInstanceClassWithContext(ctx context.Context, request *TransferInstanceClassRequest, runtime *dara.RuntimeOptions) (_result *TransferInstanceClassResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseCount) {
		query["DatabaseCount"] = request.DatabaseCount
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Transforms the payment method of a DTS instance.
//
// Description:
//
// <props="china">Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://www.aliyun.com/price/product#/dts/detail) of Data Transmission Service (DTS).
//
// <props="intl">Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/zh/product/data-transmission-service/pricing) of Data Transmission Service (DTS).
//
// - To avoid resource waste, confirm the payment method transformation before you perform the operation.
//
// - Data migration instances support only the pay-as-you-go billing method. No transformation is required.
//
// <props="china">
//
// - Serverless instances do not support payment method transformation.
//
// @param request - TransferPayTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferPayTypeResponse
func (client *Client) TransferPayTypeWithContext(ctx context.Context, request *TransferPayTypeRequest, runtime *dara.RuntimeOptions) (_result *TransferPayTypeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds tags from data migration, synchronization, and change tracking instances.
//
// Description:
//
// > After a tag is unbound from an instance, the tag is automatically deleted if it is not bound to any other instance.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades the synchronization topology of a DTS data synchronization instance from one-way synchronization to two-way synchronization.
//
// Description:
//
// <props="china">Before you use this operation, make sure that you fully understand the billing methods and [pricing](https://www.aliyun.com/price/product#/dts/detail) of ApsaraDB DTS.
//
// <props="intl">Before you use this operation, make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/zh/product/data-transmission-service/pricing) of ApsaraDB DTS.
//
// Before you begin:
//
// - The database type of both the source instance and the destination instance of the data synchronization node must be **MySQL**.
//
// - The synchronization topology of the data synchronization node must be **one-way synchronization**.
//
// - The data synchronization node must be in the **Synchronizing*	- state.
//
// - During the upgrade, data synchronization may experience a latency of approximately 5 seconds. Perform this operation during off-peak hours.
//
// @param request - UpgradeTwoWayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeTwoWayResponse
func (client *Client) UpgradeTwoWayWithContext(ctx context.Context, request *UpgradeTwoWayRequest, runtime *dara.RuntimeOptions) (_result *UpgradeTwoWayResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IP addresses of DTS servers by calling the WhiteIpList operation.
//
// Description:
//
// <props="china">If the **source or destination instance*	- is a **self-managed database*	- or a **third-party ApsaraDB database**, you need to invoke this operation to query the IP addresses of DTS servers, and then add the returned IP addresses to the security settings (typically the firewall) of the source or destination instance. For more information about how to add IP addresses, see [Add the CIDR blocks of DTS servers to the whitelist of a self-managed database for migration, synchronization, or subscribe](https://help.aliyun.com/document_detail/84900.html).
//
// <props="intl">If the **source or destination instance*	- is a **self-managed database*	- or a **third-party ApsaraDB database**, you need to invoke this operation to query the IP addresses of DTS servers, and then add the returned IP addresses to the security settings (typically the firewall) of the source or destination instance. For more information about how to add IP addresses, see [Add the CIDR blocks of DTS servers to the whitelist of a self-managed database](https://help.aliyun.com/document_detail/176627.html).
//
// > If the **source or destination database*	- is an **Alibaba Cloud database instance*	- (such as ApsaraDB RDS or ApsaraDB for MongoDB) or a **self-managed database hosted on ECS**, the system automatically adds the IP addresses of DTS servers to the security settings of the instance when you click **Authorize Whitelist and Proceed to Next Step*	- during the configuration of the source or destination instance. You do not need to manually add the IP addresses.
//
// @param request - WhiteIpListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WhiteIpListResponse
func (client *Client) WhiteIpListWithContext(ctx context.Context, request *WhiteIpListRequest, runtime *dara.RuntimeOptions) (_result *WhiteIpListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
