// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerListenersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListeners(v *DescribeLoadBalancerListenersResponseBodyListeners) *DescribeLoadBalancerListenersResponseBody
	GetListeners() *DescribeLoadBalancerListenersResponseBodyListeners
	SetPageNumber(v int32) *DescribeLoadBalancerListenersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLoadBalancerListenersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLoadBalancerListenersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeLoadBalancerListenersResponseBody
	GetTotalCount() *int32
}

type DescribeLoadBalancerListenersResponseBody struct {
	Listeners *DescribeLoadBalancerListenersResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F3B261DD-3858-4D3C-877D-303ADF374600
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 49
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLoadBalancerListenersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerListenersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenersResponseBody) GetListeners() *DescribeLoadBalancerListenersResponseBodyListeners {
	return s.Listeners
}

func (s *DescribeLoadBalancerListenersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLoadBalancerListenersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLoadBalancerListenersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoadBalancerListenersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeLoadBalancerListenersResponseBody) SetListeners(v *DescribeLoadBalancerListenersResponseBodyListeners) *DescribeLoadBalancerListenersResponseBody {
	s.Listeners = v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBody) SetPageNumber(v int32) *DescribeLoadBalancerListenersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBody) SetPageSize(v int32) *DescribeLoadBalancerListenersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBody) SetRequestId(v string) *DescribeLoadBalancerListenersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBody) SetTotalCount(v int32) *DescribeLoadBalancerListenersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBody) Validate() error {
	if s.Listeners != nil {
		if err := s.Listeners.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLoadBalancerListenersResponseBodyListeners struct {
	Listener []*DescribeLoadBalancerListenersResponseBodyListenersListener `json:"Listener,omitempty" xml:"Listener,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerListenersResponseBodyListeners) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerListenersResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) GetListener() []*DescribeLoadBalancerListenersResponseBodyListenersListener {
	return s.Listener
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetListener(v []*DescribeLoadBalancerListenersResponseBodyListenersListener) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.Listener = v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) Validate() error {
	if s.Listener != nil {
		for _, item := range s.Listener {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLoadBalancerListenersResponseBodyListenersListener struct {
	BackendServerPort *int32  `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ForwardPort       *string `json:"ForwardPort,omitempty" xml:"ForwardPort,omitempty"`
	ListenerForward   *string `json:"ListenerForward,omitempty" xml:"ListenerForward,omitempty"`
	ListenerPort      *string `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	LoadBalancerId    *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	Protocol          *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeLoadBalancerListenersResponseBodyListenersListener) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerListenersResponseBodyListenersListener) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersListener) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersListener) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersListener) GetDescription() *string {
	return s.Description
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersListener) GetForwardPort() *string {
	return s.ForwardPort
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersListener) GetListenerForward() *string {
	return s.ListenerForward
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersListener) GetListenerPort() *string {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersListener) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersListener) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersListener) GetStatus() *string {
	return s.Status
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersListener) SetBackendServerPort(v int32) *DescribeLoadBalancerListenersResponseBodyListenersListener {
	s.BackendServerPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersListener) SetCreateTime(v string) *DescribeLoadBalancerListenersResponseBodyListenersListener {
	s.CreateTime = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersListener) SetDescription(v string) *DescribeLoadBalancerListenersResponseBodyListenersListener {
	s.Description = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersListener) SetForwardPort(v string) *DescribeLoadBalancerListenersResponseBodyListenersListener {
	s.ForwardPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersListener) SetListenerForward(v string) *DescribeLoadBalancerListenersResponseBodyListenersListener {
	s.ListenerForward = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersListener) SetListenerPort(v string) *DescribeLoadBalancerListenersResponseBodyListenersListener {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersListener) SetLoadBalancerId(v string) *DescribeLoadBalancerListenersResponseBodyListenersListener {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersListener) SetProtocol(v string) *DescribeLoadBalancerListenersResponseBodyListenersListener {
	s.Protocol = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersListener) SetStatus(v string) *DescribeLoadBalancerListenersResponseBodyListenersListener {
	s.Status = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersListener) Validate() error {
	return dara.Validate(s)
}
