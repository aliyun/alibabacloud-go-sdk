// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeaturesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNames(v string) *ListFeaturesRequest
	GetNames() *string
}

type ListFeaturesRequest struct {
	// example:
	//
	// PaiConsole:IntegrateWithWorkspace
	Names *string `json:"Names,omitempty" xml:"Names,omitempty"`
}

func (s ListFeaturesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFeaturesRequest) GoString() string {
	return s.String()
}

func (s *ListFeaturesRequest) GetNames() *string {
	return s.Names
}

func (s *ListFeaturesRequest) SetNames(v string) *ListFeaturesRequest {
	s.Names = &v
	return s
}

func (s *ListFeaturesRequest) Validate() error {
	return dara.Validate(s)
}
