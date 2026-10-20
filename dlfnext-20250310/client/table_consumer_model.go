// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTableConsumer interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerId(v string) *TableConsumer
	GetConsumerId() *string
	SetLastModifiedTime(v int64) *TableConsumer
	GetLastModifiedTime() *int64
	SetNextSnapshot(v int64) *TableConsumer
	GetNextSnapshot() *int64
}

type TableConsumer struct {
	// Consumer ID
	//
	// example:
	//
	// paimon-table1-id
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// The last modified time of the consumer file in Unix milliseconds.
	//
	// example:
	//
	// 1791792221185
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// The ID of the next snapshot to be consumed.
	//
	// example:
	//
	// 123
	NextSnapshot *int64 `json:"nextSnapshot,omitempty" xml:"nextSnapshot,omitempty"`
}

func (s TableConsumer) String() string {
	return dara.Prettify(s)
}

func (s TableConsumer) GoString() string {
	return s.String()
}

func (s *TableConsumer) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *TableConsumer) GetLastModifiedTime() *int64 {
	return s.LastModifiedTime
}

func (s *TableConsumer) GetNextSnapshot() *int64 {
	return s.NextSnapshot
}

func (s *TableConsumer) SetConsumerId(v string) *TableConsumer {
	s.ConsumerId = &v
	return s
}

func (s *TableConsumer) SetLastModifiedTime(v int64) *TableConsumer {
	s.LastModifiedTime = &v
	return s
}

func (s *TableConsumer) SetNextSnapshot(v int64) *TableConsumer {
	s.NextSnapshot = &v
	return s
}

func (s *TableConsumer) Validate() error {
	return dara.Validate(s)
}
