// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceLogDeliveryStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogConfigs(v []*DescribeResourceLogDeliveryStatusResponseBodyLogConfigs) *DescribeResourceLogDeliveryStatusResponseBody
	GetLogConfigs() []*DescribeResourceLogDeliveryStatusResponseBodyLogConfigs
	SetRequestId(v string) *DescribeResourceLogDeliveryStatusResponseBody
	GetRequestId() *string
}

type DescribeResourceLogDeliveryStatusResponseBody struct {
	LogConfigs []*DescribeResourceLogDeliveryStatusResponseBodyLogConfigs `json:"LogConfigs,omitempty" xml:"LogConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19****5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeResourceLogDeliveryStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceLogDeliveryStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogDeliveryStatusResponseBody) GetLogConfigs() []*DescribeResourceLogDeliveryStatusResponseBodyLogConfigs {
	return s.LogConfigs
}

func (s *DescribeResourceLogDeliveryStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceLogDeliveryStatusResponseBody) SetLogConfigs(v []*DescribeResourceLogDeliveryStatusResponseBodyLogConfigs) *DescribeResourceLogDeliveryStatusResponseBody {
	s.LogConfigs = v
	return s
}

func (s *DescribeResourceLogDeliveryStatusResponseBody) SetRequestId(v string) *DescribeResourceLogDeliveryStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceLogDeliveryStatusResponseBody) Validate() error {
	if s.LogConfigs != nil {
		for _, item := range s.LogConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResourceLogDeliveryStatusResponseBodyLogConfigs struct {
	// example:
	//
	// export-kafka
	DeliveryName *string `json:"DeliveryName,omitempty" xml:"DeliveryName,omitempty"`
	// example:
	//
	// kafka
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// example:
	//
	// test.waf.com-waf
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeResourceLogDeliveryStatusResponseBodyLogConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceLogDeliveryStatusResponseBodyLogConfigs) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogDeliveryStatusResponseBodyLogConfigs) GetDeliveryName() *string {
	return s.DeliveryName
}

func (s *DescribeResourceLogDeliveryStatusResponseBodyLogConfigs) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *DescribeResourceLogDeliveryStatusResponseBodyLogConfigs) GetResource() *string {
	return s.Resource
}

func (s *DescribeResourceLogDeliveryStatusResponseBodyLogConfigs) GetStatus() *bool {
	return s.Status
}

func (s *DescribeResourceLogDeliveryStatusResponseBodyLogConfigs) SetDeliveryName(v string) *DescribeResourceLogDeliveryStatusResponseBodyLogConfigs {
	s.DeliveryName = &v
	return s
}

func (s *DescribeResourceLogDeliveryStatusResponseBodyLogConfigs) SetDeliveryType(v string) *DescribeResourceLogDeliveryStatusResponseBodyLogConfigs {
	s.DeliveryType = &v
	return s
}

func (s *DescribeResourceLogDeliveryStatusResponseBodyLogConfigs) SetResource(v string) *DescribeResourceLogDeliveryStatusResponseBodyLogConfigs {
	s.Resource = &v
	return s
}

func (s *DescribeResourceLogDeliveryStatusResponseBodyLogConfigs) SetStatus(v bool) *DescribeResourceLogDeliveryStatusResponseBodyLogConfigs {
	s.Status = &v
	return s
}

func (s *DescribeResourceLogDeliveryStatusResponseBodyLogConfigs) Validate() error {
	return dara.Validate(s)
}
