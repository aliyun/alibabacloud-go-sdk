// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKeywordLibsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListKeywordLibsRequest
	GetRegionId() *string
	SetTenantCode(v string) *ListKeywordLibsRequest
	GetTenantCode() *string
}

type ListKeywordLibsRequest struct {
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TenantCode *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
}

func (s ListKeywordLibsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKeywordLibsRequest) GoString() string {
	return s.String()
}

func (s *ListKeywordLibsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListKeywordLibsRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ListKeywordLibsRequest) SetRegionId(v string) *ListKeywordLibsRequest {
	s.RegionId = &v
	return s
}

func (s *ListKeywordLibsRequest) SetTenantCode(v string) *ListKeywordLibsRequest {
	s.TenantCode = &v
	return s
}

func (s *ListKeywordLibsRequest) Validate() error {
	return dara.Validate(s)
}
