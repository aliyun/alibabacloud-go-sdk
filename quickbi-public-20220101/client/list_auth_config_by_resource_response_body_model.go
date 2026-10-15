// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthConfigByResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAuthConfigByResourceResponseBody
	GetRequestId() *string
	SetResult(v []*ListAuthConfigByResourceResponseBodyResult) *ListAuthConfigByResourceResponseBody
	GetResult() []*ListAuthConfigByResourceResponseBodyResult
	SetSuccess(v bool) *ListAuthConfigByResourceResponseBody
	GetSuccess() *bool
}

type ListAuthConfigByResourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 019FA14C-ED6B-5F39-8563-855676470C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The collection of authorization details.
	Result []*ListAuthConfigByResourceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAuthConfigByResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthConfigByResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthConfigByResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthConfigByResourceResponseBody) GetResult() []*ListAuthConfigByResourceResponseBodyResult {
	return s.Result
}

func (s *ListAuthConfigByResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAuthConfigByResourceResponseBody) SetRequestId(v string) *ListAuthConfigByResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthConfigByResourceResponseBody) SetResult(v []*ListAuthConfigByResourceResponseBodyResult) *ListAuthConfigByResourceResponseBody {
	s.Result = v
	return s
}

func (s *ListAuthConfigByResourceResponseBody) SetSuccess(v bool) *ListAuthConfigByResourceResponseBody {
	s.Success = &v
	return s
}

func (s *ListAuthConfigByResourceResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAuthConfigByResourceResponseBodyResult struct {
	// The array of permissions granted.
	Authorities []*string `json:"Authorities,omitempty" xml:"Authorities,omitempty" type:"Repeated"`
	// The ID of the authorized user.
	//
	// example:
	//
	// 123123**dfsd
	AuthorizedId *string `json:"AuthorizedId,omitempty" xml:"AuthorizedId,omitempty"`
	// The name of the authorized user.
	//
	// example:
	//
	// sss
	AuthorizedName *string `json:"AuthorizedName,omitempty" xml:"AuthorizedName,omitempty"`
	// The expiration time. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1788969600000
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The time when the permission record was created, in the yyyy-MM-dd mm:hh:ss format.
	//
	// example:
	//
	// 2024-04-16 13:17:39
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// 123123**dfsd
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. Valid values:
	//
	// - dashboard: dashboard
	//
	// - screen: data screen
	//
	// - report: electron
	//
	// - ANALYSIS: ad hoc analysis
	//
	// - dashboardOfflineQuery: self-service data retrieval
	//
	// - dataForm: data form
	//
	// - cube: dataset
	//
	// - datasource: data source
	//
	// example:
	//
	// cube
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The authorization scope. Valid values:
	//
	// - 0: user
	//
	// - 1: user group
	//
	// - 2: organization
	//
	// - 3: workspace
	//
	// example:
	//
	// 0
	Scope *int32 `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ListAuthConfigByResourceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListAuthConfigByResourceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAuthConfigByResourceResponseBodyResult) GetAuthorities() []*string {
	return s.Authorities
}

func (s *ListAuthConfigByResourceResponseBodyResult) GetAuthorizedId() *string {
	return s.AuthorizedId
}

func (s *ListAuthConfigByResourceResponseBodyResult) GetAuthorizedName() *string {
	return s.AuthorizedName
}

func (s *ListAuthConfigByResourceResponseBodyResult) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *ListAuthConfigByResourceResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListAuthConfigByResourceResponseBodyResult) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListAuthConfigByResourceResponseBodyResult) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAuthConfigByResourceResponseBodyResult) GetScope() *int32 {
	return s.Scope
}

func (s *ListAuthConfigByResourceResponseBodyResult) SetAuthorities(v []*string) *ListAuthConfigByResourceResponseBodyResult {
	s.Authorities = v
	return s
}

func (s *ListAuthConfigByResourceResponseBodyResult) SetAuthorizedId(v string) *ListAuthConfigByResourceResponseBodyResult {
	s.AuthorizedId = &v
	return s
}

func (s *ListAuthConfigByResourceResponseBodyResult) SetAuthorizedName(v string) *ListAuthConfigByResourceResponseBodyResult {
	s.AuthorizedName = &v
	return s
}

func (s *ListAuthConfigByResourceResponseBodyResult) SetExpireDate(v string) *ListAuthConfigByResourceResponseBodyResult {
	s.ExpireDate = &v
	return s
}

func (s *ListAuthConfigByResourceResponseBodyResult) SetGmtCreate(v string) *ListAuthConfigByResourceResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListAuthConfigByResourceResponseBodyResult) SetResourceId(v string) *ListAuthConfigByResourceResponseBodyResult {
	s.ResourceId = &v
	return s
}

func (s *ListAuthConfigByResourceResponseBodyResult) SetResourceType(v string) *ListAuthConfigByResourceResponseBodyResult {
	s.ResourceType = &v
	return s
}

func (s *ListAuthConfigByResourceResponseBodyResult) SetScope(v int32) *ListAuthConfigByResourceResponseBodyResult {
	s.Scope = &v
	return s
}

func (s *ListAuthConfigByResourceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
