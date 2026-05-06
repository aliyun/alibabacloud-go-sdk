// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelProviderTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProviderTemplateId(v string) *DeleteModelProviderTemplateRequest
	GetProviderTemplateId() *string
}

type DeleteModelProviderTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mpt-xxxx
	ProviderTemplateId *string `json:"ProviderTemplateId,omitempty" xml:"ProviderTemplateId,omitempty"`
}

func (s DeleteModelProviderTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelProviderTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelProviderTemplateRequest) GetProviderTemplateId() *string {
	return s.ProviderTemplateId
}

func (s *DeleteModelProviderTemplateRequest) SetProviderTemplateId(v string) *DeleteModelProviderTemplateRequest {
	s.ProviderTemplateId = &v
	return s
}

func (s *DeleteModelProviderTemplateRequest) Validate() error {
	return dara.Validate(s)
}
