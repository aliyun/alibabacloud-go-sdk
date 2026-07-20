// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterCatalogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMissing(v []*string) *AlterCatalogResponseBody
	GetMissing() []*string
	SetRemoved(v []*string) *AlterCatalogResponseBody
	GetRemoved() []*string
	SetUpdated(v []*string) *AlterCatalogResponseBody
	GetUpdated() []*string
}

type AlterCatalogResponseBody struct {
	Missing []*string `json:"missing,omitempty" xml:"missing,omitempty" type:"Repeated"`
	Removed []*string `json:"removed,omitempty" xml:"removed,omitempty" type:"Repeated"`
	Updated []*string `json:"updated,omitempty" xml:"updated,omitempty" type:"Repeated"`
}

func (s AlterCatalogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AlterCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *AlterCatalogResponseBody) GetMissing() []*string {
	return s.Missing
}

func (s *AlterCatalogResponseBody) GetRemoved() []*string {
	return s.Removed
}

func (s *AlterCatalogResponseBody) GetUpdated() []*string {
	return s.Updated
}

func (s *AlterCatalogResponseBody) SetMissing(v []*string) *AlterCatalogResponseBody {
	s.Missing = v
	return s
}

func (s *AlterCatalogResponseBody) SetRemoved(v []*string) *AlterCatalogResponseBody {
	s.Removed = v
	return s
}

func (s *AlterCatalogResponseBody) SetUpdated(v []*string) *AlterCatalogResponseBody {
	s.Updated = v
	return s
}

func (s *AlterCatalogResponseBody) Validate() error {
	return dara.Validate(s)
}
