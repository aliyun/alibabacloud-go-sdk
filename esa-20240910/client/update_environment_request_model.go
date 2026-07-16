// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvironmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentName(v string) *UpdateEnvironmentRequest
	GetEnvironmentName() *string
	SetNewName(v string) *UpdateEnvironmentRequest
	GetNewName() *string
	SetReadOnly(v bool) *UpdateEnvironmentRequest
	GetReadOnly() *bool
	SetRule(v string) *UpdateEnvironmentRequest
	GetRule() *string
	SetSiteId(v int64) *UpdateEnvironmentRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *UpdateEnvironmentRequest
	GetSiteVersion() *int32
}

type UpdateEnvironmentRequest struct {
	// The environment name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 默认环境
	EnvironmentName *string `json:"EnvironmentName,omitempty" xml:"EnvironmentName,omitempty"`
	// The new environment name.
	//
	// example:
	//
	// 环境A
	NewName *string `json:"NewName,omitempty" xml:"NewName,omitempty"`
	// Specifies whether the environment is read-only.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The Wireshark rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// (http.host eq "duduko5.top")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 33993121955****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The new site version number. Only the environment with the highest priority can be modified.
	//
	// example:
	//
	// 10
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s UpdateEnvironmentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentRequest) GetEnvironmentName() *string {
	return s.EnvironmentName
}

func (s *UpdateEnvironmentRequest) GetNewName() *string {
	return s.NewName
}

func (s *UpdateEnvironmentRequest) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *UpdateEnvironmentRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateEnvironmentRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateEnvironmentRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *UpdateEnvironmentRequest) SetEnvironmentName(v string) *UpdateEnvironmentRequest {
	s.EnvironmentName = &v
	return s
}

func (s *UpdateEnvironmentRequest) SetNewName(v string) *UpdateEnvironmentRequest {
	s.NewName = &v
	return s
}

func (s *UpdateEnvironmentRequest) SetReadOnly(v bool) *UpdateEnvironmentRequest {
	s.ReadOnly = &v
	return s
}

func (s *UpdateEnvironmentRequest) SetRule(v string) *UpdateEnvironmentRequest {
	s.Rule = &v
	return s
}

func (s *UpdateEnvironmentRequest) SetSiteId(v int64) *UpdateEnvironmentRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateEnvironmentRequest) SetSiteVersion(v int32) *UpdateEnvironmentRequest {
	s.SiteVersion = &v
	return s
}

func (s *UpdateEnvironmentRequest) Validate() error {
	return dara.Validate(s)
}
