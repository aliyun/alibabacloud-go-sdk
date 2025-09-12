// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLdpsComputeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateLdpsComputeGroupResponseBody
	GetAccessDeniedDetail() *string
	SetRequestId(v string) *CreateLdpsComputeGroupResponseBody
	GetRequestId() *string
}

type CreateLdpsComputeGroupResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLdpsComputeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLdpsComputeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLdpsComputeGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateLdpsComputeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLdpsComputeGroupResponseBody) SetAccessDeniedDetail(v string) *CreateLdpsComputeGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateLdpsComputeGroupResponseBody) SetRequestId(v string) *CreateLdpsComputeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLdpsComputeGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
