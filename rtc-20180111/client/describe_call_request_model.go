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
	// The app ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// testappid
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The channel ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The time when the channel was created. You can query data of the last 30 days. The time is a UNIX timestamp in seconds.
	//
	// > The channel creation time must be exactly the same as the actual creation time (accurate to the second). You can call the [DescribeCallList](https://help.aliyun.com/document_detail/2655884.html) operation to view the channel creation time (the CreatedTS parameter in the response).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// The time when the channel was released. The time is a UNIX timestamp in seconds. If this parameter is left empty, the current time is used.
	//
	// > If the specified channel release time is later than the actual release time, the data from the creation time to the actual release time is returned.
	//
	// example:
	//
	// 1614936817
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// The extended data type to query. Valid values:
	//
	// - USER_DURATION_STAT: user duration statistics.
	//
	// example:
	//
	// USER_DURATION_STAT
	ExtDataType *string `json:"ExtDataType,omitempty" xml:"ExtDataType,omitempty"`
	// Specifies whether to query the communication experience information. Default value: true.
	//
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
