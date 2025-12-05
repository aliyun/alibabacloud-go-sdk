// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAliasesByKeyIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliases(v *ListAliasesByKeyIdResponseBodyAliases) *ListAliasesByKeyIdResponseBody
	GetAliases() *ListAliasesByKeyIdResponseBodyAliases
	SetPageNumber(v int32) *ListAliasesByKeyIdResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAliasesByKeyIdResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAliasesByKeyIdResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAliasesByKeyIdResponseBody
	GetTotalCount() *int32
}

type ListAliasesByKeyIdResponseBody struct {
	// An array that consists of aliases.
	Aliases *ListAliasesByKeyIdResponseBodyAliases `json:"Aliases,omitempty" xml:"Aliases,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 1b57992c-834b-4811-a889-f8bac1ba0353
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned CMKs.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAliasesByKeyIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAliasesByKeyIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListAliasesByKeyIdResponseBody) GetAliases() *ListAliasesByKeyIdResponseBodyAliases {
	return s.Aliases
}

func (s *ListAliasesByKeyIdResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAliasesByKeyIdResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAliasesByKeyIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAliasesByKeyIdResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAliasesByKeyIdResponseBody) SetAliases(v *ListAliasesByKeyIdResponseBodyAliases) *ListAliasesByKeyIdResponseBody {
	s.Aliases = v
	return s
}

func (s *ListAliasesByKeyIdResponseBody) SetPageNumber(v int32) *ListAliasesByKeyIdResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAliasesByKeyIdResponseBody) SetPageSize(v int32) *ListAliasesByKeyIdResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAliasesByKeyIdResponseBody) SetRequestId(v string) *ListAliasesByKeyIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAliasesByKeyIdResponseBody) SetTotalCount(v int32) *ListAliasesByKeyIdResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAliasesByKeyIdResponseBody) Validate() error {
	if s.Aliases != nil {
		if err := s.Aliases.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAliasesByKeyIdResponseBodyAliases struct {
	Alias []*ListAliasesByKeyIdResponseBodyAliasesAlias `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Repeated"`
}

func (s ListAliasesByKeyIdResponseBodyAliases) String() string {
	return dara.Prettify(s)
}

func (s ListAliasesByKeyIdResponseBodyAliases) GoString() string {
	return s.String()
}

func (s *ListAliasesByKeyIdResponseBodyAliases) GetAlias() []*ListAliasesByKeyIdResponseBodyAliasesAlias {
	return s.Alias
}

func (s *ListAliasesByKeyIdResponseBodyAliases) SetAlias(v []*ListAliasesByKeyIdResponseBodyAliasesAlias) *ListAliasesByKeyIdResponseBodyAliases {
	s.Alias = v
	return s
}

func (s *ListAliasesByKeyIdResponseBodyAliases) Validate() error {
	if s.Alias != nil {
		for _, item := range s.Alias {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAliasesByKeyIdResponseBodyAliasesAlias struct {
	// The Alibaba Cloud Resource Name (ARN) of the alias.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:123456:alias/ExampleAlias1
	AliasArn *string `json:"AliasArn,omitempty" xml:"AliasArn,omitempty"`
	// The ID of the alias.
	//
	// example:
	//
	// alias/ExampleAlias1
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The CMK to which an alias is bound.
	//
	// example:
	//
	// 08c33a6f-4e0a-4a1b-a3fa-7ddfa1d4****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s ListAliasesByKeyIdResponseBodyAliasesAlias) String() string {
	return dara.Prettify(s)
}

func (s ListAliasesByKeyIdResponseBodyAliasesAlias) GoString() string {
	return s.String()
}

func (s *ListAliasesByKeyIdResponseBodyAliasesAlias) GetAliasArn() *string {
	return s.AliasArn
}

func (s *ListAliasesByKeyIdResponseBodyAliasesAlias) GetAliasName() *string {
	return s.AliasName
}

func (s *ListAliasesByKeyIdResponseBodyAliasesAlias) GetKeyId() *string {
	return s.KeyId
}

func (s *ListAliasesByKeyIdResponseBodyAliasesAlias) SetAliasArn(v string) *ListAliasesByKeyIdResponseBodyAliasesAlias {
	s.AliasArn = &v
	return s
}

func (s *ListAliasesByKeyIdResponseBodyAliasesAlias) SetAliasName(v string) *ListAliasesByKeyIdResponseBodyAliasesAlias {
	s.AliasName = &v
	return s
}

func (s *ListAliasesByKeyIdResponseBodyAliasesAlias) SetKeyId(v string) *ListAliasesByKeyIdResponseBodyAliasesAlias {
	s.KeyId = &v
	return s
}

func (s *ListAliasesByKeyIdResponseBodyAliasesAlias) Validate() error {
	return dara.Validate(s)
}
