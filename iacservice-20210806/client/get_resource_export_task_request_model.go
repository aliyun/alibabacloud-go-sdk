// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceExportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExportVersion(v string) *GetResourceExportTaskRequest
	GetExportVersion() *string
}

type GetResourceExportTaskRequest struct {
	// example:
	//
	// v3
	ExportVersion *string `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
}

func (s GetResourceExportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceExportTaskRequest) GoString() string {
	return s.String()
}

func (s *GetResourceExportTaskRequest) GetExportVersion() *string {
	return s.ExportVersion
}

func (s *GetResourceExportTaskRequest) SetExportVersion(v string) *GetResourceExportTaskRequest {
	s.ExportVersion = &v
	return s
}

func (s *GetResourceExportTaskRequest) Validate() error {
	return dara.Validate(s)
}
