// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetComputeSourceResponseBody
	GetCode() *string
	SetComputeSourceInfo(v *GetComputeSourceResponseBodyComputeSourceInfo) *GetComputeSourceResponseBody
	GetComputeSourceInfo() *GetComputeSourceResponseBodyComputeSourceInfo
	SetHttpStatusCode(v int32) *GetComputeSourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetComputeSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetComputeSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetComputeSourceResponseBody
	GetSuccess() *bool
}

type GetComputeSourceResponseBody struct {
	// example:
	//
	// OK
	Code              *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	ComputeSourceInfo *GetComputeSourceResponseBodyComputeSourceInfo `json:"ComputeSourceInfo,omitempty" xml:"ComputeSourceInfo,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetComputeSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetComputeSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetComputeSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetComputeSourceResponseBody) GetComputeSourceInfo() *GetComputeSourceResponseBodyComputeSourceInfo {
	return s.ComputeSourceInfo
}

func (s *GetComputeSourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetComputeSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetComputeSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetComputeSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetComputeSourceResponseBody) SetCode(v string) *GetComputeSourceResponseBody {
	s.Code = &v
	return s
}

func (s *GetComputeSourceResponseBody) SetComputeSourceInfo(v *GetComputeSourceResponseBodyComputeSourceInfo) *GetComputeSourceResponseBody {
	s.ComputeSourceInfo = v
	return s
}

func (s *GetComputeSourceResponseBody) SetHttpStatusCode(v int32) *GetComputeSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetComputeSourceResponseBody) SetMessage(v string) *GetComputeSourceResponseBody {
	s.Message = &v
	return s
}

func (s *GetComputeSourceResponseBody) SetRequestId(v string) *GetComputeSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetComputeSourceResponseBody) SetSuccess(v bool) *GetComputeSourceResponseBody {
	s.Success = &v
	return s
}

func (s *GetComputeSourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetComputeSourceResponseBodyComputeSourceInfo struct {
	// example:
	//
	// true
	BindProject *bool `json:"BindProject,omitempty" xml:"BindProject,omitempty"`
	// example:
	//
	// 10101101
	BindProjectId *int64 `json:"BindProjectId,omitempty" xml:"BindProjectId,omitempty"`
	// example:
	//
	// dp_mctest
	BindProjectName *string `json:"BindProjectName,omitempty" xml:"BindProjectName,omitempty"`
	// example:
	//
	// 101101
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// test
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// example:
	//
	// compute source for xxx.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// dp_test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 1681881607000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1711881607000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 12356
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// dp_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 101101
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// test
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// MAX_COMPUTE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetComputeSourceResponseBodyComputeSourceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetComputeSourceResponseBodyComputeSourceInfo) GoString() string {
	return s.String()
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) GetBindProject() *bool {
	return s.BindProject
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) GetBindProjectId() *int64 {
	return s.BindProjectId
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) GetBindProjectName() *string {
	return s.BindProjectName
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) GetCreator() *string {
	return s.Creator
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) GetDescription() *string {
	return s.Description
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) GetId() *int64 {
	return s.Id
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) GetName() *string {
	return s.Name
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) GetOwner() *string {
	return s.Owner
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) GetType() *string {
	return s.Type
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) SetBindProject(v bool) *GetComputeSourceResponseBodyComputeSourceInfo {
	s.BindProject = &v
	return s
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) SetBindProjectId(v int64) *GetComputeSourceResponseBodyComputeSourceInfo {
	s.BindProjectId = &v
	return s
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) SetBindProjectName(v string) *GetComputeSourceResponseBodyComputeSourceInfo {
	s.BindProjectName = &v
	return s
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) SetCreator(v string) *GetComputeSourceResponseBodyComputeSourceInfo {
	s.Creator = &v
	return s
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) SetCreatorName(v string) *GetComputeSourceResponseBodyComputeSourceInfo {
	s.CreatorName = &v
	return s
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) SetDescription(v string) *GetComputeSourceResponseBodyComputeSourceInfo {
	s.Description = &v
	return s
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) SetDisplayName(v string) *GetComputeSourceResponseBodyComputeSourceInfo {
	s.DisplayName = &v
	return s
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) SetGmtCreate(v string) *GetComputeSourceResponseBodyComputeSourceInfo {
	s.GmtCreate = &v
	return s
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) SetGmtModified(v string) *GetComputeSourceResponseBodyComputeSourceInfo {
	s.GmtModified = &v
	return s
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) SetId(v int64) *GetComputeSourceResponseBodyComputeSourceInfo {
	s.Id = &v
	return s
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) SetName(v string) *GetComputeSourceResponseBodyComputeSourceInfo {
	s.Name = &v
	return s
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) SetOwner(v string) *GetComputeSourceResponseBodyComputeSourceInfo {
	s.Owner = &v
	return s
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) SetOwnerName(v string) *GetComputeSourceResponseBodyComputeSourceInfo {
	s.OwnerName = &v
	return s
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) SetType(v string) *GetComputeSourceResponseBodyComputeSourceInfo {
	s.Type = &v
	return s
}

func (s *GetComputeSourceResponseBodyComputeSourceInfo) Validate() error {
	return dara.Validate(s)
}
