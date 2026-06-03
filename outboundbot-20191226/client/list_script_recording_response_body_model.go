// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptRecordingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListScriptRecordingResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListScriptRecordingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListScriptRecordingResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListScriptRecordingResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListScriptRecordingResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListScriptRecordingResponseBody
	GetRequestId() *string
	SetScriptRecordings(v []*ListScriptRecordingResponseBodyScriptRecordings) *ListScriptRecordingResponseBody
	GetScriptRecordings() []*ListScriptRecordingResponseBodyScriptRecordings
	SetSuccess(v bool) *ListScriptRecordingResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListScriptRecordingResponseBody
	GetTotalCount() *int64
}

type ListScriptRecordingResponseBody struct {
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
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId        *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScriptRecordings []*ListScriptRecordingResponseBodyScriptRecordings `json:"ScriptRecordings,omitempty" xml:"ScriptRecordings,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 99
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListScriptRecordingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScriptRecordingResponseBody) GoString() string {
	return s.String()
}

func (s *ListScriptRecordingResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListScriptRecordingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListScriptRecordingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListScriptRecordingResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListScriptRecordingResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScriptRecordingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScriptRecordingResponseBody) GetScriptRecordings() []*ListScriptRecordingResponseBodyScriptRecordings {
	return s.ScriptRecordings
}

func (s *ListScriptRecordingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListScriptRecordingResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListScriptRecordingResponseBody) SetCode(v string) *ListScriptRecordingResponseBody {
	s.Code = &v
	return s
}

func (s *ListScriptRecordingResponseBody) SetHttpStatusCode(v int32) *ListScriptRecordingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListScriptRecordingResponseBody) SetMessage(v string) *ListScriptRecordingResponseBody {
	s.Message = &v
	return s
}

func (s *ListScriptRecordingResponseBody) SetPageNumber(v int32) *ListScriptRecordingResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListScriptRecordingResponseBody) SetPageSize(v int32) *ListScriptRecordingResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListScriptRecordingResponseBody) SetRequestId(v string) *ListScriptRecordingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScriptRecordingResponseBody) SetScriptRecordings(v []*ListScriptRecordingResponseBodyScriptRecordings) *ListScriptRecordingResponseBody {
	s.ScriptRecordings = v
	return s
}

func (s *ListScriptRecordingResponseBody) SetSuccess(v bool) *ListScriptRecordingResponseBody {
	s.Success = &v
	return s
}

func (s *ListScriptRecordingResponseBody) SetTotalCount(v int64) *ListScriptRecordingResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListScriptRecordingResponseBody) Validate() error {
	if s.ScriptRecordings != nil {
		for _, item := range s.ScriptRecordings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScriptRecordingResponseBodyScriptRecordings struct {
	// example:
	//
	// 2022-07-11T07:51:49.000+0000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1654601332000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1654601332000
	GmtUpload *int64 `json:"GmtUpload,omitempty" xml:"GmtUpload,omitempty"`
	// example:
	//
	// 1
	InnerId *string `json:"InnerId,omitempty" xml:"InnerId,omitempty"`
	// example:
	//
	// ff0fb845-9f90-46d3-9716-d36b8a1e753a
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RecordingContent *string `json:"RecordingContent,omitempty" xml:"RecordingContent,omitempty"`
	// example:
	//
	// 10
	RecordingDuration *int32 `json:"RecordingDuration,omitempty" xml:"RecordingDuration,omitempty"`
	// example:
	//
	// hello.wav
	RecordingName *string `json:"RecordingName,omitempty" xml:"RecordingName,omitempty"`
	RefId         *string `json:"RefId,omitempty" xml:"RefId,omitempty"`
	// example:
	//
	// 6019b692-fd9e-4adb-8877-cef6a297b234
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// example:
	//
	// 8
	State       *int32  `json:"State,omitempty" xml:"State,omitempty"`
	StateExtend *string `json:"StateExtend,omitempty" xml:"StateExtend,omitempty"`
	// example:
	//
	// 393674ed-3b5d-db44-0fda-615d05210178
	StorageUuid *string `json:"StorageUuid,omitempty" xml:"StorageUuid,omitempty"`
	// example:
	//
	// 0a77386e-6402-8d23-4adf-6ec13b3f404d
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListScriptRecordingResponseBodyScriptRecordings) String() string {
	return dara.Prettify(s)
}

func (s ListScriptRecordingResponseBodyScriptRecordings) GoString() string {
	return s.String()
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) GetGmtUpload() *int64 {
	return s.GmtUpload
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) GetInnerId() *string {
	return s.InnerId
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) GetRecordingContent() *string {
	return s.RecordingContent
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) GetRecordingDuration() *int32 {
	return s.RecordingDuration
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) GetRecordingName() *string {
	return s.RecordingName
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) GetRefId() *string {
	return s.RefId
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) GetState() *int32 {
	return s.State
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) GetStateExtend() *string {
	return s.StateExtend
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) GetStorageUuid() *string {
	return s.StorageUuid
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) GetUuid() *string {
	return s.Uuid
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) SetGmtCreate(v int64) *ListScriptRecordingResponseBodyScriptRecordings {
	s.GmtCreate = &v
	return s
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) SetGmtModified(v int64) *ListScriptRecordingResponseBodyScriptRecordings {
	s.GmtModified = &v
	return s
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) SetGmtUpload(v int64) *ListScriptRecordingResponseBodyScriptRecordings {
	s.GmtUpload = &v
	return s
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) SetInnerId(v string) *ListScriptRecordingResponseBodyScriptRecordings {
	s.InnerId = &v
	return s
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) SetInstanceId(v string) *ListScriptRecordingResponseBodyScriptRecordings {
	s.InstanceId = &v
	return s
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) SetRecordingContent(v string) *ListScriptRecordingResponseBodyScriptRecordings {
	s.RecordingContent = &v
	return s
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) SetRecordingDuration(v int32) *ListScriptRecordingResponseBodyScriptRecordings {
	s.RecordingDuration = &v
	return s
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) SetRecordingName(v string) *ListScriptRecordingResponseBodyScriptRecordings {
	s.RecordingName = &v
	return s
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) SetRefId(v string) *ListScriptRecordingResponseBodyScriptRecordings {
	s.RefId = &v
	return s
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) SetScriptId(v string) *ListScriptRecordingResponseBodyScriptRecordings {
	s.ScriptId = &v
	return s
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) SetState(v int32) *ListScriptRecordingResponseBodyScriptRecordings {
	s.State = &v
	return s
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) SetStateExtend(v string) *ListScriptRecordingResponseBodyScriptRecordings {
	s.StateExtend = &v
	return s
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) SetStorageUuid(v string) *ListScriptRecordingResponseBodyScriptRecordings {
	s.StorageUuid = &v
	return s
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) SetUuid(v string) *ListScriptRecordingResponseBodyScriptRecordings {
	s.Uuid = &v
	return s
}

func (s *ListScriptRecordingResponseBodyScriptRecordings) Validate() error {
	return dara.Validate(s)
}
