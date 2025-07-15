// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlashSmsApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListFlashSmsApplicationsRequest
	GetInstanceId() *string
	SetName(v string) *ListFlashSmsApplicationsRequest
	GetName() *string
	SetPageNumber(v int32) *ListFlashSmsApplicationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFlashSmsApplicationsRequest
	GetPageSize() *int32
	SetProviderId(v string) *ListFlashSmsApplicationsRequest
	GetProviderId() *string
}

type ListFlashSmsApplicationsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// Uincall
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
}

func (s ListFlashSmsApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListFlashSmsApplicationsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFlashSmsApplicationsRequest) GetName() *string {
	return s.Name
}

func (s *ListFlashSmsApplicationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFlashSmsApplicationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFlashSmsApplicationsRequest) GetProviderId() *string {
	return s.ProviderId
}

func (s *ListFlashSmsApplicationsRequest) SetInstanceId(v string) *ListFlashSmsApplicationsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFlashSmsApplicationsRequest) SetName(v string) *ListFlashSmsApplicationsRequest {
	s.Name = &v
	return s
}

func (s *ListFlashSmsApplicationsRequest) SetPageNumber(v int32) *ListFlashSmsApplicationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlashSmsApplicationsRequest) SetPageSize(v int32) *ListFlashSmsApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlashSmsApplicationsRequest) SetProviderId(v string) *ListFlashSmsApplicationsRequest {
	s.ProviderId = &v
	return s
}

func (s *ListFlashSmsApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
