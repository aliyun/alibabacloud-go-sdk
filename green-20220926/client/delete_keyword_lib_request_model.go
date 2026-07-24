// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKeywordLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLibId(v string) *DeleteKeywordLibRequest
	GetLibId() *string
	SetRegionId(v string) *DeleteKeywordLibRequest
	GetRegionId() *string
	SetTenantCode(v string) *DeleteKeywordLibRequest
	GetTenantCode() *string
}

type DeleteKeywordLibRequest struct {
	// The keyword library ID.
	//
	// example:
	//
	// customxx_xxxx
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

func (s DeleteKeywordLibRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKeywordLibRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeywordLibRequest) GetLibId() *string {
	return s.LibId
}

func (s *DeleteKeywordLibRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteKeywordLibRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *DeleteKeywordLibRequest) SetLibId(v string) *DeleteKeywordLibRequest {
	s.LibId = &v
	return s
}

func (s *DeleteKeywordLibRequest) SetRegionId(v string) *DeleteKeywordLibRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteKeywordLibRequest) SetTenantCode(v string) *DeleteKeywordLibRequest {
	s.TenantCode = &v
	return s
}

func (s *DeleteKeywordLibRequest) Validate() error {
	return dara.Validate(s)
}
