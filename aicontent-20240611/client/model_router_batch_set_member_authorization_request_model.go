// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBatchSetMemberAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedModelGroupConfig(v string) *ModelRouterBatchSetMemberAuthorizationRequest
	GetAllowedModelGroupConfig() *string
	SetUserIdList(v []*int64) *ModelRouterBatchSetMemberAuthorizationRequest
	GetUserIdList() []*int64
}

type ModelRouterBatchSetMemberAuthorizationRequest struct {
	// The authorization configuration. This parameter is a required JSON string and uses overwrite mode. Format: {"model_ids":[...],"group_ids":["mg_xxx"]}. Internal key names use a fixed underscore style and are not converted to the camelCase convention of the API.
	//
	// example:
	//
	// {"model_ids":[],"group_ids":["mg_qwen_19"]}
	AllowedModelGroupConfig *string `json:"allowedModelGroupConfig,omitempty" xml:"allowedModelGroupConfig,omitempty"`
	// The list of user IDs. This parameter is required. You can specify 1 to 50 user IDs. If more than 50 user IDs are required, call this operation in batches. All specified users must be direct members of the department.
	//
	// example:
	//
	// [30001,30002,30003]
	UserIdList []*int64 `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s ModelRouterBatchSetMemberAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBatchSetMemberAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterBatchSetMemberAuthorizationRequest) GetAllowedModelGroupConfig() *string {
	return s.AllowedModelGroupConfig
}

func (s *ModelRouterBatchSetMemberAuthorizationRequest) GetUserIdList() []*int64 {
	return s.UserIdList
}

func (s *ModelRouterBatchSetMemberAuthorizationRequest) SetAllowedModelGroupConfig(v string) *ModelRouterBatchSetMemberAuthorizationRequest {
	s.AllowedModelGroupConfig = &v
	return s
}

func (s *ModelRouterBatchSetMemberAuthorizationRequest) SetUserIdList(v []*int64) *ModelRouterBatchSetMemberAuthorizationRequest {
	s.UserIdList = v
	return s
}

func (s *ModelRouterBatchSetMemberAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
