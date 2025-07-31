// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImageLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *AddImageLibRequest
	GetComment() *string
	SetLibName(v string) *AddImageLibRequest
	GetLibName() *string
	SetRegionId(v string) *AddImageLibRequest
	GetRegionId() *string
}

type AddImageLibRequest struct {
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddImageLibRequest) String() string {
	return dara.Prettify(s)
}

func (s AddImageLibRequest) GoString() string {
	return s.String()
}

func (s *AddImageLibRequest) GetComment() *string {
	return s.Comment
}

func (s *AddImageLibRequest) GetLibName() *string {
	return s.LibName
}

func (s *AddImageLibRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddImageLibRequest) SetComment(v string) *AddImageLibRequest {
	s.Comment = &v
	return s
}

func (s *AddImageLibRequest) SetLibName(v string) *AddImageLibRequest {
	s.LibName = &v
	return s
}

func (s *AddImageLibRequest) SetRegionId(v string) *AddImageLibRequest {
	s.RegionId = &v
	return s
}

func (s *AddImageLibRequest) Validate() error {
	return dara.Validate(s)
}
