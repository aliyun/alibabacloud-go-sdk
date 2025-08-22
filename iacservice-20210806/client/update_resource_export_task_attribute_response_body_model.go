// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceExportTaskAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExportTaskId(v string) *UpdateResourceExportTaskAttributeResponseBody
	GetExportTaskId() *string
	SetExportVersion(v string) *UpdateResourceExportTaskAttributeResponseBody
	GetExportVersion() *string
	SetRequestId(v string) *UpdateResourceExportTaskAttributeResponseBody
	GetRequestId() *string
}

type UpdateResourceExportTaskAttributeResponseBody struct {
	// example:
	//
	// ex-kw161ol8te1n701e1igt8q8
	ExportTaskId *string `json:"exportTaskId,omitempty" xml:"exportTaskId,omitempty"`
	// example:
	//
	// v1
	ExportVersion *string `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
	// example:
	//
	// B43F08A7-F2A3-54D3-BDA4-69C9F32A7B9F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateResourceExportTaskAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceExportTaskAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceExportTaskAttributeResponseBody) GetExportTaskId() *string {
	return s.ExportTaskId
}

func (s *UpdateResourceExportTaskAttributeResponseBody) GetExportVersion() *string {
	return s.ExportVersion
}

func (s *UpdateResourceExportTaskAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResourceExportTaskAttributeResponseBody) SetExportTaskId(v string) *UpdateResourceExportTaskAttributeResponseBody {
	s.ExportTaskId = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeResponseBody) SetExportVersion(v string) *UpdateResourceExportTaskAttributeResponseBody {
	s.ExportVersion = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeResponseBody) SetRequestId(v string) *UpdateResourceExportTaskAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
