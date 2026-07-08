// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogDeliveryConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryConfig(v *DescribeLogDeliveryConfigResponseBodyDeliveryConfig) *DescribeLogDeliveryConfigResponseBody
	GetDeliveryConfig() *DescribeLogDeliveryConfigResponseBodyDeliveryConfig
	SetRequestId(v string) *DescribeLogDeliveryConfigResponseBody
	GetRequestId() *string
}

type DescribeLogDeliveryConfigResponseBody struct {
	// The log delivery configuration.
	DeliveryConfig *DescribeLogDeliveryConfigResponseBodyDeliveryConfig `json:"DeliveryConfig,omitempty" xml:"DeliveryConfig,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 66A98669-CC6E-4F3E-80A6-******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLogDeliveryConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogDeliveryConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogDeliveryConfigResponseBody) GetDeliveryConfig() *DescribeLogDeliveryConfigResponseBodyDeliveryConfig {
	return s.DeliveryConfig
}

func (s *DescribeLogDeliveryConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogDeliveryConfigResponseBody) SetDeliveryConfig(v *DescribeLogDeliveryConfigResponseBodyDeliveryConfig) *DescribeLogDeliveryConfigResponseBody {
	s.DeliveryConfig = v
	return s
}

func (s *DescribeLogDeliveryConfigResponseBody) SetRequestId(v string) *DescribeLogDeliveryConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogDeliveryConfigResponseBody) Validate() error {
	if s.DeliveryConfig != nil {
		if err := s.DeliveryConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLogDeliveryConfigResponseBodyDeliveryConfig struct {
	// The details of the log delivery configuration, in JSON format.
	//
	// > This parameter is the same as the **DeliveryDetail*	- parameter of the **CreateLogDeliveryConfig*	- operation. For more information, see [CreateLogDeliveryConfig](~~CreateLogDeliveryConfig~~).
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
	// The name of the log delivery configuration.
	//
	// example:
	//
	// test1
	DeliveryName *string `json:"DeliveryName,omitempty" xml:"DeliveryName,omitempty"`
	// The type of the log delivery configuration. Valid values:
	//
	// - **syslog**: The logs are delivered to a syslog service.
	//
	// - **kafka**: The logs are delivered to a Kafka service.
	//
	// example:
	//
	// syslog
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
}

func (s DescribeLogDeliveryConfigResponseBodyDeliveryConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogDeliveryConfigResponseBodyDeliveryConfig) GoString() string {
	return s.String()
}

func (s *DescribeLogDeliveryConfigResponseBodyDeliveryConfig) GetDeliveryDetail() *string {
	return s.DeliveryDetail
}

func (s *DescribeLogDeliveryConfigResponseBodyDeliveryConfig) GetDeliveryName() *string {
	return s.DeliveryName
}

func (s *DescribeLogDeliveryConfigResponseBodyDeliveryConfig) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *DescribeLogDeliveryConfigResponseBodyDeliveryConfig) SetDeliveryDetail(v string) *DescribeLogDeliveryConfigResponseBodyDeliveryConfig {
	s.DeliveryDetail = &v
	return s
}

func (s *DescribeLogDeliveryConfigResponseBodyDeliveryConfig) SetDeliveryName(v string) *DescribeLogDeliveryConfigResponseBodyDeliveryConfig {
	s.DeliveryName = &v
	return s
}

func (s *DescribeLogDeliveryConfigResponseBodyDeliveryConfig) SetDeliveryType(v string) *DescribeLogDeliveryConfigResponseBodyDeliveryConfig {
	s.DeliveryType = &v
	return s
}

func (s *DescribeLogDeliveryConfigResponseBodyDeliveryConfig) Validate() error {
	return dara.Validate(s)
}
