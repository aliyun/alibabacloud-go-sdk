// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCommonBandwidthPackageAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCommonBandwidthPackageAttributeResponseBody
	GetRequestId() *string
}

type ModifyCommonBandwidthPackageAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B450CAD8-50BC-4506-ADA7-35C6CE63E96B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCommonBandwidthPackageAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCommonBandwidthPackageAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCommonBandwidthPackageAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCommonBandwidthPackageAttributeResponseBody) SetRequestId(v string) *ModifyCommonBandwidthPackageAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCommonBandwidthPackageAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
