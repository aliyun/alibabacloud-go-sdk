// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveSnapshotDetectPornConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveSnapshotDetectPornConfigRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveSnapshotDetectPornConfigRequest
	GetDomainName() *string
	SetOrder(v string) *DescribeLiveSnapshotDetectPornConfigRequest
	GetOrder() *string
	SetOwnerId(v int64) *DescribeLiveSnapshotDetectPornConfigRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *DescribeLiveSnapshotDetectPornConfigRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeLiveSnapshotDetectPornConfigRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeLiveSnapshotDetectPornConfigRequest
	GetSecurityToken() *string
}

type DescribeLiveSnapshotDetectPornConfigRequest struct {
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
	// The order in which the entries are sorted based on creation time. Valid values:
	//
	// 	- **Asc*	- (default): ascending order
	//
	// 	- **Desc**: descending order
	//
	// Enumerated values:
	//
	// 	- asc
	//
	// 	- desc
	//
	// example:
	//
	// Asc
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

func (s DescribeLiveSnapshotDetectPornConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveSnapshotDetectPornConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotDetectPornConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveSnapshotDetectPornConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveSnapshotDetectPornConfigRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeLiveSnapshotDetectPornConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveSnapshotDetectPornConfigRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeLiveSnapshotDetectPornConfigRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveSnapshotDetectPornConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeLiveSnapshotDetectPornConfigRequest) SetAppName(v string) *DescribeLiveSnapshotDetectPornConfigRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigRequest) SetDomainName(v string) *DescribeLiveSnapshotDetectPornConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigRequest) SetOrder(v string) *DescribeLiveSnapshotDetectPornConfigRequest {
	s.Order = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigRequest) SetOwnerId(v int64) *DescribeLiveSnapshotDetectPornConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigRequest) SetPageNum(v int32) *DescribeLiveSnapshotDetectPornConfigRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigRequest) SetPageSize(v int32) *DescribeLiveSnapshotDetectPornConfigRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigRequest) SetSecurityToken(v string) *DescribeLiveSnapshotDetectPornConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigRequest) Validate() error {
	return dara.Validate(s)
}
