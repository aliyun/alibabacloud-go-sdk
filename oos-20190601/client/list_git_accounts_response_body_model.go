// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGitAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListGitAccountsResponseBody
	GetCount() *int32
	SetGitAccounts(v []*ListGitAccountsResponseBodyGitAccounts) *ListGitAccountsResponseBody
	GetGitAccounts() []*ListGitAccountsResponseBodyGitAccounts
	SetRequestId(v string) *ListGitAccountsResponseBody
	GetRequestId() *string
}

type ListGitAccountsResponseBody struct {
	// example:
	//
	// 1
	Count       *int32                                    `json:"Count,omitempty" xml:"Count,omitempty"`
	GitAccounts []*ListGitAccountsResponseBodyGitAccounts `json:"GitAccounts,omitempty" xml:"GitAccounts,omitempty" type:"Repeated"`
	// example:
	//
	// 1457F46C-7AAE-59FA-BD12-0BDB3751E6F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGitAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGitAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGitAccountsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListGitAccountsResponseBody) GetGitAccounts() []*ListGitAccountsResponseBodyGitAccounts {
	return s.GitAccounts
}

func (s *ListGitAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGitAccountsResponseBody) SetCount(v int32) *ListGitAccountsResponseBody {
	s.Count = &v
	return s
}

func (s *ListGitAccountsResponseBody) SetGitAccounts(v []*ListGitAccountsResponseBodyGitAccounts) *ListGitAccountsResponseBody {
	s.GitAccounts = v
	return s
}

func (s *ListGitAccountsResponseBody) SetRequestId(v string) *ListGitAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGitAccountsResponseBody) Validate() error {
	if s.GitAccounts != nil {
		for _, item := range s.GitAccounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGitAccountsResponseBodyGitAccounts struct {
	// example:
	//
	// true
	IsActive *bool `json:"IsActive,omitempty" xml:"IsActive,omitempty"`
	// example:
	//
	// LYH-RAIN
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
}

func (s ListGitAccountsResponseBodyGitAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListGitAccountsResponseBodyGitAccounts) GoString() string {
	return s.String()
}

func (s *ListGitAccountsResponseBodyGitAccounts) GetIsActive() *bool {
	return s.IsActive
}

func (s *ListGitAccountsResponseBodyGitAccounts) GetOwner() *string {
	return s.Owner
}

func (s *ListGitAccountsResponseBodyGitAccounts) SetIsActive(v bool) *ListGitAccountsResponseBodyGitAccounts {
	s.IsActive = &v
	return s
}

func (s *ListGitAccountsResponseBodyGitAccounts) SetOwner(v string) *ListGitAccountsResponseBodyGitAccounts {
	s.Owner = &v
	return s
}

func (s *ListGitAccountsResponseBodyGitAccounts) Validate() error {
	return dara.Validate(s)
}
