// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelGroupClientKeyItemDTO interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ModelGroupClientKeyItemDTO
	GetId() *int64
	SetKeyPreview(v string) *ModelGroupClientKeyItemDTO
	GetKeyPreview() *string
	SetName(v string) *ModelGroupClientKeyItemDTO
	GetName() *string
}

type ModelGroupClientKeyItemDTO struct {
	// example:
	//
	// 501
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// sk-ab****yz
	KeyPreview *string `json:"keyPreview,omitempty" xml:"keyPreview,omitempty"`
	// example:
	//
	// Default Key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ModelGroupClientKeyItemDTO) String() string {
	return dara.Prettify(s)
}

func (s ModelGroupClientKeyItemDTO) GoString() string {
	return s.String()
}

func (s *ModelGroupClientKeyItemDTO) GetId() *int64 {
	return s.Id
}

func (s *ModelGroupClientKeyItemDTO) GetKeyPreview() *string {
	return s.KeyPreview
}

func (s *ModelGroupClientKeyItemDTO) GetName() *string {
	return s.Name
}

func (s *ModelGroupClientKeyItemDTO) SetId(v int64) *ModelGroupClientKeyItemDTO {
	s.Id = &v
	return s
}

func (s *ModelGroupClientKeyItemDTO) SetKeyPreview(v string) *ModelGroupClientKeyItemDTO {
	s.KeyPreview = &v
	return s
}

func (s *ModelGroupClientKeyItemDTO) SetName(v string) *ModelGroupClientKeyItemDTO {
	s.Name = &v
	return s
}

func (s *ModelGroupClientKeyItemDTO) Validate() error {
	return dara.Validate(s)
}
