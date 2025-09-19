// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteListStrategyUuidCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeWhiteListStrategyUuidCountResponseBody
	GetRequestId() *string
	SetUuidCount(v int32) *DescribeWhiteListStrategyUuidCountResponseBody
	GetUuidCount() *int32
}

type DescribeWhiteListStrategyUuidCountResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of the servers on which the application whitelist policy takes effect.
	//
	// example:
	//
	// 9
	UuidCount *int32 `json:"UuidCount,omitempty" xml:"UuidCount,omitempty"`
}

func (s DescribeWhiteListStrategyUuidCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListStrategyUuidCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListStrategyUuidCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWhiteListStrategyUuidCountResponseBody) GetUuidCount() *int32 {
	return s.UuidCount
}

func (s *DescribeWhiteListStrategyUuidCountResponseBody) SetRequestId(v string) *DescribeWhiteListStrategyUuidCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWhiteListStrategyUuidCountResponseBody) SetUuidCount(v int32) *DescribeWhiteListStrategyUuidCountResponseBody {
	s.UuidCount = &v
	return s
}

func (s *DescribeWhiteListStrategyUuidCountResponseBody) Validate() error {
	return dara.Validate(s)
}
