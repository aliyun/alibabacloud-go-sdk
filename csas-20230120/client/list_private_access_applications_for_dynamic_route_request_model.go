// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateAccessApplicationsForDynamicRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicRouteIds(v []*string) *ListPrivateAccessApplicationsForDynamicRouteRequest
	GetDynamicRouteIds() []*string
}

type ListPrivateAccessApplicationsForDynamicRouteRequest struct {
	// This parameter is required.
	DynamicRouteIds []*string `json:"DynamicRouteIds,omitempty" xml:"DynamicRouteIds,omitempty" type:"Repeated"`
}

func (s ListPrivateAccessApplicationsForDynamicRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessApplicationsForDynamicRouteRequest) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationsForDynamicRouteRequest) GetDynamicRouteIds() []*string {
	return s.DynamicRouteIds
}

func (s *ListPrivateAccessApplicationsForDynamicRouteRequest) SetDynamicRouteIds(v []*string) *ListPrivateAccessApplicationsForDynamicRouteRequest {
	s.DynamicRouteIds = v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteRequest) Validate() error {
	return dara.Validate(s)
}
