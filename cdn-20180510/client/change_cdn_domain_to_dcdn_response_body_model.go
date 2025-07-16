// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeCdnDomainToDcdnResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *ChangeCdnDomainToDcdnResponseBody
	GetContent() map[string]interface{}
	SetRequestId(v string) *ChangeCdnDomainToDcdnResponseBody
	GetRequestId() *string
}

type ChangeCdnDomainToDcdnResponseBody struct {
	// The content of the migration instructions.
	//
	// example:
	//
	// {
	//
	// 		"The_domain_name_quota_for_the_target_account_has_reached_the_upper_limit": true,
	//
	// 		"Domain_name_requires_node2_architecture_to_be_enabled": true,
	//
	// 		"Internal_customer_domain_name_migration_prohibited": true,
	//
	// 		"Prohibit_the_migration_of_important_customer_domain_names": true,
	//
	// 		"Protected_domain_names_are_prohibited_from_migration": true,
	//
	// 		"Domain_names_accessed_through_TopDomain_are_prohibited_from_migration": true,
	//
	// 		"Prohibit_migration_of_customer_domain_names_for_integrated_access": true
	//
	// 	}
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E6172599-7DA7-5471-9D22-359A4E0C9B94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeCdnDomainToDcdnResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeCdnDomainToDcdnResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeCdnDomainToDcdnResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *ChangeCdnDomainToDcdnResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeCdnDomainToDcdnResponseBody) SetContent(v map[string]interface{}) *ChangeCdnDomainToDcdnResponseBody {
	s.Content = v
	return s
}

func (s *ChangeCdnDomainToDcdnResponseBody) SetRequestId(v string) *ChangeCdnDomainToDcdnResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeCdnDomainToDcdnResponseBody) Validate() error {
	return dara.Validate(s)
}
