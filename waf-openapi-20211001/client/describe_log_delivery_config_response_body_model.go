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
	// The information about the log delivery configuration.
	DeliveryConfig *DescribeLogDeliveryConfigResponseBodyDeliveryConfig `json:"DeliveryConfig,omitempty" xml:"DeliveryConfig,omitempty" type:"Struct"`
	// The request ID.
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
