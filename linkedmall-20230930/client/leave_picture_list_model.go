// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLeavePictureList interface {
	dara.Model
	String() string
	GoString() string
	SetDesc(v string) *LeavePictureList
	GetDesc() *string
	SetPicture(v string) *LeavePictureList
	GetPicture() *string
}

type LeavePictureList struct {
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// https://
	Picture *string `json:"picture,omitempty" xml:"picture,omitempty"`
}

func (s LeavePictureList) String() string {
	return dara.Prettify(s)
}

func (s LeavePictureList) GoString() string {
	return s.String()
}

func (s *LeavePictureList) GetDesc() *string {
	return s.Desc
}

func (s *LeavePictureList) GetPicture() *string {
	return s.Picture
}

func (s *LeavePictureList) SetDesc(v string) *LeavePictureList {
	s.Desc = &v
	return s
}

func (s *LeavePictureList) SetPicture(v string) *LeavePictureList {
	s.Picture = &v
	return s
}

func (s *LeavePictureList) Validate() error {
	return dara.Validate(s)
}
