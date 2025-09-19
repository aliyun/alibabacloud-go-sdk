// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageInfoListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUuids(v string) *DescribeImageInfoListRequest
	GetUuids() *string
}

type DescribeImageInfoListRequest struct {
	// The UUID of the server. Separate multiple UUIDs with commas (,).
	//
	// > You can call the [DescribeCloudCenterInstances](https://help.aliyun.com/document_detail/141932.html) operation to query the UUIDs of servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// 391abd09184cbd3743d7f5ec125d****,
	//
	// e6aeb2a5b6004479398b0bcd1160****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s DescribeImageInfoListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInfoListRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageInfoListRequest) GetUuids() *string {
	return s.Uuids
}

func (s *DescribeImageInfoListRequest) SetUuids(v string) *DescribeImageInfoListRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeImageInfoListRequest) Validate() error {
	return dara.Validate(s)
}
