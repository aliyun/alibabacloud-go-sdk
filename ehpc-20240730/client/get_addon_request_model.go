// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAddonRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonId(v string) *GetAddonRequest
	GetAddonId() *string
	SetClusterId(v string) *GetAddonRequest
	GetClusterId() *string
}

type GetAddonRequest struct {
	// The addon ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login-1.0-W4g****
	AddonId *string `json:"AddonId,omitempty" xml:"AddonId,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetAddonRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAddonRequest) GoString() string {
	return s.String()
}

func (s *GetAddonRequest) GetAddonId() *string {
	return s.AddonId
}

func (s *GetAddonRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetAddonRequest) SetAddonId(v string) *GetAddonRequest {
	s.AddonId = &v
	return s
}

func (s *GetAddonRequest) SetClusterId(v string) *GetAddonRequest {
	s.ClusterId = &v
	return s
}

func (s *GetAddonRequest) Validate() error {
	return dara.Validate(s)
}
