// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAdditionalVpcLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateAdditionalVpcLinkResponseBody
	GetAccessDeniedDetail() *string
	SetRequestId(v string) *CreateAdditionalVpcLinkResponseBody
	GetRequestId() *string
}

type CreateAdditionalVpcLinkResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAdditionalVpcLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAdditionalVpcLinkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAdditionalVpcLinkResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateAdditionalVpcLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAdditionalVpcLinkResponseBody) SetAccessDeniedDetail(v string) *CreateAdditionalVpcLinkResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateAdditionalVpcLinkResponseBody) SetRequestId(v string) *CreateAdditionalVpcLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAdditionalVpcLinkResponseBody) Validate() error {
	return dara.Validate(s)
}
