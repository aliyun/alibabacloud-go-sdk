// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMOUsageDetailExportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateMOUsageDetailExportResponseBodyData) *CreateMOUsageDetailExportResponseBody
	GetData() *CreateMOUsageDetailExportResponseBodyData
	SetMessage(v string) *CreateMOUsageDetailExportResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMOUsageDetailExportResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMOUsageDetailExportResponseBody
	GetSuccess() *bool
}

type CreateMOUsageDetailExportResponseBody struct {
	Data *CreateMOUsageDetailExportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMOUsageDetailExportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMOUsageDetailExportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMOUsageDetailExportResponseBody) GetData() *CreateMOUsageDetailExportResponseBodyData {
	return s.Data
}

func (s *CreateMOUsageDetailExportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMOUsageDetailExportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMOUsageDetailExportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMOUsageDetailExportResponseBody) SetData(v *CreateMOUsageDetailExportResponseBodyData) *CreateMOUsageDetailExportResponseBody {
	s.Data = v
	return s
}

func (s *CreateMOUsageDetailExportResponseBody) SetMessage(v string) *CreateMOUsageDetailExportResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMOUsageDetailExportResponseBody) SetRequestId(v string) *CreateMOUsageDetailExportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMOUsageDetailExportResponseBody) SetSuccess(v bool) *CreateMOUsageDetailExportResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMOUsageDetailExportResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMOUsageDetailExportResponseBodyData struct {
	// example:
	//
	// 2026-03-15T21:18:45Z
	CompletedAt *string `json:"CompletedAt,omitempty" xml:"CompletedAt,omitempty"`
	// example:
	//
	// 2026-03-15T21:14:45Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// http://***
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// example:
	//
	// picture owner error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 下载 URL 失效时间（UTC ISO8601）
	//
	// example:
	//
	// 2026-02-07T21:14:45Z
	ExpiresAt *string `json:"ExpiresAt,omitempty" xml:"ExpiresAt,omitempty"`
	// example:
	//
	// 403
	FileSize *int64                                            `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	Filters  *CreateMOUsageDetailExportResponseBodyDataFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Struct"`
	// example:
	//
	// csv
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// example:
	//
	// ddd6*****2a76
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	// example:
	//
	// 100
	RowCount *int64 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	// pending / processing / completed / failed / expired
	//
	// example:
	//
	// pending / processing / completed / failed / expired
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateMOUsageDetailExportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMOUsageDetailExportResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMOUsageDetailExportResponseBodyData) GetCompletedAt() *string {
	return s.CompletedAt
}

func (s *CreateMOUsageDetailExportResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateMOUsageDetailExportResponseBodyData) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *CreateMOUsageDetailExportResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateMOUsageDetailExportResponseBodyData) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *CreateMOUsageDetailExportResponseBodyData) GetFileSize() *int64 {
	return s.FileSize
}

func (s *CreateMOUsageDetailExportResponseBodyData) GetFilters() *CreateMOUsageDetailExportResponseBodyDataFilters {
	return s.Filters
}

func (s *CreateMOUsageDetailExportResponseBodyData) GetFormat() *string {
	return s.Format
}

func (s *CreateMOUsageDetailExportResponseBodyData) GetOssKey() *string {
	return s.OssKey
}

func (s *CreateMOUsageDetailExportResponseBodyData) GetRowCount() *int64 {
	return s.RowCount
}

func (s *CreateMOUsageDetailExportResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateMOUsageDetailExportResponseBodyData) SetCompletedAt(v string) *CreateMOUsageDetailExportResponseBodyData {
	s.CompletedAt = &v
	return s
}

func (s *CreateMOUsageDetailExportResponseBodyData) SetCreatedAt(v string) *CreateMOUsageDetailExportResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateMOUsageDetailExportResponseBodyData) SetDownloadUrl(v string) *CreateMOUsageDetailExportResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *CreateMOUsageDetailExportResponseBodyData) SetErrorMessage(v string) *CreateMOUsageDetailExportResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *CreateMOUsageDetailExportResponseBodyData) SetExpiresAt(v string) *CreateMOUsageDetailExportResponseBodyData {
	s.ExpiresAt = &v
	return s
}

func (s *CreateMOUsageDetailExportResponseBodyData) SetFileSize(v int64) *CreateMOUsageDetailExportResponseBodyData {
	s.FileSize = &v
	return s
}

func (s *CreateMOUsageDetailExportResponseBodyData) SetFilters(v *CreateMOUsageDetailExportResponseBodyDataFilters) *CreateMOUsageDetailExportResponseBodyData {
	s.Filters = v
	return s
}

func (s *CreateMOUsageDetailExportResponseBodyData) SetFormat(v string) *CreateMOUsageDetailExportResponseBodyData {
	s.Format = &v
	return s
}

func (s *CreateMOUsageDetailExportResponseBodyData) SetOssKey(v string) *CreateMOUsageDetailExportResponseBodyData {
	s.OssKey = &v
	return s
}

func (s *CreateMOUsageDetailExportResponseBodyData) SetRowCount(v int64) *CreateMOUsageDetailExportResponseBodyData {
	s.RowCount = &v
	return s
}

func (s *CreateMOUsageDetailExportResponseBodyData) SetStatus(v string) *CreateMOUsageDetailExportResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateMOUsageDetailExportResponseBodyData) Validate() error {
	if s.Filters != nil {
		if err := s.Filters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMOUsageDetailExportResponseBodyDataFilters struct {
	// example:
	//
	// sk-***
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// 2025-03-10T02:02:20Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// rds_copilot**_public_cn-******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// qwen-flash
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 2026-03-05T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateMOUsageDetailExportResponseBodyDataFilters) String() string {
	return dara.Prettify(s)
}

func (s CreateMOUsageDetailExportResponseBodyDataFilters) GoString() string {
	return s.String()
}

func (s *CreateMOUsageDetailExportResponseBodyDataFilters) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateMOUsageDetailExportResponseBodyDataFilters) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateMOUsageDetailExportResponseBodyDataFilters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateMOUsageDetailExportResponseBodyDataFilters) GetModel() *string {
	return s.Model
}

func (s *CreateMOUsageDetailExportResponseBodyDataFilters) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateMOUsageDetailExportResponseBodyDataFilters) SetApiKey(v string) *CreateMOUsageDetailExportResponseBodyDataFilters {
	s.ApiKey = &v
	return s
}

func (s *CreateMOUsageDetailExportResponseBodyDataFilters) SetEndTime(v string) *CreateMOUsageDetailExportResponseBodyDataFilters {
	s.EndTime = &v
	return s
}

func (s *CreateMOUsageDetailExportResponseBodyDataFilters) SetInstanceId(v string) *CreateMOUsageDetailExportResponseBodyDataFilters {
	s.InstanceId = &v
	return s
}

func (s *CreateMOUsageDetailExportResponseBodyDataFilters) SetModel(v string) *CreateMOUsageDetailExportResponseBodyDataFilters {
	s.Model = &v
	return s
}

func (s *CreateMOUsageDetailExportResponseBodyDataFilters) SetStartTime(v string) *CreateMOUsageDetailExportResponseBodyDataFilters {
	s.StartTime = &v
	return s
}

func (s *CreateMOUsageDetailExportResponseBodyDataFilters) Validate() error {
	return dara.Validate(s)
}
