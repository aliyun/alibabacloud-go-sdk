// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackStagingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *RollbackStagingConfigRequest
	GetDomainName() *string
}

type RollbackStagingConfigRequest struct {
	// The accelerated domain name. You can specify only one domain name in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s RollbackStagingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *RollbackStagingConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *RollbackStagingConfigRequest) SetDomainName(v string) *RollbackStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *RollbackStagingConfigRequest) Validate() error {
	return dara.Validate(s)
}
