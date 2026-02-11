// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterFromGrafanaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListClusterFromGrafanaRequest
	GetRegionId() *string
}

type ListClusterFromGrafanaRequest struct {
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListClusterFromGrafanaRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClusterFromGrafanaRequest) GoString() string {
	return s.String()
}

func (s *ListClusterFromGrafanaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListClusterFromGrafanaRequest) SetRegionId(v string) *ListClusterFromGrafanaRequest {
	s.RegionId = &v
	return s
}

func (s *ListClusterFromGrafanaRequest) Validate() error {
	return dara.Validate(s)
}
