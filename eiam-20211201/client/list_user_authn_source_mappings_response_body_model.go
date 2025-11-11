// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserAuthnSourceMappingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListUserAuthnSourceMappingsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListUserAuthnSourceMappingsResponseBody
	GetNextToken() *string
	SetPreviousToken(v string) *ListUserAuthnSourceMappingsResponseBody
	GetPreviousToken() *string
	SetRequestId(v string) *ListUserAuthnSourceMappingsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListUserAuthnSourceMappingsResponseBody
	GetTotalCount() *int32
	SetUserAuthnSourceMappings(v []*ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings) *ListUserAuthnSourceMappingsResponseBody
	GetUserAuthnSourceMappings() []*ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings
}

type ListUserAuthnSourceMappingsResponseBody struct {
	// 分页查询时每页行数。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 本次调用返回的查询凭证（Token）值，用于下一次翻页查询。
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次调用返回的查询凭证（Token）值，用于上一次翻页查询。
	//
	// example:
	//
	// PTxxxexample
	PreviousToken *string `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount              *int32                                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UserAuthnSourceMappings []*ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings `json:"UserAuthnSourceMappings,omitempty" xml:"UserAuthnSourceMappings,omitempty" type:"Repeated"`
}

func (s ListUserAuthnSourceMappingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserAuthnSourceMappingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserAuthnSourceMappingsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUserAuthnSourceMappingsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserAuthnSourceMappingsResponseBody) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListUserAuthnSourceMappingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserAuthnSourceMappingsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUserAuthnSourceMappingsResponseBody) GetUserAuthnSourceMappings() []*ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings {
	return s.UserAuthnSourceMappings
}

func (s *ListUserAuthnSourceMappingsResponseBody) SetMaxResults(v int32) *ListUserAuthnSourceMappingsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUserAuthnSourceMappingsResponseBody) SetNextToken(v string) *ListUserAuthnSourceMappingsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserAuthnSourceMappingsResponseBody) SetPreviousToken(v string) *ListUserAuthnSourceMappingsResponseBody {
	s.PreviousToken = &v
	return s
}

func (s *ListUserAuthnSourceMappingsResponseBody) SetRequestId(v string) *ListUserAuthnSourceMappingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserAuthnSourceMappingsResponseBody) SetTotalCount(v int32) *ListUserAuthnSourceMappingsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserAuthnSourceMappingsResponseBody) SetUserAuthnSourceMappings(v []*ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings) *ListUserAuthnSourceMappingsResponseBody {
	s.UserAuthnSourceMappings = v
	return s
}

func (s *ListUserAuthnSourceMappingsResponseBody) Validate() error {
	if s.UserAuthnSourceMappings != nil {
		for _, item := range s.UserAuthnSourceMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings struct {
	// 来源Idp类型
	//
	// example:
	//
	// urn:alibaba:idaas:authntype:oidc
	AuthnSourceType *string `json:"AuthnSourceType,omitempty" xml:"AuthnSourceType,omitempty"`
	// 创建时间
	//
	// example:
	//
	// 1762309642177
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 来源Idp Id
	//
	// example:
	//
	// idp_my664lwkhpicbyzirog3xxxxx
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// 实例Id
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 最近一次更新时间
	//
	// example:
	//
	// 1762309642177
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 外部ID
	//
	// example:
	//
	// xxxxxx
	UserExternalId *string `json:"UserExternalId,omitempty" xml:"UserExternalId,omitempty"`
	// 用户ID
	//
	// example:
	//
	// user_ue2jvisn35exxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings) String() string {
	return dara.Prettify(s)
}

func (s ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings) GoString() string {
	return s.String()
}

func (s *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings) GetAuthnSourceType() *string {
	return s.AuthnSourceType
}

func (s *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings) GetUserExternalId() *string {
	return s.UserExternalId
}

func (s *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings) GetUserId() *string {
	return s.UserId
}

func (s *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings) SetAuthnSourceType(v string) *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings {
	s.AuthnSourceType = &v
	return s
}

func (s *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings) SetCreateTime(v int64) *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings {
	s.CreateTime = &v
	return s
}

func (s *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings) SetIdentityProviderId(v string) *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings {
	s.IdentityProviderId = &v
	return s
}

func (s *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings) SetInstanceId(v string) *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings {
	s.InstanceId = &v
	return s
}

func (s *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings) SetUpdateTime(v int64) *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings {
	s.UpdateTime = &v
	return s
}

func (s *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings) SetUserExternalId(v string) *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings {
	s.UserExternalId = &v
	return s
}

func (s *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings) SetUserId(v string) *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings {
	s.UserId = &v
	return s
}

func (s *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings) Validate() error {
	return dara.Validate(s)
}
