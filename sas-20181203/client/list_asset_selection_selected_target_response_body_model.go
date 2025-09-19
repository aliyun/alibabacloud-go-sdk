// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetSelectionSelectedTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListAssetSelectionSelectedTargetResponseBodyData) *ListAssetSelectionSelectedTargetResponseBody
	GetData() []*ListAssetSelectionSelectedTargetResponseBodyData
	SetRequestId(v string) *ListAssetSelectionSelectedTargetResponseBody
	GetRequestId() *string
}

type ListAssetSelectionSelectedTargetResponseBody struct {
	// The data returned.
	Data []*ListAssetSelectionSelectedTargetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1EB12F73-6828-59D2-9FBF-F3713FD55128
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAssetSelectionSelectedTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAssetSelectionSelectedTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssetSelectionSelectedTargetResponseBody) GetData() []*ListAssetSelectionSelectedTargetResponseBodyData {
	return s.Data
}

func (s *ListAssetSelectionSelectedTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAssetSelectionSelectedTargetResponseBody) SetData(v []*ListAssetSelectionSelectedTargetResponseBodyData) *ListAssetSelectionSelectedTargetResponseBody {
	s.Data = v
	return s
}

func (s *ListAssetSelectionSelectedTargetResponseBody) SetRequestId(v string) *ListAssetSelectionSelectedTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAssetSelectionSelectedTargetResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAssetSelectionSelectedTargetResponseBodyData struct {
	// The ID of the asset.
	//
	// example:
	//
	// 30****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// jen****
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
}

func (s ListAssetSelectionSelectedTargetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAssetSelectionSelectedTargetResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAssetSelectionSelectedTargetResponseBodyData) GetTargetId() *string {
	return s.TargetId
}

func (s *ListAssetSelectionSelectedTargetResponseBodyData) GetTargetName() *string {
	return s.TargetName
}

func (s *ListAssetSelectionSelectedTargetResponseBodyData) SetTargetId(v string) *ListAssetSelectionSelectedTargetResponseBodyData {
	s.TargetId = &v
	return s
}

func (s *ListAssetSelectionSelectedTargetResponseBodyData) SetTargetName(v string) *ListAssetSelectionSelectedTargetResponseBodyData {
	s.TargetName = &v
	return s
}

func (s *ListAssetSelectionSelectedTargetResponseBodyData) Validate() error {
	return dara.Validate(s)
}
