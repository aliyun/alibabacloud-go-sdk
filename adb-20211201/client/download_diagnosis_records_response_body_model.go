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
	// The download ID.
	//
	// example:
	//
	// 25494
	DownloadId *int32 `json:"DownloadId,omitempty" xml:"DownloadId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 845774AC-5D43-53A2-AAB8-C73828E68508
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
