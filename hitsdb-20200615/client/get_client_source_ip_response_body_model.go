// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientSourceIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetClientSourceIpResponseBody
	GetAccessDeniedDetail() *string
	SetClientIp(v string) *GetClientSourceIpResponseBody
	GetClientIp() *string
	SetRequestId(v string) *GetClientSourceIpResponseBody
	GetRequestId() *string
}

type GetClientSourceIpResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	ClientIp           *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetClientSourceIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClientSourceIpResponseBody) GoString() string {
	return s.String()
}

func (s *GetClientSourceIpResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetClientSourceIpResponseBody) GetClientIp() *string {
	return s.ClientIp
}

func (s *GetClientSourceIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClientSourceIpResponseBody) SetAccessDeniedDetail(v string) *GetClientSourceIpResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetClientSourceIpResponseBody) SetClientIp(v string) *GetClientSourceIpResponseBody {
	s.ClientIp = &v
	return s
}

func (s *GetClientSourceIpResponseBody) SetRequestId(v string) *GetClientSourceIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClientSourceIpResponseBody) Validate() error {
	return dara.Validate(s)
}
