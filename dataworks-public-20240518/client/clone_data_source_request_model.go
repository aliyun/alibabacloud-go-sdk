// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloneDataSourceName(v string) *CloneDataSourceRequest
	GetCloneDataSourceName() *string
	SetId(v int64) *CloneDataSourceRequest
	GetId() *int64
}

type CloneDataSourceRequest struct {
	// The name of the destination data source The name can contain letters, digits, and underscores (_), and must start with a letter. It cannot exceed 60 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo_holo_datasource
	CloneDataSourceName *string `json:"CloneDataSourceName,omitempty" xml:"CloneDataSourceName,omitempty"`
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 16036
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CloneDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CloneDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CloneDataSourceRequest) GetCloneDataSourceName() *string {
	return s.CloneDataSourceName
}

func (s *CloneDataSourceRequest) GetId() *int64 {
	return s.Id
}

func (s *CloneDataSourceRequest) SetCloneDataSourceName(v string) *CloneDataSourceRequest {
	s.CloneDataSourceName = &v
	return s
}

func (s *CloneDataSourceRequest) SetId(v int64) *CloneDataSourceRequest {
	s.Id = &v
	return s
}

func (s *CloneDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
