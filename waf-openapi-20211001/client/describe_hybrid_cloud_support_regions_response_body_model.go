// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudSupportRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeHybridCloudSupportRegionsResponseBody
	GetRequestId() *string
	SetSupportRegions(v []*string) *DescribeHybridCloudSupportRegionsResponseBody
	GetSupportRegions() []*string
}

type DescribeHybridCloudSupportRegionsResponseBody struct {
	// example:
	//
	// 256959D5-3B45-54CD-A66D-F75F11E8E754
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SupportRegions []*string `json:"SupportRegions,omitempty" xml:"SupportRegions,omitempty" type:"Repeated"`
}

func (s DescribeHybridCloudSupportRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudSupportRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudSupportRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridCloudSupportRegionsResponseBody) GetSupportRegions() []*string {
	return s.SupportRegions
}

func (s *DescribeHybridCloudSupportRegionsResponseBody) SetRequestId(v string) *DescribeHybridCloudSupportRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridCloudSupportRegionsResponseBody) SetSupportRegions(v []*string) *DescribeHybridCloudSupportRegionsResponseBody {
	s.SupportRegions = v
	return s
}

func (s *DescribeHybridCloudSupportRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}
