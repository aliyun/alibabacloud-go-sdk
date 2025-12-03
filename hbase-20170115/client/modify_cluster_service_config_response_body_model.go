// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterServiceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyClusterServiceConfigResponseBody
	GetRequestId() *string
}

type ModifyClusterServiceConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterServiceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterServiceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterServiceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyClusterServiceConfigResponseBody) SetRequestId(v string) *ModifyClusterServiceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterServiceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
