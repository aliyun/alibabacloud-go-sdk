// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceLastAccessedDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsTruncated(v bool) *GetServiceLastAccessedDetailsResponseBody
	GetIsTruncated() *bool
	SetJobCompletionTime(v string) *GetServiceLastAccessedDetailsResponseBody
	GetJobCompletionTime() *string
	SetJobCreationTime(v string) *GetServiceLastAccessedDetailsResponseBody
	GetJobCreationTime() *string
	SetJobError(v *GetServiceLastAccessedDetailsResponseBodyJobError) *GetServiceLastAccessedDetailsResponseBody
	GetJobError() *GetServiceLastAccessedDetailsResponseBodyJobError
	SetJobStatus(v string) *GetServiceLastAccessedDetailsResponseBody
	GetJobStatus() *string
	SetJobType(v string) *GetServiceLastAccessedDetailsResponseBody
	GetJobType() *string
	SetMarker(v string) *GetServiceLastAccessedDetailsResponseBody
	GetMarker() *string
	SetRequestId(v string) *GetServiceLastAccessedDetailsResponseBody
	GetRequestId() *string
	SetServicesLastAccessed(v []*GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed) *GetServiceLastAccessedDetailsResponseBody
	GetServicesLastAccessed() []*GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed
}

type GetServiceLastAccessedDetailsResponseBody struct {
	// Specifies whether the returned results are truncated. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The time when the task was completed.
	//
	// example:
	//
	// 2024-06-08T02:24:48Z
	JobCompletionTime *string `json:"JobCompletionTime,omitempty" xml:"JobCompletionTime,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2024-06-08T02:20:47Z
	JobCreationTime *string `json:"JobCreationTime,omitempty" xml:"JobCreationTime,omitempty"`
	// The error message returned when the task fails. This parameter is not returned when the task is completed.
	JobError *GetServiceLastAccessedDetailsResponseBodyJobError `json:"JobError,omitempty" xml:"JobError,omitempty" type:"Struct"`
	// The status of the task. Valid values:
	//
	// - IN_PROGRESS: The task is being processed.
	//
	// - COMPLETED: The task is completed.
	//
	// - FAILED: The task failed.
	//
	// example:
	//
	// COMPLETED
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The granularity of the permission audit report. Valid values:
	//
	// - SERVICE_LEVEL: Analyzes the recent access information at the RAM code granularity for cloud services supported by permission audit.
	//
	// - ACTION_LEVEL: Analyzes the recent access information at the RAM action granularity for operation-level cloud services supported by permission audit.
	//
	// example:
	//
	// SERVICE_LEVEL
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// This parameter is returned only when IsTruncated is set to true. If true is returned, you must call this operation again and use Marker to obtain the truncated content.
	//
	// This parameter is required.
	//
	// example:
	//
	// CgxMaXN0UmFtVXNlcnMSEDE5OTQ5NTMxMDYwOTAzNzkaAjQ0IGQ=
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 697852FB-50D7-44D9-9774-530C31EAC572
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of last access information for cloud services. This list is empty if the RAM identity does not have access permissions for any cloud service.
	ServicesLastAccessed []*GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed `json:"ServicesLastAccessed,omitempty" xml:"ServicesLastAccessed,omitempty" type:"Repeated"`
}

func (s GetServiceLastAccessedDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLastAccessedDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceLastAccessedDetailsResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *GetServiceLastAccessedDetailsResponseBody) GetJobCompletionTime() *string {
	return s.JobCompletionTime
}

func (s *GetServiceLastAccessedDetailsResponseBody) GetJobCreationTime() *string {
	return s.JobCreationTime
}

func (s *GetServiceLastAccessedDetailsResponseBody) GetJobError() *GetServiceLastAccessedDetailsResponseBodyJobError {
	return s.JobError
}

func (s *GetServiceLastAccessedDetailsResponseBody) GetJobStatus() *string {
	return s.JobStatus
}

func (s *GetServiceLastAccessedDetailsResponseBody) GetJobType() *string {
	return s.JobType
}

func (s *GetServiceLastAccessedDetailsResponseBody) GetMarker() *string {
	return s.Marker
}

func (s *GetServiceLastAccessedDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceLastAccessedDetailsResponseBody) GetServicesLastAccessed() []*GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed {
	return s.ServicesLastAccessed
}

func (s *GetServiceLastAccessedDetailsResponseBody) SetIsTruncated(v bool) *GetServiceLastAccessedDetailsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBody) SetJobCompletionTime(v string) *GetServiceLastAccessedDetailsResponseBody {
	s.JobCompletionTime = &v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBody) SetJobCreationTime(v string) *GetServiceLastAccessedDetailsResponseBody {
	s.JobCreationTime = &v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBody) SetJobError(v *GetServiceLastAccessedDetailsResponseBodyJobError) *GetServiceLastAccessedDetailsResponseBody {
	s.JobError = v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBody) SetJobStatus(v string) *GetServiceLastAccessedDetailsResponseBody {
	s.JobStatus = &v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBody) SetJobType(v string) *GetServiceLastAccessedDetailsResponseBody {
	s.JobType = &v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBody) SetMarker(v string) *GetServiceLastAccessedDetailsResponseBody {
	s.Marker = &v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBody) SetRequestId(v string) *GetServiceLastAccessedDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBody) SetServicesLastAccessed(v []*GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed) *GetServiceLastAccessedDetailsResponseBody {
	s.ServicesLastAccessed = v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBody) Validate() error {
	if s.JobError != nil {
		if err := s.JobError.Validate(); err != nil {
			return err
		}
	}
	if s.ServicesLastAccessed != nil {
		for _, item := range s.ServicesLastAccessed {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetServiceLastAccessedDetailsResponseBodyJobError struct {
	// The error code.
	//
	// example:
	//
	// InternalError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// InternalError
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s GetServiceLastAccessedDetailsResponseBodyJobError) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLastAccessedDetailsResponseBodyJobError) GoString() string {
	return s.String()
}

func (s *GetServiceLastAccessedDetailsResponseBodyJobError) GetCode() *string {
	return s.Code
}

func (s *GetServiceLastAccessedDetailsResponseBodyJobError) GetMessage() *string {
	return s.Message
}

func (s *GetServiceLastAccessedDetailsResponseBodyJobError) SetCode(v string) *GetServiceLastAccessedDetailsResponseBodyJobError {
	s.Code = &v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBodyJobError) SetMessage(v string) *GetServiceLastAccessedDetailsResponseBodyJobError {
	s.Message = &v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBodyJobError) Validate() error {
	return dara.Validate(s)
}

type GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed struct {
	// The list of access policies that grant permission to access the cloud service.
	GrantPolicies []*GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedGrantPolicies `json:"GrantPolicies,omitempty" xml:"GrantPolicies,omitempty" type:"Repeated"`
	// The time when the analyzed RAM identity last attempted to access the service.
	//
	// example:
	//
	// 2024-01-29T07:48:35Z
	LastAuthenticated *string `json:"LastAuthenticated,omitempty" xml:"LastAuthenticated,omitempty"`
	// The ARN of the analyzed RAM identity.
	//
	// example:
	//
	// acs:ram::123459***:role/TestRole
	LastAuthenticatedEntity *string `json:"LastAuthenticatedEntity,omitempty" xml:"LastAuthenticatedEntity,omitempty"`
	// The RAM code of the cloud service.
	//
	// example:
	//
	// ecs
	RAMCode *string `json:"RAMCode,omitempty" xml:"RAMCode,omitempty"`
	// The list of cloud service codes.
	ServiceCodes []*string `json:"ServiceCodes,omitempty" xml:"ServiceCodes,omitempty" type:"Repeated"`
	// The number of identities that attempted to access the service.
	//
	// example:
	//
	// 1
	TotalAuthenticatedEntities *int32 `json:"TotalAuthenticatedEntities,omitempty" xml:"TotalAuthenticatedEntities,omitempty"`
	// The list of operation-level access information. This information is returned when the granularity of the created permission audit task is ACTION_LEVEL.
	TrackedActionsLastAccessed []*GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedTrackedActionsLastAccessed `json:"TrackedActionsLastAccessed,omitempty" xml:"TrackedActionsLastAccessed,omitempty" type:"Repeated"`
}

func (s GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed) GoString() string {
	return s.String()
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed) GetGrantPolicies() []*GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedGrantPolicies {
	return s.GrantPolicies
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed) GetLastAuthenticated() *string {
	return s.LastAuthenticated
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed) GetLastAuthenticatedEntity() *string {
	return s.LastAuthenticatedEntity
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed) GetRAMCode() *string {
	return s.RAMCode
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed) GetServiceCodes() []*string {
	return s.ServiceCodes
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed) GetTotalAuthenticatedEntities() *int32 {
	return s.TotalAuthenticatedEntities
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed) GetTrackedActionsLastAccessed() []*GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedTrackedActionsLastAccessed {
	return s.TrackedActionsLastAccessed
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed) SetGrantPolicies(v []*GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedGrantPolicies) *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed {
	s.GrantPolicies = v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed) SetLastAuthenticated(v string) *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed {
	s.LastAuthenticated = &v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed) SetLastAuthenticatedEntity(v string) *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed {
	s.LastAuthenticatedEntity = &v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed) SetRAMCode(v string) *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed {
	s.RAMCode = &v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed) SetServiceCodes(v []*string) *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed {
	s.ServiceCodes = v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed) SetTotalAuthenticatedEntities(v int32) *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed {
	s.TotalAuthenticatedEntities = &v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed) SetTrackedActionsLastAccessed(v []*GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedTrackedActionsLastAccessed) *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed {
	s.TrackedActionsLastAccessed = v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessed) Validate() error {
	if s.GrantPolicies != nil {
		for _, item := range s.GrantPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TrackedActionsLastAccessed != nil {
		for _, item := range s.TrackedActionsLastAccessed {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedGrantPolicies struct {
	// The ARN of the entity to which the access policy grants permission. Valid values:
	//
	// - RAM user: acs:ram::<account_id>:user/<user_name>
	//
	// - RAM role: acs:ram::<account_id>:role/<role_name>
	//
	// - RAM user group: acs:ram::<account_id>:group/<group_name>
	//
	// example:
	//
	// acs:ram::123459***:role/TestRole
	AttachedEntity *string `json:"AttachedEntity,omitempty" xml:"AttachedEntity,omitempty"`
	// The authorization scope of the access policy. Valid values:
	//
	// - Account ID (format: 123459***): account-level authorization
	//
	// - Resource group ID (format: rg-123***): resource group-level authorization
	//
	// example:
	//
	// 123459***
	AttachedScope *string `json:"AttachedScope,omitempty" xml:"AttachedScope,omitempty"`
	// The name of the access policy that grants permission.
	//
	// example:
	//
	// AdministratorAccess
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the access policy. Valid values:
	//
	// - System: system policy
	//
	// - Custom: custom policy
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedGrantPolicies) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedGrantPolicies) GoString() string {
	return s.String()
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedGrantPolicies) GetAttachedEntity() *string {
	return s.AttachedEntity
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedGrantPolicies) GetAttachedScope() *string {
	return s.AttachedScope
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedGrantPolicies) GetPolicyName() *string {
	return s.PolicyName
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedGrantPolicies) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedGrantPolicies) SetAttachedEntity(v string) *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedGrantPolicies {
	s.AttachedEntity = &v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedGrantPolicies) SetAttachedScope(v string) *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedGrantPolicies {
	s.AttachedScope = &v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedGrantPolicies) SetPolicyName(v string) *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedGrantPolicies {
	s.PolicyName = &v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedGrantPolicies) SetPolicyType(v string) *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedGrantPolicies {
	s.PolicyType = &v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedGrantPolicies) Validate() error {
	return dara.Validate(s)
}

type GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedTrackedActionsLastAccessed struct {
	// The action.
	//
	// example:
	//
	// ram:CreateAccessKey
	ActionName *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	// The ARN of the analyzed RAM identity.
	//
	// example:
	//
	// acs:ram::123459***:role/TestRole
	LastAccessedEntity *string `json:"LastAccessedEntity,omitempty" xml:"LastAccessedEntity,omitempty"`
	// The time when the analyzed RAM identity last attempted to perform the action.
	//
	// example:
	//
	// 2024-01-29T07:48:35Z
	LastAccessedTime *string `json:"LastAccessedTime,omitempty" xml:"LastAccessedTime,omitempty"`
	// The list of cloud service codes.
	ServiceCodes []*string `json:"ServiceCodes,omitempty" xml:"ServiceCodes,omitempty" type:"Repeated"`
}

func (s GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedTrackedActionsLastAccessed) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedTrackedActionsLastAccessed) GoString() string {
	return s.String()
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedTrackedActionsLastAccessed) GetActionName() *string {
	return s.ActionName
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedTrackedActionsLastAccessed) GetLastAccessedEntity() *string {
	return s.LastAccessedEntity
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedTrackedActionsLastAccessed) GetLastAccessedTime() *string {
	return s.LastAccessedTime
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedTrackedActionsLastAccessed) GetServiceCodes() []*string {
	return s.ServiceCodes
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedTrackedActionsLastAccessed) SetActionName(v string) *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedTrackedActionsLastAccessed {
	s.ActionName = &v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedTrackedActionsLastAccessed) SetLastAccessedEntity(v string) *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedTrackedActionsLastAccessed {
	s.LastAccessedEntity = &v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedTrackedActionsLastAccessed) SetLastAccessedTime(v string) *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedTrackedActionsLastAccessed {
	s.LastAccessedTime = &v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedTrackedActionsLastAccessed) SetServiceCodes(v []*string) *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedTrackedActionsLastAccessed {
	s.ServiceCodes = v
	return s
}

func (s *GetServiceLastAccessedDetailsResponseBodyServicesLastAccessedTrackedActionsLastAccessed) Validate() error {
	return dara.Validate(s)
}
