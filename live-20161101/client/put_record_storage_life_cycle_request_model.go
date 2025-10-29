// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutRecordStorageLifeCycleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStreamIds(v []*string) *PutRecordStorageLifeCycleRequest
	GetStreamIds() []*string
	SetTag(v string) *PutRecordStorageLifeCycleRequest
	GetTag() *string
	SetUnixTimestamp(v int64) *PutRecordStorageLifeCycleRequest
	GetUnixTimestamp() *int64
}

type PutRecordStorageLifeCycleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// stream-xxx
	StreamIds []*string `json:"StreamIds,omitempty" xml:"StreamIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1095days
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1758621407
	UnixTimestamp *int64 `json:"UnixTimestamp,omitempty" xml:"UnixTimestamp,omitempty"`
}

func (s PutRecordStorageLifeCycleRequest) String() string {
	return dara.Prettify(s)
}

func (s PutRecordStorageLifeCycleRequest) GoString() string {
	return s.String()
}

func (s *PutRecordStorageLifeCycleRequest) GetStreamIds() []*string {
	return s.StreamIds
}

func (s *PutRecordStorageLifeCycleRequest) GetTag() *string {
	return s.Tag
}

func (s *PutRecordStorageLifeCycleRequest) GetUnixTimestamp() *int64 {
	return s.UnixTimestamp
}

func (s *PutRecordStorageLifeCycleRequest) SetStreamIds(v []*string) *PutRecordStorageLifeCycleRequest {
	s.StreamIds = v
	return s
}

func (s *PutRecordStorageLifeCycleRequest) SetTag(v string) *PutRecordStorageLifeCycleRequest {
	s.Tag = &v
	return s
}

func (s *PutRecordStorageLifeCycleRequest) SetUnixTimestamp(v int64) *PutRecordStorageLifeCycleRequest {
	s.UnixTimestamp = &v
	return s
}

func (s *PutRecordStorageLifeCycleRequest) Validate() error {
	return dara.Validate(s)
}
