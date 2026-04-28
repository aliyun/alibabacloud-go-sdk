// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOfficePreviewUrlOption interface {
	dara.Model
	String() string
	GoString() string
	SetCopy(v bool) *GetOfficePreviewUrlOption
	GetCopy() *bool
	SetPrint(v bool) *GetOfficePreviewUrlOption
	GetPrint() *bool
}

type GetOfficePreviewUrlOption struct {
	Copy  *bool `json:"copy,omitempty" xml:"copy,omitempty"`
	Print *bool `json:"print,omitempty" xml:"print,omitempty"`
}

func (s GetOfficePreviewUrlOption) String() string {
	return dara.Prettify(s)
}

func (s GetOfficePreviewUrlOption) GoString() string {
	return s.String()
}

func (s *GetOfficePreviewUrlOption) GetCopy() *bool {
	return s.Copy
}

func (s *GetOfficePreviewUrlOption) GetPrint() *bool {
	return s.Print
}

func (s *GetOfficePreviewUrlOption) SetCopy(v bool) *GetOfficePreviewUrlOption {
	s.Copy = &v
	return s
}

func (s *GetOfficePreviewUrlOption) SetPrint(v bool) *GetOfficePreviewUrlOption {
	s.Print = &v
	return s
}

func (s *GetOfficePreviewUrlOption) Validate() error {
	return dara.Validate(s)
}
