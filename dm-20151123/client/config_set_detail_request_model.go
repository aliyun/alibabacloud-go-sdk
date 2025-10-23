// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ConfigSetDetailRequest
	GetId() *string
}

type ConfigSetDetailRequest struct {
	// example:
	//
	// xxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ConfigSetDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetDetailRequest) GoString() string {
	return s.String()
}

func (s *ConfigSetDetailRequest) GetId() *string {
	return s.Id
}

func (s *ConfigSetDetailRequest) SetId(v string) *ConfigSetDetailRequest {
	s.Id = &v
	return s
}

func (s *ConfigSetDetailRequest) Validate() error {
	return dara.Validate(s)
}
