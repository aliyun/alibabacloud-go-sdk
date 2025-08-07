// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSymbolicFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v map[string]interface{}) *GetSymbolicFilesResponseBody
	GetArgs() map[string]interface{}
	SetErrorCode(v int32) *GetSymbolicFilesResponseBody
	GetErrorCode() *int32
	SetMessage(v string) *GetSymbolicFilesResponseBody
	GetMessage() *string
	SetModel(v *GetSymbolicFilesResponseBodyModel) *GetSymbolicFilesResponseBody
	GetModel() *GetSymbolicFilesResponseBodyModel
	SetRequestId(v string) *GetSymbolicFilesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSymbolicFilesResponseBody
	GetSuccess() *bool
}

type GetSymbolicFilesResponseBody struct {
	// Args
	Args map[string]interface{} `json:"Args,omitempty" xml:"Args,omitempty"`
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// successful
	Message *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *GetSymbolicFilesResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// B3AD0FE4-36EF-1641-90B1-77618166F2ff
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSymbolicFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSymbolicFilesResponseBody) GoString() string {
	return s.String()
}

func (s *GetSymbolicFilesResponseBody) GetArgs() map[string]interface{} {
	return s.Args
}

func (s *GetSymbolicFilesResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetSymbolicFilesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSymbolicFilesResponseBody) GetModel() *GetSymbolicFilesResponseBodyModel {
	return s.Model
}

func (s *GetSymbolicFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSymbolicFilesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSymbolicFilesResponseBody) SetArgs(v map[string]interface{}) *GetSymbolicFilesResponseBody {
	s.Args = v
	return s
}

func (s *GetSymbolicFilesResponseBody) SetErrorCode(v int32) *GetSymbolicFilesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSymbolicFilesResponseBody) SetMessage(v string) *GetSymbolicFilesResponseBody {
	s.Message = &v
	return s
}

func (s *GetSymbolicFilesResponseBody) SetModel(v *GetSymbolicFilesResponseBodyModel) *GetSymbolicFilesResponseBody {
	s.Model = v
	return s
}

func (s *GetSymbolicFilesResponseBody) SetRequestId(v string) *GetSymbolicFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSymbolicFilesResponseBody) SetSuccess(v bool) *GetSymbolicFilesResponseBody {
	s.Success = &v
	return s
}

func (s *GetSymbolicFilesResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSymbolicFilesResponseBodyModel struct {
	Items []*GetSymbolicFilesResponseBodyModelItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 3
	Pages *int32 `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSymbolicFilesResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s GetSymbolicFilesResponseBodyModel) GoString() string {
	return s.String()
}

func (s *GetSymbolicFilesResponseBodyModel) GetItems() []*GetSymbolicFilesResponseBodyModelItems {
	return s.Items
}

func (s *GetSymbolicFilesResponseBodyModel) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetSymbolicFilesResponseBodyModel) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSymbolicFilesResponseBodyModel) GetPages() *int32 {
	return s.Pages
}

func (s *GetSymbolicFilesResponseBodyModel) GetTotal() *int64 {
	return s.Total
}

func (s *GetSymbolicFilesResponseBodyModel) SetItems(v []*GetSymbolicFilesResponseBodyModelItems) *GetSymbolicFilesResponseBodyModel {
	s.Items = v
	return s
}

func (s *GetSymbolicFilesResponseBodyModel) SetPageNum(v int32) *GetSymbolicFilesResponseBodyModel {
	s.PageNum = &v
	return s
}

func (s *GetSymbolicFilesResponseBodyModel) SetPageSize(v int32) *GetSymbolicFilesResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *GetSymbolicFilesResponseBodyModel) SetPages(v int32) *GetSymbolicFilesResponseBodyModel {
	s.Pages = &v
	return s
}

func (s *GetSymbolicFilesResponseBodyModel) SetTotal(v int64) *GetSymbolicFilesResponseBodyModel {
	s.Total = &v
	return s
}

func (s *GetSymbolicFilesResponseBodyModel) Validate() error {
	return dara.Validate(s)
}

type GetSymbolicFilesResponseBodyModelItems struct {
	// example:
	//
	// 1.0.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// e8a1a2b9ab653780b34383a942ac91b2
	BuildId *string `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	// example:
	//
	// EXPORT_SUCCESS
	ExportStatus *string `json:"ExportStatus,omitempty" xml:"ExportStatus,omitempty"`
	// example:
	//
	// app_so.zip
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 24781204@android/1743506690915-app_so.zip
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// example:
	//
	// APP_SO
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// example:
	//
	// 1655962713000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 392522
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// uuid
	//
	// example:
	//
	// 9634758587856312DEV
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetSymbolicFilesResponseBodyModelItems) String() string {
	return dara.Prettify(s)
}

func (s GetSymbolicFilesResponseBodyModelItems) GoString() string {
	return s.String()
}

func (s *GetSymbolicFilesResponseBodyModelItems) GetAppVersion() *string {
	return s.AppVersion
}

func (s *GetSymbolicFilesResponseBodyModelItems) GetBuildId() *string {
	return s.BuildId
}

func (s *GetSymbolicFilesResponseBodyModelItems) GetExportStatus() *string {
	return s.ExportStatus
}

func (s *GetSymbolicFilesResponseBodyModelItems) GetFileName() *string {
	return s.FileName
}

func (s *GetSymbolicFilesResponseBodyModelItems) GetFilePath() *string {
	return s.FilePath
}

func (s *GetSymbolicFilesResponseBodyModelItems) GetFileType() *string {
	return s.FileType
}

func (s *GetSymbolicFilesResponseBodyModelItems) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetSymbolicFilesResponseBodyModelItems) GetId() *int64 {
	return s.Id
}

func (s *GetSymbolicFilesResponseBodyModelItems) GetStatus() *string {
	return s.Status
}

func (s *GetSymbolicFilesResponseBodyModelItems) GetUuid() *string {
	return s.Uuid
}

func (s *GetSymbolicFilesResponseBodyModelItems) SetAppVersion(v string) *GetSymbolicFilesResponseBodyModelItems {
	s.AppVersion = &v
	return s
}

func (s *GetSymbolicFilesResponseBodyModelItems) SetBuildId(v string) *GetSymbolicFilesResponseBodyModelItems {
	s.BuildId = &v
	return s
}

func (s *GetSymbolicFilesResponseBodyModelItems) SetExportStatus(v string) *GetSymbolicFilesResponseBodyModelItems {
	s.ExportStatus = &v
	return s
}

func (s *GetSymbolicFilesResponseBodyModelItems) SetFileName(v string) *GetSymbolicFilesResponseBodyModelItems {
	s.FileName = &v
	return s
}

func (s *GetSymbolicFilesResponseBodyModelItems) SetFilePath(v string) *GetSymbolicFilesResponseBodyModelItems {
	s.FilePath = &v
	return s
}

func (s *GetSymbolicFilesResponseBodyModelItems) SetFileType(v string) *GetSymbolicFilesResponseBodyModelItems {
	s.FileType = &v
	return s
}

func (s *GetSymbolicFilesResponseBodyModelItems) SetGmtCreate(v int64) *GetSymbolicFilesResponseBodyModelItems {
	s.GmtCreate = &v
	return s
}

func (s *GetSymbolicFilesResponseBodyModelItems) SetId(v int64) *GetSymbolicFilesResponseBodyModelItems {
	s.Id = &v
	return s
}

func (s *GetSymbolicFilesResponseBodyModelItems) SetStatus(v string) *GetSymbolicFilesResponseBodyModelItems {
	s.Status = &v
	return s
}

func (s *GetSymbolicFilesResponseBodyModelItems) SetUuid(v string) *GetSymbolicFilesResponseBodyModelItems {
	s.Uuid = &v
	return s
}

func (s *GetSymbolicFilesResponseBodyModelItems) Validate() error {
	return dara.Validate(s)
}
