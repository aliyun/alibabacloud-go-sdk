// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManagedRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifier(v string) *GetManagedRuleRequest
	GetIdentifier() *string
}

type GetManagedRuleRequest struct {
	// The identifier of the managed rule.
	//
	// For more information about how to obtain the identifier of a managed rule, see [ListManagedRules](https://help.aliyun.com/document_detail/421144.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cdn-domain-https-enabled
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
}

func (s GetManagedRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetManagedRuleRequest) GoString() string {
	return s.String()
}

func (s *GetManagedRuleRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetManagedRuleRequest) SetIdentifier(v string) *GetManagedRuleRequest {
	s.Identifier = &v
	return s
}

func (s *GetManagedRuleRequest) Validate() error {
	return dara.Validate(s)
}
