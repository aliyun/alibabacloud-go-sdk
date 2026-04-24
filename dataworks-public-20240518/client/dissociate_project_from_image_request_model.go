// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateProjectFromImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DissociateProjectFromImageRequest
	GetId() *string
	SetProjectId(v int64) *DissociateProjectFromImageRequest
	GetProjectId() *int64
}

type DissociateProjectFromImageRequest struct {
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

func (s DissociateProjectFromImageRequest) String() string {
	return dara.Prettify(s)
}

func (s DissociateProjectFromImageRequest) GoString() string {
	return s.String()
}

func (s *DissociateProjectFromImageRequest) GetId() *string {
	return s.Id
}

func (s *DissociateProjectFromImageRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DissociateProjectFromImageRequest) SetId(v string) *DissociateProjectFromImageRequest {
	s.Id = &v
	return s
}

func (s *DissociateProjectFromImageRequest) SetProjectId(v int64) *DissociateProjectFromImageRequest {
	s.ProjectId = &v
	return s
}

func (s *DissociateProjectFromImageRequest) Validate() error {
	return dara.Validate(s)
}
