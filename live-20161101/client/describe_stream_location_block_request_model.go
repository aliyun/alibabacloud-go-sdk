// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamLocationBlockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeStreamLocationBlockRequest
	GetAppName() *string
	SetBlockType(v string) *DescribeStreamLocationBlockRequest
	GetBlockType() *string
	SetDomainName(v string) *DescribeStreamLocationBlockRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeStreamLocationBlockRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *DescribeStreamLocationBlockRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeStreamLocationBlockRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeStreamLocationBlockRequest
	GetRegionId() *string
	SetStreamName(v string) *DescribeStreamLocationBlockRequest
	GetStreamName() *string
}

type DescribeStreamLocationBlockRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The blocking type. Valid values:
	//
	// 	- blacklist
	//
	// 	- whitelist
	//
	// example:
	//
	// blacklist
	BlockType *string `json:"BlockType,omitempty" xml:"BlockType,omitempty"`
	// The streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Valid values: integers from 1 to 100.
	//
	// example:
	//
	// 5
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// stream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeStreamLocationBlockRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamLocationBlockRequest) GoString() string {
	return s.String()
}

func (s *DescribeStreamLocationBlockRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeStreamLocationBlockRequest) GetBlockType() *string {
	return s.BlockType
}

func (s *DescribeStreamLocationBlockRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeStreamLocationBlockRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeStreamLocationBlockRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeStreamLocationBlockRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeStreamLocationBlockRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeStreamLocationBlockRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeStreamLocationBlockRequest) SetAppName(v string) *DescribeStreamLocationBlockRequest {
	s.AppName = &v
	return s
}

func (s *DescribeStreamLocationBlockRequest) SetBlockType(v string) *DescribeStreamLocationBlockRequest {
	s.BlockType = &v
	return s
}

func (s *DescribeStreamLocationBlockRequest) SetDomainName(v string) *DescribeStreamLocationBlockRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeStreamLocationBlockRequest) SetOwnerId(v int64) *DescribeStreamLocationBlockRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeStreamLocationBlockRequest) SetPageNum(v int32) *DescribeStreamLocationBlockRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeStreamLocationBlockRequest) SetPageSize(v int32) *DescribeStreamLocationBlockRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeStreamLocationBlockRequest) SetRegionId(v string) *DescribeStreamLocationBlockRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeStreamLocationBlockRequest) SetStreamName(v string) *DescribeStreamLocationBlockRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeStreamLocationBlockRequest) Validate() error {
	return dara.Validate(s)
}
