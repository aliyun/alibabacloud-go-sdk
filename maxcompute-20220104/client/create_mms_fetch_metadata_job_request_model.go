// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMmsFetchMetadataJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *CreateMmsFetchMetadataJobRequest
	GetDbName() *string
	SetTableNames(v []*string) *CreateMmsFetchMetadataJobRequest
	GetTableNames() []*string
}

type CreateMmsFetchMetadataJobRequest struct {
	// example:
	//
	// default
	DbName     *string   `json:"dbName,omitempty" xml:"dbName,omitempty"`
	TableNames []*string `json:"tableNames,omitempty" xml:"tableNames,omitempty" type:"Repeated"`
}

func (s CreateMmsFetchMetadataJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMmsFetchMetadataJobRequest) GoString() string {
	return s.String()
}

func (s *CreateMmsFetchMetadataJobRequest) GetDbName() *string {
	return s.DbName
}

func (s *CreateMmsFetchMetadataJobRequest) GetTableNames() []*string {
	return s.TableNames
}

func (s *CreateMmsFetchMetadataJobRequest) SetDbName(v string) *CreateMmsFetchMetadataJobRequest {
	s.DbName = &v
	return s
}

func (s *CreateMmsFetchMetadataJobRequest) SetTableNames(v []*string) *CreateMmsFetchMetadataJobRequest {
	s.TableNames = v
	return s
}

func (s *CreateMmsFetchMetadataJobRequest) Validate() error {
	return dara.Validate(s)
}
