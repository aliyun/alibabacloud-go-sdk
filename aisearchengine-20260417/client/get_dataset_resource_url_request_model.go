// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetResourceUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v string) *GetDatasetResourceUrlRequest
	GetDatasetId() *string
	SetPrimaryKey(v string) *GetDatasetResourceUrlRequest
	GetPrimaryKey() *string
}

type GetDatasetResourceUrlRequest struct {
	// The dataset ID. You can view this in the dataset list on the console.
	//
	// example:
	//
	// 730
	DatasetId *string `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
	// The primary key value of the data record in the dataset.
	//
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

func (s *GetDatasetResourceUrlRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *GetDatasetResourceUrlRequest) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *GetDatasetResourceUrlRequest) SetDatasetId(v string) *GetDatasetResourceUrlRequest {
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
