// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ActivateFlowLogRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The ID of the flow log.
	//
	// This parameter is required.
	//
	// example:
	//
	// flowlog-m5evbtbpt****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
}

func (s ActivateFlowLogRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivateFlowLogRequest) GoString() string {
	return s.String()
}

func (s *ActivateFlowLogRequest) SetClientToken(v string) *ActivateFlowLogRequest {
	s.ClientToken = &v
	return s
}

func (s *ActivateFlowLogRequest) SetDryRun(v bool) *ActivateFlowLogRequest {
	s.DryRun = &v
	return s
}

func (s *ActivateFlowLogRequest) SetEcrId(v string) *ActivateFlowLogRequest {
	s.EcrId = &v
	return s
}

func (s *ActivateFlowLogRequest) SetFlowLogId(v string) *ActivateFlowLogRequest {
	s.FlowLogId = &v
	return s
}

type ActivateFlowLogResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 05130E79-588D-5C40-A718-C4863A59****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ActivateFlowLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActivateFlowLogResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateFlowLogResponseBody) SetAccessDeniedDetail(v string) *ActivateFlowLogResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ActivateFlowLogResponseBody) SetCode(v string) *ActivateFlowLogResponseBody {
	s.Code = &v
	return s
}

func (s *ActivateFlowLogResponseBody) SetDynamicCode(v string) *ActivateFlowLogResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ActivateFlowLogResponseBody) SetDynamicMessage(v string) *ActivateFlowLogResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ActivateFlowLogResponseBody) SetHttpStatusCode(v int32) *ActivateFlowLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ActivateFlowLogResponseBody) SetMessage(v string) *ActivateFlowLogResponseBody {
	s.Message = &v
	return s
}

func (s *ActivateFlowLogResponseBody) SetRequestId(v string) *ActivateFlowLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateFlowLogResponseBody) SetSuccess(v bool) *ActivateFlowLogResponseBody {
	s.Success = &v
	return s
}

type ActivateFlowLogResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateFlowLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateFlowLogResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivateFlowLogResponse) GoString() string {
	return s.String()
}

func (s *ActivateFlowLogResponse) SetHeaders(v map[string]*string) *ActivateFlowLogResponse {
	s.Headers = v
	return s
}

func (s *ActivateFlowLogResponse) SetStatusCode(v int32) *ActivateFlowLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateFlowLogResponse) SetBody(v *ActivateFlowLogResponseBody) *ActivateFlowLogResponse {
	s.Body = v
	return s
}

type AttachExpressConnectRouterChildInstanceRequest struct {
	// The VBR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-j6cwxhgg0s5nuephp****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the VBR belongs.
	//
	// >  If you want to connect to a network instance that belongs to a different account, this parameter is required.
	//
	// example:
	//
	// 190550144868****
	ChildInstanceOwnerId *int64 `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	// The region ID of the VBR.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	// The type of the network instance. Set the value to **VBR**.
	//
	// This parameter is required.
	//
	// example:
	//
	// VBR
	ChildInstanceType *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-a5xqrgbeidz1w8****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
}

func (s AttachExpressConnectRouterChildInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachExpressConnectRouterChildInstanceRequest) GoString() string {
	return s.String()
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetChildInstanceId(v string) *AttachExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetChildInstanceOwnerId(v int64) *AttachExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetChildInstanceRegionId(v string) *AttachExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetChildInstanceType(v string) *AttachExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetClientToken(v string) *AttachExpressConnectRouterChildInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetDescription(v string) *AttachExpressConnectRouterChildInstanceRequest {
	s.Description = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetDryRun(v bool) *AttachExpressConnectRouterChildInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetEcrId(v string) *AttachExpressConnectRouterChildInstanceRequest {
	s.EcrId = &v
	return s
}

type AttachExpressConnectRouterChildInstanceResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachExpressConnectRouterChildInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachExpressConnectRouterChildInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetAccessDeniedDetail(v string) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetCode(v string) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetDynamicCode(v string) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetDynamicMessage(v string) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetHttpStatusCode(v int32) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetMessage(v string) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetRequestId(v string) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetSuccess(v bool) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.Success = &v
	return s
}

type AttachExpressConnectRouterChildInstanceResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachExpressConnectRouterChildInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachExpressConnectRouterChildInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachExpressConnectRouterChildInstanceResponse) GoString() string {
	return s.String()
}

func (s *AttachExpressConnectRouterChildInstanceResponse) SetHeaders(v map[string]*string) *AttachExpressConnectRouterChildInstanceResponse {
	s.Headers = v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponse) SetStatusCode(v int32) *AttachExpressConnectRouterChildInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponse) SetBody(v *AttachExpressConnectRouterChildInstanceResponseBody) *AttachExpressConnectRouterChildInstanceResponse {
	s.Body = v
	return s
}

type CheckAddRegionToExpressConnectRouterRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- true: performs only a dry run.
	//
	// 	- false (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-fu8rszhgv7623c****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The ID of the region for which you want to check whether the CDT service is enabled for the ECR feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	FreshRegionId *string `json:"FreshRegionId,omitempty" xml:"FreshRegionId,omitempty"`
}

func (s CheckAddRegionToExpressConnectRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckAddRegionToExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *CheckAddRegionToExpressConnectRouterRequest) SetClientToken(v string) *CheckAddRegionToExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterRequest) SetDryRun(v bool) *CheckAddRegionToExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterRequest) SetEcrId(v string) *CheckAddRegionToExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterRequest) SetFreshRegionId(v string) *CheckAddRegionToExpressConnectRouterRequest {
	s.FreshRegionId = &v
	return s
}

type CheckAddRegionToExpressConnectRouterResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether the ECR is used to establish connections between regions in the Chinese mainland and regions outside China. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AnyCrossBorderLink *bool `json:"AnyCrossBorderLink,omitempty" xml:"AnyCrossBorderLink,omitempty"`
	// Indicates whether the ECR is used to establish connections between regions in the Chinese mainland. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AnyInterRegionLink *bool `json:"AnyInterRegionLink,omitempty" xml:"AnyInterRegionLink,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsInstanceId**, the request parameter **DtsInstanceId*	- is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Indicates whether the cross-border CDT service is activated for the account to which the ECR belongs. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsCdtCrossBorderEnabled *bool `json:"IsCdtCrossBorderEnabled,omitempty" xml:"IsCdtCrossBorderEnabled,omitempty"`
	// Indicates whether the CDT service is activated for the account to which the ECR belongs. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsCdtInterRegionEnabled *bool `json:"IsCdtInterRegionEnabled,omitempty" xml:"IsCdtInterRegionEnabled,omitempty"`
	// Indicates whether the account to which the ECR belongs can create cross-border connections. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsUserAllowedToCreateCrossBorderLink *bool `json:"IsUserAllowedToCreateCrossBorderLink,omitempty" xml:"IsUserAllowedToCreateCrossBorderLink,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckAddRegionToExpressConnectRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckAddRegionToExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetAnyCrossBorderLink(v bool) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.AnyCrossBorderLink = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetAnyInterRegionLink(v bool) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.AnyInterRegionLink = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetCode(v string) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetDynamicCode(v string) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetDynamicMessage(v string) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetIsCdtCrossBorderEnabled(v bool) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.IsCdtCrossBorderEnabled = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetIsCdtInterRegionEnabled(v bool) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.IsCdtInterRegionEnabled = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetIsUserAllowedToCreateCrossBorderLink(v bool) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.IsUserAllowedToCreateCrossBorderLink = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetMessage(v string) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetRequestId(v string) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetSuccess(v bool) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

type CheckAddRegionToExpressConnectRouterResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAddRegionToExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAddRegionToExpressConnectRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckAddRegionToExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *CheckAddRegionToExpressConnectRouterResponse) SetHeaders(v map[string]*string) *CheckAddRegionToExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponse) SetStatusCode(v int32) *CheckAddRegionToExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponse) SetBody(v *CheckAddRegionToExpressConnectRouterResponseBody) *CheckAddRegionToExpressConnectRouterResponse {
	s.Body = v
	return s
}

type CreateExpressConnectRouterRequest struct {
	// The autonomous system number (ASN) of the ECR. Valid values: 45104, 64512 to 65534, and 4200000000 to 4294967294. Default value: 45104. The value 65025 is reserved by Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// 45104
	AlibabaSideAsn *int64 `json:"AlibabaSideAsn,omitempty" xml:"AlibabaSideAsn,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the ECR.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The name of the ECR.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group to which the ECR belongs.
	//
	// example:
	//
	// rg-acfmvvajg5q****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The information about the tags.
	//
	// You can specify at most 20 tags in each call.
	Tag []*CreateExpressConnectRouterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateExpressConnectRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectRouterRequest) SetAlibabaSideAsn(v int64) *CreateExpressConnectRouterRequest {
	s.AlibabaSideAsn = &v
	return s
}

func (s *CreateExpressConnectRouterRequest) SetClientToken(v string) *CreateExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateExpressConnectRouterRequest) SetDescription(v string) *CreateExpressConnectRouterRequest {
	s.Description = &v
	return s
}

func (s *CreateExpressConnectRouterRequest) SetDryRun(v bool) *CreateExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *CreateExpressConnectRouterRequest) SetName(v string) *CreateExpressConnectRouterRequest {
	s.Name = &v
	return s
}

func (s *CreateExpressConnectRouterRequest) SetResourceGroupId(v string) *CreateExpressConnectRouterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateExpressConnectRouterRequest) SetTag(v []*CreateExpressConnectRouterRequestTag) *CreateExpressConnectRouterRequest {
	s.Tag = v
	return s
}

type CreateExpressConnectRouterRequestTag struct {
	// The tag key. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateExpressConnectRouterRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateExpressConnectRouterRequestTag) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectRouterRequestTag) SetKey(v string) *CreateExpressConnectRouterRequestTag {
	s.Key = &v
	return s
}

func (s *CreateExpressConnectRouterRequestTag) SetValue(v string) *CreateExpressConnectRouterRequestTag {
	s.Value = &v
	return s
}

type CreateExpressConnectRouterResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.Name
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of Name ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The ECR ID.
	//
	// example:
	//
	// ecr-fu8rszhgv7623c****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the ECR is created. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateExpressConnectRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *CreateExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetCode(v string) *CreateExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetDynamicCode(v string) *CreateExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetDynamicMessage(v string) *CreateExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetEcrId(v string) *CreateExpressConnectRouterResponseBody {
	s.EcrId = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *CreateExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetMessage(v string) *CreateExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetRequestId(v string) *CreateExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetSuccess(v bool) *CreateExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

type CreateExpressConnectRouterResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExpressConnectRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectRouterResponse) SetHeaders(v map[string]*string) *CreateExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *CreateExpressConnectRouterResponse) SetStatusCode(v int32) *CreateExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExpressConnectRouterResponse) SetBody(v *CreateExpressConnectRouterResponseBody) *CreateExpressConnectRouterResponse {
	s.Body = v
	return s
}

type CreateExpressConnectRouterAssociationRequest struct {
	// The allowed route prefixes.
	AllowedPrefixes     []*string `json:"AllowedPrefixes,omitempty" xml:"AllowedPrefixes,omitempty" type:"Repeated"`
	AllowedPrefixesMode *string   `json:"AllowedPrefixesMode,omitempty" xml:"AllowedPrefixesMode,omitempty"`
	// The region ID of the resource to be associated.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	AssociationRegionId *string `json:"AssociationRegionId,omitempty" xml:"AssociationRegionId,omitempty"`
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// example:
	//
	// cen-of3o1the3i4gwb****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to initiate an association on the TR when the ECR is associated with the TR. Valid values:
	//
	// 	- **true**: You do not need to initiate an association on the TR.
	//
	// 	- **false**: You need to initiate an association on the TR.
	//
	// example:
	//
	// true
	CreateAttachment *bool   `json:"CreateAttachment,omitempty" xml:"CreateAttachment,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The TR ID.
	//
	// example:
	//
	// tr-2ze4i71c6be454e2l****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The ID of the Alibaba Cloud account that owns the TR. Default value: ID of the Alibaba Cloud account that logs in.
	//
	// >  If you want to connect to a network instance that belongs to a different account, this parameter is required.
	//
	// example:
	//
	// 189159362009****
	TransitRouterOwnerId *int64 `json:"TransitRouterOwnerId,omitempty" xml:"TransitRouterOwnerId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp1h37fchc6jmfyln****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the Alibaba Cloud account that owns the VPC. Default value: ID of the Alibaba Cloud account that logs in.
	//
	// >  If you want to connect to a network instance that belongs to a different account, this parameter is required.
	//
	// example:
	//
	// 132193271328****
	VpcOwnerId *int64 `json:"VpcOwnerId,omitempty" xml:"VpcOwnerId,omitempty"`
}

func (s CreateExpressConnectRouterAssociationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExpressConnectRouterAssociationRequest) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectRouterAssociationRequest) SetAllowedPrefixes(v []*string) *CreateExpressConnectRouterAssociationRequest {
	s.AllowedPrefixes = v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetAllowedPrefixesMode(v string) *CreateExpressConnectRouterAssociationRequest {
	s.AllowedPrefixesMode = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetAssociationRegionId(v string) *CreateExpressConnectRouterAssociationRequest {
	s.AssociationRegionId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetCenId(v string) *CreateExpressConnectRouterAssociationRequest {
	s.CenId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetClientToken(v string) *CreateExpressConnectRouterAssociationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetCreateAttachment(v bool) *CreateExpressConnectRouterAssociationRequest {
	s.CreateAttachment = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetDescription(v string) *CreateExpressConnectRouterAssociationRequest {
	s.Description = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetDryRun(v bool) *CreateExpressConnectRouterAssociationRequest {
	s.DryRun = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetEcrId(v string) *CreateExpressConnectRouterAssociationRequest {
	s.EcrId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetTransitRouterId(v string) *CreateExpressConnectRouterAssociationRequest {
	s.TransitRouterId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetTransitRouterOwnerId(v int64) *CreateExpressConnectRouterAssociationRequest {
	s.TransitRouterOwnerId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetVpcId(v string) *CreateExpressConnectRouterAssociationRequest {
	s.VpcId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetVpcOwnerId(v int64) *CreateExpressConnectRouterAssociationRequest {
	s.VpcOwnerId = &v
	return s
}

type CreateExpressConnectRouterAssociationResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The ID of the association between the ECR and the VPC or TR.
	//
	// example:
	//
	// ecr-assoc-9p2qxx5phpca2m****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 05130E79-588D-5C40-A718-C4863A59****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateExpressConnectRouterAssociationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateExpressConnectRouterAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetAccessDeniedDetail(v string) *CreateExpressConnectRouterAssociationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetAssociationId(v string) *CreateExpressConnectRouterAssociationResponseBody {
	s.AssociationId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetCode(v string) *CreateExpressConnectRouterAssociationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetDynamicCode(v string) *CreateExpressConnectRouterAssociationResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetDynamicMessage(v string) *CreateExpressConnectRouterAssociationResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetHttpStatusCode(v int32) *CreateExpressConnectRouterAssociationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetMessage(v string) *CreateExpressConnectRouterAssociationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetRequestId(v string) *CreateExpressConnectRouterAssociationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetSuccess(v bool) *CreateExpressConnectRouterAssociationResponseBody {
	s.Success = &v
	return s
}

type CreateExpressConnectRouterAssociationResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExpressConnectRouterAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExpressConnectRouterAssociationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateExpressConnectRouterAssociationResponse) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectRouterAssociationResponse) SetHeaders(v map[string]*string) *CreateExpressConnectRouterAssociationResponse {
	s.Headers = v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponse) SetStatusCode(v int32) *CreateExpressConnectRouterAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponse) SetBody(v *CreateExpressConnectRouterAssociationResponseBody) *CreateExpressConnectRouterAssociationResponse {
	s.Body = v
	return s
}

type CreateFlowLogRequest struct {
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// myFlowlog
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// example:
	//
	// myFlowlog
	FlowLogName *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	// example:
	//
	// vbr-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// VBR
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// flowlog-logstore
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// example:
	//
	// flowlog-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 1:4096
	SamplingRate *string `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
}

func (s CreateFlowLogRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowLogRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowLogRequest) SetClientToken(v string) *CreateFlowLogRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFlowLogRequest) SetDescription(v string) *CreateFlowLogRequest {
	s.Description = &v
	return s
}

func (s *CreateFlowLogRequest) SetDryRun(v bool) *CreateFlowLogRequest {
	s.DryRun = &v
	return s
}

func (s *CreateFlowLogRequest) SetEcrId(v string) *CreateFlowLogRequest {
	s.EcrId = &v
	return s
}

func (s *CreateFlowLogRequest) SetFlowLogName(v string) *CreateFlowLogRequest {
	s.FlowLogName = &v
	return s
}

func (s *CreateFlowLogRequest) SetInstanceId(v string) *CreateFlowLogRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateFlowLogRequest) SetInstanceType(v string) *CreateFlowLogRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateFlowLogRequest) SetInterval(v int32) *CreateFlowLogRequest {
	s.Interval = &v
	return s
}

func (s *CreateFlowLogRequest) SetLogStoreName(v string) *CreateFlowLogRequest {
	s.LogStoreName = &v
	return s
}

func (s *CreateFlowLogRequest) SetProjectName(v string) *CreateFlowLogRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFlowLogRequest) SetSamplingRate(v string) *CreateFlowLogRequest {
	s.SamplingRate = &v
	return s
}

type CreateFlowLogResponseBody struct {
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// flowlog-m5evbtbpt****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFlowLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowLogResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowLogResponseBody) SetAccessDeniedDetail(v string) *CreateFlowLogResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateFlowLogResponseBody) SetCode(v string) *CreateFlowLogResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFlowLogResponseBody) SetDynamicCode(v string) *CreateFlowLogResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateFlowLogResponseBody) SetDynamicMessage(v string) *CreateFlowLogResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateFlowLogResponseBody) SetFlowLogId(v string) *CreateFlowLogResponseBody {
	s.FlowLogId = &v
	return s
}

func (s *CreateFlowLogResponseBody) SetHttpStatusCode(v int32) *CreateFlowLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateFlowLogResponseBody) SetMessage(v string) *CreateFlowLogResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFlowLogResponseBody) SetRequestId(v string) *CreateFlowLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlowLogResponseBody) SetSuccess(v bool) *CreateFlowLogResponseBody {
	s.Success = &v
	return s
}

type CreateFlowLogResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFlowLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFlowLogResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowLogResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowLogResponse) SetHeaders(v map[string]*string) *CreateFlowLogResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowLogResponse) SetStatusCode(v int32) *CreateFlowLogResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFlowLogResponse) SetBody(v *CreateFlowLogResponseBody) *CreateFlowLogResponse {
	s.Body = v
	return s
}

type DeactivateFlowLogRequest struct {
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// flowlog-m5evbtbpt****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
}

func (s DeactivateFlowLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DeactivateFlowLogRequest) GoString() string {
	return s.String()
}

func (s *DeactivateFlowLogRequest) SetClientToken(v string) *DeactivateFlowLogRequest {
	s.ClientToken = &v
	return s
}

func (s *DeactivateFlowLogRequest) SetDryRun(v bool) *DeactivateFlowLogRequest {
	s.DryRun = &v
	return s
}

func (s *DeactivateFlowLogRequest) SetEcrId(v string) *DeactivateFlowLogRequest {
	s.EcrId = &v
	return s
}

func (s *DeactivateFlowLogRequest) SetFlowLogId(v string) *DeactivateFlowLogRequest {
	s.FlowLogId = &v
	return s
}

type DeactivateFlowLogResponseBody struct {
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeactivateFlowLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeactivateFlowLogResponseBody) GoString() string {
	return s.String()
}

func (s *DeactivateFlowLogResponseBody) SetAccessDeniedDetail(v string) *DeactivateFlowLogResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeactivateFlowLogResponseBody) SetCode(v string) *DeactivateFlowLogResponseBody {
	s.Code = &v
	return s
}

func (s *DeactivateFlowLogResponseBody) SetDynamicCode(v string) *DeactivateFlowLogResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeactivateFlowLogResponseBody) SetDynamicMessage(v string) *DeactivateFlowLogResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeactivateFlowLogResponseBody) SetHttpStatusCode(v int32) *DeactivateFlowLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeactivateFlowLogResponseBody) SetMessage(v string) *DeactivateFlowLogResponseBody {
	s.Message = &v
	return s
}

func (s *DeactivateFlowLogResponseBody) SetRequestId(v string) *DeactivateFlowLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeactivateFlowLogResponseBody) SetSuccess(v bool) *DeactivateFlowLogResponseBody {
	s.Success = &v
	return s
}

type DeactivateFlowLogResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeactivateFlowLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeactivateFlowLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DeactivateFlowLogResponse) GoString() string {
	return s.String()
}

func (s *DeactivateFlowLogResponse) SetHeaders(v map[string]*string) *DeactivateFlowLogResponse {
	s.Headers = v
	return s
}

func (s *DeactivateFlowLogResponse) SetStatusCode(v int32) *DeactivateFlowLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DeactivateFlowLogResponse) SetBody(v *DeactivateFlowLogResponseBody) *DeactivateFlowLogResponse {
	s.Body = v
	return s
}

type DeleteExpressConnectRouterRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-fu8rszhgv7623c****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
}

func (s DeleteExpressConnectRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectRouterRequest) SetClientToken(v string) *DeleteExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteExpressConnectRouterRequest) SetDryRun(v bool) *DeleteExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteExpressConnectRouterRequest) SetEcrId(v string) *DeleteExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

type DeleteExpressConnectRouterResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteExpressConnectRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *DeleteExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteExpressConnectRouterResponseBody) SetCode(v string) *DeleteExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteExpressConnectRouterResponseBody) SetDynamicCode(v string) *DeleteExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteExpressConnectRouterResponseBody) SetDynamicMessage(v string) *DeleteExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *DeleteExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteExpressConnectRouterResponseBody) SetMessage(v string) *DeleteExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteExpressConnectRouterResponseBody) SetRequestId(v string) *DeleteExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExpressConnectRouterResponseBody) SetSuccess(v bool) *DeleteExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

type DeleteExpressConnectRouterResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExpressConnectRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectRouterResponse) SetHeaders(v map[string]*string) *DeleteExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *DeleteExpressConnectRouterResponse) SetStatusCode(v int32) *DeleteExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExpressConnectRouterResponse) SetBody(v *DeleteExpressConnectRouterResponseBody) *DeleteExpressConnectRouterResponse {
	s.Body = v
	return s
}

type DeleteExpressConnectRouterAssociationRequest struct {
	// The ID of the association between the ECR and the VPC or TR.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-assoc-9p2qxx5phpca2m****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to delete the association between the ECR and the VPC or TR. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	DeleteAttachment *bool `json:"DeleteAttachment,omitempty" xml:"DeleteAttachment,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
}

func (s DeleteExpressConnectRouterAssociationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteExpressConnectRouterAssociationRequest) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectRouterAssociationRequest) SetAssociationId(v string) *DeleteExpressConnectRouterAssociationRequest {
	s.AssociationId = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationRequest) SetClientToken(v string) *DeleteExpressConnectRouterAssociationRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationRequest) SetDeleteAttachment(v bool) *DeleteExpressConnectRouterAssociationRequest {
	s.DeleteAttachment = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationRequest) SetDryRun(v bool) *DeleteExpressConnectRouterAssociationRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationRequest) SetEcrId(v string) *DeleteExpressConnectRouterAssociationRequest {
	s.EcrId = &v
	return s
}

type DeleteExpressConnectRouterAssociationResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 05130E79-588D-5C40-A718-C4863A59****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteExpressConnectRouterAssociationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteExpressConnectRouterAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetAccessDeniedDetail(v string) *DeleteExpressConnectRouterAssociationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetCode(v string) *DeleteExpressConnectRouterAssociationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetDynamicCode(v string) *DeleteExpressConnectRouterAssociationResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetDynamicMessage(v string) *DeleteExpressConnectRouterAssociationResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetHttpStatusCode(v int32) *DeleteExpressConnectRouterAssociationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetMessage(v string) *DeleteExpressConnectRouterAssociationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetRequestId(v string) *DeleteExpressConnectRouterAssociationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetSuccess(v bool) *DeleteExpressConnectRouterAssociationResponseBody {
	s.Success = &v
	return s
}

type DeleteExpressConnectRouterAssociationResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExpressConnectRouterAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExpressConnectRouterAssociationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExpressConnectRouterAssociationResponse) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectRouterAssociationResponse) SetHeaders(v map[string]*string) *DeleteExpressConnectRouterAssociationResponse {
	s.Headers = v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponse) SetStatusCode(v int32) *DeleteExpressConnectRouterAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponse) SetBody(v *DeleteExpressConnectRouterAssociationResponseBody) *DeleteExpressConnectRouterAssociationResponse {
	s.Body = v
	return s
}

type DeleteFlowlogRequest struct {
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// flowlog-m5evbtbpt****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
}

func (s DeleteFlowlogRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowlogRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowlogRequest) SetClientToken(v string) *DeleteFlowlogRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFlowlogRequest) SetDryRun(v bool) *DeleteFlowlogRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteFlowlogRequest) SetEcrId(v string) *DeleteFlowlogRequest {
	s.EcrId = &v
	return s
}

func (s *DeleteFlowlogRequest) SetFlowLogId(v string) *DeleteFlowlogRequest {
	s.FlowLogId = &v
	return s
}

type DeleteFlowlogResponseBody struct {
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFlowlogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowlogResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowlogResponseBody) SetAccessDeniedDetail(v string) *DeleteFlowlogResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteFlowlogResponseBody) SetCode(v string) *DeleteFlowlogResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFlowlogResponseBody) SetDynamicCode(v string) *DeleteFlowlogResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteFlowlogResponseBody) SetDynamicMessage(v string) *DeleteFlowlogResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteFlowlogResponseBody) SetHttpStatusCode(v int32) *DeleteFlowlogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteFlowlogResponseBody) SetMessage(v string) *DeleteFlowlogResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFlowlogResponseBody) SetRequestId(v string) *DeleteFlowlogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFlowlogResponseBody) SetSuccess(v bool) *DeleteFlowlogResponseBody {
	s.Success = &v
	return s
}

type DeleteFlowlogResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFlowlogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFlowlogResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowlogResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowlogResponse) SetHeaders(v map[string]*string) *DeleteFlowlogResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowlogResponse) SetStatusCode(v int32) *DeleteFlowlogResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFlowlogResponse) SetBody(v *DeleteFlowlogResponseBody) *DeleteFlowlogResponse {
	s.Body = v
	return s
}

type DescribeDisabledExpressConnectRouterRouteEntriesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// FF9nMec/RZ6H9oqFn1pvyir/SLRlxCCyHJonbGzqL01hiM6Jb3wJowdHvjCfog7ww1b9rSHMRFJnrUBfVba68TJg==
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The maximum number of entries to read. Valid values: 1 to 2147483647. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// FFv4fzkNPW8Z+cZ+DBXXQ3Gmf3XlCgpBH43oaTYTAAcGc708Zb+pDyAGVJBo/MKsyrtZfPnX9Ztf02vgdIDyaNe8UuZdf/JJk069qxGKzqSKg=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) SetClientToken(v string) *DescribeDisabledExpressConnectRouterRouteEntriesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) SetDryRun(v bool) *DescribeDisabledExpressConnectRouterRouteEntriesRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) SetEcrId(v string) *DescribeDisabledExpressConnectRouterRouteEntriesRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) SetMaxResults(v int32) *DescribeDisabledExpressConnectRouterRouteEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) SetNextToken(v string) *DescribeDisabledExpressConnectRouterRouteEntriesRequest {
	s.NextToken = &v
	return s
}

type DescribeDisabledExpressConnectRouterRouteEntriesResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The routes that are queried.
	DisabledRouteEntryList []*DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList `json:"DisabledRouteEntryList,omitempty" xml:"DisabledRouteEntryList,omitempty" type:"Repeated"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsInstanceId**, the request parameter **DtsInstanceId*	- is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The total number of entries returned. Valid values: 1 to 2147483647. Default value: 10.
	//
	// example:
	//
	// 6
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If NextToken is empty, no next page exists.
	//
	// 	- If a value of NextToken is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// gAAAAABkDTaRFnmxUoMLVOn8YTIgYFeL2ch8j0sJs8VCIU8SS5438m3D9X1VqspCcaUEHRN9I_AfVwMhZHAhcNivifK_OtQxJQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether routes are disabled by the ECR. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of routes.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetAccessDeniedDetail(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetCode(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetDisabledRouteEntryList(v []*DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.DisabledRouteEntryList = v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetDynamicCode(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetDynamicMessage(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetHttpStatusCode(v int32) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetMaxResults(v int32) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetMessage(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetNextToken(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetRequestId(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetSuccess(v bool) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetTotalCount(v int32) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList struct {
	// The destination CIDR block of the route.
	//
	// example:
	//
	// 192.168.100.110/32
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The ECR ID.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The time when the route entry was created.
	//
	// example:
	//
	// 1682317345
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The ID of the next-hop instance.
	//
	// example:
	//
	// br-hp3u4u5f03tfuljis****
	NexthopInstanceId *string `json:"NexthopInstanceId,omitempty" xml:"NexthopInstanceId,omitempty"`
	// The region ID of the next-hop instance.
	//
	// example:
	//
	// cn-hangzhou
	NexthopInstanceRegionId *string `json:"NexthopInstanceRegionId,omitempty" xml:"NexthopInstanceRegionId,omitempty"`
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) SetDestinationCidrBlock(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) SetEcrId(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList {
	s.EcrId = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) SetGmtCreate(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) SetNexthopInstanceId(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList {
	s.NexthopInstanceId = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) SetNexthopInstanceRegionId(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList {
	s.NexthopInstanceRegionId = &v
	return s
}

type DescribeDisabledExpressConnectRouterRouteEntriesResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponse) SetHeaders(v map[string]*string) *DescribeDisabledExpressConnectRouterRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponse) SetStatusCode(v int32) *DescribeDisabledExpressConnectRouterRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponse) SetBody(v *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) *DescribeDisabledExpressConnectRouterRouteEntriesResponse {
	s.Body = v
	return s
}

type DescribeExpressConnectRouterRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// example:
	//
	// ecr-fu8rszhgv7623c****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The maximum number of entries to read. Valid values: 1 to 2147483647. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the ECR.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// e0a2dbeb69a8beeeb8194e92b702df3fd3e7bfe6ce7bfc16e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource group to which the ECR belongs.
	//
	// example:
	//
	// rg-aek2aq7f4va****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource tags. You can specify up to 20 tags.
	Tag []*DescribeExpressConnectRouterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeExpressConnectRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRequest) SetClientToken(v string) *DescribeExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectRouterRequest) SetDryRun(v bool) *DescribeExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeExpressConnectRouterRequest) SetEcrId(v string) *DescribeExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterRequest) SetMaxResults(v int32) *DescribeExpressConnectRouterRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterRequest) SetName(v string) *DescribeExpressConnectRouterRequest {
	s.Name = &v
	return s
}

func (s *DescribeExpressConnectRouterRequest) SetNextToken(v string) *DescribeExpressConnectRouterRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectRouterRequest) SetResourceGroupId(v string) *DescribeExpressConnectRouterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeExpressConnectRouterRequest) SetTag(v []*DescribeExpressConnectRouterRequestTag) *DescribeExpressConnectRouterRequest {
	s.Tag = v
	return s
}

type DescribeExpressConnectRouterRequestTag struct {
	// The tag keys.
	//
	// The tag keys cannot be an empty string. The tag keys can be up to 64 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// You can specify at most 20 tag keys.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// A tag value.
	//
	// A tag value can be a maximum of 128 characters in length. It cannot contain `http://` or `https://` and cannot start with `acs:`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeExpressConnectRouterRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRequestTag) SetKey(v string) *DescribeExpressConnectRouterRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeExpressConnectRouterRequestTag) SetValue(v string) *DescribeExpressConnectRouterRequestTag {
	s.Value = &v
	return s
}

type DescribeExpressConnectRouterResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.Name
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of Name ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The ECRs.
	EcrList []*DescribeExpressConnectRouterResponseBodyEcrList `json:"EcrList,omitempty" xml:"EcrList,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The total number of entries returned. Valid values: 1 to 2147483647. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If NextToken is empty, no next page exists.
	//
	// 	- If a value of NextToken is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// AAAAAdDWBF2w6Olxc+cMPjUtUMpttDGZkffvHCfhBKKNEyCVaq+WUEzuUWpp9+QOApNf6g==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of ECRs.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeExpressConnectRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetCode(v string) *DescribeExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetDynamicCode(v string) *DescribeExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetDynamicMessage(v string) *DescribeExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetEcrList(v []*DescribeExpressConnectRouterResponseBodyEcrList) *DescribeExpressConnectRouterResponseBody {
	s.EcrList = v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *DescribeExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetMaxResults(v int32) *DescribeExpressConnectRouterResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetMessage(v string) *DescribeExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetNextToken(v string) *DescribeExpressConnectRouterResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetRequestId(v string) *DescribeExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetSuccess(v bool) *DescribeExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetTotalCount(v int32) *DescribeExpressConnectRouterResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeExpressConnectRouterResponseBodyEcrList struct {
	// The autonomous system number (ASN) of the ECR.
	//
	// example:
	//
	// 45104
	AlibabaSideAsn *int64 `json:"AlibabaSideAsn,omitempty" xml:"AlibabaSideAsn,omitempty"`
	// The business state of the ECR. Valid values:
	//
	// 	- **Normal:*	- The ECR is running as expected.
	//
	// 	- **FinancialLocked**: The ECR is locked due to overdue payments.
	//
	// example:
	//
	// Normal
	BizStatus *string `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	// The description of the ECR.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ECR ID.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The time when the ECR was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-02-16T01:44:50Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the ECR was modified. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-02-16T01:44:50Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The name of the ECR.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the Alibaba Cloud account that owns the ECR.
	//
	// example:
	//
	// 170646818729****
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group to which the ECR belongs.
	//
	// example:
	//
	// rg-aekzuscospt****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The deployment state of the ECR. Valid values:
	//
	// 	- **ACTIVE**: The ECR is created.
	//
	// 	- **UPDATING**: The ECR is being modified.
	//
	// 	- **ASSOCIATING**: The ECR is being associated with resources.
	//
	// 	- **DISSOCIATING**: The resource is being disassociated from resources.
	//
	// 	- **LOCKED_ATTACHING**: The ECR is locked because it is being associated with an external system.
	//
	// 	- **LOCKED_DETACHING**: The ECR is locked because it is being disassociated from an external system.
	//
	// 	- **RECLAIMING**: The ECR is waiting to release resources.
	//
	// 	- **DELETING**: The ECR is being deleted.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags []*DescribeExpressConnectRouterResponseBodyEcrListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeExpressConnectRouterResponseBodyEcrList) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterResponseBodyEcrList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetAlibabaSideAsn(v int64) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.AlibabaSideAsn = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetBizStatus(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.BizStatus = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetDescription(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.Description = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetEcrId(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetGmtCreate(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetGmtModified(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.GmtModified = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetName(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.Name = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetOwnerId(v int64) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.OwnerId = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetResourceGroupId(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetStatus(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.Status = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetTags(v []*DescribeExpressConnectRouterResponseBodyEcrListTags) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.Tags = v
	return s
}

type DescribeExpressConnectRouterResponseBodyEcrListTags struct {
	// The tag key.
	//
	// example:
	//
	// test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeExpressConnectRouterResponseBodyEcrListTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterResponseBodyEcrListTags) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterResponseBodyEcrListTags) SetTagKey(v string) *DescribeExpressConnectRouterResponseBodyEcrListTags {
	s.TagKey = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrListTags) SetTagValue(v string) *DescribeExpressConnectRouterResponseBodyEcrListTags {
	s.TagValue = &v
	return s
}

type DescribeExpressConnectRouterResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectRouterResponse) SetStatusCode(v int32) *DescribeExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterResponse) SetBody(v *DescribeExpressConnectRouterResponseBody) *DescribeExpressConnectRouterResponse {
	s.Body = v
	return s
}

type DescribeExpressConnectRouterAllowedPrefixHistoryRequest struct {
	// The ID of the association between the ECR and the virtual private cloud (VPC) or transit router (TR).
	//
	// >  You must specify either **InstanceId*	- or **AssociationId**.
	//
	// example:
	//
	// ecr-assoc-9p2qxx5phpca2m****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// FF9nMec/RZ6H9oqFn1pvyir/SLRlxCCyHJonbGzqL01hiM6Jb3wJowdHvjCfog7ww1b9rSHMRFJnrUBfVba68TJg==
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The ID of the network instance that is associated with the ECR.
	//
	// >  You must specify either **InstanceId*	- or **AssociationId**.
	//
	// example:
	//
	// vbr-j6cwxhgg0s5nuephp****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **TR**
	//
	// example:
	//
	// VPC
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) SetAssociationId(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest {
	s.AssociationId = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) SetClientToken(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) SetDryRun(v bool) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) SetEcrId(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) SetInstanceId(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) SetInstanceType(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest {
	s.InstanceType = &v
	return s
}

type DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The historical route prefixes.
	AllowedPrefixHistoryList []*DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList `json:"AllowedPrefixHistoryList,omitempty" xml:"AllowedPrefixHistoryList,omitempty" type:"Repeated"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetAllowedPrefixHistoryList(v []*DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.AllowedPrefixHistoryList = v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetCode(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetDynamicCode(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetDynamicMessage(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetHttpStatusCode(v int32) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetMessage(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetRequestId(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetSuccess(v bool) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.Success = &v
	return s
}

type DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList struct {
	// The route prefix.
	AllowedPrefix []*string `json:"AllowedPrefix,omitempty" xml:"AllowedPrefix,omitempty" type:"Repeated"`
	// The time when the historical route prefix entry was created.
	//
	// example:
	//
	// 1673751163000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList) SetAllowedPrefix(v []*string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList {
	s.AllowedPrefix = v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList) SetGmtCreate(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList {
	s.GmtCreate = &v
	return s
}

type DescribeExpressConnectRouterAllowedPrefixHistoryResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponse) SetStatusCode(v int32) *DescribeExpressConnectRouterAllowedPrefixHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponse) SetBody(v *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) *DescribeExpressConnectRouterAllowedPrefixHistoryResponse {
	s.Body = v
	return s
}

type DescribeExpressConnectRouterAssociationRequest struct {
	// The ID of the association between the ECR and the VPC or TR.
	//
	// example:
	//
	// ecr-assoc-9p2qxx5phpca2m****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The type of the associated resource. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **TR**
	//
	// example:
	//
	// VPC
	AssociationNodeType *string `json:"AssociationNodeType,omitempty" xml:"AssociationNodeType,omitempty"`
	// The region ID of the VPC or TR.
	//
	// example:
	//
	// cn-hangzhou
	AssociationRegionId *string `json:"AssociationRegionId,omitempty" xml:"AssociationRegionId,omitempty"`
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// example:
	//
	// cen-of3o1the3i4gwb****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The maximum number of entries to read. Valid values: 1 to 2147483647. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAYws9fJ0Ur4MGm/5OkDoW/Zn0J0/sCjivzwX9oBcwFnWaaas/kSG+J/WzLOxJHS4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The TR ID.
	//
	// example:
	//
	// tr-2ze4i71c6be454e2l****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp1h37fchc6jmfyln****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeExpressConnectRouterAssociationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterAssociationRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetAssociationId(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.AssociationId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetAssociationNodeType(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.AssociationNodeType = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetAssociationRegionId(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.AssociationRegionId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetCenId(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.CenId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetClientToken(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetDryRun(v bool) *DescribeExpressConnectRouterAssociationRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetEcrId(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetMaxResults(v int32) *DescribeExpressConnectRouterAssociationRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetNextToken(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetTransitRouterId(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetVpcId(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.VpcId = &v
	return s
}

type DescribeExpressConnectRouterAssociationResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The associated resources.
	AssociationList []*DescribeExpressConnectRouterAssociationResponseBodyAssociationList `json:"AssociationList,omitempty" xml:"AssociationList,omitempty" type:"Repeated"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of DynamicMessage is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The total number of entries returned. Valid values: 1 to 2147483647. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// AAAAAYws9fJ0Ur4MGm/5OkDoW/Zn0J0/sCjivzwX9oBcwFnWaaas/kSG+J/WzLOxJHS4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 05130E79-588D-5C40-A718-C4863A59****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of associated resources.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeExpressConnectRouterAssociationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterAssociationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetAssociationList(v []*DescribeExpressConnectRouterAssociationResponseBodyAssociationList) *DescribeExpressConnectRouterAssociationResponseBody {
	s.AssociationList = v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetCode(v string) *DescribeExpressConnectRouterAssociationResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetDynamicCode(v string) *DescribeExpressConnectRouterAssociationResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetDynamicMessage(v string) *DescribeExpressConnectRouterAssociationResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetHttpStatusCode(v int32) *DescribeExpressConnectRouterAssociationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetMaxResults(v int32) *DescribeExpressConnectRouterAssociationResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetMessage(v string) *DescribeExpressConnectRouterAssociationResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetNextToken(v string) *DescribeExpressConnectRouterAssociationResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetRequestId(v string) *DescribeExpressConnectRouterAssociationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetSuccess(v bool) *DescribeExpressConnectRouterAssociationResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetTotalCount(v int32) *DescribeExpressConnectRouterAssociationResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeExpressConnectRouterAssociationResponseBodyAssociationList struct {
	// The allowed route prefixes.
	AllowedPrefixes     []*string `json:"AllowedPrefixes,omitempty" xml:"AllowedPrefixes,omitempty" type:"Repeated"`
	AllowedPrefixesMode *string   `json:"AllowedPrefixesMode,omitempty" xml:"AllowedPrefixesMode,omitempty"`
	// The ID of the association between the ECR and the VPC or TR.
	//
	// example:
	//
	// ecr-assoc-9p2qxx5phpca2m****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The type of the associated resource. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **TR**
	//
	// example:
	//
	// VPC
	AssociationNodeType *string `json:"AssociationNodeType,omitempty" xml:"AssociationNodeType,omitempty"`
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-5510frtx8zi54q****
	CenId       *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ECR ID.
	//
	// example:
	//
	// ecr-fu8rszhgv7623c****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The time when the association was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-01-09T12:18:23Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the association was modified. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-01-09T12:18:23Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the Alibaba Cloud account that owns the resource.
	//
	// example:
	//
	// 167509154715****
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The deployment state of the associated resource. Valid values:
	//
	// 	- **CREATING**: The resource is being created.
	//
	// 	- **ACTIVE**: The resource is created.
	//
	// 	- **INACTIVE**: The TR is pending to be associated with the ECR.
	//
	// 	- **ASSOCIATING**: The resource is being associated.
	//
	// 	- **DISSOCIATING**: The resource is being disassociated.
	//
	// 	- **UPDATING**: The resource is being modified.
	//
	// 	- **DELETING**: The resource is being deleted.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The TR ID.
	//
	// example:
	//
	// tr-2ze4i71c6be454e2l****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The ID of the Alibaba Cloud account that owns the TR.
	//
	// example:
	//
	// 189159362009****
	TransitRouterOwnerId *int64 `json:"TransitRouterOwnerId,omitempty" xml:"TransitRouterOwnerId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-2zeeaxet4i2j1a7n7****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the Alibaba Cloud account to which the VPC belongs.
	//
	// example:
	//
	// 146757288406****
	VpcOwnerId *int64 `json:"VpcOwnerId,omitempty" xml:"VpcOwnerId,omitempty"`
}

func (s DescribeExpressConnectRouterAssociationResponseBodyAssociationList) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterAssociationResponseBodyAssociationList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetAllowedPrefixes(v []*string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.AllowedPrefixes = v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetAllowedPrefixesMode(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.AllowedPrefixesMode = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetAssociationId(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.AssociationId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetAssociationNodeType(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.AssociationNodeType = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetCenId(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.CenId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetDescription(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.Description = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetEcrId(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetGmtCreate(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetGmtModified(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.GmtModified = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetOwnerId(v int64) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.OwnerId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetRegionId(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.RegionId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetStatus(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.Status = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetTransitRouterId(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetTransitRouterOwnerId(v int64) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.TransitRouterOwnerId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetVpcId(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.VpcId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetVpcOwnerId(v int64) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.VpcOwnerId = &v
	return s
}

type DescribeExpressConnectRouterAssociationResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectRouterAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectRouterAssociationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterAssociationResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAssociationResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectRouterAssociationResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponse) SetStatusCode(v int32) *DescribeExpressConnectRouterAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponse) SetBody(v *DescribeExpressConnectRouterAssociationResponseBody) *DescribeExpressConnectRouterAssociationResponse {
	s.Body = v
	return s
}

type DescribeExpressConnectRouterChildInstanceRequest struct {
	// The ID of the association between the ECR and the virtual private cloud (VPC) or transit router (TR).
	//
	// example:
	//
	// ecr-assoc-9p2qxx5phpca2m****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The ID of the network instance to be queried.
	//
	// example:
	//
	// vbr-j6cwxhgg0s5nuephp****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// The region ID of the network instance.
	//
	// example:
	//
	// cn-hangzhou
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	// The type of network instance. Set the value to VBR.
	//
	// example:
	//
	// VBR
	ChildInstanceType *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The maximum number of entries to read. Valid values: 1 to 2147483647. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If NextToken is empty, no next page exists.
	//
	// 	- If a value of NextToken is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// AAAAAYws9fJ0Ur4MGm/5OkDoW/Zn0J0/sCjivzwX9oBcwFnWaaas/kSG+J/WzLOxJHS4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeExpressConnectRouterChildInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterChildInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetAssociationId(v string) *DescribeExpressConnectRouterChildInstanceRequest {
	s.AssociationId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetChildInstanceId(v string) *DescribeExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetChildInstanceRegionId(v string) *DescribeExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetChildInstanceType(v string) *DescribeExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetClientToken(v string) *DescribeExpressConnectRouterChildInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetDryRun(v bool) *DescribeExpressConnectRouterChildInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetEcrId(v string) *DescribeExpressConnectRouterChildInstanceRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetMaxResults(v int32) *DescribeExpressConnectRouterChildInstanceRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetNextToken(v string) *DescribeExpressConnectRouterChildInstanceRequest {
	s.NextToken = &v
	return s
}

type DescribeExpressConnectRouterChildInstanceResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The VBRs.
	ChildInstanceList []*DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList `json:"ChildInstanceList,omitempty" xml:"ChildInstanceList,omitempty" type:"Repeated"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The total number of entries returned. Valid values: 1 to 2147483647. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// AAAAAYws9fJ0Ur4MGm/5OkDoW/Zn0J0/sCjivzwX9oBcwFnWaaas/kSG+J/WzLOxJHS4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 05130E79-588D-5C40-A718-C4863A59****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of associated resources.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeExpressConnectRouterChildInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterChildInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetChildInstanceList(v []*DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.ChildInstanceList = v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetCode(v string) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetDynamicCode(v string) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetDynamicMessage(v string) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetHttpStatusCode(v int32) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetMaxResults(v int32) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetMessage(v string) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetNextToken(v string) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetRequestId(v string) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetSuccess(v bool) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetTotalCount(v int32) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList struct {
	// The ID of the association between the ECR and the VPC or TR.
	//
	// example:
	//
	// ecr-assoc-9p2qxx5phpca2m****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The VBR ID.
	//
	// example:
	//
	// vbr-gw8vjq2zjux3ifsc9****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// The ID of the Alibaba Cloud account that owns the VBR.
	//
	// example:
	//
	// 112101171212****
	ChildInstanceOwnerId *int64 `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	// The region ID of the VBR.
	//
	// example:
	//
	// cn-hangzhou
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	// The type of the VBR. The value is **VBR**.
	//
	// example:
	//
	// VBR
	ChildInstanceType *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	// The description of the ECR.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ECR ID.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The time when the association was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-01-09T12:18:23Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the association was modified. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-01-09T12:18:23Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the Alibaba Cloud account that owns the VBR.
	//
	// example:
	//
	// 167509154715****
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the VBR.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The deployment state of the associated resource. Valid values:
	//
	// - **CREATING**: The resource is being created.
	//
	// - **ACTIVE**: The resource is created.
	//
	// - **ASSOCIATING**: The resource is being associated.
	//
	// - **DISSOCIATING**: The resource is being disassociated.
	//
	// - **UPDATING**: The resource is being modified.
	//
	// - **DELETING**: The resource is being deleted.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetAssociationId(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.AssociationId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetChildInstanceId(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetChildInstanceOwnerId(v int64) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetChildInstanceRegionId(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetChildInstanceType(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetDescription(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.Description = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetEcrId(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetGmtCreate(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetGmtModified(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.GmtModified = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetOwnerId(v int64) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.OwnerId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetRegionId(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.RegionId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetStatus(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.Status = &v
	return s
}

type DescribeExpressConnectRouterChildInstanceResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectRouterChildInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectRouterChildInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterChildInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterChildInstanceResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectRouterChildInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponse) SetStatusCode(v int32) *DescribeExpressConnectRouterChildInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponse) SetBody(v *DescribeExpressConnectRouterChildInstanceResponseBody) *DescribeExpressConnectRouterChildInstanceResponse {
	s.Body = v
	return s
}

type DescribeExpressConnectRouterInterRegionTransitModeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// FF9nMec/RZ6H9oqFn1pvyir/SLRlxCCyHJonbGzqL01hiM6Jb3wJowdHvjCfog7ww1b9rSHMRFJnrUBfVba68TJg==
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
}

func (s DescribeExpressConnectRouterInterRegionTransitModeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterInterRegionTransitModeRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeRequest) SetClientToken(v string) *DescribeExpressConnectRouterInterRegionTransitModeRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeRequest) SetDryRun(v bool) *DescribeExpressConnectRouterInterRegionTransitModeRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeRequest) SetEcrId(v string) *DescribeExpressConnectRouterInterRegionTransitModeRequest {
	s.EcrId = &v
	return s
}

type DescribeExpressConnectRouterInterRegionTransitModeResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The cross-region forwarding modes.
	InterRegionTransitModeList []*DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList `json:"InterRegionTransitModeList,omitempty" xml:"InterRegionTransitModeList,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeExpressConnectRouterInterRegionTransitModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterInterRegionTransitModeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetCode(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetDynamicCode(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetDynamicMessage(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetHttpStatusCode(v int32) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetInterRegionTransitModeList(v []*DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.InterRegionTransitModeList = v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetMessage(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetRequestId(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetSuccess(v bool) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.Success = &v
	return s
}

type DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList struct {
	// The cross-region forwarding mode of the ECR. Valid values:
	//
	// 	- **ECMP**: the load balancing mode.
	//
	// 	- **NearBy**: the nearby forwarding mode.
	//
	// example:
	//
	// ECMP
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The region ID of the ECR.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList) SetMode(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList {
	s.Mode = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList) SetRegionId(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList {
	s.RegionId = &v
	return s
}

type DescribeExpressConnectRouterInterRegionTransitModeResponse struct {
	Headers    map[string]*string                                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectRouterInterRegionTransitModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectRouterInterRegionTransitModeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterInterRegionTransitModeResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectRouterInterRegionTransitModeResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponse) SetStatusCode(v int32) *DescribeExpressConnectRouterInterRegionTransitModeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponse) SetBody(v *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) *DescribeExpressConnectRouterInterRegionTransitModeResponse {
	s.Body = v
	return s
}

type DescribeExpressConnectRouterRegionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- true: performs only a dry run.
	//
	// 	- false (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the ECR that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-fu8rszhgv7623c****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
}

func (s DescribeExpressConnectRouterRegionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterRegionRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRegionRequest) SetClientToken(v string) *DescribeExpressConnectRouterRegionRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionRequest) SetDryRun(v bool) *DescribeExpressConnectRouterRegionRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionRequest) SetEcrId(v string) *DescribeExpressConnectRouterRegionRequest {
	s.EcrId = &v
	return s
}

type DescribeExpressConnectRouterRegionResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The region in which the ECR feature is activated.
	RegionIdList []*string `json:"RegionIdList,omitempty" xml:"RegionIdList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeExpressConnectRouterRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterRegionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterRegionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetCode(v string) *DescribeExpressConnectRouterRegionResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetDynamicCode(v string) *DescribeExpressConnectRouterRegionResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetDynamicMessage(v string) *DescribeExpressConnectRouterRegionResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetHttpStatusCode(v int32) *DescribeExpressConnectRouterRegionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetMessage(v string) *DescribeExpressConnectRouterRegionResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetRegionIdList(v []*string) *DescribeExpressConnectRouterRegionResponseBody {
	s.RegionIdList = v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetRequestId(v string) *DescribeExpressConnectRouterRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetSuccess(v bool) *DescribeExpressConnectRouterRegionResponseBody {
	s.Success = &v
	return s
}

type DescribeExpressConnectRouterRegionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectRouterRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectRouterRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterRegionResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRegionResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectRouterRegionResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponse) SetStatusCode(v int32) *DescribeExpressConnectRouterRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponse) SetBody(v *DescribeExpressConnectRouterRegionResponseBody) *DescribeExpressConnectRouterRegionResponse {
	s.Body = v
	return s
}

type DescribeExpressConnectRouterRouteEntriesRequest struct {
	// The Autonomous System (AS) path of the route.
	//
	// example:
	//
	// [64993,64512]
	AsPath *string `json:"AsPath,omitempty" xml:"AsPath,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The community value that is carried in the Border Gateway Protocol (BGP) route.
	//
	// example:
	//
	// 9001:9263
	Community *string `json:"Community,omitempty" xml:"Community,omitempty"`
	// The destination CIDR block of the route that you want to query.
	//
	// example:
	//
	// 172.20.47.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The maximum number of entries to read. Valid values: 1 to 2147483647. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// FF9nMec/RZ6H9oqFn1pvyir/SLRlxCCyHJonbGzqL01hiM6Jb3wJowdHvjCfog7ww1b9rSHMRFJnrUBfVba68TJg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the next-hop instance.
	//
	// example:
	//
	// br-hp3u4u5f03tfuljis****
	NexthopInstanceId *string `json:"NexthopInstanceId,omitempty" xml:"NexthopInstanceId,omitempty"`
	// The region ID of the ECR.
	//
	// example:
	//
	// cn-hangzhou
	QueryRegionId *string `json:"QueryRegionId,omitempty" xml:"QueryRegionId,omitempty"`
}

func (s DescribeExpressConnectRouterRouteEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetAsPath(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.AsPath = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetClientToken(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetCommunity(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.Community = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetDestinationCidrBlock(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetDryRun(v bool) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetEcrId(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetMaxResults(v int32) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetNextToken(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetNexthopInstanceId(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.NexthopInstanceId = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetQueryRegionId(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.QueryRegionId = &v
	return s
}

type DescribeExpressConnectRouterRouteEntriesResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The total number of entries returned. Valid values: 1 to 2147483647. Default value: 10
	//
	// example:
	//
	// 6
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// gAAAAABkWwFTUMNCdWC0VMYOIylA56Hx6JUfCZlk5hQ5g_fnKmetN6303tqq5UJ2ouJzyT2fDOdzb-NqyEB5jcY8Z2euX7qHDA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The route entries.
	RouteEntriesList []*DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList `json:"RouteEntriesList,omitempty" xml:"RouteEntriesList,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of route entries.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeExpressConnectRouterRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetCode(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetDynamicCode(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetDynamicMessage(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetHttpStatusCode(v int32) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetMaxResults(v int32) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetMessage(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetNextToken(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetRequestId(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetRouteEntriesList(v []*DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.RouteEntriesList = v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetSuccess(v bool) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetTotalCount(v int32) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList struct {
	// The AS path of the route.
	//
	// example:
	//
	// [64993,64512]
	AsPath *string `json:"AsPath,omitempty" xml:"AsPath,omitempty"`
	// The community value that is carried in the BGP route.
	//
	// example:
	//
	// 9001:9263
	Community *string `json:"Community,omitempty" xml:"Community,omitempty"`
	// The destination CIDR block of the route.
	//
	// example:
	//
	// 192.168.0.0/16
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	Med                  *int64  `json:"Med,omitempty" xml:"Med,omitempty"`
	// The ID of the next-hop instance.
	//
	// example:
	//
	// br-hp3u4u5f03tfuljis****
	NexthopInstanceId *string `json:"NexthopInstanceId,omitempty" xml:"NexthopInstanceId,omitempty"`
	// The region ID of the next-hop instance.
	//
	// example:
	//
	// cn-hangzhou
	NexthopInstanceRegionId *string `json:"NexthopInstanceRegionId,omitempty" xml:"NexthopInstanceRegionId,omitempty"`
	// The state of the ECR.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) SetAsPath(v string) *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList {
	s.AsPath = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) SetCommunity(v string) *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList {
	s.Community = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) SetDestinationCidrBlock(v string) *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) SetMed(v int64) *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList {
	s.Med = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) SetNexthopInstanceId(v string) *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList {
	s.NexthopInstanceId = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) SetNexthopInstanceRegionId(v string) *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList {
	s.NexthopInstanceRegionId = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) SetStatus(v string) *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList {
	s.Status = &v
	return s
}

type DescribeExpressConnectRouterRouteEntriesResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectRouterRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectRouterRouteEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRouteEntriesResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectRouterRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponse) SetStatusCode(v int32) *DescribeExpressConnectRouterRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponse) SetBody(v *DescribeExpressConnectRouterRouteEntriesResponseBody) *DescribeExpressConnectRouterRouteEntriesResponse {
	s.Body = v
	return s
}

type DescribeFlowLogsRequest struct {
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ecr-a5xqrgbeidz1w8****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// example:
	//
	// flowlog-jqnr0veifo5d****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	// example:
	//
	// same-flowlog
	FlowLogName *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	// example:
	//
	// vbr-j6cwxhgg0s5nuephp****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// myFlowlog
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// gAAAAABkWw*******
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// myFlowlog
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DescribeFlowLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsRequest) SetClientToken(v string) *DescribeFlowLogsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetDryRun(v bool) *DescribeFlowLogsRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetEcrId(v string) *DescribeFlowLogsRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetFlowLogId(v string) *DescribeFlowLogsRequest {
	s.FlowLogId = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetFlowLogName(v string) *DescribeFlowLogsRequest {
	s.FlowLogName = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetInstanceId(v string) *DescribeFlowLogsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetLogStoreName(v string) *DescribeFlowLogsRequest {
	s.LogStoreName = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetMaxResults(v int32) *DescribeFlowLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetNextToken(v string) *DescribeFlowLogsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetProjectName(v string) *DescribeFlowLogsRequest {
	s.ProjectName = &v
	return s
}

type DescribeFlowLogsResponseBody struct {
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string                                 `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	FlowLogs       []*DescribeFlowLogsResponseBodyFlowLogs `json:"FlowLogs,omitempty" xml:"FlowLogs,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAAYws9fJ0Ur4MGm/5OkDoW/Zn0J0/sCjivzwX9oBcwFnWaaas/kSG+J/WzLOxJHS4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 11
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFlowLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsResponseBody) SetAccessDeniedDetail(v string) *DescribeFlowLogsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeFlowLogsResponseBody) SetCode(v string) *DescribeFlowLogsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeFlowLogsResponseBody) SetDynamicCode(v string) *DescribeFlowLogsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeFlowLogsResponseBody) SetDynamicMessage(v string) *DescribeFlowLogsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeFlowLogsResponseBody) SetFlowLogs(v []*DescribeFlowLogsResponseBodyFlowLogs) *DescribeFlowLogsResponseBody {
	s.FlowLogs = v
	return s
}

func (s *DescribeFlowLogsResponseBody) SetHttpStatusCode(v int32) *DescribeFlowLogsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeFlowLogsResponseBody) SetMaxResults(v int32) *DescribeFlowLogsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeFlowLogsResponseBody) SetMessage(v string) *DescribeFlowLogsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeFlowLogsResponseBody) SetNextToken(v string) *DescribeFlowLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeFlowLogsResponseBody) SetRequestId(v string) *DescribeFlowLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowLogsResponseBody) SetSuccess(v bool) *DescribeFlowLogsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFlowLogsResponseBody) SetTotalCount(v int32) *DescribeFlowLogsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeFlowLogsResponseBodyFlowLogs struct {
	// example:
	//
	// 2023-09-21T04:20Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// myFlowlog
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// ecr-h4cop1khw98*****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// example:
	//
	// flowlog-leypqehtgtia2*****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	// example:
	//
	// myFlowlog
	FlowLogName *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	// example:
	//
	// vbr-9dpty76irpf4u15*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// VBR
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 600
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// FlowLogStore
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// example:
	//
	// FlowLogProject
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aekzb3xwrexc4ry
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1:4096
	SamplingRate *string `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
	// example:
	//
	// cn-hangzhou
	SlsRegionId *string `json:"SlsRegionId,omitempty" xml:"SlsRegionId,omitempty"`
	// example:
	//
	// Active
	Status *string                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*DescribeFlowLogsResponseBodyFlowLogsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeFlowLogsResponseBodyFlowLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowLogsResponseBodyFlowLogs) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) SetCreationTime(v string) *DescribeFlowLogsResponseBodyFlowLogs {
	s.CreationTime = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) SetDescription(v string) *DescribeFlowLogsResponseBodyFlowLogs {
	s.Description = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) SetEcrId(v string) *DescribeFlowLogsResponseBodyFlowLogs {
	s.EcrId = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) SetFlowLogId(v string) *DescribeFlowLogsResponseBodyFlowLogs {
	s.FlowLogId = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) SetFlowLogName(v string) *DescribeFlowLogsResponseBodyFlowLogs {
	s.FlowLogName = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) SetInstanceId(v string) *DescribeFlowLogsResponseBodyFlowLogs {
	s.InstanceId = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) SetInstanceType(v string) *DescribeFlowLogsResponseBodyFlowLogs {
	s.InstanceType = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) SetInterval(v int32) *DescribeFlowLogsResponseBodyFlowLogs {
	s.Interval = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) SetLogStoreName(v string) *DescribeFlowLogsResponseBodyFlowLogs {
	s.LogStoreName = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) SetProjectName(v string) *DescribeFlowLogsResponseBodyFlowLogs {
	s.ProjectName = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) SetRegionId(v string) *DescribeFlowLogsResponseBodyFlowLogs {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) SetResourceGroupId(v string) *DescribeFlowLogsResponseBodyFlowLogs {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) SetSamplingRate(v string) *DescribeFlowLogsResponseBodyFlowLogs {
	s.SamplingRate = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) SetSlsRegionId(v string) *DescribeFlowLogsResponseBodyFlowLogs {
	s.SlsRegionId = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) SetStatus(v string) *DescribeFlowLogsResponseBodyFlowLogs {
	s.Status = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) SetTags(v []*DescribeFlowLogsResponseBodyFlowLogsTags) *DescribeFlowLogsResponseBodyFlowLogs {
	s.Tags = v
	return s
}

type DescribeFlowLogsResponseBodyFlowLogsTags struct {
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFlowLogsResponseBodyFlowLogsTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowLogsResponseBodyFlowLogsTags) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsResponseBodyFlowLogsTags) SetKey(v string) *DescribeFlowLogsResponseBodyFlowLogsTags {
	s.Key = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsTags) SetValue(v string) *DescribeFlowLogsResponseBodyFlowLogsTags {
	s.Value = &v
	return s
}

type DescribeFlowLogsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFlowLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFlowLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsResponse) SetHeaders(v map[string]*string) *DescribeFlowLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowLogsResponse) SetStatusCode(v int32) *DescribeFlowLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFlowLogsResponse) SetBody(v *DescribeFlowLogsResponseBody) *DescribeFlowLogsResponse {
	s.Body = v
	return s
}

type DescribeInstanceGrantedToExpressConnectRouterRequest struct {
	CallerType *string `json:"CallerType,omitempty" xml:"CallerType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The ID of the network instance.
	//
	// example:
	//
	// vbr-j6cwxhgg0s5nuephp****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the Alibaba Cloud account that owns the network instance.
	//
	// example:
	//
	// 129845258050****
	InstanceOwnerId *int64 `json:"InstanceOwnerId,omitempty" xml:"InstanceOwnerId,omitempty"`
	// The region ID of the network instance.
	//
	// example:
	//
	// cn-hangzhou
	InstanceRegionId *string `json:"InstanceRegionId,omitempty" xml:"InstanceRegionId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VBR**
	//
	// 	- **VPC**
	//
	// example:
	//
	// VBR
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum number of entries to read. Valid values: 1 to 2147483647. Default value: 20.
	//
	// example:
	//
	// 6
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// gAAAAABkyGzFbZR2NnxnyVk0EiL7F3qMBtBim8Es0mugRT3qb8yEHAMaHGanzuaHUmiEq9QRmok0RgxJAINBOJZa5KPjopEu_Q==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource group to which the network instance belongs.
	//
	// example:
	//
	// rg-aek2tsvbnfe****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	TagModels []*DescribeInstanceGrantedToExpressConnectRouterRequestTagModels `json:"TagModels,omitempty" xml:"TagModels,omitempty" type:"Repeated"`
}

func (s DescribeInstanceGrantedToExpressConnectRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceGrantedToExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetCallerType(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.CallerType = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetClientToken(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetDryRun(v bool) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetEcrId(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetInstanceId(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetInstanceOwnerId(v int64) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.InstanceOwnerId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetInstanceRegionId(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.InstanceRegionId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetInstanceType(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetMaxResults(v int32) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetNextToken(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetResourceGroupId(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetTagModels(v []*DescribeInstanceGrantedToExpressConnectRouterRequestTagModels) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.TagModels = v
	return s
}

type DescribeInstanceGrantedToExpressConnectRouterRequestTagModels struct {
	// The tag key. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value. You can specify up to 20 tag values. The tag value cannot be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeInstanceGrantedToExpressConnectRouterRequestTagModels) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceGrantedToExpressConnectRouterRequestTagModels) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequestTagModels) SetTagKey(v string) *DescribeInstanceGrantedToExpressConnectRouterRequestTagModels {
	s.TagKey = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequestTagModels) SetTagValue(v string) *DescribeInstanceGrantedToExpressConnectRouterRequestTagModels {
	s.TagValue = &v
	return s
}

type DescribeInstanceGrantedToExpressConnectRouterResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.Name
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of Name ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The network instances whose permissions are granted to the ECR.
	EcrGrantedInstanceList []*DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList `json:"EcrGrantedInstanceList,omitempty" xml:"EcrGrantedInstanceList,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The total number of entries returned. Valid values: 1 to 2147483647. Default value: 20.
	//
	// example:
	//
	// 6
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFlMqGuJ10uN3l+FX/cBrGDNXUOUifNeOuAJlT4dc3vsWD6DsNSFAC2FtpeH5QOSG2WFdyRoun7gSLCm5o69YnAQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of network instances whose permissions are granted to the ECR.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstanceGrantedToExpressConnectRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceGrantedToExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetCode(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetDynamicCode(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetDynamicMessage(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetEcrGrantedInstanceList(v []*DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.EcrGrantedInstanceList = v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetMaxResults(v int32) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetMessage(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetNextToken(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetRequestId(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetSuccess(v bool) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetTotalCount(v int32) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList struct {
	// The ECR ID.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId          *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	EcrOwnerAliUid *string `json:"EcrOwnerAliUid,omitempty" xml:"EcrOwnerAliUid,omitempty"`
	// The time when the network instance was created.
	//
	// example:
	//
	// 1669023139000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the network instance was modified.
	//
	// example:
	//
	// 1669023139000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The authorization ID.
	//
	// example:
	//
	// gr-8gdelo13mi99g1****
	GrantId *string `json:"GrantId,omitempty" xml:"GrantId,omitempty"`
	// The ID of the network instance.
	//
	// example:
	//
	// vbr-j6cwxhgg0s5nuephp****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The ID of the Alibaba Cloud enterprise account that owns the network instance.
	//
	// example:
	//
	// 26842
	NodeOwnerBid *string `json:"NodeOwnerBid,omitempty" xml:"NodeOwnerBid,omitempty"`
	// The ID of the Alibaba Cloud account that owns the network instance.
	//
	// example:
	//
	// 129845258050****
	NodeOwnerUid *int64 `json:"NodeOwnerUid,omitempty" xml:"NodeOwnerUid,omitempty"`
	// The region ID of the network instance.
	//
	// example:
	//
	// cn-hangzhou
	NodeRegionId *string `json:"NodeRegionId,omitempty" xml:"NodeRegionId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VBR**
	//
	// 	- **VPC**
	//
	// example:
	//
	// VBR
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The state of the network instance.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetEcrId(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.EcrId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetEcrOwnerAliUid(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.EcrOwnerAliUid = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetGmtCreate(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetGmtModified(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.GmtModified = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetGrantId(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.GrantId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetNodeId(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.NodeId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetNodeOwnerBid(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.NodeOwnerBid = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetNodeOwnerUid(v int64) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.NodeOwnerUid = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetNodeRegionId(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.NodeRegionId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetNodeType(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.NodeType = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetStatus(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.Status = &v
	return s
}

type DescribeInstanceGrantedToExpressConnectRouterResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceGrantedToExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceGrantedToExpressConnectRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceGrantedToExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponse) SetHeaders(v map[string]*string) *DescribeInstanceGrantedToExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponse) SetStatusCode(v int32) *DescribeInstanceGrantedToExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponse) SetBody(v *DescribeInstanceGrantedToExpressConnectRouterResponseBody) *DescribeInstanceGrantedToExpressConnectRouterResponse {
	s.Body = v
	return s
}

type DetachExpressConnectRouterChildInstanceRequest struct {
	// The VBR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-j6cwxhgg0s5nuephp****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// The type of the network instance. Set the value to **VBR**.
	//
	// This parameter is required.
	//
	// example:
	//
	// VBR
	ChildInstanceType *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
}

func (s DetachExpressConnectRouterChildInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachExpressConnectRouterChildInstanceRequest) GoString() string {
	return s.String()
}

func (s *DetachExpressConnectRouterChildInstanceRequest) SetChildInstanceId(v string) *DetachExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceRequest) SetChildInstanceType(v string) *DetachExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceRequest) SetClientToken(v string) *DetachExpressConnectRouterChildInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceRequest) SetDryRun(v bool) *DetachExpressConnectRouterChildInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceRequest) SetEcrId(v string) *DetachExpressConnectRouterChildInstanceRequest {
	s.EcrId = &v
	return s
}

type DetachExpressConnectRouterChildInstanceResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetachExpressConnectRouterChildInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachExpressConnectRouterChildInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetAccessDeniedDetail(v string) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetCode(v string) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetDynamicCode(v string) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetDynamicMessage(v string) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetHttpStatusCode(v int32) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetMessage(v string) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetRequestId(v string) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetSuccess(v bool) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.Success = &v
	return s
}

type DetachExpressConnectRouterChildInstanceResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachExpressConnectRouterChildInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachExpressConnectRouterChildInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachExpressConnectRouterChildInstanceResponse) GoString() string {
	return s.String()
}

func (s *DetachExpressConnectRouterChildInstanceResponse) SetHeaders(v map[string]*string) *DetachExpressConnectRouterChildInstanceResponse {
	s.Headers = v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponse) SetStatusCode(v int32) *DetachExpressConnectRouterChildInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponse) SetBody(v *DetachExpressConnectRouterChildInstanceResponseBody) *DetachExpressConnectRouterChildInstanceResponse {
	s.Body = v
	return s
}

type DisableExpressConnectRouterRouteEntriesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// FF9nMec/RZ6H9oqFn1pvyir/SLRlxCCyHJonbGzqL01hiM6Jb3wJowdHvjCfog7ww1b9rSHMRFJnrUBfVba68TJg==
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The destination CIDR block of the route entry in the route table of the ECR.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.153.32.16/28
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The ID of the next-hop instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-hp3u4u5f03tfuljis****
	NexthopInstanceId *string `json:"NexthopInstanceId,omitempty" xml:"NexthopInstanceId,omitempty"`
}

func (s DisableExpressConnectRouterRouteEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableExpressConnectRouterRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *DisableExpressConnectRouterRouteEntriesRequest) SetClientToken(v string) *DisableExpressConnectRouterRouteEntriesRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesRequest) SetDestinationCidrBlock(v string) *DisableExpressConnectRouterRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesRequest) SetDryRun(v bool) *DisableExpressConnectRouterRouteEntriesRequest {
	s.DryRun = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesRequest) SetEcrId(v string) *DisableExpressConnectRouterRouteEntriesRequest {
	s.EcrId = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesRequest) SetNexthopInstanceId(v string) *DisableExpressConnectRouterRouteEntriesRequest {
	s.NexthopInstanceId = &v
	return s
}

type DisableExpressConnectRouterRouteEntriesResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableExpressConnectRouterRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableExpressConnectRouterRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetAccessDeniedDetail(v string) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetCode(v string) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.Code = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetDynamicCode(v string) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetDynamicMessage(v string) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetHttpStatusCode(v int32) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetMessage(v string) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.Message = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetRequestId(v string) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetSuccess(v bool) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.Success = &v
	return s
}

type DisableExpressConnectRouterRouteEntriesResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableExpressConnectRouterRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableExpressConnectRouterRouteEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableExpressConnectRouterRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *DisableExpressConnectRouterRouteEntriesResponse) SetHeaders(v map[string]*string) *DisableExpressConnectRouterRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponse) SetStatusCode(v int32) *DisableExpressConnectRouterRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponse) SetBody(v *DisableExpressConnectRouterRouteEntriesResponseBody) *DisableExpressConnectRouterRouteEntriesResponse {
	s.Body = v
	return s
}

type EnableExpressConnectRouterRouteEntriesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// FF9nMec/RZ6H9oqFn1pvyir/SLRlxCCyHJonbGzqL01hiM6Jb3wJowdHvjCfog7ww1b9rSHMRFJnrUBfVba68TJg==
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The destination CIDR block of the route entry in the route table of the ECR.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.153.32.16/28
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The ID of the next-hop instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-hp3u4u5f03tfuljis****
	NexthopInstanceId *string `json:"NexthopInstanceId,omitempty" xml:"NexthopInstanceId,omitempty"`
}

func (s EnableExpressConnectRouterRouteEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableExpressConnectRouterRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) SetClientToken(v string) *EnableExpressConnectRouterRouteEntriesRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) SetDestinationCidrBlock(v string) *EnableExpressConnectRouterRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) SetDryRun(v bool) *EnableExpressConnectRouterRouteEntriesRequest {
	s.DryRun = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) SetEcrId(v string) *EnableExpressConnectRouterRouteEntriesRequest {
	s.EcrId = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) SetNexthopInstanceId(v string) *EnableExpressConnectRouterRouteEntriesRequest {
	s.NexthopInstanceId = &v
	return s
}

type EnableExpressConnectRouterRouteEntriesResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableExpressConnectRouterRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableExpressConnectRouterRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetAccessDeniedDetail(v string) *EnableExpressConnectRouterRouteEntriesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetCode(v string) *EnableExpressConnectRouterRouteEntriesResponseBody {
	s.Code = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetDynamicCode(v string) *EnableExpressConnectRouterRouteEntriesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetDynamicMessage(v string) *EnableExpressConnectRouterRouteEntriesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetHttpStatusCode(v int32) *EnableExpressConnectRouterRouteEntriesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetMessage(v string) *EnableExpressConnectRouterRouteEntriesResponseBody {
	s.Message = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetRequestId(v string) *EnableExpressConnectRouterRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetSuccess(v bool) *EnableExpressConnectRouterRouteEntriesResponseBody {
	s.Success = &v
	return s
}

type EnableExpressConnectRouterRouteEntriesResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableExpressConnectRouterRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableExpressConnectRouterRouteEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableExpressConnectRouterRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *EnableExpressConnectRouterRouteEntriesResponse) SetHeaders(v map[string]*string) *EnableExpressConnectRouterRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponse) SetStatusCode(v int32) *EnableExpressConnectRouterRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponse) SetBody(v *EnableExpressConnectRouterRouteEntriesResponseBody) *EnableExpressConnectRouterRouteEntriesResponse {
	s.Body = v
	return s
}

type ForceDeleteExpressConnectRouterRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-fu8rszhgv7623c****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
}

func (s ForceDeleteExpressConnectRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s ForceDeleteExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *ForceDeleteExpressConnectRouterRequest) SetClientToken(v string) *ForceDeleteExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterRequest) SetDryRun(v bool) *ForceDeleteExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterRequest) SetEcrId(v string) *ForceDeleteExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

type ForceDeleteExpressConnectRouterResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ForceDeleteExpressConnectRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ForceDeleteExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *ForceDeleteExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetCode(v string) *ForceDeleteExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetDynamicCode(v string) *ForceDeleteExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetDynamicMessage(v string) *ForceDeleteExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *ForceDeleteExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetMessage(v string) *ForceDeleteExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetRequestId(v string) *ForceDeleteExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetSuccess(v bool) *ForceDeleteExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

type ForceDeleteExpressConnectRouterResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ForceDeleteExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ForceDeleteExpressConnectRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s ForceDeleteExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *ForceDeleteExpressConnectRouterResponse) SetHeaders(v map[string]*string) *ForceDeleteExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponse) SetStatusCode(v int32) *ForceDeleteExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponse) SetBody(v *ForceDeleteExpressConnectRouterResponseBody) *ForceDeleteExpressConnectRouterResponse {
	s.Body = v
	return s
}

type GrantInstanceToExpressConnectRouterRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The ID of the Alibaba Cloud account that owns the ECR to which you want to grant permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// 121012345612****
	EcrOwnerAliUid *int64 `json:"EcrOwnerAliUid,omitempty" xml:"EcrOwnerAliUid,omitempty"`
	// The ID of the network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-j6cwxhgg0s5nuephp****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InstanceRegionId *string `json:"InstanceRegionId,omitempty" xml:"InstanceRegionId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VBR**
	//
	// 	- **VPC**
	//
	// This parameter is required.
	//
	// example:
	//
	// VBR
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s GrantInstanceToExpressConnectRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantInstanceToExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *GrantInstanceToExpressConnectRouterRequest) SetClientToken(v string) *GrantInstanceToExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterRequest) SetDryRun(v bool) *GrantInstanceToExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterRequest) SetEcrId(v string) *GrantInstanceToExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterRequest) SetEcrOwnerAliUid(v int64) *GrantInstanceToExpressConnectRouterRequest {
	s.EcrOwnerAliUid = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterRequest) SetInstanceId(v string) *GrantInstanceToExpressConnectRouterRequest {
	s.InstanceId = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterRequest) SetInstanceRegionId(v string) *GrantInstanceToExpressConnectRouterRequest {
	s.InstanceRegionId = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterRequest) SetInstanceType(v string) *GrantInstanceToExpressConnectRouterRequest {
	s.InstanceType = &v
	return s
}

type GrantInstanceToExpressConnectRouterResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GrantInstanceToExpressConnectRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantInstanceToExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *GrantInstanceToExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetCode(v string) *GrantInstanceToExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetDynamicCode(v string) *GrantInstanceToExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetDynamicMessage(v string) *GrantInstanceToExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *GrantInstanceToExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetMessage(v string) *GrantInstanceToExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetRequestId(v string) *GrantInstanceToExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetSuccess(v bool) *GrantInstanceToExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

type GrantInstanceToExpressConnectRouterResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantInstanceToExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantInstanceToExpressConnectRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantInstanceToExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *GrantInstanceToExpressConnectRouterResponse) SetHeaders(v map[string]*string) *GrantInstanceToExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponse) SetStatusCode(v int32) *GrantInstanceToExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponse) SetBody(v *GrantInstanceToExpressConnectRouterResponseBody) *GrantInstanceToExpressConnectRouterResponse {
	s.Body = v
	return s
}

type ListExpressConnectRouterSupportedRegionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **TR**
	//
	// 	- **VBR**
	//
	// 	- **VPC**
	//
	// This parameter is required.
	//
	// example:
	//
	// VBR
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s ListExpressConnectRouterSupportedRegionRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExpressConnectRouterSupportedRegionRequest) GoString() string {
	return s.String()
}

func (s *ListExpressConnectRouterSupportedRegionRequest) SetClientToken(v string) *ListExpressConnectRouterSupportedRegionRequest {
	s.ClientToken = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionRequest) SetNodeType(v string) *ListExpressConnectRouterSupportedRegionRequest {
	s.NodeType = &v
	return s
}

type ListExpressConnectRouterSupportedRegionResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- True
	//
	// 	- False
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The regions in which the ECR feature is activated.
	SupportedRegionIdList []*string `json:"SupportedRegionIdList,omitempty" xml:"SupportedRegionIdList,omitempty" type:"Repeated"`
}

func (s ListExpressConnectRouterSupportedRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExpressConnectRouterSupportedRegionResponseBody) GoString() string {
	return s.String()
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) SetCode(v string) *ListExpressConnectRouterSupportedRegionResponseBody {
	s.Code = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) SetHttpStatusCode(v int32) *ListExpressConnectRouterSupportedRegionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) SetMessage(v string) *ListExpressConnectRouterSupportedRegionResponseBody {
	s.Message = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) SetRequestId(v string) *ListExpressConnectRouterSupportedRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) SetSuccess(v bool) *ListExpressConnectRouterSupportedRegionResponseBody {
	s.Success = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) SetSupportedRegionIdList(v []*string) *ListExpressConnectRouterSupportedRegionResponseBody {
	s.SupportedRegionIdList = v
	return s
}

type ListExpressConnectRouterSupportedRegionResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExpressConnectRouterSupportedRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExpressConnectRouterSupportedRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExpressConnectRouterSupportedRegionResponse) GoString() string {
	return s.String()
}

func (s *ListExpressConnectRouterSupportedRegionResponse) SetHeaders(v map[string]*string) *ListExpressConnectRouterSupportedRegionResponse {
	s.Headers = v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponse) SetStatusCode(v int32) *ListExpressConnectRouterSupportedRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponse) SetBody(v *ListExpressConnectRouterSupportedRegionResponseBody) *ListExpressConnectRouterSupportedRegionResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The maximum number of entries to return for a single request. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAYws9fJ0Ur4MGm/5OkDoW/Zn0J0/sCjivzwX9oBcwFnWaaas/kSG+J/WzLOxJHS4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ECR IDs.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Set the value to EXPRESSCONNECTROUTER.
	//
	// This parameter is required.
	//
	// example:
	//
	// EXPRESSCONNECTROUTER
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	//
	// You can bind up to 20 tags to an ECR.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetMaxResults(v int32) *ListTagResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
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
	// The tag key.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// You can specify up to 20 tag keys in each call.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `acs:` or `aliyun`. The tag value can be an empty string.
	//
	// You can specify up to 20 tag values in each call.
	//
	// example:
	//
	// TestValue
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
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// AAAAAYws9fJ0Ur4MGm/5OkDoW/Zn0J0/sCjivzwX9oBcwFnWaaas/kSG+J/WzLOxJHS4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The tags.
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
	// The total number of records that meet the query conditions.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetCode(v string) *ListTagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetMaxResults(v int32) *ListTagResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetMessage(v string) *ListTagResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetSuccess(v bool) *ListTagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

func (s *ListTagResourcesResponseBody) SetTotalCount(v int32) *ListTagResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	// The resource ID.
	//
	// example:
	//
	// ecr-897j0jooxyr1aq****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. The value is **EXPRESSCONNECTROUTER**.
	//
	// example:
	//
	// EXPRESSCONNECTROUTER
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
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

type ModifyExpressConnectRouterRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the ECR.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-fu8rszhgv7623c****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The name of the ECR.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyExpressConnectRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterRequest) SetClientToken(v string) *ModifyExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyExpressConnectRouterRequest) SetDescription(v string) *ModifyExpressConnectRouterRequest {
	s.Description = &v
	return s
}

func (s *ModifyExpressConnectRouterRequest) SetDryRun(v bool) *ModifyExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyExpressConnectRouterRequest) SetEcrId(v string) *ModifyExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

func (s *ModifyExpressConnectRouterRequest) SetName(v string) *ModifyExpressConnectRouterRequest {
	s.Name = &v
	return s
}

type ModifyExpressConnectRouterResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.Name
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of Name ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyExpressConnectRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *ModifyExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyExpressConnectRouterResponseBody) SetCode(v string) *ModifyExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyExpressConnectRouterResponseBody) SetDynamicCode(v string) *ModifyExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ModifyExpressConnectRouterResponseBody) SetDynamicMessage(v string) *ModifyExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *ModifyExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyExpressConnectRouterResponseBody) SetMessage(v string) *ModifyExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyExpressConnectRouterResponseBody) SetRequestId(v string) *ModifyExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyExpressConnectRouterResponseBody) SetSuccess(v bool) *ModifyExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

type ModifyExpressConnectRouterResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyExpressConnectRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterResponse) SetHeaders(v map[string]*string) *ModifyExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *ModifyExpressConnectRouterResponse) SetStatusCode(v int32) *ModifyExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyExpressConnectRouterResponse) SetBody(v *ModifyExpressConnectRouterResponseBody) *ModifyExpressConnectRouterResponse {
	s.Body = v
	return s
}

type ModifyExpressConnectRouterAssociationAllowedPrefixRequest struct {
	// The allowed route prefixes.
	AllowedPrefixes     []*string `json:"AllowedPrefixes,omitempty" xml:"AllowedPrefixes,omitempty" type:"Repeated"`
	AllowedPrefixesMode *string   `json:"AllowedPrefixesMode,omitempty" xml:"AllowedPrefixesMode,omitempty"`
	// The ID of the association between the ECR and the VPC or TR.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-assoc-9p2qxx5phpca2m****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId        *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s ModifyExpressConnectRouterAssociationAllowedPrefixRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyExpressConnectRouterAssociationAllowedPrefixRequest) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) SetAllowedPrefixes(v []*string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest {
	s.AllowedPrefixes = v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) SetAllowedPrefixesMode(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest {
	s.AllowedPrefixesMode = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) SetAssociationId(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest {
	s.AssociationId = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) SetClientToken(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) SetDryRun(v bool) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) SetEcrId(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest {
	s.EcrId = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) SetOwnerAccount(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest {
	s.OwnerAccount = &v
	return s
}

type ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 05130E79-588D-5C40-A718-C4863A59****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetAccessDeniedDetail(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetCode(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetDynamicCode(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetDynamicMessage(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetHttpStatusCode(v int32) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetMessage(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetRequestId(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetSuccess(v bool) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.Success = &v
	return s
}

type ModifyExpressConnectRouterAssociationAllowedPrefixResponse struct {
	Headers    map[string]*string                                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyExpressConnectRouterAssociationAllowedPrefixResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyExpressConnectRouterAssociationAllowedPrefixResponse) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponse) SetHeaders(v map[string]*string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponse {
	s.Headers = v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponse) SetStatusCode(v int32) *ModifyExpressConnectRouterAssociationAllowedPrefixResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponse) SetBody(v *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) *ModifyExpressConnectRouterAssociationAllowedPrefixResponse {
	s.Body = v
	return s
}

type ModifyExpressConnectRouterInterRegionTransitModeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// FF9nMec/RZ6H9oqFn1pvyir/SLRlxCCyHJonbGzqL01hiM6Jb3wJowdHvjCfog7ww1b9rSHMRFJnrUBfVba68TJg==
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The cross-region forwarding modes.
	TransitModeList []*ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList `json:"TransitModeList,omitempty" xml:"TransitModeList,omitempty" type:"Repeated"`
}

func (s ModifyExpressConnectRouterInterRegionTransitModeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyExpressConnectRouterInterRegionTransitModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequest) SetClientToken(v string) *ModifyExpressConnectRouterInterRegionTransitModeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequest) SetDryRun(v bool) *ModifyExpressConnectRouterInterRegionTransitModeRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequest) SetEcrId(v string) *ModifyExpressConnectRouterInterRegionTransitModeRequest {
	s.EcrId = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequest) SetTransitModeList(v []*ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList) *ModifyExpressConnectRouterInterRegionTransitModeRequest {
	s.TransitModeList = v
	return s
}

type ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList struct {
	// The cross-domain forwarding mode of the ECR. Valid values:
	//
	// 	- **ECMP**: the load balancing mode.
	//
	// 	- **NearBy**: the nearby forwarding mode.
	//
	// example:
	//
	// ECMP
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The region ID of the ECR.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList) String() string {
	return tea.Prettify(s)
}

func (s ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList) SetMode(v string) *ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList {
	s.Mode = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList) SetRegionId(v string) *ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList {
	s.RegionId = &v
	return s
}

type ModifyExpressConnectRouterInterRegionTransitModeResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyExpressConnectRouterInterRegionTransitModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyExpressConnectRouterInterRegionTransitModeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetAccessDeniedDetail(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetCode(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetDynamicCode(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetDynamicMessage(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetHttpStatusCode(v int32) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetMessage(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetRequestId(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetSuccess(v bool) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.Success = &v
	return s
}

type ModifyExpressConnectRouterInterRegionTransitModeResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyExpressConnectRouterInterRegionTransitModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyExpressConnectRouterInterRegionTransitModeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyExpressConnectRouterInterRegionTransitModeResponse) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponse) SetHeaders(v map[string]*string) *ModifyExpressConnectRouterInterRegionTransitModeResponse {
	s.Headers = v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponse) SetStatusCode(v int32) *ModifyExpressConnectRouterInterRegionTransitModeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponse) SetBody(v *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) *ModifyExpressConnectRouterInterRegionTransitModeResponse {
	s.Body = v
	return s
}

type ModifyFlowLogAttributeRequest struct {
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// myFlowlog
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ecr-fu8rszhgv7623c****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// flowlog-m5evbtbpt****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	// example:
	//
	// myFlowlog
	FlowLogName *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	// example:
	//
	// 600
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// 1:4096
	SamplingRate *string `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
}

func (s ModifyFlowLogAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowLogAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowLogAttributeRequest) SetClientToken(v string) *ModifyFlowLogAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetDescription(v string) *ModifyFlowLogAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetDryRun(v bool) *ModifyFlowLogAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetEcrId(v string) *ModifyFlowLogAttributeRequest {
	s.EcrId = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetFlowLogId(v string) *ModifyFlowLogAttributeRequest {
	s.FlowLogId = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetFlowLogName(v string) *ModifyFlowLogAttributeRequest {
	s.FlowLogName = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetInterval(v int32) *ModifyFlowLogAttributeRequest {
	s.Interval = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetSamplingRate(v string) *ModifyFlowLogAttributeRequest {
	s.SamplingRate = &v
	return s
}

type ModifyFlowLogAttributeResponseBody struct {
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyFlowLogAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowLogAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowLogAttributeResponseBody) SetAccessDeniedDetail(v string) *ModifyFlowLogAttributeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyFlowLogAttributeResponseBody) SetCode(v string) *ModifyFlowLogAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyFlowLogAttributeResponseBody) SetDynamicCode(v string) *ModifyFlowLogAttributeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ModifyFlowLogAttributeResponseBody) SetDynamicMessage(v string) *ModifyFlowLogAttributeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyFlowLogAttributeResponseBody) SetHttpStatusCode(v int32) *ModifyFlowLogAttributeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyFlowLogAttributeResponseBody) SetMessage(v string) *ModifyFlowLogAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyFlowLogAttributeResponseBody) SetRequestId(v string) *ModifyFlowLogAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFlowLogAttributeResponseBody) SetSuccess(v bool) *ModifyFlowLogAttributeResponseBody {
	s.Success = &v
	return s
}

type ModifyFlowLogAttributeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFlowLogAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFlowLogAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowLogAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyFlowLogAttributeResponse) SetHeaders(v map[string]*string) *ModifyFlowLogAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyFlowLogAttributeResponse) SetStatusCode(v int32) *ModifyFlowLogAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFlowLogAttributeResponse) SetBody(v *ModifyFlowLogAttributeResponseBody) *ModifyFlowLogAttributeResponse {
	s.Body = v
	return s
}

type MoveResourceGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the new resource group. For more information about resource groups, see the "Resource Group" section of the [What is Resource Management?](https://help.aliyun.com/document_detail/94475.html) topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-acfmvt3xpr5****
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-897j0jooxyr1aq****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. Set the value to EXPRESSCONNECTROUTER.
	//
	// This parameter is required.
	//
	// example:
	//
	// EXPRESSCONNECTROUTER
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetClientToken(v string) *MoveResourceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *MoveResourceGroupRequest) SetDryRun(v bool) *MoveResourceGroupRequest {
	s.DryRun = &v
	return s
}

func (s *MoveResourceGroupRequest) SetNewResourceGroupId(v string) *MoveResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceType(v string) *MoveResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type MoveResourceGroupResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MoveResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBody) SetCode(v string) *MoveResourceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *MoveResourceGroupResponseBody) SetMessage(v string) *MoveResourceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *MoveResourceGroupResponseBody) SetRequestId(v string) *MoveResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveResourceGroupResponseBody) SetSuccess(v bool) *MoveResourceGroupResponseBody {
	s.Success = &v
	return s
}

type MoveResourceGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponse) SetHeaders(v map[string]*string) *MoveResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *MoveResourceGroupResponse) SetStatusCode(v int32) *MoveResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveResourceGroupResponse) SetBody(v *MoveResourceGroupResponseBody) *MoveResourceGroupResponse {
	s.Body = v
	return s
}

type RevokeInstanceFromExpressConnectRouterRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The ID of the Alibaba Cloud account that owns the ECR from which you want to revoke permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// 121012345612****
	EcrOwnerAliUid *int64 `json:"EcrOwnerAliUid,omitempty" xml:"EcrOwnerAliUid,omitempty"`
	// The ID of the network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-j6cwxhgg0s5nuephp****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InstanceRegionId *string `json:"InstanceRegionId,omitempty" xml:"InstanceRegionId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VBR**
	//
	// 	- **VPC**
	//
	// This parameter is required.
	//
	// example:
	//
	// VBR
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s RevokeInstanceFromExpressConnectRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeInstanceFromExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) SetClientToken(v string) *RevokeInstanceFromExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) SetDryRun(v bool) *RevokeInstanceFromExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) SetEcrId(v string) *RevokeInstanceFromExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) SetEcrOwnerAliUid(v int64) *RevokeInstanceFromExpressConnectRouterRequest {
	s.EcrOwnerAliUid = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) SetInstanceId(v string) *RevokeInstanceFromExpressConnectRouterRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) SetInstanceRegionId(v string) *RevokeInstanceFromExpressConnectRouterRequest {
	s.InstanceRegionId = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) SetInstanceType(v string) *RevokeInstanceFromExpressConnectRouterRequest {
	s.InstanceType = &v
	return s
}

type RevokeInstanceFromExpressConnectRouterResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RevokeInstanceFromExpressConnectRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeInstanceFromExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetCode(v string) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetDynamicCode(v string) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetDynamicMessage(v string) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetMessage(v string) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetRequestId(v string) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetSuccess(v bool) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

type RevokeInstanceFromExpressConnectRouterResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeInstanceFromExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeInstanceFromExpressConnectRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeInstanceFromExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromExpressConnectRouterResponse) SetHeaders(v map[string]*string) *RevokeInstanceFromExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponse) SetStatusCode(v int32) *RevokeInstanceFromExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponse) SetBody(v *RevokeInstanceFromExpressConnectRouterResponseBody) *RevokeInstanceFromExpressConnectRouterResponse {
	s.Body = v
	return s
}

type SynchronizeExpressConnectRouterInterRegionBandwidthRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-fu8rszhgv7623c****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
}

func (s SynchronizeExpressConnectRouterInterRegionBandwidthRequest) String() string {
	return tea.Prettify(s)
}

func (s SynchronizeExpressConnectRouterInterRegionBandwidthRequest) GoString() string {
	return s.String()
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthRequest) SetClientToken(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthRequest {
	s.ClientToken = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthRequest) SetDryRun(v bool) *SynchronizeExpressConnectRouterInterRegionBandwidthRequest {
	s.DryRun = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthRequest) SetEcrId(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthRequest {
	s.EcrId = &v
	return s
}

type SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsInstanceId**, the request parameter **DtsInstanceId*	- is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetAccessDeniedDetail(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetCode(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.Code = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetDynamicCode(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetDynamicMessage(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetHttpStatusCode(v int32) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetMessage(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.Message = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetRequestId(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetSuccess(v bool) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.Success = &v
	return s
}

type SynchronizeExpressConnectRouterInterRegionBandwidthResponse struct {
	Headers    map[string]*string                                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SynchronizeExpressConnectRouterInterRegionBandwidthResponse) String() string {
	return tea.Prettify(s)
}

func (s SynchronizeExpressConnectRouterInterRegionBandwidthResponse) GoString() string {
	return s.String()
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponse) SetHeaders(v map[string]*string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponse {
	s.Headers = v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponse) SetStatusCode(v int32) *SynchronizeExpressConnectRouterInterRegionBandwidthResponse {
	s.StatusCode = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponse) SetBody(v *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) *SynchronizeExpressConnectRouterInterRegionBandwidthResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR IDs.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Set the value to **EXPRESSCONNECTROUTER**.
	//
	// This parameter is required.
	//
	// example:
	//
	// EXPRESSCONNECTROUTER
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags to be added.
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

func (s *TagResourcesRequest) SetClientToken(v string) *TagResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *TagResourcesRequest) SetDryRun(v bool) *TagResourcesRequest {
	s.DryRun = &v
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
	// The tag key.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// You can specify up to 20 tag keys in each call.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `acs:` or `aliyun`. The tag value can be an empty string.
	//
	// You can add up to 20 tag values in each call.
	//
	// example:
	//
	// TestValue
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
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetCode(v string) *TagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *TagResourcesResponseBody) SetMessage(v string) *TagResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesResponseBody) SetSuccess(v bool) *TagResourcesResponseBody {
	s.Success = &v
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

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags. This parameter is valid only when the TagKey.N parameter is not specified. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR IDs.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Set the value to **EXPRESSCONNECTROUTER**.
	//
	// This parameter is required.
	//
	// example:
	//
	// EXPRESSCONNECTROUTER
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The keys of the tags that you want to remove from the ECRs.
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

func (s *UntagResourcesRequest) SetClientToken(v string) *UntagResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *UntagResourcesRequest) SetDryRun(v bool) *UntagResourcesRequest {
	s.DryRun = &v
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
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetCode(v string) *UntagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *UntagResourcesResponseBody) SetMessage(v string) *UntagResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) SetSuccess(v bool) *UntagResourcesResponseBody {
	s.Success = &v
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("expressconnectrouter"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Enables log delivery.
//
// @param request - ActivateFlowLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateFlowLogResponse
func (client *Client) ActivateFlowLogWithOptions(request *ActivateFlowLogRequest, runtime *util.RuntimeOptions) (_result *ActivateFlowLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.FlowLogId)) {
		body["FlowLogId"] = request.FlowLogId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ActivateFlowLog"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ActivateFlowLogResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ActivateFlowLogResponse{}
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
// Enables log delivery.
//
// @param request - ActivateFlowLogRequest
//
// @return ActivateFlowLogResponse
func (client *Client) ActivateFlowLog(request *ActivateFlowLogRequest) (_result *ActivateFlowLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActivateFlowLogResponse{}
	_body, _err := client.ActivateFlowLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates a virtual border router (VBR) with an Express Connect router (ECR).
//
// Description:
//
// Before you call the **AttachExpressConnectRouterChildInstance*	- operation to associate a VBR with an ECR, make sure that the ECR is in the **Active*	- state.
//
// @param request - AttachExpressConnectRouterChildInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachExpressConnectRouterChildInstanceResponse
func (client *Client) AttachExpressConnectRouterChildInstanceWithOptions(request *AttachExpressConnectRouterChildInstanceRequest, runtime *util.RuntimeOptions) (_result *AttachExpressConnectRouterChildInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChildInstanceId)) {
		body["ChildInstanceId"] = request.ChildInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ChildInstanceOwnerId)) {
		body["ChildInstanceOwnerId"] = request.ChildInstanceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ChildInstanceRegionId)) {
		body["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ChildInstanceType)) {
		body["ChildInstanceType"] = request.ChildInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachExpressConnectRouterChildInstance"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AttachExpressConnectRouterChildInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AttachExpressConnectRouterChildInstanceResponse{}
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
// Associates a virtual border router (VBR) with an Express Connect router (ECR).
//
// Description:
//
// Before you call the **AttachExpressConnectRouterChildInstance*	- operation to associate a VBR with an ECR, make sure that the ECR is in the **Active*	- state.
//
// @param request - AttachExpressConnectRouterChildInstanceRequest
//
// @return AttachExpressConnectRouterChildInstanceResponse
func (client *Client) AttachExpressConnectRouterChildInstance(request *AttachExpressConnectRouterChildInstanceRequest) (_result *AttachExpressConnectRouterChildInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.AttachExpressConnectRouterChildInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks the Cloud Data Transfer (CDT) service required to add a region to an Express Connect router (ECR).
//
// @param request - CheckAddRegionToExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckAddRegionToExpressConnectRouterResponse
func (client *Client) CheckAddRegionToExpressConnectRouterWithOptions(request *CheckAddRegionToExpressConnectRouterRequest, runtime *util.RuntimeOptions) (_result *CheckAddRegionToExpressConnectRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.FreshRegionId)) {
		body["FreshRegionId"] = request.FreshRegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckAddRegionToExpressConnectRouter"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CheckAddRegionToExpressConnectRouterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CheckAddRegionToExpressConnectRouterResponse{}
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
// Checks the Cloud Data Transfer (CDT) service required to add a region to an Express Connect router (ECR).
//
// @param request - CheckAddRegionToExpressConnectRouterRequest
//
// @return CheckAddRegionToExpressConnectRouterResponse
func (client *Client) CheckAddRegionToExpressConnectRouter(request *CheckAddRegionToExpressConnectRouterRequest) (_result *CheckAddRegionToExpressConnectRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckAddRegionToExpressConnectRouterResponse{}
	_body, _err := client.CheckAddRegionToExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an Express Connect Router (ECR).
//
// Description:
//
// After you create an ECR, it enters the **Active*	- state.
//
// @param request - CreateExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExpressConnectRouterResponse
func (client *Client) CreateExpressConnectRouterWithOptions(request *CreateExpressConnectRouterRequest, runtime *util.RuntimeOptions) (_result *CreateExpressConnectRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlibabaSideAsn)) {
		body["AlibabaSideAsn"] = request.AlibabaSideAsn
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateExpressConnectRouter"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateExpressConnectRouterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateExpressConnectRouterResponse{}
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
// Creates an Express Connect Router (ECR).
//
// Description:
//
// After you create an ECR, it enters the **Active*	- state.
//
// @param request - CreateExpressConnectRouterRequest
//
// @return CreateExpressConnectRouterResponse
func (client *Client) CreateExpressConnectRouter(request *CreateExpressConnectRouterRequest) (_result *CreateExpressConnectRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateExpressConnectRouterResponse{}
	_body, _err := client.CreateExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates a virtual private cloud (VPC) or a transit router (TR) with an Express Connect router (ECR).
//
// @param request - CreateExpressConnectRouterAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExpressConnectRouterAssociationResponse
func (client *Client) CreateExpressConnectRouterAssociationWithOptions(request *CreateExpressConnectRouterAssociationRequest, runtime *util.RuntimeOptions) (_result *CreateExpressConnectRouterAssociationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowedPrefixes)) {
		body["AllowedPrefixes"] = request.AllowedPrefixes
	}

	if !tea.BoolValue(util.IsUnset(request.AllowedPrefixesMode)) {
		body["AllowedPrefixesMode"] = request.AllowedPrefixesMode
	}

	if !tea.BoolValue(util.IsUnset(request.AssociationRegionId)) {
		body["AssociationRegionId"] = request.AssociationRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.CenId)) {
		body["CenId"] = request.CenId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CreateAttachment)) {
		body["CreateAttachment"] = request.CreateAttachment
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.TransitRouterId)) {
		body["TransitRouterId"] = request.TransitRouterId
	}

	if !tea.BoolValue(util.IsUnset(request.TransitRouterOwnerId)) {
		body["TransitRouterOwnerId"] = request.TransitRouterOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcOwnerId)) {
		body["VpcOwnerId"] = request.VpcOwnerId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateExpressConnectRouterAssociation"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateExpressConnectRouterAssociationResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateExpressConnectRouterAssociationResponse{}
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
// Associates a virtual private cloud (VPC) or a transit router (TR) with an Express Connect router (ECR).
//
// @param request - CreateExpressConnectRouterAssociationRequest
//
// @return CreateExpressConnectRouterAssociationResponse
func (client *Client) CreateExpressConnectRouterAssociation(request *CreateExpressConnectRouterAssociationRequest) (_result *CreateExpressConnectRouterAssociationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateExpressConnectRouterAssociationResponse{}
	_body, _err := client.CreateExpressConnectRouterAssociationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建流日志
//
// @param request - CreateFlowLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFlowLogResponse
func (client *Client) CreateFlowLogWithOptions(request *CreateFlowLogRequest, runtime *util.RuntimeOptions) (_result *CreateFlowLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.LogStoreName)) {
		query["LogStoreName"] = request.LogStoreName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.SamplingRate)) {
		query["SamplingRate"] = request.SamplingRate
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.FlowLogName)) {
		body["FlowLogName"] = request.FlowLogName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFlowLog"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateFlowLogResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateFlowLogResponse{}
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
// 创建流日志
//
// @param request - CreateFlowLogRequest
//
// @return CreateFlowLogResponse
func (client *Client) CreateFlowLog(request *CreateFlowLogRequest) (_result *CreateFlowLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlowLogResponse{}
	_body, _err := client.CreateFlowLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止流日志
//
// @param request - DeactivateFlowLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeactivateFlowLogResponse
func (client *Client) DeactivateFlowLogWithOptions(request *DeactivateFlowLogRequest, runtime *util.RuntimeOptions) (_result *DeactivateFlowLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.FlowLogId)) {
		body["FlowLogId"] = request.FlowLogId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeactivateFlowLog"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeactivateFlowLogResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeactivateFlowLogResponse{}
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
// 停止流日志
//
// @param request - DeactivateFlowLogRequest
//
// @return DeactivateFlowLogResponse
func (client *Client) DeactivateFlowLog(request *DeactivateFlowLogRequest) (_result *DeactivateFlowLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeactivateFlowLogResponse{}
	_body, _err := client.DeactivateFlowLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an Express Connect router (ECR).
//
// Description:
//
// Take note of the following items:
//
// 	- Before you call this operation, make sure that all resources are disassociated from the ECR.
//
// 	- You can delete only ECRs that are in the **Active*	- state.
//
// @param request - DeleteExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExpressConnectRouterResponse
func (client *Client) DeleteExpressConnectRouterWithOptions(request *DeleteExpressConnectRouterRequest, runtime *util.RuntimeOptions) (_result *DeleteExpressConnectRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteExpressConnectRouter"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteExpressConnectRouterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteExpressConnectRouterResponse{}
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
// Deletes an Express Connect router (ECR).
//
// Description:
//
// Take note of the following items:
//
// 	- Before you call this operation, make sure that all resources are disassociated from the ECR.
//
// 	- You can delete only ECRs that are in the **Active*	- state.
//
// @param request - DeleteExpressConnectRouterRequest
//
// @return DeleteExpressConnectRouterResponse
func (client *Client) DeleteExpressConnectRouter(request *DeleteExpressConnectRouterRequest) (_result *DeleteExpressConnectRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteExpressConnectRouterResponse{}
	_body, _err := client.DeleteExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates an Express Connect router (ECR) from a virtual private cloud (VPC) or a transit router (TR).
//
// @param request - DeleteExpressConnectRouterAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExpressConnectRouterAssociationResponse
func (client *Client) DeleteExpressConnectRouterAssociationWithOptions(request *DeleteExpressConnectRouterAssociationRequest, runtime *util.RuntimeOptions) (_result *DeleteExpressConnectRouterAssociationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssociationId)) {
		body["AssociationId"] = request.AssociationId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DeleteAttachment)) {
		body["DeleteAttachment"] = request.DeleteAttachment
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteExpressConnectRouterAssociation"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteExpressConnectRouterAssociationResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteExpressConnectRouterAssociationResponse{}
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
// Disassociates an Express Connect router (ECR) from a virtual private cloud (VPC) or a transit router (TR).
//
// @param request - DeleteExpressConnectRouterAssociationRequest
//
// @return DeleteExpressConnectRouterAssociationResponse
func (client *Client) DeleteExpressConnectRouterAssociation(request *DeleteExpressConnectRouterAssociationRequest) (_result *DeleteExpressConnectRouterAssociationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteExpressConnectRouterAssociationResponse{}
	_body, _err := client.DeleteExpressConnectRouterAssociationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除流日志
//
// @param request - DeleteFlowlogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlowlogResponse
func (client *Client) DeleteFlowlogWithOptions(request *DeleteFlowlogRequest, runtime *util.RuntimeOptions) (_result *DeleteFlowlogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FlowLogId)) {
		query["FlowLogId"] = request.FlowLogId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFlowlog"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteFlowlogResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteFlowlogResponse{}
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
// 删除流日志
//
// @param request - DeleteFlowlogRequest
//
// @return DeleteFlowlogResponse
func (client *Client) DeleteFlowlog(request *DeleteFlowlogRequest) (_result *DeleteFlowlogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFlowlogResponse{}
	_body, _err := client.DeleteFlowlogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the route entries that are disabled on an Express Connect router (ECR).
//
// @param request - DescribeDisabledExpressConnectRouterRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDisabledExpressConnectRouterRouteEntriesResponse
func (client *Client) DescribeDisabledExpressConnectRouterRouteEntriesWithOptions(request *DescribeDisabledExpressConnectRouterRouteEntriesRequest, runtime *util.RuntimeOptions) (_result *DescribeDisabledExpressConnectRouterRouteEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDisabledExpressConnectRouterRouteEntries"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDisabledExpressConnectRouterRouteEntriesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDisabledExpressConnectRouterRouteEntriesResponse{}
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
// Queries the route entries that are disabled on an Express Connect router (ECR).
//
// @param request - DescribeDisabledExpressConnectRouterRouteEntriesRequest
//
// @return DescribeDisabledExpressConnectRouterRouteEntriesResponse
func (client *Client) DescribeDisabledExpressConnectRouterRouteEntries(request *DescribeDisabledExpressConnectRouterRouteEntriesRequest) (_result *DescribeDisabledExpressConnectRouterRouteEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDisabledExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.DescribeDisabledExpressConnectRouterRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of Express Connect routers (ECRs).
//
// @param request - DescribeExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterResponse
func (client *Client) DescribeExpressConnectRouterWithOptions(request *DescribeExpressConnectRouterRequest, runtime *util.RuntimeOptions) (_result *DescribeExpressConnectRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExpressConnectRouter"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeExpressConnectRouterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeExpressConnectRouterResponse{}
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
// Queries a list of Express Connect routers (ECRs).
//
// @param request - DescribeExpressConnectRouterRequest
//
// @return DescribeExpressConnectRouterResponse
func (client *Client) DescribeExpressConnectRouter(request *DescribeExpressConnectRouterRequest) (_result *DescribeExpressConnectRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterResponse{}
	_body, _err := client.DescribeExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the historical route prefixes of an Express Connect router (ECR).
//
// @param request - DescribeExpressConnectRouterAllowedPrefixHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterAllowedPrefixHistoryResponse
func (client *Client) DescribeExpressConnectRouterAllowedPrefixHistoryWithOptions(request *DescribeExpressConnectRouterAllowedPrefixHistoryRequest, runtime *util.RuntimeOptions) (_result *DescribeExpressConnectRouterAllowedPrefixHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssociationId)) {
		body["AssociationId"] = request.AssociationId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["InstanceType"] = request.InstanceType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExpressConnectRouterAllowedPrefixHistory"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeExpressConnectRouterAllowedPrefixHistoryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeExpressConnectRouterAllowedPrefixHistoryResponse{}
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
// Queries the historical route prefixes of an Express Connect router (ECR).
//
// @param request - DescribeExpressConnectRouterAllowedPrefixHistoryRequest
//
// @return DescribeExpressConnectRouterAllowedPrefixHistoryResponse
func (client *Client) DescribeExpressConnectRouterAllowedPrefixHistory(request *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) (_result *DescribeExpressConnectRouterAllowedPrefixHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterAllowedPrefixHistoryResponse{}
	_body, _err := client.DescribeExpressConnectRouterAllowedPrefixHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the virtual private clouds (VPCs) and transit routers (TRs) associated with an Express Connect router (ECR).
//
// @param request - DescribeExpressConnectRouterAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterAssociationResponse
func (client *Client) DescribeExpressConnectRouterAssociationWithOptions(request *DescribeExpressConnectRouterAssociationRequest, runtime *util.RuntimeOptions) (_result *DescribeExpressConnectRouterAssociationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssociationId)) {
		body["AssociationId"] = request.AssociationId
	}

	if !tea.BoolValue(util.IsUnset(request.AssociationNodeType)) {
		body["AssociationNodeType"] = request.AssociationNodeType
	}

	if !tea.BoolValue(util.IsUnset(request.AssociationRegionId)) {
		body["AssociationRegionId"] = request.AssociationRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.CenId)) {
		body["CenId"] = request.CenId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.TransitRouterId)) {
		body["TransitRouterId"] = request.TransitRouterId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExpressConnectRouterAssociation"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeExpressConnectRouterAssociationResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeExpressConnectRouterAssociationResponse{}
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
// Queries the virtual private clouds (VPCs) and transit routers (TRs) associated with an Express Connect router (ECR).
//
// @param request - DescribeExpressConnectRouterAssociationRequest
//
// @return DescribeExpressConnectRouterAssociationResponse
func (client *Client) DescribeExpressConnectRouterAssociation(request *DescribeExpressConnectRouterAssociationRequest) (_result *DescribeExpressConnectRouterAssociationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterAssociationResponse{}
	_body, _err := client.DescribeExpressConnectRouterAssociationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the virtual border routers (VBRs) that are associated with an Express Connect router (ECR).
//
// @param request - DescribeExpressConnectRouterChildInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterChildInstanceResponse
func (client *Client) DescribeExpressConnectRouterChildInstanceWithOptions(request *DescribeExpressConnectRouterChildInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeExpressConnectRouterChildInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssociationId)) {
		body["AssociationId"] = request.AssociationId
	}

	if !tea.BoolValue(util.IsUnset(request.ChildInstanceId)) {
		body["ChildInstanceId"] = request.ChildInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ChildInstanceRegionId)) {
		body["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ChildInstanceType)) {
		body["ChildInstanceType"] = request.ChildInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExpressConnectRouterChildInstance"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeExpressConnectRouterChildInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeExpressConnectRouterChildInstanceResponse{}
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
// Queries the virtual border routers (VBRs) that are associated with an Express Connect router (ECR).
//
// @param request - DescribeExpressConnectRouterChildInstanceRequest
//
// @return DescribeExpressConnectRouterChildInstanceResponse
func (client *Client) DescribeExpressConnectRouterChildInstance(request *DescribeExpressConnectRouterChildInstanceRequest) (_result *DescribeExpressConnectRouterChildInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.DescribeExpressConnectRouterChildInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the cross-region forwarding modes of an Express Connect router (ECR).
//
// @param request - DescribeExpressConnectRouterInterRegionTransitModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterInterRegionTransitModeResponse
func (client *Client) DescribeExpressConnectRouterInterRegionTransitModeWithOptions(request *DescribeExpressConnectRouterInterRegionTransitModeRequest, runtime *util.RuntimeOptions) (_result *DescribeExpressConnectRouterInterRegionTransitModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExpressConnectRouterInterRegionTransitMode"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeExpressConnectRouterInterRegionTransitModeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeExpressConnectRouterInterRegionTransitModeResponse{}
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
// Queries the cross-region forwarding modes of an Express Connect router (ECR).
//
// @param request - DescribeExpressConnectRouterInterRegionTransitModeRequest
//
// @return DescribeExpressConnectRouterInterRegionTransitModeResponse
func (client *Client) DescribeExpressConnectRouterInterRegionTransitMode(request *DescribeExpressConnectRouterInterRegionTransitModeRequest) (_result *DescribeExpressConnectRouterInterRegionTransitModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterInterRegionTransitModeResponse{}
	_body, _err := client.DescribeExpressConnectRouterInterRegionTransitModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of regions in which resources related to an Express Connect router (ECR) are deployed.
//
// @param request - DescribeExpressConnectRouterRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterRegionResponse
func (client *Client) DescribeExpressConnectRouterRegionWithOptions(request *DescribeExpressConnectRouterRegionRequest, runtime *util.RuntimeOptions) (_result *DescribeExpressConnectRouterRegionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExpressConnectRouterRegion"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeExpressConnectRouterRegionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeExpressConnectRouterRegionResponse{}
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
// Queries a list of regions in which resources related to an Express Connect router (ECR) are deployed.
//
// @param request - DescribeExpressConnectRouterRegionRequest
//
// @return DescribeExpressConnectRouterRegionResponse
func (client *Client) DescribeExpressConnectRouterRegion(request *DescribeExpressConnectRouterRegionRequest) (_result *DescribeExpressConnectRouterRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterRegionResponse{}
	_body, _err := client.DescribeExpressConnectRouterRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the route entries of an Express Connect router (ECR).
//
// @param request - DescribeExpressConnectRouterRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterRouteEntriesResponse
func (client *Client) DescribeExpressConnectRouterRouteEntriesWithOptions(request *DescribeExpressConnectRouterRouteEntriesRequest, runtime *util.RuntimeOptions) (_result *DescribeExpressConnectRouterRouteEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AsPath)) {
		body["AsPath"] = request.AsPath
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Community)) {
		body["Community"] = request.Community
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationCidrBlock)) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.NexthopInstanceId)) {
		body["NexthopInstanceId"] = request.NexthopInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryRegionId)) {
		body["QueryRegionId"] = request.QueryRegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExpressConnectRouterRouteEntries"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeExpressConnectRouterRouteEntriesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeExpressConnectRouterRouteEntriesResponse{}
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
// Queries the route entries of an Express Connect router (ECR).
//
// @param request - DescribeExpressConnectRouterRouteEntriesRequest
//
// @return DescribeExpressConnectRouterRouteEntriesResponse
func (client *Client) DescribeExpressConnectRouterRouteEntries(request *DescribeExpressConnectRouterRouteEntriesRequest) (_result *DescribeExpressConnectRouterRouteEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.DescribeExpressConnectRouterRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询流日志
//
// @param request - DescribeFlowLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFlowLogsResponse
func (client *Client) DescribeFlowLogsWithOptions(request *DescribeFlowLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FlowLogId)) {
		query["FlowLogId"] = request.FlowLogId
	}

	if !tea.BoolValue(util.IsUnset(request.FlowLogName)) {
		query["FlowLogName"] = request.FlowLogName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LogStoreName)) {
		query["LogStoreName"] = request.LogStoreName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFlowLogs"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeFlowLogsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeFlowLogsResponse{}
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
// 查询流日志
//
// @param request - DescribeFlowLogsRequest
//
// @return DescribeFlowLogsResponse
func (client *Client) DescribeFlowLogs(request *DescribeFlowLogsRequest) (_result *DescribeFlowLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowLogsResponse{}
	_body, _err := client.DescribeFlowLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the network instances whose permissions are granted to an Express Connect router (ECR).
//
// @param request - DescribeInstanceGrantedToExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceGrantedToExpressConnectRouterResponse
func (client *Client) DescribeInstanceGrantedToExpressConnectRouterWithOptions(request *DescribeInstanceGrantedToExpressConnectRouterRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceGrantedToExpressConnectRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallerType)) {
		body["CallerType"] = request.CallerType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceOwnerId)) {
		body["InstanceOwnerId"] = request.InstanceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceRegionId)) {
		body["InstanceRegionId"] = request.InstanceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TagModels)) {
		body["TagModels"] = request.TagModels
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceGrantedToExpressConnectRouter"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeInstanceGrantedToExpressConnectRouterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeInstanceGrantedToExpressConnectRouterResponse{}
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
// Queries the network instances whose permissions are granted to an Express Connect router (ECR).
//
// @param request - DescribeInstanceGrantedToExpressConnectRouterRequest
//
// @return DescribeInstanceGrantedToExpressConnectRouterResponse
func (client *Client) DescribeInstanceGrantedToExpressConnectRouter(request *DescribeInstanceGrantedToExpressConnectRouterRequest) (_result *DescribeInstanceGrantedToExpressConnectRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceGrantedToExpressConnectRouterResponse{}
	_body, _err := client.DescribeInstanceGrantedToExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates a virtual border router (VBR) from an Express Connect router (ECR).
//
// Description:
//
// Before you call the **DetachExpressConnectRouterChildInstance*	- operation to uninstall a VBR from an ECR, make sure that the ECR is in the **Active*	- state.
//
// @param request - DetachExpressConnectRouterChildInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachExpressConnectRouterChildInstanceResponse
func (client *Client) DetachExpressConnectRouterChildInstanceWithOptions(request *DetachExpressConnectRouterChildInstanceRequest, runtime *util.RuntimeOptions) (_result *DetachExpressConnectRouterChildInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChildInstanceId)) {
		body["ChildInstanceId"] = request.ChildInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ChildInstanceType)) {
		body["ChildInstanceType"] = request.ChildInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachExpressConnectRouterChildInstance"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DetachExpressConnectRouterChildInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DetachExpressConnectRouterChildInstanceResponse{}
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
// Disassociates a virtual border router (VBR) from an Express Connect router (ECR).
//
// Description:
//
// Before you call the **DetachExpressConnectRouterChildInstance*	- operation to uninstall a VBR from an ECR, make sure that the ECR is in the **Active*	- state.
//
// @param request - DetachExpressConnectRouterChildInstanceRequest
//
// @return DetachExpressConnectRouterChildInstanceResponse
func (client *Client) DetachExpressConnectRouterChildInstance(request *DetachExpressConnectRouterChildInstanceRequest) (_result *DetachExpressConnectRouterChildInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.DetachExpressConnectRouterChildInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables route entries of an Express Connect router (ECR).
//
// @param request - DisableExpressConnectRouterRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableExpressConnectRouterRouteEntriesResponse
func (client *Client) DisableExpressConnectRouterRouteEntriesWithOptions(request *DisableExpressConnectRouterRouteEntriesRequest, runtime *util.RuntimeOptions) (_result *DisableExpressConnectRouterRouteEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationCidrBlock)) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.NexthopInstanceId)) {
		body["NexthopInstanceId"] = request.NexthopInstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableExpressConnectRouterRouteEntries"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DisableExpressConnectRouterRouteEntriesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DisableExpressConnectRouterRouteEntriesResponse{}
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
// Disables route entries of an Express Connect router (ECR).
//
// @param request - DisableExpressConnectRouterRouteEntriesRequest
//
// @return DisableExpressConnectRouterRouteEntriesResponse
func (client *Client) DisableExpressConnectRouterRouteEntries(request *DisableExpressConnectRouterRouteEntriesRequest) (_result *DisableExpressConnectRouterRouteEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.DisableExpressConnectRouterRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables route entries of an Express Connect router (ECR).
//
// @param request - EnableExpressConnectRouterRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableExpressConnectRouterRouteEntriesResponse
func (client *Client) EnableExpressConnectRouterRouteEntriesWithOptions(request *EnableExpressConnectRouterRouteEntriesRequest, runtime *util.RuntimeOptions) (_result *EnableExpressConnectRouterRouteEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationCidrBlock)) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.NexthopInstanceId)) {
		body["NexthopInstanceId"] = request.NexthopInstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableExpressConnectRouterRouteEntries"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &EnableExpressConnectRouterRouteEntriesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &EnableExpressConnectRouterRouteEntriesResponse{}
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
// Enables route entries of an Express Connect router (ECR).
//
// @param request - EnableExpressConnectRouterRouteEntriesRequest
//
// @return EnableExpressConnectRouterRouteEntriesResponse
func (client *Client) EnableExpressConnectRouterRouteEntries(request *EnableExpressConnectRouterRouteEntriesRequest) (_result *EnableExpressConnectRouterRouteEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.EnableExpressConnectRouterRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an Express Connect router (ECR) and disassociates the virtual private cloud (VPC), transit router (TR), and virtual border router (VBR) associated with the ECR.
//
// Description:
//
//   If you forcefully delete an ECR, all the resources associated with the ECR are disassociated at a time. Make sure that the disassociation does not affect the stability of your business.
//
// 	- You can delete only ECRs that are in the **Active*	- state.
//
// @param request - ForceDeleteExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ForceDeleteExpressConnectRouterResponse
func (client *Client) ForceDeleteExpressConnectRouterWithOptions(request *ForceDeleteExpressConnectRouterRequest, runtime *util.RuntimeOptions) (_result *ForceDeleteExpressConnectRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ForceDeleteExpressConnectRouter"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ForceDeleteExpressConnectRouterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ForceDeleteExpressConnectRouterResponse{}
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
// Deletes an Express Connect router (ECR) and disassociates the virtual private cloud (VPC), transit router (TR), and virtual border router (VBR) associated with the ECR.
//
// Description:
//
//   If you forcefully delete an ECR, all the resources associated with the ECR are disassociated at a time. Make sure that the disassociation does not affect the stability of your business.
//
// 	- You can delete only ECRs that are in the **Active*	- state.
//
// @param request - ForceDeleteExpressConnectRouterRequest
//
// @return ForceDeleteExpressConnectRouterResponse
func (client *Client) ForceDeleteExpressConnectRouter(request *ForceDeleteExpressConnectRouterRequest) (_result *ForceDeleteExpressConnectRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ForceDeleteExpressConnectRouterResponse{}
	_body, _err := client.ForceDeleteExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants permissions on a virtual private cloud (VPC) or a virtual border router (VBR) to an Express Connect router (ECR) of another account.
//
// Description:
//
// When you associate a network instance of another account with an ECR, you must grant permissions on the network instance to the ECR.
//
// @param request - GrantInstanceToExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantInstanceToExpressConnectRouterResponse
func (client *Client) GrantInstanceToExpressConnectRouterWithOptions(request *GrantInstanceToExpressConnectRouterRequest, runtime *util.RuntimeOptions) (_result *GrantInstanceToExpressConnectRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.EcrOwnerAliUid)) {
		body["EcrOwnerAliUid"] = request.EcrOwnerAliUid
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceRegionId)) {
		body["InstanceRegionId"] = request.InstanceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["InstanceType"] = request.InstanceType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GrantInstanceToExpressConnectRouter"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GrantInstanceToExpressConnectRouterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GrantInstanceToExpressConnectRouterResponse{}
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
// Grants permissions on a virtual private cloud (VPC) or a virtual border router (VBR) to an Express Connect router (ECR) of another account.
//
// Description:
//
// When you associate a network instance of another account with an ECR, you must grant permissions on the network instance to the ECR.
//
// @param request - GrantInstanceToExpressConnectRouterRequest
//
// @return GrantInstanceToExpressConnectRouterResponse
func (client *Client) GrantInstanceToExpressConnectRouter(request *GrantInstanceToExpressConnectRouterRequest) (_result *GrantInstanceToExpressConnectRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantInstanceToExpressConnectRouterResponse{}
	_body, _err := client.GrantInstanceToExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of regions in which the Express Connect router (ECR) feature is activated.
//
// @param request - ListExpressConnectRouterSupportedRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExpressConnectRouterSupportedRegionResponse
func (client *Client) ListExpressConnectRouterSupportedRegionWithOptions(request *ListExpressConnectRouterSupportedRegionRequest, runtime *util.RuntimeOptions) (_result *ListExpressConnectRouterSupportedRegionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.NodeType)) {
		body["NodeType"] = request.NodeType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExpressConnectRouterSupportedRegion"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListExpressConnectRouterSupportedRegionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListExpressConnectRouterSupportedRegionResponse{}
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
// Queries a list of regions in which the Express Connect router (ECR) feature is activated.
//
// @param request - ListExpressConnectRouterSupportedRegionRequest
//
// @return ListExpressConnectRouterSupportedRegionResponse
func (client *Client) ListExpressConnectRouterSupportedRegion(request *ListExpressConnectRouterSupportedRegionRequest) (_result *ListExpressConnectRouterSupportedRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExpressConnectRouterSupportedRegionResponse{}
	_body, _err := client.ListExpressConnectRouterSupportedRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of tags that are added to an Express Connect router (ECR).
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2023-09-01"),
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
// Queries a list of tags that are added to an Express Connect router (ECR).
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
// Modifies the properties such as the name of an Express Connect router (ECR).
//
// Description:
//
// You can modify only properties of ECRs in the **Active*	- state.
//
// @param request - ModifyExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyExpressConnectRouterResponse
func (client *Client) ModifyExpressConnectRouterWithOptions(request *ModifyExpressConnectRouterRequest, runtime *util.RuntimeOptions) (_result *ModifyExpressConnectRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyExpressConnectRouter"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyExpressConnectRouterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyExpressConnectRouterResponse{}
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
// Modifies the properties such as the name of an Express Connect router (ECR).
//
// Description:
//
// You can modify only properties of ECRs in the **Active*	- state.
//
// @param request - ModifyExpressConnectRouterRequest
//
// @return ModifyExpressConnectRouterResponse
func (client *Client) ModifyExpressConnectRouter(request *ModifyExpressConnectRouterRequest) (_result *ModifyExpressConnectRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyExpressConnectRouterResponse{}
	_body, _err := client.ModifyExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the route prefixes of a virtual private cloud (VPC) or a transit router (TR) that is associated with an Express Connect router (ECR).
//
// @param request - ModifyExpressConnectRouterAssociationAllowedPrefixRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyExpressConnectRouterAssociationAllowedPrefixResponse
func (client *Client) ModifyExpressConnectRouterAssociationAllowedPrefixWithOptions(request *ModifyExpressConnectRouterAssociationAllowedPrefixRequest, runtime *util.RuntimeOptions) (_result *ModifyExpressConnectRouterAssociationAllowedPrefixResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowedPrefixes)) {
		body["AllowedPrefixes"] = request.AllowedPrefixes
	}

	if !tea.BoolValue(util.IsUnset(request.AllowedPrefixesMode)) {
		body["AllowedPrefixesMode"] = request.AllowedPrefixesMode
	}

	if !tea.BoolValue(util.IsUnset(request.AssociationId)) {
		body["AssociationId"] = request.AssociationId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		body["OwnerAccount"] = request.OwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyExpressConnectRouterAssociationAllowedPrefix"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyExpressConnectRouterAssociationAllowedPrefixResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyExpressConnectRouterAssociationAllowedPrefixResponse{}
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
// Modifies the route prefixes of a virtual private cloud (VPC) or a transit router (TR) that is associated with an Express Connect router (ECR).
//
// @param request - ModifyExpressConnectRouterAssociationAllowedPrefixRequest
//
// @return ModifyExpressConnectRouterAssociationAllowedPrefixResponse
func (client *Client) ModifyExpressConnectRouterAssociationAllowedPrefix(request *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) (_result *ModifyExpressConnectRouterAssociationAllowedPrefixResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyExpressConnectRouterAssociationAllowedPrefixResponse{}
	_body, _err := client.ModifyExpressConnectRouterAssociationAllowedPrefixWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the cross-region forwarding mode of an Express Connect router (ECR).
//
// @param request - ModifyExpressConnectRouterInterRegionTransitModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyExpressConnectRouterInterRegionTransitModeResponse
func (client *Client) ModifyExpressConnectRouterInterRegionTransitModeWithOptions(request *ModifyExpressConnectRouterInterRegionTransitModeRequest, runtime *util.RuntimeOptions) (_result *ModifyExpressConnectRouterInterRegionTransitModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.TransitModeList)) {
		body["TransitModeList"] = request.TransitModeList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyExpressConnectRouterInterRegionTransitMode"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyExpressConnectRouterInterRegionTransitModeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyExpressConnectRouterInterRegionTransitModeResponse{}
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
// Modifies the cross-region forwarding mode of an Express Connect router (ECR).
//
// @param request - ModifyExpressConnectRouterInterRegionTransitModeRequest
//
// @return ModifyExpressConnectRouterInterRegionTransitModeResponse
func (client *Client) ModifyExpressConnectRouterInterRegionTransitMode(request *ModifyExpressConnectRouterInterRegionTransitModeRequest) (_result *ModifyExpressConnectRouterInterRegionTransitModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyExpressConnectRouterInterRegionTransitModeResponse{}
	_body, _err := client.ModifyExpressConnectRouterInterRegionTransitModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改流日志
//
// @param request - ModifyFlowLogAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFlowLogAttributeResponse
func (client *Client) ModifyFlowLogAttributeWithOptions(request *ModifyFlowLogAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyFlowLogAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FlowLogId)) {
		query["FlowLogId"] = request.FlowLogId
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.SamplingRate)) {
		query["SamplingRate"] = request.SamplingRate
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.FlowLogName)) {
		body["FlowLogName"] = request.FlowLogName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFlowLogAttribute"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyFlowLogAttributeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyFlowLogAttributeResponse{}
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
// 修改流日志
//
// @param request - ModifyFlowLogAttributeRequest
//
// @return ModifyFlowLogAttributeResponse
func (client *Client) ModifyFlowLogAttribute(request *ModifyFlowLogAttributeRequest) (_result *ModifyFlowLogAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFlowLogAttributeResponse{}
	_body, _err := client.ModifyFlowLogAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the resource group to which an Express Connect router (ECR) belongs.
//
// @param request - MoveResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *util.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		body["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveResourceGroup"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &MoveResourceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &MoveResourceGroupResponse{}
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
// Modifies the resource group to which an Express Connect router (ECR) belongs.
//
// @param request - MoveResourceGroupRequest
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes permissions on a virtual private cloud (VPC) or a virtual border router (VBR) from an Express Connect router (ECR) owned by another account.
//
// @param request - RevokeInstanceFromExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeInstanceFromExpressConnectRouterResponse
func (client *Client) RevokeInstanceFromExpressConnectRouterWithOptions(request *RevokeInstanceFromExpressConnectRouterRequest, runtime *util.RuntimeOptions) (_result *RevokeInstanceFromExpressConnectRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.EcrOwnerAliUid)) {
		body["EcrOwnerAliUid"] = request.EcrOwnerAliUid
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceRegionId)) {
		body["InstanceRegionId"] = request.InstanceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["InstanceType"] = request.InstanceType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeInstanceFromExpressConnectRouter"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RevokeInstanceFromExpressConnectRouterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RevokeInstanceFromExpressConnectRouterResponse{}
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
// Revokes permissions on a virtual private cloud (VPC) or a virtual border router (VBR) from an Express Connect router (ECR) owned by another account.
//
// @param request - RevokeInstanceFromExpressConnectRouterRequest
//
// @return RevokeInstanceFromExpressConnectRouterResponse
func (client *Client) RevokeInstanceFromExpressConnectRouter(request *RevokeInstanceFromExpressConnectRouterRequest) (_result *RevokeInstanceFromExpressConnectRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeInstanceFromExpressConnectRouterResponse{}
	_body, _err := client.RevokeInstanceFromExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Synchronizes the forwarding bandwidth limit between regions for an Express Connect router (ECR).
//
// Description:
//
// Updates are allowed only when the ECR is in the **Active*	- state.
//
// @param request - SynchronizeExpressConnectRouterInterRegionBandwidthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SynchronizeExpressConnectRouterInterRegionBandwidthResponse
func (client *Client) SynchronizeExpressConnectRouterInterRegionBandwidthWithOptions(request *SynchronizeExpressConnectRouterInterRegionBandwidthRequest, runtime *util.RuntimeOptions) (_result *SynchronizeExpressConnectRouterInterRegionBandwidthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SynchronizeExpressConnectRouterInterRegionBandwidth"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SynchronizeExpressConnectRouterInterRegionBandwidthResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SynchronizeExpressConnectRouterInterRegionBandwidthResponse{}
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
// Synchronizes the forwarding bandwidth limit between regions for an Express Connect router (ECR).
//
// Description:
//
// Updates are allowed only when the ECR is in the **Active*	- state.
//
// @param request - SynchronizeExpressConnectRouterInterRegionBandwidthRequest
//
// @return SynchronizeExpressConnectRouterInterRegionBandwidthResponse
func (client *Client) SynchronizeExpressConnectRouterInterRegionBandwidth(request *SynchronizeExpressConnectRouterInterRegionBandwidthRequest) (_result *SynchronizeExpressConnectRouterInterRegionBandwidthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SynchronizeExpressConnectRouterInterRegionBandwidthResponse{}
	_body, _err := client.SynchronizeExpressConnectRouterInterRegionBandwidthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tags to an Express Connect router (ECR). You can add tags to only one ECR each time you call this operation. You can add multiple tags at a time.
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2023-09-01"),
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
// Adds tags to an Express Connect router (ECR). You can add tags to only one ECR each time you call this operation. You can add multiple tags at a time.
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

// Summary:
//
// Removes tags from an Express Connect router (ECR).
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		body["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		body["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2023-09-01"),
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
// Removes tags from an Express Connect router (ECR).
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
