// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDLPartitionInput interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int32) *DLPartitionInput
	GetCreateTime() *int32
	SetLastAccessTime(v int32) *DLPartitionInput
	GetLastAccessTime() *int32
	SetParameters(v map[string]*string) *DLPartitionInput
	GetParameters() map[string]*string
	SetStorageDescriptor(v *DLStorageDescriptor) *DLPartitionInput
	GetStorageDescriptor() *DLStorageDescriptor
	SetValues(v []*string) *DLPartitionInput
	GetValues() []*string
}

type DLPartitionInput struct {
	CreateTime        *int32               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LastAccessTime    *int32               `json:"LastAccessTime,omitempty" xml:"LastAccessTime,omitempty"`
	Parameters        map[string]*string   `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	StorageDescriptor *DLStorageDescriptor `json:"StorageDescriptor,omitempty" xml:"StorageDescriptor,omitempty"`
	Values            []*string            `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DLPartitionInput) String() string {
	return dara.Prettify(s)
}

func (s DLPartitionInput) GoString() string {
	return s.String()
}

func (s *DLPartitionInput) GetCreateTime() *int32 {
	return s.CreateTime
}

func (s *DLPartitionInput) GetLastAccessTime() *int32 {
	return s.LastAccessTime
}

func (s *DLPartitionInput) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *DLPartitionInput) GetStorageDescriptor() *DLStorageDescriptor {
	return s.StorageDescriptor
}

func (s *DLPartitionInput) GetValues() []*string {
	return s.Values
}

func (s *DLPartitionInput) SetCreateTime(v int32) *DLPartitionInput {
	s.CreateTime = &v
	return s
}

func (s *DLPartitionInput) SetLastAccessTime(v int32) *DLPartitionInput {
	s.LastAccessTime = &v
	return s
}

func (s *DLPartitionInput) SetParameters(v map[string]*string) *DLPartitionInput {
	s.Parameters = v
	return s
}

func (s *DLPartitionInput) SetStorageDescriptor(v *DLStorageDescriptor) *DLPartitionInput {
	s.StorageDescriptor = v
	return s
}

func (s *DLPartitionInput) SetValues(v []*string) *DLPartitionInput {
	s.Values = v
	return s
}

func (s *DLPartitionInput) Validate() error {
	if s.StorageDescriptor != nil {
		if err := s.StorageDescriptor.Validate(); err != nil {
			return err
		}
	}
	return nil
}
