// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveSnapshotConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveSnapshotConfigRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveSnapshotConfigRequest
	GetDomainName() *string
	SetOrder(v string) *DescribeLiveSnapshotConfigRequest
	GetOrder() *string
	SetOwnerId(v int64) *DescribeLiveSnapshotConfigRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *DescribeLiveSnapshotConfigRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeLiveSnapshotConfigRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeLiveSnapshotConfigRequest
	GetSecurityToken() *string
}

type DescribeLiveSnapshotConfigRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The sort order. Valid values:
	//
	// 	- **asc*	- (default): ascending order
	//
	// 	- **desc**: descending order
	//
	// example:
	//
	// asc
	Order   *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Valid values: **5 to 30**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeLiveSnapshotConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveSnapshotConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveSnapshotConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveSnapshotConfigRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeLiveSnapshotConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveSnapshotConfigRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeLiveSnapshotConfigRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveSnapshotConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeLiveSnapshotConfigRequest) SetAppName(v string) *DescribeLiveSnapshotConfigRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveSnapshotConfigRequest) SetDomainName(v string) *DescribeLiveSnapshotConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveSnapshotConfigRequest) SetOrder(v string) *DescribeLiveSnapshotConfigRequest {
	s.Order = &v
	return s
}

func (s *DescribeLiveSnapshotConfigRequest) SetOwnerId(v int64) *DescribeLiveSnapshotConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveSnapshotConfigRequest) SetPageNum(v int32) *DescribeLiveSnapshotConfigRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveSnapshotConfigRequest) SetPageSize(v int32) *DescribeLiveSnapshotConfigRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveSnapshotConfigRequest) SetSecurityToken(v string) *DescribeLiveSnapshotConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveSnapshotConfigRequest) Validate() error {
	return dara.Validate(s)
}
