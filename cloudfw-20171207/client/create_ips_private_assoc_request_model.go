// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpsPrivateAssocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateIpsPrivateAssocRequest
	GetLang() *string
	SetResourceId(v string) *CreateIpsPrivateAssocRequest
	GetResourceId() *string
}

type CreateIpsPrivateAssocRequest struct {
	// The language type for the request and response messages. Valid values:
	//
	// - en: English.
	//
	// - zh: Chinese.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The instance ID. This parameter is required. If this parameter is not specified, the API returns error code -103201. Only NAT gateway instance IDs (in the format ngw-*) that are protected by Cloud Firewall are accepted. Other resource types such as vpc-	- or eip-	- are rejected.
	//
	// example:
	//
	// ngw-c5vhmjdfp5t****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s CreateIpsPrivateAssocRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIpsPrivateAssocRequest) GoString() string {
	return s.String()
}

func (s *CreateIpsPrivateAssocRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateIpsPrivateAssocRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateIpsPrivateAssocRequest) SetLang(v string) *CreateIpsPrivateAssocRequest {
	s.Lang = &v
	return s
}

func (s *CreateIpsPrivateAssocRequest) SetResourceId(v string) *CreateIpsPrivateAssocRequest {
	s.ResourceId = &v
	return s
}

func (s *CreateIpsPrivateAssocRequest) Validate() error {
	return dara.Validate(s)
}
