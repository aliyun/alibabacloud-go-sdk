// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExportTaskId(v string) *CreateResourceExportTaskResponseBody
	GetExportTaskId() *string
	SetExportVersion(v string) *CreateResourceExportTaskResponseBody
	GetExportVersion() *string
	SetRequestId(v string) *CreateResourceExportTaskResponseBody
	GetRequestId() *string
}

type CreateResourceExportTaskResponseBody struct {
	// The resource export task ID.
	//
	// example:
	//
	// ex-4a1ec8b7003d24528326821be
	ExportTaskId *string `json:"exportTaskId,omitempty" xml:"exportTaskId,omitempty"`
	// The resource export version.
	//
	// example:
	//
	// v1
	ExportVersion *string `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CFD8C2A8-5BE7-56D2-916D-64039B8E06E3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateResourceExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceExportTaskResponseBody) GetExportTaskId() *string {
	return s.ExportTaskId
}

func (s *CreateResourceExportTaskResponseBody) GetExportVersion() *string {
	return s.ExportVersion
}

func (s *CreateResourceExportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResourceExportTaskResponseBody) SetExportTaskId(v string) *CreateResourceExportTaskResponseBody {
	s.ExportTaskId = &v
	return s
}

func (s *CreateResourceExportTaskResponseBody) SetExportVersion(v string) *CreateResourceExportTaskResponseBody {
	s.ExportVersion = &v
	return s
}

func (s *CreateResourceExportTaskResponseBody) SetRequestId(v string) *CreateResourceExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceExportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
