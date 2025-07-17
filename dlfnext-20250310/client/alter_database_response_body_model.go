// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMissing(v []*string) *AlterDatabaseResponseBody
	GetMissing() []*string
	SetRemoved(v []*string) *AlterDatabaseResponseBody
	GetRemoved() []*string
	SetUpdated(v []*string) *AlterDatabaseResponseBody
	GetUpdated() []*string
}

type AlterDatabaseResponseBody struct {
	Missing []*string `json:"missing,omitempty" xml:"missing,omitempty" type:"Repeated"`
	Removed []*string `json:"removed,omitempty" xml:"removed,omitempty" type:"Repeated"`
	Updated []*string `json:"updated,omitempty" xml:"updated,omitempty" type:"Repeated"`
}

func (s AlterDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AlterDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *AlterDatabaseResponseBody) GetMissing() []*string {
	return s.Missing
}

func (s *AlterDatabaseResponseBody) GetRemoved() []*string {
	return s.Removed
}

func (s *AlterDatabaseResponseBody) GetUpdated() []*string {
	return s.Updated
}

func (s *AlterDatabaseResponseBody) SetMissing(v []*string) *AlterDatabaseResponseBody {
	s.Missing = v
	return s
}

func (s *AlterDatabaseResponseBody) SetRemoved(v []*string) *AlterDatabaseResponseBody {
	s.Removed = v
	return s
}

func (s *AlterDatabaseResponseBody) SetUpdated(v []*string) *AlterDatabaseResponseBody {
	s.Updated = v
	return s
}

func (s *AlterDatabaseResponseBody) Validate() error {
	return dara.Validate(s)
}
