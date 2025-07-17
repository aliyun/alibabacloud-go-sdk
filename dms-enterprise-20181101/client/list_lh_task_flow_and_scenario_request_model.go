// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLhTaskFlowAndScenarioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSpaceId(v int64) *ListLhTaskFlowAndScenarioRequest
	GetSpaceId() *int64
	SetTid(v int64) *ListLhTaskFlowAndScenarioRequest
	GetTid() *int64
	SetUserId(v int64) *ListLhTaskFlowAndScenarioRequest
	GetUserId() *int64
}

type ListLhTaskFlowAndScenarioRequest struct {
	// The ID of the workspace. You can call the [GetLhSpaceByName](https://help.aliyun.com/document_detail/424379.html) operation to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 24
	SpaceId *int64 `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The ID of the user. You can call the [ListUsers](https://help.aliyun.com/document_detail/141938.html) or [GetUser](https://help.aliyun.com/document_detail/147098.html) operation to obtain the user ID.
	//
	// example:
	//
	// 51****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListLhTaskFlowAndScenarioRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLhTaskFlowAndScenarioRequest) GoString() string {
	return s.String()
}

func (s *ListLhTaskFlowAndScenarioRequest) GetSpaceId() *int64 {
	return s.SpaceId
}

func (s *ListLhTaskFlowAndScenarioRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListLhTaskFlowAndScenarioRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *ListLhTaskFlowAndScenarioRequest) SetSpaceId(v int64) *ListLhTaskFlowAndScenarioRequest {
	s.SpaceId = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioRequest) SetTid(v int64) *ListLhTaskFlowAndScenarioRequest {
	s.Tid = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioRequest) SetUserId(v int64) *ListLhTaskFlowAndScenarioRequest {
	s.UserId = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioRequest) Validate() error {
	return dara.Validate(s)
}
