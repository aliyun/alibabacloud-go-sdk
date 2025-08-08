// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetWorkspaceResponseBodyData) *GetWorkspaceResponseBody
	GetData() *GetWorkspaceResponseBodyData
	SetMessage(v string) *GetWorkspaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkspaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkspaceResponseBody
	GetSuccess() *bool
}

type GetWorkspaceResponseBody struct {
	Data *GetWorkspaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// UnknownError
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FE8EE2F1-4880-46BC-A704-5CF63EAF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBody) GetData() *GetWorkspaceResponseBodyData {
	return s.Data
}

func (s *GetWorkspaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkspaceResponseBody) SetData(v *GetWorkspaceResponseBodyData) *GetWorkspaceResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkspaceResponseBody) SetMessage(v string) *GetWorkspaceResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetRequestId(v string) *GetWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetSuccess(v bool) *GetWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetWorkspaceResponseBodyData struct {
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 12345****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 12345****
	ServiceAccountId *string `json:"ServiceAccountId,omitempty" xml:"ServiceAccountId,omitempty"`
	// example:
	//
	// 3322****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// example:
	//
	// vpc-xxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// 863020290155****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// workspace_xxx
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s GetWorkspaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetWorkspaceResponseBodyData) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetWorkspaceResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetWorkspaceResponseBodyData) GetServiceAccountId() *string {
	return s.ServiceAccountId
}

func (s *GetWorkspaceResponseBodyData) GetTid() *int64 {
	return s.Tid
}

func (s *GetWorkspaceResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *GetWorkspaceResponseBodyData) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *GetWorkspaceResponseBodyData) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *GetWorkspaceResponseBodyData) SetDescription(v string) *GetWorkspaceResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetWorkspaceResponseBodyData) SetOwnerId(v string) *GetWorkspaceResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *GetWorkspaceResponseBodyData) SetRegionId(v string) *GetWorkspaceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetWorkspaceResponseBodyData) SetServiceAccountId(v string) *GetWorkspaceResponseBodyData {
	s.ServiceAccountId = &v
	return s
}

func (s *GetWorkspaceResponseBodyData) SetTid(v int64) *GetWorkspaceResponseBodyData {
	s.Tid = &v
	return s
}

func (s *GetWorkspaceResponseBodyData) SetVpcId(v string) *GetWorkspaceResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *GetWorkspaceResponseBodyData) SetWorkspaceId(v int64) *GetWorkspaceResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetWorkspaceResponseBodyData) SetWorkspaceName(v string) *GetWorkspaceResponseBodyData {
	s.WorkspaceName = &v
	return s
}

func (s *GetWorkspaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
