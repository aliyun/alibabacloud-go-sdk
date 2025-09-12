// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLindormV2InstanceParameterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateLindormV2InstanceParameterResponseBody
	GetAccessDeniedDetail() *string
	SetRequestId(v string) *UpdateLindormV2InstanceParameterResponseBody
	GetRequestId() *string
}

type UpdateLindormV2InstanceParameterResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLindormV2InstanceParameterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLindormV2InstanceParameterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLindormV2InstanceParameterResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateLindormV2InstanceParameterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLindormV2InstanceParameterResponseBody) SetAccessDeniedDetail(v string) *UpdateLindormV2InstanceParameterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateLindormV2InstanceParameterResponseBody) SetRequestId(v string) *UpdateLindormV2InstanceParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLindormV2InstanceParameterResponseBody) Validate() error {
	return dara.Validate(s)
}
