// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenComputePreCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *OpenComputePreCheckResponseBody
	GetAccessDeniedDetail() *string
	SetRequestId(v string) *OpenComputePreCheckResponseBody
	GetRequestId() *string
}

type OpenComputePreCheckResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenComputePreCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenComputePreCheckResponseBody) GoString() string {
	return s.String()
}

func (s *OpenComputePreCheckResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *OpenComputePreCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenComputePreCheckResponseBody) SetAccessDeniedDetail(v string) *OpenComputePreCheckResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *OpenComputePreCheckResponseBody) SetRequestId(v string) *OpenComputePreCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenComputePreCheckResponseBody) Validate() error {
	return dara.Validate(s)
}
