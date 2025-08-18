// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSiteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteSiteRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DeleteSiteRequest
	GetSecurityToken() *string
	SetSiteId(v int64) *DeleteSiteRequest
	GetSiteId() *int64
}

type DeleteSiteRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteSiteRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSiteRequest) GoString() string {
	return s.String()
}

func (s *DeleteSiteRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSiteRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteSiteRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteSiteRequest) SetOwnerId(v int64) *DeleteSiteRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSiteRequest) SetSecurityToken(v string) *DeleteSiteRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteSiteRequest) SetSiteId(v int64) *DeleteSiteRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteSiteRequest) Validate() error {
	return dara.Validate(s)
}
