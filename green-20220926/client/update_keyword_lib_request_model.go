// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKeywordLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLibId(v string) *UpdateKeywordLibRequest
	GetLibId() *string
	SetLibName(v string) *UpdateKeywordLibRequest
	GetLibName() *string
	SetRegionId(v string) *UpdateKeywordLibRequest
	GetRegionId() *string
}

type UpdateKeywordLibRequest struct {
	// Library ID.
	//
	// example:
	//
	// custom_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// Keyword library name.
	//
	// example:
	//
	// 测试库
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateKeywordLibRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKeywordLibRequest) GoString() string {
	return s.String()
}

func (s *UpdateKeywordLibRequest) GetLibId() *string {
	return s.LibId
}

func (s *UpdateKeywordLibRequest) GetLibName() *string {
	return s.LibName
}

func (s *UpdateKeywordLibRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateKeywordLibRequest) SetLibId(v string) *UpdateKeywordLibRequest {
	s.LibId = &v
	return s
}

func (s *UpdateKeywordLibRequest) SetLibName(v string) *UpdateKeywordLibRequest {
	s.LibName = &v
	return s
}

func (s *UpdateKeywordLibRequest) SetRegionId(v string) *UpdateKeywordLibRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateKeywordLibRequest) Validate() error {
	return dara.Validate(s)
}
