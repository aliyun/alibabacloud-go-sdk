// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExportConfigJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadUrl(v string) *GetExportConfigJobResponseBody
	GetDownloadUrl() *string
	SetExpireTime(v int64) *GetExportConfigJobResponseBody
	GetExpireTime() *int64
	SetJobId(v string) *GetExportConfigJobResponseBody
	GetJobId() *string
	SetMessage(v string) *GetExportConfigJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetExportConfigJobResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetExportConfigJobResponseBody
	GetStatus() *string
}

type GetExportConfigJobResponseBody struct {
	// example:
	//
	// https://bastionhost-cn-hangzhou-164***.oss-cn-hangzhou.aliyuncs.com/bastionhost-cn-2******
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// example:
	//
	// 1679393152
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 2
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 98DBE5C2-7D7A-5393-9E5A-71074336D33B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetExportConfigJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExportConfigJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetExportConfigJobResponseBody) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetExportConfigJobResponseBody) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GetExportConfigJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetExportConfigJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetExportConfigJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExportConfigJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetExportConfigJobResponseBody) SetDownloadUrl(v string) *GetExportConfigJobResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *GetExportConfigJobResponseBody) SetExpireTime(v int64) *GetExportConfigJobResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GetExportConfigJobResponseBody) SetJobId(v string) *GetExportConfigJobResponseBody {
	s.JobId = &v
	return s
}

func (s *GetExportConfigJobResponseBody) SetMessage(v string) *GetExportConfigJobResponseBody {
	s.Message = &v
	return s
}

func (s *GetExportConfigJobResponseBody) SetRequestId(v string) *GetExportConfigJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExportConfigJobResponseBody) SetStatus(v string) *GetExportConfigJobResponseBody {
	s.Status = &v
	return s
}

func (s *GetExportConfigJobResponseBody) Validate() error {
	return dara.Validate(s)
}
