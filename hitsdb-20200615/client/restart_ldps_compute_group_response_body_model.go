// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartLdpsComputeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RestartLdpsComputeGroupResponseBody
	GetAccessDeniedDetail() *string
	SetRequestId(v string) *RestartLdpsComputeGroupResponseBody
	GetRequestId() *string
}

type RestartLdpsComputeGroupResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartLdpsComputeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartLdpsComputeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RestartLdpsComputeGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RestartLdpsComputeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartLdpsComputeGroupResponseBody) SetAccessDeniedDetail(v string) *RestartLdpsComputeGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RestartLdpsComputeGroupResponseBody) SetRequestId(v string) *RestartLdpsComputeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartLdpsComputeGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
