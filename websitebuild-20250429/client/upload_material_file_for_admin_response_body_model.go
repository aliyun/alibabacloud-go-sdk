// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadMaterialFileForAdminResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UploadMaterialFileForAdminResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *UploadMaterialFileForAdminResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *UploadMaterialFileForAdminResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *UploadMaterialFileForAdminResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UploadMaterialFileForAdminResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *UploadMaterialFileForAdminResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *UploadMaterialFileForAdminResponseBodyModule) *UploadMaterialFileForAdminResponseBody
	GetModule() *UploadMaterialFileForAdminResponseBodyModule
	SetRequestId(v string) *UploadMaterialFileForAdminResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *UploadMaterialFileForAdminResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *UploadMaterialFileForAdminResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *UploadMaterialFileForAdminResponseBody
	GetSynchro() *bool
}

type UploadMaterialFileForAdminResponseBody struct {
	// The detailed reason why access is denied.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether a retry is allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The application name.
	//
	// example:
	//
	// or
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic message.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The file object.
	Module *UploadMaterialFileForAdminResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The error code.
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// System exception
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Indicates whether the request is processed synchronously.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s UploadMaterialFileForAdminResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadMaterialFileForAdminResponseBody) GoString() string {
	return s.String()
}

func (s *UploadMaterialFileForAdminResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UploadMaterialFileForAdminResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *UploadMaterialFileForAdminResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *UploadMaterialFileForAdminResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UploadMaterialFileForAdminResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UploadMaterialFileForAdminResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *UploadMaterialFileForAdminResponseBody) GetModule() *UploadMaterialFileForAdminResponseBodyModule {
	return s.Module
}

func (s *UploadMaterialFileForAdminResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadMaterialFileForAdminResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *UploadMaterialFileForAdminResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *UploadMaterialFileForAdminResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *UploadMaterialFileForAdminResponseBody) SetAccessDeniedDetail(v string) *UploadMaterialFileForAdminResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBody) SetAllowRetry(v bool) *UploadMaterialFileForAdminResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBody) SetAppName(v string) *UploadMaterialFileForAdminResponseBody {
	s.AppName = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBody) SetDynamicCode(v string) *UploadMaterialFileForAdminResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBody) SetDynamicMessage(v string) *UploadMaterialFileForAdminResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBody) SetErrorArgs(v []interface{}) *UploadMaterialFileForAdminResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *UploadMaterialFileForAdminResponseBody) SetModule(v *UploadMaterialFileForAdminResponseBodyModule) *UploadMaterialFileForAdminResponseBody {
	s.Module = v
	return s
}

func (s *UploadMaterialFileForAdminResponseBody) SetRequestId(v string) *UploadMaterialFileForAdminResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBody) SetRootErrorCode(v string) *UploadMaterialFileForAdminResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBody) SetRootErrorMsg(v string) *UploadMaterialFileForAdminResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBody) SetSynchro(v bool) *UploadMaterialFileForAdminResponseBody {
	s.Synchro = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadMaterialFileForAdminResponseBodyModule struct {
	// The business instance ID.
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The file content type.
	//
	// example:
	//
	// image/png
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The creation time. The time is in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2025-11-03 02:05:01
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the file was moved to the recycle bin or deleted. The time is in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2025-11-03 02:05:01
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// ddddwww
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The unique ID of the file.
	//
	// example:
	//
	// d75e1e9693cc460da218419d67735567
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The file path.
	//
	// example:
	//
	// https://xxx.xxx.cn/original-data/pdf/mndj_report/cd649b2cc2102c0df57bfa1ae62931d6.zip
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The height of the file.
	//
	// example:
	//
	// 1280
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The file name.
	//
	// example:
	//
	// FileName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The file status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The file size.
	//
	// example:
	//
	// 23M
	StorageSize *string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// The file suffix.
	//
	// example:
	//
	// png
	Suffix *string `json:"Suffix,omitempty" xml:"Suffix,omitempty"`
	// The file type.
	//
	// example:
	//
	// IMAGE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The width of the file.
	//
	// example:
	//
	// 1920
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s UploadMaterialFileForAdminResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s UploadMaterialFileForAdminResponseBodyModule) GoString() string {
	return s.String()
}

func (s *UploadMaterialFileForAdminResponseBodyModule) GetBizId() *string {
	return s.BizId
}

func (s *UploadMaterialFileForAdminResponseBodyModule) GetContentType() *string {
	return s.ContentType
}

func (s *UploadMaterialFileForAdminResponseBodyModule) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UploadMaterialFileForAdminResponseBodyModule) GetDeletedTime() *string {
	return s.DeletedTime
}

func (s *UploadMaterialFileForAdminResponseBodyModule) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UploadMaterialFileForAdminResponseBodyModule) GetFileId() *string {
	return s.FileId
}

func (s *UploadMaterialFileForAdminResponseBodyModule) GetFileUrl() *string {
	return s.FileUrl
}

func (s *UploadMaterialFileForAdminResponseBodyModule) GetHeight() *int32 {
	return s.Height
}

func (s *UploadMaterialFileForAdminResponseBodyModule) GetName() *string {
	return s.Name
}

func (s *UploadMaterialFileForAdminResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *UploadMaterialFileForAdminResponseBodyModule) GetStorageSize() *string {
	return s.StorageSize
}

func (s *UploadMaterialFileForAdminResponseBodyModule) GetSuffix() *string {
	return s.Suffix
}

func (s *UploadMaterialFileForAdminResponseBodyModule) GetType() *string {
	return s.Type
}

func (s *UploadMaterialFileForAdminResponseBodyModule) GetWidth() *int32 {
	return s.Width
}

func (s *UploadMaterialFileForAdminResponseBodyModule) SetBizId(v string) *UploadMaterialFileForAdminResponseBodyModule {
	s.BizId = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBodyModule) SetContentType(v string) *UploadMaterialFileForAdminResponseBodyModule {
	s.ContentType = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBodyModule) SetCreateTime(v string) *UploadMaterialFileForAdminResponseBodyModule {
	s.CreateTime = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBodyModule) SetDeletedTime(v string) *UploadMaterialFileForAdminResponseBodyModule {
	s.DeletedTime = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBodyModule) SetDirectoryId(v string) *UploadMaterialFileForAdminResponseBodyModule {
	s.DirectoryId = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBodyModule) SetFileId(v string) *UploadMaterialFileForAdminResponseBodyModule {
	s.FileId = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBodyModule) SetFileUrl(v string) *UploadMaterialFileForAdminResponseBodyModule {
	s.FileUrl = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBodyModule) SetHeight(v int32) *UploadMaterialFileForAdminResponseBodyModule {
	s.Height = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBodyModule) SetName(v string) *UploadMaterialFileForAdminResponseBodyModule {
	s.Name = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBodyModule) SetStatus(v string) *UploadMaterialFileForAdminResponseBodyModule {
	s.Status = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBodyModule) SetStorageSize(v string) *UploadMaterialFileForAdminResponseBodyModule {
	s.StorageSize = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBodyModule) SetSuffix(v string) *UploadMaterialFileForAdminResponseBodyModule {
	s.Suffix = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBodyModule) SetType(v string) *UploadMaterialFileForAdminResponseBodyModule {
	s.Type = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBodyModule) SetWidth(v int32) *UploadMaterialFileForAdminResponseBodyModule {
	s.Width = &v
	return s
}

func (s *UploadMaterialFileForAdminResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
