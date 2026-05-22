// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKeylessServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetKeylessServerRequest
	GetId() *string
	SetSiteId(v int64) *GetKeylessServerRequest
	GetSiteId() *int64
}

type GetKeylessServerRequest struct {
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
	// 5407498413****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetKeylessServerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKeylessServerRequest) GoString() string {
	return s.String()
}

func (s *GetKeylessServerRequest) GetId() *string {
	return s.Id
}

func (s *GetKeylessServerRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetKeylessServerRequest) SetId(v string) *GetKeylessServerRequest {
	s.Id = &v
	return s
}

func (s *GetKeylessServerRequest) SetSiteId(v int64) *GetKeylessServerRequest {
	s.SiteId = &v
	return s
}

func (s *GetKeylessServerRequest) Validate() error {
	return dara.Validate(s)
}
