// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApplicationResourceServerIdentifierRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *SetApplicationResourceServerIdentifierRequest
	GetApplicationId() *string
	SetClientToken(v string) *SetApplicationResourceServerIdentifierRequest
	GetClientToken() *string
	SetInstanceId(v string) *SetApplicationResourceServerIdentifierRequest
	GetInstanceId() *string
	SetResourceServerIdentifier(v string) *SetApplicationResourceServerIdentifierRequest
	GetResourceServerIdentifier() *string
}

type SetApplicationResourceServerIdentifierRequest struct {
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// ResourceServer唯一标识，对应ResourceServer受众
	//
	// This parameter is required.
	//
	// example:
	//
	// http://gateway.com
	ResourceServerIdentifier *string `json:"ResourceServerIdentifier,omitempty" xml:"ResourceServerIdentifier,omitempty"`
}

func (s SetApplicationResourceServerIdentifierRequest) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationResourceServerIdentifierRequest) GoString() string {
	return s.String()
}

func (s *SetApplicationResourceServerIdentifierRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *SetApplicationResourceServerIdentifierRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SetApplicationResourceServerIdentifierRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetApplicationResourceServerIdentifierRequest) GetResourceServerIdentifier() *string {
	return s.ResourceServerIdentifier
}

func (s *SetApplicationResourceServerIdentifierRequest) SetApplicationId(v string) *SetApplicationResourceServerIdentifierRequest {
	s.ApplicationId = &v
	return s
}

func (s *SetApplicationResourceServerIdentifierRequest) SetClientToken(v string) *SetApplicationResourceServerIdentifierRequest {
	s.ClientToken = &v
	return s
}

func (s *SetApplicationResourceServerIdentifierRequest) SetInstanceId(v string) *SetApplicationResourceServerIdentifierRequest {
	s.InstanceId = &v
	return s
}

func (s *SetApplicationResourceServerIdentifierRequest) SetResourceServerIdentifier(v string) *SetApplicationResourceServerIdentifierRequest {
	s.ResourceServerIdentifier = &v
	return s
}

func (s *SetApplicationResourceServerIdentifierRequest) Validate() error {
	return dara.Validate(s)
}
