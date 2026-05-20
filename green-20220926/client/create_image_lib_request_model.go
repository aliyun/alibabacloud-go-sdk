// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateImageLibRequest
	GetComment() *string
	SetLibName(v string) *CreateImageLibRequest
	GetLibName() *string
	SetRegionId(v string) *CreateImageLibRequest
	GetRegionId() *string
}

type CreateImageLibRequest struct {
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateImageLibRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageLibRequest) GoString() string {
	return s.String()
}

func (s *CreateImageLibRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateImageLibRequest) GetLibName() *string {
	return s.LibName
}

func (s *CreateImageLibRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateImageLibRequest) SetComment(v string) *CreateImageLibRequest {
	s.Comment = &v
	return s
}

func (s *CreateImageLibRequest) SetLibName(v string) *CreateImageLibRequest {
	s.LibName = &v
	return s
}

func (s *CreateImageLibRequest) SetRegionId(v string) *CreateImageLibRequest {
	s.RegionId = &v
	return s
}

func (s *CreateImageLibRequest) Validate() error {
	return dara.Validate(s)
}
