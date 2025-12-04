// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenLdpsColumnarIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *OpenLdpsColumnarIndexResponseBody
	GetAccessDeniedDetail() *string
	SetRequestId(v string) *OpenLdpsColumnarIndexResponseBody
	GetRequestId() *string
}

type OpenLdpsColumnarIndexResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenLdpsColumnarIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenLdpsColumnarIndexResponseBody) GoString() string {
	return s.String()
}

func (s *OpenLdpsColumnarIndexResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *OpenLdpsColumnarIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenLdpsColumnarIndexResponseBody) SetAccessDeniedDetail(v string) *OpenLdpsColumnarIndexResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *OpenLdpsColumnarIndexResponseBody) SetRequestId(v string) *OpenLdpsColumnarIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenLdpsColumnarIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
