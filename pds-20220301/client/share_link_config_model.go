// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShareLinkConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnableShareLinkOfficeEdit(v bool) *ShareLinkConfig
	GetEnableShareLinkOfficeEdit() *bool
}

type ShareLinkConfig struct {
	EnableShareLinkOfficeEdit *bool `json:"enable_share_link_office_edit,omitempty" xml:"enable_share_link_office_edit,omitempty"`
}

func (s ShareLinkConfig) String() string {
	return dara.Prettify(s)
}

func (s ShareLinkConfig) GoString() string {
	return s.String()
}

func (s *ShareLinkConfig) GetEnableShareLinkOfficeEdit() *bool {
	return s.EnableShareLinkOfficeEdit
}

func (s *ShareLinkConfig) SetEnableShareLinkOfficeEdit(v bool) *ShareLinkConfig {
	s.EnableShareLinkOfficeEdit = &v
	return s
}

func (s *ShareLinkConfig) Validate() error {
	return dara.Validate(s)
}
