// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBatchSlsDispatchStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeBatchSlsDispatchStatusRequest
	GetLang() *string
	SetPageNo(v int32) *DescribeBatchSlsDispatchStatusRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeBatchSlsDispatchStatusRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeBatchSlsDispatchStatusRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *DescribeBatchSlsDispatchStatusRequest
	GetSourceIp() *string
}

type DescribeBatchSlsDispatchStatusRequest struct {
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeBatchSlsDispatchStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchSlsDispatchStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeBatchSlsDispatchStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeBatchSlsDispatchStatusRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeBatchSlsDispatchStatusRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBatchSlsDispatchStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeBatchSlsDispatchStatusRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeBatchSlsDispatchStatusRequest) SetLang(v string) *DescribeBatchSlsDispatchStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusRequest) SetPageNo(v int32) *DescribeBatchSlsDispatchStatusRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusRequest) SetPageSize(v int32) *DescribeBatchSlsDispatchStatusRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusRequest) SetResourceGroupId(v string) *DescribeBatchSlsDispatchStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusRequest) SetSourceIp(v string) *DescribeBatchSlsDispatchStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusRequest) Validate() error {
	return dara.Validate(s)
}
