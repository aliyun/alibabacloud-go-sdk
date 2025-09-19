// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAssetSelectionConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateAssetSelectionConfigResponseBodyData) *CreateAssetSelectionConfigResponseBody
	GetData() *CreateAssetSelectionConfigResponseBodyData
	SetRequestId(v string) *CreateAssetSelectionConfigResponseBody
	GetRequestId() *string
}

type CreateAssetSelectionConfigResponseBody struct {
	// The returned data.
	Data *CreateAssetSelectionConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A3D7C47D-3F11-57BB-90E8-E5C20C619F37
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAssetSelectionConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAssetSelectionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAssetSelectionConfigResponseBody) GetData() *CreateAssetSelectionConfigResponseBodyData {
	return s.Data
}

func (s *CreateAssetSelectionConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAssetSelectionConfigResponseBody) SetData(v *CreateAssetSelectionConfigResponseBodyData) *CreateAssetSelectionConfigResponseBody {
	s.Data = v
	return s
}

func (s *CreateAssetSelectionConfigResponseBody) SetRequestId(v string) *CreateAssetSelectionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAssetSelectionConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateAssetSelectionConfigResponseBodyData struct {
	// The business type that is selected for the asset. Valid values:
	//
	// 	- **VIRUS_SCAN_CYCLE_CONFIG**: virus detection configuration
	//
	// 	- **VIRUS_SCAN_ONCE_TASK**: one-time scan for virus detection
	//
	// example:
	//
	// VIRUS_SCAN_CYCLE_CONFIG
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The operating system of the asset. Valid values:
	//
	// 	- **windows**: the Windows operating system
	//
	// 	- **linux**: the Linux operating system
	//
	// example:
	//
	// all
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The ID of the current asset selection. It can be used to query and modify the asset that is selected.
	//
	// example:
	//
	// 53e93435-d694-4c03-9ce7-da12bee1****
	SelectionKey *int64 `json:"SelectionKey,omitempty" xml:"SelectionKey,omitempty"`
	// The dimension based on which the asset is selected. Valid values:
	//
	// 	- **instance**: The asset is selected by server.
	//
	// 	- **group**: The asset is selected by group.
	//
	// 	- **vpc**: The asset is selected by VPC.
	//
	// example:
	//
	// group
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateAssetSelectionConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAssetSelectionConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAssetSelectionConfigResponseBodyData) GetBusinessType() *string {
	return s.BusinessType
}

func (s *CreateAssetSelectionConfigResponseBodyData) GetPlatform() *string {
	return s.Platform
}

func (s *CreateAssetSelectionConfigResponseBodyData) GetSelectionKey() *int64 {
	return s.SelectionKey
}

func (s *CreateAssetSelectionConfigResponseBodyData) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateAssetSelectionConfigResponseBodyData) SetBusinessType(v string) *CreateAssetSelectionConfigResponseBodyData {
	s.BusinessType = &v
	return s
}

func (s *CreateAssetSelectionConfigResponseBodyData) SetPlatform(v string) *CreateAssetSelectionConfigResponseBodyData {
	s.Platform = &v
	return s
}

func (s *CreateAssetSelectionConfigResponseBodyData) SetSelectionKey(v int64) *CreateAssetSelectionConfigResponseBodyData {
	s.SelectionKey = &v
	return s
}

func (s *CreateAssetSelectionConfigResponseBodyData) SetTargetType(v string) *CreateAssetSelectionConfigResponseBodyData {
	s.TargetType = &v
	return s
}

func (s *CreateAssetSelectionConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
