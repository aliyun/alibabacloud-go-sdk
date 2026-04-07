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
	// The metadata ID of the dataset file.
	//
	// This parameter is required.
	//
	// example:
	//
	// 07914c9534586e4e7aa6e9dbca5009082df*******8a0d857b33296c59bf6
	DatasetFileMetaId *string `json:"DatasetFileMetaId,omitempty" xml:"DatasetFileMetaId,omitempty"`
	// The description of the reason why the metadata operation failed.
	//
	// This parameter is required.
	//
	// example:
	//
	// Not Found
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The file URI.
	//
	// example:
	//
	// oss://l*****-test/dataset/1653421.jpg
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
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
