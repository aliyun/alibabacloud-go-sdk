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
	// List of department IDs. Separate multiple IDs with commas (,).
	//
	// You can call the [GetAllDepartment](https://help.aliyun.com/document_detail/2717975.html) API and check the **DepartmentId*	- field in the response to obtain department IDs.
	//
	// > When this parameter is not empty:
	//
	// > - If GroupIds is not empty, the system prioritizes querying data metrics for the skill groups specified by GroupIds.
	//
	// > - If GroupIds is empty, the system prioritizes querying data metrics for the departments specified by this parameter.
	//
	// example:
	//
	// 2332****,2334****
	DepIds *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// List of skill group IDs. Separate multiple IDs with commas (,).
	//
	// You can call the [QuerySkillGroups](https://help.aliyun.com/document_detail/2717970.html) API and check the **SkillGroupId*	- field in the response to obtain skill group IDs.
	//
	// > When this parameter is not empty, the system prioritizes querying data metrics for the specified skill groups.
	//
	// example:
	//
	// 2323****,2324****
	GroupIds *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	// Artificial Intelligence Cloud Call Service (AICCS) instance ID.
	//
	// You can obtain it from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
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
