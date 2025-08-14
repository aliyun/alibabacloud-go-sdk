// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateCenBandwidthPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateCenBandwidthPackageResponseBody
	GetRequestId() *string
}

type AssociateCenBandwidthPackageResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0C2EE7A8-74D4-4081-8236-CEBDE3BBCF50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateCenBandwidthPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateCenBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateCenBandwidthPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateCenBandwidthPackageResponseBody) SetRequestId(v string) *AssociateCenBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateCenBandwidthPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
