// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogDeliveryConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryConfigs(v []*DescribeLogDeliveryConfigsResponseBodyDeliveryConfigs) *DescribeLogDeliveryConfigsResponseBody
	GetDeliveryConfigs() []*DescribeLogDeliveryConfigsResponseBodyDeliveryConfigs
	SetMaxResults(v int32) *DescribeLogDeliveryConfigsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeLogDeliveryConfigsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeLogDeliveryConfigsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeLogDeliveryConfigsResponseBody
	GetTotalCount() *int32
}

type DescribeLogDeliveryConfigsResponseBody struct {
	// The delivery configuration information.
	DeliveryConfigs []*DescribeLogDeliveryConfigsResponseBodyDeliveryConfigs `json:"DeliveryConfigs,omitempty" xml:"DeliveryConfigs,omitempty" type:"Repeated"`
	// The number of entries to return per page during paging in a paged query. Valid values: 1 to 50. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for querying the next page. If a next page exists, this field returns a value.
	//
	// > If this parameter returns a value, it indicates that more pages exist. You can use the returned **NextToken*	- value as a request parameter to obtain the data on the next page. Repeat this process until no value is returned, which indicates that all data has been retrieved.
	//
	// example:
	//
	// AAAAAGBgV9tolsLfijC4wam2htS*****D/46H3X2wIS
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2EFCFE18-78F8-5079-B312-07***48B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLogDeliveryConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogDeliveryConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogDeliveryConfigsResponseBody) GetDeliveryConfigs() []*DescribeLogDeliveryConfigsResponseBodyDeliveryConfigs {
	return s.DeliveryConfigs
}

func (s *DescribeLogDeliveryConfigsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeLogDeliveryConfigsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeLogDeliveryConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogDeliveryConfigsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeLogDeliveryConfigsResponseBody) SetDeliveryConfigs(v []*DescribeLogDeliveryConfigsResponseBodyDeliveryConfigs) *DescribeLogDeliveryConfigsResponseBody {
	s.DeliveryConfigs = v
	return s
}

func (s *DescribeLogDeliveryConfigsResponseBody) SetMaxResults(v int32) *DescribeLogDeliveryConfigsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeLogDeliveryConfigsResponseBody) SetNextToken(v string) *DescribeLogDeliveryConfigsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeLogDeliveryConfigsResponseBody) SetRequestId(v string) *DescribeLogDeliveryConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogDeliveryConfigsResponseBody) SetTotalCount(v int32) *DescribeLogDeliveryConfigsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLogDeliveryConfigsResponseBody) Validate() error {
	if s.DeliveryConfigs != nil {
		for _, item := range s.DeliveryConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLogDeliveryConfigsResponseBodyDeliveryConfigs struct {
	// The content of the delivery configuration. The content is a JSON string constructed from a series of parameters.
	//
	// > This parameter is the same as the **DeliveryDetail*	- request parameter of the [CreateLogDeliveryConfig](~~CreateLogDeliveryConfig~~) operation. For more information, see the **Delivery configuration content parameters*	- section of the [CreateLogDeliveryConfig](~~CreateLogDeliveryConfig~~) operation.
	//
	// example:
	//
	// {
	//
	//   "rfcVersion": "rfc3164",
	//
	//   "protocol": "tcp",
	//
	//   "servers": [
	//
	//     {
	//
	//       "address": "1.1.1.1",
	//
	//       "port": 20
	//
	//     }
	//
	//   ]
	//
	// }
	DeliveryDetail *string `json:"DeliveryDetail,omitempty" xml:"DeliveryDetail,omitempty"`
	// The name of the delivery configuration.
	//
	// example:
	//
	// test1
	DeliveryName *string `json:"DeliveryName,omitempty" xml:"DeliveryName,omitempty"`
	// The type of the delivery configuration. Valid values:
	//
	// - **syslog**: Delivers logs to a Syslog service.
	//
	// - **kafka**: Delivers logs to a Kafka service.
	//
	// example:
	//
	// syslog
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
}

func (s DescribeLogDeliveryConfigsResponseBodyDeliveryConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogDeliveryConfigsResponseBodyDeliveryConfigs) GoString() string {
	return s.String()
}

func (s *DescribeLogDeliveryConfigsResponseBodyDeliveryConfigs) GetDeliveryDetail() *string {
	return s.DeliveryDetail
}

func (s *DescribeLogDeliveryConfigsResponseBodyDeliveryConfigs) GetDeliveryName() *string {
	return s.DeliveryName
}

func (s *DescribeLogDeliveryConfigsResponseBodyDeliveryConfigs) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *DescribeLogDeliveryConfigsResponseBodyDeliveryConfigs) SetDeliveryDetail(v string) *DescribeLogDeliveryConfigsResponseBodyDeliveryConfigs {
	s.DeliveryDetail = &v
	return s
}

func (s *DescribeLogDeliveryConfigsResponseBodyDeliveryConfigs) SetDeliveryName(v string) *DescribeLogDeliveryConfigsResponseBodyDeliveryConfigs {
	s.DeliveryName = &v
	return s
}

func (s *DescribeLogDeliveryConfigsResponseBodyDeliveryConfigs) SetDeliveryType(v string) *DescribeLogDeliveryConfigsResponseBodyDeliveryConfigs {
	s.DeliveryType = &v
	return s
}

func (s *DescribeLogDeliveryConfigsResponseBodyDeliveryConfigs) Validate() error {
	return dara.Validate(s)
}
