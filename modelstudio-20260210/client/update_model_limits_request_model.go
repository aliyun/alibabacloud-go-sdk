// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelLimitsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *UpdateModelLimitsRequest
	GetWorkspaceId() *string
	SetWorkspaceLimits(v []*UpdateModelLimitsRequestWorkspaceLimits) *UpdateModelLimitsRequest
	GetWorkspaceLimits() []*UpdateModelLimitsRequestWorkspaceLimits
}

type UpdateModelLimitsRequest struct {
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-ac3ef438bec22dc5
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// The throttling values for the workspace.
	WorkspaceLimits []*UpdateModelLimitsRequestWorkspaceLimits `json:"workspaceLimits,omitempty" xml:"workspaceLimits,omitempty" type:"Repeated"`
}

func (s UpdateModelLimitsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelLimitsRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelLimitsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateModelLimitsRequest) GetWorkspaceLimits() []*UpdateModelLimitsRequestWorkspaceLimits {
	return s.WorkspaceLimits
}

func (s *UpdateModelLimitsRequest) SetWorkspaceId(v string) *UpdateModelLimitsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateModelLimitsRequest) SetWorkspaceLimits(v []*UpdateModelLimitsRequestWorkspaceLimits) *UpdateModelLimitsRequest {
	s.WorkspaceLimits = v
	return s
}

func (s *UpdateModelLimitsRequest) Validate() error {
	if s.WorkspaceLimits != nil {
		for _, item := range s.WorkspaceLimits {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateModelLimitsRequestWorkspaceLimits struct {
	// The model.
	//
	// example:
	//
	// qwen-plus
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The throttling operation type. Valid values:
	//
	// - **OVERLAY**: Sets or overwrites the throttling configuration.
	//
	// - **DELETE**: Deletes the throttling configuration (restores to no throttling).
	//
	// example:
	//
	// OVERLAY
	OperationType *string `json:"operationType,omitempty" xml:"operationType,omitempty"`
	// The request throttling value.
	//
	// example:
	//
	// 10
	RequestLimit *int64 `json:"requestLimit,omitempty" xml:"requestLimit,omitempty"`
	// The time period for request throttling. Unit: seconds.
	//
	// example:
	//
	// 1
	RequestLimitPeriod *int64 `json:"requestLimitPeriod,omitempty" xml:"requestLimitPeriod,omitempty"`
	// The usage throttling value.
	//
	// example:
	//
	// 10
	UsageLimit *int64 `json:"usageLimit,omitempty" xml:"usageLimit,omitempty"`
	// The time period for usage throttling. Unit: seconds.
	//
	// example:
	//
	// 1
	UsageLimitPeriod *int64 `json:"usageLimitPeriod,omitempty" xml:"usageLimitPeriod,omitempty"`
}

func (s UpdateModelLimitsRequestWorkspaceLimits) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelLimitsRequestWorkspaceLimits) GoString() string {
	return s.String()
}

func (s *UpdateModelLimitsRequestWorkspaceLimits) GetModel() *string {
	return s.Model
}

func (s *UpdateModelLimitsRequestWorkspaceLimits) GetOperationType() *string {
	return s.OperationType
}

func (s *UpdateModelLimitsRequestWorkspaceLimits) GetRequestLimit() *int64 {
	return s.RequestLimit
}

func (s *UpdateModelLimitsRequestWorkspaceLimits) GetRequestLimitPeriod() *int64 {
	return s.RequestLimitPeriod
}

func (s *UpdateModelLimitsRequestWorkspaceLimits) GetUsageLimit() *int64 {
	return s.UsageLimit
}

func (s *UpdateModelLimitsRequestWorkspaceLimits) GetUsageLimitPeriod() *int64 {
	return s.UsageLimitPeriod
}

func (s *UpdateModelLimitsRequestWorkspaceLimits) SetModel(v string) *UpdateModelLimitsRequestWorkspaceLimits {
	s.Model = &v
	return s
}

func (s *UpdateModelLimitsRequestWorkspaceLimits) SetOperationType(v string) *UpdateModelLimitsRequestWorkspaceLimits {
	s.OperationType = &v
	return s
}

func (s *UpdateModelLimitsRequestWorkspaceLimits) SetRequestLimit(v int64) *UpdateModelLimitsRequestWorkspaceLimits {
	s.RequestLimit = &v
	return s
}

func (s *UpdateModelLimitsRequestWorkspaceLimits) SetRequestLimitPeriod(v int64) *UpdateModelLimitsRequestWorkspaceLimits {
	s.RequestLimitPeriod = &v
	return s
}

func (s *UpdateModelLimitsRequestWorkspaceLimits) SetUsageLimit(v int64) *UpdateModelLimitsRequestWorkspaceLimits {
	s.UsageLimit = &v
	return s
}

func (s *UpdateModelLimitsRequestWorkspaceLimits) SetUsageLimitPeriod(v int64) *UpdateModelLimitsRequestWorkspaceLimits {
	s.UsageLimitPeriod = &v
	return s
}

func (s *UpdateModelLimitsRequestWorkspaceLimits) Validate() error {
	return dara.Validate(s)
}
