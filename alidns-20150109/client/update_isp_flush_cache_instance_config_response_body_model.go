// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIspFlushCacheInstanceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateIspFlushCacheInstanceConfigResponseBody
	GetRequestId() *string
}

type UpdateIspFlushCacheInstanceConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIspFlushCacheInstanceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIspFlushCacheInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIspFlushCacheInstanceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIspFlushCacheInstanceConfigResponseBody) SetRequestId(v string) *UpdateIspFlushCacheInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIspFlushCacheInstanceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
