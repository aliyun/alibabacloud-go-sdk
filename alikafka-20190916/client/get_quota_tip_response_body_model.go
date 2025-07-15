// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaTipResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetQuotaTipResponseBody
	GetCode() *int32
	SetMessage(v string) *GetQuotaTipResponseBody
	GetMessage() *string
	SetQuotaData(v *GetQuotaTipResponseBodyQuotaData) *GetQuotaTipResponseBody
	GetQuotaData() *GetQuotaTipResponseBodyQuotaData
	SetRequestId(v string) *GetQuotaTipResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQuotaTipResponseBody
	GetSuccess() *bool
}

type GetQuotaTipResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional message. This message is typically used to describe API call failures for troubleshooting.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The quota.
	QuotaData *GetQuotaTipResponseBodyQuotaData `json:"QuotaData,omitempty" xml:"QuotaData,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0178A3A7-E87B-5E50-A16F-3E62F534****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQuotaTipResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaTipResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaTipResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetQuotaTipResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQuotaTipResponseBody) GetQuotaData() *GetQuotaTipResponseBodyQuotaData {
	return s.QuotaData
}

func (s *GetQuotaTipResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQuotaTipResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQuotaTipResponseBody) SetCode(v int32) *GetQuotaTipResponseBody {
	s.Code = &v
	return s
}

func (s *GetQuotaTipResponseBody) SetMessage(v string) *GetQuotaTipResponseBody {
	s.Message = &v
	return s
}

func (s *GetQuotaTipResponseBody) SetQuotaData(v *GetQuotaTipResponseBodyQuotaData) *GetQuotaTipResponseBody {
	s.QuotaData = v
	return s
}

func (s *GetQuotaTipResponseBody) SetRequestId(v string) *GetQuotaTipResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuotaTipResponseBody) SetSuccess(v bool) *GetQuotaTipResponseBody {
	s.Success = &v
	return s
}

func (s *GetQuotaTipResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetQuotaTipResponseBodyQuotaData struct {
	// The number of available groups.
	//
	// example:
	//
	// 50
	GroupLeft *int32 `json:"GroupLeft,omitempty" xml:"GroupLeft,omitempty"`
	// The number of used groups.
	//
	// example:
	//
	// 50
	GroupUsed *int32 `json:"GroupUsed,omitempty" xml:"GroupUsed,omitempty"`
	// The method that you use to purchase partitions. Valid values:
	//
	// 	- 0: indicates that the instance is purchased based on topics.
	//
	// 	- 1: indicates that the instance is purchased based on partitions.
	//
	// example:
	//
	// 1
	IsPartitionBuy *int32 `json:"IsPartitionBuy,omitempty" xml:"IsPartitionBuy,omitempty"`
	// The number of available partitions.
	//
	// example:
	//
	// 1050
	PartitionLeft *int32 `json:"PartitionLeft,omitempty" xml:"PartitionLeft,omitempty"`
	// The number of purchased partitions.
	//
	// example:
	//
	// 100
	PartitionNumOfBuy *int32 `json:"PartitionNumOfBuy,omitempty" xml:"PartitionNumOfBuy,omitempty"`
	// The quota of partitions.
	//
	// example:
	//
	// 1100
	PartitionQuota *int32 `json:"PartitionQuota,omitempty" xml:"PartitionQuota,omitempty"`
	// The number of used partitions.
	//
	// example:
	//
	// 50
	PartitionUsed *int32 `json:"PartitionUsed,omitempty" xml:"PartitionUsed,omitempty"`
	// The number of available topics.
	//
	// example:
	//
	// 20
	TopicLeft *int32 `json:"TopicLeft,omitempty" xml:"TopicLeft,omitempty"`
	// The number of purchased topics.
	//
	// example:
	//
	// 50
	TopicNumOfBuy *int32 `json:"TopicNumOfBuy,omitempty" xml:"TopicNumOfBuy,omitempty"`
	// The quota of topics.
	//
	// example:
	//
	// 50
	TopicQuota *int32 `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
	// The number of used topics.
	//
	// example:
	//
	// 30
	TopicUsed *int32 `json:"TopicUsed,omitempty" xml:"TopicUsed,omitempty"`
}

func (s GetQuotaTipResponseBodyQuotaData) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaTipResponseBodyQuotaData) GoString() string {
	return s.String()
}

func (s *GetQuotaTipResponseBodyQuotaData) GetGroupLeft() *int32 {
	return s.GroupLeft
}

func (s *GetQuotaTipResponseBodyQuotaData) GetGroupUsed() *int32 {
	return s.GroupUsed
}

func (s *GetQuotaTipResponseBodyQuotaData) GetIsPartitionBuy() *int32 {
	return s.IsPartitionBuy
}

func (s *GetQuotaTipResponseBodyQuotaData) GetPartitionLeft() *int32 {
	return s.PartitionLeft
}

func (s *GetQuotaTipResponseBodyQuotaData) GetPartitionNumOfBuy() *int32 {
	return s.PartitionNumOfBuy
}

func (s *GetQuotaTipResponseBodyQuotaData) GetPartitionQuota() *int32 {
	return s.PartitionQuota
}

func (s *GetQuotaTipResponseBodyQuotaData) GetPartitionUsed() *int32 {
	return s.PartitionUsed
}

func (s *GetQuotaTipResponseBodyQuotaData) GetTopicLeft() *int32 {
	return s.TopicLeft
}

func (s *GetQuotaTipResponseBodyQuotaData) GetTopicNumOfBuy() *int32 {
	return s.TopicNumOfBuy
}

func (s *GetQuotaTipResponseBodyQuotaData) GetTopicQuota() *int32 {
	return s.TopicQuota
}

func (s *GetQuotaTipResponseBodyQuotaData) GetTopicUsed() *int32 {
	return s.TopicUsed
}

func (s *GetQuotaTipResponseBodyQuotaData) SetGroupLeft(v int32) *GetQuotaTipResponseBodyQuotaData {
	s.GroupLeft = &v
	return s
}

func (s *GetQuotaTipResponseBodyQuotaData) SetGroupUsed(v int32) *GetQuotaTipResponseBodyQuotaData {
	s.GroupUsed = &v
	return s
}

func (s *GetQuotaTipResponseBodyQuotaData) SetIsPartitionBuy(v int32) *GetQuotaTipResponseBodyQuotaData {
	s.IsPartitionBuy = &v
	return s
}

func (s *GetQuotaTipResponseBodyQuotaData) SetPartitionLeft(v int32) *GetQuotaTipResponseBodyQuotaData {
	s.PartitionLeft = &v
	return s
}

func (s *GetQuotaTipResponseBodyQuotaData) SetPartitionNumOfBuy(v int32) *GetQuotaTipResponseBodyQuotaData {
	s.PartitionNumOfBuy = &v
	return s
}

func (s *GetQuotaTipResponseBodyQuotaData) SetPartitionQuota(v int32) *GetQuotaTipResponseBodyQuotaData {
	s.PartitionQuota = &v
	return s
}

func (s *GetQuotaTipResponseBodyQuotaData) SetPartitionUsed(v int32) *GetQuotaTipResponseBodyQuotaData {
	s.PartitionUsed = &v
	return s
}

func (s *GetQuotaTipResponseBodyQuotaData) SetTopicLeft(v int32) *GetQuotaTipResponseBodyQuotaData {
	s.TopicLeft = &v
	return s
}

func (s *GetQuotaTipResponseBodyQuotaData) SetTopicNumOfBuy(v int32) *GetQuotaTipResponseBodyQuotaData {
	s.TopicNumOfBuy = &v
	return s
}

func (s *GetQuotaTipResponseBodyQuotaData) SetTopicQuota(v int32) *GetQuotaTipResponseBodyQuotaData {
	s.TopicQuota = &v
	return s
}

func (s *GetQuotaTipResponseBodyQuotaData) SetTopicUsed(v int32) *GetQuotaTipResponseBodyQuotaData {
	s.TopicUsed = &v
	return s
}

func (s *GetQuotaTipResponseBodyQuotaData) Validate() error {
	return dara.Validate(s)
}
