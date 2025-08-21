// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRenderingProjectInstanceStateMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *GetRenderingProjectInstanceStateMetricsRequest
	GetProjectId() *string
}

type GetRenderingProjectInstanceStateMetricsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// project-422bc38dfgh5eb44149f135ef76304f63b
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetRenderingProjectInstanceStateMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRenderingProjectInstanceStateMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetRenderingProjectInstanceStateMetricsRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetRenderingProjectInstanceStateMetricsRequest) SetProjectId(v string) *GetRenderingProjectInstanceStateMetricsRequest {
	s.ProjectId = &v
	return s
}

func (s *GetRenderingProjectInstanceStateMetricsRequest) Validate() error {
	return dara.Validate(s)
}
