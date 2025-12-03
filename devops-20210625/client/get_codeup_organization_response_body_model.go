// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCodeupOrganizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetCodeupOrganizationResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetCodeupOrganizationResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetCodeupOrganizationResponseBody
	GetRequestId() *string
	SetResult(v *GetCodeupOrganizationResponseBodyResult) *GetCodeupOrganizationResponseBody
	GetResult() *GetCodeupOrganizationResponseBodyResult
	SetSuccess(v bool) *GetCodeupOrganizationResponseBody
	GetSuccess() *bool
}

type GetCodeupOrganizationResponseBody struct {
	// example:
	//
	// InvalidTagGroup.IdNotFound
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetCodeupOrganizationResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetCodeupOrganizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCodeupOrganizationResponseBody) GoString() string {
	return s.String()
}

func (s *GetCodeupOrganizationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetCodeupOrganizationResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetCodeupOrganizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCodeupOrganizationResponseBody) GetResult() *GetCodeupOrganizationResponseBodyResult {
	return s.Result
}

func (s *GetCodeupOrganizationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCodeupOrganizationResponseBody) SetErrorCode(v string) *GetCodeupOrganizationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetCodeupOrganizationResponseBody) SetErrorMessage(v string) *GetCodeupOrganizationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetCodeupOrganizationResponseBody) SetRequestId(v string) *GetCodeupOrganizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCodeupOrganizationResponseBody) SetResult(v *GetCodeupOrganizationResponseBodyResult) *GetCodeupOrganizationResponseBody {
	s.Result = v
	return s
}

func (s *GetCodeupOrganizationResponseBody) SetSuccess(v bool) *GetCodeupOrganizationResponseBody {
	s.Success = &v
	return s
}

func (s *GetCodeupOrganizationResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCodeupOrganizationResponseBodyResult struct {
	// example:
	//
	// 2022-03-12 12:00:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 3624
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 100003
	NamespaceId *int64 `json:"namespaceId,omitempty" xml:"namespaceId,omitempty"`
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// test-codeup
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// ORG_MEMBER
	UserRole *string `json:"userRole,omitempty" xml:"userRole,omitempty"`
}

func (s GetCodeupOrganizationResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetCodeupOrganizationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetCodeupOrganizationResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetCodeupOrganizationResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *GetCodeupOrganizationResponseBodyResult) GetNamespaceId() *int64 {
	return s.NamespaceId
}

func (s *GetCodeupOrganizationResponseBodyResult) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetCodeupOrganizationResponseBodyResult) GetPath() *string {
	return s.Path
}

func (s *GetCodeupOrganizationResponseBodyResult) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetCodeupOrganizationResponseBodyResult) GetUserRole() *string {
	return s.UserRole
}

func (s *GetCodeupOrganizationResponseBodyResult) SetCreatedAt(v string) *GetCodeupOrganizationResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetId(v int64) *GetCodeupOrganizationResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetNamespaceId(v int64) *GetCodeupOrganizationResponseBodyResult {
	s.NamespaceId = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetOrganizationId(v string) *GetCodeupOrganizationResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetPath(v string) *GetCodeupOrganizationResponseBodyResult {
	s.Path = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetUpdatedAt(v string) *GetCodeupOrganizationResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetUserRole(v string) *GetCodeupOrganizationResponseBodyResult {
	s.UserRole = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
