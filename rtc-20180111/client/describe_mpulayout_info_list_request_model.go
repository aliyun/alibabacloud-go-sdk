// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMPULayoutInfoListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeMPULayoutInfoListRequest
	GetAppId() *string
	SetLayoutId(v int64) *DescribeMPULayoutInfoListRequest
	GetLayoutId() *int64
	SetName(v string) *DescribeMPULayoutInfoListRequest
	GetName() *string
	SetOwnerId(v int64) *DescribeMPULayoutInfoListRequest
	GetOwnerId() *int64
	SetPageNum(v int64) *DescribeMPULayoutInfoListRequest
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeMPULayoutInfoListRequest
	GetPageSize() *int64
}

type DescribeMPULayoutInfoListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 2
	LayoutId *int64 `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// example:
	//
	// LayoutName
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeMPULayoutInfoListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMPULayoutInfoListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoListRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeMPULayoutInfoListRequest) GetLayoutId() *int64 {
	return s.LayoutId
}

func (s *DescribeMPULayoutInfoListRequest) GetName() *string {
	return s.Name
}

func (s *DescribeMPULayoutInfoListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeMPULayoutInfoListRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeMPULayoutInfoListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeMPULayoutInfoListRequest) SetAppId(v string) *DescribeMPULayoutInfoListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeMPULayoutInfoListRequest) SetLayoutId(v int64) *DescribeMPULayoutInfoListRequest {
	s.LayoutId = &v
	return s
}

func (s *DescribeMPULayoutInfoListRequest) SetName(v string) *DescribeMPULayoutInfoListRequest {
	s.Name = &v
	return s
}

func (s *DescribeMPULayoutInfoListRequest) SetOwnerId(v int64) *DescribeMPULayoutInfoListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMPULayoutInfoListRequest) SetPageNum(v int64) *DescribeMPULayoutInfoListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeMPULayoutInfoListRequest) SetPageSize(v int64) *DescribeMPULayoutInfoListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMPULayoutInfoListRequest) Validate() error {
	return dara.Validate(s)
}
