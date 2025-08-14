// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEditingProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *GetEditingProjectRequest
	GetProjectId() *string
	SetRequestSource(v string) *GetEditingProjectRequest
	GetRequestSource() *string
}

type GetEditingProjectRequest struct {
	// The ID of the online editing project.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****fb2101bf24b2754cb318787dc****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the request source. Valid values:
	//
	// \\- OpenAPI (default): Timeline conversion is not performed.
	//
	// \\- WebSDK: If you specify this value, the project timeline is automatically converted into the frontend style, and the materials in the timeline are associated with the project to enable preview by using frontend web SDKs.
	//
	// example:
	//
	// WebSDK
	RequestSource *string `json:"RequestSource,omitempty" xml:"RequestSource,omitempty"`
}

func (s GetEditingProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *GetEditingProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetEditingProjectRequest) GetRequestSource() *string {
	return s.RequestSource
}

func (s *GetEditingProjectRequest) SetProjectId(v string) *GetEditingProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *GetEditingProjectRequest) SetRequestSource(v string) *GetEditingProjectRequest {
	s.RequestSource = &v
	return s
}

func (s *GetEditingProjectRequest) Validate() error {
	return dara.Validate(s)
}
