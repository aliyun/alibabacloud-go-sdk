// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKeywordLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLibId(v string) *UpdateKeywordLibRequest
	GetLibId() *string
	SetLibName(v string) *UpdateKeywordLibRequest
	GetLibName() *string
	SetRegionId(v string) *UpdateKeywordLibRequest
	GetRegionId() *string
	SetTenantCode(v string) *UpdateKeywordLibRequest
	GetTenantCode() *string
}

type UpdateKeywordLibRequest struct {
	// The ID of the keyword library.
	//
	// example:
	//
	// custom_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// The name of the keyword library.
	//
	// example:
	//
	// TestLibrary.
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The keyword library code.
	//
	// - desensitize: desensitization keyword library
	//
	// example:
	//
	// desensitize
	TenantCode *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
}

func (s UpdateKeywordLibRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKeywordLibRequest) GoString() string {
	return s.String()
}

func (s *UpdateKeywordLibRequest) GetLibId() *string {
	return s.LibId
}

func (s *UpdateKeywordLibRequest) GetLibName() *string {
	return s.LibName
}

func (s *UpdateKeywordLibRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateKeywordLibRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *UpdateKeywordLibRequest) SetLibId(v string) *UpdateKeywordLibRequest {
	s.LibId = &v
	return s
}

func (s *UpdateKeywordLibRequest) SetLibName(v string) *UpdateKeywordLibRequest {
	s.LibName = &v
	return s
}

func (s *UpdateKeywordLibRequest) SetRegionId(v string) *UpdateKeywordLibRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateKeywordLibRequest) SetTenantCode(v string) *UpdateKeywordLibRequest {
	s.TenantCode = &v
	return s
}

func (s *UpdateKeywordLibRequest) Validate() error {
	return dara.Validate(s)
}
