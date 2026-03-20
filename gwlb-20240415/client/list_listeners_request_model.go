// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListListenersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerIds(v []*string) *ListListenersRequest
	GetListenerIds() []*string
	SetLoadBalancerIds(v []*string) *ListListenersRequest
	GetLoadBalancerIds() []*string
	SetMaxResults(v int32) *ListListenersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListListenersRequest
	GetNextToken() *string
	SetSkip(v int32) *ListListenersRequest
	GetSkip() *int32
	SetTag(v []*ListListenersRequestTag) *ListListenersRequest
	GetTag() []*ListListenersRequestTag
}

type ListListenersRequest struct {
	// The listener IDs. You can specify at most 20 listener IDs.
	ListenerIds []*string `json:"ListenerIds,omitempty" xml:"ListenerIds,omitempty" type:"Repeated"`
	// The GWLB instance IDs. You can specify at most 20 instance IDs.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Repeated"`
	// The maximum number of results to be returned from a single query when the NextToken parameter is used in the query. Valid values: 1 to 1000. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// d209f4e63ec942c967c50c888a13****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries to be skipped in the call.
	//
	// example:
	//
	// 10
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The tags. You can specify at most 20 tags in each call.
	Tag []*ListListenersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListListenersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListListenersRequest) GoString() string {
	return s.String()
}

func (s *ListListenersRequest) GetListenerIds() []*string {
	return s.ListenerIds
}

func (s *ListListenersRequest) GetLoadBalancerIds() []*string {
	return s.LoadBalancerIds
}

func (s *ListListenersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListListenersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListListenersRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListListenersRequest) GetTag() []*ListListenersRequestTag {
	return s.Tag
}

func (s *ListListenersRequest) SetListenerIds(v []*string) *ListListenersRequest {
	s.ListenerIds = v
	return s
}

func (s *ListListenersRequest) SetLoadBalancerIds(v []*string) *ListListenersRequest {
	s.LoadBalancerIds = v
	return s
}

func (s *ListListenersRequest) SetMaxResults(v int32) *ListListenersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListListenersRequest) SetNextToken(v string) *ListListenersRequest {
	s.NextToken = &v
	return s
}

func (s *ListListenersRequest) SetSkip(v int32) *ListListenersRequest {
	s.Skip = &v
	return s
}

func (s *ListListenersRequest) SetTag(v []*ListListenersRequestTag) *ListListenersRequest {
	s.Tag = v
	return s
}

func (s *ListListenersRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListListenersRequestTag struct {
	// The tag key. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. The tag key cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// tagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 256 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// tagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListListenersRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListListenersRequestTag) GoString() string {
	return s.String()
}

func (s *ListListenersRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListListenersRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListListenersRequestTag) SetKey(v string) *ListListenersRequestTag {
	s.Key = &v
	return s
}

func (s *ListListenersRequestTag) SetValue(v string) *ListListenersRequestTag {
	s.Value = &v
	return s
}

func (s *ListListenersRequestTag) Validate() error {
	return dara.Validate(s)
}
