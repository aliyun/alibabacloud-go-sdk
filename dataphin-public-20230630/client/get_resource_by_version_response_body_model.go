// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceByVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetResourceByVersionResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetResourceByVersionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetResourceByVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetResourceByVersionResponseBody
	GetRequestId() *string
	SetResourceInfo(v *GetResourceByVersionResponseBodyResourceInfo) *GetResourceByVersionResponseBody
	GetResourceInfo() *GetResourceByVersionResponseBodyResourceInfo
	SetSuccess(v bool) *GetResourceByVersionResponseBody
	GetSuccess() *bool
}

type GetResourceByVersionResponseBody struct {
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
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceInfo *GetResourceByVersionResponseBodyResourceInfo `json:"ResourceInfo,omitempty" xml:"ResourceInfo,omitempty" type:"Struct"`
	Success      *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetResourceByVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceByVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceByVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetResourceByVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetResourceByVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetResourceByVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceByVersionResponseBody) GetResourceInfo() *GetResourceByVersionResponseBodyResourceInfo {
	return s.ResourceInfo
}

func (s *GetResourceByVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetResourceByVersionResponseBody) SetCode(v string) *GetResourceByVersionResponseBody {
	s.Code = &v
	return s
}

func (s *GetResourceByVersionResponseBody) SetHttpStatusCode(v int32) *GetResourceByVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetResourceByVersionResponseBody) SetMessage(v string) *GetResourceByVersionResponseBody {
	s.Message = &v
	return s
}

func (s *GetResourceByVersionResponseBody) SetRequestId(v string) *GetResourceByVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceByVersionResponseBody) SetResourceInfo(v *GetResourceByVersionResponseBodyResourceInfo) *GetResourceByVersionResponseBody {
	s.ResourceInfo = v
	return s
}

func (s *GetResourceByVersionResponseBody) SetSuccess(v bool) *GetResourceByVersionResponseBody {
	s.Success = &v
	return s
}

func (s *GetResourceByVersionResponseBody) Validate() error {
	if s.ResourceInfo != nil {
		if err := s.ResourceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceByVersionResponseBodyResourceInfo struct {
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

func (s GetResourceByVersionResponseBodyResourceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetResourceByVersionResponseBodyResourceInfo) GoString() string {
	return s.String()
}

func (s *GetResourceByVersionResponseBodyResourceInfo) GetComputeEngineType() *string {
	return s.ComputeEngineType
}

func (s *GetResourceByVersionResponseBodyResourceInfo) GetCreator() *string {
	return s.Creator
}

func (s *GetResourceByVersionResponseBodyResourceInfo) GetDescription() *string {
	return s.Description
}

func (s *GetResourceByVersionResponseBodyResourceInfo) GetDirectory() *string {
	return s.Directory
}

func (s *GetResourceByVersionResponseBodyResourceInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetResourceByVersionResponseBodyResourceInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetResourceByVersionResponseBodyResourceInfo) GetId() *int64 {
	return s.Id
}

func (s *GetResourceByVersionResponseBodyResourceInfo) GetLastModifier() *string {
	return s.LastModifier
}

func (s *GetResourceByVersionResponseBodyResourceInfo) GetName() *string {
	return s.Name
}

func (s *GetResourceByVersionResponseBodyResourceInfo) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetResourceByVersionResponseBodyResourceInfo) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetResourceByVersionResponseBodyResourceInfo) GetSize() *int64 {
	return s.Size
}

func (s *GetResourceByVersionResponseBodyResourceInfo) GetStorageAddress() *string {
	return s.StorageAddress
}

func (s *GetResourceByVersionResponseBodyResourceInfo) SetComputeEngineType(v string) *GetResourceByVersionResponseBodyResourceInfo {
	s.ComputeEngineType = &v
	return s
}

func (s *GetResourceByVersionResponseBodyResourceInfo) SetCreator(v string) *GetResourceByVersionResponseBodyResourceInfo {
	s.Creator = &v
	return s
}

func (s *GetResourceByVersionResponseBodyResourceInfo) SetDescription(v string) *GetResourceByVersionResponseBodyResourceInfo {
	s.Description = &v
	return s
}

func (s *GetResourceByVersionResponseBodyResourceInfo) SetDirectory(v string) *GetResourceByVersionResponseBodyResourceInfo {
	s.Directory = &v
	return s
}

func (s *GetResourceByVersionResponseBodyResourceInfo) SetGmtCreate(v string) *GetResourceByVersionResponseBodyResourceInfo {
	s.GmtCreate = &v
	return s
}

func (s *GetResourceByVersionResponseBodyResourceInfo) SetGmtModified(v string) *GetResourceByVersionResponseBodyResourceInfo {
	s.GmtModified = &v
	return s
}

func (s *GetResourceByVersionResponseBodyResourceInfo) SetId(v int64) *GetResourceByVersionResponseBodyResourceInfo {
	s.Id = &v
	return s
}

func (s *GetResourceByVersionResponseBodyResourceInfo) SetLastModifier(v string) *GetResourceByVersionResponseBodyResourceInfo {
	s.LastModifier = &v
	return s
}

func (s *GetResourceByVersionResponseBodyResourceInfo) SetName(v string) *GetResourceByVersionResponseBodyResourceInfo {
	s.Name = &v
	return s
}

func (s *GetResourceByVersionResponseBodyResourceInfo) SetProjectId(v int64) *GetResourceByVersionResponseBodyResourceInfo {
	s.ProjectId = &v
	return s
}

func (s *GetResourceByVersionResponseBodyResourceInfo) SetResourceType(v string) *GetResourceByVersionResponseBodyResourceInfo {
	s.ResourceType = &v
	return s
}

func (s *GetResourceByVersionResponseBodyResourceInfo) SetSize(v int64) *GetResourceByVersionResponseBodyResourceInfo {
	s.Size = &v
	return s
}

func (s *GetResourceByVersionResponseBodyResourceInfo) SetStorageAddress(v string) *GetResourceByVersionResponseBodyResourceInfo {
	s.StorageAddress = &v
	return s
}

func (s *GetResourceByVersionResponseBodyResourceInfo) Validate() error {
	return dara.Validate(s)
}
