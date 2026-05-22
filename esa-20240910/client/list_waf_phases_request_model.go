// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafPhasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *ListWafPhasesRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListWafPhasesRequest
	GetSiteVersion() *int32
}

type ListWafPhasesRequest struct {
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Site version.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListWafPhasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWafPhasesRequest) GoString() string {
	return s.String()
}

func (s *ListWafPhasesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListWafPhasesRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListWafPhasesRequest) SetSiteId(v int64) *ListWafPhasesRequest {
	s.SiteId = &v
	return s
}

func (s *ListWafPhasesRequest) SetSiteVersion(v int32) *ListWafPhasesRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListWafPhasesRequest) Validate() error {
	return dara.Validate(s)
}
