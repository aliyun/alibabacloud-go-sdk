// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTieredCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTieredCacheResponseBody
	GetRequestId() *string
}

type UpdateTieredCacheResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTieredCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTieredCacheResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTieredCacheResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTieredCacheResponseBody) SetRequestId(v string) *UpdateTieredCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTieredCacheResponseBody) Validate() error {
	return dara.Validate(s)
}
