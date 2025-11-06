// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiSchemaUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetApiSchemaUsageRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *GetApiSchemaUsageRequest
	GetSiteVersion() *int32
}

type GetApiSchemaUsageRequest struct {
	// example:
	//
	// 1159101787****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s GetApiSchemaUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApiSchemaUsageRequest) GoString() string {
	return s.String()
}

func (s *GetApiSchemaUsageRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetApiSchemaUsageRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetApiSchemaUsageRequest) SetSiteId(v int64) *GetApiSchemaUsageRequest {
	s.SiteId = &v
	return s
}

func (s *GetApiSchemaUsageRequest) SetSiteVersion(v int32) *GetApiSchemaUsageRequest {
	s.SiteVersion = &v
	return s
}

func (s *GetApiSchemaUsageRequest) Validate() error {
	return dara.Validate(s)
}
