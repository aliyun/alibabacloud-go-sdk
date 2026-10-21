// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyWorkspaceAcrRamAuthorizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VerifyWorkspaceAcrRamAuthorizationResponseBody
	GetCode() *string
	SetData(v *VerifyWorkspaceAcrRamAuthorizationResponseBodyData) *VerifyWorkspaceAcrRamAuthorizationResponseBody
	GetData() *VerifyWorkspaceAcrRamAuthorizationResponseBodyData
	SetHttpStatusCode(v int32) *VerifyWorkspaceAcrRamAuthorizationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *VerifyWorkspaceAcrRamAuthorizationResponseBody
	GetMessage() *string
	SetRequestId(v string) *VerifyWorkspaceAcrRamAuthorizationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *VerifyWorkspaceAcrRamAuthorizationResponseBody
	GetSuccess() *bool
}

type VerifyWorkspaceAcrRamAuthorizationResponseBody struct {
	// The business status code.
	//
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	//
	// This parameter is required.
	Data *VerifyWorkspaceAcrRamAuthorizationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-1234-1234-123456789012
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s VerifyWorkspaceAcrRamAuthorizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyWorkspaceAcrRamAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBody) GetCode() *string {
	return s.Code
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBody) GetData() *VerifyWorkspaceAcrRamAuthorizationResponseBodyData {
	return s.Data
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBody) SetCode(v string) *VerifyWorkspaceAcrRamAuthorizationResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBody) SetData(v *VerifyWorkspaceAcrRamAuthorizationResponseBodyData) *VerifyWorkspaceAcrRamAuthorizationResponseBody {
	s.Data = v
	return s
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBody) SetHttpStatusCode(v int32) *VerifyWorkspaceAcrRamAuthorizationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBody) SetMessage(v string) *VerifyWorkspaceAcrRamAuthorizationResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBody) SetRequestId(v string) *VerifyWorkspaceAcrRamAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBody) SetSuccess(v bool) *VerifyWorkspaceAcrRamAuthorizationResponseBody {
	s.Success = &v
	return s
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifyWorkspaceAcrRamAuthorizationResponseBodyData struct {
	// The ID of the ACR Enterprise instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-1234567890abcdef
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// The current attachment status of the policy for the target repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// UNAUTHORIZED
	AuthorizationStatus *string `json:"authorizationStatus,omitempty" xml:"authorizationStatus,omitempty"`
	// The access prerequisite status, which does not indicate Secret Ready.
	//
	// This parameter is required.
	//
	// example:
	//
	// ELIGIBLE
	EligibilityStatus *string `json:"eligibilityStatus,omitempty" xml:"eligibilityStatus,omitempty"`
	// The stable reason code for unauthorized or unmet conditions. This field is omitted if no reason exists.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// AcrRamUnauthorized
	ReasonCode *string `json:"reasonCode,omitempty" xml:"reasonCode,omitempty"`
	// The user-readable reason. This field is omitted if no reason exists.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// Authorize the Workspace role for this ACR instance.
	ReasonMessage *string `json:"reasonMessage,omitempty" xml:"reasonMessage,omitempty"`
	// The name of the shared role selected by the backend. This field is not editable on the frontend.
	//
	// This parameter is required.
	//
	// example:
	//
	// AliyunAgentCoreWorkspace-ws-1234567890abcdef12345
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	// The source of the shared role. This does not indicate that authorization is complete.
	//
	// This parameter is required.
	//
	// example:
	//
	// WORKSPACE_SHARED
	RoleSource *string `json:"roleSource,omitempty" xml:"roleSource,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-1234567890abcdef12345
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s VerifyWorkspaceAcrRamAuthorizationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s VerifyWorkspaceAcrRamAuthorizationResponseBodyData) GoString() string {
	return s.String()
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBodyData) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBodyData) GetAuthorizationStatus() *string {
	return s.AuthorizationStatus
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBodyData) GetEligibilityStatus() *string {
	return s.EligibilityStatus
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBodyData) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBodyData) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBodyData) GetRoleName() *string {
	return s.RoleName
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBodyData) GetRoleSource() *string {
	return s.RoleSource
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBodyData) SetAcrInstanceId(v string) *VerifyWorkspaceAcrRamAuthorizationResponseBodyData {
	s.AcrInstanceId = &v
	return s
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBodyData) SetAuthorizationStatus(v string) *VerifyWorkspaceAcrRamAuthorizationResponseBodyData {
	s.AuthorizationStatus = &v
	return s
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBodyData) SetEligibilityStatus(v string) *VerifyWorkspaceAcrRamAuthorizationResponseBodyData {
	s.EligibilityStatus = &v
	return s
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBodyData) SetReasonCode(v string) *VerifyWorkspaceAcrRamAuthorizationResponseBodyData {
	s.ReasonCode = &v
	return s
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBodyData) SetReasonMessage(v string) *VerifyWorkspaceAcrRamAuthorizationResponseBodyData {
	s.ReasonMessage = &v
	return s
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBodyData) SetRoleName(v string) *VerifyWorkspaceAcrRamAuthorizationResponseBodyData {
	s.RoleName = &v
	return s
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBodyData) SetRoleSource(v string) *VerifyWorkspaceAcrRamAuthorizationResponseBodyData {
	s.RoleSource = &v
	return s
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBodyData) SetWorkspaceId(v string) *VerifyWorkspaceAcrRamAuthorizationResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
