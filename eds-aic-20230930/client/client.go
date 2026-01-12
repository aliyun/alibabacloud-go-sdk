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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("eds-aic"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Attaches an Android Debug Bridge (ADB) key pair to one or more cloud phone instances.
//
// Description:
//
//	  You can attach to an ADB key pair only to cloud phone instances in the Running state.
//
//		- After you attach an ADB key pair, make sure the private key of the ADB key pair is copied to the ~/.android directory (macOS or Linux operating systems) or the C:\\Users\\Username.android directory (Windows operating systems). In addition, you must run the adb kill-server command to restart the ADB process to ensure correct ADB connection. Otherwise, ADB connection may fail due to authentication exceptions.
//
// @param request - AttachKeyPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachKeyPairResponse
func (client *Client) AttachKeyPairWithOptions(request *AttachKeyPairRequest, runtime *dara.RuntimeOptions) (_result *AttachKeyPairResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.KeyPairId) {
		query["KeyPairId"] = request.KeyPairId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachKeyPair"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachKeyPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches an Android Debug Bridge (ADB) key pair to one or more cloud phone instances.
//
// Description:
//
//	  You can attach to an ADB key pair only to cloud phone instances in the Running state.
//
//		- After you attach an ADB key pair, make sure the private key of the ADB key pair is copied to the ~/.android directory (macOS or Linux operating systems) or the C:\\Users\\Username.android directory (Windows operating systems). In addition, you must run the adb kill-server command to restart the ADB process to ensure correct ADB connection. Otherwise, ADB connection may fail due to authentication exceptions.
//
// @param request - AttachKeyPairRequest
//
// @return AttachKeyPairResponse
func (client *Client) AttachKeyPair(request *AttachKeyPairRequest) (_result *AttachKeyPairResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachKeyPairResponse{}
	_body, _err := client.AttachKeyPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Authorize/unauthorize Android instances for users.
//
// Description:
//
// Instance states that support user assignment: Available, Shutting Down, Stopped, Starting, Backing Up, Restoring, Backup Failed, Restore Failed.
//
// Instance states that support unassignment: Available, Shutting Down, Stopped, Starting, Backing Up, Restoring, Backup Failed, Restore Failed, Expired, Overdue, Deleted.
//
// @param request - AuthorizeAndroidInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeAndroidInstanceResponse
func (client *Client) AuthorizeAndroidInstanceWithOptions(request *AuthorizeAndroidInstanceRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeAndroidInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIds) {
		query["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	if !dara.IsNil(request.AuthorizeUserId) {
		query["AuthorizeUserId"] = request.AuthorizeUserId
	}

	if !dara.IsNil(request.UnAuthorizeUserId) {
		query["UnAuthorizeUserId"] = request.UnAuthorizeUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeAndroidInstance"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeAndroidInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorize/unauthorize Android instances for users.
//
// Description:
//
// Instance states that support user assignment: Available, Shutting Down, Stopped, Starting, Backing Up, Restoring, Backup Failed, Restore Failed.
//
// Instance states that support unassignment: Available, Shutting Down, Stopped, Starting, Backing Up, Restoring, Backup Failed, Restore Failed, Expired, Overdue, Deleted.
//
// @param request - AuthorizeAndroidInstanceRequest
//
// @return AuthorizeAndroidInstanceResponse
func (client *Client) AuthorizeAndroidInstance(request *AuthorizeAndroidInstanceRequest) (_result *AuthorizeAndroidInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AuthorizeAndroidInstanceResponse{}
	_body, _err := client.AuthorizeAndroidInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 整机备份
//
// @param request - BackupAndroidInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BackupAndroidInstanceResponse
func (client *Client) BackupAndroidInstanceWithOptions(request *BackupAndroidInstanceRequest, runtime *dara.RuntimeOptions) (_result *BackupAndroidInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIdList) {
		query["AndroidInstanceIdList"] = request.AndroidInstanceIdList
	}

	if !dara.IsNil(request.BackupFileName) {
		query["BackupFileName"] = request.BackupFileName
	}

	if !dara.IsNil(request.BackupFilePath) {
		query["BackupFilePath"] = request.BackupFilePath
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.UploadEndpoint) {
		query["UploadEndpoint"] = request.UploadEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BackupAndroidInstance"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BackupAndroidInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 整机备份
//
// @param request - BackupAndroidInstanceRequest
//
// @return BackupAndroidInstanceResponse
func (client *Client) BackupAndroidInstance(request *BackupAndroidInstanceRequest) (_result *BackupAndroidInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BackupAndroidInstanceResponse{}
	_body, _err := client.BackupAndroidInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 应用备份
//
// @param request - BackupAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BackupAppResponse
func (client *Client) BackupAppWithOptions(request *BackupAppRequest, runtime *dara.RuntimeOptions) (_result *BackupAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIdList) {
		query["AndroidInstanceIdList"] = request.AndroidInstanceIdList
	}

	if !dara.IsNil(request.BackupFileName) {
		query["BackupFileName"] = request.BackupFileName
	}

	if !dara.IsNil(request.BackupFilePath) {
		query["BackupFilePath"] = request.BackupFilePath
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SourceAppList) {
		query["SourceAppList"] = request.SourceAppList
	}

	if !dara.IsNil(request.UploadEndpoint) {
		query["UploadEndpoint"] = request.UploadEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BackupApp"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BackupAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 应用备份
//
// @param request - BackupAppRequest
//
// @return BackupAppResponse
func (client *Client) BackupApp(request *BackupAppRequest) (_result *BackupAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BackupAppResponse{}
	_body, _err := client.BackupAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates and uploads backup files.
//
// Description:
//
// Currently, this operation allows you to upload only backup files generated by cloud phones to Object Storage Service (OSS) buckets.
//
// @param request - BackupFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BackupFileResponse
func (client *Client) BackupFileWithOptions(request *BackupFileRequest, runtime *dara.RuntimeOptions) (_result *BackupFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIdList) {
		query["AndroidInstanceIdList"] = request.AndroidInstanceIdList
	}

	if !dara.IsNil(request.BackupAll) {
		query["BackupAll"] = request.BackupAll
	}

	if !dara.IsNil(request.BackupFileName) {
		query["BackupFileName"] = request.BackupFileName
	}

	if !dara.IsNil(request.BackupFilePath) {
		query["BackupFilePath"] = request.BackupFilePath
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExcludeSourceFilePathList) {
		query["ExcludeSourceFilePathList"] = request.ExcludeSourceFilePathList
	}

	if !dara.IsNil(request.SourceAppList) {
		query["SourceAppList"] = request.SourceAppList
	}

	if !dara.IsNil(request.SourceFilePathList) {
		query["SourceFilePathList"] = request.SourceFilePathList
	}

	if !dara.IsNil(request.UploadEndpoint) {
		query["UploadEndpoint"] = request.UploadEndpoint
	}

	if !dara.IsNil(request.UploadType) {
		query["UploadType"] = request.UploadType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BackupFile"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BackupFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates and uploads backup files.
//
// Description:
//
// Currently, this operation allows you to upload only backup files generated by cloud phones to Object Storage Service (OSS) buckets.
//
// @param request - BackupFileRequest
//
// @return BackupFileResponse
func (client *Client) BackupFile(request *BackupFileRequest) (_result *BackupFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BackupFileResponse{}
	_body, _err := client.BackupFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves connection tickets in batch.
//
// @param request - BatchGetAcpConnectionTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetAcpConnectionTicketResponse
func (client *Client) BatchGetAcpConnectionTicketWithOptions(request *BatchGetAcpConnectionTicketRequest, runtime *dara.RuntimeOptions) (_result *BatchGetAcpConnectionTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.InstanceGroupId) {
		query["InstanceGroupId"] = request.InstanceGroupId
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.InstanceTasks) {
		query["InstanceTasks"] = request.InstanceTasks
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchGetAcpConnectionTicket"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchGetAcpConnectionTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves connection tickets in batch.
//
// @param request - BatchGetAcpConnectionTicketRequest
//
// @return BatchGetAcpConnectionTicketResponse
func (client *Client) BatchGetAcpConnectionTicket(request *BatchGetAcpConnectionTicketRequest) (_result *BatchGetAcpConnectionTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchGetAcpConnectionTicketResponse{}
	_body, _err := client.BatchGetAcpConnectionTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改云手机矩阵的配置
//
// @param request - ChangeCloudPhoneNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeCloudPhoneNodeResponse
func (client *Client) ChangeCloudPhoneNodeWithOptions(request *ChangeCloudPhoneNodeRequest, runtime *dara.RuntimeOptions) (_result *ChangeCloudPhoneNodeResponse, _err error) {
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

	if !dara.IsNil(request.DisplayConfig) {
		query["DisplayConfig"] = request.DisplayConfig
	}

	if !dara.IsNil(request.DownBandwidthLimit) {
		query["DownBandwidthLimit"] = request.DownBandwidthLimit
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.PhoneCount) {
		query["PhoneCount"] = request.PhoneCount
	}

	if !dara.IsNil(request.PhoneDataVolume) {
		query["PhoneDataVolume"] = request.PhoneDataVolume
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.ShareDataVolume) {
		query["ShareDataVolume"] = request.ShareDataVolume
	}

	if !dara.IsNil(request.SwapSize) {
		query["SwapSize"] = request.SwapSize
	}

	if !dara.IsNil(request.UpBandwidthLimit) {
		query["UpBandwidthLimit"] = request.UpBandwidthLimit
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeCloudPhoneNode"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeCloudPhoneNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改云手机矩阵的配置
//
// @param request - ChangeCloudPhoneNodeRequest
//
// @return ChangeCloudPhoneNodeResponse
func (client *Client) ChangeCloudPhoneNode(request *ChangeCloudPhoneNodeRequest) (_result *ChangeCloudPhoneNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeCloudPhoneNodeResponse{}
	_body, _err := client.ChangeCloudPhoneNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Check the resource inventory.
//
// @param request - CheckResourceStockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckResourceStockResponse
func (client *Client) CheckResourceStockWithOptions(request *CheckResourceStockRequest, runtime *dara.RuntimeOptions) (_result *CheckResourceStockResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcpSpecId) {
		query["AcpSpecId"] = request.AcpSpecId
	}

	if !dara.IsNil(request.Amount) {
		query["Amount"] = request.Amount
	}

	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.GpuAcceleration) {
		query["GpuAcceleration"] = request.GpuAcceleration
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckResourceStock"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckResourceStockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Check the resource inventory.
//
// @param request - CheckResourceStockRequest
//
// @return CheckResourceStockResponse
func (client *Client) CheckResourceStock(request *CheckResourceStockRequest) (_result *CheckResourceStockResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckResourceStockResponse{}
	_body, _err := client.CheckResourceStockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates pay-as-you-go or subscription instance groups.
//
// Description:
//
// Before creating an instance group, ensure you understand the [billing methods](https://help.aliyun.com/document_detail/2807121.html) supported by Cloud Phone.
//
//   - If the billing method of an instance group is PrePaid, AutoPay is set to false by default. In this case, you need to go to [Expenses and Costs](https://usercenter2-intl.aliyun.com/order/list) to manually complete the payment.
//
//   - You can also set AutoPay to true based on your business requirements.
//
// @param tmpReq - CreateAndroidInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAndroidInstanceGroupResponse
func (client *Client) CreateAndroidInstanceGroupWithOptions(tmpReq *CreateAndroidInstanceGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateAndroidInstanceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAndroidInstanceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NetworkInfo) {
		request.NetworkInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NetworkInfo, dara.String("NetworkInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		query["Amount"] = request.Amount
	}

	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.BandwidthPackageId) {
		query["BandwidthPackageId"] = request.BandwidthPackageId
	}

	if !dara.IsNil(request.BandwidthPackageType) {
		query["BandwidthPackageType"] = request.BandwidthPackageType
	}

	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableIpv6) {
		query["EnableIpv6"] = request.EnableIpv6
	}

	if !dara.IsNil(request.GpuAcceleration) {
		query["GpuAcceleration"] = request.GpuAcceleration
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InstanceGroupName) {
		query["InstanceGroupName"] = request.InstanceGroupName
	}

	if !dara.IsNil(request.InstanceGroupSpec) {
		query["InstanceGroupSpec"] = request.InstanceGroupSpec
	}

	if !dara.IsNil(request.Ipv6Bandwidth) {
		query["Ipv6Bandwidth"] = request.Ipv6Bandwidth
	}

	if !dara.IsNil(request.KeyPairId) {
		query["KeyPairId"] = request.KeyPairId
	}

	if !dara.IsNil(request.NetworkInfoShrink) {
		query["NetworkInfo"] = request.NetworkInfoShrink
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.NumberOfInstances) {
		query["NumberOfInstances"] = request.NumberOfInstances
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.StreamMode) {
		query["StreamMode"] = request.StreamMode
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAndroidInstanceGroup"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAndroidInstanceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates pay-as-you-go or subscription instance groups.
//
// Description:
//
// Before creating an instance group, ensure you understand the [billing methods](https://help.aliyun.com/document_detail/2807121.html) supported by Cloud Phone.
//
//   - If the billing method of an instance group is PrePaid, AutoPay is set to false by default. In this case, you need to go to [Expenses and Costs](https://usercenter2-intl.aliyun.com/order/list) to manually complete the payment.
//
//   - You can also set AutoPay to true based on your business requirements.
//
// @param request - CreateAndroidInstanceGroupRequest
//
// @return CreateAndroidInstanceGroupResponse
func (client *Client) CreateAndroidInstanceGroup(request *CreateAndroidInstanceGroupRequest) (_result *CreateAndroidInstanceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAndroidInstanceGroupResponse{}
	_body, _err := client.CreateAndroidInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an Android application.
//
// Description:
//
// When creating an app, you can provide app information to the system in one of the following ways:
//
//   - Way 1: Apps from the Application Center
//
//   - You can use one of the following methods:
//
//   - Method 1: Pass in the `FileName` and `FilePath` parameters at the same time.
//
//   - Method 2: Pass in the `OssAppUrl` parameter
//
//   - Rule: If your app is from the Alibaba Cloud Workspace Application Center, you must use either Method 1 or Method 2. If both are used, Method 1 takes priority.
//
//   - Condition: Before you proceed, log on to the [Elastic Desktop Service (EDS) Enterprise console](https://eds.console.aliyun.com/osshelp) and follow the on-screen instructions to upload the app file to the Application Center to obtain the values of the `FileName`, `FilePath`, and `OssAppUrl` parameters.
//
//   - Way 2: Custom apps
//
//   - Pass in the `CustomAppInfo` parameter.
//
//   - Rule: If you pass in the `CustomAppInfo` parameter, all six fields within it are required.
//
// >  If Way 1 and Way 2 are adopted simultaneously, the information from Way 2 takes priority.
//
// @param tmpReq - CreateAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppResponse
func (client *Client) CreateAppWithOptions(tmpReq *CreateAppRequest, runtime *dara.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CustomAppInfo) {
		request.CustomAppInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomAppInfo, dara.String("CustomAppInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.CustomAppInfoShrink) {
		query["CustomAppInfo"] = request.CustomAppInfoShrink
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.IconUrl) {
		query["IconUrl"] = request.IconUrl
	}

	if !dara.IsNil(request.InstallParam) {
		query["InstallParam"] = request.InstallParam
	}

	if !dara.IsNil(request.OssAppUrl) {
		query["OssAppUrl"] = request.OssAppUrl
	}

	if !dara.IsNil(request.SignApk) {
		query["SignApk"] = request.SignApk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApp"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an Android application.
//
// Description:
//
// When creating an app, you can provide app information to the system in one of the following ways:
//
//   - Way 1: Apps from the Application Center
//
//   - You can use one of the following methods:
//
//   - Method 1: Pass in the `FileName` and `FilePath` parameters at the same time.
//
//   - Method 2: Pass in the `OssAppUrl` parameter
//
//   - Rule: If your app is from the Alibaba Cloud Workspace Application Center, you must use either Method 1 or Method 2. If both are used, Method 1 takes priority.
//
//   - Condition: Before you proceed, log on to the [Elastic Desktop Service (EDS) Enterprise console](https://eds.console.aliyun.com/osshelp) and follow the on-screen instructions to upload the app file to the Application Center to obtain the values of the `FileName`, `FilePath`, and `OssAppUrl` parameters.
//
//   - Way 2: Custom apps
//
//   - Pass in the `CustomAppInfo` parameter.
//
//   - Rule: If you pass in the `CustomAppInfo` parameter, all six fields within it are required.
//
// >  If Way 1 and Way 2 are adopted simultaneously, the information from Way 2 takes priority.
//
// @param request - CreateAppRequest
//
// @return CreateAppResponse
func (client *Client) CreateApp(request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a cloud phone matrix.
//
// @param tmpReq - CreateCloudPhoneNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudPhoneNodeResponse
func (client *Client) CreateCloudPhoneNodeWithOptions(tmpReq *CreateCloudPhoneNodeRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudPhoneNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateCloudPhoneNodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DisplayConfig) {
		request.DisplayConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DisplayConfig, dara.String("DisplayConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NetworkInfo) {
		request.NetworkInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NetworkInfo, dara.String("NetworkInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.BandwidthPackageId) {
		query["BandwidthPackageId"] = request.BandwidthPackageId
	}

	if !dara.IsNil(request.BandwidthPackageType) {
		query["BandwidthPackageType"] = request.BandwidthPackageType
	}

	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.Count) {
		query["Count"] = request.Count
	}

	if !dara.IsNil(request.DownBandwidthLimit) {
		query["DownBandwidthLimit"] = request.DownBandwidthLimit
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.IsSingleImgDisk) {
		query["IsSingleImgDisk"] = request.IsSingleImgDisk
	}

	if !dara.IsNil(request.NetworkId) {
		query["NetworkId"] = request.NetworkId
	}

	if !dara.IsNil(request.NetworkInfoShrink) {
		query["NetworkInfo"] = request.NetworkInfoShrink
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.NodeName) {
		query["NodeName"] = request.NodeName
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PhoneCount) {
		query["PhoneCount"] = request.PhoneCount
	}

	if !dara.IsNil(request.PhoneDataVolume) {
		query["PhoneDataVolume"] = request.PhoneDataVolume
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.ResolutionHeight) {
		query["ResolutionHeight"] = request.ResolutionHeight
	}

	if !dara.IsNil(request.ResolutionWidth) {
		query["ResolutionWidth"] = request.ResolutionWidth
	}

	if !dara.IsNil(request.ServerShareDataVolume) {
		query["ServerShareDataVolume"] = request.ServerShareDataVolume
	}

	if !dara.IsNil(request.ServerType) {
		query["ServerType"] = request.ServerType
	}

	if !dara.IsNil(request.StreamMode) {
		query["StreamMode"] = request.StreamMode
	}

	if !dara.IsNil(request.SwapSize) {
		query["SwapSize"] = request.SwapSize
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UpBandwidthLimit) {
		query["UpBandwidthLimit"] = request.UpBandwidthLimit
	}

	if !dara.IsNil(request.UseTemplate) {
		query["UseTemplate"] = request.UseTemplate
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DisplayConfigShrink) {
		body["DisplayConfig"] = request.DisplayConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudPhoneNode"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudPhoneNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a cloud phone matrix.
//
// @param request - CreateCloudPhoneNodeRequest
//
// @return CreateCloudPhoneNodeResponse
func (client *Client) CreateCloudPhoneNode(request *CreateCloudPhoneNodeRequest) (_result *CreateCloudPhoneNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCloudPhoneNodeResponse{}
	_body, _err := client.CreateCloudPhoneNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom image from a cloud phone instance.
//
// @param request - CreateCustomImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomImageResponse
func (client *Client) CreateCustomImageWithOptions(request *CreateCustomImageRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ImageName) {
		body["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomImage"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom image from a cloud phone instance.
//
// @param request - CreateCustomImageRequest
//
// @return CreateCustomImageResponse
func (client *Client) CreateCustomImage(request *CreateCustomImageRequest) (_result *CreateCustomImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCustomImageResponse{}
	_body, _err := client.CreateCustomImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an Android Debug Bridge (ADB) key pair. The system retains the public key and provides a PEM-encoded private key in PKCS#8 format, adhering to the ADB connection specification. You must securely store the private key.
//
// Description:
//
// In addition to using the CreateKeyPair operation to generate a key pair, you can also create one by using the ADB tool and upload it to the Cloud Phone console. The usage of this key pair is identical to that of a system-generated key pair.
//
// Each tenant can create up to 500 key pairs.
//
// @param request - CreateKeyPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKeyPairResponse
func (client *Client) CreateKeyPairWithOptions(request *CreateKeyPairRequest, runtime *dara.RuntimeOptions) (_result *CreateKeyPairResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyPairName) {
		query["KeyPairName"] = request.KeyPairName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateKeyPair"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateKeyPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an Android Debug Bridge (ADB) key pair. The system retains the public key and provides a PEM-encoded private key in PKCS#8 format, adhering to the ADB connection specification. You must securely store the private key.
//
// Description:
//
// In addition to using the CreateKeyPair operation to generate a key pair, you can also create one by using the ADB tool and upload it to the Cloud Phone console. The usage of this key pair is identical to that of a system-generated key pair.
//
// Each tenant can create up to 500 key pairs.
//
// @param request - CreateKeyPairRequest
//
// @return CreateKeyPairResponse
func (client *Client) CreateKeyPair(request *CreateKeyPairRequest) (_result *CreateKeyPairResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateKeyPairResponse{}
	_body, _err := client.CreateKeyPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a policy.
//
// @param tmpReq - CreatePolicyGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyGroupResponse
func (client *Client) CreatePolicyGroupWithOptions(tmpReq *CreatePolicyGroupRequest, runtime *dara.RuntimeOptions) (_result *CreatePolicyGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePolicyGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NetRedirectPolicy) {
		request.NetRedirectPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NetRedirectPolicy, dara.String("NetRedirectPolicy"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Watermark) {
		request.WatermarkShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Watermark, dara.String("Watermark"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CameraRedirect) {
		body["CameraRedirect"] = request.CameraRedirect
	}

	if !dara.IsNil(request.Clipboard) {
		body["Clipboard"] = request.Clipboard
	}

	if !dara.IsNil(request.Html5FileTransfer) {
		body["Html5FileTransfer"] = request.Html5FileTransfer
	}

	if !dara.IsNil(request.LocalDrive) {
		body["LocalDrive"] = request.LocalDrive
	}

	if !dara.IsNil(request.LockResolution) {
		body["LockResolution"] = request.LockResolution
	}

	if !dara.IsNil(request.NetRedirectPolicyShrink) {
		body["NetRedirectPolicy"] = request.NetRedirectPolicyShrink
	}

	if !dara.IsNil(request.PolicyGroupName) {
		body["PolicyGroupName"] = request.PolicyGroupName
	}

	if !dara.IsNil(request.PolicyType) {
		body["PolicyType"] = request.PolicyType
	}

	if !dara.IsNil(request.ResolutionHeight) {
		body["ResolutionHeight"] = request.ResolutionHeight
	}

	if !dara.IsNil(request.ResolutionWidth) {
		body["ResolutionWidth"] = request.ResolutionWidth
	}

	if !dara.IsNil(request.WatermarkShrink) {
		body["Watermark"] = request.WatermarkShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePolicyGroup"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePolicyGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a policy.
//
// @param request - CreatePolicyGroupRequest
//
// @return CreatePolicyGroupResponse
func (client *Client) CreatePolicyGroup(request *CreatePolicyGroupRequest) (_result *CreatePolicyGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePolicyGroupResponse{}
	_body, _err := client.CreatePolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a screenshot of a cloud phone instance.
//
// Description:
//
// You can call this operation to create a screenshot of a cloud phone instance and upload it to the default Object Storage Service (OSS) bucket. The operation returns a task ID, which you can use with the DescribeTasks operation to get the download link for the screenshot.
//
// @param request - CreateScreenshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScreenshotResponse
func (client *Client) CreateScreenshotWithOptions(request *CreateScreenshotRequest, runtime *dara.RuntimeOptions) (_result *CreateScreenshotResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIdList) {
		query["AndroidInstanceIdList"] = request.AndroidInstanceIdList
	}

	if !dara.IsNil(request.OssBucketName) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.ScreenshotId) {
		query["ScreenshotId"] = request.ScreenshotId
	}

	if !dara.IsNil(request.SkipCheckPolicyConfig) {
		query["SkipCheckPolicyConfig"] = request.SkipCheckPolicyConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScreenshot"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScreenshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a screenshot of a cloud phone instance.
//
// Description:
//
// You can call this operation to create a screenshot of a cloud phone instance and upload it to the default Object Storage Service (OSS) bucket. The operation returns a task ID, which you can use with the DescribeTasks operation to get the download link for the screenshot.
//
// @param request - CreateScreenshotRequest
//
// @return CreateScreenshotResponse
func (client *Client) CreateScreenshot(request *CreateScreenshotRequest) (_result *CreateScreenshotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateScreenshotResponse{}
	_body, _err := client.CreateScreenshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建系统属性模板
//
// @param tmpReq - CreateSystemPropertyTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSystemPropertyTemplateResponse
func (client *Client) CreateSystemPropertyTemplateWithOptions(tmpReq *CreateSystemPropertyTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateSystemPropertyTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSystemPropertyTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SystemPropertyInfo) {
		request.SystemPropertyInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SystemPropertyInfo, dara.String("SystemPropertyInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableAuto) {
		query["EnableAuto"] = request.EnableAuto
	}

	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.SystemPropertyInfoShrink) {
		query["SystemPropertyInfo"] = request.SystemPropertyInfoShrink
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSystemPropertyTemplate"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSystemPropertyTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建系统属性模板
//
// @param request - CreateSystemPropertyTemplateRequest
//
// @return CreateSystemPropertyTemplateResponse
func (client *Client) CreateSystemPropertyTemplate(request *CreateSystemPropertyTemplateRequest) (_result *CreateSystemPropertyTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSystemPropertyTemplateResponse{}
	_body, _err := client.CreateSystemPropertyTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete an instance group.
//
// Description:
//
// You can delete only pay-as-you-go instance groups.
//
// You can delete subscription instance groups only after they expire.
//
// @param request - DeleteAndroidInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAndroidInstanceGroupResponse
func (client *Client) DeleteAndroidInstanceGroupWithOptions(request *DeleteAndroidInstanceGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteAndroidInstanceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceGroupIds) {
		query["InstanceGroupIds"] = request.InstanceGroupIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAndroidInstanceGroup"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAndroidInstanceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete an instance group.
//
// Description:
//
// You can delete only pay-as-you-go instance groups.
//
// You can delete subscription instance groups only after they expire.
//
// @param request - DeleteAndroidInstanceGroupRequest
//
// @return DeleteAndroidInstanceGroupResponse
func (client *Client) DeleteAndroidInstanceGroup(request *DeleteAndroidInstanceGroupRequest) (_result *DeleteAndroidInstanceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAndroidInstanceGroupResponse{}
	_body, _err := client.DeleteAndroidInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an application. Before you delete an application, make sure that the application is not installed on any instances.
//
// @param request - DeleteAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppsResponse
func (client *Client) DeleteAppsWithOptions(request *DeleteAppsRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppIdList) {
		query["AppIdList"] = request.AppIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApps"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an application. Before you delete an application, make sure that the application is not installed on any instances.
//
// @param request - DeleteAppsRequest
//
// @return DeleteAppsResponse
func (client *Client) DeleteApps(request *DeleteAppsRequest) (_result *DeleteAppsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAppsResponse{}
	_body, _err := client.DeleteAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除备份文件
//
// @param request - DeleteBackupFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBackupFileResponse
func (client *Client) DeleteBackupFileWithOptions(request *DeleteBackupFileRequest, runtime *dara.RuntimeOptions) (_result *DeleteBackupFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupFileIdList) {
		query["BackupFileIdList"] = request.BackupFileIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBackupFile"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBackupFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除备份文件
//
// @param request - DeleteBackupFileRequest
//
// @return DeleteBackupFileResponse
func (client *Client) DeleteBackupFile(request *DeleteBackupFileRequest) (_result *DeleteBackupFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBackupFileResponse{}
	_body, _err := client.DeleteBackupFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a cloud phone matrix.
//
// Description:
//
// Before you proceed, make sure that the cloud phone matrix that you want to delete expired.
//
// @param request - DeleteCloudPhoneNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudPhoneNodesResponse
func (client *Client) DeleteCloudPhoneNodesWithOptions(request *DeleteCloudPhoneNodesRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudPhoneNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeIds) {
		body["NodeIds"] = request.NodeIds
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloudPhoneNodes"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudPhoneNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a cloud phone matrix.
//
// Description:
//
// Before you proceed, make sure that the cloud phone matrix that you want to delete expired.
//
// @param request - DeleteCloudPhoneNodesRequest
//
// @return DeleteCloudPhoneNodesResponse
func (client *Client) DeleteCloudPhoneNodes(request *DeleteCloudPhoneNodesRequest) (_result *DeleteCloudPhoneNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCloudPhoneNodesResponse{}
	_body, _err := client.DeleteCloudPhoneNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom image.
//
// Description:
//
// You cannot delete an image that is currently in use by an instance group.
//
// @param tmpReq - DeleteImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteImagesResponse
func (client *Client) DeleteImagesWithOptions(tmpReq *DeleteImagesRequest, runtime *dara.RuntimeOptions) (_result *DeleteImagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteImagesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ImageIds) {
		request.ImageIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageIds, dara.String("ImageIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageIdsShrink) {
		body["ImageIds"] = request.ImageIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteImages"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom image.
//
// Description:
//
// You cannot delete an image that is currently in use by an instance group.
//
// @param request - DeleteImagesRequest
//
// @return DeleteImagesResponse
func (client *Client) DeleteImages(request *DeleteImagesRequest) (_result *DeleteImagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteImagesResponse{}
	_body, _err := client.DeleteImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes Android Debug Bridge (ADB) key pairs.
//
// Description:
//
//	  If a cloud phone instance is currently associated with the ADB key pair you intend to delete, the ADB key pair cannot be deleted.
//
//		- Once an ADB key pair is deleted, it cannot be retrieved or queried by using the DescribeKeyPairs operation.
//
// @param request - DeleteKeyPairsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteKeyPairsResponse
func (client *Client) DeleteKeyPairsWithOptions(request *DeleteKeyPairsRequest, runtime *dara.RuntimeOptions) (_result *DeleteKeyPairsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyPairIds) {
		query["KeyPairIds"] = request.KeyPairIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteKeyPairs"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteKeyPairsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes Android Debug Bridge (ADB) key pairs.
//
// Description:
//
//	  If a cloud phone instance is currently associated with the ADB key pair you intend to delete, the ADB key pair cannot be deleted.
//
//		- Once an ADB key pair is deleted, it cannot be retrieved or queried by using the DescribeKeyPairs operation.
//
// @param request - DeleteKeyPairsRequest
//
// @return DeleteKeyPairsResponse
func (client *Client) DeleteKeyPairs(request *DeleteKeyPairsRequest) (_result *DeleteKeyPairsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteKeyPairsResponse{}
	_body, _err := client.DeleteKeyPairsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a policy.
//
// @param request - DeletePolicyGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyGroupResponse
func (client *Client) DeletePolicyGroupWithOptions(request *DeletePolicyGroupRequest, runtime *dara.RuntimeOptions) (_result *DeletePolicyGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyGroupIds) {
		query["PolicyGroupIds"] = request.PolicyGroupIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolicyGroup"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolicyGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a policy.
//
// @param request - DeletePolicyGroupRequest
//
// @return DeletePolicyGroupResponse
func (client *Client) DeletePolicyGroup(request *DeletePolicyGroupRequest) (_result *DeletePolicyGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePolicyGroupResponse{}
	_body, _err := client.DeletePolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除系统属性模板
//
// @param request - DeleteSystemPropertyTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSystemPropertyTemplatesResponse
func (client *Client) DeleteSystemPropertyTemplatesWithOptions(request *DeleteSystemPropertyTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DeleteSystemPropertyTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateIds) {
		query["TemplateIds"] = request.TemplateIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSystemPropertyTemplates"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSystemPropertyTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除系统属性模板
//
// @param request - DeleteSystemPropertyTemplatesRequest
//
// @return DeleteSystemPropertyTemplatesResponse
func (client *Client) DeleteSystemPropertyTemplates(request *DeleteSystemPropertyTemplatesRequest) (_result *DeleteSystemPropertyTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSystemPropertyTemplatesResponse{}
	_body, _err := client.DeleteSystemPropertyTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an instance group.
//
// @param request - DescribeAndroidInstanceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAndroidInstanceGroupsResponse
func (client *Client) DescribeAndroidInstanceGroupsWithOptions(request *DescribeAndroidInstanceGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAndroidInstanceGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.InstanceGroupIds) {
		query["InstanceGroupIds"] = request.InstanceGroupIds
	}

	if !dara.IsNil(request.InstanceGroupName) {
		query["InstanceGroupName"] = request.InstanceGroupName
	}

	if !dara.IsNil(request.KeyPairId) {
		query["KeyPairId"] = request.KeyPairId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.SaleMode) {
		query["SaleMode"] = request.SaleMode
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAndroidInstanceGroups"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAndroidInstanceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an instance group.
//
// @param request - DescribeAndroidInstanceGroupsRequest
//
// @return DescribeAndroidInstanceGroupsResponse
func (client *Client) DescribeAndroidInstanceGroups(request *DescribeAndroidInstanceGroupsRequest) (_result *DescribeAndroidInstanceGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAndroidInstanceGroupsResponse{}
	_body, _err := client.DescribeAndroidInstanceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries cloud phone instances.
//
// @param request - DescribeAndroidInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAndroidInstancesResponse
func (client *Client) DescribeAndroidInstancesWithOptions(request *DescribeAndroidInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAndroidInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIds) {
		query["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	if !dara.IsNil(request.AndroidInstanceName) {
		query["AndroidInstanceName"] = request.AndroidInstanceName
	}

	if !dara.IsNil(request.AppManagePolicyId) {
		query["AppManagePolicyId"] = request.AppManagePolicyId
	}

	if !dara.IsNil(request.AuthorizedUserId) {
		query["AuthorizedUserId"] = request.AuthorizedUserId
	}

	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.InstanceGroupId) {
		query["InstanceGroupId"] = request.InstanceGroupId
	}

	if !dara.IsNil(request.InstanceGroupIds) {
		query["InstanceGroupIds"] = request.InstanceGroupIds
	}

	if !dara.IsNil(request.InstanceGroupName) {
		query["InstanceGroupName"] = request.InstanceGroupName
	}

	if !dara.IsNil(request.KeyPairId) {
		query["KeyPairId"] = request.KeyPairId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.NodeName) {
		query["NodeName"] = request.NodeName
	}

	if !dara.IsNil(request.OfficeSiteIds) {
		query["OfficeSiteIds"] = request.OfficeSiteIds
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.QosRuleIds) {
		query["QosRuleIds"] = request.QosRuleIds
	}

	if !dara.IsNil(request.SaleMode) {
		query["SaleMode"] = request.SaleMode
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAndroidInstances"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAndroidInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries cloud phone instances.
//
// @param request - DescribeAndroidInstancesRequest
//
// @return DescribeAndroidInstancesResponse
func (client *Client) DescribeAndroidInstances(request *DescribeAndroidInstancesRequest) (_result *DescribeAndroidInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAndroidInstancesResponse{}
	_body, _err := client.DescribeAndroidInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries applications.
//
// @param request - DescribeAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppsResponse
func (client *Client) DescribeAppsWithOptions(request *DescribeAppsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppIdList) {
		query["AppIdList"] = request.AppIdList
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AppType) {
		query["AppType"] = request.AppType
	}

	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.InstallationStatus) {
		query["InstallationStatus"] = request.InstallationStatus
	}

	if !dara.IsNil(request.MD5) {
		query["MD5"] = request.MD5
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApps"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries applications.
//
// @param request - DescribeAppsRequest
//
// @return DescribeAppsResponse
func (client *Client) DescribeApps(request *DescribeAppsRequest) (_result *DescribeAppsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppsResponse{}
	_body, _err := client.DescribeAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries backup files.
//
// Description:
//
// Currently, this operation allows you to query only backup files generated by cloud phones that are stored in Object Storage Service (OSS) buckets.
//
// @param request - DescribeBackupFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupFilesResponse
func (client *Client) DescribeBackupFilesWithOptions(request *DescribeBackupFilesRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceId) {
		query["AndroidInstanceId"] = request.AndroidInstanceId
	}

	if !dara.IsNil(request.AndroidInstanceName) {
		query["AndroidInstanceName"] = request.AndroidInstanceName
	}

	if !dara.IsNil(request.BackupAll) {
		query["BackupAll"] = request.BackupAll
	}

	if !dara.IsNil(request.BackupFileId) {
		query["BackupFileId"] = request.BackupFileId
	}

	if !dara.IsNil(request.BackupFileName) {
		query["BackupFileName"] = request.BackupFileName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.InstanceGroupId) {
		query["InstanceGroupId"] = request.InstanceGroupId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SaleMode) {
		query["SaleMode"] = request.SaleMode
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StatusList) {
		query["StatusList"] = request.StatusList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupFiles"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries backup files.
//
// Description:
//
// Currently, this operation allows you to query only backup files generated by cloud phones that are stored in Object Storage Service (OSS) buckets.
//
// @param request - DescribeBackupFilesRequest
//
// @return DescribeBackupFilesResponse
func (client *Client) DescribeBackupFiles(request *DescribeBackupFilesRequest) (_result *DescribeBackupFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupFilesResponse{}
	_body, _err := client.DescribeBackupFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询bucket信息
//
// @param request - DescribeBucketsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBucketsResponse
func (client *Client) DescribeBucketsWithOptions(request *DescribeBucketsRequest, runtime *dara.RuntimeOptions) (_result *DescribeBucketsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileType) {
		query["FileType"] = request.FileType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBuckets"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBucketsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询bucket信息
//
// @param request - DescribeBucketsRequest
//
// @return DescribeBucketsResponse
func (client *Client) DescribeBuckets(request *DescribeBucketsRequest) (_result *DescribeBucketsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBucketsResponse{}
	_body, _err := client.DescribeBucketsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a cloud phone matrix.
//
// @param request - DescribeCloudPhoneNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudPhoneNodesResponse
func (client *Client) DescribeCloudPhoneNodesWithOptions(request *DescribeCloudPhoneNodesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudPhoneNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BandwidthPackageId) {
		query["BandwidthPackageId"] = request.BandwidthPackageId
	}

	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NodeIds) {
		query["NodeIds"] = request.NodeIds
	}

	if !dara.IsNil(request.NodeName) {
		query["NodeName"] = request.NodeName
	}

	if !dara.IsNil(request.ServerType) {
		query["ServerType"] = request.ServerType
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudPhoneNodes"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudPhoneNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a cloud phone matrix.
//
// @param request - DescribeCloudPhoneNodesRequest
//
// @return DescribeCloudPhoneNodesResponse
func (client *Client) DescribeCloudPhoneNodes(request *DescribeCloudPhoneNodesRequest) (_result *DescribeCloudPhoneNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudPhoneNodesResponse{}
	_body, _err := client.DescribeCloudPhoneNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询显示设置
//
// @param request - DescribeDisplayConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDisplayConfigResponse
func (client *Client) DescribeDisplayConfigWithOptions(request *DescribeDisplayConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDisplayConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIds) {
		body["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDisplayConfig"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDisplayConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询显示设置
//
// @param request - DescribeDisplayConfigRequest
//
// @return DescribeDisplayConfigResponse
func (client *Client) DescribeDisplayConfig(request *DescribeDisplayConfigRequest) (_result *DescribeDisplayConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDisplayConfigResponse{}
	_body, _err := client.DescribeDisplayConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries images.
//
// @param request - DescribeImageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImageListResponse
func (client *Client) DescribeImageListWithOptions(request *DescribeImageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeImageListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageBizTags) {
		query["ImageBizTags"] = request.ImageBizTags
	}

	if !dara.IsNil(request.ImagePackageType) {
		query["ImagePackageType"] = request.ImagePackageType
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.SystemType) {
		query["SystemType"] = request.SystemType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageName) {
		body["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.ImageType) {
		body["ImageType"] = request.ImageType
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeImageList"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeImageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries images.
//
// @param request - DescribeImageListRequest
//
// @return DescribeImageListResponse
func (client *Client) DescribeImageList(request *DescribeImageListRequest) (_result *DescribeImageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeImageListResponse{}
	_body, _err := client.DescribeImageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution results of commands.
//
// @param request - DescribeInvocationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInvocationsResponse
func (client *Client) DescribeInvocationsWithOptions(request *DescribeInvocationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInvocationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.InvocationId) {
		query["InvocationId"] = request.InvocationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInvocations"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution results of commands.
//
// @param request - DescribeInvocationsRequest
//
// @return DescribeInvocationsResponse
func (client *Client) DescribeInvocations(request *DescribeInvocationsRequest) (_result *DescribeInvocationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.DescribeInvocationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries one or more key pairs.
//
// @param request - DescribeKeyPairsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKeyPairsResponse
func (client *Client) DescribeKeyPairsWithOptions(request *DescribeKeyPairsRequest, runtime *dara.RuntimeOptions) (_result *DescribeKeyPairsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyPairIds) {
		query["KeyPairIds"] = request.KeyPairIds
	}

	if !dara.IsNil(request.KeyPairName) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeKeyPairs"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeKeyPairsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries one or more key pairs.
//
// @param request - DescribeKeyPairsRequest
//
// @return DescribeKeyPairsResponse
func (client *Client) DescribeKeyPairs(request *DescribeKeyPairsRequest) (_result *DescribeKeyPairsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeKeyPairsResponse{}
	_body, _err := client.DescribeKeyPairsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询指定监控项的最新监控数据
//
// @param request - DescribeMetricLastRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricLastResponse
func (client *Client) DescribeMetricLastWithOptions(request *DescribeMetricLastRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricLastResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIds) {
		body["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Length) {
		body["Length"] = request.Length
	}

	if !dara.IsNil(request.MetricNames) {
		body["MetricNames"] = request.MetricNames
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
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
		Action:      dara.String("DescribeMetricLast"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricLastResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定监控项的最新监控数据
//
// @param request - DescribeMetricLastRequest
//
// @return DescribeMetricLastResponse
func (client *Client) DescribeMetricLast(request *DescribeMetricLastRequest) (_result *DescribeMetricLastResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMetricLastResponse{}
	_body, _err := client.DescribeMetricLastWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询指定监控项的监控数据
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIds) {
		body["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceIds) {
		body["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.Length) {
		body["Length"] = request.Length
	}

	if !dara.IsNil(request.MetricNames) {
		body["MetricNames"] = request.MetricNames
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Period) {
		body["Period"] = request.Period
	}

	if !dara.IsNil(request.ProcessInfos) {
		body["ProcessInfos"] = request.ProcessInfos
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricList"),
		Version:     dara.String("2023-09-30"),
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
// 查询指定监控项的监控数据
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
// 查询指定监控项的最新监控数据
//
// @param request - DescribeMetricTopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricTopResponse
func (client *Client) DescribeMetricTopWithOptions(request *DescribeMetricTopRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricTopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIds) {
		body["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceIds) {
		body["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.Length) {
		body["Length"] = request.Length
	}

	if !dara.IsNil(request.MetricNames) {
		body["MetricNames"] = request.MetricNames
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Period) {
		body["Period"] = request.Period
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricTop"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricTopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定监控项的最新监控数据
//
// @param request - DescribeMetricTopRequest
//
// @return DescribeMetricTopResponse
func (client *Client) DescribeMetricTop(request *DescribeMetricTopRequest) (_result *DescribeMetricTopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMetricTopResponse{}
	_body, _err := client.DescribeMetricTopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query available regions.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.SaleMode) {
		query["SaleMode"] = request.SaleMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2023-09-30"),
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
// Query available regions.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query available specifications.
//
// @param request - DescribeSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSpecResponse
func (client *Client) DescribeSpecWithOptions(request *DescribeSpecRequest, runtime *dara.RuntimeOptions) (_result *DescribeSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.MatrixSpec) {
		query["MatrixSpec"] = request.MatrixSpec
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SaleMode) {
		query["SaleMode"] = request.SaleMode
	}

	if !dara.IsNil(request.SpecIds) {
		query["SpecIds"] = request.SpecIds
	}

	if !dara.IsNil(request.SpecStatus) {
		query["SpecStatus"] = request.SpecStatus
	}

	if !dara.IsNil(request.SpecType) {
		query["SpecType"] = request.SpecType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSpec"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query available specifications.
//
// @param request - DescribeSpecRequest
//
// @return DescribeSpecResponse
func (client *Client) DescribeSpec(request *DescribeSpecRequest) (_result *DescribeSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSpecResponse{}
	_body, _err := client.DescribeSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询系统属性模板
//
// @param request - DescribeSystemPropertyTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSystemPropertyTemplatesResponse
func (client *Client) DescribeSystemPropertyTemplatesWithOptions(request *DescribeSystemPropertyTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSystemPropertyTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TemplateIds) {
		query["TemplateIds"] = request.TemplateIds
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSystemPropertyTemplates"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSystemPropertyTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询系统属性模板
//
// @param request - DescribeSystemPropertyTemplatesRequest
//
// @return DescribeSystemPropertyTemplatesResponse
func (client *Client) DescribeSystemPropertyTemplates(request *DescribeSystemPropertyTemplatesRequest) (_result *DescribeSystemPropertyTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSystemPropertyTemplatesResponse{}
	_body, _err := client.DescribeSystemPropertyTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries tasks created for a cloud phone instance.
//
// Description:
//
//	  You can call the DescribeTasks operation to query the tasks created for one or more cloud phone instances.
//
//		- The system currently supports various tasks, including starting, stopping, restarting, and resetting cloud phone instances; backing up and restoring data; installing apps; and executing remote commands.
//
//		- You can use the Level field to specify the type of task. If Level is set to 1, it represents a batch task. If Level is set to 2, it represents an instance-level task.
//
// **Example**
//
// Assume you restart two cloud phone instances with the instance IDs acp-25nt4kk9whhok\\*\\*\\*\\	- and acp-j2taq887orj8l\\*\\*\\*\\*, and the returned request ID is B8ED2BA9-0C6A-5643-818F-B5D60A64\\*\\*\\*\\*. If you want to check the operation outcomes of the two cloud phone instances, you can call the DescribeTasks operation. You need to set the InvokeId request parameter to B8ED2BA9-0C6A-5643-818F-B5D60A64\\*\\*\\*\\*. If you only want to check the cloud phone instance with the ID acp-25nt4kk9whhok\\*\\*\\*\\*, you must set the ParentTaskId request parameter to the ID of the batch task and the AndroidInstanceId request parameter to acp-25nt4kk9whhok\\*\\*\\*\\	- when calling the DescribeTasks operation.
//
// @param request - DescribeTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTasksResponse
func (client *Client) DescribeTasksWithOptions(request *DescribeTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeTasksResponse, _err error) {
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

	if !dara.IsNil(request.InvokeId) {
		query["InvokeId"] = request.InvokeId
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Param) {
		query["Param"] = request.Param
	}

	if !dara.IsNil(request.ParentTaskId) {
		query["ParentTaskId"] = request.ParentTaskId
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.TaskIds) {
		query["TaskIds"] = request.TaskIds
	}

	if !dara.IsNil(request.TaskStatus) {
		query["TaskStatus"] = request.TaskStatus
	}

	if !dara.IsNil(request.TaskStatuses) {
		query["TaskStatuses"] = request.TaskStatuses
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.TaskTypes) {
		query["TaskTypes"] = request.TaskTypes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTasks"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tasks created for a cloud phone instance.
//
// Description:
//
//	  You can call the DescribeTasks operation to query the tasks created for one or more cloud phone instances.
//
//		- The system currently supports various tasks, including starting, stopping, restarting, and resetting cloud phone instances; backing up and restoring data; installing apps; and executing remote commands.
//
//		- You can use the Level field to specify the type of task. If Level is set to 1, it represents a batch task. If Level is set to 2, it represents an instance-level task.
//
// **Example**
//
// Assume you restart two cloud phone instances with the instance IDs acp-25nt4kk9whhok\\*\\*\\*\\	- and acp-j2taq887orj8l\\*\\*\\*\\*, and the returned request ID is B8ED2BA9-0C6A-5643-818F-B5D60A64\\*\\*\\*\\*. If you want to check the operation outcomes of the two cloud phone instances, you can call the DescribeTasks operation. You need to set the InvokeId request parameter to B8ED2BA9-0C6A-5643-818F-B5D60A64\\*\\*\\*\\*. If you only want to check the cloud phone instance with the ID acp-25nt4kk9whhok\\*\\*\\*\\*, you must set the ParentTaskId request parameter to the ID of the batch task and the AndroidInstanceId request parameter to acp-25nt4kk9whhok\\*\\*\\*\\	- when calling the DescribeTasks operation.
//
// @param request - DescribeTasksRequest
//
// @return DescribeTasksResponse
func (client *Client) DescribeTasks(request *DescribeTasksRequest) (_result *DescribeTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTasksResponse{}
	_body, _err := client.DescribeTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detaches an Android Debug Bridge (ADB) key pair from one or more cloud phone instances.
//
// Description:
//
//	After you detach an ADB key pair from a cloud phone instance, the ADB connection will fail. This occurs because the system can no longer authenticate using a valid ADB public key, leading to authentication errors.
//
// @param request - DetachKeyPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachKeyPairResponse
func (client *Client) DetachKeyPairWithOptions(request *DetachKeyPairRequest, runtime *dara.RuntimeOptions) (_result *DetachKeyPairResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.KeyPairId) {
		query["KeyPairId"] = request.KeyPairId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachKeyPair"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachKeyPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches an Android Debug Bridge (ADB) key pair from one or more cloud phone instances.
//
// Description:
//
//	After you detach an ADB key pair from a cloud phone instance, the ADB connection will fail. This occurs because the system can no longer authenticate using a valid ADB public key, leading to authentication errors.
//
// @param request - DetachKeyPairRequest
//
// @return DetachKeyPairResponse
func (client *Client) DetachKeyPair(request *DetachKeyPairRequest) (_result *DetachKeyPairResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachKeyPairResponse{}
	_body, _err := client.DetachKeyPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实例断开连接
//
// @param request - DisconnectAndroidInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisconnectAndroidInstanceResponse
func (client *Client) DisconnectAndroidInstanceWithOptions(request *DisconnectAndroidInstanceRequest, runtime *dara.RuntimeOptions) (_result *DisconnectAndroidInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisconnectAndroidInstance"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisconnectAndroidInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实例断开连接
//
// @param request - DisconnectAndroidInstanceRequest
//
// @return DisconnectAndroidInstanceResponse
func (client *Client) DisconnectAndroidInstance(request *DisconnectAndroidInstanceRequest) (_result *DisconnectAndroidInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisconnectAndroidInstanceResponse{}
	_body, _err := client.DisconnectAndroidInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Distributes an image.
//
// Description:
//
// After you distribute an image in supported regions, the distribution cannot be canceled.
//
// @param request - DistributeImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DistributeImageResponse
func (client *Client) DistributeImageWithOptions(request *DistributeImageRequest, runtime *dara.RuntimeOptions) (_result *DistributeImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DistributeRegionList) {
		body["DistributeRegionList"] = request.DistributeRegionList
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DistributeImage"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DistributeImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Distributes an image.
//
// Description:
//
// After you distribute an image in supported regions, the distribution cannot be canceled.
//
// @param request - DistributeImageRequest
//
// @return DistributeImageResponse
func (client *Client) DistributeImage(request *DistributeImageRequest) (_result *DistributeImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DistributeImageResponse{}
	_body, _err := client.DistributeImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Downgrades an instance group. Currently, this operation allows you to only delete specific cloud phone instances from an instance group.
//
// Description:
//
// This operation only allows you to scale down an instance group.
//
// @param request - DowngradeAndroidInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DowngradeAndroidInstanceGroupResponse
func (client *Client) DowngradeAndroidInstanceGroupWithOptions(request *DowngradeAndroidInstanceGroupRequest, runtime *dara.RuntimeOptions) (_result *DowngradeAndroidInstanceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIds) {
		query["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.InstanceGroupId) {
		query["InstanceGroupId"] = request.InstanceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DowngradeAndroidInstanceGroup"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DowngradeAndroidInstanceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Downgrades an instance group. Currently, this operation allows you to only delete specific cloud phone instances from an instance group.
//
// Description:
//
// This operation only allows you to scale down an instance group.
//
// @param request - DowngradeAndroidInstanceGroupRequest
//
// @return DowngradeAndroidInstanceGroupResponse
func (client *Client) DowngradeAndroidInstanceGroup(request *DowngradeAndroidInstanceGroupRequest) (_result *DowngradeAndroidInstanceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DowngradeAndroidInstanceGroupResponse{}
	_body, _err := client.DowngradeAndroidInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 结束协同
//
// @param request - EndCoordinationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EndCoordinationResponse
func (client *Client) EndCoordinationWithOptions(request *EndCoordinationRequest, runtime *dara.RuntimeOptions) (_result *EndCoordinationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CoordinatorUserId) {
		query["CoordinatorUserId"] = request.CoordinatorUserId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerUserId) {
		query["OwnerUserId"] = request.OwnerUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EndCoordination"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EndCoordinationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 结束协同
//
// @param request - EndCoordinationRequest
//
// @return EndCoordinationResponse
func (client *Client) EndCoordination(request *EndCoordinationRequest) (_result *EndCoordinationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EndCoordinationResponse{}
	_body, _err := client.EndCoordinationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 存储扩容
//
// @param request - ExpandDataVolumeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExpandDataVolumeResponse
func (client *Client) ExpandDataVolumeWithOptions(request *ExpandDataVolumeRequest, runtime *dara.RuntimeOptions) (_result *ExpandDataVolumeResponse, _err error) {
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

	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.NodeIds) {
		query["NodeIds"] = request.NodeIds
	}

	if !dara.IsNil(request.PhoneDataVolume) {
		query["PhoneDataVolume"] = request.PhoneDataVolume
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.ShareDataVolume) {
		query["ShareDataVolume"] = request.ShareDataVolume
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExpandDataVolume"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExpandDataVolumeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 存储扩容
//
// @param request - ExpandDataVolumeRequest
//
// @return ExpandDataVolumeResponse
func (client *Client) ExpandDataVolume(request *ExpandDataVolumeRequest) (_result *ExpandDataVolumeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExpandDataVolumeResponse{}
	_body, _err := client.ExpandDataVolumeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 扩容实例的独立机身存储
//
// @param request - ExpandPhoneDataVolumeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExpandPhoneDataVolumeResponse
func (client *Client) ExpandPhoneDataVolumeWithOptions(request *ExpandPhoneDataVolumeRequest, runtime *dara.RuntimeOptions) (_result *ExpandPhoneDataVolumeResponse, _err error) {
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

	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.PhoneDataVolume) {
		query["PhoneDataVolume"] = request.PhoneDataVolume
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExpandPhoneDataVolume"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExpandPhoneDataVolumeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 扩容实例的独立机身存储
//
// @param request - ExpandPhoneDataVolumeRequest
//
// @return ExpandPhoneDataVolumeResponse
func (client *Client) ExpandPhoneDataVolume(request *ExpandPhoneDataVolumeRequest) (_result *ExpandPhoneDataVolumeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExpandPhoneDataVolumeResponse{}
	_body, _err := client.ExpandPhoneDataVolumeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Pulls a file from a cloud phone instance and stores it in Object Storage Service (OSS).
//
// Description:
//
// Currently, this operation allows you to retrieve files or folders from cloud phone instances and save them directly to OSS.
//
// @param request - FetchFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FetchFileResponse
func (client *Client) FetchFileWithOptions(request *FetchFileRequest, runtime *dara.RuntimeOptions) (_result *FetchFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIdList) {
		query["AndroidInstanceIdList"] = request.AndroidInstanceIdList
	}

	if !dara.IsNil(request.SourceFilePath) {
		query["SourceFilePath"] = request.SourceFilePath
	}

	if !dara.IsNil(request.UploadEndpoint) {
		query["UploadEndpoint"] = request.UploadEndpoint
	}

	if !dara.IsNil(request.UploadType) {
		query["UploadType"] = request.UploadType
	}

	if !dara.IsNil(request.UploadUrl) {
		query["UploadUrl"] = request.UploadUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FetchFile"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FetchFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pulls a file from a cloud phone instance and stores it in Object Storage Service (OSS).
//
// Description:
//
// Currently, this operation allows you to retrieve files or folders from cloud phone instances and save them directly to OSS.
//
// @param request - FetchFileRequest
//
// @return FetchFileResponse
func (client *Client) FetchFile(request *FetchFileRequest) (_result *FetchFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FetchFileResponse{}
	_body, _err := client.FetchFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a collaboration code for the cloud phone being accessed by using the current convenience account, and shares this code with other convenience accounts to allow them to access the same cloud phone.
//
// Description:
//
// You can call this operation to generate a collaboration code for a cloud phone accessed by your current account and share this code with other convenience users to allow them to access the same cloud phone over the desktop, mobile, or web client.
//
// @param request - GenerateCoordinationCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateCoordinationCodeResponse
func (client *Client) GenerateCoordinationCodeWithOptions(request *GenerateCoordinationCodeRequest, runtime *dara.RuntimeOptions) (_result *GenerateCoordinationCodeResponse, _err error) {
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

	if !dara.IsNil(request.OwnerUserId) {
		query["OwnerUserId"] = request.OwnerUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateCoordinationCode"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateCoordinationCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a collaboration code for the cloud phone being accessed by using the current convenience account, and shares this code with other convenience accounts to allow them to access the same cloud phone.
//
// Description:
//
// You can call this operation to generate a collaboration code for a cloud phone accessed by your current account and share this code with other convenience users to allow them to access the same cloud phone over the desktop, mobile, or web client.
//
// @param request - GenerateCoordinationCodeRequest
//
// @return GenerateCoordinationCodeResponse
func (client *Client) GenerateCoordinationCode(request *GenerateCoordinationCodeRequest) (_result *GenerateCoordinationCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateCoordinationCodeResponse{}
	_body, _err := client.GenerateCoordinationCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取属性模板信息
//
// @param request - GetInstancePropertiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstancePropertiesResponse
func (client *Client) GetInstancePropertiesWithOptions(request *GetInstancePropertiesRequest, runtime *dara.RuntimeOptions) (_result *GetInstancePropertiesResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceProperties"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstancePropertiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取属性模板信息
//
// @param request - GetInstancePropertiesRequest
//
// @return GetInstancePropertiesResponse
func (client *Client) GetInstanceProperties(request *GetInstancePropertiesRequest) (_result *GetInstancePropertiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstancePropertiesResponse{}
	_body, _err := client.GetInstancePropertiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导入自定义镜像
//
// @param request - ImportImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportImageResponse
func (client *Client) ImportImageWithOptions(request *ImportImageRequest, runtime *dara.RuntimeOptions) (_result *ImportImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageDescription) {
		query["ImageDescription"] = request.ImageDescription
	}

	if !dara.IsNil(request.ImageFileURL) {
		query["ImageFileURL"] = request.ImageFileURL
	}

	if !dara.IsNil(request.ImageName) {
		query["ImageName"] = request.ImageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportImage"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导入自定义镜像
//
// @param request - ImportImageRequest
//
// @return ImportImageResponse
func (client *Client) ImportImage(request *ImportImageRequest) (_result *ImportImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportImageResponse{}
	_body, _err := client.ImportImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports the public key of an Android Debug Bridge (ADB) key pair.
//
// Description:
//
// To avoid authorization errors that could cause ADB connection failures, you must import the public key of an ADB key pair.
//
// @param request - ImportKeyPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportKeyPairResponse
func (client *Client) ImportKeyPairWithOptions(request *ImportKeyPairRequest, runtime *dara.RuntimeOptions) (_result *ImportKeyPairResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyPairName) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !dara.IsNil(request.PublicKeyBody) {
		query["PublicKeyBody"] = request.PublicKeyBody
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportKeyPair"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportKeyPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports the public key of an Android Debug Bridge (ADB) key pair.
//
// Description:
//
// To avoid authorization errors that could cause ADB connection failures, you must import the public key of an ADB key pair.
//
// @param request - ImportKeyPairRequest
//
// @return ImportKeyPairResponse
func (client *Client) ImportKeyPair(request *ImportKeyPairRequest) (_result *ImportKeyPairResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportKeyPairResponse{}
	_body, _err := client.ImportKeyPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Installs an app on multiple cloud phone instances at the same time.
//
// Description:
//
// This operation runs asynchronously. To check the operation result, visit the Task Center. To retrieve task details, call the [DescribeTasks](~~DescribeTasks~~) operation.
//
// @param request - InstallAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallAppResponse
func (client *Client) InstallAppWithOptions(request *InstallAppRequest, runtime *dara.RuntimeOptions) (_result *InstallAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppIdList) {
		query["AppIdList"] = request.AppIdList
	}

	if !dara.IsNil(request.InstanceGroupIdList) {
		query["InstanceGroupIdList"] = request.InstanceGroupIdList
	}

	if !dara.IsNil(request.InstanceIdList) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallApp"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs an app on multiple cloud phone instances at the same time.
//
// Description:
//
// This operation runs asynchronously. To check the operation result, visit the Task Center. To retrieve task details, call the [DescribeTasks](~~DescribeTasks~~) operation.
//
// @param request - InstallAppRequest
//
// @return InstallAppResponse
func (client *Client) InstallApp(request *InstallAppRequest) (_result *InstallAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InstallAppResponse{}
	_body, _err := client.InstallAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 安装监控插件
//
// @param request - InstallMonitorAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallMonitorAgentResponse
func (client *Client) InstallMonitorAgentWithOptions(request *InstallMonitorAgentRequest, runtime *dara.RuntimeOptions) (_result *InstallMonitorAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIds) {
		body["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	if !dara.IsNil(request.SaleMode) {
		body["SaleMode"] = request.SaleMode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallMonitorAgent"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallMonitorAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 安装监控插件
//
// @param request - InstallMonitorAgentRequest
//
// @return InstallMonitorAgentResponse
func (client *Client) InstallMonitorAgent(request *InstallMonitorAgentRequest) (_result *InstallMonitorAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InstallMonitorAgentResponse{}
	_body, _err := client.InstallMonitorAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实例诊断
//
// @param request - InstanceHealerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstanceHealerResponse
func (client *Client) InstanceHealerWithOptions(request *InstanceHealerRequest, runtime *dara.RuntimeOptions) (_result *InstanceHealerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIdList) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	if !dara.IsNil(request.Strategy) {
		query["Strategy"] = request.Strategy
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstanceHealer"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InstanceHealerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实例诊断
//
// @param request - InstanceHealerRequest
//
// @return InstanceHealerResponse
func (client *Client) InstanceHealer(request *InstanceHealerRequest) (_result *InstanceHealerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InstanceHealerResponse{}
	_body, _err := client.InstanceHealerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询ADB端口连接信息
//
// @param request - ListInstanceAdbAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceAdbAttributesResponse
func (client *Client) ListInstanceAdbAttributesWithOptions(request *ListInstanceAdbAttributesRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceAdbAttributesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExternalIp) {
		query["ExternalIp"] = request.ExternalIp
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.InternalIp) {
		query["InternalIp"] = request.InternalIp
	}

	if !dara.IsNil(request.InternalPort) {
		query["InternalPort"] = request.InternalPort
	}

	if !dara.IsNil(request.IpProtocol) {
		query["IpProtocol"] = request.IpProtocol
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceAdbAttributes"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceAdbAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询ADB端口连接信息
//
// @param request - ListInstanceAdbAttributesRequest
//
// @return ListInstanceAdbAttributesResponse
func (client *Client) ListInstanceAdbAttributes(request *ListInstanceAdbAttributesRequest) (_result *ListInstanceAdbAttributesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceAdbAttributesResponse{}
	_body, _err := client.ListInstanceAdbAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries policies.
//
// @param request - ListPolicyGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicyGroupsResponse
func (client *Client) ListPolicyGroupsWithOptions(request *ListPolicyGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListPolicyGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PolicyGroupIds) {
		body["PolicyGroupIds"] = request.PolicyGroupIds
	}

	if !dara.IsNil(request.PolicyGroupName) {
		body["PolicyGroupName"] = request.PolicyGroupName
	}

	if !dara.IsNil(request.PolicyType) {
		body["PolicyType"] = request.PolicyType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicyGroups"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPolicyGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries policies.
//
// @param request - ListPolicyGroupsRequest
//
// @return ListPolicyGroupsResponse
func (client *Client) ListPolicyGroups(request *ListPolicyGroupsRequest) (_result *ListPolicyGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPolicyGroupsResponse{}
	_body, _err := client.ListPolicyGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资源标签
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
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

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
		Version:     dara.String("2023-09-30"),
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
// 查询资源标签
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
// Modifies attributes of a cloud phone instance. Currently, this operation allows you to modify only the name of a cloud phone instance.
//
// @param request - ModifyAndroidInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAndroidInstanceResponse
func (client *Client) ModifyAndroidInstanceWithOptions(request *ModifyAndroidInstanceRequest, runtime *dara.RuntimeOptions) (_result *ModifyAndroidInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceId) {
		query["AndroidInstanceId"] = request.AndroidInstanceId
	}

	if !dara.IsNil(request.DownBandwidthLimit) {
		query["DownBandwidthLimit"] = request.DownBandwidthLimit
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.NewAndroidInstanceName) {
		query["NewAndroidInstanceName"] = request.NewAndroidInstanceName
	}

	if !dara.IsNil(request.UpBandwidthLimit) {
		query["UpBandwidthLimit"] = request.UpBandwidthLimit
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAndroidInstance"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAndroidInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies attributes of a cloud phone instance. Currently, this operation allows you to modify only the name of a cloud phone instance.
//
// @param request - ModifyAndroidInstanceRequest
//
// @return ModifyAndroidInstanceResponse
func (client *Client) ModifyAndroidInstance(request *ModifyAndroidInstanceRequest) (_result *ModifyAndroidInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAndroidInstanceResponse{}
	_body, _err := client.ModifyAndroidInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies attributes of an instance group.
//
// @param request - ModifyAndroidInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAndroidInstanceGroupResponse
func (client *Client) ModifyAndroidInstanceGroupWithOptions(request *ModifyAndroidInstanceGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyAndroidInstanceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceGroupId) {
		query["InstanceGroupId"] = request.InstanceGroupId
	}

	if !dara.IsNil(request.NewInstanceGroupName) {
		query["NewInstanceGroupName"] = request.NewInstanceGroupName
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.StreamMode) {
		query["StreamMode"] = request.StreamMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAndroidInstanceGroup"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAndroidInstanceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies attributes of an instance group.
//
// @param request - ModifyAndroidInstanceGroupRequest
//
// @return ModifyAndroidInstanceGroupResponse
func (client *Client) ModifyAndroidInstanceGroup(request *ModifyAndroidInstanceGroupRequest) (_result *ModifyAndroidInstanceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAndroidInstanceGroupResponse{}
	_body, _err := client.ModifyAndroidInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modify attributes of an application.
//
// @param request - ModifyAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppResponse
func (client *Client) ModifyAppWithOptions(request *ModifyAppRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.IconUrl) {
		query["IconUrl"] = request.IconUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApp"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify attributes of an application.
//
// @param request - ModifyAppRequest
//
// @return ModifyAppResponse
func (client *Client) ModifyApp(request *ModifyAppRequest) (_result *ModifyAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAppResponse{}
	_body, _err := client.ModifyAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a cloud phone matrix. Currently, you can only modify the name of a cloud phone matrix.
//
// @param request - ModifyCloudPhoneNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCloudPhoneNodeResponse
func (client *Client) ModifyCloudPhoneNodeWithOptions(request *ModifyCloudPhoneNodeRequest, runtime *dara.RuntimeOptions) (_result *ModifyCloudPhoneNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewNodeName) {
		query["NewNodeName"] = request.NewNodeName
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.StreamMode) {
		query["StreamMode"] = request.StreamMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCloudPhoneNode"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCloudPhoneNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a cloud phone matrix. Currently, you can only modify the name of a cloud phone matrix.
//
// @param request - ModifyCloudPhoneNodeRequest
//
// @return ModifyCloudPhoneNodeResponse
func (client *Client) ModifyCloudPhoneNode(request *ModifyCloudPhoneNodeRequest) (_result *ModifyCloudPhoneNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCloudPhoneNodeResponse{}
	_body, _err := client.ModifyCloudPhoneNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改显示设置
//
// @param tmpReq - ModifyDisplayConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDisplayConfigResponse
func (client *Client) ModifyDisplayConfigWithOptions(tmpReq *ModifyDisplayConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyDisplayConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyDisplayConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DisplayConfig) {
		request.DisplayConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DisplayConfig, dara.String("DisplayConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIds) {
		body["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	if !dara.IsNil(request.DisplayConfigShrink) {
		body["DisplayConfig"] = request.DisplayConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDisplayConfig"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDisplayConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改显示设置
//
// @param request - ModifyDisplayConfigRequest
//
// @return ModifyDisplayConfigResponse
func (client *Client) ModifyDisplayConfig(request *ModifyDisplayConfigRequest) (_result *ModifyDisplayConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDisplayConfigResponse{}
	_body, _err := client.ModifyDisplayConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the billing method. Currently, this operation only allows you to change the billing method from pay-as-you-go to subscription.
//
// @param request - ModifyInstanceChargeTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceChargeTypeResponse
func (client *Client) ModifyInstanceChargeTypeWithOptions(request *ModifyInstanceChargeTypeRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceChargeTypeResponse, _err error) {
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

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.InstanceGroupIds) {
		query["InstanceGroupIds"] = request.InstanceGroupIds
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceChargeType"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceChargeTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the billing method. Currently, this operation only allows you to change the billing method from pay-as-you-go to subscription.
//
// @param request - ModifyInstanceChargeTypeRequest
//
// @return ModifyInstanceChargeTypeResponse
func (client *Client) ModifyInstanceChargeType(request *ModifyInstanceChargeTypeRequest) (_result *ModifyInstanceChargeTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceChargeTypeResponse{}
	_body, _err := client.ModifyInstanceChargeTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies Android Debug Bridge (ADB) key pairs.
//
// @param request - ModifyKeyPairNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyKeyPairNameResponse
func (client *Client) ModifyKeyPairNameWithOptions(request *ModifyKeyPairNameRequest, runtime *dara.RuntimeOptions) (_result *ModifyKeyPairNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyPairId) {
		query["KeyPairId"] = request.KeyPairId
	}

	if !dara.IsNil(request.NewKeyPairName) {
		query["NewKeyPairName"] = request.NewKeyPairName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyKeyPairName"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyKeyPairNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies Android Debug Bridge (ADB) key pairs.
//
// @param request - ModifyKeyPairNameRequest
//
// @return ModifyKeyPairNameResponse
func (client *Client) ModifyKeyPairName(request *ModifyKeyPairNameRequest) (_result *ModifyKeyPairNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyKeyPairNameResponse{}
	_body, _err := client.ModifyKeyPairNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a policy.
//
// @param tmpReq - ModifyPolicyGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPolicyGroupResponse
func (client *Client) ModifyPolicyGroupWithOptions(tmpReq *ModifyPolicyGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyPolicyGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyPolicyGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NetRedirectPolicy) {
		request.NetRedirectPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NetRedirectPolicy, dara.String("NetRedirectPolicy"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Watermark) {
		request.WatermarkShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Watermark, dara.String("Watermark"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CameraRedirect) {
		body["CameraRedirect"] = request.CameraRedirect
	}

	if !dara.IsNil(request.Clipboard) {
		body["Clipboard"] = request.Clipboard
	}

	if !dara.IsNil(request.Html5FileTransfer) {
		body["Html5FileTransfer"] = request.Html5FileTransfer
	}

	if !dara.IsNil(request.LocalDrive) {
		body["LocalDrive"] = request.LocalDrive
	}

	if !dara.IsNil(request.LockResolution) {
		body["LockResolution"] = request.LockResolution
	}

	if !dara.IsNil(request.NetRedirectPolicyShrink) {
		body["NetRedirectPolicy"] = request.NetRedirectPolicyShrink
	}

	if !dara.IsNil(request.PolicyGroupId) {
		body["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.PolicyGroupName) {
		body["PolicyGroupName"] = request.PolicyGroupName
	}

	if !dara.IsNil(request.ResolutionHeight) {
		body["ResolutionHeight"] = request.ResolutionHeight
	}

	if !dara.IsNil(request.ResolutionWidth) {
		body["ResolutionWidth"] = request.ResolutionWidth
	}

	if !dara.IsNil(request.WatermarkShrink) {
		body["Watermark"] = request.WatermarkShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPolicyGroup"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPolicyGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a policy.
//
// @param request - ModifyPolicyGroupRequest
//
// @return ModifyPolicyGroupResponse
func (client *Client) ModifyPolicyGroup(request *ModifyPolicyGroupRequest) (_result *ModifyPolicyGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyPolicyGroupResponse{}
	_body, _err := client.ModifyPolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改属性模板
//
// @param tmpReq - ModifySystemPropertyTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySystemPropertyTemplateResponse
func (client *Client) ModifySystemPropertyTemplateWithOptions(tmpReq *ModifySystemPropertyTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifySystemPropertyTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifySystemPropertyTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SystemPropertyInfo) {
		request.SystemPropertyInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SystemPropertyInfo, dara.String("SystemPropertyInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableAuto) {
		query["EnableAuto"] = request.EnableAuto
	}

	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.SystemPropertyInfoShrink) {
		query["SystemPropertyInfo"] = request.SystemPropertyInfoShrink
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySystemPropertyTemplate"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySystemPropertyTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改属性模板
//
// @param request - ModifySystemPropertyTemplateRequest
//
// @return ModifySystemPropertyTemplateResponse
func (client *Client) ModifySystemPropertyTemplate(request *ModifySystemPropertyTemplateRequest) (_result *ModifySystemPropertyTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySystemPropertyTemplateResponse{}
	_body, _err := client.ModifySystemPropertyTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Operates apps in a cloud phone, such as opening, closing, and reopening apps.
//
// Description:
//
// This operation runs asynchronously. To check the operation result, visit the Task Center. To retrieve task details, call the [DescribeTasks](~~DescribeTasks~~) operation.
//
// @param request - OperateAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateAppResponse
func (client *Client) OperateAppWithOptions(request *OperateAppRequest, runtime *dara.RuntimeOptions) (_result *OperateAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.InstanceIdList) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	if !dara.IsNil(request.OperateType) {
		query["OperateType"] = request.OperateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateApp"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Operates apps in a cloud phone, such as opening, closing, and reopening apps.
//
// Description:
//
// This operation runs asynchronously. To check the operation result, visit the Task Center. To retrieve task details, call the [DescribeTasks](~~DescribeTasks~~) operation.
//
// @param request - OperateAppRequest
//
// @return OperateAppResponse
func (client *Client) OperateApp(request *OperateAppRequest) (_result *OperateAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateAppResponse{}
	_body, _err := client.OperateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts one or more cloud phone instances.
//
// Description:
//
// Before you restart a cloud phone instance, make sure it is in one of the following states: **Available, Abnormal, Backup failure, and Restoration failure**.
//
// @param request - RebootAndroidInstancesInGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootAndroidInstancesInGroupResponse
func (client *Client) RebootAndroidInstancesInGroupWithOptions(request *RebootAndroidInstancesInGroupRequest, runtime *dara.RuntimeOptions) (_result *RebootAndroidInstancesInGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIds) {
		query["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	if !dara.IsNil(request.ForceStop) {
		query["ForceStop"] = request.ForceStop
	}

	if !dara.IsNil(request.SaleMode) {
		query["SaleMode"] = request.SaleMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebootAndroidInstancesInGroup"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RebootAndroidInstancesInGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts one or more cloud phone instances.
//
// Description:
//
// Before you restart a cloud phone instance, make sure it is in one of the following states: **Available, Abnormal, Backup failure, and Restoration failure**.
//
// @param request - RebootAndroidInstancesInGroupRequest
//
// @return RebootAndroidInstancesInGroupResponse
func (client *Client) RebootAndroidInstancesInGroup(request *RebootAndroidInstancesInGroupRequest) (_result *RebootAndroidInstancesInGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RebootAndroidInstancesInGroupResponse{}
	_body, _err := client.RebootAndroidInstancesInGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 整机恢复
//
// @param request - RecoverAndroidInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecoverAndroidInstanceResponse
func (client *Client) RecoverAndroidInstanceWithOptions(request *RecoverAndroidInstanceRequest, runtime *dara.RuntimeOptions) (_result *RecoverAndroidInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIdList) {
		query["AndroidInstanceIdList"] = request.AndroidInstanceIdList
	}

	if !dara.IsNil(request.BackupFileId) {
		query["BackupFileId"] = request.BackupFileId
	}

	if !dara.IsNil(request.BackupFilePath) {
		query["BackupFilePath"] = request.BackupFilePath
	}

	if !dara.IsNil(request.UploadEndpoint) {
		query["UploadEndpoint"] = request.UploadEndpoint
	}

	if !dara.IsNil(request.UploadType) {
		query["UploadType"] = request.UploadType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecoverAndroidInstance"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecoverAndroidInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 整机恢复
//
// @param request - RecoverAndroidInstanceRequest
//
// @return RecoverAndroidInstanceResponse
func (client *Client) RecoverAndroidInstance(request *RecoverAndroidInstanceRequest) (_result *RecoverAndroidInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RecoverAndroidInstanceResponse{}
	_body, _err := client.RecoverAndroidInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 恢复应用
//
// @param request - RecoverAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecoverAppResponse
func (client *Client) RecoverAppWithOptions(request *RecoverAppRequest, runtime *dara.RuntimeOptions) (_result *RecoverAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIdList) {
		query["AndroidInstanceIdList"] = request.AndroidInstanceIdList
	}

	if !dara.IsNil(request.BackupFileId) {
		query["BackupFileId"] = request.BackupFileId
	}

	if !dara.IsNil(request.BackupFilePath) {
		query["BackupFilePath"] = request.BackupFilePath
	}

	if !dara.IsNil(request.UploadEndpoint) {
		query["UploadEndpoint"] = request.UploadEndpoint
	}

	if !dara.IsNil(request.UploadType) {
		query["UploadType"] = request.UploadType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecoverApp"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecoverAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 恢复应用
//
// @param request - RecoverAppRequest
//
// @return RecoverAppResponse
func (client *Client) RecoverApp(request *RecoverAppRequest) (_result *RecoverAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RecoverAppResponse{}
	_body, _err := client.RecoverAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restores backup files.
//
// Description:
//
// Currently, this operation allows you to restore only backup files generated by cloud phones that are stored in Object Storage Service (OSS) buckets.
//
// @param request - RecoveryFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecoveryFileResponse
func (client *Client) RecoveryFileWithOptions(request *RecoveryFileRequest, runtime *dara.RuntimeOptions) (_result *RecoveryFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIdList) {
		query["AndroidInstanceIdList"] = request.AndroidInstanceIdList
	}

	if !dara.IsNil(request.BackupAll) {
		query["BackupAll"] = request.BackupAll
	}

	if !dara.IsNil(request.BackupFileId) {
		query["BackupFileId"] = request.BackupFileId
	}

	if !dara.IsNil(request.BackupFilePath) {
		query["BackupFilePath"] = request.BackupFilePath
	}

	if !dara.IsNil(request.UploadEndpoint) {
		query["UploadEndpoint"] = request.UploadEndpoint
	}

	if !dara.IsNil(request.UploadType) {
		query["UploadType"] = request.UploadType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecoveryFile"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecoveryFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores backup files.
//
// Description:
//
// Currently, this operation allows you to restore only backup files generated by cloud phones that are stored in Object Storage Service (OSS) buckets.
//
// @param request - RecoveryFileRequest
//
// @return RecoveryFileResponse
func (client *Client) RecoveryFile(request *RecoveryFileRequest) (_result *RecoveryFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RecoveryFileResponse{}
	_body, _err := client.RecoveryFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renews instance groups.
//
// @param request - RenewAndroidInstanceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewAndroidInstanceGroupsResponse
func (client *Client) RenewAndroidInstanceGroupsWithOptions(request *RenewAndroidInstanceGroupsRequest, runtime *dara.RuntimeOptions) (_result *RenewAndroidInstanceGroupsResponse, _err error) {
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

	if !dara.IsNil(request.InstanceGroupIds) {
		query["InstanceGroupIds"] = request.InstanceGroupIds
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewAndroidInstanceGroups"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewAndroidInstanceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews instance groups.
//
// @param request - RenewAndroidInstanceGroupsRequest
//
// @return RenewAndroidInstanceGroupsResponse
func (client *Client) RenewAndroidInstanceGroups(request *RenewAndroidInstanceGroupsRequest) (_result *RenewAndroidInstanceGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewAndroidInstanceGroupsResponse{}
	_body, _err := client.RenewAndroidInstanceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renews a cloud mobile matrix.
//
// @param request - RenewCloudPhoneNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewCloudPhoneNodesResponse
func (client *Client) RenewCloudPhoneNodesWithOptions(request *RenewCloudPhoneNodesRequest, runtime *dara.RuntimeOptions) (_result *RenewCloudPhoneNodesResponse, _err error) {
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

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenew) {
		body["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.NodeIds) {
		body["NodeIds"] = request.NodeIds
	}

	if !dara.IsNil(request.Period) {
		body["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		body["PeriodUnit"] = request.PeriodUnit
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewCloudPhoneNodes"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewCloudPhoneNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews a cloud mobile matrix.
//
// @param request - RenewCloudPhoneNodesRequest
//
// @return RenewCloudPhoneNodesResponse
func (client *Client) RenewCloudPhoneNodes(request *RenewCloudPhoneNodesRequest) (_result *RenewCloudPhoneNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewCloudPhoneNodesResponse{}
	_body, _err := client.RenewCloudPhoneNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets one or more cloud phone instances.
//
// Description:
//
// Before you reset a cloud phone instance, make sure it is in one of the following states: **Available, Stopped, Abnormal, Backup failure, and Restoration failure**.
//
// @param request - ResetAndroidInstancesInGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAndroidInstancesInGroupResponse
func (client *Client) ResetAndroidInstancesInGroupWithOptions(request *ResetAndroidInstancesInGroupRequest, runtime *dara.RuntimeOptions) (_result *ResetAndroidInstancesInGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIds) {
		query["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	if !dara.IsNil(request.SaleMode) {
		query["SaleMode"] = request.SaleMode
	}

	if !dara.IsNil(request.SettingResetType) {
		query["SettingResetType"] = request.SettingResetType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetAndroidInstancesInGroup"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetAndroidInstancesInGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets one or more cloud phone instances.
//
// Description:
//
// Before you reset a cloud phone instance, make sure it is in one of the following states: **Available, Stopped, Abnormal, Backup failure, and Restoration failure**.
//
// @param request - ResetAndroidInstancesInGroupRequest
//
// @return ResetAndroidInstancesInGroupResponse
func (client *Client) ResetAndroidInstancesInGroup(request *ResetAndroidInstancesInGroupRequest) (_result *ResetAndroidInstancesInGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetAndroidInstancesInGroupResponse{}
	_body, _err := client.ResetAndroidInstancesInGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Executes a command on a cloud phone instance.
//
// @param request - RunCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCommandResponse
func (client *Client) RunCommandWithOptions(request *RunCommandRequest, runtime *dara.RuntimeOptions) (_result *RunCommandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentType) {
		query["AgentType"] = request.AgentType
	}

	if !dara.IsNil(request.CommandContent) {
		query["CommandContent"] = request.CommandContent
	}

	if !dara.IsNil(request.ContentEncoding) {
		query["ContentEncoding"] = request.ContentEncoding
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunCommand"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes a command on a cloud phone instance.
//
// @param request - RunCommandRequest
//
// @return RunCommandResponse
func (client *Client) RunCommand(request *RunCommandRequest) (_result *RunCommandResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunCommandResponse{}
	_body, _err := client.RunCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Pushes files from Object Storage Service (OSS) buckets to cloud phone instances.
//
// Description:
//
// Currently, this operation allows you to only push files or folders from OSS buckets to cloud phone instances.
//
// @param request - SendFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendFileResponse
func (client *Client) SendFileWithOptions(request *SendFileRequest, runtime *dara.RuntimeOptions) (_result *SendFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIdList) {
		query["AndroidInstanceIdList"] = request.AndroidInstanceIdList
	}

	if !dara.IsNil(request.AutoInstall) {
		query["AutoInstall"] = request.AutoInstall
	}

	if !dara.IsNil(request.SourceFilePath) {
		query["SourceFilePath"] = request.SourceFilePath
	}

	if !dara.IsNil(request.TargetFileName) {
		query["TargetFileName"] = request.TargetFileName
	}

	if !dara.IsNil(request.UploadEndpoint) {
		query["UploadEndpoint"] = request.UploadEndpoint
	}

	if !dara.IsNil(request.UploadType) {
		query["UploadType"] = request.UploadType
	}

	if !dara.IsNil(request.UploadUrl) {
		query["UploadUrl"] = request.UploadUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendFile"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pushes files from Object Storage Service (OSS) buckets to cloud phone instances.
//
// Description:
//
// Currently, this operation allows you to only push files or folders from OSS buckets to cloud phone instances.
//
// @param request - SendFileRequest
//
// @return SendFileResponse
func (client *Client) SendFile(request *SendFileRequest) (_result *SendFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendFileResponse{}
	_body, _err := client.SendFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发送属性模板
//
// @param request - SendSystemPropertyTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendSystemPropertyTemplateResponse
func (client *Client) SendSystemPropertyTemplateWithOptions(request *SendSystemPropertyTemplateRequest, runtime *dara.RuntimeOptions) (_result *SendSystemPropertyTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateIds) {
		query["TemplateIds"] = request.TemplateIds
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIds) {
		body["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendSystemPropertyTemplate"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendSystemPropertyTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送属性模板
//
// @param request - SendSystemPropertyTemplateRequest
//
// @return SendSystemPropertyTemplateResponse
func (client *Client) SendSystemPropertyTemplate(request *SendSystemPropertyTemplateRequest) (_result *SendSystemPropertyTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendSystemPropertyTemplateResponse{}
	_body, _err := client.SendSystemPropertyTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the authentication status for cloud phone instances. If you enable Android Debug Bridge (ADB) authentication for cloud phone instances, the system will verify the validity of the ADB key pairs provided by end users when they connect to the instances over ADB. To ensure successful authentication and a proper connection, we recommend that you attach ADB key pairs to cloud phone instances. If you disable ADB authentication for cloud phone instances, the system will no longer verify the validity of any ADB key pairs. As a result, end users can connect to the cloud phone instances over ADB without authentication, provided the network connection is functioning properly.
//
// Description:
//
// Before you call this operation, make sure that the desired cloud phone instance is in the Running state.
//
// @param request - SetAdbSecureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAdbSecureResponse
func (client *Client) SetAdbSecureWithOptions(request *SetAdbSecureRequest, runtime *dara.RuntimeOptions) (_result *SetAdbSecureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetAdbSecure"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetAdbSecureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the authentication status for cloud phone instances. If you enable Android Debug Bridge (ADB) authentication for cloud phone instances, the system will verify the validity of the ADB key pairs provided by end users when they connect to the instances over ADB. To ensure successful authentication and a proper connection, we recommend that you attach ADB key pairs to cloud phone instances. If you disable ADB authentication for cloud phone instances, the system will no longer verify the validity of any ADB key pairs. As a result, end users can connect to the cloud phone instances over ADB without authentication, provided the network connection is functioning properly.
//
// Description:
//
// Before you call this operation, make sure that the desired cloud phone instance is in the Running state.
//
// @param request - SetAdbSecureRequest
//
// @return SetAdbSecureResponse
func (client *Client) SetAdbSecure(request *SetAdbSecureRequest) (_result *SetAdbSecureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetAdbSecureResponse{}
	_body, _err := client.SetAdbSecureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Start instances.
//
// Description:
//
// Only supports starting when the instance is in the **Stopped, Backup Failed, or Recovery Failed*	- state.
//
// @param request - StartAndroidInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAndroidInstanceResponse
func (client *Client) StartAndroidInstanceWithOptions(request *StartAndroidInstanceRequest, runtime *dara.RuntimeOptions) (_result *StartAndroidInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIds) {
		query["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	if !dara.IsNil(request.SaleMode) {
		query["SaleMode"] = request.SaleMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartAndroidInstance"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartAndroidInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Start instances.
//
// Description:
//
// Only supports starting when the instance is in the **Stopped, Backup Failed, or Recovery Failed*	- state.
//
// @param request - StartAndroidInstanceRequest
//
// @return StartAndroidInstanceResponse
func (client *Client) StartAndroidInstance(request *StartAndroidInstanceRequest) (_result *StartAndroidInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartAndroidInstanceResponse{}
	_body, _err := client.StartAndroidInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开启实例ADB端口并创建端口转发条目
//
// @param request - StartInstanceAdbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartInstanceAdbResponse
func (client *Client) StartInstanceAdbWithOptions(request *StartInstanceAdbRequest, runtime *dara.RuntimeOptions) (_result *StartInstanceAdbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartInstanceAdb"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartInstanceAdbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启实例ADB端口并创建端口转发条目
//
// @param request - StartInstanceAdbRequest
//
// @return StartInstanceAdbResponse
func (client *Client) StartInstanceAdb(request *StartInstanceAdbRequest) (_result *StartInstanceAdbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartInstanceAdbResponse{}
	_body, _err := client.StartInstanceAdbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a cloud phone instance.
//
// Description:
//
// Before you stop a cloud phone instance, make sure it is in one of the following states: **Available, Backup failure, and Restoration failure**.
//
// @param request - StopAndroidInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAndroidInstanceResponse
func (client *Client) StopAndroidInstanceWithOptions(request *StopAndroidInstanceRequest, runtime *dara.RuntimeOptions) (_result *StopAndroidInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIds) {
		query["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	if !dara.IsNil(request.ForceStop) {
		query["ForceStop"] = request.ForceStop
	}

	if !dara.IsNil(request.SaleMode) {
		query["SaleMode"] = request.SaleMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopAndroidInstance"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopAndroidInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a cloud phone instance.
//
// Description:
//
// Before you stop a cloud phone instance, make sure it is in one of the following states: **Available, Backup failure, and Restoration failure**.
//
// @param request - StopAndroidInstanceRequest
//
// @return StopAndroidInstanceResponse
func (client *Client) StopAndroidInstance(request *StopAndroidInstanceRequest) (_result *StopAndroidInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopAndroidInstanceResponse{}
	_body, _err := client.StopAndroidInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止实例ADB端口并删除端口转发条目
//
// @param request - StopInstanceAdbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopInstanceAdbResponse
func (client *Client) StopInstanceAdbWithOptions(request *StopInstanceAdbRequest, runtime *dara.RuntimeOptions) (_result *StopInstanceAdbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopInstanceAdb"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopInstanceAdbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止实例ADB端口并删除端口转发条目
//
// @param request - StopInstanceAdbRequest
//
// @return StopInstanceAdbResponse
func (client *Client) StopInstanceAdb(request *StopInstanceAdbRequest) (_result *StopInstanceAdbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopInstanceAdbResponse{}
	_body, _err := client.StopInstanceAdbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 给资源打标签
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
		Version:     dara.String("2023-09-30"),
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
// 给资源打标签
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
// Uninstalls an app from multiple cloud phone instances.
//
// Description:
//
// This operation runs asynchronously. To check the operation result, you can visit the Task Center. To retrieve task details, call the [DescribeTasks](~~DescribeTasks~~) operation.
//
// @param request - UninstallAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallAppResponse
func (client *Client) UninstallAppWithOptions(request *UninstallAppRequest, runtime *dara.RuntimeOptions) (_result *UninstallAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppIdList) {
		query["AppIdList"] = request.AppIdList
	}

	if !dara.IsNil(request.InstanceGroupIdList) {
		query["InstanceGroupIdList"] = request.InstanceGroupIdList
	}

	if !dara.IsNil(request.InstanceIdList) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UninstallApp"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UninstallAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uninstalls an app from multiple cloud phone instances.
//
// Description:
//
// This operation runs asynchronously. To check the operation result, you can visit the Task Center. To retrieve task details, call the [DescribeTasks](~~DescribeTasks~~) operation.
//
// @param request - UninstallAppRequest
//
// @return UninstallAppResponse
func (client *Client) UninstallApp(request *UninstallAppRequest) (_result *UninstallAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UninstallAppResponse{}
	_body, _err := client.UninstallAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 卸载监控插件
//
// @param request - UninstallMonitorAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallMonitorAgentResponse
func (client *Client) UninstallMonitorAgentWithOptions(request *UninstallMonitorAgentRequest, runtime *dara.RuntimeOptions) (_result *UninstallMonitorAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AndroidInstanceIds) {
		body["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	if !dara.IsNil(request.SaleMode) {
		body["SaleMode"] = request.SaleMode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UninstallMonitorAgent"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UninstallMonitorAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 卸载监控插件
//
// @param request - UninstallMonitorAgentRequest
//
// @return UninstallMonitorAgentResponse
func (client *Client) UninstallMonitorAgent(request *UninstallMonitorAgentRequest) (_result *UninstallMonitorAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UninstallMonitorAgentResponse{}
	_body, _err := client.UninstallMonitorAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除资源标签
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
		Version:     dara.String("2023-09-30"),
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
// 删除资源标签
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
// Updates the name of a custom image.
//
// @param request - UpdateCustomImageNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomImageNameResponse
func (client *Client) UpdateCustomImageNameWithOptions(request *UpdateCustomImageNameRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomImageNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageName) {
		body["ImageName"] = request.ImageName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomImageName"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomImageNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the name of a custom image.
//
// @param request - UpdateCustomImageNameRequest
//
// @return UpdateCustomImageNameResponse
func (client *Client) UpdateCustomImageName(request *UpdateCustomImageNameRequest) (_result *UpdateCustomImageNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCustomImageNameResponse{}
	_body, _err := client.UpdateCustomImageNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the image of an instance group.
//
// Description:
//
// Before you call this operation, make sure the image is in the Available state and the region of the image is included in the region list of the desired instance group. In addition, the instance group itself is available.
//
// @param request - UpdateInstanceGroupImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceGroupImageResponse
func (client *Client) UpdateInstanceGroupImageWithOptions(request *UpdateInstanceGroupImageRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceGroupImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InstanceGroupIds) {
		body["InstanceGroupIds"] = request.InstanceGroupIds
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceGroupImage"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceGroupImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the image of an instance group.
//
// Description:
//
// Before you call this operation, make sure the image is in the Available state and the region of the image is included in the region list of the desired instance group. In addition, the instance group itself is available.
//
// @param request - UpdateInstanceGroupImageRequest
//
// @return UpdateInstanceGroupImageResponse
func (client *Client) UpdateInstanceGroupImage(request *UpdateInstanceGroupImageRequest) (_result *UpdateInstanceGroupImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateInstanceGroupImageResponse{}
	_body, _err := client.UpdateInstanceGroupImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新实例镜像
//
// @param request - UpdateInstanceImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceImageResponse
func (client *Client) UpdateInstanceImageWithOptions(request *UpdateInstanceImageRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InstanceIdList) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	if !dara.IsNil(request.Reset) {
		query["Reset"] = request.Reset
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceImage"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新实例镜像
//
// @param request - UpdateInstanceImageRequest
//
// @return UpdateInstanceImageResponse
func (client *Client) UpdateInstanceImage(request *UpdateInstanceImageRequest) (_result *UpdateInstanceImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateInstanceImageResponse{}
	_body, _err := client.UpdateInstanceImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades an instance group. Currently, this operation allows you to only increase the number of instances in an instance group.
//
// Description:
//
// Currently, this operation allows you to only increase the size of an instance group.
//
// @param request - UpgradeAndroidInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeAndroidInstanceGroupResponse
func (client *Client) UpgradeAndroidInstanceGroupWithOptions(request *UpgradeAndroidInstanceGroupRequest, runtime *dara.RuntimeOptions) (_result *UpgradeAndroidInstanceGroupResponse, _err error) {
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

	if !dara.IsNil(request.IncreaseNumberOfInstance) {
		query["IncreaseNumberOfInstance"] = request.IncreaseNumberOfInstance
	}

	if !dara.IsNil(request.InstanceGroupId) {
		query["InstanceGroupId"] = request.InstanceGroupId
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeAndroidInstanceGroup"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeAndroidInstanceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades an instance group. Currently, this operation allows you to only increase the number of instances in an instance group.
//
// Description:
//
// Currently, this operation allows you to only increase the size of an instance group.
//
// @param request - UpgradeAndroidInstanceGroupRequest
//
// @return UpgradeAndroidInstanceGroupResponse
func (client *Client) UpgradeAndroidInstanceGroup(request *UpgradeAndroidInstanceGroupRequest) (_result *UpgradeAndroidInstanceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeAndroidInstanceGroupResponse{}
	_body, _err := client.UpgradeAndroidInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
