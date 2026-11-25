// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDimension(v string) *GetContextDatabaseUsageResponseBody
	GetDimension() *string
	SetEndTime(v string) *GetContextDatabaseUsageResponseBody
	GetEndTime() *string
	SetItems(v []*GetContextDatabaseUsageResponseBodyItems) *GetContextDatabaseUsageResponseBody
	GetItems() []*GetContextDatabaseUsageResponseBodyItems
	SetRequestId(v string) *GetContextDatabaseUsageResponseBody
	GetRequestId() *string
	SetStartTime(v string) *GetContextDatabaseUsageResponseBody
	GetStartTime() *string
	SetWorkspaceId(v string) *GetContextDatabaseUsageResponseBody
	GetWorkspaceId() *string
}

type GetContextDatabaseUsageResponseBody struct {
	// The query dimension. Valid values:
	//
	// - `Workspace`: A single workspace.
	//
	// - `Account`: All workspaces under the caller account.
	//
	// example:
	//
	// Workspace
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The end time of the query range passed in the request, exclusive.
	//
	// example:
	//
	// 2026-08-01T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The usage list. It returns all seven metering items in a fixed order. If an item has no data within the query range, its usage is 0.
	Items []*GetContextDatabaseUsageResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The ID of the API request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query range passed in the request, inclusive.
	//
	// example:
	//
	// 2026-07-01T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the queried workspace. This field has a value only when `Dimension` is set to `Workspace`. It is empty for the account dimension.
	//
	// example:
	//
	// ws-x9z4ctzed9gfu2w
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetContextDatabaseUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseUsageResponseBody) GetDimension() *string {
	return s.Dimension
}

func (s *GetContextDatabaseUsageResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetContextDatabaseUsageResponseBody) GetItems() []*GetContextDatabaseUsageResponseBodyItems {
	return s.Items
}

func (s *GetContextDatabaseUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContextDatabaseUsageResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetContextDatabaseUsageResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetContextDatabaseUsageResponseBody) SetDimension(v string) *GetContextDatabaseUsageResponseBody {
	s.Dimension = &v
	return s
}

func (s *GetContextDatabaseUsageResponseBody) SetEndTime(v string) *GetContextDatabaseUsageResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetContextDatabaseUsageResponseBody) SetItems(v []*GetContextDatabaseUsageResponseBodyItems) *GetContextDatabaseUsageResponseBody {
	s.Items = v
	return s
}

func (s *GetContextDatabaseUsageResponseBody) SetRequestId(v string) *GetContextDatabaseUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContextDatabaseUsageResponseBody) SetStartTime(v string) *GetContextDatabaseUsageResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetContextDatabaseUsageResponseBody) SetWorkspaceId(v string) *GetContextDatabaseUsageResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetContextDatabaseUsageResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetContextDatabaseUsageResponseBodyItems struct {
	// The metering item.
	//
	// example:
	//
	// KnowledgeBaseRead
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The usage of the item within the query range. The unit is specified by `Unit`. The number of calls, tokens, and context evolutions are accumulated over the range. Storage is calculated by averaging the occupancy snapshots of each workspace within the range and rounding to the nearest integer byte. For the account dimension, the average values of all workspaces are summed.
	//
	// example:
	//
	// 15
	Quantity *int64 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The metering unit. Valid values:
	//
	// - `Count`: The number of calls.
	//
	// - `Token`: The number of tokens.
	//
	// - `MicroCredit`: One millionth of a credit.
	//
	// - `Byte`: Bytes.
	//
	// example:
	//
	// Count
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s GetContextDatabaseUsageResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseUsageResponseBodyItems) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseUsageResponseBodyItems) GetItem() *string {
	return s.Item
}

func (s *GetContextDatabaseUsageResponseBodyItems) GetQuantity() *int64 {
	return s.Quantity
}

func (s *GetContextDatabaseUsageResponseBodyItems) GetUnit() *string {
	return s.Unit
}

func (s *GetContextDatabaseUsageResponseBodyItems) SetItem(v string) *GetContextDatabaseUsageResponseBodyItems {
	s.Item = &v
	return s
}

func (s *GetContextDatabaseUsageResponseBodyItems) SetQuantity(v int64) *GetContextDatabaseUsageResponseBodyItems {
	s.Quantity = &v
	return s
}

func (s *GetContextDatabaseUsageResponseBodyItems) SetUnit(v string) *GetContextDatabaseUsageResponseBodyItems {
	s.Unit = &v
	return s
}

func (s *GetContextDatabaseUsageResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
