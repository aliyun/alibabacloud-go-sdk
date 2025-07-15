// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveRecordConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveRecordConfigRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveRecordConfigRequest
	GetDomainName() *string
	SetOrder(v string) *DescribeLiveRecordConfigRequest
	GetOrder() *string
	SetOwnerId(v int64) *DescribeLiveRecordConfigRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *DescribeLiveRecordConfigRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeLiveRecordConfigRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeLiveRecordConfigRequest
	GetSecurityToken() *string
	SetStreamName(v string) *DescribeLiveRecordConfigRequest
	GetStreamName() *string
}

type DescribeLiveRecordConfigRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The name of the main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The order in which the entries are sorted based on creation time. Valid values:
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
	// 5
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveRecordConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveRecordConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveRecordConfigRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeLiveRecordConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveRecordConfigRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeLiveRecordConfigRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveRecordConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeLiveRecordConfigRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveRecordConfigRequest) SetAppName(v string) *DescribeLiveRecordConfigRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveRecordConfigRequest) SetDomainName(v string) *DescribeLiveRecordConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveRecordConfigRequest) SetOrder(v string) *DescribeLiveRecordConfigRequest {
	s.Order = &v
	return s
}

func (s *DescribeLiveRecordConfigRequest) SetOwnerId(v int64) *DescribeLiveRecordConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveRecordConfigRequest) SetPageNum(v int32) *DescribeLiveRecordConfigRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveRecordConfigRequest) SetPageSize(v int32) *DescribeLiveRecordConfigRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveRecordConfigRequest) SetSecurityToken(v string) *DescribeLiveRecordConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveRecordConfigRequest) SetStreamName(v string) *DescribeLiveRecordConfigRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveRecordConfigRequest) Validate() error {
	return dara.Validate(s)
}
