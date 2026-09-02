// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUuidVulNumClassifyStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]*DataValue) *DescribeUuidVulNumClassifyStatisticResponseBody
	GetData() map[string]*DataValue
	SetRequestId(v string) *DescribeUuidVulNumClassifyStatisticResponseBody
	GetRequestId() *string
}

type DescribeUuidVulNumClassifyStatisticResponseBody struct {
	Data map[string]*DataValue `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// FDF7B8D9-8493-4B90-8D13-E0C1FFCE5F97
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUuidVulNumClassifyStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUuidVulNumClassifyStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUuidVulNumClassifyStatisticResponseBody) GetData() map[string]*DataValue {
	return s.Data
}

func (s *DescribeUuidVulNumClassifyStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUuidVulNumClassifyStatisticResponseBody) SetData(v map[string]*DataValue) *DescribeUuidVulNumClassifyStatisticResponseBody {
	s.Data = v
	return s
}

func (s *DescribeUuidVulNumClassifyStatisticResponseBody) SetRequestId(v string) *DescribeUuidVulNumClassifyStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUuidVulNumClassifyStatisticResponseBody) Validate() error {
	return dara.Validate(s)
}
