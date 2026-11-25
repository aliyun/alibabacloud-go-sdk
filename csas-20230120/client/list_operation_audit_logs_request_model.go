// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationAuditLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListOperationAuditLogsRequest
	GetCurrentPage() *int32
	SetEndTime(v string) *ListOperationAuditLogsRequest
	GetEndTime() *string
	SetEventType(v string) *ListOperationAuditLogsRequest
	GetEventType() *string
	SetOperationFunc(v string) *ListOperationAuditLogsRequest
	GetOperationFunc() *string
	SetOperationStatus(v string) *ListOperationAuditLogsRequest
	GetOperationStatus() *string
	SetOperationType(v string) *ListOperationAuditLogsRequest
	GetOperationType() *string
	SetOperatorId(v string) *ListOperationAuditLogsRequest
	GetOperatorId() *string
	SetPageSize(v int32) *ListOperationAuditLogsRequest
	GetPageSize() *int32
	SetStartTime(v string) *ListOperationAuditLogsRequest
	GetStartTime() *string
}

type ListOperationAuditLogsRequest struct {
	// The page number of the current page for a paged query. Pages start from 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end time of the query. The value is a UNIX timestamp in seconds. The value must be later than StartTime.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1787550343
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The source type of the event. Valid values:
	//
	// - **console**: The event is triggered by a console call.
	//
	// - **sdk**: The event is triggered by an SDK call.
	//
	// example:
	//
	// console
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The operation feature module. The value must be the English code of the feature module. If you specify other values, no records are queried. Valid values:
	//
	// - **PrivateAccess**: internal network access
	//
	// - **OfficeNetworkAccess**: office network access
	//
	// - **AppAcceleration**: application acceleration
	//
	// - **InternetAccess**: Internet access
	//
	// - **OfficeDataProtection**: office data protection
	//
	// - **IdentityAccessManagement**: identity and access management
	//
	// - **DeviceManagement**: endpoint management
	//
	// - **ApprovalCenter**: approval center
	//
	// - **SoftwareManagement**: software management
	//
	// - **LogAnalysis**: log analysis
	//
	// - **Setting**: settings
	//
	// - **DigitalWatermark**: digital watermarking
	//
	// - **DynamicDecision**: dynamic decision-making
	//
	// - **InternetBehaviorManagement**: Internet behavior management
	//
	// - **AgentOfficeSecurity**: agent office security
	//
	// - **NetworkAccess**: network access
	//
	// - **RiskManagement**: risk management
	//
	// - **EndpointProtection**: endpoint protection
	//
	// - **Overview**: overview page
	//
	// - **ITManagement**: IT management
	//
	// - **InstanceManagement**: instance management
	//
	// example:
	//
	// OfficeDataProtection
	OperationFunc *string `json:"OperationFunc,omitempty" xml:"OperationFunc,omitempty"`
	// The operation status. Valid values:
	//
	// - **success**: The operation is successful. Equivalent values are true and Success.
	//
	// - **failure**: The operation fails. Equivalent values are fail, failed, false, and Failure.
	//
	// If this parameter is not specified, only records of successful operations are returned.
	//
	// example:
	//
	// success
	OperationStatus *string `json:"OperationStatus,omitempty" xml:"OperationStatus,omitempty"`
	// The operation type. The value must be exactly the same as the original operation type recorded in the log. The OperationType in the returned result is localized based on the request language and may differ from this filter value.
	//
	// example:
	//
	// Synchronize classification and grading rules
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The Alibaba Cloud account ID (AliUid) of the operator.
	//
	// example:
	//
	// 1234****
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// The number of entries to return on each page for a paged query. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start time of the query. The value is a UNIX timestamp in seconds. The value must be earlier than EndTime. The interval between StartTime and EndTime cannot exceed 30 days, and StartTime cannot be more than 31 days before the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1786945543
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListOperationAuditLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOperationAuditLogsRequest) GoString() string {
	return s.String()
}

func (s *ListOperationAuditLogsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListOperationAuditLogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListOperationAuditLogsRequest) GetEventType() *string {
	return s.EventType
}

func (s *ListOperationAuditLogsRequest) GetOperationFunc() *string {
	return s.OperationFunc
}

func (s *ListOperationAuditLogsRequest) GetOperationStatus() *string {
	return s.OperationStatus
}

func (s *ListOperationAuditLogsRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *ListOperationAuditLogsRequest) GetOperatorId() *string {
	return s.OperatorId
}

func (s *ListOperationAuditLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOperationAuditLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListOperationAuditLogsRequest) SetCurrentPage(v int32) *ListOperationAuditLogsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListOperationAuditLogsRequest) SetEndTime(v string) *ListOperationAuditLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListOperationAuditLogsRequest) SetEventType(v string) *ListOperationAuditLogsRequest {
	s.EventType = &v
	return s
}

func (s *ListOperationAuditLogsRequest) SetOperationFunc(v string) *ListOperationAuditLogsRequest {
	s.OperationFunc = &v
	return s
}

func (s *ListOperationAuditLogsRequest) SetOperationStatus(v string) *ListOperationAuditLogsRequest {
	s.OperationStatus = &v
	return s
}

func (s *ListOperationAuditLogsRequest) SetOperationType(v string) *ListOperationAuditLogsRequest {
	s.OperationType = &v
	return s
}

func (s *ListOperationAuditLogsRequest) SetOperatorId(v string) *ListOperationAuditLogsRequest {
	s.OperatorId = &v
	return s
}

func (s *ListOperationAuditLogsRequest) SetPageSize(v int32) *ListOperationAuditLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListOperationAuditLogsRequest) SetStartTime(v string) *ListOperationAuditLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListOperationAuditLogsRequest) Validate() error {
	return dara.Validate(s)
}
