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
	RequestId         *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	DownloadUrl                   *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	ResourceInventoryGenerateTime *int64  `json:"ResourceInventoryGenerateTime,omitempty" xml:"ResourceInventoryGenerateTime,omitempty"`
	Status                        *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
