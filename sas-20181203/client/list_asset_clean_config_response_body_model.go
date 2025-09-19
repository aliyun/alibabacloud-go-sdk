// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetCleanConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListAssetCleanConfigResponseBody
	GetCount() *int32
	SetData(v []*ListAssetCleanConfigResponseBodyData) *ListAssetCleanConfigResponseBody
	GetData() []*ListAssetCleanConfigResponseBodyData
	SetRequestId(v string) *ListAssetCleanConfigResponseBody
	GetRequestId() *string
}

type ListAssetCleanConfigResponseBody struct {
	// The number of cleanup configurations.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The asset cleanup configurations.
	Data []*ListAssetCleanConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 20456DD5-5CBF-5015-9173-12CA4246B***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAssetCleanConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAssetCleanConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssetCleanConfigResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListAssetCleanConfigResponseBody) GetData() []*ListAssetCleanConfigResponseBodyData {
	return s.Data
}

func (s *ListAssetCleanConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAssetCleanConfigResponseBody) SetCount(v int32) *ListAssetCleanConfigResponseBody {
	s.Count = &v
	return s
}

func (s *ListAssetCleanConfigResponseBody) SetData(v []*ListAssetCleanConfigResponseBodyData) *ListAssetCleanConfigResponseBody {
	s.Data = v
	return s
}

func (s *ListAssetCleanConfigResponseBody) SetRequestId(v string) *ListAssetCleanConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAssetCleanConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAssetCleanConfigResponseBodyData struct {
	// The number of days before hosts whose provider cannot be identified are automatically cleaned after they enter the offline state. Valid value: an integer that ranges from 1 to 30.
	//
	// example:
	//
	// 7
	CleanDays *int32 `json:"CleanDays,omitempty" xml:"CleanDays,omitempty"`
	// Indicates whether the configuration takes effect. Valid values:
	//
	// 	- **0**: The configuration does not take effect.
	//
	// 	- **1**: The configuration takes effect.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of hosts that are cleaned.
	//
	// 	- The value is set to **1**, which indicates hosts whose provider cannot be identified.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAssetCleanConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAssetCleanConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAssetCleanConfigResponseBodyData) GetCleanDays() *int32 {
	return s.CleanDays
}

func (s *ListAssetCleanConfigResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ListAssetCleanConfigResponseBodyData) GetType() *int32 {
	return s.Type
}

func (s *ListAssetCleanConfigResponseBodyData) SetCleanDays(v int32) *ListAssetCleanConfigResponseBodyData {
	s.CleanDays = &v
	return s
}

func (s *ListAssetCleanConfigResponseBodyData) SetStatus(v int32) *ListAssetCleanConfigResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListAssetCleanConfigResponseBodyData) SetType(v int32) *ListAssetCleanConfigResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListAssetCleanConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
