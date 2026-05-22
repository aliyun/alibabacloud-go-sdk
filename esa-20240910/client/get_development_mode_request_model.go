// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDevelopmentModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetDevelopmentModeRequest
	GetSiteId() *int64
}

type GetDevelopmentModeRequest struct {
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5407498413****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetDevelopmentModeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDevelopmentModeRequest) GoString() string {
	return s.String()
}

func (s *GetDevelopmentModeRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetDevelopmentModeRequest) SetSiteId(v int64) *GetDevelopmentModeRequest {
	s.SiteId = &v
	return s
}

func (s *GetDevelopmentModeRequest) Validate() error {
	return dara.Validate(s)
}
