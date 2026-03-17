// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProbeAccessPointNetworkQualityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ProbeAccessPointNetworkQualityResponseBody
	GetRequestId() *string
}

type ProbeAccessPointNetworkQualityResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E26DBAAE-A796-4A48-98B4-B45AFCD1F299
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ProbeAccessPointNetworkQualityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ProbeAccessPointNetworkQualityResponseBody) GoString() string {
	return s.String()
}

func (s *ProbeAccessPointNetworkQualityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ProbeAccessPointNetworkQualityResponseBody) SetRequestId(v string) *ProbeAccessPointNetworkQualityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProbeAccessPointNetworkQualityResponseBody) Validate() error {
	return dara.Validate(s)
}
