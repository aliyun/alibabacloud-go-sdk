// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIAgentPhoneNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNumber(v string) *ListAIAgentPhoneNumberRequest
	GetNumber() *string
	SetPageNumber(v int64) *ListAIAgentPhoneNumberRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAIAgentPhoneNumberRequest
	GetPageSize() *int64
	SetStatus(v int32) *ListAIAgentPhoneNumberRequest
	GetStatus() *int32
}

type ListAIAgentPhoneNumberRequest struct {
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status   *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAIAgentPhoneNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAIAgentPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *ListAIAgentPhoneNumberRequest) GetNumber() *string {
	return s.Number
}

func (s *ListAIAgentPhoneNumberRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAIAgentPhoneNumberRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAIAgentPhoneNumberRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListAIAgentPhoneNumberRequest) SetNumber(v string) *ListAIAgentPhoneNumberRequest {
	s.Number = &v
	return s
}

func (s *ListAIAgentPhoneNumberRequest) SetPageNumber(v int64) *ListAIAgentPhoneNumberRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAIAgentPhoneNumberRequest) SetPageSize(v int64) *ListAIAgentPhoneNumberRequest {
	s.PageSize = &v
	return s
}

func (s *ListAIAgentPhoneNumberRequest) SetStatus(v int32) *ListAIAgentPhoneNumberRequest {
	s.Status = &v
	return s
}

func (s *ListAIAgentPhoneNumberRequest) Validate() error {
	return dara.Validate(s)
}
