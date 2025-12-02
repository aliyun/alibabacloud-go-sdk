// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *UpdateImageLibRequest
	GetComment() *string
	SetFreeInspection(v int32) *UpdateImageLibRequest
	GetFreeInspection() *int32
	SetLibId(v string) *UpdateImageLibRequest
	GetLibId() *string
	SetLibName(v string) *UpdateImageLibRequest
	GetLibName() *string
	SetRegionId(v string) *UpdateImageLibRequest
	GetRegionId() *string
}

type UpdateImageLibRequest struct {
	// Comment information for the library.
	//
	// example:
	//
	// 备注
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Exemption from review configuration.
	//
	// example:
	//
	// 0
	FreeInspection *int32 `json:"FreeInspection,omitempty" xml:"FreeInspection,omitempty"`
	// Library ID.
	//
	// example:
	//
	// custom_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// Library name.
	//
	// example:
	//
	// 测试图库
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateImageLibRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageLibRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageLibRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateImageLibRequest) GetFreeInspection() *int32 {
	return s.FreeInspection
}

func (s *UpdateImageLibRequest) GetLibId() *string {
	return s.LibId
}

func (s *UpdateImageLibRequest) GetLibName() *string {
	return s.LibName
}

func (s *UpdateImageLibRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateImageLibRequest) SetComment(v string) *UpdateImageLibRequest {
	s.Comment = &v
	return s
}

func (s *UpdateImageLibRequest) SetFreeInspection(v int32) *UpdateImageLibRequest {
	s.FreeInspection = &v
	return s
}

func (s *UpdateImageLibRequest) SetLibId(v string) *UpdateImageLibRequest {
	s.LibId = &v
	return s
}

func (s *UpdateImageLibRequest) SetLibName(v string) *UpdateImageLibRequest {
	s.LibName = &v
	return s
}

func (s *UpdateImageLibRequest) SetRegionId(v string) *UpdateImageLibRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateImageLibRequest) Validate() error {
	return dara.Validate(s)
}
