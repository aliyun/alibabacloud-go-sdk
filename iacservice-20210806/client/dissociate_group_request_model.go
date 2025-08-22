// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DissociateGroupRequest
	GetClientToken() *string
	SetResourceIds(v []*string) *DissociateGroupRequest
	GetResourceIds() []*string
	SetResourceType(v string) *DissociateGroupRequest
	GetResourceType() *string
}

type DissociateGroupRequest struct {
	// example:
	//
	// 2daf4227f747cbf11a5501f18cc5e004
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// This parameter is required.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// example:
	//
	// Task
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s DissociateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DissociateGroupRequest) GoString() string {
	return s.String()
}

func (s *DissociateGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DissociateGroupRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *DissociateGroupRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DissociateGroupRequest) SetClientToken(v string) *DissociateGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DissociateGroupRequest) SetResourceIds(v []*string) *DissociateGroupRequest {
	s.ResourceIds = v
	return s
}

func (s *DissociateGroupRequest) SetResourceType(v string) *DissociateGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *DissociateGroupRequest) Validate() error {
	return dara.Validate(s)
}
