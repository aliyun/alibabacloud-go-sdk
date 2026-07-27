// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMOUsageDetailExportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeMOUsageDetailExportResponseBodyData) *DescribeMOUsageDetailExportResponseBody
	GetData() *DescribeMOUsageDetailExportResponseBodyData
	SetMessage(v string) *DescribeMOUsageDetailExportResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeMOUsageDetailExportResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMOUsageDetailExportResponseBody
	GetSuccess() *bool
}

type DescribeMOUsageDetailExportResponseBody struct {
	Data *DescribeMOUsageDetailExportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s DescribeMOUsageDetailExportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMOUsageDetailExportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMOUsageDetailExportResponseBody) GetData() *DescribeMOUsageDetailExportResponseBodyData {
	return s.Data
}

func (s *DescribeMOUsageDetailExportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMOUsageDetailExportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMOUsageDetailExportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMOUsageDetailExportResponseBody) SetData(v *DescribeMOUsageDetailExportResponseBodyData) *DescribeMOUsageDetailExportResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMOUsageDetailExportResponseBody) SetMessage(v string) *DescribeMOUsageDetailExportResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMOUsageDetailExportResponseBody) SetRequestId(v string) *DescribeMOUsageDetailExportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMOUsageDetailExportResponseBody) SetSuccess(v bool) *DescribeMOUsageDetailExportResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMOUsageDetailExportResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMOUsageDetailExportResponseBodyData struct {
	// example:
	//
	// 2026-02-01T01:50:03Z
	CompletedAt *string `json:"CompletedAt,omitempty" xml:"CompletedAt,omitempty"`
	// example:
	//
	// 2026-02-01T01:45:03Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// OSS 预签名下载 URL；status=expired 时为 null
	//
	// example:
	//
	// http://***
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// example:
	//
	// tenant auth error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 下载 URL 失效时间（UTC ISO8601）
	//
	// example:
	//
	// 2026-02-04T01:45:03Z
	ExpiresAt *string `json:"ExpiresAt,omitempty" xml:"ExpiresAt,omitempty"`
	// example:
	//
	// 252
	FileSize *int64                                              `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	Filters  *DescribeMOUsageDetailExportResponseBodyDataFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Struct"`
	// example:
	//
	// csv
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// example:
	//
	// 1
	RowCount *int64 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	// none / pending / processing / completed / failed / expired
	//
	// example:
	//
	// pending / processing / completed / failed / expired
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMOUsageDetailExportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMOUsageDetailExportResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMOUsageDetailExportResponseBodyData) GetCompletedAt() *string {
	return s.CompletedAt
}

func (s *DescribeMOUsageDetailExportResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DescribeMOUsageDetailExportResponseBodyData) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DescribeMOUsageDetailExportResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeMOUsageDetailExportResponseBodyData) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *DescribeMOUsageDetailExportResponseBodyData) GetFileSize() *int64 {
	return s.FileSize
}

func (s *DescribeMOUsageDetailExportResponseBodyData) GetFilters() *DescribeMOUsageDetailExportResponseBodyDataFilters {
	return s.Filters
}

func (s *DescribeMOUsageDetailExportResponseBodyData) GetFormat() *string {
	return s.Format
}

func (s *DescribeMOUsageDetailExportResponseBodyData) GetRowCount() *int64 {
	return s.RowCount
}

func (s *DescribeMOUsageDetailExportResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeMOUsageDetailExportResponseBodyData) SetCompletedAt(v string) *DescribeMOUsageDetailExportResponseBodyData {
	s.CompletedAt = &v
	return s
}

func (s *DescribeMOUsageDetailExportResponseBodyData) SetCreatedAt(v string) *DescribeMOUsageDetailExportResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *DescribeMOUsageDetailExportResponseBodyData) SetDownloadUrl(v string) *DescribeMOUsageDetailExportResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeMOUsageDetailExportResponseBodyData) SetErrorMessage(v string) *DescribeMOUsageDetailExportResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMOUsageDetailExportResponseBodyData) SetExpiresAt(v string) *DescribeMOUsageDetailExportResponseBodyData {
	s.ExpiresAt = &v
	return s
}

func (s *DescribeMOUsageDetailExportResponseBodyData) SetFileSize(v int64) *DescribeMOUsageDetailExportResponseBodyData {
	s.FileSize = &v
	return s
}

func (s *DescribeMOUsageDetailExportResponseBodyData) SetFilters(v *DescribeMOUsageDetailExportResponseBodyDataFilters) *DescribeMOUsageDetailExportResponseBodyData {
	s.Filters = v
	return s
}

func (s *DescribeMOUsageDetailExportResponseBodyData) SetFormat(v string) *DescribeMOUsageDetailExportResponseBodyData {
	s.Format = &v
	return s
}

func (s *DescribeMOUsageDetailExportResponseBodyData) SetRowCount(v int64) *DescribeMOUsageDetailExportResponseBodyData {
	s.RowCount = &v
	return s
}

func (s *DescribeMOUsageDetailExportResponseBodyData) SetStatus(v string) *DescribeMOUsageDetailExportResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeMOUsageDetailExportResponseBodyData) Validate() error {
	if s.Filters != nil {
		if err := s.Filters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMOUsageDetailExportResponseBodyDataFilters struct {
	// example:
	//
	// sk-rds-**
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// 2026-01-30T01:45:03Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// rds_copilot**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// glm-5
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 2026-01-28T01:45:03Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeMOUsageDetailExportResponseBodyDataFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeMOUsageDetailExportResponseBodyDataFilters) GoString() string {
	return s.String()
}

func (s *DescribeMOUsageDetailExportResponseBodyDataFilters) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeMOUsageDetailExportResponseBodyDataFilters) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeMOUsageDetailExportResponseBodyDataFilters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMOUsageDetailExportResponseBodyDataFilters) GetModel() *string {
	return s.Model
}

func (s *DescribeMOUsageDetailExportResponseBodyDataFilters) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeMOUsageDetailExportResponseBodyDataFilters) SetApiKey(v string) *DescribeMOUsageDetailExportResponseBodyDataFilters {
	s.ApiKey = &v
	return s
}

func (s *DescribeMOUsageDetailExportResponseBodyDataFilters) SetEndTime(v string) *DescribeMOUsageDetailExportResponseBodyDataFilters {
	s.EndTime = &v
	return s
}

func (s *DescribeMOUsageDetailExportResponseBodyDataFilters) SetInstanceId(v string) *DescribeMOUsageDetailExportResponseBodyDataFilters {
	s.InstanceId = &v
	return s
}

func (s *DescribeMOUsageDetailExportResponseBodyDataFilters) SetModel(v string) *DescribeMOUsageDetailExportResponseBodyDataFilters {
	s.Model = &v
	return s
}

func (s *DescribeMOUsageDetailExportResponseBodyDataFilters) SetStartTime(v string) *DescribeMOUsageDetailExportResponseBodyDataFilters {
	s.StartTime = &v
	return s
}

func (s *DescribeMOUsageDetailExportResponseBodyDataFilters) Validate() error {
	return dara.Validate(s)
}
