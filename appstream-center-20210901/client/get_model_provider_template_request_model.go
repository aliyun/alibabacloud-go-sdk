// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelProviderTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProviderTemplateId(v string) *GetModelProviderTemplateRequest
	GetProviderTemplateId() *string
}

type GetModelProviderTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mpt-xxxx
	ProviderTemplateId *string `json:"ProviderTemplateId,omitempty" xml:"ProviderTemplateId,omitempty"`
}

func (s GetModelProviderTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModelProviderTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetModelProviderTemplateRequest) GetProviderTemplateId() *string {
	return s.ProviderTemplateId
}

func (s *GetModelProviderTemplateRequest) SetProviderTemplateId(v string) *GetModelProviderTemplateRequest {
	s.ProviderTemplateId = &v
	return s
}

func (s *GetModelProviderTemplateRequest) Validate() error {
	return dara.Validate(s)
}
