// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePerformanceDataCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v string) *UpdatePerformanceDataCollectionRequest
	GetEnable() *string
	SetSiteId(v int64) *UpdatePerformanceDataCollectionRequest
	GetSiteId() *int64
}

type UpdatePerformanceDataCollectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 34003500310****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdatePerformanceDataCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePerformanceDataCollectionRequest) GoString() string {
	return s.String()
}

func (s *UpdatePerformanceDataCollectionRequest) GetEnable() *string {
	return s.Enable
}

func (s *UpdatePerformanceDataCollectionRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdatePerformanceDataCollectionRequest) SetEnable(v string) *UpdatePerformanceDataCollectionRequest {
	s.Enable = &v
	return s
}

func (s *UpdatePerformanceDataCollectionRequest) SetSiteId(v int64) *UpdatePerformanceDataCollectionRequest {
	s.SiteId = &v
	return s
}

func (s *UpdatePerformanceDataCollectionRequest) Validate() error {
	return dara.Validate(s)
}
