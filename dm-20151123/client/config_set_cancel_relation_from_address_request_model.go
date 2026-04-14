// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetCancelRelationFromAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromAddress(v string) *ConfigSetCancelRelationFromAddressRequest
	GetFromAddress() *string
	SetId(v string) *ConfigSetCancelRelationFromAddressRequest
	GetId() *string
}

type ConfigSetCancelRelationFromAddressRequest struct {
	// example:
	//
	// xxx@xxx.com
	FromAddress *string `json:"FromAddress,omitempty" xml:"FromAddress,omitempty"`
	// example:
	//
	// xxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ConfigSetCancelRelationFromAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetCancelRelationFromAddressRequest) GoString() string {
	return s.String()
}

func (s *ConfigSetCancelRelationFromAddressRequest) GetFromAddress() *string {
	return s.FromAddress
}

func (s *ConfigSetCancelRelationFromAddressRequest) GetId() *string {
	return s.Id
}

func (s *ConfigSetCancelRelationFromAddressRequest) SetFromAddress(v string) *ConfigSetCancelRelationFromAddressRequest {
	s.FromAddress = &v
	return s
}

func (s *ConfigSetCancelRelationFromAddressRequest) SetId(v string) *ConfigSetCancelRelationFromAddressRequest {
	s.Id = &v
	return s
}

func (s *ConfigSetCancelRelationFromAddressRequest) Validate() error {
	return dara.Validate(s)
}
