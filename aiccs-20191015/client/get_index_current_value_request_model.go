// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIndexCurrentValueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDepIds(v string) *GetIndexCurrentValueRequest
	GetDepIds() *string
	SetGroupIds(v string) *GetIndexCurrentValueRequest
	GetGroupIds() *string
	SetInstanceId(v string) *GetIndexCurrentValueRequest
	GetInstanceId() *string
}

type GetIndexCurrentValueRequest struct {
	// The list of department IDs. Separate multiple IDs with commas (,).
	//
	// Call the [GetAllDepartment](https://help.aliyun.com/document_detail/2717975.html) operation and check the **DepartmentId*	- parameter in the response to obtain the department ID.
	//
	// > When this parameter is not empty:
	//
	// > - If GroupIds is not empty, the query is performed based on the skill group list corresponding to GroupIds.
	//
	// > - If GroupIds is empty, the query is performed based on the department list corresponding to this parameter.
	//
	// example:
	//
	// 2332****,2334****
	DepIds *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// The list of skill group IDs. Separate multiple IDs with commas (,).
	//
	// Call the [QuerySkillGroups](https://help.aliyun.com/document_detail/2717970.html) operation and check the **SkillGroupId*	- parameter in the response to obtain the skill group ID.
	//
	// > When this parameter is not empty, the query is performed based on the skill group list corresponding to this parameter.
	//
	// example:
	//
	// 2323****,2324****
	GroupIds *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	// The AICCS instance ID. You can obtain the instance ID from <b>Instance Management</b> in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// > The AICCS instance ID is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetIndexCurrentValueRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIndexCurrentValueRequest) GoString() string {
	return s.String()
}

func (s *GetIndexCurrentValueRequest) GetDepIds() *string {
	return s.DepIds
}

func (s *GetIndexCurrentValueRequest) GetGroupIds() *string {
	return s.GroupIds
}

func (s *GetIndexCurrentValueRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetIndexCurrentValueRequest) SetDepIds(v string) *GetIndexCurrentValueRequest {
	s.DepIds = &v
	return s
}

func (s *GetIndexCurrentValueRequest) SetGroupIds(v string) *GetIndexCurrentValueRequest {
	s.GroupIds = &v
	return s
}

func (s *GetIndexCurrentValueRequest) SetInstanceId(v string) *GetIndexCurrentValueRequest {
	s.InstanceId = &v
	return s
}

func (s *GetIndexCurrentValueRequest) Validate() error {
	return dara.Validate(s)
}
