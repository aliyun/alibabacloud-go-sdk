// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifySiteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *VerifySiteRequest
	GetSiteId() *int64
}

type VerifySiteRequest struct {
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s VerifySiteRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifySiteRequest) GoString() string {
	return s.String()
}

func (s *VerifySiteRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *VerifySiteRequest) SetSiteId(v int64) *VerifySiteRequest {
	s.SiteId = &v
	return s
}

func (s *VerifySiteRequest) Validate() error {
	return dara.Validate(s)
}
