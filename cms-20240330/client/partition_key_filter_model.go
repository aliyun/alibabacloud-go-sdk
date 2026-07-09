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
	// Returns only items whose partition key value equals this string.
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
