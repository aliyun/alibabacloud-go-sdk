// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPartitionError interface {
	dara.Model
	String() string
	GoString() string
	SetErrorDetail(v string) *PartitionError
	GetErrorDetail() *string
	SetValues(v []*string) *PartitionError
	GetValues() []*string
}

type PartitionError struct {
	ErrorDetail *string   `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty"`
	Values      []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s PartitionError) String() string {
	return dara.Prettify(s)
}

func (s PartitionError) GoString() string {
	return s.String()
}

func (s *PartitionError) GetErrorDetail() *string {
	return s.ErrorDetail
}

func (s *PartitionError) GetValues() []*string {
	return s.Values
}

func (s *PartitionError) SetErrorDetail(v string) *PartitionError {
	s.ErrorDetail = &v
	return s
}

func (s *PartitionError) SetValues(v []*string) *PartitionError {
	s.Values = v
	return s
}

func (s *PartitionError) Validate() error {
	return dara.Validate(s)
}
