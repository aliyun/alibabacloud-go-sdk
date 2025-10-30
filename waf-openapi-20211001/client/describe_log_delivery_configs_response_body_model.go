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
	// The information about the log delivery configuration.
	DeliveryConfigs []*DescribeLogDeliveryConfigsResponseBodyDeliveryConfigs `json:"DeliveryConfigs,omitempty" xml:"DeliveryConfigs,omitempty" type:"Repeated"`
	MaxResults      *int32                                                   `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string                                                  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2EFCFE18-78F8-5079-B312-07***48B
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The content of the log delivery configuration. The value is a JSON string that contains multiple parameters.
	//
	// >  This parameter is the same as the **DeliveryDetail*	- parameter of the **CreateLogDeliveryConfig*	- operation. For more information, see **Parameter description for log delivery configuration*	- of the [CreateLogDeliveryConfig](~~CreateLogDeliveryConfig~~) operation.
	DeliveryDetail *string `json:"DeliveryDetail,omitempty" xml:"DeliveryDetail,omitempty"`
	// The name of the log delivery configuration.
	//
	// example:
	//
	// test1
	DeliveryName *string `json:"DeliveryName,omitempty" xml:"DeliveryName,omitempty"`
	// The type of the log delivery configuration. Valid values:
	//
	// 	- **syslog**: Logs are delivered to a syslog service.
	//
	// 	- **kafka**: Logs are delivered to a Kafka service.
	//
	// example:
	//
	// kafka
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
