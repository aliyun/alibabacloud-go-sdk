// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneSqlInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CloneSqlInstanceRequest
	GetName() *string
	SetTargetFolderId(v int64) *CloneSqlInstanceRequest
	GetTargetFolderId() *int64
}

type CloneSqlInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TargetFolderId *int64 `json:"targetFolderId,omitempty" xml:"targetFolderId,omitempty"`
}

func (s CloneSqlInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CloneSqlInstanceRequest) GoString() string {
	return s.String()
}

func (s *CloneSqlInstanceRequest) GetName() *string {
	return s.Name
}

func (s *CloneSqlInstanceRequest) GetTargetFolderId() *int64 {
	return s.TargetFolderId
}

func (s *CloneSqlInstanceRequest) SetName(v string) *CloneSqlInstanceRequest {
	s.Name = &v
	return s
}

func (s *CloneSqlInstanceRequest) SetTargetFolderId(v int64) *CloneSqlInstanceRequest {
	s.TargetFolderId = &v
	return s
}

func (s *CloneSqlInstanceRequest) Validate() error {
	return dara.Validate(s)
}
