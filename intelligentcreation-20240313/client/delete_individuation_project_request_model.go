// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIndividuationProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *DeleteIndividuationProjectRequest
	GetProjectId() *string
}

type DeleteIndividuationProjectRequest struct {
	// example:
	//
	// 840015278620459008
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s DeleteIndividuationProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIndividuationProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteIndividuationProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DeleteIndividuationProjectRequest) SetProjectId(v string) *DeleteIndividuationProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteIndividuationProjectRequest) Validate() error {
	return dara.Validate(s)
}
