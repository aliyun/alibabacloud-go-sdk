// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteVanityNSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *UpdateSiteVanityNSRequest
	GetSiteId() *int64
	SetVanityNSList(v string) *UpdateSiteVanityNSRequest
	GetVanityNSList() *string
}

type UpdateSiteVanityNSRequest struct {
	// The website ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The custom nameserver names. You can specify two to five custom nameserver names. Separate multiple names with commas (,).
	//
	// example:
	//
	// ns1.example.com,ns2.example.com
	VanityNSList *string `json:"VanityNSList,omitempty" xml:"VanityNSList,omitempty"`
}

func (s UpdateSiteVanityNSRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteVanityNSRequest) GoString() string {
	return s.String()
}

func (s *UpdateSiteVanityNSRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateSiteVanityNSRequest) GetVanityNSList() *string {
	return s.VanityNSList
}

func (s *UpdateSiteVanityNSRequest) SetSiteId(v int64) *UpdateSiteVanityNSRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateSiteVanityNSRequest) SetVanityNSList(v string) *UpdateSiteVanityNSRequest {
	s.VanityNSList = &v
	return s
}

func (s *UpdateSiteVanityNSRequest) Validate() error {
	return dara.Validate(s)
}
