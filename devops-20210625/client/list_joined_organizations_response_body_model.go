// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJoinedOrganizationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListJoinedOrganizationsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListJoinedOrganizationsResponseBody
	GetErrorMessage() *string
	SetOrganizations(v []*ListJoinedOrganizationsResponseBodyOrganizations) *ListJoinedOrganizationsResponseBody
	GetOrganizations() []*ListJoinedOrganizationsResponseBodyOrganizations
	SetRequestId(v string) *ListJoinedOrganizationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListJoinedOrganizationsResponseBody
	GetSuccess() *bool
}

type ListJoinedOrganizationsResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ”“
	ErrorMessage  *string                                             `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Organizations []*ListJoinedOrganizationsResponseBodyOrganizations `json:"organizations,omitempty" xml:"organizations,omitempty" type:"Repeated"`
	// example:
	//
	// 11D0EE6E-5803-5D4C-A652-E672BE1F3D8E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListJoinedOrganizationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJoinedOrganizationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJoinedOrganizationsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListJoinedOrganizationsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListJoinedOrganizationsResponseBody) GetOrganizations() []*ListJoinedOrganizationsResponseBodyOrganizations {
	return s.Organizations
}

func (s *ListJoinedOrganizationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJoinedOrganizationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListJoinedOrganizationsResponseBody) SetErrorCode(v string) *ListJoinedOrganizationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListJoinedOrganizationsResponseBody) SetErrorMessage(v string) *ListJoinedOrganizationsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListJoinedOrganizationsResponseBody) SetOrganizations(v []*ListJoinedOrganizationsResponseBodyOrganizations) *ListJoinedOrganizationsResponseBody {
	s.Organizations = v
	return s
}

func (s *ListJoinedOrganizationsResponseBody) SetRequestId(v string) *ListJoinedOrganizationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJoinedOrganizationsResponseBody) SetSuccess(v bool) *ListJoinedOrganizationsResponseBody {
	s.Success = &v
	return s
}

func (s *ListJoinedOrganizationsResponseBody) Validate() error {
	if s.Organizations != nil {
		for _, item := range s.Organizations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJoinedOrganizationsResponseBodyOrganizations struct {
	// example:
	//
	// 65f25d0fa54c216dcf6b1dbd
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	IsOrgAdmin *bool   `json:"isOrgAdmin,omitempty" xml:"isOrgAdmin,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListJoinedOrganizationsResponseBodyOrganizations) String() string {
	return dara.Prettify(s)
}

func (s ListJoinedOrganizationsResponseBodyOrganizations) GoString() string {
	return s.String()
}

func (s *ListJoinedOrganizationsResponseBodyOrganizations) GetId() *string {
	return s.Id
}

func (s *ListJoinedOrganizationsResponseBodyOrganizations) GetIsOrgAdmin() *bool {
	return s.IsOrgAdmin
}

func (s *ListJoinedOrganizationsResponseBodyOrganizations) GetName() *string {
	return s.Name
}

func (s *ListJoinedOrganizationsResponseBodyOrganizations) SetId(v string) *ListJoinedOrganizationsResponseBodyOrganizations {
	s.Id = &v
	return s
}

func (s *ListJoinedOrganizationsResponseBodyOrganizations) SetIsOrgAdmin(v bool) *ListJoinedOrganizationsResponseBodyOrganizations {
	s.IsOrgAdmin = &v
	return s
}

func (s *ListJoinedOrganizationsResponseBodyOrganizations) SetName(v string) *ListJoinedOrganizationsResponseBodyOrganizations {
	s.Name = &v
	return s
}

func (s *ListJoinedOrganizationsResponseBodyOrganizations) Validate() error {
	return dara.Validate(s)
}
