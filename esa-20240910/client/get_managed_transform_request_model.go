// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManagedTransformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetManagedTransformRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *GetManagedTransformRequest
	GetSiteVersion() *int32
}

type GetManagedTransformRequest struct {
	// This parameter is required.
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s GetManagedTransformRequest) String() string {
	return dara.Prettify(s)
}

func (s GetManagedTransformRequest) GoString() string {
	return s.String()
}

func (s *GetManagedTransformRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetManagedTransformRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetManagedTransformRequest) SetSiteId(v int64) *GetManagedTransformRequest {
	s.SiteId = &v
	return s
}

func (s *GetManagedTransformRequest) SetSiteVersion(v int32) *GetManagedTransformRequest {
	s.SiteVersion = &v
	return s
}

func (s *GetManagedTransformRequest) Validate() error {
	return dara.Validate(s)
}
