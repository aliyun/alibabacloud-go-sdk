// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateProjectToImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *AssociateProjectToImageRequest
	GetId() *string
	SetProjectId(v int64) *AssociateProjectToImageRequest
	GetProjectId() *int64
}

type AssociateProjectToImageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Custom_image_xxxx_xxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100000001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s AssociateProjectToImageRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateProjectToImageRequest) GoString() string {
	return s.String()
}

func (s *AssociateProjectToImageRequest) GetId() *string {
	return s.Id
}

func (s *AssociateProjectToImageRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *AssociateProjectToImageRequest) SetId(v string) *AssociateProjectToImageRequest {
	s.Id = &v
	return s
}

func (s *AssociateProjectToImageRequest) SetProjectId(v int64) *AssociateProjectToImageRequest {
	s.ProjectId = &v
	return s
}

func (s *AssociateProjectToImageRequest) Validate() error {
	return dara.Validate(s)
}
