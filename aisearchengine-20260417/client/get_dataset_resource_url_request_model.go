// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetResourceUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v int64) *GetDatasetResourceUrlRequest
	GetDatasetId() *int64
	SetPrimaryKey(v string) *GetDatasetResourceUrlRequest
	GetPrimaryKey() *string
}

type GetDatasetResourceUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 730
	DatasetId *int64 `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 01KQCJBPM9JVDTXWV50G2AKXXX
	PrimaryKey *string `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
}

func (s GetDatasetResourceUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResourceUrlRequest) GoString() string {
	return s.String()
}

func (s *GetDatasetResourceUrlRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *GetDatasetResourceUrlRequest) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *GetDatasetResourceUrlRequest) SetDatasetId(v int64) *GetDatasetResourceUrlRequest {
	s.DatasetId = &v
	return s
}

func (s *GetDatasetResourceUrlRequest) SetPrimaryKey(v string) *GetDatasetResourceUrlRequest {
	s.PrimaryKey = &v
	return s
}

func (s *GetDatasetResourceUrlRequest) Validate() error {
	return dara.Validate(s)
}
