// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApiKeyListQry interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v int64) *ApiKeyListQry
	GetClientId() *int64
	SetKeyword(v string) *ApiKeyListQry
	GetKeyword() *string
	SetPage(v int32) *ApiKeyListQry
	GetPage() *int32
	SetPageSize(v int32) *ApiKeyListQry
	GetPageSize() *int32
	SetStatus(v int32) *ApiKeyListQry
	GetStatus() *int32
}

type ApiKeyListQry struct {
	// example:
	//
	// 1
	ClientId *int64  `json:"clientId,omitempty" xml:"clientId,omitempty"`
	Keyword  *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ApiKeyListQry) String() string {
	return dara.Prettify(s)
}

func (s ApiKeyListQry) GoString() string {
	return s.String()
}

func (s *ApiKeyListQry) GetClientId() *int64 {
	return s.ClientId
}

func (s *ApiKeyListQry) GetKeyword() *string {
	return s.Keyword
}

func (s *ApiKeyListQry) GetPage() *int32 {
	return s.Page
}

func (s *ApiKeyListQry) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ApiKeyListQry) GetStatus() *int32 {
	return s.Status
}

func (s *ApiKeyListQry) SetClientId(v int64) *ApiKeyListQry {
	s.ClientId = &v
	return s
}

func (s *ApiKeyListQry) SetKeyword(v string) *ApiKeyListQry {
	s.Keyword = &v
	return s
}

func (s *ApiKeyListQry) SetPage(v int32) *ApiKeyListQry {
	s.Page = &v
	return s
}

func (s *ApiKeyListQry) SetPageSize(v int32) *ApiKeyListQry {
	s.PageSize = &v
	return s
}

func (s *ApiKeyListQry) SetStatus(v int32) *ApiKeyListQry {
	s.Status = &v
	return s
}

func (s *ApiKeyListQry) Validate() error {
	return dara.Validate(s)
}
