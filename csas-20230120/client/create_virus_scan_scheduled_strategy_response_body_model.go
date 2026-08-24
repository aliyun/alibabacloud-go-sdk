// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirusScanScheduledStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVirusScanScheduledStrategyResponseBody
	GetRequestId() *string
	SetStrategyId(v string) *CreateVirusScanScheduledStrategyResponseBody
	GetStrategyId() *string
}

type CreateVirusScanScheduledStrategyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the created scheduled virus scan policy.
	//
	// example:
	//
	// vc-strategy-8a3f6c2e91b7****
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s CreateVirusScanScheduledStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVirusScanScheduledStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVirusScanScheduledStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVirusScanScheduledStrategyResponseBody) GetStrategyId() *string {
	return s.StrategyId
}

func (s *CreateVirusScanScheduledStrategyResponseBody) SetRequestId(v string) *CreateVirusScanScheduledStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVirusScanScheduledStrategyResponseBody) SetStrategyId(v string) *CreateVirusScanScheduledStrategyResponseBody {
	s.StrategyId = &v
	return s
}

func (s *CreateVirusScanScheduledStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
