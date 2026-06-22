// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageRegistryRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListImageRegistryRegionRequest
	GetLang() *string
}

type ListImageRegistryRegionRequest struct {
	// The language type for requesting and receiving messages. Default value: zh. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ListImageRegistryRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImageRegistryRegionRequest) GoString() string {
	return s.String()
}

func (s *ListImageRegistryRegionRequest) GetLang() *string {
	return s.Lang
}

func (s *ListImageRegistryRegionRequest) SetLang(v string) *ListImageRegistryRegionRequest {
	s.Lang = &v
	return s
}

func (s *ListImageRegistryRegionRequest) Validate() error {
	return dara.Validate(s)
}
