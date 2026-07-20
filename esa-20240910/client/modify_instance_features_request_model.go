// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceFeaturesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyInstanceFeaturesRequest
	GetInstanceId() *string
	SetSiteFeatures(v []*ModifyInstanceFeaturesRequestSiteFeatures) *ModifyInstanceFeaturesRequest
	GetSiteFeatures() []*ModifyInstanceFeaturesRequestSiteFeatures
}

type ModifyInstanceFeaturesRequest struct {
	// The plan instance ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// esa-site-b6ga97vfo64g
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of site feature configurations.
	//
	// This parameter is required.
	SiteFeatures []*ModifyInstanceFeaturesRequestSiteFeatures `json:"SiteFeatures,omitempty" xml:"SiteFeatures,omitempty" type:"Repeated"`
}

func (s ModifyInstanceFeaturesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceFeaturesRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceFeaturesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceFeaturesRequest) GetSiteFeatures() []*ModifyInstanceFeaturesRequestSiteFeatures {
	return s.SiteFeatures
}

func (s *ModifyInstanceFeaturesRequest) SetInstanceId(v string) *ModifyInstanceFeaturesRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceFeaturesRequest) SetSiteFeatures(v []*ModifyInstanceFeaturesRequestSiteFeatures) *ModifyInstanceFeaturesRequest {
	s.SiteFeatures = v
	return s
}

func (s *ModifyInstanceFeaturesRequest) Validate() error {
	if s.SiteFeatures != nil {
		for _, item := range s.SiteFeatures {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyInstanceFeaturesRequestSiteFeatures struct {
	// The site feature configurations to modify.
	//
	// example:
	//
	// network_optimization|smart_routing,loadbalance
	Features *string `json:"Features,omitempty" xml:"Features,omitempty"`
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the ID.
	//
	// example:
	//
	// 151097616427232
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ModifyInstanceFeaturesRequestSiteFeatures) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceFeaturesRequestSiteFeatures) GoString() string {
	return s.String()
}

func (s *ModifyInstanceFeaturesRequestSiteFeatures) GetFeatures() *string {
	return s.Features
}

func (s *ModifyInstanceFeaturesRequestSiteFeatures) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ModifyInstanceFeaturesRequestSiteFeatures) SetFeatures(v string) *ModifyInstanceFeaturesRequestSiteFeatures {
	s.Features = &v
	return s
}

func (s *ModifyInstanceFeaturesRequestSiteFeatures) SetSiteId(v int64) *ModifyInstanceFeaturesRequestSiteFeatures {
	s.SiteId = &v
	return s
}

func (s *ModifyInstanceFeaturesRequestSiteFeatures) Validate() error {
	return dara.Validate(s)
}
