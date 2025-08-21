// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnloadRegionSDGRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationRegionIds(v []*string) *UnloadRegionSDGRequest
	GetDestinationRegionIds() []*string
	SetNamespaces(v []*string) *UnloadRegionSDGRequest
	GetNamespaces() []*string
	SetSDGId(v string) *UnloadRegionSDGRequest
	GetSDGId() *string
}

type UnloadRegionSDGRequest struct {
	// The destination nodes.
	//
	// This parameter is required.
	DestinationRegionIds []*string `json:"DestinationRegionIds,omitempty" xml:"DestinationRegionIds,omitempty" type:"Repeated"`
	// An array that consists of queried namespaces.
	Namespaces []*string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// Deletes the shared data group (SDG) ID of the preloaded data.
	//
	// This parameter is required.
	//
	// example:
	//
	// sdg-xxxx
	SDGId *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
}

func (s UnloadRegionSDGRequest) String() string {
	return dara.Prettify(s)
}

func (s UnloadRegionSDGRequest) GoString() string {
	return s.String()
}

func (s *UnloadRegionSDGRequest) GetDestinationRegionIds() []*string {
	return s.DestinationRegionIds
}

func (s *UnloadRegionSDGRequest) GetNamespaces() []*string {
	return s.Namespaces
}

func (s *UnloadRegionSDGRequest) GetSDGId() *string {
	return s.SDGId
}

func (s *UnloadRegionSDGRequest) SetDestinationRegionIds(v []*string) *UnloadRegionSDGRequest {
	s.DestinationRegionIds = v
	return s
}

func (s *UnloadRegionSDGRequest) SetNamespaces(v []*string) *UnloadRegionSDGRequest {
	s.Namespaces = v
	return s
}

func (s *UnloadRegionSDGRequest) SetSDGId(v string) *UnloadRegionSDGRequest {
	s.SDGId = &v
	return s
}

func (s *UnloadRegionSDGRequest) Validate() error {
	return dara.Validate(s)
}
