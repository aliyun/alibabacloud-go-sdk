// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAliYunSafeCenterResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateSimilarSecurityEventsQueryTaskRequest(v *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest) *GetAliYunSafeCenterResultRequest
	GetCreateSimilarSecurityEventsQueryTaskRequest() *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest
	SetDescribeInstancesFullStatusRequest(v *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest) *GetAliYunSafeCenterResultRequest
	GetDescribeInstancesFullStatusRequest() *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest
	SetDescribeSecurityEventOperationStatusRequest(v *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest) *GetAliYunSafeCenterResultRequest
	GetDescribeSecurityEventOperationStatusRequest() *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest
	SetDescribeSimilarSecurityEventsRequest(v *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest) *GetAliYunSafeCenterResultRequest
	GetDescribeSimilarSecurityEventsRequest() *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest
	SetGetAssetDetailByUuidRequest(v *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest) *GetAliYunSafeCenterResultRequest
	GetGetAssetDetailByUuidRequest() *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest
	SetHandleSecurityEventsRequest(v *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) *GetAliYunSafeCenterResultRequest
	GetHandleSecurityEventsRequest() *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest
	SetHandleSimilarSecurityEventsRequest(v *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) *GetAliYunSafeCenterResultRequest
	GetHandleSimilarSecurityEventsRequest() *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest
	SetInterfaceCode(v string) *GetAliYunSafeCenterResultRequest
	GetInterfaceCode() *string
	SetListInstancesRequest(v *GetAliYunSafeCenterResultRequestListInstancesRequest) *GetAliYunSafeCenterResultRequest
	GetListInstancesRequest() *GetAliYunSafeCenterResultRequestListInstancesRequest
	SetRegionId(v string) *GetAliYunSafeCenterResultRequest
	GetRegionId() *string
}

type GetAliYunSafeCenterResultRequest struct {
	// Creates a node to query security alerting events triggered by the same rule or alerting type.
	CreateSimilarSecurityEventsQueryTaskRequest *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest `json:"CreateSimilarSecurityEventsQueryTaskRequest,omitempty" xml:"CreateSimilarSecurityEventsQueryTaskRequest,omitempty" type:"Struct"`
	// Queries the running status of ECS instances.
	DescribeInstancesFullStatusRequest *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest `json:"DescribeInstancesFullStatusRequest,omitempty" xml:"DescribeInstancesFullStatusRequest,omitempty" type:"Struct"`
	// Queries whether the list of security alerting events that match the same IP rule or same alerting type as the alerting event to be handled is empty.
	DescribeSecurityEventOperationStatusRequest *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest `json:"DescribeSecurityEventOperationStatusRequest,omitempty" xml:"DescribeSecurityEventOperationStatusRequest,omitempty" type:"Struct"`
	// Queries identical security alert events in Security Center.
	DescribeSimilarSecurityEventsRequest *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest `json:"DescribeSimilarSecurityEventsRequest,omitempty" xml:"DescribeSimilarSecurityEventsRequest,omitempty" type:"Struct"`
	// The request parameters for querying the Security Center Agent status.
	GetAssetDetailByUuidRequest *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest `json:"GetAssetDetailByUuidRequest,omitempty" xml:"GetAssetDetailByUuidRequest,omitempty" type:"Struct"`
	// Handles security alert events.
	HandleSecurityEventsRequest *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest `json:"HandleSecurityEventsRequest,omitempty" xml:"HandleSecurityEventsRequest,omitempty" type:"Struct"`
	// Handles security alert events in batches based on the same IP rule or type.
	HandleSimilarSecurityEventsRequest *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest `json:"HandleSimilarSecurityEventsRequest,omitempty" xml:"HandleSimilarSecurityEventsRequest,omitempty" type:"Struct"`
	// The code of the public API operation.
	//
	// - **GetAssetDetailByUuid**: Retrieves the Agent status. Request parameter: GetAssetDetailByUuidRequest.
	//
	// - **DescribeSimilarSecurityEvents**: Retrieves the list of instance IDs for identical security alerting events. Request parameter: DescribeSimilarSecurityEventsRequest.
	//
	// - **CreateSimilarSecurityEventsQueryTask**: Creates a node to query security alerting events triggered by the same rule or alerting type. Request parameter: CreateSimilarSecurityEventsQueryTaskRequest.
	//
	// - **DescribeSecurityEventOperationStatus**: Queries whether the list of security alerting events that match the same IP rule or same alerting type as the alerting event to be handled is empty. Request parameter: DescribeSecurityEventOperationStatusRequest.
	//
	// - **HandleSimilarSecurityEvents**: Handles security alerting events in batches based on the same IP rule or type. Request parameter: HandleSimilarSecurityEventsRequest.
	//
	// HandleSecurityEvents: Handles security alerting events. Request parameter: HandleSecurityEventsRequest.
	//
	// - **DescribeInstancesFullStatus**: Queries the running status of ECS instances. Request parameter: DescribeInstancesFullStatusRequest.
	//
	// - **ListInstances**: Queries the running status of simple application servers. Request parameter: ListInstancesRequest.
	//
	// - **StartConfigRuleEvaluation**: Re-evaluates security check rules.
	//
	// > Each API operation name corresponds to its own request parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// ListInstanceStatus
	InterfaceCode *string `json:"InterfaceCode,omitempty" xml:"InterfaceCode,omitempty"`
	// Queries the running status of simple application servers.
	ListInstancesRequest *GetAliYunSafeCenterResultRequestListInstancesRequest `json:"ListInstancesRequest,omitempty" xml:"ListInstancesRequest,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-guangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAliYunSafeCenterResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultRequest) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultRequest) GetCreateSimilarSecurityEventsQueryTaskRequest() *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest {
	return s.CreateSimilarSecurityEventsQueryTaskRequest
}

func (s *GetAliYunSafeCenterResultRequest) GetDescribeInstancesFullStatusRequest() *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest {
	return s.DescribeInstancesFullStatusRequest
}

func (s *GetAliYunSafeCenterResultRequest) GetDescribeSecurityEventOperationStatusRequest() *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest {
	return s.DescribeSecurityEventOperationStatusRequest
}

func (s *GetAliYunSafeCenterResultRequest) GetDescribeSimilarSecurityEventsRequest() *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest {
	return s.DescribeSimilarSecurityEventsRequest
}

func (s *GetAliYunSafeCenterResultRequest) GetGetAssetDetailByUuidRequest() *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest {
	return s.GetAssetDetailByUuidRequest
}

func (s *GetAliYunSafeCenterResultRequest) GetHandleSecurityEventsRequest() *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	return s.HandleSecurityEventsRequest
}

func (s *GetAliYunSafeCenterResultRequest) GetHandleSimilarSecurityEventsRequest() *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest {
	return s.HandleSimilarSecurityEventsRequest
}

func (s *GetAliYunSafeCenterResultRequest) GetInterfaceCode() *string {
	return s.InterfaceCode
}

func (s *GetAliYunSafeCenterResultRequest) GetListInstancesRequest() *GetAliYunSafeCenterResultRequestListInstancesRequest {
	return s.ListInstancesRequest
}

func (s *GetAliYunSafeCenterResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAliYunSafeCenterResultRequest) SetCreateSimilarSecurityEventsQueryTaskRequest(v *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest) *GetAliYunSafeCenterResultRequest {
	s.CreateSimilarSecurityEventsQueryTaskRequest = v
	return s
}

func (s *GetAliYunSafeCenterResultRequest) SetDescribeInstancesFullStatusRequest(v *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest) *GetAliYunSafeCenterResultRequest {
	s.DescribeInstancesFullStatusRequest = v
	return s
}

func (s *GetAliYunSafeCenterResultRequest) SetDescribeSecurityEventOperationStatusRequest(v *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest) *GetAliYunSafeCenterResultRequest {
	s.DescribeSecurityEventOperationStatusRequest = v
	return s
}

func (s *GetAliYunSafeCenterResultRequest) SetDescribeSimilarSecurityEventsRequest(v *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest) *GetAliYunSafeCenterResultRequest {
	s.DescribeSimilarSecurityEventsRequest = v
	return s
}

func (s *GetAliYunSafeCenterResultRequest) SetGetAssetDetailByUuidRequest(v *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest) *GetAliYunSafeCenterResultRequest {
	s.GetAssetDetailByUuidRequest = v
	return s
}

func (s *GetAliYunSafeCenterResultRequest) SetHandleSecurityEventsRequest(v *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) *GetAliYunSafeCenterResultRequest {
	s.HandleSecurityEventsRequest = v
	return s
}

func (s *GetAliYunSafeCenterResultRequest) SetHandleSimilarSecurityEventsRequest(v *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) *GetAliYunSafeCenterResultRequest {
	s.HandleSimilarSecurityEventsRequest = v
	return s
}

func (s *GetAliYunSafeCenterResultRequest) SetInterfaceCode(v string) *GetAliYunSafeCenterResultRequest {
	s.InterfaceCode = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequest) SetListInstancesRequest(v *GetAliYunSafeCenterResultRequestListInstancesRequest) *GetAliYunSafeCenterResultRequest {
	s.ListInstancesRequest = v
	return s
}

func (s *GetAliYunSafeCenterResultRequest) SetRegionId(v string) *GetAliYunSafeCenterResultRequest {
	s.RegionId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequest) Validate() error {
	if s.CreateSimilarSecurityEventsQueryTaskRequest != nil {
		if err := s.CreateSimilarSecurityEventsQueryTaskRequest.Validate(); err != nil {
			return err
		}
	}
	if s.DescribeInstancesFullStatusRequest != nil {
		if err := s.DescribeInstancesFullStatusRequest.Validate(); err != nil {
			return err
		}
	}
	if s.DescribeSecurityEventOperationStatusRequest != nil {
		if err := s.DescribeSecurityEventOperationStatusRequest.Validate(); err != nil {
			return err
		}
	}
	if s.DescribeSimilarSecurityEventsRequest != nil {
		if err := s.DescribeSimilarSecurityEventsRequest.Validate(); err != nil {
			return err
		}
	}
	if s.GetAssetDetailByUuidRequest != nil {
		if err := s.GetAssetDetailByUuidRequest.Validate(); err != nil {
			return err
		}
	}
	if s.HandleSecurityEventsRequest != nil {
		if err := s.HandleSecurityEventsRequest.Validate(); err != nil {
			return err
		}
	}
	if s.HandleSimilarSecurityEventsRequest != nil {
		if err := s.HandleSimilarSecurityEventsRequest.Validate(); err != nil {
			return err
		}
	}
	if s.ListInstancesRequest != nil {
		if err := s.ListInstancesRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security alert event.
	//
	// example:
	//
	// 629755508
	SecurityEventId *int64 `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
	// The code of the alerting event that has the same type or rule hits.
	//
	// example:
	//
	// default
	SimilarEventScenarioCode *string `json:"SimilarEventScenarioCode,omitempty" xml:"SimilarEventScenarioCode,omitempty"`
}

func (s GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest) GetSecurityEventId() *int64 {
	return s.SecurityEventId
}

func (s *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest) GetSimilarEventScenarioCode() *string {
	return s.SimilarEventScenarioCode
}

func (s *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest) SetRegionId(v string) *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest {
	s.RegionId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest) SetSecurityEventId(v int64) *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest {
	s.SecurityEventId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest) SetSimilarEventScenarioCode(v string) *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest {
	s.SimilarEventScenarioCode = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest) Validate() error {
	return dara.Validate(s)
}

type GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest struct {
	// The list of instance IDs.
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest) SetInstanceId(v []*string) *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest {
	s.InstanceId = v
	return s
}

func (s *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest) SetRegionId(v string) *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest) Validate() error {
	return dara.Validate(s)
}

type GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest struct {
	// The region ID. Example: ap-southeast-1.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of security alert event IDs.
	//
	// > You must specify either TaskId or SecurityEventIds.N. At least one of these parameters is required for a successful call.
	SecurityEventIds []*string `json:"SecurityEventIds,omitempty" xml:"SecurityEventIds,omitempty" type:"Repeated"`
	// The ID of the task for handling security alert events.
	//
	// > You must specify either TaskId or SecurityEventIds. At least one of these parameters is required for a successful call.
	//
	// example:
	//
	// 0BC3B4E600002A9F000048BCDCE7E710
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest) GetSecurityEventIds() []*string {
	return s.SecurityEventIds
}

func (s *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest) SetRegionId(v string) *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest) SetSecurityEventIds(v []*string) *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest {
	s.SecurityEventIds = v
	return s
}

func (s *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest) SetTaskId(v int64) *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest) Validate() error {
	return dara.Validate(s)
}

type GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the query task. You can call the CreateSimilarSecurityEventsQueryTask operation to obtain this parameter.
	//
	// example:
	//
	// 1689135
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest) SetRegionId(v string) *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest {
	s.RegionId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest) SetTaskId(v int64) *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest {
	s.TaskId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest) Validate() error {
	return dara.Validate(s)
}

type GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The UUID of the asset to query.
	//
	// example:
	//
	// 9A75F21D3993C0A2B094A4AB132890B2
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest) SetRegionId(v string) *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest {
	s.RegionId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest) SetUuid(v string) *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest {
	s.Uuid = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest) Validate() error {
	return dara.Validate(s)
}

type GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest struct {
	// The alert rule type.
	//
	// example:
	//
	// default
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The MD5 hash of the file.
	//
	// example:
	//
	// fbbb90731fbb6df57c933173182d01a5
	FileMd5 *string `json:"FileMd5,omitempty" xml:"FileMd5,omitempty"`
	// The path of the sensitive file.
	//
	// example:
	//
	// flyfish-lfp-wy.release
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// api-shared-vpc-002
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The user IP address.
	//
	// example:
	//
	// 123.56.127.180
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// Specifies whether to add to the whitelist in batches.
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	MarkBatch *string `json:"MarkBatch,omitempty" xml:"MarkBatch,omitempty"`
	// The whitelist rule configuration. The value is in JSON format and contains the following fields:
	//
	// - **field**: The whitelist field.
	//
	// - **operate**: The whitelist method. Valid values:
	//
	//   - **notContains**: Does not contain.
	//
	//   - **contains**: Contains.
	//
	//   - **regex**: Regular expression match.
	//
	//   - **strEqual**: Equals.
	//
	//   - **strNotEqual**: Does not equal.
	//
	// - **fieldValue**: The match value.
	//
	// - **uuid**: The scope of the whitelist rule. Valid values:
	//
	//   - **part**: Only the current asset.
	//
	//   - **ALL**: All assets.
	//
	// > Call the DescribeSecurityEventOperations operation to obtain the field whitelist field.
	//
	// example:
	//
	// [{"uuid":"part","field":"gmtModified","operate":"contains","fieldValue":"asd"},{"uuid":"part","field":"loginUser","operate":"contains","fieldValue":"vff"}]
	MarkMissParam *string `json:"MarkMissParam,omitempty" xml:"MarkMissParam,omitempty"`
	// The method for handling the security alert event. Valid values:
	//
	// - **block_ip**: Block.
	//
	// - **advance_mark_mis_info**: Add to whitelist.
	//
	// - **ignore**: Ignore.
	//
	// - **manual_handled**: Manually handled.
	//
	// - **kill_process**: Terminate process.
	//
	// - **cleanup**: Deep scan and cleanup.
	//
	// - **kill_and_quara**: Virus scan and quarantine.
	//
	// - **disable_malicious_defense**: Disable malicious behavior defense.
	//
	// - **client_problem_check**: Troubleshoot.
	//
	// - **quara**: Quarantine.
	//
	// example:
	//
	// block_ip
	OperationCode *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	// The configuration of the sub-operation for handling security alert events.
	//
	// When OperationCode is set to kill_and_quara, specify the parameter type "subOperation":${code}.
	//
	// Valid code values:
	//
	// - Quarantined file: quaraFileByMd5andPath
	//
	// - Kill process and quarantined file by process ID and path: killAndQuaraFileByPidAndMd5andPath
	//
	// - Kill process only: killByMd5andPath
	//
	// - Kill process and quarantined file: killAndQuaraFileByMd5andPath
	//
	// - Kill container process by process ID and path: killProcessByPidandPathandCmdline
	//
	// - Kill container process by file MD5 and path: killContainerProcessByMd5AndPath
	//
	// When OperationCode is set to block_ip, the parameter is:
	//
	// - Expiration time: expireTime:${timestamp}
	//
	// > This parameter is required only when OperationCode is set to `kill_and_quara` or `block_ip`. For other values of OperationCode, this parameter can be left empty. ${timestamp} indicates the timestamp of the deadline for blocking this IP address.
	//
	// example:
	//
	// {\\"expireTime\\":1719588943551,\\"subOperation\\":\\"killAndQuaraFileByMd5andPath\\"}
	OperationParams *string `json:"OperationParams,omitempty" xml:"OperationParams,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks.
	//
	// example:
	//
	// 11
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The collection of IDs of the security alert events to handle.
	//
	// Example:
	SecurityEventIds []*string `json:"SecurityEventIds,omitempty" xml:"SecurityEventIds,omitempty" type:"Repeated"`
}

func (s GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetAlertType() *string {
	return s.AlertType
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetFileMd5() *string {
	return s.FileMd5
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetIp() *string {
	return s.Ip
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetMarkBatch() *string {
	return s.MarkBatch
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetMarkMissParam() *string {
	return s.MarkMissParam
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetOperationCode() *string {
	return s.OperationCode
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetOperationParams() *string {
	return s.OperationParams
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetRemark() *string {
	return s.Remark
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetSecurityEventIds() []*string {
	return s.SecurityEventIds
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetAlertType(v string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.AlertType = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetFileMd5(v string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.FileMd5 = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetFilePath(v string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.FilePath = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetInstanceId(v string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetIp(v string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.Ip = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetMarkBatch(v string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.MarkBatch = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetMarkMissParam(v string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.MarkMissParam = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetOperationCode(v string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.OperationCode = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetOperationParams(v string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.OperationParams = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetRegionId(v string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.RegionId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetRemark(v string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.Remark = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetSecurityEventIds(v []*string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.SecurityEventIds = v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) Validate() error {
	return dara.Validate(s)
}

type GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest struct {
	// The alerting type.
	//
	// example:
	//
	// default
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cfw_elasticity_public_cn-g4t3nkh3i00b
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the instance.
	//
	// example:
	//
	// 203.10*.44.71
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The operation type for batch handling similar security alert events.
	//
	// > You can call the DescribeSecurityEventOperations operation to obtain this parameter.
	//
	// example:
	//
	// offline_handled
	OperationCode *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	// The configuration of the sub-operation for handling alerting events. The value is in JSON format.
	//
	// > This parameter is required only when **OperationCode*	- is set to **kill_and_quara**, **block_ip**, or **virus_quara**. For other values of **OperationCode**, this parameter can be left empty.
	//
	// > When **OperationCode*	- is set to **block_ip**, the following field is included:
	//
	// > - **expireTime**: The lock expiration time. Unit: milliseconds.
	//
	// >
	//
	// > When **OperationCode*	- is set to **kill_and_quara**, the following field is included:
	//
	// > - **subOperation**: The method for killing and quarantining. Valid values:
	//
	// >     - **killAndQuaraFileByMd5andPath**: Terminates the process and quarantines the file.
	//
	// >     - **killByMd5andPath**: Terminates the running process.
	//
	// >
	//
	// > When **OperationCode*	- is set to **virus_quara**, the following field is included:
	//
	// > - **subOperation**: The method for killing and quarantining. Valid values:
	//
	// >    - **quaraFileByMd5andPath**: Quarantines the source file of the process.
	//
	// example:
	//
	// {\\"expireTime\\":1767687685917}
	OperationParams *string `json:"OperationParams,omitempty" xml:"OperationParams,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 220.2*3.155.93
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the task for batch handling all security alert events of the same type.
	//
	// > You can call the CreateSimilarSecurityEventsQueryTask operation to obtain this parameter.
	//
	// example:
	//
	// 12221
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) GetAlertType() *string {
	return s.AlertType
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) GetIp() *string {
	return s.Ip
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) GetOperationCode() *string {
	return s.OperationCode
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) GetOperationParams() *string {
	return s.OperationParams
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) SetAlertType(v string) *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest {
	s.AlertType = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) SetInstanceId(v string) *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) SetIp(v string) *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest {
	s.Ip = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) SetOperationCode(v string) *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest {
	s.OperationCode = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) SetOperationParams(v string) *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest {
	s.OperationParams = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) SetRegionId(v string) *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest {
	s.RegionId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) SetSourceIp(v string) *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) SetTaskId(v int64) *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest {
	s.TaskId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) Validate() error {
	return dara.Validate(s)
}

type GetAliYunSafeCenterResultRequestListInstancesRequest struct {
	// The instance IDs of simple application servers. The value is a JSON array that can contain up to 100 IDs. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// ["2ad1ae67295445f598017499dc****", "2ad1ae67295445f598017123dc****"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAliYunSafeCenterResultRequestListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultRequestListInstancesRequest) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultRequestListInstancesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *GetAliYunSafeCenterResultRequestListInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAliYunSafeCenterResultRequestListInstancesRequest) SetInstanceIds(v string) *GetAliYunSafeCenterResultRequestListInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestListInstancesRequest) SetRegionId(v string) *GetAliYunSafeCenterResultRequestListInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestListInstancesRequest) Validate() error {
	return dara.Validate(s)
}
