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
	// The returned data.
	Data *CreateMOUsageDetailExportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result.
	//
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
	// The completion time.
	//
	// example:
	//
	// 2026-03-15T21:18:45Z
	CompletedAt *string `json:"CompletedAt,omitempty" xml:"CompletedAt,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-03-15T21:14:45Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The download URL.
	//
	// example:
	//
	// http://***
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The error message.
	//
	// example:
	//
	// picture owner error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The expiration time in ISO 8601 format.
	//
	// example:
	//
	// 2026-02-07T21:14:45Z
	ExpiresAt *string `json:"ExpiresAt,omitempty" xml:"ExpiresAt,omitempty"`
	// The file size in bytes.
	//
	// example:
	//
	// 403
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The query filter conditions.
	Filters *CreateMOUsageDetailExportResponseBodyDataFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Struct"`
	// The target conversion format.
	//
	// example:
	//
	// csv
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The OssKey passed when calling the operation. (This field is deprecated)
	//
	// example:
	//
	// ddd6*****2a76
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	// The total number of data rows.
	//
	// example:
	//
	// 100
	RowCount *int64 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	// The module status.
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
	// The API key to use.
	//
	// example:
	//
	// sk-***
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The inspection end time in the YYYY-MM-DDTHH:mm:ssZ format.
	//
	// example:
	//
	// 2025-03-10T02:02:20Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance name.
	//
	// example:
	//
	// rds_copilot**_public_cn-******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The model to call.
	//
	// example:
	//
	// qwen-flash
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The inspection start time in the YYYY-MM-DDTHH:mm:ssZ format.
	//
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
