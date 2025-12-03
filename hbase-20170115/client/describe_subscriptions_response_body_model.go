// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSubscriptionsResponseBody
	GetRequestId() *string
	SetSubscriptions(v []*DescribeSubscriptionsResponseBodySubscriptions) *DescribeSubscriptionsResponseBody
	GetSubscriptions() []*DescribeSubscriptionsResponseBodySubscriptions
}

type DescribeSubscriptionsResponseBody struct {
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Subscriptions []*DescribeSubscriptionsResponseBodySubscriptions `json:"Subscriptions,omitempty" xml:"Subscriptions,omitempty" type:"Repeated"`
}

func (s DescribeSubscriptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSubscriptionsResponseBody) GetSubscriptions() []*DescribeSubscriptionsResponseBodySubscriptions {
	return s.Subscriptions
}

func (s *DescribeSubscriptionsResponseBody) SetRequestId(v string) *DescribeSubscriptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubscriptionsResponseBody) SetSubscriptions(v []*DescribeSubscriptionsResponseBodySubscriptions) *DescribeSubscriptionsResponseBody {
	s.Subscriptions = v
	return s
}

func (s *DescribeSubscriptionsResponseBody) Validate() error {
	if s.Subscriptions != nil {
		for _, item := range s.Subscriptions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSubscriptionsResponseBodySubscriptions struct {
	DBInstances             []*DescribeSubscriptionsResponseBodySubscriptionsDBInstances `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
	Mapping                 *string                                                      `json:"Mapping,omitempty" xml:"Mapping,omitempty"`
	SubscriptionDescription *string                                                      `json:"SubscriptionDescription,omitempty" xml:"SubscriptionDescription,omitempty"`
	SubscriptionId          *string                                                      `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
	SubscriptionStatus      *string                                                      `json:"SubscriptionStatus,omitempty" xml:"SubscriptionStatus,omitempty"`
	SubscriptionType        *string                                                      `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s DescribeSubscriptionsResponseBodySubscriptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionsResponseBodySubscriptions) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionsResponseBodySubscriptions) GetDBInstances() []*DescribeSubscriptionsResponseBodySubscriptionsDBInstances {
	return s.DBInstances
}

func (s *DescribeSubscriptionsResponseBodySubscriptions) GetMapping() *string {
	return s.Mapping
}

func (s *DescribeSubscriptionsResponseBodySubscriptions) GetSubscriptionDescription() *string {
	return s.SubscriptionDescription
}

func (s *DescribeSubscriptionsResponseBodySubscriptions) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *DescribeSubscriptionsResponseBodySubscriptions) GetSubscriptionStatus() *string {
	return s.SubscriptionStatus
}

func (s *DescribeSubscriptionsResponseBodySubscriptions) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *DescribeSubscriptionsResponseBodySubscriptions) SetDBInstances(v []*DescribeSubscriptionsResponseBodySubscriptionsDBInstances) *DescribeSubscriptionsResponseBodySubscriptions {
	s.DBInstances = v
	return s
}

func (s *DescribeSubscriptionsResponseBodySubscriptions) SetMapping(v string) *DescribeSubscriptionsResponseBodySubscriptions {
	s.Mapping = &v
	return s
}

func (s *DescribeSubscriptionsResponseBodySubscriptions) SetSubscriptionDescription(v string) *DescribeSubscriptionsResponseBodySubscriptions {
	s.SubscriptionDescription = &v
	return s
}

func (s *DescribeSubscriptionsResponseBodySubscriptions) SetSubscriptionId(v string) *DescribeSubscriptionsResponseBodySubscriptions {
	s.SubscriptionId = &v
	return s
}

func (s *DescribeSubscriptionsResponseBodySubscriptions) SetSubscriptionStatus(v string) *DescribeSubscriptionsResponseBodySubscriptions {
	s.SubscriptionStatus = &v
	return s
}

func (s *DescribeSubscriptionsResponseBodySubscriptions) SetSubscriptionType(v string) *DescribeSubscriptionsResponseBodySubscriptions {
	s.SubscriptionType = &v
	return s
}

func (s *DescribeSubscriptionsResponseBodySubscriptions) Validate() error {
	if s.DBInstances != nil {
		for _, item := range s.DBInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSubscriptionsResponseBodySubscriptionsDBInstances struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Role         *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s DescribeSubscriptionsResponseBodySubscriptionsDBInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionsResponseBodySubscriptionsDBInstances) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionsResponseBodySubscriptionsDBInstances) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSubscriptionsResponseBodySubscriptionsDBInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSubscriptionsResponseBodySubscriptionsDBInstances) GetRole() *string {
	return s.Role
}

func (s *DescribeSubscriptionsResponseBodySubscriptionsDBInstances) SetDBInstanceId(v string) *DescribeSubscriptionsResponseBodySubscriptionsDBInstances {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSubscriptionsResponseBodySubscriptionsDBInstances) SetRegionId(v string) *DescribeSubscriptionsResponseBodySubscriptionsDBInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeSubscriptionsResponseBodySubscriptionsDBInstances) SetRole(v string) *DescribeSubscriptionsResponseBodySubscriptionsDBInstances {
	s.Role = &v
	return s
}

func (s *DescribeSubscriptionsResponseBodySubscriptionsDBInstances) Validate() error {
	return dara.Validate(s)
}
