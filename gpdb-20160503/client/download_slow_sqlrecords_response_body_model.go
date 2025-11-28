// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadSlowSQLRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DownloadSlowSQLRecordsResponseBody
	GetDBInstanceId() *string
	SetRequestId(v string) *DownloadSlowSQLRecordsResponseBody
	GetRequestId() *string
}

type DownloadSlowSQLRecordsResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DownloadSlowSQLRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadSlowSQLRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadSlowSQLRecordsResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DownloadSlowSQLRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadSlowSQLRecordsResponseBody) SetDBInstanceId(v string) *DownloadSlowSQLRecordsResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DownloadSlowSQLRecordsResponseBody) SetRequestId(v string) *DownloadSlowSQLRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadSlowSQLRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}
