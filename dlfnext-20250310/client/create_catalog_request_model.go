// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCatalogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateCatalogRequest
	GetName() *string
	SetOptions(v map[string]*string) *CreateCatalogRequest
	GetOptions() map[string]*string
	SetType(v string) *CreateCatalogRequest
	GetType() *string
}

type CreateCatalogRequest struct {
	// example:
	//
	// catalog_demo
	Name    *string            `json:"name,omitempty" xml:"name,omitempty"`
	Options map[string]*string `json:"options,omitempty" xml:"options,omitempty"`
	Type    *string            `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateCatalogRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCatalogRequest) GoString() string {
	return s.String()
}

func (s *CreateCatalogRequest) GetName() *string {
	return s.Name
}

func (s *CreateCatalogRequest) GetOptions() map[string]*string {
	return s.Options
}

func (s *CreateCatalogRequest) GetType() *string {
	return s.Type
}

func (s *CreateCatalogRequest) SetName(v string) *CreateCatalogRequest {
	s.Name = &v
	return s
}

func (s *CreateCatalogRequest) SetOptions(v map[string]*string) *CreateCatalogRequest {
	s.Options = v
	return s
}

func (s *CreateCatalogRequest) SetType(v string) *CreateCatalogRequest {
	s.Type = &v
	return s
}

func (s *CreateCatalogRequest) Validate() error {
	return dara.Validate(s)
}
