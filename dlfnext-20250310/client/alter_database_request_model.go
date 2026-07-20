// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRemovals(v []*string) *AlterDatabaseRequest
	GetRemovals() []*string
	SetUpdates(v map[string]*string) *AlterDatabaseRequest
	GetUpdates() map[string]*string
}

type AlterDatabaseRequest struct {
	Removals []*string          `json:"removals,omitempty" xml:"removals,omitempty" type:"Repeated"`
	Updates  map[string]*string `json:"updates,omitempty" xml:"updates,omitempty"`
}

func (s AlterDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s AlterDatabaseRequest) GoString() string {
	return s.String()
}

func (s *AlterDatabaseRequest) GetRemovals() []*string {
	return s.Removals
}

func (s *AlterDatabaseRequest) GetUpdates() map[string]*string {
	return s.Updates
}

func (s *AlterDatabaseRequest) SetRemovals(v []*string) *AlterDatabaseRequest {
	s.Removals = v
	return s
}

func (s *AlterDatabaseRequest) SetUpdates(v map[string]*string) *AlterDatabaseRequest {
	s.Updates = v
	return s
}

func (s *AlterDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
