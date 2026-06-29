// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSOverseasAttackCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *DescribeDDoSOverseasAttackCountRequest
	GetSiteId() *int64
}

type DescribeDDoSOverseasAttackCountRequest struct {
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1039004047346160
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DescribeDDoSOverseasAttackCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSOverseasAttackCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDoSOverseasAttackCountRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribeDDoSOverseasAttackCountRequest) SetSiteId(v int64) *DescribeDDoSOverseasAttackCountRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeDDoSOverseasAttackCountRequest) Validate() error {
	return dara.Validate(s)
}
