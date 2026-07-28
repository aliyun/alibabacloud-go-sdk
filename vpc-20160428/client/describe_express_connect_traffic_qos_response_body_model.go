// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectTrafficQosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v string) *DescribeExpressConnectTrafficQosResponseBody
	GetCount() *string
	SetMaxResults(v int32) *DescribeExpressConnectTrafficQosResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeExpressConnectTrafficQosResponseBody
	GetNextToken() *string
	SetQosList(v []*DescribeExpressConnectTrafficQosResponseBodyQosList) *DescribeExpressConnectTrafficQosResponseBody
	GetQosList() []*DescribeExpressConnectTrafficQosResponseBodyQosList
	SetRequestId(v string) *DescribeExpressConnectTrafficQosResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeExpressConnectTrafficQosResponseBody
	GetTotalCount() *int32
}

type DescribeExpressConnectTrafficQosResponseBody struct {
	// The number of entries on the current page.
	//
	// example:
	//
	// 1
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The number of entries per page for paginated queries. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Valid values:
	//
	// - Leave this parameter empty for the first query or if no subsequent query is required.
	//
	// - If a next query is to be sent, set the value to the **NextToken*	- value returned in the previous API call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of QoS policies.
	QosList []*DescribeExpressConnectTrafficQosResponseBodyQosList `json:"QosList,omitempty" xml:"QosList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4CF20CC7-D1FC-425B-A15B-DF7C8E2131A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeExpressConnectTrafficQosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectTrafficQosResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectTrafficQosResponseBody) GetCount() *string {
	return s.Count
}

func (s *DescribeExpressConnectTrafficQosResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeExpressConnectTrafficQosResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeExpressConnectTrafficQosResponseBody) GetQosList() []*DescribeExpressConnectTrafficQosResponseBodyQosList {
	return s.QosList
}

func (s *DescribeExpressConnectTrafficQosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExpressConnectTrafficQosResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeExpressConnectTrafficQosResponseBody) SetCount(v string) *DescribeExpressConnectTrafficQosResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBody) SetMaxResults(v int32) *DescribeExpressConnectTrafficQosResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBody) SetNextToken(v string) *DescribeExpressConnectTrafficQosResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBody) SetQosList(v []*DescribeExpressConnectTrafficQosResponseBodyQosList) *DescribeExpressConnectTrafficQosResponseBody {
	s.QosList = v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBody) SetRequestId(v string) *DescribeExpressConnectTrafficQosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBody) SetTotalCount(v int32) *DescribeExpressConnectTrafficQosResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBody) Validate() error {
	if s.QosList != nil {
		for _, item := range s.QosList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeExpressConnectTrafficQosResponseBodyQosList struct {
	// The list of associated instances.
	AssociatedInstanceList []*DescribeExpressConnectTrafficQosResponseBodyQosListAssociatedInstanceList `json:"AssociatedInstanceList,omitempty" xml:"AssociatedInstanceList,omitempty" type:"Repeated"`
	// The overall configuration progress of the QoS policy. Valid values: **0*	- to **100**.
	//
	// example:
	//
	// 100
	Progressing *int32 `json:"Progressing,omitempty" xml:"Progressing,omitempty"`
	// The description of the QoS policy.
	//
	// The description is **0*	- to **256*	- characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// qos-test
	QosDescription *string `json:"QosDescription,omitempty" xml:"QosDescription,omitempty"`
	// The QoS policy ID.
	//
	// example:
	//
	// qos-pksbqfmotl5hzq****
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The name of the QoS policy.
	//
	// The name is **0*	- to **128*	- characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// qos-test
	QosName *string `json:"QosName,omitempty" xml:"QosName,omitempty"`
	// The list of QoS queues.
	QueueList []*DescribeExpressConnectTrafficQosResponseBodyQosListQueueList `json:"QueueList,omitempty" xml:"QueueList,omitempty" type:"Repeated"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmz7vtyl4f***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the QoS policy. Valid values:
	//
	// - **Normal**: available.
	//
	// - **Configuring**: being configured.
	//
	// > A QoS policy in the Configuring state restricts most create, update, and delete operations on QoS policies, QoS queues, and QoS rules.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the resource.
	Tags []*DescribeExpressConnectTrafficQosResponseBodyQosListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeExpressConnectTrafficQosResponseBodyQosList) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectTrafficQosResponseBodyQosList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosList) GetAssociatedInstanceList() []*DescribeExpressConnectTrafficQosResponseBodyQosListAssociatedInstanceList {
	return s.AssociatedInstanceList
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosList) GetProgressing() *int32 {
	return s.Progressing
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosList) GetQosDescription() *string {
	return s.QosDescription
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosList) GetQosId() *string {
	return s.QosId
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosList) GetQosName() *string {
	return s.QosName
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosList) GetQueueList() []*DescribeExpressConnectTrafficQosResponseBodyQosListQueueList {
	return s.QueueList
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosList) GetStatus() *string {
	return s.Status
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosList) GetTags() []*DescribeExpressConnectTrafficQosResponseBodyQosListTags {
	return s.Tags
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosList) SetAssociatedInstanceList(v []*DescribeExpressConnectTrafficQosResponseBodyQosListAssociatedInstanceList) *DescribeExpressConnectTrafficQosResponseBodyQosList {
	s.AssociatedInstanceList = v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosList) SetProgressing(v int32) *DescribeExpressConnectTrafficQosResponseBodyQosList {
	s.Progressing = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosList) SetQosDescription(v string) *DescribeExpressConnectTrafficQosResponseBodyQosList {
	s.QosDescription = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosList) SetQosId(v string) *DescribeExpressConnectTrafficQosResponseBodyQosList {
	s.QosId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosList) SetQosName(v string) *DescribeExpressConnectTrafficQosResponseBodyQosList {
	s.QosName = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosList) SetQueueList(v []*DescribeExpressConnectTrafficQosResponseBodyQosListQueueList) *DescribeExpressConnectTrafficQosResponseBodyQosList {
	s.QueueList = v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosList) SetResourceGroupId(v string) *DescribeExpressConnectTrafficQosResponseBodyQosList {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosList) SetStatus(v string) *DescribeExpressConnectTrafficQosResponseBodyQosList {
	s.Status = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosList) SetTags(v []*DescribeExpressConnectTrafficQosResponseBodyQosListTags) *DescribeExpressConnectTrafficQosResponseBodyQosList {
	s.Tags = v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosList) Validate() error {
	if s.AssociatedInstanceList != nil {
		for _, item := range s.AssociatedInstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.QueueList != nil {
		for _, item := range s.QueueList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeExpressConnectTrafficQosResponseBodyQosListAssociatedInstanceList struct {
	// The ID of the associated instance.
	//
	// example:
	//
	// pc-bp159zj8zujwy3p07****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The configuration progress of the associated instance. Valid values: **0*	- to **100**.
	//
	// example:
	//
	// 100
	InstanceProgressing *int32 `json:"InstanceProgressing,omitempty" xml:"InstanceProgressing,omitempty"`
	// The status of the associated instance. Valid values:
	//
	// - **Normal**: available.
	//
	// - **Configuring**: being configured.
	//
	// - **Deleting**: being deleted.
	//
	// example:
	//
	// Normal
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The type of the associated instance. Valid values: **PHYSICALCONNECTION**: Express Connect circuit.
	//
	// example:
	//
	// PHYSICALCONNECTION
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeExpressConnectTrafficQosResponseBodyQosListAssociatedInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectTrafficQosResponseBodyQosListAssociatedInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListAssociatedInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListAssociatedInstanceList) GetInstanceProgressing() *int32 {
	return s.InstanceProgressing
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListAssociatedInstanceList) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListAssociatedInstanceList) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListAssociatedInstanceList) SetInstanceId(v string) *DescribeExpressConnectTrafficQosResponseBodyQosListAssociatedInstanceList {
	s.InstanceId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListAssociatedInstanceList) SetInstanceProgressing(v int32) *DescribeExpressConnectTrafficQosResponseBodyQosListAssociatedInstanceList {
	s.InstanceProgressing = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListAssociatedInstanceList) SetInstanceStatus(v string) *DescribeExpressConnectTrafficQosResponseBodyQosListAssociatedInstanceList {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListAssociatedInstanceList) SetInstanceType(v string) *DescribeExpressConnectTrafficQosResponseBodyQosListAssociatedInstanceList {
	s.InstanceType = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListAssociatedInstanceList) Validate() error {
	return dara.Validate(s)
}

type DescribeExpressConnectTrafficQosResponseBodyQosListQueueList struct {
	// The bandwidth percentage of the QoS queue.
	//
	// 	- When the QoS queue type is **Medium**, this parameter is required. Valid values: **1*	- to **100**.
	//
	// 	- When the QoS queue type is **Default**, this parameter is set to "-".
	//
	// example:
	//
	// 100
	BandwidthPercent *string `json:"BandwidthPercent,omitempty" xml:"BandwidthPercent,omitempty"`
	// The QoS policy ID.
	//
	// example:
	//
	// qos-pksbqfmotl5hzq****
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The description of the QoS queue.
	//
	// The description is **0*	- to **256*	- characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// qos-queue-test
	QueueDescription *string `json:"QueueDescription,omitempty" xml:"QueueDescription,omitempty"`
	// The QoS queue ID.
	//
	// example:
	//
	// qos-queue-9nyx2u7n71s2rc****
	QueueId *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	// The name of the QoS queue.
	//
	// The name is **0*	- to **128*	- characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// qos-queue-test
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The QoS queue type. Valid values:
	//
	// - **High**: high-priority queue.
	//
	// - **Medium**: medium-priority queue.
	//
	// - **Default**: default-priority queue.
	//
	// > The default-priority queue cannot be created.
	//
	// example:
	//
	// High
	QueueType *string `json:"QueueType,omitempty" xml:"QueueType,omitempty"`
	// The status of the QoS queue. Valid values:
	//
	// - **Normal**: available.
	//
	// - **Configuring**: being configured.
	//
	// - **Deleting**: being deleted.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeExpressConnectTrafficQosResponseBodyQosListQueueList) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectTrafficQosResponseBodyQosListQueueList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListQueueList) GetBandwidthPercent() *string {
	return s.BandwidthPercent
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListQueueList) GetQosId() *string {
	return s.QosId
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListQueueList) GetQueueDescription() *string {
	return s.QueueDescription
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListQueueList) GetQueueId() *string {
	return s.QueueId
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListQueueList) GetQueueName() *string {
	return s.QueueName
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListQueueList) GetQueueType() *string {
	return s.QueueType
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListQueueList) GetStatus() *string {
	return s.Status
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListQueueList) SetBandwidthPercent(v string) *DescribeExpressConnectTrafficQosResponseBodyQosListQueueList {
	s.BandwidthPercent = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListQueueList) SetQosId(v string) *DescribeExpressConnectTrafficQosResponseBodyQosListQueueList {
	s.QosId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListQueueList) SetQueueDescription(v string) *DescribeExpressConnectTrafficQosResponseBodyQosListQueueList {
	s.QueueDescription = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListQueueList) SetQueueId(v string) *DescribeExpressConnectTrafficQosResponseBodyQosListQueueList {
	s.QueueId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListQueueList) SetQueueName(v string) *DescribeExpressConnectTrafficQosResponseBodyQosListQueueList {
	s.QueueName = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListQueueList) SetQueueType(v string) *DescribeExpressConnectTrafficQosResponseBodyQosListQueueList {
	s.QueueType = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListQueueList) SetStatus(v string) *DescribeExpressConnectTrafficQosResponseBodyQosListQueueList {
	s.Status = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListQueueList) Validate() error {
	return dara.Validate(s)
}

type DescribeExpressConnectTrafficQosResponseBodyQosListTags struct {
	// The tag key of the resource.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeExpressConnectTrafficQosResponseBodyQosListTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectTrafficQosResponseBodyQosListTags) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListTags) GetKey() *string {
	return s.Key
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListTags) GetValue() *string {
	return s.Value
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListTags) SetKey(v string) *DescribeExpressConnectTrafficQosResponseBodyQosListTags {
	s.Key = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListTags) SetValue(v string) *DescribeExpressConnectTrafficQosResponseBodyQosListTags {
	s.Value = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponseBodyQosListTags) Validate() error {
	return dara.Validate(s)
}
