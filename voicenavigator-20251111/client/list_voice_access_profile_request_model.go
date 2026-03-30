// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoiceAccessProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListVoiceAccessProfileRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListVoiceAccessProfileRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVoiceAccessProfileRequest
	GetPageSize() *int32
}

type ListVoiceAccessProfileRequest struct {
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListVoiceAccessProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceAccessProfileRequest) GoString() string {
	return s.String()
}

func (s *ListVoiceAccessProfileRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVoiceAccessProfileRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVoiceAccessProfileRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVoiceAccessProfileRequest) SetInstanceId(v string) *ListVoiceAccessProfileRequest {
	s.InstanceId = &v
	return s
}

func (s *ListVoiceAccessProfileRequest) SetPageNumber(v int32) *ListVoiceAccessProfileRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVoiceAccessProfileRequest) SetPageSize(v int32) *ListVoiceAccessProfileRequest {
	s.PageSize = &v
	return s
}

func (s *ListVoiceAccessProfileRequest) Validate() error {
	return dara.Validate(s)
}
