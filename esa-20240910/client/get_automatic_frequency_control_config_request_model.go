// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutomaticFrequencyControlConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetAutomaticFrequencyControlConfigRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *GetAutomaticFrequencyControlConfigRequest
	GetSiteVersion() *int32
}

type GetAutomaticFrequencyControlConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s GetAutomaticFrequencyControlConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAutomaticFrequencyControlConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAutomaticFrequencyControlConfigRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetAutomaticFrequencyControlConfigRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetAutomaticFrequencyControlConfigRequest) SetSiteId(v int64) *GetAutomaticFrequencyControlConfigRequest {
	s.SiteId = &v
	return s
}

func (s *GetAutomaticFrequencyControlConfigRequest) SetSiteVersion(v int32) *GetAutomaticFrequencyControlConfigRequest {
	s.SiteVersion = &v
	return s
}

func (s *GetAutomaticFrequencyControlConfigRequest) Validate() error {
	return dara.Validate(s)
}
