// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySiteFeaturesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNewInstanceId(v string) *ModifySiteFeaturesRequest
	GetNewInstanceId() *string
	SetSiteFeatures(v string) *ModifySiteFeaturesRequest
	GetSiteFeatures() *string
	SetSiteId(v int64) *ModifySiteFeaturesRequest
	GetSiteId() *int64
}

type ModifySiteFeaturesRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// esa-site-bl39ryjtineo
	NewInstanceId *string `json:"NewInstanceId,omitempty" xml:"NewInstanceId,omitempty"`
	// The site feature information to be cleared.
	//
	// This parameter is required.
	//
	// example:
	//
	// network_optimization|smart_routing,loadbalance
	SiteFeatures *string `json:"SiteFeatures,omitempty" xml:"SiteFeatures,omitempty"`
	// The site ID. You can obtain the ID by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1067072706415168
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ModifySiteFeaturesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySiteFeaturesRequest) GoString() string {
	return s.String()
}

func (s *ModifySiteFeaturesRequest) GetNewInstanceId() *string {
	return s.NewInstanceId
}

func (s *ModifySiteFeaturesRequest) GetSiteFeatures() *string {
	return s.SiteFeatures
}

func (s *ModifySiteFeaturesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ModifySiteFeaturesRequest) SetNewInstanceId(v string) *ModifySiteFeaturesRequest {
	s.NewInstanceId = &v
	return s
}

func (s *ModifySiteFeaturesRequest) SetSiteFeatures(v string) *ModifySiteFeaturesRequest {
	s.SiteFeatures = &v
	return s
}

func (s *ModifySiteFeaturesRequest) SetSiteId(v int64) *ModifySiteFeaturesRequest {
	s.SiteId = &v
	return s
}

func (s *ModifySiteFeaturesRequest) Validate() error {
	return dara.Validate(s)
}
