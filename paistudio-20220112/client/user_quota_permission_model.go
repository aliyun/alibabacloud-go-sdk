// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserQuotaPermission interface {
	dara.Model
	String() string
	GoString() string
	SetPermissions(v []*string) *UserQuotaPermission
	GetPermissions() []*string
	SetQuotaId(v string) *UserQuotaPermission
	GetQuotaId() *string
}

type UserQuotaPermission struct {
	Permissions []*string `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	QuotaId     *string   `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
}

func (s UserQuotaPermission) String() string {
	return dara.Prettify(s)
}

func (s UserQuotaPermission) GoString() string {
	return s.String()
}

func (s *UserQuotaPermission) GetPermissions() []*string {
	return s.Permissions
}

func (s *UserQuotaPermission) GetQuotaId() *string {
	return s.QuotaId
}

func (s *UserQuotaPermission) SetPermissions(v []*string) *UserQuotaPermission {
	s.Permissions = v
	return s
}

func (s *UserQuotaPermission) SetQuotaId(v string) *UserQuotaPermission {
	s.QuotaId = &v
	return s
}

func (s *UserQuotaPermission) Validate() error {
	return dara.Validate(s)
}
