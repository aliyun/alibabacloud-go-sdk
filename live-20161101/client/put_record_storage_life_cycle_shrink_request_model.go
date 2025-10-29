// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutRecordStorageLifeCycleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStreamIdsShrink(v string) *PutRecordStorageLifeCycleShrinkRequest
	GetStreamIdsShrink() *string
	SetTag(v string) *PutRecordStorageLifeCycleShrinkRequest
	GetTag() *string
	SetUnixTimestamp(v int64) *PutRecordStorageLifeCycleShrinkRequest
	GetUnixTimestamp() *int64
}

type PutRecordStorageLifeCycleShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// stream-xxx
	StreamIdsShrink *string `json:"StreamIds,omitempty" xml:"StreamIds,omitempty"`
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

func (s PutRecordStorageLifeCycleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PutRecordStorageLifeCycleShrinkRequest) GoString() string {
	return s.String()
}

func (s *PutRecordStorageLifeCycleShrinkRequest) GetStreamIdsShrink() *string {
	return s.StreamIdsShrink
}

func (s *PutRecordStorageLifeCycleShrinkRequest) GetTag() *string {
	return s.Tag
}

func (s *PutRecordStorageLifeCycleShrinkRequest) GetUnixTimestamp() *int64 {
	return s.UnixTimestamp
}

func (s *PutRecordStorageLifeCycleShrinkRequest) SetStreamIdsShrink(v string) *PutRecordStorageLifeCycleShrinkRequest {
	s.StreamIdsShrink = &v
	return s
}

func (s *PutRecordStorageLifeCycleShrinkRequest) SetTag(v string) *PutRecordStorageLifeCycleShrinkRequest {
	s.Tag = &v
	return s
}

func (s *PutRecordStorageLifeCycleShrinkRequest) SetUnixTimestamp(v int64) *PutRecordStorageLifeCycleShrinkRequest {
	s.UnixTimestamp = &v
	return s
}

func (s *PutRecordStorageLifeCycleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
