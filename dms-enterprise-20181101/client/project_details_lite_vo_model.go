// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProjectDetailsLiteVO interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ProjectDetailsLiteVO
	GetId() *int64
	SetProjectName(v string) *ProjectDetailsLiteVO
	GetProjectName() *string
}

type ProjectDetailsLiteVO struct {
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ProjectDetailsLiteVO) String() string {
	return dara.Prettify(s)
}

func (s ProjectDetailsLiteVO) GoString() string {
	return s.String()
}

func (s *ProjectDetailsLiteVO) GetId() *int64 {
	return s.Id
}

func (s *ProjectDetailsLiteVO) GetProjectName() *string {
	return s.ProjectName
}

func (s *ProjectDetailsLiteVO) SetId(v int64) *ProjectDetailsLiteVO {
	s.Id = &v
	return s
}

func (s *ProjectDetailsLiteVO) SetProjectName(v string) *ProjectDetailsLiteVO {
	s.ProjectName = &v
	return s
}

func (s *ProjectDetailsLiteVO) Validate() error {
	return dara.Validate(s)
}
