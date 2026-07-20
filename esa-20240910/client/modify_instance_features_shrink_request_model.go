// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceFeaturesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyInstanceFeaturesShrinkRequest
	GetInstanceId() *string
	SetSiteFeaturesShrink(v string) *ModifyInstanceFeaturesShrinkRequest
	GetSiteFeaturesShrink() *string
}

type ModifyInstanceFeaturesShrinkRequest struct {
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
	SiteFeaturesShrink *string `json:"SiteFeatures,omitempty" xml:"SiteFeatures,omitempty"`
}

func (s ModifyInstanceFeaturesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceFeaturesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceFeaturesShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceFeaturesShrinkRequest) GetSiteFeaturesShrink() *string {
	return s.SiteFeaturesShrink
}

func (s *ModifyInstanceFeaturesShrinkRequest) SetInstanceId(v string) *ModifyInstanceFeaturesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceFeaturesShrinkRequest) SetSiteFeaturesShrink(v string) *ModifyInstanceFeaturesShrinkRequest {
	s.SiteFeaturesShrink = &v
	return s
}

func (s *ModifyInstanceFeaturesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
