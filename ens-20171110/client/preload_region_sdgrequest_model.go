// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreloadRegionSDGRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationRegionIds(v []*string) *PreloadRegionSDGRequest
	GetDestinationRegionIds() []*string
	SetDiskType(v string) *PreloadRegionSDGRequest
	GetDiskType() *string
	SetNamespaces(v []*string) *PreloadRegionSDGRequest
	GetNamespaces() []*string
	SetRedundantNum(v int32) *PreloadRegionSDGRequest
	GetRedundantNum() *int32
	SetSDGId(v string) *PreloadRegionSDGRequest
	GetSDGId() *string
}

type PreloadRegionSDGRequest struct {
	// The IDs of the destination nodes.
	//
	// This parameter is required.
	DestinationRegionIds []*string `json:"DestinationRegionIds,omitempty" xml:"DestinationRegionIds,omitempty" type:"Repeated"`
	DiskType             *string   `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// An array that consists of queried namespaces.
	Namespaces []*string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
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

func (s PreloadRegionSDGRequest) String() string {
	return dara.Prettify(s)
}

func (s PreloadRegionSDGRequest) GoString() string {
	return s.String()
}

func (s *PreloadRegionSDGRequest) GetDestinationRegionIds() []*string {
	return s.DestinationRegionIds
}

func (s *PreloadRegionSDGRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *PreloadRegionSDGRequest) GetNamespaces() []*string {
	return s.Namespaces
}

func (s *PreloadRegionSDGRequest) GetRedundantNum() *int32 {
	return s.RedundantNum
}

func (s *PreloadRegionSDGRequest) GetSDGId() *string {
	return s.SDGId
}

func (s *PreloadRegionSDGRequest) SetDestinationRegionIds(v []*string) *PreloadRegionSDGRequest {
	s.DestinationRegionIds = v
	return s
}

func (s *PreloadRegionSDGRequest) SetDiskType(v string) *PreloadRegionSDGRequest {
	s.DiskType = &v
	return s
}

func (s *PreloadRegionSDGRequest) SetNamespaces(v []*string) *PreloadRegionSDGRequest {
	s.Namespaces = v
	return s
}

func (s *PreloadRegionSDGRequest) SetRedundantNum(v int32) *PreloadRegionSDGRequest {
	s.RedundantNum = &v
	return s
}

func (s *PreloadRegionSDGRequest) SetSDGId(v string) *PreloadRegionSDGRequest {
	s.SDGId = &v
	return s
}

func (s *PreloadRegionSDGRequest) Validate() error {
	return dara.Validate(s)
}
