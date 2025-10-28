// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSynchronizeResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceIds(v string) *SynchronizeResourceRequest
	GetResourceIds() *string
	SetType(v string) *SynchronizeResourceRequest
	GetType() *string
}

type SynchronizeResourceRequest struct {
	// The ID of the resource. This parameter is required only when you set the Type parameter to `ecs`. If you specify multiple IDs, separate them with commas (,). You can synchronize up to 50 resources at a time.
	//
	// example:
	//
	// i-bp17c***5q8x,i-bp1**5q8x
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The type of the resource to be synchronized. Valid values: `ecs, slb, vpc, and all`. These values are case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SynchronizeResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s SynchronizeResourceRequest) GoString() string {
	return s.String()
}

func (s *SynchronizeResourceRequest) GetResourceIds() *string {
	return s.ResourceIds
}

func (s *SynchronizeResourceRequest) GetType() *string {
	return s.Type
}

func (s *SynchronizeResourceRequest) SetResourceIds(v string) *SynchronizeResourceRequest {
	s.ResourceIds = &v
	return s
}

func (s *SynchronizeResourceRequest) SetType(v string) *SynchronizeResourceRequest {
	s.Type = &v
	return s
}

func (s *SynchronizeResourceRequest) Validate() error {
	return dara.Validate(s)
}
