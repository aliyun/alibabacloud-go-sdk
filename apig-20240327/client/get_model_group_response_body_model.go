// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetModelGroupResponseBody
	GetCode() *string
	SetData(v *GetModelGroupResponseBodyData) *GetModelGroupResponseBody
	GetData() *GetModelGroupResponseBodyData
	SetMessage(v string) *GetModelGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetModelGroupResponseBody
	GetRequestId() *string
}

type GetModelGroupResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetModelGroupResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetModelGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModelGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetModelGroupResponseBody) GetData() *GetModelGroupResponseBodyData {
	return s.Data
}

func (s *GetModelGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetModelGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModelGroupResponseBody) SetCode(v string) *GetModelGroupResponseBody {
	s.Code = &v
	return s
}

func (s *GetModelGroupResponseBody) SetData(v *GetModelGroupResponseBodyData) *GetModelGroupResponseBody {
	s.Data = v
	return s
}

func (s *GetModelGroupResponseBody) SetMessage(v string) *GetModelGroupResponseBody {
	s.Message = &v
	return s
}

func (s *GetModelGroupResponseBody) SetRequestId(v string) *GetModelGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetModelGroupResponseBodyData struct {
	// The summary of the latest group configuration deployment, which is independent of API publishing and entitlement status. This field is not returned when no deployment task exists, and its presence does not indicate success. It represents the current version only when the revision matches the current version. List operations do not guarantee that this field is populated.
	ConfigurationDeployment *GetModelGroupResponseBodyDataConfigurationDeployment `json:"configurationDeployment,omitempty" xml:"configurationDeployment,omitempty" type:"Struct"`
	// The description of the model group.
	//
	// example:
	//
	// A model group for conversational scenarios
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The entitlement status information.
	EntitlementState *GetModelGroupResponseBodyDataEntitlementState `json:"entitlementState,omitempty" xml:"entitlementState,omitempty" type:"Struct"`
	// The ID of the gateway instance.
	//
	// example:
	//
	// gw-xxxxxxxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The list of model card IDs included in the model group.
	ModelCardIds []*string `json:"modelCardIds,omitempty" xml:"modelCardIds,omitempty" type:"Repeated"`
	// The ID of the model group.
	//
	// example:
	//
	// mg-xxxxxxxx
	ModelGroupId *string `json:"modelGroupId,omitempty" xml:"modelGroupId,omitempty"`
	// The name of the model group.
	//
	// example:
	//
	// chat-group
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The last update time in ISO 8601 (RFC 3339) format in the UTC time zone, such as 2026-05-15T10:00:00Z.
	//
	// example:
	//
	// 2026-05-15T10:00:00Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The version number of the model group.
	//
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetModelGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetModelGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetModelGroupResponseBodyData) GetConfigurationDeployment() *GetModelGroupResponseBodyDataConfigurationDeployment {
	return s.ConfigurationDeployment
}

func (s *GetModelGroupResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetModelGroupResponseBodyData) GetEntitlementState() *GetModelGroupResponseBodyDataEntitlementState {
	return s.EntitlementState
}

func (s *GetModelGroupResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetModelGroupResponseBodyData) GetModelCardIds() []*string {
	return s.ModelCardIds
}

func (s *GetModelGroupResponseBodyData) GetModelGroupId() *string {
	return s.ModelGroupId
}

func (s *GetModelGroupResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetModelGroupResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetModelGroupResponseBodyData) GetVersion() *int64 {
	return s.Version
}

func (s *GetModelGroupResponseBodyData) SetConfigurationDeployment(v *GetModelGroupResponseBodyDataConfigurationDeployment) *GetModelGroupResponseBodyData {
	s.ConfigurationDeployment = v
	return s
}

func (s *GetModelGroupResponseBodyData) SetDescription(v string) *GetModelGroupResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetModelGroupResponseBodyData) SetEntitlementState(v *GetModelGroupResponseBodyDataEntitlementState) *GetModelGroupResponseBodyData {
	s.EntitlementState = v
	return s
}

func (s *GetModelGroupResponseBodyData) SetGatewayId(v string) *GetModelGroupResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetModelGroupResponseBodyData) SetModelCardIds(v []*string) *GetModelGroupResponseBodyData {
	s.ModelCardIds = v
	return s
}

func (s *GetModelGroupResponseBodyData) SetModelGroupId(v string) *GetModelGroupResponseBodyData {
	s.ModelGroupId = &v
	return s
}

func (s *GetModelGroupResponseBodyData) SetName(v string) *GetModelGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetModelGroupResponseBodyData) SetUpdateTime(v string) *GetModelGroupResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetModelGroupResponseBodyData) SetVersion(v int64) *GetModelGroupResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetModelGroupResponseBodyData) Validate() error {
	if s.ConfigurationDeployment != nil {
		if err := s.ConfigurationDeployment.Validate(); err != nil {
			return err
		}
	}
	if s.EntitlementState != nil {
		if err := s.EntitlementState.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetModelGroupResponseBodyDataConfigurationDeployment struct {
	// The operation ID of the group configuration task, which is used for tracking only. It cannot be used for GetLogicalModelDeployment or RetryLogicalModelDeployment.
	//
	// example:
	//
	// op-xxxxxxxx
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
	// The model group version corresponding to the task, which may be earlier than the current version.
	//
	// example:
	//
	// 1
	Revision *int64 `json:"revision,omitempty" xml:"revision,omitempty"`
	// The configuration deployment status. Valid values: Pending, Running, Succeeded, Failed, and PartiallyFailed.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetModelGroupResponseBodyDataConfigurationDeployment) String() string {
	return dara.Prettify(s)
}

func (s GetModelGroupResponseBodyDataConfigurationDeployment) GoString() string {
	return s.String()
}

func (s *GetModelGroupResponseBodyDataConfigurationDeployment) GetOperationId() *string {
	return s.OperationId
}

func (s *GetModelGroupResponseBodyDataConfigurationDeployment) GetRevision() *int64 {
	return s.Revision
}

func (s *GetModelGroupResponseBodyDataConfigurationDeployment) GetStatus() *string {
	return s.Status
}

func (s *GetModelGroupResponseBodyDataConfigurationDeployment) SetOperationId(v string) *GetModelGroupResponseBodyDataConfigurationDeployment {
	s.OperationId = &v
	return s
}

func (s *GetModelGroupResponseBodyDataConfigurationDeployment) SetRevision(v int64) *GetModelGroupResponseBodyDataConfigurationDeployment {
	s.Revision = &v
	return s
}

func (s *GetModelGroupResponseBodyDataConfigurationDeployment) SetStatus(v string) *GetModelGroupResponseBodyDataConfigurationDeployment {
	s.Status = &v
	return s
}

func (s *GetModelGroupResponseBodyDataConfigurationDeployment) Validate() error {
	return dara.Validate(s)
}

type GetModelGroupResponseBodyDataEntitlementState struct {
	// The applied entitlement revision.
	//
	// example:
	//
	// 1
	AppliedRevision *int64 `json:"appliedRevision,omitempty" xml:"appliedRevision,omitempty"`
	// The desired entitlement revision.
	//
	// example:
	//
	// 1
	DesiredRevision *int64 `json:"desiredRevision,omitempty" xml:"desiredRevision,omitempty"`
	// The error message when entitlement fails. This field is not returned when no error occurs.
	//
	// example:
	//
	// Entitlement synchronization failed
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The entitlement status. Valid values: Pending, Applying, Committed, and Failed. Committed indicates only that the control plane has successfully submitted and read back the status, and does not mean that all nodes have acknowledged it.
	//
	// example:
	//
	// Committed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetModelGroupResponseBodyDataEntitlementState) String() string {
	return dara.Prettify(s)
}

func (s GetModelGroupResponseBodyDataEntitlementState) GoString() string {
	return s.String()
}

func (s *GetModelGroupResponseBodyDataEntitlementState) GetAppliedRevision() *int64 {
	return s.AppliedRevision
}

func (s *GetModelGroupResponseBodyDataEntitlementState) GetDesiredRevision() *int64 {
	return s.DesiredRevision
}

func (s *GetModelGroupResponseBodyDataEntitlementState) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetModelGroupResponseBodyDataEntitlementState) GetStatus() *string {
	return s.Status
}

func (s *GetModelGroupResponseBodyDataEntitlementState) SetAppliedRevision(v int64) *GetModelGroupResponseBodyDataEntitlementState {
	s.AppliedRevision = &v
	return s
}

func (s *GetModelGroupResponseBodyDataEntitlementState) SetDesiredRevision(v int64) *GetModelGroupResponseBodyDataEntitlementState {
	s.DesiredRevision = &v
	return s
}

func (s *GetModelGroupResponseBodyDataEntitlementState) SetErrorMessage(v string) *GetModelGroupResponseBodyDataEntitlementState {
	s.ErrorMessage = &v
	return s
}

func (s *GetModelGroupResponseBodyDataEntitlementState) SetStatus(v string) *GetModelGroupResponseBodyDataEntitlementState {
	s.Status = &v
	return s
}

func (s *GetModelGroupResponseBodyDataEntitlementState) Validate() error {
	return dara.Validate(s)
}
