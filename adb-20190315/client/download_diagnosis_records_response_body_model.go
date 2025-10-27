// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadDiagnosisRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadId(v int32) *DownloadDiagnosisRecordsResponseBody
	GetDownloadId() *int32
	SetRequestId(v string) *DownloadDiagnosisRecordsResponseBody
	GetRequestId() *string
}

type DownloadDiagnosisRecordsResponseBody struct {
	// The ID of the download task.
	//
	// example:
	//
	// 68
	DownloadId *int32 `json:"DownloadId,omitempty" xml:"DownloadId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D4ACF4E0-2952-3A87-9A2C-474058******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DownloadDiagnosisRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadDiagnosisRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadDiagnosisRecordsResponseBody) GetDownloadId() *int32 {
	return s.DownloadId
}

func (s *DownloadDiagnosisRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadDiagnosisRecordsResponseBody) SetDownloadId(v int32) *DownloadDiagnosisRecordsResponseBody {
	s.DownloadId = &v
	return s
}

func (s *DownloadDiagnosisRecordsResponseBody) SetRequestId(v string) *DownloadDiagnosisRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadDiagnosisRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}
