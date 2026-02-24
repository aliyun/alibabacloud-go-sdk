// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRemediationTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateIdentifier(v string) *GetRemediationTemplateRequest
	GetTemplateIdentifier() *string
}

type GetRemediationTemplateRequest struct {
	// The identifier of the automatic remediation template.
	//
	// For information about how to obtain the remediation template identifier, see [Compliance library](https://help.aliyun.com/document_detail/2337741.html).
	//
	// if can be null:
	// true
	//
	// example:
	//
	// ACS-ALB-BulkyEnableDeletionProtection
	TemplateIdentifier *string `json:"TemplateIdentifier,omitempty" xml:"TemplateIdentifier,omitempty"`
}

func (s GetRemediationTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRemediationTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetRemediationTemplateRequest) GetTemplateIdentifier() *string {
	return s.TemplateIdentifier
}

func (s *GetRemediationTemplateRequest) SetTemplateIdentifier(v string) *GetRemediationTemplateRequest {
	s.TemplateIdentifier = &v
	return s
}

func (s *GetRemediationTemplateRequest) Validate() error {
	return dara.Validate(s)
}
