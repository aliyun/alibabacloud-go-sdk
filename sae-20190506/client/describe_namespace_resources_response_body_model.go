// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNamespaceResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeNamespaceResourcesResponseBody
	GetCode() *string
	SetData(v *DescribeNamespaceResourcesResponseBodyData) *DescribeNamespaceResourcesResponseBody
	GetData() *DescribeNamespaceResourcesResponseBodyData
	SetErrorCode(v string) *DescribeNamespaceResourcesResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeNamespaceResourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeNamespaceResourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeNamespaceResourcesResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeNamespaceResourcesResponseBody
	GetTraceId() *string
}

type DescribeNamespaceResourcesResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// - **2xx**: The request was successful.
	//
	// - **3xx**: The request was redirected.
	//
	// - **4xx**: The request was invalid.
	//
	// - **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *DescribeNamespaceResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. This parameter is returned only if the request fails. For more information, see the **Error codes*	- section of this topic.
	//
	// - A successful request does not return the **ErrorCode*	- field.
	//
	// - A failed request returns the **ErrorCode*	- field. For more information, see the list of **error codes*	- in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message. Valid values:
	//
	// - If the request is successful, **success*	- is returned.
	//
	// - If the request fails, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the namespace resources were queried successfully.
	//
	// - **true**: The query was successful.
	//
	// - **false**: The query failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID. You can use this ID to query the details of the call.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeNamespaceResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceResourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeNamespaceResourcesResponseBody) GetData() *DescribeNamespaceResourcesResponseBodyData {
	return s.Data
}

func (s *DescribeNamespaceResourcesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeNamespaceResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeNamespaceResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNamespaceResourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeNamespaceResourcesResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeNamespaceResourcesResponseBody) SetCode(v string) *DescribeNamespaceResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBody) SetData(v *DescribeNamespaceResourcesResponseBodyData) *DescribeNamespaceResourcesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeNamespaceResourcesResponseBody) SetErrorCode(v string) *DescribeNamespaceResourcesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBody) SetMessage(v string) *DescribeNamespaceResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBody) SetRequestId(v string) *DescribeNamespaceResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBody) SetSuccess(v bool) *DescribeNamespaceResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBody) SetTraceId(v string) *DescribeNamespaceResourcesResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNamespaceResourcesResponseBodyData struct {
	// The version of the APM Java agent.
	ApmJavaAgentVersion *string `json:"ApmJavaAgentVersion,omitempty" xml:"ApmJavaAgentVersion,omitempty"`
	// The number of applications.
	//
	// example:
	//
	// 1
	AppCount *int64 `json:"AppCount,omitempty" xml:"AppCount,omitempty"`
	// The region of the namespace.
	//
	// example:
	//
	// cn-shanghai
	BelongRegion *string `json:"BelongRegion,omitempty" xml:"BelongRegion,omitempty"`
	// The description of the namespace.
	//
	// example:
	//
	// decs
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the jump server application.
	//
	// example:
	//
	// 5ec46468-6b26-4a3c-9f7c-bf88761a****
	JumpServerAppId *string `json:"JumpServerAppId,omitempty" xml:"JumpServerAppId,omitempty"`
	// The IP address of the jump server.
	//
	// example:
	//
	// 120.78.XXX.XX
	JumpServerIp *string `json:"JumpServerIp,omitempty" xml:"JumpServerIp,omitempty"`
	// The release order ID.
	//
	// example:
	//
	// afedb3c4-63f8-4a3d-aaa3-2c30363f****
	LastChangeOrderId *string `json:"LastChangeOrderId,omitempty" xml:"LastChangeOrderId,omitempty"`
	// Indicates whether a release order is running in the namespace. Valid values:
	//
	// - **true**: A release order is running.
	//
	// - **false**: No release order is running.
	//
	// example:
	//
	// true
	LastChangeOrderRunning *bool `json:"LastChangeOrderRunning,omitempty" xml:"LastChangeOrderRunning,omitempty"`
	// The status of the last release order. Valid values:
	//
	// - **READY**: The release order is ready.
	//
	// - **RUNNING**: The release order is running.
	//
	// - **SUCCESS**: The release order was successful.
	//
	// - **FAIL**: The release order failed.
	//
	// - **ABORT**: The release order was aborted.
	//
	// - **WAIT_BATCH_CONFIRM**: The release order is waiting for manual batch confirmation.
	//
	// - **AUTO_BATCH_WAIT**: The release order is in an automatic batch-wait state.
	//
	// - **SYSTEM_FAIL**: A system error occurred.
	//
	// - **WAIT_APPROVAL**: The release order is pending approval.
	//
	// - **APPROVED**: The release order is approved and pending execution.
	//
	// example:
	//
	// SUCCESS
	LastChangeOrderStatus *string `json:"LastChangeOrderStatus,omitempty" xml:"LastChangeOrderStatus,omitempty"`
	// The short-format namespace ID.
	//
	// example:
	//
	// test
	NameSpaceShortId *string `json:"NameSpaceShortId,omitempty" xml:"NameSpaceShortId,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// cn-shangha:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The namespace name.
	//
	// example:
	//
	// test
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// Indicates whether the notification for the release order has expired. Valid values:
	//
	// - **true**: The notification has expired.
	//
	// - **false**: The notification has not expired.
	//
	// example:
	//
	// true
	NotificationExpired *bool `json:"NotificationExpired,omitempty" xml:"NotificationExpired,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-wz969ngg2e49q5i4****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The configuration for collecting logs to Simple Log Service (SLS).
	//
	// - To use SLS resources that are automatically created by SAE: `[{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]`.
	//
	// - To use custom SLS resources: `[{"projectName":"test-sls","logType":"stdout","logDir":"","logstoreName":"sae","logtailName":""},{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]`.
	//
	// The configuration includes the following parameters:
	//
	// - **projectName**: The project name in SLS.
	//
	// - **logDir**: The log path.
	//
	// - **logType**: The log type. **stdout*	- indicates the standard output of the container. Only one stdout entry is allowed. If you omit this parameter, file logs are collected.
	//
	// - **logstoreName**: The Logstore name in SLS.
	//
	// - **logtailName**: The Logtail name in SLS. If you omit this parameter, SAE creates a Logtail.
	//
	// You do not need to set this parameter if the SLS log collection configuration is unchanged for subsequent deployments. To disable log collection, set this parameter to an empty string ("").
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	// The tenant ID of the SAE namespace.
	//
	// example:
	//
	// 838cad95-973f-48fe-830b-2a8546d7****
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// test@aliyun.com
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-2ze559r1z1bpwqxwp****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The vSwitch name.
	//
	// example:
	//
	// test
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-2ze0i263cnn311nvj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VPC name.
	//
	// example:
	//
	// test
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeNamespaceResourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceResourcesResponseBodyData) GetApmJavaAgentVersion() *string {
	return s.ApmJavaAgentVersion
}

func (s *DescribeNamespaceResourcesResponseBodyData) GetAppCount() *int64 {
	return s.AppCount
}

func (s *DescribeNamespaceResourcesResponseBodyData) GetBelongRegion() *string {
	return s.BelongRegion
}

func (s *DescribeNamespaceResourcesResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeNamespaceResourcesResponseBodyData) GetJumpServerAppId() *string {
	return s.JumpServerAppId
}

func (s *DescribeNamespaceResourcesResponseBodyData) GetJumpServerIp() *string {
	return s.JumpServerIp
}

func (s *DescribeNamespaceResourcesResponseBodyData) GetLastChangeOrderId() *string {
	return s.LastChangeOrderId
}

func (s *DescribeNamespaceResourcesResponseBodyData) GetLastChangeOrderRunning() *bool {
	return s.LastChangeOrderRunning
}

func (s *DescribeNamespaceResourcesResponseBodyData) GetLastChangeOrderStatus() *string {
	return s.LastChangeOrderStatus
}

func (s *DescribeNamespaceResourcesResponseBodyData) GetNameSpaceShortId() *string {
	return s.NameSpaceShortId
}

func (s *DescribeNamespaceResourcesResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeNamespaceResourcesResponseBodyData) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *DescribeNamespaceResourcesResponseBodyData) GetNotificationExpired() *bool {
	return s.NotificationExpired
}

func (s *DescribeNamespaceResourcesResponseBodyData) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeNamespaceResourcesResponseBodyData) GetSlsConfigs() *string {
	return s.SlsConfigs
}

func (s *DescribeNamespaceResourcesResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeNamespaceResourcesResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *DescribeNamespaceResourcesResponseBodyData) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeNamespaceResourcesResponseBodyData) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeNamespaceResourcesResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeNamespaceResourcesResponseBodyData) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetApmJavaAgentVersion(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.ApmJavaAgentVersion = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetAppCount(v int64) *DescribeNamespaceResourcesResponseBodyData {
	s.AppCount = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetBelongRegion(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.BelongRegion = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetDescription(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetJumpServerAppId(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.JumpServerAppId = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetJumpServerIp(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.JumpServerIp = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetLastChangeOrderId(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.LastChangeOrderId = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetLastChangeOrderRunning(v bool) *DescribeNamespaceResourcesResponseBodyData {
	s.LastChangeOrderRunning = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetLastChangeOrderStatus(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.LastChangeOrderStatus = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetNameSpaceShortId(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.NameSpaceShortId = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetNamespaceId(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetNamespaceName(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.NamespaceName = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetNotificationExpired(v bool) *DescribeNamespaceResourcesResponseBodyData {
	s.NotificationExpired = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetSecurityGroupId(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetSlsConfigs(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.SlsConfigs = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetTenantId(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetUserId(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.UserId = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetVSwitchId(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.VSwitchId = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetVSwitchName(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.VSwitchName = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetVpcId(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetVpcName(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.VpcName = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
