// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKeywordLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLibId(v string) *GetKeywordLibRequest
	GetLibId() *string
	SetRegionId(v string) *GetKeywordLibRequest
	GetRegionId() *string
	SetTenantCode(v string) *GetKeywordLibRequest
	GetTenantCode() *string
}

type GetKeywordLibRequest struct {
	// The keyword library ID.
	//
	// example:
	//
	// customxx_xxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The library code.
	//
	// - desensitize: masking library
	//
	// example:
	//
	// desensitize
	TenantCode *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
}

func (s GetKeywordLibRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKeywordLibRequest) GoString() string {
	return s.String()
}

func (s *GetKeywordLibRequest) GetLibId() *string {
	return s.LibId
}

func (s *GetKeywordLibRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetKeywordLibRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *GetKeywordLibRequest) SetLibId(v string) *GetKeywordLibRequest {
	s.LibId = &v
	return s
}

func (s *GetKeywordLibRequest) SetRegionId(v string) *GetKeywordLibRequest {
	s.RegionId = &v
	return s
}

func (s *GetKeywordLibRequest) SetTenantCode(v string) *GetKeywordLibRequest {
	s.TenantCode = &v
	return s
}

func (s *GetKeywordLibRequest) Validate() error {
	return dara.Validate(s)
}
