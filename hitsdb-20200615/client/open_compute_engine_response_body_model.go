// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenComputeEngineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *OpenComputeEngineResponseBody
	GetAccessDeniedDetail() *string
	SetRequestId(v string) *OpenComputeEngineResponseBody
	GetRequestId() *string
}

type OpenComputeEngineResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenComputeEngineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenComputeEngineResponseBody) GoString() string {
	return s.String()
}

func (s *OpenComputeEngineResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *OpenComputeEngineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenComputeEngineResponseBody) SetAccessDeniedDetail(v string) *OpenComputeEngineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *OpenComputeEngineResponseBody) SetRequestId(v string) *OpenComputeEngineResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenComputeEngineResponseBody) Validate() error {
	return dara.Validate(s)
}
