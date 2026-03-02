// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDependLibraryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *RemoveDependLibraryRequest
	GetCompanyId() *int64
}

type RemoveDependLibraryRequest struct {
	// This parameter is required.
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
}

func (s RemoveDependLibraryRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveDependLibraryRequest) GoString() string {
	return s.String()
}

func (s *RemoveDependLibraryRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *RemoveDependLibraryRequest) SetCompanyId(v int64) *RemoveDependLibraryRequest {
	s.CompanyId = &v
	return s
}

func (s *RemoveDependLibraryRequest) Validate() error {
	return dara.Validate(s)
}
