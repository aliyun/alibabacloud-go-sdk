// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceInventoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetResourceInventoryResponseBody
	GetRequestId() *string
	SetResourceInventory(v *GetResourceInventoryResponseBodyResourceInventory) *GetResourceInventoryResponseBody
	GetResourceInventory() *GetResourceInventoryResponseBodyResourceInventory
}

type GetResourceInventoryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 89E2F38F-4EE4-545A-BD56-92E007ECFEE9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resource inventory.
	ResourceInventory *GetResourceInventoryResponseBodyResourceInventory `json:"ResourceInventory,omitempty" xml:"ResourceInventory,omitempty" type:"Struct"`
}

func (s GetResourceInventoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceInventoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceInventoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceInventoryResponseBody) GetResourceInventory() *GetResourceInventoryResponseBodyResourceInventory {
	return s.ResourceInventory
}

func (s *GetResourceInventoryResponseBody) SetRequestId(v string) *GetResourceInventoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceInventoryResponseBody) SetResourceInventory(v *GetResourceInventoryResponseBodyResourceInventory) *GetResourceInventoryResponseBody {
	s.ResourceInventory = v
	return s
}

func (s *GetResourceInventoryResponseBody) Validate() error {
	if s.ResourceInventory != nil {
		if err := s.ResourceInventory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceInventoryResponseBodyResourceInventory struct {
	// The download URL of the resource inventory.
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The time when the resource inventory was generated. The value is a timestamp.
	//
	// example:
	//
	// 1687674634220
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

func (s GetResourceInventoryResponseBodyResourceInventory) String() string {
	return dara.Prettify(s)
}

func (s GetResourceInventoryResponseBodyResourceInventory) GoString() string {
	return s.String()
}

func (s *GetResourceInventoryResponseBodyResourceInventory) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetResourceInventoryResponseBodyResourceInventory) GetResourceInventoryGenerateTime() *int64 {
	return s.ResourceInventoryGenerateTime
}

func (s *GetResourceInventoryResponseBodyResourceInventory) GetStatus() *string {
	return s.Status
}

func (s *GetResourceInventoryResponseBodyResourceInventory) SetDownloadUrl(v string) *GetResourceInventoryResponseBodyResourceInventory {
	s.DownloadUrl = &v
	return s
}

func (s *GetResourceInventoryResponseBodyResourceInventory) SetResourceInventoryGenerateTime(v int64) *GetResourceInventoryResponseBodyResourceInventory {
	s.ResourceInventoryGenerateTime = &v
	return s
}

func (s *GetResourceInventoryResponseBodyResourceInventory) SetStatus(v string) *GetResourceInventoryResponseBodyResourceInventory {
	s.Status = &v
	return s
}

func (s *GetResourceInventoryResponseBodyResourceInventory) Validate() error {
	return dara.Validate(s)
}
