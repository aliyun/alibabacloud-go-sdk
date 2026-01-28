// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceOrg interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GrafanaWorkspaceOrg
	GetId() *int64
	SetName(v string) *GrafanaWorkspaceOrg
	GetName() *string
}

type GrafanaWorkspaceOrg struct {
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// main org
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GrafanaWorkspaceOrg) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceOrg) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceOrg) GetId() *int64 {
	return s.Id
}

func (s *GrafanaWorkspaceOrg) GetName() *string {
	return s.Name
}

func (s *GrafanaWorkspaceOrg) SetId(v int64) *GrafanaWorkspaceOrg {
	s.Id = &v
	return s
}

func (s *GrafanaWorkspaceOrg) SetName(v string) *GrafanaWorkspaceOrg {
	s.Name = &v
	return s
}

func (s *GrafanaWorkspaceOrg) Validate() error {
	return dara.Validate(s)
}
