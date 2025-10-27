// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetSelectionConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAssetSelectionConfigResponseBodyData) *GetAssetSelectionConfigResponseBody
	GetData() *GetAssetSelectionConfigResponseBodyData
	SetRequestId(v string) *GetAssetSelectionConfigResponseBody
	GetRequestId() *string
}

type GetAssetSelectionConfigResponseBody struct {
	// The data returned.
	Data *GetAssetSelectionConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C5936B67-3EDF-53ED-A542-02543972449A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAssetSelectionConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAssetSelectionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAssetSelectionConfigResponseBody) GetData() *GetAssetSelectionConfigResponseBodyData {
	return s.Data
}

func (s *GetAssetSelectionConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAssetSelectionConfigResponseBody) SetData(v *GetAssetSelectionConfigResponseBodyData) *GetAssetSelectionConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetAssetSelectionConfigResponseBody) SetRequestId(v string) *GetAssetSelectionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAssetSelectionConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAssetSelectionConfigResponseBodyData struct {
	// The operating system of the asset. Valid values:
	//
	// 	- **windows**
	//
	// 	- **linux**
	//
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The ID of the current asset selection. It can be used to query and modify the asset that is selected.
	//
	// example:
	//
	// 657c8411-4e89-446c-ab66-d45d1331****
	SelectionKey *string `json:"SelectionKey,omitempty" xml:"SelectionKey,omitempty"`
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
	// instance
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s GetAssetSelectionConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAssetSelectionConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAssetSelectionConfigResponseBodyData) GetPlatform() *string {
	return s.Platform
}

func (s *GetAssetSelectionConfigResponseBodyData) GetSelectionKey() *string {
	return s.SelectionKey
}

func (s *GetAssetSelectionConfigResponseBodyData) GetTargetType() *string {
	return s.TargetType
}

func (s *GetAssetSelectionConfigResponseBodyData) SetPlatform(v string) *GetAssetSelectionConfigResponseBodyData {
	s.Platform = &v
	return s
}

func (s *GetAssetSelectionConfigResponseBodyData) SetSelectionKey(v string) *GetAssetSelectionConfigResponseBodyData {
	s.SelectionKey = &v
	return s
}

func (s *GetAssetSelectionConfigResponseBodyData) SetTargetType(v string) *GetAssetSelectionConfigResponseBodyData {
	s.TargetType = &v
	return s
}

func (s *GetAssetSelectionConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
