// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWuyingServerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPassword(v string) *ModifyWuyingServerAttributeRequest
	GetPassword() *string
	SetProductType(v string) *ModifyWuyingServerAttributeRequest
	GetProductType() *string
	SetWuyingServerId(v string) *ModifyWuyingServerAttributeRequest
	GetWuyingServerId() *string
	SetWuyingServerName(v string) *ModifyWuyingServerAttributeRequest
	GetWuyingServerName() *string
}

type ModifyWuyingServerAttributeRequest struct {
	// Workstation login password.
	//
	// example:
	//
	// yourPassword
	Password    *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The ID of the workstation.
	//
	// example:
	//
	// ws-0bw2f11****dial
	WuyingServerId *string `json:"WuyingServerId,omitempty" xml:"WuyingServerId,omitempty"`
	// The name.
	//
	// example:
	//
	// exampleServerName
	WuyingServerName *string `json:"WuyingServerName,omitempty" xml:"WuyingServerName,omitempty"`
}

func (s ModifyWuyingServerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWuyingServerAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyWuyingServerAttributeRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyWuyingServerAttributeRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ModifyWuyingServerAttributeRequest) GetWuyingServerId() *string {
	return s.WuyingServerId
}

func (s *ModifyWuyingServerAttributeRequest) GetWuyingServerName() *string {
	return s.WuyingServerName
}

func (s *ModifyWuyingServerAttributeRequest) SetPassword(v string) *ModifyWuyingServerAttributeRequest {
	s.Password = &v
	return s
}

func (s *ModifyWuyingServerAttributeRequest) SetProductType(v string) *ModifyWuyingServerAttributeRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyWuyingServerAttributeRequest) SetWuyingServerId(v string) *ModifyWuyingServerAttributeRequest {
	s.WuyingServerId = &v
	return s
}

func (s *ModifyWuyingServerAttributeRequest) SetWuyingServerName(v string) *ModifyWuyingServerAttributeRequest {
	s.WuyingServerName = &v
	return s
}

func (s *ModifyWuyingServerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
