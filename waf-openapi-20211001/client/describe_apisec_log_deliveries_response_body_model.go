// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecLogDeliveriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryConfigs(v []*DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs) *DescribeApisecLogDeliveriesResponseBody
	GetDeliveryConfigs() []*DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs
	SetRequestId(v string) *DescribeApisecLogDeliveriesResponseBody
	GetRequestId() *string
}

type DescribeApisecLogDeliveriesResponseBody struct {
	// The configurations of API security log subscription.
	DeliveryConfigs []*DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs `json:"DeliveryConfigs,omitempty" xml:"DeliveryConfigs,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 2EFCFE18-78F8-5079-B312-07***48B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApisecLogDeliveriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecLogDeliveriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisecLogDeliveriesResponseBody) GetDeliveryConfigs() []*DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs {
	return s.DeliveryConfigs
}

func (s *DescribeApisecLogDeliveriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisecLogDeliveriesResponseBody) SetDeliveryConfigs(v []*DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs) *DescribeApisecLogDeliveriesResponseBody {
	s.DeliveryConfigs = v
	return s
}

func (s *DescribeApisecLogDeliveriesResponseBody) SetRequestId(v string) *DescribeApisecLogDeliveriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisecLogDeliveriesResponseBody) Validate() error {
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

type DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs struct {
	// The type of the log subscription. Valid values:
	//
	// 	- **risk**: risk information.
	//
	// 	- **event**: attack event information.
	//
	// 	- **asset**: asset information.
	//
	// example:
	//
	// risk
	AssertKey *string `json:"AssertKey,omitempty" xml:"AssertKey,omitempty"`
	// The ID of the region where logs are stored.
	//
	// example:
	//
	// cn-hangzhou
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// The name of the Logstore in Simple Log Service.
	//
	// example:
	//
	// apisec-logstore***
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The name of the project in Simple Log Service.
	//
	// example:
	//
	// apisec-project-14316572********
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The status of API security log subscription. Valid values:
	//
	// 	- **true**: enabled.
	//
	// 	- **false**: disabled.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs) GoString() string {
	return s.String()
}

func (s *DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs) GetAssertKey() *string {
	return s.AssertKey
}

func (s *DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs) GetStatus() *bool {
	return s.Status
}

func (s *DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs) SetAssertKey(v string) *DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs {
	s.AssertKey = &v
	return s
}

func (s *DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs) SetLogRegionId(v string) *DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs {
	s.LogRegionId = &v
	return s
}

func (s *DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs) SetLogStoreName(v string) *DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs {
	s.LogStoreName = &v
	return s
}

func (s *DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs) SetProjectName(v string) *DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs {
	s.ProjectName = &v
	return s
}

func (s *DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs) SetStatus(v bool) *DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs {
	s.Status = &v
	return s
}

func (s *DescribeApisecLogDeliveriesResponseBodyDeliveryConfigs) Validate() error {
	return dara.Validate(s)
}
