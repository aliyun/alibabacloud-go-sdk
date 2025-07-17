// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMultimodalModerationResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReqId(v string) *DescribeMultimodalModerationResultRequest
	GetReqId() *string
}

type DescribeMultimodalModerationResultRequest struct {
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
}

func (s DescribeMultimodalModerationResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultimodalModerationResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeMultimodalModerationResultRequest) GetReqId() *string {
	return s.ReqId
}

func (s *DescribeMultimodalModerationResultRequest) SetReqId(v string) *DescribeMultimodalModerationResultRequest {
	s.ReqId = &v
	return s
}

func (s *DescribeMultimodalModerationResultRequest) Validate() error {
	return dara.Validate(s)
}
