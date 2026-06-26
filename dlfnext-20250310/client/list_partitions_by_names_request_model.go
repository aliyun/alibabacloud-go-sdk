// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPartitionsByNamesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSpecs(v []map[string]*string) *ListPartitionsByNamesRequest
	GetSpecs() []map[string]*string
}

type ListPartitionsByNamesRequest struct {
	// 分区规格列表。
	Specs []map[string]*string `json:"specs,omitempty" xml:"specs,omitempty" type:"Repeated"`
}

func (s ListPartitionsByNamesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPartitionsByNamesRequest) GoString() string {
	return s.String()
}

func (s *ListPartitionsByNamesRequest) GetSpecs() []map[string]*string {
	return s.Specs
}

func (s *ListPartitionsByNamesRequest) SetSpecs(v []map[string]*string) *ListPartitionsByNamesRequest {
	s.Specs = v
	return s
}

func (s *ListPartitionsByNamesRequest) Validate() error {
	return dara.Validate(s)
}
