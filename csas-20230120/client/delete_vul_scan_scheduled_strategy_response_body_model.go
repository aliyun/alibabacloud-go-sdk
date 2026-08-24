// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVulScanScheduledStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVulScanScheduledStrategyResponseBody
	GetRequestId() *string
}

type DeleteVulScanScheduledStrategyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVulScanScheduledStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVulScanScheduledStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVulScanScheduledStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVulScanScheduledStrategyResponseBody) SetRequestId(v string) *DeleteVulScanScheduledStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVulScanScheduledStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
