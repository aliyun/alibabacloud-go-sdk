// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLdpsComputeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateLdpsComputeGroupResponseBody
	GetAccessDeniedDetail() *string
	SetRequestId(v string) *UpdateLdpsComputeGroupResponseBody
	GetRequestId() *string
}

type UpdateLdpsComputeGroupResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLdpsComputeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLdpsComputeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLdpsComputeGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateLdpsComputeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLdpsComputeGroupResponseBody) SetAccessDeniedDetail(v string) *UpdateLdpsComputeGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateLdpsComputeGroupResponseBody) SetRequestId(v string) *UpdateLdpsComputeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLdpsComputeGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
