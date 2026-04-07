// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloneVoiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *ListCloneVoiceRequest
	GetBusinessUnitId() *string
	SetPageNumber(v int32) *ListCloneVoiceRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCloneVoiceRequest
	GetPageSize() *int32
	SetStatus(v string) *ListCloneVoiceRequest
	GetStatus() *string
}

type ListCloneVoiceRequest struct {
	// example:
	//
	// llm-3pptowd2olrctsvc
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// Published
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCloneVoiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloneVoiceRequest) GoString() string {
	return s.String()
}

func (s *ListCloneVoiceRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *ListCloneVoiceRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCloneVoiceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloneVoiceRequest) GetStatus() *string {
	return s.Status
}

func (s *ListCloneVoiceRequest) SetBusinessUnitId(v string) *ListCloneVoiceRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *ListCloneVoiceRequest) SetPageNumber(v int32) *ListCloneVoiceRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCloneVoiceRequest) SetPageSize(v int32) *ListCloneVoiceRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloneVoiceRequest) SetStatus(v string) *ListCloneVoiceRequest {
	s.Status = &v
	return s
}

func (s *ListCloneVoiceRequest) Validate() error {
	return dara.Validate(s)
}
