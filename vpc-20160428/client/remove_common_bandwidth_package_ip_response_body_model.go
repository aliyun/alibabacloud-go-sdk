// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveCommonBandwidthPackageIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveCommonBandwidthPackageIpResponseBody
	GetRequestId() *string
}

type RemoveCommonBandwidthPackageIpResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveCommonBandwidthPackageIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveCommonBandwidthPackageIpResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveCommonBandwidthPackageIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveCommonBandwidthPackageIpResponseBody) SetRequestId(v string) *RemoveCommonBandwidthPackageIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveCommonBandwidthPackageIpResponseBody) Validate() error {
	return dara.Validate(s)
}
