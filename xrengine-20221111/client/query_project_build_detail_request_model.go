// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProjectBuildDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *QueryProjectBuildDetailRequest
	GetProjectId() *string
}

type QueryProjectBuildDetailRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s QueryProjectBuildDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectBuildDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryProjectBuildDetailRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *QueryProjectBuildDetailRequest) SetProjectId(v string) *QueryProjectBuildDetailRequest {
	s.ProjectId = &v
	return s
}

func (s *QueryProjectBuildDetailRequest) Validate() error {
	return dara.Validate(s)
}
