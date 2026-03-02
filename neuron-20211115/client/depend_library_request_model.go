// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDependLibraryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *DependLibraryRequest
	GetCompanyId() *int64
}

type DependLibraryRequest struct {
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
}

func (s DependLibraryRequest) String() string {
	return dara.Prettify(s)
}

func (s DependLibraryRequest) GoString() string {
	return s.String()
}

func (s *DependLibraryRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *DependLibraryRequest) SetCompanyId(v int64) *DependLibraryRequest {
	s.CompanyId = &v
	return s
}

func (s *DependLibraryRequest) Validate() error {
	return dara.Validate(s)
}
