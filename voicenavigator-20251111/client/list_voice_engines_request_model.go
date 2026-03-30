// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoiceEnginesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListVoiceEnginesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVoiceEnginesRequest
	GetPageSize() *int32
}

type ListVoiceEnginesRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListVoiceEnginesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceEnginesRequest) GoString() string {
	return s.String()
}

func (s *ListVoiceEnginesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVoiceEnginesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVoiceEnginesRequest) SetPageNumber(v int32) *ListVoiceEnginesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVoiceEnginesRequest) SetPageSize(v int32) *ListVoiceEnginesRequest {
	s.PageSize = &v
	return s
}

func (s *ListVoiceEnginesRequest) Validate() error {
	return dara.Validate(s)
}
