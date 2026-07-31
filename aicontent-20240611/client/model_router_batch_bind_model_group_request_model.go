// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBatchBindModelGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedModelGroupConfig(v string) *ModelRouterBatchBindModelGroupRequest
	GetAllowedModelGroupConfig() *string
	SetClientIdList(v []*int64) *ModelRouterBatchBindModelGroupRequest
	GetClientIdList() []*int64
}

type ModelRouterBatchBindModelGroupRequest struct {
	// The authorization configuration (JSON string). Internal key names use a fixed underscore style: {"model_ids":[...],"group_ids":["mg_xxx"]}
	//
	// This parameter is required.
	//
	// example:
	//
	// {"model_ids":[],"group_ids":["mg_pro"]}
	AllowedModelGroupConfig *string `json:"allowedModelGroupConfig,omitempty" xml:"allowedModelGroupConfig,omitempty"`
	// The array of department IDs. You can specify 1 to 50 IDs. If more than 50, call this operation in batches.
	//
	// This parameter is required.
	//
	// example:
	//
	// [1001, 1002, 1003]
	ClientIdList []*int64 `json:"clientIdList,omitempty" xml:"clientIdList,omitempty" type:"Repeated"`
}

func (s ModelRouterBatchBindModelGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBatchBindModelGroupRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterBatchBindModelGroupRequest) GetAllowedModelGroupConfig() *string {
	return s.AllowedModelGroupConfig
}

func (s *ModelRouterBatchBindModelGroupRequest) GetClientIdList() []*int64 {
	return s.ClientIdList
}

func (s *ModelRouterBatchBindModelGroupRequest) SetAllowedModelGroupConfig(v string) *ModelRouterBatchBindModelGroupRequest {
	s.AllowedModelGroupConfig = &v
	return s
}

func (s *ModelRouterBatchBindModelGroupRequest) SetClientIdList(v []*int64) *ModelRouterBatchBindModelGroupRequest {
	s.ClientIdList = v
	return s
}

func (s *ModelRouterBatchBindModelGroupRequest) Validate() error {
	return dara.Validate(s)
}
