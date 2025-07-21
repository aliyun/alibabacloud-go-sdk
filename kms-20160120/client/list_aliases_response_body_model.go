// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAliasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliases(v *ListAliasesResponseBodyAliases) *ListAliasesResponseBody
	GetAliases() *ListAliasesResponseBodyAliases
	SetPageNumber(v int32) *ListAliasesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAliasesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAliasesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAliasesResponseBody
	GetTotalCount() *int32
}

type ListAliasesResponseBody struct {
	// The alias of the user.
	Aliases *ListAliasesResponseBodyAliases `json:"Aliases,omitempty" xml:"Aliases,omitempty" type:"Struct"`
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
	// The ID of the request.
	//
	// example:
	//
	// 1b57992c-834b-4811-a889-f8bac1ba0353
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned aliases.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAliasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAliasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAliasesResponseBody) GetAliases() *ListAliasesResponseBodyAliases {
	return s.Aliases
}

func (s *ListAliasesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAliasesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAliasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAliasesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAliasesResponseBody) SetAliases(v *ListAliasesResponseBodyAliases) *ListAliasesResponseBody {
	s.Aliases = v
	return s
}

func (s *ListAliasesResponseBody) SetPageNumber(v int32) *ListAliasesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAliasesResponseBody) SetPageSize(v int32) *ListAliasesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAliasesResponseBody) SetRequestId(v string) *ListAliasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAliasesResponseBody) SetTotalCount(v int32) *ListAliasesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAliasesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAliasesResponseBodyAliases struct {
	Alias []*ListAliasesResponseBodyAliasesAlias `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Repeated"`
}

func (s ListAliasesResponseBodyAliases) String() string {
	return dara.Prettify(s)
}

func (s ListAliasesResponseBodyAliases) GoString() string {
	return s.String()
}

func (s *ListAliasesResponseBodyAliases) GetAlias() []*ListAliasesResponseBodyAliasesAlias {
	return s.Alias
}

func (s *ListAliasesResponseBodyAliases) SetAlias(v []*ListAliasesResponseBodyAliasesAlias) *ListAliasesResponseBodyAliases {
	s.Alias = v
	return s
}

func (s *ListAliasesResponseBodyAliases) Validate() error {
	return dara.Validate(s)
}

type ListAliasesResponseBodyAliasesAlias struct {
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
	// The CMK to which the alias belongs.
	//
	// example:
	//
	// 08c33a6f-4e0a-4a1b-a3fa-7ddfa1d****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s ListAliasesResponseBodyAliasesAlias) String() string {
	return dara.Prettify(s)
}

func (s ListAliasesResponseBodyAliasesAlias) GoString() string {
	return s.String()
}

func (s *ListAliasesResponseBodyAliasesAlias) GetAliasArn() *string {
	return s.AliasArn
}

func (s *ListAliasesResponseBodyAliasesAlias) GetAliasName() *string {
	return s.AliasName
}

func (s *ListAliasesResponseBodyAliasesAlias) GetKeyId() *string {
	return s.KeyId
}

func (s *ListAliasesResponseBodyAliasesAlias) SetAliasArn(v string) *ListAliasesResponseBodyAliasesAlias {
	s.AliasArn = &v
	return s
}

func (s *ListAliasesResponseBodyAliasesAlias) SetAliasName(v string) *ListAliasesResponseBodyAliasesAlias {
	s.AliasName = &v
	return s
}

func (s *ListAliasesResponseBodyAliasesAlias) SetKeyId(v string) *ListAliasesResponseBodyAliasesAlias {
	s.KeyId = &v
	return s
}

func (s *ListAliasesResponseBodyAliasesAlias) Validate() error {
	return dara.Validate(s)
}
