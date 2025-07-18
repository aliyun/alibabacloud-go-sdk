// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadDiagnosisRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DownloadDiagnosisRecordsResponseBody
	GetDBInstanceId() *string
	SetDownloadId(v string) *DownloadDiagnosisRecordsResponseBody
	GetDownloadId() *string
	SetRequestId(v string) *DownloadDiagnosisRecordsResponseBody
	GetRequestId() *string
}

type DownloadDiagnosisRecordsResponseBody struct {
	// The ID of the instance.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the download task.
	//
	// example:
	//
	// 11
	DownloadId *string `json:"DownloadId,omitempty" xml:"DownloadId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DownloadDiagnosisRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadDiagnosisRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadDiagnosisRecordsResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DownloadDiagnosisRecordsResponseBody) GetDownloadId() *string {
	return s.DownloadId
}

func (s *DownloadDiagnosisRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadDiagnosisRecordsResponseBody) SetDBInstanceId(v string) *DownloadDiagnosisRecordsResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DownloadDiagnosisRecordsResponseBody) SetDownloadId(v string) *DownloadDiagnosisRecordsResponseBody {
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
