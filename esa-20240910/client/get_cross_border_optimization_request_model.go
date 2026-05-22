// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCrossBorderOptimizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetCrossBorderOptimizationRequest
	GetSiteId() *int64
}

type GetCrossBorderOptimizationRequest struct {
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 340035003106221
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetCrossBorderOptimizationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCrossBorderOptimizationRequest) GoString() string {
	return s.String()
}

func (s *GetCrossBorderOptimizationRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetCrossBorderOptimizationRequest) SetSiteId(v int64) *GetCrossBorderOptimizationRequest {
	s.SiteId = &v
	return s
}

func (s *GetCrossBorderOptimizationRequest) Validate() error {
	return dara.Validate(s)
}
