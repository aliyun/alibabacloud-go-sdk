// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySkillGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelType(v int32) *QuerySkillGroupsRequest
	GetChannelType() *int32
	SetClientToken(v string) *QuerySkillGroupsRequest
	GetClientToken() *string
	SetDepartmentId(v int64) *QuerySkillGroupsRequest
	GetDepartmentId() *int64
	SetInstanceId(v string) *QuerySkillGroupsRequest
	GetInstanceId() *string
	SetPageNo(v int32) *QuerySkillGroupsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QuerySkillGroupsRequest
	GetPageSize() *int32
}

type QuerySkillGroupsRequest struct {
	// The channel type of the skill group. Valid values:
	//
	// - **0**: Returns all skill groups.
	//
	// - **1**: Hotline skill group.
	//
	// - **2**: Online skill group.
	//
	// - **3**: Online and hotline skill group.
	//
	// - **4**: Ticket skill group.
	//
	// - **5**: Hotline and ticket skill group.
	//
	// - **6**: Online and ticket skill group.
	//
	// - **7**: Online, hotline, and ticket skill group.
	//
	// example:
	//
	// 2
	ChannelType *int32 `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// A unique ID for the customer request. Used for idempotency validation. You can generate it using UUID.
	//
	// example:
	//
	// 46c1341e-2648-447a-****-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The department ID.
	//
	// example:
	//
	// 1023****
	DepartmentId *int64 `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// The Artificial Intelligence Cloud Call Service (AICCS) instance ID.
	//
	// You can obtain it from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The current page number. The value must be greater than **0**. Default Value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Page size. The value must be greater than **0**. Default value: **20**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QuerySkillGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySkillGroupsRequest) GoString() string {
	return s.String()
}

func (s *QuerySkillGroupsRequest) GetChannelType() *int32 {
	return s.ChannelType
}

func (s *QuerySkillGroupsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *QuerySkillGroupsRequest) GetDepartmentId() *int64 {
	return s.DepartmentId
}

func (s *QuerySkillGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QuerySkillGroupsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QuerySkillGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySkillGroupsRequest) SetChannelType(v int32) *QuerySkillGroupsRequest {
	s.ChannelType = &v
	return s
}

func (s *QuerySkillGroupsRequest) SetClientToken(v string) *QuerySkillGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *QuerySkillGroupsRequest) SetDepartmentId(v int64) *QuerySkillGroupsRequest {
	s.DepartmentId = &v
	return s
}

func (s *QuerySkillGroupsRequest) SetInstanceId(v string) *QuerySkillGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *QuerySkillGroupsRequest) SetPageNo(v int32) *QuerySkillGroupsRequest {
	s.PageNo = &v
	return s
}

func (s *QuerySkillGroupsRequest) SetPageSize(v int32) *QuerySkillGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySkillGroupsRequest) Validate() error {
	return dara.Validate(s)
}
