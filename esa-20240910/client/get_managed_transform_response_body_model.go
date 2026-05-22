// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManagedTransformResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddClientGeolocationHeader(v string) *GetManagedTransformResponseBody
	GetAddClientGeolocationHeader() *string
	SetAddRealClientIpHeader(v string) *GetManagedTransformResponseBody
	GetAddRealClientIpHeader() *string
	SetRealClientIpHeaderName(v string) *GetManagedTransformResponseBody
	GetRealClientIpHeaderName() *string
	SetRequestId(v string) *GetManagedTransformResponseBody
	GetRequestId() *string
	SetSiteVersion(v int32) *GetManagedTransformResponseBody
	GetSiteVersion() *int32
}

type GetManagedTransformResponseBody struct {
	AddClientGeolocationHeader *string `json:"AddClientGeolocationHeader,omitempty" xml:"AddClientGeolocationHeader,omitempty"`
	AddRealClientIpHeader      *string `json:"AddRealClientIpHeader,omitempty" xml:"AddRealClientIpHeader,omitempty"`
	RealClientIpHeaderName     *string `json:"RealClientIpHeaderName,omitempty" xml:"RealClientIpHeaderName,omitempty"`
	RequestId                  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SiteVersion                *int32  `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s GetManagedTransformResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetManagedTransformResponseBody) GoString() string {
	return s.String()
}

func (s *GetManagedTransformResponseBody) GetAddClientGeolocationHeader() *string {
	return s.AddClientGeolocationHeader
}

func (s *GetManagedTransformResponseBody) GetAddRealClientIpHeader() *string {
	return s.AddRealClientIpHeader
}

func (s *GetManagedTransformResponseBody) GetRealClientIpHeaderName() *string {
	return s.RealClientIpHeaderName
}

func (s *GetManagedTransformResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetManagedTransformResponseBody) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetManagedTransformResponseBody) SetAddClientGeolocationHeader(v string) *GetManagedTransformResponseBody {
	s.AddClientGeolocationHeader = &v
	return s
}

func (s *GetManagedTransformResponseBody) SetAddRealClientIpHeader(v string) *GetManagedTransformResponseBody {
	s.AddRealClientIpHeader = &v
	return s
}

func (s *GetManagedTransformResponseBody) SetRealClientIpHeaderName(v string) *GetManagedTransformResponseBody {
	s.RealClientIpHeaderName = &v
	return s
}

func (s *GetManagedTransformResponseBody) SetRequestId(v string) *GetManagedTransformResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetManagedTransformResponseBody) SetSiteVersion(v int32) *GetManagedTransformResponseBody {
	s.SiteVersion = &v
	return s
}

func (s *GetManagedTransformResponseBody) Validate() error {
	return dara.Validate(s)
}
