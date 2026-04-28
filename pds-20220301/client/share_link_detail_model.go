// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShareLinkDetail interface {
	dara.Model
	String() string
	GoString() string
	SetEnableOfficeEditable(v bool) *ShareLinkDetail
	GetEnableOfficeEditable() *bool
}

type ShareLinkDetail struct {
	EnableOfficeEditable *bool `json:"enable_office_editable,omitempty" xml:"enable_office_editable,omitempty"`
}

func (s ShareLinkDetail) String() string {
	return dara.Prettify(s)
}

func (s ShareLinkDetail) GoString() string {
	return s.String()
}

func (s *ShareLinkDetail) GetEnableOfficeEditable() *bool {
	return s.EnableOfficeEditable
}

func (s *ShareLinkDetail) SetEnableOfficeEditable(v bool) *ShareLinkDetail {
	s.EnableOfficeEditable = &v
	return s
}

func (s *ShareLinkDetail) Validate() error {
	return dara.Validate(s)
}
