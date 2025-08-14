// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCenBandwidthPackageAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCenBandwidthPackageAttributeResponseBody
	GetRequestId() *string
}

type ModifyCenBandwidthPackageAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 13526224-5780-4426-8ADF-BC8B08700F23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCenBandwidthPackageAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenBandwidthPackageAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCenBandwidthPackageAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCenBandwidthPackageAttributeResponseBody) SetRequestId(v string) *ModifyCenBandwidthPackageAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCenBandwidthPackageAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
