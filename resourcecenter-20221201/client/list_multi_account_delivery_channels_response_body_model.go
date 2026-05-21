// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiAccountDeliveryChannelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannels(v []*ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels) *ListMultiAccountDeliveryChannelsResponseBody
	GetDeliveryChannels() []*ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels
	SetMaxResults(v int32) *ListMultiAccountDeliveryChannelsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListMultiAccountDeliveryChannelsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMultiAccountDeliveryChannelsResponseBody
	GetRequestId() *string
}

type ListMultiAccountDeliveryChannelsResponseBody struct {
	// The delivery channels.
	DeliveryChannels []*ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels `json:"DeliveryChannels,omitempty" xml:"DeliveryChannels,omitempty" type:"Repeated"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 17502A1B-7026-5551-8E1C-CCABD0CBC***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMultiAccountDeliveryChannelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMultiAccountDeliveryChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultiAccountDeliveryChannelsResponseBody) GetDeliveryChannels() []*ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels {
	return s.DeliveryChannels
}

func (s *ListMultiAccountDeliveryChannelsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMultiAccountDeliveryChannelsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMultiAccountDeliveryChannelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMultiAccountDeliveryChannelsResponseBody) SetDeliveryChannels(v []*ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels) *ListMultiAccountDeliveryChannelsResponseBody {
	s.DeliveryChannels = v
	return s
}

func (s *ListMultiAccountDeliveryChannelsResponseBody) SetMaxResults(v int32) *ListMultiAccountDeliveryChannelsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMultiAccountDeliveryChannelsResponseBody) SetNextToken(v string) *ListMultiAccountDeliveryChannelsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountDeliveryChannelsResponseBody) SetRequestId(v string) *ListMultiAccountDeliveryChannelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultiAccountDeliveryChannelsResponseBody) Validate() error {
	if s.DeliveryChannels != nil {
		for _, item := range s.DeliveryChannels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels struct {
	// The time when the delivery channel was created.
	//
	// example:
	//
	// 2023-08-17T00:23:55Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the delivery channel.
	DeliveryChannelDescription *string `json:"DeliveryChannelDescription,omitempty" xml:"DeliveryChannelDescription,omitempty"`
	// The ID of the delivery channel.
	//
	// example:
	//
	// dc-0bzhsqpnk***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The name of the delivery channel.
	//
	// example:
	//
	// test-multi-account-delivery-channel
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
}

func (s ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels) String() string {
	return dara.Prettify(s)
}

func (s ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels) GoString() string {
	return s.String()
}

func (s *ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels) GetDeliveryChannelDescription() *string {
	return s.DeliveryChannelDescription
}

func (s *ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels) GetDeliveryChannelName() *string {
	return s.DeliveryChannelName
}

func (s *ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels) SetCreateTime(v string) *ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels {
	s.CreateTime = &v
	return s
}

func (s *ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelDescription(v string) *ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelDescription = &v
	return s
}

func (s *ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelId(v string) *ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelId = &v
	return s
}

func (s *ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelName(v string) *ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelName = &v
	return s
}

func (s *ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels) Validate() error {
	return dara.Validate(s)
}
