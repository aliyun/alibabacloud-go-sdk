// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdviceServiceEnabledResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DescribeAdviceServiceEnabledResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAdviceServiceEnabledResponseBody
	GetRequestId() *string
	SetResult(v bool) *DescribeAdviceServiceEnabledResponseBody
	GetResult() *bool
}

type DescribeAdviceServiceEnabledResponseBody struct {
	// The returned message. Valid values:
	//
	// 	- If the request was successful, a **Success*	- message is returned.
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E031AABF-BD56-5966-A063-4283EF18DB45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the suggestion feature is enabled. Valid values:
	//
	// 	- **True**
	//
	// 	- **False**
	//
	// example:
	//
	// False
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeAdviceServiceEnabledResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdviceServiceEnabledResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdviceServiceEnabledResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAdviceServiceEnabledResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAdviceServiceEnabledResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DescribeAdviceServiceEnabledResponseBody) SetMessage(v string) *DescribeAdviceServiceEnabledResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAdviceServiceEnabledResponseBody) SetRequestId(v string) *DescribeAdviceServiceEnabledResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdviceServiceEnabledResponseBody) SetResult(v bool) *DescribeAdviceServiceEnabledResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeAdviceServiceEnabledResponseBody) Validate() error {
	return dara.Validate(s)
}
