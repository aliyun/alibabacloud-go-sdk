// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransportLayerApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v int64) *DeleteTransportLayerApplicationRequest
	GetApplicationId() *int64
	SetSiteId(v int64) *DeleteTransportLayerApplicationRequest
	GetSiteId() *int64
}

type DeleteTransportLayerApplicationRequest struct {
	// Application ID, which can be obtained by calling the [ListTransportLayerApplications](~~ListTransportLayerApplications~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 170996390868****
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteTransportLayerApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransportLayerApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteTransportLayerApplicationRequest) GetApplicationId() *int64 {
	return s.ApplicationId
}

func (s *DeleteTransportLayerApplicationRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteTransportLayerApplicationRequest) SetApplicationId(v int64) *DeleteTransportLayerApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *DeleteTransportLayerApplicationRequest) SetSiteId(v int64) *DeleteTransportLayerApplicationRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteTransportLayerApplicationRequest) Validate() error {
	return dara.Validate(s)
}
