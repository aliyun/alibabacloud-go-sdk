// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListNetworkServicesResponseBody
	GetMaxResults() *int32
	SetNetworkServices(v []*ListNetworkServicesResponseBodyNetworkServices) *ListNetworkServicesResponseBody
	GetNetworkServices() []*ListNetworkServicesResponseBodyNetworkServices
	SetNextToken(v string) *ListNetworkServicesResponseBody
	GetNextToken() *string
	SetQueues(v []*ListNetworkServicesResponseBodyQueues) *ListNetworkServicesResponseBody
	GetQueues() []*ListNetworkServicesResponseBodyQueues
	SetRequestId(v string) *ListNetworkServicesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListNetworkServicesResponseBody
	GetTotalCount() *int32
}

type ListNetworkServicesResponseBody struct {
	// 一次获取的最大记录数。
	//
	// example:
	//
	// 20
	MaxResults      *int32                                            `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NetworkServices []*ListNetworkServicesResponseBodyNetworkServices `json:"networkServices,omitempty" xml:"networkServices,omitempty" type:"Repeated"`
	// 下一页TOKEN。
	//
	// example:
	//
	// 1
	NextToken *string                                  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Queues    []*ListNetworkServicesResponseBodyQueues `json:"queues,omitempty" xml:"queues,omitempty" type:"Repeated"`
	// example:
	//
	// 18C7775A-7995-128A-A10C-9116EA87****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 记录总数。
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListNetworkServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworkServicesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNetworkServicesResponseBody) GetNetworkServices() []*ListNetworkServicesResponseBodyNetworkServices {
	return s.NetworkServices
}

func (s *ListNetworkServicesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNetworkServicesResponseBody) GetQueues() []*ListNetworkServicesResponseBodyQueues {
	return s.Queues
}

func (s *ListNetworkServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNetworkServicesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNetworkServicesResponseBody) SetMaxResults(v int32) *ListNetworkServicesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNetworkServicesResponseBody) SetNetworkServices(v []*ListNetworkServicesResponseBodyNetworkServices) *ListNetworkServicesResponseBody {
	s.NetworkServices = v
	return s
}

func (s *ListNetworkServicesResponseBody) SetNextToken(v string) *ListNetworkServicesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNetworkServicesResponseBody) SetQueues(v []*ListNetworkServicesResponseBodyQueues) *ListNetworkServicesResponseBody {
	s.Queues = v
	return s
}

func (s *ListNetworkServicesResponseBody) SetRequestId(v string) *ListNetworkServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNetworkServicesResponseBody) SetTotalCount(v int32) *ListNetworkServicesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNetworkServicesResponseBody) Validate() error {
	if s.NetworkServices != nil {
		for _, item := range s.NetworkServices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Queues != nil {
		for _, item := range s.Queues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNetworkServicesResponseBodyNetworkServices struct {
	// example:
	//
	// vpc_test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// ns-d7a1d02045fb****
	NetworkServiceId *string `json:"networkServiceId,omitempty" xml:"networkServiceId,omitempty"`
	// example:
	//
	// running
	State             *string                                                          `json:"state,omitempty" xml:"state,omitempty"`
	StateChangeReason *ListNetworkServicesResponseBodyNetworkServicesStateChangeReason `json:"stateChangeReason,omitempty" xml:"stateChangeReason,omitempty" type:"Struct"`
	// example:
	//
	// NetworkService
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// VPC id。
	//
	// example:
	//
	// vpc-bp1vt6r7o1w4tw7j6****
	VpcId      *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	VswitchIds []*string `json:"vswitchIds,omitempty" xml:"vswitchIds,omitempty" type:"Repeated"`
	// example:
	//
	// w-d2d82aa09155****
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListNetworkServicesResponseBodyNetworkServices) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkServicesResponseBodyNetworkServices) GoString() string {
	return s.String()
}

func (s *ListNetworkServicesResponseBodyNetworkServices) GetName() *string {
	return s.Name
}

func (s *ListNetworkServicesResponseBodyNetworkServices) GetNetworkServiceId() *string {
	return s.NetworkServiceId
}

func (s *ListNetworkServicesResponseBodyNetworkServices) GetState() *string {
	return s.State
}

func (s *ListNetworkServicesResponseBodyNetworkServices) GetStateChangeReason() *ListNetworkServicesResponseBodyNetworkServicesStateChangeReason {
	return s.StateChangeReason
}

func (s *ListNetworkServicesResponseBodyNetworkServices) GetType() *string {
	return s.Type
}

func (s *ListNetworkServicesResponseBodyNetworkServices) GetVpcId() *string {
	return s.VpcId
}

func (s *ListNetworkServicesResponseBodyNetworkServices) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *ListNetworkServicesResponseBodyNetworkServices) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListNetworkServicesResponseBodyNetworkServices) SetName(v string) *ListNetworkServicesResponseBodyNetworkServices {
	s.Name = &v
	return s
}

func (s *ListNetworkServicesResponseBodyNetworkServices) SetNetworkServiceId(v string) *ListNetworkServicesResponseBodyNetworkServices {
	s.NetworkServiceId = &v
	return s
}

func (s *ListNetworkServicesResponseBodyNetworkServices) SetState(v string) *ListNetworkServicesResponseBodyNetworkServices {
	s.State = &v
	return s
}

func (s *ListNetworkServicesResponseBodyNetworkServices) SetStateChangeReason(v *ListNetworkServicesResponseBodyNetworkServicesStateChangeReason) *ListNetworkServicesResponseBodyNetworkServices {
	s.StateChangeReason = v
	return s
}

func (s *ListNetworkServicesResponseBodyNetworkServices) SetType(v string) *ListNetworkServicesResponseBodyNetworkServices {
	s.Type = &v
	return s
}

func (s *ListNetworkServicesResponseBodyNetworkServices) SetVpcId(v string) *ListNetworkServicesResponseBodyNetworkServices {
	s.VpcId = &v
	return s
}

func (s *ListNetworkServicesResponseBodyNetworkServices) SetVswitchIds(v []*string) *ListNetworkServicesResponseBodyNetworkServices {
	s.VswitchIds = v
	return s
}

func (s *ListNetworkServicesResponseBodyNetworkServices) SetWorkspaceId(v string) *ListNetworkServicesResponseBodyNetworkServices {
	s.WorkspaceId = &v
	return s
}

func (s *ListNetworkServicesResponseBodyNetworkServices) Validate() error {
	if s.StateChangeReason != nil {
		if err := s.StateChangeReason.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNetworkServicesResponseBodyNetworkServicesStateChangeReason struct {
	// example:
	//
	// code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListNetworkServicesResponseBodyNetworkServicesStateChangeReason) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkServicesResponseBodyNetworkServicesStateChangeReason) GoString() string {
	return s.String()
}

func (s *ListNetworkServicesResponseBodyNetworkServicesStateChangeReason) GetCode() *string {
	return s.Code
}

func (s *ListNetworkServicesResponseBodyNetworkServicesStateChangeReason) GetMessage() *string {
	return s.Message
}

func (s *ListNetworkServicesResponseBodyNetworkServicesStateChangeReason) SetCode(v string) *ListNetworkServicesResponseBodyNetworkServicesStateChangeReason {
	s.Code = &v
	return s
}

func (s *ListNetworkServicesResponseBodyNetworkServicesStateChangeReason) SetMessage(v string) *ListNetworkServicesResponseBodyNetworkServicesStateChangeReason {
	s.Message = &v
	return s
}

func (s *ListNetworkServicesResponseBodyNetworkServicesStateChangeReason) Validate() error {
	return dara.Validate(s)
}

type ListNetworkServicesResponseBodyQueues struct {
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// ns-d7a1d02045fb****
	NetworkServiceId *string `json:"networkServiceId,omitempty" xml:"networkServiceId,omitempty"`
	// example:
	//
	// running
	State             *string                                                 `json:"state,omitempty" xml:"state,omitempty"`
	StateChangeReason *ListNetworkServicesResponseBodyQueuesStateChangeReason `json:"stateChangeReason,omitempty" xml:"stateChangeReason,omitempty" type:"Struct"`
	// example:
	//
	// NetworkService
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// vpc-uf6k2anfa9nzbm4cj****
	VpcId      *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	VswitchIds []*string `json:"vswitchIds,omitempty" xml:"vswitchIds,omitempty" type:"Repeated"`
	// example:
	//
	// w-d2d82aa09155****
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListNetworkServicesResponseBodyQueues) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkServicesResponseBodyQueues) GoString() string {
	return s.String()
}

func (s *ListNetworkServicesResponseBodyQueues) GetName() *string {
	return s.Name
}

func (s *ListNetworkServicesResponseBodyQueues) GetNetworkServiceId() *string {
	return s.NetworkServiceId
}

func (s *ListNetworkServicesResponseBodyQueues) GetState() *string {
	return s.State
}

func (s *ListNetworkServicesResponseBodyQueues) GetStateChangeReason() *ListNetworkServicesResponseBodyQueuesStateChangeReason {
	return s.StateChangeReason
}

func (s *ListNetworkServicesResponseBodyQueues) GetType() *string {
	return s.Type
}

func (s *ListNetworkServicesResponseBodyQueues) GetVpcId() *string {
	return s.VpcId
}

func (s *ListNetworkServicesResponseBodyQueues) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *ListNetworkServicesResponseBodyQueues) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListNetworkServicesResponseBodyQueues) SetName(v string) *ListNetworkServicesResponseBodyQueues {
	s.Name = &v
	return s
}

func (s *ListNetworkServicesResponseBodyQueues) SetNetworkServiceId(v string) *ListNetworkServicesResponseBodyQueues {
	s.NetworkServiceId = &v
	return s
}

func (s *ListNetworkServicesResponseBodyQueues) SetState(v string) *ListNetworkServicesResponseBodyQueues {
	s.State = &v
	return s
}

func (s *ListNetworkServicesResponseBodyQueues) SetStateChangeReason(v *ListNetworkServicesResponseBodyQueuesStateChangeReason) *ListNetworkServicesResponseBodyQueues {
	s.StateChangeReason = v
	return s
}

func (s *ListNetworkServicesResponseBodyQueues) SetType(v string) *ListNetworkServicesResponseBodyQueues {
	s.Type = &v
	return s
}

func (s *ListNetworkServicesResponseBodyQueues) SetVpcId(v string) *ListNetworkServicesResponseBodyQueues {
	s.VpcId = &v
	return s
}

func (s *ListNetworkServicesResponseBodyQueues) SetVswitchIds(v []*string) *ListNetworkServicesResponseBodyQueues {
	s.VswitchIds = v
	return s
}

func (s *ListNetworkServicesResponseBodyQueues) SetWorkspaceId(v string) *ListNetworkServicesResponseBodyQueues {
	s.WorkspaceId = &v
	return s
}

func (s *ListNetworkServicesResponseBodyQueues) Validate() error {
	if s.StateChangeReason != nil {
		if err := s.StateChangeReason.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNetworkServicesResponseBodyQueuesStateChangeReason struct {
	// example:
	//
	// code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListNetworkServicesResponseBodyQueuesStateChangeReason) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkServicesResponseBodyQueuesStateChangeReason) GoString() string {
	return s.String()
}

func (s *ListNetworkServicesResponseBodyQueuesStateChangeReason) GetCode() *string {
	return s.Code
}

func (s *ListNetworkServicesResponseBodyQueuesStateChangeReason) GetMessage() *string {
	return s.Message
}

func (s *ListNetworkServicesResponseBodyQueuesStateChangeReason) SetCode(v string) *ListNetworkServicesResponseBodyQueuesStateChangeReason {
	s.Code = &v
	return s
}

func (s *ListNetworkServicesResponseBodyQueuesStateChangeReason) SetMessage(v string) *ListNetworkServicesResponseBodyQueuesStateChangeReason {
	s.Message = &v
	return s
}

func (s *ListNetworkServicesResponseBodyQueuesStateChangeReason) Validate() error {
	return dara.Validate(s)
}
