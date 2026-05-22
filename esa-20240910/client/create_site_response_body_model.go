// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSiteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNameServerList(v string) *CreateSiteResponseBody
	GetNameServerList() *string
	SetRequestId(v string) *CreateSiteResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *CreateSiteResponseBody
	GetSiteId() *int64
	SetVerifyCode(v string) *CreateSiteResponseBody
	GetVerifyCode() *string
}

type CreateSiteResponseBody struct {
	NameServerList *string `json:"NameServerList,omitempty" xml:"NameServerList,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SiteId         *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	VerifyCode     *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s CreateSiteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSiteResponseBody) GetNameServerList() *string {
	return s.NameServerList
}

func (s *CreateSiteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSiteResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateSiteResponseBody) GetVerifyCode() *string {
	return s.VerifyCode
}

func (s *CreateSiteResponseBody) SetNameServerList(v string) *CreateSiteResponseBody {
	s.NameServerList = &v
	return s
}

func (s *CreateSiteResponseBody) SetRequestId(v string) *CreateSiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSiteResponseBody) SetSiteId(v int64) *CreateSiteResponseBody {
	s.SiteId = &v
	return s
}

func (s *CreateSiteResponseBody) SetVerifyCode(v string) *CreateSiteResponseBody {
	s.VerifyCode = &v
	return s
}

func (s *CreateSiteResponseBody) Validate() error {
	return dara.Validate(s)
}
