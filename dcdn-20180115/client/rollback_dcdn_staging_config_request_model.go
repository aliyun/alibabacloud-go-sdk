// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackDcdnStagingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *RollbackDcdnStagingConfigRequest
	GetDomainName() *string
}

type RollbackDcdnStagingConfigRequest struct {
	// The accelerated domain name. You can specify only one domain name in each call.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s RollbackDcdnStagingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackDcdnStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *RollbackDcdnStagingConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *RollbackDcdnStagingConfigRequest) SetDomainName(v string) *RollbackDcdnStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *RollbackDcdnStagingConfigRequest) Validate() error {
	return dara.Validate(s)
}
