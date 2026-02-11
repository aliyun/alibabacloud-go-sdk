// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTraceAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListTraceAppsRequest
	GetRegionId() *string
}

type ListTraceAppsRequest struct {
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListTraceAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTraceAppsRequest) GoString() string {
	return s.String()
}

func (s *ListTraceAppsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTraceAppsRequest) SetRegionId(v string) *ListTraceAppsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTraceAppsRequest) Validate() error {
	return dara.Validate(s)
}
