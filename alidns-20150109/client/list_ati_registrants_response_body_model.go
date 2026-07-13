// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAtiRegistrantsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ListAtiRegistrantsResponseBodyAccessDeniedDetail) *ListAtiRegistrantsResponseBody
	GetAccessDeniedDetail() *ListAtiRegistrantsResponseBodyAccessDeniedDetail
	SetMaxResults(v int32) *ListAtiRegistrantsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAtiRegistrantsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListAtiRegistrantsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAtiRegistrantsResponseBody
	GetPageSize() *int32
	SetRegistrants(v *ListAtiRegistrantsResponseBodyRegistrants) *ListAtiRegistrantsResponseBody
	GetRegistrants() *ListAtiRegistrantsResponseBodyRegistrants
	SetRequestId(v string) *ListAtiRegistrantsResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *ListAtiRegistrantsResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *ListAtiRegistrantsResponseBody
	GetTotalPages() *int32
}

type ListAtiRegistrantsResponseBody struct {
	AccessDeniedDetail *ListAtiRegistrantsResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 500
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 4698691
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 2
	PageSize    *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Registrants *ListAtiRegistrantsResponseBodyRegistrants `json:"Registrants,omitempty" xml:"Registrants,omitempty" type:"Struct"`
	// example:
	//
	// B57C121B-A45F-44D8-A9B2-XXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 224
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// example:
	//
	// 11
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListAtiRegistrantsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAtiRegistrantsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAtiRegistrantsResponseBody) GetAccessDeniedDetail() *ListAtiRegistrantsResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ListAtiRegistrantsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAtiRegistrantsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAtiRegistrantsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAtiRegistrantsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAtiRegistrantsResponseBody) GetRegistrants() *ListAtiRegistrantsResponseBodyRegistrants {
	return s.Registrants
}

func (s *ListAtiRegistrantsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAtiRegistrantsResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *ListAtiRegistrantsResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListAtiRegistrantsResponseBody) SetAccessDeniedDetail(v *ListAtiRegistrantsResponseBodyAccessDeniedDetail) *ListAtiRegistrantsResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ListAtiRegistrantsResponseBody) SetMaxResults(v int32) *ListAtiRegistrantsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAtiRegistrantsResponseBody) SetNextToken(v string) *ListAtiRegistrantsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAtiRegistrantsResponseBody) SetPageNumber(v int32) *ListAtiRegistrantsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAtiRegistrantsResponseBody) SetPageSize(v int32) *ListAtiRegistrantsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAtiRegistrantsResponseBody) SetRegistrants(v *ListAtiRegistrantsResponseBodyRegistrants) *ListAtiRegistrantsResponseBody {
	s.Registrants = v
	return s
}

func (s *ListAtiRegistrantsResponseBody) SetRequestId(v string) *ListAtiRegistrantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAtiRegistrantsResponseBody) SetTotalItems(v int32) *ListAtiRegistrantsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *ListAtiRegistrantsResponseBody) SetTotalPages(v int32) *ListAtiRegistrantsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *ListAtiRegistrantsResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Registrants != nil {
		if err := s.Registrants.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAtiRegistrantsResponseBodyAccessDeniedDetail struct {
	// example:
	//
	// AddRspDomainServerHoldStatusForGatewayOte
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// example:
	//
	// 2015555733387XXXX
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// example:
	//
	// 10469733312XXX
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// example:
	//
	// AQEAAAAAaNIARXXXXUQwNjE0LUQzN0XXXXVEQy1BQzExLTMzXXXXNTkxRjk1Ng==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// example:
	//
	// DlpSend
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListAtiRegistrantsResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ListAtiRegistrantsResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ListAtiRegistrantsResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ListAtiRegistrantsResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ListAtiRegistrantsResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ListAtiRegistrantsResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ListAtiRegistrantsResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ListAtiRegistrantsResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ListAtiRegistrantsResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListAtiRegistrantsResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ListAtiRegistrantsResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ListAtiRegistrantsResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ListAtiRegistrantsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ListAtiRegistrantsResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ListAtiRegistrantsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ListAtiRegistrantsResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ListAtiRegistrantsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ListAtiRegistrantsResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ListAtiRegistrantsResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ListAtiRegistrantsResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ListAtiRegistrantsResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ListAtiRegistrantsResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ListAtiRegistrantsResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ListAtiRegistrantsResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ListAtiRegistrantsResponseBodyRegistrants struct {
	Registrant []*ListAtiRegistrantsResponseBodyRegistrantsRegistrant `json:"Registrant,omitempty" xml:"Registrant,omitempty" type:"Repeated"`
}

func (s ListAtiRegistrantsResponseBodyRegistrants) String() string {
	return dara.Prettify(s)
}

func (s ListAtiRegistrantsResponseBodyRegistrants) GoString() string {
	return s.String()
}

func (s *ListAtiRegistrantsResponseBodyRegistrants) GetRegistrant() []*ListAtiRegistrantsResponseBodyRegistrantsRegistrant {
	return s.Registrant
}

func (s *ListAtiRegistrantsResponseBodyRegistrants) SetRegistrant(v []*ListAtiRegistrantsResponseBodyRegistrantsRegistrant) *ListAtiRegistrantsResponseBodyRegistrants {
	s.Registrant = v
	return s
}

func (s *ListAtiRegistrantsResponseBodyRegistrants) Validate() error {
	if s.Registrant != nil {
		for _, item := range s.Registrant {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAtiRegistrantsResponseBodyRegistrantsRegistrant struct {
	Cc              *string `json:"Cc,omitempty" xml:"Cc,omitempty"`
	City            *string `json:"City,omitempty" xml:"City,omitempty"`
	CreateTimestamp *string `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	DocumentCode    *string `json:"DocumentCode,omitempty" xml:"DocumentCode,omitempty"`
	DocumentType    *string `json:"DocumentType,omitempty" xml:"DocumentType,omitempty"`
	Email           *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegistrantId    *string `json:"RegistrantId,omitempty" xml:"RegistrantId,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTimestamp *string `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s ListAtiRegistrantsResponseBodyRegistrantsRegistrant) String() string {
	return dara.Prettify(s)
}

func (s ListAtiRegistrantsResponseBodyRegistrantsRegistrant) GoString() string {
	return s.String()
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) GetCc() *string {
	return s.Cc
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) GetCity() *string {
	return s.City
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) GetCreateTimestamp() *string {
	return s.CreateTimestamp
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) GetDocumentCode() *string {
	return s.DocumentCode
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) GetDocumentType() *string {
	return s.DocumentType
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) GetEmail() *string {
	return s.Email
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) GetName() *string {
	return s.Name
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) GetRegistrantId() *string {
	return s.RegistrantId
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) GetState() *string {
	return s.State
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) GetStatus() *string {
	return s.Status
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) GetUpdateTimestamp() *string {
	return s.UpdateTimestamp
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) SetCc(v string) *ListAtiRegistrantsResponseBodyRegistrantsRegistrant {
	s.Cc = &v
	return s
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) SetCity(v string) *ListAtiRegistrantsResponseBodyRegistrantsRegistrant {
	s.City = &v
	return s
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) SetCreateTimestamp(v string) *ListAtiRegistrantsResponseBodyRegistrantsRegistrant {
	s.CreateTimestamp = &v
	return s
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) SetDocumentCode(v string) *ListAtiRegistrantsResponseBodyRegistrantsRegistrant {
	s.DocumentCode = &v
	return s
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) SetDocumentType(v string) *ListAtiRegistrantsResponseBodyRegistrantsRegistrant {
	s.DocumentType = &v
	return s
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) SetEmail(v string) *ListAtiRegistrantsResponseBodyRegistrantsRegistrant {
	s.Email = &v
	return s
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) SetName(v string) *ListAtiRegistrantsResponseBodyRegistrantsRegistrant {
	s.Name = &v
	return s
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) SetRegistrantId(v string) *ListAtiRegistrantsResponseBodyRegistrantsRegistrant {
	s.RegistrantId = &v
	return s
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) SetState(v string) *ListAtiRegistrantsResponseBodyRegistrantsRegistrant {
	s.State = &v
	return s
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) SetStatus(v string) *ListAtiRegistrantsResponseBodyRegistrantsRegistrant {
	s.Status = &v
	return s
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) SetUpdateTimestamp(v string) *ListAtiRegistrantsResponseBodyRegistrantsRegistrant {
	s.UpdateTimestamp = &v
	return s
}

func (s *ListAtiRegistrantsResponseBodyRegistrantsRegistrant) Validate() error {
	return dara.Validate(s)
}
