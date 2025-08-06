// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodUserTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeVodUserTagsRequest
	GetOwnerId() *int64
}

type DescribeVodUserTagsRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVodUserTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodUserTagsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodUserTagsRequest) SetOwnerId(v int64) *DescribeVodUserTagsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodUserTagsRequest) Validate() error {
	return dara.Validate(s)
}
