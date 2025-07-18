// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnterpriseAccelerateTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEnterpriseAccelerateTargetResponseBody
	GetRequestId() *string
}

type DeleteEnterpriseAccelerateTargetResponseBody struct {
	// example:
	//
	// 655CE28F-2C0C-5801-A31E-C16BF54BD225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEnterpriseAccelerateTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnterpriseAccelerateTargetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseAccelerateTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEnterpriseAccelerateTargetResponseBody) SetRequestId(v string) *DeleteEnterpriseAccelerateTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEnterpriseAccelerateTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
