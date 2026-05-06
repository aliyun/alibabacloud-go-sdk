// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKeylessServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteKeylessServerRequest
	GetId() *string
	SetSiteId(v int64) *DeleteKeylessServerRequest
	GetSiteId() *int64
}

type DeleteKeylessServerRequest struct {
	// Keyless server ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// baba39055622c008b90285a8838e****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11223***
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteKeylessServerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKeylessServerRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeylessServerRequest) GetId() *string {
	return s.Id
}

func (s *DeleteKeylessServerRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteKeylessServerRequest) SetId(v string) *DeleteKeylessServerRequest {
	s.Id = &v
	return s
}

func (s *DeleteKeylessServerRequest) SetSiteId(v int64) *DeleteKeylessServerRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteKeylessServerRequest) Validate() error {
	return dara.Validate(s)
}
