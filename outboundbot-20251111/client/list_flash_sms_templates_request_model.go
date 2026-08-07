// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlashSmsTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListFlashSmsTemplatesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListFlashSmsTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFlashSmsTemplatesRequest
	GetPageSize() *int32
	SetProviderId(v string) *ListFlashSmsTemplatesRequest
	GetProviderId() *string
}

type ListFlashSmsTemplatesRequest struct {
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
	// 供应商ID
	//
	// example:
	//
	// Uincall
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
}

func (s ListFlashSmsTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListFlashSmsTemplatesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFlashSmsTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFlashSmsTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFlashSmsTemplatesRequest) GetProviderId() *string {
	return s.ProviderId
}

func (s *ListFlashSmsTemplatesRequest) SetInstanceId(v string) *ListFlashSmsTemplatesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFlashSmsTemplatesRequest) SetPageNumber(v int32) *ListFlashSmsTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlashSmsTemplatesRequest) SetPageSize(v int32) *ListFlashSmsTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlashSmsTemplatesRequest) SetProviderId(v string) *ListFlashSmsTemplatesRequest {
	s.ProviderId = &v
	return s
}

func (s *ListFlashSmsTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
