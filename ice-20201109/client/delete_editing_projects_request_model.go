// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEditingProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectIds(v string) *DeleteEditingProjectsRequest
	GetProjectIds() *string
}

type DeleteEditingProjectsRequest struct {
	// The ID of the online editing project. You can specify multiple IDs separated with commas (,).
	//
	// example:
	//
	// ****fb2101bf24bf41cb318787dc****,****87dcfb2101bf24bf41cb3187****
	ProjectIds *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
}

func (s DeleteEditingProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEditingProjectsRequest) GoString() string {
	return s.String()
}

func (s *DeleteEditingProjectsRequest) GetProjectIds() *string {
	return s.ProjectIds
}

func (s *DeleteEditingProjectsRequest) SetProjectIds(v string) *DeleteEditingProjectsRequest {
	s.ProjectIds = &v
	return s
}

func (s *DeleteEditingProjectsRequest) Validate() error {
	return dara.Validate(s)
}
