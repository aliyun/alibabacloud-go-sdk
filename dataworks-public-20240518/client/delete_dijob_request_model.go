// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDIJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDIJobId(v int64) *DeleteDIJobRequest
	GetDIJobId() *int64
	SetId(v int64) *DeleteDIJobRequest
	GetId() *int64
	SetProjectId(v int64) *DeleteDIJobRequest
	GetProjectId() *int64
}

type DeleteDIJobRequest struct {
	// Deprecated
	//
	// This parameter is deprecated. Use the Id parameter instead.
	//
	// example:
	//
	// 11126
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// 11126
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 108864
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteDIJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDIJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteDIJobRequest) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *DeleteDIJobRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteDIJobRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteDIJobRequest) SetDIJobId(v int64) *DeleteDIJobRequest {
	s.DIJobId = &v
	return s
}

func (s *DeleteDIJobRequest) SetId(v int64) *DeleteDIJobRequest {
	s.Id = &v
	return s
}

func (s *DeleteDIJobRequest) SetProjectId(v int64) *DeleteDIJobRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteDIJobRequest) Validate() error {
	return dara.Validate(s)
}
