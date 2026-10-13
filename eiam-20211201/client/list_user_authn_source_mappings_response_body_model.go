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
	// The maximum number of rows per page in a paged query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token returned by this call.
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The pagination token returned by this call.
	//
	// example:
	//
	// PTxxxexample
	PreviousToken *string `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries in the list.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of third-party account login bindings.
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
	// The authentication source type of the source IdP.
	//
	// example:
	//
	// urn:alibaba:idaas:authntype:oidc
	AuthnSourceType *string `json:"AuthnSourceType,omitempty" xml:"AuthnSourceType,omitempty"`
	// The creation time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1762309642177
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The association information of the third-party account.
	//
	// example:
	//
	// {\\"userId\\":\\"xxxx\\",\\"name\\":\\"xxx\\",\\"bindTime\\":\\"1766050298872\\",\\"description\\":\\"bind request id: reqpre_xxx\\"}
	ExternalData *string `json:"ExternalData,omitempty" xml:"ExternalData,omitempty"`
	// The source IdP ID.
	//
	// example:
	//
	// idp_my664lwkhpicbyzirog3xxxxx
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The last update time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1762309642177
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The external ID.
	//
	// example:
	//
	// user_123131xxasasa
	UserExternalId *string `json:"UserExternalId,omitempty" xml:"UserExternalId,omitempty"`
	// The user ID.
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

func (s *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings) GetExternalData() *string {
	return s.ExternalData
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

func (s *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings) SetExternalData(v string) *ListUserAuthnSourceMappingsResponseBodyUserAuthnSourceMappings {
	s.ExternalData = &v
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
