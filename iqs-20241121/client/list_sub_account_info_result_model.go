// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubAccountInfoResult interface {
	dara.Model
	String() string
	GoString() string
	SetSubAccountInfos(v []*ListSubAccountInfoResultSubAccountInfos) *ListSubAccountInfoResult
	GetSubAccountInfos() []*ListSubAccountInfoResultSubAccountInfos
}

type ListSubAccountInfoResult struct {
	SubAccountInfos []*ListSubAccountInfoResultSubAccountInfos `json:"subAccountInfos,omitempty" xml:"subAccountInfos,omitempty" type:"Repeated"`
}

func (s ListSubAccountInfoResult) String() string {
	return dara.Prettify(s)
}

func (s ListSubAccountInfoResult) GoString() string {
	return s.String()
}

func (s *ListSubAccountInfoResult) GetSubAccountInfos() []*ListSubAccountInfoResultSubAccountInfos {
	return s.SubAccountInfos
}

func (s *ListSubAccountInfoResult) SetSubAccountInfos(v []*ListSubAccountInfoResultSubAccountInfos) *ListSubAccountInfoResult {
	s.SubAccountInfos = v
	return s
}

func (s *ListSubAccountInfoResult) Validate() error {
	if s.SubAccountInfos != nil {
		for _, item := range s.SubAccountInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSubAccountInfoResultSubAccountInfos struct {
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
}

func (s ListSubAccountInfoResultSubAccountInfos) String() string {
	return dara.Prettify(s)
}

func (s ListSubAccountInfoResultSubAccountInfos) GoString() string {
	return s.String()
}

func (s *ListSubAccountInfoResultSubAccountInfos) GetAccountName() *string {
	return s.AccountName
}

func (s *ListSubAccountInfoResultSubAccountInfos) SetAccountName(v string) *ListSubAccountInfoResultSubAccountInfos {
	s.AccountName = &v
	return s
}

func (s *ListSubAccountInfoResultSubAccountInfos) Validate() error {
	return dara.Validate(s)
}
