// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRetcodeAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListRetcodeAppsRequest
	GetRegionId() *string
	SetSecurityToken(v string) *ListRetcodeAppsRequest
	GetSecurityToken() *string
}

type ListRetcodeAppsRequest struct {
	// This parameter is required.
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ListRetcodeAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRetcodeAppsRequest) GoString() string {
	return s.String()
}

func (s *ListRetcodeAppsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListRetcodeAppsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ListRetcodeAppsRequest) SetRegionId(v string) *ListRetcodeAppsRequest {
	s.RegionId = &v
	return s
}

func (s *ListRetcodeAppsRequest) SetSecurityToken(v string) *ListRetcodeAppsRequest {
	s.SecurityToken = &v
	return s
}

func (s *ListRetcodeAppsRequest) Validate() error {
	return dara.Validate(s)
}
