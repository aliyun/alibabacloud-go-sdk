// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeCallRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeCallRequest
	GetChannelId() *string
	SetCreatedTs(v int64) *DescribeCallRequest
	GetCreatedTs() *int64
	SetDestroyedTs(v int64) *DescribeCallRequest
	GetDestroyedTs() *int64
	SetExtDataType(v string) *DescribeCallRequest
	GetExtDataType() *string
	SetQueryExpInfo(v bool) *DescribeCallRequest
	GetQueryExpInfo() *bool
}

type DescribeCallRequest struct {
	// APP ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// testappid
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1614936817
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// example:
	//
	// USER_DURATION_STAT
	ExtDataType *string `json:"ExtDataType,omitempty" xml:"ExtDataType,omitempty"`
	// example:
	//
	// false
	QueryExpInfo *bool `json:"QueryExpInfo,omitempty" xml:"QueryExpInfo,omitempty"`
}

func (s DescribeCallRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCallRequest) GoString() string {
	return s.String()
}

func (s *DescribeCallRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeCallRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeCallRequest) GetCreatedTs() *int64 {
	return s.CreatedTs
}

func (s *DescribeCallRequest) GetDestroyedTs() *int64 {
	return s.DestroyedTs
}

func (s *DescribeCallRequest) GetExtDataType() *string {
	return s.ExtDataType
}

func (s *DescribeCallRequest) GetQueryExpInfo() *bool {
	return s.QueryExpInfo
}

func (s *DescribeCallRequest) SetAppId(v string) *DescribeCallRequest {
	s.AppId = &v
	return s
}

func (s *DescribeCallRequest) SetChannelId(v string) *DescribeCallRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeCallRequest) SetCreatedTs(v int64) *DescribeCallRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeCallRequest) SetDestroyedTs(v int64) *DescribeCallRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeCallRequest) SetExtDataType(v string) *DescribeCallRequest {
	s.ExtDataType = &v
	return s
}

func (s *DescribeCallRequest) SetQueryExpInfo(v bool) *DescribeCallRequest {
	s.QueryExpInfo = &v
	return s
}

func (s *DescribeCallRequest) Validate() error {
	return dara.Validate(s)
}
