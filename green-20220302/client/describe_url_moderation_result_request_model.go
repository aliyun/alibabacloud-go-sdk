// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUrlModerationResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReqId(v string) *DescribeUrlModerationResultRequest
	GetReqId() *string
}

type DescribeUrlModerationResultRequest struct {
	// The ReqId field returned by an asynchronous URL moderation operation.
	//
	// example:
	//
	// B0963D30-BAB4-562F-9ED0-7A23AEC51C7C
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
}

func (s DescribeUrlModerationResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUrlModerationResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeUrlModerationResultRequest) GetReqId() *string {
	return s.ReqId
}

func (s *DescribeUrlModerationResultRequest) SetReqId(v string) *DescribeUrlModerationResultRequest {
	s.ReqId = &v
	return s
}

func (s *DescribeUrlModerationResultRequest) Validate() error {
	return dara.Validate(s)
}
