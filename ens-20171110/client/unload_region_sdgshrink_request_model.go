// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnloadRegionSDGShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationRegionIdsShrink(v string) *UnloadRegionSDGShrinkRequest
	GetDestinationRegionIdsShrink() *string
	SetNamespacesShrink(v string) *UnloadRegionSDGShrinkRequest
	GetNamespacesShrink() *string
	SetSDGId(v string) *UnloadRegionSDGShrinkRequest
	GetSDGId() *string
}

type UnloadRegionSDGShrinkRequest struct {
	// The list of destination edge zones.
	//
	// This parameter is required.
	DestinationRegionIdsShrink *string `json:"DestinationRegionIds,omitempty" xml:"DestinationRegionIds,omitempty"`
	// The list of namespaces.
	NamespacesShrink *string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty"`
	// The ID of the SDG from which to delete preloaded data.
	//
	// This parameter is required.
	//
	// example:
	//
	// sdg-xxxx
	SDGId *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
}

func (s UnloadRegionSDGShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UnloadRegionSDGShrinkRequest) GoString() string {
	return s.String()
}

func (s *UnloadRegionSDGShrinkRequest) GetDestinationRegionIdsShrink() *string {
	return s.DestinationRegionIdsShrink
}

func (s *UnloadRegionSDGShrinkRequest) GetNamespacesShrink() *string {
	return s.NamespacesShrink
}

func (s *UnloadRegionSDGShrinkRequest) GetSDGId() *string {
	return s.SDGId
}

func (s *UnloadRegionSDGShrinkRequest) SetDestinationRegionIdsShrink(v string) *UnloadRegionSDGShrinkRequest {
	s.DestinationRegionIdsShrink = &v
	return s
}

func (s *UnloadRegionSDGShrinkRequest) SetNamespacesShrink(v string) *UnloadRegionSDGShrinkRequest {
	s.NamespacesShrink = &v
	return s
}

func (s *UnloadRegionSDGShrinkRequest) SetSDGId(v string) *UnloadRegionSDGShrinkRequest {
	s.SDGId = &v
	return s
}

func (s *UnloadRegionSDGShrinkRequest) Validate() error {
	return dara.Validate(s)
}
