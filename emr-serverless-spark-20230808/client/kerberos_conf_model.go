// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKerberosConf interface {
	dara.Model
	String() string
	GoString() string
	SetCreator(v string) *KerberosConf
	GetCreator() *string
	SetEnabled(v bool) *KerberosConf
	GetEnabled() *bool
	SetGmtCreated(v string) *KerberosConf
	GetGmtCreated() *string
	SetGmtModified(v string) *KerberosConf
	GetGmtModified() *string
	SetKerberosConfId(v string) *KerberosConf
	GetKerberosConfId() *string
	SetKeytabs(v []*string) *KerberosConf
	GetKeytabs() []*string
	SetKrb5Conf(v string) *KerberosConf
	GetKrb5Conf() *string
	SetName(v string) *KerberosConf
	GetName() *string
	SetNetworkServiceId(v string) *KerberosConf
	GetNetworkServiceId() *string
	SetWorkspaceId(v string) *KerberosConf
	GetWorkspaceId() *string
}

type KerberosConf struct {
	Creator          *string   `json:"creator,omitempty" xml:"creator,omitempty"`
	Enabled          *bool     `json:"enabled,omitempty" xml:"enabled,omitempty"`
	GmtCreated       *string   `json:"gmtCreated,omitempty" xml:"gmtCreated,omitempty"`
	GmtModified      *string   `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	KerberosConfId   *string   `json:"kerberosConfId,omitempty" xml:"kerberosConfId,omitempty"`
	Keytabs          []*string `json:"keytabs,omitempty" xml:"keytabs,omitempty" type:"Repeated"`
	Krb5Conf         *string   `json:"krb5Conf,omitempty" xml:"krb5Conf,omitempty"`
	Name             *string   `json:"name,omitempty" xml:"name,omitempty"`
	NetworkServiceId *string   `json:"networkServiceId,omitempty" xml:"networkServiceId,omitempty"`
	WorkspaceId      *string   `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s KerberosConf) String() string {
	return dara.Prettify(s)
}

func (s KerberosConf) GoString() string {
	return s.String()
}

func (s *KerberosConf) GetCreator() *string {
	return s.Creator
}

func (s *KerberosConf) GetEnabled() *bool {
	return s.Enabled
}

func (s *KerberosConf) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *KerberosConf) GetGmtModified() *string {
	return s.GmtModified
}

func (s *KerberosConf) GetKerberosConfId() *string {
	return s.KerberosConfId
}

func (s *KerberosConf) GetKeytabs() []*string {
	return s.Keytabs
}

func (s *KerberosConf) GetKrb5Conf() *string {
	return s.Krb5Conf
}

func (s *KerberosConf) GetName() *string {
	return s.Name
}

func (s *KerberosConf) GetNetworkServiceId() *string {
	return s.NetworkServiceId
}

func (s *KerberosConf) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *KerberosConf) SetCreator(v string) *KerberosConf {
	s.Creator = &v
	return s
}

func (s *KerberosConf) SetEnabled(v bool) *KerberosConf {
	s.Enabled = &v
	return s
}

func (s *KerberosConf) SetGmtCreated(v string) *KerberosConf {
	s.GmtCreated = &v
	return s
}

func (s *KerberosConf) SetGmtModified(v string) *KerberosConf {
	s.GmtModified = &v
	return s
}

func (s *KerberosConf) SetKerberosConfId(v string) *KerberosConf {
	s.KerberosConfId = &v
	return s
}

func (s *KerberosConf) SetKeytabs(v []*string) *KerberosConf {
	s.Keytabs = v
	return s
}

func (s *KerberosConf) SetKrb5Conf(v string) *KerberosConf {
	s.Krb5Conf = &v
	return s
}

func (s *KerberosConf) SetName(v string) *KerberosConf {
	s.Name = &v
	return s
}

func (s *KerberosConf) SetNetworkServiceId(v string) *KerberosConf {
	s.NetworkServiceId = &v
	return s
}

func (s *KerberosConf) SetWorkspaceId(v string) *KerberosConf {
	s.WorkspaceId = &v
	return s
}

func (s *KerberosConf) Validate() error {
	return dara.Validate(s)
}
