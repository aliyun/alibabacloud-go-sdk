// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudProtectableCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProtectableCount(v int32) *DescribeHybridCloudProtectableCountResponseBody
	GetProtectableCount() *int32
	SetRequestId(v string) *DescribeHybridCloudProtectableCountResponseBody
	GetRequestId() *string
}

type DescribeHybridCloudProtectableCountResponseBody struct {
	// The number of protection nodes that can be added to the hybrid cloud cluster.
	//
	// example:
	//
	// 1
	ProtectableCount *int32 `json:"ProtectableCount,omitempty" xml:"ProtectableCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6087EA47-C10F-5A0A-A405-DB5B241**B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHybridCloudProtectableCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudProtectableCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudProtectableCountResponseBody) GetProtectableCount() *int32 {
	return s.ProtectableCount
}

func (s *DescribeHybridCloudProtectableCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridCloudProtectableCountResponseBody) SetProtectableCount(v int32) *DescribeHybridCloudProtectableCountResponseBody {
	s.ProtectableCount = &v
	return s
}

func (s *DescribeHybridCloudProtectableCountResponseBody) SetRequestId(v string) *DescribeHybridCloudProtectableCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridCloudProtectableCountResponseBody) Validate() error {
	return dara.Validate(s)
}
