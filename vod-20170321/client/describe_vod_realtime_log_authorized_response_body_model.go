// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodRealtimeLogAuthorizedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizedStatus(v string) *DescribeVodRealtimeLogAuthorizedResponseBody
	GetAuthorizedStatus() *string
	SetRequestId(v string) *DescribeVodRealtimeLogAuthorizedResponseBody
	GetRequestId() *string
}

type DescribeVodRealtimeLogAuthorizedResponseBody struct {
	AuthorizedStatus *string `json:"AuthorizedStatus,omitempty" xml:"AuthorizedStatus,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodRealtimeLogAuthorizedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodRealtimeLogAuthorizedResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodRealtimeLogAuthorizedResponseBody) GetAuthorizedStatus() *string {
	return s.AuthorizedStatus
}

func (s *DescribeVodRealtimeLogAuthorizedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodRealtimeLogAuthorizedResponseBody) SetAuthorizedStatus(v string) *DescribeVodRealtimeLogAuthorizedResponseBody {
	s.AuthorizedStatus = &v
	return s
}

func (s *DescribeVodRealtimeLogAuthorizedResponseBody) SetRequestId(v string) *DescribeVodRealtimeLogAuthorizedResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodRealtimeLogAuthorizedResponseBody) Validate() error {
	return dara.Validate(s)
}
