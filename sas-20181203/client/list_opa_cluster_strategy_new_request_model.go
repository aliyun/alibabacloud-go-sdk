// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOpaClusterStrategyNewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListOpaClusterStrategyNewRequest
	GetCurrentPage() *int32
	SetImageName(v []*string) *ListOpaClusterStrategyNewRequest
	GetImageName() []*string
	SetLabel(v []*string) *ListOpaClusterStrategyNewRequest
	GetLabel() []*string
	SetPageSize(v int32) *ListOpaClusterStrategyNewRequest
	GetPageSize() *int32
	SetStrategyName(v []*string) *ListOpaClusterStrategyNewRequest
	GetStrategyName() []*string
}

type ListOpaClusterStrategyNewRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The image names.
	ImageName []*string `json:"ImageName,omitempty" xml:"ImageName,omitempty" type:"Repeated"`
	// The tags that are added to the container.
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The rule names.
	StrategyName []*string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty" type:"Repeated"`
}

func (s ListOpaClusterStrategyNewRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOpaClusterStrategyNewRequest) GoString() string {
	return s.String()
}

func (s *ListOpaClusterStrategyNewRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListOpaClusterStrategyNewRequest) GetImageName() []*string {
	return s.ImageName
}

func (s *ListOpaClusterStrategyNewRequest) GetLabel() []*string {
	return s.Label
}

func (s *ListOpaClusterStrategyNewRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOpaClusterStrategyNewRequest) GetStrategyName() []*string {
	return s.StrategyName
}

func (s *ListOpaClusterStrategyNewRequest) SetCurrentPage(v int32) *ListOpaClusterStrategyNewRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListOpaClusterStrategyNewRequest) SetImageName(v []*string) *ListOpaClusterStrategyNewRequest {
	s.ImageName = v
	return s
}

func (s *ListOpaClusterStrategyNewRequest) SetLabel(v []*string) *ListOpaClusterStrategyNewRequest {
	s.Label = v
	return s
}

func (s *ListOpaClusterStrategyNewRequest) SetPageSize(v int32) *ListOpaClusterStrategyNewRequest {
	s.PageSize = &v
	return s
}

func (s *ListOpaClusterStrategyNewRequest) SetStrategyName(v []*string) *ListOpaClusterStrategyNewRequest {
	s.StrategyName = v
	return s
}

func (s *ListOpaClusterStrategyNewRequest) Validate() error {
	return dara.Validate(s)
}
