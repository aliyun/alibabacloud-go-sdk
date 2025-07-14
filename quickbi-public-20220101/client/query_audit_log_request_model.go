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
	// if can be null:
	// true
	AccessSourceFlag *string `json:"AccessSourceFlag,omitempty" xml:"AccessSourceFlag,omitempty"`
	// End date of the query, format ("yyyyMMdd").
	//
	// This parameter is required.
	//
	// example:
	//
	// 20240604
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Log type:
	//
	// - dataView - Access
	//
	// - function - Operation
	//
	// - permission - Permission
	//
	// This parameter is required.
	//
	// example:
	//
	// function
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// Operator\\"s user ID.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0***
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// Permission/Access/Operation type, empty - default all;
	//
	// Refer to the audit log code values, send multiple values separated by English commas.
	//
	// example:
	//
	// MODIFY
	OperatorTypes *string `json:"OperatorTypes,omitempty" xml:"OperatorTypes,omitempty"`
	// Resource type, refer to the work type.
	//
	// example:
	//
	// cube
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Start date of the query, format ("yyyyMMdd"), cannot be earlier than 90 days from the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20240504
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// if can be null:
	// true
	UserAccessDevice *string `json:"UserAccessDevice,omitempty" xml:"UserAccessDevice,omitempty"`
	// Workspace ID, the ID of the workspace to which the logs to be queried belong.
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
