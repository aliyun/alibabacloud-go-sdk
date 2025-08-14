// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCenBandwidthPackageSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCenBandwidthPackageSpecResponseBody
	GetRequestId() *string
}

type ModifyCenBandwidthPackageSpecResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0C2EE7A8-74D4-4081-8236-CEBDE3BBCF50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCenBandwidthPackageSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenBandwidthPackageSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCenBandwidthPackageSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCenBandwidthPackageSpecResponseBody) SetRequestId(v string) *ModifyCenBandwidthPackageSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCenBandwidthPackageSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
