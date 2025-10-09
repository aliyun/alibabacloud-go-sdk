// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetRoutineRouteRequest
	GetConfigId() *int64
	SetSiteId(v int64) *GetRoutineRouteRequest
	GetSiteId() *int64
}

type GetRoutineRouteRequest struct {
	// The ID of the configuration. You can call the [ListSiteRoutes](https://help.aliyun.com/document_detail/2879631.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetRoutineRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineRouteRequest) GoString() string {
	return s.String()
}

func (s *GetRoutineRouteRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetRoutineRouteRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetRoutineRouteRequest) SetConfigId(v int64) *GetRoutineRouteRequest {
	s.ConfigId = &v
	return s
}

func (s *GetRoutineRouteRequest) SetSiteId(v int64) *GetRoutineRouteRequest {
	s.SiteId = &v
	return s
}

func (s *GetRoutineRouteRequest) Validate() error {
	return dara.Validate(s)
}
