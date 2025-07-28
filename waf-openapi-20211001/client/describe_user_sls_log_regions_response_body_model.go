// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserSlsLogRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogRegions(v []*string) *DescribeUserSlsLogRegionsResponseBody
	GetLogRegions() []*string
	SetRequestId(v string) *DescribeUserSlsLogRegionsResponseBody
	GetRequestId() *string
}

type DescribeUserSlsLogRegionsResponseBody struct {
	// The region IDs.
	LogRegions []*string `json:"LogRegions,omitempty" xml:"LogRegions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9D11AC3A-A10C-56E7-A342-E87EC892****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserSlsLogRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserSlsLogRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserSlsLogRegionsResponseBody) GetLogRegions() []*string {
	return s.LogRegions
}

func (s *DescribeUserSlsLogRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserSlsLogRegionsResponseBody) SetLogRegions(v []*string) *DescribeUserSlsLogRegionsResponseBody {
	s.LogRegions = v
	return s
}

func (s *DescribeUserSlsLogRegionsResponseBody) SetRequestId(v string) *DescribeUserSlsLogRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserSlsLogRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}
