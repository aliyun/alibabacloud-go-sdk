// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetResourceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetResourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetResourceResponseBody
	GetRequestId() *string
	SetResourceInfo(v *GetResourceResponseBodyResourceInfo) *GetResourceResponseBody
	GetResourceInfo() *GetResourceResponseBodyResourceInfo
	SetSuccess(v bool) *GetResourceResponseBody
	GetSuccess() *bool
}

type GetResourceResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceInfo *GetResourceResponseBodyResourceInfo `json:"ResourceInfo,omitempty" xml:"ResourceInfo,omitempty" type:"Struct"`
	Success      *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetResourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceResponseBody) GetResourceInfo() *GetResourceResponseBodyResourceInfo {
	return s.ResourceInfo
}

func (s *GetResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetResourceResponseBody) SetCode(v string) *GetResourceResponseBody {
	s.Code = &v
	return s
}

func (s *GetResourceResponseBody) SetHttpStatusCode(v int32) *GetResourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetResourceResponseBody) SetMessage(v string) *GetResourceResponseBody {
	s.Message = &v
	return s
}

func (s *GetResourceResponseBody) SetRequestId(v string) *GetResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceResponseBody) SetResourceInfo(v *GetResourceResponseBodyResourceInfo) *GetResourceResponseBody {
	s.ResourceInfo = v
	return s
}

func (s *GetResourceResponseBody) SetSuccess(v bool) *GetResourceResponseBody {
	s.Success = &v
	return s
}

func (s *GetResourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetResourceResponseBodyResourceInfo struct {
	// example:
	//
	// MAX_COMPUTE
	ComputeEngineType *string `json:"ComputeEngineType,omitempty" xml:"ComputeEngineType,omitempty"`
	// example:
	//
	// 30011021
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// /
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// example:
	//
	// 2025-06-10 10:01:01
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2025-06-10 10:01:01
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 10200121011
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 30011021
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// example:
	//
	// udf_sleep.jar
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1030111021
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// JAR
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 102400
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 300011448/711833/cdcd1c44-f1ee-46bb-b318-1d123dbabf6c
	StorageAddress *string `json:"StorageAddress,omitempty" xml:"StorageAddress,omitempty"`
}

func (s GetResourceResponseBodyResourceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetResourceResponseBodyResourceInfo) GoString() string {
	return s.String()
}

func (s *GetResourceResponseBodyResourceInfo) GetComputeEngineType() *string {
	return s.ComputeEngineType
}

func (s *GetResourceResponseBodyResourceInfo) GetCreator() *string {
	return s.Creator
}

func (s *GetResourceResponseBodyResourceInfo) GetDescription() *string {
	return s.Description
}

func (s *GetResourceResponseBodyResourceInfo) GetDirectory() *string {
	return s.Directory
}

func (s *GetResourceResponseBodyResourceInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetResourceResponseBodyResourceInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetResourceResponseBodyResourceInfo) GetId() *int64 {
	return s.Id
}

func (s *GetResourceResponseBodyResourceInfo) GetLastModifier() *string {
	return s.LastModifier
}

func (s *GetResourceResponseBodyResourceInfo) GetName() *string {
	return s.Name
}

func (s *GetResourceResponseBodyResourceInfo) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetResourceResponseBodyResourceInfo) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetResourceResponseBodyResourceInfo) GetSize() *int64 {
	return s.Size
}

func (s *GetResourceResponseBodyResourceInfo) GetStorageAddress() *string {
	return s.StorageAddress
}

func (s *GetResourceResponseBodyResourceInfo) SetComputeEngineType(v string) *GetResourceResponseBodyResourceInfo {
	s.ComputeEngineType = &v
	return s
}

func (s *GetResourceResponseBodyResourceInfo) SetCreator(v string) *GetResourceResponseBodyResourceInfo {
	s.Creator = &v
	return s
}

func (s *GetResourceResponseBodyResourceInfo) SetDescription(v string) *GetResourceResponseBodyResourceInfo {
	s.Description = &v
	return s
}

func (s *GetResourceResponseBodyResourceInfo) SetDirectory(v string) *GetResourceResponseBodyResourceInfo {
	s.Directory = &v
	return s
}

func (s *GetResourceResponseBodyResourceInfo) SetGmtCreate(v string) *GetResourceResponseBodyResourceInfo {
	s.GmtCreate = &v
	return s
}

func (s *GetResourceResponseBodyResourceInfo) SetGmtModified(v string) *GetResourceResponseBodyResourceInfo {
	s.GmtModified = &v
	return s
}

func (s *GetResourceResponseBodyResourceInfo) SetId(v int64) *GetResourceResponseBodyResourceInfo {
	s.Id = &v
	return s
}

func (s *GetResourceResponseBodyResourceInfo) SetLastModifier(v string) *GetResourceResponseBodyResourceInfo {
	s.LastModifier = &v
	return s
}

func (s *GetResourceResponseBodyResourceInfo) SetName(v string) *GetResourceResponseBodyResourceInfo {
	s.Name = &v
	return s
}

func (s *GetResourceResponseBodyResourceInfo) SetProjectId(v int64) *GetResourceResponseBodyResourceInfo {
	s.ProjectId = &v
	return s
}

func (s *GetResourceResponseBodyResourceInfo) SetResourceType(v string) *GetResourceResponseBodyResourceInfo {
	s.ResourceType = &v
	return s
}

func (s *GetResourceResponseBodyResourceInfo) SetSize(v int64) *GetResourceResponseBodyResourceInfo {
	s.Size = &v
	return s
}

func (s *GetResourceResponseBodyResourceInfo) SetStorageAddress(v string) *GetResourceResponseBodyResourceInfo {
	s.StorageAddress = &v
	return s
}

func (s *GetResourceResponseBodyResourceInfo) Validate() error {
	return dara.Validate(s)
}
