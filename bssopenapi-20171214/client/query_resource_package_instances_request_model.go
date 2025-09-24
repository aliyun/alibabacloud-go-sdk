// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryResourcePackageInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpiryTimeEnd(v string) *QueryResourcePackageInstancesRequest
	GetExpiryTimeEnd() *string
	SetExpiryTimeStart(v string) *QueryResourcePackageInstancesRequest
	GetExpiryTimeStart() *string
	SetIncludePartner(v bool) *QueryResourcePackageInstancesRequest
	GetIncludePartner() *bool
	SetOwnerId(v int64) *QueryResourcePackageInstancesRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *QueryResourcePackageInstancesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryResourcePackageInstancesRequest
	GetPageSize() *int32
	SetProductCode(v string) *QueryResourcePackageInstancesRequest
	GetProductCode() *string
}

type QueryResourcePackageInstancesRequest struct {
	// The end of the expiration time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2020-02-02T13:00:00Z
	ExpiryTimeEnd *string `json:"ExpiryTimeEnd,omitempty" xml:"ExpiryTimeEnd,omitempty"`
	// The beginning of the expiration time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2020-02-02T12:00:00Z
	ExpiryTimeStart *string `json:"ExpiryTimeStart,omitempty" xml:"ExpiryTimeStart,omitempty"`
	// Specifies whether partners are involved.
	//
	// example:
	//
	// true
	IncludePartner *bool  `json:"IncludePartner,omitempty" xml:"IncludePartner,omitempty"`
	OwnerId        *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page. Default value: 20. Maximum value: 300.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The code of the service.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s QueryResourcePackageInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryResourcePackageInstancesRequest) GoString() string {
	return s.String()
}

func (s *QueryResourcePackageInstancesRequest) GetExpiryTimeEnd() *string {
	return s.ExpiryTimeEnd
}

func (s *QueryResourcePackageInstancesRequest) GetExpiryTimeStart() *string {
	return s.ExpiryTimeStart
}

func (s *QueryResourcePackageInstancesRequest) GetIncludePartner() *bool {
	return s.IncludePartner
}

func (s *QueryResourcePackageInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryResourcePackageInstancesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryResourcePackageInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryResourcePackageInstancesRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *QueryResourcePackageInstancesRequest) SetExpiryTimeEnd(v string) *QueryResourcePackageInstancesRequest {
	s.ExpiryTimeEnd = &v
	return s
}

func (s *QueryResourcePackageInstancesRequest) SetExpiryTimeStart(v string) *QueryResourcePackageInstancesRequest {
	s.ExpiryTimeStart = &v
	return s
}

func (s *QueryResourcePackageInstancesRequest) SetIncludePartner(v bool) *QueryResourcePackageInstancesRequest {
	s.IncludePartner = &v
	return s
}

func (s *QueryResourcePackageInstancesRequest) SetOwnerId(v int64) *QueryResourcePackageInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryResourcePackageInstancesRequest) SetPageNum(v int32) *QueryResourcePackageInstancesRequest {
	s.PageNum = &v
	return s
}

func (s *QueryResourcePackageInstancesRequest) SetPageSize(v int32) *QueryResourcePackageInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *QueryResourcePackageInstancesRequest) SetProductCode(v string) *QueryResourcePackageInstancesRequest {
	s.ProductCode = &v
	return s
}

func (s *QueryResourcePackageInstancesRequest) Validate() error {
	return dara.Validate(s)
}
