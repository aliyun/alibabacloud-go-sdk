// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAuditLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessSourceFlag(v string) *QueryAuditLogRequest
	GetAccessSourceFlag() *string
	SetEndDate(v string) *QueryAuditLogRequest
	GetEndDate() *string
	SetLogType(v string) *QueryAuditLogRequest
	GetLogType() *string
	SetOperatorId(v string) *QueryAuditLogRequest
	GetOperatorId() *string
	SetOperatorTypes(v string) *QueryAuditLogRequest
	GetOperatorTypes() *string
	SetResourceType(v string) *QueryAuditLogRequest
	GetResourceType() *string
	SetStartDate(v string) *QueryAuditLogRequest
	GetStartDate() *string
	SetUserAccessDevice(v string) *QueryAuditLogRequest
	GetUserAccessDevice() *string
	SetWorkspaceId(v string) *QueryAuditLogRequest
	GetWorkspaceId() *string
}

type QueryAuditLogRequest struct {
	// The access source. Valid values:
	//
	// - COMMON: standard access
	//
	// - IMBEDDED: embedded report
	//
	// - PUBLIC: public report
	//
	// - IMBEDDED_COMPONENT: embedded card
	//
	// if can be null:
	// true
	//
	// example:
	//
	// PUBLIC
	AccessSourceFlag *string `json:"AccessSourceFlag,omitempty" xml:"AccessSourceFlag,omitempty"`
	// The end date for the query. Use the yyyyMMdd format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20240604
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The log type. Valid values:
	//
	// - dataView: access logs
	//
	// - function: operation logs
	//
	// - permission: permission logs
	//
	// This parameter is required.
	//
	// example:
	//
	// function
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The user ID of the operator.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0***
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// The permission, access, or operation type. If left empty, all types are queried by default.
	//
	// For valid values, see audit log codes. To specify multiple types, separate them with commas.
	//
	// example:
	//
	// MODIFY
	OperatorTypes *string `json:"OperatorTypes,omitempty" xml:"OperatorTypes,omitempty"`
	// The resource type. For more information, see work types.
	//
	// example:
	//
	// cube
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The start date for the query. Use the yyyyMMdd format. The date cannot be more than 90 days before the current date.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20240504
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The device used for access. Valid values:
	//
	// - MOBILE: mobile device
	//
	// - PC: PC
	//
	// if can be null:
	// true
	//
	// example:
	//
	// PC
	UserAccessDevice *string `json:"UserAccessDevice,omitempty" xml:"UserAccessDevice,omitempty"`
	// The ID of the workspace that contains the logs to query.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryAuditLogRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAuditLogRequest) GoString() string {
	return s.String()
}

func (s *QueryAuditLogRequest) GetAccessSourceFlag() *string {
	return s.AccessSourceFlag
}

func (s *QueryAuditLogRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryAuditLogRequest) GetLogType() *string {
	return s.LogType
}

func (s *QueryAuditLogRequest) GetOperatorId() *string {
	return s.OperatorId
}

func (s *QueryAuditLogRequest) GetOperatorTypes() *string {
	return s.OperatorTypes
}

func (s *QueryAuditLogRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *QueryAuditLogRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryAuditLogRequest) GetUserAccessDevice() *string {
	return s.UserAccessDevice
}

func (s *QueryAuditLogRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryAuditLogRequest) SetAccessSourceFlag(v string) *QueryAuditLogRequest {
	s.AccessSourceFlag = &v
	return s
}

func (s *QueryAuditLogRequest) SetEndDate(v string) *QueryAuditLogRequest {
	s.EndDate = &v
	return s
}

func (s *QueryAuditLogRequest) SetLogType(v string) *QueryAuditLogRequest {
	s.LogType = &v
	return s
}

func (s *QueryAuditLogRequest) SetOperatorId(v string) *QueryAuditLogRequest {
	s.OperatorId = &v
	return s
}

func (s *QueryAuditLogRequest) SetOperatorTypes(v string) *QueryAuditLogRequest {
	s.OperatorTypes = &v
	return s
}

func (s *QueryAuditLogRequest) SetResourceType(v string) *QueryAuditLogRequest {
	s.ResourceType = &v
	return s
}

func (s *QueryAuditLogRequest) SetStartDate(v string) *QueryAuditLogRequest {
	s.StartDate = &v
	return s
}

func (s *QueryAuditLogRequest) SetUserAccessDevice(v string) *QueryAuditLogRequest {
	s.UserAccessDevice = &v
	return s
}

func (s *QueryAuditLogRequest) SetWorkspaceId(v string) *QueryAuditLogRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryAuditLogRequest) Validate() error {
	return dara.Validate(s)
}
