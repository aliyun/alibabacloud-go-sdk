// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeliveryChannelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannels(v []*ListDeliveryChannelsResponseBodyDeliveryChannels) *ListDeliveryChannelsResponseBody
	GetDeliveryChannels() []*ListDeliveryChannelsResponseBodyDeliveryChannels
	SetMaxResults(v int32) *ListDeliveryChannelsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDeliveryChannelsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDeliveryChannelsResponseBody
	GetRequestId() *string
}

type ListDeliveryChannelsResponseBody struct {
	// The delivery channels.
	DeliveryChannels []*ListDeliveryChannelsResponseBodyDeliveryChannels `json:"DeliveryChannels,omitempty" xml:"DeliveryChannels,omitempty" type:"Repeated"`
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

func (s ListDeliveryChannelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeliveryChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeliveryChannelsResponseBody) GetDeliveryChannels() []*ListDeliveryChannelsResponseBodyDeliveryChannels {
	return s.DeliveryChannels
}

func (s *ListDeliveryChannelsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDeliveryChannelsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDeliveryChannelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeliveryChannelsResponseBody) SetDeliveryChannels(v []*ListDeliveryChannelsResponseBodyDeliveryChannels) *ListDeliveryChannelsResponseBody {
	s.DeliveryChannels = v
	return s
}

func (s *ListDeliveryChannelsResponseBody) SetMaxResults(v int32) *ListDeliveryChannelsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDeliveryChannelsResponseBody) SetNextToken(v string) *ListDeliveryChannelsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDeliveryChannelsResponseBody) SetRequestId(v string) *ListDeliveryChannelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeliveryChannelsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDeliveryChannelsResponseBodyDeliveryChannels struct {
	// The time when the delivery channel was created.
	//
	// example:
	//
	// 2024-06-20T08:46:29Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the delivery channel.
	DeliveryChannelDescription *string `json:"DeliveryChannelDescription,omitempty" xml:"DeliveryChannelDescription,omitempty"`
	// The ID of the delivery channel.
	//
	// example:
	//
	// dc-4m6ffpkgu***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The name of the delivery channel.
	//
	// example:
	//
	// test-delivery-channel
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
}

func (s ListDeliveryChannelsResponseBodyDeliveryChannels) String() string {
	return dara.Prettify(s)
}

func (s ListDeliveryChannelsResponseBodyDeliveryChannels) GoString() string {
	return s.String()
}

func (s *ListDeliveryChannelsResponseBodyDeliveryChannels) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDeliveryChannelsResponseBodyDeliveryChannels) GetDeliveryChannelDescription() *string {
	return s.DeliveryChannelDescription
}

func (s *ListDeliveryChannelsResponseBodyDeliveryChannels) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *ListDeliveryChannelsResponseBodyDeliveryChannels) GetDeliveryChannelName() *string {
	return s.DeliveryChannelName
}

func (s *ListDeliveryChannelsResponseBodyDeliveryChannels) SetCreateTime(v string) *ListDeliveryChannelsResponseBodyDeliveryChannels {
	s.CreateTime = &v
	return s
}

func (s *ListDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelDescription(v string) *ListDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelDescription = &v
	return s
}

func (s *ListDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelId(v string) *ListDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelId = &v
	return s
}

func (s *ListDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelName(v string) *ListDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelName = &v
	return s
}

func (s *ListDeliveryChannelsResponseBodyDeliveryChannels) Validate() error {
	return dara.Validate(s)
}
