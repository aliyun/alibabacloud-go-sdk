// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLifecycleActionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLifecycleActions(v []*DescribeLifecycleActionsResponseBodyLifecycleActions) *DescribeLifecycleActionsResponseBody
	GetLifecycleActions() []*DescribeLifecycleActionsResponseBodyLifecycleActions
	SetMaxResults(v int32) *DescribeLifecycleActionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeLifecycleActionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeLifecycleActionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeLifecycleActionsResponseBody
	GetTotalCount() *int32
}

type DescribeLifecycleActionsResponseBody struct {
	// The actions of the lifecycle hook.
	LifecycleActions []*DescribeLifecycleActionsResponseBodyLifecycleActions `json:"LifecycleActions,omitempty" xml:"LifecycleActions,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 3
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token returned in this call.
	//
	// example:
	//
	// AAAAAcSz4VTb1Nq****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 42A742EB-FCF3-459E-9C62-E107048C17E3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the queried lifecycle actions.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLifecycleActionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLifecycleActionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleActionsResponseBody) GetLifecycleActions() []*DescribeLifecycleActionsResponseBodyLifecycleActions {
	return s.LifecycleActions
}

func (s *DescribeLifecycleActionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeLifecycleActionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeLifecycleActionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLifecycleActionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeLifecycleActionsResponseBody) SetLifecycleActions(v []*DescribeLifecycleActionsResponseBodyLifecycleActions) *DescribeLifecycleActionsResponseBody {
	s.LifecycleActions = v
	return s
}

func (s *DescribeLifecycleActionsResponseBody) SetMaxResults(v int32) *DescribeLifecycleActionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeLifecycleActionsResponseBody) SetNextToken(v string) *DescribeLifecycleActionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeLifecycleActionsResponseBody) SetRequestId(v string) *DescribeLifecycleActionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLifecycleActionsResponseBody) SetTotalCount(v int32) *DescribeLifecycleActionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLifecycleActionsResponseBody) Validate() error {
	if s.LifecycleActions != nil {
		for _, item := range s.LifecycleActions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLifecycleActionsResponseBodyLifecycleActions struct {
	// The IDs of the ECS instances on which the lifecycle hook takes effect
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The subsequent action that Auto Scaling performs after the lifecycle hook times out. Valid values:
	//
	// 	- CONTINUE: Auto Scaling continues to respond to a scale-in or scale-out request.
	//
	// 	- ABANDON: Auto Scaling releases ECS instances that are created during scale-out events, or removes ECS instances from the scaling group during scale-in events.
	//
	// example:
	//
	// CONTINUE
	LifecycleActionResult *string `json:"LifecycleActionResult,omitempty" xml:"LifecycleActionResult,omitempty"`
	// The status of the lifecycle hook action.
	//
	// example:
	//
	// Pending
	LifecycleActionStatus *string `json:"LifecycleActionStatus,omitempty" xml:"LifecycleActionStatus,omitempty"`
	// The token of the lifecycle hook action.
	//
	// example:
	//
	// 9C2E9DA7-F794-449A-ACF6-CEE24444F7BB
	LifecycleActionToken *string `json:"LifecycleActionToken,omitempty" xml:"LifecycleActionToken,omitempty"`
	// The ID of the lifecycle hook.
	//
	// example:
	//
	// ash-bp18uoft0deax0f7****
	LifecycleHookId *string `json:"LifecycleHookId,omitempty" xml:"LifecycleHookId,omitempty"`
}

func (s DescribeLifecycleActionsResponseBodyLifecycleActions) String() string {
	return dara.Prettify(s)
}

func (s DescribeLifecycleActionsResponseBodyLifecycleActions) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActions) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActions) GetLifecycleActionResult() *string {
	return s.LifecycleActionResult
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActions) GetLifecycleActionStatus() *string {
	return s.LifecycleActionStatus
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActions) GetLifecycleActionToken() *string {
	return s.LifecycleActionToken
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActions) GetLifecycleHookId() *string {
	return s.LifecycleHookId
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActions) SetInstanceIds(v []*string) *DescribeLifecycleActionsResponseBodyLifecycleActions {
	s.InstanceIds = v
	return s
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActions) SetLifecycleActionResult(v string) *DescribeLifecycleActionsResponseBodyLifecycleActions {
	s.LifecycleActionResult = &v
	return s
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActions) SetLifecycleActionStatus(v string) *DescribeLifecycleActionsResponseBodyLifecycleActions {
	s.LifecycleActionStatus = &v
	return s
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActions) SetLifecycleActionToken(v string) *DescribeLifecycleActionsResponseBodyLifecycleActions {
	s.LifecycleActionToken = &v
	return s
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActions) SetLifecycleHookId(v string) *DescribeLifecycleActionsResponseBodyLifecycleActions {
	s.LifecycleHookId = &v
	return s
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActions) Validate() error {
	return dara.Validate(s)
}
