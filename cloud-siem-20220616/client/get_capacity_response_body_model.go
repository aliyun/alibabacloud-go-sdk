// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCapacityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetCapacityResponseBodyData) *GetCapacityResponseBody
	GetData() *GetCapacityResponseBodyData
	SetRequestId(v string) *GetCapacityResponseBody
	GetRequestId() *string
}

type GetCapacityResponseBody struct {
	// The information about the storage capacity.
	Data *GetCapacityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 27D27DCB-D76B-5064-8B3B-0900DEF7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCapacityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCapacityResponseBody) GoString() string {
	return s.String()
}

func (s *GetCapacityResponseBody) GetData() *GetCapacityResponseBodyData {
	return s.Data
}

func (s *GetCapacityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCapacityResponseBody) SetData(v *GetCapacityResponseBodyData) *GetCapacityResponseBody {
	s.Data = v
	return s
}

func (s *GetCapacityResponseBody) SetRequestId(v string) *GetCapacityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCapacityResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCapacityResponseBodyData struct {
	// Indicates whether the Logstores for the threat analysis feature exist on the user side. Valid values:
	//
	// 	- true: The logs are in the normal state. The log analysis feature is available.
	//
	// 	- false: The logs are being cleared. The log analysis feature is unavailable.
	//
	// example:
	//
	// true
	ExistLogStore *bool `json:"ExistLogStore,omitempty" xml:"ExistLogStore,omitempty"`
	// The purchased storage capacity of the threat analysis feature. Unit: GB.
	//
	// example:
	//
	// 9000
	PreservedCapacity *int64 `json:"PreservedCapacity,omitempty" xml:"PreservedCapacity,omitempty"`
	// The billable storage capacity of the threat analysis feature. Unit: GB.
	//
	// example:
	//
	// 10
	UsedCapacity *float64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
}

func (s GetCapacityResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCapacityResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCapacityResponseBodyData) GetExistLogStore() *bool {
	return s.ExistLogStore
}

func (s *GetCapacityResponseBodyData) GetPreservedCapacity() *int64 {
	return s.PreservedCapacity
}

func (s *GetCapacityResponseBodyData) GetUsedCapacity() *float64 {
	return s.UsedCapacity
}

func (s *GetCapacityResponseBodyData) SetExistLogStore(v bool) *GetCapacityResponseBodyData {
	s.ExistLogStore = &v
	return s
}

func (s *GetCapacityResponseBodyData) SetPreservedCapacity(v int64) *GetCapacityResponseBodyData {
	s.PreservedCapacity = &v
	return s
}

func (s *GetCapacityResponseBodyData) SetUsedCapacity(v float64) *GetCapacityResponseBodyData {
	s.UsedCapacity = &v
	return s
}

func (s *GetCapacityResponseBodyData) Validate() error {
	return dara.Validate(s)
}
