// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroup(v *GetDataServiceGroupResponseBodyGroup) *GetDataServiceGroupResponseBody
	GetGroup() *GetDataServiceGroupResponseBodyGroup
	SetRequestId(v string) *GetDataServiceGroupResponseBody
	GetRequestId() *string
}

type GetDataServiceGroupResponseBody struct {
	// The details of the business process.
	Group *GetDataServiceGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataServiceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataServiceGroupResponseBody) GetGroup() *GetDataServiceGroupResponseBodyGroup {
	return s.Group
}

func (s *GetDataServiceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataServiceGroupResponseBody) SetGroup(v *GetDataServiceGroupResponseBodyGroup) *GetDataServiceGroupResponseBody {
	s.Group = v
	return s
}

func (s *GetDataServiceGroupResponseBody) SetRequestId(v string) *GetDataServiceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataServiceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataServiceGroupResponseBodyGroup struct {
	// The ID of the API group that is associated with the business process in the API Gateway console.
	//
	// example:
	//
	// 100abc
	ApiGatewayGroupId *string `json:"ApiGatewayGroupId,omitempty" xml:"ApiGatewayGroupId,omitempty"`
	// The time when the business process was created.
	//
	// example:
	//
	// 2020-09-24T18:37:51+0800
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The user identifier (UID) of the creator of the business process. The value of this parameter may be empty for creators of some existing business processes.
	//
	// example:
	//
	// 10001
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The description of the business process.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The business process ID.
	//
	// example:
	//
	// ds_123abc
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the business process.
	//
	// example:
	//
	// Test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The time when the business process was last modified.
	//
	// example:
	//
	// 2020-09-24T18:37:51+0800
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 10002
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10003
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetDataServiceGroupResponseBodyGroup) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *GetDataServiceGroupResponseBodyGroup) GetApiGatewayGroupId() *string {
	return s.ApiGatewayGroupId
}

func (s *GetDataServiceGroupResponseBodyGroup) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *GetDataServiceGroupResponseBodyGroup) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetDataServiceGroupResponseBodyGroup) GetDescription() *string {
	return s.Description
}

func (s *GetDataServiceGroupResponseBodyGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *GetDataServiceGroupResponseBodyGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *GetDataServiceGroupResponseBodyGroup) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetDataServiceGroupResponseBodyGroup) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDataServiceGroupResponseBodyGroup) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetDataServiceGroupResponseBodyGroup) SetApiGatewayGroupId(v string) *GetDataServiceGroupResponseBodyGroup {
	s.ApiGatewayGroupId = &v
	return s
}

func (s *GetDataServiceGroupResponseBodyGroup) SetCreatedTime(v string) *GetDataServiceGroupResponseBodyGroup {
	s.CreatedTime = &v
	return s
}

func (s *GetDataServiceGroupResponseBodyGroup) SetCreatorId(v string) *GetDataServiceGroupResponseBodyGroup {
	s.CreatorId = &v
	return s
}

func (s *GetDataServiceGroupResponseBodyGroup) SetDescription(v string) *GetDataServiceGroupResponseBodyGroup {
	s.Description = &v
	return s
}

func (s *GetDataServiceGroupResponseBodyGroup) SetGroupId(v string) *GetDataServiceGroupResponseBodyGroup {
	s.GroupId = &v
	return s
}

func (s *GetDataServiceGroupResponseBodyGroup) SetGroupName(v string) *GetDataServiceGroupResponseBodyGroup {
	s.GroupName = &v
	return s
}

func (s *GetDataServiceGroupResponseBodyGroup) SetModifiedTime(v string) *GetDataServiceGroupResponseBodyGroup {
	s.ModifiedTime = &v
	return s
}

func (s *GetDataServiceGroupResponseBodyGroup) SetProjectId(v int64) *GetDataServiceGroupResponseBodyGroup {
	s.ProjectId = &v
	return s
}

func (s *GetDataServiceGroupResponseBodyGroup) SetTenantId(v int64) *GetDataServiceGroupResponseBodyGroup {
	s.TenantId = &v
	return s
}

func (s *GetDataServiceGroupResponseBodyGroup) Validate() error {
	return dara.Validate(s)
}
