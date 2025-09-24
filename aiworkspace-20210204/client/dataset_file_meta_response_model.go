// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatasetFileMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetFileMetaId(v string) *DatasetFileMetaResponse
	GetDatasetFileMetaId() *string
	SetResult(v string) *DatasetFileMetaResponse
	GetResult() *string
	SetUri(v string) *DatasetFileMetaResponse
	GetUri() *string
}

type DatasetFileMetaResponse struct {
	// This parameter is required.
	DatasetFileMetaId *string `json:"DatasetFileMetaId,omitempty" xml:"DatasetFileMetaId,omitempty"`
	// This parameter is required.
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Uri    *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DatasetFileMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s DatasetFileMetaResponse) GoString() string {
	return s.String()
}

func (s *DatasetFileMetaResponse) GetDatasetFileMetaId() *string {
	return s.DatasetFileMetaId
}

func (s *DatasetFileMetaResponse) GetResult() *string {
	return s.Result
}

func (s *DatasetFileMetaResponse) GetUri() *string {
	return s.Uri
}

func (s *DatasetFileMetaResponse) SetDatasetFileMetaId(v string) *DatasetFileMetaResponse {
	s.DatasetFileMetaId = &v
	return s
}

func (s *DatasetFileMetaResponse) SetResult(v string) *DatasetFileMetaResponse {
	s.Result = &v
	return s
}

func (s *DatasetFileMetaResponse) SetUri(v string) *DatasetFileMetaResponse {
	s.Uri = &v
	return s
}

func (s *DatasetFileMetaResponse) Validate() error {
	return dara.Validate(s)
}
