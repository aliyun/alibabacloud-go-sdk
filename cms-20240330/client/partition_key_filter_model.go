// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPartitionKeyFilter interface {
	dara.Model
	String() string
	GoString() string
	SetEq(v string) *PartitionKeyFilter
	GetEq() *string
}

type PartitionKeyFilter struct {
	// The exact match value.
	//
	// example:
	//
	// workspace
	Eq *string `json:"eq,omitempty" xml:"eq,omitempty"`
}

func (s PartitionKeyFilter) String() string {
	return dara.Prettify(s)
}

func (s PartitionKeyFilter) GoString() string {
	return s.String()
}

func (s *PartitionKeyFilter) GetEq() *string {
	return s.Eq
}

func (s *PartitionKeyFilter) SetEq(v string) *PartitionKeyFilter {
	s.Eq = &v
	return s
}

func (s *PartitionKeyFilter) Validate() error {
	return dara.Validate(s)
}
