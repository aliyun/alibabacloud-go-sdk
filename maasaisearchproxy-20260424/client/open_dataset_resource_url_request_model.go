// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenDatasetResourceUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v int64) *OpenDatasetResourceUrlRequest
	GetDatasetId() *int64
	SetPrimaryKey(v string) *OpenDatasetResourceUrlRequest
	GetPrimaryKey() *string
}

type OpenDatasetResourceUrlRequest struct {
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

func (s OpenDatasetResourceUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenDatasetResourceUrlRequest) GoString() string {
	return s.String()
}

func (s *OpenDatasetResourceUrlRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *OpenDatasetResourceUrlRequest) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *OpenDatasetResourceUrlRequest) SetDatasetId(v int64) *OpenDatasetResourceUrlRequest {
	s.DatasetId = &v
	return s
}

func (s *OpenDatasetResourceUrlRequest) SetPrimaryKey(v string) *OpenDatasetResourceUrlRequest {
	s.PrimaryKey = &v
	return s
}

func (s *OpenDatasetResourceUrlRequest) Validate() error {
	return dara.Validate(s)
}
