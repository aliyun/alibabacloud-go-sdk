// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCnameFlatteningRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetCnameFlatteningRequest
	GetSiteId() *int64
}

type GetCnameFlatteningRequest struct {
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetCnameFlatteningRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCnameFlatteningRequest) GoString() string {
	return s.String()
}

func (s *GetCnameFlatteningRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetCnameFlatteningRequest) SetSiteId(v int64) *GetCnameFlatteningRequest {
	s.SiteId = &v
	return s
}

func (s *GetCnameFlatteningRequest) Validate() error {
	return dara.Validate(s)
}
