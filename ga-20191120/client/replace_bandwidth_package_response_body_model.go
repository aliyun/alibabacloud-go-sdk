// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceBandwidthPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReplaceBandwidthPackageResponseBody
	GetRequestId() *string
}

type ReplaceBandwidthPackageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A0EA8CCA-F081-4338-9790-A1C791CCA779
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReplaceBandwidthPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplaceBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceBandwidthPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplaceBandwidthPackageResponseBody) SetRequestId(v string) *ReplaceBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceBandwidthPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
