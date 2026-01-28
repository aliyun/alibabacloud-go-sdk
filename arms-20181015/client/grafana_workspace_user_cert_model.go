// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceUserCert interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GrafanaWorkspaceUserCert
	GetId() *string
	SetName(v string) *GrafanaWorkspaceUserCert
	GetName() *string
}

type GrafanaWorkspaceUserCert struct {
	// example:
	//
	// 8096753
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// cert-7700050 [grafana.tongtong-max.cn]
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GrafanaWorkspaceUserCert) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceUserCert) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceUserCert) GetId() *string {
	return s.Id
}

func (s *GrafanaWorkspaceUserCert) GetName() *string {
	return s.Name
}

func (s *GrafanaWorkspaceUserCert) SetId(v string) *GrafanaWorkspaceUserCert {
	s.Id = &v
	return s
}

func (s *GrafanaWorkspaceUserCert) SetName(v string) *GrafanaWorkspaceUserCert {
	s.Name = &v
	return s
}

func (s *GrafanaWorkspaceUserCert) Validate() error {
	return dara.Validate(s)
}
