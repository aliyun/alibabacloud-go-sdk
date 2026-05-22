// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCacheTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCaseInsensitive(v string) *GetCacheTagResponseBody
	GetCaseInsensitive() *string
	SetRequestId(v string) *GetCacheTagResponseBody
	GetRequestId() *string
	SetSiteVersion(v int32) *GetCacheTagResponseBody
	GetSiteVersion() *int32
	SetTagName(v string) *GetCacheTagResponseBody
	GetTagName() *string
}

type GetCacheTagResponseBody struct {
	CaseInsensitive *string `json:"CaseInsensitive,omitempty" xml:"CaseInsensitive,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SiteVersion     *int32  `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	TagName         *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s GetCacheTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCacheTagResponseBody) GoString() string {
	return s.String()
}

func (s *GetCacheTagResponseBody) GetCaseInsensitive() *string {
	return s.CaseInsensitive
}

func (s *GetCacheTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCacheTagResponseBody) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetCacheTagResponseBody) GetTagName() *string {
	return s.TagName
}

func (s *GetCacheTagResponseBody) SetCaseInsensitive(v string) *GetCacheTagResponseBody {
	s.CaseInsensitive = &v
	return s
}

func (s *GetCacheTagResponseBody) SetRequestId(v string) *GetCacheTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCacheTagResponseBody) SetSiteVersion(v int32) *GetCacheTagResponseBody {
	s.SiteVersion = &v
	return s
}

func (s *GetCacheTagResponseBody) SetTagName(v string) *GetCacheTagResponseBody {
	s.TagName = &v
	return s
}

func (s *GetCacheTagResponseBody) Validate() error {
	return dara.Validate(s)
}
