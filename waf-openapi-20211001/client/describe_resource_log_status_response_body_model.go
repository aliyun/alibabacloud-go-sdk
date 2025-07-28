// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceLogStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeResourceLogStatusResponseBody
	GetRequestId() *string
	SetResult(v []*DescribeResourceLogStatusResponseBodyResult) *DescribeResourceLogStatusResponseBody
	GetResult() []*DescribeResourceLogStatusResponseBodyResult
}

type DescribeResourceLogStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0DABF8AB-2321-5F8D-A8D7-922D757FBFFE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result []*DescribeResourceLogStatusResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeResourceLogStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceLogStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceLogStatusResponseBody) GetResult() []*DescribeResourceLogStatusResponseBodyResult {
	return s.Result
}

func (s *DescribeResourceLogStatusResponseBody) SetRequestId(v string) *DescribeResourceLogStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceLogStatusResponseBody) SetResult(v []*DescribeResourceLogStatusResponseBodyResult) *DescribeResourceLogStatusResponseBody {
	s.Result = v
	return s
}

func (s *DescribeResourceLogStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceLogStatusResponseBodyResult struct {
	// The protected object.
	//
	// example:
	//
	// alb-wewbb23dfsetetcic****
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// Indicates whether the log collection feature is enabled for the protected object. Valid values:
	//
	// 	- **true:*	- The log collection feature is enabled.
	//
	// 	- **false:*	- The log collection feature is disabled.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeResourceLogStatusResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceLogStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogStatusResponseBodyResult) GetResource() *string {
	return s.Resource
}

func (s *DescribeResourceLogStatusResponseBodyResult) GetStatus() *bool {
	return s.Status
}

func (s *DescribeResourceLogStatusResponseBodyResult) SetResource(v string) *DescribeResourceLogStatusResponseBodyResult {
	s.Resource = &v
	return s
}

func (s *DescribeResourceLogStatusResponseBodyResult) SetStatus(v bool) *DescribeResourceLogStatusResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeResourceLogStatusResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
