// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoiceAccessProfilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListVoiceAccessProfilesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListVoiceAccessProfilesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVoiceAccessProfilesRequest
	GetPageSize() *int32
}

type ListVoiceAccessProfilesRequest struct {
	// 实例ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 页码，从1开始
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页记录数
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListVoiceAccessProfilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceAccessProfilesRequest) GoString() string {
	return s.String()
}

func (s *ListVoiceAccessProfilesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVoiceAccessProfilesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVoiceAccessProfilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVoiceAccessProfilesRequest) SetInstanceId(v string) *ListVoiceAccessProfilesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListVoiceAccessProfilesRequest) SetPageNumber(v int32) *ListVoiceAccessProfilesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVoiceAccessProfilesRequest) SetPageSize(v int32) *ListVoiceAccessProfilesRequest {
	s.PageSize = &v
	return s
}

func (s *ListVoiceAccessProfilesRequest) Validate() error {
	return dara.Validate(s)
}
