// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChanges(v []*FullSchemaChange) *AlterTableRequest
	GetChanges() []*FullSchemaChange
}

type AlterTableRequest struct {
	Changes []*FullSchemaChange `json:"changes,omitempty" xml:"changes,omitempty" type:"Repeated"`
}

func (s AlterTableRequest) String() string {
	return dara.Prettify(s)
}

func (s AlterTableRequest) GoString() string {
	return s.String()
}

func (s *AlterTableRequest) GetChanges() []*FullSchemaChange {
	return s.Changes
}

func (s *AlterTableRequest) SetChanges(v []*FullSchemaChange) *AlterTableRequest {
	s.Changes = v
	return s
}

func (s *AlterTableRequest) Validate() error {
	if s.Changes != nil {
		for _, item := range s.Changes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
