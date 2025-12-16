// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowQueryStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSlowQueryStatusResponseBody
	GetRequestId() *string
	SetResult(v *DescribeSlowQueryStatusResponseBodyResult) *DescribeSlowQueryStatusResponseBody
	GetResult() *DescribeSlowQueryStatusResponseBodyResult
}

type DescribeSlowQueryStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5C1C1C45-C64A-AD30-565F-140871D57E5E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The return result.
	Result *DescribeSlowQueryStatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeSlowQueryStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowQueryStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowQueryStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlowQueryStatusResponseBody) GetResult() *DescribeSlowQueryStatusResponseBodyResult {
	return s.Result
}

func (s *DescribeSlowQueryStatusResponseBody) SetRequestId(v string) *DescribeSlowQueryStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowQueryStatusResponseBody) SetResult(v *DescribeSlowQueryStatusResponseBodyResult) *DescribeSlowQueryStatusResponseBody {
	s.Result = v
	return s
}

func (s *DescribeSlowQueryStatusResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSlowQueryStatusResponseBodyResult struct {
	// The ID of the application.
	//
	// example:
	//
	// 100298370
	AppGroupId *string `json:"appGroupId,omitempty" xml:"appGroupId,omitempty"`
	// The network type of the slow query optimization service. Valid values:
	//
	// 	- outer: the Internet
	//
	// 	- internal: the internal network
	//
	// example:
	//
	// internal
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The status of the slow query optimization service. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// 	- n/a
	//
	// example:
	//
	// disabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeSlowQueryStatusResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowQueryStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeSlowQueryStatusResponseBodyResult) GetAppGroupId() *string {
	return s.AppGroupId
}

func (s *DescribeSlowQueryStatusResponseBodyResult) GetRegion() *string {
	return s.Region
}

func (s *DescribeSlowQueryStatusResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *DescribeSlowQueryStatusResponseBodyResult) SetAppGroupId(v string) *DescribeSlowQueryStatusResponseBodyResult {
	s.AppGroupId = &v
	return s
}

func (s *DescribeSlowQueryStatusResponseBodyResult) SetRegion(v string) *DescribeSlowQueryStatusResponseBodyResult {
	s.Region = &v
	return s
}

func (s *DescribeSlowQueryStatusResponseBodyResult) SetStatus(v string) *DescribeSlowQueryStatusResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeSlowQueryStatusResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
