// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseInstancePublicConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseInstancePublicConnectionResponseBody
	GetRequestId() *string
}

type ReleaseInstancePublicConnectionResponseBody struct {
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseInstancePublicConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseInstancePublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseInstancePublicConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseInstancePublicConnectionResponseBody) SetRequestId(v string) *ReleaseInstancePublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseInstancePublicConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
