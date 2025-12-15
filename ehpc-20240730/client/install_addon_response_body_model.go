// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAddonResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddonId(v string) *InstallAddonResponseBody
	GetAddonId() *string
	SetClusterId(v string) *InstallAddonResponseBody
	GetClusterId() *string
	SetRequestId(v string) *InstallAddonResponseBody
	GetRequestId() *string
}

type InstallAddonResponseBody struct {
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
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B745C159-3155-4B94-95D0-4B73D4D2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallAddonResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallAddonResponseBody) GoString() string {
	return s.String()
}

func (s *InstallAddonResponseBody) GetAddonId() *string {
	return s.AddonId
}

func (s *InstallAddonResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *InstallAddonResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallAddonResponseBody) SetAddonId(v string) *InstallAddonResponseBody {
	s.AddonId = &v
	return s
}

func (s *InstallAddonResponseBody) SetClusterId(v string) *InstallAddonResponseBody {
	s.ClusterId = &v
	return s
}

func (s *InstallAddonResponseBody) SetRequestId(v string) *InstallAddonResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallAddonResponseBody) Validate() error {
	return dara.Validate(s)
}
