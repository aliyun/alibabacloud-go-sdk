// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoiceAccessProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *ListVoiceAccessProfileRequest
	GetBusinessUnitId() *string
	SetPageNumber(v int32) *ListVoiceAccessProfileRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVoiceAccessProfileRequest
	GetPageSize() *int32
}

type ListVoiceAccessProfileRequest struct {
	// example:
	//
	// llm-c11iig67g863rih8
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
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

func (s *ListVoiceAccessProfileRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *ListVoiceAccessProfileRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVoiceAccessProfileRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVoiceAccessProfileRequest) SetBusinessUnitId(v string) *ListVoiceAccessProfileRequest {
	s.BusinessUnitId = &v
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
