// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertBandwidthPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConvertInstanceId(v string) *ConvertBandwidthPackageResponseBody
	GetConvertInstanceId() *string
	SetRequestId(v string) *ConvertBandwidthPackageResponseBody
	GetRequestId() *string
}

type ConvertBandwidthPackageResponseBody struct {
	// The ID of the Internet Shared Bandwidth instance.
	//
	// example:
	//
	// bwp-s6lmotmkkf567b****
	ConvertInstanceId *string `json:"ConvertInstanceId,omitempty" xml:"ConvertInstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 455AC20C-7061-446A-BDBD-B3BEE0856304
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConvertBandwidthPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConvertBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertBandwidthPackageResponseBody) GetConvertInstanceId() *string {
	return s.ConvertInstanceId
}

func (s *ConvertBandwidthPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConvertBandwidthPackageResponseBody) SetConvertInstanceId(v string) *ConvertBandwidthPackageResponseBody {
	s.ConvertInstanceId = &v
	return s
}

func (s *ConvertBandwidthPackageResponseBody) SetRequestId(v string) *ConvertBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConvertBandwidthPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
