// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloneVoiceModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListCloneVoiceModelsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCloneVoiceModelsRequest
	GetPageSize() *int32
}

type ListCloneVoiceModelsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCloneVoiceModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloneVoiceModelsRequest) GoString() string {
	return s.String()
}

func (s *ListCloneVoiceModelsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCloneVoiceModelsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloneVoiceModelsRequest) SetPageNumber(v int32) *ListCloneVoiceModelsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCloneVoiceModelsRequest) SetPageSize(v int32) *ListCloneVoiceModelsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloneVoiceModelsRequest) Validate() error {
	return dara.Validate(s)
}
