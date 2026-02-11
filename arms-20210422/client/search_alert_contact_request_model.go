// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAlertContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactIds(v string) *SearchAlertContactRequest
	GetContactIds() *string
	SetContactName(v string) *SearchAlertContactRequest
	GetContactName() *string
	SetCurrentPage(v string) *SearchAlertContactRequest
	GetCurrentPage() *string
	SetEmail(v string) *SearchAlertContactRequest
	GetEmail() *string
	SetPageSize(v string) *SearchAlertContactRequest
	GetPageSize() *string
	SetPhone(v string) *SearchAlertContactRequest
	GetPhone() *string
	SetRegionId(v string) *SearchAlertContactRequest
	GetRegionId() *string
}

type SearchAlertContactRequest struct {
	ContactIds  *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Phone       *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SearchAlertContactRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertContactRequest) GoString() string {
	return s.String()
}

func (s *SearchAlertContactRequest) GetContactIds() *string {
	return s.ContactIds
}

func (s *SearchAlertContactRequest) GetContactName() *string {
	return s.ContactName
}

func (s *SearchAlertContactRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *SearchAlertContactRequest) GetEmail() *string {
	return s.Email
}

func (s *SearchAlertContactRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *SearchAlertContactRequest) GetPhone() *string {
	return s.Phone
}

func (s *SearchAlertContactRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchAlertContactRequest) SetContactIds(v string) *SearchAlertContactRequest {
	s.ContactIds = &v
	return s
}

func (s *SearchAlertContactRequest) SetContactName(v string) *SearchAlertContactRequest {
	s.ContactName = &v
	return s
}

func (s *SearchAlertContactRequest) SetCurrentPage(v string) *SearchAlertContactRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchAlertContactRequest) SetEmail(v string) *SearchAlertContactRequest {
	s.Email = &v
	return s
}

func (s *SearchAlertContactRequest) SetPageSize(v string) *SearchAlertContactRequest {
	s.PageSize = &v
	return s
}

func (s *SearchAlertContactRequest) SetPhone(v string) *SearchAlertContactRequest {
	s.Phone = &v
	return s
}

func (s *SearchAlertContactRequest) SetRegionId(v string) *SearchAlertContactRequest {
	s.RegionId = &v
	return s
}

func (s *SearchAlertContactRequest) Validate() error {
	return dara.Validate(s)
}
