// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetRelationFromAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromAddress(v string) *ConfigSetRelationFromAddressRequest
	GetFromAddress() *string
	SetId(v string) *ConfigSetRelationFromAddressRequest
	GetId() *string
}

type ConfigSetRelationFromAddressRequest struct {
	// example:
	//
	// xxx@xxx.com
	FromAddress *string `json:"FromAddress,omitempty" xml:"FromAddress,omitempty"`
	// example:
	//
	// xxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ConfigSetRelationFromAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetRelationFromAddressRequest) GoString() string {
	return s.String()
}

func (s *ConfigSetRelationFromAddressRequest) GetFromAddress() *string {
	return s.FromAddress
}

func (s *ConfigSetRelationFromAddressRequest) GetId() *string {
	return s.Id
}

func (s *ConfigSetRelationFromAddressRequest) SetFromAddress(v string) *ConfigSetRelationFromAddressRequest {
	s.FromAddress = &v
	return s
}

func (s *ConfigSetRelationFromAddressRequest) SetId(v string) *ConfigSetRelationFromAddressRequest {
	s.Id = &v
	return s
}

func (s *ConfigSetRelationFromAddressRequest) Validate() error {
	return dara.Validate(s)
}
