// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCommonBandwidthPackageSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCommonBandwidthPackageSpecResponseBody
	GetRequestId() *string
}

type ModifyCommonBandwidthPackageSpecResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7F129000-F929-4AF5-BE8D-BAE434C795306
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCommonBandwidthPackageSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCommonBandwidthPackageSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCommonBandwidthPackageSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCommonBandwidthPackageSpecResponseBody) SetRequestId(v string) *ModifyCommonBandwidthPackageSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCommonBandwidthPackageSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
