// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlashSmsAccessProfilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListFlashSmsAccessProfilesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListFlashSmsAccessProfilesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFlashSmsAccessProfilesRequest
	GetPageSize() *int32
}

type ListFlashSmsAccessProfilesRequest struct {
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFlashSmsAccessProfilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsAccessProfilesRequest) GoString() string {
	return s.String()
}

func (s *ListFlashSmsAccessProfilesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFlashSmsAccessProfilesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFlashSmsAccessProfilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFlashSmsAccessProfilesRequest) SetInstanceId(v string) *ListFlashSmsAccessProfilesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFlashSmsAccessProfilesRequest) SetPageNumber(v int32) *ListFlashSmsAccessProfilesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlashSmsAccessProfilesRequest) SetPageSize(v int32) *ListFlashSmsAccessProfilesRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlashSmsAccessProfilesRequest) Validate() error {
	return dara.Validate(s)
}
