// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressCloudConnectionBandwidthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyExpressCloudConnectionBandwidthResponseBody
	GetRequestId() *string
}

type ModifyExpressCloudConnectionBandwidthResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E6385514-B0CC-48E3-B9F9-F7BFF64460A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyExpressCloudConnectionBandwidthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressCloudConnectionBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyExpressCloudConnectionBandwidthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyExpressCloudConnectionBandwidthResponseBody) SetRequestId(v string) *ModifyExpressCloudConnectionBandwidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyExpressCloudConnectionBandwidthResponseBody) Validate() error {
	return dara.Validate(s)
}
