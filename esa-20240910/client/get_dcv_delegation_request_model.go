// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDcvDelegationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetDcvDelegationRequest
	GetSiteId() *int64
}

type GetDcvDelegationRequest struct {
	// The site ID. You can obtain the site ID by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetDcvDelegationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDcvDelegationRequest) GoString() string {
	return s.String()
}

func (s *GetDcvDelegationRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetDcvDelegationRequest) SetSiteId(v int64) *GetDcvDelegationRequest {
	s.SiteId = &v
	return s
}

func (s *GetDcvDelegationRequest) Validate() error {
	return dara.Validate(s)
}
