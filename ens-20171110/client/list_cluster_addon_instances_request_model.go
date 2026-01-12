// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterAddonInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListClusterAddonInstancesRequest
	GetClusterId() *string
}

type ListClusterAddonInstancesRequest struct {
	// Cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListClusterAddonInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClusterAddonInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListClusterAddonInstancesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClusterAddonInstancesRequest) SetClusterId(v string) *ListClusterAddonInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterAddonInstancesRequest) Validate() error {
	return dara.Validate(s)
}
