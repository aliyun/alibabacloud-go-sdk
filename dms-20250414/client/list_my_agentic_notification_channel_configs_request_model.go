// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyAgenticNotificationChannelConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v string) *ListMyAgenticNotificationChannelConfigsRequest
	GetChannel() *string
	SetMaxResults(v int32) *ListMyAgenticNotificationChannelConfigsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMyAgenticNotificationChannelConfigsRequest
	GetNextToken() *string
}

type ListMyAgenticNotificationChannelConfigsRequest struct {
	// The notification channel type used to filter results. If you do not specify this parameter, all personal notification channels are returned. Currently, only FEISHU (Lark) is supported. Specifying other values results in a parameter error.
	//
	// example:
	//
	// FEISHU
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The maximum number of records to return per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The paging token for the paged query. Do not specify this parameter for the first query. To query the next page, specify the NextToken value returned in the previous response.
	//
	// example:
	//
	// MjA
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListMyAgenticNotificationChannelConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMyAgenticNotificationChannelConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListMyAgenticNotificationChannelConfigsRequest) GetChannel() *string {
	return s.Channel
}

func (s *ListMyAgenticNotificationChannelConfigsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMyAgenticNotificationChannelConfigsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMyAgenticNotificationChannelConfigsRequest) SetChannel(v string) *ListMyAgenticNotificationChannelConfigsRequest {
	s.Channel = &v
	return s
}

func (s *ListMyAgenticNotificationChannelConfigsRequest) SetMaxResults(v int32) *ListMyAgenticNotificationChannelConfigsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMyAgenticNotificationChannelConfigsRequest) SetNextToken(v string) *ListMyAgenticNotificationChannelConfigsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMyAgenticNotificationChannelConfigsRequest) Validate() error {
	return dara.Validate(s)
}
