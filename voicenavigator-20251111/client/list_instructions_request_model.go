// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstructionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListInstructionsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListInstructionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstructionsRequest
	GetPageSize() *int32
	SetProviderId(v string) *ListInstructionsRequest
	GetProviderId() *string
}

type ListInstructionsRequest struct {
	// example:
	//
	// 36e9a4cd-a821-4f78-86fa-a9a4aefeea2e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// ********
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
}

func (s ListInstructionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstructionsRequest) GoString() string {
	return s.String()
}

func (s *ListInstructionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstructionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstructionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstructionsRequest) GetProviderId() *string {
	return s.ProviderId
}

func (s *ListInstructionsRequest) SetInstanceId(v string) *ListInstructionsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstructionsRequest) SetPageNumber(v int32) *ListInstructionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstructionsRequest) SetPageSize(v int32) *ListInstructionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstructionsRequest) SetProviderId(v string) *ListInstructionsRequest {
	s.ProviderId = &v
	return s
}

func (s *ListInstructionsRequest) Validate() error {
	return dara.Validate(s)
}
