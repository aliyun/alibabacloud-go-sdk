// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateAccessTagsForDynamicRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicRouteIds(v []*string) *ListPrivateAccessTagsForDynamicRouteRequest
	GetDynamicRouteIds() []*string
}

type ListPrivateAccessTagsForDynamicRouteRequest struct {
	// This parameter is required.
	DynamicRouteIds []*string `json:"DynamicRouteIds,omitempty" xml:"DynamicRouteIds,omitempty" type:"Repeated"`
}

func (s ListPrivateAccessTagsForDynamicRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessTagsForDynamicRouteRequest) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessTagsForDynamicRouteRequest) GetDynamicRouteIds() []*string {
	return s.DynamicRouteIds
}

func (s *ListPrivateAccessTagsForDynamicRouteRequest) SetDynamicRouteIds(v []*string) *ListPrivateAccessTagsForDynamicRouteRequest {
	s.DynamicRouteIds = v
	return s
}

func (s *ListPrivateAccessTagsForDynamicRouteRequest) Validate() error {
	return dara.Validate(s)
}
