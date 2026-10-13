// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoiceCloneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListVoiceCloneRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListVoiceCloneRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVoiceCloneRequest
	GetPageSize() *int32
	SetStatus(v string) *ListVoiceCloneRequest
	GetStatus() *string
}

type ListVoiceCloneRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Pages start from 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The filter condition for the voice status. Valid values: Enabled (available) and Disabled (unavailable).
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListVoiceCloneRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceCloneRequest) GoString() string {
	return s.String()
}

func (s *ListVoiceCloneRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVoiceCloneRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVoiceCloneRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVoiceCloneRequest) GetStatus() *string {
	return s.Status
}

func (s *ListVoiceCloneRequest) SetInstanceId(v string) *ListVoiceCloneRequest {
	s.InstanceId = &v
	return s
}

func (s *ListVoiceCloneRequest) SetPageNumber(v int32) *ListVoiceCloneRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVoiceCloneRequest) SetPageSize(v int32) *ListVoiceCloneRequest {
	s.PageSize = &v
	return s
}

func (s *ListVoiceCloneRequest) SetStatus(v string) *ListVoiceCloneRequest {
	s.Status = &v
	return s
}

func (s *ListVoiceCloneRequest) Validate() error {
	return dara.Validate(s)
}
