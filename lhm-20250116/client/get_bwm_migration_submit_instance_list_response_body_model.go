// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBwmMigrationSubmitInstanceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetBwmMigrationSubmitInstanceListResponseBodyData) *GetBwmMigrationSubmitInstanceListResponseBody
	GetData() []*GetBwmMigrationSubmitInstanceListResponseBodyData
	SetEmpty(v bool) *GetBwmMigrationSubmitInstanceListResponseBody
	GetEmpty() *bool
	SetErrCode(v string) *GetBwmMigrationSubmitInstanceListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetBwmMigrationSubmitInstanceListResponseBody
	GetErrMessage() *string
	SetNotEmpty(v bool) *GetBwmMigrationSubmitInstanceListResponseBody
	GetNotEmpty() *bool
	SetPageIndex(v int32) *GetBwmMigrationSubmitInstanceListResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *GetBwmMigrationSubmitInstanceListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetBwmMigrationSubmitInstanceListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBwmMigrationSubmitInstanceListResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *GetBwmMigrationSubmitInstanceListResponseBody
	GetTotalCount() *int32
	SetTotalPages(v int32) *GetBwmMigrationSubmitInstanceListResponseBody
	GetTotalPages() *int32
}

type GetBwmMigrationSubmitInstanceListResponseBody struct {
	// The response data.
	Data []*GetBwmMigrationSubmitInstanceListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Indicates whether the result is empty.
	Empty *bool `json:"empty,omitempty" xml:"empty,omitempty"`
	// The error code. An empty string is returned if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. An empty string is returned if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// Indicates whether the result is not empty.
	NotEmpty *bool `json:"notEmpty,omitempty" xml:"notEmpty,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The page size.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure. If the call fails, use errCode and errMessage to troubleshoot the issue.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 10
	TotalPages *int32 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
}

func (s GetBwmMigrationSubmitInstanceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBwmMigrationSubmitInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) GetData() []*GetBwmMigrationSubmitInstanceListResponseBodyData {
	return s.Data
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) GetEmpty() *bool {
	return s.Empty
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) GetNotEmpty() *bool {
	return s.NotEmpty
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) SetData(v []*GetBwmMigrationSubmitInstanceListResponseBodyData) *GetBwmMigrationSubmitInstanceListResponseBody {
	s.Data = v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) SetEmpty(v bool) *GetBwmMigrationSubmitInstanceListResponseBody {
	s.Empty = &v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) SetErrCode(v string) *GetBwmMigrationSubmitInstanceListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) SetErrMessage(v string) *GetBwmMigrationSubmitInstanceListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) SetNotEmpty(v bool) *GetBwmMigrationSubmitInstanceListResponseBody {
	s.NotEmpty = &v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) SetPageIndex(v int32) *GetBwmMigrationSubmitInstanceListResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) SetPageSize(v int32) *GetBwmMigrationSubmitInstanceListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) SetRequestId(v string) *GetBwmMigrationSubmitInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) SetSuccess(v bool) *GetBwmMigrationSubmitInstanceListResponseBody {
	s.Success = &v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) SetTotalCount(v int32) *GetBwmMigrationSubmitInstanceListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) SetTotalPages(v int32) *GetBwmMigrationSubmitInstanceListResponseBody {
	s.TotalPages = &v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBwmMigrationSubmitInstanceListResponseBodyData struct {
	// The reason for the conversion failure.
	//
	// example:
	//
	// None
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	// The conversion execution time.
	//
	// example:
	//
	// 2023-01-01 00:00:00
	GmtConvert *string `json:"gmtConvert,omitempty" xml:"gmtConvert,omitempty"`
	// The UUID of the instance.
	//
	// example:
	//
	// 12345
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// instance-1
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The time when the source metadata was last updated.
	//
	// example:
	//
	// 2023-01-01 00:00:00
	SrcMetaGmtUpdate *string `json:"srcMetaGmtUpdate,omitempty" xml:"srcMetaGmtUpdate,omitempty"`
	// The scheduling information of the source.
	//
	// example:
	//
	// source-info
	SrcMetaInfo *string `json:"srcMetaInfo,omitempty" xml:"srcMetaInfo,omitempty"`
	// The execution status of the instance. Valid values:
	//
	// - NOT_START: Not started.
	//
	// - READY: Pending execution.
	//
	// - RUNNING: Running.
	//
	// - ALL_SUCCESS: All succeeded.
	//
	// - PARTIAL_SUCCESS: Partially succeeded.
	//
	// - FAILURE: Failed.
	//
	// - MANUAL: Manually uploaded.
	//
	// If the status code cannot be recognized, the value defaults to NOT_START.
	//
	// example:
	//
	// NOT_START
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetBwmMigrationSubmitInstanceListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBwmMigrationSubmitInstanceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBwmMigrationSubmitInstanceListResponseBodyData) GetDetail() *string {
	return s.Detail
}

func (s *GetBwmMigrationSubmitInstanceListResponseBodyData) GetGmtConvert() *string {
	return s.GmtConvert
}

func (s *GetBwmMigrationSubmitInstanceListResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetBwmMigrationSubmitInstanceListResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetBwmMigrationSubmitInstanceListResponseBodyData) GetSrcMetaGmtUpdate() *string {
	return s.SrcMetaGmtUpdate
}

func (s *GetBwmMigrationSubmitInstanceListResponseBodyData) GetSrcMetaInfo() *string {
	return s.SrcMetaInfo
}

func (s *GetBwmMigrationSubmitInstanceListResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetBwmMigrationSubmitInstanceListResponseBodyData) SetDetail(v string) *GetBwmMigrationSubmitInstanceListResponseBodyData {
	s.Detail = &v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListResponseBodyData) SetGmtConvert(v string) *GetBwmMigrationSubmitInstanceListResponseBodyData {
	s.GmtConvert = &v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListResponseBodyData) SetInstanceId(v string) *GetBwmMigrationSubmitInstanceListResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListResponseBodyData) SetInstanceName(v string) *GetBwmMigrationSubmitInstanceListResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListResponseBodyData) SetSrcMetaGmtUpdate(v string) *GetBwmMigrationSubmitInstanceListResponseBodyData {
	s.SrcMetaGmtUpdate = &v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListResponseBodyData) SetSrcMetaInfo(v string) *GetBwmMigrationSubmitInstanceListResponseBodyData {
	s.SrcMetaInfo = &v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListResponseBodyData) SetStatus(v string) *GetBwmMigrationSubmitInstanceListResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
