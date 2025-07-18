// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadSQLLogsRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadId(v int64) *DownloadSQLLogsRecordsResponseBody
	GetDownloadId() *int64
	SetRequestId(v string) *DownloadSQLLogsRecordsResponseBody
	GetRequestId() *string
}

type DownloadSQLLogsRecordsResponseBody struct {
	// The ID of the download task.
	//
	// example:
	//
	// 11
	DownloadId *int64 `json:"DownloadId,omitempty" xml:"DownloadId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7565770E-7C45-462D-BA4A-8A5396F2CAD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DownloadSQLLogsRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadSQLLogsRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadSQLLogsRecordsResponseBody) GetDownloadId() *int64 {
	return s.DownloadId
}

func (s *DownloadSQLLogsRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadSQLLogsRecordsResponseBody) SetDownloadId(v int64) *DownloadSQLLogsRecordsResponseBody {
	s.DownloadId = &v
	return s
}

func (s *DownloadSQLLogsRecordsResponseBody) SetRequestId(v string) *DownloadSQLLogsRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadSQLLogsRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}
