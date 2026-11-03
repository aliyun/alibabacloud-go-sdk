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
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirect.
	//
	// - **4xx**: Request error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response result.
	Data *DescribeNamespaceResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// - Successful request: The **ErrorCode*	- field is not returned.
	//
	// - Failed request: The **ErrorCode*	- field is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// InvalidParameter.Obviously
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information. Valid values:
	//
	// - **success**: The request is normal.
	//
	// - Specific error code: The request is abnormal.
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
	// Indicates whether the namespace resource information was queried successfully. Valid values:
	//
	// - **true**: The query was successful.
	//
	// - **false**: The query failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID, which is used to accurately query call information.
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
	//
	// example:
	//
	// 2.5.0
	ApmJavaAgentVersion *string `json:"ApmJavaAgentVersion,omitempty" xml:"ApmJavaAgentVersion,omitempty"`
	// The number of applications.
	//
	// example:
	//
	// 1
	AppCount *int64 `json:"AppCount,omitempty" xml:"AppCount,omitempty"`
	// The region to which the namespace belongs.
	//
	// example:
	//
	// cn-shanghai
	BelongRegion *string `json:"BelongRegion,omitempty" xml:"BelongRegion,omitempty"`
	// The namespace description.
	//
	// example:
	//
	// decs
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The jump server application ID.
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
	// The change order ID.
	//
	// example:
	//
	// afedb3c4-63f8-4a3d-aaa3-2c30363f****
	LastChangeOrderId *string `json:"LastChangeOrderId,omitempty" xml:"LastChangeOrderId,omitempty"`
	// Indicates whether a change order is running in the namespace. Valid values:
	//
	// - **true**: A change order is running.
	//
	// - **false**: No change order is running.
	//
	// example:
	//
	// true
	LastChangeOrderRunning *bool `json:"LastChangeOrderRunning,omitempty" xml:"LastChangeOrderRunning,omitempty"`
	// The status of the latest change order. Valid values:
	//
	//   - **READY**: The latest change order is ready.
	//
	//   - **RUNNING**: The latest change order is running.
	//
	//   - **SUCCESS**: The latest change order is successful.
	//
	//   - **FAIL**: The latest change order failed.
	//
	//   - **ABORT**: The latest change order is aborted.
	//
	//   - **WAIT_BATCH_CONFIRM**: The latest change order is waiting for manual batch confirmation.
	//
	//   - **AUTO_BATCH_WAIT**: The latest change order is in the automatic batch waiting state.
	//
	//   - **SYSTEM_FAIL**: A system fault occurred.
	//
	//   - **WAIT_APPROVAL**: The latest change order is waiting for approval.
	//
	//   - **APPROVED**: The latest change order is approved and waiting to be executed.
	//
	// example:
	//
	// SUCCESS
	LastChangeOrderStatus *string `json:"LastChangeOrderStatus,omitempty" xml:"LastChangeOrderStatus,omitempty"`
	// The short namespace ID.
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
	// Indicates whether the change order notification has expired. Valid values:
	//
	// - **true**: The change order notification has expired.
	//
	// - **false**: The change order notification has not expired.
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
	// The configurations for collecting logs to Simple Log Service (SLS).
	//
	// - To use SLS resources through automatic creation by SAE: `[{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]`.
	//
	// - To use custom SLS resources: `[{"projectName":"test-sls","logType":"stdout","logDir":"","logstoreName":"sae","logtailName":""},{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]`.
	//
	// Metric description:
	//
	// - **projectName**: The name of the SLS project.
	//
	// - **logDir**: The log path.
	//
	// - **logType**: The log type. **stdout*	- indicates container standard output logs. You can configure only one such log in the Settings. If you do not configure it, file logs are collected.
	//
	// - **logstoreName**: The name of the SLS Logstore.
	//
	// - **logtailName**: The name of the SLS Logtail. If this parameter is not specified, a new Logtail is created.
	//
	// If the SLS collection configurations remain unchanged during multiple deployments, you do not need to modify the Settings for this parameter (that is, the **SlsConfigs*	- field is not required in the request). If you no longer need to use the SLS collection feature, you must change the Settings of this field to an empty character string in the request (that is, the value of the **SlsConfigs*	- field is "").
	//
	// example:
	//
	// [{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]
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
