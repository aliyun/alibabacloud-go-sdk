// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOptionValueForProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtensionCode(v string) *GetOptionValueForProjectRequest
	GetExtensionCode() *string
	SetProjectId(v string) *GetOptionValueForProjectRequest
	GetProjectId() *string
}

type GetOptionValueForProjectRequest struct {
	// The unique code of the extension.
	//
	// example:
	//
	// ce4*********086da5
	ExtensionCode *string `json:"ExtensionCode,omitempty" xml:"ExtensionCode,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 234
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetOptionValueForProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOptionValueForProjectRequest) GoString() string {
	return s.String()
}

func (s *GetOptionValueForProjectRequest) GetExtensionCode() *string {
	return s.ExtensionCode
}

func (s *GetOptionValueForProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetOptionValueForProjectRequest) SetExtensionCode(v string) *GetOptionValueForProjectRequest {
	s.ExtensionCode = &v
	return s
}

func (s *GetOptionValueForProjectRequest) SetProjectId(v string) *GetOptionValueForProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *GetOptionValueForProjectRequest) Validate() error {
	return dara.Validate(s)
}
