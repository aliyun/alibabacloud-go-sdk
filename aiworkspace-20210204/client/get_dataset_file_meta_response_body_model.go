// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetFileMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetFileMeta(v *DatasetFileMetaContentGet) *GetDatasetFileMetaResponseBody
	GetDatasetFileMeta() *DatasetFileMetaContentGet
	SetDatasetId(v string) *GetDatasetFileMetaResponseBody
	GetDatasetId() *string
	SetDatasetVersion(v string) *GetDatasetFileMetaResponseBody
	GetDatasetVersion() *string
	SetRequestId(v string) *GetDatasetFileMetaResponseBody
	GetRequestId() *string
	SetWorkspaceId(v string) *GetDatasetFileMetaResponseBody
	GetWorkspaceId() *string
}

type GetDatasetFileMetaResponseBody struct {
	// The queried metadata records of dataset files.
	DatasetFileMeta *DatasetFileMetaContentGet `json:"DatasetFileMeta,omitempty" xml:"DatasetFileMeta,omitempty"`
	// The dataset ID.
	//
	// example:
	//
	// d-rbvg5wz****c9ks92
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The dataset version.
	//
	// example:
	//
	// v1
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 1234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDatasetFileMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetFileMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasetFileMetaResponseBody) GetDatasetFileMeta() *DatasetFileMetaContentGet {
	return s.DatasetFileMeta
}

func (s *GetDatasetFileMetaResponseBody) GetDatasetId() *string {
	return s.DatasetId
}

func (s *GetDatasetFileMetaResponseBody) GetDatasetVersion() *string {
	return s.DatasetVersion
}

func (s *GetDatasetFileMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatasetFileMetaResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetDatasetFileMetaResponseBody) SetDatasetFileMeta(v *DatasetFileMetaContentGet) *GetDatasetFileMetaResponseBody {
	s.DatasetFileMeta = v
	return s
}

func (s *GetDatasetFileMetaResponseBody) SetDatasetId(v string) *GetDatasetFileMetaResponseBody {
	s.DatasetId = &v
	return s
}

func (s *GetDatasetFileMetaResponseBody) SetDatasetVersion(v string) *GetDatasetFileMetaResponseBody {
	s.DatasetVersion = &v
	return s
}

func (s *GetDatasetFileMetaResponseBody) SetRequestId(v string) *GetDatasetFileMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatasetFileMetaResponseBody) SetWorkspaceId(v string) *GetDatasetFileMetaResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetDatasetFileMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
