// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthenticationTokensResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntities(v []*ListAuthenticationTokensResponseBodyEntities) *ListAuthenticationTokensResponseBody
	GetEntities() []*ListAuthenticationTokensResponseBodyEntities
	SetMaxResults(v int64) *ListAuthenticationTokensResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListAuthenticationTokensResponseBody
	GetNextToken() *string
	SetTotalCount(v int64) *ListAuthenticationTokensResponseBody
	GetTotalCount() *int64
}

type ListAuthenticationTokensResponseBody struct {
	// 资源实体列表。
	Entities []*ListAuthenticationTokensResponseBodyEntities `json:"entities,omitempty" xml:"entities,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListAuthenticationTokensResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthenticationTokensResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthenticationTokensResponseBody) GetEntities() []*ListAuthenticationTokensResponseBodyEntities {
	return s.Entities
}

func (s *ListAuthenticationTokensResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListAuthenticationTokensResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAuthenticationTokensResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAuthenticationTokensResponseBody) SetEntities(v []*ListAuthenticationTokensResponseBodyEntities) *ListAuthenticationTokensResponseBody {
	s.Entities = v
	return s
}

func (s *ListAuthenticationTokensResponseBody) SetMaxResults(v int64) *ListAuthenticationTokensResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAuthenticationTokensResponseBody) SetNextToken(v string) *ListAuthenticationTokensResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAuthenticationTokensResponseBody) SetTotalCount(v int64) *ListAuthenticationTokensResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAuthenticationTokensResponseBody) Validate() error {
	if s.Entities != nil {
		for _, item := range s.Entities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAuthenticationTokensResponseBodyEntities struct {
	// example:
	//
	// atntkn_01kqflm0sxxx8nmdc1cb5dskxxxxx
	AuthenticationTokenId *string `json:"authenticationTokenId,omitempty" xml:"authenticationTokenId,omitempty"`
	// example:
	//
	// jwt
	AuthenticationTokenType *string `json:"authenticationTokenType,omitempty" xml:"authenticationTokenType,omitempty"`
	// example:
	//
	// test_jwt_subject
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// example:
	//
	// custom
	ConsumerType *string `json:"consumerType,omitempty" xml:"consumerType,omitempty"`
	// example:
	//
	// 1649830225000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// app_ngtkgrrxxxxktg5eao6z4xxxxx
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// application
	CreatorType *string `json:"creatorType,omitempty" xml:"creatorType,omitempty"`
	// example:
	//
	// atp_01kr2cmj5gxxx4fvmls2e93dxxxxx
	CredentialProviderId *string `json:"credentialProviderId,omitempty" xml:"credentialProviderId,omitempty"`
	// example:
	//
	// 1772693568000
	ExpirationTime *int64 `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
	// EIAM实例ID。
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// false
	Revoked *bool `json:"revoked,omitempty" xml:"revoked,omitempty"`
	// example:
	//
	// 1649830225000
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListAuthenticationTokensResponseBodyEntities) String() string {
	return dara.Prettify(s)
}

func (s ListAuthenticationTokensResponseBodyEntities) GoString() string {
	return s.String()
}

func (s *ListAuthenticationTokensResponseBodyEntities) GetAuthenticationTokenId() *string {
	return s.AuthenticationTokenId
}

func (s *ListAuthenticationTokensResponseBodyEntities) GetAuthenticationTokenType() *string {
	return s.AuthenticationTokenType
}

func (s *ListAuthenticationTokensResponseBodyEntities) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *ListAuthenticationTokensResponseBodyEntities) GetConsumerType() *string {
	return s.ConsumerType
}

func (s *ListAuthenticationTokensResponseBodyEntities) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListAuthenticationTokensResponseBodyEntities) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListAuthenticationTokensResponseBodyEntities) GetCreatorType() *string {
	return s.CreatorType
}

func (s *ListAuthenticationTokensResponseBodyEntities) GetCredentialProviderId() *string {
	return s.CredentialProviderId
}

func (s *ListAuthenticationTokensResponseBodyEntities) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *ListAuthenticationTokensResponseBodyEntities) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAuthenticationTokensResponseBodyEntities) GetRevoked() *bool {
	return s.Revoked
}

func (s *ListAuthenticationTokensResponseBodyEntities) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListAuthenticationTokensResponseBodyEntities) SetAuthenticationTokenId(v string) *ListAuthenticationTokensResponseBodyEntities {
	s.AuthenticationTokenId = &v
	return s
}

func (s *ListAuthenticationTokensResponseBodyEntities) SetAuthenticationTokenType(v string) *ListAuthenticationTokensResponseBodyEntities {
	s.AuthenticationTokenType = &v
	return s
}

func (s *ListAuthenticationTokensResponseBodyEntities) SetConsumerId(v string) *ListAuthenticationTokensResponseBodyEntities {
	s.ConsumerId = &v
	return s
}

func (s *ListAuthenticationTokensResponseBodyEntities) SetConsumerType(v string) *ListAuthenticationTokensResponseBodyEntities {
	s.ConsumerType = &v
	return s
}

func (s *ListAuthenticationTokensResponseBodyEntities) SetCreateTime(v int64) *ListAuthenticationTokensResponseBodyEntities {
	s.CreateTime = &v
	return s
}

func (s *ListAuthenticationTokensResponseBodyEntities) SetCreatorId(v string) *ListAuthenticationTokensResponseBodyEntities {
	s.CreatorId = &v
	return s
}

func (s *ListAuthenticationTokensResponseBodyEntities) SetCreatorType(v string) *ListAuthenticationTokensResponseBodyEntities {
	s.CreatorType = &v
	return s
}

func (s *ListAuthenticationTokensResponseBodyEntities) SetCredentialProviderId(v string) *ListAuthenticationTokensResponseBodyEntities {
	s.CredentialProviderId = &v
	return s
}

func (s *ListAuthenticationTokensResponseBodyEntities) SetExpirationTime(v int64) *ListAuthenticationTokensResponseBodyEntities {
	s.ExpirationTime = &v
	return s
}

func (s *ListAuthenticationTokensResponseBodyEntities) SetInstanceId(v string) *ListAuthenticationTokensResponseBodyEntities {
	s.InstanceId = &v
	return s
}

func (s *ListAuthenticationTokensResponseBodyEntities) SetRevoked(v bool) *ListAuthenticationTokensResponseBodyEntities {
	s.Revoked = &v
	return s
}

func (s *ListAuthenticationTokensResponseBodyEntities) SetUpdateTime(v int64) *ListAuthenticationTokensResponseBodyEntities {
	s.UpdateTime = &v
	return s
}

func (s *ListAuthenticationTokensResponseBodyEntities) Validate() error {
	return dara.Validate(s)
}
