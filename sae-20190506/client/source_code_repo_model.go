// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSourceCodeRepo interface {
	dara.Model
	String() string
	GoString() string
	SetFullName(v string) *SourceCodeRepo
	GetFullName() *string
	SetId(v string) *SourceCodeRepo
	GetId() *string
	SetName(v string) *SourceCodeRepo
	GetName() *string
}

type SourceCodeRepo struct {
	FullName *string `json:"FullName,omitempty" xml:"FullName,omitempty"`
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SourceCodeRepo) String() string {
	return dara.Prettify(s)
}

func (s SourceCodeRepo) GoString() string {
	return s.String()
}

func (s *SourceCodeRepo) GetFullName() *string {
	return s.FullName
}

func (s *SourceCodeRepo) GetId() *string {
	return s.Id
}

func (s *SourceCodeRepo) GetName() *string {
	return s.Name
}

func (s *SourceCodeRepo) SetFullName(v string) *SourceCodeRepo {
	s.FullName = &v
	return s
}

func (s *SourceCodeRepo) SetId(v string) *SourceCodeRepo {
	s.Id = &v
	return s
}

func (s *SourceCodeRepo) SetName(v string) *SourceCodeRepo {
	s.Name = &v
	return s
}

func (s *SourceCodeRepo) Validate() error {
	return dara.Validate(s)
}
