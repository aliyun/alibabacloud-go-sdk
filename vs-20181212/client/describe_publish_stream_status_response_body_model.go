// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePublishStreamStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribePublishStreamStatusResponseBody
	GetCode() *int64
	SetMessage(v int64) *DescribePublishStreamStatusResponseBody
	GetMessage() *int64
	SetRequestId(v string) *DescribePublishStreamStatusResponseBody
	GetRequestId() *string
}

type DescribePublishStreamStatusResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *int64  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePublishStreamStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePublishStreamStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePublishStreamStatusResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribePublishStreamStatusResponseBody) GetMessage() *int64 {
	return s.Message
}

func (s *DescribePublishStreamStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePublishStreamStatusResponseBody) SetCode(v int64) *DescribePublishStreamStatusResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePublishStreamStatusResponseBody) SetMessage(v int64) *DescribePublishStreamStatusResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePublishStreamStatusResponseBody) SetRequestId(v string) *DescribePublishStreamStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePublishStreamStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
