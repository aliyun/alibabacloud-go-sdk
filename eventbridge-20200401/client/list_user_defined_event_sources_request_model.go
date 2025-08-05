// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserDefinedEventSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventBusName(v string) *ListUserDefinedEventSourcesRequest
	GetEventBusName() *string
	SetLimit(v int32) *ListUserDefinedEventSourcesRequest
	GetLimit() *int32
	SetNamePrefix(v string) *ListUserDefinedEventSourcesRequest
	GetNamePrefix() *string
	SetNextToken(v string) *ListUserDefinedEventSourcesRequest
	GetNextToken() *string
}

type ListUserDefinedEventSourcesRequest struct {
	// The name of the event bus.
	//
	// example:
	//
	// testBus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The maximum number of entries to be returned in a call. You can use this parameter and NextToken to implement paging. Note: Up to 100 entries can be returned in a call.
	//
	// example:
	//
	// 50
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The name of the event source.
	//
	// example:
	//
	// testName
	NamePrefix *string `json:"NamePrefix,omitempty" xml:"NamePrefix,omitempty"`
	// If you configure Limit and excess return values exist, this parameter is returned.
	//
	// example:
	//
	// 100
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListUserDefinedEventSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserDefinedEventSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *ListUserDefinedEventSourcesRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListUserDefinedEventSourcesRequest) GetNamePrefix() *string {
	return s.NamePrefix
}

func (s *ListUserDefinedEventSourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserDefinedEventSourcesRequest) SetEventBusName(v string) *ListUserDefinedEventSourcesRequest {
	s.EventBusName = &v
	return s
}

func (s *ListUserDefinedEventSourcesRequest) SetLimit(v int32) *ListUserDefinedEventSourcesRequest {
	s.Limit = &v
	return s
}

func (s *ListUserDefinedEventSourcesRequest) SetNamePrefix(v string) *ListUserDefinedEventSourcesRequest {
	s.NamePrefix = &v
	return s
}

func (s *ListUserDefinedEventSourcesRequest) SetNextToken(v string) *ListUserDefinedEventSourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserDefinedEventSourcesRequest) Validate() error {
	return dara.Validate(s)
}
