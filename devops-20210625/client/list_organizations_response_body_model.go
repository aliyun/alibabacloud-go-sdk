// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListOrganizationsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListOrganizationsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListOrganizationsResponseBody
	GetRequestId() *string
	SetResult(v []*ListOrganizationsResponseBodyResult) *ListOrganizationsResponseBody
	GetResult() []*ListOrganizationsResponseBodyResult
	SetSuccess(v bool) *ListOrganizationsResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListOrganizationsResponseBody
	GetTotal() *int64
}

type ListOrganizationsResponseBody struct {
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 95FE5953-FF5B-5F80-94AD-FFF97D990FE0
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListOrganizationsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListOrganizationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListOrganizationsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListOrganizationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOrganizationsResponseBody) GetResult() []*ListOrganizationsResponseBodyResult {
	return s.Result
}

func (s *ListOrganizationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListOrganizationsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListOrganizationsResponseBody) SetErrorCode(v string) *ListOrganizationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListOrganizationsResponseBody) SetErrorMessage(v string) *ListOrganizationsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListOrganizationsResponseBody) SetRequestId(v string) *ListOrganizationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrganizationsResponseBody) SetResult(v []*ListOrganizationsResponseBodyResult) *ListOrganizationsResponseBody {
	s.Result = v
	return s
}

func (s *ListOrganizationsResponseBody) SetSuccess(v bool) *ListOrganizationsResponseBody {
	s.Success = &v
	return s
}

func (s *ListOrganizationsResponseBody) SetTotal(v int64) *ListOrganizationsResponseBody {
	s.Total = &v
	return s
}

func (s *ListOrganizationsResponseBody) Validate() error {
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

type ListOrganizationsResponseBodyResult struct {
	// example:
	//
	// 60
	AccessLevel *int32 `json:"accessLevel,omitempty" xml:"accessLevel,omitempty"`
	// example:
	//
	// 5240
	Id                *int64  `json:"id,omitempty" xml:"id,omitempty"`
	NamespaceId       *string `json:"namespaceId,omitempty" xml:"namespaceId,omitempty"`
	OrganizationAlias *string `json:"organizationAlias,omitempty" xml:"organizationAlias,omitempty"`
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId   *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	OrganizationName *string `json:"organizationName,omitempty" xml:"organizationName,omitempty"`
	// example:
	//
	// ORG_ADMIN
	OrganizationRole *string `json:"organizationRole,omitempty" xml:"organizationRole,omitempty"`
}

func (s ListOrganizationsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListOrganizationsResponseBodyResult) GetAccessLevel() *int32 {
	return s.AccessLevel
}

func (s *ListOrganizationsResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *ListOrganizationsResponseBodyResult) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListOrganizationsResponseBodyResult) GetOrganizationAlias() *string {
	return s.OrganizationAlias
}

func (s *ListOrganizationsResponseBodyResult) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListOrganizationsResponseBodyResult) GetOrganizationName() *string {
	return s.OrganizationName
}

func (s *ListOrganizationsResponseBodyResult) GetOrganizationRole() *string {
	return s.OrganizationRole
}

func (s *ListOrganizationsResponseBodyResult) SetAccessLevel(v int32) *ListOrganizationsResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *ListOrganizationsResponseBodyResult) SetId(v int64) *ListOrganizationsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListOrganizationsResponseBodyResult) SetNamespaceId(v string) *ListOrganizationsResponseBodyResult {
	s.NamespaceId = &v
	return s
}

func (s *ListOrganizationsResponseBodyResult) SetOrganizationAlias(v string) *ListOrganizationsResponseBodyResult {
	s.OrganizationAlias = &v
	return s
}

func (s *ListOrganizationsResponseBodyResult) SetOrganizationId(v string) *ListOrganizationsResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *ListOrganizationsResponseBodyResult) SetOrganizationName(v string) *ListOrganizationsResponseBodyResult {
	s.OrganizationName = &v
	return s
}

func (s *ListOrganizationsResponseBodyResult) SetOrganizationRole(v string) *ListOrganizationsResponseBodyResult {
	s.OrganizationRole = &v
	return s
}

func (s *ListOrganizationsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
