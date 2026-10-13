// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDSEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListDSEntityRequest
	GetAgentKey() *string
	SetEntityType(v string) *ListDSEntityRequest
	GetEntityType() *string
	SetInstanceId(v string) *ListDSEntityRequest
	GetInstanceId() *string
	SetKeyword(v string) *ListDSEntityRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListDSEntityRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDSEntityRequest
	GetPageSize() *int32
}

type ListDSEntityRequest struct {
	// The business space key. If this parameter is not specified, the default business space is accessed. You can obtain the key value on the business management page of the primary account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// If left empty, all custom entities are returned (default).
	//
	// example:
	//
	// synonyms
	//
	// regex
	//
	// system
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The chatbot ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The filter keyword. Uses contains matching. Scope: entity name (future expansion: entity members, synonyms).
	//
	// example:
	//
	// entity
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDSEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDSEntityRequest) GoString() string {
	return s.String()
}

func (s *ListDSEntityRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListDSEntityRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *ListDSEntityRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDSEntityRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListDSEntityRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDSEntityRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDSEntityRequest) SetAgentKey(v string) *ListDSEntityRequest {
	s.AgentKey = &v
	return s
}

func (s *ListDSEntityRequest) SetEntityType(v string) *ListDSEntityRequest {
	s.EntityType = &v
	return s
}

func (s *ListDSEntityRequest) SetInstanceId(v string) *ListDSEntityRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDSEntityRequest) SetKeyword(v string) *ListDSEntityRequest {
	s.Keyword = &v
	return s
}

func (s *ListDSEntityRequest) SetPageNumber(v int32) *ListDSEntityRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDSEntityRequest) SetPageSize(v int32) *ListDSEntityRequest {
	s.PageSize = &v
	return s
}

func (s *ListDSEntityRequest) Validate() error {
	return dara.Validate(s)
}
