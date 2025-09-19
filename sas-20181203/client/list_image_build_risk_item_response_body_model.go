// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageBuildRiskItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListImageBuildRiskItemResponseBodyData) *ListImageBuildRiskItemResponseBody
	GetData() []*ListImageBuildRiskItemResponseBodyData
	SetRequestId(v string) *ListImageBuildRiskItemResponseBody
	GetRequestId() *string
}

type ListImageBuildRiskItemResponseBody struct {
	// The response parameters.
	Data []*ListImageBuildRiskItemResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// A3D7C47D-3F11-57BB-90E8-E5C20C619F37
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListImageBuildRiskItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImageBuildRiskItemResponseBody) GoString() string {
	return s.String()
}

func (s *ListImageBuildRiskItemResponseBody) GetData() []*ListImageBuildRiskItemResponseBodyData {
	return s.Data
}

func (s *ListImageBuildRiskItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImageBuildRiskItemResponseBody) SetData(v []*ListImageBuildRiskItemResponseBodyData) *ListImageBuildRiskItemResponseBody {
	s.Data = v
	return s
}

func (s *ListImageBuildRiskItemResponseBody) SetRequestId(v string) *ListImageBuildRiskItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImageBuildRiskItemResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListImageBuildRiskItemResponseBodyData struct {
	// The type key of the risky build command.
	//
	// example:
	//
	// key
	ItemKey *string `json:"ItemKey,omitempty" xml:"ItemKey,omitempty"`
	// The type name of the risky build command.
	//
	// example:
	//
	// itemName.
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
}

func (s ListImageBuildRiskItemResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListImageBuildRiskItemResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListImageBuildRiskItemResponseBodyData) GetItemKey() *string {
	return s.ItemKey
}

func (s *ListImageBuildRiskItemResponseBodyData) GetItemName() *string {
	return s.ItemName
}

func (s *ListImageBuildRiskItemResponseBodyData) SetItemKey(v string) *ListImageBuildRiskItemResponseBodyData {
	s.ItemKey = &v
	return s
}

func (s *ListImageBuildRiskItemResponseBodyData) SetItemName(v string) *ListImageBuildRiskItemResponseBodyData {
	s.ItemName = &v
	return s
}

func (s *ListImageBuildRiskItemResponseBodyData) Validate() error {
	return dara.Validate(s)
}
