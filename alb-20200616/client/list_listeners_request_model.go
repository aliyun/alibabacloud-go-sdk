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
	SetListenerProtocol(v string) *ListListenersRequest
	GetListenerProtocol() *string
	SetLoadBalancerIds(v []*string) *ListListenersRequest
	GetLoadBalancerIds() []*string
	SetMaxResults(v int32) *ListListenersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListListenersRequest
	GetNextToken() *string
	SetTag(v []*ListListenersRequestTag) *ListListenersRequest
	GetTag() []*ListListenersRequestTag
}

type ListListenersRequest struct {
	// The listener IDs. You can specify at most 20 listener IDs.
	ListenerIds []*string `json:"ListenerIds,omitempty" xml:"ListenerIds,omitempty" type:"Repeated"`
	// The listener protocol. Valid values:
	//
	// 	- **HTTP**
	//
	// 	- **HTTPS**
	//
	// 	- **QUIC**
	//
	// example:
	//
	// HTTP
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The ALB instance ID. You can specify at most 20 instance IDs.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Repeated"`
	// The maximum number of entries to return. This parameter is optional. Valid values: **1 to 100**. If you do not specify this parameter, the default value **20*	- is used.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- If a value is returned for NextToken, you must specify the token that is obtained from the previous query as the value of **NextToken**.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The tags.
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

func (s *ListListenersRequest) GetListenerProtocol() *string {
	return s.ListenerProtocol
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

func (s *ListListenersRequest) GetTag() []*ListListenersRequestTag {
	return s.Tag
}

func (s *ListListenersRequest) SetListenerIds(v []*string) *ListListenersRequest {
	s.ListenerIds = v
	return s
}

func (s *ListListenersRequest) SetListenerProtocol(v string) *ListListenersRequest {
	s.ListenerProtocol = &v
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
	// The tag key. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// product
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
