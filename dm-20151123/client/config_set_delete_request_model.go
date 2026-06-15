// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v string) *ConfigSetDeleteRequest
	GetIds() *string
	SetIsForce(v bool) *ConfigSetDeleteRequest
	GetIsForce() *bool
}

type ConfigSetDeleteRequest struct {
	// The IDs of the configuration sets. Separate multiple IDs with commas. This parameter is required.
	//
	// example:
	//
	// xxx
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// Specifies whether to force delete the associations with sender addresses.
	//
	// example:
	//
	// true
	IsForce *bool `json:"IsForce,omitempty" xml:"IsForce,omitempty"`
}

func (s ConfigSetDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetDeleteRequest) GoString() string {
	return s.String()
}

func (s *ConfigSetDeleteRequest) GetIds() *string {
	return s.Ids
}

func (s *ConfigSetDeleteRequest) GetIsForce() *bool {
	return s.IsForce
}

func (s *ConfigSetDeleteRequest) SetIds(v string) *ConfigSetDeleteRequest {
	s.Ids = &v
	return s
}

func (s *ConfigSetDeleteRequest) SetIsForce(v bool) *ConfigSetDeleteRequest {
	s.IsForce = &v
	return s
}

func (s *ConfigSetDeleteRequest) Validate() error {
	return dara.Validate(s)
}
