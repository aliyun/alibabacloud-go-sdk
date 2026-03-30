// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultDomainName(v string) *SetDefaultDomainRequest
	GetDefaultDomainName() *string
}

type SetDefaultDomainRequest struct {
	// The default domain name.
	//
	// The default domain name is in the format of `<AccountAlias>.onaliyun.com`. `<AccountAlias>` indicates the account alias. By default, the value of AccountAlias is the ID of the Alibaba Cloud account. The default domain name must end with `.onaliyun.com`.
	//
	// The default domain name can contain up to 64 characters in length. The default domain name can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// >  The default domain name cannot start or end with a hyphen (-) and cannot contain two consecutive hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// examplecompany.onaliyun.com
	DefaultDomainName *string `json:"DefaultDomainName,omitempty" xml:"DefaultDomainName,omitempty"`
}

func (s SetDefaultDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultDomainRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultDomainRequest) GetDefaultDomainName() *string {
	return s.DefaultDomainName
}

func (s *SetDefaultDomainRequest) SetDefaultDomainName(v string) *SetDefaultDomainRequest {
	s.DefaultDomainName = &v
	return s
}

func (s *SetDefaultDomainRequest) Validate() error {
	return dara.Validate(s)
}
