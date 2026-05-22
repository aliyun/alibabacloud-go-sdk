// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteCustomLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetSiteCustomLogResponseBody
	GetConfigId() *int64
	SetIsExist(v bool) *GetSiteCustomLogResponseBody
	GetIsExist() *bool
	SetLogCustomField(v *GetSiteCustomLogResponseBodyLogCustomField) *GetSiteCustomLogResponseBody
	GetLogCustomField() *GetSiteCustomLogResponseBodyLogCustomField
	SetRequestId(v string) *GetSiteCustomLogResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *GetSiteCustomLogResponseBody
	GetSiteId() *int64
}

type GetSiteCustomLogResponseBody struct {
	ConfigId       *int64                                      `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	IsExist        *bool                                       `json:"IsExist,omitempty" xml:"IsExist,omitempty"`
	LogCustomField *GetSiteCustomLogResponseBodyLogCustomField `json:"LogCustomField,omitempty" xml:"LogCustomField,omitempty" type:"Struct"`
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SiteId         *int64                                      `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetSiteCustomLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSiteCustomLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetSiteCustomLogResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetSiteCustomLogResponseBody) GetIsExist() *bool {
	return s.IsExist
}

func (s *GetSiteCustomLogResponseBody) GetLogCustomField() *GetSiteCustomLogResponseBodyLogCustomField {
	return s.LogCustomField
}

func (s *GetSiteCustomLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSiteCustomLogResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetSiteCustomLogResponseBody) SetConfigId(v int64) *GetSiteCustomLogResponseBody {
	s.ConfigId = &v
	return s
}

func (s *GetSiteCustomLogResponseBody) SetIsExist(v bool) *GetSiteCustomLogResponseBody {
	s.IsExist = &v
	return s
}

func (s *GetSiteCustomLogResponseBody) SetLogCustomField(v *GetSiteCustomLogResponseBodyLogCustomField) *GetSiteCustomLogResponseBody {
	s.LogCustomField = v
	return s
}

func (s *GetSiteCustomLogResponseBody) SetRequestId(v string) *GetSiteCustomLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSiteCustomLogResponseBody) SetSiteId(v int64) *GetSiteCustomLogResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetSiteCustomLogResponseBody) Validate() error {
	if s.LogCustomField != nil {
		if err := s.LogCustomField.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSiteCustomLogResponseBodyLogCustomField struct {
	Cookies         []*string `json:"Cookies,omitempty" xml:"Cookies,omitempty" type:"Repeated"`
	RequestHeaders  []*string `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	ResponseHeaders []*string `json:"ResponseHeaders,omitempty" xml:"ResponseHeaders,omitempty" type:"Repeated"`
}

func (s GetSiteCustomLogResponseBodyLogCustomField) String() string {
	return dara.Prettify(s)
}

func (s GetSiteCustomLogResponseBodyLogCustomField) GoString() string {
	return s.String()
}

func (s *GetSiteCustomLogResponseBodyLogCustomField) GetCookies() []*string {
	return s.Cookies
}

func (s *GetSiteCustomLogResponseBodyLogCustomField) GetRequestHeaders() []*string {
	return s.RequestHeaders
}

func (s *GetSiteCustomLogResponseBodyLogCustomField) GetResponseHeaders() []*string {
	return s.ResponseHeaders
}

func (s *GetSiteCustomLogResponseBodyLogCustomField) SetCookies(v []*string) *GetSiteCustomLogResponseBodyLogCustomField {
	s.Cookies = v
	return s
}

func (s *GetSiteCustomLogResponseBodyLogCustomField) SetRequestHeaders(v []*string) *GetSiteCustomLogResponseBodyLogCustomField {
	s.RequestHeaders = v
	return s
}

func (s *GetSiteCustomLogResponseBodyLogCustomField) SetResponseHeaders(v []*string) *GetSiteCustomLogResponseBodyLogCustomField {
	s.ResponseHeaders = v
	return s
}

func (s *GetSiteCustomLogResponseBodyLogCustomField) Validate() error {
	return dara.Validate(s)
}
