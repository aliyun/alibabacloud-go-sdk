// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListListenersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListeners(v []*ListListenersResponseBodyListeners) *ListListenersResponseBody
	GetListeners() []*ListListenersResponseBodyListeners
	SetMaxResults(v int32) *ListListenersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListListenersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListListenersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListListenersResponseBody
	GetTotalCount() *int32
}

type ListListenersResponseBody struct {
	// The GWLB listeners.
	Listeners []*ListListenersResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	// The maximum number of results to be returned from a single query when the NextToken parameter is used in the query. Valid values: 1 to 1000. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// 5c281c0a0d6bfb6355ed088c2108aca8e0b5e8707e68****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7DBFC67C-A272-5952-8287-6C3EBE4E04D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 8
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListListenersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListListenersResponseBody) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBody) GetListeners() []*ListListenersResponseBodyListeners {
	return s.Listeners
}

func (s *ListListenersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListListenersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListListenersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListListenersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListListenersResponseBody) SetListeners(v []*ListListenersResponseBodyListeners) *ListListenersResponseBody {
	s.Listeners = v
	return s
}

func (s *ListListenersResponseBody) SetMaxResults(v int32) *ListListenersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListListenersResponseBody) SetNextToken(v string) *ListListenersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListListenersResponseBody) SetRequestId(v string) *ListListenersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListListenersResponseBody) SetTotalCount(v int32) *ListListenersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListListenersResponseBody) Validate() error {
	if s.Listeners != nil {
		for _, item := range s.Listeners {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListListenersResponseBodyListeners struct {
	// The description of the listener.
	//
	// example:
	//
	// listener-description
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// The listener ID.
	//
	// example:
	//
	// lsn-vu7folhh5ntm8u****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The status of the listener. Valid values:
	//
	// 	- **Provisioning**: The listener is being created.
	//
	// 	- **Running**: The listener is running.
	//
	// 	- **Configuring**: The listener is being configured.
	//
	// 	- **Deleting**: The listener is being deleted.
	//
	// example:
	//
	// Running
	ListenerStatus *string `json:"ListenerStatus,omitempty" xml:"ListenerStatus,omitempty"`
	// The GWLB instance ID.
	//
	// example:
	//
	// gwlb-uf6hbeh795xlqln7g****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The server group ID.
	//
	// example:
	//
	// sgp-5yapcb422i51ru****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The tags.
	Tags []*ListListenersResponseBodyListenersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The timeout period of an idle TCP connection. Unit: seconds.
	//
	// example:
	//
	// 350
	TcpIdleTimeout *int32 `json:"TcpIdleTimeout,omitempty" xml:"TcpIdleTimeout,omitempty"`
}

func (s ListListenersResponseBodyListeners) String() string {
	return dara.Prettify(s)
}

func (s ListListenersResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListeners) GetListenerDescription() *string {
	return s.ListenerDescription
}

func (s *ListListenersResponseBodyListeners) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListListenersResponseBodyListeners) GetListenerStatus() *string {
	return s.ListenerStatus
}

func (s *ListListenersResponseBodyListeners) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *ListListenersResponseBodyListeners) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *ListListenersResponseBodyListeners) GetTags() []*ListListenersResponseBodyListenersTags {
	return s.Tags
}

func (s *ListListenersResponseBodyListeners) GetTcpIdleTimeout() *int32 {
	return s.TcpIdleTimeout
}

func (s *ListListenersResponseBodyListeners) SetListenerDescription(v string) *ListListenersResponseBodyListeners {
	s.ListenerDescription = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerId(v string) *ListListenersResponseBodyListeners {
	s.ListenerId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerStatus(v string) *ListListenersResponseBodyListeners {
	s.ListenerStatus = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetLoadBalancerId(v string) *ListListenersResponseBodyListeners {
	s.LoadBalancerId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetServerGroupId(v string) *ListListenersResponseBodyListeners {
	s.ServerGroupId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetTags(v []*ListListenersResponseBodyListenersTags) *ListListenersResponseBodyListeners {
	s.Tags = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetTcpIdleTimeout(v int32) *ListListenersResponseBodyListeners {
	s.TcpIdleTimeout = &v
	return s
}

func (s *ListListenersResponseBodyListeners) Validate() error {
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

type ListListenersResponseBodyListenersTags struct {
	// The tag key.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListListenersResponseBodyListenersTags) String() string {
	return dara.Prettify(s)
}

func (s ListListenersResponseBodyListenersTags) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersTags) GetKey() *string {
	return s.Key
}

func (s *ListListenersResponseBodyListenersTags) GetValue() *string {
	return s.Value
}

func (s *ListListenersResponseBodyListenersTags) SetKey(v string) *ListListenersResponseBodyListenersTags {
	s.Key = &v
	return s
}

func (s *ListListenersResponseBodyListenersTags) SetValue(v string) *ListListenersResponseBodyListenersTags {
	s.Value = &v
	return s
}

func (s *ListListenersResponseBodyListenersTags) Validate() error {
	return dara.Validate(s)
}
