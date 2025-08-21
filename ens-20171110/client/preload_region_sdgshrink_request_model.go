// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreloadRegionSDGShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationRegionIdsShrink(v string) *PreloadRegionSDGShrinkRequest
	GetDestinationRegionIdsShrink() *string
	SetDiskType(v string) *PreloadRegionSDGShrinkRequest
	GetDiskType() *string
	SetNamespacesShrink(v string) *PreloadRegionSDGShrinkRequest
	GetNamespacesShrink() *string
	SetRedundantNum(v int32) *PreloadRegionSDGShrinkRequest
	GetRedundantNum() *int32
	SetSDGId(v string) *PreloadRegionSDGShrinkRequest
	GetSDGId() *string
}

type PreloadRegionSDGShrinkRequest struct {
	// The IDs of the destination nodes.
	//
	// This parameter is required.
	DestinationRegionIdsShrink *string `json:"DestinationRegionIds,omitempty" xml:"DestinationRegionIds,omitempty"`
	DiskType                   *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// An array that consists of queried namespaces.
	NamespacesShrink *string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty"`
	// The number of redundant replicas to support rapid deployment.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	RedundantNum *int32 `json:"RedundantNum,omitempty" xml:"RedundantNum,omitempty"`
	// The ID of the SDG for which data is preloaded.
	//
	// This parameter is required.
	//
	// example:
	//
	// sdg-xxxx
	SDGId *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
}

func (s PreloadRegionSDGShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PreloadRegionSDGShrinkRequest) GoString() string {
	return s.String()
}

func (s *PreloadRegionSDGShrinkRequest) GetDestinationRegionIdsShrink() *string {
	return s.DestinationRegionIdsShrink
}

func (s *PreloadRegionSDGShrinkRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *PreloadRegionSDGShrinkRequest) GetNamespacesShrink() *string {
	return s.NamespacesShrink
}

func (s *PreloadRegionSDGShrinkRequest) GetRedundantNum() *int32 {
	return s.RedundantNum
}

func (s *PreloadRegionSDGShrinkRequest) GetSDGId() *string {
	return s.SDGId
}

func (s *PreloadRegionSDGShrinkRequest) SetDestinationRegionIdsShrink(v string) *PreloadRegionSDGShrinkRequest {
	s.DestinationRegionIdsShrink = &v
	return s
}

func (s *PreloadRegionSDGShrinkRequest) SetDiskType(v string) *PreloadRegionSDGShrinkRequest {
	s.DiskType = &v
	return s
}

func (s *PreloadRegionSDGShrinkRequest) SetNamespacesShrink(v string) *PreloadRegionSDGShrinkRequest {
	s.NamespacesShrink = &v
	return s
}

func (s *PreloadRegionSDGShrinkRequest) SetRedundantNum(v int32) *PreloadRegionSDGShrinkRequest {
	s.RedundantNum = &v
	return s
}

func (s *PreloadRegionSDGShrinkRequest) SetSDGId(v string) *PreloadRegionSDGShrinkRequest {
	s.SDGId = &v
	return s
}

func (s *PreloadRegionSDGShrinkRequest) Validate() error {
	return dara.Validate(s)
}
