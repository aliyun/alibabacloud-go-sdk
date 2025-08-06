// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodRealtimeLogAuthorizedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeVodRealtimeLogAuthorizedRequest
	GetOwnerId() *int64
}

type DescribeVodRealtimeLogAuthorizedRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVodRealtimeLogAuthorizedRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodRealtimeLogAuthorizedRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodRealtimeLogAuthorizedRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodRealtimeLogAuthorizedRequest) SetOwnerId(v int64) *DescribeVodRealtimeLogAuthorizedRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodRealtimeLogAuthorizedRequest) Validate() error {
	return dara.Validate(s)
}
