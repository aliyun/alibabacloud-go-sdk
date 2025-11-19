// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCredentialsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*CredentialListItem) *ListCredentialsOutput
	GetItems() []*CredentialListItem
	SetPageNumber(v string) *ListCredentialsOutput
	GetPageNumber() *string
	SetPageSize(v string) *ListCredentialsOutput
	GetPageSize() *string
	SetTotal(v string) *ListCredentialsOutput
	GetTotal() *string
}

type ListCredentialsOutput struct {
	Items      []*CredentialListItem `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	PageNumber *string               `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *string               `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total      *string               `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListCredentialsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialsOutput) GoString() string {
	return s.String()
}

func (s *ListCredentialsOutput) GetItems() []*CredentialListItem {
	return s.Items
}

func (s *ListCredentialsOutput) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListCredentialsOutput) GetPageSize() *string {
	return s.PageSize
}

func (s *ListCredentialsOutput) GetTotal() *string {
	return s.Total
}

func (s *ListCredentialsOutput) SetItems(v []*CredentialListItem) *ListCredentialsOutput {
	s.Items = v
	return s
}

func (s *ListCredentialsOutput) SetPageNumber(v string) *ListCredentialsOutput {
	s.PageNumber = &v
	return s
}

func (s *ListCredentialsOutput) SetPageSize(v string) *ListCredentialsOutput {
	s.PageSize = &v
	return s
}

func (s *ListCredentialsOutput) SetTotal(v string) *ListCredentialsOutput {
	s.Total = &v
	return s
}

func (s *ListCredentialsOutput) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
