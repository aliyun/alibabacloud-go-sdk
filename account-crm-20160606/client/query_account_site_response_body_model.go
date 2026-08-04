// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountSiteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAccountSiteResponseBody
	GetCode() *string
	SetRequestId(v string) *QueryAccountSiteResponseBody
	GetRequestId() *string
	SetSite(v string) *QueryAccountSiteResponseBody
	GetSite() *string
	SetSuccess(v bool) *QueryAccountSiteResponseBody
	GetSuccess() *bool
}

type QueryAccountSiteResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Site      *string `json:"Site,omitempty" xml:"Site,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAccountSiteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountSiteResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccountSiteResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAccountSiteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAccountSiteResponseBody) GetSite() *string {
	return s.Site
}

func (s *QueryAccountSiteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAccountSiteResponseBody) SetCode(v string) *QueryAccountSiteResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAccountSiteResponseBody) SetRequestId(v string) *QueryAccountSiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAccountSiteResponseBody) SetSite(v string) *QueryAccountSiteResponseBody {
	s.Site = &v
	return s
}

func (s *QueryAccountSiteResponseBody) SetSuccess(v bool) *QueryAccountSiteResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAccountSiteResponseBody) Validate() error {
	return dara.Validate(s)
}
