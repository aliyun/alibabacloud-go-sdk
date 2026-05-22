// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPurgeQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetPurgeQuotaRequest
	GetSiteId() *int64
	SetType(v string) *GetPurgeQuotaRequest
	GetType() *string
}

type GetPurgeQuotaRequest struct {
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetPurgeQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPurgeQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetPurgeQuotaRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetPurgeQuotaRequest) GetType() *string {
	return s.Type
}

func (s *GetPurgeQuotaRequest) SetSiteId(v int64) *GetPurgeQuotaRequest {
	s.SiteId = &v
	return s
}

func (s *GetPurgeQuotaRequest) SetType(v string) *GetPurgeQuotaRequest {
	s.Type = &v
	return s
}

func (s *GetPurgeQuotaRequest) Validate() error {
	return dara.Validate(s)
}
