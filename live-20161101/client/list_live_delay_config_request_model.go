// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveDelayConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *ListLiveDelayConfigRequest
	GetDomain() *string
	SetOwnerId(v int64) *ListLiveDelayConfigRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *ListLiveDelayConfigRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListLiveDelayConfigRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListLiveDelayConfigRequest
	GetRegionId() *string
}

type ListLiveDelayConfigRequest struct {
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Valid values: 5 to 30. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListLiveDelayConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLiveDelayConfigRequest) GoString() string {
	return s.String()
}

func (s *ListLiveDelayConfigRequest) GetDomain() *string {
	return s.Domain
}

func (s *ListLiveDelayConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListLiveDelayConfigRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListLiveDelayConfigRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLiveDelayConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLiveDelayConfigRequest) SetDomain(v string) *ListLiveDelayConfigRequest {
	s.Domain = &v
	return s
}

func (s *ListLiveDelayConfigRequest) SetOwnerId(v int64) *ListLiveDelayConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ListLiveDelayConfigRequest) SetPageNum(v int32) *ListLiveDelayConfigRequest {
	s.PageNum = &v
	return s
}

func (s *ListLiveDelayConfigRequest) SetPageSize(v int32) *ListLiveDelayConfigRequest {
	s.PageSize = &v
	return s
}

func (s *ListLiveDelayConfigRequest) SetRegionId(v string) *ListLiveDelayConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ListLiveDelayConfigRequest) Validate() error {
	return dara.Validate(s)
}
