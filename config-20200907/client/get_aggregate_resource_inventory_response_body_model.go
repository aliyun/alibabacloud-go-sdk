// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceInventoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAggregateResourceInventoryResponseBody
	GetRequestId() *string
	SetResourceInventory(v *GetAggregateResourceInventoryResponseBodyResourceInventory) *GetAggregateResourceInventoryResponseBody
	GetResourceInventory() *GetAggregateResourceInventoryResponseBodyResourceInventory
}

type GetAggregateResourceInventoryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1A6D3604-EF1A-5798-A576-2A5FB855493C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resource inventory.
	ResourceInventory *GetAggregateResourceInventoryResponseBodyResourceInventory `json:"ResourceInventory,omitempty" xml:"ResourceInventory,omitempty" type:"Struct"`
}

func (s GetAggregateResourceInventoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceInventoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceInventoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregateResourceInventoryResponseBody) GetResourceInventory() *GetAggregateResourceInventoryResponseBodyResourceInventory {
	return s.ResourceInventory
}

func (s *GetAggregateResourceInventoryResponseBody) SetRequestId(v string) *GetAggregateResourceInventoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateResourceInventoryResponseBody) SetResourceInventory(v *GetAggregateResourceInventoryResponseBodyResourceInventory) *GetAggregateResourceInventoryResponseBody {
	s.ResourceInventory = v
	return s
}

func (s *GetAggregateResourceInventoryResponseBody) Validate() error {
	if s.ResourceInventory != nil {
		if err := s.ResourceInventory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAggregateResourceInventoryResponseBodyResourceInventory struct {
	// The download URL of the resource inventory.
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The time when the resource inventory was generated. The value is a timestamp.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1687674774123
	ResourceInventoryGenerateTime *int64 `json:"ResourceInventoryGenerateTime,omitempty" xml:"ResourceInventoryGenerateTime,omitempty"`
	// The generation status of the resource inventory. Valid values:
	//
	// 	- CREATING: The resource inventory is being generated.
	//
	// 	- COMPLETE: The resource inventory is generated.
	//
	// example:
	//
	// COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAggregateResourceInventoryResponseBodyResourceInventory) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceInventoryResponseBodyResourceInventory) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceInventoryResponseBodyResourceInventory) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetAggregateResourceInventoryResponseBodyResourceInventory) GetResourceInventoryGenerateTime() *int64 {
	return s.ResourceInventoryGenerateTime
}

func (s *GetAggregateResourceInventoryResponseBodyResourceInventory) GetStatus() *string {
	return s.Status
}

func (s *GetAggregateResourceInventoryResponseBodyResourceInventory) SetDownloadUrl(v string) *GetAggregateResourceInventoryResponseBodyResourceInventory {
	s.DownloadUrl = &v
	return s
}

func (s *GetAggregateResourceInventoryResponseBodyResourceInventory) SetResourceInventoryGenerateTime(v int64) *GetAggregateResourceInventoryResponseBodyResourceInventory {
	s.ResourceInventoryGenerateTime = &v
	return s
}

func (s *GetAggregateResourceInventoryResponseBodyResourceInventory) SetStatus(v string) *GetAggregateResourceInventoryResponseBodyResourceInventory {
	s.Status = &v
	return s
}

func (s *GetAggregateResourceInventoryResponseBodyResourceInventory) Validate() error {
	return dara.Validate(s)
}
