// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetSelectionSelectedTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSelectionKey(v string) *ListAssetSelectionSelectedTargetRequest
	GetSelectionKey() *string
	SetTargetList(v []*string) *ListAssetSelectionSelectedTargetRequest
	GetTargetList() []*string
}

type ListAssetSelectionSelectedTargetRequest struct {
	// The unique ID of the asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2e6ab33d-4e00-4581-ac16-0dd1f9ad****
	SelectionKey *string `json:"SelectionKey,omitempty" xml:"SelectionKey,omitempty"`
	// The details of queries.
	TargetList []*string `json:"TargetList,omitempty" xml:"TargetList,omitempty" type:"Repeated"`
}

func (s ListAssetSelectionSelectedTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAssetSelectionSelectedTargetRequest) GoString() string {
	return s.String()
}

func (s *ListAssetSelectionSelectedTargetRequest) GetSelectionKey() *string {
	return s.SelectionKey
}

func (s *ListAssetSelectionSelectedTargetRequest) GetTargetList() []*string {
	return s.TargetList
}

func (s *ListAssetSelectionSelectedTargetRequest) SetSelectionKey(v string) *ListAssetSelectionSelectedTargetRequest {
	s.SelectionKey = &v
	return s
}

func (s *ListAssetSelectionSelectedTargetRequest) SetTargetList(v []*string) *ListAssetSelectionSelectedTargetRequest {
	s.TargetList = v
	return s
}

func (s *ListAssetSelectionSelectedTargetRequest) Validate() error {
	return dara.Validate(s)
}
