// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListAuthorizedUsersRequest
	GetAppId() *string
	SetAppInstanceGroupId(v string) *ListAuthorizedUsersRequest
	GetAppInstanceGroupId() *string
	SetAppInstanceGroupSetId(v string) *ListAuthorizedUsersRequest
	GetAppInstanceGroupSetId() *string
	SetAppInstancePersistentId(v string) *ListAuthorizedUsersRequest
	GetAppInstancePersistentId() *string
	SetEndUserId(v string) *ListAuthorizedUsersRequest
	GetEndUserId() *string
	SetPageNumber(v int32) *ListAuthorizedUsersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAuthorizedUsersRequest
	GetPageSize() *int32
	SetProductType(v string) *ListAuthorizedUsersRequest
	GetProductType() *string
	SetUserIdFuzzy(v string) *ListAuthorizedUsersRequest
	GetUserIdFuzzy() *string
}

type ListAuthorizedUsersRequest struct {
	// The application ID used to filter authorization relationships.
	//
	// Set this parameter when querying authorized users of a specific application. This parameter is not required when querying cloud browser groups or delivery group sets.
	//
	// example:
	//
	// -
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The delivery group ID. When querying cloud browsers, set this parameter to the browser group ID.
	//
	// Specify either this parameter or `AppInstanceGroupSetId`, but not both.
	//
	// example:
	//
	// big-3jm9d0abc00example
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The delivery group set ID.
	//
	// Specify either this parameter or `AppInstanceGroupId`, but not both. When querying by set, omit `AppId` and `AppInstancePersistentId`.
	//
	// example:
	//
	// set-3jm9d0abc00example
	AppInstanceGroupSetId *string `json:"AppInstanceGroupSetId,omitempty" xml:"AppInstanceGroupSetId,omitempty"`
	// The persistent session ID used to filter authorization relationships. This parameter applies to delivery groups that use session-based authorization.
	//
	// This parameter is not required when querying delivery group sets.
	//
	// example:
	//
	// ai-3jm9d0abc00example
	AppInstancePersistentId *string `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	// Performs an exact match by authorized username. If this parameter is not specified, results are not filtered by exact username.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The page number. This parameter is required. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of records per page. This parameter is required. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type. Set this parameter to `CloudBrowser` when querying authorized users of cloud browsers.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudBrowser
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// Performs a fuzzy match by text contained in the authorized username.
	//
	// example:
	//
	// ali
	UserIdFuzzy *string `json:"UserIdFuzzy,omitempty" xml:"UserIdFuzzy,omitempty"`
}

func (s ListAuthorizedUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUsersRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUsersRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListAuthorizedUsersRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListAuthorizedUsersRequest) GetAppInstanceGroupSetId() *string {
	return s.AppInstanceGroupSetId
}

func (s *ListAuthorizedUsersRequest) GetAppInstancePersistentId() *string {
	return s.AppInstancePersistentId
}

func (s *ListAuthorizedUsersRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ListAuthorizedUsersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAuthorizedUsersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAuthorizedUsersRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListAuthorizedUsersRequest) GetUserIdFuzzy() *string {
	return s.UserIdFuzzy
}

func (s *ListAuthorizedUsersRequest) SetAppId(v string) *ListAuthorizedUsersRequest {
	s.AppId = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetAppInstanceGroupId(v string) *ListAuthorizedUsersRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetAppInstanceGroupSetId(v string) *ListAuthorizedUsersRequest {
	s.AppInstanceGroupSetId = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetAppInstancePersistentId(v string) *ListAuthorizedUsersRequest {
	s.AppInstancePersistentId = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetEndUserId(v string) *ListAuthorizedUsersRequest {
	s.EndUserId = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetPageNumber(v int32) *ListAuthorizedUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetPageSize(v int32) *ListAuthorizedUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetProductType(v string) *ListAuthorizedUsersRequest {
	s.ProductType = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetUserIdFuzzy(v string) *ListAuthorizedUsersRequest {
	s.UserIdFuzzy = &v
	return s
}

func (s *ListAuthorizedUsersRequest) Validate() error {
	return dara.Validate(s)
}
