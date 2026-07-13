// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAtiRegistrantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeAtiRegistrantResponseBodyAccessDeniedDetail) *DescribeAtiRegistrantResponseBody
	GetAccessDeniedDetail() *DescribeAtiRegistrantResponseBodyAccessDeniedDetail
	SetCc(v string) *DescribeAtiRegistrantResponseBody
	GetCc() *string
	SetCity(v string) *DescribeAtiRegistrantResponseBody
	GetCity() *string
	SetCreateTimestamp(v string) *DescribeAtiRegistrantResponseBody
	GetCreateTimestamp() *string
	SetDocumentCode(v string) *DescribeAtiRegistrantResponseBody
	GetDocumentCode() *string
	SetDocumentType(v string) *DescribeAtiRegistrantResponseBody
	GetDocumentType() *string
	SetEmail(v string) *DescribeAtiRegistrantResponseBody
	GetEmail() *string
	SetName(v string) *DescribeAtiRegistrantResponseBody
	GetName() *string
	SetPhone(v string) *DescribeAtiRegistrantResponseBody
	GetPhone() *string
	SetRegistrantId(v string) *DescribeAtiRegistrantResponseBody
	GetRegistrantId() *string
	SetRejectReason(v string) *DescribeAtiRegistrantResponseBody
	GetRejectReason() *string
	SetRequestId(v string) *DescribeAtiRegistrantResponseBody
	GetRequestId() *string
	SetState(v string) *DescribeAtiRegistrantResponseBody
	GetState() *string
	SetStatus(v string) *DescribeAtiRegistrantResponseBody
	GetStatus() *string
	SetStreet(v string) *DescribeAtiRegistrantResponseBody
	GetStreet() *string
	SetUpdateTimestamp(v string) *DescribeAtiRegistrantResponseBody
	GetUpdateTimestamp() *string
}

type DescribeAtiRegistrantResponseBody struct {
	// The details of the access denial. This field is returned only when RAM authentication fails.
	AccessDeniedDetail *DescribeAtiRegistrantResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The country.
	//
	// example:
	//
	// 中国
	Cc *string `json:"Cc,omitempty" xml:"Cc,omitempty"`
	// The city. Default value: Hangzhou.
	//
	// example:
	//
	// 杭州市
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The creation time (timestamp).
	//
	// example:
	//
	// 1533773400000
	CreateTimestamp *string `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The document number of the registrant. The number can be up to 50 characters in length.
	//
	// example:
	//
	// 110123456789122341
	DocumentCode *string `json:"DocumentCode,omitempty" xml:"DocumentCode,omitempty"`
	// The document type of the registrant. For more information, see the appendix on document types.
	//
	// example:
	//
	// SFZ
	DocumentType *string `json:"DocumentType,omitempty" xml:"DocumentType,omitempty"`
	// The email address. The address can be up to 300 characters in length.
	//
	// example:
	//
	// example@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The name of the registrant. The name can be up to 255 characters in length.
	//
	// example:
	//
	// 张xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The phone number of the registrant. The number can be up to 128 characters in length. If the country is China and the number is not a mobile phone number, the area code must match the city.
	//
	// example:
	//
	// 13112345678
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The ID of the real-name verified registrant.
	//
	// example:
	//
	// 2072277378616354816
	RegistrantId *string `json:"RegistrantId,omitempty" xml:"RegistrantId,omitempty"`
	// The reason why the real-name verification was rejected.
	//
	// example:
	//
	// 以实际返回为准
	RejectReason *string `json:"RejectReason,omitempty" xml:"RejectReason,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B57C121B-A45F-44D8-A9B2-13E5A5044195
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the public recursive service.
	//
	// example:
	//
	// 浙江省
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The real-name verification status. Valid values:
	//
	// - Approved.
	//
	// - Under review.
	//
	// - Rejected.
	//
	// example:
	//
	// Approved
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The street address.
	//
	// example:
	//
	// xx区xx街道
	Street *string `json:"Street,omitempty" xml:"Street,omitempty"`
	// The update time (timestamp).
	//
	// example:
	//
	// 1527690629357
	UpdateTimestamp *string `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s DescribeAtiRegistrantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiRegistrantResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAtiRegistrantResponseBody) GetAccessDeniedDetail() *DescribeAtiRegistrantResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeAtiRegistrantResponseBody) GetCc() *string {
	return s.Cc
}

func (s *DescribeAtiRegistrantResponseBody) GetCity() *string {
	return s.City
}

func (s *DescribeAtiRegistrantResponseBody) GetCreateTimestamp() *string {
	return s.CreateTimestamp
}

func (s *DescribeAtiRegistrantResponseBody) GetDocumentCode() *string {
	return s.DocumentCode
}

func (s *DescribeAtiRegistrantResponseBody) GetDocumentType() *string {
	return s.DocumentType
}

func (s *DescribeAtiRegistrantResponseBody) GetEmail() *string {
	return s.Email
}

func (s *DescribeAtiRegistrantResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeAtiRegistrantResponseBody) GetPhone() *string {
	return s.Phone
}

func (s *DescribeAtiRegistrantResponseBody) GetRegistrantId() *string {
	return s.RegistrantId
}

func (s *DescribeAtiRegistrantResponseBody) GetRejectReason() *string {
	return s.RejectReason
}

func (s *DescribeAtiRegistrantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAtiRegistrantResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeAtiRegistrantResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeAtiRegistrantResponseBody) GetStreet() *string {
	return s.Street
}

func (s *DescribeAtiRegistrantResponseBody) GetUpdateTimestamp() *string {
	return s.UpdateTimestamp
}

func (s *DescribeAtiRegistrantResponseBody) SetAccessDeniedDetail(v *DescribeAtiRegistrantResponseBodyAccessDeniedDetail) *DescribeAtiRegistrantResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeAtiRegistrantResponseBody) SetCc(v string) *DescribeAtiRegistrantResponseBody {
	s.Cc = &v
	return s
}

func (s *DescribeAtiRegistrantResponseBody) SetCity(v string) *DescribeAtiRegistrantResponseBody {
	s.City = &v
	return s
}

func (s *DescribeAtiRegistrantResponseBody) SetCreateTimestamp(v string) *DescribeAtiRegistrantResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeAtiRegistrantResponseBody) SetDocumentCode(v string) *DescribeAtiRegistrantResponseBody {
	s.DocumentCode = &v
	return s
}

func (s *DescribeAtiRegistrantResponseBody) SetDocumentType(v string) *DescribeAtiRegistrantResponseBody {
	s.DocumentType = &v
	return s
}

func (s *DescribeAtiRegistrantResponseBody) SetEmail(v string) *DescribeAtiRegistrantResponseBody {
	s.Email = &v
	return s
}

func (s *DescribeAtiRegistrantResponseBody) SetName(v string) *DescribeAtiRegistrantResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeAtiRegistrantResponseBody) SetPhone(v string) *DescribeAtiRegistrantResponseBody {
	s.Phone = &v
	return s
}

func (s *DescribeAtiRegistrantResponseBody) SetRegistrantId(v string) *DescribeAtiRegistrantResponseBody {
	s.RegistrantId = &v
	return s
}

func (s *DescribeAtiRegistrantResponseBody) SetRejectReason(v string) *DescribeAtiRegistrantResponseBody {
	s.RejectReason = &v
	return s
}

func (s *DescribeAtiRegistrantResponseBody) SetRequestId(v string) *DescribeAtiRegistrantResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAtiRegistrantResponseBody) SetState(v string) *DescribeAtiRegistrantResponseBody {
	s.State = &v
	return s
}

func (s *DescribeAtiRegistrantResponseBody) SetStatus(v string) *DescribeAtiRegistrantResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeAtiRegistrantResponseBody) SetStreet(v string) *DescribeAtiRegistrantResponseBody {
	s.Street = &v
	return s
}

func (s *DescribeAtiRegistrantResponseBody) SetUpdateTimestamp(v string) *DescribeAtiRegistrantResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeAtiRegistrantResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAtiRegistrantResponseBodyAccessDeniedDetail struct {
	// The unauthorized operation that was attempted.
	//
	// example:
	//
	// RemoveRspDomainServerHoldStatusForGatewayOte
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authorization principal.
	//
	// example:
	//
	// 2015555733387XXXX
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The owner ID of the authorization principal.
	//
	// example:
	//
	// 10469733312XXX
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The identity type.
	//
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encrypted diagnostic message.
	//
	// example:
	//
	// AQEAAAAAaNIARXXXXUQwNjE0LUQzN0XXXXVEQy1BQzExLTMzXXXXNTkxRjk1Ng==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The reason for the authentication failure. Valid values:
	//
	// - ExplicitDeny: explicit deny.
	//
	// - ImplicitDeny: implicit deny.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// DlpSend
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribeAtiRegistrantResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiRegistrantResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeAtiRegistrantResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeAtiRegistrantResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeAtiRegistrantResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeAtiRegistrantResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeAtiRegistrantResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeAtiRegistrantResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeAtiRegistrantResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeAtiRegistrantResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeAtiRegistrantResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeAtiRegistrantResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeAtiRegistrantResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeAtiRegistrantResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeAtiRegistrantResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeAtiRegistrantResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeAtiRegistrantResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeAtiRegistrantResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeAtiRegistrantResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeAtiRegistrantResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeAtiRegistrantResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeAtiRegistrantResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeAtiRegistrantResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeAtiRegistrantResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
