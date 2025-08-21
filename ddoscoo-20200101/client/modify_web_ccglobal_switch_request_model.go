// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebCCGlobalSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCcGlobalSwitch(v string) *ModifyWebCCGlobalSwitchRequest
	GetCcGlobalSwitch() *string
	SetDomain(v string) *ModifyWebCCGlobalSwitchRequest
	GetDomain() *string
}

type ModifyWebCCGlobalSwitchRequest struct {
	// Specifies whether the HTTP flood mitigation feature is enabled. Valid values:
	//
	// 	- **open**
	//
	// 	- **close**
	//
	// This parameter is required.
	//
	// example:
	//
	// open
	CcGlobalSwitch *string `json:"CcGlobalSwitch,omitempty" xml:"CcGlobalSwitch,omitempty"`
	// The domain name of the website.
	//
	// >  A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s ModifyWebCCGlobalSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebCCGlobalSwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebCCGlobalSwitchRequest) GetCcGlobalSwitch() *string {
	return s.CcGlobalSwitch
}

func (s *ModifyWebCCGlobalSwitchRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyWebCCGlobalSwitchRequest) SetCcGlobalSwitch(v string) *ModifyWebCCGlobalSwitchRequest {
	s.CcGlobalSwitch = &v
	return s
}

func (s *ModifyWebCCGlobalSwitchRequest) SetDomain(v string) *ModifyWebCCGlobalSwitchRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebCCGlobalSwitchRequest) Validate() error {
	return dara.Validate(s)
}
