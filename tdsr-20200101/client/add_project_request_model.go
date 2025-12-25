// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessId(v int64) *AddProjectRequest
	GetBusinessId() *int64
	SetName(v string) *AddProjectRequest
	GetName() *string
}

type AddProjectRequest struct {
	// example:
	//
	// 5432****
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s AddProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s AddProjectRequest) GoString() string {
	return s.String()
}

func (s *AddProjectRequest) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *AddProjectRequest) GetName() *string {
	return s.Name
}

func (s *AddProjectRequest) SetBusinessId(v int64) *AddProjectRequest {
	s.BusinessId = &v
	return s
}

func (s *AddProjectRequest) SetName(v string) *AddProjectRequest {
	s.Name = &v
	return s
}

func (s *AddProjectRequest) Validate() error {
	return dara.Validate(s)
}
