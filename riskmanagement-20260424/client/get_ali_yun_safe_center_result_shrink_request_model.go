// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAliYunSafeCenterResultShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateSimilarSecurityEventsQueryTaskRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest
	GetCreateSimilarSecurityEventsQueryTaskRequestShrink() *string
	SetDescribeInstancesFullStatusRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest
	GetDescribeInstancesFullStatusRequestShrink() *string
	SetDescribeSecurityEventOperationStatusRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest
	GetDescribeSecurityEventOperationStatusRequestShrink() *string
	SetDescribeSimilarSecurityEventsRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest
	GetDescribeSimilarSecurityEventsRequestShrink() *string
	SetGetAssetDetailByUuidRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest
	GetGetAssetDetailByUuidRequestShrink() *string
	SetHandleSecurityEventsRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest
	GetHandleSecurityEventsRequestShrink() *string
	SetHandleSimilarSecurityEventsRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest
	GetHandleSimilarSecurityEventsRequestShrink() *string
	SetInterfaceCode(v string) *GetAliYunSafeCenterResultShrinkRequest
	GetInterfaceCode() *string
	SetListInstancesRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest
	GetListInstancesRequestShrink() *string
	SetRegionId(v string) *GetAliYunSafeCenterResultShrinkRequest
	GetRegionId() *string
}

type GetAliYunSafeCenterResultShrinkRequest struct {
	// Creates a node to query security alerting events triggered by the same rule or alerting type.
	CreateSimilarSecurityEventsQueryTaskRequestShrink *string `json:"CreateSimilarSecurityEventsQueryTaskRequest,omitempty" xml:"CreateSimilarSecurityEventsQueryTaskRequest,omitempty"`
	// Queries the running status of ECS instances.
	DescribeInstancesFullStatusRequestShrink *string `json:"DescribeInstancesFullStatusRequest,omitempty" xml:"DescribeInstancesFullStatusRequest,omitempty"`
	// Queries whether the list of security alerting events that match the same IP rule or same alerting type as the alerting event to be handled is empty.
	DescribeSecurityEventOperationStatusRequestShrink *string `json:"DescribeSecurityEventOperationStatusRequest,omitempty" xml:"DescribeSecurityEventOperationStatusRequest,omitempty"`
	// Queries identical security alert events in Security Center.
	DescribeSimilarSecurityEventsRequestShrink *string `json:"DescribeSimilarSecurityEventsRequest,omitempty" xml:"DescribeSimilarSecurityEventsRequest,omitempty"`
	// The request parameters for querying the Security Center Agent status.
	GetAssetDetailByUuidRequestShrink *string `json:"GetAssetDetailByUuidRequest,omitempty" xml:"GetAssetDetailByUuidRequest,omitempty"`
	// Handles security alert events.
	HandleSecurityEventsRequestShrink *string `json:"HandleSecurityEventsRequest,omitempty" xml:"HandleSecurityEventsRequest,omitempty"`
	// Handles security alert events in batches based on the same IP rule or type.
	HandleSimilarSecurityEventsRequestShrink *string `json:"HandleSimilarSecurityEventsRequest,omitempty" xml:"HandleSimilarSecurityEventsRequest,omitempty"`
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
	ListInstancesRequestShrink *string `json:"ListInstancesRequest,omitempty" xml:"ListInstancesRequest,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-guangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAliYunSafeCenterResultShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultShrinkRequest) GetCreateSimilarSecurityEventsQueryTaskRequestShrink() *string {
	return s.CreateSimilarSecurityEventsQueryTaskRequestShrink
}

func (s *GetAliYunSafeCenterResultShrinkRequest) GetDescribeInstancesFullStatusRequestShrink() *string {
	return s.DescribeInstancesFullStatusRequestShrink
}

func (s *GetAliYunSafeCenterResultShrinkRequest) GetDescribeSecurityEventOperationStatusRequestShrink() *string {
	return s.DescribeSecurityEventOperationStatusRequestShrink
}

func (s *GetAliYunSafeCenterResultShrinkRequest) GetDescribeSimilarSecurityEventsRequestShrink() *string {
	return s.DescribeSimilarSecurityEventsRequestShrink
}

func (s *GetAliYunSafeCenterResultShrinkRequest) GetGetAssetDetailByUuidRequestShrink() *string {
	return s.GetAssetDetailByUuidRequestShrink
}

func (s *GetAliYunSafeCenterResultShrinkRequest) GetHandleSecurityEventsRequestShrink() *string {
	return s.HandleSecurityEventsRequestShrink
}

func (s *GetAliYunSafeCenterResultShrinkRequest) GetHandleSimilarSecurityEventsRequestShrink() *string {
	return s.HandleSimilarSecurityEventsRequestShrink
}

func (s *GetAliYunSafeCenterResultShrinkRequest) GetInterfaceCode() *string {
	return s.InterfaceCode
}

func (s *GetAliYunSafeCenterResultShrinkRequest) GetListInstancesRequestShrink() *string {
	return s.ListInstancesRequestShrink
}

func (s *GetAliYunSafeCenterResultShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAliYunSafeCenterResultShrinkRequest) SetCreateSimilarSecurityEventsQueryTaskRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest {
	s.CreateSimilarSecurityEventsQueryTaskRequestShrink = &v
	return s
}

func (s *GetAliYunSafeCenterResultShrinkRequest) SetDescribeInstancesFullStatusRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest {
	s.DescribeInstancesFullStatusRequestShrink = &v
	return s
}

func (s *GetAliYunSafeCenterResultShrinkRequest) SetDescribeSecurityEventOperationStatusRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest {
	s.DescribeSecurityEventOperationStatusRequestShrink = &v
	return s
}

func (s *GetAliYunSafeCenterResultShrinkRequest) SetDescribeSimilarSecurityEventsRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest {
	s.DescribeSimilarSecurityEventsRequestShrink = &v
	return s
}

func (s *GetAliYunSafeCenterResultShrinkRequest) SetGetAssetDetailByUuidRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest {
	s.GetAssetDetailByUuidRequestShrink = &v
	return s
}

func (s *GetAliYunSafeCenterResultShrinkRequest) SetHandleSecurityEventsRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest {
	s.HandleSecurityEventsRequestShrink = &v
	return s
}

func (s *GetAliYunSafeCenterResultShrinkRequest) SetHandleSimilarSecurityEventsRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest {
	s.HandleSimilarSecurityEventsRequestShrink = &v
	return s
}

func (s *GetAliYunSafeCenterResultShrinkRequest) SetInterfaceCode(v string) *GetAliYunSafeCenterResultShrinkRequest {
	s.InterfaceCode = &v
	return s
}

func (s *GetAliYunSafeCenterResultShrinkRequest) SetListInstancesRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest {
	s.ListInstancesRequestShrink = &v
	return s
}

func (s *GetAliYunSafeCenterResultShrinkRequest) SetRegionId(v string) *GetAliYunSafeCenterResultShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetAliYunSafeCenterResultShrinkRequest) Validate() error {
	return dara.Validate(s)
}
