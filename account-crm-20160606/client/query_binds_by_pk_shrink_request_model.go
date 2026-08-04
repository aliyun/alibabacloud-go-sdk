// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBindsByPkShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *QueryBindsByPkShrinkRequest
	GetAppName() *string
	SetPk(v string) *QueryBindsByPkShrinkRequest
	GetPk() *string
	SetTenantIdsShrink(v string) *QueryBindsByPkShrinkRequest
	GetTenantIdsShrink() *string
}

type QueryBindsByPkShrinkRequest struct {
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// This parameter is required.
	TenantIdsShrink *string `json:"TenantIds,omitempty" xml:"TenantIds,omitempty"`
}

func (s QueryBindsByPkShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryBindsByPkShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryBindsByPkShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *QueryBindsByPkShrinkRequest) GetPk() *string {
	return s.Pk
}

func (s *QueryBindsByPkShrinkRequest) GetTenantIdsShrink() *string {
	return s.TenantIdsShrink
}

func (s *QueryBindsByPkShrinkRequest) SetAppName(v string) *QueryBindsByPkShrinkRequest {
	s.AppName = &v
	return s
}

func (s *QueryBindsByPkShrinkRequest) SetPk(v string) *QueryBindsByPkShrinkRequest {
	s.Pk = &v
	return s
}

func (s *QueryBindsByPkShrinkRequest) SetTenantIdsShrink(v string) *QueryBindsByPkShrinkRequest {
	s.TenantIdsShrink = &v
	return s
}

func (s *QueryBindsByPkShrinkRequest) Validate() error {
	return dara.Validate(s)
}
