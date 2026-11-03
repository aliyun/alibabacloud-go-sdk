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
	// The language of the request and response.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the Internet NAT gateway instance to be associated. This parameter is required. If this parameter is not specified, the ErrorParamsNotEnough error (HTTP 400, Parameters are insufficient.) is returned.
	//
	// > The backend does not verify the ID format. Instead, it queries the instance in the private network asset table of Cloud Firewall under the current account. If the instance is not found, the ErrorParamsInvalid error (HTTP 400, Invalid Params) is returned. Common causes include: the resource type is not a NAT gateway, the resource is not managed by Cloud Firewall, or the asynchronous asset synchronization for a newly created NAT gateway is not completed.
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
