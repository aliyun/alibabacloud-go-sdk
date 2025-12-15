// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnInstallAddonRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonId(v string) *UnInstallAddonRequest
	GetAddonId() *string
	SetClusterId(v string) *UnInstallAddonRequest
	GetClusterId() *string
}

type UnInstallAddonRequest struct {
	// The addon ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login-1.0-W2g****
	AddonId *string `json:"AddonId,omitempty" xml:"AddonId,omitempty"`
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s UnInstallAddonRequest) String() string {
	return dara.Prettify(s)
}

func (s UnInstallAddonRequest) GoString() string {
	return s.String()
}

func (s *UnInstallAddonRequest) GetAddonId() *string {
	return s.AddonId
}

func (s *UnInstallAddonRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UnInstallAddonRequest) SetAddonId(v string) *UnInstallAddonRequest {
	s.AddonId = &v
	return s
}

func (s *UnInstallAddonRequest) SetClusterId(v string) *UnInstallAddonRequest {
	s.ClusterId = &v
	return s
}

func (s *UnInstallAddonRequest) Validate() error {
	return dara.Validate(s)
}
