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
	// The language of the request and response messages.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the Internet NAT gateway instance to associate. This parameter is required. If this parameter is not specified, ErrorParamsNotEnough is returned (HTTP 400, Parameters are insufficient.).
	//
	// > The backend does not validate the ID format. Instead, it queries the instance in the Cloud Firewall private network asset table for the current account. If the instance is not found, ErrorParamsInvalid is returned (HTTP 400, Invalid Params). Common scenarios include the resource type not being a NAT gateway, the resource not being managed by Cloud Firewall, or a newly created NAT gateway for which asynchronous asset synchronization has not yet completed.
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
