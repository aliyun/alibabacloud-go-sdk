// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGrafanaWorkspaceAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v int64) *DeleteGrafanaWorkspaceAccountRequest
	GetAccountId() *int64
	SetGrafanaWorkspaceId(v string) *DeleteGrafanaWorkspaceAccountRequest
	GetGrafanaWorkspaceId() *string
	SetRegionId(v string) *DeleteGrafanaWorkspaceAccountRequest
	GetRegionId() *string
}

type DeleteGrafanaWorkspaceAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1781058524719073
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// grafana-cn-zky4h2wey01
	GrafanaWorkspaceId *string `json:"GrafanaWorkspaceId,omitempty" xml:"GrafanaWorkspaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteGrafanaWorkspaceAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGrafanaWorkspaceAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteGrafanaWorkspaceAccountRequest) GetAccountId() *int64 {
	return s.AccountId
}

func (s *DeleteGrafanaWorkspaceAccountRequest) GetGrafanaWorkspaceId() *string {
	return s.GrafanaWorkspaceId
}

func (s *DeleteGrafanaWorkspaceAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteGrafanaWorkspaceAccountRequest) SetAccountId(v int64) *DeleteGrafanaWorkspaceAccountRequest {
	s.AccountId = &v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountRequest) SetGrafanaWorkspaceId(v string) *DeleteGrafanaWorkspaceAccountRequest {
	s.GrafanaWorkspaceId = &v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountRequest) SetRegionId(v string) *DeleteGrafanaWorkspaceAccountRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountRequest) Validate() error {
	return dara.Validate(s)
}
