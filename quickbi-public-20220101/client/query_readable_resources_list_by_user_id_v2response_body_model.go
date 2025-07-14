// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReadableResourcesListByUserIdV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryReadableResourcesListByUserIdV2ResponseBody
	GetRequestId() *string
	SetResult(v []*QueryReadableResourcesListByUserIdV2ResponseBodyResult) *QueryReadableResourcesListByUserIdV2ResponseBody
	GetResult() []*QueryReadableResourcesListByUserIdV2ResponseBodyResult
	SetSuccess(v bool) *QueryReadableResourcesListByUserIdV2ResponseBody
	GetSuccess() *bool
}

type QueryReadableResourcesListByUserIdV2ResponseBody struct {
	// example:
	//
	// D787E1A3-A93C********05DF8D885
	RequestId *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryReadableResourcesListByUserIdV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryReadableResourcesListByUserIdV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryReadableResourcesListByUserIdV2ResponseBody) GoString() string {
	return s.String()
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBody) GetResult() []*QueryReadableResourcesListByUserIdV2ResponseBodyResult {
	return s.Result
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBody) SetRequestId(v string) *QueryReadableResourcesListByUserIdV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBody) SetResult(v []*QueryReadableResourcesListByUserIdV2ResponseBodyResult) *QueryReadableResourcesListByUserIdV2ResponseBody {
	s.Result = v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBody) SetSuccess(v bool) *QueryReadableResourcesListByUserIdV2ResponseBody {
	s.Success = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryReadableResourcesListByUserIdV2ResponseBodyResult struct {
	// example:
	//
	// 1611023338000
	CreateTime  *string                                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	Directory   *QueryReadableResourcesListByUserIdV2ResponseBodyResultDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	ModifyName  *string                                                          `json:"ModifyName,omitempty" xml:"ModifyName,omitempty"`
	// example:
	//
	// 1530078690000
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
	OwnerId   *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 0
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	ThirdPartAuthFlag *int32  `json:"ThirdPartAuthFlag,omitempty" xml:"ThirdPartAuthFlag,omitempty"`
	WorkName          *string `json:"WorkName,omitempty" xml:"WorkName,omitempty"`
	// example:
	//
	// PAGE
	WorkType *string `json:"WorkType,omitempty" xml:"WorkType,omitempty"`
	// example:
	//
	// ccd3428c-****-****-a608-26bae29dffee
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
	// example:
	//
	// c5f86ad2-ef53-4c51-8720-162ecfdb****
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryReadableResourcesListByUserIdV2ResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryReadableResourcesListByUserIdV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) GetDirectory() *QueryReadableResourcesListByUserIdV2ResponseBodyResultDirectory {
	return s.Directory
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) GetModifyName() *string {
	return s.ModifyName
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) GetOwnerId() *string {
	return s.OwnerId
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) GetOwnerName() *string {
	return s.OwnerName
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) GetSecurityLevel() *string {
	return s.SecurityLevel
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) GetThirdPartAuthFlag() *int32 {
	return s.ThirdPartAuthFlag
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) GetWorkName() *string {
	return s.WorkName
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) GetWorkType() *string {
	return s.WorkType
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) GetWorksId() *string {
	return s.WorksId
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) SetCreateTime(v string) *QueryReadableResourcesListByUserIdV2ResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) SetDescription(v string) *QueryReadableResourcesListByUserIdV2ResponseBodyResult {
	s.Description = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) SetDirectory(v *QueryReadableResourcesListByUserIdV2ResponseBodyResultDirectory) *QueryReadableResourcesListByUserIdV2ResponseBodyResult {
	s.Directory = v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) SetModifyName(v string) *QueryReadableResourcesListByUserIdV2ResponseBodyResult {
	s.ModifyName = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) SetModifyTime(v string) *QueryReadableResourcesListByUserIdV2ResponseBodyResult {
	s.ModifyTime = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) SetOwnerId(v string) *QueryReadableResourcesListByUserIdV2ResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) SetOwnerName(v string) *QueryReadableResourcesListByUserIdV2ResponseBodyResult {
	s.OwnerName = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) SetSecurityLevel(v string) *QueryReadableResourcesListByUserIdV2ResponseBodyResult {
	s.SecurityLevel = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) SetStatus(v int32) *QueryReadableResourcesListByUserIdV2ResponseBodyResult {
	s.Status = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) SetThirdPartAuthFlag(v int32) *QueryReadableResourcesListByUserIdV2ResponseBodyResult {
	s.ThirdPartAuthFlag = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) SetWorkName(v string) *QueryReadableResourcesListByUserIdV2ResponseBodyResult {
	s.WorkName = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) SetWorkType(v string) *QueryReadableResourcesListByUserIdV2ResponseBodyResult {
	s.WorkType = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) SetWorksId(v string) *QueryReadableResourcesListByUserIdV2ResponseBodyResult {
	s.WorksId = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) SetWorkspaceId(v string) *QueryReadableResourcesListByUserIdV2ResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) SetWorkspaceName(v string) *QueryReadableResourcesListByUserIdV2ResponseBodyResult {
	s.WorkspaceName = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type QueryReadableResourcesListByUserIdV2ResponseBodyResultDirectory struct {
	// example:
	//
	// a3eecab7-618d-4f9f-*****
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 88b680****
	PathId   *string `json:"PathId,omitempty" xml:"PathId,omitempty"`
	PathName *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
}

func (s QueryReadableResourcesListByUserIdV2ResponseBodyResultDirectory) String() string {
	return dara.Prettify(s)
}

func (s QueryReadableResourcesListByUserIdV2ResponseBodyResultDirectory) GoString() string {
	return s.String()
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResultDirectory) GetId() *string {
	return s.Id
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResultDirectory) GetName() *string {
	return s.Name
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResultDirectory) GetPathId() *string {
	return s.PathId
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResultDirectory) GetPathName() *string {
	return s.PathName
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResultDirectory) SetId(v string) *QueryReadableResourcesListByUserIdV2ResponseBodyResultDirectory {
	s.Id = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResultDirectory) SetName(v string) *QueryReadableResourcesListByUserIdV2ResponseBodyResultDirectory {
	s.Name = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResultDirectory) SetPathId(v string) *QueryReadableResourcesListByUserIdV2ResponseBodyResultDirectory {
	s.PathId = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResultDirectory) SetPathName(v string) *QueryReadableResourcesListByUserIdV2ResponseBodyResultDirectory {
	s.PathName = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2ResponseBodyResultDirectory) Validate() error {
	return dara.Validate(s)
}
