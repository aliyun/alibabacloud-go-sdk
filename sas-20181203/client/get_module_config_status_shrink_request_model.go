// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModuleConfigStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModuleNamesShrink(v string) *GetModuleConfigStatusShrinkRequest
	GetModuleNamesShrink() *string
}

type GetModuleConfigStatusShrinkRequest struct {
	// The service modules that you want to query.
	//
	// This parameter is required.
	ModuleNamesShrink *string `json:"ModuleNames,omitempty" xml:"ModuleNames,omitempty"`
}

func (s GetModuleConfigStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModuleConfigStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetModuleConfigStatusShrinkRequest) GetModuleNamesShrink() *string {
	return s.ModuleNamesShrink
}

func (s *GetModuleConfigStatusShrinkRequest) SetModuleNamesShrink(v string) *GetModuleConfigStatusShrinkRequest {
	s.ModuleNamesShrink = &v
	return s
}

func (s *GetModuleConfigStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}
