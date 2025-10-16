// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCredentialsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *CredentialListItem) *ListCredentialsOutput
	GetItems() *CredentialListItem
	SetPageNum(v string) *ListCredentialsOutput
	GetPageNum() *string
	SetPageSize(v string) *ListCredentialsOutput
	GetPageSize() *string
	SetTotal(v string) *ListCredentialsOutput
	GetTotal() *string
}

type ListCredentialsOutput struct {
	Items    *CredentialListItem `json:"items,omitempty" xml:"items,omitempty"`
	PageNum  *string             `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *string             `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *string             `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListCredentialsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialsOutput) GoString() string {
	return s.String()
}

func (s *ListCredentialsOutput) GetItems() *CredentialListItem {
	return s.Items
}

func (s *ListCredentialsOutput) GetPageNum() *string {
	return s.PageNum
}

func (s *ListCredentialsOutput) GetPageSize() *string {
	return s.PageSize
}

func (s *ListCredentialsOutput) GetTotal() *string {
	return s.Total
}

func (s *ListCredentialsOutput) SetItems(v *CredentialListItem) *ListCredentialsOutput {
	s.Items = v
	return s
}

func (s *ListCredentialsOutput) SetPageNum(v string) *ListCredentialsOutput {
	s.PageNum = &v
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
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}
