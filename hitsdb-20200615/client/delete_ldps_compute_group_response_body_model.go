// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLdpsComputeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteLdpsComputeGroupResponseBody
	GetAccessDeniedDetail() *string
	SetRequestId(v string) *DeleteLdpsComputeGroupResponseBody
	GetRequestId() *string
}

type DeleteLdpsComputeGroupResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLdpsComputeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLdpsComputeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLdpsComputeGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteLdpsComputeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLdpsComputeGroupResponseBody) SetAccessDeniedDetail(v string) *DeleteLdpsComputeGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteLdpsComputeGroupResponseBody) SetRequestId(v string) *DeleteLdpsComputeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLdpsComputeGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
