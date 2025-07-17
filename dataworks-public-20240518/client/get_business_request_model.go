// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBusinessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessId(v int64) *GetBusinessRequest
	GetBusinessId() *int64
	SetProjectId(v int64) *GetBusinessRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *GetBusinessRequest
	GetProjectIdentifier() *string
}

type GetBusinessRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1000000111
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
}

func (s GetBusinessRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBusinessRequest) GoString() string {
	return s.String()
}

func (s *GetBusinessRequest) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *GetBusinessRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetBusinessRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *GetBusinessRequest) SetBusinessId(v int64) *GetBusinessRequest {
	s.BusinessId = &v
	return s
}

func (s *GetBusinessRequest) SetProjectId(v int64) *GetBusinessRequest {
	s.ProjectId = &v
	return s
}

func (s *GetBusinessRequest) SetProjectIdentifier(v string) *GetBusinessRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *GetBusinessRequest) Validate() error {
	return dara.Validate(s)
}
