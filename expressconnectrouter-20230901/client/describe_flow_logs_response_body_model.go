// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeFlowLogsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DescribeFlowLogsResponseBody
	GetCode() *string
	SetDynamicCode(v string) *DescribeFlowLogsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeFlowLogsResponseBody
	GetDynamicMessage() *string
	SetFlowLogs(v []*DescribeFlowLogsResponseBodyFlowLogs) *DescribeFlowLogsResponseBody
	GetFlowLogs() []*DescribeFlowLogsResponseBodyFlowLogs
	SetHttpStatusCode(v int32) *DescribeFlowLogsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *DescribeFlowLogsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *DescribeFlowLogsResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeFlowLogsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeFlowLogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeFlowLogsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *DescribeFlowLogsResponseBody
	GetTotalCount() *int32
}

type DescribeFlowLogsResponseBody struct {
	// The details of the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code that indicates whether the call was successful. A value of 200 indicates that the call was successful. Other values indicate that the call failed. For more information about the failure causes, see the error codes.
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
	// The dynamic error message, which is used to replace `%s` in the **ErrMessage*	- response parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the specified **DtsJobId*	- request parameter is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The list of flow logs.
	FlowLogs []*DescribeFlowLogsResponseBodyFlowLogs `json:"FlowLogs,omitempty" xml:"FlowLogs,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The maximum number of entries returned. Valid values: 1 to 2147483647. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The message returned by the operation.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The token for the next query. Valid values:
	//
	// - If **NextToken*	- is empty, no next query is available.
	//
	// - If **NextToken*	- has a return value, the value indicates the token for the next query.
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
	// Indicates whether the call was successful. Valid values:
	//
	// - **True**: The call was successful.
	//
	// - **False**: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries that meet the query conditions.
	//
	// example:
	//
	// 11
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFlowLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeFlowLogsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeFlowLogsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeFlowLogsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeFlowLogsResponseBody) GetFlowLogs() []*DescribeFlowLogsResponseBodyFlowLogs {
	return s.FlowLogs
}

func (s *DescribeFlowLogsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeFlowLogsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeFlowLogsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeFlowLogsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeFlowLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFlowLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeFlowLogsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
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

func (s *DescribeFlowLogsResponseBody) Validate() error {
	if s.FlowLogs != nil {
		for _, item := range s.FlowLogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFlowLogsResponseBodyFlowLogs struct {
	// The time when the flow log was created. The time is displayed in UTC and follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format.
	//
	// example:
	//
	// 2023-09-21T04:20Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the Express Connect Router (ECR) instance.
	//
	// example:
	//
	// myFlowlog
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID of the Express Connect Router (ECR).
	//
	// example:
	//
	// ecr-h4cop1khw98*****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The instance ID of the flow log instance.
	//
	// example:
	//
	// ecr-flog-leypqehtgtia2*****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	// The name of the flow log.
	//
	// example:
	//
	// myFlowlog
	FlowLogName *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	// The ID of the authorized instance.
	//
	// example:
	//
	// vbr-9dpty76irpf4u15*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// - **VBR**: virtual border router (VBR) instance.
	//
	// example:
	//
	// VBR
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The capture window duration for the traffic captured by the flow log. Unit: seconds. Valid values:
	//
	// - **60*	-
	//
	// - **600**
	//
	// Default value: **600**.
	//
	// example:
	//
	// 600
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Logstore that stores the captured traffic.
	//
	// example:
	//
	// FlowLogStore
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The name of the Simple Log Service project that stores the captured traffic.
	//
	// example:
	//
	// FlowLogProject
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The region ID of the Express Connect Router.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Express Connect Router (ECR) instance belongs.
	//
	// example:
	//
	// rg-aekzb3xwrexc4ry
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The sampling ratio. Valid values:
	//
	// - **1:4096**
	//
	// - **1:2048**
	//
	// - **1:1024**
	//
	// Default value: **1:4096**.
	//
	// example:
	//
	// 1:4096
	SamplingRate *string `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
	// The region ID of Simple Log Service.
	//
	// example:
	//
	// cn-hangzhou
	SlsRegionId *string `json:"SlsRegionId,omitempty" xml:"SlsRegionId,omitempty"`
	// The status of the flow log. Valid values:
	//
	// - **Active**: The flow log is started.
	//
	// - **Inactive**: The flow log is not started.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the Express Connect Router (ECR) instance.
	Tags []*DescribeFlowLogsResponseBodyFlowLogsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeFlowLogsResponseBodyFlowLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogsResponseBodyFlowLogs) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) GetDescription() *string {
	return s.Description
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) GetEcrId() *string {
	return s.EcrId
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) GetFlowLogName() *string {
	return s.FlowLogName
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) GetSamplingRate() *string {
	return s.SamplingRate
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) GetSlsRegionId() *string {
	return s.SlsRegionId
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) GetStatus() *string {
	return s.Status
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) GetTags() []*DescribeFlowLogsResponseBodyFlowLogsTags {
	return s.Tags
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

func (s *DescribeFlowLogsResponseBodyFlowLogs) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFlowLogsResponseBodyFlowLogsTags struct {
	// The tag key of the instance. If you specify this parameter, it cannot be an empty string.
	//
	// > The tag key can be up to 64 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// > The tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`. If you specify this parameter, it can be an empty string.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFlowLogsResponseBodyFlowLogsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogsResponseBodyFlowLogsTags) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsResponseBodyFlowLogsTags) GetKey() *string {
	return s.Key
}

func (s *DescribeFlowLogsResponseBodyFlowLogsTags) GetValue() *string {
	return s.Value
}

func (s *DescribeFlowLogsResponseBodyFlowLogsTags) SetKey(v string) *DescribeFlowLogsResponseBodyFlowLogsTags {
	s.Key = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsTags) SetValue(v string) *DescribeFlowLogsResponseBodyFlowLogsTags {
	s.Value = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsTags) Validate() error {
	return dara.Validate(s)
}
