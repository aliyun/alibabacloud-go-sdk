// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelResourceExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExportTaskId(v string) *CancelResourceExportTaskResponseBody
	GetExportTaskId() *string
	SetExportVersion(v string) *CancelResourceExportTaskResponseBody
	GetExportVersion() *string
	SetRequestId(v string) *CancelResourceExportTaskResponseBody
	GetRequestId() *string
}

type CancelResourceExportTaskResponseBody struct {
	// example:
	//
	// ex-3b6cb9fa4751a6e5cdc6460282
	ExportTaskId *string `json:"exportTaskId,omitempty" xml:"exportTaskId,omitempty"`
	// example:
	//
	// v1
	ExportVersion *string `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
	// example:
	//
	// 136B3926-DD90-5DB2-96EC-8BAD6407D1C9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CancelResourceExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelResourceExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelResourceExportTaskResponseBody) GetExportTaskId() *string {
	return s.ExportTaskId
}

func (s *CancelResourceExportTaskResponseBody) GetExportVersion() *string {
	return s.ExportVersion
}

func (s *CancelResourceExportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelResourceExportTaskResponseBody) SetExportTaskId(v string) *CancelResourceExportTaskResponseBody {
	s.ExportTaskId = &v
	return s
}

func (s *CancelResourceExportTaskResponseBody) SetExportVersion(v string) *CancelResourceExportTaskResponseBody {
	s.ExportVersion = &v
	return s
}

func (s *CancelResourceExportTaskResponseBody) SetRequestId(v string) *CancelResourceExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelResourceExportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
