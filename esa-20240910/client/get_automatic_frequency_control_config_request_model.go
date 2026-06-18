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
	// The ID of the site. Call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version of the site. If versioning is enabled, use this parameter to specify the site version. The default value is 0.
	//
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
