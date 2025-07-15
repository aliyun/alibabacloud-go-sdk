// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCasterVideoResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *ModifyCasterVideoResourceResponseBody
	GetCasterId() *string
	SetRequestId(v string) *ModifyCasterVideoResourceResponseBody
	GetRequestId() *string
	SetResourceId(v string) *ModifyCasterVideoResourceResponseBody
	GetResourceId() *string
}

type ModifyCasterVideoResourceResponseBody struct {
	// The ID of the production studio. This parameter is used in the requests of the following operations: DescribeCasterVideoResources, AddCasterLayout, and DescribeCasterLayouts.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CF60DB6A-7FD6-426E-9288-122CC1A52FA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// 98461064-1c94-4dc1-85ce-94098764****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s ModifyCasterVideoResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCasterVideoResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCasterVideoResourceResponseBody) GetCasterId() *string {
	return s.CasterId
}

func (s *ModifyCasterVideoResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCasterVideoResourceResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *ModifyCasterVideoResourceResponseBody) SetCasterId(v string) *ModifyCasterVideoResourceResponseBody {
	s.CasterId = &v
	return s
}

func (s *ModifyCasterVideoResourceResponseBody) SetRequestId(v string) *ModifyCasterVideoResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCasterVideoResourceResponseBody) SetResourceId(v string) *ModifyCasterVideoResourceResponseBody {
	s.ResourceId = &v
	return s
}

func (s *ModifyCasterVideoResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
