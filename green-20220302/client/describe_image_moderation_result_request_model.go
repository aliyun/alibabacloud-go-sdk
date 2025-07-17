// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageModerationResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReqId(v string) *DescribeImageModerationResultRequest
	GetReqId() *string
}

type DescribeImageModerationResultRequest struct {
	// The ReqId field returned by the asynchronous Image Moderation 2.0 API.
	//
	// example:
	//
	// B0963D30-BAB4-562F-9ED0-7A23AEC51C7C
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
}

func (s DescribeImageModerationResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageModerationResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageModerationResultRequest) GetReqId() *string {
	return s.ReqId
}

func (s *DescribeImageModerationResultRequest) SetReqId(v string) *DescribeImageModerationResultRequest {
	s.ReqId = &v
	return s
}

func (s *DescribeImageModerationResultRequest) Validate() error {
	return dara.Validate(s)
}
