// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterCatalogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRemovals(v []*string) *AlterCatalogRequest
	GetRemovals() []*string
	SetUpdates(v map[string]*string) *AlterCatalogRequest
	GetUpdates() map[string]*string
}

type AlterCatalogRequest struct {
	Removals []*string          `json:"removals,omitempty" xml:"removals,omitempty" type:"Repeated"`
	Updates  map[string]*string `json:"updates,omitempty" xml:"updates,omitempty"`
}

func (s AlterCatalogRequest) String() string {
	return dara.Prettify(s)
}

func (s AlterCatalogRequest) GoString() string {
	return s.String()
}

func (s *AlterCatalogRequest) GetRemovals() []*string {
	return s.Removals
}

func (s *AlterCatalogRequest) GetUpdates() map[string]*string {
	return s.Updates
}

func (s *AlterCatalogRequest) SetRemovals(v []*string) *AlterCatalogRequest {
	s.Removals = v
	return s
}

func (s *AlterCatalogRequest) SetUpdates(v map[string]*string) *AlterCatalogRequest {
	s.Updates = v
	return s
}

func (s *AlterCatalogRequest) Validate() error {
	return dara.Validate(s)
}
