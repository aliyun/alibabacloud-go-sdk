// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPerformanceDataCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetPerformanceDataCollectionRequest
	GetSiteId() *int64
}

type GetPerformanceDataCollectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetPerformanceDataCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPerformanceDataCollectionRequest) GoString() string {
	return s.String()
}

func (s *GetPerformanceDataCollectionRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetPerformanceDataCollectionRequest) SetSiteId(v int64) *GetPerformanceDataCollectionRequest {
	s.SiteId = &v
	return s
}

func (s *GetPerformanceDataCollectionRequest) Validate() error {
	return dara.Validate(s)
}
